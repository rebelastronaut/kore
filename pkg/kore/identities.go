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

	orgv1 "github.com/appvia/kore/pkg/apis/org/v1"
	"github.com/appvia/kore/pkg/kore/authentication"
	"github.com/appvia/kore/pkg/persistence"
	"github.com/appvia/kore/pkg/persistence/model"
	"github.com/appvia/kore/pkg/utils"
	"github.com/appvia/kore/pkg/utils/validation"
	"golang.org/x/crypto/bcrypt"

	log "github.com/sirupsen/logrus"
)

var (
	// AccountLocal is a local basic auth account
	AccountLocal = "basicauth"
	// AccountToken is a api token account
	AccountToken = "token"
	// AccountSSO is a openid account
	AccountSSO = "sso"
	// SupportedAccounts is a list of supported accounts
	SupportedAccounts = []string{
		AccountLocal,
		AccountToken,
		AccountSSO,
	}
	// RefreshTokenAudience is the audience for kore refresh tokens
	RefreshTokenAudience = "kore-refresh"
	// KoreTokenAudience is the audience for kore identity tokens
	KoreTokenAudience = "kore"
)

// IdentitiesListOptions are search options for listing
type IdentitiesListOptions struct {
	// IdentityTypes is a collection of type to search for
	IdentityTypes []string
	// User is a specific user to user to look for
	User string
}

// Identities is the contract to interact with user identities
type Identities interface {
	// AssociateIDPUser is used to associate an internal user to an idp user
	AssociateIDPUser(ctx context.Context, update *orgv1.UpdateIDPIdentity) error
	// Delete is called to delete an associated identity of a user
	Delete(ctx context.Context, user string, identity string) error
	// IssueRefreshToken creates a new refresh token for the specified username and password,
	// returning true if user valid, false if not, and the created token
	IssueRefreshToken(ctx context.Context, username string, password string) (bool, []byte)
	// ExchangeRefreshToken is used to issue a token for the specified refresh token, returning true
	// and the token if the supplied token is valid and the user is enabled, false otherwise
	ExchangeRefreshToken(ctx context.Context, refreshToken []byte) (bool, []byte)
	// IssueIDToken is used to issue a token for the current user context in kore
	IssueIDToken(ctx context.Context, audience string) ([]byte, error)
	// List returns a list of all the identities managed in kore
	List(ctx context.Context, options IdentitiesListOptions) (*orgv1.IdentityList, error)
	// UpdateUserBasicAuth is used to update a basic auth profile in kore
	UpdateUserBasicAuth(ctx context.Context, update *orgv1.UpdateBasicAuthIdentity) error
}

type idImpl struct {
	*hubImpl
}

// AssociateIDPUser is used to associate an internal user to an idp user
func (h *idImpl) AssociateIDPUser(ctx context.Context, update *orgv1.UpdateIDPIdentity) error {
	return nil
}

func (h *idImpl) IssueRefreshToken(ctx context.Context, username string, password string) (bool, []byte) {
	u, err := h.Persist().Users().Get(ctx, username)
	if err != nil {
		log.WithField("user", username).WithError(err).Error("trying to retrieve user to issue refresh token")
		return false, nil
	}
	if u.Disabled {
		log.WithField("user", username).Debug("refresh token not issued as user is disabled")
		return false, nil
	}

	identity, err := h.Persist().Identities().Get(ctx,
		persistence.Filter.WithUser(username),
		persistence.Filter.WithProvider(AccountLocal),
	)
	if err != nil {
		log.WithField("user", username).WithError(err).Error("trying to retrieve the identity to issue refresh token")
		return false, nil
	}
	if identity.ProviderToken == "" {
		log.WithField("user", username).Debug("refresh token not issued as user identity has no password (provider token) set")
		return false, nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(identity.ProviderToken), []byte(password)); err != nil {
		log.WithField("user", username).Debug("refresh token not issued as user provided incorrect password")
		return false, nil
	}

	minted, err := h.Tokens().Issue(ctx, IssueOptions{
		Audience: RefreshTokenAudience,
		Scopes:   []string{},
		Duration: 24 * time.Hour * 28,
		Email:    u.Email,
		Username: u.Username,
	})
	if err != nil {
		log.WithField("user", username).WithError(err).Error("trying to mint refresh token")

		return false, nil
	}

	log.WithFields(log.Fields{
		"email":    u.Email,
		"username": u.Username,
	}).Info("created refresh token for user")

	return true, minted
}

func (h *idImpl) ExchangeRefreshToken(ctx context.Context, refreshToken []byte) (bool, []byte) {
	// First, check refresh token is signed by us, and has the right audience
	valid, claims := h.Tokens().Validate(ctx, refreshToken, RefreshTokenAudience)
	if !valid {
		log.Debug("invalid refresh token presented, rejecting exchange")
		return false, nil
	}

	// check the user is still valid in kore
	username, found := claims.GetUserClaim(Userclaim)
	if !found {
		log.Debug("no username claim on refresh token, rejecting exchange")
		return false, nil
	}

	identity, found, err := h.GetUserIdentity(ctx, username, WithAuthMethod("jwt"))
	if err != nil {
		log.WithField("user", username).WithError(err).Error("trying to retrieve user identity, rejecting exchange")
		return false, nil
	}
	if !found {
		log.WithField("user", username).Debug("user not found in kore (may have been deleted since refresh token was issued), rejecting exchange")
		return false, nil
	}
	if identity.Disabled() {
		log.WithField("user", username).Debug("user disabled, rejecting exchange")
		return false, nil
	}

	minted, err := h.Tokens().Issue(ctx, IssueOptions{
		Audience: KoreTokenAudience,
		Scopes:   []string{},
		Duration: 30 * time.Minute,
		Email:    identity.Email(),
		Username: identity.Username(),
	})
	if err != nil {
		log.WithField("user", username).WithError(err).Error("trying to mint id token")
		return false, nil
	}
	log.WithFields(log.Fields{
		"email":    identity.Email(),
		"username": identity.Username(),
	}).Debug("created access token for user")
	return true, minted
}

