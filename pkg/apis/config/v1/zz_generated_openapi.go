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

package v1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/appvia/kore/pkg/apis/config/v1.Allocation":         schema_pkg_apis_config_v1_Allocation(ref),
		"github.com/appvia/kore/pkg/apis/config/v1.AllocationSpec":     schema_pkg_apis_config_v1_AllocationSpec(ref),
		"github.com/appvia/kore/pkg/apis/config/v1.AllocationStatus":   schema_pkg_apis_config_v1_AllocationStatus(ref),
		"github.com/appvia/kore/pkg/apis/config/v1.Plan":               schema_pkg_apis_config_v1_Plan(ref),
		"github.com/appvia/kore/pkg/apis/config/v1.PlanPolicy":         schema_pkg_apis_config_v1_PlanPolicy(ref),
		"github.com/appvia/kore/pkg/apis/config/v1.PlanPolicyProperty": schema_pkg_apis_config_v1_PlanPolicyProperty(ref),
		"github.com/appvia/kore/pkg/apis/config/v1.PlanPolicySpec":     schema_pkg_apis_config_v1_PlanPolicySpec(ref),
		"github.com/appvia/kore/pkg/apis/config/v1.PlanPolicyStatus":   schema_pkg_apis_config_v1_PlanPolicyStatus(ref),
		"github.com/appvia/kore/pkg/apis/config/v1.PlanSpec":           schema_pkg_apis_config_v1_PlanSpec(ref),
		"github.com/appvia/kore/pkg/apis/config/v1.PlanStatus":         schema_pkg_apis_config_v1_PlanStatus(ref),
		"github.com/appvia/kore/pkg/apis/config/v1.Secret":             schema_pkg_apis_config_v1_Secret(ref),
		"github.com/appvia/kore/pkg/apis/config/v1.SecretSpec":         schema_pkg_apis_config_v1_SecretSpec(ref),
		"github.com/appvia/kore/pkg/apis/config/v1.SecretStatus":       schema_pkg_apis_config_v1_SecretStatus(ref),
	}
}

func schema_pkg_apis_config_v1_Allocation(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Allocation is the Schema for the allocations API",
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
							Ref: ref("github.com/appvia/kore/pkg/apis/config/v1.AllocationSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/appvia/kore/pkg/apis/config/v1.AllocationStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/config/v1.AllocationSpec", "github.com/appvia/kore/pkg/apis/config/v1.AllocationStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_config_v1_AllocationSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AllocationSpec defines the desired state of Allocation",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name is the name of the resource being shared",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"summary": {
						SchemaProps: spec.SchemaProps{
							Description: "Summary is the summary of the resource being shared",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"resource": {
						SchemaProps: spec.SchemaProps{
							Description: "Resource is the resource which is being shared with another team",
							Ref:         ref("github.com/appvia/kore/pkg/apis/core/v1.Ownership"),
						},
					},
					"teams": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Teams is a collection of teams the allocation is permitted to use",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
				Required: []string{"name", "summary", "resource", "teams"},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/core/v1.Ownership"},
	}
}

func schema_pkg_apis_config_v1_AllocationStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AllocationStatus defines the observed state of Allocation",
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

func schema_pkg_apis_config_v1_Plan(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Plan is the Schema for the plans API",
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
							Ref: ref("github.com/appvia/kore/pkg/apis/config/v1.PlanSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/appvia/kore/pkg/apis/config/v1.PlanStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/config/v1.PlanSpec", "github.com/appvia/kore/pkg/apis/config/v1.PlanStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_config_v1_PlanPolicy(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PlanPolicy is the Schema for the plan policies API",
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
							Ref: ref("github.com/appvia/kore/pkg/apis/config/v1.PlanPolicySpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/appvia/kore/pkg/apis/config/v1.PlanPolicyStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/config/v1.PlanPolicySpec", "github.com/appvia/kore/pkg/apis/config/v1.PlanPolicyStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_config_v1_PlanPolicyProperty(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PlanPolicyProperty defines a JSON schema for a given property",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name is the name of the property",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"allowUpdate": {
						SchemaProps: spec.SchemaProps{
							Description: "AllowUpdate will allow the parameter to be modified by the teams",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"disallowUpdate": {
						SchemaProps: spec.SchemaProps{
							Description: "DisallowUpdate will forbid modification of the parameter, even if it was allowed by an other policy",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
				Required: []string{"name", "allowUpdate", "disallowUpdate"},
			},
		},
	}
}

func schema_pkg_apis_config_v1_PlanPolicySpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PlanPolicySpec defines Plan JSON Schema extensions",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind refers to the cluster type this is a plan policy for",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"labels": {
						SchemaProps: spec.SchemaProps{
							Description: "Labels is a collection of labels for this plan policy",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"summary": {
						SchemaProps: spec.SchemaProps{
							Description: "Summary provides a short title summary for the plan policy",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"description": {
						SchemaProps: spec.SchemaProps{
							Description: "Description provides a detailed description of the plan policy",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"properties": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": []interface{}{
									"name",
								},
								"x-kubernetes-list-type": "map",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Properties are the",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/appvia/kore/pkg/apis/config/v1.PlanPolicyProperty"),
									},
								},
							},
						},
					},
				},
				Required: []string{"kind", "summary", "description", "properties"},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/config/v1.PlanPolicyProperty"},
	}
}

