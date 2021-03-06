// Code generated by go-swagger; DO NOT EDIT.

package costs

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

// NewGetAssetsParams creates a new GetAssetsParams object
// with the default values initialized.
func NewGetAssetsParams() *GetAssetsParams {
	var ()
	return &GetAssetsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAssetsParamsWithTimeout creates a new GetAssetsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAssetsParamsWithTimeout(timeout time.Duration) *GetAssetsParams {
	var ()
	return &GetAssetsParams{

		timeout: timeout,
	}
}

// NewGetAssetsParamsWithContext creates a new GetAssetsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAssetsParamsWithContext(ctx context.Context) *GetAssetsParams {
	var ()
	return &GetAssetsParams{

		Context: ctx,
	}
}

// NewGetAssetsParamsWithHTTPClient creates a new GetAssetsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAssetsParamsWithHTTPClient(client *http.Client) *GetAssetsParams {
	var ()
	return &GetAssetsParams{
		HTTPClient: client,
	}
}

/*GetAssetsParams contains all the parameters to send to the API endpoint
for the get assets operation typically these are written to a http.Request
*/
type GetAssetsParams struct {

	/*Asset
	  Identifier of an asset to filter assets for

	*/
	Asset *string
	/*Provider
	  Cloud provider (e.g. gcp, aws, azure) to return asset metadata for

	*/
	Provider string
	/*Team
	  Identifier of a team to filter assets for

	*/
	Team *string
	/*WithDeleted
	  Set to true to include deleted assets

	*/
	WithDeleted *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get assets params
func (o *GetAssetsParams) WithTimeout(timeout time.Duration) *GetAssetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get assets params
func (o *GetAssetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get assets params
func (o *GetAssetsParams) WithContext(ctx context.Context) *GetAssetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get assets params
func (o *GetAssetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get assets params
func (o *GetAssetsParams) WithHTTPClient(client *http.Client) *GetAssetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get assets params
func (o *GetAssetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAsset adds the asset to the get assets params
func (o *GetAssetsParams) WithAsset(asset *string) *GetAssetsParams {
	o.SetAsset(asset)
	return o
}

// SetAsset adds the asset to the get assets params
func (o *GetAssetsParams) SetAsset(asset *string) {
	o.Asset = asset
}

// WithProvider adds the provider to the get assets params
func (o *GetAssetsParams) WithProvider(provider string) *GetAssetsParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the get assets params
func (o *GetAssetsParams) SetProvider(provider string) {
	o.Provider = provider
}

// WithTeam adds the team to the get assets params
func (o *GetAssetsParams) WithTeam(team *string) *GetAssetsParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the get assets params
func (o *GetAssetsParams) SetTeam(team *string) {
	o.Team = team
}

// WithWithDeleted adds the withDeleted to the get assets params
func (o *GetAssetsParams) WithWithDeleted(withDeleted *string) *GetAssetsParams {
	o.SetWithDeleted(withDeleted)
	return o
}

// SetWithDeleted adds the withDeleted to the get assets params
func (o *GetAssetsParams) SetWithDeleted(withDeleted *string) {
	o.WithDeleted = withDeleted
}

// WriteToRequest writes these params to a swagger request
func (o *GetAssetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Asset != nil {

		// query param asset
		var qrAsset string
		if o.Asset != nil {
			qrAsset = *o.Asset
		}
		qAsset := qrAsset
		if qAsset != "" {
			if err := r.SetQueryParam("asset", qAsset); err != nil {
				return err
			}
		}

	}

	// path param provider
	if err := r.SetPathParam("provider", o.Provider); err != nil {
		return err
	}

	if o.Team != nil {

		// query param team
		var qrTeam string
		if o.Team != nil {
			qrTeam = *o.Team
		}
		qTeam := qrTeam
		if qTeam != "" {
			if err := r.SetQueryParam("team", qTeam); err != nil {
				return err
			}
		}

	}

	if o.WithDeleted != nil {

		// query param with_deleted
		var qrWithDeleted string
		if o.WithDeleted != nil {
			qrWithDeleted = *o.WithDeleted
		}
		qWithDeleted := qrWithDeleted
		if qWithDeleted != "" {
			if err := r.SetQueryParam("with_deleted", qWithDeleted); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
