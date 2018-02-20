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

// GetNumberNumForbiddenBodyError get number num forbidden body error
// swagger:model getNumberNumForbiddenBodyError
type GetNumberNumForbiddenBodyError struct {

	// Forbidden Message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this get number num forbidden body error
func (m *GetNumberNumForbiddenBodyError) Validate(formats strfmt.Registry) error {
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

func (m *GetNumberNumForbiddenBodyError) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetNumberNumForbiddenBodyError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetNumberNumForbiddenBodyError) UnmarshalBinary(b []byte) error {
	var res GetNumberNumForbiddenBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}