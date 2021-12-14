// Code generated by go-swagger; DO NOT EDIT.

package download_server_logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new download server logs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for download server logs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CheckServerLogs(params *CheckServerLogsParams, authInfo runtime.ClientAuthInfoWriter) (*CheckServerLogsOK, *CheckServerLogsNotFound, *CheckServerLogsInternalServerError, error)
	CheckServerLogsShort(params *CheckServerLogsParams, authInfo runtime.ClientAuthInfoWriter) (*CheckServerLogsOK, error)
	DownloadServerLogs(params *DownloadServerLogsParams, authInfo runtime.ClientAuthInfoWriter) (*DownloadServerLogsOK, *DownloadServerLogsNotFound, *DownloadServerLogsInternalServerError, error)
	DownloadServerLogsShort(params *DownloadServerLogsParams, authInfo runtime.ClientAuthInfoWriter) (*DownloadServerLogsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CheckServerLogs checks dedicated server log files existence

  Required permission: ADMIN:NAMESPACE:{namespace}:DSLM:LOG [READ]

Required scope: social

This endpoint will check log file existence before download file.
*/
func (a *Client) CheckServerLogs(params *CheckServerLogsParams, authInfo runtime.ClientAuthInfoWriter) (*CheckServerLogsOK, *CheckServerLogsNotFound, *CheckServerLogsInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCheckServerLogsParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "checkServerLogs",
		Method:             "GET",
		PathPattern:        "/dslogmanager/namespaces/{namespace}/servers/{podName}/logs/exists",
		ProducesMediaTypes: []string{"application/json", "text/x-log"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CheckServerLogsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}

	switch v := result.(type) {

	case *CheckServerLogsOK:
		return v, nil, nil, nil

	case *CheckServerLogsNotFound:
		return nil, v, nil, nil

	case *CheckServerLogsInternalServerError:
		return nil, nil, v, nil

	default:
		return nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) CheckServerLogsShort(params *CheckServerLogsParams, authInfo runtime.ClientAuthInfoWriter) (*CheckServerLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCheckServerLogsParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "checkServerLogs",
		Method:             "GET",
		PathPattern:        "/dslogmanager/namespaces/{namespace}/servers/{podName}/logs/exists",
		ProducesMediaTypes: []string{"application/json", "text/x-log"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CheckServerLogsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *CheckServerLogsOK:
		return v, nil
	case *CheckServerLogsNotFound:
		return nil, v
	case *CheckServerLogsInternalServerError:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  DownloadServerLogs downloads dedicated server log files

  Required permission: ADMIN:NAMESPACE:{namespace}:DSLM:LOG [READ]

Required scope: social

This endpoint will download dedicated server&#39;s log file (.log).
*/
func (a *Client) DownloadServerLogs(params *DownloadServerLogsParams, authInfo runtime.ClientAuthInfoWriter) (*DownloadServerLogsOK, *DownloadServerLogsNotFound, *DownloadServerLogsInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDownloadServerLogsParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "downloadServerLogs",
		Method:             "GET",
		PathPattern:        "/dslogmanager/namespaces/{namespace}/servers/{podName}/logs/download",
		ProducesMediaTypes: []string{"application/json", "text/x-log"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DownloadServerLogsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}

	switch v := result.(type) {

	case *DownloadServerLogsOK:
		return v, nil, nil, nil

	case *DownloadServerLogsNotFound:
		return nil, v, nil, nil

	case *DownloadServerLogsInternalServerError:
		return nil, nil, v, nil

	default:
		return nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) DownloadServerLogsShort(params *DownloadServerLogsParams, authInfo runtime.ClientAuthInfoWriter) (*DownloadServerLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDownloadServerLogsParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "downloadServerLogs",
		Method:             "GET",
		PathPattern:        "/dslogmanager/namespaces/{namespace}/servers/{podName}/logs/download",
		ProducesMediaTypes: []string{"application/json", "text/x-log"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DownloadServerLogsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *DownloadServerLogsOK:
		return v, nil
	case *DownloadServerLogsNotFound:
		return nil, v
	case *DownloadServerLogsInternalServerError:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
