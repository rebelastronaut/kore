// Code generated by go-swagger; DO NOT EDIT.

package costs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new costs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for costs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetAssets(params *GetAssetsParams, authInfo runtime.ClientAuthInfoWriter) (*GetAssetsOK, error)

	GetCostSummary(params *GetCostSummaryParams, authInfo runtime.ClientAuthInfoWriter) (*GetCostSummaryOK, error)

	GetTeamCostSummary(params *GetTeamCostSummaryParams, authInfo runtime.ClientAuthInfoWriter) (*GetTeamCostSummaryOK, error)

	ListCosts(params *ListCostsParams, authInfo runtime.ClientAuthInfoWriter) (*ListCostsOK, error)

	PostCosts(params *PostCostsParams, authInfo runtime.ClientAuthInfoWriter) (*PostCostsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAssets returns details of the assets known to kore which should be monitored for costs by a costs provider
*/
func (a *Client) GetAssets(params *GetAssetsParams, authInfo runtime.ClientAuthInfoWriter) (*GetAssetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAssetsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAssets",
		Method:             "GET",
		PathPattern:        "/api/v1alpha1/costs/assets/{provider}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAssetsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAssetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAssets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCostSummary returns a summary of all costs known to kore for the specified time period
*/
func (a *Client) GetCostSummary(params *GetCostSummaryParams, authInfo runtime.ClientAuthInfoWriter) (*GetCostSummaryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCostSummaryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCostSummary",
		Method:             "GET",
		PathPattern:        "/api/v1alpha1/costs/summary/{from}/{to}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCostSummaryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCostSummaryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetCostSummary: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTeamCostSummary returns a summary of all costs known to kore for the specified time period
*/
func (a *Client) GetTeamCostSummary(params *GetTeamCostSummaryParams, authInfo runtime.ClientAuthInfoWriter) (*GetTeamCostSummaryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTeamCostSummaryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTeamCostSummary",
		Method:             "GET",
		PathPattern:        "/api/v1alpha1/costs/teamsummary/{teamIdentifier}/{from}/{to}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTeamCostSummaryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTeamCostSummaryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetTeamCostSummary: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListCosts returns a list of actual costs
*/
func (a *Client) ListCosts(params *ListCostsParams, authInfo runtime.ClientAuthInfoWriter) (*ListCostsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListCostsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListCosts",
		Method:             "GET",
		PathPattern:        "/api/v1alpha1/costs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListCostsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListCostsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListCosts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostCosts persists one or more asset costs
*/
func (a *Client) PostCosts(params *PostCostsParams, authInfo runtime.ClientAuthInfoWriter) (*PostCostsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCostsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostCosts",
		Method:             "POST",
		PathPattern:        "/api/v1alpha1/costs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostCostsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostCostsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostCosts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
