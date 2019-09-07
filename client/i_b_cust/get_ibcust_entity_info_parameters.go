// Code generated by go-swagger; DO NOT EDIT.

package i_b_cust

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

// NewGetIbcustEntityInfoParams creates a new GetIbcustEntityInfoParams object
// with the default values initialized.
func NewGetIbcustEntityInfoParams() *GetIbcustEntityInfoParams {

	return &GetIbcustEntityInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIbcustEntityInfoParamsWithTimeout creates a new GetIbcustEntityInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIbcustEntityInfoParamsWithTimeout(timeout time.Duration) *GetIbcustEntityInfoParams {

	return &GetIbcustEntityInfoParams{

		timeout: timeout,
	}
}

// NewGetIbcustEntityInfoParamsWithContext creates a new GetIbcustEntityInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIbcustEntityInfoParamsWithContext(ctx context.Context) *GetIbcustEntityInfoParams {

	return &GetIbcustEntityInfoParams{

		Context: ctx,
	}
}

// NewGetIbcustEntityInfoParamsWithHTTPClient creates a new GetIbcustEntityInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIbcustEntityInfoParamsWithHTTPClient(client *http.Client) *GetIbcustEntityInfoParams {

	return &GetIbcustEntityInfoParams{
		HTTPClient: client,
	}
}

/*GetIbcustEntityInfoParams contains all the parameters to send to the API endpoint
for the get ibcust entity info operation typically these are written to a http.Request
*/
type GetIbcustEntityInfoParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get ibcust entity info params
func (o *GetIbcustEntityInfoParams) WithTimeout(timeout time.Duration) *GetIbcustEntityInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ibcust entity info params
func (o *GetIbcustEntityInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ibcust entity info params
func (o *GetIbcustEntityInfoParams) WithContext(ctx context.Context) *GetIbcustEntityInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ibcust entity info params
func (o *GetIbcustEntityInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ibcust entity info params
func (o *GetIbcustEntityInfoParams) WithHTTPClient(client *http.Client) *GetIbcustEntityInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ibcust entity info params
func (o *GetIbcustEntityInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetIbcustEntityInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
