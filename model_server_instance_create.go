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

// checks if the ServerInstanceCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerInstanceCreate{}

// ServerInstanceCreate struct for ServerInstanceCreate
type ServerInstanceCreate struct {
	// The server instance label.
	Label *string `json:"label,omitempty"`
	GroupId *int32 `json:"groupId,omitempty"`
	// The server type ID.
	ServerTypeId *int32 `json:"serverTypeId,omitempty"`
	// The template id of the operating system to deploy on the server. Can be null in which case no OS will be deployed but all operations will continue as normal. 
	TemplateId *int32 `json:"templateId,omitempty"`
	Tags []string `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServerInstanceCreate ServerInstanceCreate

// NewServerInstanceCreate instantiates a new ServerInstanceCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerInstanceCreate() *ServerInstanceCreate {
	this := ServerInstanceCreate{}
	return &this
}

// NewServerInstanceCreateWithDefaults instantiates a new ServerInstanceCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerInstanceCreateWithDefaults() *ServerInstanceCreate {
	this := ServerInstanceCreate{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ServerInstanceCreate) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceCreate) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ServerInstanceCreate) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ServerInstanceCreate) SetLabel(v string) {
	o.Label = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *ServerInstanceCreate) GetGroupId() int32 {
	if o == nil || IsNil(o.GroupId) {
		var ret int32
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceCreate) GetGroupIdOk() (*int32, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ServerInstanceCreate) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given int32 and assigns it to the GroupId field.
func (o *ServerInstanceCreate) SetGroupId(v int32) {
	o.GroupId = &v
}

// GetServerTypeId returns the ServerTypeId field value if set, zero value otherwise.
func (o *ServerInstanceCreate) GetServerTypeId() int32 {
	if o == nil || IsNil(o.ServerTypeId) {
		var ret int32
		return ret
	}
	return *o.ServerTypeId
}

// GetServerTypeIdOk returns a tuple with the ServerTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceCreate) GetServerTypeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ServerTypeId) {
		return nil, false
	}
	return o.ServerTypeId, true
}

// HasServerTypeId returns a boolean if a field has been set.
func (o *ServerInstanceCreate) HasServerTypeId() bool {
	if o != nil && !IsNil(o.ServerTypeId) {
		return true
	}

	return false
}

// SetServerTypeId gets a reference to the given int32 and assigns it to the ServerTypeId field.
func (o *ServerInstanceCreate) SetServerTypeId(v int32) {
	o.ServerTypeId = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *ServerInstanceCreate) GetTemplateId() int32 {
	if o == nil || IsNil(o.TemplateId) {
		var ret int32
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceCreate) GetTemplateIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *ServerInstanceCreate) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given int32 and assigns it to the TemplateId field.
func (o *ServerInstanceCreate) SetTemplateId(v int32) {
	o.TemplateId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ServerInstanceCreate) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceCreate) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ServerInstanceCreate) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ServerInstanceCreate) SetTags(v []string) {
	o.Tags = v
}

func (o ServerInstanceCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerInstanceCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !IsNil(o.ServerTypeId) {
		toSerialize["serverTypeId"] = o.ServerTypeId
	}
	if !IsNil(o.TemplateId) {
		toSerialize["templateId"] = o.TemplateId
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServerInstanceCreate) UnmarshalJSON(data []byte) (err error) {
	varServerInstanceCreate := _ServerInstanceCreate{}

	err = json.Unmarshal(data, &varServerInstanceCreate)

	if err != nil {
		return err
	}

	*o = ServerInstanceCreate(varServerInstanceCreate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "groupId")
		delete(additionalProperties, "serverTypeId")
		delete(additionalProperties, "templateId")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServerInstanceCreate struct {
	value *ServerInstanceCreate
	isSet bool
}

func (v NullableServerInstanceCreate) Get() *ServerInstanceCreate {
	return v.value
}

func (v *NullableServerInstanceCreate) Set(val *ServerInstanceCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableServerInstanceCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableServerInstanceCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerInstanceCreate(val *ServerInstanceCreate) *NullableServerInstanceCreate {
	return &NullableServerInstanceCreate{value: val, isSet: true}
}

func (v NullableServerInstanceCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerInstanceCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


