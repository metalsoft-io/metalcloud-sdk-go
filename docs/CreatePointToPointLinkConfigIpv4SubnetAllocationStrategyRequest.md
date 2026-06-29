# CreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | [**PointToPointAllocationStrategyKind**](PointToPointAllocationStrategyKind.md) |  | 
**Scope** | [**CreateResourceScope**](CreateResourceScope.md) |  | 
**SubnetId** | **int64** | ID of a pre-existing IPAM subnet to use for the link. Must be a /31. The subnet is tracked in the IPAM allocation registry but its lifecycle stays external (it is not deleted when the link is). | 
**InterfaceABinding** | Pointer to [**PointToPointInterfaceBinding**](PointToPointInterfaceBinding.md) | Which interface receives the .0 IP. AUTO &#x3D; system picks by interface id ordering; A_FIRST &#x3D; interface A gets .0; B_FIRST &#x3D; interface B gets .0. | [optional] [default to POINTTOPOINTINTERFACEBINDING_AUTO]
**SubnetPoolIds** | **[]int64** | SubnetPool IDs to carve /31 prefixes from. The first pool with availability is used. | 
**PrefixLength** | Pointer to **int32** | Prefix length to carve from the pool. Defaults to 31 for IPv4 P2P. | [optional] [default to 31]

## Methods

### NewCreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest

`func NewCreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest(kind PointToPointAllocationStrategyKind, scope CreateResourceScope, subnetId int64, subnetPoolIds []int64, ) *CreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest`

NewCreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest instantiates a new CreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequestWithDefaults

`func NewCreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequestWithDefaults() *CreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest`

NewCreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequestWithDefaults instantiates a new CreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest) GetKind() PointToPointAllocationStrategyKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest) GetKindOk() (*PointToPointAllocationStrategyKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest) SetKind(v PointToPointAllocationStrategyKind)`

SetKind sets Kind field to given value.


### GetScope

`func (o *CreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest) GetScope() CreateResourceScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest) GetScopeOk() (*CreateResourceScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest) SetScope(v CreateResourceScope)`

SetScope sets Scope field to given value.


### GetSubnetId

`func (o *CreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest) GetSubnetId() int64`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *CreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest) GetSubnetIdOk() (*int64, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *CreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest) SetSubnetId(v int64)`

SetSubnetId sets SubnetId field to given value.


### GetInterfaceABinding

`func (o *CreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest) GetInterfaceABinding() PointToPointInterfaceBinding`

GetInterfaceABinding returns the InterfaceABinding field if non-nil, zero value otherwise.

### GetInterfaceABindingOk

`func (o *CreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest) GetInterfaceABindingOk() (*PointToPointInterfaceBinding, bool)`

GetInterfaceABindingOk returns a tuple with the InterfaceABinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceABinding

`func (o *CreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest) SetInterfaceABinding(v PointToPointInterfaceBinding)`

SetInterfaceABinding sets InterfaceABinding field to given value.

### HasInterfaceABinding

`func (o *CreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest) HasInterfaceABinding() bool`

HasInterfaceABinding returns a boolean if a field has been set.

### GetSubnetPoolIds

`func (o *CreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest) GetSubnetPoolIds() []int64`

GetSubnetPoolIds returns the SubnetPoolIds field if non-nil, zero value otherwise.

### GetSubnetPoolIdsOk

`func (o *CreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest) GetSubnetPoolIdsOk() (*[]int64, bool)`

GetSubnetPoolIdsOk returns a tuple with the SubnetPoolIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetPoolIds

`func (o *CreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest) SetSubnetPoolIds(v []int64)`

SetSubnetPoolIds sets SubnetPoolIds field to given value.


### GetPrefixLength

`func (o *CreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *CreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *CreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.

### HasPrefixLength

`func (o *CreatePointToPointLinkConfigIpv4SubnetAllocationStrategyRequest) HasPrefixLength() bool`

HasPrefixLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


