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

// checks if the VMInstanceGroupPaginatedList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VMInstanceGroupPaginatedList{}

// VMInstanceGroupPaginatedList struct for VMInstanceGroupPaginatedList
type VMInstanceGroupPaginatedList struct {
	Data []VMInstanceGroup `json:"data"`
	// Metadata about the pagination of the response
	Meta PaginatedResponseMeta `json:"meta"`
	// Links to navigate through the paginated results
	Links PaginatedResponseLinks `json:"links"`
	AdditionalProperties map[string]interface{}
}

type _VMInstanceGroupPaginatedList VMInstanceGroupPaginatedList

// NewVMInstanceGroupPaginatedList instantiates a new VMInstanceGroupPaginatedList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVMInstanceGroupPaginatedList(data []VMInstanceGroup, meta PaginatedResponseMeta, links PaginatedResponseLinks) *VMInstanceGroupPaginatedList {
	this := VMInstanceGroupPaginatedList{}
	this.Data = data
	this.Meta = meta
	this.Links = links
	return &this
}

// NewVMInstanceGroupPaginatedListWithDefaults instantiates a new VMInstanceGroupPaginatedList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVMInstanceGroupPaginatedListWithDefaults() *VMInstanceGroupPaginatedList {
	this := VMInstanceGroupPaginatedList{}
	return &this
}

// GetData returns the Data field value
func (o *VMInstanceGroupPaginatedList) GetData() []VMInstanceGroup {
	if o == nil {
		var ret []VMInstanceGroup
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *VMInstanceGroupPaginatedList) GetDataOk() ([]VMInstanceGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *VMInstanceGroupPaginatedList) SetData(v []VMInstanceGroup) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *VMInstanceGroupPaginatedList) GetMeta() PaginatedResponseMeta {
	if o == nil {
		var ret PaginatedResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *VMInstanceGroupPaginatedList) GetMetaOk() (*PaginatedResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *VMInstanceGroupPaginatedList) SetMeta(v PaginatedResponseMeta) {
	o.Meta = v
}

// GetLinks returns the Links field value
func (o *VMInstanceGroupPaginatedList) GetLinks() PaginatedResponseLinks {
	if o == nil {
		var ret PaginatedResponseLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *VMInstanceGroupPaginatedList) GetLinksOk() (*PaginatedResponseLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *VMInstanceGroupPaginatedList) SetLinks(v PaginatedResponseLinks) {
	o.Links = v
}

func (o VMInstanceGroupPaginatedList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VMInstanceGroupPaginatedList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["meta"] = o.Meta
	toSerialize["links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VMInstanceGroupPaginatedList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"meta",
		"links",
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

	varVMInstanceGroupPaginatedList := _VMInstanceGroupPaginatedList{}

	err = json.Unmarshal(data, &varVMInstanceGroupPaginatedList)

	if err != nil {
		return err
	}

	*o = VMInstanceGroupPaginatedList(varVMInstanceGroupPaginatedList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "meta")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVMInstanceGroupPaginatedList struct {
	value *VMInstanceGroupPaginatedList
	isSet bool
}

func (v NullableVMInstanceGroupPaginatedList) Get() *VMInstanceGroupPaginatedList {
	return v.value
}

func (v *NullableVMInstanceGroupPaginatedList) Set(val *VMInstanceGroupPaginatedList) {
	v.value = val
	v.isSet = true
}

func (v NullableVMInstanceGroupPaginatedList) IsSet() bool {
	return v.isSet
}

func (v *NullableVMInstanceGroupPaginatedList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVMInstanceGroupPaginatedList(val *VMInstanceGroupPaginatedList) *NullableVMInstanceGroupPaginatedList {
	return &NullableVMInstanceGroupPaginatedList{value: val, isSet: true}
}

func (v NullableVMInstanceGroupPaginatedList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVMInstanceGroupPaginatedList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


