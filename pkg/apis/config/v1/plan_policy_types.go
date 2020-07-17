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
)

// PlanPolicySpec defines Plan JSON Schema extensions
// +k8s:openapi-gen=true
type PlanPolicySpec struct {
	// Kind refers to the cluster type this is a plan policy for
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Kind string `json:"kind"`
	// Labels is a collection of labels for this plan policy
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`
	// Summary provides a short title summary for the plan policy
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Summary string `json:"summary"`
	// Description provides a detailed description of the plan policy
	// +kubebuilder:validation:Optional
	Description string `json:"description"`
	// Properties are the
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinItems=1
	// +listType=map
	// +listMapKey=name
	Properties []PlanPolicyProperty `json:"properties"`
}

// PlanPolicyStatus defines the observed state of Plan Policy
// +k8s:openapi-gen=true
type PlanPolicyStatus struct {
	// Conditions is a set of condition which has caused an error
	// +kubebuilder:validation:Optional
	Conditions []corev1.Condition `json:"conditions"`
	// Status is overall status of the plan policy
	Status corev1.Status `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PlanPolicy is the Schema for the plan policies API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=planpolicies
type PlanPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PlanPolicySpec   `json:"spec,omitempty"`
	Status PlanPolicyStatus `json:"status,omitempty"`
}

func (p PlanPolicy) CreateAllocation(teams []string) *Allocation {
	annotations := map[string]string{}
	if p.Annotations["kore.appvia.io/readonly"] == "true" {
		annotations["kore.appvia.io/readonly"] = "true"
	}
	return &Allocation{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "planpolicy-" + p.Name,
			Annotations: annotations,
		},
		Spec: AllocationSpec{
			Name:    p.Name,
			Summary: p.Spec.Description,
			Resource: corev1.Ownership{
				Group:     GroupVersion.Group,
				Version:   GroupVersion.Version,
				Kind:      "PlanPolicy",
				Namespace: p.Namespace,
				Name:      p.Name,
			},
			Teams: teams,
		},
	}
}

// PlanPolicyProperty defines a JSON schema for a given property
// +k8s:openapi-gen=true
type PlanPolicyProperty struct {
	// Name is the name of the property
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Name string `json:"name"`
	// AllowUpdate will allow the parameter to be modified by the teams
	AllowUpdate bool `json:"allowUpdate"`
	// DisallowUpdate will forbid modification of the parameter, even if it was allowed by an other policy
	DisallowUpdate bool `json:"disallowUpdate"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PlanPolicyList contains a list of Plan Policies
type PlanPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PlanPolicy `json:"items"`
}
