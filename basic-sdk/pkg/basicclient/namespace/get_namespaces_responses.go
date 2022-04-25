// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package namespace

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

// GetNamespacesReader is a Reader for the GetNamespaces structure.
type GetNamespacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNamespacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNamespacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetNamespacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetNamespacesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /basic/v1/admin/namespaces returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetNamespacesOK creates a GetNamespacesOK with default headers values
func NewGetNamespacesOK() *GetNamespacesOK {
	return &GetNamespacesOK{}
}

/*GetNamespacesOK handles this case with default header values.

  Successful operation
*/
type GetNamespacesOK struct {
	Payload []*basicclientmodels.NamespaceInfo
}

func (o *GetNamespacesOK) Error() string {
	return fmt.Sprintf("[GET /basic/v1/admin/namespaces][%d] getNamespacesOK  %+v", 200, o.Payload)
}

func (o *GetNamespacesOK) GetPayload() []*basicclientmodels.NamespaceInfo {
	return o.Payload
}

func (o *GetNamespacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNamespacesUnauthorized creates a GetNamespacesUnauthorized with default headers values
func NewGetNamespacesUnauthorized() *GetNamespacesUnauthorized {
	return &GetNamespacesUnauthorized{}
}

/*GetNamespacesUnauthorized handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20001</td><td>unauthorized</td></tr></table>
*/
type GetNamespacesUnauthorized struct {
	Payload *basicclientmodels.ErrorEntity
}

func (o *GetNamespacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /basic/v1/admin/namespaces][%d] getNamespacesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetNamespacesUnauthorized) GetPayload() *basicclientmodels.ErrorEntity {
	return o.Payload
}

func (o *GetNamespacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(basicclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNamespacesForbidden creates a GetNamespacesForbidden with default headers values
func NewGetNamespacesForbidden() *GetNamespacesForbidden {
	return &GetNamespacesForbidden{}
}

/*GetNamespacesForbidden handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20013</td><td>insufficient permission</td></tr></table>
*/
type GetNamespacesForbidden struct {
	Payload *basicclientmodels.ErrorEntity
}

func (o *GetNamespacesForbidden) Error() string {
	return fmt.Sprintf("[GET /basic/v1/admin/namespaces][%d] getNamespacesForbidden  %+v", 403, o.Payload)
}

func (o *GetNamespacesForbidden) GetPayload() *basicclientmodels.ErrorEntity {
	return o.Payload
}

func (o *GetNamespacesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(basicclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
