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

// NewListServiceProvidersParams creates a new ListServiceProvidersParams object
// with the default values initialized.
func NewListServiceProvidersParams() *ListServiceProvidersParams {
	var ()
	return &ListServiceProvidersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListServiceProvidersParamsWithTimeout creates a new ListServiceProvidersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListServiceProvidersParamsWithTimeout(timeout time.Duration) *ListServiceProvidersParams {
	var ()
	return &ListServiceProvidersParams{

		timeout: timeout,
	}
}

// NewListServiceProvidersParamsWithContext creates a new ListServiceProvidersParams object
// with the default values initialized, and the ability to set a context for a request
func NewListServiceProvidersParamsWithContext(ctx context.Context) *ListServiceProvidersParams {
	var ()
	return &ListServiceProvidersParams{

		Context: ctx,
	}
}

// NewListServiceProvidersParamsWithHTTPClient creates a new ListServiceProvidersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListServiceProvidersParamsWithHTTPClient(client *http.Client) *ListServiceProvidersParams {
	var ()
	return &ListServiceProvidersParams{
		HTTPClient: client,
	}
}

/*ListServiceProvidersParams contains all the parameters to send to the API endpoint
for the list service providers operation typically these are written to a http.Request
*/
type ListServiceProvidersParams struct {

	/*Kind
	  Filters service providers for a specific kind

	*/
	Kind *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list service providers params
func (o *ListServiceProvidersParams) WithTimeout(timeout time.Duration) *ListServiceProvidersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list service providers params
func (o *ListServiceProvidersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list service providers params
func (o *ListServiceProvidersParams) WithContext(ctx context.Context) *ListServiceProvidersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list service providers params
func (o *ListServiceProvidersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list service providers params
func (o *ListServiceProvidersParams) WithHTTPClient(client *http.Client) *ListServiceProvidersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list service providers params
func (o *ListServiceProvidersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKind adds the kind to the list service providers params
func (o *ListServiceProvidersParams) WithKind(kind *string) *ListServiceProvidersParams {
	o.SetKind(kind)
	return o
}

// SetKind adds the kind to the list service providers params
func (o *ListServiceProvidersParams) SetKind(kind *string) {
	o.Kind = kind
}

// WriteToRequest writes these params to a swagger request
func (o *ListServiceProvidersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Kind != nil {

		// query param kind
		var qrKind string
		if o.Kind != nil {
			qrKind = *o.Kind
		}
		qKind := qrKind
		if qKind != "" {
			if err := r.SetQueryParam("kind", qKind); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
