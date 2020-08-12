// Code generated by go-swagger; DO NOT EDIT.

package costs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/appvia/kore/pkg/apiclient/models"
)

// GetTeamCostSummaryReader is a Reader for the GetTeamCostSummary structure.
type GetTeamCostSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTeamCostSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTeamCostSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetTeamCostSummaryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTeamCostSummaryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTeamCostSummaryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTeamCostSummaryOK creates a GetTeamCostSummaryOK with default headers values
func NewGetTeamCostSummaryOK() *GetTeamCostSummaryOK {
	return &GetTeamCostSummaryOK{}
}

/*GetTeamCostSummaryOK handles this case with default header values.

A summary of costs known to the system for the team
*/
type GetTeamCostSummaryOK struct {
	Payload *models.V1beta1TeamCostSummary
}

func (o *GetTeamCostSummaryOK) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/costs/teamsummary/{teamIdentifier}/{from}/{to}][%d] getTeamCostSummaryOK  %+v", 200, o.Payload)
}

func (o *GetTeamCostSummaryOK) GetPayload() *models.V1beta1TeamCostSummary {
	return o.Payload
}

func (o *GetTeamCostSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1beta1TeamCostSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamCostSummaryUnauthorized creates a GetTeamCostSummaryUnauthorized with default headers values
func NewGetTeamCostSummaryUnauthorized() *GetTeamCostSummaryUnauthorized {
	return &GetTeamCostSummaryUnauthorized{}
}

/*GetTeamCostSummaryUnauthorized handles this case with default header values.

If not authenticated
*/
type GetTeamCostSummaryUnauthorized struct {
}

func (o *GetTeamCostSummaryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/costs/teamsummary/{teamIdentifier}/{from}/{to}][%d] getTeamCostSummaryUnauthorized ", 401)
}

func (o *GetTeamCostSummaryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTeamCostSummaryForbidden creates a GetTeamCostSummaryForbidden with default headers values
func NewGetTeamCostSummaryForbidden() *GetTeamCostSummaryForbidden {
	return &GetTeamCostSummaryForbidden{}
}

/*GetTeamCostSummaryForbidden handles this case with default header values.

If authenticated but not authorized
*/
type GetTeamCostSummaryForbidden struct {
}

func (o *GetTeamCostSummaryForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/costs/teamsummary/{teamIdentifier}/{from}/{to}][%d] getTeamCostSummaryForbidden ", 403)
}

func (o *GetTeamCostSummaryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTeamCostSummaryInternalServerError creates a GetTeamCostSummaryInternalServerError with default headers values
func NewGetTeamCostSummaryInternalServerError() *GetTeamCostSummaryInternalServerError {
	return &GetTeamCostSummaryInternalServerError{}
}

/*GetTeamCostSummaryInternalServerError handles this case with default header values.

A generic API error containing the cause of the error
*/
type GetTeamCostSummaryInternalServerError struct {
	Payload *models.ApiserverError
}

func (o *GetTeamCostSummaryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/costs/teamsummary/{teamIdentifier}/{from}/{to}][%d] getTeamCostSummaryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTeamCostSummaryInternalServerError) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *GetTeamCostSummaryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
