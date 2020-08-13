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
	"fmt"

	corev1 "github.com/appvia/kore/pkg/apis/core/v1"
	"github.com/appvia/kore/pkg/utils/kubernetes"
)

type ProviderData struct {
	Resources []corev1.Ownership `json:"resources,omitempty"`
}

type AppConfiguration struct {
	Resources kubernetes.Objects
	Values    map[string]interface{}
	Secrets   map[string]interface{}
}

func (a AppConfiguration) CompileResources(params ResourceParams) (kubernetes.Objects, error) {
	var compiledResources kubernetes.Objects
	for _, r := range a.Resources {
		compiled, err := compileResource(r.DeepCopyObject(), params)
		if err != nil {
			return nil, fmt.Errorf("compiling resource %v failed: %w", r, err)
		}
		compiledResources = append(compiledResources, compiled)
	}
	return compiledResources, nil
}
