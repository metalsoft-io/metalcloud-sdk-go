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

// checks if the CreateVlanLogicalNetworkVlanProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateVlanLogicalNetworkVlanProperties{}

// CreateVlanLogicalNetworkVlanProperties struct for CreateVlanLogicalNetworkVlanProperties
type CreateVlanLogicalNetworkVlanProperties struct {
	VlanAllocationStrategies []CreateVlanAllocationStrategy `json:"vlanAllocationStrategies"`
	AdditionalProperties map[string]interface{}
}

type _CreateVlanLogicalNetworkVlanProperties CreateVlanLogicalNetworkVlanProperties

// NewCreateVlanLogicalNetworkVlanProperties instantiates a new CreateVlanLogicalNetworkVlanProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVlanLogicalNetworkVlanProperties(vlanAllocationStrategies []CreateVlanAllocationStrategy) *CreateVlanLogicalNetworkVlanProperties {
	this := CreateVlanLogicalNetworkVlanProperties{}
	this.VlanAllocationStrategies = vlanAllocationStrategies
	return &this
}

// NewCreateVlanLogicalNetworkVlanPropertiesWithDefaults instantiates a new CreateVlanLogicalNetworkVlanProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVlanLogicalNetworkVlanPropertiesWithDefaults() *CreateVlanLogicalNetworkVlanProperties {
	this := CreateVlanLogicalNetworkVlanProperties{}
	return &this
}

// GetVlanAllocationStrategies returns the VlanAllocationStrategies field value
func (o *CreateVlanLogicalNetworkVlanProperties) GetVlanAllocationStrategies() []CreateVlanAllocationStrategy {
	if o == nil {
		var ret []CreateVlanAllocationStrategy
		return ret
	}

	return o.VlanAllocationStrategies
}

// GetVlanAllocationStrategiesOk returns a tuple with the VlanAllocationStrategies field value
// and a boolean to check if the value has been set.
func (o *CreateVlanLogicalNetworkVlanProperties) GetVlanAllocationStrategiesOk() ([]CreateVlanAllocationStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return o.VlanAllocationStrategies, true
}

// SetVlanAllocationStrategies sets field value
func (o *CreateVlanLogicalNetworkVlanProperties) SetVlanAllocationStrategies(v []CreateVlanAllocationStrategy) {
	o.VlanAllocationStrategies = v
}

func (o CreateVlanLogicalNetworkVlanProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateVlanLogicalNetworkVlanProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["vlanAllocationStrategies"] = o.VlanAllocationStrategies

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateVlanLogicalNetworkVlanProperties) UnmarshalJSON(data []byte) (err error) {
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

	varCreateVlanLogicalNetworkVlanProperties := _CreateVlanLogicalNetworkVlanProperties{}

	err = json.Unmarshal(data, &varCreateVlanLogicalNetworkVlanProperties)

	if err != nil {
		return err
	}

	*o = CreateVlanLogicalNetworkVlanProperties(varCreateVlanLogicalNetworkVlanProperties)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "vlanAllocationStrategies")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateVlanLogicalNetworkVlanProperties struct {
	value *CreateVlanLogicalNetworkVlanProperties
	isSet bool
}

func (v NullableCreateVlanLogicalNetworkVlanProperties) Get() *CreateVlanLogicalNetworkVlanProperties {
	return v.value
}

func (v *NullableCreateVlanLogicalNetworkVlanProperties) Set(val *CreateVlanLogicalNetworkVlanProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVlanLogicalNetworkVlanProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVlanLogicalNetworkVlanProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVlanLogicalNetworkVlanProperties(val *CreateVlanLogicalNetworkVlanProperties) *NullableCreateVlanLogicalNetworkVlanProperties {
	return &NullableCreateVlanLogicalNetworkVlanProperties{value: val, isSet: true}
}

func (v NullableCreateVlanLogicalNetworkVlanProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVlanLogicalNetworkVlanProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


