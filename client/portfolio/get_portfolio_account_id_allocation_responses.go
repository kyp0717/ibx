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

// GetPortfolioAccountIDAllocationReader is a Reader for the GetPortfolioAccountIDAllocation structure.
type GetPortfolioAccountIDAllocationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPortfolioAccountIDAllocationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPortfolioAccountIDAllocationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPortfolioAccountIDAllocationOK creates a GetPortfolioAccountIDAllocationOK with default headers values
func NewGetPortfolioAccountIDAllocationOK() *GetPortfolioAccountIDAllocationOK {
	return &GetPortfolioAccountIDAllocationOK{}
}

/*GetPortfolioAccountIDAllocationOK handles this case with default header values.

returns an object of three different allocations
*/
type GetPortfolioAccountIDAllocationOK struct {
	Payload models.Allocation
}

func (o *GetPortfolioAccountIDAllocationOK) Error() string {
	return fmt.Sprintf("[GET /portfolio/{accountId}/allocation][%d] getPortfolioAccountIdAllocationOK  %+v", 200, o.Payload)
}

func (o *GetPortfolioAccountIDAllocationOK) GetPayload() models.Allocation {
	return o.Payload
}

func (o *GetPortfolioAccountIDAllocationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
