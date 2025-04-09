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

// checks if the NetworkDevicePortCounters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkDevicePortCounters{}

// NetworkDevicePortCounters struct for NetworkDevicePortCounters
type NetworkDevicePortCounters struct {
	OctetsIn float32 `json:"octets_in"`
	OctetsOut float32 `json:"octets_out"`
	PacketsIn float32 `json:"packets_in"`
	PacketsOut float32 `json:"packets_out"`
	ErrorsIn float32 `json:"errors_in"`
	ErrorsOut float32 `json:"errors_out"`
	AdditionalProperties map[string]interface{}
}

type _NetworkDevicePortCounters NetworkDevicePortCounters

// NewNetworkDevicePortCounters instantiates a new NetworkDevicePortCounters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkDevicePortCounters(octetsIn float32, octetsOut float32, packetsIn float32, packetsOut float32, errorsIn float32, errorsOut float32) *NetworkDevicePortCounters {
	this := NetworkDevicePortCounters{}
	this.OctetsIn = octetsIn
	this.OctetsOut = octetsOut
	this.PacketsIn = packetsIn
	this.PacketsOut = packetsOut
	this.ErrorsIn = errorsIn
	this.ErrorsOut = errorsOut
	return &this
}

// NewNetworkDevicePortCountersWithDefaults instantiates a new NetworkDevicePortCounters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkDevicePortCountersWithDefaults() *NetworkDevicePortCounters {
	this := NetworkDevicePortCounters{}
	return &this
}

// GetOctetsIn returns the OctetsIn field value
func (o *NetworkDevicePortCounters) GetOctetsIn() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.OctetsIn
}

// GetOctetsInOk returns a tuple with the OctetsIn field value
// and a boolean to check if the value has been set.
func (o *NetworkDevicePortCounters) GetOctetsInOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OctetsIn, true
}

// SetOctetsIn sets field value
func (o *NetworkDevicePortCounters) SetOctetsIn(v float32) {
	o.OctetsIn = v
}

// GetOctetsOut returns the OctetsOut field value
func (o *NetworkDevicePortCounters) GetOctetsOut() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.OctetsOut
}

// GetOctetsOutOk returns a tuple with the OctetsOut field value
// and a boolean to check if the value has been set.
func (o *NetworkDevicePortCounters) GetOctetsOutOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OctetsOut, true
}

// SetOctetsOut sets field value
func (o *NetworkDevicePortCounters) SetOctetsOut(v float32) {
	o.OctetsOut = v
}

// GetPacketsIn returns the PacketsIn field value
func (o *NetworkDevicePortCounters) GetPacketsIn() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PacketsIn
}

// GetPacketsInOk returns a tuple with the PacketsIn field value
// and a boolean to check if the value has been set.
func (o *NetworkDevicePortCounters) GetPacketsInOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PacketsIn, true
}

// SetPacketsIn sets field value
func (o *NetworkDevicePortCounters) SetPacketsIn(v float32) {
	o.PacketsIn = v
}

// GetPacketsOut returns the PacketsOut field value
func (o *NetworkDevicePortCounters) GetPacketsOut() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PacketsOut
}

// GetPacketsOutOk returns a tuple with the PacketsOut field value
// and a boolean to check if the value has been set.
func (o *NetworkDevicePortCounters) GetPacketsOutOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PacketsOut, true
}

// SetPacketsOut sets field value
func (o *NetworkDevicePortCounters) SetPacketsOut(v float32) {
	o.PacketsOut = v
}

// GetErrorsIn returns the ErrorsIn field value
func (o *NetworkDevicePortCounters) GetErrorsIn() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ErrorsIn
}

// GetErrorsInOk returns a tuple with the ErrorsIn field value
// and a boolean to check if the value has been set.
func (o *NetworkDevicePortCounters) GetErrorsInOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorsIn, true
}

// SetErrorsIn sets field value
func (o *NetworkDevicePortCounters) SetErrorsIn(v float32) {
	o.ErrorsIn = v
}

// GetErrorsOut returns the ErrorsOut field value
func (o *NetworkDevicePortCounters) GetErrorsOut() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ErrorsOut
}

// GetErrorsOutOk returns a tuple with the ErrorsOut field value
// and a boolean to check if the value has been set.
func (o *NetworkDevicePortCounters) GetErrorsOutOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorsOut, true
}

// SetErrorsOut sets field value
func (o *NetworkDevicePortCounters) SetErrorsOut(v float32) {
	o.ErrorsOut = v
}

func (o NetworkDevicePortCounters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkDevicePortCounters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["octets_in"] = o.OctetsIn
	toSerialize["octets_out"] = o.OctetsOut
	toSerialize["packets_in"] = o.PacketsIn
	toSerialize["packets_out"] = o.PacketsOut
	toSerialize["errors_in"] = o.ErrorsIn
	toSerialize["errors_out"] = o.ErrorsOut

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkDevicePortCounters) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"octets_in",
		"octets_out",
		"packets_in",
		"packets_out",
		"errors_in",
		"errors_out",
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

	varNetworkDevicePortCounters := _NetworkDevicePortCounters{}

	err = json.Unmarshal(data, &varNetworkDevicePortCounters)

	if err != nil {
		return err
	}

	*o = NetworkDevicePortCounters(varNetworkDevicePortCounters)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "octets_in")
		delete(additionalProperties, "octets_out")
		delete(additionalProperties, "packets_in")
		delete(additionalProperties, "packets_out")
		delete(additionalProperties, "errors_in")
		delete(additionalProperties, "errors_out")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkDevicePortCounters struct {
	value *NetworkDevicePortCounters
	isSet bool
}

func (v NullableNetworkDevicePortCounters) Get() *NetworkDevicePortCounters {
	return v.value
}

func (v *NullableNetworkDevicePortCounters) Set(val *NetworkDevicePortCounters) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkDevicePortCounters) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkDevicePortCounters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkDevicePortCounters(val *NetworkDevicePortCounters) *NullableNetworkDevicePortCounters {
	return &NullableNetworkDevicePortCounters{value: val, isSet: true}
}

func (v NullableNetworkDevicePortCounters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkDevicePortCounters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


