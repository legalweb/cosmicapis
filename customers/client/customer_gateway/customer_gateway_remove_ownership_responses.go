// Code generated by go-swagger; DO NOT EDIT.

package customer_gateway

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"lwebco.de/cosmic-apis-spec/customers/models"
)

// CustomerGatewayRemoveOwnershipReader is a Reader for the CustomerGatewayRemoveOwnership structure.
type CustomerGatewayRemoveOwnershipReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerGatewayRemoveOwnershipReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerGatewayRemoveOwnershipOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerGatewayRemoveOwnershipDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerGatewayRemoveOwnershipOK creates a CustomerGatewayRemoveOwnershipOK with default headers values
func NewCustomerGatewayRemoveOwnershipOK() *CustomerGatewayRemoveOwnershipOK {
	return &CustomerGatewayRemoveOwnershipOK{}
}

/*CustomerGatewayRemoveOwnershipOK handles this case with default header values.

A successful response.
*/
type CustomerGatewayRemoveOwnershipOK struct {
	Payload models.V1RemoveOwnershipResponse
}

func (o *CustomerGatewayRemoveOwnershipOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/ownerships/{ownership_id}][%d] customerGatewayRemoveOwnershipOK  %+v", 200, o.Payload)
}

func (o *CustomerGatewayRemoveOwnershipOK) GetPayload() models.V1RemoveOwnershipResponse {
	return o.Payload
}

func (o *CustomerGatewayRemoveOwnershipOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerGatewayRemoveOwnershipDefault creates a CustomerGatewayRemoveOwnershipDefault with default headers values
func NewCustomerGatewayRemoveOwnershipDefault(code int) *CustomerGatewayRemoveOwnershipDefault {
	return &CustomerGatewayRemoveOwnershipDefault{
		_statusCode: code,
	}
}

/*CustomerGatewayRemoveOwnershipDefault handles this case with default header values.

An unexpected error response
*/
type CustomerGatewayRemoveOwnershipDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the customer gateway remove ownership default response
func (o *CustomerGatewayRemoveOwnershipDefault) Code() int {
	return o._statusCode
}

func (o *CustomerGatewayRemoveOwnershipDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/ownerships/{ownership_id}][%d] CustomerGateway_RemoveOwnership default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerGatewayRemoveOwnershipDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *CustomerGatewayRemoveOwnershipDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
