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
	apiextv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// ServiceDeploymentGVK is the GroupVersionKind for ServiceDeployment
var ServiceDeploymentGVK = schema.GroupVersionKind{
	Group:   GroupVersion.Group,
	Version: GroupVersion.Version,
	Kind:    "ServiceDeployment",
}

// ServiceDeploymentSpec defines the desired state of a service satalog
// +k8s:openapi-gen=true
type ServiceDeploymentSpec struct {
	// DisplayName overrides the name to display
	DisplayName string `json:"displayName,omitempty"`
	// Summary provides a short title summary for the deployment
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Summary string `json:"summary"`
	// Description is a detailed description of the service deployment
	// +kubebuilder:validation:Optional
	Description string `json:"description,omitempty"`
	// Kind refers to the service type
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Kind string `json:"kind"`
	// Plan is the name of the service plan which is used to create the services
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	Plan string `json:"plan"`
	// Configuration are the configuration values for the created service
	// It will contain values from the plan + overrides by the user
	// This will provide a simple interface to calculate diffs between plan and service configuration
	// +kubebuilder:validation:Type=object
	// +kubebuilder:validation:Optional
	Configuration *apiextv1.JSON `json:"configuration,omitempty"`
	// ConfigurationFrom is a way to load configuration values from alternative sources, e.g. from secrets
	// The values from these sources will override any existing keys defined in Configuration
	// +kubebuilder:validation:Optional
	ConfigurationFrom corev1.ConfigurationFromSourceList `json:"configurationFrom,omitempty"`
	// ClusterSelector defines in which clusters should we install the given service
	// +kubebuilder:validation:Required
	ClusterSelector ClusterSelector `json:"clusterSelector"`
	// ClusterNamespace is the target namespace in the clusters where there the service will be created
	// +kubebuilder:validation:Optional
	ClusterNamespace string `json:"clusterNamespace,omitempty"`
	// ServiceName is the name of the service in each cluster
	// If empty it defaults to the name of the service deployment
	// +kubebuilder:validation:Optional
	ServiceName string `json:"serviceName,omitempty"`
}

// ClusterSelector is a way to define conditions to identify a group of clusters
// +k8s:openapi-gen=true
type ClusterSelector struct {
	// Kinds defines the cluster kinds this deployment applies to
	// If empty, the cluster kind is not filtered
	// +listType=set
	Kinds []string `json:"kinds,omitempty"`
	// Kinds defines the teams this deployment applies to
	// If empty, the team is not filtered
	// +listType=set
	Teams []string `json:"teams,omitempty"`
	// LabelSelector is a cluster label selector
	metav1.LabelSelector `json:",inline"`
}

// ServiceDeploymentStatus defines the observed state of a service deployment
// +k8s:openapi-gen=true
type ServiceDeploymentStatus struct {
	// Status is the overall status of the service
	// +kubebuilder:validation:Optional
	Status corev1.Status `json:"status,omitempty"`
	// Message is the description of the current status
	// +kubebuilder:validation:Optional
	Message string `json:"message,omitempty"`
	// Components is a collection of component statuses
	// +kubebuilder:validation:Optional
	Components corev1.Components `json:"components,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceDeployment is a template for a service deployment
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=servicedeployments
type ServiceDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceDeploymentSpec   `json:"spec,omitempty"`
	Status ServiceDeploymentStatus `json:"status,omitempty"`
}

func NewServiceDeployment(name, namespace string) *ServiceDeployment {
	return &ServiceDeployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       ServiceDeploymentGVK.Kind,
			APIVersion: GroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
	}
}

func (s *ServiceDeployment) GetStatus() (status corev1.Status, message string) {
	return s.Status.Status, s.Status.Message
}

func (s *ServiceDeployment) SetStatus(status corev1.Status, message string) {
	s.Status.Status = status
	s.Status.Message = message
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceDeploymentList contains a list of service deployments
type ServiceDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceDeployment `json:"items"`
}
