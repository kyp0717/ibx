// Code generated by go-swagger; DO NOT EDIT.

package pn_l

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new pn l API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for pn l API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetAccountPnlPartitioned pns l for the selected account

Returns an object containing PnLfor the selected account and its models (if any).
*/
func (a *Client) GetAccountPnlPartitioned(params *GetAccountPnlPartitionedParams) (*GetAccountPnlPartitionedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccountPnlPartitionedParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAccountPnlPartitioned",
		Method:             "GET",
		PathPattern:        "/iserver/account/pnl/partitioned",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccountPnlPartitionedReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAccountPnlPartitionedOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAccountPnlPartitioned: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
