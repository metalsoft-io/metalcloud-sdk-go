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

// checks if the UpdateNetworkFabricDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkFabricDto{}

// UpdateNetworkFabricDto struct for UpdateNetworkFabricDto
type UpdateNetworkFabricDto struct {
	// The network fabric name
	Name string `json:"name"`
	// Network fabric description
	Description *string `json:"description,omitempty"`
	// The type of network fabric
	FabricType string `json:"fabricType"`
	FabricConfiguration NetworkFabricFabricConfiguration `json:"fabricConfiguration"`
	// Revision number of the entity
	Revision float32 `json:"revision"`
	// The network fabric ID.
	Id string `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _UpdateNetworkFabricDto UpdateNetworkFabricDto

// NewUpdateNetworkFabricDto instantiates a new UpdateNetworkFabricDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkFabricDto(name string, fabricType string, fabricConfiguration NetworkFabricFabricConfiguration, revision float32, id string) *UpdateNetworkFabricDto {
	this := UpdateNetworkFabricDto{}
	this.Name = name
	this.FabricType = fabricType
	this.FabricConfiguration = fabricConfiguration
	this.Revision = revision
	this.Id = id
	return &this
}

// NewUpdateNetworkFabricDtoWithDefaults instantiates a new UpdateNetworkFabricDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkFabricDtoWithDefaults() *UpdateNetworkFabricDto {
	this := UpdateNetworkFabricDto{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateNetworkFabricDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFabricDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateNetworkFabricDto) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateNetworkFabricDto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFabricDto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateNetworkFabricDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateNetworkFabricDto) SetDescription(v string) {
	o.Description = &v
}

// GetFabricType returns the FabricType field value
func (o *UpdateNetworkFabricDto) GetFabricType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FabricType
}

// GetFabricTypeOk returns a tuple with the FabricType field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFabricDto) GetFabricTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FabricType, true
}

// SetFabricType sets field value
func (o *UpdateNetworkFabricDto) SetFabricType(v string) {
	o.FabricType = v
}

// GetFabricConfiguration returns the FabricConfiguration field value
func (o *UpdateNetworkFabricDto) GetFabricConfiguration() NetworkFabricFabricConfiguration {
	if o == nil {
		var ret NetworkFabricFabricConfiguration
		return ret
	}

	return o.FabricConfiguration
}

// GetFabricConfigurationOk returns a tuple with the FabricConfiguration field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFabricDto) GetFabricConfigurationOk() (*NetworkFabricFabricConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FabricConfiguration, true
}

// SetFabricConfiguration sets field value
func (o *UpdateNetworkFabricDto) SetFabricConfiguration(v NetworkFabricFabricConfiguration) {
	o.FabricConfiguration = v
}

// GetRevision returns the Revision field value
func (o *UpdateNetworkFabricDto) GetRevision() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFabricDto) GetRevisionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *UpdateNetworkFabricDto) SetRevision(v float32) {
	o.Revision = v
}

// GetId returns the Id field value
func (o *UpdateNetworkFabricDto) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFabricDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateNetworkFabricDto) SetId(v string) {
	o.Id = v
}

func (o UpdateNetworkFabricDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkFabricDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["fabricType"] = o.FabricType
	toSerialize["fabricConfiguration"] = o.FabricConfiguration
	toSerialize["revision"] = o.Revision
	toSerialize["id"] = o.Id

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateNetworkFabricDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"fabricType",
		"fabricConfiguration",
		"revision",
		"id",
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

	varUpdateNetworkFabricDto := _UpdateNetworkFabricDto{}

	err = json.Unmarshal(data, &varUpdateNetworkFabricDto)

	if err != nil {
		return err
	}

	*o = UpdateNetworkFabricDto(varUpdateNetworkFabricDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "fabricType")
		delete(additionalProperties, "fabricConfiguration")
		delete(additionalProperties, "revision")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateNetworkFabricDto struct {
	value *UpdateNetworkFabricDto
	isSet bool
}

func (v NullableUpdateNetworkFabricDto) Get() *UpdateNetworkFabricDto {
	return v.value
}

func (v *NullableUpdateNetworkFabricDto) Set(val *UpdateNetworkFabricDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkFabricDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkFabricDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkFabricDto(val *UpdateNetworkFabricDto) *NullableUpdateNetworkFabricDto {
	return &NullableUpdateNetworkFabricDto{value: val, isSet: true}
}

func (v NullableUpdateNetworkFabricDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkFabricDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


