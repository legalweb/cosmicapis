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

// CustomerGatewaySetNotificationPreferencesReader is a Reader for the CustomerGatewaySetNotificationPreferences structure.
type CustomerGatewaySetNotificationPreferencesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerGatewaySetNotificationPreferencesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerGatewaySetNotificationPreferencesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerGatewaySetNotificationPreferencesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerGatewaySetNotificationPreferencesOK creates a CustomerGatewaySetNotificationPreferencesOK with default headers values
func NewCustomerGatewaySetNotificationPreferencesOK() *CustomerGatewaySetNotificationPreferencesOK {
	return &CustomerGatewaySetNotificationPreferencesOK{}
}

/*CustomerGatewaySetNotificationPreferencesOK handles this case with default header values.

A successful response.
*/
type CustomerGatewaySetNotificationPreferencesOK struct {
	Payload models.V1SetNotificationPreferencesResponse
}

func (o *CustomerGatewaySetNotificationPreferencesOK) Error() string {
	return fmt.Sprintf("[POST /v1/settings/notifications][%d] customerGatewaySetNotificationPreferencesOK  %+v", 200, o.Payload)
}

func (o *CustomerGatewaySetNotificationPreferencesOK) GetPayload() models.V1SetNotificationPreferencesResponse {
	return o.Payload
}

func (o *CustomerGatewaySetNotificationPreferencesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerGatewaySetNotificationPreferencesDefault creates a CustomerGatewaySetNotificationPreferencesDefault with default headers values
func NewCustomerGatewaySetNotificationPreferencesDefault(code int) *CustomerGatewaySetNotificationPreferencesDefault {
	return &CustomerGatewaySetNotificationPreferencesDefault{
		_statusCode: code,
	}
}

/*CustomerGatewaySetNotificationPreferencesDefault handles this case with default header values.

An unexpected error response
*/
type CustomerGatewaySetNotificationPreferencesDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the customer gateway set notification preferences default response
func (o *CustomerGatewaySetNotificationPreferencesDefault) Code() int {
	return o._statusCode
}

func (o *CustomerGatewaySetNotificationPreferencesDefault) Error() string {
	return fmt.Sprintf("[POST /v1/settings/notifications][%d] CustomerGateway_SetNotificationPreferences default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerGatewaySetNotificationPreferencesDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *CustomerGatewaySetNotificationPreferencesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
