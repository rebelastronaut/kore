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

// ListAWSOrganizationsReader is a Reader for the ListAWSOrganizations structure.
type ListAWSOrganizationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAWSOrganizationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAWSOrganizationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListAWSOrganizationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAWSOrganizationsOK creates a ListAWSOrganizationsOK with default headers values
func NewListAWSOrganizationsOK() *ListAWSOrganizationsOK {
	return &ListAWSOrganizationsOK{}
}

/*ListAWSOrganizationsOK handles this case with default header values.

Contains the former team definition from the kore
*/
type ListAWSOrganizationsOK struct {
	Payload *models.V1alpha1AWSOrganizationList
}

func (o *ListAWSOrganizationsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/teams/{team}/awsorganizations][%d] listAWSOrganizationsOK  %+v", 200, o.Payload)
}

func (o *ListAWSOrganizationsOK) GetPayload() *models.V1alpha1AWSOrganizationList {
	return o.Payload
}

func (o *ListAWSOrganizationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1alpha1AWSOrganizationList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAWSOrganizationsDefault creates a ListAWSOrganizationsDefault with default headers values
func NewListAWSOrganizationsDefault(code int) *ListAWSOrganizationsDefault {
	return &ListAWSOrganizationsDefault{
		_statusCode: code,
	}
}

/*ListAWSOrganizationsDefault handles this case with default header values.

A generic API error containing the cause of the error
*/
type ListAWSOrganizationsDefault struct {
	_statusCode int

	Payload *models.ApiserverError
}

// Code gets the status code for the list a w s organizations default response
func (o *ListAWSOrganizationsDefault) Code() int {
	return o._statusCode
}

func (o *ListAWSOrganizationsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/teams/{team}/awsorganizations][%d] ListAWSOrganizations default  %+v", o._statusCode, o.Payload)
}

func (o *ListAWSOrganizationsDefault) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *ListAWSOrganizationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
