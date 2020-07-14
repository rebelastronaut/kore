// Code generated by go-swagger; DO NOT EDIT.

package monitoring

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/appvia/kore/pkg/apiclient/models"
)

// SilenceAlertReader is a Reader for the SilenceAlert structure.
type SilenceAlertReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SilenceAlertReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSilenceAlertOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSilenceAlertBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSilenceAlertUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSilenceAlertForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSilenceAlertInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSilenceAlertOK creates a SilenceAlertOK with default headers values
func NewSilenceAlertOK() *SilenceAlertOK {
	return &SilenceAlertOK{}
}

/*SilenceAlertOK handles this case with default header values.

The alert has been successfully silenced
*/
type SilenceAlertOK struct {
}

func (o *SilenceAlertOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/monitoring/alerts/silence/{uid}][%d] silenceAlertOK ", 200)
}

func (o *SilenceAlertOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSilenceAlertBadRequest creates a SilenceAlertBadRequest with default headers values
func NewSilenceAlertBadRequest() *SilenceAlertBadRequest {
	return &SilenceAlertBadRequest{}
}

/*SilenceAlertBadRequest handles this case with default header values.

Validation error of supplied parameters/body
*/
type SilenceAlertBadRequest struct {
	Payload *models.ValidationError
}

func (o *SilenceAlertBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/monitoring/alerts/silence/{uid}][%d] silenceAlertBadRequest  %+v", 400, o.Payload)
}

func (o *SilenceAlertBadRequest) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *SilenceAlertBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSilenceAlertUnauthorized creates a SilenceAlertUnauthorized with default headers values
func NewSilenceAlertUnauthorized() *SilenceAlertUnauthorized {
	return &SilenceAlertUnauthorized{}
}

/*SilenceAlertUnauthorized handles this case with default header values.

If not authenticated
*/
type SilenceAlertUnauthorized struct {
}

func (o *SilenceAlertUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/monitoring/alerts/silence/{uid}][%d] silenceAlertUnauthorized ", 401)
}

func (o *SilenceAlertUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSilenceAlertForbidden creates a SilenceAlertForbidden with default headers values
func NewSilenceAlertForbidden() *SilenceAlertForbidden {
	return &SilenceAlertForbidden{}
}

/*SilenceAlertForbidden handles this case with default header values.

If authenticated but not authorized
*/
type SilenceAlertForbidden struct {
}

func (o *SilenceAlertForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/monitoring/alerts/silence/{uid}][%d] silenceAlertForbidden ", 403)
}

func (o *SilenceAlertForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSilenceAlertInternalServerError creates a SilenceAlertInternalServerError with default headers values
func NewSilenceAlertInternalServerError() *SilenceAlertInternalServerError {
	return &SilenceAlertInternalServerError{}
}

/*SilenceAlertInternalServerError handles this case with default header values.

A generic API error containing the cause of the error
*/
type SilenceAlertInternalServerError struct {
	Payload *models.ApiserverError
}

func (o *SilenceAlertInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/monitoring/alerts/silence/{uid}][%d] silenceAlertInternalServerError  %+v", 500, o.Payload)
}

func (o *SilenceAlertInternalServerError) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *SilenceAlertInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
