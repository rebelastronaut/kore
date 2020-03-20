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

// NewUpdateGKEParams creates a new UpdateGKEParams object
// with the default values initialized.
func NewUpdateGKEParams() *UpdateGKEParams {
	var ()
	return &UpdateGKEParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateGKEParamsWithTimeout creates a new UpdateGKEParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateGKEParamsWithTimeout(timeout time.Duration) *UpdateGKEParams {
	var ()
	return &UpdateGKEParams{

		timeout: timeout,
	}
}

// NewUpdateGKEParamsWithContext creates a new UpdateGKEParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateGKEParamsWithContext(ctx context.Context) *UpdateGKEParams {
	var ()
	return &UpdateGKEParams{

		Context: ctx,
	}
}

// NewUpdateGKEParamsWithHTTPClient creates a new UpdateGKEParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateGKEParamsWithHTTPClient(client *http.Client) *UpdateGKEParams {
	var ()
	return &UpdateGKEParams{
		HTTPClient: client,
	}
}

/*UpdateGKEParams contains all the parameters to send to the API endpoint
for the update g k e operation typically these are written to a http.Request
*/
type UpdateGKEParams struct {

	/*Name
	  Is name the of the GKE cluster you are acting upon

	*/
	Name string
	/*Team
	  Is the name of the team you are acting within

	*/
	Team string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update g k e params
func (o *UpdateGKEParams) WithTimeout(timeout time.Duration) *UpdateGKEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update g k e params
func (o *UpdateGKEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update g k e params
func (o *UpdateGKEParams) WithContext(ctx context.Context) *UpdateGKEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update g k e params
func (o *UpdateGKEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update g k e params
func (o *UpdateGKEParams) WithHTTPClient(client *http.Client) *UpdateGKEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update g k e params
func (o *UpdateGKEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the update g k e params
func (o *UpdateGKEParams) WithName(name string) *UpdateGKEParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update g k e params
func (o *UpdateGKEParams) SetName(name string) {
	o.Name = name
}

// WithTeam adds the team to the update g k e params
func (o *UpdateGKEParams) WithTeam(team string) *UpdateGKEParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the update g k e params
func (o *UpdateGKEParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateGKEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param team
	if err := r.SetPathParam("team", o.Team); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
