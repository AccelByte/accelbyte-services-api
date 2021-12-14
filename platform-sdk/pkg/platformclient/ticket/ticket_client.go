// Code generated by go-swagger; DO NOT EDIT.

package ticket

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new ticket API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ticket API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AcquireUserTicket(params *AcquireUserTicketParams, authInfo runtime.ClientAuthInfoWriter) (*AcquireUserTicketOK, *AcquireUserTicketNotFound, *AcquireUserTicketConflict, *AcquireUserTicketUnprocessableEntity, error)
	AcquireUserTicketShort(params *AcquireUserTicketParams, authInfo runtime.ClientAuthInfoWriter) (*AcquireUserTicketOK, error)
	DecreaseTicketSale(params *DecreaseTicketSaleParams, authInfo runtime.ClientAuthInfoWriter) (*DecreaseTicketSaleNoContent, *DecreaseTicketSaleNotFound, *DecreaseTicketSaleUnprocessableEntity, error)
	DecreaseTicketSaleShort(params *DecreaseTicketSaleParams, authInfo runtime.ClientAuthInfoWriter) (*DecreaseTicketSaleNoContent, error)
	GetTicketBoothID(params *GetTicketBoothIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetTicketBoothIDOK, *GetTicketBoothIDNotFound, error)
	GetTicketBoothIDShort(params *GetTicketBoothIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetTicketBoothIDOK, error)
	GetTicketDynamic(params *GetTicketDynamicParams, authInfo runtime.ClientAuthInfoWriter) (*GetTicketDynamicOK, *GetTicketDynamicNotFound, error)
	GetTicketDynamicShort(params *GetTicketDynamicParams, authInfo runtime.ClientAuthInfoWriter) (*GetTicketDynamicOK, error)
	IncreaseTicketSale(params *IncreaseTicketSaleParams, authInfo runtime.ClientAuthInfoWriter) (*IncreaseTicketSaleOK, *IncreaseTicketSaleNotFound, *IncreaseTicketSaleUnprocessableEntity, error)
	IncreaseTicketSaleShort(params *IncreaseTicketSaleParams, authInfo runtime.ClientAuthInfoWriter) (*IncreaseTicketSaleOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AcquireUserTicket acquires ticket

  &lt;b&gt;[SERVICE COMMUNICATION ONLY]&lt;/b&gt; Acquire ticket(code/key) based on booth name.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:{namespace}:USER:{userId}:TICKET&#34;, action=1 (CREATE)&lt;/li&gt;&lt;li&gt;&lt;i&gt;Returns&lt;/i&gt;: acquire result&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) AcquireUserTicket(params *AcquireUserTicketParams, authInfo runtime.ClientAuthInfoWriter) (*AcquireUserTicketOK, *AcquireUserTicketNotFound, *AcquireUserTicketConflict, *AcquireUserTicketUnprocessableEntity, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAcquireUserTicketParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "acquireUserTicket",
		Method:             "POST",
		PathPattern:        "/admin/namespaces/{namespace}/users/{userId}/tickets/{boothName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AcquireUserTicketReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, nil, err
	}

	switch v := result.(type) {

	case *AcquireUserTicketOK:
		return v, nil, nil, nil, nil

	case *AcquireUserTicketNotFound:
		return nil, v, nil, nil, nil

	case *AcquireUserTicketConflict:
		return nil, nil, v, nil, nil

	case *AcquireUserTicketUnprocessableEntity:
		return nil, nil, nil, v, nil

	default:
		return nil, nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) AcquireUserTicketShort(params *AcquireUserTicketParams, authInfo runtime.ClientAuthInfoWriter) (*AcquireUserTicketOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAcquireUserTicketParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "acquireUserTicket",
		Method:             "POST",
		PathPattern:        "/admin/namespaces/{namespace}/users/{userId}/tickets/{boothName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AcquireUserTicketReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *AcquireUserTicketOK:
		return v, nil
	case *AcquireUserTicketNotFound:
		return nil, v
	case *AcquireUserTicketConflict:
		return nil, v
	case *AcquireUserTicketUnprocessableEntity:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  DecreaseTicketSale decreases ticket sale

  &lt;b&gt;[SERVICE COMMUNICATION ONLY]&lt;/b&gt; Decrease ticket(code/key) sale if requested orderNo is already increased.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:{namespace}:TICKET&#34;, action=4 (UPDATE)&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) DecreaseTicketSale(params *DecreaseTicketSaleParams, authInfo runtime.ClientAuthInfoWriter) (*DecreaseTicketSaleNoContent, *DecreaseTicketSaleNotFound, *DecreaseTicketSaleUnprocessableEntity, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDecreaseTicketSaleParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "decreaseTicketSale",
		Method:             "PUT",
		PathPattern:        "/admin/namespaces/{namespace}/tickets/{boothName}/decrement",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DecreaseTicketSaleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}

	switch v := result.(type) {

	case *DecreaseTicketSaleNoContent:
		return v, nil, nil, nil

	case *DecreaseTicketSaleNotFound:
		return nil, v, nil, nil

	case *DecreaseTicketSaleUnprocessableEntity:
		return nil, nil, v, nil

	default:
		return nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) DecreaseTicketSaleShort(params *DecreaseTicketSaleParams, authInfo runtime.ClientAuthInfoWriter) (*DecreaseTicketSaleNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDecreaseTicketSaleParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "decreaseTicketSale",
		Method:             "PUT",
		PathPattern:        "/admin/namespaces/{namespace}/tickets/{boothName}/decrement",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DecreaseTicketSaleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *DecreaseTicketSaleNoContent:
		return v, nil
	case *DecreaseTicketSaleNotFound:
		return nil, v
	case *DecreaseTicketSaleUnprocessableEntity:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetTicketBoothID gets ticket booth ID

  Get ticket(code/key) booth ID.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:{namespace}:TICKET&#34;, action=2 (READ)&lt;/li&gt;&lt;li&gt;&lt;i&gt;Returns&lt;/i&gt;: ticket booth id&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) GetTicketBoothID(params *GetTicketBoothIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetTicketBoothIDOK, *GetTicketBoothIDNotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTicketBoothIDParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTicketBoothID",
		Method:             "GET",
		PathPattern:        "/admin/namespaces/{namespace}/tickets/{boothName}/id",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTicketBoothIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *GetTicketBoothIDOK:
		return v, nil, nil

	case *GetTicketBoothIDNotFound:
		return nil, v, nil

	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) GetTicketBoothIDShort(params *GetTicketBoothIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetTicketBoothIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTicketBoothIDParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTicketBoothID",
		Method:             "GET",
		PathPattern:        "/admin/namespaces/{namespace}/tickets/{boothName}/id",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTicketBoothIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *GetTicketBoothIDOK:
		return v, nil
	case *GetTicketBoothIDNotFound:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  GetTicketDynamic gets ticket dynamic

  &lt;b&gt;[SERVICE COMMUNICATION ONLY]&lt;/b&gt; Get ticket(code/key) dynamic based on booth name.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:{namespace}:TICKET&#34;, action=2 (READ)&lt;/li&gt;&lt;li&gt;&lt;i&gt;Returns&lt;/i&gt;: ticket dynamic&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) GetTicketDynamic(params *GetTicketDynamicParams, authInfo runtime.ClientAuthInfoWriter) (*GetTicketDynamicOK, *GetTicketDynamicNotFound, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTicketDynamicParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTicketDynamic",
		Method:             "GET",
		PathPattern:        "/admin/namespaces/{namespace}/tickets/{boothName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTicketDynamicReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}

	switch v := result.(type) {

	case *GetTicketDynamicOK:
		return v, nil, nil

	case *GetTicketDynamicNotFound:
		return nil, v, nil

	default:
		return nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) GetTicketDynamicShort(params *GetTicketDynamicParams, authInfo runtime.ClientAuthInfoWriter) (*GetTicketDynamicOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTicketDynamicParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTicketDynamic",
		Method:             "GET",
		PathPattern:        "/admin/namespaces/{namespace}/tickets/{boothName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTicketDynamicReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *GetTicketDynamicOK:
		return v, nil
	case *GetTicketDynamicNotFound:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

/*
  IncreaseTicketSale increases ticket sale

  &lt;b&gt;[SERVICE COMMUNICATION ONLY]&lt;/b&gt; increase ticket(code/key) sale.&lt;br&gt;Other detail info: &lt;ul&gt;&lt;li&gt;&lt;i&gt;Required permission&lt;/i&gt;: resource=&#34;ADMIN:NAMESPACE:{namespace}:TICKET&#34;, action=4 (UPDATE)&lt;/li&gt;&lt;li&gt;&lt;i&gt;Returns&lt;/i&gt;: Ticket sale increment result&lt;/li&gt;&lt;/ul&gt;
*/
func (a *Client) IncreaseTicketSale(params *IncreaseTicketSaleParams, authInfo runtime.ClientAuthInfoWriter) (*IncreaseTicketSaleOK, *IncreaseTicketSaleNotFound, *IncreaseTicketSaleUnprocessableEntity, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIncreaseTicketSaleParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "increaseTicketSale",
		Method:             "PUT",
		PathPattern:        "/admin/namespaces/{namespace}/tickets/{boothName}/increment",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IncreaseTicketSaleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}

	switch v := result.(type) {

	case *IncreaseTicketSaleOK:
		return v, nil, nil, nil

	case *IncreaseTicketSaleNotFound:
		return nil, v, nil, nil

	case *IncreaseTicketSaleUnprocessableEntity:
		return nil, nil, v, nil

	default:
		return nil, nil, nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

func (a *Client) IncreaseTicketSaleShort(params *IncreaseTicketSaleParams, authInfo runtime.ClientAuthInfoWriter) (*IncreaseTicketSaleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIncreaseTicketSaleParams()
	}

	if params.Context == nil {
		params.Context = context.Background()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "increaseTicketSale",
		Method:             "PUT",
		PathPattern:        "/admin/namespaces/{namespace}/tickets/{boothName}/increment",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IncreaseTicketSaleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch v := result.(type) {

	case *IncreaseTicketSaleOK:
		return v, nil
	case *IncreaseTicketSaleNotFound:
		return nil, v
	case *IncreaseTicketSaleUnprocessableEntity:
		return nil, v

	default:
		return nil, fmt.Errorf("Unexpected Type %v", reflect.TypeOf(v))
	}
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
