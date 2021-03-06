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
	"crypto/tls"
	"encoding/pem"
	"io"
	"net/url"
	"strings"

	"github.com/appvia/kore/pkg/kore"
	"github.com/appvia/kore/pkg/version"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// New creates a configuration
func New(reader io.Reader) (*Config, error) {
	config := &Config{}

	if err := yaml.NewDecoder(reader).Decode(config); err != nil {
		return nil, err
	}

	if config.FeatureGates == nil {
		config.FeatureGates = map[string]bool{}
	}
	for fg, enabled := range kore.DefaultFeatureGates() {
		if _, isSet := config.FeatureGates[fg]; !isSet {
			config.FeatureGates[fg] = enabled
		}
	}

	return config, nil
}

// NewEmpty returns an empty configuration
func NewEmpty() *Config {
	return &Config{
		AuthInfos:    make(map[string]*AuthInfo),
		Profiles:     make(map[string]*Profile),
		Servers:      make(map[string]*Server),
		Version:      version.Release,
		FeatureGates: kore.DefaultFeatureGates(),
	}
}

// IsValid checks if the configuration is valid
func (c *Config) IsValid() error {
	return nil
}

// CreateProfile is used to create a profile
func (c *Config) CreateProfile(name, endpoint string) {
	var ca []byte

	u, _ := url.Parse(endpoint)
	if u != nil && u.Scheme == "https" && u.Hostname() == "localhost" {
		ca = c.getUntrustedCA(strings.TrimPrefix(endpoint, "https://"))
	}

	c.AddProfile(name, &Profile{
		Server:   name,
		AuthInfo: name,
	})
	c.AddServer(name, &Server{Endpoint: endpoint, CACertificate: string(ca)})
}

func (c *Config) getUntrustedCA(url string) (ca []byte) {
	conn, err := tls.Dial("tcp", url, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		log.Debugf("failed to connect to %s: %s", url, err.Error())
		return nil
	}

	if err := conn.Handshake(); err != nil {
		log.Debugf("SSL handshake failed with %s: %s", url, err.Error())
		return nil
	}

	l := len(conn.ConnectionState().PeerCertificates)
	caCert := conn.ConnectionState().PeerCertificates[0]
	cert := conn.ConnectionState().PeerCertificates[l-1]
	for _, domain := range cert.DNSNames {
		if domain == "localhost" {
			block := &pem.Block{
				Type:  "CERTIFICATE",
				Bytes: caCert.Raw,
			}
			return pem.EncodeToMemory(block)
		}
	}

	return nil
}

// ListProfiles returns a list of profile names
func (c *Config) ListProfiles() []string {
	if c.Profiles == nil {
		return nil
	}
	var list []string

	for k := range c.Profiles {
		list = append(list, k)
	}

	return list
}

// GetProfile returns the profile
func (c *Config) GetProfile(name string) *Profile {
	if !c.HasProfile(name) {
		return &Profile{}
	}

	return c.Profiles[name]
}

// GetProfileAuthMethod returns the method of authentication for a profile
func (c *Config) GetProfileAuthMethod(name string) string {
	if !c.HasProfile(name) {
		return ""
	}
	if !c.HasAuthInfo(c.Profiles[name].AuthInfo) {
		return ""
	}
	auth := c.AuthInfos[c.Profiles[name].AuthInfo]
	switch {
	case auth.BasicAuth != nil:
		return "basicauth"
	case auth.OIDC != nil:
		return "sso"
	case auth.Token != nil:
		return "token"
	case auth.KoreIdentity != nil:
		return "idtoken"
	}

	return "none"
}

// GetServer returns the endpoint for the profile
func (c *Config) GetServer(name string) *Server {
	if !c.HasProfile(name) {
		return &Server{}
	}

	return c.Servers[c.Profiles[name].Server]
}

// GetAuthInfo returns the auth for a profile
func (c *Config) GetAuthInfo(name string) *AuthInfo {
	ct := c.Profiles[name]
	if ct == nil {
		return &AuthInfo{}
	}

	a := c.AuthInfos[ct.AuthInfo]

	if a == nil {
		return &AuthInfo{}
	}

	return a
}

// HasAuth checks if we have auth enabled
func (c *Config) HasAuth(name string) bool {
	a := c.GetAuthInfo(name)
	if a.OIDC != nil || a.BasicAuth != nil || a.Token != nil {
		return true
	}

	return false
}

// AddProfile adds a profile to the config
func (c *Config) AddProfile(name string, ctx *Profile) {
	if c.Profiles == nil {
		c.Profiles = make(map[string]*Profile)
	}
	c.Profiles[name] = ctx
}

// AddServer adds a server
func (c *Config) AddServer(name string, server *Server) {
	if c.Servers == nil {
		c.Servers = make(map[string]*Server)
	}
	c.Servers[name] = server
}

// AddAuthInfo adds a authentication
func (c *Config) AddAuthInfo(name string, auth *AuthInfo) {
	if c.AuthInfos == nil {
		c.AuthInfos = make(map[string]*AuthInfo)
	}
	c.AuthInfos[name] = auth
}

// HasValidProfile checks we have a current context
func (c *Config) HasValidProfile(name string) error {
	if name == "" {
		return ErrNoProfileSelected
	}
	if !c.HasServer(c.Profiles[name].Server) {
		return ErrNoProfileEndpoint
	}

	return nil
}

// HasProfile checks if the context exists in the config
func (c *Config) HasProfile(name string) bool {
	_, found := c.Profiles[name]

	return found
}

// HasServer checks if the context exists in the config
func (c *Config) HasServer(name string) bool {
	_, found := c.Servers[name]

	return found
}

// HasAuthInfo checks if the context exists in the config
func (c *Config) HasAuthInfo(name string) bool {
	_, found := c.AuthInfos[name]

	return found
}

// IsOIDCProviderConfigured checks if an OIDC provider is available
func (c *Config) IsOIDCProviderConfigured(name string) bool {
	info, found := c.AuthInfos[name]
	if !found {
		return false
	}

	return len(info.OIDC.ClientID) > 0 &&
		len(info.OIDC.ClientSecret) > 0 &&
		len(info.OIDC.AuthorizeURL) > 0
}

// RemoveServer removes a server instance
func (c *Config) RemoveServer(name string) {
	delete(c.Servers, name)
}

// RemoveUserInfo removes the user info
func (c *Config) RemoveUserInfo(name string) {
	delete(c.AuthInfos, name)
}

// RemoveProfile removes the profile
func (c *Config) RemoveProfile(name string) {
	p, found := c.Profiles[name]
	if !found {
		return
	}
	c.RemoveServer(p.Server)
	c.RemoveUserInfo(p.AuthInfo)

	delete(c.Profiles, name)
}

// Update writes the config to the file
func (c *Config) Update(w io.Writer) error {
	return yaml.NewEncoder(w).Encode(c)
}
