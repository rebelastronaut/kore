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

// PurgeRuleAlertsReader is a Reader for the PurgeRuleAlerts structure.
type PurgeRuleAlertsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PurgeRuleAlertsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPurgeRuleAlertsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPurgeRuleAlertsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPurgeRuleAlertsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPurgeRuleAlertsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPurgeRuleAlertsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPurgeRuleAlertsOK creates a PurgeRuleAlertsOK with default headers values
func NewPurgeRuleAlertsOK() *PurgeRuleAlertsOK {
	return &PurgeRuleAlertsOK{}
}

/*PurgeRuleAlertsOK handles this case with default header values.

The history has been purged
*/
type PurgeRuleAlertsOK struct {
}

func (o *PurgeRuleAlertsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1alpha1/monitoring/rules/{group}/{version}/{kind}/{namespace}/{resource}/{name}/purge][%d] purgeRuleAlertsOK ", 200)
}

func (o *PurgeRuleAlertsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPurgeRuleAlertsBadRequest creates a PurgeRuleAlertsBadRequest with default headers values
func NewPurgeRuleAlertsBadRequest() *PurgeRuleAlertsBadRequest {
	return &PurgeRuleAlertsBadRequest{}
}

/*PurgeRuleAlertsBadRequest handles this case with default header values.

Validation error of supplied parameters/body
*/
type PurgeRuleAlertsBadRequest struct {
	Payload *models.ValidationError
}

func (o *PurgeRuleAlertsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1alpha1/monitoring/rules/{group}/{version}/{kind}/{namespace}/{resource}/{name}/purge][%d] purgeRuleAlertsBadRequest  %+v", 400, o.Payload)
}

func (o *PurgeRuleAlertsBadRequest) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *PurgeRuleAlertsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPurgeRuleAlertsUnauthorized creates a PurgeRuleAlertsUnauthorized with default headers values
func NewPurgeRuleAlertsUnauthorized() *PurgeRuleAlertsUnauthorized {
	return &PurgeRuleAlertsUnauthorized{}
}

/*PurgeRuleAlertsUnauthorized handles this case with default header values.

If not authenticated
*/
type PurgeRuleAlertsUnauthorized struct {
}

func (o *PurgeRuleAlertsUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1alpha1/monitoring/rules/{group}/{version}/{kind}/{namespace}/{resource}/{name}/purge][%d] purgeRuleAlertsUnauthorized ", 401)
}

func (o *PurgeRuleAlertsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPurgeRuleAlertsForbidden creates a PurgeRuleAlertsForbidden with default headers values
func NewPurgeRuleAlertsForbidden() *PurgeRuleAlertsForbidden {
	return &PurgeRuleAlertsForbidden{}
}

/*PurgeRuleAlertsForbidden handles this case with default header values.

If authenticated but not authorized
*/
type PurgeRuleAlertsForbidden struct {
}

func (o *PurgeRuleAlertsForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1alpha1/monitoring/rules/{group}/{version}/{kind}/{namespace}/{resource}/{name}/purge][%d] purgeRuleAlertsForbidden ", 403)
}

func (o *PurgeRuleAlertsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPurgeRuleAlertsInternalServerError creates a PurgeRuleAlertsInternalServerError with default headers values
func NewPurgeRuleAlertsInternalServerError() *PurgeRuleAlertsInternalServerError {
	return &PurgeRuleAlertsInternalServerError{}
}

/*PurgeRuleAlertsInternalServerError handles this case with default header values.

A generic API error containing the cause of the error
*/
type PurgeRuleAlertsInternalServerError struct {
	Payload *models.ApiserverError
}

func (o *PurgeRuleAlertsInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1alpha1/monitoring/rules/{group}/{version}/{kind}/{namespace}/{resource}/{name}/purge][%d] purgeRuleAlertsInternalServerError  %+v", 500, o.Payload)
}

func (o *PurgeRuleAlertsInternalServerError) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *PurgeRuleAlertsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
