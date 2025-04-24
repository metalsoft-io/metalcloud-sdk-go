# PaginatedIpv6SubnetAllocationStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Ipv6SubnetAllocationStrategy1DataItem**](Ipv6SubnetAllocationStrategy1DataItem.md) |  | 
**Meta** | [**PaginatedResponseMeta**](PaginatedResponseMeta.md) | Metadata about the pagination of the response | 
**Links** | [**PaginatedResponseLinks**](PaginatedResponseLinks.md) | Links to navigate through the paginated results | 

## Methods

### NewPaginatedIpv6SubnetAllocationStrategy

`func NewPaginatedIpv6SubnetAllocationStrategy(data []Ipv6SubnetAllocationStrategy1DataItem, meta PaginatedResponseMeta, links PaginatedResponseLinks, ) *PaginatedIpv6SubnetAllocationStrategy`

NewPaginatedIpv6SubnetAllocationStrategy instantiates a new PaginatedIpv6SubnetAllocationStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedIpv6SubnetAllocationStrategyWithDefaults

`func NewPaginatedIpv6SubnetAllocationStrategyWithDefaults() *PaginatedIpv6SubnetAllocationStrategy`

NewPaginatedIpv6SubnetAllocationStrategyWithDefaults instantiates a new PaginatedIpv6SubnetAllocationStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PaginatedIpv6SubnetAllocationStrategy) GetData() []Ipv6SubnetAllocationStrategy1DataItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaginatedIpv6SubnetAllocationStrategy) GetDataOk() (*[]Ipv6SubnetAllocationStrategy1DataItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaginatedIpv6SubnetAllocationStrategy) SetData(v []Ipv6SubnetAllocationStrategy1DataItem)`

SetData sets Data field to given value.


### GetMeta

`func (o *PaginatedIpv6SubnetAllocationStrategy) GetMeta() PaginatedResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PaginatedIpv6SubnetAllocationStrategy) GetMetaOk() (*PaginatedResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PaginatedIpv6SubnetAllocationStrategy) SetMeta(v PaginatedResponseMeta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *PaginatedIpv6SubnetAllocationStrategy) GetLinks() PaginatedResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaginatedIpv6SubnetAllocationStrategy) GetLinksOk() (*PaginatedResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaginatedIpv6SubnetAllocationStrategy) SetLinks(v PaginatedResponseLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


