# CreateAutoIpv6PointToPointAllocationStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | [**PointToPointAllocationStrategyKind**](PointToPointAllocationStrategyKind.md) |  | 
**Scope** | [**CreateResourceScope**](CreateResourceScope.md) |  | 
**SubnetPoolIds** | **[]int64** | SubnetPool IDs to carve /127 prefixes from. The first pool with availability is used. | 
**PrefixLength** | Pointer to **int32** | Prefix length to carve from the pool. Defaults to 127 for IPv6 P2P. | [optional] [default to 127]

## Methods

### NewCreateAutoIpv6PointToPointAllocationStrategy

`func NewCreateAutoIpv6PointToPointAllocationStrategy(kind PointToPointAllocationStrategyKind, scope CreateResourceScope, subnetPoolIds []int64, ) *CreateAutoIpv6PointToPointAllocationStrategy`

NewCreateAutoIpv6PointToPointAllocationStrategy instantiates a new CreateAutoIpv6PointToPointAllocationStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAutoIpv6PointToPointAllocationStrategyWithDefaults

`func NewCreateAutoIpv6PointToPointAllocationStrategyWithDefaults() *CreateAutoIpv6PointToPointAllocationStrategy`

NewCreateAutoIpv6PointToPointAllocationStrategyWithDefaults instantiates a new CreateAutoIpv6PointToPointAllocationStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CreateAutoIpv6PointToPointAllocationStrategy) GetKind() PointToPointAllocationStrategyKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateAutoIpv6PointToPointAllocationStrategy) GetKindOk() (*PointToPointAllocationStrategyKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateAutoIpv6PointToPointAllocationStrategy) SetKind(v PointToPointAllocationStrategyKind)`

SetKind sets Kind field to given value.


### GetScope

`func (o *CreateAutoIpv6PointToPointAllocationStrategy) GetScope() CreateResourceScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateAutoIpv6PointToPointAllocationStrategy) GetScopeOk() (*CreateResourceScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateAutoIpv6PointToPointAllocationStrategy) SetScope(v CreateResourceScope)`

SetScope sets Scope field to given value.


### GetSubnetPoolIds

`func (o *CreateAutoIpv6PointToPointAllocationStrategy) GetSubnetPoolIds() []int64`

GetSubnetPoolIds returns the SubnetPoolIds field if non-nil, zero value otherwise.

### GetSubnetPoolIdsOk

`func (o *CreateAutoIpv6PointToPointAllocationStrategy) GetSubnetPoolIdsOk() (*[]int64, bool)`

GetSubnetPoolIdsOk returns a tuple with the SubnetPoolIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetPoolIds

`func (o *CreateAutoIpv6PointToPointAllocationStrategy) SetSubnetPoolIds(v []int64)`

SetSubnetPoolIds sets SubnetPoolIds field to given value.


### GetPrefixLength

`func (o *CreateAutoIpv6PointToPointAllocationStrategy) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *CreateAutoIpv6PointToPointAllocationStrategy) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *CreateAutoIpv6PointToPointAllocationStrategy) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.

### HasPrefixLength

`func (o *CreateAutoIpv6PointToPointAllocationStrategy) HasPrefixLength() bool`

HasPrefixLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


