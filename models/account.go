// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Account account
// swagger:model Account
type Account struct {
	EditAccount

	AccountAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Account) UnmarshalJSON(raw []byte) error {

	var aO0 EditAccount
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.EditAccount = aO0

	var aO1 AccountAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.AccountAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Account) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.EditAccount)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.AccountAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this account
func (m *Account) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.EditAccount.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.AccountAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Account) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Account) UnmarshalBinary(b []byte) error {
	var res Account
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
