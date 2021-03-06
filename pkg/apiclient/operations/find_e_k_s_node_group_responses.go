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

// FindEKSNodeGroupReader is a Reader for the FindEKSNodeGroup structure.
type FindEKSNodeGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindEKSNodeGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindEKSNodeGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindEKSNodeGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindEKSNodeGroupOK creates a FindEKSNodeGroupOK with default headers values
func NewFindEKSNodeGroupOK() *FindEKSNodeGroupOK {
	return &FindEKSNodeGroupOK{}
}

/*FindEKSNodeGroupOK handles this case with default header values.

Contains the former team definition from the kore
*/
type FindEKSNodeGroupOK struct {
	Payload *models.V1alpha1EKSNodeGroup
}

func (o *FindEKSNodeGroupOK) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/teams/{team}/eksnodegroups/{name}][%d] findEKSNodeGroupOK  %+v", 200, o.Payload)
}

func (o *FindEKSNodeGroupOK) GetPayload() *models.V1alpha1EKSNodeGroup {
	return o.Payload
}

func (o *FindEKSNodeGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1alpha1EKSNodeGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindEKSNodeGroupDefault creates a FindEKSNodeGroupDefault with default headers values
func NewFindEKSNodeGroupDefault(code int) *FindEKSNodeGroupDefault {
	return &FindEKSNodeGroupDefault{
		_statusCode: code,
	}
}

/*FindEKSNodeGroupDefault handles this case with default header values.

A generic API error containing the cause of the error
*/
type FindEKSNodeGroupDefault struct {
	_statusCode int

	Payload *models.ApiserverError
}

// Code gets the status code for the find e k s node group default response
func (o *FindEKSNodeGroupDefault) Code() int {
	return o._statusCode
}

func (o *FindEKSNodeGroupDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/teams/{team}/eksnodegroups/{name}][%d] findEKSNodeGroup default  %+v", o._statusCode, o.Payload)
}

func (o *FindEKSNodeGroupDefault) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *FindEKSNodeGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
