// Code generated by go-swagger; DO NOT EDIT.

package headscale_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/headscale/v0.22.3/models"
)

// HeadscaleServiceDeleteRouteReader is a Reader for the HeadscaleServiceDeleteRoute structure.
type HeadscaleServiceDeleteRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeadscaleServiceDeleteRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHeadscaleServiceDeleteRouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHeadscaleServiceDeleteRouteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHeadscaleServiceDeleteRouteOK creates a HeadscaleServiceDeleteRouteOK with default headers values
func NewHeadscaleServiceDeleteRouteOK() *HeadscaleServiceDeleteRouteOK {
	return &HeadscaleServiceDeleteRouteOK{}
}

/*
HeadscaleServiceDeleteRouteOK describes a response with status code 200, with default header values.

A successful response.
*/
type HeadscaleServiceDeleteRouteOK struct {
	Payload models.V1DeleteRouteResponse
}

// IsSuccess returns true when this headscale service delete route o k response has a 2xx status code
func (o *HeadscaleServiceDeleteRouteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this headscale service delete route o k response has a 3xx status code
func (o *HeadscaleServiceDeleteRouteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this headscale service delete route o k response has a 4xx status code
func (o *HeadscaleServiceDeleteRouteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this headscale service delete route o k response has a 5xx status code
func (o *HeadscaleServiceDeleteRouteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this headscale service delete route o k response a status code equal to that given
func (o *HeadscaleServiceDeleteRouteOK) IsCode(code int) bool {
	return code == 200
}

func (o *HeadscaleServiceDeleteRouteOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/routes/{routeId}][%d] headscaleServiceDeleteRouteOK  %+v", 200, o.Payload)
}

func (o *HeadscaleServiceDeleteRouteOK) String() string {
	return fmt.Sprintf("[DELETE /api/v1/routes/{routeId}][%d] headscaleServiceDeleteRouteOK  %+v", 200, o.Payload)
}

func (o *HeadscaleServiceDeleteRouteOK) GetPayload() models.V1DeleteRouteResponse {
	return o.Payload
}

func (o *HeadscaleServiceDeleteRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHeadscaleServiceDeleteRouteDefault creates a HeadscaleServiceDeleteRouteDefault with default headers values
func NewHeadscaleServiceDeleteRouteDefault(code int) *HeadscaleServiceDeleteRouteDefault {
	return &HeadscaleServiceDeleteRouteDefault{
		_statusCode: code,
	}
}

/*
HeadscaleServiceDeleteRouteDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HeadscaleServiceDeleteRouteDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the headscale service delete route default response
func (o *HeadscaleServiceDeleteRouteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this headscale service delete route default response has a 2xx status code
func (o *HeadscaleServiceDeleteRouteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this headscale service delete route default response has a 3xx status code
func (o *HeadscaleServiceDeleteRouteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this headscale service delete route default response has a 4xx status code
func (o *HeadscaleServiceDeleteRouteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this headscale service delete route default response has a 5xx status code
func (o *HeadscaleServiceDeleteRouteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this headscale service delete route default response a status code equal to that given
func (o *HeadscaleServiceDeleteRouteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *HeadscaleServiceDeleteRouteDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/routes/{routeId}][%d] HeadscaleService_DeleteRoute default  %+v", o._statusCode, o.Payload)
}

func (o *HeadscaleServiceDeleteRouteDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v1/routes/{routeId}][%d] HeadscaleService_DeleteRoute default  %+v", o._statusCode, o.Payload)
}

func (o *HeadscaleServiceDeleteRouteDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *HeadscaleServiceDeleteRouteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
