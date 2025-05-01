# CreateLogicalNetworkProfileVlanAllocationStrategyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | [**AllocationStrategyKind**](AllocationStrategyKind.md) |  | 
**Scope** | [**CreateResourceScope**](CreateResourceScope.md) |  | 
**GranularityLevel** | Pointer to [**NullableVlanAllocationGranularityLevel**](VlanAllocationGranularityLevel.md) |  | [optional] 
**VlanId** | **int32** |  | 

## Methods

### NewCreateLogicalNetworkProfileVlanAllocationStrategyRequest

`func NewCreateLogicalNetworkProfileVlanAllocationStrategyRequest(kind AllocationStrategyKind, scope CreateResourceScope, vlanId int32, ) *CreateLogicalNetworkProfileVlanAllocationStrategyRequest`

NewCreateLogicalNetworkProfileVlanAllocationStrategyRequest instantiates a new CreateLogicalNetworkProfileVlanAllocationStrategyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLogicalNetworkProfileVlanAllocationStrategyRequestWithDefaults

`func NewCreateLogicalNetworkProfileVlanAllocationStrategyRequestWithDefaults() *CreateLogicalNetworkProfileVlanAllocationStrategyRequest`

NewCreateLogicalNetworkProfileVlanAllocationStrategyRequestWithDefaults instantiates a new CreateLogicalNetworkProfileVlanAllocationStrategyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CreateLogicalNetworkProfileVlanAllocationStrategyRequest) GetKind() AllocationStrategyKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateLogicalNetworkProfileVlanAllocationStrategyRequest) GetKindOk() (*AllocationStrategyKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateLogicalNetworkProfileVlanAllocationStrategyRequest) SetKind(v AllocationStrategyKind)`

SetKind sets Kind field to given value.


### GetScope

`func (o *CreateLogicalNetworkProfileVlanAllocationStrategyRequest) GetScope() CreateResourceScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateLogicalNetworkProfileVlanAllocationStrategyRequest) GetScopeOk() (*CreateResourceScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateLogicalNetworkProfileVlanAllocationStrategyRequest) SetScope(v CreateResourceScope)`

SetScope sets Scope field to given value.


### GetGranularityLevel

`func (o *CreateLogicalNetworkProfileVlanAllocationStrategyRequest) GetGranularityLevel() VlanAllocationGranularityLevel`

GetGranularityLevel returns the GranularityLevel field if non-nil, zero value otherwise.

### GetGranularityLevelOk

`func (o *CreateLogicalNetworkProfileVlanAllocationStrategyRequest) GetGranularityLevelOk() (*VlanAllocationGranularityLevel, bool)`

GetGranularityLevelOk returns a tuple with the GranularityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularityLevel

`func (o *CreateLogicalNetworkProfileVlanAllocationStrategyRequest) SetGranularityLevel(v VlanAllocationGranularityLevel)`

SetGranularityLevel sets GranularityLevel field to given value.

### HasGranularityLevel

`func (o *CreateLogicalNetworkProfileVlanAllocationStrategyRequest) HasGranularityLevel() bool`

HasGranularityLevel returns a boolean if a field has been set.

### SetGranularityLevelNil

`func (o *CreateLogicalNetworkProfileVlanAllocationStrategyRequest) SetGranularityLevelNil(b bool)`

 SetGranularityLevelNil sets the value for GranularityLevel to be an explicit nil

### UnsetGranularityLevel
`func (o *CreateLogicalNetworkProfileVlanAllocationStrategyRequest) UnsetGranularityLevel()`

UnsetGranularityLevel ensures that no value is present for GranularityLevel, not even an explicit nil
### GetVlanId

`func (o *CreateLogicalNetworkProfileVlanAllocationStrategyRequest) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *CreateLogicalNetworkProfileVlanAllocationStrategyRequest) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *CreateLogicalNetworkProfileVlanAllocationStrategyRequest) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


