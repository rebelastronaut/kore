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

package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/appvia/kore/pkg/utils"

	"gopkg.in/yaml.v2"
)

// GetClientConfigurationPath returns the path to the client config
func GetClientConfigurationPath() string {
	// @step: retrieve the configuration path from env of default path
	path := os.ExpandEnv(os.Getenv(DefaultKoreConfigPathEnv))
	if path == "" {
		path = os.ExpandEnv(DefaultKoreConfigPath)
	}

	return path
}

// GetOrCreateClientConfiguration is responsible for retrieving the client configuration
func GetOrCreateClientConfiguration() (*Config, error) {
	path := GetClientConfigurationPath()

	// @step: check the file exists else create it
	if found, err := utils.FileExists(path); err != nil {
		return nil, err
	} else if !found {
		// @step: we need to write an empty file for now
		if err := UpdateConfig(NewEmpty(), path); err != nil {
			return nil, err
		}

		return NewEmpty(), nil
	}

	// @step: open the configuration file for reading
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return New(file)
}

// UpdateConfig is responsible for writing the configuration to disk
func UpdateConfig(config *Config, path string) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), os.FileMode(0750)); err != nil {
		return err
	}

	return ioutil.WriteFile(path, data, os.FileMode(0640))
}
