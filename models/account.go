// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/shopspring/decimal"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Account account
// swagger:model Account
type Account struct {
	EditAccount

	// Signifies whether the account is prepaid, postpaid or trial
	//
	// Enum: [prepaid postpaid trial]
	AccountType string `json:"account_type,omitempty"`

	// Whether auto-recharge has been enabled.
	//
	AutoRecharge bool `json:"auto_recharge,omitempty"`

	// Date when this account was created
	// Format: date-time
	CreatedTime strfmt.DateTime `json:"created_time,omitempty"`

	// Account's credit balance in US dollars.
	//   - For postpaid accounts this value will be `null`.
	//   - For subaccounts this value will reflect balance of parent account
	//
	CreditBalance decimal.Decimal `json:"credit_balance,omitempty"`

	// There is generally only one parent account in list. Rest are child subaccounts. Possible values:
	//   - `true`: If the account is a parent account
	//   - `false`: If the account is a subaccount of the parent account
	//
	IsParent bool `json:"is_parent,omitempty"`

	// Token password for this account. Used as password in Basic Authentication
	//
	Token string `json:"token,omitempty"`

	// Alphanumeric user identification. Used as username for Basic Authentication
	//
	UID string `json:"uid,omitempty"`

	// Date when this account was last updated
	// Format: date-time
	UpdatedTime strfmt.DateTime `json:"updated_time,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Account) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 EditAccount
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.EditAccount = aO0

	// AO1
	var dataAO1 struct {
		AccountType string `json:"account_type,omitempty"`

		AutoRecharge bool `json:"auto_recharge,omitempty"`

		CreatedTime strfmt.DateTime `json:"created_time,omitempty"`

		CreditBalance decimal.Decimal `json:"credit_balance,omitempty"`

		IsParent bool `json:"is_parent,omitempty"`

		Token string `json:"token,omitempty"`

		UID string `json:"uid,omitempty"`

		UpdatedTime strfmt.DateTime `json:"updated_time,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AccountType = dataAO1.AccountType

	m.AutoRecharge = dataAO1.AutoRecharge

	m.CreatedTime = dataAO1.CreatedTime

	m.CreditBalance = dataAO1.CreditBalance

	m.IsParent = dataAO1.IsParent

	m.Token = dataAO1.Token

	m.UID = dataAO1.UID

	m.UpdatedTime = dataAO1.UpdatedTime

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Account) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.EditAccount)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		AccountType string `json:"account_type,omitempty"`

		AutoRecharge bool `json:"auto_recharge,omitempty"`

		CreatedTime strfmt.DateTime `json:"created_time,omitempty"`

		CreditBalance decimal.Decimal `json:"credit_balance,omitempty"`

		IsParent bool `json:"is_parent,omitempty"`

		Token string `json:"token,omitempty"`

		UID string `json:"uid,omitempty"`

		UpdatedTime strfmt.DateTime `json:"updated_time,omitempty"`
	}

	dataAO1.AccountType = m.AccountType

	dataAO1.AutoRecharge = m.AutoRecharge

	dataAO1.CreatedTime = m.CreatedTime

	dataAO1.CreditBalance = m.CreditBalance

	dataAO1.IsParent = m.IsParent

	dataAO1.Token = m.Token

	dataAO1.UID = m.UID

	dataAO1.UpdatedTime = m.UpdatedTime

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this account
func (m *Account) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with EditAccount
	if err := m.EditAccount.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var accountTypeAccountTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["prepaid","postpaid","trial"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountTypeAccountTypePropEnum = append(accountTypeAccountTypePropEnum, v)
	}
}

// property enum
func (m *Account) validateAccountTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, accountTypeAccountTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Account) validateAccountType(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAccountTypeEnum("account_type", "body", m.AccountType); err != nil {
		return err
	}

	return nil
}

func (m *Account) validateCreatedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("created_time", "body", "date-time", m.CreatedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Account) validateUpdatedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_time", "body", "date-time", m.UpdatedTime.String(), formats); err != nil {
		return err
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
