package core

import (
	"context"
	"fmt"
	"net/url"

	goharborv1 "github.com/plotly/harbor-operator/apis/goharbor.io/v1beta1"
	harbormetav1 "github.com/plotly/harbor-operator/apis/meta/v1alpha1"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

const (
	RedisDSNKey         = "_REDIS_URL_CORE"
	RegistryRedisDSNKey = "_REDIS_URL_REG"
)

func (r *Reconciler) GetSecret(ctx context.Context, core *goharborv1.Core) (*corev1.Secret, error) {
	name := r.NormalizeName(ctx, core.GetName())
	namespace := core.GetNamespace()

	var redisPassword string

	if core.Spec.Redis.PasswordRef != "" {
		var passwordSecret corev1.Secret

		err := r.Client.Get(ctx, types.NamespacedName{
			Namespace: namespace,
			Name:      core.Spec.Redis.PasswordRef,
		}, &passwordSecret)
		if err != nil {
			return nil, errors.Wrap(err, "cannot get redis password")
		}

		password, ok := passwordSecret.Data[harbormetav1.RedisPasswordKey]
		if !ok {
			return nil, errors.Errorf("%s not found in secret %s", harbormetav1.RedisPasswordKey, core.Spec.Redis.PasswordRef)
		}

		redisPassword = string(password)
	}
	// support redis password contains special character by using url.QueryEscape
	coreCacheDSN := core.Spec.Redis.GetDSNStringWithRawPassword(url.QueryEscape(redisPassword))

	var registryPassword string

	if core.Spec.Components.Registry.Redis != nil && core.Spec.Components.Registry.Redis.PasswordRef != "" {
		var passwordSecret corev1.Secret

		err := r.Client.Get(ctx, types.NamespacedName{
			Namespace: namespace,
			Name:      core.Spec.Components.Registry.Redis.PasswordRef,
		}, &passwordSecret)
		if err != nil {
			return nil, errors.Wrap(err, "cannot get registry redis password")
		}

		password, ok := passwordSecret.Data[harbormetav1.RedisPasswordKey]
		if !ok {
			return nil, errors.Errorf("%s not found in secret %s", harbormetav1.RedisPasswordKey, core.Spec.Components.Registry.Redis.PasswordRef)
		}

		registryPassword = string(password)
	}
	// support redis password contains special character by using url.QueryEscape
	registryCacheDSN := core.Spec.Components.Registry.Redis.GetDSNStringWithRawPassword(url.QueryEscape(registryPassword))

	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		StringData: map[string]string{
			RedisDSNKey:         fmt.Sprintf("%s?idle_timeout_seconds=%.0f", coreCacheDSN, core.Spec.Redis.IdleTimeout.Duration.Seconds()),
			RegistryRedisDSNKey: fmt.Sprintf("%s?idle_timeout_seconds=%.0f", registryCacheDSN, core.Spec.Redis.IdleTimeout.Duration.Seconds()),
		},
	}, nil
}
