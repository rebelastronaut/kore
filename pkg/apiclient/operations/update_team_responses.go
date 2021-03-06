// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/appvia/kore/pkg/apiclient/models"
)

// UpdateTeamReader is a Reader for the UpdateTeam structure.
type UpdateTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewUpdateTeamNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewUpdateTeamBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateTeamUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateTeamForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateTeamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateTeamInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateTeamOK creates a UpdateTeamOK with default headers values
func NewUpdateTeamOK() *UpdateTeamOK {
	return &UpdateTeamOK{}
}

/*UpdateTeamOK handles this case with default header values.

Contains the team definition from the kore
*/
type UpdateTeamOK struct {
	Payload *models.V1Team
}

func (o *UpdateTeamOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/teams/{team}][%d] updateTeamOK  %+v", 200, o.Payload)
}

func (o *UpdateTeamOK) GetPayload() *models.V1Team {
	return o.Payload
}

func (o *UpdateTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Team)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamNotModified creates a UpdateTeamNotModified with default headers values
func NewUpdateTeamNotModified() *UpdateTeamNotModified {
	return &UpdateTeamNotModified{}
}

/*UpdateTeamNotModified handles this case with default header values.

Indicates the request was processed but no changes applied
*/
type UpdateTeamNotModified struct {
	Payload *models.V1Team
}

func (o *UpdateTeamNotModified) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/teams/{team}][%d] updateTeamNotModified  %+v", 304, o.Payload)
}

func (o *UpdateTeamNotModified) GetPayload() *models.V1Team {
	return o.Payload
}

func (o *UpdateTeamNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Team)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamBadRequest creates a UpdateTeamBadRequest with default headers values
func NewUpdateTeamBadRequest() *UpdateTeamBadRequest {
	return &UpdateTeamBadRequest{}
}

/*UpdateTeamBadRequest handles this case with default header values.

Validation error of supplied parameters/body
*/
type UpdateTeamBadRequest struct {
	Payload *models.ValidationError
}

func (o *UpdateTeamBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/teams/{team}][%d] updateTeamBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateTeamBadRequest) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdateTeamBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamUnauthorized creates a UpdateTeamUnauthorized with default headers values
func NewUpdateTeamUnauthorized() *UpdateTeamUnauthorized {
	return &UpdateTeamUnauthorized{}
}

/*UpdateTeamUnauthorized handles this case with default header values.

If not authenticated
*/
type UpdateTeamUnauthorized struct {
}

func (o *UpdateTeamUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/teams/{team}][%d] updateTeamUnauthorized ", 401)
}

func (o *UpdateTeamUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateTeamForbidden creates a UpdateTeamForbidden with default headers values
func NewUpdateTeamForbidden() *UpdateTeamForbidden {
	return &UpdateTeamForbidden{}
}

/*UpdateTeamForbidden handles this case with default header values.

If authenticated but not authorized
*/
type UpdateTeamForbidden struct {
}

func (o *UpdateTeamForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/teams/{team}][%d] updateTeamForbidden ", 403)
}

func (o *UpdateTeamForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateTeamNotFound creates a UpdateTeamNotFound with default headers values
func NewUpdateTeamNotFound() *UpdateTeamNotFound {
	return &UpdateTeamNotFound{}
}

/*UpdateTeamNotFound handles this case with default header values.

Team does not exist
*/
type UpdateTeamNotFound struct {
}

func (o *UpdateTeamNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/teams/{team}][%d] updateTeamNotFound ", 404)
}

func (o *UpdateTeamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateTeamInternalServerError creates a UpdateTeamInternalServerError with default headers values
func NewUpdateTeamInternalServerError() *UpdateTeamInternalServerError {
	return &UpdateTeamInternalServerError{}
}

/*UpdateTeamInternalServerError handles this case with default header values.

A generic API error containing the cause of the error
*/
type UpdateTeamInternalServerError struct {
	Payload *models.ApiserverError
}

func (o *UpdateTeamInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/teams/{team}][%d] updateTeamInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateTeamInternalServerError) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *UpdateTeamInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
