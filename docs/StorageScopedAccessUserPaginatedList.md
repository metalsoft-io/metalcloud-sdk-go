# StorageScopedAccessUserPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]StorageScopedAccessUser**](StorageScopedAccessUser.md) |  | 
**Meta** | [**PaginatedResponseMeta**](PaginatedResponseMeta.md) | Metadata about the pagination of the response | 
**Links** | Pointer to [**PaginatedResponseLinks**](PaginatedResponseLinks.md) | Links to navigate through the paginated results | [optional] 

## Methods

### NewStorageScopedAccessUserPaginatedList

`func NewStorageScopedAccessUserPaginatedList(data []StorageScopedAccessUser, meta PaginatedResponseMeta, ) *StorageScopedAccessUserPaginatedList`

NewStorageScopedAccessUserPaginatedList instantiates a new StorageScopedAccessUserPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageScopedAccessUserPaginatedListWithDefaults

`func NewStorageScopedAccessUserPaginatedListWithDefaults() *StorageScopedAccessUserPaginatedList`

NewStorageScopedAccessUserPaginatedListWithDefaults instantiates a new StorageScopedAccessUserPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *StorageScopedAccessUserPaginatedList) GetData() []StorageScopedAccessUser`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *StorageScopedAccessUserPaginatedList) GetDataOk() (*[]StorageScopedAccessUser, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *StorageScopedAccessUserPaginatedList) SetData(v []StorageScopedAccessUser)`

SetData sets Data field to given value.


### GetMeta

`func (o *StorageScopedAccessUserPaginatedList) GetMeta() PaginatedResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *StorageScopedAccessUserPaginatedList) GetMetaOk() (*PaginatedResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *StorageScopedAccessUserPaginatedList) SetMeta(v PaginatedResponseMeta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *StorageScopedAccessUserPaginatedList) GetLinks() PaginatedResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StorageScopedAccessUserPaginatedList) GetLinksOk() (*PaginatedResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StorageScopedAccessUserPaginatedList) SetLinks(v PaginatedResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StorageScopedAccessUserPaginatedList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


