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

// Relation relation
//
// swagger:model Relation
type Relation struct {

	// Permission of target on origin
	// Required: true
	GrantToOrigin GrantType `json:"grantToOrigin"`

	// Permission of origin on target
	// Required: true
	GrantToTarget GrantType `json:"grantToTarget"`

	// last changed date
	// Required: true
	// Format: date-time
	LastChangedDate *strfmt.DateTime `json:"lastChangedDate"`

	// origin
	Origin *Person `json:"origin,omitempty"`

	// origin Uid
	// Required: true
	// Format: uuid
	OriginUID *strfmt.UUID `json:"originUid"`

	// target
	Target *Person `json:"target,omitempty"`

	// target Uid
	// Required: true
	// Format: uuid
	TargetUID *strfmt.UUID `json:"targetUid"`

	// Type of the relation, defined as {origin} is {type} (of) {target}
	// Required: true
	Type RelationType `json:"type"`

	// uid
	// Required: true
	// Format: uuid
	UID *strfmt.UUID `json:"uid"`

	// valid from
	// Required: true
	// Format: date-time
	ValidFrom *strfmt.DateTime `json:"validFrom"`

	// valid to
	// Format: date-time
	ValidTo *strfmt.DateTime `json:"validTo,omitempty"`
}

// Validate validates this relation
func (m *Relation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGrantToOrigin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGrantToTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastChangedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrigin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUID(formats); err != nil {
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

func (m *Relation) validateGrantToOrigin(formats strfmt.Registry) error {

	return nil
}

func (m *Relation) validateGrantToTarget(formats strfmt.Registry) error {

	return nil
}

func (m *Relation) validateLastChangedDate(formats strfmt.Registry) error {

	if err := validate.Required("lastChangedDate", "body", m.LastChangedDate); err != nil {
		return err
	}

	if err := validate.FormatOf("lastChangedDate", "body", "date-time", m.LastChangedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Relation) validateOrigin(formats strfmt.Registry) error {
	if swag.IsZero(m.Origin) { // not required
		return nil
	}

	if m.Origin != nil {
		if err := m.Origin.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("origin")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("origin")
			}
			return err
		}
	}

	return nil
}

func (m *Relation) validateOriginUID(formats strfmt.Registry) error {

	if err := validate.Required("originUid", "body", m.OriginUID); err != nil {
		return err
	}

	if err := validate.FormatOf("originUid", "body", "uuid", m.OriginUID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Relation) validateTarget(formats strfmt.Registry) error {
	if swag.IsZero(m.Target) { // not required
		return nil
	}

	if m.Target != nil {
		if err := m.Target.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

func (m *Relation) validateTargetUID(formats strfmt.Registry) error {

	if err := validate.Required("targetUid", "body", m.TargetUID); err != nil {
		return err
	}

	if err := validate.FormatOf("targetUid", "body", "uuid", m.TargetUID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Relation) validateType(formats strfmt.Registry) error {

	return nil
}

func (m *Relation) validateUID(formats strfmt.Registry) error {

	if err := validate.Required("uid", "body", m.UID); err != nil {
		return err
	}

	if err := validate.FormatOf("uid", "body", "uuid", m.UID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Relation) validateValidFrom(formats strfmt.Registry) error {

	if err := validate.Required("validFrom", "body", m.ValidFrom); err != nil {
		return err
	}

	if err := validate.FormatOf("validFrom", "body", "date-time", m.ValidFrom.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Relation) validateValidTo(formats strfmt.Registry) error {
	if swag.IsZero(m.ValidTo) { // not required
		return nil
	}

	if err := validate.FormatOf("validTo", "body", "date-time", m.ValidTo.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this relation based on the context it is used
func (m *Relation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGrantToOrigin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGrantToTarget(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrigin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTarget(ctx, formats); err != nil {
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

func (m *Relation) contextValidateGrantToOrigin(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *Relation) contextValidateGrantToTarget(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *Relation) contextValidateOrigin(ctx context.Context, formats strfmt.Registry) error {

	if m.Origin != nil {

		if swag.IsZero(m.Origin) { // not required
			return nil
		}

		if err := m.Origin.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("origin")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("origin")
			}
			return err
		}
	}

	return nil
}

func (m *Relation) contextValidateTarget(ctx context.Context, formats strfmt.Registry) error {

	if m.Target != nil {

		if swag.IsZero(m.Target) { // not required
			return nil
		}

		if err := m.Target.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

func (m *Relation) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *Relation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Relation) UnmarshalBinary(b []byte) error {
	var res Relation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
