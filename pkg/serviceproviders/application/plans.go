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

package application

import (
	servicev1 "github.com/appvia/kore/pkg/apis/services/v1"
	"github.com/appvia/kore/pkg/kore"
	"github.com/appvia/kore/pkg/utils/jsonutils"
	apiextv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// HelmAppClusterAutoscaler is the plan name for instances of the cluster autoscaler
	HelmAppClusterAutoscaler = "helm-app-cluster-autoscaler"
	// HelmAppKoreMonitoring is the plan name of the kore monitoring plan
	HelmAppKoreMonitoring = "helm-app-kore-monitoring"
)

// GetDefaultPlans returns a collection of plans for the resources
func GetDefaultPlans() []servicev1.ServicePlan {
	return []servicev1.ServicePlan{
		{
			TypeMeta: metav1.TypeMeta{
				Kind:       "ServicePlan",
				APIVersion: servicev1.GroupVersion.String(),
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      HelmAppClusterAutoscaler,
				Namespace: "kore",
				Annotations: map[string]string{
					kore.AnnotationSystem:      kore.AnnotationValueTrue,
					kore.AnnotationInstallOnce: kore.AnnotationValueTrue,
					kore.AnnotationReadOnly:    kore.AnnotationValueTrue,
				},
			},
			Spec: servicev1.ServicePlanSpec{
				Kind:        "helm-app",
				Summary:     "Autoscaler plan",
				Description: "Cluster Autoscaler Plan",
				Configuration: &apiextv1.JSON{
					Raw: []byte(`{
						"source": {
							"helm": {
								"url": "https://kubernetes-charts.storage.googleapis.com/",
								"name": "cluster-autoscaler",
								"version": "7.3.2"
							}
						},
						"resourceKinds": [
							{
								"group": "apps",
								"kind": "Deployment"
							}
						],
						"resourceSelector": {
							"matchLabels": {
								"app.kubernetes.io/name": "aws-cluster-autoscaler"
							}
						}
					}`),
				},
				Schema: string(jsonutils.MustCompact([]byte(helmAppSchemaV1))),
			},
		},
		{
			TypeMeta: metav1.TypeMeta{
				Kind:       "ServicePlan",
				APIVersion: servicev1.GroupVersion.String(),
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      HelmAppKoreMonitoring,
				Namespace: "kore",
				Annotations: map[string]string{
					kore.AnnotationSystem:      kore.AnnotationValueTrue,
					kore.AnnotationInstallOnce: kore.AnnotationValueTrue,
					kore.AnnotationReadOnly:    kore.AnnotationValueTrue,
				},
			},
			Spec: servicev1.ServicePlanSpec{
				Kind:        "helm-app",
				Summary:     "Managed Cluster Monitoring",
				DisplayName: "Kore Cluster Monitoring",
				Description: "Kore Monitoring service provides a monitoring stack used to ensure the health of the clusters and the services",
				Configuration: &apiextv1.JSON{
					Raw: []byte(`{
						"source": {
							"helm": {
								"url": "https://storage.googleapis.com/kore-charts",
								"name": "kore-monitoring",
								"version": "0.0.1"
							}
						},
						"resourceKinds": [
							{
								"group": "apps",
								"kind": "Deployment"
							}
						],
						"resourceSelector": {
							"matchLabels": {
								"app.kubernetes.io/name": "monitoring-operator"
							}
						}
					}`),
				},
				Schema: string(jsonutils.MustCompact([]byte(helmAppSchemaV1))),
			},
		},
	}
}
