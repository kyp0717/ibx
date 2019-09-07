// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Contract Contains all details of the contract, including rules you can use when placing orders
// swagger:model contract
type Contract struct {

	// category
	Category string `json:"category,omitempty"`

	// company name
	CompanyName string `json:"companyName,omitempty"`
	// company name
	Company_Name string `json:"company_name,omitempty"`
	// same as that in request
	ConID string `json:"con_id,omitempty"`

	// currency
	Currency string `json:"currency,omitempty"`

	// exchange
	Exchange string `json:"exchange,omitempty"`

	// industry
	Industry string `json:"industry,omitempty"`

	// for example STK
	InstrumentType string `json:"instrument_type,omitempty"`

	// for exmple FB
	LocalSymbol string `json:"local_symbol,omitempty"`

	// true means you can trade outside RTH(regular trading hours)
	Rth bool `json:"r_t_h,omitempty"`

	// rules
	Rules *ContractRules `json:"rules,omitempty"`
}

// Validate validates this contract
func (m *Contract) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Contract) validateRules(formats strfmt.Registry) error {

	if swag.IsZero(m.Rules) { // not required
		return nil
	}

	if m.Rules != nil {
		if err := m.Rules.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rules")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Contract) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Contract) UnmarshalBinary(b []byte) error {
	var res Contract
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ContractRules contract rules
// swagger:model ContractRules
type ContractRules struct {

	// default quantity you can use to place an order
	DefaultSize float64 `json:"defaultSize,omitempty"`

	// display size
	DisplaySize string `json:"displaySize,omitempty"`

	// increment
	Increment string `json:"increment,omitempty"`

	// default limit price you can use to prefill your order
	LimitPrice float64 `json:"limitPrice,omitempty"`

	// order types
	OrderTypes []string `json:"orderTypes"`

	// order types outside
	OrderTypesOutside []string `json:"orderTypesOutside"`

	// if you can preview the order or not with the whatif end-point
	Preview bool `json:"preview,omitempty"`

	// size increment
	SizeIncrement float64 `json:"sizeIncrement,omitempty"`

	// default stop price you can use to prefill your order
	Stopprice float64 `json:"stopprice,omitempty"`

	// tif types
	TifTypes []string `json:"tifTypes"`
}

// Validate validates this contract rules
func (m *ContractRules) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContractRules) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContractRules) UnmarshalBinary(b []byte) error {
	var res ContractRules
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
