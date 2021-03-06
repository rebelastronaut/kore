// Code generated by go-swagger; DO NOT EDIT.

package monitoring

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

	"github.com/appvia/kore/pkg/apiclient/models"
)

// NewUpdateRuleParams creates a new UpdateRuleParams object
// with the default values initialized.
func NewUpdateRuleParams() *UpdateRuleParams {
	var ()
	return &UpdateRuleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRuleParamsWithTimeout creates a new UpdateRuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateRuleParamsWithTimeout(timeout time.Duration) *UpdateRuleParams {
	var ()
	return &UpdateRuleParams{

		timeout: timeout,
	}
}

// NewUpdateRuleParamsWithContext creates a new UpdateRuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateRuleParamsWithContext(ctx context.Context) *UpdateRuleParams {
	var ()
	return &UpdateRuleParams{

		Context: ctx,
	}
}

// NewUpdateRuleParamsWithHTTPClient creates a new UpdateRuleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateRuleParamsWithHTTPClient(client *http.Client) *UpdateRuleParams {
	var ()
	return &UpdateRuleParams{
		HTTPClient: client,
	}
}

/*UpdateRuleParams contains all the parameters to send to the API endpoint
for the update rule operation typically these are written to a http.Request
*/
type UpdateRuleParams struct {

	/*Body
	  The specification for a rule in the kore

	*/
	Body *models.V1beta1AlertRule
	/*Group
	  Is the group of the kind

	*/
	Group string
	/*Kind
	  Is the kind of the resource

	*/
	Kind string
	/*Name
	  Is the name of the alerting rule

	*/
	Name string
	/*Namespace
	  Is the namespace of the resource

	*/
	Namespace string
	/*Resource
	  Is the name of the resource

	*/
	Resource string
	/*Version
	  Is the version of the kind

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update rule params
func (o *UpdateRuleParams) WithTimeout(timeout time.Duration) *UpdateRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update rule params
func (o *UpdateRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update rule params
func (o *UpdateRuleParams) WithContext(ctx context.Context) *UpdateRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update rule params
func (o *UpdateRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update rule params
func (o *UpdateRuleParams) WithHTTPClient(client *http.Client) *UpdateRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update rule params
func (o *UpdateRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update rule params
func (o *UpdateRuleParams) WithBody(body *models.V1beta1AlertRule) *UpdateRuleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update rule params
func (o *UpdateRuleParams) SetBody(body *models.V1beta1AlertRule) {
	o.Body = body
}

// WithGroup adds the group to the update rule params
func (o *UpdateRuleParams) WithGroup(group string) *UpdateRuleParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the update rule params
func (o *UpdateRuleParams) SetGroup(group string) {
	o.Group = group
}

// WithKind adds the kind to the update rule params
func (o *UpdateRuleParams) WithKind(kind string) *UpdateRuleParams {
	o.SetKind(kind)
	return o
}

// SetKind adds the kind to the update rule params
func (o *UpdateRuleParams) SetKind(kind string) {
	o.Kind = kind
}

// WithName adds the name to the update rule params
func (o *UpdateRuleParams) WithName(name string) *UpdateRuleParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update rule params
func (o *UpdateRuleParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the update rule params
func (o *UpdateRuleParams) WithNamespace(namespace string) *UpdateRuleParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the update rule params
func (o *UpdateRuleParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithResource adds the resource to the update rule params
func (o *UpdateRuleParams) WithResource(resource string) *UpdateRuleParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the update rule params
func (o *UpdateRuleParams) SetResource(resource string) {
	o.Resource = resource
}

// WithVersion adds the version to the update rule params
func (o *UpdateRuleParams) WithVersion(version string) *UpdateRuleParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the update rule params
func (o *UpdateRuleParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param group
	if err := r.SetPathParam("group", o.Group); err != nil {
		return err
	}

	// path param kind
	if err := r.SetPathParam("kind", o.Kind); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param resource
	if err := r.SetPathParam("resource", o.Resource); err != nil {
		return err
	}

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
