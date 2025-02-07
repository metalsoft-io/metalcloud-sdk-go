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

// checks if the ServerDefaultCredentialsPaginatedList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerDefaultCredentialsPaginatedList{}

// ServerDefaultCredentialsPaginatedList struct for ServerDefaultCredentialsPaginatedList
type ServerDefaultCredentialsPaginatedList struct {
	Data []ServerDefaultCredentials `json:"data"`
	// Metadata about the pagination of the response
	Meta PaginatedResponseMeta `json:"meta"`
	// Links to navigate through the paginated results
	Links PaginatedResponseLinks `json:"links"`
	AdditionalProperties map[string]interface{}
}

type _ServerDefaultCredentialsPaginatedList ServerDefaultCredentialsPaginatedList

// NewServerDefaultCredentialsPaginatedList instantiates a new ServerDefaultCredentialsPaginatedList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerDefaultCredentialsPaginatedList(data []ServerDefaultCredentials, meta PaginatedResponseMeta, links PaginatedResponseLinks) *ServerDefaultCredentialsPaginatedList {
	this := ServerDefaultCredentialsPaginatedList{}
	this.Data = data
	this.Meta = meta
	this.Links = links
	return &this
}

// NewServerDefaultCredentialsPaginatedListWithDefaults instantiates a new ServerDefaultCredentialsPaginatedList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerDefaultCredentialsPaginatedListWithDefaults() *ServerDefaultCredentialsPaginatedList {
	this := ServerDefaultCredentialsPaginatedList{}
	return &this
}

// GetData returns the Data field value
func (o *ServerDefaultCredentialsPaginatedList) GetData() []ServerDefaultCredentials {
	if o == nil {
		var ret []ServerDefaultCredentials
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ServerDefaultCredentialsPaginatedList) GetDataOk() ([]ServerDefaultCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ServerDefaultCredentialsPaginatedList) SetData(v []ServerDefaultCredentials) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *ServerDefaultCredentialsPaginatedList) GetMeta() PaginatedResponseMeta {
	if o == nil {
		var ret PaginatedResponseMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *ServerDefaultCredentialsPaginatedList) GetMetaOk() (*PaginatedResponseMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *ServerDefaultCredentialsPaginatedList) SetMeta(v PaginatedResponseMeta) {
	o.Meta = v
}

// GetLinks returns the Links field value
func (o *ServerDefaultCredentialsPaginatedList) GetLinks() PaginatedResponseLinks {
	if o == nil {
		var ret PaginatedResponseLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ServerDefaultCredentialsPaginatedList) GetLinksOk() (*PaginatedResponseLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ServerDefaultCredentialsPaginatedList) SetLinks(v PaginatedResponseLinks) {
	o.Links = v
}

func (o ServerDefaultCredentialsPaginatedList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerDefaultCredentialsPaginatedList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["meta"] = o.Meta
	toSerialize["links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServerDefaultCredentialsPaginatedList) UnmarshalJSON(data []byte) (err error) {
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

	varServerDefaultCredentialsPaginatedList := _ServerDefaultCredentialsPaginatedList{}

	err = json.Unmarshal(data, &varServerDefaultCredentialsPaginatedList)

	if err != nil {
		return err
	}

	*o = ServerDefaultCredentialsPaginatedList(varServerDefaultCredentialsPaginatedList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "meta")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServerDefaultCredentialsPaginatedList struct {
	value *ServerDefaultCredentialsPaginatedList
	isSet bool
}

func (v NullableServerDefaultCredentialsPaginatedList) Get() *ServerDefaultCredentialsPaginatedList {
	return v.value
}

func (v *NullableServerDefaultCredentialsPaginatedList) Set(val *ServerDefaultCredentialsPaginatedList) {
	v.value = val
	v.isSet = true
}

func (v NullableServerDefaultCredentialsPaginatedList) IsSet() bool {
	return v.isSet
}

func (v *NullableServerDefaultCredentialsPaginatedList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerDefaultCredentialsPaginatedList(val *ServerDefaultCredentialsPaginatedList) *NullableServerDefaultCredentialsPaginatedList {
	return &NullableServerDefaultCredentialsPaginatedList{value: val, isSet: true}
}

func (v NullableServerDefaultCredentialsPaginatedList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerDefaultCredentialsPaginatedList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


