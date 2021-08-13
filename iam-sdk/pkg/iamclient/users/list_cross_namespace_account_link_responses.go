// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ListCrossNamespaceAccountLinkReader is a Reader for the ListCrossNamespaceAccountLink structure.
type ListCrossNamespaceAccountLinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCrossNamespaceAccountLinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCrossNamespaceAccountLinkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListCrossNamespaceAccountLinkBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListCrossNamespaceAccountLinkUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListCrossNamespaceAccountLinkForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewListCrossNamespaceAccountLinkNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /iam/namespaces/{namespace}/users/{userId}/crosslink returns an error %d: %s", response.Code(), string(data))
	}
}

// NewListCrossNamespaceAccountLinkOK creates a ListCrossNamespaceAccountLinkOK with default headers values
func NewListCrossNamespaceAccountLinkOK() *ListCrossNamespaceAccountLinkOK {
	return &ListCrossNamespaceAccountLinkOK{}
}

/*ListCrossNamespaceAccountLinkOK handles this case with default header values.

  Operation succeeded
*/
type ListCrossNamespaceAccountLinkOK struct {
}

func (o *ListCrossNamespaceAccountLinkOK) Error() string {
	return fmt.Sprintf("[POST /iam/namespaces/{namespace}/users/{userId}/crosslink][%d] listCrossNamespaceAccountLinkOK ", 200)
}

func (o *ListCrossNamespaceAccountLinkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListCrossNamespaceAccountLinkBadRequest creates a ListCrossNamespaceAccountLinkBadRequest with default headers values
func NewListCrossNamespaceAccountLinkBadRequest() *ListCrossNamespaceAccountLinkBadRequest {
	return &ListCrossNamespaceAccountLinkBadRequest{}
}

/*ListCrossNamespaceAccountLinkBadRequest handles this case with default header values.

  Invalid request
*/
type ListCrossNamespaceAccountLinkBadRequest struct {
}

func (o *ListCrossNamespaceAccountLinkBadRequest) Error() string {
	return fmt.Sprintf("[POST /iam/namespaces/{namespace}/users/{userId}/crosslink][%d] listCrossNamespaceAccountLinkBadRequest ", 400)
}

func (o *ListCrossNamespaceAccountLinkBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListCrossNamespaceAccountLinkUnauthorized creates a ListCrossNamespaceAccountLinkUnauthorized with default headers values
func NewListCrossNamespaceAccountLinkUnauthorized() *ListCrossNamespaceAccountLinkUnauthorized {
	return &ListCrossNamespaceAccountLinkUnauthorized{}
}

/*ListCrossNamespaceAccountLinkUnauthorized handles this case with default header values.

  Unauthorized access
*/
type ListCrossNamespaceAccountLinkUnauthorized struct {
}

func (o *ListCrossNamespaceAccountLinkUnauthorized) Error() string {
	return fmt.Sprintf("[POST /iam/namespaces/{namespace}/users/{userId}/crosslink][%d] listCrossNamespaceAccountLinkUnauthorized ", 401)
}

func (o *ListCrossNamespaceAccountLinkUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListCrossNamespaceAccountLinkForbidden creates a ListCrossNamespaceAccountLinkForbidden with default headers values
func NewListCrossNamespaceAccountLinkForbidden() *ListCrossNamespaceAccountLinkForbidden {
	return &ListCrossNamespaceAccountLinkForbidden{}
}

/*ListCrossNamespaceAccountLinkForbidden handles this case with default header values.

  Forbidden
*/
type ListCrossNamespaceAccountLinkForbidden struct {
}

func (o *ListCrossNamespaceAccountLinkForbidden) Error() string {
	return fmt.Sprintf("[POST /iam/namespaces/{namespace}/users/{userId}/crosslink][%d] listCrossNamespaceAccountLinkForbidden ", 403)
}

func (o *ListCrossNamespaceAccountLinkForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListCrossNamespaceAccountLinkNotFound creates a ListCrossNamespaceAccountLinkNotFound with default headers values
func NewListCrossNamespaceAccountLinkNotFound() *ListCrossNamespaceAccountLinkNotFound {
	return &ListCrossNamespaceAccountLinkNotFound{}
}

/*ListCrossNamespaceAccountLinkNotFound handles this case with default header values.

  Data not found
*/
type ListCrossNamespaceAccountLinkNotFound struct {
}

func (o *ListCrossNamespaceAccountLinkNotFound) Error() string {
	return fmt.Sprintf("[POST /iam/namespaces/{namespace}/users/{userId}/crosslink][%d] listCrossNamespaceAccountLinkNotFound ", 404)
}

func (o *ListCrossNamespaceAccountLinkNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
