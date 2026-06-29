# CreateManualIpv6PointToPointAllocationStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | [**PointToPointAllocationStrategyKind**](PointToPointAllocationStrategyKind.md) |  | 
**Scope** | [**CreateResourceScope**](CreateResourceScope.md) |  | 
**SubnetId** | **int64** | ID of a pre-existing IPAM subnet to use for the link. Must be a /127. The subnet is tracked in the IPAM allocation registry but its lifecycle stays external (it is not deleted when the link is). | 
**InterfaceABinding** | Pointer to [**PointToPointInterfaceBinding**](PointToPointInterfaceBinding.md) |  | [optional] [default to POINTTOPOINTINTERFACEBINDING_AUTO]

## Methods

### NewCreateManualIpv6PointToPointAllocationStrategy

`func NewCreateManualIpv6PointToPointAllocationStrategy(kind PointToPointAllocationStrategyKind, scope CreateResourceScope, subnetId int64, ) *CreateManualIpv6PointToPointAllocationStrategy`

NewCreateManualIpv6PointToPointAllocationStrategy instantiates a new CreateManualIpv6PointToPointAllocationStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateManualIpv6PointToPointAllocationStrategyWithDefaults

`func NewCreateManualIpv6PointToPointAllocationStrategyWithDefaults() *CreateManualIpv6PointToPointAllocationStrategy`

NewCreateManualIpv6PointToPointAllocationStrategyWithDefaults instantiates a new CreateManualIpv6PointToPointAllocationStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CreateManualIpv6PointToPointAllocationStrategy) GetKind() PointToPointAllocationStrategyKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateManualIpv6PointToPointAllocationStrategy) GetKindOk() (*PointToPointAllocationStrategyKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateManualIpv6PointToPointAllocationStrategy) SetKind(v PointToPointAllocationStrategyKind)`

SetKind sets Kind field to given value.


### GetScope

`func (o *CreateManualIpv6PointToPointAllocationStrategy) GetScope() CreateResourceScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateManualIpv6PointToPointAllocationStrategy) GetScopeOk() (*CreateResourceScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateManualIpv6PointToPointAllocationStrategy) SetScope(v CreateResourceScope)`

SetScope sets Scope field to given value.


### GetSubnetId

`func (o *CreateManualIpv6PointToPointAllocationStrategy) GetSubnetId() int64`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *CreateManualIpv6PointToPointAllocationStrategy) GetSubnetIdOk() (*int64, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *CreateManualIpv6PointToPointAllocationStrategy) SetSubnetId(v int64)`

SetSubnetId sets SubnetId field to given value.


### GetInterfaceABinding

`func (o *CreateManualIpv6PointToPointAllocationStrategy) GetInterfaceABinding() PointToPointInterfaceBinding`

GetInterfaceABinding returns the InterfaceABinding field if non-nil, zero value otherwise.

### GetInterfaceABindingOk

`func (o *CreateManualIpv6PointToPointAllocationStrategy) GetInterfaceABindingOk() (*PointToPointInterfaceBinding, bool)`

GetInterfaceABindingOk returns a tuple with the InterfaceABinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceABinding

`func (o *CreateManualIpv6PointToPointAllocationStrategy) SetInterfaceABinding(v PointToPointInterfaceBinding)`

SetInterfaceABinding sets InterfaceABinding field to given value.

### HasInterfaceABinding

`func (o *CreateManualIpv6PointToPointAllocationStrategy) HasInterfaceABinding() bool`

HasInterfaceABinding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


