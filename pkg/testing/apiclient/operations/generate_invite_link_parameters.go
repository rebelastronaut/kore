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

// NewGenerateInviteLinkParams creates a new GenerateInviteLinkParams object
// with the default values initialized.
func NewGenerateInviteLinkParams() *GenerateInviteLinkParams {
	var (
		expireDefault = string("1h")
	)
	return &GenerateInviteLinkParams{
		Expire: &expireDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGenerateInviteLinkParamsWithTimeout creates a new GenerateInviteLinkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGenerateInviteLinkParamsWithTimeout(timeout time.Duration) *GenerateInviteLinkParams {
	var (
		expireDefault = string("1h")
	)
	return &GenerateInviteLinkParams{
		Expire: &expireDefault,

		timeout: timeout,
	}
}

// NewGenerateInviteLinkParamsWithContext creates a new GenerateInviteLinkParams object
// with the default values initialized, and the ability to set a context for a request
func NewGenerateInviteLinkParamsWithContext(ctx context.Context) *GenerateInviteLinkParams {
	var (
		expireDefault = string("1h")
	)
	return &GenerateInviteLinkParams{
		Expire: &expireDefault,

		Context: ctx,
	}
}

// NewGenerateInviteLinkParamsWithHTTPClient creates a new GenerateInviteLinkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGenerateInviteLinkParamsWithHTTPClient(client *http.Client) *GenerateInviteLinkParams {
	var (
		expireDefault = string("1h")
	)
	return &GenerateInviteLinkParams{
		Expire:     &expireDefault,
		HTTPClient: client,
	}
}

/*GenerateInviteLinkParams contains all the parameters to send to the API endpoint
for the generate invite link operation typically these are written to a http.Request
*/
type GenerateInviteLinkParams struct {

	/*Expire
	  The expiration of the generated link

	*/
	Expire *string
	/*Team
	  The name of the team you are creating an invition link

	*/
	Team string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the generate invite link params
func (o *GenerateInviteLinkParams) WithTimeout(timeout time.Duration) *GenerateInviteLinkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the generate invite link params
func (o *GenerateInviteLinkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the generate invite link params
func (o *GenerateInviteLinkParams) WithContext(ctx context.Context) *GenerateInviteLinkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the generate invite link params
func (o *GenerateInviteLinkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the generate invite link params
func (o *GenerateInviteLinkParams) WithHTTPClient(client *http.Client) *GenerateInviteLinkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the generate invite link params
func (o *GenerateInviteLinkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpire adds the expire to the generate invite link params
func (o *GenerateInviteLinkParams) WithExpire(expire *string) *GenerateInviteLinkParams {
	o.SetExpire(expire)
	return o
}

// SetExpire adds the expire to the generate invite link params
func (o *GenerateInviteLinkParams) SetExpire(expire *string) {
	o.Expire = expire
}

// WithTeam adds the team to the generate invite link params
func (o *GenerateInviteLinkParams) WithTeam(team string) *GenerateInviteLinkParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the generate invite link params
func (o *GenerateInviteLinkParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *GenerateInviteLinkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Expire != nil {

		// query param expire
		var qrExpire string
		if o.Expire != nil {
			qrExpire = *o.Expire
		}
		qExpire := qrExpire
		if qExpire != "" {
			if err := r.SetQueryParam("expire", qExpire); err != nil {
				return err
			}
		}

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
