// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetNumberNumNotFoundBodyError get number num not found body error
// swagger:model getNumberNumNotFoundBodyError
type GetNumberNumNotFoundBodyError struct {

	// Error Message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this get number num not found body error
func (m *GetNumberNumNotFoundBodyError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetNumberNumNotFoundBodyError) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetNumberNumNotFoundBodyError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetNumberNumNotFoundBodyError) UnmarshalBinary(b []byte) error {
	var res GetNumberNumNotFoundBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
