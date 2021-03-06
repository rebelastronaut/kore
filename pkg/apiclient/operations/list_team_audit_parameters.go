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

// NewListTeamAuditParams creates a new ListTeamAuditParams object
// with the default values initialized.
func NewListTeamAuditParams() *ListTeamAuditParams {
	var (
		sinceDefault = string("60m")
	)
	return &ListTeamAuditParams{
		Since: &sinceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListTeamAuditParamsWithTimeout creates a new ListTeamAuditParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListTeamAuditParamsWithTimeout(timeout time.Duration) *ListTeamAuditParams {
	var (
		sinceDefault = string("60m")
	)
	return &ListTeamAuditParams{
		Since: &sinceDefault,

		timeout: timeout,
	}
}

// NewListTeamAuditParamsWithContext creates a new ListTeamAuditParams object
// with the default values initialized, and the ability to set a context for a request
func NewListTeamAuditParamsWithContext(ctx context.Context) *ListTeamAuditParams {
	var (
		sinceDefault = string("60m")
	)
	return &ListTeamAuditParams{
		Since: &sinceDefault,

		Context: ctx,
	}
}

// NewListTeamAuditParamsWithHTTPClient creates a new ListTeamAuditParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListTeamAuditParamsWithHTTPClient(client *http.Client) *ListTeamAuditParams {
	var (
		sinceDefault = string("60m")
	)
	return &ListTeamAuditParams{
		Since:      &sinceDefault,
		HTTPClient: client,
	}
}

/*ListTeamAuditParams contains all the parameters to send to the API endpoint
for the list team audit operation typically these are written to a http.Request
*/
type ListTeamAuditParams struct {

	/*Since
	  The duration to retrieve from the audit log

	*/
	Since *string
	/*Team
	  Is the name of the team you are acting within

	*/
	Team string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list team audit params
func (o *ListTeamAuditParams) WithTimeout(timeout time.Duration) *ListTeamAuditParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list team audit params
func (o *ListTeamAuditParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list team audit params
func (o *ListTeamAuditParams) WithContext(ctx context.Context) *ListTeamAuditParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list team audit params
func (o *ListTeamAuditParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list team audit params
func (o *ListTeamAuditParams) WithHTTPClient(client *http.Client) *ListTeamAuditParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list team audit params
func (o *ListTeamAuditParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSince adds the since to the list team audit params
func (o *ListTeamAuditParams) WithSince(since *string) *ListTeamAuditParams {
	o.SetSince(since)
	return o
}

// SetSince adds the since to the list team audit params
func (o *ListTeamAuditParams) SetSince(since *string) {
	o.Since = since
}

// WithTeam adds the team to the list team audit params
func (o *ListTeamAuditParams) WithTeam(team string) *ListTeamAuditParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the list team audit params
func (o *ListTeamAuditParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *ListTeamAuditParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Since != nil {

		// query param since
		var qrSince string
		if o.Since != nil {
			qrSince = *o.Since
		}
		qSince := qrSince
		if qSince != "" {
			if err := r.SetQueryParam("since", qSince); err != nil {
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
