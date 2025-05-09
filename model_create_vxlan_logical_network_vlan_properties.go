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

// checks if the CreateVxlanLogicalNetworkVlanProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateVxlanLogicalNetworkVlanProperties{}

// CreateVxlanLogicalNetworkVlanProperties struct for CreateVxlanLogicalNetworkVlanProperties
type CreateVxlanLogicalNetworkVlanProperties struct {
	VlanAllocationStrategies []CreateVlanAllocationStrategy `json:"vlanAllocationStrategies"`
	AdditionalProperties map[string]interface{}
}

type _CreateVxlanLogicalNetworkVlanProperties CreateVxlanLogicalNetworkVlanProperties

// NewCreateVxlanLogicalNetworkVlanProperties instantiates a new CreateVxlanLogicalNetworkVlanProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVxlanLogicalNetworkVlanProperties(vlanAllocationStrategies []CreateVlanAllocationStrategy) *CreateVxlanLogicalNetworkVlanProperties {
	this := CreateVxlanLogicalNetworkVlanProperties{}
	this.VlanAllocationStrategies = vlanAllocationStrategies
	return &this
}

// NewCreateVxlanLogicalNetworkVlanPropertiesWithDefaults instantiates a new CreateVxlanLogicalNetworkVlanProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVxlanLogicalNetworkVlanPropertiesWithDefaults() *CreateVxlanLogicalNetworkVlanProperties {
	this := CreateVxlanLogicalNetworkVlanProperties{}
	return &this
}

// GetVlanAllocationStrategies returns the VlanAllocationStrategies field value
func (o *CreateVxlanLogicalNetworkVlanProperties) GetVlanAllocationStrategies() []CreateVlanAllocationStrategy {
	if o == nil {
		var ret []CreateVlanAllocationStrategy
		return ret
	}

	return o.VlanAllocationStrategies
}

// GetVlanAllocationStrategiesOk returns a tuple with the VlanAllocationStrategies field value
// and a boolean to check if the value has been set.
func (o *CreateVxlanLogicalNetworkVlanProperties) GetVlanAllocationStrategiesOk() ([]CreateVlanAllocationStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return o.VlanAllocationStrategies, true
}

// SetVlanAllocationStrategies sets field value
func (o *CreateVxlanLogicalNetworkVlanProperties) SetVlanAllocationStrategies(v []CreateVlanAllocationStrategy) {
	o.VlanAllocationStrategies = v
}

func (o CreateVxlanLogicalNetworkVlanProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateVxlanLogicalNetworkVlanProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["vlanAllocationStrategies"] = o.VlanAllocationStrategies

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateVxlanLogicalNetworkVlanProperties) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"vlanAllocationStrategies",
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

	varCreateVxlanLogicalNetworkVlanProperties := _CreateVxlanLogicalNetworkVlanProperties{}

	err = json.Unmarshal(data, &varCreateVxlanLogicalNetworkVlanProperties)

	if err != nil {
		return err
	}

	*o = CreateVxlanLogicalNetworkVlanProperties(varCreateVxlanLogicalNetworkVlanProperties)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "vlanAllocationStrategies")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateVxlanLogicalNetworkVlanProperties struct {
	value *CreateVxlanLogicalNetworkVlanProperties
	isSet bool
}

func (v NullableCreateVxlanLogicalNetworkVlanProperties) Get() *CreateVxlanLogicalNetworkVlanProperties {
	return v.value
}

func (v *NullableCreateVxlanLogicalNetworkVlanProperties) Set(val *CreateVxlanLogicalNetworkVlanProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVxlanLogicalNetworkVlanProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVxlanLogicalNetworkVlanProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVxlanLogicalNetworkVlanProperties(val *CreateVxlanLogicalNetworkVlanProperties) *NullableCreateVxlanLogicalNetworkVlanProperties {
	return &NullableCreateVxlanLogicalNetworkVlanProperties{value: val, isSet: true}
}

func (v NullableCreateVxlanLogicalNetworkVlanProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVxlanLogicalNetworkVlanProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


