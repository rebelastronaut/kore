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

package login

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/appvia/kore/pkg/apiserver"
	"github.com/appvia/kore/pkg/apiserver/types"
	"github.com/appvia/kore/pkg/client"
	"github.com/appvia/kore/pkg/client/config"
	restconfig "github.com/appvia/kore/pkg/client/config"
	cmdutil "github.com/appvia/kore/pkg/cmd/utils"
	"github.com/appvia/kore/pkg/kore"
	"github.com/appvia/kore/pkg/utils"
	"github.com/manifoldco/promptui"

	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

var (
	loginLongDescription = `
Is used to authenticate yourself to the currently selected profile. Login
performs an oauth2 authentication flow and retrieve your identity token.
`

	loginExamples = `
$ kore login    # will login and update the current profile
$ kore login local -a http://127.0.0.1:8080  # create a profile and login
`
)

const (
	// DefaultKoreURL is the default value for the kore api
	DefaultKoreURL = "http://localhost:10080"
)

// LoginOptions are the options for logging in
type LoginOptions struct {
	cmdutil.Factory
	cmdutil.DefaultHandler
	// Name is used when creating a profile
	Name string
	// Endpoint is an optional endpoint
	Endpoint string
	// Force is used to force an operation
	Force bool
	// LocalUser indicates we are using local users
	LocalUser bool
	// Port is the local port to use for http server
	Port int
}

// NewCmdLogin creates and returns a login command
func NewCmdLogin(factory cmdutil.Factory) *cobra.Command {
	o := &LoginOptions{Factory: factory}

	command := &cobra.Command{
		Use:     "login",
		Short:   "Authenticate yourself and retrieve a token for Kore",
		Long:    loginLongDescription,
		Example: loginExamples,
		Run:     cmdutil.DefaultRunFunc(o),
	}

	flags := command.Flags()
	flags.StringVarP(&o.Endpoint, "api-url", "a", "", "specify the kore api server to login `URL`")
	flags.BoolVarP(&o.Force, "force", "f", false, "must be set when you want to override the api-server on an existing profile `BOOL`")
	flags.IntVarP(&o.Port, "port", "p", 3001, "sets the local port used for redirection when authenticating `PORT`")
	flags.BoolVarP(&o.LocalUser, "local", "l", false, "indicate we are using a local login `BOOL`")

	return command
}

// Validate is used to validate the parameters
func (o *LoginOptions) Validate() error {
	current := o.Client().CurrentProfile()

	if o.Name != "" {
		current = o.Name
		o.Client().OverrideProfile(current)
	}

	if o.Name != "" && o.Config().HasProfile(current) && !o.Force {
		return fmt.Errorf("profile name already used (note: you can use the --force option to force the update)")
	}

	if o.Client().CurrentProfile() == "" {
		return fmt.Errorf("please specify a name for the profile")
	}

	return nil
}

// Run performs the command action to login
func (o *LoginOptions) Run() error {
	var err error

	current := o.Client().CurrentProfile()

	if !o.Config().HasProfile(current) || !o.Config().HasServer(current) {
		// @step: set the default profile
		o.Config().CurrentProfile = current

		if o.Endpoint == "" {
			o.Endpoint = DefaultKoreURL

			if err := (cmdutil.Prompts{
				&cmdutil.Prompt{
					Id:     "Please enter the Kore API (e.g https://api.example.com)",
					Value:  &o.Endpoint,
					ErrMsg: "invalid endpoint",
					Validate: func(v string) error {
						if !utils.IsValidURL(v) {
							return errors.New("invalid endpoint url")
						}
						return nil
					},
				},
			}).Collect(); err != nil {
				return err
			}
		}
	}

	// @step: do we even have a profile?
	if !o.Config().HasProfile(current) {
		o.Config().CreateProfile(current, o.Endpoint)
	}
	if !o.Config().HasAuthInfo(current) {
		o.Config().AddAuthInfo(current, &config.AuthInfo{})
	}

	method := o.Config().GetProfileAuthMethod(current)
	if method == "none" {
		// we need to ask if this is sso or basicauth auth
		if o.LocalUser {
			return o.RunIDAuth()
		}
		_, method, err = (&promptui.Select{
			Label:        "Which method are you using to login?",
			Items:        []string{"sso", "idtoken"},
			HideHelp:     true,
			HideSelected: true,
		}).Run()
		if err != nil {
			return err
		}
	} else {
		method = o.Config().GetProfileAuthMethod(current)
	}

	// @step: else we do have a profile so need see if basicauth or sso
	switch method {
	case "idtoken":
		return o.RunIDAuth()
	case "sso":
		return o.RunOAuth()
	case "token", "basicauth":
		return errors.New(method + " authentication does not require login")
	default:
		return errors.New("unknown authentication method")
	}
}

