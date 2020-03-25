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

package korectl

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"

	clustersv1 "github.com/appvia/kore/pkg/apis/clusters/v1"
	configv1 "github.com/appvia/kore/pkg/apis/config/v1"
	corev1 "github.com/appvia/kore/pkg/apis/core/v1"
	"github.com/appvia/kore/pkg/apiserver/types"
	"github.com/appvia/kore/pkg/utils"

	yml "github.com/ghodss/yaml"
	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

// Document defines a rest endpoint
type Document struct {
	// Endpoint is the rest storage endpoint
	Endpoint string
	// Object the resource to send
	Object *unstructured.Unstructured
}

var (
	hostnameRegex = regexp.MustCompile(`^https?://([0-9a-zA-Z\.]+)|([0-9]{1,3}\.){3,3}[0-9]{1,3}(:[0-9]+)?$`)
)

// IsValidHostname checks the endpoint is valid
func IsValidHostname(endpoint string) bool {
	return hostnameRegex.MatchString(endpoint)
}

// RemoveDoubleQuote removes all double quotes from string
func RemoveDoubleQuote(v string) string {
	return strings.ReplaceAll(v, "\"", "")
}

// ReadInContent reads from a file or stdin
func ReadInContent(path string) ([]byte, error) {
	if path == "-" {
		return ioutil.ReadAll(os.Stdin)
	}

	return ioutil.ReadFile(path)
}

// GetWhoAmI returns an whoami
func GetWhoAmI(config *Config) (*types.WhoAmI, error) {
	who := &types.WhoAmI{}

	return who, NewRequest().
		WithConfig(config).
		WithEndpoint("/whoami").
		WithRuntimeObject(who).
		Get()
}

// GetCluster returns the cluster object
func GetCluster(config *Config, team, name string) (*clustersv1.Kubernetes, error) {
	cluster := &clustersv1.Kubernetes{}

	return cluster, GetTeamResource(config, team, "clusters", name, cluster)
}

// CreateTeamResource checks if a resources exists in the team
func CreateTeamResource(config *Config, team, kind, name string, object runtime.Object) error {
	kind = strings.ToLower(utils.ToPlural(kind))

	return NewRequest().
		WithConfig(config).
		PathParameter("team", true).
		PathParameter("kind", true).
		PathParameter("name", true).
		WithInject("team", team).
		WithInject("kind", kind).
		WithInject("name", name).
		WithEndpoint("/teams/{team}/{kind}/{name}").
		WithRuntimeObject(object).
		Update()
}

// ResourceExists checks if a team resource exists
func ResourceExists(config *Config, kind, name string) (bool, error) {
	kind = strings.ToLower(utils.ToPlural(kind))

	return NewRequest().
		WithConfig(config).
		PathParameter("kind", true).
		PathParameter("name", true).
		WithInject("kind", kind).
		WithInject("name", name).
		WithEndpoint("{kind}/{name}").
		Exists()
}

// TeamResourceExists checks if a resources exists in the team
func TeamResourceExists(config *Config, team, kind, name string) (bool, error) {
	kind = strings.ToLower(utils.ToPlural(kind))

	return NewRequest().
		WithConfig(config).
		PathParameter("team", true).
		PathParameter("kind", true).
		PathParameter("name", true).
		WithInject("team", team).
		WithInject("kind", kind).
		WithInject("name", name).
		WithEndpoint("/teams/{team}/{kind}/{name}").
		Exists()
}

// GetTeamResourceList returns a collection of resources - essentially minus the name
func GetTeamResourceList(config *Config, team, kind string, object runtime.Object) error {
	kind = strings.ToLower(utils.ToPlural(kind))

	return NewRequest().
		WithConfig(config).
		PathParameter("team", true).
		PathParameter("kind", true).
		WithInject("team", team).
		WithInject("kind", kind).
		WithEndpoint("/teams/{team}/{kind}").
		WithRuntimeObject(object).
		Get()
}

// GetTeamAllocation returns an allocation for a team
func GetTeamAllocation(config *Config, team, name string) (*configv1.Allocation, error) {
	o := &configv1.Allocation{}

	return o, GetTeamResource(config, team, "allocation", name, o)
}

