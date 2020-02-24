// Code generated by go-swagger; DO NOT EDIT.

package introducers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new introducers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for introducers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AddIntroducerOwner(params *AddIntroducerOwnerParams, authInfo runtime.ClientAuthInfoWriter) (*AddIntroducerOwnerOK, error)

	CreateIntroducer(params *CreateIntroducerParams, authInfo runtime.ClientAuthInfoWriter) (*CreateIntroducerOK, error)

	GetIntroducer(params *GetIntroducerParams, authInfo runtime.ClientAuthInfoWriter) (*GetIntroducerOK, error)

	GetIntroducers(params *GetIntroducersParams, authInfo runtime.ClientAuthInfoWriter) (*GetIntroducersOK, error)

	RemoveIntroducerOwner(params *RemoveIntroducerOwnerParams, authInfo runtime.ClientAuthInfoWriter) (*RemoveIntroducerOwnerOK, error)

	UpdateIntroducerAddress(params *UpdateIntroducerAddressParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateIntroducerAddressOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddIntroducerOwner add introducer owner API
*/
func (a *Client) AddIntroducerOwner(params *AddIntroducerOwnerParams, authInfo runtime.ClientAuthInfoWriter) (*AddIntroducerOwnerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddIntroducerOwnerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddIntroducerOwner",
		Method:             "POST",
		PathPattern:        "/v1/introducers/{introducer_id}/owners",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddIntroducerOwnerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddIntroducerOwnerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AddIntroducerOwner: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateIntroducer create introducer API
*/
func (a *Client) CreateIntroducer(params *CreateIntroducerParams, authInfo runtime.ClientAuthInfoWriter) (*CreateIntroducerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateIntroducerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateIntroducer",
		Method:             "POST",
		PathPattern:        "/v1/introducers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateIntroducerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateIntroducerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateIntroducer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetIntroducer get introducer API
*/
func (a *Client) GetIntroducer(params *GetIntroducerParams, authInfo runtime.ClientAuthInfoWriter) (*GetIntroducerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIntroducerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetIntroducer",
		Method:             "GET",
		PathPattern:        "/v1/introducers/{introducer_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIntroducerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIntroducerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetIntroducer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetIntroducers gets introducers will perform a search for introducers with the given query params see get introducers request
*/
func (a *Client) GetIntroducers(params *GetIntroducersParams, authInfo runtime.ClientAuthInfoWriter) (*GetIntroducersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIntroducersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetIntroducers",
		Method:             "GET",
		PathPattern:        "/v1/introducers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIntroducersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIntroducersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetIntroducers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RemoveIntroducerOwner remove introducer owner API
*/
func (a *Client) RemoveIntroducerOwner(params *RemoveIntroducerOwnerParams, authInfo runtime.ClientAuthInfoWriter) (*RemoveIntroducerOwnerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveIntroducerOwnerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RemoveIntroducerOwner",
		Method:             "DELETE",
		PathPattern:        "/v1/introducers/{introducer_id}/owners/{account_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RemoveIntroducerOwnerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RemoveIntroducerOwnerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RemoveIntroducerOwner: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateIntroducerAddress update introducer address API
*/
func (a *Client) UpdateIntroducerAddress(params *UpdateIntroducerAddressParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateIntroducerAddressOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateIntroducerAddressParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateIntroducerAddress",
		Method:             "POST",
		PathPattern:        "/v1/introducers/{introducer_id}/update_address",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateIntroducerAddressReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateIntroducerAddressOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateIntroducerAddress: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
