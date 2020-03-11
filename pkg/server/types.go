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

package server

import (
	"context"

	api "github.com/appvia/kore/pkg/apiserver"
	"github.com/appvia/kore/pkg/kore"
	"github.com/appvia/kore/pkg/services/users"
)

// Interface is the contract to the server
type Interface interface {
	// Run is responsible for starting the services
	Run(context.Context) error
	// Stop is responsible for trying to stop services
	Stop(context.Context) error
}

// KubernetesAPI is the configuration for the kubernetes api
type KubernetesAPI struct {
	// InCluster indicates we are running in cluster
	InCluster bool `json:"inCluster"`
	// MasterAPIURL specifies the kube-apiserver url
	MasterAPIURL string `json:"masterAPIUrl"`
	// Token is kubernetes token to authenticate to the api
	Token string `json:"token"`
	// KubeConfig is the kubeconfig path
	KubeConfig string
	// SkipTLSVerify indicates we skip tls
	SkipTLSVerify bool
}

// Config is the configuration of the various components
type Config struct {
	// APIServer is the config for the api server
	APIServer api.Config `json:"apiServer"`
	// Kubernetes is configuration for the api
	Kubernetes KubernetesAPI `json:"kubernetes"`
	// Kore is the configuration for the kore bridge
	Kore kore.Config `json:"kore"`
	// UsersMgr are the user management service options
	UsersMgr users.Config `json:"usersMgr"`
}
