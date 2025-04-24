# CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | [**AllocationStrategyKind**](AllocationStrategyKind.md) |  | 
**Scope** | [**CreateResourceScope**](CreateResourceScope.md) |  | 
**SubnetId** | **int32** |  | 
**SubnetPoolIds** | **[]int32** |  | 
**PrefixLength** | **int32** |  | 

## Methods

### NewCreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest

`func NewCreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest(kind AllocationStrategyKind, scope CreateResourceScope, subnetId int32, subnetPoolIds []int32, prefixLength int32, ) *CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest`

NewCreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest instantiates a new CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequestWithDefaults

`func NewCreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequestWithDefaults() *CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest`

NewCreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequestWithDefaults instantiates a new CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest) GetKind() AllocationStrategyKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest) GetKindOk() (*AllocationStrategyKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest) SetKind(v AllocationStrategyKind)`

SetKind sets Kind field to given value.


### GetScope

`func (o *CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest) GetScope() CreateResourceScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest) GetScopeOk() (*CreateResourceScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest) SetScope(v CreateResourceScope)`

SetScope sets Scope field to given value.


### GetSubnetId

`func (o *CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest) GetSubnetId() int32`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest) GetSubnetIdOk() (*int32, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest) SetSubnetId(v int32)`

SetSubnetId sets SubnetId field to given value.


### GetSubnetPoolIds

`func (o *CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest) GetSubnetPoolIds() []int32`

GetSubnetPoolIds returns the SubnetPoolIds field if non-nil, zero value otherwise.

### GetSubnetPoolIdsOk

`func (o *CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest) GetSubnetPoolIdsOk() (*[]int32, bool)`

GetSubnetPoolIdsOk returns a tuple with the SubnetPoolIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetPoolIds

`func (o *CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest) SetSubnetPoolIds(v []int32)`

SetSubnetPoolIds sets SubnetPoolIds field to given value.


### GetPrefixLength

`func (o *CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *CreateLogicalNetworkProfileIpv4SubnetAllocationStrategyRequest) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


