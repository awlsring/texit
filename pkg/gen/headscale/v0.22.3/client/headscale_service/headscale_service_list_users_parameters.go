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

// NewHeadscaleServiceListUsersParams creates a new HeadscaleServiceListUsersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHeadscaleServiceListUsersParams() *HeadscaleServiceListUsersParams {
	return &HeadscaleServiceListUsersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHeadscaleServiceListUsersParamsWithTimeout creates a new HeadscaleServiceListUsersParams object
// with the ability to set a timeout on a request.
func NewHeadscaleServiceListUsersParamsWithTimeout(timeout time.Duration) *HeadscaleServiceListUsersParams {
	return &HeadscaleServiceListUsersParams{
		timeout: timeout,
	}
}

// NewHeadscaleServiceListUsersParamsWithContext creates a new HeadscaleServiceListUsersParams object
// with the ability to set a context for a request.
func NewHeadscaleServiceListUsersParamsWithContext(ctx context.Context) *HeadscaleServiceListUsersParams {
	return &HeadscaleServiceListUsersParams{
		Context: ctx,
	}
}

// NewHeadscaleServiceListUsersParamsWithHTTPClient creates a new HeadscaleServiceListUsersParams object
// with the ability to set a custom HTTPClient for a request.
func NewHeadscaleServiceListUsersParamsWithHTTPClient(client *http.Client) *HeadscaleServiceListUsersParams {
	return &HeadscaleServiceListUsersParams{
		HTTPClient: client,
	}
}

/*
HeadscaleServiceListUsersParams contains all the parameters to send to the API endpoint

	for the headscale service list users operation.

	Typically these are written to a http.Request.
*/
type HeadscaleServiceListUsersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the headscale service list users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadscaleServiceListUsersParams) WithDefaults() *HeadscaleServiceListUsersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the headscale service list users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadscaleServiceListUsersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the headscale service list users params
func (o *HeadscaleServiceListUsersParams) WithTimeout(timeout time.Duration) *HeadscaleServiceListUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the headscale service list users params
func (o *HeadscaleServiceListUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the headscale service list users params
func (o *HeadscaleServiceListUsersParams) WithContext(ctx context.Context) *HeadscaleServiceListUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the headscale service list users params
func (o *HeadscaleServiceListUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the headscale service list users params
func (o *HeadscaleServiceListUsersParams) WithHTTPClient(client *http.Client) *HeadscaleServiceListUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the headscale service list users params
func (o *HeadscaleServiceListUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *HeadscaleServiceListUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