// GetTeamResource returns a team object
func GetTeamResource(config *Config, team, kind, name string, object interface{}) error {
	kind = strings.ToLower(utils.ToPlural(kind))

	return NewRequest().
		WithConfig(config).
		PathParameter("team", true).
		PathParameter("kind", true).
		PathParameter("name", true).
		WithInject("team", team).
		WithInject("kind", kind).
		WithInject("name", name).
		WithEndpoint("/teams/{team}/{kind}/{name}").
		WithRuntimeObject(object).
		Get()
}

// GetResource returns a global resource object
func GetResource(config *Config, kind, name string, object runtime.Object) error {
	kind = strings.ToLower(utils.ToPlural(kind))

	return NewRequest().
		WithConfig(config).
		PathParameter("kind", true).
		PathParameter("name", true).
		WithInject("kind", kind).
		WithInject("name", name).
		WithEndpoint("/{kind}/{name}").
		WithRuntimeObject(object).
		Get()
}

// GetResourceList returns a list of global resource types
func GetResourceList(config *Config, team, kind, name string, object runtime.Object) error {
	kind = strings.ToLower(utils.ToPlural(kind))

	return NewRequest().
		WithConfig(config).
		PathParameter("kind", true).
		PathParameter("name", true).
		WithInject("kind", kind).
		WithInject("name", name).
		WithEndpoint("/{kind}/{name}").
		WithRuntimeObject(object).
		Get()
}

// WaitOnResource indicates we should wait for the resource to transition to fail or success
func WaitOnResource(ctx context.Context, config *Config, team, kind, name string, interval, timeout time.Duration) (bool, error) {
	var success bool

	u := make(map[string]interface{})
	var max int

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		<-sig
		cancel()
	}()

	err := utils.WaitUntilComplete(ctx, timeout, interval, func() (bool, error) {
		err := GetTeamResource(config, team, kind, name, &u)
		if err != nil {
			return false, nil
		}

		status, ok := u["status"].(map[string]interface{})
		if !ok {
			return false, nil
		}
		state, ok := status["status"].(string)
		if !ok {
			return false, nil
		}

		if state == string(corev1.FailureStatus) {
			if max > 3 {
				return false, errors.New("resource has failed to provision")
			}
			max++
		}
		if state == string(corev1.SuccessStatus) {
			success = true

			return true, nil
		}

		return false, nil
	})

	return success, err
}

// ParseDocument returns a collection of parsed documents and the api endpoints
func ParseDocument(src io.Reader, namespace string) ([]*Document, error) {
	var list []*Document

	// @step: read in the content of the file
	content, err := ioutil.ReadAll(src)
	if err != nil {
		return nil, err
	}
	// @step: split the yaml documents up
	splitter := regexp.MustCompile("(?m)^---\n")
	documents := splitter.Split(string(content), -1)

	for _, x := range documents {
		if x == "" {
			continue
		}

		doc, err := yml.YAMLToJSON([]byte(x))
		if err != nil {
			return nil, err
		}

		// @step: attempt to read the document into an unstructured
		u := &unstructured.Unstructured{}
		if err := u.UnmarshalJSON(doc); err != nil {
			return nil, err
		}

		// @checks
		// - ensure we have a name
		// - ensure we have a api kind
		if u.GetName() == "" {
			return nil, errors.New("resource must have a name")
		}
		if u.GetKind() == "" {
			return nil, errors.New("resource must have an api kind")
		}
		if u.GetAPIVersion() == "" {
			return nil, errors.New("resource requires an api group")
		}

		// @step: we pluralize the kind and use that route the resource
		kind := strings.ToLower(utils.ToPlural(u.GetKind()))

		team := u.GetNamespace()
		if !IsGlobalResource(kind) {
			if namespace != "" {
				if team != "" && team != namespace {
					return nil, errors.New("resource name and team selected are different")
				}
				team = namespace
			}
			if team == "" {
				return nil, errors.New("all resources must have a team namespace")
			}
		}
		u.SetNamespace(team)
		team = strings.ToLower(team)
		name := strings.ToLower(u.GetName())

		remapping := map[string]string{
			"kubernetes": "clusters",
		}
		for k, v := range remapping {
			if k == kind {
				kind = v
			}
		}

		item := &Document{Object: u}
		switch IsGlobalResource(kind) {
		case true:
			item.Endpoint = fmt.Sprintf("%s/%s", kind, name)
		default:
			item.Endpoint = fmt.Sprintf("teams/%s/%s/%s", team, kind, name)
		}

		list = append(list, item)
	}

	return list, nil
}

