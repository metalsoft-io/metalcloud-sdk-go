# ResourcePoolPaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ResourcePoolWithStats**](ResourcePoolWithStats.md) |  | 
**Meta** | [**PaginatedResponseMeta**](PaginatedResponseMeta.md) | Metadata about the pagination of the response | 
**Links** | Pointer to [**PaginatedResponseLinks**](PaginatedResponseLinks.md) | Links to navigate through the paginated results | [optional] 

## Methods

### NewResourcePoolPaginatedList

`func NewResourcePoolPaginatedList(data []ResourcePoolWithStats, meta PaginatedResponseMeta, ) *ResourcePoolPaginatedList`

NewResourcePoolPaginatedList instantiates a new ResourcePoolPaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcePoolPaginatedListWithDefaults

`func NewResourcePoolPaginatedListWithDefaults() *ResourcePoolPaginatedList`

NewResourcePoolPaginatedListWithDefaults instantiates a new ResourcePoolPaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResourcePoolPaginatedList) GetData() []ResourcePoolWithStats`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResourcePoolPaginatedList) GetDataOk() (*[]ResourcePoolWithStats, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResourcePoolPaginatedList) SetData(v []ResourcePoolWithStats)`

SetData sets Data field to given value.


### GetMeta

`func (o *ResourcePoolPaginatedList) GetMeta() PaginatedResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ResourcePoolPaginatedList) GetMetaOk() (*PaginatedResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ResourcePoolPaginatedList) SetMeta(v PaginatedResponseMeta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *ResourcePoolPaginatedList) GetLinks() PaginatedResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ResourcePoolPaginatedList) GetLinksOk() (*PaginatedResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ResourcePoolPaginatedList) SetLinks(v PaginatedResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ResourcePoolPaginatedList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


