// Code generated by go-swagger; DO NOT EDIT.

package introducers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"lwebco.de/cosmic-apis-spec/introducers/models"
)

// AddIntroducerOwnerReader is a Reader for the AddIntroducerOwner structure.
type AddIntroducerOwnerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddIntroducerOwnerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddIntroducerOwnerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddIntroducerOwnerOK creates a AddIntroducerOwnerOK with default headers values
func NewAddIntroducerOwnerOK() *AddIntroducerOwnerOK {
	return &AddIntroducerOwnerOK{}
}

/*AddIntroducerOwnerOK handles this case with default header values.

A successful response.
*/
type AddIntroducerOwnerOK struct {
	Payload models.V1AddIntroducerOwnerResponse
}

func (o *AddIntroducerOwnerOK) Error() string {
	return fmt.Sprintf("[POST /v1/introducers/{introducer_id}/owners][%d] addIntroducerOwnerOK  %+v", 200, o.Payload)
}

func (o *AddIntroducerOwnerOK) GetPayload() models.V1AddIntroducerOwnerResponse {
	return o.Payload
}

func (o *AddIntroducerOwnerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
