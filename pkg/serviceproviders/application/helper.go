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
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"text/template"

	"github.com/appvia/kore/pkg/utils/configuration"
	servicesv1 "github.com/appvia/kore/pkg/apis/services/v1"
	corev1 "github.com/appvia/kore/pkg/apis/core/v1"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	applicationv1beta "sigs.k8s.io/application/api/v1beta1"
	"github.com/appvia/kore/pkg/kore"
	koreschema "github.com/appvia/kore/pkg/schema"
	"github.com/appvia/kore/pkg/utils/kubernetes"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ResourceParams struct {
	Release Release
	Values  map[string]interface{}
	Secrets map[string]interface{}
}

type Release struct {
	Name      string
	Namespace string
}

func NewResourceParams(service *servicesv1.Service, config *AppConfiguration) ResourceParams {
	return ResourceParams{
		Release: Release{
			Name:      service.Name,
			Namespace: service.Spec.ClusterNamespace,
		},
		Values:  config.Values,
		Secrets: config.Secrets,
	}
}

func CreateSystemServiceFromPlan(servicePlan servicesv1.ServicePlan, cluster corev1.Ownership, name, namespace string) servicesv1.Service {
	config := &AppConfiguration{}
	if err := servicePlan.Spec.GetConfiguration(config); err != nil {
		// This should not happen
		panic(err)
	}

	clusterNamespace := servicePlan.Name
	for _, resource := range config.Resources {
		if ns, ok := resource.(*v1.Namespace); ok {
			clusterNamespace = ns.Name
		}
	}

	var priority string
	if servicePlan.Name == "app-"+kore.AppAppManager {
		priority = "1"
	}

	return servicesv1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: servicesv1.GroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Annotations: map[string]string{
				kore.AnnotationPriority: priority,
				kore.AnnotationSystem:   kore.AnnotationValueTrue,
				kore.AnnotationReadOnly: kore.AnnotationValueTrue,
			},
		},
		Spec: servicesv1.ServiceSpec{
			Kind:             servicePlan.Spec.Kind,
			Plan:             servicePlan.Name,
			Cluster:          cluster,
			ClusterNamespace: clusterNamespace,
			Configuration:    servicePlan.Spec.Configuration,
		},
	}
}

func compileResource(obj runtime.Object, params ResourceParams) (runtime.Object, error) {
	document, err := json.Marshal(obj)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal object %w", err)
	}

	tmpl, err := template.
		New("document").
		Funcs(template.FuncMap{
			"json": func(v interface{}) (interface{}, error) {
				val, err := json.Marshal(v)
				if err != nil {
					return nil, err
				}
				return string(val), nil
			},
			"jsonb64": func(v interface{}) (interface{}, error) {
				val, err := json.Marshal(v)
				if err != nil {
					return nil, err
				}
				return base64.StdEncoding.EncodeToString(val), nil
			},
			"sha1": func(v interface{}) (interface{}, error) {
				val, err := json.Marshal(v)
				if err != nil {
					return nil, err
				}
				h := sha1.New()
				_, err = h.Write(val)
				if err != nil {
					return nil, err
				}
				bs := h.Sum(nil)
				return base64.StdEncoding.EncodeToString(bs), nil
			},
		}).
		Parse(string(document))
	if err != nil {
		return nil, fmt.Errorf("failed to parse document as template %w", err)
	}
	tmplBuf := bytes.NewBuffer(make([]byte, 0, 16384))
	if err := tmpl.Execute(tmplBuf, params); err != nil {
		return nil, fmt.Errorf("failed to apply templating %w", err)
	}

	o, err := koreschema.DecodeJSON(tmplBuf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("failed to decode templated output %v %w", tmplBuf.String(), err)
	}

	return o, nil
}

