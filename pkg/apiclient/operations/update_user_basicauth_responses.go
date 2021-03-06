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

// UpdateUserBasicauthReader is a Reader for the UpdateUserBasicauth structure.
type UpdateUserBasicauthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserBasicauthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUserBasicauthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateUserBasicauthUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateUserBasicauthForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateUserBasicauthNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateUserBasicauthInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateUserBasicauthOK creates a UpdateUserBasicauthOK with default headers values
func NewUpdateUserBasicauthOK() *UpdateUserBasicauthOK {
	return &UpdateUserBasicauthOK{}
}

/*UpdateUserBasicauthOK handles this case with default header values.

Contains the identities definitions from the kore
*/
type UpdateUserBasicauthOK struct {
}

func (o *UpdateUserBasicauthOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/identities/{user}/basicauth][%d] updateUserBasicauthOK ", 200)
}

func (o *UpdateUserBasicauthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserBasicauthUnauthorized creates a UpdateUserBasicauthUnauthorized with default headers values
func NewUpdateUserBasicauthUnauthorized() *UpdateUserBasicauthUnauthorized {
	return &UpdateUserBasicauthUnauthorized{}
}

/*UpdateUserBasicauthUnauthorized handles this case with default header values.

If not authenticated
*/
type UpdateUserBasicauthUnauthorized struct {
}

func (o *UpdateUserBasicauthUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/identities/{user}/basicauth][%d] updateUserBasicauthUnauthorized ", 401)
}

func (o *UpdateUserBasicauthUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserBasicauthForbidden creates a UpdateUserBasicauthForbidden with default headers values
func NewUpdateUserBasicauthForbidden() *UpdateUserBasicauthForbidden {
	return &UpdateUserBasicauthForbidden{}
}

/*UpdateUserBasicauthForbidden handles this case with default header values.

If authenticated but not authorized
*/
type UpdateUserBasicauthForbidden struct {
}

func (o *UpdateUserBasicauthForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/identities/{user}/basicauth][%d] updateUserBasicauthForbidden ", 403)
}

func (o *UpdateUserBasicauthForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserBasicauthNotFound creates a UpdateUserBasicauthNotFound with default headers values
func NewUpdateUserBasicauthNotFound() *UpdateUserBasicauthNotFound {
	return &UpdateUserBasicauthNotFound{}
}

/*UpdateUserBasicauthNotFound handles this case with default header values.

User does not exist
*/
type UpdateUserBasicauthNotFound struct {
}

func (o *UpdateUserBasicauthNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/identities/{user}/basicauth][%d] updateUserBasicauthNotFound ", 404)
}

func (o *UpdateUserBasicauthNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserBasicauthInternalServerError creates a UpdateUserBasicauthInternalServerError with default headers values
func NewUpdateUserBasicauthInternalServerError() *UpdateUserBasicauthInternalServerError {
	return &UpdateUserBasicauthInternalServerError{}
}

/*UpdateUserBasicauthInternalServerError handles this case with default header values.

A generic API error containing the cause of the error
*/
type UpdateUserBasicauthInternalServerError struct {
	Payload *models.ApiserverError
}

func (o *UpdateUserBasicauthInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/identities/{user}/basicauth][%d] updateUserBasicauthInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateUserBasicauthInternalServerError) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *UpdateUserBasicauthInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
