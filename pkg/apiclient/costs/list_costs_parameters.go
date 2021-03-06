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

// NewListCostsParams creates a new ListCostsParams object
// with the default values initialized.
func NewListCostsParams() *ListCostsParams {
	var ()
	return &ListCostsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListCostsParamsWithTimeout creates a new ListCostsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListCostsParamsWithTimeout(timeout time.Duration) *ListCostsParams {
	var ()
	return &ListCostsParams{

		timeout: timeout,
	}
}

// NewListCostsParamsWithContext creates a new ListCostsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListCostsParamsWithContext(ctx context.Context) *ListCostsParams {
	var ()
	return &ListCostsParams{

		Context: ctx,
	}
}

// NewListCostsParamsWithHTTPClient creates a new ListCostsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListCostsParamsWithHTTPClient(client *http.Client) *ListCostsParams {
	var ()
	return &ListCostsParams{
		HTTPClient: client,
	}
}

/*ListCostsParams contains all the parameters to send to the API endpoint
for the list costs operation typically these are written to a http.Request
*/
type ListCostsParams struct {

	/*Account
	  Account/project/subscription to return costs for

	*/
	Account *string
	/*Asset
	  Identifier of an asset to filter costs for

	*/
	Asset *string
	/*From
	  Start of time range to return costs for

	*/
	From *string
	/*Invoice
	  Invoice to return costs for, in the formay YYYYMM

	*/
	Invoice *string
	/*Provider
	  Cloud provider (e.g. gcp, aws, azure) to return costs for

	*/
	Provider *string
	/*Team
	  Identifier of a team to filter costs for

	*/
	Team *string
	/*To
	  End of time range to return costs for

	*/
	To *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list costs params
func (o *ListCostsParams) WithTimeout(timeout time.Duration) *ListCostsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list costs params
func (o *ListCostsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list costs params
func (o *ListCostsParams) WithContext(ctx context.Context) *ListCostsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list costs params
func (o *ListCostsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list costs params
func (o *ListCostsParams) WithHTTPClient(client *http.Client) *ListCostsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list costs params
func (o *ListCostsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccount adds the account to the list costs params
func (o *ListCostsParams) WithAccount(account *string) *ListCostsParams {
	o.SetAccount(account)
	return o
}

// SetAccount adds the account to the list costs params
func (o *ListCostsParams) SetAccount(account *string) {
	o.Account = account
}

// WithAsset adds the asset to the list costs params
func (o *ListCostsParams) WithAsset(asset *string) *ListCostsParams {
	o.SetAsset(asset)
	return o
}

// SetAsset adds the asset to the list costs params
func (o *ListCostsParams) SetAsset(asset *string) {
	o.Asset = asset
}

// WithFrom adds the from to the list costs params
func (o *ListCostsParams) WithFrom(from *string) *ListCostsParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the list costs params
func (o *ListCostsParams) SetFrom(from *string) {
	o.From = from
}

// WithInvoice adds the invoice to the list costs params
func (o *ListCostsParams) WithInvoice(invoice *string) *ListCostsParams {
	o.SetInvoice(invoice)
	return o
}

// SetInvoice adds the invoice to the list costs params
func (o *ListCostsParams) SetInvoice(invoice *string) {
	o.Invoice = invoice
}

// WithProvider adds the provider to the list costs params
func (o *ListCostsParams) WithProvider(provider *string) *ListCostsParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the list costs params
func (o *ListCostsParams) SetProvider(provider *string) {
	o.Provider = provider
}

// WithTeam adds the team to the list costs params
func (o *ListCostsParams) WithTeam(team *string) *ListCostsParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the list costs params
func (o *ListCostsParams) SetTeam(team *string) {
	o.Team = team
}

// WithTo adds the to to the list costs params
func (o *ListCostsParams) WithTo(to *string) *ListCostsParams {
	o.SetTo(to)
	return o
}

// SetTo adds the to to the list costs params
func (o *ListCostsParams) SetTo(to *string) {
	o.To = to
}

// WriteToRequest writes these params to a swagger request
func (o *ListCostsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Account != nil {

		// query param account
		var qrAccount string
		if o.Account != nil {
			qrAccount = *o.Account
		}
		qAccount := qrAccount
		if qAccount != "" {
			if err := r.SetQueryParam("account", qAccount); err != nil {
				return err
			}
		}

	}

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

	if o.From != nil {

		// query param from
		var qrFrom string
		if o.From != nil {
			qrFrom = *o.From
		}
		qFrom := qrFrom
		if qFrom != "" {
			if err := r.SetQueryParam("from", qFrom); err != nil {
				return err
			}
		}

	}

	if o.Invoice != nil {

		// query param invoice
		var qrInvoice string
		if o.Invoice != nil {
			qrInvoice = *o.Invoice
		}
		qInvoice := qrInvoice
		if qInvoice != "" {
			if err := r.SetQueryParam("invoice", qInvoice); err != nil {
				return err
			}
		}

	}

	if o.Provider != nil {

		// query param provider
		var qrProvider string
		if o.Provider != nil {
			qrProvider = *o.Provider
		}
		qProvider := qrProvider
		if qProvider != "" {
			if err := r.SetQueryParam("provider", qProvider); err != nil {
				return err
			}
		}

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

	if o.To != nil {

		// query param to
		var qrTo string
		if o.To != nil {
			qrTo = *o.To
		}
		qTo := qrTo
		if qTo != "" {
			if err := r.SetQueryParam("to", qTo); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
