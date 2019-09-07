// Code generated by go-swagger; DO NOT EDIT.

package portfolio

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "learn2/models"
)

// GetPortfolioAccountIDPositionConidReader is a Reader for the GetPortfolioAccountIDPositionConid structure.
type GetPortfolioAccountIDPositionConidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPortfolioAccountIDPositionConidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPortfolioAccountIDPositionConidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPortfolioAccountIDPositionConidOK creates a GetPortfolioAccountIDPositionConidOK with default headers values
func NewGetPortfolioAccountIDPositionConidOK() *GetPortfolioAccountIDPositionConidOK {
	return &GetPortfolioAccountIDPositionConidOK{}
}

/*GetPortfolioAccountIDPositionConidOK handles this case with default header values.

returns a list containing only one position for the conid
*/
type GetPortfolioAccountIDPositionConidOK struct {
	Payload models.Position
}

func (o *GetPortfolioAccountIDPositionConidOK) Error() string {
	return fmt.Sprintf("[GET /portfolio/{accountId}/position/{conid}][%d] getPortfolioAccountIdPositionConidOK  %+v", 200, o.Payload)
}

func (o *GetPortfolioAccountIDPositionConidOK) GetPayload() models.Position {
	return o.Payload
}

func (o *GetPortfolioAccountIDPositionConidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
