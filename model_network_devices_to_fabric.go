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

// checks if the NetworkDevicesToFabric type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkDevicesToFabric{}

// NetworkDevicesToFabric struct for NetworkDevicesToFabric
type NetworkDevicesToFabric struct {
	// The network devices IDs to add to the fabric
	NetworkDeviceIds []float32 `json:"networkDeviceIds"`
	AdditionalProperties map[string]interface{}
}

type _NetworkDevicesToFabric NetworkDevicesToFabric

// NewNetworkDevicesToFabric instantiates a new NetworkDevicesToFabric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkDevicesToFabric(networkDeviceIds []float32) *NetworkDevicesToFabric {
	this := NetworkDevicesToFabric{}
	this.NetworkDeviceIds = networkDeviceIds
	return &this
}

// NewNetworkDevicesToFabricWithDefaults instantiates a new NetworkDevicesToFabric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkDevicesToFabricWithDefaults() *NetworkDevicesToFabric {
	this := NetworkDevicesToFabric{}
	return &this
}

// GetNetworkDeviceIds returns the NetworkDeviceIds field value
func (o *NetworkDevicesToFabric) GetNetworkDeviceIds() []float32 {
	if o == nil {
		var ret []float32
		return ret
	}

	return o.NetworkDeviceIds
}

// GetNetworkDeviceIdsOk returns a tuple with the NetworkDeviceIds field value
// and a boolean to check if the value has been set.
func (o *NetworkDevicesToFabric) GetNetworkDeviceIdsOk() ([]float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetworkDeviceIds, true
}

// SetNetworkDeviceIds sets field value
func (o *NetworkDevicesToFabric) SetNetworkDeviceIds(v []float32) {
	o.NetworkDeviceIds = v
}

func (o NetworkDevicesToFabric) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkDevicesToFabric) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["networkDeviceIds"] = o.NetworkDeviceIds

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkDevicesToFabric) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"networkDeviceIds",
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

	varNetworkDevicesToFabric := _NetworkDevicesToFabric{}

	err = json.Unmarshal(data, &varNetworkDevicesToFabric)

	if err != nil {
		return err
	}

	*o = NetworkDevicesToFabric(varNetworkDevicesToFabric)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "networkDeviceIds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkDevicesToFabric struct {
	value *NetworkDevicesToFabric
	isSet bool
}

func (v NullableNetworkDevicesToFabric) Get() *NetworkDevicesToFabric {
	return v.value
}

func (v *NullableNetworkDevicesToFabric) Set(val *NetworkDevicesToFabric) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkDevicesToFabric) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkDevicesToFabric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkDevicesToFabric(val *NetworkDevicesToFabric) *NullableNetworkDevicesToFabric {
	return &NullableNetworkDevicesToFabric{value: val, isSet: true}
}

func (v NullableNetworkDevicesToFabric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkDevicesToFabric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


