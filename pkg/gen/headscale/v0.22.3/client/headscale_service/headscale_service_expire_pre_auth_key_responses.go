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

// HeadscaleServiceExpirePreAuthKeyReader is a Reader for the HeadscaleServiceExpirePreAuthKey structure.
type HeadscaleServiceExpirePreAuthKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeadscaleServiceExpirePreAuthKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHeadscaleServiceExpirePreAuthKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHeadscaleServiceExpirePreAuthKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHeadscaleServiceExpirePreAuthKeyOK creates a HeadscaleServiceExpirePreAuthKeyOK with default headers values
func NewHeadscaleServiceExpirePreAuthKeyOK() *HeadscaleServiceExpirePreAuthKeyOK {
	return &HeadscaleServiceExpirePreAuthKeyOK{}
}

/*
HeadscaleServiceExpirePreAuthKeyOK describes a response with status code 200, with default header values.

A successful response.
*/
type HeadscaleServiceExpirePreAuthKeyOK struct {
	Payload models.V1ExpirePreAuthKeyResponse
}

// IsSuccess returns true when this headscale service expire pre auth key o k response has a 2xx status code
func (o *HeadscaleServiceExpirePreAuthKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this headscale service expire pre auth key o k response has a 3xx status code
func (o *HeadscaleServiceExpirePreAuthKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this headscale service expire pre auth key o k response has a 4xx status code
func (o *HeadscaleServiceExpirePreAuthKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this headscale service expire pre auth key o k response has a 5xx status code
func (o *HeadscaleServiceExpirePreAuthKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this headscale service expire pre auth key o k response a status code equal to that given
func (o *HeadscaleServiceExpirePreAuthKeyOK) IsCode(code int) bool {
	return code == 200
}

func (o *HeadscaleServiceExpirePreAuthKeyOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/preauthkey/expire][%d] headscaleServiceExpirePreAuthKeyOK  %+v", 200, o.Payload)
}

func (o *HeadscaleServiceExpirePreAuthKeyOK) String() string {
	return fmt.Sprintf("[POST /api/v1/preauthkey/expire][%d] headscaleServiceExpirePreAuthKeyOK  %+v", 200, o.Payload)
}

func (o *HeadscaleServiceExpirePreAuthKeyOK) GetPayload() models.V1ExpirePreAuthKeyResponse {
	return o.Payload
}

func (o *HeadscaleServiceExpirePreAuthKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHeadscaleServiceExpirePreAuthKeyDefault creates a HeadscaleServiceExpirePreAuthKeyDefault with default headers values
func NewHeadscaleServiceExpirePreAuthKeyDefault(code int) *HeadscaleServiceExpirePreAuthKeyDefault {
	return &HeadscaleServiceExpirePreAuthKeyDefault{
		_statusCode: code,
	}
}

/*
HeadscaleServiceExpirePreAuthKeyDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HeadscaleServiceExpirePreAuthKeyDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the headscale service expire pre auth key default response
func (o *HeadscaleServiceExpirePreAuthKeyDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this headscale service expire pre auth key default response has a 2xx status code
func (o *HeadscaleServiceExpirePreAuthKeyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this headscale service expire pre auth key default response has a 3xx status code
func (o *HeadscaleServiceExpirePreAuthKeyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this headscale service expire pre auth key default response has a 4xx status code
func (o *HeadscaleServiceExpirePreAuthKeyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this headscale service expire pre auth key default response has a 5xx status code
func (o *HeadscaleServiceExpirePreAuthKeyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this headscale service expire pre auth key default response a status code equal to that given
func (o *HeadscaleServiceExpirePreAuthKeyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *HeadscaleServiceExpirePreAuthKeyDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/preauthkey/expire][%d] HeadscaleService_ExpirePreAuthKey default  %+v", o._statusCode, o.Payload)
}

func (o *HeadscaleServiceExpirePreAuthKeyDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/preauthkey/expire][%d] HeadscaleService_ExpirePreAuthKey default  %+v", o._statusCode, o.Payload)
}

func (o *HeadscaleServiceExpirePreAuthKeyDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *HeadscaleServiceExpirePreAuthKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}