func schema_pkg_apis_config_v1_PlanPolicyStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PlanPolicyStatus defines the observed state of Plan Policy",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Conditions is a set of condition which has caused an error",
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
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status is overall status of the plan policy",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"conditions", "status"},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/core/v1.Condition"},
	}
}

func schema_pkg_apis_config_v1_PlanSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PlanSpec defines the desired state of Plan",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Resource refers to the resource type this is a plan for",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"labels": {
						SchemaProps: spec.SchemaProps{
							Description: "Labels is a collection of labels for this plan",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"description": {
						SchemaProps: spec.SchemaProps{
							Description: "Description provides a summary of the configuration provided by this plan",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"summary": {
						SchemaProps: spec.SchemaProps{
							Description: "Summary provides a short title summary for the plan",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"configuration": {
						SchemaProps: spec.SchemaProps{
							Description: "Configuration are the key+value pairs describing a cluster configuration",
							Ref:         ref("k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1.JSON"),
						},
					},
				},
				Required: []string{"kind", "description", "summary", "configuration"},
			},
		},
		Dependencies: []string{
			"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1.JSON"},
	}
}

func schema_pkg_apis_config_v1_PlanStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PlanStatus defines the observed state of Plan",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Conditions is a set of condition which has caused an error",
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
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status is overall status of the workspace",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"conditions", "status"},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/core/v1.Condition"},
	}
}

func schema_pkg_apis_config_v1_Secret(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Secret is the Schema for the plans API",
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
							Ref: ref("github.com/appvia/kore/pkg/apis/config/v1.SecretSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/appvia/kore/pkg/apis/config/v1.SecretStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/config/v1.SecretSpec", "github.com/appvia/kore/pkg/apis/config/v1.SecretStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_config_v1_SecretSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SecretSpec defines the desired state of Plan",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "Type refers to the secret type",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"description": {
						SchemaProps: spec.SchemaProps{
							Description: "Description provides a summary of the secret",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"data": {
						SchemaProps: spec.SchemaProps{
							Description: "Values are the key values to the plan",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
				Required: []string{"type", "description"},
			},
		},
	}
}

func schema_pkg_apis_config_v1_SecretStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SecretStatus defines the observed state of Plan",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"systemManaged": {
						SchemaProps: spec.SchemaProps{
							Description: "SystemManaged indicates the secret is managed by kore and cannot be changed",
							Type:        []string{"boolean"},
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
							Description: "Conditions is a set of condition which has caused an error",
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
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status is overall status of the workspace",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"verified": {
						SchemaProps: spec.SchemaProps{
							Description: "Verified indicates if the secret has been verified as working",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/core/v1.Condition"},
	}
}
