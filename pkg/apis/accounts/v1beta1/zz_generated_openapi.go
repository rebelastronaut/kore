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

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1beta1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/appvia/kore/pkg/apis/accounts/v1beta1.AccountManagement":       schema_pkg_apis_accounts_v1beta1_AccountManagement(ref),
		"github.com/appvia/kore/pkg/apis/accounts/v1beta1.AccountManagementSpec":   schema_pkg_apis_accounts_v1beta1_AccountManagementSpec(ref),
		"github.com/appvia/kore/pkg/apis/accounts/v1beta1.AccountManagementStatus": schema_pkg_apis_accounts_v1beta1_AccountManagementStatus(ref),
	}
}

func schema_pkg_apis_accounts_v1beta1_AccountManagement(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AccountManagement is the Schema for the accounts API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/appvia/kore/pkg/apis/accounts/v1beta1.AccountManagementSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/appvia/kore/pkg/apis/accounts/v1beta1.AccountManagementStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/accounts/v1beta1.AccountManagementSpec", "github.com/appvia/kore/pkg/apis/accounts/v1beta1.AccountManagementStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_accounts_v1beta1_AccountManagementSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AccountManagementSpec defines the desired state of accounting for a provider I've a feeling this will probably need provider specific attributes are some point",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"provider": {
						SchemaProps: spec.SchemaProps{
							Description: "Provider is the name of provider which maps to the cluster kind",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"managed": {
						SchemaProps: spec.SchemaProps{
							Description: "Managed indicates if kore to manage the accounts - is this one actually required???",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"rules": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Rules is a set of rules for this provider",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/appvia/kore/pkg/apis/accounts/v1beta1.AccountsRule"),
									},
								},
							},
						},
					},
					"organization": {
						SchemaProps: spec.SchemaProps{
							Description: "Organization is the underlying organizational resource (only require if more than one)",
							Ref:         ref("github.com/appvia/kore/pkg/apis/core/v1.Ownership"),
						},
					},
				},
				Required: []string{"provider", "managed"},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/accounts/v1beta1.AccountsRule", "github.com/appvia/kore/pkg/apis/core/v1.Ownership"},
	}
}

func schema_pkg_apis_accounts_v1beta1_AccountManagementStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AccountManagementStatus defines the observed state of Allocation",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status is the general status of the resource",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Conditions is a collection of potential issues",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/appvia/kore/pkg/apis/core/v1.Condition"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/core/v1.Condition"},
	}
}
