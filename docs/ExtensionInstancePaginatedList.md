# ExtensionInstancePaginatedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ExtensionInstance**](ExtensionInstance.md) |  | 
**Meta** | [**PaginatedResponseMeta**](PaginatedResponseMeta.md) | Metadata about the pagination of the response | 
**Links** | [**PaginatedResponseLinks**](PaginatedResponseLinks.md) | Links to navigate through the paginated results | 

## Methods

### NewExtensionInstancePaginatedList

`func NewExtensionInstancePaginatedList(data []ExtensionInstance, meta PaginatedResponseMeta, links PaginatedResponseLinks, ) *ExtensionInstancePaginatedList`

NewExtensionInstancePaginatedList instantiates a new ExtensionInstancePaginatedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionInstancePaginatedListWithDefaults

`func NewExtensionInstancePaginatedListWithDefaults() *ExtensionInstancePaginatedList`

NewExtensionInstancePaginatedListWithDefaults instantiates a new ExtensionInstancePaginatedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ExtensionInstancePaginatedList) GetData() []ExtensionInstance`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ExtensionInstancePaginatedList) GetDataOk() (*[]ExtensionInstance, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ExtensionInstancePaginatedList) SetData(v []ExtensionInstance)`

SetData sets Data field to given value.


### GetMeta

`func (o *ExtensionInstancePaginatedList) GetMeta() PaginatedResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ExtensionInstancePaginatedList) GetMetaOk() (*PaginatedResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ExtensionInstancePaginatedList) SetMeta(v PaginatedResponseMeta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *ExtensionInstancePaginatedList) GetLinks() PaginatedResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExtensionInstancePaginatedList) GetLinksOk() (*PaginatedResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExtensionInstancePaginatedList) SetLinks(v PaginatedResponseLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


