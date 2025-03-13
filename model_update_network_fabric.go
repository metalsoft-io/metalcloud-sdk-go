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

// checks if the UpdateNetworkFabric type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkFabric{}

// UpdateNetworkFabric struct for UpdateNetworkFabric
type UpdateNetworkFabric struct {
	// The ID of the site where the entity is located.
	SiteId *float32 `json:"siteId,omitempty"`
	// The network fabric name
	Name *string `json:"name,omitempty"`
	// Network fabric description
	Description *string `json:"description,omitempty"`
	FabricConfiguration NetworkFabricFabricConfiguration `json:"fabricConfiguration"`
	AdditionalProperties map[string]interface{}
}

type _UpdateNetworkFabric UpdateNetworkFabric

// NewUpdateNetworkFabric instantiates a new UpdateNetworkFabric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkFabric(fabricConfiguration NetworkFabricFabricConfiguration) *UpdateNetworkFabric {
	this := UpdateNetworkFabric{}
	this.FabricConfiguration = fabricConfiguration
	return &this
}

// NewUpdateNetworkFabricWithDefaults instantiates a new UpdateNetworkFabric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkFabricWithDefaults() *UpdateNetworkFabric {
	this := UpdateNetworkFabric{}
	return &this
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *UpdateNetworkFabric) GetSiteId() float32 {
	if o == nil || IsNil(o.SiteId) {
		var ret float32
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFabric) GetSiteIdOk() (*float32, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *UpdateNetworkFabric) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given float32 and assigns it to the SiteId field.
func (o *UpdateNetworkFabric) SetSiteId(v float32) {
	o.SiteId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateNetworkFabric) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFabric) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateNetworkFabric) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateNetworkFabric) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateNetworkFabric) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFabric) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateNetworkFabric) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateNetworkFabric) SetDescription(v string) {
	o.Description = &v
}

// GetFabricConfiguration returns the FabricConfiguration field value
func (o *UpdateNetworkFabric) GetFabricConfiguration() NetworkFabricFabricConfiguration {
	if o == nil {
		var ret NetworkFabricFabricConfiguration
		return ret
	}

	return o.FabricConfiguration
}

// GetFabricConfigurationOk returns a tuple with the FabricConfiguration field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFabric) GetFabricConfigurationOk() (*NetworkFabricFabricConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FabricConfiguration, true
}

// SetFabricConfiguration sets field value
func (o *UpdateNetworkFabric) SetFabricConfiguration(v NetworkFabricFabricConfiguration) {
	o.FabricConfiguration = v
}

func (o UpdateNetworkFabric) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkFabric) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SiteId) {
		toSerialize["siteId"] = o.SiteId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["fabricConfiguration"] = o.FabricConfiguration

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateNetworkFabric) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fabricConfiguration",
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

	varUpdateNetworkFabric := _UpdateNetworkFabric{}

	err = json.Unmarshal(data, &varUpdateNetworkFabric)

	if err != nil {
		return err
	}

	*o = UpdateNetworkFabric(varUpdateNetworkFabric)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "siteId")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "fabricConfiguration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateNetworkFabric struct {
	value *UpdateNetworkFabric
	isSet bool
}

func (v NullableUpdateNetworkFabric) Get() *UpdateNetworkFabric {
	return v.value
}

func (v *NullableUpdateNetworkFabric) Set(val *UpdateNetworkFabric) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkFabric) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkFabric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkFabric(val *UpdateNetworkFabric) *NullableUpdateNetworkFabric {
	return &NullableUpdateNetworkFabric{value: val, isSet: true}
}

func (v NullableUpdateNetworkFabric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkFabric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


