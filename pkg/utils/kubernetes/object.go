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
	"bytes"
	"fmt"
	"regexp"
	"strings"

	applicationv1beta "sigs.k8s.io/application/api/v1beta1"
	"sigs.k8s.io/yaml"

	corev1 "github.com/appvia/kore/pkg/apis/core/v1"
	koreschema "github.com/appvia/kore/pkg/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// Object is a Kubernetes object
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . Object
type Object interface {
	runtime.Object
	metav1.Object
}

// ObjectWithStatus is a Kubernetes object where you can set/get the status and manage the status components
type ObjectWithStatus interface {
	Object
	GetStatus() (status corev1.Status, message string)
	SetStatus(status corev1.Status, message string)
}

type ObjectWithStatusComponents interface {
	Object
	StatusComponents() *corev1.Components
}

// NewObject creates a new object given the GVK definition
func NewObject(gvk schema.GroupVersionKind) (Object, error) {
	ro, err := koreschema.GetScheme().New(gvk)
	if err != nil {
		return nil, err
	}

	if o, ok := ro.(Object); ok {
		return o, nil
	}

	return nil, fmt.Errorf("%T object doesn't implement kubernetes.Object", ro)
}

type Objects []runtime.Object

func (o Objects) Application() *applicationv1beta.Application {
	for _, res := range o {
		if app, ok := res.(*applicationv1beta.Application); ok {
			return app
		}
	}
	return nil
}

func (o Objects) MarshalYAML() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0, 16384))
	for _, obj := range o {
		yamlData, err := yaml.Marshal(obj)
		if err != nil {
			return nil, err
		}
		buf.WriteString("---\n")
		buf.Write(yamlData)
		buf.WriteRune('\n')
	}
	return buf.Bytes(), nil
}

func (o *Objects) UnmarshalYAML(data []byte) error {
	if len(data) == 0 {
		return nil
	}

	documents := regexp.MustCompile("(?m)^---\n").Split(string(data), -1)

	var objects []runtime.Object
	for _, document := range documents {
		if strings.TrimSpace(document) == "" {
			continue
		}

		obj, err := koreschema.DecodeYAML([]byte(document))
		if err != nil {
			return err
		}

		objects = append(objects, obj)
	}

	*o = objects
	return nil
}
