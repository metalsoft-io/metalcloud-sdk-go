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

// checks if the CreateEndpointInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateEndpointInterface{}

// CreateEndpointInterface struct for CreateEndpointInterface
type CreateEndpointInterface struct {
	// Network device interface id
	NetworkDeviceInterfaceId float32 `json:"networkDeviceInterfaceId"`
	// Device interface mac address
	MacAddress *string `json:"macAddress,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateEndpointInterface CreateEndpointInterface

// NewCreateEndpointInterface instantiates a new CreateEndpointInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEndpointInterface(networkDeviceInterfaceId float32) *CreateEndpointInterface {
	this := CreateEndpointInterface{}
	this.NetworkDeviceInterfaceId = networkDeviceInterfaceId
	return &this
}

// NewCreateEndpointInterfaceWithDefaults instantiates a new CreateEndpointInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEndpointInterfaceWithDefaults() *CreateEndpointInterface {
	this := CreateEndpointInterface{}
	return &this
}

// GetNetworkDeviceInterfaceId returns the NetworkDeviceInterfaceId field value
func (o *CreateEndpointInterface) GetNetworkDeviceInterfaceId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NetworkDeviceInterfaceId
}

// GetNetworkDeviceInterfaceIdOk returns a tuple with the NetworkDeviceInterfaceId field value
// and a boolean to check if the value has been set.
func (o *CreateEndpointInterface) GetNetworkDeviceInterfaceIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkDeviceInterfaceId, true
}

// SetNetworkDeviceInterfaceId sets field value
func (o *CreateEndpointInterface) SetNetworkDeviceInterfaceId(v float32) {
	o.NetworkDeviceInterfaceId = v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *CreateEndpointInterface) GetMacAddress() string {
	if o == nil || IsNil(o.MacAddress) {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEndpointInterface) GetMacAddressOk() (*string, bool) {
	if o == nil || IsNil(o.MacAddress) {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *CreateEndpointInterface) HasMacAddress() bool {
	if o != nil && !IsNil(o.MacAddress) {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *CreateEndpointInterface) SetMacAddress(v string) {
	o.MacAddress = &v
}

func (o CreateEndpointInterface) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateEndpointInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["networkDeviceInterfaceId"] = o.NetworkDeviceInterfaceId
	if !IsNil(o.MacAddress) {
		toSerialize["macAddress"] = o.MacAddress
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateEndpointInterface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"networkDeviceInterfaceId",
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

	varCreateEndpointInterface := _CreateEndpointInterface{}

	err = json.Unmarshal(data, &varCreateEndpointInterface)

	if err != nil {
		return err
	}

	*o = CreateEndpointInterface(varCreateEndpointInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "networkDeviceInterfaceId")
		delete(additionalProperties, "macAddress")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateEndpointInterface struct {
	value *CreateEndpointInterface
	isSet bool
}

func (v NullableCreateEndpointInterface) Get() *CreateEndpointInterface {
	return v.value
}

func (v *NullableCreateEndpointInterface) Set(val *CreateEndpointInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEndpointInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEndpointInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEndpointInterface(val *CreateEndpointInterface) *NullableCreateEndpointInterface {
	return &NullableCreateEndpointInterface{value: val, isSet: true}
}

func (v NullableCreateEndpointInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEndpointInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


