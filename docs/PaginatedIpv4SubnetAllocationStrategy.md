# PaginatedIpv4SubnetAllocationStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Ipv4SubnetAllocationStrategy**](Ipv4SubnetAllocationStrategy.md) |  | 
**Meta** | [**PaginatedResponseMeta**](PaginatedResponseMeta.md) | Metadata about the pagination of the response | 
**Links** | [**PaginatedResponseLinks**](PaginatedResponseLinks.md) | Links to navigate through the paginated results | 

## Methods

### NewPaginatedIpv4SubnetAllocationStrategy

`func NewPaginatedIpv4SubnetAllocationStrategy(data []Ipv4SubnetAllocationStrategy, meta PaginatedResponseMeta, links PaginatedResponseLinks, ) *PaginatedIpv4SubnetAllocationStrategy`

NewPaginatedIpv4SubnetAllocationStrategy instantiates a new PaginatedIpv4SubnetAllocationStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedIpv4SubnetAllocationStrategyWithDefaults

`func NewPaginatedIpv4SubnetAllocationStrategyWithDefaults() *PaginatedIpv4SubnetAllocationStrategy`

NewPaginatedIpv4SubnetAllocationStrategyWithDefaults instantiates a new PaginatedIpv4SubnetAllocationStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PaginatedIpv4SubnetAllocationStrategy) GetData() []Ipv4SubnetAllocationStrategy`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaginatedIpv4SubnetAllocationStrategy) GetDataOk() (*[]Ipv4SubnetAllocationStrategy, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaginatedIpv4SubnetAllocationStrategy) SetData(v []Ipv4SubnetAllocationStrategy)`

SetData sets Data field to given value.


### GetMeta

`func (o *PaginatedIpv4SubnetAllocationStrategy) GetMeta() PaginatedResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PaginatedIpv4SubnetAllocationStrategy) GetMetaOk() (*PaginatedResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PaginatedIpv4SubnetAllocationStrategy) SetMeta(v PaginatedResponseMeta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *PaginatedIpv4SubnetAllocationStrategy) GetLinks() PaginatedResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaginatedIpv4SubnetAllocationStrategy) GetLinksOk() (*PaginatedResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaginatedIpv4SubnetAllocationStrategy) SetLinks(v PaginatedResponseLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


