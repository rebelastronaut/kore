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

// DeleteServiceCatalogReader is a Reader for the DeleteServiceCatalog structure.
type DeleteServiceCatalogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteServiceCatalogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteServiceCatalogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteServiceCatalogBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteServiceCatalogUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteServiceCatalogForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteServiceCatalogNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteServiceCatalogInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteServiceCatalogOK creates a DeleteServiceCatalogOK with default headers values
func NewDeleteServiceCatalogOK() *DeleteServiceCatalogOK {
	return &DeleteServiceCatalogOK{}
}

/*DeleteServiceCatalogOK handles this case with default header values.

Contains the service catalog definition
*/
type DeleteServiceCatalogOK struct {
	Payload *models.V1ServiceCatalog
}

func (o *DeleteServiceCatalogOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1alpha1/servicecatalogs/{name}][%d] deleteServiceCatalogOK  %+v", 200, o.Payload)
}

func (o *DeleteServiceCatalogOK) GetPayload() *models.V1ServiceCatalog {
	return o.Payload
}

func (o *DeleteServiceCatalogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ServiceCatalog)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteServiceCatalogBadRequest creates a DeleteServiceCatalogBadRequest with default headers values
func NewDeleteServiceCatalogBadRequest() *DeleteServiceCatalogBadRequest {
	return &DeleteServiceCatalogBadRequest{}
}

/*DeleteServiceCatalogBadRequest handles this case with default header values.

Validation error of supplied parameters/body
*/
type DeleteServiceCatalogBadRequest struct {
	Payload *models.ValidationError
}

func (o *DeleteServiceCatalogBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1alpha1/servicecatalogs/{name}][%d] deleteServiceCatalogBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteServiceCatalogBadRequest) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *DeleteServiceCatalogBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteServiceCatalogUnauthorized creates a DeleteServiceCatalogUnauthorized with default headers values
func NewDeleteServiceCatalogUnauthorized() *DeleteServiceCatalogUnauthorized {
	return &DeleteServiceCatalogUnauthorized{}
}

/*DeleteServiceCatalogUnauthorized handles this case with default header values.

If not authenticated
*/
type DeleteServiceCatalogUnauthorized struct {
}

func (o *DeleteServiceCatalogUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1alpha1/servicecatalogs/{name}][%d] deleteServiceCatalogUnauthorized ", 401)
}

func (o *DeleteServiceCatalogUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServiceCatalogForbidden creates a DeleteServiceCatalogForbidden with default headers values
func NewDeleteServiceCatalogForbidden() *DeleteServiceCatalogForbidden {
	return &DeleteServiceCatalogForbidden{}
}

/*DeleteServiceCatalogForbidden handles this case with default header values.

If authenticated but not authorized
*/
type DeleteServiceCatalogForbidden struct {
}

func (o *DeleteServiceCatalogForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1alpha1/servicecatalogs/{name}][%d] deleteServiceCatalogForbidden ", 403)
}

func (o *DeleteServiceCatalogForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServiceCatalogNotFound creates a DeleteServiceCatalogNotFound with default headers values
func NewDeleteServiceCatalogNotFound() *DeleteServiceCatalogNotFound {
	return &DeleteServiceCatalogNotFound{}
}

/*DeleteServiceCatalogNotFound handles this case with default header values.

the service catalog with the given name doesn't exist
*/
type DeleteServiceCatalogNotFound struct {
}

func (o *DeleteServiceCatalogNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1alpha1/servicecatalogs/{name}][%d] deleteServiceCatalogNotFound ", 404)
}

func (o *DeleteServiceCatalogNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServiceCatalogInternalServerError creates a DeleteServiceCatalogInternalServerError with default headers values
func NewDeleteServiceCatalogInternalServerError() *DeleteServiceCatalogInternalServerError {
	return &DeleteServiceCatalogInternalServerError{}
}

/*DeleteServiceCatalogInternalServerError handles this case with default header values.

A generic API error containing the cause of the error
*/
type DeleteServiceCatalogInternalServerError struct {
	Payload *models.ApiserverError
}

func (o *DeleteServiceCatalogInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1alpha1/servicecatalogs/{name}][%d] deleteServiceCatalogInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteServiceCatalogInternalServerError) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *DeleteServiceCatalogInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
