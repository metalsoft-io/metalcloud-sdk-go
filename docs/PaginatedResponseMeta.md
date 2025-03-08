# PaginatedResponseMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemsPerPage** | **int32** | The number of items per page to return in the response | 
**TotalItems** | **int32** | Total number of items matching the query | 
**CurrentPage** | **int32** | Current page number (1-based indexing) | 
**TotalPages** | **int32** | Total number of pages based on the total number of items and the number of items per page | 
**SortBy** | Pointer to **[][]string** | Array of [field, direction] pairs for sorting. Each pair must contain exactly 2 strings. | [optional] 
**SearchBy** | Pointer to **[]string** |  | [optional] 
**Search** | Pointer to **string** | Text to search for in searchable fields | [optional] 
**Select** | Pointer to **[]string** | List of fields to include in the response | [optional] 
**Filter** | Pointer to **map[string]interface{}** | Key-value pairs of filters applied to the query | [optional] 

## Methods

### NewPaginatedResponseMeta

`func NewPaginatedResponseMeta(itemsPerPage int32, totalItems int32, currentPage int32, totalPages int32, ) *PaginatedResponseMeta`

NewPaginatedResponseMeta instantiates a new PaginatedResponseMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedResponseMetaWithDefaults

`func NewPaginatedResponseMetaWithDefaults() *PaginatedResponseMeta`

NewPaginatedResponseMetaWithDefaults instantiates a new PaginatedResponseMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemsPerPage

`func (o *PaginatedResponseMeta) GetItemsPerPage() int32`

GetItemsPerPage returns the ItemsPerPage field if non-nil, zero value otherwise.

### GetItemsPerPageOk

`func (o *PaginatedResponseMeta) GetItemsPerPageOk() (*int32, bool)`

GetItemsPerPageOk returns a tuple with the ItemsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPerPage

`func (o *PaginatedResponseMeta) SetItemsPerPage(v int32)`

SetItemsPerPage sets ItemsPerPage field to given value.


### GetTotalItems

`func (o *PaginatedResponseMeta) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *PaginatedResponseMeta) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *PaginatedResponseMeta) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.


### GetCurrentPage

`func (o *PaginatedResponseMeta) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *PaginatedResponseMeta) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *PaginatedResponseMeta) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.


### GetTotalPages

`func (o *PaginatedResponseMeta) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *PaginatedResponseMeta) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *PaginatedResponseMeta) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.


### GetSortBy

`func (o *PaginatedResponseMeta) GetSortBy() [][]string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *PaginatedResponseMeta) GetSortByOk() (*[][]string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *PaginatedResponseMeta) SetSortBy(v [][]string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *PaginatedResponseMeta) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### GetSearchBy

`func (o *PaginatedResponseMeta) GetSearchBy() []string`

GetSearchBy returns the SearchBy field if non-nil, zero value otherwise.

### GetSearchByOk

`func (o *PaginatedResponseMeta) GetSearchByOk() (*[]string, bool)`

GetSearchByOk returns a tuple with the SearchBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBy

`func (o *PaginatedResponseMeta) SetSearchBy(v []string)`

SetSearchBy sets SearchBy field to given value.

### HasSearchBy

`func (o *PaginatedResponseMeta) HasSearchBy() bool`

HasSearchBy returns a boolean if a field has been set.

### GetSearch

`func (o *PaginatedResponseMeta) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *PaginatedResponseMeta) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *PaginatedResponseMeta) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *PaginatedResponseMeta) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetSelect

`func (o *PaginatedResponseMeta) GetSelect() []string`

GetSelect returns the Select field if non-nil, zero value otherwise.

### GetSelectOk

`func (o *PaginatedResponseMeta) GetSelectOk() (*[]string, bool)`

GetSelectOk returns a tuple with the Select field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelect

`func (o *PaginatedResponseMeta) SetSelect(v []string)`

SetSelect sets Select field to given value.

### HasSelect

`func (o *PaginatedResponseMeta) HasSelect() bool`

HasSelect returns a boolean if a field has been set.

### GetFilter

`func (o *PaginatedResponseMeta) GetFilter() map[string]interface{}`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *PaginatedResponseMeta) GetFilterOk() (*map[string]interface{}, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *PaginatedResponseMeta) SetFilter(v map[string]interface{})`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *PaginatedResponseMeta) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


