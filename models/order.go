// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Order contains all the order related info
// swagger:model order
type Order struct {

	// account id
	Acct string `json:"acct,omitempty"`

	// back-ground color
	BgColor string `json:"bgColor,omitempty"`

	// company name
	CompanyName string `json:"companyName,omitempty"`

	// conid
	Conid int64 `json:"conid,omitempty"`

	// description1
	Description1 string `json:"description1,omitempty"`

	// fg color
	FgColor string `json:"fgColor,omitempty"`

	// filled quantity
	FilledQuantity string `json:"filledQuantity,omitempty"`

	// for example NASDAQ.NMS
	ListingExchange string `json:"listingExchange,omitempty"`

	// order desc
	OrderDesc string `json:"orderDesc,omitempty"`

	// order Id
	OrderID int64 `json:"orderId,omitempty"`

	// User defined string used to identify the order. Value is set using "cOID" field while placing an order.
	OrderRef string `json:"order_ref,omitempty"`

	// for example Limit
	OrigOrderType string `json:"origOrderType,omitempty"`

	// Only exists in child order of bracket
	ParentID int64 `json:"parentId,omitempty"`

	// price
	Price float64 `json:"price,omitempty"`

	// remaining quantity
	RemainingQuantity string `json:"remainingQuantity,omitempty"`

	// for example STK
	SecType string `json:"secType,omitempty"`

	// BUY or SELL
	Side string `json:"side,omitempty"`

	// PendingSubmit - Indicates the order was sent, but confirmation has not been received that it has been received by the destination.
	//                 Occurs most commonly if an exchange is closed.
	// PendingCancel - Indicates that a request has been sent to cancel an order but confirmation has not been received of its cancellation. PreSubmitted - Indicates that a simulated order type has been accepted by the IBKR system and that this order has yet to be elected.
	//                The order is held in the IBKR system until the election criteria are met. At that time the order is transmitted to the order destination as specified.
	// Submitted - Indicates that the order has been accepted at the order destination and is working. Cancelled - Indicates that the balance of the order has been confirmed cancelled by the IB system.
	//             This could occur unexpectedly when IB or the destination has rejected the order.
	// Filled - Indicates that the order has been completely filled.  Inactive - Indicates the order is not working, for instance if the order was invalid and triggered an error message,
	//            or if the order was to short a security and shares have not yet been located.
	//
	Status string `json:"status,omitempty"`

	// for exmple FB
	Ticker string `json:"ticker,omitempty"`
}

// Validate validates this order
func (m *Order) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Order) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Order) UnmarshalBinary(b []byte) error {
	var res Order
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
