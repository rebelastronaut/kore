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

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Component) DeepCopyInto(out *Component) {
	*out = *in
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(Ownership)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Component.
func (in *Component) DeepCopy() *Component {
	if in == nil {
		return nil
	}
	out := new(Component)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Components) DeepCopyInto(out *Components) {
	{
		in := &in
		*out = make(Components, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Component)
				(*in).DeepCopyInto(*out)
			}
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Components.
func (in Components) DeepCopy() Components {
	if in == nil {
		return nil
	}
	out := new(Components)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationFromSource) DeepCopyInto(out *ConfigurationFromSource) {
	*out = *in
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		*out = new(OptionalSecretKeySelector)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationFromSource.
func (in *ConfigurationFromSource) DeepCopy() *ConfigurationFromSource {
	if in == nil {
		return nil
	}
	out := new(ConfigurationFromSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GithubIDP) DeepCopyInto(out *GithubIDP) {
	*out = *in
	if in.Orgs != nil {
		in, out := &in.Orgs, &out.Orgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GithubIDP.
func (in *GithubIDP) DeepCopy() *GithubIDP {
	if in == nil {
		return nil
	}
	out := new(GithubIDP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GoogleIDP) DeepCopyInto(out *GoogleIDP) {
	*out = *in
	if in.Domains != nil {
		in, out := &in.Domains, &out.Domains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GoogleIDP.
func (in *GoogleIDP) DeepCopy() *GoogleIDP {
	if in == nil {
		return nil
	}
	out := new(GoogleIDP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IDP) DeepCopyInto(out *IDP) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IDP.
func (in *IDP) DeepCopy() *IDP {
	if in == nil {
		return nil
	}
	out := new(IDP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IDP) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IDPClient) DeepCopyInto(out *IDPClient) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IDPClient.
func (in *IDPClient) DeepCopy() *IDPClient {
	if in == nil {
		return nil
	}
	out := new(IDPClient)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IDPClient) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IDPClientList) DeepCopyInto(out *IDPClientList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IDPClient, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IDPClientList.
func (in *IDPClientList) DeepCopy() *IDPClientList {
	if in == nil {
		return nil
	}
	out := new(IDPClientList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IDPClientList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IDPClientSpec) DeepCopyInto(out *IDPClientSpec) {
	*out = *in
	if in.RedirectURIs != nil {
		in, out := &in.RedirectURIs, &out.RedirectURIs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IDPClientSpec.
func (in *IDPClientSpec) DeepCopy() *IDPClientSpec {
	if in == nil {
		return nil
	}
	out := new(IDPClientSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IDPClientStatus) DeepCopyInto(out *IDPClientStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IDPClientStatus.
func (in *IDPClientStatus) DeepCopy() *IDPClientStatus {
	if in == nil {
		return nil
	}
	out := new(IDPClientStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IDPConfig) DeepCopyInto(out *IDPConfig) {
	*out = *in
	if in.Github != nil {
		in, out := &in.Github, &out.Github
		*out = new(GithubIDP)
		(*in).DeepCopyInto(*out)
	}
	if in.Google != nil {
		in, out := &in.Google, &out.Google
		*out = new(GoogleIDP)
		(*in).DeepCopyInto(*out)
	}
	if in.SAML != nil {
		in, out := &in.SAML, &out.SAML
		*out = new(SAMLIDP)
		(*in).DeepCopyInto(*out)
	}
	if in.OIDC != nil {
		in, out := &in.OIDC, &out.OIDC
		*out = new(OIDCIDP)
		(*in).DeepCopyInto(*out)
	}
	if in.OIDCDirect != nil {
		in, out := &in.OIDCDirect, &out.OIDCDirect
		*out = new(StaticOIDCIDP)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IDPConfig.
func (in *IDPConfig) DeepCopy() *IDPConfig {
	if in == nil {
		return nil
	}
	out := new(IDPConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IDPList) DeepCopyInto(out *IDPList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IDP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IDPList.
func (in *IDPList) DeepCopy() *IDPList {
	if in == nil {
		return nil
	}
	out := new(IDPList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IDPList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IDPSpec) DeepCopyInto(out *IDPSpec) {
	*out = *in
	in.Config.DeepCopyInto(&out.Config)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IDPSpec.
func (in *IDPSpec) DeepCopy() *IDPSpec {
	if in == nil {
		return nil
	}
	out := new(IDPSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IDPStatus) DeepCopyInto(out *IDPStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IDPStatus.
func (in *IDPStatus) DeepCopy() *IDPStatus {
	if in == nil {
		return nil
	}
	out := new(IDPStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OIDCIDP) DeepCopyInto(out *OIDCIDP) {
	*out = *in
	if in.ClientScopes != nil {
		in, out := &in.ClientScopes, &out.ClientScopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.UserClaims != nil {
		in, out := &in.UserClaims, &out.UserClaims
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OIDCIDP.
func (in *OIDCIDP) DeepCopy() *OIDCIDP {
	if in == nil {
		return nil
	}
	out := new(OIDCIDP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OptionalSecretKeySelector) DeepCopyInto(out *OptionalSecretKeySelector) {
	*out = *in
	out.SecretKeySelector = in.SecretKeySelector
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OptionalSecretKeySelector.
func (in *OptionalSecretKeySelector) DeepCopy() *OptionalSecretKeySelector {
	if in == nil {
		return nil
	}
	out := new(OptionalSecretKeySelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ownership) DeepCopyInto(out *Ownership) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ownership.
func (in *Ownership) DeepCopy() *Ownership {
	if in == nil {
		return nil
	}
	out := new(Ownership)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSelector) DeepCopyInto(out *ResourceSelector) {
	*out = *in
	out.Resource = in.Resource
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSelector.
func (in *ResourceSelector) DeepCopy() *ResourceSelector {
	if in == nil {
		return nil
	}
	out := new(ResourceSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SAMLIDP) DeepCopyInto(out *SAMLIDP) {
	*out = *in
	if in.CAData != nil {
		in, out := &in.CAData, &out.CAData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.AllowedGroups != nil {
		in, out := &in.AllowedGroups, &out.AllowedGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SAMLIDP.
func (in *SAMLIDP) DeepCopy() *SAMLIDP {
	if in == nil {
		return nil
	}
	out := new(SAMLIDP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretKeySelector) DeepCopyInto(out *SecretKeySelector) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretKeySelector.
func (in *SecretKeySelector) DeepCopy() *SecretKeySelector {
	if in == nil {
		return nil
	}
	out := new(SecretKeySelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticOIDCIDP) DeepCopyInto(out *StaticOIDCIDP) {
	*out = *in
	in.OIDCIDP.DeepCopyInto(&out.OIDCIDP)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticOIDCIDP.
func (in *StaticOIDCIDP) DeepCopy() *StaticOIDCIDP {
	if in == nil {
		return nil
	}
	out := new(StaticOIDCIDP)
	in.DeepCopyInto(out)
	return out
}
