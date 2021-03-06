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

// RemoveConfigReader is a Reader for the RemoveConfig structure.
type RemoveConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRemoveConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRemoveConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRemoveConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveConfigOK creates a RemoveConfigOK with default headers values
func NewRemoveConfigOK() *RemoveConfigOK {
	return &RemoveConfigOK{}
}

/*RemoveConfigOK handles this case with default header values.

Contains the former config definition
*/
type RemoveConfigOK struct {
	Payload *models.V1Config
}

func (o *RemoveConfigOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1alpha1/configs/{config}][%d] removeConfigOK  %+v", 200, o.Payload)
}

func (o *RemoveConfigOK) GetPayload() *models.V1Config {
	return o.Payload
}

func (o *RemoveConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Config)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveConfigUnauthorized creates a RemoveConfigUnauthorized with default headers values
func NewRemoveConfigUnauthorized() *RemoveConfigUnauthorized {
	return &RemoveConfigUnauthorized{}
}

/*RemoveConfigUnauthorized handles this case with default header values.

If not authenticated
*/
type RemoveConfigUnauthorized struct {
}

func (o *RemoveConfigUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1alpha1/configs/{config}][%d] removeConfigUnauthorized ", 401)
}

func (o *RemoveConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveConfigForbidden creates a RemoveConfigForbidden with default headers values
func NewRemoveConfigForbidden() *RemoveConfigForbidden {
	return &RemoveConfigForbidden{}
}

/*RemoveConfigForbidden handles this case with default header values.

If authenticated but not authorized
*/
type RemoveConfigForbidden struct {
}

func (o *RemoveConfigForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1alpha1/configs/{config}][%d] removeConfigForbidden ", 403)
}

func (o *RemoveConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveConfigInternalServerError creates a RemoveConfigInternalServerError with default headers values
func NewRemoveConfigInternalServerError() *RemoveConfigInternalServerError {
	return &RemoveConfigInternalServerError{}
}

/*RemoveConfigInternalServerError handles this case with default header values.

A generic API error containing the cause of the error
*/
type RemoveConfigInternalServerError struct {
	Payload *models.ApiserverError
}

func (o *RemoveConfigInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1alpha1/configs/{config}][%d] removeConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *RemoveConfigInternalServerError) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *RemoveConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
