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

// SearchVisibility search visibility
//
// swagger:model SearchVisibility
type SearchVisibility string

func NewSearchVisibility(value SearchVisibility) *SearchVisibility {
	return &value
}

// Pointer returns a pointer to a freshly-allocated SearchVisibility.
func (m SearchVisibility) Pointer() *SearchVisibility {
	return &m
}

const (

	// SearchVisibilityGlobal captures enum value "Global"
	SearchVisibilityGlobal SearchVisibility = "Global"

	// SearchVisibilityDistrict captures enum value "District"
	SearchVisibilityDistrict SearchVisibility = "District"

	// SearchVisibilityHidden captures enum value "Hidden"
	SearchVisibilityHidden SearchVisibility = "Hidden"
)

// for schema
var searchVisibilityEnum []interface{}

func init() {
	var res []SearchVisibility
	if err := json.Unmarshal([]byte(`["Global","District","Hidden"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		searchVisibilityEnum = append(searchVisibilityEnum, v)
	}
}

func (m SearchVisibility) validateSearchVisibilityEnum(path, location string, value SearchVisibility) error {
	if err := validate.EnumCase(path, location, value, searchVisibilityEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this search visibility
func (m SearchVisibility) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSearchVisibilityEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this search visibility based on context it is used
func (m SearchVisibility) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
