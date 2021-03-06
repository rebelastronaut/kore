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

// ListConfigReader is a Reader for the ListConfig structure.
type ListConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListConfigOK creates a ListConfigOK with default headers values
func NewListConfigOK() *ListConfigOK {
	return &ListConfigOK{}
}

/*ListConfigOK handles this case with default header values.

A list of all the config values
*/
type ListConfigOK struct {
	Payload *models.V1ConfigList
}

func (o *ListConfigOK) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/configs][%d] listConfigOK  %+v", 200, o.Payload)
}

func (o *ListConfigOK) GetPayload() *models.V1ConfigList {
	return o.Payload
}

func (o *ListConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ConfigList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListConfigUnauthorized creates a ListConfigUnauthorized with default headers values
func NewListConfigUnauthorized() *ListConfigUnauthorized {
	return &ListConfigUnauthorized{}
}

/*ListConfigUnauthorized handles this case with default header values.

If not authenticated
*/
type ListConfigUnauthorized struct {
}

func (o *ListConfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/configs][%d] listConfigUnauthorized ", 401)
}

func (o *ListConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListConfigForbidden creates a ListConfigForbidden with default headers values
func NewListConfigForbidden() *ListConfigForbidden {
	return &ListConfigForbidden{}
}

/*ListConfigForbidden handles this case with default header values.

If authenticated but not authorized
*/
type ListConfigForbidden struct {
}

func (o *ListConfigForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/configs][%d] listConfigForbidden ", 403)
}

func (o *ListConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListConfigInternalServerError creates a ListConfigInternalServerError with default headers values
func NewListConfigInternalServerError() *ListConfigInternalServerError {
	return &ListConfigInternalServerError{}
}

/*ListConfigInternalServerError handles this case with default header values.

A generic API error containing the cause of the error
*/
type ListConfigInternalServerError struct {
	Payload *models.ApiserverError
}

func (o *ListConfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/configs][%d] listConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *ListConfigInternalServerError) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *ListConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
