// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ArrayMetaResponse array meta response
// swagger:model ArrayMetaResponse
type ArrayMetaResponse struct {
	MetaResponse

	ArrayMetaResponseAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ArrayMetaResponse) UnmarshalJSON(raw []byte) error {

	var aO0 MetaResponse
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MetaResponse = aO0

	var aO1 ArrayMetaResponseAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.ArrayMetaResponseAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ArrayMetaResponse) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.MetaResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.ArrayMetaResponseAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this array meta response
func (m *ArrayMetaResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.MetaResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.ArrayMetaResponseAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ArrayMetaResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArrayMetaResponse) UnmarshalBinary(b []byte) error {
	var res ArrayMetaResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
