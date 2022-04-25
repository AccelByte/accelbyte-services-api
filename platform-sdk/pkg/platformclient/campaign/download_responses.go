// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package campaign

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DownloadReader is a Reader for the Download structure.
type DownloadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DownloadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDownloadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /platform/admin/namespaces/{namespace}/codes/campaigns/{campaignId}/codes.csv returns an error %d: %s", response.Code(), string(data))
	}
}

// NewDownloadOK creates a DownloadOK with default headers values
func NewDownloadOK() *DownloadOK {
	return &DownloadOK{}
}

/*DownloadOK handles this case with default header values.

  Successful operation
*/
type DownloadOK struct {
}

func (o *DownloadOK) Error() string {
	return fmt.Sprintf("[GET /platform/admin/namespaces/{namespace}/codes/campaigns/{campaignId}/codes.csv][%d] downloadOK ", 200)
}

func (o *DownloadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
