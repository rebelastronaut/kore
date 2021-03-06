// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewRemoveInviteParams creates a new RemoveInviteParams object
// with the default values initialized.
func NewRemoveInviteParams() *RemoveInviteParams {
	var ()
	return &RemoveInviteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveInviteParamsWithTimeout creates a new RemoveInviteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemoveInviteParamsWithTimeout(timeout time.Duration) *RemoveInviteParams {
	var ()
	return &RemoveInviteParams{

		timeout: timeout,
	}
}

// NewRemoveInviteParamsWithContext creates a new RemoveInviteParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemoveInviteParamsWithContext(ctx context.Context) *RemoveInviteParams {
	var ()
	return &RemoveInviteParams{

		Context: ctx,
	}
}

// NewRemoveInviteParamsWithHTTPClient creates a new RemoveInviteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemoveInviteParamsWithHTTPClient(client *http.Client) *RemoveInviteParams {
	var ()
	return &RemoveInviteParams{
		HTTPClient: client,
	}
}

/*RemoveInviteParams contains all the parameters to send to the API endpoint
for the remove invite operation typically these are written to a http.Request
*/
type RemoveInviteParams struct {

	/*Team
	  The name of the team you are deleting the invitation

	*/
	Team string
	/*User
	  The username of the user whos invitation you are removing

	*/
	User string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remove invite params
func (o *RemoveInviteParams) WithTimeout(timeout time.Duration) *RemoveInviteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove invite params
func (o *RemoveInviteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove invite params
func (o *RemoveInviteParams) WithContext(ctx context.Context) *RemoveInviteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove invite params
func (o *RemoveInviteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove invite params
func (o *RemoveInviteParams) WithHTTPClient(client *http.Client) *RemoveInviteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove invite params
func (o *RemoveInviteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTeam adds the team to the remove invite params
func (o *RemoveInviteParams) WithTeam(team string) *RemoveInviteParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the remove invite params
func (o *RemoveInviteParams) SetTeam(team string) {
	o.Team = team
}

// WithUser adds the user to the remove invite params
func (o *RemoveInviteParams) WithUser(user string) *RemoveInviteParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the remove invite params
func (o *RemoveInviteParams) SetUser(user string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveInviteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param team
	if err := r.SetPathParam("team", o.Team); err != nil {
		return err
	}

	// path param user
	if err := r.SetPathParam("user", o.User); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
