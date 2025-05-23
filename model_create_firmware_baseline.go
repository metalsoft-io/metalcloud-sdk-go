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

// checks if the CreateFirmwareBaseline type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFirmwareBaseline{}

// CreateFirmwareBaseline struct for CreateFirmwareBaseline
type CreateFirmwareBaseline struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Catalog []string `json:"catalog,omitempty"`
	Level BaselineLevelType `json:"level"`
	LevelFilter []string `json:"levelFilter"`
	AdditionalProperties map[string]interface{}
}

type _CreateFirmwareBaseline CreateFirmwareBaseline

// NewCreateFirmwareBaseline instantiates a new CreateFirmwareBaseline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFirmwareBaseline(name string, level BaselineLevelType, levelFilter []string) *CreateFirmwareBaseline {
	this := CreateFirmwareBaseline{}
	this.Name = name
	this.Level = level
	this.LevelFilter = levelFilter
	return &this
}

// NewCreateFirmwareBaselineWithDefaults instantiates a new CreateFirmwareBaseline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFirmwareBaselineWithDefaults() *CreateFirmwareBaseline {
	this := CreateFirmwareBaseline{}
	return &this
}

// GetName returns the Name field value
func (o *CreateFirmwareBaseline) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateFirmwareBaseline) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateFirmwareBaseline) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateFirmwareBaseline) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFirmwareBaseline) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateFirmwareBaseline) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateFirmwareBaseline) SetDescription(v string) {
	o.Description = &v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *CreateFirmwareBaseline) GetCatalog() []string {
	if o == nil || IsNil(o.Catalog) {
		var ret []string
		return ret
	}
	return o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFirmwareBaseline) GetCatalogOk() ([]string, bool) {
	if o == nil || IsNil(o.Catalog) {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *CreateFirmwareBaseline) HasCatalog() bool {
	if o != nil && !IsNil(o.Catalog) {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given []string and assigns it to the Catalog field.
func (o *CreateFirmwareBaseline) SetCatalog(v []string) {
	o.Catalog = v
}

// GetLevel returns the Level field value
func (o *CreateFirmwareBaseline) GetLevel() BaselineLevelType {
	if o == nil {
		var ret BaselineLevelType
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *CreateFirmwareBaseline) GetLevelOk() (*BaselineLevelType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *CreateFirmwareBaseline) SetLevel(v BaselineLevelType) {
	o.Level = v
}

// GetLevelFilter returns the LevelFilter field value
func (o *CreateFirmwareBaseline) GetLevelFilter() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.LevelFilter
}

// GetLevelFilterOk returns a tuple with the LevelFilter field value
// and a boolean to check if the value has been set.
func (o *CreateFirmwareBaseline) GetLevelFilterOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LevelFilter, true
}

// SetLevelFilter sets field value
func (o *CreateFirmwareBaseline) SetLevelFilter(v []string) {
	o.LevelFilter = v
}

func (o CreateFirmwareBaseline) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFirmwareBaseline) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Catalog) {
		toSerialize["catalog"] = o.Catalog
	}
	toSerialize["level"] = o.Level
	toSerialize["levelFilter"] = o.LevelFilter

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateFirmwareBaseline) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"level",
		"levelFilter",
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

	varCreateFirmwareBaseline := _CreateFirmwareBaseline{}

	err = json.Unmarshal(data, &varCreateFirmwareBaseline)

	if err != nil {
		return err
	}

	*o = CreateFirmwareBaseline(varCreateFirmwareBaseline)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "catalog")
		delete(additionalProperties, "level")
		delete(additionalProperties, "levelFilter")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateFirmwareBaseline struct {
	value *CreateFirmwareBaseline
	isSet bool
}

func (v NullableCreateFirmwareBaseline) Get() *CreateFirmwareBaseline {
	return v.value
}

func (v *NullableCreateFirmwareBaseline) Set(val *CreateFirmwareBaseline) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFirmwareBaseline) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFirmwareBaseline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFirmwareBaseline(val *CreateFirmwareBaseline) *NullableCreateFirmwareBaseline {
	return &NullableCreateFirmwareBaseline{value: val, isSet: true}
}

func (v NullableCreateFirmwareBaseline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFirmwareBaseline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


