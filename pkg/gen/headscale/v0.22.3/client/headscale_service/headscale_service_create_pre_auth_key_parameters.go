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

	"github.com/awlsring/texit/pkg/gen/headscale/v0.22.3/models"
)

// NewHeadscaleServiceCreatePreAuthKeyParams creates a new HeadscaleServiceCreatePreAuthKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHeadscaleServiceCreatePreAuthKeyParams() *HeadscaleServiceCreatePreAuthKeyParams {
	return &HeadscaleServiceCreatePreAuthKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHeadscaleServiceCreatePreAuthKeyParamsWithTimeout creates a new HeadscaleServiceCreatePreAuthKeyParams object
// with the ability to set a timeout on a request.
func NewHeadscaleServiceCreatePreAuthKeyParamsWithTimeout(timeout time.Duration) *HeadscaleServiceCreatePreAuthKeyParams {
	return &HeadscaleServiceCreatePreAuthKeyParams{
		timeout: timeout,
	}
}

// NewHeadscaleServiceCreatePreAuthKeyParamsWithContext creates a new HeadscaleServiceCreatePreAuthKeyParams object
// with the ability to set a context for a request.
func NewHeadscaleServiceCreatePreAuthKeyParamsWithContext(ctx context.Context) *HeadscaleServiceCreatePreAuthKeyParams {
	return &HeadscaleServiceCreatePreAuthKeyParams{
		Context: ctx,
	}
}

// NewHeadscaleServiceCreatePreAuthKeyParamsWithHTTPClient creates a new HeadscaleServiceCreatePreAuthKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewHeadscaleServiceCreatePreAuthKeyParamsWithHTTPClient(client *http.Client) *HeadscaleServiceCreatePreAuthKeyParams {
	return &HeadscaleServiceCreatePreAuthKeyParams{
		HTTPClient: client,
	}
}

/*
HeadscaleServiceCreatePreAuthKeyParams contains all the parameters to send to the API endpoint

	for the headscale service create pre auth key operation.

	Typically these are written to a http.Request.
*/
type HeadscaleServiceCreatePreAuthKeyParams struct {

	// Body.
	Body *models.V1CreatePreAuthKeyRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the headscale service create pre auth key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadscaleServiceCreatePreAuthKeyParams) WithDefaults() *HeadscaleServiceCreatePreAuthKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the headscale service create pre auth key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadscaleServiceCreatePreAuthKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the headscale service create pre auth key params
func (o *HeadscaleServiceCreatePreAuthKeyParams) WithTimeout(timeout time.Duration) *HeadscaleServiceCreatePreAuthKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the headscale service create pre auth key params
func (o *HeadscaleServiceCreatePreAuthKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the headscale service create pre auth key params
func (o *HeadscaleServiceCreatePreAuthKeyParams) WithContext(ctx context.Context) *HeadscaleServiceCreatePreAuthKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the headscale service create pre auth key params
func (o *HeadscaleServiceCreatePreAuthKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the headscale service create pre auth key params
func (o *HeadscaleServiceCreatePreAuthKeyParams) WithHTTPClient(client *http.Client) *HeadscaleServiceCreatePreAuthKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the headscale service create pre auth key params
func (o *HeadscaleServiceCreatePreAuthKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the headscale service create pre auth key params
func (o *HeadscaleServiceCreatePreAuthKeyParams) WithBody(body *models.V1CreatePreAuthKeyRequest) *HeadscaleServiceCreatePreAuthKeyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the headscale service create pre auth key params
func (o *HeadscaleServiceCreatePreAuthKeyParams) SetBody(body *models.V1CreatePreAuthKeyRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *HeadscaleServiceCreatePreAuthKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