// IsGlobalResource is a global resource
func IsGlobalResource(name string) bool {
	return utils.Contains(name, []string{"teams", "users", "plans"})
}

// GetCaches is responsible for checking if are caches are up to date
func GetCaches(config *Config) error {
	/*
		content, err := GetSwaggerCache(*config)
		if err != nil {
			return err
		}
		_, err = fastjson.Parse(string(content))
		if err != nil {
			return err
		}
	*/

	return nil
}

// GetOrCreateKubeConfig is used to retrieve the kubeconfig path
func GetOrCreateKubeConfig() (string, error) {
	path := func() string {
		p := os.ExpandEnv(os.Getenv("$KUBECONFIG"))
		if p != "" {
			return p
		}

		return os.ExpandEnv("${HOME}/.kube/config")
	}()

	_, err := utils.EnsureFileExists(path)
	if err != nil {
		return "", err
	}

	return path, nil
}

// BuildResourcesFromSwagger builds a list of global and namespaces resources
// from the swagger

// GetSwaggerCache is responsible for updating the swagger cache
func GetSwaggerCache(config Config) ([]byte, error) {
	log.Debug("checking for cached resources file")

	current, err := config.GetSwagger()
	if err != nil {
		log.WithError(err).Error("failed read in cached swagger")

		return nil, err
	}
	// @step: we need to check if the swagger is up to date
	checksum, err := GetFileChecksum(config.GetSwaggerCachedFile())
	if err != nil {
		log.WithError(err).Warn("failed checking the swagger cache file")

		return current, nil
	}
	latest, err := config.GetSwaggerChecksum()
	if err != nil {
		log.WithError(err).Debug("failed retrieving latest swagger cache file")
	}
	if checksum == latest {
		return current, nil
	}

	// @step: else we need to download the latest
	update, err := config.GetSwaggerFromAPI()
	if err != nil {
		log.WithError(err).Debug("failed to retrieve swagger from apiserver")

		return current, nil
	}
	if err := config.UpdateSwaggerCache(update); err != nil {
		log.WithError(err).Error("failed to update the swagger cache")
	}

	return update, nil
}

// GetFileChecksum returns the checksum of a content of a file
func GetFileChecksum(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	h := sha256.New()
	_, _ = h.Write(content)

	return hex.EncodeToString(h.Sum(nil)), nil
}

// GetOrCreateClientConfiguration is responsible for retrieving the client configuration
func GetOrCreateClientConfiguration() (*Config, error) {
	config := &Config{}
	var content []byte

	configPath := os.ExpandEnv(HubConfig)

	if _, err := os.Stat(configPath); err == nil {
		content, err = ioutil.ReadFile(configPath)
		if err != nil {
			return config, err
		}
	} else if os.IsNotExist(err) {
		err := config.Update()
		if err != nil {
			return config, err
		}
	} else {
		return config, err
	}

	if strings.TrimSpace(string(content)) == "" {
		return config, nil
	}

	// @step: parse the configuration
	return config, yaml.NewDecoder(bytes.NewReader(content)).Decode(config)
}

func formatLongDescription(desc string) string {
	var res strings.Builder
	for n, line := range strings.Split(strings.Trim(desc, " \n\t"), "\n") {
		if n == 0 {
			res.WriteString(strings.TrimSpace(line))
		} else {
			res.WriteRune('\n')
			res.WriteString("   ")
			res.WriteString(line)
		}
	}
	return res.String()
}
