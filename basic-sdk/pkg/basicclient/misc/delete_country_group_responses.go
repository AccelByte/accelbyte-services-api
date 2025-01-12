// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package misc

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/basic-sdk/pkg/basicclientmodels"
)

// DeleteCountryGroupReader is a Reader for the DeleteCountryGroup structure.
type DeleteCountryGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCountryGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCountryGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteCountryGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteCountryGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteCountryGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteCountryGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested DELETE /basic/v1/admin/namespaces/{namespace}/misc/countrygroups/{countryGroupCode} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewDeleteCountryGroupOK creates a DeleteCountryGroupOK with default headers values
func NewDeleteCountryGroupOK() *DeleteCountryGroupOK {
	return &DeleteCountryGroupOK{}
}

/*DeleteCountryGroupOK handles this case with default header values.

  Successful operation
*/
type DeleteCountryGroupOK struct {
}

func (o *DeleteCountryGroupOK) Error() string {
	return fmt.Sprintf("[DELETE /basic/v1/admin/namespaces/{namespace}/misc/countrygroups/{countryGroupCode}][%d] deleteCountryGroupOK ", 200)
}

func (o *DeleteCountryGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	return nil
}

// NewDeleteCountryGroupBadRequest creates a DeleteCountryGroupBadRequest with default headers values
func NewDeleteCountryGroupBadRequest() *DeleteCountryGroupBadRequest {
	return &DeleteCountryGroupBadRequest{}
}

/*DeleteCountryGroupBadRequest handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20002</td><td>validation error</td></tr></table>
*/
type DeleteCountryGroupBadRequest struct {
	Payload *basicclientmodels.ValidationErrorEntity
}

func (o *DeleteCountryGroupBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /basic/v1/admin/namespaces/{namespace}/misc/countrygroups/{countryGroupCode}][%d] deleteCountryGroupBadRequest  %+v", 400, o.ToJSONString())
}

func (o *DeleteCountryGroupBadRequest) ToJSONString() string {
	if o.Payload == nil {
		return "{}"
	}

	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)

		return fmt.Sprintf("Failed to marshal the payload: %+v", o.Payload)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *DeleteCountryGroupBadRequest) GetPayload() *basicclientmodels.ValidationErrorEntity {
	return o.Payload
}

func (o *DeleteCountryGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(basicclientmodels.ValidationErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCountryGroupUnauthorized creates a DeleteCountryGroupUnauthorized with default headers values
func NewDeleteCountryGroupUnauthorized() *DeleteCountryGroupUnauthorized {
	return &DeleteCountryGroupUnauthorized{}
}

/*DeleteCountryGroupUnauthorized handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20001</td><td>unauthorized</td></tr></table>
*/
type DeleteCountryGroupUnauthorized struct {
	Payload *basicclientmodels.ErrorEntity
}

func (o *DeleteCountryGroupUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /basic/v1/admin/namespaces/{namespace}/misc/countrygroups/{countryGroupCode}][%d] deleteCountryGroupUnauthorized  %+v", 401, o.ToJSONString())
}

func (o *DeleteCountryGroupUnauthorized) ToJSONString() string {
	if o.Payload == nil {
		return "{}"
	}

	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)

		return fmt.Sprintf("Failed to marshal the payload: %+v", o.Payload)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *DeleteCountryGroupUnauthorized) GetPayload() *basicclientmodels.ErrorEntity {
	return o.Payload
}

func (o *DeleteCountryGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(basicclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCountryGroupForbidden creates a DeleteCountryGroupForbidden with default headers values
func NewDeleteCountryGroupForbidden() *DeleteCountryGroupForbidden {
	return &DeleteCountryGroupForbidden{}
}

/*DeleteCountryGroupForbidden handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20013</td><td>insufficient permission</td></tr></table>
*/
type DeleteCountryGroupForbidden struct {
	Payload *basicclientmodels.ErrorEntity
}

func (o *DeleteCountryGroupForbidden) Error() string {
	return fmt.Sprintf("[DELETE /basic/v1/admin/namespaces/{namespace}/misc/countrygroups/{countryGroupCode}][%d] deleteCountryGroupForbidden  %+v", 403, o.ToJSONString())
}

func (o *DeleteCountryGroupForbidden) ToJSONString() string {
	if o.Payload == nil {
		return "{}"
	}

	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)

		return fmt.Sprintf("Failed to marshal the payload: %+v", o.Payload)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *DeleteCountryGroupForbidden) GetPayload() *basicclientmodels.ErrorEntity {
	return o.Payload
}

func (o *DeleteCountryGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(basicclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCountryGroupNotFound creates a DeleteCountryGroupNotFound with default headers values
func NewDeleteCountryGroupNotFound() *DeleteCountryGroupNotFound {
	return &DeleteCountryGroupNotFound{}
}

/*DeleteCountryGroupNotFound handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>11233</td><td>Unable to {action}: Country group with code [{countryGroupCode}] is not found</td></tr></table>
*/
type DeleteCountryGroupNotFound struct {
	Payload *basicclientmodels.ErrorEntity
}

func (o *DeleteCountryGroupNotFound) Error() string {
	return fmt.Sprintf("[DELETE /basic/v1/admin/namespaces/{namespace}/misc/countrygroups/{countryGroupCode}][%d] deleteCountryGroupNotFound  %+v", 404, o.ToJSONString())
}

func (o *DeleteCountryGroupNotFound) ToJSONString() string {
	if o.Payload == nil {
		return "{}"
	}

	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)

		return fmt.Sprintf("Failed to marshal the payload: %+v", o.Payload)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *DeleteCountryGroupNotFound) GetPayload() *basicclientmodels.ErrorEntity {
	return o.Payload
}

func (o *DeleteCountryGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(basicclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
