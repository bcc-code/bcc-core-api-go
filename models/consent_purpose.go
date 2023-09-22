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

// ConsentPurpose consent purpose
//
// swagger:model ConsentPurpose
type ConsentPurpose string

func NewConsentPurpose(value ConsentPurpose) *ConsentPurpose {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ConsentPurpose.
func (m ConsentPurpose) Pointer() *ConsentPurpose {
	return &m
}

const (

	// ConsentPurposeDataSharing captures enum value "DataSharing"
	ConsentPurposeDataSharing ConsentPurpose = "DataSharing"

	// ConsentPurposeTracking captures enum value "Tracking"
	ConsentPurposeTracking ConsentPurpose = "Tracking"
)

// for schema
var consentPurposeEnum []interface{}

func init() {
	var res []ConsentPurpose
	if err := json.Unmarshal([]byte(`["DataSharing","Tracking"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consentPurposeEnum = append(consentPurposeEnum, v)
	}
}

func (m ConsentPurpose) validateConsentPurposeEnum(path, location string, value ConsentPurpose) error {
	if err := validate.EnumCase(path, location, value, consentPurposeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this consent purpose
func (m ConsentPurpose) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateConsentPurposeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this consent purpose based on context it is used
func (m ConsentPurpose) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
