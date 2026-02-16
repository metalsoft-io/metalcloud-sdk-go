# CreateManualVniAllocationStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | [**AllocationStrategyKind**](AllocationStrategyKind.md) |  | 
**Scope** | [**CreateResourceScope**](CreateResourceScope.md) |  | 
**GranularityLevel** | Pointer to [**NullableVniAllocationGranularityLevel**](VniAllocationGranularityLevel.md) |  | [optional] 
**Vni** | **int32** |  | 

## Methods

### NewCreateManualVniAllocationStrategy

`func NewCreateManualVniAllocationStrategy(kind AllocationStrategyKind, scope CreateResourceScope, vni int32, ) *CreateManualVniAllocationStrategy`

NewCreateManualVniAllocationStrategy instantiates a new CreateManualVniAllocationStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateManualVniAllocationStrategyWithDefaults

`func NewCreateManualVniAllocationStrategyWithDefaults() *CreateManualVniAllocationStrategy`

NewCreateManualVniAllocationStrategyWithDefaults instantiates a new CreateManualVniAllocationStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CreateManualVniAllocationStrategy) GetKind() AllocationStrategyKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateManualVniAllocationStrategy) GetKindOk() (*AllocationStrategyKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateManualVniAllocationStrategy) SetKind(v AllocationStrategyKind)`

SetKind sets Kind field to given value.


### GetScope

`func (o *CreateManualVniAllocationStrategy) GetScope() CreateResourceScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateManualVniAllocationStrategy) GetScopeOk() (*CreateResourceScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateManualVniAllocationStrategy) SetScope(v CreateResourceScope)`

SetScope sets Scope field to given value.


### GetGranularityLevel

`func (o *CreateManualVniAllocationStrategy) GetGranularityLevel() VniAllocationGranularityLevel`

GetGranularityLevel returns the GranularityLevel field if non-nil, zero value otherwise.

### GetGranularityLevelOk

`func (o *CreateManualVniAllocationStrategy) GetGranularityLevelOk() (*VniAllocationGranularityLevel, bool)`

GetGranularityLevelOk returns a tuple with the GranularityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularityLevel

`func (o *CreateManualVniAllocationStrategy) SetGranularityLevel(v VniAllocationGranularityLevel)`

SetGranularityLevel sets GranularityLevel field to given value.

### HasGranularityLevel

`func (o *CreateManualVniAllocationStrategy) HasGranularityLevel() bool`

HasGranularityLevel returns a boolean if a field has been set.

### SetGranularityLevelNil

`func (o *CreateManualVniAllocationStrategy) SetGranularityLevelNil(b bool)`

 SetGranularityLevelNil sets the value for GranularityLevel to be an explicit nil

### UnsetGranularityLevel
`func (o *CreateManualVniAllocationStrategy) UnsetGranularityLevel()`

UnsetGranularityLevel ensures that no value is present for GranularityLevel, not even an explicit nil
### GetVni

`func (o *CreateManualVniAllocationStrategy) GetVni() int32`

GetVni returns the Vni field if non-nil, zero value otherwise.

### GetVniOk

`func (o *CreateManualVniAllocationStrategy) GetVniOk() (*int32, bool)`

GetVniOk returns a tuple with the Vni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVni

`func (o *CreateManualVniAllocationStrategy) SetVni(v int32)`

SetVni sets Vni field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


