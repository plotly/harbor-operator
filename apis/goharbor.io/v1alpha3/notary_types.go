package v1alpha3

import (
	harbormetav1 "github.com/plotly/harbor-operator/apis/meta/v1alpha1"
)

type NotaryLoggingSpec struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default="info"
	Level harbormetav1.NotaryLogLevel `json:"level,omitempty"`
}

type NotaryStorageSpec struct {
	// +kubebuilder:validation:Required
	Postgres harbormetav1.PostgresConnectionWithParameters `json:"postgres"`

	// TODO Add support for mysql and memory
}

func (n *NotaryStorageSpec) GetPasswordFieldKey() string {
	return harbormetav1.PostgresqlPasswordKey
}
