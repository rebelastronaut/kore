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

package v1

import (
	corev1 "github.com/appvia/kore/pkg/apis/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// ServiceCatalogGVK is the GroupVersionKind for ServiceCatalog
var ServiceCatalogGVK = schema.GroupVersionKind{
	Group:   GroupVersion.Group,
	Version: GroupVersion.Version,
	Kind:    "ServiceCatalog",
}

// ServiceCatalogSpec defines the desired state of a service satalog
// +k8s:openapi-gen=true
type ServiceCatalogSpec struct {
	// DisplayName overrides the name to display
	DisplayName string `json:"displayName,omitempty"`
	// Summary provides a short title summary for the catalog
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Summary string `json:"summary"`
	// Description is a detailed description of the service catalog
	// +kubebuilder:validation:Optional
	Description string `json:"description,omitempty"`
	// URL is the URL of the service catalog
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	URL string `json:"url"`
	// ServiceKindPrefix is the prefix to add to all created service kinds
	ServiceKindPrefix string `json:"serviceKindPrefix"`
}

// ServiceCatalogStatus defines the observed state of a service catalog
// +k8s:openapi-gen=true
type ServiceCatalogStatus struct {
	// Status is the overall status of the service
	// +kubebuilder:validation:Optional
	Status corev1.Status `json:"status,omitempty"`
	// Message is the description of the current status
	// +kubebuilder:validation:Optional
	Message string `json:"message,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceCatalog is a template for a service catalog
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=servicecatalogs
type ServiceCatalog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceCatalogSpec   `json:"spec,omitempty"`
	Status ServiceCatalogStatus `json:"status,omitempty"`
}

func NewServiceCatalog(name, namespace string) *ServiceCatalog {
	return &ServiceCatalog{
		TypeMeta: metav1.TypeMeta{
			Kind:       ServiceCatalogGVK.Kind,
			APIVersion: GroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
	}
}

func (s *ServiceCatalog) GetStatus() (status corev1.Status, message string) {
	return s.Status.Status, s.Status.Message
}

func (s *ServiceCatalog) SetStatus(status corev1.Status, message string) {
	s.Status.Status = status
	s.Status.Message = message
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceCatalogList contains a list of service catalogs
type ServiceCatalogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceCatalog `json:"items"`
}