// RunIDAuth performs a identity flow login using basicauth credentials
func (o *LoginOptions) RunIDAuth() error {
	// @step: check if we have a token
	current := o.Client().CurrentProfile()

	auth := o.Config().GetAuthInfo(current)
	token := utils.StringValue(auth.IdentityToken)

	var username, password string

	if token != "" {
		claims, err := utils.NewJWTTokenFromBytes([]byte(token))
		if err == nil {
			username, _ = claims.GetUserClaim(kore.Userclaim)
		}
	}

	if err := (cmdutil.Prompts{
		&cmdutil.Prompt{
			Id:     "Please enter your username",
			Value:  &username,
			ErrMsg: "invalid username",
		},
		&cmdutil.Prompt{
			Id:     "Please confirm password for " + username,
			Value:  &password,
			Mask:   true,
			ErrMsg: "invalid password",
		},
	}).Collect(); err != nil {
		return err
	}

	issued := &types.IssuedToken{}
	auth.BasicAuth = &config.BasicAuth{
		Username: username,
		Password: password,
	}

	// @step: exchange the credentials for idtoken
	err := o.ClientWithEndpoint("/login/authorize/{user}").
		Parameters(client.PathParameter("user", username)).
		Result(issued).
		Update().
		Error()
	if err != nil {
		if client.IsNotAuthorized(err) {
			return errors.New("authentication denied, please recheck your credentials")
		}

		return err
	}

	// @step: do not save the authentication setting
	// @note if this probably has been be reviewed as it relys on orders
	// https://github.com/appvia/kore/blob/master/pkg/client/client.go#L249-L259
	auth.BasicAuth = nil
	auth.IdentityToken = utils.StringPtr(string(issued.Token))

	return o.UpdateConfig()
}

// RunOAuth performs a traditional oauth login
func (o *LoginOptions) RunOAuth() error {
	var err error

	current := o.Client().CurrentProfile()

	// @check we have the minimum required for authentication
	auth := o.Config().GetAuthInfo(current)
	if auth.OIDC == nil {
		auth.OIDC = &config.OIDC{}
	}

	// @step: we make done channels to signal events
	doneCh := make(chan struct{})
	errCh := make(chan error)

	token := &apiserver.AuthorizationResponse{}

	// @step: we create a local http server to order to handle the callback
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		token, err = handleLoginCallback(req, w)
		if err != nil {
			errCh <- err
			return
		}
		doneCh <- struct{}{}
	})

	// @step: we need to start the http server in the background
	go func() {
		listenAddress := fmt.Sprintf("127.0.0.1:%d", o.Port)
		if err := http.ListenAndServe(listenAddress, nil); err != nil {
			errCh <- fmt.Errorf("trying to start local http server: %s", err)
		}
	}()

	o.Println("Attempting to authenticate to Kore: %s [%s]\n",
		o.Config().GetServer(current).Endpoint,
		o.Client().CurrentProfile(),
	)

	// @step: open a browser to the to the api server
	redirectURL := fmt.Sprintf("%s/oauth/authorize?redirect_url=http://localhost:%d",
		o.Config().GetServer(current).Endpoint, o.Port)

	if err := open.Run(redirectURL); err != nil {
		return fmt.Errorf("trying to open web browser, error: %s", err)
	}

	// @step: we wait for either a done or error or timeout
	select {
	case <-doneCh:
	case err := <-errCh:
		return fmt.Errorf("trying to authorize the client: %s", err)
	case <-time.After(30 * time.Second):
		return errors.New("authorization request timed out waiting to complete")
	}

	auth = o.Config().GetAuthInfo(current)
	auth.OIDC = &restconfig.OIDC{
		AccessToken:  token.AccessToken,
		AuthorizeURL: token.AuthorizationURL,
		ClientID:     token.ClientID,
		ClientSecret: token.ClientSecret,
		IDToken:      token.IDToken,
		RefreshToken: token.RefreshToken,
		TokenURL:     token.TokenEndpointURL,
	}

	// @step: update the local configuration on disk
	if err := o.UpdateConfig(); err != nil {
		return fmt.Errorf("trying to update the client configuration: %s", err)
	}

	o.Println("Successfully authenticated to %s", current)

	return nil
}

// handleLoginCallback is used to handle the callback from the api server
func handleLoginCallback(req *http.Request, resp http.ResponseWriter) (*apiserver.AuthorizationResponse, error) {
	// @step: check we have a token in the return
	if req.URL.RawQuery == "" {
		return nil, errors.New("no token found in the authorization request")
	}
	if !strings.HasPrefix(req.URL.RawQuery, "token=") {
		return nil, errors.New("invalid token response from apiserver")
	}
	raw := strings.TrimPrefix(req.URL.RawQuery, "token=")

	// @step: extract and decode the token
	decoded, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		return nil, err
	}

	g, err := gzip.NewReader(bytes.NewReader(decoded))
	if err != nil {
		return nil, err
	}
	g.Close()

	rb, err := ioutil.ReadAll(g)
	if err != nil {
		return nil, err
	}

	token := &apiserver.AuthorizationResponse{}

	if err := json.NewDecoder(bytes.NewReader(rb)).Decode(token); err != nil {
		return nil, err
	}

	// @step: send back the html
	shutdown := `<html><body><script>window.close();</script></body></html>`
	if _, err := resp.Write([]byte(shutdown)); err != nil {
		return nil, err
	}

	return token, nil
}
