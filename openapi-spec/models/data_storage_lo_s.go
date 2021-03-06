// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DataStorageLoS DataStorageLoS can be used to describe a service option covering storage provisioning and availability.
// swagger:model DataStorageLoS
type DataStorageLoS struct {

	// is space efficient
	IsSpaceEfficient bool `json:"isSpaceEfficient,omitempty"`

	// provisioning policy
	// Enum: [Thin Thick]
	ProvisioningPolicy string `json:"provisioningPolicy,omitempty"`

	// recovery time objective
	RecoveryTimeObjective int64 `json:"recoveryTimeObjective,omitempty"`
}

// Validate validates this data storage lo s
func (m *DataStorageLoS) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProvisioningPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var dataStorageLoSTypeProvisioningPolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Thin","Thick"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dataStorageLoSTypeProvisioningPolicyPropEnum = append(dataStorageLoSTypeProvisioningPolicyPropEnum, v)
	}
}

const (

	// DataStorageLoSProvisioningPolicyThin captures enum value "Thin"
	DataStorageLoSProvisioningPolicyThin string = "Thin"

	// DataStorageLoSProvisioningPolicyThick captures enum value "Thick"
	DataStorageLoSProvisioningPolicyThick string = "Thick"
)

// prop value enum
func (m *DataStorageLoS) validateProvisioningPolicyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, dataStorageLoSTypeProvisioningPolicyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DataStorageLoS) validateProvisioningPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.ProvisioningPolicy) { // not required
		return nil
	}

	// value enum
	if err := m.validateProvisioningPolicyEnum("provisioningPolicy", "body", m.ProvisioningPolicy); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataStorageLoS) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataStorageLoS) UnmarshalBinary(b []byte) error {
	var res DataStorageLoS
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
