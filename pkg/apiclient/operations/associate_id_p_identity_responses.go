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

// AssociateIDPIdentityReader is a Reader for the AssociateIDPIdentity structure.
type AssociateIDPIdentityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssociateIDPIdentityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssociateIDPIdentityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAssociateIDPIdentityUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAssociateIDPIdentityForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAssociateIDPIdentityNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAssociateIDPIdentityInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAssociateIDPIdentityOK creates a AssociateIDPIdentityOK with default headers values
func NewAssociateIDPIdentityOK() *AssociateIDPIdentityOK {
	return &AssociateIDPIdentityOK{}
}

/*AssociateIDPIdentityOK handles this case with default header values.

Contains the identities definitions from the kore
*/
type AssociateIDPIdentityOK struct {
}

func (o *AssociateIDPIdentityOK) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/identities/{user}/associate][%d] associateIdPIdentityOK ", 200)
}

func (o *AssociateIDPIdentityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAssociateIDPIdentityUnauthorized creates a AssociateIDPIdentityUnauthorized with default headers values
func NewAssociateIDPIdentityUnauthorized() *AssociateIDPIdentityUnauthorized {
	return &AssociateIDPIdentityUnauthorized{}
}

/*AssociateIDPIdentityUnauthorized handles this case with default header values.

If not authenticated
*/
type AssociateIDPIdentityUnauthorized struct {
}

func (o *AssociateIDPIdentityUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/identities/{user}/associate][%d] associateIdPIdentityUnauthorized ", 401)
}

func (o *AssociateIDPIdentityUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAssociateIDPIdentityForbidden creates a AssociateIDPIdentityForbidden with default headers values
func NewAssociateIDPIdentityForbidden() *AssociateIDPIdentityForbidden {
	return &AssociateIDPIdentityForbidden{}
}

/*AssociateIDPIdentityForbidden handles this case with default header values.

If authenticated but not authorized
*/
type AssociateIDPIdentityForbidden struct {
}

func (o *AssociateIDPIdentityForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/identities/{user}/associate][%d] associateIdPIdentityForbidden ", 403)
}

func (o *AssociateIDPIdentityForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAssociateIDPIdentityNotFound creates a AssociateIDPIdentityNotFound with default headers values
func NewAssociateIDPIdentityNotFound() *AssociateIDPIdentityNotFound {
	return &AssociateIDPIdentityNotFound{}
}

/*AssociateIDPIdentityNotFound handles this case with default header values.

User does not exist
*/
type AssociateIDPIdentityNotFound struct {
}

func (o *AssociateIDPIdentityNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/identities/{user}/associate][%d] associateIdPIdentityNotFound ", 404)
}

func (o *AssociateIDPIdentityNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAssociateIDPIdentityInternalServerError creates a AssociateIDPIdentityInternalServerError with default headers values
func NewAssociateIDPIdentityInternalServerError() *AssociateIDPIdentityInternalServerError {
	return &AssociateIDPIdentityInternalServerError{}
}

/*AssociateIDPIdentityInternalServerError handles this case with default header values.

A generic API error containing the cause of the error
*/
type AssociateIDPIdentityInternalServerError struct {
	Payload *models.ApiserverError
}

func (o *AssociateIDPIdentityInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/identities/{user}/associate][%d] associateIdPIdentityInternalServerError  %+v", 500, o.Payload)
}

func (o *AssociateIDPIdentityInternalServerError) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *AssociateIDPIdentityInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
