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

// checks if the ServerInstanceGroupPaginatedList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerInstanceGroupPaginatedList{}

// ServerInstanceGroupPaginatedList struct for ServerInstanceGroupPaginatedList
type ServerInstanceGroupPaginatedList struct {
	Data []ServerInstanceGroup `json:"data"`
	// Metadata about the pagination of the response
	Meta PaginatedResponseMeta `json:"meta"`
	// Links to navigate through the paginated results
	Links PaginatedResponseLinks `json:"links"`
	AdditionalProperties map[string]interface{}
}

type _ServerInstanceGroupPaginatedList ServerInstanceGroupPaginatedList

// NewServerInstanceGroupPaginatedList instantiates a new ServerInstanceGroupPaginatedList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerInstanceGroupPaginatedList(data []ServerInstanceGroup, meta PaginatedResponseMeta, links PaginatedResponseLinks) *ServerInstanceGroupPaginatedList {
	this := ServerInstanceGroupPaginatedList{}
	this.Data = data
	this.Meta = meta
	this.Links = links
	return &this
}

// NewServerInstanceGroupPaginatedListWithDefaults instantiates a new ServerInstanceGroupPaginatedList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerInstanceGroupPaginatedListWithDefaults() *ServerInstanceGroupPaginatedList {
	this := ServerInstanceGroupPaginatedList{}
	return &this
}

// GetData returns the Data field value
func (o *ServerInstanceGroupPaginatedList) GetData() []ServerInstanceGroup {
	if o == nil {
		var ret []ServerInstanceGroup
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupPaginatedList) GetDataOk() ([]ServerInstanceGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ServerInstanceGroupPaginatedList) SetData(v []ServerInstanceGroup) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *ServerInstanceGroupPaginatedList) GetMeta() PaginatedResponseMeta {
	if o == nil {
		var ret PaginatedResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupPaginatedList) GetMetaOk() (*PaginatedResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *ServerInstanceGroupPaginatedList) SetMeta(v PaginatedResponseMeta) {
	o.Meta = v
}

// GetLinks returns the Links field value
func (o *ServerInstanceGroupPaginatedList) GetLinks() PaginatedResponseLinks {
	if o == nil {
		var ret PaginatedResponseLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ServerInstanceGroupPaginatedList) GetLinksOk() (*PaginatedResponseLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ServerInstanceGroupPaginatedList) SetLinks(v PaginatedResponseLinks) {
	o.Links = v
}

func (o ServerInstanceGroupPaginatedList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerInstanceGroupPaginatedList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["meta"] = o.Meta
	toSerialize["links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServerInstanceGroupPaginatedList) UnmarshalJSON(data []byte) (err error) {
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

	varServerInstanceGroupPaginatedList := _ServerInstanceGroupPaginatedList{}

	err = json.Unmarshal(data, &varServerInstanceGroupPaginatedList)

	if err != nil {
		return err
	}

	*o = ServerInstanceGroupPaginatedList(varServerInstanceGroupPaginatedList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "meta")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServerInstanceGroupPaginatedList struct {
	value *ServerInstanceGroupPaginatedList
	isSet bool
}

func (v NullableServerInstanceGroupPaginatedList) Get() *ServerInstanceGroupPaginatedList {
	return v.value
}

func (v *NullableServerInstanceGroupPaginatedList) Set(val *ServerInstanceGroupPaginatedList) {
	v.value = val
	v.isSet = true
}

func (v NullableServerInstanceGroupPaginatedList) IsSet() bool {
	return v.isSet
}

func (v *NullableServerInstanceGroupPaginatedList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerInstanceGroupPaginatedList(val *ServerInstanceGroupPaginatedList) *NullableServerInstanceGroupPaginatedList {
	return &NullableServerInstanceGroupPaginatedList{value: val, isSet: true}
}

func (v NullableServerInstanceGroupPaginatedList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerInstanceGroupPaginatedList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


