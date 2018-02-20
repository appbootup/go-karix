// Code generated by go-swagger; DO NOT EDIT.

package number

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new number API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for number API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteNumberNum unrents number from your account

Unrent number from your account
*/
func (a *Client) DeleteNumberNum(params *DeleteNumberNumParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNumberNumNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNumberNumParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteNumberNum",
		Method:             "DELETE",
		PathPattern:        "/number/{num}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteNumberNumReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteNumberNumNoContent), nil

}

/*
GetNumberNum gets details of a number

Get details of a number
*/
func (a *Client) GetNumberNum(params *GetNumberNumParams, authInfo runtime.ClientAuthInfoWriter) (*GetNumberNumOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNumberNumParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNumberNum",
		Method:             "GET",
		PathPattern:        "/number/{num}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNumberNumReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNumberNumOK), nil

}

/*
PatchNumberNum edits phone number belonging to your account

Edit phone number belonging to your account
*/
func (a *Client) PatchNumberNum(params *PatchNumberNumParams, authInfo runtime.ClientAuthInfoWriter) (*PatchNumberNumOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchNumberNumParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchNumberNum",
		Method:             "PATCH",
		PathPattern:        "/number/{num}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchNumberNumReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchNumberNumOK), nil

}

/*
GetNumber gets details of all phone numbers linked to your account

Get details of all phone numbers linked to your account.
*/
func (a *Client) GetNumber(params *GetNumberParams, authInfo runtime.ClientAuthInfoWriter) (*GetNumberOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNumberParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNumber",
		Method:             "GET",
		PathPattern:        "/number/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNumberReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNumberOK), nil

}

/*
RentNumber rents a phone number

Rent a phone number.
Refer to Number Search API to find available phone numbers

*/
func (a *Client) RentNumber(params *RentNumberParams, authInfo runtime.ClientAuthInfoWriter) (*RentNumberCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRentNumberParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "rentNumber",
		Method:             "POST",
		PathPattern:        "/number/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RentNumberReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RentNumberCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
