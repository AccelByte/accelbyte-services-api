// Code generated by go-swagger; DO NOT EDIT.

package user_statistic

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/social-sdk/pkg/socialclientmodels"
)

// BulkFetchStatItems1Reader is a Reader for the BulkFetchStatItems1 structure.
type BulkFetchStatItems1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BulkFetchStatItems1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBulkFetchStatItems1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewBulkFetchStatItems1UnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /v1/public/namespaces/{namespace}/statitems/bulk returns an error %d: %s", response.Code(), string(data))
	}
}

// NewBulkFetchStatItems1OK creates a BulkFetchStatItems1OK with default headers values
func NewBulkFetchStatItems1OK() *BulkFetchStatItems1OK {
	return &BulkFetchStatItems1OK{}
}

/*BulkFetchStatItems1OK handles this case with default header values.

  successful operation
*/
type BulkFetchStatItems1OK struct {
	Payload []*socialclientmodels.UserStatItemInfo
}

func (o *BulkFetchStatItems1OK) Error() string {
	return fmt.Sprintf("[GET /v1/public/namespaces/{namespace}/statitems/bulk][%d] bulkFetchStatItems1OK  %+v", 200, o.Payload)
}

func (o *BulkFetchStatItems1OK) GetPayload() []*socialclientmodels.UserStatItemInfo {
	return o.Payload
}

func (o *BulkFetchStatItems1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkFetchStatItems1UnprocessableEntity creates a BulkFetchStatItems1UnprocessableEntity with default headers values
func NewBulkFetchStatItems1UnprocessableEntity() *BulkFetchStatItems1UnprocessableEntity {
	return &BulkFetchStatItems1UnprocessableEntity{}
}

/*BulkFetchStatItems1UnprocessableEntity handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>20002</td><td>validation error</td></tr></table>
*/
type BulkFetchStatItems1UnprocessableEntity struct {
	Payload *socialclientmodels.ValidationErrorEntity
}

func (o *BulkFetchStatItems1UnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /v1/public/namespaces/{namespace}/statitems/bulk][%d] bulkFetchStatItems1UnprocessableEntity  %+v", 422, o.Payload)
}

func (o *BulkFetchStatItems1UnprocessableEntity) GetPayload() *socialclientmodels.ValidationErrorEntity {
	return o.Payload
}

func (o *BulkFetchStatItems1UnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(socialclientmodels.ValidationErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}