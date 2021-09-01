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

// BulkUpdateUserStatItem2Reader is a Reader for the BulkUpdateUserStatItem2 structure.
type BulkUpdateUserStatItem2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BulkUpdateUserStatItem2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBulkUpdateUserStatItem2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewBulkUpdateUserStatItem2UnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /v2/public/namespaces/{namespace}/users/{userId}/statitems/value/bulk returns an error %d: %s", response.Code(), string(data))
	}
}

// NewBulkUpdateUserStatItem2OK creates a BulkUpdateUserStatItem2OK with default headers values
func NewBulkUpdateUserStatItem2OK() *BulkUpdateUserStatItem2OK {
	return &BulkUpdateUserStatItem2OK{}
}

/*BulkUpdateUserStatItem2OK handles this case with default header values.

  successful operation
*/
type BulkUpdateUserStatItem2OK struct {
	Payload []*socialclientmodels.BulkStatItemOperationResult
}

func (o *BulkUpdateUserStatItem2OK) Error() string {
	return fmt.Sprintf("[PUT /v2/public/namespaces/{namespace}/users/{userId}/statitems/value/bulk][%d] bulkUpdateUserStatItem2OK  %+v", 200, o.Payload)
}

func (o *BulkUpdateUserStatItem2OK) GetPayload() []*socialclientmodels.BulkStatItemOperationResult {
	return o.Payload
}

func (o *BulkUpdateUserStatItem2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkUpdateUserStatItem2UnprocessableEntity creates a BulkUpdateUserStatItem2UnprocessableEntity with default headers values
func NewBulkUpdateUserStatItem2UnprocessableEntity() *BulkUpdateUserStatItem2UnprocessableEntity {
	return &BulkUpdateUserStatItem2UnprocessableEntity{}
}

/*BulkUpdateUserStatItem2UnprocessableEntity handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20002</td><td>validation error</td></tr></table>
*/
type BulkUpdateUserStatItem2UnprocessableEntity struct {
	Payload *socialclientmodels.ValidationErrorEntity
}

func (o *BulkUpdateUserStatItem2UnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /v2/public/namespaces/{namespace}/users/{userId}/statitems/value/bulk][%d] bulkUpdateUserStatItem2UnprocessableEntity  %+v", 422, o.Payload)
}

func (o *BulkUpdateUserStatItem2UnprocessableEntity) GetPayload() *socialclientmodels.ValidationErrorEntity {
	return o.Payload
}

func (o *BulkUpdateUserStatItem2UnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(socialclientmodels.ValidationErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}