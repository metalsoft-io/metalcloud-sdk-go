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

// checks if the RedundancyImplementation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RedundancyImplementation{}

// RedundancyImplementation struct for RedundancyImplementation
type RedundancyImplementation struct {
	// The type of redundancy implementation
	ImplementationType NetworkEndpointGroupRedundancyImplementationType `json:"implementationType"`
	AdditionalProperties map[string]interface{}
}

type _RedundancyImplementation RedundancyImplementation

// NewRedundancyImplementation instantiates a new RedundancyImplementation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedundancyImplementation(implementationType NetworkEndpointGroupRedundancyImplementationType) *RedundancyImplementation {
	this := RedundancyImplementation{}
	this.ImplementationType = implementationType
	return &this
}

// NewRedundancyImplementationWithDefaults instantiates a new RedundancyImplementation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedundancyImplementationWithDefaults() *RedundancyImplementation {
	this := RedundancyImplementation{}
	return &this
}

// GetImplementationType returns the ImplementationType field value
func (o *RedundancyImplementation) GetImplementationType() NetworkEndpointGroupRedundancyImplementationType {
	if o == nil {
		var ret NetworkEndpointGroupRedundancyImplementationType
		return ret
	}

	return o.ImplementationType
}

// GetImplementationTypeOk returns a tuple with the ImplementationType field value
// and a boolean to check if the value has been set.
func (o *RedundancyImplementation) GetImplementationTypeOk() (*NetworkEndpointGroupRedundancyImplementationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImplementationType, true
}

// SetImplementationType sets field value
func (o *RedundancyImplementation) SetImplementationType(v NetworkEndpointGroupRedundancyImplementationType) {
	o.ImplementationType = v
}

func (o RedundancyImplementation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RedundancyImplementation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["implementationType"] = o.ImplementationType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RedundancyImplementation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"implementationType",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRedundancyImplementation := _RedundancyImplementation{}

	err = json.Unmarshal(data, &varRedundancyImplementation)

	if err != nil {
		return err
	}

	*o = RedundancyImplementation(varRedundancyImplementation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "implementationType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRedundancyImplementation struct {
	value *RedundancyImplementation
	isSet bool
}

func (v NullableRedundancyImplementation) Get() *RedundancyImplementation {
	return v.value
}

func (v *NullableRedundancyImplementation) Set(val *RedundancyImplementation) {
	v.value = val
	v.isSet = true
}

func (v NullableRedundancyImplementation) IsSet() bool {
	return v.isSet
}

func (v *NullableRedundancyImplementation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedundancyImplementation(val *RedundancyImplementation) *NullableRedundancyImplementation {
	return &NullableRedundancyImplementation{value: val, isSet: true}
}

func (v NullableRedundancyImplementation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedundancyImplementation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


