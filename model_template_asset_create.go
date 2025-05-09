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

// checks if the TemplateAssetCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateAssetCreate{}

// TemplateAssetCreate struct for TemplateAssetCreate
type TemplateAssetCreate struct {
	// The ID of the OS template that this template asset belongs to
	TemplateId int32 `json:"templateId"`
	// The template asset usage
	Usage string `json:"usage"`
	File TemplateAssetFile `json:"file"`
	// The tags associated with the template asset
	Tags []string `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TemplateAssetCreate TemplateAssetCreate

// NewTemplateAssetCreate instantiates a new TemplateAssetCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateAssetCreate(templateId int32, usage string, file TemplateAssetFile) *TemplateAssetCreate {
	this := TemplateAssetCreate{}
	this.TemplateId = templateId
	this.Usage = usage
	this.File = file
	return &this
}

// NewTemplateAssetCreateWithDefaults instantiates a new TemplateAssetCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateAssetCreateWithDefaults() *TemplateAssetCreate {
	this := TemplateAssetCreate{}
	return &this
}

// GetTemplateId returns the TemplateId field value
func (o *TemplateAssetCreate) GetTemplateId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value
// and a boolean to check if the value has been set.
func (o *TemplateAssetCreate) GetTemplateIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateId, true
}

// SetTemplateId sets field value
func (o *TemplateAssetCreate) SetTemplateId(v int32) {
	o.TemplateId = v
}

// GetUsage returns the Usage field value
func (o *TemplateAssetCreate) GetUsage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
func (o *TemplateAssetCreate) GetUsageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Usage, true
}

// SetUsage sets field value
func (o *TemplateAssetCreate) SetUsage(v string) {
	o.Usage = v
}

// GetFile returns the File field value
func (o *TemplateAssetCreate) GetFile() TemplateAssetFile {
	if o == nil {
		var ret TemplateAssetFile
		return ret
	}

	return o.File
}

// GetFileOk returns a tuple with the File field value
// and a boolean to check if the value has been set.
func (o *TemplateAssetCreate) GetFileOk() (*TemplateAssetFile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.File, true
}

// SetFile sets field value
func (o *TemplateAssetCreate) SetFile(v TemplateAssetFile) {
	o.File = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *TemplateAssetCreate) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateAssetCreate) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TemplateAssetCreate) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *TemplateAssetCreate) SetTags(v []string) {
	o.Tags = v
}

func (o TemplateAssetCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateAssetCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["templateId"] = o.TemplateId
	toSerialize["usage"] = o.Usage
	toSerialize["file"] = o.File
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TemplateAssetCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"templateId",
		"usage",
		"file",
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

	varTemplateAssetCreate := _TemplateAssetCreate{}

	err = json.Unmarshal(data, &varTemplateAssetCreate)

	if err != nil {
		return err
	}

	*o = TemplateAssetCreate(varTemplateAssetCreate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "templateId")
		delete(additionalProperties, "usage")
		delete(additionalProperties, "file")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTemplateAssetCreate struct {
	value *TemplateAssetCreate
	isSet bool
}

func (v NullableTemplateAssetCreate) Get() *TemplateAssetCreate {
	return v.value
}

func (v *NullableTemplateAssetCreate) Set(val *TemplateAssetCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateAssetCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateAssetCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateAssetCreate(val *TemplateAssetCreate) *NullableTemplateAssetCreate {
	return &NullableTemplateAssetCreate{value: val, isSet: true}
}

func (v NullableTemplateAssetCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateAssetCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


