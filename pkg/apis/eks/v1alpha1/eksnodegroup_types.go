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

// EKSNodeGroupSpec defines the desired state of EKSNodeGroup
// +k8s:openapi-gen=true
type EKSNodeGroupSpec struct {
	// AMIType is the AWS Machine Image type. We use a sensible default.
	AMIType string `json:"amiType"`
	// Cluster refers to the cluster this object belongs to
	// +kubebuilder:validation:Required
	Cluster corev1.Ownership `json:"cluster,omitempty"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Required
	DiskSize int64 `json:"diskSize"`
	// InstanceType is the EC2 machine type
	// +kubebuilder:validation:Required
	InstanceType string `json:"instanceType,omitempty"`
	// Labels are any custom kubernetes labels to apply to nodes
	Labels map[string]string `json:"labels,omitempty"`
	// Version is the Kubernetes version to run for the kubelet
	Version string `json:"version,omitempty"`
	// ReleaseVersion is release version of the managed node ami
	ReleaseVersion string `json:"releaseVersion,omitempty"`
	// DesiredSize is the number of nodes to attempt to use
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Required
	DesiredSize int64 `json:"desiredSize"`
	// MaxSize is the most nodes the nodegroups can grow to
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Maximum=100
	MaxSize int64 `json:"maxSize"`
	// MinSize is the least nodes the nodegroups can shrink to
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Minimum=1
	MinSize int64 `json:"minSize"`
	// Subnets is the VPC networks to use for the nodes
	// +kubebuilder:validation:Required
	// +listType=set
	Subnets []string `json:"subnets"`
	// Tags are the AWS metadata to apply to the node group
	// +kubebuilder:validation:Required
	Tags map[string]string `json:"tags,omitempty"`
	// Region is the AWS location to launch node group within, must match the region of the cluster
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// SSHSourceSecurityGroups is the security groups that are allowed SSH access (port 22) to the worker nodes
	// +listType=set
	SSHSourceSecurityGroups []string `json:"sshSourceSecurityGroups,omitempty"`
	// EC2SSHKey is the Amazon EC2 SSH key that provides access for SSH communication with
	// the worker nodes in the managed node group
	// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html
	// +kubebuilder:validation:Required
	EC2SSHKey string `json:"eC2SSHKey"`
	// EnableAutoscaler indicates if the node pool should be configured with
	// autoscaling turned on
	// +kubebuilder:validation:Optional
	EnableAutoscaler bool `json:"enableAutoscaler"`
	// Credentials is a reference to an AWSCredentials object to use for authentication
	// +kubebuilder:validation:Required
	// +k8s:openapi-gen=false
	Credentials core.Ownership `json:"credentials"`
}

// EKSNodeGroupStatus defines the observed state of EKSNodeGroup
// +k8s:openapi-gen=true
type EKSNodeGroupStatus struct {
	// Conditions is the status of the components
	Conditions core.Components `json:"conditions,omitempty"`
	// NodeIAMRole is the IAM role assumed by the worker nodes themselves
	NodeIAMRole string `json:"nodeIAMRole,omitempty"`
	// AutoScalingGroupName is the name of the Auto Scaling Groups belonging to this node group
	// +listType=set
	AutoScalingGroupNames []string `json:"autoScalingGroupNames,omitempty"`
	// Status provides a overall status
	Status core.Status `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EKSNodeGroup is the Schema for the eksnodegroups API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=eksnodegroups,scope=Namespaced
// +kubebuilder:printcolumn:name="Description",type="string",JSONPath=".spec.description",description="A description of the EKS cluster nodegroup"
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.status",description="The overall status of the cluster nodegroup"
type EKSNodeGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EKSNodeGroupSpec   `json:"spec,omitempty"`
	Status EKSNodeGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EKSNodeGroupList contains a list of EKSNodeGroup
type EKSNodeGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EKSNodeGroup `json:"items"`
}
