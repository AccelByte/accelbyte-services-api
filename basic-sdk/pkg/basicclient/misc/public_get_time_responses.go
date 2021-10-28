// Code generated by go-swagger; DO NOT EDIT.

package misc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/basic-sdk/pkg/basicclientmodels"
)

// PublicGetTimeReader is a Reader for the PublicGetTime structure.
type PublicGetTimeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicGetTimeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicGetTimeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /v1/public/misc/time returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPublicGetTimeOK creates a PublicGetTimeOK with default headers values
func NewPublicGetTimeOK() *PublicGetTimeOK {
	return &PublicGetTimeOK{}
}

/*PublicGetTimeOK handles this case with default header values.

  Success retrieve server time
*/
type PublicGetTimeOK struct {
	Payload *basicclientmodels.RetrieveTimeResponse
}

func (o *PublicGetTimeOK) Error() string {
	return fmt.Sprintf("[GET /v1/public/misc/time][%d] publicGetTimeOK  %+v", 200, o.Payload)
}

func (o *PublicGetTimeOK) GetPayload() *basicclientmodels.RetrieveTimeResponse {
	return o.Payload
}

func (o *PublicGetTimeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(basicclientmodels.RetrieveTimeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
