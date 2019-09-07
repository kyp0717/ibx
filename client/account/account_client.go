// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new account API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

// Client for account API
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetIserverAccounts brokerages accounts
Returns a list of accounts the user has trading access to,
their respective aliases and the currently selected account.
Note this endpoint must be called before modifying an order or querying open orders.
*/
func (a *Client) GetIserverAccounts(params *GetIserverAccountsParams) (*GetIserverAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIserverAccountsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetIserverAccounts",
		Method:             "GET",
		PathPattern:        "/iserver/accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIserverAccountsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIserverAccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetIserverAccounts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPortfolioAccountIDLedger accounts ledger

Information regarding settled cash, cash balances, etc. in
the account's base currency and any other cash balances hold in other currencies.  /portfolio/accounts or /portfolio/subaccounts must be called prior to this endpoint. The list of supported currencies is available at https://www.interactivebrokers.com/en/index.php?f=3185.
*/
func (a *Client) GetPortfolioAccountIDLedger(params *GetPortfolioAccountIDLedgerParams) (*GetPortfolioAccountIDLedgerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPortfolioAccountIDLedgerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPortfolioAccountIDLedger",
		Method:             "GET",
		PathPattern:        "/portfolio/{accountId}/ledger",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPortfolioAccountIDLedgerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPortfolioAccountIDLedgerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPortfolioAccountIDLedger: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPortfolioAccountIDMeta accounts information

Account information related to account Id /portfolio/accounts or /portfolio/subaccounts must be called prior to this endpoint.
*/
func (a *Client) GetPortfolioAccountIDMeta(params *GetPortfolioAccountIDMetaParams) (*GetPortfolioAccountIDMetaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPortfolioAccountIDMetaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPortfolioAccountIDMeta",
		Method:             "GET",
		PathPattern:        "/portfolio/{accountId}/meta",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPortfolioAccountIDMetaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPortfolioAccountIDMetaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPortfolioAccountIDMeta: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPortfolioAccountIDSummary accounts summary

Returns information about margin, cash balances and other information
related to specified account. See also /portfolio/{accountId}/ledger.
/portfolio/accounts or /portfolio/subaccounts must be called prior to this endpoint.
*/
func (a *Client) GetPortfolioAccountIDSummary(params *GetPortfolioAccountIDSummaryParams) (*GetPortfolioAccountIDSummaryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPortfolioAccountIDSummaryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPortfolioAccountIDSummary",
		Method:             "GET",
		PathPattern:        "/portfolio/{accountId}/summary",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPortfolioAccountIDSummaryReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPortfolioAccountIDSummaryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPortfolioAccountIDSummary: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPortfolioAccounts portfolios accounts

In non-tiered account structures, returns a list of accounts for which the user can view position and account information. This endpoint must be called prior  to calling other /portfolio endpoints for those accounts. For querying a list of accounts  which the user can trade, see /iserver/accounts. For a list of subaccounts in tiered  account structures (e.g. financial advisor or ibroker accounts) see /portfolio/subaccounts.
*/
func (a *Client) GetPortfolioAccounts(params *GetPortfolioAccountsParams) (*GetPortfolioAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPortfolioAccountsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPortfolioAccounts",
		Method:             "GET",
		PathPattern:        "/portfolio/accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPortfolioAccountsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPortfolioAccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPortfolioAccounts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPortfolioSubaccounts lists of sub accounts

Used in tiered account structures (such as financial advisor and ibroker accounts)  to return a list of sub-accounts for which the user can view position and  account-related information. This endpoint must be called prior to calling other  /portfolio endpoints for those subaccounts.  To query a list of accounts the user can trade, see /iserver/accounts.
*/
func (a *Client) GetPortfolioSubaccounts(params *GetPortfolioSubaccountsParams) (*GetPortfolioSubaccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPortfolioSubaccountsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPortfolioSubaccounts",
		Method:             "GET",
		PathPattern:        "/portfolio/subaccounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPortfolioSubaccountsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPortfolioSubaccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPortfolioSubaccounts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostIserverAccount updates currently selected account to the provided account

If an user has multiple accounts, and user wants to get orders, trades, etc. of an account other than currently selected account, then user can update the currently selected account using this API and then can fetch required information for the newly updated account.
*/
func (a *Client) PostIserverAccount(params *PostIserverAccountParams) (*PostIserverAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostIserverAccountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostIserverAccount",
		Method:             "POST",
		PathPattern:        "/iserver/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostIserverAccountReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostIserverAccountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostIserverAccount: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
