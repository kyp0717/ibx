// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "learn2/models"
)

// GetPortfolioAccountIDLedgerReader is a Reader for the GetPortfolioAccountIDLedger structure.
type GetPortfolioAccountIDLedgerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPortfolioAccountIDLedgerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPortfolioAccountIDLedgerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPortfolioAccountIDLedgerOK creates a GetPortfolioAccountIDLedgerOK with default headers values
func NewGetPortfolioAccountIDLedgerOK() *GetPortfolioAccountIDLedgerOK {
	return &GetPortfolioAccountIDLedgerOK{}
}

/*GetPortfolioAccountIDLedgerOK handles this case with default header values.

200 means successful
*/
type GetPortfolioAccountIDLedgerOK struct {
	Payload *GetPortfolioAccountIDLedgerOKBody
}

func (o *GetPortfolioAccountIDLedgerOK) Error() string {
	return fmt.Sprintf("[GET /portfolio/{accountId}/ledger][%d] getPortfolioAccountIdLedgerOK  %+v", 200, o.Payload)
}

func (o *GetPortfolioAccountIDLedgerOK) GetPayload() *GetPortfolioAccountIDLedgerOKBody {
	return o.Payload
}

func (o *GetPortfolioAccountIDLedgerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetPortfolioAccountIDLedgerOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetPortfolioAccountIDLedgerOKBody get portfolio account ID ledger o k body
swagger:model GetPortfolioAccountIDLedgerOKBody
*/
type GetPortfolioAccountIDLedgerOKBody struct {

	// b a s e
	BASE *models.Ledger `json:"BASE,omitempty"`
}

// Validate validates this get portfolio account ID ledger o k body
func (o *GetPortfolioAccountIDLedgerOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBASE(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPortfolioAccountIDLedgerOKBody) validateBASE(formats strfmt.Registry) error {

	if swag.IsZero(o.BASE) { // not required
		return nil
	}

	if o.BASE != nil {
		if err := o.BASE.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getPortfolioAccountIdLedgerOK" + "." + "BASE")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPortfolioAccountIDLedgerOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPortfolioAccountIDLedgerOKBody) UnmarshalBinary(b []byte) error {
	var res GetPortfolioAccountIDLedgerOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
