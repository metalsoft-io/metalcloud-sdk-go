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

// checks if the GenericGUISettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenericGUISettings{}

// GenericGUISettings struct for GenericGUISettings
type GenericGUISettings struct {
	// Row index of the object.
	RowIndex float32 `json:"rowIndex"`
	// Column index of the object.
	ColumnIndex float32 `json:"columnIndex"`
	// Whether to show the object children in the GUI.
	ShowWidgetChildren bool `json:"showWidgetChildren"`
	// Random instance ID.
	RandomInstanceID string `json:"randomInstanceID"`
	// User agent.
	UserAgent string `json:"userAgent"`
	// tags.
	Tags []string `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GenericGUISettings GenericGUISettings

// NewGenericGUISettings instantiates a new GenericGUISettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericGUISettings(rowIndex float32, columnIndex float32, showWidgetChildren bool, randomInstanceID string, userAgent string) *GenericGUISettings {
	this := GenericGUISettings{}
	this.RowIndex = rowIndex
	this.ColumnIndex = columnIndex
	this.ShowWidgetChildren = showWidgetChildren
	this.RandomInstanceID = randomInstanceID
	this.UserAgent = userAgent
	return &this
}

// NewGenericGUISettingsWithDefaults instantiates a new GenericGUISettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericGUISettingsWithDefaults() *GenericGUISettings {
	this := GenericGUISettings{}
	return &this
}

// GetRowIndex returns the RowIndex field value
func (o *GenericGUISettings) GetRowIndex() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RowIndex
}

// GetRowIndexOk returns a tuple with the RowIndex field value
// and a boolean to check if the value has been set.
func (o *GenericGUISettings) GetRowIndexOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RowIndex, true
}

// SetRowIndex sets field value
func (o *GenericGUISettings) SetRowIndex(v float32) {
	o.RowIndex = v
}

// GetColumnIndex returns the ColumnIndex field value
func (o *GenericGUISettings) GetColumnIndex() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ColumnIndex
}

// GetColumnIndexOk returns a tuple with the ColumnIndex field value
// and a boolean to check if the value has been set.
func (o *GenericGUISettings) GetColumnIndexOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ColumnIndex, true
}

// SetColumnIndex sets field value
func (o *GenericGUISettings) SetColumnIndex(v float32) {
	o.ColumnIndex = v
}

// GetShowWidgetChildren returns the ShowWidgetChildren field value
func (o *GenericGUISettings) GetShowWidgetChildren() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShowWidgetChildren
}

// GetShowWidgetChildrenOk returns a tuple with the ShowWidgetChildren field value
// and a boolean to check if the value has been set.
func (o *GenericGUISettings) GetShowWidgetChildrenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShowWidgetChildren, true
}

// SetShowWidgetChildren sets field value
func (o *GenericGUISettings) SetShowWidgetChildren(v bool) {
	o.ShowWidgetChildren = v
}

// GetRandomInstanceID returns the RandomInstanceID field value
func (o *GenericGUISettings) GetRandomInstanceID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RandomInstanceID
}

// GetRandomInstanceIDOk returns a tuple with the RandomInstanceID field value
// and a boolean to check if the value has been set.
func (o *GenericGUISettings) GetRandomInstanceIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RandomInstanceID, true
}

// SetRandomInstanceID sets field value
func (o *GenericGUISettings) SetRandomInstanceID(v string) {
	o.RandomInstanceID = v
}

// GetUserAgent returns the UserAgent field value
func (o *GenericGUISettings) GetUserAgent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value
// and a boolean to check if the value has been set.
func (o *GenericGUISettings) GetUserAgentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserAgent, true
}

// SetUserAgent sets field value
func (o *GenericGUISettings) SetUserAgent(v string) {
	o.UserAgent = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GenericGUISettings) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericGUISettings) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GenericGUISettings) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GenericGUISettings) SetTags(v []string) {
	o.Tags = v
}

func (o GenericGUISettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenericGUISettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rowIndex"] = o.RowIndex
	toSerialize["columnIndex"] = o.ColumnIndex
	toSerialize["showWidgetChildren"] = o.ShowWidgetChildren
	toSerialize["randomInstanceID"] = o.RandomInstanceID
	toSerialize["userAgent"] = o.UserAgent
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GenericGUISettings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"rowIndex",
		"columnIndex",
		"showWidgetChildren",
		"randomInstanceID",
		"userAgent",
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

	varGenericGUISettings := _GenericGUISettings{}

	err = json.Unmarshal(data, &varGenericGUISettings)

	if err != nil {
		return err
	}

	*o = GenericGUISettings(varGenericGUISettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rowIndex")
		delete(additionalProperties, "columnIndex")
		delete(additionalProperties, "showWidgetChildren")
		delete(additionalProperties, "randomInstanceID")
		delete(additionalProperties, "userAgent")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGenericGUISettings struct {
	value *GenericGUISettings
	isSet bool
}

func (v NullableGenericGUISettings) Get() *GenericGUISettings {
	return v.value
}

func (v *NullableGenericGUISettings) Set(val *GenericGUISettings) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericGUISettings) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericGUISettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericGUISettings(val *GenericGUISettings) *NullableGenericGUISettings {
	return &NullableGenericGUISettings{value: val, isSet: true}
}

func (v NullableGenericGUISettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericGUISettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


