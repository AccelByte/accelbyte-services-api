// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/dsmc-sdk/pkg/dsmcclientmodels"
)

// GetDefaultProviderReader is a Reader for the GetDefaultProvider structure.
type GetDefaultProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDefaultProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDefaultProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDefaultProviderOK creates a GetDefaultProviderOK with default headers values
func NewGetDefaultProviderOK() *GetDefaultProviderOK {
	return &GetDefaultProviderOK{}
}

/*GetDefaultProviderOK handles this case with default header values.

Default provider got
*/
type GetDefaultProviderOK struct {
	Payload *dsmcclientmodels.ModelsDefaultProvider
}

func (o *GetDefaultProviderOK) Error() string {
	return fmt.Sprintf("[GET /dsmcontroller/public/provider/default][%d] getDefaultProviderOK  %+v", 200, o.Payload)
}

func (o *GetDefaultProviderOK) GetPayload() *dsmcclientmodels.ModelsDefaultProvider {
	return o.Payload
}

func (o *GetDefaultProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ModelsDefaultProvider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}