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

// checks if the NetworkDevicePort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkDevicePort{}

// NetworkDevicePort struct for NetworkDevicePort
type NetworkDevicePort struct {
	PortName string `json:"port_name"`
	Enabled bool `json:"enabled"`
	Active bool `json:"active"`
	LinkSpeed float32 `json:"link_speed"`
	LinkDuplex LinkDuplex `json:"link_duplex"`
	UtilizationIn float32 `json:"utilization_in"`
	UtilizationOut float32 `json:"utilization_out"`
	Counters NetworkDevicePortCounters `json:"counters"`
	AdditionalProperties map[string]interface{}
}

type _NetworkDevicePort NetworkDevicePort

// NewNetworkDevicePort instantiates a new NetworkDevicePort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkDevicePort(portName string, enabled bool, active bool, linkSpeed float32, linkDuplex LinkDuplex, utilizationIn float32, utilizationOut float32, counters NetworkDevicePortCounters) *NetworkDevicePort {
	this := NetworkDevicePort{}
	this.PortName = portName
	this.Enabled = enabled
	this.Active = active
	this.LinkSpeed = linkSpeed
	this.LinkDuplex = linkDuplex
	this.UtilizationIn = utilizationIn
	this.UtilizationOut = utilizationOut
	this.Counters = counters
	return &this
}

// NewNetworkDevicePortWithDefaults instantiates a new NetworkDevicePort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkDevicePortWithDefaults() *NetworkDevicePort {
	this := NetworkDevicePort{}
	return &this
}

// GetPortName returns the PortName field value
func (o *NetworkDevicePort) GetPortName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PortName
}

// GetPortNameOk returns a tuple with the PortName field value
// and a boolean to check if the value has been set.
func (o *NetworkDevicePort) GetPortNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PortName, true
}

// SetPortName sets field value
func (o *NetworkDevicePort) SetPortName(v string) {
	o.PortName = v
}

// GetEnabled returns the Enabled field value
func (o *NetworkDevicePort) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *NetworkDevicePort) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *NetworkDevicePort) SetEnabled(v bool) {
	o.Enabled = v
}

// GetActive returns the Active field value
func (o *NetworkDevicePort) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *NetworkDevicePort) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *NetworkDevicePort) SetActive(v bool) {
	o.Active = v
}

// GetLinkSpeed returns the LinkSpeed field value
func (o *NetworkDevicePort) GetLinkSpeed() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.LinkSpeed
}

// GetLinkSpeedOk returns a tuple with the LinkSpeed field value
// and a boolean to check if the value has been set.
func (o *NetworkDevicePort) GetLinkSpeedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkSpeed, true
}

// SetLinkSpeed sets field value
func (o *NetworkDevicePort) SetLinkSpeed(v float32) {
	o.LinkSpeed = v
}

// GetLinkDuplex returns the LinkDuplex field value
func (o *NetworkDevicePort) GetLinkDuplex() LinkDuplex {
	if o == nil {
		var ret LinkDuplex
		return ret
	}

	return o.LinkDuplex
}

// GetLinkDuplexOk returns a tuple with the LinkDuplex field value
// and a boolean to check if the value has been set.
func (o *NetworkDevicePort) GetLinkDuplexOk() (*LinkDuplex, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkDuplex, true
}

// SetLinkDuplex sets field value
func (o *NetworkDevicePort) SetLinkDuplex(v LinkDuplex) {
	o.LinkDuplex = v
}

// GetUtilizationIn returns the UtilizationIn field value
func (o *NetworkDevicePort) GetUtilizationIn() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UtilizationIn
}

// GetUtilizationInOk returns a tuple with the UtilizationIn field value
// and a boolean to check if the value has been set.
func (o *NetworkDevicePort) GetUtilizationInOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UtilizationIn, true
}

// SetUtilizationIn sets field value
func (o *NetworkDevicePort) SetUtilizationIn(v float32) {
	o.UtilizationIn = v
}

// GetUtilizationOut returns the UtilizationOut field value
func (o *NetworkDevicePort) GetUtilizationOut() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UtilizationOut
}

// GetUtilizationOutOk returns a tuple with the UtilizationOut field value
// and a boolean to check if the value has been set.
func (o *NetworkDevicePort) GetUtilizationOutOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UtilizationOut, true
}

// SetUtilizationOut sets field value
func (o *NetworkDevicePort) SetUtilizationOut(v float32) {
	o.UtilizationOut = v
}

// GetCounters returns the Counters field value
func (o *NetworkDevicePort) GetCounters() NetworkDevicePortCounters {
	if o == nil {
		var ret NetworkDevicePortCounters
		return ret
	}

	return o.Counters
}

// GetCountersOk returns a tuple with the Counters field value
// and a boolean to check if the value has been set.
func (o *NetworkDevicePort) GetCountersOk() (*NetworkDevicePortCounters, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Counters, true
}

// SetCounters sets field value
func (o *NetworkDevicePort) SetCounters(v NetworkDevicePortCounters) {
	o.Counters = v
}

func (o NetworkDevicePort) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkDevicePort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["port_name"] = o.PortName
	toSerialize["enabled"] = o.Enabled
	toSerialize["active"] = o.Active
	toSerialize["link_speed"] = o.LinkSpeed
	toSerialize["link_duplex"] = o.LinkDuplex
	toSerialize["utilization_in"] = o.UtilizationIn
	toSerialize["utilization_out"] = o.UtilizationOut
	toSerialize["counters"] = o.Counters

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkDevicePort) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"port_name",
		"enabled",
		"active",
		"link_speed",
		"link_duplex",
		"utilization_in",
		"utilization_out",
		"counters",
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

	varNetworkDevicePort := _NetworkDevicePort{}

	err = json.Unmarshal(data, &varNetworkDevicePort)

	if err != nil {
		return err
	}

	*o = NetworkDevicePort(varNetworkDevicePort)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "port_name")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "active")
		delete(additionalProperties, "link_speed")
		delete(additionalProperties, "link_duplex")
		delete(additionalProperties, "utilization_in")
		delete(additionalProperties, "utilization_out")
		delete(additionalProperties, "counters")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkDevicePort struct {
	value *NetworkDevicePort
	isSet bool
}

func (v NullableNetworkDevicePort) Get() *NetworkDevicePort {
	return v.value
}

func (v *NullableNetworkDevicePort) Set(val *NetworkDevicePort) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkDevicePort) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkDevicePort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkDevicePort(val *NetworkDevicePort) *NullableNetworkDevicePort {
	return &NullableNetworkDevicePort{value: val, isSet: true}
}

func (v NullableNetworkDevicePort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkDevicePort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


