// Code generated by go-swagger; DO NOT EDIT.

package security

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
	"github.com/go-openapi/swag"
)

// NewListSecurityScansParams creates a new ListSecurityScansParams object
// with the default values initialized.
func NewListSecurityScansParams() *ListSecurityScansParams {
	var (
		latestOnlyDefault = bool(true)
	)
	return &ListSecurityScansParams{
		LatestOnly: &latestOnlyDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListSecurityScansParamsWithTimeout creates a new ListSecurityScansParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListSecurityScansParamsWithTimeout(timeout time.Duration) *ListSecurityScansParams {
	var (
		latestOnlyDefault = bool(true)
	)
	return &ListSecurityScansParams{
		LatestOnly: &latestOnlyDefault,

		timeout: timeout,
	}
}

// NewListSecurityScansParamsWithContext creates a new ListSecurityScansParams object
// with the default values initialized, and the ability to set a context for a request
func NewListSecurityScansParamsWithContext(ctx context.Context) *ListSecurityScansParams {
	var (
		latestOnlyDefault = bool(true)
	)
	return &ListSecurityScansParams{
		LatestOnly: &latestOnlyDefault,

		Context: ctx,
	}
}

// NewListSecurityScansParamsWithHTTPClient creates a new ListSecurityScansParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListSecurityScansParamsWithHTTPClient(client *http.Client) *ListSecurityScansParams {
	var (
		latestOnlyDefault = bool(true)
	)
	return &ListSecurityScansParams{
		LatestOnly: &latestOnlyDefault,
		HTTPClient: client,
	}
}

/*ListSecurityScansParams contains all the parameters to send to the API endpoint
for the list security scans operation typically these are written to a http.Request
*/
type ListSecurityScansParams struct {

	/*LatestOnly
	  Set to false to retrieve full history

	*/
	LatestOnly *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list security scans params
func (o *ListSecurityScansParams) WithTimeout(timeout time.Duration) *ListSecurityScansParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list security scans params
func (o *ListSecurityScansParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list security scans params
func (o *ListSecurityScansParams) WithContext(ctx context.Context) *ListSecurityScansParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list security scans params
func (o *ListSecurityScansParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list security scans params
func (o *ListSecurityScansParams) WithHTTPClient(client *http.Client) *ListSecurityScansParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list security scans params
func (o *ListSecurityScansParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLatestOnly adds the latestOnly to the list security scans params
func (o *ListSecurityScansParams) WithLatestOnly(latestOnly *bool) *ListSecurityScansParams {
	o.SetLatestOnly(latestOnly)
	return o
}

// SetLatestOnly adds the latestOnly to the list security scans params
func (o *ListSecurityScansParams) SetLatestOnly(latestOnly *bool) {
	o.LatestOnly = latestOnly
}

// WriteToRequest writes these params to a swagger request
func (o *ListSecurityScansParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.LatestOnly != nil {

		// query param latestOnly
		var qrLatestOnly bool
		if o.LatestOnly != nil {
			qrLatestOnly = *o.LatestOnly
		}
		qLatestOnly := swag.FormatBool(qrLatestOnly)
		if qLatestOnly != "" {
			if err := r.SetQueryParam("latestOnly", qLatestOnly); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
