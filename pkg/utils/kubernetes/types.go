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

package kubernetes

import (
	"fmt"
)

// DependentReference is an object reference to a dependent object in the same namespace
type DependentReference struct {
	// API version of the dependent
	APIVersion string `json:"apiVersion"`
	// Kind of the dependent
	Kind string `json:"kind"`
	// Name of the dependent
	Name string `json:"name"`
}

func (d DependentReference) String() string {
	return fmt.Sprintf("%s/%s/%s", d.APIVersion, d.Kind, d.Name)
}

// DependentReferenceFromObject creates a DependentReference from the given object
func DependentReferenceFromObject(o Object) DependentReference {
	return DependentReference{
		APIVersion: o.GetObjectKind().GroupVersionKind().GroupVersion().String(),
		Kind:       o.GetObjectKind().GroupVersionKind().Kind,
		Name:       o.GetName(),
	}
}

// KubernetesAPI is the configuration for the kubernetes api
type KubernetesAPI struct {
	// InCluster indicates we are running in cluster
	InCluster bool
	// MasterAPIURL specifies the kube-apiserver url
	MasterAPIURL string
	// Token is kubernetes token to authenticate to the api
	Token string
	// KubeConfig is the kubeconfig path
	KubeConfig string
	// SkipTLSVerify indicates we skip tls
	SkipTLSVerify bool
}
