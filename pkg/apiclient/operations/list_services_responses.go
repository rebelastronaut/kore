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

// ListServicesReader is a Reader for the ListServices structure.
type ListServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListServicesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListServicesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListServicesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListServicesOK creates a ListServicesOK with default headers values
func NewListServicesOK() *ListServicesOK {
	return &ListServicesOK{}
}

/*ListServicesOK handles this case with default header values.

List of all services for a team
*/
type ListServicesOK struct {
	Payload *models.V1ServiceList
}

func (o *ListServicesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/teams/{team}/services][%d] listServicesOK  %+v", 200, o.Payload)
}

func (o *ListServicesOK) GetPayload() *models.V1ServiceList {
	return o.Payload
}

func (o *ListServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ServiceList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServicesUnauthorized creates a ListServicesUnauthorized with default headers values
func NewListServicesUnauthorized() *ListServicesUnauthorized {
	return &ListServicesUnauthorized{}
}

/*ListServicesUnauthorized handles this case with default header values.

If not authenticated
*/
type ListServicesUnauthorized struct {
}

func (o *ListServicesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/teams/{team}/services][%d] listServicesUnauthorized ", 401)
}

func (o *ListServicesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListServicesForbidden creates a ListServicesForbidden with default headers values
func NewListServicesForbidden() *ListServicesForbidden {
	return &ListServicesForbidden{}
}

/*ListServicesForbidden handles this case with default header values.

If authenticated but not authorized
*/
type ListServicesForbidden struct {
}

func (o *ListServicesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/teams/{team}/services][%d] listServicesForbidden ", 403)
}

func (o *ListServicesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListServicesInternalServerError creates a ListServicesInternalServerError with default headers values
func NewListServicesInternalServerError() *ListServicesInternalServerError {
	return &ListServicesInternalServerError{}
}

/*ListServicesInternalServerError handles this case with default header values.

A generic API error containing the cause of the error
*/
type ListServicesInternalServerError struct {
	Payload *models.ApiserverError
}

func (o *ListServicesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/teams/{team}/services][%d] listServicesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListServicesInternalServerError) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *ListServicesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
