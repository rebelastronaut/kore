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

// ListClustersReader is a Reader for the ListClusters structure.
type ListClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewListClustersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListClustersOK creates a ListClustersOK with default headers values
func NewListClustersOK() *ListClustersOK {
	return &ListClustersOK{}
}

/*ListClustersOK handles this case with default header values.

Contains the former team definition from the kore
*/
type ListClustersOK struct {
	Payload *apimodels.V1KubernetesList
}

func (o *ListClustersOK) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/teams/{team}/clusters][%d] listClustersOK  %+v", 200, o.Payload)
}

func (o *ListClustersOK) GetPayload() *apimodels.V1KubernetesList {
	return o.Payload
}

func (o *ListClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(apimodels.V1KubernetesList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClustersInternalServerError creates a ListClustersInternalServerError with default headers values
func NewListClustersInternalServerError() *ListClustersInternalServerError {
	return &ListClustersInternalServerError{}
}

/*ListClustersInternalServerError handles this case with default header values.

A generic API error containing the cause of the error
*/
type ListClustersInternalServerError struct {
	Payload *apimodels.ApiserverError
}

func (o *ListClustersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/teams/{team}/clusters][%d] listClustersInternalServerError  %+v", 500, o.Payload)
}

func (o *ListClustersInternalServerError) GetPayload() *apimodels.ApiserverError {
	return o.Payload
}

func (o *ListClustersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(apimodels.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
