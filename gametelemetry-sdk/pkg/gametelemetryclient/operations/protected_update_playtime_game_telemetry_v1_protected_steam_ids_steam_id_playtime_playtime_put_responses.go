// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/gametelemetry-sdk/pkg/gametelemetryclientmodels"
)

// ProtectedUpdatePlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimePlaytimePutReader is a Reader for the ProtectedUpdatePlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimePlaytimePut structure.
type ProtectedUpdatePlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimePlaytimePutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProtectedUpdatePlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimePlaytimePutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProtectedUpdatePlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimePlaytimePutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewProtectedUpdatePlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimePlaytimePutUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /game-telemetry/v1/protected/steamIds/{steamId}/playtime/{playtime} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewProtectedUpdatePlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimePlaytimePutOK creates a ProtectedUpdatePlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimePlaytimePutOK with default headers values
func NewProtectedUpdatePlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimePlaytimePutOK() *ProtectedUpdatePlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimePlaytimePutOK {
	return &ProtectedUpdatePlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimePlaytimePutOK{}
}

/*ProtectedUpdatePlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimePlaytimePutOK handles this case with default header values.

  Successful Response
*/
type ProtectedUpdatePlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimePlaytimePutOK struct {
	Payload interface{}
}

func (o *ProtectedUpdatePlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimePlaytimePutOK) Error() string {
	return fmt.Sprintf("[PUT /game-telemetry/v1/protected/steamIds/{steamId}/playtime/{playtime}][%d] protectedUpdatePlaytimeGameTelemetryV1ProtectedSteamIdsSteamIdPlaytimePlaytimePutOK  %+v", 200, o.Payload)
}

func (o *ProtectedUpdatePlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimePlaytimePutOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ProtectedUpdatePlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimePlaytimePutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProtectedUpdatePlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimePlaytimePutUnprocessableEntity creates a ProtectedUpdatePlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimePlaytimePutUnprocessableEntity with default headers values
func NewProtectedUpdatePlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimePlaytimePutUnprocessableEntity() *ProtectedUpdatePlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimePlaytimePutUnprocessableEntity {
	return &ProtectedUpdatePlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimePlaytimePutUnprocessableEntity{}
}

/*ProtectedUpdatePlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimePlaytimePutUnprocessableEntity handles this case with default header values.

  Validation Error
*/
type ProtectedUpdatePlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimePlaytimePutUnprocessableEntity struct {
	Payload *gametelemetryclientmodels.HTTPValidationError
}

func (o *ProtectedUpdatePlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimePlaytimePutUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /game-telemetry/v1/protected/steamIds/{steamId}/playtime/{playtime}][%d] protectedUpdatePlaytimeGameTelemetryV1ProtectedSteamIdsSteamIdPlaytimePlaytimePutUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ProtectedUpdatePlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimePlaytimePutUnprocessableEntity) GetPayload() *gametelemetryclientmodels.HTTPValidationError {
	return o.Payload
}

func (o *ProtectedUpdatePlaytimeGameTelemetryV1ProtectedSteamIdsSteamIDPlaytimePlaytimePutUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(gametelemetryclientmodels.HTTPValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
