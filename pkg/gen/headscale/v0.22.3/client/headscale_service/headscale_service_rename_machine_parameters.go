// Code generated by go-swagger; DO NOT EDIT.

package headscale_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewHeadscaleServiceRenameMachineParams creates a new HeadscaleServiceRenameMachineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHeadscaleServiceRenameMachineParams() *HeadscaleServiceRenameMachineParams {
	return &HeadscaleServiceRenameMachineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHeadscaleServiceRenameMachineParamsWithTimeout creates a new HeadscaleServiceRenameMachineParams object
// with the ability to set a timeout on a request.
func NewHeadscaleServiceRenameMachineParamsWithTimeout(timeout time.Duration) *HeadscaleServiceRenameMachineParams {
	return &HeadscaleServiceRenameMachineParams{
		timeout: timeout,
	}
}

// NewHeadscaleServiceRenameMachineParamsWithContext creates a new HeadscaleServiceRenameMachineParams object
// with the ability to set a context for a request.
func NewHeadscaleServiceRenameMachineParamsWithContext(ctx context.Context) *HeadscaleServiceRenameMachineParams {
	return &HeadscaleServiceRenameMachineParams{
		Context: ctx,
	}
}

// NewHeadscaleServiceRenameMachineParamsWithHTTPClient creates a new HeadscaleServiceRenameMachineParams object
// with the ability to set a custom HTTPClient for a request.
func NewHeadscaleServiceRenameMachineParamsWithHTTPClient(client *http.Client) *HeadscaleServiceRenameMachineParams {
	return &HeadscaleServiceRenameMachineParams{
		HTTPClient: client,
	}
}

/*
HeadscaleServiceRenameMachineParams contains all the parameters to send to the API endpoint

	for the headscale service rename machine operation.

	Typically these are written to a http.Request.
*/
type HeadscaleServiceRenameMachineParams struct {

	// MachineID.
	//
	// Format: uint64
	MachineID string

	// NewName.
	NewName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the headscale service rename machine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadscaleServiceRenameMachineParams) WithDefaults() *HeadscaleServiceRenameMachineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the headscale service rename machine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadscaleServiceRenameMachineParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the headscale service rename machine params
func (o *HeadscaleServiceRenameMachineParams) WithTimeout(timeout time.Duration) *HeadscaleServiceRenameMachineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the headscale service rename machine params
func (o *HeadscaleServiceRenameMachineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the headscale service rename machine params
func (o *HeadscaleServiceRenameMachineParams) WithContext(ctx context.Context) *HeadscaleServiceRenameMachineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the headscale service rename machine params
func (o *HeadscaleServiceRenameMachineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the headscale service rename machine params
func (o *HeadscaleServiceRenameMachineParams) WithHTTPClient(client *http.Client) *HeadscaleServiceRenameMachineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the headscale service rename machine params
func (o *HeadscaleServiceRenameMachineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMachineID adds the machineID to the headscale service rename machine params
func (o *HeadscaleServiceRenameMachineParams) WithMachineID(machineID string) *HeadscaleServiceRenameMachineParams {
	o.SetMachineID(machineID)
	return o
}

// SetMachineID adds the machineId to the headscale service rename machine params
func (o *HeadscaleServiceRenameMachineParams) SetMachineID(machineID string) {
	o.MachineID = machineID
}

// WithNewName adds the newName to the headscale service rename machine params
func (o *HeadscaleServiceRenameMachineParams) WithNewName(newName string) *HeadscaleServiceRenameMachineParams {
	o.SetNewName(newName)
	return o
}

// SetNewName adds the newName to the headscale service rename machine params
func (o *HeadscaleServiceRenameMachineParams) SetNewName(newName string) {
	o.NewName = newName
}

// WriteToRequest writes these params to a swagger request
func (o *HeadscaleServiceRenameMachineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param machineId
	if err := r.SetPathParam("machineId", o.MachineID); err != nil {
		return err
	}

	// path param newName
	if err := r.SetPathParam("newName", o.NewName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}