# CreateAutoIpv4SubnetAllocationStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | [**AllocationStrategyKind**](AllocationStrategyKind.md) |  | 
**Scope** | [**CreateResourceScope**](CreateResourceScope.md) |  | 
**GatewayPlacement** | Pointer to [**SubnetGatewayPlacement**](SubnetGatewayPlacement.md) |  | [optional] [default to DEFAULT]
**SubnetPoolIds** | **[]int32** |  | 
**PrefixLength** | **int32** |  | 

## Methods

### NewCreateAutoIpv4SubnetAllocationStrategy

`func NewCreateAutoIpv4SubnetAllocationStrategy(kind AllocationStrategyKind, scope CreateResourceScope, subnetPoolIds []int32, prefixLength int32, ) *CreateAutoIpv4SubnetAllocationStrategy`

NewCreateAutoIpv4SubnetAllocationStrategy instantiates a new CreateAutoIpv4SubnetAllocationStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAutoIpv4SubnetAllocationStrategyWithDefaults

`func NewCreateAutoIpv4SubnetAllocationStrategyWithDefaults() *CreateAutoIpv4SubnetAllocationStrategy`

NewCreateAutoIpv4SubnetAllocationStrategyWithDefaults instantiates a new CreateAutoIpv4SubnetAllocationStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CreateAutoIpv4SubnetAllocationStrategy) GetKind() AllocationStrategyKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateAutoIpv4SubnetAllocationStrategy) GetKindOk() (*AllocationStrategyKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateAutoIpv4SubnetAllocationStrategy) SetKind(v AllocationStrategyKind)`

SetKind sets Kind field to given value.


### GetScope

`func (o *CreateAutoIpv4SubnetAllocationStrategy) GetScope() CreateResourceScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateAutoIpv4SubnetAllocationStrategy) GetScopeOk() (*CreateResourceScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateAutoIpv4SubnetAllocationStrategy) SetScope(v CreateResourceScope)`

SetScope sets Scope field to given value.


### GetGatewayPlacement

`func (o *CreateAutoIpv4SubnetAllocationStrategy) GetGatewayPlacement() SubnetGatewayPlacement`

GetGatewayPlacement returns the GatewayPlacement field if non-nil, zero value otherwise.

### GetGatewayPlacementOk

`func (o *CreateAutoIpv4SubnetAllocationStrategy) GetGatewayPlacementOk() (*SubnetGatewayPlacement, bool)`

GetGatewayPlacementOk returns a tuple with the GatewayPlacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPlacement

`func (o *CreateAutoIpv4SubnetAllocationStrategy) SetGatewayPlacement(v SubnetGatewayPlacement)`

SetGatewayPlacement sets GatewayPlacement field to given value.

### HasGatewayPlacement

`func (o *CreateAutoIpv4SubnetAllocationStrategy) HasGatewayPlacement() bool`

HasGatewayPlacement returns a boolean if a field has been set.

### GetSubnetPoolIds

`func (o *CreateAutoIpv4SubnetAllocationStrategy) GetSubnetPoolIds() []int32`

GetSubnetPoolIds returns the SubnetPoolIds field if non-nil, zero value otherwise.

### GetSubnetPoolIdsOk

`func (o *CreateAutoIpv4SubnetAllocationStrategy) GetSubnetPoolIdsOk() (*[]int32, bool)`

GetSubnetPoolIdsOk returns a tuple with the SubnetPoolIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetPoolIds

`func (o *CreateAutoIpv4SubnetAllocationStrategy) SetSubnetPoolIds(v []int32)`

SetSubnetPoolIds sets SubnetPoolIds field to given value.


### GetPrefixLength

`func (o *CreateAutoIpv4SubnetAllocationStrategy) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *CreateAutoIpv4SubnetAllocationStrategy) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *CreateAutoIpv4SubnetAllocationStrategy) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


