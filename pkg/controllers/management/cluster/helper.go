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

package cluster

import (
	corev1 "github.com/appvia/kore/pkg/apis/core/v1"
	"github.com/appvia/kore/pkg/utils/kubernetes"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// GetObjectStatus attempts to inspect the resource for a status
func GetObjectStatus(object runtime.Object) (corev1.Status, error) {
	var status corev1.Status

	return status, kubernetes.GetRuntimeField(object, "status.status", &status)
}

// GetObjectReasonForFailure try's to get the reason for failure
func GetObjectReasonForFailure(object runtime.Object) (corev1.Condition, error) {
	c, err := GetObjectComponents(object)
	if err == nil {
		return c, nil
	}

	return GetObjectConditions(object)
}

// GetObjectComponents attempts to inspect the resource components
func GetObjectComponents(object runtime.Object) (corev1.Condition, error) {
	var components corev1.Components
	if err := kubernetes.GetRuntimeField(object, "status.components", &components); err != nil {
		return corev1.Condition{}, err
	}

	if components != nil && len(components) > 0 {
		return corev1.Condition{
			Detail:  components[0].Detail,
			Message: components[0].Message,
		}, nil
	}

	return corev1.Condition{}, kubernetes.ErrFieldNotFound
}

// GetObjectConditions returns the conditions on a resource
func GetObjectConditions(object runtime.Object) (corev1.Condition, error) {
	var conditions []corev1.Condition

	if err := kubernetes.GetRuntimeField(object, "status.conditions", &conditions); err != nil {
		return corev1.Condition{}, err
	}

	return conditions[0], nil
}

// IsDeleting check if the resource is being deleted
func IsDeleting(object runtime.Object) bool {
	mo, _ := object.(metav1.Object)

	return !mo.GetDeletionTimestamp().IsZero()
}

// ComponentToUnstructured converts the component to a runtime reference
func ComponentToUnstructured(component *corev1.Component) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   component.Resource.Group,
		Version: component.Resource.Version,
		Kind:    component.Resource.Kind,
	})
	u.SetNamespace(component.Resource.Namespace)
	u.SetName(component.Resource.Name)

	return u
}
