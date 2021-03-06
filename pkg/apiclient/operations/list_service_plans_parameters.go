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

// NewListServicePlansParams creates a new ListServicePlansParams object
// with the default values initialized.
func NewListServicePlansParams() *ListServicePlansParams {
	var ()
	return &ListServicePlansParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListServicePlansParamsWithTimeout creates a new ListServicePlansParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListServicePlansParamsWithTimeout(timeout time.Duration) *ListServicePlansParams {
	var ()
	return &ListServicePlansParams{

		timeout: timeout,
	}
}

// NewListServicePlansParamsWithContext creates a new ListServicePlansParams object
// with the default values initialized, and the ability to set a context for a request
func NewListServicePlansParamsWithContext(ctx context.Context) *ListServicePlansParams {
	var ()
	return &ListServicePlansParams{

		Context: ctx,
	}
}

// NewListServicePlansParamsWithHTTPClient creates a new ListServicePlansParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListServicePlansParamsWithHTTPClient(client *http.Client) *ListServicePlansParams {
	var ()
	return &ListServicePlansParams{
		HTTPClient: client,
	}
}

/*ListServicePlansParams contains all the parameters to send to the API endpoint
for the list service plans operation typically these are written to a http.Request
*/
type ListServicePlansParams struct {

	/*Kind
	  Filters service plans for a specific kind

	*/
	Kind *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list service plans params
func (o *ListServicePlansParams) WithTimeout(timeout time.Duration) *ListServicePlansParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list service plans params
func (o *ListServicePlansParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list service plans params
func (o *ListServicePlansParams) WithContext(ctx context.Context) *ListServicePlansParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list service plans params
func (o *ListServicePlansParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list service plans params
func (o *ListServicePlansParams) WithHTTPClient(client *http.Client) *ListServicePlansParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list service plans params
func (o *ListServicePlansParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKind adds the kind to the list service plans params
func (o *ListServicePlansParams) WithKind(kind *string) *ListServicePlansParams {
	o.SetKind(kind)
	return o
}

// SetKind adds the kind to the list service plans params
func (o *ListServicePlansParams) SetKind(kind *string) {
	o.Kind = kind
}

// WriteToRequest writes these params to a swagger request
func (o *ListServicePlansParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
