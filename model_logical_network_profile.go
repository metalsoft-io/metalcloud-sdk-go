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

// checks if the LogicalNetworkProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogicalNetworkProfile{}

// LogicalNetworkProfile struct for LogicalNetworkProfile
type LogicalNetworkProfile struct {
	// ID of the logical network profile
	Id float32 `json:"id"`
	// Revision number of the logical network profile
	Revision float32 `json:"revision"`
	// Label for the logical network profile
	Label *string `json:"label,omitempty"`
	// Name of the logical network profile
	Name *string `json:"name,omitempty"`
	// Description of the logical network profile
	Description *string `json:"description,omitempty"`
	// Annotations for the logical network profile
	Annotations map[string]interface{} `json:"annotations,omitempty"`
	// Type of the logical network profile
	LogicalNetworkType string `json:"logicalNetworkType"`
	AdditionalProperties map[string]interface{}
}

type _LogicalNetworkProfile LogicalNetworkProfile

// NewLogicalNetworkProfile instantiates a new LogicalNetworkProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogicalNetworkProfile(id float32, revision float32, logicalNetworkType string) *LogicalNetworkProfile {
	this := LogicalNetworkProfile{}
	this.Id = id
	this.Revision = revision
	this.LogicalNetworkType = logicalNetworkType
	return &this
}

// NewLogicalNetworkProfileWithDefaults instantiates a new LogicalNetworkProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogicalNetworkProfileWithDefaults() *LogicalNetworkProfile {
	this := LogicalNetworkProfile{}
	return &this
}

// GetId returns the Id field value
func (o *LogicalNetworkProfile) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LogicalNetworkProfile) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LogicalNetworkProfile) SetId(v float32) {
	o.Id = v
}

// GetRevision returns the Revision field value
func (o *LogicalNetworkProfile) GetRevision() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *LogicalNetworkProfile) GetRevisionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *LogicalNetworkProfile) SetRevision(v float32) {
	o.Revision = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *LogicalNetworkProfile) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogicalNetworkProfile) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *LogicalNetworkProfile) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *LogicalNetworkProfile) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LogicalNetworkProfile) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogicalNetworkProfile) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LogicalNetworkProfile) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LogicalNetworkProfile) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LogicalNetworkProfile) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogicalNetworkProfile) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LogicalNetworkProfile) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LogicalNetworkProfile) SetDescription(v string) {
	o.Description = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *LogicalNetworkProfile) GetAnnotations() map[string]interface{} {
	if o == nil || IsNil(o.Annotations) {
		var ret map[string]interface{}
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogicalNetworkProfile) GetAnnotationsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Annotations) {
		return map[string]interface{}{}, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *LogicalNetworkProfile) HasAnnotations() bool {
	if o != nil && !IsNil(o.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]interface{} and assigns it to the Annotations field.
func (o *LogicalNetworkProfile) SetAnnotations(v map[string]interface{}) {
	o.Annotations = v
}

// GetLogicalNetworkType returns the LogicalNetworkType field value
func (o *LogicalNetworkProfile) GetLogicalNetworkType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogicalNetworkType
}

// GetLogicalNetworkTypeOk returns a tuple with the LogicalNetworkType field value
// and a boolean to check if the value has been set.
func (o *LogicalNetworkProfile) GetLogicalNetworkTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogicalNetworkType, true
}

// SetLogicalNetworkType sets field value
func (o *LogicalNetworkProfile) SetLogicalNetworkType(v string) {
	o.LogicalNetworkType = v
}

func (o LogicalNetworkProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogicalNetworkProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["revision"] = o.Revision
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Annotations) {
		toSerialize["annotations"] = o.Annotations
	}
	toSerialize["logicalNetworkType"] = o.LogicalNetworkType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LogicalNetworkProfile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"revision",
		"logicalNetworkType",
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

	varLogicalNetworkProfile := _LogicalNetworkProfile{}

	err = json.Unmarshal(data, &varLogicalNetworkProfile)

	if err != nil {
		return err
	}

	*o = LogicalNetworkProfile(varLogicalNetworkProfile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "revision")
		delete(additionalProperties, "label")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "annotations")
		delete(additionalProperties, "logicalNetworkType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLogicalNetworkProfile struct {
	value *LogicalNetworkProfile
	isSet bool
}

func (v NullableLogicalNetworkProfile) Get() *LogicalNetworkProfile {
	return v.value
}

func (v *NullableLogicalNetworkProfile) Set(val *LogicalNetworkProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableLogicalNetworkProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableLogicalNetworkProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogicalNetworkProfile(val *LogicalNetworkProfile) *NullableLogicalNetworkProfile {
	return &NullableLogicalNetworkProfile{value: val, isSet: true}
}

func (v NullableLogicalNetworkProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogicalNetworkProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


