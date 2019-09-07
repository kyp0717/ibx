// Code generated by go-swagger; DO NOT EDIT.

package portfolio

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

// NewGetPortfolioAccountIDAllocationParams creates a new GetPortfolioAccountIDAllocationParams object
// with the default values initialized.
func NewGetPortfolioAccountIDAllocationParams() *GetPortfolioAccountIDAllocationParams {
	var ()
	return &GetPortfolioAccountIDAllocationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPortfolioAccountIDAllocationParamsWithTimeout creates a new GetPortfolioAccountIDAllocationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPortfolioAccountIDAllocationParamsWithTimeout(timeout time.Duration) *GetPortfolioAccountIDAllocationParams {
	var ()
	return &GetPortfolioAccountIDAllocationParams{

		timeout: timeout,
	}
}

// NewGetPortfolioAccountIDAllocationParamsWithContext creates a new GetPortfolioAccountIDAllocationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPortfolioAccountIDAllocationParamsWithContext(ctx context.Context) *GetPortfolioAccountIDAllocationParams {
	var ()
	return &GetPortfolioAccountIDAllocationParams{

		Context: ctx,
	}
}

// NewGetPortfolioAccountIDAllocationParamsWithHTTPClient creates a new GetPortfolioAccountIDAllocationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPortfolioAccountIDAllocationParamsWithHTTPClient(client *http.Client) *GetPortfolioAccountIDAllocationParams {
	var ()
	return &GetPortfolioAccountIDAllocationParams{
		HTTPClient: client,
	}
}

/*GetPortfolioAccountIDAllocationParams contains all the parameters to send to the API endpoint
for the get portfolio account ID allocation operation typically these are written to a http.Request
*/
type GetPortfolioAccountIDAllocationParams struct {

	/*AccountID
	  account id

	*/
	AccountID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get portfolio account ID allocation params
func (o *GetPortfolioAccountIDAllocationParams) WithTimeout(timeout time.Duration) *GetPortfolioAccountIDAllocationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get portfolio account ID allocation params
func (o *GetPortfolioAccountIDAllocationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get portfolio account ID allocation params
func (o *GetPortfolioAccountIDAllocationParams) WithContext(ctx context.Context) *GetPortfolioAccountIDAllocationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get portfolio account ID allocation params
func (o *GetPortfolioAccountIDAllocationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get portfolio account ID allocation params
func (o *GetPortfolioAccountIDAllocationParams) WithHTTPClient(client *http.Client) *GetPortfolioAccountIDAllocationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get portfolio account ID allocation params
func (o *GetPortfolioAccountIDAllocationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get portfolio account ID allocation params
func (o *GetPortfolioAccountIDAllocationParams) WithAccountID(accountID string) *GetPortfolioAccountIDAllocationParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get portfolio account ID allocation params
func (o *GetPortfolioAccountIDAllocationParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPortfolioAccountIDAllocationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
