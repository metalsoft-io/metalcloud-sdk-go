# CreateLogicalNetworkProfileVniAllocationStrategyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | [**AllocationStrategyKind**](AllocationStrategyKind.md) |  | 
**Scope** | [**CreateResourceScope**](CreateResourceScope.md) |  | 
**Vni** | **int32** |  | 

## Methods

### NewCreateLogicalNetworkProfileVniAllocationStrategyRequest

`func NewCreateLogicalNetworkProfileVniAllocationStrategyRequest(kind AllocationStrategyKind, scope CreateResourceScope, vni int32, ) *CreateLogicalNetworkProfileVniAllocationStrategyRequest`

NewCreateLogicalNetworkProfileVniAllocationStrategyRequest instantiates a new CreateLogicalNetworkProfileVniAllocationStrategyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLogicalNetworkProfileVniAllocationStrategyRequestWithDefaults

`func NewCreateLogicalNetworkProfileVniAllocationStrategyRequestWithDefaults() *CreateLogicalNetworkProfileVniAllocationStrategyRequest`

NewCreateLogicalNetworkProfileVniAllocationStrategyRequestWithDefaults instantiates a new CreateLogicalNetworkProfileVniAllocationStrategyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CreateLogicalNetworkProfileVniAllocationStrategyRequest) GetKind() AllocationStrategyKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateLogicalNetworkProfileVniAllocationStrategyRequest) GetKindOk() (*AllocationStrategyKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateLogicalNetworkProfileVniAllocationStrategyRequest) SetKind(v AllocationStrategyKind)`

SetKind sets Kind field to given value.


### GetScope

`func (o *CreateLogicalNetworkProfileVniAllocationStrategyRequest) GetScope() CreateResourceScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateLogicalNetworkProfileVniAllocationStrategyRequest) GetScopeOk() (*CreateResourceScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateLogicalNetworkProfileVniAllocationStrategyRequest) SetScope(v CreateResourceScope)`

SetScope sets Scope field to given value.


### GetVni

`func (o *CreateLogicalNetworkProfileVniAllocationStrategyRequest) GetVni() int32`

GetVni returns the Vni field if non-nil, zero value otherwise.

### GetVniOk

`func (o *CreateLogicalNetworkProfileVniAllocationStrategyRequest) GetVniOk() (*int32, bool)`

GetVniOk returns a tuple with the Vni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVni

`func (o *CreateLogicalNetworkProfileVniAllocationStrategyRequest) SetVni(v int32)`

SetVni sets Vni field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


