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

// GetNamespaceReader is a Reader for the GetNamespace structure.
type GetNamespaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNamespaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNamespaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetNamespaceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNamespaceOK creates a GetNamespaceOK with default headers values
func NewGetNamespaceOK() *GetNamespaceOK {
	return &GetNamespaceOK{}
}

/*GetNamespaceOK handles this case with default header values.

Contains the former team definition from the kore
*/
type GetNamespaceOK struct {
	Payload *models.V1NamespaceClaim
}

func (o *GetNamespaceOK) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/teams/{team}/namespaceclaims/{name}][%d] getNamespaceOK  %+v", 200, o.Payload)
}

func (o *GetNamespaceOK) GetPayload() *models.V1NamespaceClaim {
	return o.Payload
}

func (o *GetNamespaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1NamespaceClaim)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNamespaceInternalServerError creates a GetNamespaceInternalServerError with default headers values
func NewGetNamespaceInternalServerError() *GetNamespaceInternalServerError {
	return &GetNamespaceInternalServerError{}
}

/*GetNamespaceInternalServerError handles this case with default header values.

A generic API error containing the cause of the error
*/
type GetNamespaceInternalServerError struct {
	Payload *models.ApiserverError
}

func (o *GetNamespaceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/teams/{team}/namespaceclaims/{name}][%d] getNamespaceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetNamespaceInternalServerError) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *GetNamespaceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
