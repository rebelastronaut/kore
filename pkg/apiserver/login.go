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

package apiserver

import (
	"net/http"
	"time"

	"github.com/appvia/kore/pkg/apiserver/filters"
	"github.com/appvia/kore/pkg/apiserver/types"
	"github.com/appvia/kore/pkg/kore"
	"github.com/appvia/kore/pkg/utils"

	restful "github.com/emicklei/go-restful"
	log "github.com/sirupsen/logrus"
)

func init() {
	RegisterHandler(&loginHandler{})
}

type loginHandler struct {
	kore.Interface
	// default handler
	DefaultHandler
}

// Path returns the handler path
func (l *loginHandler) Path() string {
	return "login"
}

// Register is responsible for handling the registration
func (l *loginHandler) Register(i kore.Interface, builder utils.PathBuilder) (*restful.WebService, error) {
	path := builder.Add(l.Path())
	log.WithFields(log.Fields{
		"path": path.Base(),
	}).Info("registering the login webservice with container")

	l.Interface = i

	ws := &restful.WebService{}
	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)
	ws.Path(path.Base())

	ws.Route(
		withAllNonValidationErrors(ws.GET("/methods")).To(l.getLoginMethods).
			Doc("Retrieve the supported authentication methods").
			Operation("GetLoginMethods").
			Returns(http.StatusOK, "Details of which login providers are configured", []string{}),
	)

	ws.Route(
		withAllErrors(ws.POST("")).To(l.login).
			Filter(filters.NewRateLimiter(filters.RateConfig{Period: 60 * time.Second, Limit: 5})).
			Doc("Retrieve a refresh token using the specified local credentials").
			Operation("Login").
			Reads(types.LocalUser{}).
			Returns(http.StatusOK, "An access token and a refresh token to access Kore", types.IssuedToken{}),
	)

	ws.Route(
		withAllNonValidationErrors(ws.GET("/token")).To(l.getToken).
			// higher rate limit as multiple in-flight API requests at the same time may all cause a token refresh,
			// lower risk as only accepts a token, so no possibility of credential stuffing. For UI token login,
			// all requests will come from the same server, so this is in effect a global limit not a per client
			// limit. @TODO: Consider changing or possibly removing these limits for this endpoint.
			Filter(filters.NewRateLimiter(filters.RateConfig{Period: 30 * time.Second, Limit: 30})).
			Doc("Retrieve a new token for the user identified by the specified refresh token").
			Param(ws.QueryParameter("refresh", "The refresh token to exchange")).
			Operation("GetToken").
			Returns(http.StatusOK, "An access token which can be used for accessing Kore", types.IssuedToken{}),
	)

	return ws, nil
}

func (l *loginHandler) getLoginMethods(req *restful.Request, resp *restful.Response) {
	handleErrors(req, resp, func() error {
		return resp.WriteHeaderAndEntity(http.StatusOK, l.Config().Authenticators)
	})
}

func (l *loginHandler) login(req *restful.Request, resp *restful.Response) {
	handleErrors(req, resp, func() error {
		localUser := &types.LocalUser{}
		if err := req.ReadEntity(localUser); err != nil {
			return err
		}
		valid, refreshToken := l.Users().Identities().IssueRefreshToken(req.Request.Context(), localUser.Username, localUser.Password)
		if !valid {
			resp.WriteHeader(http.StatusUnauthorized)
			return nil
		}
		valid, token := l.Users().Identities().ExchangeRefreshToken(req.Request.Context(), refreshToken)
		if !valid {
			resp.WriteHeader(http.StatusUnauthorized)
			return nil
		}

		return resp.WriteHeaderAndEntity(http.StatusOK, &types.IssuedToken{
			RefreshToken: string(refreshToken),
			Token:        string(token),
		})
	})
}

func (l *loginHandler) getToken(req *restful.Request, resp *restful.Response) {
	handleErrors(req, resp, func() error {
		refreshToken := req.QueryParameter("refresh")
		// @question: wonder if we can start to use scope somehow?
		valid, issued := l.Users().Identities().ExchangeRefreshToken(req.Request.Context(), []byte(refreshToken))
		if !valid {
			resp.WriteHeader(http.StatusUnauthorized)
			return nil
		}

		return resp.WriteHeaderAndEntity(http.StatusOK, &types.IssuedToken{Token: string(issued)})
	})
}

// EnableAuthentication returns false for the login controller - it does authentication, but its operations are anonymous
func (l *loginHandler) EnableAuthentication() bool {
	return false
}

// EnableAdminsOnly indicates if we need to be an admin user
func (l *loginHandler) EnableAdminsOnly() bool {
	return false
}

// Name returns the name of the handler
func (l *loginHandler) Name() string {
	return "login"
}
