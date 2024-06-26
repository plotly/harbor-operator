package v1alpha3

import (
	"github.com/plotly/harbor-operator/pkg/convert"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

var _ conversion.Convertible = &Exporter{}

func (e *Exporter) ConvertTo(dstRaw conversion.Hub) error {
	return convert.ConverterObject(e).To(dstRaw)
}

func (e *Exporter) ConvertFrom(srcRaw conversion.Hub) error {
	return convert.ConverterObject(e).From(srcRaw)
}
