// Code generated by go-swagger; DO NOT EDIT.

package scanner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// GetScannerParamsReader is a Reader for the GetScannerParams structure.
type GetScannerParamsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScannerParamsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScannerParamsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetScannerParamsOK creates a GetScannerParamsOK with default headers values
func NewGetScannerParamsOK() *GetScannerParamsOK {
	return &GetScannerParamsOK{}
}

/*GetScannerParamsOK handles this case with default header values.

An object contains lists
*/
type GetScannerParamsOK struct {
	Payload *GetScannerParamsOKBody
}

func (o *GetScannerParamsOK) Error() string {
	return fmt.Sprintf("[GET /iserver/scanner/params][%d] getScannerParamsOK  %+v", 200, o.Payload)
}

func (o *GetScannerParamsOK) GetPayload() *GetScannerParamsOKBody {
	return o.Payload
}

func (o *GetScannerParamsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetScannerParamsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*FilterListItems0 filter list items0
swagger:model FilterListItems0
*/
type FilterListItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// display name
	DisplayName string `json:"display_name,omitempty"`

	// group
	Group string `json:"group,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this filter list items0
func (o *FilterListItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *FilterListItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *FilterListItems0) UnmarshalBinary(b []byte) error {
	var res FilterListItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetScannerParamsOKBody get iserver scanner params o k body
swagger:model GetScannerParamsOKBody
*/
type GetScannerParamsOKBody struct {

	// filter list
	FilterList []*FilterListItems0 `json:"filter_list"`

	// instrument list
	InstrumentList []*InstrumentListItems0 `json:"instrument_list"`

	// location tree
	LocationTree []*LocationTreeItems0 `json:"location_tree"`

	// scan type list
	ScanTypeList []*ScanTypeListItems0 `json:"scan_type_list"`
}

// Validate validates this get iserver scanner params o k body
func (o *GetScannerParamsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFilterList(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateInstrumentList(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLocationTree(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateScanTypeList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetScannerParamsOKBody) validateFilterList(formats strfmt.Registry) error {

	if swag.IsZero(o.FilterList) { // not required
		return nil
	}

	for i := 0; i < len(o.FilterList); i++ {
		if swag.IsZero(o.FilterList[i]) { // not required
			continue
		}

		if o.FilterList[i] != nil {
			if err := o.FilterList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getScannerParamsOK" + "." + "filter_list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetScannerParamsOKBody) validateInstrumentList(formats strfmt.Registry) error {

	if swag.IsZero(o.InstrumentList) { // not required
		return nil
	}

	for i := 0; i < len(o.InstrumentList); i++ {
		if swag.IsZero(o.InstrumentList[i]) { // not required
			continue
		}

		if o.InstrumentList[i] != nil {
			if err := o.InstrumentList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getScannerParamsOK" + "." + "instrument_list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetScannerParamsOKBody) validateLocationTree(formats strfmt.Registry) error {

	if swag.IsZero(o.LocationTree) { // not required
		return nil
	}

	for i := 0; i < len(o.LocationTree); i++ {
		if swag.IsZero(o.LocationTree[i]) { // not required
			continue
		}

		if o.LocationTree[i] != nil {
			if err := o.LocationTree[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getScannerParamsOK" + "." + "location_tree" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetScannerParamsOKBody) validateScanTypeList(formats strfmt.Registry) error {

	if swag.IsZero(o.ScanTypeList) { // not required
		return nil
	}

	for i := 0; i < len(o.ScanTypeList); i++ {
		if swag.IsZero(o.ScanTypeList[i]) { // not required
			continue
		}

		if o.ScanTypeList[i] != nil {
			if err := o.ScanTypeList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getScannerParamsOK" + "." + "scan_type_list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetScannerParamsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetScannerParamsOKBody) UnmarshalBinary(b []byte) error {
	var res GetScannerParamsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*InstrumentListItems0 instrument list items0
swagger:model InstrumentListItems0
*/
type InstrumentListItems0 struct {

	// display name
	DisplayName string `json:"display_name,omitempty"`

	// filters
	Filters []string `json:"filters"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this instrument list items0
func (o *InstrumentListItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InstrumentListItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InstrumentListItems0) UnmarshalBinary(b []byte) error {
	var res InstrumentListItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LocationTreeItems0 location tree items0
swagger:model LocationTreeItems0
*/
type LocationTreeItems0 struct {

	// display name
	DisplayName string `json:"display_name,omitempty"`

	// locations
	Locations []*LocationTreeItems0LocationsItems0 `json:"locations"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this location tree items0
func (o *LocationTreeItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLocations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LocationTreeItems0) validateLocations(formats strfmt.Registry) error {

	if swag.IsZero(o.Locations) { // not required
		return nil
	}

	for i := 0; i < len(o.Locations); i++ {
		if swag.IsZero(o.Locations[i]) { // not required
			continue
		}

		if o.Locations[i] != nil {
			if err := o.Locations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("locations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *LocationTreeItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LocationTreeItems0) UnmarshalBinary(b []byte) error {
	var res LocationTreeItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LocationTreeItems0LocationsItems0 location tree items0 locations items0
swagger:model LocationTreeItems0LocationsItems0
*/
type LocationTreeItems0LocationsItems0 struct {

	// display name
	DisplayName string `json:"display_name,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this location tree items0 locations items0
func (o *LocationTreeItems0LocationsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LocationTreeItems0LocationsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LocationTreeItems0LocationsItems0) UnmarshalBinary(b []byte) error {
	var res LocationTreeItems0LocationsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ScanTypeListItems0 scan type list items0
swagger:model ScanTypeListItems0
*/
type ScanTypeListItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// display name
	DisplayName string `json:"display_name,omitempty"`

	// instruments
	Instruments []string `json:"instruments"`
}

// Validate validates this scan type list items0
func (o *ScanTypeListItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ScanTypeListItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ScanTypeListItems0) UnmarshalBinary(b []byte) error {
	var res ScanTypeListItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
