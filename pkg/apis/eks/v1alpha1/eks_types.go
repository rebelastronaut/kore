/**
 * Copyright 2020 Appvia Ltd <info@appvia.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package v1alpha1

import (
	core "github.com/appvia/kore/pkg/apis/core/v1"
	corev1 "github.com/appvia/kore/pkg/apis/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EKSSpec defines the desired state of EKSCluster
// +k8s:openapi-gen=true
type EKSSpec struct {
	// AuthorizedMasterNetworks is the network ranges which are permitted
	// to access the EKS control plane endpoint i.e the managed one (not the
	// authentication proxy)
	AuthorizedMasterNetworks []string `json:"authorizedMasterNetworks,omitempty"`
	// Cluster refers to the cluster this object belongs to
	// +kubebuilder:validation:Required
	Cluster corev1.Ownership `json:"cluster,omitempty"`
	// Version is the Kubernetes version to use
	// +kubebuilder:validation:MinLength=3
	// +kubebuilder:validation:Required
	Version string `json:"version,omitempty"`
	// Region is the AWS region to launch this cluster within
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// EnablePrivateNetwork indicates if the EKS should have public networking
	// enabled or not - note at present peering is not provided but must be
	// implemented outside of Kore
	// +kubebuilder:validation:Optional
	EnablePrivateNetwork *bool `json:"enablePrivateNetwork,omitempty"`
	// SubnetIds is a list of subnet IDs
	// +kubebuilder:validation:Required
	SubnetIDs []string `json:"subnetIDs"`
	// SecurityGroupIds is a list of security group IDs
	// +kubebuilder:validation:Required
	SecurityGroupIDs []string `json:"securityGroupIDs,omitempty"`
	// Credentials is a reference to an EKSCredentials object to use for authentication
	// +k8s:openapi-gen=false
	Credentials core.Ownership `json:"credentials"`
	// Tags is a collection of tags to apply to the AWS resources which make up this cluster
	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty"`
}

// EKSStatus defines the observed state of EKS cluster
// +k8s:openapi-gen=true
type EKSStatus struct {
	// Conditions is the status of the components
	Conditions core.Components `json:"conditions,omitempty"`
	// CACertificate is the certificate for this cluster
	CACertificate string `json:"caCertificate,omitempty"`
	// Endpoint is the endpoint of the cluster
	Endpoint string `json:"endpoint,omitempty"`
	// ARN is the AWS ARN of the EKS cluster resource
	ARN string `json:"arn,omitempty"`
	// RoleARN is the role ARN which provides permissions to EKS
	RoleARN string `json:"roleARN,omitempty"`
	// OIDCProviderURL is the OIDC provider URL (used for providing IAM roles for service accounts)
	OIDCProviderURL string `json:"oidcProviderURL,omitempty"`
	// Status provides a overall status
	Status core.Status `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EKS is the Schema for the eksclusters API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=eks,scope=Namespaced
// +kubebuilder:printcolumn:name="Description",type="string",JSONPath=".spec.description",description="A description of the EKS cluster"
// +kubebuilder:printcolumn:name="Endpoint",type="string",JSONPath=".status.endpoint",description="The endpoint of the eks cluster"
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.status",description="The overall status of the cluster"
type EKS struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EKSSpec   `json:"spec,omitempty"`
	Status EKSStatus `json:"status,omitempty"`
}

// Ownership returns a owner reference
func (e *EKS) Ownership() corev1.Ownership {
	return corev1.Ownership{
		Group:     GroupVersion.Group,
		Version:   GroupVersion.Version,
		Kind:      "EKS",
		Namespace: e.Namespace,
		Name:      e.Name,
	}
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EKSList contains a list of EKS clusters
type EKSList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EKS `json:"items"`
}
