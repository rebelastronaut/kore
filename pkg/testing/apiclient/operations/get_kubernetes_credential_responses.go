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

// GetKubernetesCredentialReader is a Reader for the GetKubernetesCredential structure.
type GetKubernetesCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKubernetesCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKubernetesCredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetKubernetesCredentialInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetKubernetesCredentialOK creates a GetKubernetesCredentialOK with default headers values
func NewGetKubernetesCredentialOK() *GetKubernetesCredentialOK {
	return &GetKubernetesCredentialOK{}
}

/*GetKubernetesCredentialOK handles this case with default header values.

Contains the former team definition from the kore
*/
type GetKubernetesCredentialOK struct {
	Payload *apimodels.V1KubernetesCredentials
}

func (o *GetKubernetesCredentialOK) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/teams/{team}/kubernetescredentials/{name}][%d] getKubernetesCredentialOK  %+v", 200, o.Payload)
}

func (o *GetKubernetesCredentialOK) GetPayload() *apimodels.V1KubernetesCredentials {
	return o.Payload
}

func (o *GetKubernetesCredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(apimodels.V1KubernetesCredentials)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKubernetesCredentialInternalServerError creates a GetKubernetesCredentialInternalServerError with default headers values
func NewGetKubernetesCredentialInternalServerError() *GetKubernetesCredentialInternalServerError {
	return &GetKubernetesCredentialInternalServerError{}
}

/*GetKubernetesCredentialInternalServerError handles this case with default header values.

A generic API error containing the cause of the error
*/
type GetKubernetesCredentialInternalServerError struct {
	Payload *apimodels.ApiserverError
}

func (o *GetKubernetesCredentialInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/teams/{team}/kubernetescredentials/{name}][%d] getKubernetesCredentialInternalServerError  %+v", 500, o.Payload)
}

func (o *GetKubernetesCredentialInternalServerError) GetPayload() *apimodels.ApiserverError {
	return o.Payload
}

func (o *GetKubernetesCredentialInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(apimodels.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
