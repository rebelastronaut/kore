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
	"context"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"io/ioutil"

	"github.com/appvia/kore/pkg/apiserver/plugins/identity"
	"github.com/appvia/kore/pkg/kore"
	"github.com/appvia/kore/pkg/kore/authentication"
	"github.com/appvia/kore/pkg/utils"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

const (
	// UserClaim is the jwt claims used to hold the subject
	UserClaim = "preferred_username"
)

type authImpl struct {
	kore.Interface
	// key is the public key used to verify
	key interface{}
	// config is the service config
	config Config
}

// New returns an jwt authenticator
func New(h kore.Interface, config Config) (identity.Plugin, error) {
	// @step: verify the configuration
	if err := config.IsValid(); err != nil {
		return nil, err
	}
	log.Info("initializing the jwt authentication plugin")

	var key interface{}

	if config.PublicKey != "" {
		pubkey, err := base64.StdEncoding.DecodeString(config.PublicKey)
		if err != nil {
			return nil, fmt.Errorf("Unable to parse public key from config: %v", err)
		}

		key, err = x509.ParsePKIXPublicKey([]byte(pubkey))
		if err != nil {
			return nil, err
		}
	}
	if config.PublicKeyPath != "" {
		content, err := ioutil.ReadFile(config.PublicKeyPath)
		if err != nil {
			return nil, err
		}
		key, err = jwt.ParseRSAPublicKeyFromPEM(content)
		if err != nil {
			return nil, err
		}
	}

	return &authImpl{Interface: h, config: config, key: key}, nil
}

// Admit is called to authenticate the inbound request
func (o *authImpl) Admit(ctx context.Context, req identity.Requestor) (authentication.Identity, bool) {
	// @step: verify the authorization token
	bearer, found := utils.GetBearerToken(req.Headers().Get("Authorization"))
	if !found {
		return nil, false
	}

	c := make(jwt.MapClaims)

	// @step: parse and extract the identity
	token, err := jwt.ParseWithClaims(bearer, &c, func(token *jwt.Token) (interface{}, error) {
		return o.key, nil
	})
	if err != nil {
		return nil, false
	}
	if !token.Valid {
		return nil, false
	}

	claims := utils.NewClaims(c)

	// @step: check the audience if required
	if o.config.HasAudience() {
		aud, found := claims.GetAudience()
		if !found {
			log.Warn("no audience in the presented token")

			return nil, false
		}
		if aud != o.config.Audience {
			log.Warn("invalid audience presented in the token")

			return nil, false
		}
	}

	// @note: i think it's fine to hardcode the claim here as where issuing the token anyhow
	username, found := claims.GetUserClaim(UserClaim)
	if !found {
		return nil, false
	}

	identity, found, err := o.GetUserIdentity(ctx, username, kore.WithAuthMethod("jwt"))
	if err != nil {
		log.WithError(err).Error("trying to retrieve user identity")

		return nil, false
	}
	if !found {
		return nil, false
	}

	log.WithField("user", username).Debug("user passed JWT authentication")
	return identity, true
}

// Name returns the plugin name
func (o *authImpl) Name() string {
	return "localjwt"
}
