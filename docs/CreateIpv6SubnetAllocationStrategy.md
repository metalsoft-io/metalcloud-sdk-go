# CreateIpv6SubnetAllocationStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | [**AllocationStrategyKind**](AllocationStrategyKind.md) |  | 
**Scope** | [**CreateResourceScope**](CreateResourceScope.md) |  | 
**GatewayPlacement** | Pointer to [**SubnetGatewayPlacement**](SubnetGatewayPlacement.md) |  | [optional] [default to SUBNETGATEWAYPLACEMENT_DEFAULT]
**SubnetId** | **int32** |  | 
**SubnetPoolIds** | **[]int32** |  | 
**PrefixLength** | **int32** |  | 

## Methods

### NewCreateIpv6SubnetAllocationStrategy

`func NewCreateIpv6SubnetAllocationStrategy(kind AllocationStrategyKind, scope CreateResourceScope, subnetId int32, subnetPoolIds []int32, prefixLength int32, ) *CreateIpv6SubnetAllocationStrategy`

NewCreateIpv6SubnetAllocationStrategy instantiates a new CreateIpv6SubnetAllocationStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIpv6SubnetAllocationStrategyWithDefaults

`func NewCreateIpv6SubnetAllocationStrategyWithDefaults() *CreateIpv6SubnetAllocationStrategy`

NewCreateIpv6SubnetAllocationStrategyWithDefaults instantiates a new CreateIpv6SubnetAllocationStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CreateIpv6SubnetAllocationStrategy) GetKind() AllocationStrategyKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateIpv6SubnetAllocationStrategy) GetKindOk() (*AllocationStrategyKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateIpv6SubnetAllocationStrategy) SetKind(v AllocationStrategyKind)`

SetKind sets Kind field to given value.


### GetScope

`func (o *CreateIpv6SubnetAllocationStrategy) GetScope() CreateResourceScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateIpv6SubnetAllocationStrategy) GetScopeOk() (*CreateResourceScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateIpv6SubnetAllocationStrategy) SetScope(v CreateResourceScope)`

SetScope sets Scope field to given value.


### GetGatewayPlacement

`func (o *CreateIpv6SubnetAllocationStrategy) GetGatewayPlacement() SubnetGatewayPlacement`

GetGatewayPlacement returns the GatewayPlacement field if non-nil, zero value otherwise.

### GetGatewayPlacementOk

`func (o *CreateIpv6SubnetAllocationStrategy) GetGatewayPlacementOk() (*SubnetGatewayPlacement, bool)`

GetGatewayPlacementOk returns a tuple with the GatewayPlacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPlacement

`func (o *CreateIpv6SubnetAllocationStrategy) SetGatewayPlacement(v SubnetGatewayPlacement)`

SetGatewayPlacement sets GatewayPlacement field to given value.

### HasGatewayPlacement

`func (o *CreateIpv6SubnetAllocationStrategy) HasGatewayPlacement() bool`

HasGatewayPlacement returns a boolean if a field has been set.

### GetSubnetId

`func (o *CreateIpv6SubnetAllocationStrategy) GetSubnetId() int32`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *CreateIpv6SubnetAllocationStrategy) GetSubnetIdOk() (*int32, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *CreateIpv6SubnetAllocationStrategy) SetSubnetId(v int32)`

SetSubnetId sets SubnetId field to given value.


### GetSubnetPoolIds

`func (o *CreateIpv6SubnetAllocationStrategy) GetSubnetPoolIds() []int32`

GetSubnetPoolIds returns the SubnetPoolIds field if non-nil, zero value otherwise.

### GetSubnetPoolIdsOk

`func (o *CreateIpv6SubnetAllocationStrategy) GetSubnetPoolIdsOk() (*[]int32, bool)`

GetSubnetPoolIdsOk returns a tuple with the SubnetPoolIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetPoolIds

`func (o *CreateIpv6SubnetAllocationStrategy) SetSubnetPoolIds(v []int32)`

SetSubnetPoolIds sets SubnetPoolIds field to given value.


### GetPrefixLength

`func (o *CreateIpv6SubnetAllocationStrategy) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *CreateIpv6SubnetAllocationStrategy) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *CreateIpv6SubnetAllocationStrategy) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


