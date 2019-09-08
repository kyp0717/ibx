// Code generated by go-swagger; DO NOT EDIT.

package market_data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetMarketdataSnapshotParams creates a new GetMarketdataSnapshotParams object
// with the default values initialized.
func NewGetMarketdataSnapshotParams() *GetMarketdataSnapshotParams {
	var ()
	return &GetMarketdataSnapshotParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMarketdataSnapshotParamsWithTimeout creates a new GetMarketdataSnapshotParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMarketdataSnapshotParamsWithTimeout(timeout time.Duration) *GetMarketdataSnapshotParams {
	var ()
	return &GetMarketdataSnapshotParams{

		timeout: timeout,
	}
}

// NewGetMarketdataSnapshotParamsWithContext creates a new GetMarketdataSnapshotParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMarketdataSnapshotParamsWithContext(ctx context.Context) *GetMarketdataSnapshotParams {
	var ()
	return &GetMarketdataSnapshotParams{

		Context: ctx,
	}
}

// NewGetMarketdataSnapshotParamsWithHTTPClient creates a new GetMarketdataSnapshotParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMarketdataSnapshotParamsWithHTTPClient(client *http.Client) *GetMarketdataSnapshotParams {
	var ()
	return &GetMarketdataSnapshotParams{
		HTTPClient: client,
	}
}

/*GetMarketdataSnapshotParams contains all the parameters to send to the API endpoint
for the get iserver marketdata snapshot operation typically these are written to a http.Request
*/
type GetMarketdataSnapshotParams struct {

	/*Conids
	  list of conids separated by comma

	*/
	Conids string
	/*Fields
	  list of fields separated by comma

	*/
	Fields *string
	/*Since
	  time period since which updates are required. uses epoch time with milliseconds.

	*/
	Since *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get iserver marketdata snapshot params
func (o *GetMarketdataSnapshotParams) WithTimeout(timeout time.Duration) *GetMarketdataSnapshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get iserver marketdata snapshot params
func (o *GetMarketdataSnapshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get iserver marketdata snapshot params
func (o *GetMarketdataSnapshotParams) WithContext(ctx context.Context) *GetMarketdataSnapshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get iserver marketdata snapshot params
func (o *GetMarketdataSnapshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get iserver marketdata snapshot params
func (o *GetMarketdataSnapshotParams) WithHTTPClient(client *http.Client) *GetMarketdataSnapshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get iserver marketdata snapshot params
func (o *GetMarketdataSnapshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConids adds the conids to the get iserver marketdata snapshot params
func (o *GetMarketdataSnapshotParams) WithConids(conids string) *GetMarketdataSnapshotParams {
	o.SetConids(conids)
	return o
}

// SetConids adds the conids to the get iserver marketdata snapshot params
func (o *GetMarketdataSnapshotParams) SetConids(conids string) {
	o.Conids = conids
}

// WithFields adds the fields to the get iserver marketdata snapshot params
func (o *GetMarketdataSnapshotParams) WithFields(fields *string) *GetMarketdataSnapshotParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get iserver marketdata snapshot params
func (o *GetMarketdataSnapshotParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithSince adds the since to the get iserver marketdata snapshot params
func (o *GetMarketdataSnapshotParams) WithSince(since *int64) *GetMarketdataSnapshotParams {
	o.SetSince(since)
	return o
}

// SetSince adds the since to the get iserver marketdata snapshot params
func (o *GetMarketdataSnapshotParams) SetSince(since *int64) {
	o.Since = since
}

// WriteToRequest writes these params to a swagger request
func (o *GetMarketdataSnapshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param conids
	qrConids := o.Conids
	qConids := qrConids
	if qConids != "" {
		if err := r.SetQueryParam("conids", qConids); err != nil {
			return err
		}
	}

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	if o.Since != nil {

		// query param since
		var qrSince int64
		if o.Since != nil {
			qrSince = *o.Since
		}
		qSince := swag.FormatInt64(qrSince)
		if qSince != "" {
			if err := r.SetQueryParam("since", qSince); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
