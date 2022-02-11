// Code generated by go-swagger; DO NOT EDIT.

package i_a_p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SyncTwitchDropsEntitlementReader is a Reader for the SyncTwitchDropsEntitlement structure.
type SyncTwitchDropsEntitlementReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SyncTwitchDropsEntitlementReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSyncTwitchDropsEntitlementNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /public/namespaces/{namespace}/users/{userId}/iap/twitch/sync returns an error %d: %s", response.Code(), string(data))
	}
}

// NewSyncTwitchDropsEntitlementNoContent creates a SyncTwitchDropsEntitlementNoContent with default headers values
func NewSyncTwitchDropsEntitlementNoContent() *SyncTwitchDropsEntitlementNoContent {
	return &SyncTwitchDropsEntitlementNoContent{}
}

/*SyncTwitchDropsEntitlementNoContent handles this case with default header values.

  Sync Successful
*/
type SyncTwitchDropsEntitlementNoContent struct {
}

func (o *SyncTwitchDropsEntitlementNoContent) Error() string {
	return fmt.Sprintf("[PUT /public/namespaces/{namespace}/users/{userId}/iap/twitch/sync][%d] syncTwitchDropsEntitlementNoContent ", 204)
}

func (o *SyncTwitchDropsEntitlementNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
