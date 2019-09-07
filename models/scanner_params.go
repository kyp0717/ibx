// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ScannerParams scanner params
// swagger:model scanner-params
type ScannerParams struct {

	// filter
	Filter []*ScannerParamsFilterItems0 `json:"filter"`

	// for example-STK
	Instrument string `json:"instrument,omitempty"`

	// location
	Location string `json:"location,omitempty"`

	// size
	Size string `json:"size,omitempty"`

	// for example-TOP_PERC_GAIN
	Type string `json:"type,omitempty"`
}

// Validate validates this scanner params
func (m *ScannerParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScannerParams) validateFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.Filter) { // not required
		return nil
	}

	for i := 0; i < len(m.Filter); i++ {
		if swag.IsZero(m.Filter[i]) { // not required
			continue
		}

		if m.Filter[i] != nil {
			if err := m.Filter[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filter" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScannerParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScannerParams) UnmarshalBinary(b []byte) error {
	var res ScannerParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ScannerParamsFilterItems0 scanner params filter items0
// swagger:model ScannerParamsFilterItems0
type ScannerParamsFilterItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// value
	Value float64 `json:"value,omitempty"`
}

// Validate validates this scanner params filter items0
func (m *ScannerParamsFilterItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ScannerParamsFilterItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScannerParamsFilterItems0) UnmarshalBinary(b []byte) error {
	var res ScannerParamsFilterItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
