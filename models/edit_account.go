// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EditAccount edit account
// swagger:model EditAccount
type EditAccount struct {

	// Name of the account. Must be unique within the parent account.
	//
	// Max Length: 200
	Name string `json:"name,omitempty"`

	// Status of your account. Possible values are:
	//   - enabled: Account is ready to be used
	//   - suspended: Account has been temporarily suspended
	//   - disabled: Account has been permanently disabled and
	//             cannot be revived. All resources allocated
	//             to the subaccount like phonenumbers are also
	//             deleted.
	//
	Status string `json:"status,omitempty"`
}

// Validate validates this edit account
func (m *EditAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EditAccount) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 200); err != nil {
		return err
	}

	return nil
}

var editAccountTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","suspended","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		editAccountTypeStatusPropEnum = append(editAccountTypeStatusPropEnum, v)
	}
}

const (

	// EditAccountStatusEnabled captures enum value "enabled"
	EditAccountStatusEnabled string = "enabled"

	// EditAccountStatusSuspended captures enum value "suspended"
	EditAccountStatusSuspended string = "suspended"

	// EditAccountStatusDisabled captures enum value "disabled"
	EditAccountStatusDisabled string = "disabled"
)

// prop value enum
func (m *EditAccount) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, editAccountTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *EditAccount) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EditAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EditAccount) UnmarshalBinary(b []byte) error {
	var res EditAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
