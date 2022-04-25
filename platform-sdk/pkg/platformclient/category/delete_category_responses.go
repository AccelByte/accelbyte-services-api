// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package category

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

// DeleteCategoryReader is a Reader for the DeleteCategory structure.
type DeleteCategoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCategoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCategoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteCategoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewDeleteCategoryConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested DELETE /platform/admin/namespaces/{namespace}/categories/{categoryPath} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewDeleteCategoryOK creates a DeleteCategoryOK with default headers values
func NewDeleteCategoryOK() *DeleteCategoryOK {
	return &DeleteCategoryOK{}
}

/*DeleteCategoryOK handles this case with default header values.

  successful operation
*/
type DeleteCategoryOK struct {
	Payload *platformclientmodels.FullCategoryInfo
}

func (o *DeleteCategoryOK) Error() string {
	return fmt.Sprintf("[DELETE /platform/admin/namespaces/{namespace}/categories/{categoryPath}][%d] deleteCategoryOK  %+v", 200, o.Payload)
}

func (o *DeleteCategoryOK) GetPayload() *platformclientmodels.FullCategoryInfo {
	return o.Payload
}

func (o *DeleteCategoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.FullCategoryInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCategoryNotFound creates a DeleteCategoryNotFound with default headers values
func NewDeleteCategoryNotFound() *DeleteCategoryNotFound {
	return &DeleteCategoryNotFound{}
}

/*DeleteCategoryNotFound handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>30241</td><td>Category [{categoryPath}] does not exist in namespace [{namespace}]</td></tr></table>
*/
type DeleteCategoryNotFound struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *DeleteCategoryNotFound) Error() string {
	return fmt.Sprintf("[DELETE /platform/admin/namespaces/{namespace}/categories/{categoryPath}][%d] deleteCategoryNotFound  %+v", 404, o.Payload)
}

func (o *DeleteCategoryNotFound) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *DeleteCategoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCategoryConflict creates a DeleteCategoryConflict with default headers values
func NewDeleteCategoryConflict() *DeleteCategoryConflict {
	return &DeleteCategoryConflict{}
}

/*DeleteCategoryConflict handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>30272</td><td>Category [{categoryPath}] is not empty in namespace [{namespace}]</td></tr><tr><td>30173</td><td>Published store can't modify content</td></tr></table>
*/
type DeleteCategoryConflict struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *DeleteCategoryConflict) Error() string {
	return fmt.Sprintf("[DELETE /platform/admin/namespaces/{namespace}/categories/{categoryPath}][%d] deleteCategoryConflict  %+v", 409, o.Payload)
}

func (o *DeleteCategoryConflict) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *DeleteCategoryConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
