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

package localjwt

import (
	"errors"
)

// IsValid checks the configuration is valid
func (c Config) IsValid() error {
	if c.PublicKey == "" && c.PublicKeyPath == "" {
		return errors.New("no public key configured")
	}

	return nil
}

// HasAudience checks if we have a audience requiremnt
func (c *Config) HasAudience() bool {
	return c.Audience != ""
}
