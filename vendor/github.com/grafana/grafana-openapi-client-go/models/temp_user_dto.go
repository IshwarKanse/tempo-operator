// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TempUserDTO temp user DTO
//
// swagger:model TempUserDTO
type TempUserDTO struct {

	// code
	Code string `json:"code,omitempty"`

	// created on
	// Format: date-time
	CreatedOn strfmt.DateTime `json:"createdOn,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// email sent
	EmailSent bool `json:"emailSent,omitempty"`

	// email sent on
	// Format: date-time
	EmailSentOn strfmt.DateTime `json:"emailSentOn,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// invited by email
	InvitedByEmail string `json:"invitedByEmail,omitempty"`

	// invited by login
	InvitedByLogin string `json:"invitedByLogin,omitempty"`

	// invited by name
	InvitedByName string `json:"invitedByName,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// org Id
	OrgID int64 `json:"orgId,omitempty"`

	// role
	// Enum: [None Viewer Editor Admin]
	Role string `json:"role,omitempty"`

	// status
	Status TempUserStatus `json:"status,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this temp user DTO
func (m *TempUserDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmailSentOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TempUserDTO) validateCreatedOn(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("createdOn", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TempUserDTO) validateEmailSentOn(formats strfmt.Registry) error {
	if swag.IsZero(m.EmailSentOn) { // not required
		return nil
	}

	if err := validate.FormatOf("emailSentOn", "body", "date-time", m.EmailSentOn.String(), formats); err != nil {
		return err
	}

	return nil
}

var tempUserDtoTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","Viewer","Editor","Admin"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tempUserDtoTypeRolePropEnum = append(tempUserDtoTypeRolePropEnum, v)
	}
}

const (

	// TempUserDTORoleNone captures enum value "None"
	TempUserDTORoleNone string = "None"

	// TempUserDTORoleViewer captures enum value "Viewer"
	TempUserDTORoleViewer string = "Viewer"

	// TempUserDTORoleEditor captures enum value "Editor"
	TempUserDTORoleEditor string = "Editor"

	// TempUserDTORoleAdmin captures enum value "Admin"
	TempUserDTORoleAdmin string = "Admin"
)

// prop value enum
func (m *TempUserDTO) validateRoleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, tempUserDtoTypeRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TempUserDTO) validateRole(formats strfmt.Registry) error {
	if swag.IsZero(m.Role) { // not required
		return nil
	}

	// value enum
	if err := m.validateRoleEnum("role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

func (m *TempUserDTO) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

// ContextValidate validate this temp user DTO based on the context it is used
func (m *TempUserDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TempUserDTO) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TempUserDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TempUserDTO) UnmarshalBinary(b []byte) error {
	var res TempUserDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
