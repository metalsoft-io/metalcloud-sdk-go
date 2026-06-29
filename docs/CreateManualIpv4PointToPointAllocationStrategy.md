# CreateManualIpv4PointToPointAllocationStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | [**PointToPointAllocationStrategyKind**](PointToPointAllocationStrategyKind.md) |  | 
**Scope** | [**CreateResourceScope**](CreateResourceScope.md) |  | 
**SubnetId** | **int64** | ID of a pre-existing IPAM subnet to use for the link. Must be a /31. The subnet is tracked in the IPAM allocation registry but its lifecycle stays external (it is not deleted when the link is). | 
**InterfaceABinding** | Pointer to [**PointToPointInterfaceBinding**](PointToPointInterfaceBinding.md) | Which interface receives the .0 IP. AUTO &#x3D; system picks by interface id ordering; A_FIRST &#x3D; interface A gets .0; B_FIRST &#x3D; interface B gets .0. | [optional] [default to POINTTOPOINTINTERFACEBINDING_AUTO]

## Methods

### NewCreateManualIpv4PointToPointAllocationStrategy

`func NewCreateManualIpv4PointToPointAllocationStrategy(kind PointToPointAllocationStrategyKind, scope CreateResourceScope, subnetId int64, ) *CreateManualIpv4PointToPointAllocationStrategy`

NewCreateManualIpv4PointToPointAllocationStrategy instantiates a new CreateManualIpv4PointToPointAllocationStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateManualIpv4PointToPointAllocationStrategyWithDefaults

`func NewCreateManualIpv4PointToPointAllocationStrategyWithDefaults() *CreateManualIpv4PointToPointAllocationStrategy`

NewCreateManualIpv4PointToPointAllocationStrategyWithDefaults instantiates a new CreateManualIpv4PointToPointAllocationStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CreateManualIpv4PointToPointAllocationStrategy) GetKind() PointToPointAllocationStrategyKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateManualIpv4PointToPointAllocationStrategy) GetKindOk() (*PointToPointAllocationStrategyKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateManualIpv4PointToPointAllocationStrategy) SetKind(v PointToPointAllocationStrategyKind)`

SetKind sets Kind field to given value.


### GetScope

`func (o *CreateManualIpv4PointToPointAllocationStrategy) GetScope() CreateResourceScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateManualIpv4PointToPointAllocationStrategy) GetScopeOk() (*CreateResourceScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateManualIpv4PointToPointAllocationStrategy) SetScope(v CreateResourceScope)`

SetScope sets Scope field to given value.


### GetSubnetId

`func (o *CreateManualIpv4PointToPointAllocationStrategy) GetSubnetId() int64`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *CreateManualIpv4PointToPointAllocationStrategy) GetSubnetIdOk() (*int64, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *CreateManualIpv4PointToPointAllocationStrategy) SetSubnetId(v int64)`

SetSubnetId sets SubnetId field to given value.


### GetInterfaceABinding

`func (o *CreateManualIpv4PointToPointAllocationStrategy) GetInterfaceABinding() PointToPointInterfaceBinding`

GetInterfaceABinding returns the InterfaceABinding field if non-nil, zero value otherwise.

### GetInterfaceABindingOk

`func (o *CreateManualIpv4PointToPointAllocationStrategy) GetInterfaceABindingOk() (*PointToPointInterfaceBinding, bool)`

GetInterfaceABindingOk returns a tuple with the InterfaceABinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceABinding

`func (o *CreateManualIpv4PointToPointAllocationStrategy) SetInterfaceABinding(v PointToPointInterfaceBinding)`

SetInterfaceABinding sets InterfaceABinding field to given value.

### HasInterfaceABinding

`func (o *CreateManualIpv4PointToPointAllocationStrategy) HasInterfaceABinding() bool`

HasInterfaceABinding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


