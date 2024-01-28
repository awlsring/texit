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

// HeadscaleServiceGetUserReader is a Reader for the HeadscaleServiceGetUser structure.
type HeadscaleServiceGetUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeadscaleServiceGetUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHeadscaleServiceGetUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHeadscaleServiceGetUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHeadscaleServiceGetUserOK creates a HeadscaleServiceGetUserOK with default headers values
func NewHeadscaleServiceGetUserOK() *HeadscaleServiceGetUserOK {
	return &HeadscaleServiceGetUserOK{}
}

/*
HeadscaleServiceGetUserOK describes a response with status code 200, with default header values.

A successful response.
*/
type HeadscaleServiceGetUserOK struct {
	Payload *models.V1GetUserResponse
}

// IsSuccess returns true when this headscale service get user o k response has a 2xx status code
func (o *HeadscaleServiceGetUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this headscale service get user o k response has a 3xx status code
func (o *HeadscaleServiceGetUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this headscale service get user o k response has a 4xx status code
func (o *HeadscaleServiceGetUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this headscale service get user o k response has a 5xx status code
func (o *HeadscaleServiceGetUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this headscale service get user o k response a status code equal to that given
func (o *HeadscaleServiceGetUserOK) IsCode(code int) bool {
	return code == 200
}

func (o *HeadscaleServiceGetUserOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/{name}][%d] headscaleServiceGetUserOK  %+v", 200, o.Payload)
}

func (o *HeadscaleServiceGetUserOK) String() string {
	return fmt.Sprintf("[GET /api/v1/user/{name}][%d] headscaleServiceGetUserOK  %+v", 200, o.Payload)
}

func (o *HeadscaleServiceGetUserOK) GetPayload() *models.V1GetUserResponse {
	return o.Payload
}

func (o *HeadscaleServiceGetUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GetUserResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHeadscaleServiceGetUserDefault creates a HeadscaleServiceGetUserDefault with default headers values
func NewHeadscaleServiceGetUserDefault(code int) *HeadscaleServiceGetUserDefault {
	return &HeadscaleServiceGetUserDefault{
		_statusCode: code,
	}
}

/*
HeadscaleServiceGetUserDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HeadscaleServiceGetUserDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the headscale service get user default response
func (o *HeadscaleServiceGetUserDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this headscale service get user default response has a 2xx status code
func (o *HeadscaleServiceGetUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this headscale service get user default response has a 3xx status code
func (o *HeadscaleServiceGetUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this headscale service get user default response has a 4xx status code
func (o *HeadscaleServiceGetUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this headscale service get user default response has a 5xx status code
func (o *HeadscaleServiceGetUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this headscale service get user default response a status code equal to that given
func (o *HeadscaleServiceGetUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *HeadscaleServiceGetUserDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/{name}][%d] HeadscaleService_GetUser default  %+v", o._statusCode, o.Payload)
}

func (o *HeadscaleServiceGetUserDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/user/{name}][%d] HeadscaleService_GetUser default  %+v", o._statusCode, o.Payload)
}

func (o *HeadscaleServiceGetUserDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *HeadscaleServiceGetUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
