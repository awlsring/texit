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

// HeadscaleServiceDebugCreateMachineReader is a Reader for the HeadscaleServiceDebugCreateMachine structure.
type HeadscaleServiceDebugCreateMachineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeadscaleServiceDebugCreateMachineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHeadscaleServiceDebugCreateMachineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHeadscaleServiceDebugCreateMachineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHeadscaleServiceDebugCreateMachineOK creates a HeadscaleServiceDebugCreateMachineOK with default headers values
func NewHeadscaleServiceDebugCreateMachineOK() *HeadscaleServiceDebugCreateMachineOK {
	return &HeadscaleServiceDebugCreateMachineOK{}
}

/*
HeadscaleServiceDebugCreateMachineOK describes a response with status code 200, with default header values.

A successful response.
*/
type HeadscaleServiceDebugCreateMachineOK struct {
	Payload *models.V1DebugCreateMachineResponse
}

// IsSuccess returns true when this headscale service debug create machine o k response has a 2xx status code
func (o *HeadscaleServiceDebugCreateMachineOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this headscale service debug create machine o k response has a 3xx status code
func (o *HeadscaleServiceDebugCreateMachineOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this headscale service debug create machine o k response has a 4xx status code
func (o *HeadscaleServiceDebugCreateMachineOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this headscale service debug create machine o k response has a 5xx status code
func (o *HeadscaleServiceDebugCreateMachineOK) IsServerError() bool {
	return false
}

// IsCode returns true when this headscale service debug create machine o k response a status code equal to that given
func (o *HeadscaleServiceDebugCreateMachineOK) IsCode(code int) bool {
	return code == 200
}

func (o *HeadscaleServiceDebugCreateMachineOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/debug/machine][%d] headscaleServiceDebugCreateMachineOK  %+v", 200, o.Payload)
}

func (o *HeadscaleServiceDebugCreateMachineOK) String() string {
	return fmt.Sprintf("[POST /api/v1/debug/machine][%d] headscaleServiceDebugCreateMachineOK  %+v", 200, o.Payload)
}

func (o *HeadscaleServiceDebugCreateMachineOK) GetPayload() *models.V1DebugCreateMachineResponse {
	return o.Payload
}

func (o *HeadscaleServiceDebugCreateMachineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1DebugCreateMachineResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHeadscaleServiceDebugCreateMachineDefault creates a HeadscaleServiceDebugCreateMachineDefault with default headers values
func NewHeadscaleServiceDebugCreateMachineDefault(code int) *HeadscaleServiceDebugCreateMachineDefault {
	return &HeadscaleServiceDebugCreateMachineDefault{
		_statusCode: code,
	}
}

/*
HeadscaleServiceDebugCreateMachineDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HeadscaleServiceDebugCreateMachineDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the headscale service debug create machine default response
func (o *HeadscaleServiceDebugCreateMachineDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this headscale service debug create machine default response has a 2xx status code
func (o *HeadscaleServiceDebugCreateMachineDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this headscale service debug create machine default response has a 3xx status code
func (o *HeadscaleServiceDebugCreateMachineDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this headscale service debug create machine default response has a 4xx status code
func (o *HeadscaleServiceDebugCreateMachineDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this headscale service debug create machine default response has a 5xx status code
func (o *HeadscaleServiceDebugCreateMachineDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this headscale service debug create machine default response a status code equal to that given
func (o *HeadscaleServiceDebugCreateMachineDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *HeadscaleServiceDebugCreateMachineDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/debug/machine][%d] HeadscaleService_DebugCreateMachine default  %+v", o._statusCode, o.Payload)
}

func (o *HeadscaleServiceDebugCreateMachineDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/debug/machine][%d] HeadscaleService_DebugCreateMachine default  %+v", o._statusCode, o.Payload)
}

func (o *HeadscaleServiceDebugCreateMachineDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *HeadscaleServiceDebugCreateMachineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}