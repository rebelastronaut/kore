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

// GetLoginMethodsReader is a Reader for the GetLoginMethods structure.
type GetLoginMethodsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLoginMethodsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLoginMethodsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetLoginMethodsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLoginMethodsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLoginMethodsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLoginMethodsOK creates a GetLoginMethodsOK with default headers values
func NewGetLoginMethodsOK() *GetLoginMethodsOK {
	return &GetLoginMethodsOK{}
}

/*GetLoginMethodsOK handles this case with default header values.

Details of which login providers are configured
*/
type GetLoginMethodsOK struct {
	Payload []string
}

func (o *GetLoginMethodsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/login/methods][%d] getLoginMethodsOK  %+v", 200, o.Payload)
}

func (o *GetLoginMethodsOK) GetPayload() []string {
	return o.Payload
}

func (o *GetLoginMethodsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLoginMethodsUnauthorized creates a GetLoginMethodsUnauthorized with default headers values
func NewGetLoginMethodsUnauthorized() *GetLoginMethodsUnauthorized {
	return &GetLoginMethodsUnauthorized{}
}

/*GetLoginMethodsUnauthorized handles this case with default header values.

If not authenticated
*/
type GetLoginMethodsUnauthorized struct {
}

func (o *GetLoginMethodsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/login/methods][%d] getLoginMethodsUnauthorized ", 401)
}

func (o *GetLoginMethodsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLoginMethodsForbidden creates a GetLoginMethodsForbidden with default headers values
func NewGetLoginMethodsForbidden() *GetLoginMethodsForbidden {
	return &GetLoginMethodsForbidden{}
}

/*GetLoginMethodsForbidden handles this case with default header values.

If authenticated but not authorized
*/
type GetLoginMethodsForbidden struct {
}

func (o *GetLoginMethodsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/login/methods][%d] getLoginMethodsForbidden ", 403)
}

func (o *GetLoginMethodsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLoginMethodsInternalServerError creates a GetLoginMethodsInternalServerError with default headers values
func NewGetLoginMethodsInternalServerError() *GetLoginMethodsInternalServerError {
	return &GetLoginMethodsInternalServerError{}
}

/*GetLoginMethodsInternalServerError handles this case with default header values.

A generic API error containing the cause of the error
*/
type GetLoginMethodsInternalServerError struct {
	Payload *models.ApiserverError
}

func (o *GetLoginMethodsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/login/methods][%d] getLoginMethodsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLoginMethodsInternalServerError) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *GetLoginMethodsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
