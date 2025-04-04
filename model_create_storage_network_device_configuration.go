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

// checks if the CreateStorageNetworkDeviceConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateStorageNetworkDeviceConfiguration{}

// CreateStorageNetworkDeviceConfiguration struct for CreateStorageNetworkDeviceConfiguration
type CreateStorageNetworkDeviceConfiguration struct {
	// Id of the network equipment
	NetworkEquipmentId float32 `json:"networkEquipmentId"`
	// Identifier of the storage physical interface
	StoragePhysicalInterfaceIdentifier string `json:"storagePhysicalInterfaceIdentifier"`
	// Identifier of the network equipment interface
	NetworkEquipmentInterfaceIdentifier string `json:"networkEquipmentInterfaceIdentifier"`
	// Array of VLANS for the network equipment interface
	NetworkEquipmentInterfaceVlans []float32 `json:"networkEquipmentInterfaceVlans"`
	AdditionalProperties map[string]interface{}
}

type _CreateStorageNetworkDeviceConfiguration CreateStorageNetworkDeviceConfiguration

// NewCreateStorageNetworkDeviceConfiguration instantiates a new CreateStorageNetworkDeviceConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateStorageNetworkDeviceConfiguration(networkEquipmentId float32, storagePhysicalInterfaceIdentifier string, networkEquipmentInterfaceIdentifier string, networkEquipmentInterfaceVlans []float32) *CreateStorageNetworkDeviceConfiguration {
	this := CreateStorageNetworkDeviceConfiguration{}
	this.NetworkEquipmentId = networkEquipmentId
	this.StoragePhysicalInterfaceIdentifier = storagePhysicalInterfaceIdentifier
	this.NetworkEquipmentInterfaceIdentifier = networkEquipmentInterfaceIdentifier
	this.NetworkEquipmentInterfaceVlans = networkEquipmentInterfaceVlans
	return &this
}

// NewCreateStorageNetworkDeviceConfigurationWithDefaults instantiates a new CreateStorageNetworkDeviceConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateStorageNetworkDeviceConfigurationWithDefaults() *CreateStorageNetworkDeviceConfiguration {
	this := CreateStorageNetworkDeviceConfiguration{}
	return &this
}

// GetNetworkEquipmentId returns the NetworkEquipmentId field value
func (o *CreateStorageNetworkDeviceConfiguration) GetNetworkEquipmentId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NetworkEquipmentId
}

// GetNetworkEquipmentIdOk returns a tuple with the NetworkEquipmentId field value
// and a boolean to check if the value has been set.
func (o *CreateStorageNetworkDeviceConfiguration) GetNetworkEquipmentIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkEquipmentId, true
}

// SetNetworkEquipmentId sets field value
func (o *CreateStorageNetworkDeviceConfiguration) SetNetworkEquipmentId(v float32) {
	o.NetworkEquipmentId = v
}

// GetStoragePhysicalInterfaceIdentifier returns the StoragePhysicalInterfaceIdentifier field value
func (o *CreateStorageNetworkDeviceConfiguration) GetStoragePhysicalInterfaceIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StoragePhysicalInterfaceIdentifier
}

// GetStoragePhysicalInterfaceIdentifierOk returns a tuple with the StoragePhysicalInterfaceIdentifier field value
// and a boolean to check if the value has been set.
func (o *CreateStorageNetworkDeviceConfiguration) GetStoragePhysicalInterfaceIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StoragePhysicalInterfaceIdentifier, true
}

// SetStoragePhysicalInterfaceIdentifier sets field value
func (o *CreateStorageNetworkDeviceConfiguration) SetStoragePhysicalInterfaceIdentifier(v string) {
	o.StoragePhysicalInterfaceIdentifier = v
}

// GetNetworkEquipmentInterfaceIdentifier returns the NetworkEquipmentInterfaceIdentifier field value
func (o *CreateStorageNetworkDeviceConfiguration) GetNetworkEquipmentInterfaceIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkEquipmentInterfaceIdentifier
}

// GetNetworkEquipmentInterfaceIdentifierOk returns a tuple with the NetworkEquipmentInterfaceIdentifier field value
// and a boolean to check if the value has been set.
func (o *CreateStorageNetworkDeviceConfiguration) GetNetworkEquipmentInterfaceIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkEquipmentInterfaceIdentifier, true
}

// SetNetworkEquipmentInterfaceIdentifier sets field value
func (o *CreateStorageNetworkDeviceConfiguration) SetNetworkEquipmentInterfaceIdentifier(v string) {
	o.NetworkEquipmentInterfaceIdentifier = v
}

// GetNetworkEquipmentInterfaceVlans returns the NetworkEquipmentInterfaceVlans field value
func (o *CreateStorageNetworkDeviceConfiguration) GetNetworkEquipmentInterfaceVlans() []float32 {
	if o == nil {
		var ret []float32
		return ret
	}

	return o.NetworkEquipmentInterfaceVlans
}

// GetNetworkEquipmentInterfaceVlansOk returns a tuple with the NetworkEquipmentInterfaceVlans field value
// and a boolean to check if the value has been set.
func (o *CreateStorageNetworkDeviceConfiguration) GetNetworkEquipmentInterfaceVlansOk() ([]float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetworkEquipmentInterfaceVlans, true
}

// SetNetworkEquipmentInterfaceVlans sets field value
func (o *CreateStorageNetworkDeviceConfiguration) SetNetworkEquipmentInterfaceVlans(v []float32) {
	o.NetworkEquipmentInterfaceVlans = v
}

func (o CreateStorageNetworkDeviceConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateStorageNetworkDeviceConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["networkEquipmentId"] = o.NetworkEquipmentId
	toSerialize["storagePhysicalInterfaceIdentifier"] = o.StoragePhysicalInterfaceIdentifier
	toSerialize["networkEquipmentInterfaceIdentifier"] = o.NetworkEquipmentInterfaceIdentifier
	toSerialize["networkEquipmentInterfaceVlans"] = o.NetworkEquipmentInterfaceVlans

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateStorageNetworkDeviceConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"networkEquipmentId",
		"storagePhysicalInterfaceIdentifier",
		"networkEquipmentInterfaceIdentifier",
		"networkEquipmentInterfaceVlans",
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

	varCreateStorageNetworkDeviceConfiguration := _CreateStorageNetworkDeviceConfiguration{}

	err = json.Unmarshal(data, &varCreateStorageNetworkDeviceConfiguration)

	if err != nil {
		return err
	}

	*o = CreateStorageNetworkDeviceConfiguration(varCreateStorageNetworkDeviceConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "networkEquipmentId")
		delete(additionalProperties, "storagePhysicalInterfaceIdentifier")
		delete(additionalProperties, "networkEquipmentInterfaceIdentifier")
		delete(additionalProperties, "networkEquipmentInterfaceVlans")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateStorageNetworkDeviceConfiguration struct {
	value *CreateStorageNetworkDeviceConfiguration
	isSet bool
}

func (v NullableCreateStorageNetworkDeviceConfiguration) Get() *CreateStorageNetworkDeviceConfiguration {
	return v.value
}

func (v *NullableCreateStorageNetworkDeviceConfiguration) Set(val *CreateStorageNetworkDeviceConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateStorageNetworkDeviceConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateStorageNetworkDeviceConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateStorageNetworkDeviceConfiguration(val *CreateStorageNetworkDeviceConfiguration) *NullableCreateStorageNetworkDeviceConfiguration {
	return &NullableCreateStorageNetworkDeviceConfiguration{value: val, isSet: true}
}

func (v NullableCreateStorageNetworkDeviceConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateStorageNetworkDeviceConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


