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

// checks if the ServerRegistrationBiosProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerRegistrationBiosProfile{}

// ServerRegistrationBiosProfile struct for ServerRegistrationBiosProfile
type ServerRegistrationBiosProfile struct {
	// BIOS attribute name
	AttributeName string `json:"attributeName"`
	// BIOS attribute value
	AttributeValue string `json:"attributeValue"`
	// Server vendor type
	ServerVendor string `json:"serverVendor"`
	// Server model
	ServerModel *string `json:"serverModel,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServerRegistrationBiosProfile ServerRegistrationBiosProfile

// NewServerRegistrationBiosProfile instantiates a new ServerRegistrationBiosProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerRegistrationBiosProfile(attributeName string, attributeValue string, serverVendor string) *ServerRegistrationBiosProfile {
	this := ServerRegistrationBiosProfile{}
	this.AttributeName = attributeName
	this.AttributeValue = attributeValue
	this.ServerVendor = serverVendor
	return &this
}

// NewServerRegistrationBiosProfileWithDefaults instantiates a new ServerRegistrationBiosProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerRegistrationBiosProfileWithDefaults() *ServerRegistrationBiosProfile {
	this := ServerRegistrationBiosProfile{}
	return &this
}

// GetAttributeName returns the AttributeName field value
func (o *ServerRegistrationBiosProfile) GetAttributeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeName
}

// GetAttributeNameOk returns a tuple with the AttributeName field value
// and a boolean to check if the value has been set.
func (o *ServerRegistrationBiosProfile) GetAttributeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeName, true
}

// SetAttributeName sets field value
func (o *ServerRegistrationBiosProfile) SetAttributeName(v string) {
	o.AttributeName = v
}

// GetAttributeValue returns the AttributeValue field value
func (o *ServerRegistrationBiosProfile) GetAttributeValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeValue
}

// GetAttributeValueOk returns a tuple with the AttributeValue field value
// and a boolean to check if the value has been set.
func (o *ServerRegistrationBiosProfile) GetAttributeValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeValue, true
}

// SetAttributeValue sets field value
func (o *ServerRegistrationBiosProfile) SetAttributeValue(v string) {
	o.AttributeValue = v
}

// GetServerVendor returns the ServerVendor field value
func (o *ServerRegistrationBiosProfile) GetServerVendor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerVendor
}

// GetServerVendorOk returns a tuple with the ServerVendor field value
// and a boolean to check if the value has been set.
func (o *ServerRegistrationBiosProfile) GetServerVendorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerVendor, true
}

// SetServerVendor sets field value
func (o *ServerRegistrationBiosProfile) SetServerVendor(v string) {
	o.ServerVendor = v
}

// GetServerModel returns the ServerModel field value if set, zero value otherwise.
func (o *ServerRegistrationBiosProfile) GetServerModel() string {
	if o == nil || IsNil(o.ServerModel) {
		var ret string
		return ret
	}
	return *o.ServerModel
}

// GetServerModelOk returns a tuple with the ServerModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerRegistrationBiosProfile) GetServerModelOk() (*string, bool) {
	if o == nil || IsNil(o.ServerModel) {
		return nil, false
	}
	return o.ServerModel, true
}

// HasServerModel returns a boolean if a field has been set.
func (o *ServerRegistrationBiosProfile) HasServerModel() bool {
	if o != nil && !IsNil(o.ServerModel) {
		return true
	}

	return false
}

// SetServerModel gets a reference to the given string and assigns it to the ServerModel field.
func (o *ServerRegistrationBiosProfile) SetServerModel(v string) {
	o.ServerModel = &v
}

func (o ServerRegistrationBiosProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerRegistrationBiosProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attributeName"] = o.AttributeName
	toSerialize["attributeValue"] = o.AttributeValue
	toSerialize["serverVendor"] = o.ServerVendor
	if !IsNil(o.ServerModel) {
		toSerialize["serverModel"] = o.ServerModel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServerRegistrationBiosProfile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attributeName",
		"attributeValue",
		"serverVendor",
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

	varServerRegistrationBiosProfile := _ServerRegistrationBiosProfile{}

	err = json.Unmarshal(data, &varServerRegistrationBiosProfile)

	if err != nil {
		return err
	}

	*o = ServerRegistrationBiosProfile(varServerRegistrationBiosProfile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "attributeName")
		delete(additionalProperties, "attributeValue")
		delete(additionalProperties, "serverVendor")
		delete(additionalProperties, "serverModel")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServerRegistrationBiosProfile struct {
	value *ServerRegistrationBiosProfile
	isSet bool
}

func (v NullableServerRegistrationBiosProfile) Get() *ServerRegistrationBiosProfile {
	return v.value
}

func (v *NullableServerRegistrationBiosProfile) Set(val *ServerRegistrationBiosProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableServerRegistrationBiosProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableServerRegistrationBiosProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerRegistrationBiosProfile(val *ServerRegistrationBiosProfile) *NullableServerRegistrationBiosProfile {
	return &NullableServerRegistrationBiosProfile{value: val, isSet: true}
}

func (v NullableServerRegistrationBiosProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerRegistrationBiosProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


