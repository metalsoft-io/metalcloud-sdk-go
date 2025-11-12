# CreateManualIpv4SubnetAllocationStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | [**AllocationStrategyKind**](AllocationStrategyKind.md) |  | 
**Scope** | [**CreateResourceScope**](CreateResourceScope.md) |  | 
**GatewayPlacement** | Pointer to [**SubnetGatewayPlacement**](SubnetGatewayPlacement.md) |  | [optional] [default to DEFAULT]
**SubnetId** | **int32** |  | 

## Methods

### NewCreateManualIpv4SubnetAllocationStrategy

`func NewCreateManualIpv4SubnetAllocationStrategy(kind AllocationStrategyKind, scope CreateResourceScope, subnetId int32, ) *CreateManualIpv4SubnetAllocationStrategy`

NewCreateManualIpv4SubnetAllocationStrategy instantiates a new CreateManualIpv4SubnetAllocationStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateManualIpv4SubnetAllocationStrategyWithDefaults

`func NewCreateManualIpv4SubnetAllocationStrategyWithDefaults() *CreateManualIpv4SubnetAllocationStrategy`

NewCreateManualIpv4SubnetAllocationStrategyWithDefaults instantiates a new CreateManualIpv4SubnetAllocationStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CreateManualIpv4SubnetAllocationStrategy) GetKind() AllocationStrategyKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateManualIpv4SubnetAllocationStrategy) GetKindOk() (*AllocationStrategyKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateManualIpv4SubnetAllocationStrategy) SetKind(v AllocationStrategyKind)`

SetKind sets Kind field to given value.


### GetScope

`func (o *CreateManualIpv4SubnetAllocationStrategy) GetScope() CreateResourceScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateManualIpv4SubnetAllocationStrategy) GetScopeOk() (*CreateResourceScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateManualIpv4SubnetAllocationStrategy) SetScope(v CreateResourceScope)`

SetScope sets Scope field to given value.


### GetGatewayPlacement

`func (o *CreateManualIpv4SubnetAllocationStrategy) GetGatewayPlacement() SubnetGatewayPlacement`

GetGatewayPlacement returns the GatewayPlacement field if non-nil, zero value otherwise.

### GetGatewayPlacementOk

`func (o *CreateManualIpv4SubnetAllocationStrategy) GetGatewayPlacementOk() (*SubnetGatewayPlacement, bool)`

GetGatewayPlacementOk returns a tuple with the GatewayPlacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPlacement

`func (o *CreateManualIpv4SubnetAllocationStrategy) SetGatewayPlacement(v SubnetGatewayPlacement)`

SetGatewayPlacement sets GatewayPlacement field to given value.

### HasGatewayPlacement

`func (o *CreateManualIpv4SubnetAllocationStrategy) HasGatewayPlacement() bool`

HasGatewayPlacement returns a boolean if a field has been set.

### GetSubnetId

`func (o *CreateManualIpv4SubnetAllocationStrategy) GetSubnetId() int32`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *CreateManualIpv4SubnetAllocationStrategy) GetSubnetIdOk() (*int32, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *CreateManualIpv4SubnetAllocationStrategy) SetSubnetId(v int32)`

SetSubnetId sets SubnetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


