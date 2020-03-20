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

// NewGetTeamMembersParams creates a new GetTeamMembersParams object
// with the default values initialized.
func NewGetTeamMembersParams() *GetTeamMembersParams {
	var ()
	return &GetTeamMembersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTeamMembersParamsWithTimeout creates a new GetTeamMembersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTeamMembersParamsWithTimeout(timeout time.Duration) *GetTeamMembersParams {
	var ()
	return &GetTeamMembersParams{

		timeout: timeout,
	}
}

// NewGetTeamMembersParamsWithContext creates a new GetTeamMembersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTeamMembersParamsWithContext(ctx context.Context) *GetTeamMembersParams {
	var ()
	return &GetTeamMembersParams{

		Context: ctx,
	}
}

// NewGetTeamMembersParamsWithHTTPClient creates a new GetTeamMembersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTeamMembersParamsWithHTTPClient(client *http.Client) *GetTeamMembersParams {
	var ()
	return &GetTeamMembersParams{
		HTTPClient: client,
	}
}

/*GetTeamMembersParams contains all the parameters to send to the API endpoint
for the get team members operation typically these are written to a http.Request
*/
type GetTeamMembersParams struct {

	/*Team
	  Is the name of the team you are acting within

	*/
	Team string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get team members params
func (o *GetTeamMembersParams) WithTimeout(timeout time.Duration) *GetTeamMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get team members params
func (o *GetTeamMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get team members params
func (o *GetTeamMembersParams) WithContext(ctx context.Context) *GetTeamMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get team members params
func (o *GetTeamMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get team members params
func (o *GetTeamMembersParams) WithHTTPClient(client *http.Client) *GetTeamMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get team members params
func (o *GetTeamMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTeam adds the team to the get team members params
func (o *GetTeamMembersParams) WithTeam(team string) *GetTeamMembersParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the get team members params
func (o *GetTeamMembersParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *GetTeamMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param team
	if err := r.SetPathParam("team", o.Team); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
