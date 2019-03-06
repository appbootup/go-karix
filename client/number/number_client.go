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
DeleteNumber unrents number from your account

Unrent number from your account
*/
func (a *Client) DeleteNumber(params *DeleteNumberParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNumberNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNumberParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteNumber",
		Method:             "DELETE",
		PathPattern:        "/number/{num}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteNumberReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteNumberNoContent), nil

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
GetNumberDetails gets details of a number

Get details of a number
*/
func (a *Client) GetNumberDetails(params *GetNumberDetailsParams, authInfo runtime.ClientAuthInfoWriter) (*GetNumberDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNumberDetailsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNumberDetails",
		Method:             "GET",
		PathPattern:        "/number/{num}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNumberDetailsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNumberDetailsOK), nil

}

/*
PatchNumberDetails edits phone number belonging to your account

Edit phone number belonging to your account
*/
func (a *Client) PatchNumberDetails(params *PatchNumberDetailsParams, authInfo runtime.ClientAuthInfoWriter) (*PatchNumberDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchNumberDetailsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchNumberDetails",
		Method:             "PATCH",
		PathPattern:        "/number/{num}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchNumberDetailsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchNumberDetailsOK), nil

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
