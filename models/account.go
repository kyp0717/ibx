// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Account account information
// swagger:model account
type Account struct {

	// account alias
	AccountAlias string `json:"accountAlias,omitempty"`

	// account Id
	AccountID string `json:"accountId,omitempty"`

	// account status
	AccountStatus float64 `json:"accountStatus,omitempty"`

	// account title
	AccountTitle string `json:"accountTitle,omitempty"`

	// account van
	AccountVan string `json:"accountVan,omitempty"`

	// covestor
	Covestor bool `json:"covestor,omitempty"`

	// currency
	Currency string `json:"currency,omitempty"`

	// desc
	Desc string `json:"desc,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// faclient
	Faclient bool `json:"faclient,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// master
	Master *AccountMaster `json:"master,omitempty"`

	// parent
	Parent string `json:"parent,omitempty"`

	// trading type
	TradingType string `json:"tradingType,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this account
func (m *Account) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMaster(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Account) validateMaster(formats strfmt.Registry) error {

	if swag.IsZero(m.Master) { // not required
		return nil
	}

	if m.Master != nil {
		if err := m.Master.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("master")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Account) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Account) UnmarshalBinary(b []byte) error {
	var res Account
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AccountMaster account master
// swagger:model AccountMaster
type AccountMaster struct {

	// official title
	OfficialTitle string `json:"officialTitle,omitempty"`

	// title
	Title string `json:"title,omitempty"`
}

// Validate validates this account master
func (m *AccountMaster) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountMaster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountMaster) UnmarshalBinary(b []byte) error {
	var res AccountMaster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
