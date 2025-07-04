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

// checks if the LogicalNetworkVxlanProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogicalNetworkVxlanProperties{}

// LogicalNetworkVxlanProperties struct for LogicalNetworkVxlanProperties
type LogicalNetworkVxlanProperties struct {
	Vnis []VniAllocation `json:"vnis"`
	VniAllocationStrategies []VniAllocationStrategy `json:"vniAllocationStrategies"`
	AdditionalProperties map[string]interface{}
}

type _LogicalNetworkVxlanProperties LogicalNetworkVxlanProperties

// NewLogicalNetworkVxlanProperties instantiates a new LogicalNetworkVxlanProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogicalNetworkVxlanProperties(vnis []VniAllocation, vniAllocationStrategies []VniAllocationStrategy) *LogicalNetworkVxlanProperties {
	this := LogicalNetworkVxlanProperties{}
	this.Vnis = vnis
	this.VniAllocationStrategies = vniAllocationStrategies
	return &this
}

// NewLogicalNetworkVxlanPropertiesWithDefaults instantiates a new LogicalNetworkVxlanProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogicalNetworkVxlanPropertiesWithDefaults() *LogicalNetworkVxlanProperties {
	this := LogicalNetworkVxlanProperties{}
	return &this
}

// GetVnis returns the Vnis field value
func (o *LogicalNetworkVxlanProperties) GetVnis() []VniAllocation {
	if o == nil {
		var ret []VniAllocation
		return ret
	}

	return o.Vnis
}

// GetVnisOk returns a tuple with the Vnis field value
// and a boolean to check if the value has been set.
func (o *LogicalNetworkVxlanProperties) GetVnisOk() ([]VniAllocation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vnis, true
}

// SetVnis sets field value
func (o *LogicalNetworkVxlanProperties) SetVnis(v []VniAllocation) {
	o.Vnis = v
}

// GetVniAllocationStrategies returns the VniAllocationStrategies field value
func (o *LogicalNetworkVxlanProperties) GetVniAllocationStrategies() []VniAllocationStrategy {
	if o == nil {
		var ret []VniAllocationStrategy
		return ret
	}

	return o.VniAllocationStrategies
}

// GetVniAllocationStrategiesOk returns a tuple with the VniAllocationStrategies field value
// and a boolean to check if the value has been set.
func (o *LogicalNetworkVxlanProperties) GetVniAllocationStrategiesOk() ([]VniAllocationStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return o.VniAllocationStrategies, true
}

// SetVniAllocationStrategies sets field value
func (o *LogicalNetworkVxlanProperties) SetVniAllocationStrategies(v []VniAllocationStrategy) {
	o.VniAllocationStrategies = v
}

func (o LogicalNetworkVxlanProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogicalNetworkVxlanProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["vnis"] = o.Vnis
	toSerialize["vniAllocationStrategies"] = o.VniAllocationStrategies

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LogicalNetworkVxlanProperties) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"vnis",
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

	varLogicalNetworkVxlanProperties := _LogicalNetworkVxlanProperties{}

	err = json.Unmarshal(data, &varLogicalNetworkVxlanProperties)

	if err != nil {
		return err
	}

	*o = LogicalNetworkVxlanProperties(varLogicalNetworkVxlanProperties)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "vnis")
		delete(additionalProperties, "vniAllocationStrategies")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLogicalNetworkVxlanProperties struct {
	value *LogicalNetworkVxlanProperties
	isSet bool
}

func (v NullableLogicalNetworkVxlanProperties) Get() *LogicalNetworkVxlanProperties {
	return v.value
}

func (v *NullableLogicalNetworkVxlanProperties) Set(val *LogicalNetworkVxlanProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableLogicalNetworkVxlanProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableLogicalNetworkVxlanProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogicalNetworkVxlanProperties(val *LogicalNetworkVxlanProperties) *NullableLogicalNetworkVxlanProperties {
	return &NullableLogicalNetworkVxlanProperties{value: val, isSet: true}
}

func (v NullableLogicalNetworkVxlanProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogicalNetworkVxlanProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


