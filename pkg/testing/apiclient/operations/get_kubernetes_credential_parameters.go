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

// NewGetKubernetesCredentialParams creates a new GetKubernetesCredentialParams object
// with the default values initialized.
func NewGetKubernetesCredentialParams() *GetKubernetesCredentialParams {
	var ()
	return &GetKubernetesCredentialParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKubernetesCredentialParamsWithTimeout creates a new GetKubernetesCredentialParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKubernetesCredentialParamsWithTimeout(timeout time.Duration) *GetKubernetesCredentialParams {
	var ()
	return &GetKubernetesCredentialParams{

		timeout: timeout,
	}
}

// NewGetKubernetesCredentialParamsWithContext creates a new GetKubernetesCredentialParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetKubernetesCredentialParamsWithContext(ctx context.Context) *GetKubernetesCredentialParams {
	var ()
	return &GetKubernetesCredentialParams{

		Context: ctx,
	}
}

// NewGetKubernetesCredentialParamsWithHTTPClient creates a new GetKubernetesCredentialParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetKubernetesCredentialParamsWithHTTPClient(client *http.Client) *GetKubernetesCredentialParams {
	var ()
	return &GetKubernetesCredentialParams{
		HTTPClient: client,
	}
}

/*GetKubernetesCredentialParams contains all the parameters to send to the API endpoint
for the get kubernetes credential operation typically these are written to a http.Request
*/
type GetKubernetesCredentialParams struct {

	/*Name
	  Is name the of the kubernetes credentials you are acting upon

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

// WithTimeout adds the timeout to the get kubernetes credential params
func (o *GetKubernetesCredentialParams) WithTimeout(timeout time.Duration) *GetKubernetesCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get kubernetes credential params
func (o *GetKubernetesCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get kubernetes credential params
func (o *GetKubernetesCredentialParams) WithContext(ctx context.Context) *GetKubernetesCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get kubernetes credential params
func (o *GetKubernetesCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get kubernetes credential params
func (o *GetKubernetesCredentialParams) WithHTTPClient(client *http.Client) *GetKubernetesCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get kubernetes credential params
func (o *GetKubernetesCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get kubernetes credential params
func (o *GetKubernetesCredentialParams) WithName(name string) *GetKubernetesCredentialParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get kubernetes credential params
func (o *GetKubernetesCredentialParams) SetName(name string) {
	o.Name = name
}

// WithTeam adds the team to the get kubernetes credential params
func (o *GetKubernetesCredentialParams) WithTeam(team string) *GetKubernetesCredentialParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the get kubernetes credential params
func (o *GetKubernetesCredentialParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *GetKubernetesCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
