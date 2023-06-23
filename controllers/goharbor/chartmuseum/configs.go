package chartmuseum

import (
	"context"
	"crypto/sha256"
	"fmt"

	goharborv1 "github.com/plotly/harbor-operator/apis/goharbor.io/v1beta1"
	conftemplate "github.com/plotly/harbor-operator/pkg/config/template"
	"github.com/plotly/harbor-operator/pkg/resources/checksum"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const ConfigName = "config.yaml"

func (r *Reconciler) GetConfigMap(ctx context.Context, chartMuseum *goharborv1.ChartMuseum) (*corev1.ConfigMap, error) {
	templateConfig, err := r.ConfigStore.GetItemValue(conftemplate.ConfigTemplateKey)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get template")
	}

	content, err := r.GetTemplatedConfig(ctx, templateConfig, chartMuseum)
	if err != nil {
		return nil, err
	}

	name := r.NormalizeName(ctx, chartMuseum.GetName())
	namespace := chartMuseum.GetNamespace()

	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Annotations: map[string]string{
				checksum.GetStaticID("template"): fmt.Sprintf("%x", sha256.Sum256([]byte(templateConfig))),
			},
		},
		BinaryData: map[string][]byte{
			ConfigName: content,
		},
	}, nil
}
