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

// checks if the NetworkDevicePortStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkDevicePortStatus{}

// NetworkDevicePortStatus struct for NetworkDevicePortStatus
type NetworkDevicePortStatus struct {
	// The ports of the network device that will have their status changed
	Ports []string `json:"ports"`
	// The new status of the ports
	Status bool `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _NetworkDevicePortStatus NetworkDevicePortStatus

// NewNetworkDevicePortStatus instantiates a new NetworkDevicePortStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkDevicePortStatus(ports []string, status bool) *NetworkDevicePortStatus {
	this := NetworkDevicePortStatus{}
	this.Ports = ports
	this.Status = status
	return &this
}

// NewNetworkDevicePortStatusWithDefaults instantiates a new NetworkDevicePortStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkDevicePortStatusWithDefaults() *NetworkDevicePortStatus {
	this := NetworkDevicePortStatus{}
	return &this
}

// GetPorts returns the Ports field value
func (o *NetworkDevicePortStatus) GetPorts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value
// and a boolean to check if the value has been set.
func (o *NetworkDevicePortStatus) GetPortsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ports, true
}

// SetPorts sets field value
func (o *NetworkDevicePortStatus) SetPorts(v []string) {
	o.Ports = v
}

// GetStatus returns the Status field value
func (o *NetworkDevicePortStatus) GetStatus() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *NetworkDevicePortStatus) GetStatusOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *NetworkDevicePortStatus) SetStatus(v bool) {
	o.Status = v
}

func (o NetworkDevicePortStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkDevicePortStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ports"] = o.Ports
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkDevicePortStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ports",
		"status",
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

	varNetworkDevicePortStatus := _NetworkDevicePortStatus{}

	err = json.Unmarshal(data, &varNetworkDevicePortStatus)

	if err != nil {
		return err
	}

	*o = NetworkDevicePortStatus(varNetworkDevicePortStatus)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ports")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkDevicePortStatus struct {
	value *NetworkDevicePortStatus
	isSet bool
}

func (v NullableNetworkDevicePortStatus) Get() *NetworkDevicePortStatus {
	return v.value
}

func (v *NullableNetworkDevicePortStatus) Set(val *NetworkDevicePortStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkDevicePortStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkDevicePortStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkDevicePortStatus(val *NetworkDevicePortStatus) *NullableNetworkDevicePortStatus {
	return &NullableNetworkDevicePortStatus{value: val, isSet: true}
}

func (v NullableNetworkDevicePortStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkDevicePortStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


