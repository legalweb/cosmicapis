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

// GetIntroducersReader is a Reader for the GetIntroducers structure.
type GetIntroducersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntroducersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntroducersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetIntroducersOK creates a GetIntroducersOK with default headers values
func NewGetIntroducersOK() *GetIntroducersOK {
	return &GetIntroducersOK{}
}

/*GetIntroducersOK handles this case with default header values.

A successful response.
*/
type GetIntroducersOK struct {
	Payload *models.V1GetIntroducersResponse
}

func (o *GetIntroducersOK) Error() string {
	return fmt.Sprintf("[GET /v1/introducers][%d] getIntroducersOK  %+v", 200, o.Payload)
}

func (o *GetIntroducersOK) GetPayload() *models.V1GetIntroducersResponse {
	return o.Payload
}

func (o *GetIntroducersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GetIntroducersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
