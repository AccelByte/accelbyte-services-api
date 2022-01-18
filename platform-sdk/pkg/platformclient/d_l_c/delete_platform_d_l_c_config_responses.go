// Code generated by go-swagger; DO NOT EDIT.

package d_l_c

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeletePlatformDLCConfigReader is a Reader for the DeletePlatformDLCConfig structure.
type DeletePlatformDLCConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePlatformDLCConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeletePlatformDLCConfigNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested DELETE /admin/namespaces/{namespace}/dlc/config/platformMap returns an error %d: %s", response.Code(), string(data))
	}
}

// NewDeletePlatformDLCConfigNoContent creates a DeletePlatformDLCConfigNoContent with default headers values
func NewDeletePlatformDLCConfigNoContent() *DeletePlatformDLCConfigNoContent {
	return &DeletePlatformDLCConfigNoContent{}
}

/*DeletePlatformDLCConfigNoContent handles this case with default header values.

  Delete successfully
*/
type DeletePlatformDLCConfigNoContent struct {
}

func (o *DeletePlatformDLCConfigNoContent) Error() string {
	return fmt.Sprintf("[DELETE /admin/namespaces/{namespace}/dlc/config/platformMap][%d] deletePlatformDLCConfigNoContent ", 204)
}

func (o *DeletePlatformDLCConfigNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
