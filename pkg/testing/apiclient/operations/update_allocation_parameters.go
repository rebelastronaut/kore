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

// NewUpdateAllocationParams creates a new UpdateAllocationParams object
// with the default values initialized.
func NewUpdateAllocationParams() *UpdateAllocationParams {
	var ()
	return &UpdateAllocationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAllocationParamsWithTimeout creates a new UpdateAllocationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateAllocationParamsWithTimeout(timeout time.Duration) *UpdateAllocationParams {
	var ()
	return &UpdateAllocationParams{

		timeout: timeout,
	}
}

// NewUpdateAllocationParamsWithContext creates a new UpdateAllocationParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateAllocationParamsWithContext(ctx context.Context) *UpdateAllocationParams {
	var ()
	return &UpdateAllocationParams{

		Context: ctx,
	}
}

// NewUpdateAllocationParamsWithHTTPClient creates a new UpdateAllocationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateAllocationParamsWithHTTPClient(client *http.Client) *UpdateAllocationParams {
	var ()
	return &UpdateAllocationParams{
		HTTPClient: client,
	}
}

/*UpdateAllocationParams contains all the parameters to send to the API endpoint
for the update allocation operation typically these are written to a http.Request
*/
type UpdateAllocationParams struct {

	/*Name
	  Is the name of the allocation you wish to update

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

// WithTimeout adds the timeout to the update allocation params
func (o *UpdateAllocationParams) WithTimeout(timeout time.Duration) *UpdateAllocationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update allocation params
func (o *UpdateAllocationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update allocation params
func (o *UpdateAllocationParams) WithContext(ctx context.Context) *UpdateAllocationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update allocation params
func (o *UpdateAllocationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update allocation params
func (o *UpdateAllocationParams) WithHTTPClient(client *http.Client) *UpdateAllocationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update allocation params
func (o *UpdateAllocationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the update allocation params
func (o *UpdateAllocationParams) WithName(name string) *UpdateAllocationParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update allocation params
func (o *UpdateAllocationParams) SetName(name string) {
	o.Name = name
}

// WithTeam adds the team to the update allocation params
func (o *UpdateAllocationParams) WithTeam(team string) *UpdateAllocationParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the update allocation params
func (o *UpdateAllocationParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAllocationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
