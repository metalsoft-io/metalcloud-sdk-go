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

// LogicalNetworkDataItem - struct for LogicalNetworkDataItem
type LogicalNetworkDataItem struct {
	VlanLogicalNetwork *VlanLogicalNetwork
	VxlanLogicalNetwork *VxlanLogicalNetwork
}

// VlanLogicalNetworkAsLogicalNetworkDataItem is a convenience function that returns VlanLogicalNetwork wrapped in LogicalNetworkDataItem
func VlanLogicalNetworkAsLogicalNetworkDataItem(v *VlanLogicalNetwork) LogicalNetworkDataItem {
	return LogicalNetworkDataItem{
		VlanLogicalNetwork: v,
	}
}

// VxlanLogicalNetworkAsLogicalNetworkDataItem is a convenience function that returns VxlanLogicalNetwork wrapped in LogicalNetworkDataItem
func VxlanLogicalNetworkAsLogicalNetworkDataItem(v *VxlanLogicalNetwork) LogicalNetworkDataItem {
	return LogicalNetworkDataItem{
		VxlanLogicalNetwork: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *LogicalNetworkDataItem) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'vlan'
	if jsonDict["kind"] == "vlan" {
		// try to unmarshal JSON data into VlanLogicalNetwork
		err = json.Unmarshal(data, &dst.VlanLogicalNetwork)
		if err == nil {
			return nil // data stored in dst.VlanLogicalNetwork, return on the first match
		} else {
			dst.VlanLogicalNetwork = nil
			return fmt.Errorf("failed to unmarshal LogicalNetworkDataItem as VlanLogicalNetwork: %s", err.Error())
		}
	}

	// check if the discriminator value is 'vxlan'
	if jsonDict["kind"] == "vxlan" {
		// try to unmarshal JSON data into VxlanLogicalNetwork
		err = json.Unmarshal(data, &dst.VxlanLogicalNetwork)
		if err == nil {
			return nil // data stored in dst.VxlanLogicalNetwork, return on the first match
		} else {
			dst.VxlanLogicalNetwork = nil
			return fmt.Errorf("failed to unmarshal LogicalNetworkDataItem as VxlanLogicalNetwork: %s", err.Error())
		}
	}

	// check if the discriminator value is 'VlanLogicalNetwork'
	if jsonDict["kind"] == "VlanLogicalNetwork" {
		// try to unmarshal JSON data into VlanLogicalNetwork
		err = json.Unmarshal(data, &dst.VlanLogicalNetwork)
		if err == nil {
			return nil // data stored in dst.VlanLogicalNetwork, return on the first match
		} else {
			dst.VlanLogicalNetwork = nil
			return fmt.Errorf("failed to unmarshal LogicalNetworkDataItem as VlanLogicalNetwork: %s", err.Error())
		}
	}

	// check if the discriminator value is 'VxlanLogicalNetwork'
	if jsonDict["kind"] == "VxlanLogicalNetwork" {
		// try to unmarshal JSON data into VxlanLogicalNetwork
		err = json.Unmarshal(data, &dst.VxlanLogicalNetwork)
		if err == nil {
			return nil // data stored in dst.VxlanLogicalNetwork, return on the first match
		} else {
			dst.VxlanLogicalNetwork = nil
			return fmt.Errorf("failed to unmarshal LogicalNetworkDataItem as VxlanLogicalNetwork: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src LogicalNetworkDataItem) MarshalJSON() ([]byte, error) {
	if src.VlanLogicalNetwork != nil {
		return json.Marshal(&src.VlanLogicalNetwork)
	}

	if src.VxlanLogicalNetwork != nil {
		return json.Marshal(&src.VxlanLogicalNetwork)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *LogicalNetworkDataItem) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.VlanLogicalNetwork != nil {
		return obj.VlanLogicalNetwork
	}

	if obj.VxlanLogicalNetwork != nil {
		return obj.VxlanLogicalNetwork
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj LogicalNetworkDataItem) GetActualInstanceValue() (interface{}) {
	if obj.VlanLogicalNetwork != nil {
		return *obj.VlanLogicalNetwork
	}

	if obj.VxlanLogicalNetwork != nil {
		return *obj.VxlanLogicalNetwork
	}

	// all schemas are nil
	return nil
}

type NullableLogicalNetworkDataItem struct {
	value *LogicalNetworkDataItem
	isSet bool
}

func (v NullableLogicalNetworkDataItem) Get() *LogicalNetworkDataItem {
	return v.value
}

func (v *NullableLogicalNetworkDataItem) Set(val *LogicalNetworkDataItem) {
	v.value = val
	v.isSet = true
}

func (v NullableLogicalNetworkDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullableLogicalNetworkDataItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogicalNetworkDataItem(val *LogicalNetworkDataItem) *NullableLogicalNetworkDataItem {
	return &NullableLogicalNetworkDataItem{value: val, isSet: true}
}

func (v NullableLogicalNetworkDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogicalNetworkDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


