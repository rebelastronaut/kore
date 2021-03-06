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

// NewListServiceKindsParams creates a new ListServiceKindsParams object
// with the default values initialized.
func NewListServiceKindsParams() *ListServiceKindsParams {
	var ()
	return &ListServiceKindsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListServiceKindsParamsWithTimeout creates a new ListServiceKindsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListServiceKindsParamsWithTimeout(timeout time.Duration) *ListServiceKindsParams {
	var ()
	return &ListServiceKindsParams{

		timeout: timeout,
	}
}

// NewListServiceKindsParamsWithContext creates a new ListServiceKindsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListServiceKindsParamsWithContext(ctx context.Context) *ListServiceKindsParams {
	var ()
	return &ListServiceKindsParams{

		Context: ctx,
	}
}

// NewListServiceKindsParamsWithHTTPClient creates a new ListServiceKindsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListServiceKindsParamsWithHTTPClient(client *http.Client) *ListServiceKindsParams {
	var ()
	return &ListServiceKindsParams{
		HTTPClient: client,
	}
}

/*ListServiceKindsParams contains all the parameters to send to the API endpoint
for the list service kinds operation typically these are written to a http.Request
*/
type ListServiceKindsParams struct {

	/*Enabled
	  Filters service kinds for enabled/disabled status

	*/
	Enabled *string
	/*Platform
	  Filters service kinds for a specific service platform

	*/
	Platform *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list service kinds params
func (o *ListServiceKindsParams) WithTimeout(timeout time.Duration) *ListServiceKindsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list service kinds params
func (o *ListServiceKindsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list service kinds params
func (o *ListServiceKindsParams) WithContext(ctx context.Context) *ListServiceKindsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list service kinds params
func (o *ListServiceKindsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list service kinds params
func (o *ListServiceKindsParams) WithHTTPClient(client *http.Client) *ListServiceKindsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list service kinds params
func (o *ListServiceKindsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnabled adds the enabled to the list service kinds params
func (o *ListServiceKindsParams) WithEnabled(enabled *string) *ListServiceKindsParams {
	o.SetEnabled(enabled)
	return o
}

// SetEnabled adds the enabled to the list service kinds params
func (o *ListServiceKindsParams) SetEnabled(enabled *string) {
	o.Enabled = enabled
}

// WithPlatform adds the platform to the list service kinds params
func (o *ListServiceKindsParams) WithPlatform(platform *string) *ListServiceKindsParams {
	o.SetPlatform(platform)
	return o
}

// SetPlatform adds the platform to the list service kinds params
func (o *ListServiceKindsParams) SetPlatform(platform *string) {
	o.Platform = platform
}

// WriteToRequest writes these params to a swagger request
func (o *ListServiceKindsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Enabled != nil {

		// query param enabled
		var qrEnabled string
		if o.Enabled != nil {
			qrEnabled = *o.Enabled
		}
		qEnabled := qrEnabled
		if qEnabled != "" {
			if err := r.SetQueryParam("enabled", qEnabled); err != nil {
				return err
			}
		}

	}

	if o.Platform != nil {

		// query param platform
		var qrPlatform string
		if o.Platform != nil {
			qrPlatform = *o.Platform
		}
		qPlatform := qrPlatform
		if qPlatform != "" {
			if err := r.SetQueryParam("platform", qPlatform); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
