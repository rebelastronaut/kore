// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "github.com/appvia/kore/pkg/apis/core/v1"
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKS) DeepCopyInto(out *EKS) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKS.
func (in *EKS) DeepCopy() *EKS {
	if in == nil {
		return nil
	}
	out := new(EKS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EKS) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSCredentials) DeepCopyInto(out *EKSCredentials) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSCredentials.
func (in *EKSCredentials) DeepCopy() *EKSCredentials {
	if in == nil {
		return nil
	}
	out := new(EKSCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EKSCredentials) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSCredentialsList) DeepCopyInto(out *EKSCredentialsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EKSCredentials, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSCredentialsList.
func (in *EKSCredentialsList) DeepCopy() *EKSCredentialsList {
	if in == nil {
		return nil
	}
	out := new(EKSCredentialsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EKSCredentialsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSCredentialsSpec) DeepCopyInto(out *EKSCredentialsSpec) {
	*out = *in
	if in.CredentialsRef != nil {
		in, out := &in.CredentialsRef, &out.CredentialsRef
		*out = new(v1.SecretReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSCredentialsSpec.
func (in *EKSCredentialsSpec) DeepCopy() *EKSCredentialsSpec {
	if in == nil {
		return nil
	}
	out := new(EKSCredentialsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSCredentialsStatus) DeepCopyInto(out *EKSCredentialsStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]corev1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.Verified != nil {
		in, out := &in.Verified, &out.Verified
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSCredentialsStatus.
func (in *EKSCredentialsStatus) DeepCopy() *EKSCredentialsStatus {
	if in == nil {
		return nil
	}
	out := new(EKSCredentialsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSList) DeepCopyInto(out *EKSList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EKS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSList.
func (in *EKSList) DeepCopy() *EKSList {
	if in == nil {
		return nil
	}
	out := new(EKSList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EKSList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSNodeGroup) DeepCopyInto(out *EKSNodeGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSNodeGroup.
func (in *EKSNodeGroup) DeepCopy() *EKSNodeGroup {
	if in == nil {
		return nil
	}
	out := new(EKSNodeGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EKSNodeGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSNodeGroupList) DeepCopyInto(out *EKSNodeGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EKSNodeGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSNodeGroupList.
func (in *EKSNodeGroupList) DeepCopy() *EKSNodeGroupList {
	if in == nil {
		return nil
	}
	out := new(EKSNodeGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EKSNodeGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSNodeGroupSpec) DeepCopyInto(out *EKSNodeGroupSpec) {
	*out = *in
	out.Cluster = in.Cluster
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SSHSourceSecurityGroups != nil {
		in, out := &in.SSHSourceSecurityGroups, &out.SSHSourceSecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Credentials = in.Credentials
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSNodeGroupSpec.
func (in *EKSNodeGroupSpec) DeepCopy() *EKSNodeGroupSpec {
	if in == nil {
		return nil
	}
	out := new(EKSNodeGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSNodeGroupStatus) DeepCopyInto(out *EKSNodeGroupStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(corev1.Components, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1.Component)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSNodeGroupStatus.
func (in *EKSNodeGroupStatus) DeepCopy() *EKSNodeGroupStatus {
	if in == nil {
		return nil
	}
	out := new(EKSNodeGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSSpec) DeepCopyInto(out *EKSSpec) {
	*out = *in
	if in.AuthorizedMasterNetworks != nil {
		in, out := &in.AuthorizedMasterNetworks, &out.AuthorizedMasterNetworks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Cluster = in.Cluster
	if in.SubnetIDs != nil {
		in, out := &in.SubnetIDs, &out.SubnetIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroupIDs != nil {
		in, out := &in.SecurityGroupIDs, &out.SecurityGroupIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Credentials = in.Credentials
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSSpec.
func (in *EKSSpec) DeepCopy() *EKSSpec {
	if in == nil {
		return nil
	}
	out := new(EKSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSStatus) DeepCopyInto(out *EKSStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(corev1.Components, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1.Component)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSStatus.
func (in *EKSStatus) DeepCopy() *EKSStatus {
	if in == nil {
		return nil
	}
	out := new(EKSStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSVPC) DeepCopyInto(out *EKSVPC) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSVPC.
func (in *EKSVPC) DeepCopy() *EKSVPC {
	if in == nil {
		return nil
	}
	out := new(EKSVPC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EKSVPC) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSVPCList) DeepCopyInto(out *EKSVPCList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EKSVPC, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSVPCList.
func (in *EKSVPCList) DeepCopy() *EKSVPCList {
	if in == nil {
		return nil
	}
	out := new(EKSVPCList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EKSVPCList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSVPCSpec) DeepCopyInto(out *EKSVPCSpec) {
	*out = *in
	out.Credentials = in.Credentials
	out.Cluster = in.Cluster
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSVPCSpec.
func (in *EKSVPCSpec) DeepCopy() *EKSVPCSpec {
	if in == nil {
		return nil
	}
	out := new(EKSVPCSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EKSVPCStatus) DeepCopyInto(out *EKSVPCStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(corev1.Components, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1.Component)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	in.Infra.DeepCopyInto(&out.Infra)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EKSVPCStatus.
func (in *EKSVPCStatus) DeepCopy() *EKSVPCStatus {
	if in == nil {
		return nil
	}
	out := new(EKSVPCStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Infra) DeepCopyInto(out *Infra) {
	*out = *in
	if in.AvailabilityZoneIDs != nil {
		in, out := &in.AvailabilityZoneIDs, &out.AvailabilityZoneIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AvailabilityZoneNames != nil {
		in, out := &in.AvailabilityZoneNames, &out.AvailabilityZoneNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PrivateSubnetIDs != nil {
		in, out := &in.PrivateSubnetIDs, &out.PrivateSubnetIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PublicSubnetIDs != nil {
		in, out := &in.PublicSubnetIDs, &out.PublicSubnetIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroupIDs != nil {
		in, out := &in.SecurityGroupIDs, &out.SecurityGroupIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PublicIPV4EgressAddresses != nil {
		in, out := &in.PublicIPV4EgressAddresses, &out.PublicIPV4EgressAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Infra.
func (in *Infra) DeepCopy() *Infra {
	if in == nil {
		return nil
	}
	out := new(Infra)
	in.DeepCopyInto(out)
	return out
}
