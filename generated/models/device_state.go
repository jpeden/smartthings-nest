// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeviceState device state
//
// swagger:model DeviceState
type DeviceState struct {

	// device cookie
	DeviceCookie Cookie `json:"deviceCookie,omitempty"`

	// device error
	DeviceError []*DeviceStateDeviceErrorItems0 `json:"deviceError,omitempty"`

	// external device Id
	// Max Length: 256
	ExternalDeviceID string `json:"externalDeviceId,omitempty"`

	// states
	States []*DeviceStateStatesItems0 `json:"states"`
}

// Validate validates this device state
func (m *DeviceState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeviceCookie(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalDeviceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceState) validateDeviceCookie(formats strfmt.Registry) error {

	if swag.IsZero(m.DeviceCookie) { // not required
		return nil
	}

	if err := m.DeviceCookie.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("deviceCookie")
		}
		return err
	}

	return nil
}

func (m *DeviceState) validateDeviceError(formats strfmt.Registry) error {

	if swag.IsZero(m.DeviceError) { // not required
		return nil
	}

	for i := 0; i < len(m.DeviceError); i++ {
		if swag.IsZero(m.DeviceError[i]) { // not required
			continue
		}

		if m.DeviceError[i] != nil {
			if err := m.DeviceError[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deviceError" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeviceState) validateExternalDeviceID(formats strfmt.Registry) error {

	if swag.IsZero(m.ExternalDeviceID) { // not required
		return nil
	}

	if err := validate.MaxLength("externalDeviceId", "body", string(m.ExternalDeviceID), 256); err != nil {
		return err
	}

	return nil
}

func (m *DeviceState) validateStates(formats strfmt.Registry) error {

	if swag.IsZero(m.States) { // not required
		return nil
	}

	for i := 0; i < len(m.States); i++ {
		if swag.IsZero(m.States[i]) { // not required
			continue
		}

		if m.States[i] != nil {
			if err := m.States[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("states" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceState) UnmarshalBinary(b []byte) error {
	var res DeviceState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DeviceStateDeviceErrorItems0 device state device error items0
//
// swagger:model DeviceStateDeviceErrorItems0
type DeviceStateDeviceErrorItems0 struct {

	// detail
	// Max Length: 256
	Detail string `json:"detail,omitempty"`

	// error enum
	// Required: true
	// Max Length: 50
	ErrorEnum *string `json:"errorEnum"`
}

// Validate validates this device state device error items0
func (m *DeviceStateDeviceErrorItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDetail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrorEnum(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceStateDeviceErrorItems0) validateDetail(formats strfmt.Registry) error {

	if swag.IsZero(m.Detail) { // not required
		return nil
	}

	if err := validate.MaxLength("detail", "body", string(m.Detail), 256); err != nil {
		return err
	}

	return nil
}

func (m *DeviceStateDeviceErrorItems0) validateErrorEnum(formats strfmt.Registry) error {

	if err := validate.Required("errorEnum", "body", m.ErrorEnum); err != nil {
		return err
	}

	if err := validate.MaxLength("errorEnum", "body", string(*m.ErrorEnum), 50); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceStateDeviceErrorItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceStateDeviceErrorItems0) UnmarshalBinary(b []byte) error {
	var res DeviceStateDeviceErrorItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DeviceStateStatesItems0 device state states items0
//
// swagger:model DeviceStateStatesItems0
type DeviceStateStatesItems0 struct {

	// attribute
	// Max Length: 30
	Attribute string `json:"attribute,omitempty"`

	// capability
	// Max Length: 30
	// Pattern: ^st\.
	Capability string `json:"capability,omitempty"`

	// component
	// Max Length: 30
	Component string `json:"component,omitempty"`

	// timestamp
	// Maximum: 2.147483647e+12
	// Minimum: -2.147483648e+12
	Timestamp *int64 `json:"timestamp,omitempty"`

	// value
	Value interface{} `json:"value,omitempty"`

	// device state states items0 additional properties
	DeviceStateStatesItems0AdditionalProperties map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *DeviceStateStatesItems0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// attribute
		// Max Length: 30
		Attribute string `json:"attribute,omitempty"`

		// capability
		// Max Length: 30
		// Pattern: ^st\.
		Capability string `json:"capability,omitempty"`

		// component
		// Max Length: 30
		Component string `json:"component,omitempty"`

		// timestamp
		// Maximum: 2.147483647e+12
		// Minimum: -2.147483648e+12
		Timestamp *int64 `json:"timestamp,omitempty"`

		// value
		Value interface{} `json:"value,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv DeviceStateStatesItems0

	rcv.Attribute = stage1.Attribute
	rcv.Capability = stage1.Capability
	rcv.Component = stage1.Component
	rcv.Timestamp = stage1.Timestamp
	rcv.Value = stage1.Value
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "attribute")
	delete(stage2, "capability")
	delete(stage2, "component")
	delete(stage2, "timestamp")
	delete(stage2, "value")
	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]interface{})
		for k, v := range stage2 {
			var toadd interface{}
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		m.DeviceStateStatesItems0AdditionalProperties = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m DeviceStateStatesItems0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// attribute
		// Max Length: 30
		Attribute string `json:"attribute,omitempty"`

		// capability
		// Max Length: 30
		// Pattern: ^st\.
		Capability string `json:"capability,omitempty"`

		// component
		// Max Length: 30
		Component string `json:"component,omitempty"`

		// timestamp
		// Maximum: 2.147483647e+12
		// Minimum: -2.147483648e+12
		Timestamp *int64 `json:"timestamp,omitempty"`

		// value
		Value interface{} `json:"value,omitempty"`
	}

	stage1.Attribute = m.Attribute
	stage1.Capability = m.Capability
	stage1.Component = m.Component
	stage1.Timestamp = m.Timestamp
	stage1.Value = m.Value

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.DeviceStateStatesItems0AdditionalProperties) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.DeviceStateStatesItems0AdditionalProperties)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 {
		return additional, nil
	}

	// concatenate the 2 objects
	props[len(props)-1] = ','
	return append(props, additional[1:]...), nil
}

// Validate validates this device state states items0
func (m *DeviceStateStatesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttribute(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCapability(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceStateStatesItems0) validateAttribute(formats strfmt.Registry) error {

	if swag.IsZero(m.Attribute) { // not required
		return nil
	}

	if err := validate.MaxLength("attribute", "body", string(m.Attribute), 30); err != nil {
		return err
	}

	return nil
}

func (m *DeviceStateStatesItems0) validateCapability(formats strfmt.Registry) error {

	if swag.IsZero(m.Capability) { // not required
		return nil
	}

	if err := validate.MaxLength("capability", "body", string(m.Capability), 30); err != nil {
		return err
	}

	if err := validate.Pattern("capability", "body", string(m.Capability), `^st\.`); err != nil {
		return err
	}

	return nil
}

func (m *DeviceStateStatesItems0) validateComponent(formats strfmt.Registry) error {

	if swag.IsZero(m.Component) { // not required
		return nil
	}

	if err := validate.MaxLength("component", "body", string(m.Component), 30); err != nil {
		return err
	}

	return nil
}

func (m *DeviceStateStatesItems0) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.MinimumInt("timestamp", "body", int64(*m.Timestamp), -2.147483648e+12, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("timestamp", "body", int64(*m.Timestamp), 2.147483647e+12, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceStateStatesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceStateStatesItems0) UnmarshalBinary(b []byte) error {
	var res DeviceStateStatesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}