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

package authproxy

import (
	"errors"
)

// IsValid checks the configuation of the proxy
func (c Config) IsValid() error {
	if c.ClientID == "" {
		return errors.New("no client id")
	}
	if c.DiscoveryURL == "" && c.SigningCA == "" {
		return errors.New("neither disovery-url or signing ca are not defined")
	}
	if len(c.UserClaims) <= 0 {
		return errors.New("user claims are empty")
	}
	if c.TLSCert != "" && c.TLSKey == "" {
		return errors.New("no tls private key")
	}
	if c.TLSKey != "" && c.TLSCert == "" {
		return errors.New("no tls certificate")
	}
	if c.UpstreamURL == "" {
		return errors.New("no upstream url")
	}

	return nil
}

// HasTLS checks if we have tls
func (c Config) HasTLS() bool {
	return c.TLSCert != "" && c.TLSKey != ""
}
