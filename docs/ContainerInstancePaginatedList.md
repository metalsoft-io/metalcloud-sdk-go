# ContainerInstancePaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ContainerInstance**](ContainerInstance.md) |  | 
**Meta** | [**PaginatedResponseMeta**](PaginatedResponseMeta.md) | Metadata about the pagination of the response | 
**Links** | Pointer to [**PaginatedResponseLinks**](PaginatedResponseLinks.md) | Links to navigate through the paginated results | [optional] 

## Methods

### NewContainerInstancePaginatedList

`func NewContainerInstancePaginatedList(data []ContainerInstance, meta PaginatedResponseMeta, ) *ContainerInstancePaginatedList`

NewContainerInstancePaginatedList instantiates a new ContainerInstancePaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerInstancePaginatedListWithDefaults

`func NewContainerInstancePaginatedListWithDefaults() *ContainerInstancePaginatedList`

NewContainerInstancePaginatedListWithDefaults instantiates a new ContainerInstancePaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ContainerInstancePaginatedList) GetData() []ContainerInstance`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ContainerInstancePaginatedList) GetDataOk() (*[]ContainerInstance, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ContainerInstancePaginatedList) SetData(v []ContainerInstance)`

SetData sets Data field to given value.


### GetMeta

`func (o *ContainerInstancePaginatedList) GetMeta() PaginatedResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ContainerInstancePaginatedList) GetMetaOk() (*PaginatedResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ContainerInstancePaginatedList) SetMeta(v PaginatedResponseMeta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *ContainerInstancePaginatedList) GetLinks() PaginatedResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ContainerInstancePaginatedList) GetLinksOk() (*PaginatedResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ContainerInstancePaginatedList) SetLinks(v PaginatedResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ContainerInstancePaginatedList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


