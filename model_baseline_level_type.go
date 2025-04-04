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

// BaselineLevelType the model 'BaselineLevelType'
type BaselineLevelType string

// List of BaselineLevelType
const (
	BASELINELEVELTYPE_DATACENTER BaselineLevelType = "datacenter"
	BASELINELEVELTYPE_SERVER_TYPE BaselineLevelType = "serverType"
	BASELINELEVELTYPE_OS_TEMPLATE BaselineLevelType = "osTemplate"
)

// All allowed values of BaselineLevelType enum
var AllowedBaselineLevelTypeEnumValues = []BaselineLevelType{
	"datacenter",
	"serverType",
	"osTemplate",
}

func (v *BaselineLevelType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BaselineLevelType(value)
	for _, existing := range AllowedBaselineLevelTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BaselineLevelType", value)
}

// NewBaselineLevelTypeFromValue returns a pointer to a valid BaselineLevelType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBaselineLevelTypeFromValue(v string) (*BaselineLevelType, error) {
	ev := BaselineLevelType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BaselineLevelType: valid values are %v", v, AllowedBaselineLevelTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BaselineLevelType) IsValid() bool {
	for _, existing := range AllowedBaselineLevelTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BaselineLevelType value
func (v BaselineLevelType) Ptr() *BaselineLevelType {
	return &v
}

type NullableBaselineLevelType struct {
	value *BaselineLevelType
	isSet bool
}

func (v NullableBaselineLevelType) Get() *BaselineLevelType {
	return v.value
}

func (v *NullableBaselineLevelType) Set(val *BaselineLevelType) {
	v.value = val
	v.isSet = true
}

func (v NullableBaselineLevelType) IsSet() bool {
	return v.isSet
}

func (v *NullableBaselineLevelType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaselineLevelType(val *BaselineLevelType) *NullableBaselineLevelType {
	return &NullableBaselineLevelType{value: val, isSet: true}
}

func (v NullableBaselineLevelType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaselineLevelType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

