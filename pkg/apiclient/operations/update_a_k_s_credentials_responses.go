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

// UpdateAKSCredentialsReader is a Reader for the UpdateAKSCredentials structure.
type UpdateAKSCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAKSCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAKSCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateAKSCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateAKSCredentialsOK creates a UpdateAKSCredentialsOK with default headers values
func NewUpdateAKSCredentialsOK() *UpdateAKSCredentialsOK {
	return &UpdateAKSCredentialsOK{}
}

/*UpdateAKSCredentialsOK handles this case with default header values.

Contains the final definition
*/
type UpdateAKSCredentialsOK struct {
	Payload *models.V1alpha1AKSCredentials
}

func (o *UpdateAKSCredentialsOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/teams/{team}/akscredentials/{name}][%d] updateAKSCredentialsOK  %+v", 200, o.Payload)
}

func (o *UpdateAKSCredentialsOK) GetPayload() *models.V1alpha1AKSCredentials {
	return o.Payload
}

func (o *UpdateAKSCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1alpha1AKSCredentials)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAKSCredentialsDefault creates a UpdateAKSCredentialsDefault with default headers values
func NewUpdateAKSCredentialsDefault(code int) *UpdateAKSCredentialsDefault {
	return &UpdateAKSCredentialsDefault{
		_statusCode: code,
	}
}

/*UpdateAKSCredentialsDefault handles this case with default header values.

A generic API error containing the cause of the error
*/
type UpdateAKSCredentialsDefault struct {
	_statusCode int

	Payload *models.ApiserverError
}

// Code gets the status code for the update a k s credentials default response
func (o *UpdateAKSCredentialsDefault) Code() int {
	return o._statusCode
}

func (o *UpdateAKSCredentialsDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/teams/{team}/akscredentials/{name}][%d] UpdateAKSCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateAKSCredentialsDefault) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *UpdateAKSCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
