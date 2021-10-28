// Code generated by go-swagger; DO NOT EDIT.

package campaign

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
)

// BulkEnableCodesReader is a Reader for the BulkEnableCodes structure.
type BulkEnableCodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BulkEnableCodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBulkEnableCodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /admin/namespaces/{namespace}/codes/campaigns/{campaignId}/enable/bulk returns an error %d: %s", response.Code(), string(data))
	}
}

// NewBulkEnableCodesOK creates a BulkEnableCodesOK with default headers values
func NewBulkEnableCodesOK() *BulkEnableCodesOK {
	return &BulkEnableCodesOK{}
}

/*BulkEnableCodesOK handles this case with default header values.

  successful operation
*/
type BulkEnableCodesOK struct {
	Payload *platformclientmodels.BulkOperationResult
}

func (o *BulkEnableCodesOK) Error() string {
	return fmt.Sprintf("[PUT /admin/namespaces/{namespace}/codes/campaigns/{campaignId}/enable/bulk][%d] bulkEnableCodesOK  %+v", 200, o.Payload)
}

func (o *BulkEnableCodesOK) GetPayload() *platformclientmodels.BulkOperationResult {
	return o.Payload
}

func (o *BulkEnableCodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.BulkOperationResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
