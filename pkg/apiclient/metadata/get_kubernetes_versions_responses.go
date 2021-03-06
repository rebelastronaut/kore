// Code generated by go-swagger; DO NOT EDIT.

package metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/appvia/kore/pkg/apiclient/models"
)

// GetKubernetesVersionsReader is a Reader for the GetKubernetesVersions structure.
type GetKubernetesVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKubernetesVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKubernetesVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetKubernetesVersionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetKubernetesVersionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetKubernetesVersionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetKubernetesVersionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetKubernetesVersionsOK creates a GetKubernetesVersionsOK with default headers values
func NewGetKubernetesVersionsOK() *GetKubernetesVersionsOK {
	return &GetKubernetesVersionsOK{}
}

/*GetKubernetesVersionsOK handles this case with default header values.

A list of supported kubernetes versions
*/
type GetKubernetesVersionsOK struct {
	Payload []string
}

func (o *GetKubernetesVersionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/metadata/k8s/{provider}/regions/{region}/versions][%d] getKubernetesVersionsOK  %+v", 200, o.Payload)
}

func (o *GetKubernetesVersionsOK) GetPayload() []string {
	return o.Payload
}

func (o *GetKubernetesVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKubernetesVersionsUnauthorized creates a GetKubernetesVersionsUnauthorized with default headers values
func NewGetKubernetesVersionsUnauthorized() *GetKubernetesVersionsUnauthorized {
	return &GetKubernetesVersionsUnauthorized{}
}

/*GetKubernetesVersionsUnauthorized handles this case with default header values.

If not authenticated
*/
type GetKubernetesVersionsUnauthorized struct {
}

func (o *GetKubernetesVersionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/metadata/k8s/{provider}/regions/{region}/versions][%d] getKubernetesVersionsUnauthorized ", 401)
}

func (o *GetKubernetesVersionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesVersionsForbidden creates a GetKubernetesVersionsForbidden with default headers values
func NewGetKubernetesVersionsForbidden() *GetKubernetesVersionsForbidden {
	return &GetKubernetesVersionsForbidden{}
}

/*GetKubernetesVersionsForbidden handles this case with default header values.

If authenticated but not authorized
*/
type GetKubernetesVersionsForbidden struct {
}

func (o *GetKubernetesVersionsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/metadata/k8s/{provider}/regions/{region}/versions][%d] getKubernetesVersionsForbidden ", 403)
}

func (o *GetKubernetesVersionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesVersionsNotFound creates a GetKubernetesVersionsNotFound with default headers values
func NewGetKubernetesVersionsNotFound() *GetKubernetesVersionsNotFound {
	return &GetKubernetesVersionsNotFound{}
}

/*GetKubernetesVersionsNotFound handles this case with default header values.

provider or region doesn't exist
*/
type GetKubernetesVersionsNotFound struct {
}

func (o *GetKubernetesVersionsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/metadata/k8s/{provider}/regions/{region}/versions][%d] getKubernetesVersionsNotFound ", 404)
}

func (o *GetKubernetesVersionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetKubernetesVersionsInternalServerError creates a GetKubernetesVersionsInternalServerError with default headers values
func NewGetKubernetesVersionsInternalServerError() *GetKubernetesVersionsInternalServerError {
	return &GetKubernetesVersionsInternalServerError{}
}

/*GetKubernetesVersionsInternalServerError handles this case with default header values.

A generic API error containing the cause of the error
*/
type GetKubernetesVersionsInternalServerError struct {
	Payload *models.ApiserverError
}

func (o *GetKubernetesVersionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/metadata/k8s/{provider}/regions/{region}/versions][%d] getKubernetesVersionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetKubernetesVersionsInternalServerError) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *GetKubernetesVersionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
