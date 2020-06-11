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
	"encoding/json"
	"fmt"

	apiextv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// ServicePlanGVK is the GroupVersionKind for ServicePlan
var ServicePlanGVK = schema.GroupVersionKind{
	Group:   GroupVersion.Group,
	Version: GroupVersion.Version,
	Kind:    "ServicePlan",
}

// ServicePlanSpec defines the desired state of Service plan
// +k8s:openapi-gen=true
type ServicePlanSpec struct {
	// Kind refers to the service type this is a plan for
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Kind string `json:"kind"`
	// ServiceAccessDisabled is true if service access is disabled for services using this plan
	// It only has an effect if service access is enabled on the service kind
	// +kubebuilder:validation:Optional
	ServiceAccessDisabled bool `json:"serviceAccessDisabled,omitempty"`
	// DisplayName refers to the display name of the service type
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Optional
	DisplayName string `json:"displayName,omitempty"`
	// Labels is a collection of labels for this plan
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`
	// Summary provides a short title summary for the plan
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Summary string `json:"summary"`
	// Description is a detailed description of the service plan
	// +kubebuilder:validation:Optional
	Description string `json:"description,omitempty"`
	// Configuration are the key+value pairs describing a service configuration
	// +kubebuilder:validation:Type=object
	// +kubebuilder:validation:Optional
	Configuration *apiextv1.JSON `json:"configuration,omitempty"`
	// Schema is the JSON schema for the plan
	// +kubebuilder:validation:Optional
	Schema string `json:"schema,omitempty"`
	// CredentialSchema is the JSON schema for credentials created for service using this plan
	// +kubebuilder:validation:Optional
	CredentialSchema string `json:"credentialSchema,omitempty"`
	// ProviderData is provider specific data
	// +kubebuilder:validation:Type=object
	// +kubebuilder:validation:Optional
	ProviderData *apiextv1.JSON `json:"providerData,omitempty"`
}

func (s *ServicePlanSpec) GetConfiguration(v interface{}) error {
	if s.Configuration == nil {
		return nil
	}

	if err := json.Unmarshal(s.Configuration.Raw, v); err != nil {
		return fmt.Errorf("failed to unmarshal service plan configuration: %w", err)
	}
	return nil
}

func (s *ServicePlanSpec) SetConfiguration(v interface{}) error {
	if v == nil {
		s.Configuration = nil
		return nil
	}

	raw, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("failed to marshal service plan configuration: %w", err)
	}
	s.Configuration = &apiextv1.JSON{Raw: raw}
	return nil
}

func (s *ServicePlanSpec) GetProviderData(v interface{}) error {
	if s.ProviderData == nil {
		return nil
	}

	if err := json.Unmarshal(s.ProviderData.Raw, v); err != nil {
		return fmt.Errorf("failed to unmarshal service plan data: %w", err)
	}
	return nil
}

func (s *ServicePlanSpec) SetProviderData(v interface{}) error {
	if v == nil {
		s.ProviderData = nil
		return nil
	}

	raw, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("failed to marshal service kind provider data: %w", err)
	}
	s.ProviderData = &apiextv1.JSON{Raw: raw}
	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServicePlan is a template for a service
// +k8s:openapi-gen=true
// +kubebuilder:resource:path=serviceplans
type ServicePlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ServicePlanSpec `json:"spec,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServicePlanList contains a list of service plans
type ServicePlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServicePlan `json:"items"`
}
