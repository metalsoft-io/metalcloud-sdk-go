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

// checks if the PaginatedResponseMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedResponseMeta{}

// PaginatedResponseMeta Metadata about the pagination of the response
type PaginatedResponseMeta struct {
	// The number of items per page to return in the response
	ItemsPerPage int32 `json:"itemsPerPage"`
	// Total number of items matching the query
	TotalItems *int32 `json:"totalItems,omitempty"`
	// Current page number (1-based indexing)
	CurrentPage *int32 `json:"currentPage,omitempty"`
	// Total number of pages based on the total number of items and the number of items per page
	TotalPages *int32 `json:"totalPages,omitempty"`
	// Array of [field, direction] pairs for sorting. Each pair must contain exactly 2 strings.
	SortBy [][]string `json:"sortBy,omitempty"`
	SearchBy []string `json:"searchBy,omitempty"`
	// Text to search for in searchable fields
	Search *string `json:"search,omitempty"`
	// List of fields to include in the response
	Select []string `json:"select,omitempty"`
	// Key-value pairs of filters applied to the query
	Filter map[string]interface{} `json:"filter,omitempty"`
	// Cursor to navigate to the next page
	Cursor *string `json:"cursor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedResponseMeta PaginatedResponseMeta

// NewPaginatedResponseMeta instantiates a new PaginatedResponseMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedResponseMeta(itemsPerPage int32) *PaginatedResponseMeta {
	this := PaginatedResponseMeta{}
	this.ItemsPerPage = itemsPerPage
	return &this
}

// NewPaginatedResponseMetaWithDefaults instantiates a new PaginatedResponseMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedResponseMetaWithDefaults() *PaginatedResponseMeta {
	this := PaginatedResponseMeta{}
	return &this
}

// GetItemsPerPage returns the ItemsPerPage field value
func (o *PaginatedResponseMeta) GetItemsPerPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ItemsPerPage
}

// GetItemsPerPageOk returns a tuple with the ItemsPerPage field value
// and a boolean to check if the value has been set.
func (o *PaginatedResponseMeta) GetItemsPerPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemsPerPage, true
}

// SetItemsPerPage sets field value
func (o *PaginatedResponseMeta) SetItemsPerPage(v int32) {
	o.ItemsPerPage = v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *PaginatedResponseMeta) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedResponseMeta) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *PaginatedResponseMeta) HasTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *PaginatedResponseMeta) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetCurrentPage returns the CurrentPage field value if set, zero value otherwise.
func (o *PaginatedResponseMeta) GetCurrentPage() int32 {
	if o == nil || IsNil(o.CurrentPage) {
		var ret int32
		return ret
	}
	return *o.CurrentPage
}

// GetCurrentPageOk returns a tuple with the CurrentPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedResponseMeta) GetCurrentPageOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrentPage) {
		return nil, false
	}
	return o.CurrentPage, true
}

// HasCurrentPage returns a boolean if a field has been set.
func (o *PaginatedResponseMeta) HasCurrentPage() bool {
	if o != nil && !IsNil(o.CurrentPage) {
		return true
	}

	return false
}

// SetCurrentPage gets a reference to the given int32 and assigns it to the CurrentPage field.
func (o *PaginatedResponseMeta) SetCurrentPage(v int32) {
	o.CurrentPage = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *PaginatedResponseMeta) GetTotalPages() int32 {
	if o == nil || IsNil(o.TotalPages) {
		var ret int32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedResponseMeta) GetTotalPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPages) {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *PaginatedResponseMeta) HasTotalPages() bool {
	if o != nil && !IsNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given int32 and assigns it to the TotalPages field.
func (o *PaginatedResponseMeta) SetTotalPages(v int32) {
	o.TotalPages = &v
}

// GetSortBy returns the SortBy field value if set, zero value otherwise.
func (o *PaginatedResponseMeta) GetSortBy() [][]string {
	if o == nil || IsNil(o.SortBy) {
		var ret [][]string
		return ret
	}
	return o.SortBy
}

// GetSortByOk returns a tuple with the SortBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedResponseMeta) GetSortByOk() ([][]string, bool) {
	if o == nil || IsNil(o.SortBy) {
		return nil, false
	}
	return o.SortBy, true
}

// HasSortBy returns a boolean if a field has been set.
func (o *PaginatedResponseMeta) HasSortBy() bool {
	if o != nil && !IsNil(o.SortBy) {
		return true
	}

	return false
}

// SetSortBy gets a reference to the given [][]string and assigns it to the SortBy field.
func (o *PaginatedResponseMeta) SetSortBy(v [][]string) {
	o.SortBy = v
}

// GetSearchBy returns the SearchBy field value if set, zero value otherwise.
func (o *PaginatedResponseMeta) GetSearchBy() []string {
	if o == nil || IsNil(o.SearchBy) {
		var ret []string
		return ret
	}
	return o.SearchBy
}

// GetSearchByOk returns a tuple with the SearchBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedResponseMeta) GetSearchByOk() ([]string, bool) {
	if o == nil || IsNil(o.SearchBy) {
		return nil, false
	}
	return o.SearchBy, true
}

// HasSearchBy returns a boolean if a field has been set.
func (o *PaginatedResponseMeta) HasSearchBy() bool {
	if o != nil && !IsNil(o.SearchBy) {
		return true
	}

	return false
}

// SetSearchBy gets a reference to the given []string and assigns it to the SearchBy field.
func (o *PaginatedResponseMeta) SetSearchBy(v []string) {
	o.SearchBy = v
}

// GetSearch returns the Search field value if set, zero value otherwise.
func (o *PaginatedResponseMeta) GetSearch() string {
	if o == nil || IsNil(o.Search) {
		var ret string
		return ret
	}
	return *o.Search
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedResponseMeta) GetSearchOk() (*string, bool) {
	if o == nil || IsNil(o.Search) {
		return nil, false
	}
	return o.Search, true
}

// HasSearch returns a boolean if a field has been set.
func (o *PaginatedResponseMeta) HasSearch() bool {
	if o != nil && !IsNil(o.Search) {
		return true
	}

	return false
}

// SetSearch gets a reference to the given string and assigns it to the Search field.
func (o *PaginatedResponseMeta) SetSearch(v string) {
	o.Search = &v
}

// GetSelect returns the Select field value if set, zero value otherwise.
func (o *PaginatedResponseMeta) GetSelect() []string {
	if o == nil || IsNil(o.Select) {
		var ret []string
		return ret
	}
	return o.Select
}

// GetSelectOk returns a tuple with the Select field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedResponseMeta) GetSelectOk() ([]string, bool) {
	if o == nil || IsNil(o.Select) {
		return nil, false
	}
	return o.Select, true
}

// HasSelect returns a boolean if a field has been set.
func (o *PaginatedResponseMeta) HasSelect() bool {
	if o != nil && !IsNil(o.Select) {
		return true
	}

	return false
}

// SetSelect gets a reference to the given []string and assigns it to the Select field.
func (o *PaginatedResponseMeta) SetSelect(v []string) {
	o.Select = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *PaginatedResponseMeta) GetFilter() map[string]interface{} {
	if o == nil || IsNil(o.Filter) {
		var ret map[string]interface{}
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedResponseMeta) GetFilterOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Filter) {
		return map[string]interface{}{}, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *PaginatedResponseMeta) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given map[string]interface{} and assigns it to the Filter field.
func (o *PaginatedResponseMeta) SetFilter(v map[string]interface{}) {
	o.Filter = v
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *PaginatedResponseMeta) GetCursor() string {
	if o == nil || IsNil(o.Cursor) {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedResponseMeta) GetCursorOk() (*string, bool) {
	if o == nil || IsNil(o.Cursor) {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *PaginatedResponseMeta) HasCursor() bool {
	if o != nil && !IsNil(o.Cursor) {
		return true
	}

	return false
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *PaginatedResponseMeta) SetCursor(v string) {
	o.Cursor = &v
}

func (o PaginatedResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedResponseMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["itemsPerPage"] = o.ItemsPerPage
	if !IsNil(o.TotalItems) {
		toSerialize["totalItems"] = o.TotalItems
	}
	if !IsNil(o.CurrentPage) {
		toSerialize["currentPage"] = o.CurrentPage
	}
	if !IsNil(o.TotalPages) {
		toSerialize["totalPages"] = o.TotalPages
	}
	if !IsNil(o.SortBy) {
		toSerialize["sortBy"] = o.SortBy
	}
	if !IsNil(o.SearchBy) {
		toSerialize["searchBy"] = o.SearchBy
	}
	if !IsNil(o.Search) {
		toSerialize["search"] = o.Search
	}
	if !IsNil(o.Select) {
		toSerialize["select"] = o.Select
	}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.Cursor) {
		toSerialize["cursor"] = o.Cursor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaginatedResponseMeta) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"itemsPerPage",
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

	varPaginatedResponseMeta := _PaginatedResponseMeta{}

	err = json.Unmarshal(data, &varPaginatedResponseMeta)

	if err != nil {
		return err
	}

	*o = PaginatedResponseMeta(varPaginatedResponseMeta)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "itemsPerPage")
		delete(additionalProperties, "totalItems")
		delete(additionalProperties, "currentPage")
		delete(additionalProperties, "totalPages")
		delete(additionalProperties, "sortBy")
		delete(additionalProperties, "searchBy")
		delete(additionalProperties, "search")
		delete(additionalProperties, "select")
		delete(additionalProperties, "filter")
		delete(additionalProperties, "cursor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginatedResponseMeta struct {
	value *PaginatedResponseMeta
	isSet bool
}

func (v NullablePaginatedResponseMeta) Get() *PaginatedResponseMeta {
	return v.value
}

func (v *NullablePaginatedResponseMeta) Set(val *PaginatedResponseMeta) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedResponseMeta) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedResponseMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedResponseMeta(val *PaginatedResponseMeta) *NullablePaginatedResponseMeta {
	return &NullablePaginatedResponseMeta{value: val, isSet: true}
}

func (v NullablePaginatedResponseMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedResponseMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


