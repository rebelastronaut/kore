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

// NewDeleteAWSOrganizationParams creates a new DeleteAWSOrganizationParams object
// with the default values initialized.
func NewDeleteAWSOrganizationParams() *DeleteAWSOrganizationParams {
	var ()
	return &DeleteAWSOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAWSOrganizationParamsWithTimeout creates a new DeleteAWSOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAWSOrganizationParamsWithTimeout(timeout time.Duration) *DeleteAWSOrganizationParams {
	var ()
	return &DeleteAWSOrganizationParams{

		timeout: timeout,
	}
}

// NewDeleteAWSOrganizationParamsWithContext creates a new DeleteAWSOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAWSOrganizationParamsWithContext(ctx context.Context) *DeleteAWSOrganizationParams {
	var ()
	return &DeleteAWSOrganizationParams{

		Context: ctx,
	}
}

// NewDeleteAWSOrganizationParamsWithHTTPClient creates a new DeleteAWSOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAWSOrganizationParamsWithHTTPClient(client *http.Client) *DeleteAWSOrganizationParams {
	var ()
	return &DeleteAWSOrganizationParams{
		HTTPClient: client,
	}
}

/*DeleteAWSOrganizationParams contains all the parameters to send to the API endpoint
for the delete a w s organization operation typically these are written to a http.Request
*/
type DeleteAWSOrganizationParams struct {

	/*Name
	  Is name the of the resource you are acting on

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

// WithTimeout adds the timeout to the delete a w s organization params
func (o *DeleteAWSOrganizationParams) WithTimeout(timeout time.Duration) *DeleteAWSOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete a w s organization params
func (o *DeleteAWSOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete a w s organization params
func (o *DeleteAWSOrganizationParams) WithContext(ctx context.Context) *DeleteAWSOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete a w s organization params
func (o *DeleteAWSOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete a w s organization params
func (o *DeleteAWSOrganizationParams) WithHTTPClient(client *http.Client) *DeleteAWSOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete a w s organization params
func (o *DeleteAWSOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete a w s organization params
func (o *DeleteAWSOrganizationParams) WithName(name string) *DeleteAWSOrganizationParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete a w s organization params
func (o *DeleteAWSOrganizationParams) SetName(name string) {
	o.Name = name
}

// WithTeam adds the team to the delete a w s organization params
func (o *DeleteAWSOrganizationParams) WithTeam(team string) *DeleteAWSOrganizationParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the delete a w s organization params
func (o *DeleteAWSOrganizationParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAWSOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
