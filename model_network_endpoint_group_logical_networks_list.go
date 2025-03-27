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

// checks if the NetworkEndpointGroupLogicalNetworksList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkEndpointGroupLogicalNetworksList{}

// NetworkEndpointGroupLogicalNetworksList struct for NetworkEndpointGroupLogicalNetworksList
type NetworkEndpointGroupLogicalNetworksList struct {
	Data []NetworkEndpointGroupLogicalNetworkDto `json:"data"`
	// Reference links
	Links []Link `json:"links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkEndpointGroupLogicalNetworksList NetworkEndpointGroupLogicalNetworksList

// NewNetworkEndpointGroupLogicalNetworksList instantiates a new NetworkEndpointGroupLogicalNetworksList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkEndpointGroupLogicalNetworksList(data []NetworkEndpointGroupLogicalNetworkDto) *NetworkEndpointGroupLogicalNetworksList {
	this := NetworkEndpointGroupLogicalNetworksList{}
	this.Data = data
	return &this
}

// NewNetworkEndpointGroupLogicalNetworksListWithDefaults instantiates a new NetworkEndpointGroupLogicalNetworksList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkEndpointGroupLogicalNetworksListWithDefaults() *NetworkEndpointGroupLogicalNetworksList {
	this := NetworkEndpointGroupLogicalNetworksList{}
	return &this
}

// GetData returns the Data field value
func (o *NetworkEndpointGroupLogicalNetworksList) GetData() []NetworkEndpointGroupLogicalNetworkDto {
	if o == nil {
		var ret []NetworkEndpointGroupLogicalNetworkDto
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *NetworkEndpointGroupLogicalNetworksList) GetDataOk() ([]NetworkEndpointGroupLogicalNetworkDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *NetworkEndpointGroupLogicalNetworksList) SetData(v []NetworkEndpointGroupLogicalNetworkDto) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *NetworkEndpointGroupLogicalNetworksList) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkEndpointGroupLogicalNetworksList) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *NetworkEndpointGroupLogicalNetworksList) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *NetworkEndpointGroupLogicalNetworksList) SetLinks(v []Link) {
	o.Links = v
}

func (o NetworkEndpointGroupLogicalNetworksList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkEndpointGroupLogicalNetworksList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkEndpointGroupLogicalNetworksList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varNetworkEndpointGroupLogicalNetworksList := _NetworkEndpointGroupLogicalNetworksList{}

	err = json.Unmarshal(data, &varNetworkEndpointGroupLogicalNetworksList)

	if err != nil {
		return err
	}

	*o = NetworkEndpointGroupLogicalNetworksList(varNetworkEndpointGroupLogicalNetworksList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkEndpointGroupLogicalNetworksList struct {
	value *NetworkEndpointGroupLogicalNetworksList
	isSet bool
}

func (v NullableNetworkEndpointGroupLogicalNetworksList) Get() *NetworkEndpointGroupLogicalNetworksList {
	return v.value
}

func (v *NullableNetworkEndpointGroupLogicalNetworksList) Set(val *NetworkEndpointGroupLogicalNetworksList) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkEndpointGroupLogicalNetworksList) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkEndpointGroupLogicalNetworksList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkEndpointGroupLogicalNetworksList(val *NetworkEndpointGroupLogicalNetworksList) *NullableNetworkEndpointGroupLogicalNetworksList {
	return &NullableNetworkEndpointGroupLogicalNetworksList{value: val, isSet: true}
}

func (v NullableNetworkEndpointGroupLogicalNetworksList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkEndpointGroupLogicalNetworksList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


