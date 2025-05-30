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

// checks if the VxlanLogicalNetworkConfigVxlanProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VxlanLogicalNetworkConfigVxlanProperties{}

// VxlanLogicalNetworkConfigVxlanProperties struct for VxlanLogicalNetworkConfigVxlanProperties
type VxlanLogicalNetworkConfigVxlanProperties struct {
	VniAllocationStrategies []VniAllocationStrategy `json:"vniAllocationStrategies"`
	AdditionalProperties map[string]interface{}
}

type _VxlanLogicalNetworkConfigVxlanProperties VxlanLogicalNetworkConfigVxlanProperties

// NewVxlanLogicalNetworkConfigVxlanProperties instantiates a new VxlanLogicalNetworkConfigVxlanProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVxlanLogicalNetworkConfigVxlanProperties(vniAllocationStrategies []VniAllocationStrategy) *VxlanLogicalNetworkConfigVxlanProperties {
	this := VxlanLogicalNetworkConfigVxlanProperties{}
	this.VniAllocationStrategies = vniAllocationStrategies
	return &this
}

// NewVxlanLogicalNetworkConfigVxlanPropertiesWithDefaults instantiates a new VxlanLogicalNetworkConfigVxlanProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVxlanLogicalNetworkConfigVxlanPropertiesWithDefaults() *VxlanLogicalNetworkConfigVxlanProperties {
	this := VxlanLogicalNetworkConfigVxlanProperties{}
	return &this
}

// GetVniAllocationStrategies returns the VniAllocationStrategies field value
func (o *VxlanLogicalNetworkConfigVxlanProperties) GetVniAllocationStrategies() []VniAllocationStrategy {
	if o == nil {
		var ret []VniAllocationStrategy
		return ret
	}

	return o.VniAllocationStrategies
}

// GetVniAllocationStrategiesOk returns a tuple with the VniAllocationStrategies field value
// and a boolean to check if the value has been set.
func (o *VxlanLogicalNetworkConfigVxlanProperties) GetVniAllocationStrategiesOk() ([]VniAllocationStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return o.VniAllocationStrategies, true
}

// SetVniAllocationStrategies sets field value
func (o *VxlanLogicalNetworkConfigVxlanProperties) SetVniAllocationStrategies(v []VniAllocationStrategy) {
	o.VniAllocationStrategies = v
}

func (o VxlanLogicalNetworkConfigVxlanProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VxlanLogicalNetworkConfigVxlanProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["vniAllocationStrategies"] = o.VniAllocationStrategies

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VxlanLogicalNetworkConfigVxlanProperties) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"vniAllocationStrategies",
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

	varVxlanLogicalNetworkConfigVxlanProperties := _VxlanLogicalNetworkConfigVxlanProperties{}

	err = json.Unmarshal(data, &varVxlanLogicalNetworkConfigVxlanProperties)

	if err != nil {
		return err
	}

	*o = VxlanLogicalNetworkConfigVxlanProperties(varVxlanLogicalNetworkConfigVxlanProperties)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "vniAllocationStrategies")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVxlanLogicalNetworkConfigVxlanProperties struct {
	value *VxlanLogicalNetworkConfigVxlanProperties
	isSet bool
}

func (v NullableVxlanLogicalNetworkConfigVxlanProperties) Get() *VxlanLogicalNetworkConfigVxlanProperties {
	return v.value
}

func (v *NullableVxlanLogicalNetworkConfigVxlanProperties) Set(val *VxlanLogicalNetworkConfigVxlanProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableVxlanLogicalNetworkConfigVxlanProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableVxlanLogicalNetworkConfigVxlanProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVxlanLogicalNetworkConfigVxlanProperties(val *VxlanLogicalNetworkConfigVxlanProperties) *NullableVxlanLogicalNetworkConfigVxlanProperties {
	return &NullableVxlanLogicalNetworkConfigVxlanProperties{value: val, isSet: true}
}

func (v NullableVxlanLogicalNetworkConfigVxlanProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVxlanLogicalNetworkConfigVxlanProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


