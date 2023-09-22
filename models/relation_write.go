// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RelationWrite relation write
//
// swagger:model RelationWrite
type RelationWrite struct {

	// Permission of target on origin
	// Required: true
	GrantToOrigin GrantType `json:"grantToOrigin"`

	// Permission of origin on target
	// Required: true
	GrantToTarget GrantType `json:"grantToTarget"`

	// origin Uid
	// Required: true
	// Format: uuid
	OriginUID *strfmt.UUID `json:"originUid"`

	// target Uid
	// Required: true
	// Format: uuid
	TargetUID *strfmt.UUID `json:"targetUid"`

	// Type of the relation, defined as {origin} is {type} (of) {target}
	// Required: true
	Type RelationType `json:"type"`

	// valid from
	// Required: true
	// Format: date-time
	ValidFrom *strfmt.DateTime `json:"validFrom"`

	// valid to
	// Format: date-time
	ValidTo *strfmt.DateTime `json:"validTo,omitempty"`
}

// Validate validates this relation write
func (m *RelationWrite) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGrantToOrigin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGrantToTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RelationWrite) validateGrantToOrigin(formats strfmt.Registry) error {

	return nil
}

func (m *RelationWrite) validateGrantToTarget(formats strfmt.Registry) error {

	return nil
}

func (m *RelationWrite) validateOriginUID(formats strfmt.Registry) error {

	if err := validate.Required("originUid", "body", m.OriginUID); err != nil {
		return err
	}

	if err := validate.FormatOf("originUid", "body", "uuid", m.OriginUID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RelationWrite) validateTargetUID(formats strfmt.Registry) error {

	if err := validate.Required("targetUid", "body", m.TargetUID); err != nil {
		return err
	}

	if err := validate.FormatOf("targetUid", "body", "uuid", m.TargetUID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RelationWrite) validateType(formats strfmt.Registry) error {

	return nil
}

func (m *RelationWrite) validateValidFrom(formats strfmt.Registry) error {

	if err := validate.Required("validFrom", "body", m.ValidFrom); err != nil {
		return err
	}

	if err := validate.FormatOf("validFrom", "body", "date-time", m.ValidFrom.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RelationWrite) validateValidTo(formats strfmt.Registry) error {
	if swag.IsZero(m.ValidTo) { // not required
		return nil
	}

	if err := validate.FormatOf("validTo", "body", "date-time", m.ValidTo.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this relation write based on the context it is used
func (m *RelationWrite) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGrantToOrigin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGrantToTarget(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RelationWrite) contextValidateGrantToOrigin(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *RelationWrite) contextValidateGrantToTarget(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *RelationWrite) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *RelationWrite) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RelationWrite) UnmarshalBinary(b []byte) error {
	var res RelationWrite
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}