func ensureResource(ctx kore.Context, client client.Client, original runtime.Object) error {
	var current runtime.Object
	if koreschema.GetScheme().Recognizes(original.GetObjectKind().GroupVersionKind()) {
		var err error
		if current, err = koreschema.GetScheme().New(original.GetObjectKind().GroupVersionKind()); err != nil {
			return err
		}
	} else {
		current = &unstructured.Unstructured{}
	}

	currentMeta, err := meta.Accessor(current)
	if err != nil {
		return err
	}

	originalMeta, err := meta.Accessor(original)
	if err != nil {
		return err
	}

	current.GetObjectKind().SetGroupVersionKind(original.GetObjectKind().GroupVersionKind())
	currentMeta.SetName(originalMeta.GetName())
	currentMeta.SetNamespace(originalMeta.GetNamespace())

	exists, err := kubernetes.GetIfExists(ctx, client, current)
	if err != nil {
		return fmt.Errorf("failed to get resource %q: %w", kubernetes.MustGetRuntimeSelfLink(current), err)
	}

	if exists {
		// The runtime client doesn't set the GVK on the result object
		current.GetObjectKind().SetGroupVersionKind(original.GetObjectKind().GroupVersionKind())
	}

	updated, err := kubernetes.UpdateIfChangedSinceLastUpdate(ctx, client, original, current)
	if updated {
		ctx.Logger().WithField("resource", kubernetes.MustGetRuntimeSelfLink(original)).Debug("resource has changed")
	}
	return err
}

func getAppConfiguration(ctx kore.Context, service *servicesv1.Service) (*AppConfiguration, error) {
	switch service.Spec.Kind {
	case ServiceKindApp:
		config := &AppConfiguration{}
		if err := configuration.ParseObjectConfiguration(ctx, ctx.Client(), service, config); err != nil {
			return nil, err
		}
		return config, nil
	case ServiceKindHelmApp:
		helmConfig := &HelmAppConfiguration{}
		if err := configuration.ParseObjectConfiguration(ctx, ctx.Client(), service, helmConfig); err != nil {
			return nil, err
		}

		var chart map[string]interface{}
		if helmConfig.Source.HelmRepository != nil {
			chart = map[string]interface{}{
				"repository": helmConfig.Source.HelmRepository.URL,
				"version":    helmConfig.Source.HelmRepository.Version,
				"name":       helmConfig.Source.HelmRepository.Name,
			}
		}
		if helmConfig.Source.GitRepository != nil {
			if chart != nil {
				return nil, fmt.Errorf("only one Helm chart source should be defined")
			}
			chart = map[string]interface{}{
				"git":  helmConfig.Source.GitRepository.URL,
				"path": helmConfig.Source.GitRepository.Path,
				"ref":  helmConfig.Source.GitRepository.Ref,
			}
		}
		if chart == nil {
			return nil, fmt.Errorf("one Helm chart source must be defined")
		}

		values := map[string]interface{}{}
		if len(helmConfig.Values) > 0 {
			values = helmConfig.Values
		}

		helmRelease := &unstructured.Unstructured{Object: map[string]interface{}{
			"apiVersion": "helm.fluxcd.io/v1",
			"kind":       "HelmRelease",
			"metadata": map[string]interface{}{
				"name":      service.Name,
				"namespace": service.Spec.ClusterNamespace,
			},
			"spec": map[string]interface{}{
				"releaseName": service.Name,
				"chart":       chart,
				"values":      values,
			},
		}}

		app := &applicationv1beta.Application{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Application",
				APIVersion: applicationv1beta.GroupVersion.String(),
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      service.Name,
				Namespace: service.Spec.ClusterNamespace,
			},
			Spec: applicationv1beta.ApplicationSpec{
				ComponentGroupKinds: helmConfig.ResourceKinds,
			},
		}
		if helmConfig.ResourceSelector != nil {
			app.Spec.Selector = &metav1.LabelSelector{
				MatchLabels: helmConfig.ResourceSelector.MatchLabels,
			}
		} else {
			app.Spec.Selector = &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app.kubernetes.io/name": service.Name,
				},
			}
		}

		return &AppConfiguration{
			Resources: []runtime.Object{
				helmRelease,
				app,
			},
			Values: nil,
		}, nil
	default:
		panic(fmt.Errorf("unexpected service kind: %s", service.Spec.Kind))
	}
}
