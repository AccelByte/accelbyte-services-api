// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ListProvidersReader is a Reader for the ListProviders structure.
type ListProvidersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProvidersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProvidersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /dsmcontroller/public/providers returns an error %d: %s", response.Code(), string(data))
	}
}

// NewListProvidersOK creates a ListProvidersOK with default headers values
func NewListProvidersOK() *ListProvidersOK {
	return &ListProvidersOK{}
}

/*ListProvidersOK handles this case with default header values.

  Providers listed
*/
type ListProvidersOK struct {
	Payload []string
}

func (o *ListProvidersOK) Error() string {
	return fmt.Sprintf("[GET /dsmcontroller/public/providers][%d] listProvidersOK  %+v", 200, o.Payload)
}

func (o *ListProvidersOK) GetPayload() []string {
	return o.Payload
}

func (o *ListProvidersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
