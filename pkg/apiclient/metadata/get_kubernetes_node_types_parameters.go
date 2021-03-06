// Code generated by go-swagger; DO NOT EDIT.

package metadata

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

// NewGetKubernetesNodeTypesParams creates a new GetKubernetesNodeTypesParams object
// with the default values initialized.
func NewGetKubernetesNodeTypesParams() *GetKubernetesNodeTypesParams {
	var ()
	return &GetKubernetesNodeTypesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKubernetesNodeTypesParamsWithTimeout creates a new GetKubernetesNodeTypesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKubernetesNodeTypesParamsWithTimeout(timeout time.Duration) *GetKubernetesNodeTypesParams {
	var ()
	return &GetKubernetesNodeTypesParams{

		timeout: timeout,
	}
}

// NewGetKubernetesNodeTypesParamsWithContext creates a new GetKubernetesNodeTypesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetKubernetesNodeTypesParamsWithContext(ctx context.Context) *GetKubernetesNodeTypesParams {
	var ()
	return &GetKubernetesNodeTypesParams{

		Context: ctx,
	}
}

// NewGetKubernetesNodeTypesParamsWithHTTPClient creates a new GetKubernetesNodeTypesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetKubernetesNodeTypesParamsWithHTTPClient(client *http.Client) *GetKubernetesNodeTypesParams {
	var ()
	return &GetKubernetesNodeTypesParams{
		HTTPClient: client,
	}
}

/*GetKubernetesNodeTypesParams contains all the parameters to send to the API endpoint
for the get kubernetes node types operation typically these are written to a http.Request
*/
type GetKubernetesNodeTypesParams struct {

	/*Provider
	  The kubernetes provider to retrieve instance types/prices for

	*/
	Provider string
	/*Region
	  The region to retrieve instance types/prices for

	*/
	Region string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get kubernetes node types params
func (o *GetKubernetesNodeTypesParams) WithTimeout(timeout time.Duration) *GetKubernetesNodeTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get kubernetes node types params
func (o *GetKubernetesNodeTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get kubernetes node types params
func (o *GetKubernetesNodeTypesParams) WithContext(ctx context.Context) *GetKubernetesNodeTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get kubernetes node types params
func (o *GetKubernetesNodeTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get kubernetes node types params
func (o *GetKubernetesNodeTypesParams) WithHTTPClient(client *http.Client) *GetKubernetesNodeTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get kubernetes node types params
func (o *GetKubernetesNodeTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProvider adds the provider to the get kubernetes node types params
func (o *GetKubernetesNodeTypesParams) WithProvider(provider string) *GetKubernetesNodeTypesParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the get kubernetes node types params
func (o *GetKubernetesNodeTypesParams) SetProvider(provider string) {
	o.Provider = provider
}

// WithRegion adds the region to the get kubernetes node types params
func (o *GetKubernetesNodeTypesParams) WithRegion(region string) *GetKubernetesNodeTypesParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the get kubernetes node types params
func (o *GetKubernetesNodeTypesParams) SetRegion(region string) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *GetKubernetesNodeTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param provider
	if err := r.SetPathParam("provider", o.Provider); err != nil {
		return err
	}

	// path param region
	if err := r.SetPathParam("region", o.Region); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
