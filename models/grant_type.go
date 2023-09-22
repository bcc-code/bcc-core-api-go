// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// GrantType grant type
//
// swagger:model GrantType
type GrantType string

func NewGrantType(value GrantType) *GrantType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated GrantType.
func (m GrantType) Pointer() *GrantType {
	return &m
}

const (

	// GrantTypeDefault captures enum value "Default"
	GrantTypeDefault GrantType = "Default"

	// GrantTypeView captures enum value "View"
	GrantTypeView GrantType = "View"

	// GrantTypeAdministrate captures enum value "Administrate"
	GrantTypeAdministrate GrantType = "Administrate"

	// GrantTypeRepresent captures enum value "Represent"
	GrantTypeRepresent GrantType = "Represent"

	// GrantTypeNone captures enum value "None"
	GrantTypeNone GrantType = "None"
)

// for schema
var grantTypeEnum []interface{}

func init() {
	var res []GrantType
	if err := json.Unmarshal([]byte(`["Default","View","Administrate","Represent","None"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		grantTypeEnum = append(grantTypeEnum, v)
	}
}

func (m GrantType) validateGrantTypeEnum(path, location string, value GrantType) error {
	if err := validate.EnumCase(path, location, value, grantTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this grant type
func (m GrantType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateGrantTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this grant type based on context it is used
func (m GrantType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}