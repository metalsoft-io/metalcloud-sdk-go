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
)

// checks if the GenericGUISettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenericGUISettings{}

// GenericGUISettings struct for GenericGUISettings
type GenericGUISettings struct {
	// Row index of the object.
	RowIndex *float32 `json:"rowIndex,omitempty"`
	// Column index of the object.
	ColumnIndex *float32 `json:"columnIndex,omitempty"`
	// Whether to show the object children in the GUI.
	ShowWidgetChildren *bool `json:"showWidgetChildren,omitempty"`
	// Random instance ID.
	RandomInstanceID *string `json:"randomInstanceID,omitempty"`
	// User agent.
	UserAgent *string `json:"userAgent,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GenericGUISettings GenericGUISettings

// NewGenericGUISettings instantiates a new GenericGUISettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericGUISettings() *GenericGUISettings {
	this := GenericGUISettings{}
	return &this
}

// NewGenericGUISettingsWithDefaults instantiates a new GenericGUISettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericGUISettingsWithDefaults() *GenericGUISettings {
	this := GenericGUISettings{}
	return &this
}

// GetRowIndex returns the RowIndex field value if set, zero value otherwise.
func (o *GenericGUISettings) GetRowIndex() float32 {
	if o == nil || IsNil(o.RowIndex) {
		var ret float32
		return ret
	}
	return *o.RowIndex
}

// GetRowIndexOk returns a tuple with the RowIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericGUISettings) GetRowIndexOk() (*float32, bool) {
	if o == nil || IsNil(o.RowIndex) {
		return nil, false
	}
	return o.RowIndex, true
}

// HasRowIndex returns a boolean if a field has been set.
func (o *GenericGUISettings) HasRowIndex() bool {
	if o != nil && !IsNil(o.RowIndex) {
		return true
	}

	return false
}

// SetRowIndex gets a reference to the given float32 and assigns it to the RowIndex field.
func (o *GenericGUISettings) SetRowIndex(v float32) {
	o.RowIndex = &v
}

// GetColumnIndex returns the ColumnIndex field value if set, zero value otherwise.
func (o *GenericGUISettings) GetColumnIndex() float32 {
	if o == nil || IsNil(o.ColumnIndex) {
		var ret float32
		return ret
	}
	return *o.ColumnIndex
}

// GetColumnIndexOk returns a tuple with the ColumnIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericGUISettings) GetColumnIndexOk() (*float32, bool) {
	if o == nil || IsNil(o.ColumnIndex) {
		return nil, false
	}
	return o.ColumnIndex, true
}

// HasColumnIndex returns a boolean if a field has been set.
func (o *GenericGUISettings) HasColumnIndex() bool {
	if o != nil && !IsNil(o.ColumnIndex) {
		return true
	}

	return false
}

// SetColumnIndex gets a reference to the given float32 and assigns it to the ColumnIndex field.
func (o *GenericGUISettings) SetColumnIndex(v float32) {
	o.ColumnIndex = &v
}

// GetShowWidgetChildren returns the ShowWidgetChildren field value if set, zero value otherwise.
func (o *GenericGUISettings) GetShowWidgetChildren() bool {
	if o == nil || IsNil(o.ShowWidgetChildren) {
		var ret bool
		return ret
	}
	return *o.ShowWidgetChildren
}

// GetShowWidgetChildrenOk returns a tuple with the ShowWidgetChildren field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericGUISettings) GetShowWidgetChildrenOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowWidgetChildren) {
		return nil, false
	}
	return o.ShowWidgetChildren, true
}

// HasShowWidgetChildren returns a boolean if a field has been set.
func (o *GenericGUISettings) HasShowWidgetChildren() bool {
	if o != nil && !IsNil(o.ShowWidgetChildren) {
		return true
	}

	return false
}

// SetShowWidgetChildren gets a reference to the given bool and assigns it to the ShowWidgetChildren field.
func (o *GenericGUISettings) SetShowWidgetChildren(v bool) {
	o.ShowWidgetChildren = &v
}

// GetRandomInstanceID returns the RandomInstanceID field value if set, zero value otherwise.
func (o *GenericGUISettings) GetRandomInstanceID() string {
	if o == nil || IsNil(o.RandomInstanceID) {
		var ret string
		return ret
	}
	return *o.RandomInstanceID
}

// GetRandomInstanceIDOk returns a tuple with the RandomInstanceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericGUISettings) GetRandomInstanceIDOk() (*string, bool) {
	if o == nil || IsNil(o.RandomInstanceID) {
		return nil, false
	}
	return o.RandomInstanceID, true
}

// HasRandomInstanceID returns a boolean if a field has been set.
func (o *GenericGUISettings) HasRandomInstanceID() bool {
	if o != nil && !IsNil(o.RandomInstanceID) {
		return true
	}

	return false
}

// SetRandomInstanceID gets a reference to the given string and assigns it to the RandomInstanceID field.
func (o *GenericGUISettings) SetRandomInstanceID(v string) {
	o.RandomInstanceID = &v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise.
func (o *GenericGUISettings) GetUserAgent() string {
	if o == nil || IsNil(o.UserAgent) {
		var ret string
		return ret
	}
	return *o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericGUISettings) GetUserAgentOk() (*string, bool) {
	if o == nil || IsNil(o.UserAgent) {
		return nil, false
	}
	return o.UserAgent, true
}

// HasUserAgent returns a boolean if a field has been set.
func (o *GenericGUISettings) HasUserAgent() bool {
	if o != nil && !IsNil(o.UserAgent) {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given string and assigns it to the UserAgent field.
func (o *GenericGUISettings) SetUserAgent(v string) {
	o.UserAgent = &v
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
	if !IsNil(o.RowIndex) {
		toSerialize["rowIndex"] = o.RowIndex
	}
	if !IsNil(o.ColumnIndex) {
		toSerialize["columnIndex"] = o.ColumnIndex
	}
	if !IsNil(o.ShowWidgetChildren) {
		toSerialize["showWidgetChildren"] = o.ShowWidgetChildren
	}
	if !IsNil(o.RandomInstanceID) {
		toSerialize["randomInstanceID"] = o.RandomInstanceID
	}
	if !IsNil(o.UserAgent) {
		toSerialize["userAgent"] = o.UserAgent
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GenericGUISettings) UnmarshalJSON(data []byte) (err error) {
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


