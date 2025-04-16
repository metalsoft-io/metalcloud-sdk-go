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

// checks if the FileSharePaginatedList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileSharePaginatedList{}

// FileSharePaginatedList struct for FileSharePaginatedList
type FileSharePaginatedList struct {
	Data []FileShare `json:"data"`
	// Metadata about the pagination of the response
	Meta PaginatedResponseMeta `json:"meta"`
	// Links to navigate through the paginated results
	Links PaginatedResponseLinks `json:"links"`
	AdditionalProperties map[string]interface{}
}

type _FileSharePaginatedList FileSharePaginatedList

// NewFileSharePaginatedList instantiates a new FileSharePaginatedList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileSharePaginatedList(data []FileShare, meta PaginatedResponseMeta, links PaginatedResponseLinks) *FileSharePaginatedList {
	this := FileSharePaginatedList{}
	this.Data = data
	this.Meta = meta
	this.Links = links
	return &this
}

// NewFileSharePaginatedListWithDefaults instantiates a new FileSharePaginatedList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileSharePaginatedListWithDefaults() *FileSharePaginatedList {
	this := FileSharePaginatedList{}
	return &this
}

// GetData returns the Data field value
func (o *FileSharePaginatedList) GetData() []FileShare {
	if o == nil {
		var ret []FileShare
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *FileSharePaginatedList) GetDataOk() ([]FileShare, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *FileSharePaginatedList) SetData(v []FileShare) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *FileSharePaginatedList) GetMeta() PaginatedResponseMeta {
	if o == nil {
		var ret PaginatedResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *FileSharePaginatedList) GetMetaOk() (*PaginatedResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *FileSharePaginatedList) SetMeta(v PaginatedResponseMeta) {
	o.Meta = v
}

// GetLinks returns the Links field value
func (o *FileSharePaginatedList) GetLinks() PaginatedResponseLinks {
	if o == nil {
		var ret PaginatedResponseLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *FileSharePaginatedList) GetLinksOk() (*PaginatedResponseLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *FileSharePaginatedList) SetLinks(v PaginatedResponseLinks) {
	o.Links = v
}

func (o FileSharePaginatedList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileSharePaginatedList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["meta"] = o.Meta
	toSerialize["links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FileSharePaginatedList) UnmarshalJSON(data []byte) (err error) {
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

	varFileSharePaginatedList := _FileSharePaginatedList{}

	err = json.Unmarshal(data, &varFileSharePaginatedList)

	if err != nil {
		return err
	}

	*o = FileSharePaginatedList(varFileSharePaginatedList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "meta")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFileSharePaginatedList struct {
	value *FileSharePaginatedList
	isSet bool
}

func (v NullableFileSharePaginatedList) Get() *FileSharePaginatedList {
	return v.value
}

func (v *NullableFileSharePaginatedList) Set(val *FileSharePaginatedList) {
	v.value = val
	v.isSet = true
}

func (v NullableFileSharePaginatedList) IsSet() bool {
	return v.isSet
}

func (v *NullableFileSharePaginatedList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileSharePaginatedList(val *FileSharePaginatedList) *NullableFileSharePaginatedList {
	return &NullableFileSharePaginatedList{value: val, isSet: true}
}

func (v NullableFileSharePaginatedList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileSharePaginatedList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


