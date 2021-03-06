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

// DeleteResourceRulesReader is a Reader for the DeleteResourceRules structure.
type DeleteResourceRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteResourceRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteResourceRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteResourceRulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteResourceRulesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteResourceRulesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteResourceRulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteResourceRulesOK creates a DeleteResourceRulesOK with default headers values
func NewDeleteResourceRulesOK() *DeleteResourceRulesOK {
	return &DeleteResourceRulesOK{}
}

/*DeleteResourceRulesOK handles this case with default header values.

The rules have been deleted
*/
type DeleteResourceRulesOK struct {
}

func (o *DeleteResourceRulesOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1alpha1/monitoring/rules/{group}/{version}/{kind}/{namespace}/{resource}][%d] deleteResourceRulesOK ", 200)
}

func (o *DeleteResourceRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteResourceRulesBadRequest creates a DeleteResourceRulesBadRequest with default headers values
func NewDeleteResourceRulesBadRequest() *DeleteResourceRulesBadRequest {
	return &DeleteResourceRulesBadRequest{}
}

/*DeleteResourceRulesBadRequest handles this case with default header values.

Validation error of supplied parameters/body
*/
type DeleteResourceRulesBadRequest struct {
	Payload *models.ValidationError
}

func (o *DeleteResourceRulesBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1alpha1/monitoring/rules/{group}/{version}/{kind}/{namespace}/{resource}][%d] deleteResourceRulesBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteResourceRulesBadRequest) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *DeleteResourceRulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteResourceRulesUnauthorized creates a DeleteResourceRulesUnauthorized with default headers values
func NewDeleteResourceRulesUnauthorized() *DeleteResourceRulesUnauthorized {
	return &DeleteResourceRulesUnauthorized{}
}

/*DeleteResourceRulesUnauthorized handles this case with default header values.

If not authenticated
*/
type DeleteResourceRulesUnauthorized struct {
}

func (o *DeleteResourceRulesUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1alpha1/monitoring/rules/{group}/{version}/{kind}/{namespace}/{resource}][%d] deleteResourceRulesUnauthorized ", 401)
}

func (o *DeleteResourceRulesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteResourceRulesForbidden creates a DeleteResourceRulesForbidden with default headers values
func NewDeleteResourceRulesForbidden() *DeleteResourceRulesForbidden {
	return &DeleteResourceRulesForbidden{}
}

/*DeleteResourceRulesForbidden handles this case with default header values.

If authenticated but not authorized
*/
type DeleteResourceRulesForbidden struct {
}

func (o *DeleteResourceRulesForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1alpha1/monitoring/rules/{group}/{version}/{kind}/{namespace}/{resource}][%d] deleteResourceRulesForbidden ", 403)
}

func (o *DeleteResourceRulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteResourceRulesInternalServerError creates a DeleteResourceRulesInternalServerError with default headers values
func NewDeleteResourceRulesInternalServerError() *DeleteResourceRulesInternalServerError {
	return &DeleteResourceRulesInternalServerError{}
}

/*DeleteResourceRulesInternalServerError handles this case with default header values.

A generic API error containing the cause of the error
*/
type DeleteResourceRulesInternalServerError struct {
	Payload *models.ApiserverError
}

func (o *DeleteResourceRulesInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1alpha1/monitoring/rules/{group}/{version}/{kind}/{namespace}/{resource}][%d] deleteResourceRulesInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteResourceRulesInternalServerError) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *DeleteResourceRulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
