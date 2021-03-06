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

// UpdatePlanReader is a Reader for the UpdatePlan structure.
type UpdatePlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdatePlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdatePlanBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdatePlanUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdatePlanForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdatePlanInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdatePlanOK creates a UpdatePlanOK with default headers values
func NewUpdatePlanOK() *UpdatePlanOK {
	return &UpdatePlanOK{}
}

/*UpdatePlanOK handles this case with default header values.

Contains the plan definition
*/
type UpdatePlanOK struct {
	Payload *models.V1Plan
}

func (o *UpdatePlanOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/plans/{name}][%d] updatePlanOK  %+v", 200, o.Payload)
}

func (o *UpdatePlanOK) GetPayload() *models.V1Plan {
	return o.Payload
}

func (o *UpdatePlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Plan)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePlanBadRequest creates a UpdatePlanBadRequest with default headers values
func NewUpdatePlanBadRequest() *UpdatePlanBadRequest {
	return &UpdatePlanBadRequest{}
}

/*UpdatePlanBadRequest handles this case with default header values.

Validation error of supplied parameters/body
*/
type UpdatePlanBadRequest struct {
	Payload *models.ValidationError
}

func (o *UpdatePlanBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/plans/{name}][%d] updatePlanBadRequest  %+v", 400, o.Payload)
}

func (o *UpdatePlanBadRequest) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdatePlanBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePlanUnauthorized creates a UpdatePlanUnauthorized with default headers values
func NewUpdatePlanUnauthorized() *UpdatePlanUnauthorized {
	return &UpdatePlanUnauthorized{}
}

/*UpdatePlanUnauthorized handles this case with default header values.

If not authenticated
*/
type UpdatePlanUnauthorized struct {
}

func (o *UpdatePlanUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/plans/{name}][%d] updatePlanUnauthorized ", 401)
}

func (o *UpdatePlanUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePlanForbidden creates a UpdatePlanForbidden with default headers values
func NewUpdatePlanForbidden() *UpdatePlanForbidden {
	return &UpdatePlanForbidden{}
}

/*UpdatePlanForbidden handles this case with default header values.

If authenticated but not authorized
*/
type UpdatePlanForbidden struct {
}

func (o *UpdatePlanForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/plans/{name}][%d] updatePlanForbidden ", 403)
}

func (o *UpdatePlanForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePlanInternalServerError creates a UpdatePlanInternalServerError with default headers values
func NewUpdatePlanInternalServerError() *UpdatePlanInternalServerError {
	return &UpdatePlanInternalServerError{}
}

/*UpdatePlanInternalServerError handles this case with default header values.

A generic API error containing the cause of the error
*/
type UpdatePlanInternalServerError struct {
	Payload *models.ApiserverError
}

func (o *UpdatePlanInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/plans/{name}][%d] updatePlanInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdatePlanInternalServerError) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *UpdatePlanInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
