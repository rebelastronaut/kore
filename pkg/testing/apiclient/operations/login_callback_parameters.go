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

// NewLoginCallbackParams creates a new LoginCallbackParams object
// with the default values initialized.
func NewLoginCallbackParams() *LoginCallbackParams {
	var ()
	return &LoginCallbackParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLoginCallbackParamsWithTimeout creates a new LoginCallbackParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLoginCallbackParamsWithTimeout(timeout time.Duration) *LoginCallbackParams {
	var ()
	return &LoginCallbackParams{

		timeout: timeout,
	}
}

// NewLoginCallbackParamsWithContext creates a new LoginCallbackParams object
// with the default values initialized, and the ability to set a context for a request
func NewLoginCallbackParamsWithContext(ctx context.Context) *LoginCallbackParams {
	var ()
	return &LoginCallbackParams{

		Context: ctx,
	}
}

// NewLoginCallbackParamsWithHTTPClient creates a new LoginCallbackParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLoginCallbackParamsWithHTTPClient(client *http.Client) *LoginCallbackParams {
	var ()
	return &LoginCallbackParams{
		HTTPClient: client,
	}
}

/*LoginCallbackParams contains all the parameters to send to the API endpoint
for the login callback operation typically these are written to a http.Request
*/
type LoginCallbackParams struct {

	/*Code
	  The authorization code returned from the identity provider

	*/
	Code string
	/*State
	  The state parameter which was passed on authorization request

	*/
	State string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the login callback params
func (o *LoginCallbackParams) WithTimeout(timeout time.Duration) *LoginCallbackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the login callback params
func (o *LoginCallbackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the login callback params
func (o *LoginCallbackParams) WithContext(ctx context.Context) *LoginCallbackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the login callback params
func (o *LoginCallbackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the login callback params
func (o *LoginCallbackParams) WithHTTPClient(client *http.Client) *LoginCallbackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the login callback params
func (o *LoginCallbackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCode adds the code to the login callback params
func (o *LoginCallbackParams) WithCode(code string) *LoginCallbackParams {
	o.SetCode(code)
	return o
}

// SetCode adds the code to the login callback params
func (o *LoginCallbackParams) SetCode(code string) {
	o.Code = code
}

// WithState adds the state to the login callback params
func (o *LoginCallbackParams) WithState(state string) *LoginCallbackParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the login callback params
func (o *LoginCallbackParams) SetState(state string) {
	o.State = state
}

// WriteToRequest writes these params to a swagger request
func (o *LoginCallbackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param code
	qrCode := o.Code
	qCode := qrCode
	if qCode != "" {
		if err := r.SetQueryParam("code", qCode); err != nil {
			return err
		}
	}

	// query param state
	qrState := o.State
	qState := qrState
	if qState != "" {
		if err := r.SetQueryParam("state", qState); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
