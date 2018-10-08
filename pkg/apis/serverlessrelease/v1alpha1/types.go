package v1alpha1

import (
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ServerlessRelease struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata, omitempty"`

	Spec   ServerlessReleaseSpec   `json:"spec"`
	Status ServerlessReleaseStatus `json:"status"`
}

type ServerlessReleaseSpec struct {
	PackagePath string   `json:"packagePath"`
	ReleaseName string   `json:"releaseName,omitempty"`
	Stage       string   `json:"stage"`
	Region      string   `json:"region"`
	Verbose     bool     `json:"verbose"`
	Env         []EnvVar `json:"env"`
}

type ServerlessReleaseStatus struct {
	ReleaseStatus string `json:"releaseStatus"`
}

// EnvVar represents an environment variable present when deploying.
type EnvVar struct {
	// Name of the environment variable.
	Name string `json:"name"`

	// Value for the variable. Defaults to ""
	Value string `json:"value,omitempty"`

	// Source for the environment variable's value. Cannot be used if value is not empty.
	ValueFrom *EnvVarSource `json:"valueFrom,omitempty"`
}

// EnvVarSource represents a source for the value of an EnvVar.
type EnvVarSource struct {
	// Selects a key of a ConfigMap.
	ConfigMapKeyRef *apiv1.ConfigMapKeySelector `json:"configMapKeyRef,omitempty"`

	// Selects a key of a secret in the pod's namespace
	SecretKeyRef *apiv1.SecretKeySelector `json:"secretKeyRef,omitempty"`
}
