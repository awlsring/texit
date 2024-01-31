// Code generated by go-swagger; DO NOT EDIT.

package headscale_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/awlsring/texit/pkg/gen/headscale/v0.22.3/models"
)

// HeadscaleServiceCreatePreAuthKeyReader is a Reader for the HeadscaleServiceCreatePreAuthKey structure.
type HeadscaleServiceCreatePreAuthKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeadscaleServiceCreatePreAuthKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHeadscaleServiceCreatePreAuthKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHeadscaleServiceCreatePreAuthKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHeadscaleServiceCreatePreAuthKeyOK creates a HeadscaleServiceCreatePreAuthKeyOK with default headers values
func NewHeadscaleServiceCreatePreAuthKeyOK() *HeadscaleServiceCreatePreAuthKeyOK {
	return &HeadscaleServiceCreatePreAuthKeyOK{}
}

/*
HeadscaleServiceCreatePreAuthKeyOK describes a response with status code 200, with default header values.

A successful response.
*/
type HeadscaleServiceCreatePreAuthKeyOK struct {
	Payload *models.V1CreatePreAuthKeyResponse
}

// IsSuccess returns true when this headscale service create pre auth key o k response has a 2xx status code
func (o *HeadscaleServiceCreatePreAuthKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this headscale service create pre auth key o k response has a 3xx status code
func (o *HeadscaleServiceCreatePreAuthKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this headscale service create pre auth key o k response has a 4xx status code
func (o *HeadscaleServiceCreatePreAuthKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this headscale service create pre auth key o k response has a 5xx status code
func (o *HeadscaleServiceCreatePreAuthKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this headscale service create pre auth key o k response a status code equal to that given
func (o *HeadscaleServiceCreatePreAuthKeyOK) IsCode(code int) bool {
	return code == 200
}

func (o *HeadscaleServiceCreatePreAuthKeyOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/preauthkey][%d] headscaleServiceCreatePreAuthKeyOK  %+v", 200, o.Payload)
}

func (o *HeadscaleServiceCreatePreAuthKeyOK) String() string {
	return fmt.Sprintf("[POST /api/v1/preauthkey][%d] headscaleServiceCreatePreAuthKeyOK  %+v", 200, o.Payload)
}

func (o *HeadscaleServiceCreatePreAuthKeyOK) GetPayload() *models.V1CreatePreAuthKeyResponse {
	return o.Payload
}

func (o *HeadscaleServiceCreatePreAuthKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1CreatePreAuthKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHeadscaleServiceCreatePreAuthKeyDefault creates a HeadscaleServiceCreatePreAuthKeyDefault with default headers values
func NewHeadscaleServiceCreatePreAuthKeyDefault(code int) *HeadscaleServiceCreatePreAuthKeyDefault {
	return &HeadscaleServiceCreatePreAuthKeyDefault{
		_statusCode: code,
	}
}

/*
HeadscaleServiceCreatePreAuthKeyDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HeadscaleServiceCreatePreAuthKeyDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the headscale service create pre auth key default response
func (o *HeadscaleServiceCreatePreAuthKeyDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this headscale service create pre auth key default response has a 2xx status code
func (o *HeadscaleServiceCreatePreAuthKeyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this headscale service create pre auth key default response has a 3xx status code
func (o *HeadscaleServiceCreatePreAuthKeyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this headscale service create pre auth key default response has a 4xx status code
func (o *HeadscaleServiceCreatePreAuthKeyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this headscale service create pre auth key default response has a 5xx status code
func (o *HeadscaleServiceCreatePreAuthKeyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this headscale service create pre auth key default response a status code equal to that given
func (o *HeadscaleServiceCreatePreAuthKeyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *HeadscaleServiceCreatePreAuthKeyDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/preauthkey][%d] HeadscaleService_CreatePreAuthKey default  %+v", o._statusCode, o.Payload)
}

func (o *HeadscaleServiceCreatePreAuthKeyDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/preauthkey][%d] HeadscaleService_CreatePreAuthKey default  %+v", o._statusCode, o.Payload)
}

func (o *HeadscaleServiceCreatePreAuthKeyDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *HeadscaleServiceCreatePreAuthKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
