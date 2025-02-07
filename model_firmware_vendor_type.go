/*
MetalSoft REST API

MetalSoft REST API documentation

API version: 2.0
Contact: support@metalsoft.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
	"fmt"
)

// FirmwareVendorType the model 'FirmwareVendorType'
type FirmwareVendorType string

// List of FirmwareVendorType
const (
	DELL FirmwareVendorType = "dell"
	LENOVO FirmwareVendorType = "lenovo"
	HP FirmwareVendorType = "hp"
)

// All allowed values of FirmwareVendorType enum
var AllowedFirmwareVendorTypeEnumValues = []FirmwareVendorType{
	"dell",
	"lenovo",
	"hp",
}

func (v *FirmwareVendorType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FirmwareVendorType(value)
	for _, existing := range AllowedFirmwareVendorTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FirmwareVendorType", value)
}

// NewFirmwareVendorTypeFromValue returns a pointer to a valid FirmwareVendorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFirmwareVendorTypeFromValue(v string) (*FirmwareVendorType, error) {
	ev := FirmwareVendorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FirmwareVendorType: valid values are %v", v, AllowedFirmwareVendorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FirmwareVendorType) IsValid() bool {
	for _, existing := range AllowedFirmwareVendorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FirmwareVendorType value
func (v FirmwareVendorType) Ptr() *FirmwareVendorType {
	return &v
}

type NullableFirmwareVendorType struct {
	value *FirmwareVendorType
	isSet bool
}

func (v NullableFirmwareVendorType) Get() *FirmwareVendorType {
	return v.value
}

func (v *NullableFirmwareVendorType) Set(val *FirmwareVendorType) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareVendorType) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareVendorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareVendorType(val *FirmwareVendorType) *NullableFirmwareVendorType {
	return &NullableFirmwareVendorType{value: val, isSet: true}
}

func (v NullableFirmwareVendorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareVendorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

