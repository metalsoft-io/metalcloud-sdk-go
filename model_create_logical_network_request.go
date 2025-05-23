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

// CreateLogicalNetworkRequest - struct for CreateLogicalNetworkRequest
type CreateLogicalNetworkRequest struct {
	CreateVlanLogicalNetwork *CreateVlanLogicalNetwork
	CreateVxlanLogicalNetwork *CreateVxlanLogicalNetwork
}

// CreateVlanLogicalNetworkAsCreateLogicalNetworkRequest is a convenience function that returns CreateVlanLogicalNetwork wrapped in CreateLogicalNetworkRequest
func CreateVlanLogicalNetworkAsCreateLogicalNetworkRequest(v *CreateVlanLogicalNetwork) CreateLogicalNetworkRequest {
	return CreateLogicalNetworkRequest{
		CreateVlanLogicalNetwork: v,
	}
}

// CreateVxlanLogicalNetworkAsCreateLogicalNetworkRequest is a convenience function that returns CreateVxlanLogicalNetwork wrapped in CreateLogicalNetworkRequest
func CreateVxlanLogicalNetworkAsCreateLogicalNetworkRequest(v *CreateVxlanLogicalNetwork) CreateLogicalNetworkRequest {
	return CreateLogicalNetworkRequest{
		CreateVxlanLogicalNetwork: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateLogicalNetworkRequest) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'vlan'
	if jsonDict["kind"] == "vlan" {
		// try to unmarshal JSON data into CreateVlanLogicalNetwork
		err = json.Unmarshal(data, &dst.CreateVlanLogicalNetwork)
		if err == nil {
			return nil // data stored in dst.CreateVlanLogicalNetwork, return on the first match
		} else {
			dst.CreateVlanLogicalNetwork = nil
			return fmt.Errorf("failed to unmarshal CreateLogicalNetworkRequest as CreateVlanLogicalNetwork: %s", err.Error())
		}
	}

	// check if the discriminator value is 'vxlan'
	if jsonDict["kind"] == "vxlan" {
		// try to unmarshal JSON data into CreateVxlanLogicalNetwork
		err = json.Unmarshal(data, &dst.CreateVxlanLogicalNetwork)
		if err == nil {
			return nil // data stored in dst.CreateVxlanLogicalNetwork, return on the first match
		} else {
			dst.CreateVxlanLogicalNetwork = nil
			return fmt.Errorf("failed to unmarshal CreateLogicalNetworkRequest as CreateVxlanLogicalNetwork: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CreateVlanLogicalNetwork'
	if jsonDict["kind"] == "CreateVlanLogicalNetwork" {
		// try to unmarshal JSON data into CreateVlanLogicalNetwork
		err = json.Unmarshal(data, &dst.CreateVlanLogicalNetwork)
		if err == nil {
			return nil // data stored in dst.CreateVlanLogicalNetwork, return on the first match
		} else {
			dst.CreateVlanLogicalNetwork = nil
			return fmt.Errorf("failed to unmarshal CreateLogicalNetworkRequest as CreateVlanLogicalNetwork: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CreateVxlanLogicalNetwork'
	if jsonDict["kind"] == "CreateVxlanLogicalNetwork" {
		// try to unmarshal JSON data into CreateVxlanLogicalNetwork
		err = json.Unmarshal(data, &dst.CreateVxlanLogicalNetwork)
		if err == nil {
			return nil // data stored in dst.CreateVxlanLogicalNetwork, return on the first match
		} else {
			dst.CreateVxlanLogicalNetwork = nil
			return fmt.Errorf("failed to unmarshal CreateLogicalNetworkRequest as CreateVxlanLogicalNetwork: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateLogicalNetworkRequest) MarshalJSON() ([]byte, error) {
	if src.CreateVlanLogicalNetwork != nil {
		return json.Marshal(&src.CreateVlanLogicalNetwork)
	}

	if src.CreateVxlanLogicalNetwork != nil {
		return json.Marshal(&src.CreateVxlanLogicalNetwork)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateLogicalNetworkRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CreateVlanLogicalNetwork != nil {
		return obj.CreateVlanLogicalNetwork
	}

	if obj.CreateVxlanLogicalNetwork != nil {
		return obj.CreateVxlanLogicalNetwork
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj CreateLogicalNetworkRequest) GetActualInstanceValue() (interface{}) {
	if obj.CreateVlanLogicalNetwork != nil {
		return *obj.CreateVlanLogicalNetwork
	}

	if obj.CreateVxlanLogicalNetwork != nil {
		return *obj.CreateVxlanLogicalNetwork
	}

	// all schemas are nil
	return nil
}

type NullableCreateLogicalNetworkRequest struct {
	value *CreateLogicalNetworkRequest
	isSet bool
}

func (v NullableCreateLogicalNetworkRequest) Get() *CreateLogicalNetworkRequest {
	return v.value
}

func (v *NullableCreateLogicalNetworkRequest) Set(val *CreateLogicalNetworkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLogicalNetworkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLogicalNetworkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLogicalNetworkRequest(val *CreateLogicalNetworkRequest) *NullableCreateLogicalNetworkRequest {
	return &NullableCreateLogicalNetworkRequest{value: val, isSet: true}
}

func (v NullableCreateLogicalNetworkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLogicalNetworkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


