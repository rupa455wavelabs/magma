// Code generated by go-swagger; DO NOT EDIT.

package federated_lte_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetFegLTEParams creates a new GetFegLTEParams object
// with the default values initialized.
func NewGetFegLTEParams() *GetFegLTEParams {

	return &GetFegLTEParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFegLTEParamsWithTimeout creates a new GetFegLTEParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFegLTEParamsWithTimeout(timeout time.Duration) *GetFegLTEParams {

	return &GetFegLTEParams{

		timeout: timeout,
	}
}

// NewGetFegLTEParamsWithContext creates a new GetFegLTEParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFegLTEParamsWithContext(ctx context.Context) *GetFegLTEParams {

	return &GetFegLTEParams{

		Context: ctx,
	}
}

// NewGetFegLTEParamsWithHTTPClient creates a new GetFegLTEParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFegLTEParamsWithHTTPClient(client *http.Client) *GetFegLTEParams {

	return &GetFegLTEParams{
		HTTPClient: client,
	}
}

/*GetFegLTEParams contains all the parameters to send to the API endpoint
for the get feg LTE operation typically these are written to a http.Request
*/
type GetFegLTEParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get feg LTE params
func (o *GetFegLTEParams) WithTimeout(timeout time.Duration) *GetFegLTEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get feg LTE params
func (o *GetFegLTEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get feg LTE params
func (o *GetFegLTEParams) WithContext(ctx context.Context) *GetFegLTEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get feg LTE params
func (o *GetFegLTEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get feg LTE params
func (o *GetFegLTEParams) WithHTTPClient(client *http.Client) *GetFegLTEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get feg LTE params
func (o *GetFegLTEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetFegLTEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}