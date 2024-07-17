// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GettableAPIAlertingConfig gettable Api alerting config
//
// swagger:model GettableApiAlertingConfig
type GettableAPIAlertingConfig struct {

	// global
	Global *GlobalConfig `json:"global,omitempty"`

	// inhibit rules
	InhibitRules []*InhibitRule `json:"inhibit_rules"`

	// mute time provenances
	MuteTimeProvenances map[string]Provenance `json:"muteTimeProvenances,omitempty"`

	// mute time intervals
	MuteTimeIntervals []*MuteTimeInterval `json:"mute_time_intervals"`

	// Override with our superset receiver type
	Receivers []*GettableAPIReceiver `json:"receivers"`

	// route
	Route *Route `json:"route,omitempty"`

	// templates
	Templates []string `json:"templates"`
}

// Validate validates this gettable Api alerting config
func (m *GettableAPIAlertingConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGlobal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInhibitRules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMuteTimeProvenances(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMuteTimeIntervals(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReceivers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoute(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GettableAPIAlertingConfig) validateGlobal(formats strfmt.Registry) error {
	if swag.IsZero(m.Global) { // not required
		return nil
	}

	if m.Global != nil {
		if err := m.Global.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("global")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("global")
			}
			return err
		}
	}

	return nil
}

func (m *GettableAPIAlertingConfig) validateInhibitRules(formats strfmt.Registry) error {
	if swag.IsZero(m.InhibitRules) { // not required
		return nil
	}

	for i := 0; i < len(m.InhibitRules); i++ {
		if swag.IsZero(m.InhibitRules[i]) { // not required
			continue
		}

		if m.InhibitRules[i] != nil {
			if err := m.InhibitRules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inhibit_rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inhibit_rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GettableAPIAlertingConfig) validateMuteTimeProvenances(formats strfmt.Registry) error {
	if swag.IsZero(m.MuteTimeProvenances) { // not required
		return nil
	}

	for k := range m.MuteTimeProvenances {

		if val, ok := m.MuteTimeProvenances[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("muteTimeProvenances" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("muteTimeProvenances" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *GettableAPIAlertingConfig) validateMuteTimeIntervals(formats strfmt.Registry) error {
	if swag.IsZero(m.MuteTimeIntervals) { // not required
		return nil
	}

	for i := 0; i < len(m.MuteTimeIntervals); i++ {
		if swag.IsZero(m.MuteTimeIntervals[i]) { // not required
			continue
		}

		if m.MuteTimeIntervals[i] != nil {
			if err := m.MuteTimeIntervals[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mute_time_intervals" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mute_time_intervals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GettableAPIAlertingConfig) validateReceivers(formats strfmt.Registry) error {
	if swag.IsZero(m.Receivers) { // not required
		return nil
	}

	for i := 0; i < len(m.Receivers); i++ {
		if swag.IsZero(m.Receivers[i]) { // not required
			continue
		}

		if m.Receivers[i] != nil {
			if err := m.Receivers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("receivers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("receivers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GettableAPIAlertingConfig) validateRoute(formats strfmt.Registry) error {
	if swag.IsZero(m.Route) { // not required
		return nil
	}

	if m.Route != nil {
		if err := m.Route.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("route")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("route")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this gettable Api alerting config based on the context it is used
func (m *GettableAPIAlertingConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGlobal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInhibitRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMuteTimeProvenances(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMuteTimeIntervals(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReceivers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoute(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GettableAPIAlertingConfig) contextValidateGlobal(ctx context.Context, formats strfmt.Registry) error {

	if m.Global != nil {

		if swag.IsZero(m.Global) { // not required
			return nil
		}

		if err := m.Global.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("global")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("global")
			}
			return err
		}
	}

	return nil
}

func (m *GettableAPIAlertingConfig) contextValidateInhibitRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InhibitRules); i++ {

		if m.InhibitRules[i] != nil {

			if swag.IsZero(m.InhibitRules[i]) { // not required
				return nil
			}

			if err := m.InhibitRules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inhibit_rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inhibit_rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GettableAPIAlertingConfig) contextValidateMuteTimeProvenances(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.MuteTimeProvenances {

		if val, ok := m.MuteTimeProvenances[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *GettableAPIAlertingConfig) contextValidateMuteTimeIntervals(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MuteTimeIntervals); i++ {

		if m.MuteTimeIntervals[i] != nil {

			if swag.IsZero(m.MuteTimeIntervals[i]) { // not required
				return nil
			}

			if err := m.MuteTimeIntervals[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mute_time_intervals" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mute_time_intervals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GettableAPIAlertingConfig) contextValidateReceivers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Receivers); i++ {

		if m.Receivers[i] != nil {

			if swag.IsZero(m.Receivers[i]) { // not required
				return nil
			}

			if err := m.Receivers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("receivers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("receivers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GettableAPIAlertingConfig) contextValidateRoute(ctx context.Context, formats strfmt.Registry) error {

	if m.Route != nil {

		if swag.IsZero(m.Route) { // not required
			return nil
		}

		if err := m.Route.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("route")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("route")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GettableAPIAlertingConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GettableAPIAlertingConfig) UnmarshalBinary(b []byte) error {
	var res GettableAPIAlertingConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
