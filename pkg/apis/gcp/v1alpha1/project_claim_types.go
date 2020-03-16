package v1alpha1

import (
	core "github.com/appvia/kore/pkg/apis/core/v1"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ProjectClaimSpec defines the desired state of ProjectClaim
// +k8s:openapi-gen=true
type ProjectClaimSpec struct {
	// Organization is a reference to the gcp admin project to use
	// +kubebuilder:validation:Required
	Organization core.Ownership `json:"organization"`
}

// ProjectClaimStatus defines the observed state of GCP Project
// +k8s:openapi-gen=true
type ProjectClaimStatus struct {
	// CredentialRef is the reference to the credentials secret
	CredentialRef *v1.SecretReference `json:"credentialRef,omitempty"`
	// ProjectID is the  project id
	ProjectID string `json:"projectID,omitempty"`
	// Status provides a overall status
	Status core.Status `json:"status,omitempty"`
	// Conditions is a set of components conditions
	Conditions *core.Components `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ProjectClaim is the Schema for the gcpprojects API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=projectclaims,scope=Namespaced
type ProjectClaim struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProjectClaimSpec   `json:"spec,omitempty"`
	Status ProjectClaimStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ProjectClaimList contains a list of ProjectClaim
type ProjectClaimList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectClaim `json:"items"`
}
