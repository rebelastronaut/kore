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

// ListGKECredentialsReader is a Reader for the ListGKECredentials structure.
type ListGKECredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGKECredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListGKECredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewListGKECredentialsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListGKECredentialsOK creates a ListGKECredentialsOK with default headers values
func NewListGKECredentialsOK() *ListGKECredentialsOK {
	return &ListGKECredentialsOK{}
}

/*ListGKECredentialsOK handles this case with default header values.

Contains the former team definition from the kore
*/
type ListGKECredentialsOK struct {
	Payload *apimodels.V1alpha1GKECredentialsList
}

func (o *ListGKECredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/teams/{team}/gkecredentials][%d] listGKECredentialsOK  %+v", 200, o.Payload)
}

func (o *ListGKECredentialsOK) GetPayload() *apimodels.V1alpha1GKECredentialsList {
	return o.Payload
}

func (o *ListGKECredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(apimodels.V1alpha1GKECredentialsList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGKECredentialsInternalServerError creates a ListGKECredentialsInternalServerError with default headers values
func NewListGKECredentialsInternalServerError() *ListGKECredentialsInternalServerError {
	return &ListGKECredentialsInternalServerError{}
}

/*ListGKECredentialsInternalServerError handles this case with default header values.

A generic API error containing the cause of the error
*/
type ListGKECredentialsInternalServerError struct {
	Payload *apimodels.ApiserverError
}

func (o *ListGKECredentialsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/teams/{team}/gkecredentials][%d] listGKECredentialsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListGKECredentialsInternalServerError) GetPayload() *apimodels.ApiserverError {
	return o.Payload
}

func (o *ListGKECredentialsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(apimodels.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
