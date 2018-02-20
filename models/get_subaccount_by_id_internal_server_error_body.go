// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetSubaccountByIDInternalServerErrorBody get subaccount by Id internal server error body
// swagger:model getSubaccountByIdInternalServerErrorBody
type GetSubaccountByIDInternalServerErrorBody struct {

	// error
	Error *GetSubaccountByIDInternalServerErrorBodyError `json:"error,omitempty"`

	// meta
	Meta *MetaResponse `json:"meta,omitempty"`
}

// Validate validates this get subaccount by Id internal server error body
func (m *GetSubaccountByIDInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMeta(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetSubaccountByIDInternalServerErrorBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {

		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			}
			return err
		}

	}

	return nil
}

func (m *GetSubaccountByIDInternalServerErrorBody) validateMeta(formats strfmt.Registry) error {

	if swag.IsZero(m.Meta) { // not required
		return nil
	}

	if m.Meta != nil {

		if err := m.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetSubaccountByIDInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetSubaccountByIDInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetSubaccountByIDInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