// IssueIDToken is used to issue a token for a identity in kore
func (h *idImpl) IssueIDToken(ctx context.Context, audience string) ([]byte, error) {
	user := authentication.MustGetIdentity(ctx)

	minted, err := h.Tokens().Issue(ctx, IssueOptions{
		Audience: audience,
		Scopes:   []string{},
		Duration: 24 * time.Hour,
		Email:    user.Email(),
		Username: user.Username(),
	})
	if err != nil {
		log.WithField("user", user.Username()).WithError(err).Error("trying to mint id token")

		return nil, err
	}
	log.WithFields(log.Fields{
		"email":    user.Email(),
		"username": user.Username(),
	}).Debug("created access token for user")

	return minted, nil
}

// Delete is called to delete an associated identity of a user
func (h *idImpl) Delete(ctx context.Context, username string, identity string) error {
	user := authentication.MustGetIdentity(ctx)

	// @step: you must be the user or an admin to perform this
	if !user.IsGlobalAdmin() && user.Username() != username {
		return NewErrNotAllowed("must be administrator or the user to delete credential")
	}

	// @step: check the identity type and username
	if !utils.Contains(identity, SupportedAccounts) {
		return validation.NewError("invalid identity").
			WithFieldError("identity", validation.InvalidValue, "identity type does not exist")
	}
	if !UsernameRegex.MatchString(username) {
		return validation.NewError("invalid username").
			WithFieldError("username", validation.InvalidValue, "username is invalid")
	}

	// @step: cannot delete the admin user
	if username == HubAdminUser && identity == AccountLocal {
		return NewErrNotAllowed("cannot delete the admin user identity")
	}

	// @step: check the user exists
	_, err := h.Persist().Users().Get(ctx, username)
	if err != nil {
		if persistence.IsNotFound(err) {
			return ErrNotFound
		}
		log.WithError(err).Error("trying to check if user exists")

		return err
	}

	// @step: retrieve the identity if any
	ident, err := h.Persist().Identities().Get(ctx,
		persistence.Filter.WithUser(username),
		persistence.Filter.WithProvider(identity),
	)
	if err != nil {
		if persistence.IsNotFound(err) {
			return NewErrNotAllowed("user does not have this identity type")
		}
	}

	return h.Persist().Identities().Delete(ctx, ident)
}

// List returns a list of all the identities managed in kore
func (h *idImpl) List(ctx context.Context, options IdentitiesListOptions) (*orgv1.IdentityList, error) {
	user := authentication.MustGetIdentity(ctx)
	var filters []persistence.ListFunc

	// @step: validate inputs
	for _, x := range options.IdentityTypes {
		if !utils.Contains(x, SupportedAccounts) {
			return nil, validation.NewError("invalid identity type").
				WithFieldError("type", validation.InvalidValue, "must be a valid identity type")
		}
		filters = append(filters, persistence.Filter.WithProvider(x))
	}

	if options.User != "" {
		if !UsernameRegex.MatchString(options.User) {
			return nil, validation.NewError("invalid username").
				WithFieldError("username", validation.InvalidValue, "username is invalid")
		}
		filters = append(filters, persistence.Filter.WithUser(options.User))
	}

	if options.User == "" && !user.IsGlobalAdmin() {
		return nil, ErrUnauthorized
	}

	list, err := h.Persist().Identities().List(ctx, filters...)
	if err != nil {
		return nil, err
	}

	return DefaultConvertor.FromIdentityModelList(list), err
}

// UpdateUserBasicAuth is used to update a basic auth profile in kore
func (h *idImpl) UpdateUserBasicAuth(ctx context.Context, update *orgv1.UpdateBasicAuthIdentity) error {
	user := authentication.MustGetIdentity(ctx)

	if !user.IsGlobalAdmin() && user.Username() != update.Username {
		return NewErrNotAllowed("must be administrator or the user to update credential")
	}

	logger := log.WithFields(log.Fields{
		"username": update.Username,
	})

	// @step: check the user exists
	u, err := h.Persist().Users().Get(ctx, update.Username)
	if err != nil {
		if persistence.IsNotFound(err) {
			return ErrNotFound
		}
		logger.WithError(err).Error("trying to check if user exists")

		return err
	}

	// @step: if the user has zero identites only the admin can it up
	identity, err := h.Persist().Identities().Get(ctx,
		persistence.Filter.WithUser(update.Username),
		persistence.Filter.WithProvider(AccountLocal),
	)
	if err != nil {
		if !persistence.IsNotFound(err) {
			logger.WithError(err).Error("trying to retrieve the identity")

			return err
		}
		logger.Info("setting up basicauth identity for user")
	}

	if identity == nil {
		identity = &model.Identity{
			Provider: AccountLocal,
			UserID:   u.ID,
		}
	}

	// @step: encrypt the token
	hashed, err := bcrypt.GenerateFromPassword([]byte(update.Password), bcrypt.DefaultCost)
	if err != nil {
		log.WithError(err).Error("trying to hash the password")

		return err
	}
	identity.ProviderToken = string(hashed)

	// @step: update the credentials
	if err := h.Persist().Identities().Update(ctx, identity); err != nil {
		logger.WithError(err).Error("trying to update the credential")

		return err
	}
	logger.Info("updated the basicauth credential for user")

	return nil
}
