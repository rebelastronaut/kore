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

// GetSecurityScanReader is a Reader for the GetSecurityScan structure.
type GetSecurityScanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSecurityScanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSecurityScanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSecurityScanUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSecurityScanForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSecurityScanNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSecurityScanInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSecurityScanOK creates a GetSecurityScanOK with default headers values
func NewGetSecurityScanOK() *GetSecurityScanOK {
	return &GetSecurityScanOK{}
}

/*GetSecurityScanOK handles this case with default header values.

Security scan
*/
type GetSecurityScanOK struct {
	Payload *models.V1ScanResult
}

func (o *GetSecurityScanOK) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/securityscans/{id}][%d] getSecurityScanOK  %+v", 200, o.Payload)
}

func (o *GetSecurityScanOK) GetPayload() *models.V1ScanResult {
	return o.Payload
}

func (o *GetSecurityScanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ScanResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSecurityScanUnauthorized creates a GetSecurityScanUnauthorized with default headers values
func NewGetSecurityScanUnauthorized() *GetSecurityScanUnauthorized {
	return &GetSecurityScanUnauthorized{}
}

/*GetSecurityScanUnauthorized handles this case with default header values.

If not authenticated
*/
type GetSecurityScanUnauthorized struct {
}

func (o *GetSecurityScanUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/securityscans/{id}][%d] getSecurityScanUnauthorized ", 401)
}

func (o *GetSecurityScanUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSecurityScanForbidden creates a GetSecurityScanForbidden with default headers values
func NewGetSecurityScanForbidden() *GetSecurityScanForbidden {
	return &GetSecurityScanForbidden{}
}

/*GetSecurityScanForbidden handles this case with default header values.

If authenticated but not authorized
*/
type GetSecurityScanForbidden struct {
}

func (o *GetSecurityScanForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/securityscans/{id}][%d] getSecurityScanForbidden ", 403)
}

func (o *GetSecurityScanForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSecurityScanNotFound creates a GetSecurityScanNotFound with default headers values
func NewGetSecurityScanNotFound() *GetSecurityScanNotFound {
	return &GetSecurityScanNotFound{}
}

/*GetSecurityScanNotFound handles this case with default header values.

No current security scan exists for the ID
*/
type GetSecurityScanNotFound struct {
}

func (o *GetSecurityScanNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/securityscans/{id}][%d] getSecurityScanNotFound ", 404)
}

func (o *GetSecurityScanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSecurityScanInternalServerError creates a GetSecurityScanInternalServerError with default headers values
func NewGetSecurityScanInternalServerError() *GetSecurityScanInternalServerError {
	return &GetSecurityScanInternalServerError{}
}

/*GetSecurityScanInternalServerError handles this case with default header values.

A generic API error containing the cause of the error
*/
type GetSecurityScanInternalServerError struct {
	Payload *models.ApiserverError
}

func (o *GetSecurityScanInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/securityscans/{id}][%d] getSecurityScanInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSecurityScanInternalServerError) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *GetSecurityScanInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
