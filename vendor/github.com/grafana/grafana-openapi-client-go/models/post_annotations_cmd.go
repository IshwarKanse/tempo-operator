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

// PostAnnotationsCmd post annotations cmd
//
// swagger:model PostAnnotationsCmd
type PostAnnotationsCmd struct {

	// dashboard Id
	DashboardID int64 `json:"dashboardId,omitempty"`

	// dashboard UID
	DashboardUID string `json:"dashboardUID,omitempty"`

	// data
	Data JSON `json:"data,omitempty"`

	// panel Id
	PanelID int64 `json:"panelId,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// text
	// Required: true
	Text *string `json:"text"`

	// time
	Time int64 `json:"time,omitempty"`

	// time end
	TimeEnd int64 `json:"timeEnd,omitempty"`
}

// Validate validates this post annotations cmd
func (m *PostAnnotationsCmd) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateText(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostAnnotationsCmd) validateText(formats strfmt.Registry) error {

	if err := validate.Required("text", "body", m.Text); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post annotations cmd based on context it is used
func (m *PostAnnotationsCmd) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostAnnotationsCmd) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostAnnotationsCmd) UnmarshalBinary(b []byte) error {
	var res PostAnnotationsCmd
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}