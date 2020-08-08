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
		ws.PUT("/authorize/{user}").To(l.loginHandler).
			Filter(filters.NewRateLimiter(filters.RateConfig{Period: 60 * time.Second, Limit: 5})).
			Doc("Used login and authorize an account in kore").
			Operation("AuthorizeUser").
			Param(ws.PathParameter("user", "The username you are trying to authorize")).
			Param(ws.QueryParameter("scopes", "A list of requested scopes being request").Required(false)).
			Returns(http.StatusOK, "Contains the access token on successfully authentication", types.IssuedToken{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	return ws, nil
}

// loginHandler is responsible for issuing a local token for local users
func (l *loginHandler) loginHandler(req *restful.Request, resp *restful.Response) {
	handleErrors(req, resp, func() error {
		// @question: wonder if we can start to use scope somehow?
		issued, err := l.Users().Identities().IssueIDToken(req.Request.Context(), "kore")
		if err != nil {
			return err
		}

		return resp.WriteHeaderAndEntity(http.StatusOK, &types.IssuedToken{Token: issued})
	})
}

// EnableAdminsOnly indicates if we need to be an admin user
func (l *loginHandler) EnableAdminsOnly() bool {
	return false
}

// Name returns the name of the handler
func (l *loginHandler) Name() string {
	return "login"
}
