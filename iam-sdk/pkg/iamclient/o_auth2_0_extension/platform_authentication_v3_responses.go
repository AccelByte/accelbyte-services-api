// Code generated by go-swagger; DO NOT EDIT.

package o_auth2_0_extension

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PlatformAuthenticationV3Reader is a Reader for the PlatformAuthenticationV3 structure.
type PlatformAuthenticationV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PlatformAuthenticationV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 302:
		result := NewPlatformAuthenticationV3Found()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /iam/v3/platforms/{platformId}/authenticate returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPlatformAuthenticationV3Found creates a PlatformAuthenticationV3Found with default headers values
func NewPlatformAuthenticationV3Found() *PlatformAuthenticationV3Found {
	return &PlatformAuthenticationV3Found{}
}

/*PlatformAuthenticationV3Found handles this case with default header values.

  Found. Redirect to clients redirection URL with either code or error on the query parameter
*/
type PlatformAuthenticationV3Found struct {
	/*The Location header
	 */
	Location string
}

func (o *PlatformAuthenticationV3Found) Error() string {
	return fmt.Sprintf("[GET /iam/v3/platforms/{platformId}/authenticate][%d] platformAuthenticationV3Found ", 302)
}

func (o *PlatformAuthenticationV3Found) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	return nil
}
