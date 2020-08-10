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

// NewGetLoginMethodsParams creates a new GetLoginMethodsParams object
// with the default values initialized.
func NewGetLoginMethodsParams() *GetLoginMethodsParams {

	return &GetLoginMethodsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLoginMethodsParamsWithTimeout creates a new GetLoginMethodsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLoginMethodsParamsWithTimeout(timeout time.Duration) *GetLoginMethodsParams {

	return &GetLoginMethodsParams{

		timeout: timeout,
	}
}

// NewGetLoginMethodsParamsWithContext creates a new GetLoginMethodsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLoginMethodsParamsWithContext(ctx context.Context) *GetLoginMethodsParams {

	return &GetLoginMethodsParams{

		Context: ctx,
	}
}

// NewGetLoginMethodsParamsWithHTTPClient creates a new GetLoginMethodsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLoginMethodsParamsWithHTTPClient(client *http.Client) *GetLoginMethodsParams {

	return &GetLoginMethodsParams{
		HTTPClient: client,
	}
}

/*GetLoginMethodsParams contains all the parameters to send to the API endpoint
for the get login methods operation typically these are written to a http.Request
*/
type GetLoginMethodsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get login methods params
func (o *GetLoginMethodsParams) WithTimeout(timeout time.Duration) *GetLoginMethodsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get login methods params
func (o *GetLoginMethodsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get login methods params
func (o *GetLoginMethodsParams) WithContext(ctx context.Context) *GetLoginMethodsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get login methods params
func (o *GetLoginMethodsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get login methods params
func (o *GetLoginMethodsParams) WithHTTPClient(client *http.Client) *GetLoginMethodsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get login methods params
func (o *GetLoginMethodsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLoginMethodsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
