// Code generated by go-swagger; DO NOT EDIT.

package order

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

// NewGetAccountOrdersParams creates a new GetAccountOrdersParams object
// with the default values initialized.
func NewGetAccountOrdersParams() *GetAccountOrdersParams {

	return &GetAccountOrdersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountOrdersParamsWithTimeout creates a new GetAccountOrdersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccountOrdersParamsWithTimeout(timeout time.Duration) *GetAccountOrdersParams {

	return &GetAccountOrdersParams{

		timeout: timeout,
	}
}

// NewGetAccountOrdersParamsWithContext creates a new GetAccountOrdersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAccountOrdersParamsWithContext(ctx context.Context) *GetAccountOrdersParams {

	return &GetAccountOrdersParams{

		Context: ctx,
	}
}

// NewGetAccountOrdersParamsWithHTTPClient creates a new GetAccountOrdersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAccountOrdersParamsWithHTTPClient(client *http.Client) *GetAccountOrdersParams {

	return &GetAccountOrdersParams{
		HTTPClient: client,
	}
}

/*GetAccountOrdersParams contains all the parameters to send to the API endpoint
for the get iserver account orders operation typically these are written to a http.Request
*/
type GetAccountOrdersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get iserver account orders params
func (o *GetAccountOrdersParams) WithTimeout(timeout time.Duration) *GetAccountOrdersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get iserver account orders params
func (o *GetAccountOrdersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get iserver account orders params
func (o *GetAccountOrdersParams) WithContext(ctx context.Context) *GetAccountOrdersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get iserver account orders params
func (o *GetAccountOrdersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get iserver account orders params
func (o *GetAccountOrdersParams) WithHTTPClient(client *http.Client) *GetAccountOrdersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get iserver account orders params
func (o *GetAccountOrdersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountOrdersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
