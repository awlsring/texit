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

// HeadscaleServiceDisableRouteReader is a Reader for the HeadscaleServiceDisableRoute structure.
type HeadscaleServiceDisableRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeadscaleServiceDisableRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHeadscaleServiceDisableRouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHeadscaleServiceDisableRouteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHeadscaleServiceDisableRouteOK creates a HeadscaleServiceDisableRouteOK with default headers values
func NewHeadscaleServiceDisableRouteOK() *HeadscaleServiceDisableRouteOK {
	return &HeadscaleServiceDisableRouteOK{}
}

/*
HeadscaleServiceDisableRouteOK describes a response with status code 200, with default header values.

A successful response.
*/
type HeadscaleServiceDisableRouteOK struct {
	Payload models.V1DisableRouteResponse
}

// IsSuccess returns true when this headscale service disable route o k response has a 2xx status code
func (o *HeadscaleServiceDisableRouteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this headscale service disable route o k response has a 3xx status code
func (o *HeadscaleServiceDisableRouteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this headscale service disable route o k response has a 4xx status code
func (o *HeadscaleServiceDisableRouteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this headscale service disable route o k response has a 5xx status code
func (o *HeadscaleServiceDisableRouteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this headscale service disable route o k response a status code equal to that given
func (o *HeadscaleServiceDisableRouteOK) IsCode(code int) bool {
	return code == 200
}

func (o *HeadscaleServiceDisableRouteOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/routes/{routeId}/disable][%d] headscaleServiceDisableRouteOK  %+v", 200, o.Payload)
}

func (o *HeadscaleServiceDisableRouteOK) String() string {
	return fmt.Sprintf("[POST /api/v1/routes/{routeId}/disable][%d] headscaleServiceDisableRouteOK  %+v", 200, o.Payload)
}

func (o *HeadscaleServiceDisableRouteOK) GetPayload() models.V1DisableRouteResponse {
	return o.Payload
}

func (o *HeadscaleServiceDisableRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHeadscaleServiceDisableRouteDefault creates a HeadscaleServiceDisableRouteDefault with default headers values
func NewHeadscaleServiceDisableRouteDefault(code int) *HeadscaleServiceDisableRouteDefault {
	return &HeadscaleServiceDisableRouteDefault{
		_statusCode: code,
	}
}

/*
HeadscaleServiceDisableRouteDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HeadscaleServiceDisableRouteDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the headscale service disable route default response
func (o *HeadscaleServiceDisableRouteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this headscale service disable route default response has a 2xx status code
func (o *HeadscaleServiceDisableRouteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this headscale service disable route default response has a 3xx status code
func (o *HeadscaleServiceDisableRouteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this headscale service disable route default response has a 4xx status code
func (o *HeadscaleServiceDisableRouteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this headscale service disable route default response has a 5xx status code
func (o *HeadscaleServiceDisableRouteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this headscale service disable route default response a status code equal to that given
func (o *HeadscaleServiceDisableRouteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *HeadscaleServiceDisableRouteDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/routes/{routeId}/disable][%d] HeadscaleService_DisableRoute default  %+v", o._statusCode, o.Payload)
}

func (o *HeadscaleServiceDisableRouteDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/routes/{routeId}/disable][%d] HeadscaleService_DisableRoute default  %+v", o._statusCode, o.Payload)
}

func (o *HeadscaleServiceDisableRouteDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *HeadscaleServiceDisableRouteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
