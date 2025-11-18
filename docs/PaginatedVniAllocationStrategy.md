# PaginatedVniAllocationStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]VniAllocationStrategy2DataItem**](VniAllocationStrategy2DataItem.md) |  | 
**Meta** | [**PaginatedResponseMeta**](PaginatedResponseMeta.md) | Metadata about the pagination of the response | 
**Links** | [**PaginatedResponseLinks**](PaginatedResponseLinks.md) | Links to navigate through the paginated results | 

## Methods

### NewPaginatedVniAllocationStrategy

`func NewPaginatedVniAllocationStrategy(data []VniAllocationStrategy2DataItem, meta PaginatedResponseMeta, links PaginatedResponseLinks, ) *PaginatedVniAllocationStrategy`

NewPaginatedVniAllocationStrategy instantiates a new PaginatedVniAllocationStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedVniAllocationStrategyWithDefaults

`func NewPaginatedVniAllocationStrategyWithDefaults() *PaginatedVniAllocationStrategy`

NewPaginatedVniAllocationStrategyWithDefaults instantiates a new PaginatedVniAllocationStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PaginatedVniAllocationStrategy) GetData() []VniAllocationStrategy2DataItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaginatedVniAllocationStrategy) GetDataOk() (*[]VniAllocationStrategy2DataItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaginatedVniAllocationStrategy) SetData(v []VniAllocationStrategy2DataItem)`

SetData sets Data field to given value.


### GetMeta

`func (o *PaginatedVniAllocationStrategy) GetMeta() PaginatedResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PaginatedVniAllocationStrategy) GetMetaOk() (*PaginatedResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PaginatedVniAllocationStrategy) SetMeta(v PaginatedResponseMeta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *PaginatedVniAllocationStrategy) GetLinks() PaginatedResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaginatedVniAllocationStrategy) GetLinksOk() (*PaginatedResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaginatedVniAllocationStrategy) SetLinks(v PaginatedResponseLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


