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

// MaritalStatus marital status
//
// swagger:model MaritalStatus
type MaritalStatus string

func NewMaritalStatus(value MaritalStatus) *MaritalStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated MaritalStatus.
func (m MaritalStatus) Pointer() *MaritalStatus {
	return &m
}

const (

	// MaritalStatusSingle captures enum value "Single"
	MaritalStatusSingle MaritalStatus = "Single"

	// MaritalStatusMarried captures enum value "Married"
	MaritalStatusMarried MaritalStatus = "Married"

	// MaritalStatusWidowed captures enum value "Widowed"
	MaritalStatusWidowed MaritalStatus = "Widowed"

	// MaritalStatusSeparated captures enum value "Separated"
	MaritalStatusSeparated MaritalStatus = "Separated"

	// MaritalStatusSingleParent captures enum value "SingleParent"
	MaritalStatusSingleParent MaritalStatus = "SingleParent"

	// MaritalStatusUnknown captures enum value "Unknown"
	MaritalStatusUnknown MaritalStatus = "Unknown"
)

// for schema
var maritalStatusEnum []interface{}

func init() {
	var res []MaritalStatus
	if err := json.Unmarshal([]byte(`["Single","Married","Widowed","Separated","SingleParent","Unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		maritalStatusEnum = append(maritalStatusEnum, v)
	}
}

func (m MaritalStatus) validateMaritalStatusEnum(path, location string, value MaritalStatus) error {
	if err := validate.EnumCase(path, location, value, maritalStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this marital status
func (m MaritalStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateMaritalStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this marital status based on context it is used
func (m MaritalStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
