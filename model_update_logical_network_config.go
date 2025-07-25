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

// checks if the UpdateLogicalNetworkConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateLogicalNetworkConfig{}

// UpdateLogicalNetworkConfig struct for UpdateLogicalNetworkConfig
type UpdateLogicalNetworkConfig struct {
	Vlan map[string]interface{} `json:"vlan"`
	Vxlan map[string]interface{} `json:"vxlan"`
	Ipv4 map[string]interface{} `json:"ipv4"`
	Ipv6 map[string]interface{} `json:"ipv6"`
	// Maximum Transmission Unit (MTU) in bytes
	Mtu NullableInt32 `json:"mtu,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateLogicalNetworkConfig UpdateLogicalNetworkConfig

// NewUpdateLogicalNetworkConfig instantiates a new UpdateLogicalNetworkConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateLogicalNetworkConfig(vlan map[string]interface{}, vxlan map[string]interface{}, ipv4 map[string]interface{}, ipv6 map[string]interface{}) *UpdateLogicalNetworkConfig {
	this := UpdateLogicalNetworkConfig{}
	this.Vlan = vlan
	this.Vxlan = vxlan
	this.Ipv4 = ipv4
	this.Ipv6 = ipv6
	return &this
}

// NewUpdateLogicalNetworkConfigWithDefaults instantiates a new UpdateLogicalNetworkConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateLogicalNetworkConfigWithDefaults() *UpdateLogicalNetworkConfig {
	this := UpdateLogicalNetworkConfig{}
	return &this
}

// GetVlan returns the Vlan field value
func (o *UpdateLogicalNetworkConfig) GetVlan() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value
// and a boolean to check if the value has been set.
func (o *UpdateLogicalNetworkConfig) GetVlanOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Vlan, true
}

// SetVlan sets field value
func (o *UpdateLogicalNetworkConfig) SetVlan(v map[string]interface{}) {
	o.Vlan = v
}

// GetVxlan returns the Vxlan field value
func (o *UpdateLogicalNetworkConfig) GetVxlan() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Vxlan
}

// GetVxlanOk returns a tuple with the Vxlan field value
// and a boolean to check if the value has been set.
func (o *UpdateLogicalNetworkConfig) GetVxlanOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Vxlan, true
}

// SetVxlan sets field value
func (o *UpdateLogicalNetworkConfig) SetVxlan(v map[string]interface{}) {
	o.Vxlan = v
}

// GetIpv4 returns the Ipv4 field value
func (o *UpdateLogicalNetworkConfig) GetIpv4() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value
// and a boolean to check if the value has been set.
func (o *UpdateLogicalNetworkConfig) GetIpv4Ok() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Ipv4, true
}

// SetIpv4 sets field value
func (o *UpdateLogicalNetworkConfig) SetIpv4(v map[string]interface{}) {
	o.Ipv4 = v
}

// GetIpv6 returns the Ipv6 field value
func (o *UpdateLogicalNetworkConfig) GetIpv6() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value
// and a boolean to check if the value has been set.
func (o *UpdateLogicalNetworkConfig) GetIpv6Ok() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Ipv6, true
}

// SetIpv6 sets field value
func (o *UpdateLogicalNetworkConfig) SetIpv6(v map[string]interface{}) {
	o.Ipv6 = v
}

// GetMtu returns the Mtu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateLogicalNetworkConfig) GetMtu() int32 {
	if o == nil || IsNil(o.Mtu.Get()) {
		var ret int32
		return ret
	}
	return *o.Mtu.Get()
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateLogicalNetworkConfig) GetMtuOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mtu.Get(), o.Mtu.IsSet()
}

// HasMtu returns a boolean if a field has been set.
func (o *UpdateLogicalNetworkConfig) HasMtu() bool {
	if o != nil && o.Mtu.IsSet() {
		return true
	}

	return false
}

// SetMtu gets a reference to the given NullableInt32 and assigns it to the Mtu field.
func (o *UpdateLogicalNetworkConfig) SetMtu(v int32) {
	o.Mtu.Set(&v)
}
// SetMtuNil sets the value for Mtu to be an explicit nil
func (o *UpdateLogicalNetworkConfig) SetMtuNil() {
	o.Mtu.Set(nil)
}

// UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
func (o *UpdateLogicalNetworkConfig) UnsetMtu() {
	o.Mtu.Unset()
}

func (o UpdateLogicalNetworkConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateLogicalNetworkConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["vlan"] = o.Vlan
	toSerialize["vxlan"] = o.Vxlan
	toSerialize["ipv4"] = o.Ipv4
	toSerialize["ipv6"] = o.Ipv6
	if o.Mtu.IsSet() {
		toSerialize["mtu"] = o.Mtu.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateLogicalNetworkConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"vlan",
		"vxlan",
		"ipv4",
		"ipv6",
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

	varUpdateLogicalNetworkConfig := _UpdateLogicalNetworkConfig{}

	err = json.Unmarshal(data, &varUpdateLogicalNetworkConfig)

	if err != nil {
		return err
	}

	*o = UpdateLogicalNetworkConfig(varUpdateLogicalNetworkConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "vlan")
		delete(additionalProperties, "vxlan")
		delete(additionalProperties, "ipv4")
		delete(additionalProperties, "ipv6")
		delete(additionalProperties, "mtu")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateLogicalNetworkConfig struct {
	value *UpdateLogicalNetworkConfig
	isSet bool
}

func (v NullableUpdateLogicalNetworkConfig) Get() *UpdateLogicalNetworkConfig {
	return v.value
}

func (v *NullableUpdateLogicalNetworkConfig) Set(val *UpdateLogicalNetworkConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateLogicalNetworkConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateLogicalNetworkConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateLogicalNetworkConfig(val *UpdateLogicalNetworkConfig) *NullableUpdateLogicalNetworkConfig {
	return &NullableUpdateLogicalNetworkConfig{value: val, isSet: true}
}

func (v NullableUpdateLogicalNetworkConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateLogicalNetworkConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


