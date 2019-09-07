// Code generated by go-swagger; DO NOT EDIT.

package f_y_i

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

// NewGetFyiDeliveryoptionsParams creates a new GetFyiDeliveryoptionsParams object
// with the default values initialized.
func NewGetFyiDeliveryoptionsParams() *GetFyiDeliveryoptionsParams {

	return &GetFyiDeliveryoptionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFyiDeliveryoptionsParamsWithTimeout creates a new GetFyiDeliveryoptionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFyiDeliveryoptionsParamsWithTimeout(timeout time.Duration) *GetFyiDeliveryoptionsParams {

	return &GetFyiDeliveryoptionsParams{

		timeout: timeout,
	}
}

// NewGetFyiDeliveryoptionsParamsWithContext creates a new GetFyiDeliveryoptionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFyiDeliveryoptionsParamsWithContext(ctx context.Context) *GetFyiDeliveryoptionsParams {

	return &GetFyiDeliveryoptionsParams{

		Context: ctx,
	}
}

// NewGetFyiDeliveryoptionsParamsWithHTTPClient creates a new GetFyiDeliveryoptionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFyiDeliveryoptionsParamsWithHTTPClient(client *http.Client) *GetFyiDeliveryoptionsParams {

	return &GetFyiDeliveryoptionsParams{
		HTTPClient: client,
	}
}

/*GetFyiDeliveryoptionsParams contains all the parameters to send to the API endpoint
for the get fyi deliveryoptions operation typically these are written to a http.Request
*/
type GetFyiDeliveryoptionsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get fyi deliveryoptions params
func (o *GetFyiDeliveryoptionsParams) WithTimeout(timeout time.Duration) *GetFyiDeliveryoptionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get fyi deliveryoptions params
func (o *GetFyiDeliveryoptionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get fyi deliveryoptions params
func (o *GetFyiDeliveryoptionsParams) WithContext(ctx context.Context) *GetFyiDeliveryoptionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get fyi deliveryoptions params
func (o *GetFyiDeliveryoptionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get fyi deliveryoptions params
func (o *GetFyiDeliveryoptionsParams) WithHTTPClient(client *http.Client) *GetFyiDeliveryoptionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get fyi deliveryoptions params
func (o *GetFyiDeliveryoptionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetFyiDeliveryoptionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
