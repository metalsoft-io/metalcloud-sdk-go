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

// checks if the NetworkDevicePorts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkDevicePorts{}

// NetworkDevicePorts struct for NetworkDevicePorts
type NetworkDevicePorts struct {
	SwitchId string `json:"switch_id"`
	Ports []NetworkDevicePort `json:"ports"`
	AdditionalProperties map[string]interface{}
}

type _NetworkDevicePorts NetworkDevicePorts

// NewNetworkDevicePorts instantiates a new NetworkDevicePorts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkDevicePorts(switchId string, ports []NetworkDevicePort) *NetworkDevicePorts {
	this := NetworkDevicePorts{}
	this.SwitchId = switchId
	this.Ports = ports
	return &this
}

// NewNetworkDevicePortsWithDefaults instantiates a new NetworkDevicePorts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkDevicePortsWithDefaults() *NetworkDevicePorts {
	this := NetworkDevicePorts{}
	return &this
}

// GetSwitchId returns the SwitchId field value
func (o *NetworkDevicePorts) GetSwitchId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SwitchId
}

// GetSwitchIdOk returns a tuple with the SwitchId field value
// and a boolean to check if the value has been set.
func (o *NetworkDevicePorts) GetSwitchIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SwitchId, true
}

// SetSwitchId sets field value
func (o *NetworkDevicePorts) SetSwitchId(v string) {
	o.SwitchId = v
}

// GetPorts returns the Ports field value
func (o *NetworkDevicePorts) GetPorts() []NetworkDevicePort {
	if o == nil {
		var ret []NetworkDevicePort
		return ret
	}

	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value
// and a boolean to check if the value has been set.
func (o *NetworkDevicePorts) GetPortsOk() ([]NetworkDevicePort, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ports, true
}

// SetPorts sets field value
func (o *NetworkDevicePorts) SetPorts(v []NetworkDevicePort) {
	o.Ports = v
}

func (o NetworkDevicePorts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkDevicePorts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["switch_id"] = o.SwitchId
	toSerialize["ports"] = o.Ports

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkDevicePorts) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"switch_id",
		"ports",
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

	varNetworkDevicePorts := _NetworkDevicePorts{}

	err = json.Unmarshal(data, &varNetworkDevicePorts)

	if err != nil {
		return err
	}

	*o = NetworkDevicePorts(varNetworkDevicePorts)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "switch_id")
		delete(additionalProperties, "ports")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkDevicePorts struct {
	value *NetworkDevicePorts
	isSet bool
}

func (v NullableNetworkDevicePorts) Get() *NetworkDevicePorts {
	return v.value
}

func (v *NullableNetworkDevicePorts) Set(val *NetworkDevicePorts) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkDevicePorts) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkDevicePorts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkDevicePorts(val *NetworkDevicePorts) *NullableNetworkDevicePorts {
	return &NullableNetworkDevicePorts{value: val, isSet: true}
}

func (v NullableNetworkDevicePorts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkDevicePorts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


