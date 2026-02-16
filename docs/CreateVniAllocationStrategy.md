# CreateVniAllocationStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | [**AllocationStrategyKind**](AllocationStrategyKind.md) |  | 
**Scope** | [**CreateResourceScope**](CreateResourceScope.md) |  | 
**GranularityLevel** | Pointer to [**NullableVniAllocationGranularityLevel**](VniAllocationGranularityLevel.md) |  | [optional] 
**Vni** | **int32** |  | 

## Methods

### NewCreateVniAllocationStrategy

`func NewCreateVniAllocationStrategy(kind AllocationStrategyKind, scope CreateResourceScope, vni int32, ) *CreateVniAllocationStrategy`

NewCreateVniAllocationStrategy instantiates a new CreateVniAllocationStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVniAllocationStrategyWithDefaults

`func NewCreateVniAllocationStrategyWithDefaults() *CreateVniAllocationStrategy`

NewCreateVniAllocationStrategyWithDefaults instantiates a new CreateVniAllocationStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CreateVniAllocationStrategy) GetKind() AllocationStrategyKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateVniAllocationStrategy) GetKindOk() (*AllocationStrategyKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateVniAllocationStrategy) SetKind(v AllocationStrategyKind)`

SetKind sets Kind field to given value.


### GetScope

`func (o *CreateVniAllocationStrategy) GetScope() CreateResourceScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateVniAllocationStrategy) GetScopeOk() (*CreateResourceScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateVniAllocationStrategy) SetScope(v CreateResourceScope)`

SetScope sets Scope field to given value.


### GetGranularityLevel

`func (o *CreateVniAllocationStrategy) GetGranularityLevel() VniAllocationGranularityLevel`

GetGranularityLevel returns the GranularityLevel field if non-nil, zero value otherwise.

### GetGranularityLevelOk

`func (o *CreateVniAllocationStrategy) GetGranularityLevelOk() (*VniAllocationGranularityLevel, bool)`

GetGranularityLevelOk returns a tuple with the GranularityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularityLevel

`func (o *CreateVniAllocationStrategy) SetGranularityLevel(v VniAllocationGranularityLevel)`

SetGranularityLevel sets GranularityLevel field to given value.

### HasGranularityLevel

`func (o *CreateVniAllocationStrategy) HasGranularityLevel() bool`

HasGranularityLevel returns a boolean if a field has been set.

### SetGranularityLevelNil

`func (o *CreateVniAllocationStrategy) SetGranularityLevelNil(b bool)`

 SetGranularityLevelNil sets the value for GranularityLevel to be an explicit nil

### UnsetGranularityLevel
`func (o *CreateVniAllocationStrategy) UnsetGranularityLevel()`

UnsetGranularityLevel ensures that no value is present for GranularityLevel, not even an explicit nil
### GetVni

`func (o *CreateVniAllocationStrategy) GetVni() int32`

GetVni returns the Vni field if non-nil, zero value otherwise.

### GetVniOk

`func (o *CreateVniAllocationStrategy) GetVniOk() (*int32, bool)`

GetVniOk returns a tuple with the Vni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVni

`func (o *CreateVniAllocationStrategy) SetVni(v int32)`

SetVni sets Vni field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


