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

// UpdateEKSVPCReader is a Reader for the UpdateEKSVPC structure.
type UpdateEKSVPCReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateEKSVPCReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateEKSVPCOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateEKSVPCDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateEKSVPCOK creates a UpdateEKSVPCOK with default headers values
func NewUpdateEKSVPCOK() *UpdateEKSVPCOK {
	return &UpdateEKSVPCOK{}
}

/*UpdateEKSVPCOK handles this case with default header values.

Contains the former team definition from the kore
*/
type UpdateEKSVPCOK struct {
	Payload *models.V1alpha1EKSVPC
}

func (o *UpdateEKSVPCOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/teams/{team}/eksvpcs/{name}][%d] updateEKSVPCOK  %+v", 200, o.Payload)
}

func (o *UpdateEKSVPCOK) GetPayload() *models.V1alpha1EKSVPC {
	return o.Payload
}

func (o *UpdateEKSVPCOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1alpha1EKSVPC)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEKSVPCDefault creates a UpdateEKSVPCDefault with default headers values
func NewUpdateEKSVPCDefault(code int) *UpdateEKSVPCDefault {
	return &UpdateEKSVPCDefault{
		_statusCode: code,
	}
}

/*UpdateEKSVPCDefault handles this case with default header values.

A generic API error containing the cause of the error
*/
type UpdateEKSVPCDefault struct {
	_statusCode int

	Payload *models.ApiserverError
}

// Code gets the status code for the update e k s v p c default response
func (o *UpdateEKSVPCDefault) Code() int {
	return o._statusCode
}

func (o *UpdateEKSVPCDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/teams/{team}/eksvpcs/{name}][%d] updateEKSVPC default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateEKSVPCDefault) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *UpdateEKSVPCDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
