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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AuditEventSpec defines the desired state of User
// +k8s:openapi-gen=false
type AuditEventSpec struct {
	// CreatedAt is the timestamp of record creation
	CreatedAt metav1.Time `json:"createdAt,omitempty"`
	// Type is the type of event
	Type string `json:"type,omitempty"`
	// Team is the team whom event may be associated to
	Team string `json:"team,omitempty"`
	// User is the user which the event is related
	User string `json:"user,omitempty"`
	// Message is event message itself
	Message string `json:"message,omitempty"`
	// Resource is the name of the resource in question namespace/name
	Resource string `json:"resource,omitempty"`
	// ResourceUID is a unique id for the resource
	ResourceUID string `json:"resourceUID,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AuditEvent is the Schema for the audit API
// +k8s:openapi-gen=false
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=auditevents
type AuditEvent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec AuditEventSpec `json:"spec,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AuditEventList contains a list of audit event
type AuditEventList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AuditEvent `json:"items"`
}
