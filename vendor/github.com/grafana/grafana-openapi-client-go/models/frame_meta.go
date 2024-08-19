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

// FrameMeta FrameMeta matches:
//
// https://github.com/grafana/grafana/blob/master/packages/grafana-data/src/types/data.ts#L11
// NOTE -- in javascript this can accept any `[key: string]: any;` however
// this interface only exposes the values we want to be exposed
//
// swagger:model FrameMeta
type FrameMeta struct {

	// Channel is the path to a stream in grafana live that has real-time updates for this data.
	Channel string `json:"channel,omitempty"`

	// Custom datasource specific values.
	Custom interface{} `json:"custom,omitempty"`

	// data topic
	DataTopic DataTopic `json:"dataTopic,omitempty"`

	// ExecutedQueryString is the raw query sent to the underlying system. All macros and templating
	// have been applied.  When metadata contains this value, it will be shown in the query inspector.
	ExecutedQueryString string `json:"executedQueryString,omitempty"`

	// Notices provide additional information about the data in the Frame that
	// Grafana can display to the user in the user interface.
	Notices []*Notice `json:"notices"`

	// Path is a browsable path on the datasource.
	Path string `json:"path,omitempty"`

	// PathSeparator defines the separator pattern to decode a hierarchy. The default separator is '/'.
	PathSeparator string `json:"pathSeparator,omitempty"`

	// PreferredVisualizationPluginId sets the panel plugin id to use to render the data when using Explore. If
	// the plugin cannot be found will fall back to PreferredVisualization.
	PreferredVisualisationPluginID string `json:"preferredVisualisationPluginId,omitempty"`

	// preferred visualisation type
	PreferredVisualisationType VisType `json:"preferredVisualisationType,omitempty"`

	// Stats is an array of query result statistics.
	Stats []*QueryStat `json:"stats"`

	// type
	Type FrameType `json:"type,omitempty"`

	// type version
	TypeVersion FrameTypeVersion `json:"typeVersion,omitempty"`

	// Array of field indices which values create a unique id for each row. Ideally this should be globally unique ID
	// but that isn't guarantied. Should help with keeping track and deduplicating rows in visualizations, especially
	// with streaming data with frequent updates.
	// Example: TraceID in Tempo, table name + primary key in SQL
	UniqueRowIDFields []int64 `json:"uniqueRowIdFields"`
}

// Validate validates this frame meta
func (m *FrameMeta) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataTopic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreferredVisualisationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypeVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FrameMeta) validateDataTopic(formats strfmt.Registry) error {
	if swag.IsZero(m.DataTopic) { // not required
		return nil
	}

	if err := m.DataTopic.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("dataTopic")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("dataTopic")
		}
		return err
	}

	return nil
}

func (m *FrameMeta) validateNotices(formats strfmt.Registry) error {
	if swag.IsZero(m.Notices) { // not required
		return nil
	}

	for i := 0; i < len(m.Notices); i++ {
		if swag.IsZero(m.Notices[i]) { // not required
			continue
		}

		if m.Notices[i] != nil {
			if err := m.Notices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("notices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("notices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FrameMeta) validatePreferredVisualisationType(formats strfmt.Registry) error {
	if swag.IsZero(m.PreferredVisualisationType) { // not required
		return nil
	}

	if err := m.PreferredVisualisationType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("preferredVisualisationType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("preferredVisualisationType")
		}
		return err
	}

	return nil
}

func (m *FrameMeta) validateStats(formats strfmt.Registry) error {
	if swag.IsZero(m.Stats) { // not required
		return nil
	}

	for i := 0; i < len(m.Stats); i++ {
		if swag.IsZero(m.Stats[i]) { // not required
			continue
		}

		if m.Stats[i] != nil {
			if err := m.Stats[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stats" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("stats" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FrameMeta) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

func (m *FrameMeta) validateTypeVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.TypeVersion) { // not required
		return nil
	}

	if err := m.TypeVersion.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("typeVersion")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("typeVersion")
		}
		return err
	}

	return nil
}

// ContextValidate validate this frame meta based on the context it is used
func (m *FrameMeta) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDataTopic(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNotices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePreferredVisualisationType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTypeVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FrameMeta) contextValidateDataTopic(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.DataTopic) { // not required
		return nil
	}

	if err := m.DataTopic.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("dataTopic")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("dataTopic")
		}
		return err
	}

	return nil
}

func (m *FrameMeta) contextValidateNotices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Notices); i++ {

		if m.Notices[i] != nil {

			if swag.IsZero(m.Notices[i]) { // not required
				return nil
			}

			if err := m.Notices[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("notices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("notices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FrameMeta) contextValidatePreferredVisualisationType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.PreferredVisualisationType) { // not required
		return nil
	}

	if err := m.PreferredVisualisationType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("preferredVisualisationType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("preferredVisualisationType")
		}
		return err
	}

	return nil
}

func (m *FrameMeta) contextValidateStats(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Stats); i++ {

		if m.Stats[i] != nil {

			if swag.IsZero(m.Stats[i]) { // not required
				return nil
			}

			if err := m.Stats[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stats" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("stats" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FrameMeta) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

func (m *FrameMeta) contextValidateTypeVersion(ctx context.Context, formats strfmt.Registry) error {

	if err := m.TypeVersion.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("typeVersion")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("typeVersion")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FrameMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FrameMeta) UnmarshalBinary(b []byte) error {
	var res FrameMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
