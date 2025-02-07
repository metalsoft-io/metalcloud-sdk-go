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

// checks if the VirtualFabricDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualFabricDto{}

// VirtualFabricDto struct for VirtualFabricDto
type VirtualFabricDto struct {
	DefaultNetworkProfileId float32 `json:"defaultNetworkProfileId"`
	GnmiMonitoringEnabled *bool `json:"gnmiMonitoringEnabled,omitempty"`
	SyslogMonitoringEnabled *bool `json:"syslogMonitoringEnabled,omitempty"`
	ZeroTouchEnabled *bool `json:"zeroTouchEnabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualFabricDto VirtualFabricDto

// NewVirtualFabricDto instantiates a new VirtualFabricDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualFabricDto(defaultNetworkProfileId float32) *VirtualFabricDto {
	this := VirtualFabricDto{}
	this.DefaultNetworkProfileId = defaultNetworkProfileId
	return &this
}

// NewVirtualFabricDtoWithDefaults instantiates a new VirtualFabricDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualFabricDtoWithDefaults() *VirtualFabricDto {
	this := VirtualFabricDto{}
	return &this
}

// GetDefaultNetworkProfileId returns the DefaultNetworkProfileId field value
func (o *VirtualFabricDto) GetDefaultNetworkProfileId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DefaultNetworkProfileId
}

// GetDefaultNetworkProfileIdOk returns a tuple with the DefaultNetworkProfileId field value
// and a boolean to check if the value has been set.
func (o *VirtualFabricDto) GetDefaultNetworkProfileIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultNetworkProfileId, true
}

// SetDefaultNetworkProfileId sets field value
func (o *VirtualFabricDto) SetDefaultNetworkProfileId(v float32) {
	o.DefaultNetworkProfileId = v
}

// GetGnmiMonitoringEnabled returns the GnmiMonitoringEnabled field value if set, zero value otherwise.
func (o *VirtualFabricDto) GetGnmiMonitoringEnabled() bool {
	if o == nil || IsNil(o.GnmiMonitoringEnabled) {
		var ret bool
		return ret
	}
	return *o.GnmiMonitoringEnabled
}

// GetGnmiMonitoringEnabledOk returns a tuple with the GnmiMonitoringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualFabricDto) GetGnmiMonitoringEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.GnmiMonitoringEnabled) {
		return nil, false
	}
	return o.GnmiMonitoringEnabled, true
}

// HasGnmiMonitoringEnabled returns a boolean if a field has been set.
func (o *VirtualFabricDto) HasGnmiMonitoringEnabled() bool {
	if o != nil && !IsNil(o.GnmiMonitoringEnabled) {
		return true
	}

	return false
}

// SetGnmiMonitoringEnabled gets a reference to the given bool and assigns it to the GnmiMonitoringEnabled field.
func (o *VirtualFabricDto) SetGnmiMonitoringEnabled(v bool) {
	o.GnmiMonitoringEnabled = &v
}

// GetSyslogMonitoringEnabled returns the SyslogMonitoringEnabled field value if set, zero value otherwise.
func (o *VirtualFabricDto) GetSyslogMonitoringEnabled() bool {
	if o == nil || IsNil(o.SyslogMonitoringEnabled) {
		var ret bool
		return ret
	}
	return *o.SyslogMonitoringEnabled
}

// GetSyslogMonitoringEnabledOk returns a tuple with the SyslogMonitoringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualFabricDto) GetSyslogMonitoringEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SyslogMonitoringEnabled) {
		return nil, false
	}
	return o.SyslogMonitoringEnabled, true
}

// HasSyslogMonitoringEnabled returns a boolean if a field has been set.
func (o *VirtualFabricDto) HasSyslogMonitoringEnabled() bool {
	if o != nil && !IsNil(o.SyslogMonitoringEnabled) {
		return true
	}

	return false
}

// SetSyslogMonitoringEnabled gets a reference to the given bool and assigns it to the SyslogMonitoringEnabled field.
func (o *VirtualFabricDto) SetSyslogMonitoringEnabled(v bool) {
	o.SyslogMonitoringEnabled = &v
}

// GetZeroTouchEnabled returns the ZeroTouchEnabled field value if set, zero value otherwise.
func (o *VirtualFabricDto) GetZeroTouchEnabled() bool {
	if o == nil || IsNil(o.ZeroTouchEnabled) {
		var ret bool
		return ret
	}
	return *o.ZeroTouchEnabled
}

// GetZeroTouchEnabledOk returns a tuple with the ZeroTouchEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualFabricDto) GetZeroTouchEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ZeroTouchEnabled) {
		return nil, false
	}
	return o.ZeroTouchEnabled, true
}

// HasZeroTouchEnabled returns a boolean if a field has been set.
func (o *VirtualFabricDto) HasZeroTouchEnabled() bool {
	if o != nil && !IsNil(o.ZeroTouchEnabled) {
		return true
	}

	return false
}

// SetZeroTouchEnabled gets a reference to the given bool and assigns it to the ZeroTouchEnabled field.
func (o *VirtualFabricDto) SetZeroTouchEnabled(v bool) {
	o.ZeroTouchEnabled = &v
}

func (o VirtualFabricDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualFabricDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["defaultNetworkProfileId"] = o.DefaultNetworkProfileId
	if !IsNil(o.GnmiMonitoringEnabled) {
		toSerialize["gnmiMonitoringEnabled"] = o.GnmiMonitoringEnabled
	}
	if !IsNil(o.SyslogMonitoringEnabled) {
		toSerialize["syslogMonitoringEnabled"] = o.SyslogMonitoringEnabled
	}
	if !IsNil(o.ZeroTouchEnabled) {
		toSerialize["zeroTouchEnabled"] = o.ZeroTouchEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VirtualFabricDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"defaultNetworkProfileId",
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

	varVirtualFabricDto := _VirtualFabricDto{}

	err = json.Unmarshal(data, &varVirtualFabricDto)

	if err != nil {
		return err
	}

	*o = VirtualFabricDto(varVirtualFabricDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "defaultNetworkProfileId")
		delete(additionalProperties, "gnmiMonitoringEnabled")
		delete(additionalProperties, "syslogMonitoringEnabled")
		delete(additionalProperties, "zeroTouchEnabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualFabricDto struct {
	value *VirtualFabricDto
	isSet bool
}

func (v NullableVirtualFabricDto) Get() *VirtualFabricDto {
	return v.value
}

func (v *NullableVirtualFabricDto) Set(val *VirtualFabricDto) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualFabricDto) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualFabricDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualFabricDto(val *VirtualFabricDto) *NullableVirtualFabricDto {
	return &NullableVirtualFabricDto{value: val, isSet: true}
}

func (v NullableVirtualFabricDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualFabricDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


