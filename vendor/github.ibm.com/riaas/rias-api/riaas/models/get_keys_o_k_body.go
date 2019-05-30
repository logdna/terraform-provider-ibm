// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetKeysOKBody KeyCollection
// swagger:model getKeysOKBody
type GetKeysOKBody struct {
	Pagination

	// Collection of keys
	Keys []*Key `json:"keys"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GetKeysOKBody) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Pagination
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Pagination = aO0

	// now for regular properties
	var propsGetKeysOKBody struct {
		Keys []*Key `json:"keys"`
	}
	if err := swag.ReadJSON(raw, &propsGetKeysOKBody); err != nil {
		return err
	}
	m.Keys = propsGetKeysOKBody.Keys

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GetKeysOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.Pagination)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsGetKeysOKBody struct {
		Keys []*Key `json:"keys"`
	}
	propsGetKeysOKBody.Keys = m.Keys

	jsonDataPropsGetKeysOKBody, errGetKeysOKBody := swag.WriteJSON(propsGetKeysOKBody)
	if errGetKeysOKBody != nil {
		return nil, errGetKeysOKBody
	}
	_parts = append(_parts, jsonDataPropsGetKeysOKBody)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get keys o k body
func (m *GetKeysOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Pagination
	if err := m.Pagination.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeys(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetKeysOKBody) validateKeys(formats strfmt.Registry) error {

	if swag.IsZero(m.Keys) { // not required
		return nil
	}

	for i := 0; i < len(m.Keys); i++ {
		if swag.IsZero(m.Keys[i]) { // not required
			continue
		}

		if m.Keys[i] != nil {
			if err := m.Keys[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("keys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetKeysOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetKeysOKBody) UnmarshalBinary(b []byte) error {
	var res GetKeysOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}