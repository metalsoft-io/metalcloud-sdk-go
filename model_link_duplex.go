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

// LinkDuplex the model 'LinkDuplex'
type LinkDuplex string

// List of LinkDuplex
const (
	LINKDUPLEX_HALF LinkDuplex = "half"
	LINKDUPLEX_FULL LinkDuplex = "full"
)

// All allowed values of LinkDuplex enum
var AllowedLinkDuplexEnumValues = []LinkDuplex{
	"half",
	"full",
}

func (v *LinkDuplex) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LinkDuplex(value)
	for _, existing := range AllowedLinkDuplexEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LinkDuplex", value)
}

// NewLinkDuplexFromValue returns a pointer to a valid LinkDuplex
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLinkDuplexFromValue(v string) (*LinkDuplex, error) {
	ev := LinkDuplex(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LinkDuplex: valid values are %v", v, AllowedLinkDuplexEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LinkDuplex) IsValid() bool {
	for _, existing := range AllowedLinkDuplexEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LinkDuplex value
func (v LinkDuplex) Ptr() *LinkDuplex {
	return &v
}

type NullableLinkDuplex struct {
	value *LinkDuplex
	isSet bool
}

func (v NullableLinkDuplex) Get() *LinkDuplex {
	return v.value
}

func (v *NullableLinkDuplex) Set(val *LinkDuplex) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkDuplex) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkDuplex) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkDuplex(val *LinkDuplex) *NullableLinkDuplex {
	return &NullableLinkDuplex{value: val, isSet: true}
}

func (v NullableLinkDuplex) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkDuplex) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

