// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/appvia/kore/pkg/testing/apimodels"
)

// GetTeamReader is a Reader for the GetTeam structure.
type GetTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetTeamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTeamInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTeamOK creates a GetTeamOK with default headers values
func NewGetTeamOK() *GetTeamOK {
	return &GetTeamOK{}
}

/*GetTeamOK handles this case with default header values.

Contains the team definintion from the kore
*/
type GetTeamOK struct {
	Payload *apimodels.V1Team
}

func (o *GetTeamOK) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/teams/{team}][%d] getTeamOK  %+v", 200, o.Payload)
}

func (o *GetTeamOK) GetPayload() *apimodels.V1Team {
	return o.Payload
}

func (o *GetTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(apimodels.V1Team)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamNotFound creates a GetTeamNotFound with default headers values
func NewGetTeamNotFound() *GetTeamNotFound {
	return &GetTeamNotFound{}
}

/*GetTeamNotFound handles this case with default header values.

Team does not exist
*/
type GetTeamNotFound struct {
}

func (o *GetTeamNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/teams/{team}][%d] getTeamNotFound ", 404)
}

func (o *GetTeamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTeamInternalServerError creates a GetTeamInternalServerError with default headers values
func NewGetTeamInternalServerError() *GetTeamInternalServerError {
	return &GetTeamInternalServerError{}
}

/*GetTeamInternalServerError handles this case with default header values.

A generic API error containing the cause of the error
*/
type GetTeamInternalServerError struct {
	Payload *apimodels.ApiserverError
}

func (o *GetTeamInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/teams/{team}][%d] getTeamInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTeamInternalServerError) GetPayload() *apimodels.ApiserverError {
	return o.Payload
}

func (o *GetTeamInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(apimodels.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
