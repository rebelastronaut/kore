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

// GetServiceCatalogReader is a Reader for the GetServiceCatalog structure.
type GetServiceCatalogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceCatalogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServiceCatalogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetServiceCatalogUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetServiceCatalogForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetServiceCatalogNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetServiceCatalogInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetServiceCatalogOK creates a GetServiceCatalogOK with default headers values
func NewGetServiceCatalogOK() *GetServiceCatalogOK {
	return &GetServiceCatalogOK{}
}

/*GetServiceCatalogOK handles this case with default header values.

Contains the service catalog definition
*/
type GetServiceCatalogOK struct {
	Payload *models.V1ServiceCatalog
}

func (o *GetServiceCatalogOK) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/servicecatalogs/{name}][%d] getServiceCatalogOK  %+v", 200, o.Payload)
}

func (o *GetServiceCatalogOK) GetPayload() *models.V1ServiceCatalog {
	return o.Payload
}

func (o *GetServiceCatalogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ServiceCatalog)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceCatalogUnauthorized creates a GetServiceCatalogUnauthorized with default headers values
func NewGetServiceCatalogUnauthorized() *GetServiceCatalogUnauthorized {
	return &GetServiceCatalogUnauthorized{}
}

/*GetServiceCatalogUnauthorized handles this case with default header values.

If not authenticated
*/
type GetServiceCatalogUnauthorized struct {
}

func (o *GetServiceCatalogUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/servicecatalogs/{name}][%d] getServiceCatalogUnauthorized ", 401)
}

func (o *GetServiceCatalogUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServiceCatalogForbidden creates a GetServiceCatalogForbidden with default headers values
func NewGetServiceCatalogForbidden() *GetServiceCatalogForbidden {
	return &GetServiceCatalogForbidden{}
}

/*GetServiceCatalogForbidden handles this case with default header values.

If authenticated but not authorized
*/
type GetServiceCatalogForbidden struct {
}

func (o *GetServiceCatalogForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/servicecatalogs/{name}][%d] getServiceCatalogForbidden ", 403)
}

func (o *GetServiceCatalogForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServiceCatalogNotFound creates a GetServiceCatalogNotFound with default headers values
func NewGetServiceCatalogNotFound() *GetServiceCatalogNotFound {
	return &GetServiceCatalogNotFound{}
}

/*GetServiceCatalogNotFound handles this case with default header values.

the service catalog with the given name doesn't exist
*/
type GetServiceCatalogNotFound struct {
}

func (o *GetServiceCatalogNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/servicecatalogs/{name}][%d] getServiceCatalogNotFound ", 404)
}

func (o *GetServiceCatalogNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServiceCatalogInternalServerError creates a GetServiceCatalogInternalServerError with default headers values
func NewGetServiceCatalogInternalServerError() *GetServiceCatalogInternalServerError {
	return &GetServiceCatalogInternalServerError{}
}

/*GetServiceCatalogInternalServerError handles this case with default header values.

A generic API error containing the cause of the error
*/
type GetServiceCatalogInternalServerError struct {
	Payload *models.ApiserverError
}

func (o *GetServiceCatalogInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/servicecatalogs/{name}][%d] getServiceCatalogInternalServerError  %+v", 500, o.Payload)
}

func (o *GetServiceCatalogInternalServerError) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *GetServiceCatalogInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
