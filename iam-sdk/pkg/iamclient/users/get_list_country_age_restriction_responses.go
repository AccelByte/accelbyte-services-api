// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclientmodels"
)

// GetListCountryAgeRestrictionReader is a Reader for the GetListCountryAgeRestriction structure.
type GetListCountryAgeRestrictionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetListCountryAgeRestrictionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetListCountryAgeRestrictionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetListCountryAgeRestrictionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetListCountryAgeRestrictionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetListCountryAgeRestrictionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /iam/v2/admin/namespaces/{namespace}/countries/agerestrictions returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetListCountryAgeRestrictionOK creates a GetListCountryAgeRestrictionOK with default headers values
func NewGetListCountryAgeRestrictionOK() *GetListCountryAgeRestrictionOK {
	return &GetListCountryAgeRestrictionOK{}
}

/*GetListCountryAgeRestrictionOK handles this case with default header values.

  OK
*/
type GetListCountryAgeRestrictionOK struct {
	Payload []*iamclientmodels.AccountcommonCountryAgeRestriction
}

func (o *GetListCountryAgeRestrictionOK) Error() string {
	return fmt.Sprintf("[GET /iam/v2/admin/namespaces/{namespace}/countries/agerestrictions][%d] getListCountryAgeRestrictionOK  %+v", 200, o.ToString())
}

func (o *GetListCountryAgeRestrictionOK) ToString() string {
	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *GetListCountryAgeRestrictionOK) GetPayload() []*iamclientmodels.AccountcommonCountryAgeRestriction {
	return o.Payload
}

func (o *GetListCountryAgeRestrictionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListCountryAgeRestrictionUnauthorized creates a GetListCountryAgeRestrictionUnauthorized with default headers values
func NewGetListCountryAgeRestrictionUnauthorized() *GetListCountryAgeRestrictionUnauthorized {
	return &GetListCountryAgeRestrictionUnauthorized{}
}

/*GetListCountryAgeRestrictionUnauthorized handles this case with default header values.

  Unauthorized access
*/
type GetListCountryAgeRestrictionUnauthorized struct {
}

func (o *GetListCountryAgeRestrictionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /iam/v2/admin/namespaces/{namespace}/countries/agerestrictions][%d] getListCountryAgeRestrictionUnauthorized ", 401)
}

func (o *GetListCountryAgeRestrictionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListCountryAgeRestrictionForbidden creates a GetListCountryAgeRestrictionForbidden with default headers values
func NewGetListCountryAgeRestrictionForbidden() *GetListCountryAgeRestrictionForbidden {
	return &GetListCountryAgeRestrictionForbidden{}
}

/*GetListCountryAgeRestrictionForbidden handles this case with default header values.

  Forbidden
*/
type GetListCountryAgeRestrictionForbidden struct {
}

func (o *GetListCountryAgeRestrictionForbidden) Error() string {
	return fmt.Sprintf("[GET /iam/v2/admin/namespaces/{namespace}/countries/agerestrictions][%d] getListCountryAgeRestrictionForbidden ", 403)
}

func (o *GetListCountryAgeRestrictionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetListCountryAgeRestrictionNotFound creates a GetListCountryAgeRestrictionNotFound with default headers values
func NewGetListCountryAgeRestrictionNotFound() *GetListCountryAgeRestrictionNotFound {
	return &GetListCountryAgeRestrictionNotFound{}
}

/*GetListCountryAgeRestrictionNotFound handles this case with default header values.

  Data not found
*/
type GetListCountryAgeRestrictionNotFound struct {
}

func (o *GetListCountryAgeRestrictionNotFound) Error() string {
	return fmt.Sprintf("[GET /iam/v2/admin/namespaces/{namespace}/countries/agerestrictions][%d] getListCountryAgeRestrictionNotFound ", 404)
}

func (o *GetListCountryAgeRestrictionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
