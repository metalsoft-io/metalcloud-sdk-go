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

// AllocationStrategyKind the model 'AllocationStrategyKind'
type AllocationStrategyKind string

// List of AllocationStrategyKind
const (
	ALLOCATIONSTRATEGYKIND_MANUAL AllocationStrategyKind = "manual"
	ALLOCATIONSTRATEGYKIND_AUTO AllocationStrategyKind = "auto"
)

// All allowed values of AllocationStrategyKind enum
var AllowedAllocationStrategyKindEnumValues = []AllocationStrategyKind{
	"manual",
	"auto",
}

func (v *AllocationStrategyKind) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AllocationStrategyKind(value)
	for _, existing := range AllowedAllocationStrategyKindEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AllocationStrategyKind", value)
}

// NewAllocationStrategyKindFromValue returns a pointer to a valid AllocationStrategyKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAllocationStrategyKindFromValue(v string) (*AllocationStrategyKind, error) {
	ev := AllocationStrategyKind(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AllocationStrategyKind: valid values are %v", v, AllowedAllocationStrategyKindEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AllocationStrategyKind) IsValid() bool {
	for _, existing := range AllowedAllocationStrategyKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AllocationStrategyKind value
func (v AllocationStrategyKind) Ptr() *AllocationStrategyKind {
	return &v
}

type NullableAllocationStrategyKind struct {
	value *AllocationStrategyKind
	isSet bool
}

func (v NullableAllocationStrategyKind) Get() *AllocationStrategyKind {
	return v.value
}

func (v *NullableAllocationStrategyKind) Set(val *AllocationStrategyKind) {
	v.value = val
	v.isSet = true
}

func (v NullableAllocationStrategyKind) IsSet() bool {
	return v.isSet
}

func (v *NullableAllocationStrategyKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllocationStrategyKind(val *AllocationStrategyKind) *NullableAllocationStrategyKind {
	return &NullableAllocationStrategyKind{value: val, isSet: true}
}

func (v NullableAllocationStrategyKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllocationStrategyKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

