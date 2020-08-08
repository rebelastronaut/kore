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

// AuthorizeUserReader is a Reader for the AuthorizeUser structure.
type AuthorizeUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthorizeUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthorizeUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAuthorizeUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAuthorizeUserOK creates a AuthorizeUserOK with default headers values
func NewAuthorizeUserOK() *AuthorizeUserOK {
	return &AuthorizeUserOK{}
}

/*AuthorizeUserOK handles this case with default header values.

Contains the access token on successfully authentication
*/
type AuthorizeUserOK struct {
	Payload *models.TypesIssuedToken
}

func (o *AuthorizeUserOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/login/authorize/{user}][%d] authorizeUserOK  %+v", 200, o.Payload)
}

func (o *AuthorizeUserOK) GetPayload() *models.TypesIssuedToken {
	return o.Payload
}

func (o *AuthorizeUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TypesIssuedToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthorizeUserDefault creates a AuthorizeUserDefault with default headers values
func NewAuthorizeUserDefault(code int) *AuthorizeUserDefault {
	return &AuthorizeUserDefault{
		_statusCode: code,
	}
}

/*AuthorizeUserDefault handles this case with default header values.

A generic API error containing the cause of the error
*/
type AuthorizeUserDefault struct {
	_statusCode int

	Payload *models.ApiserverError
}

// Code gets the status code for the authorize user default response
func (o *AuthorizeUserDefault) Code() int {
	return o._statusCode
}

func (o *AuthorizeUserDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/login/authorize/{user}][%d] AuthorizeUser default  %+v", o._statusCode, o.Payload)
}

func (o *AuthorizeUserDefault) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *AuthorizeUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
