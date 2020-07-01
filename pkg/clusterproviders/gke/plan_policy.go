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

package gke

import (
	configv1 "github.com/appvia/kore/pkg/apis/config/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var planPolicy = configv1.PlanPolicy{
	ObjectMeta: metav1.ObjectMeta{
		Name: "default-gke",
		Annotations: map[string]string{
			"kore.appvia.io/readonly": "true",
		},
	},
	Spec: configv1.PlanPolicySpec{
		Kind:        Kind,
		Summary:     "Default plan policy for GKE clusters",
		Description: "This policy defines which plan properties can be edited by default for GKE clusters",
		Properties: []configv1.PlanPolicyProperty{
			{Name: "authProxyAllowedIPs", AllowUpdate: true},
			{Name: "clusterUsers", AllowUpdate: true},
			{Name: "defaultTeamRole", AllowUpdate: true},
			{Name: "description", AllowUpdate: true},
			{Name: "diskSize", AllowUpdate: true},
			{Name: "domain", AllowUpdate: true},
			{Name: "maxSize", AllowUpdate: true},
			{Name: "size", AllowUpdate: true},
			{Name: "releaseChannel", AllowUpdate: true},
			{Name: "version", AllowUpdate: true},
			{Name: "nodePools", AllowUpdate: true},
		},
	},
}
