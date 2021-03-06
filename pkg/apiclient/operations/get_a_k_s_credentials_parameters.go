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

// NewGetAKSCredentialsParams creates a new GetAKSCredentialsParams object
// with the default values initialized.
func NewGetAKSCredentialsParams() *GetAKSCredentialsParams {
	var ()
	return &GetAKSCredentialsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAKSCredentialsParamsWithTimeout creates a new GetAKSCredentialsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAKSCredentialsParamsWithTimeout(timeout time.Duration) *GetAKSCredentialsParams {
	var ()
	return &GetAKSCredentialsParams{

		timeout: timeout,
	}
}

// NewGetAKSCredentialsParamsWithContext creates a new GetAKSCredentialsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAKSCredentialsParamsWithContext(ctx context.Context) *GetAKSCredentialsParams {
	var ()
	return &GetAKSCredentialsParams{

		Context: ctx,
	}
}

// NewGetAKSCredentialsParamsWithHTTPClient creates a new GetAKSCredentialsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAKSCredentialsParamsWithHTTPClient(client *http.Client) *GetAKSCredentialsParams {
	var ()
	return &GetAKSCredentialsParams{
		HTTPClient: client,
	}
}

/*GetAKSCredentialsParams contains all the parameters to send to the API endpoint
for the get a k s credentials operation typically these are written to a http.Request
*/
type GetAKSCredentialsParams struct {

	/*Name
	  Is name the of the AKS Credentials you are acting upon

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

// WithTimeout adds the timeout to the get a k s credentials params
func (o *GetAKSCredentialsParams) WithTimeout(timeout time.Duration) *GetAKSCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get a k s credentials params
func (o *GetAKSCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get a k s credentials params
func (o *GetAKSCredentialsParams) WithContext(ctx context.Context) *GetAKSCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get a k s credentials params
func (o *GetAKSCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get a k s credentials params
func (o *GetAKSCredentialsParams) WithHTTPClient(client *http.Client) *GetAKSCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get a k s credentials params
func (o *GetAKSCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get a k s credentials params
func (o *GetAKSCredentialsParams) WithName(name string) *GetAKSCredentialsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get a k s credentials params
func (o *GetAKSCredentialsParams) SetName(name string) {
	o.Name = name
}

// WithTeam adds the team to the get a k s credentials params
func (o *GetAKSCredentialsParams) WithTeam(team string) *GetAKSCredentialsParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the get a k s credentials params
func (o *GetAKSCredentialsParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *GetAKSCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
