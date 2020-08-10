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

package kore

import (
	"context"
	"time"

	"github.com/appvia/kore/pkg/utils"
	jwt "github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

const (
	// Userclaim is the default claim we use for username
	Userclaim = "preferred_username"
)

// Tokens provides an interface for issuing kore managed tokens
type Tokens interface {
	// Issue is used to issue a token from kore
	Issue(ctx context.Context, options IssueOptions) ([]byte, error)
	// Validate is used to verify a token has been signed by kore, returning true
	// plus the claims if valid, false otherwise
	Validate(ctx context.Context, token []byte, audience string) (bool, *utils.Claims)
}

// IssueOptions are options for creating tokens
type IssueOptions struct {
	// Audience is a collection of audiences
	Audience string
	// Scopes is a collection of scope to be included
	Scopes []string
	// Groups is a collection of groups to include
	Groups []string
	// Duration is the time to live for the tokne
	Duration time.Duration
	// Username is the username to include
	Username string
	// Email is the email address to include
	Email string
}

type tokenImpl struct {
	CertificateIface
	ConfigIface
}

// Issue is used to issue a token from kore
func (t *tokenImpl) Issue(ctx context.Context, options IssueOptions) ([]byte, error) {
	claims := utils.NewClaims(jwt.MapClaims{
		"aud":     options.Audience,
		"email":   options.Email,
		"exp":     float64(time.Now().UTC().Add(options.Duration).Unix()),
		"iss":     t.Config().PublicAPIURL,
		"nbf":     time.Now().UTC().Add(-60 * time.Second).Unix(),
		"scopes":  options.Scopes,
		Userclaim: options.Username,
	})

	minted, err := claims.Sign(t.CertificateAuthorityKey())
	if err != nil {
		log.WithField("username", options.Username).WithError(err).Error("trying to mint token")

		return nil, err
	}

	return minted, nil
}

func (t *tokenImpl) Validate(ctx context.Context, token []byte, audience string) (bool, *utils.Claims) {
	c := make(jwt.MapClaims)

	// @step: parse and extract the identity
	parsedToken, err := jwt.ParseWithClaims(string(token), &c, func(token *jwt.Token) (interface{}, error) {
		return jwt.ParseRSAPublicKeyFromPEM(t.CertificateAuthority())
	})
	if err != nil {
		log.WithError(err).Warn("error attempting to validate token")
		return false, nil
	}
	if !parsedToken.Valid {
		log.Debug("invalid token presented (e.g. signature failure)")
		return false, nil
	}

	claims := utils.NewClaims(c)

	// @step: check the audience
	aud, found := claims.GetAudience()
	if !found {
		log.Debug("no audience in the presented token")

		return false, nil
	}
	if aud != audience {
		log.Debug("invalid audience presented in the token")

		return false, nil
	}
	return true, claims
}
