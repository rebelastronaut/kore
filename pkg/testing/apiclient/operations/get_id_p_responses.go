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

// GetIDPReader is a Reader for the GetIDP structure.
type GetIDPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIDPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIDPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetIDPNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetIDPDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIDPOK creates a GetIDPOK with default headers values
func NewGetIDPOK() *GetIDPOK {
	return &GetIDPOK{}
}

/*GetIDPOK handles this case with default header values.

the specified identity provider
*/
type GetIDPOK struct {
	Payload *apimodels.V1IDP
}

func (o *GetIDPOK) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/idp/configured/{name}][%d] getIdPOK  %+v", 200, o.Payload)
}

func (o *GetIDPOK) GetPayload() *apimodels.V1IDP {
	return o.Payload
}

func (o *GetIDPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(apimodels.V1IDP)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIDPNotFound creates a GetIDPNotFound with default headers values
func NewGetIDPNotFound() *GetIDPNotFound {
	return &GetIDPNotFound{}
}

/*GetIDPNotFound handles this case with default header values.

Indicate the class was not found in the kore
*/
type GetIDPNotFound struct {
}

func (o *GetIDPNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/idp/configured/{name}][%d] getIdPNotFound ", 404)
}

func (o *GetIDPNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetIDPDefault creates a GetIDPDefault with default headers values
func NewGetIDPDefault(code int) *GetIDPDefault {
	return &GetIDPDefault{
		_statusCode: code,
	}
}

/*GetIDPDefault handles this case with default header values.

A generic API error containing the cause of the error
*/
type GetIDPDefault struct {
	_statusCode int

	Payload *apimodels.ApiserverError
}

// Code gets the status code for the get ID p default response
func (o *GetIDPDefault) Code() int {
	return o._statusCode
}

func (o *GetIDPDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/idp/configured/{name}][%d] GetIDP default  %+v", o._statusCode, o.Payload)
}

func (o *GetIDPDefault) GetPayload() *apimodels.ApiserverError {
	return o.Payload
}

func (o *GetIDPDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(apimodels.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
