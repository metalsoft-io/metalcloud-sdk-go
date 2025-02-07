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

// FirmwareBinaryUpdateSeverity the model 'FirmwareBinaryUpdateSeverity'
type FirmwareBinaryUpdateSeverity string

// List of FirmwareBinaryUpdateSeverity
const (
	CRITICAL FirmwareBinaryUpdateSeverity = "critical"
	RECOMMENDED FirmwareBinaryUpdateSeverity = "recommended"
	OPTIONAL FirmwareBinaryUpdateSeverity = "optional"
	UNKNOWN FirmwareBinaryUpdateSeverity = "unknown"
)

// All allowed values of FirmwareBinaryUpdateSeverity enum
var AllowedFirmwareBinaryUpdateSeverityEnumValues = []FirmwareBinaryUpdateSeverity{
	"critical",
	"recommended",
	"optional",
	"unknown",
}

func (v *FirmwareBinaryUpdateSeverity) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FirmwareBinaryUpdateSeverity(value)
	for _, existing := range AllowedFirmwareBinaryUpdateSeverityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FirmwareBinaryUpdateSeverity", value)
}

// NewFirmwareBinaryUpdateSeverityFromValue returns a pointer to a valid FirmwareBinaryUpdateSeverity
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFirmwareBinaryUpdateSeverityFromValue(v string) (*FirmwareBinaryUpdateSeverity, error) {
	ev := FirmwareBinaryUpdateSeverity(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FirmwareBinaryUpdateSeverity: valid values are %v", v, AllowedFirmwareBinaryUpdateSeverityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FirmwareBinaryUpdateSeverity) IsValid() bool {
	for _, existing := range AllowedFirmwareBinaryUpdateSeverityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FirmwareBinaryUpdateSeverity value
func (v FirmwareBinaryUpdateSeverity) Ptr() *FirmwareBinaryUpdateSeverity {
	return &v
}

type NullableFirmwareBinaryUpdateSeverity struct {
	value *FirmwareBinaryUpdateSeverity
	isSet bool
}

func (v NullableFirmwareBinaryUpdateSeverity) Get() *FirmwareBinaryUpdateSeverity {
	return v.value
}

func (v *NullableFirmwareBinaryUpdateSeverity) Set(val *FirmwareBinaryUpdateSeverity) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareBinaryUpdateSeverity) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareBinaryUpdateSeverity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareBinaryUpdateSeverity(val *FirmwareBinaryUpdateSeverity) *NullableFirmwareBinaryUpdateSeverity {
	return &NullableFirmwareBinaryUpdateSeverity{value: val, isSet: true}
}

func (v NullableFirmwareBinaryUpdateSeverity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareBinaryUpdateSeverity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

