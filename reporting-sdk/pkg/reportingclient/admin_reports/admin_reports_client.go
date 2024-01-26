// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated ; DO NOT EDIT.

package admin_reports

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new admin reports API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for admin reports API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ListReports(params *ListReportsParams, authInfo runtime.ClientAuthInfoWriter) (*ListReportsOK, *ListReportsInternalServerError, error)
	ListReportsShort(params *ListReportsParams, authInfo runtime.ClientAuthInfoWriter) (*ListReportsOK, error)
	AdminSubmitReport(params *AdminSubmitReportParams, authInfo runtime.ClientAuthInfoWriter) (*AdminSubmitReportCreated, *AdminSubmitReportBadRequest, *AdminSubmitReportConflict, *AdminSubmitReportInternalServerError, error)
	AdminSubmitReportShort(params *AdminSubmitReportParams, authInfo runtime.ClientAuthInfoWriter) (*AdminSubmitReportCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Deprecated: 2022-08-10 - Use ListReportsShort instead.

ListReports list reports
Required permission: ADMIN:NAMESPACE:{namespace}:TICKET [READ]
Reports list can be ordered by:
- createdAt
- updatedAt
*/
func (a *Client) ListReports(params *ListReportsParams, authInfo runtime.ClientAuthInfoWriter) (*ListReportsOK, *ListReportsInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListReportsParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	if params.XFlightId != nil {
		params.SetFlightId(*params.XFlightId)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listReports",
		Method:             "GET",
		PathPattern:        "/reporting/v1/admin/namespaces/{namespace}/reports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListReportsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *ListReportsOK:
		return v, nil, nil

	case *ListReportsInternalServerError:
		return nil, v, nil

	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
ListReportsShort list reports
Required permission: ADMIN:NAMESPACE:{namespace}:TICKET [READ]
Reports list can be ordered by:
- createdAt
- updatedAt
*/
func (a *Client) ListReportsShort(params *ListReportsParams, authInfo runtime.ClientAuthInfoWriter) (*ListReportsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListReportsParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listReports",
		Method:             "GET",
		PathPattern:        "/reporting/v1/admin/namespaces/{namespace}/reports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListReportsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *ListReportsOK:
		return v, nil
	case *ListReportsInternalServerError:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
Deprecated: 2022-08-10 - Use AdminSubmitReportShort instead.

AdminSubmitReport submit a report by admin
Submit a report and will return ticket for reported object.
New ticket will be created if no OPEN ticket present for reported object (based by objectId and objectType) in a namespace.

Admin can only submit report once for each different user / object reported in the same OPEN ticket.
Reporting the same user / object in the same OPEN ticket will return HTTP code 409 (conflict).

Fill the 'reason' field with a 'reason title'
Supported category: - UGC - USER - CHAT - EXTENSION
*/
func (a *Client) AdminSubmitReport(params *AdminSubmitReportParams, authInfo runtime.ClientAuthInfoWriter) (*AdminSubmitReportCreated, *AdminSubmitReportBadRequest, *AdminSubmitReportConflict, *AdminSubmitReportInternalServerError, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminSubmitReportParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	if params.XFlightId != nil {
		params.SetFlightId(*params.XFlightId)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "adminSubmitReport",
		Method:             "POST",
		PathPattern:        "/reporting/v1/admin/namespaces/{namespace}/reports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AdminSubmitReportReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *AdminSubmitReportCreated:
		return v, nil, nil, nil, nil

	case *AdminSubmitReportBadRequest:
		return nil, v, nil, nil, nil

	case *AdminSubmitReportConflict:
		return nil, nil, v, nil, nil

	case *AdminSubmitReportInternalServerError:
		return nil, nil, nil, v, nil

	default:
		return nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
AdminSubmitReportShort submit a report by admin
Submit a report and will return ticket for reported object.
New ticket will be created if no OPEN ticket present for reported object (based by objectId and objectType) in a namespace.

Admin can only submit report once for each different user / object reported in the same OPEN ticket.
Reporting the same user / object in the same OPEN ticket will return HTTP code 409 (conflict).

Fill the 'reason' field with a 'reason title'
Supported category: - UGC - USER - CHAT - EXTENSION
*/
func (a *Client) AdminSubmitReportShort(params *AdminSubmitReportParams, authInfo runtime.ClientAuthInfoWriter) (*AdminSubmitReportCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdminSubmitReportParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	if params.RetryPolicy != nil {
		params.SetHTTPClientTransport(params.RetryPolicy)
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "adminSubmitReport",
		Method:             "POST",
		PathPattern:        "/reporting/v1/admin/namespaces/{namespace}/reports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AdminSubmitReportReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *AdminSubmitReportCreated:
		return v, nil
	case *AdminSubmitReportBadRequest:
		return nil, v
	case *AdminSubmitReportConflict:
		return nil, v
	case *AdminSubmitReportInternalServerError:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
