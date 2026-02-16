# CreateAutoVniAllocationStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | [**AllocationStrategyKind**](AllocationStrategyKind.md) |  | 
**Scope** | [**CreateResourceScope**](CreateResourceScope.md) |  | 
**GranularityLevel** | Pointer to [**NullableVniAllocationGranularityLevel**](VniAllocationGranularityLevel.md) |  | [optional] 

## Methods

### NewCreateAutoVniAllocationStrategy

`func NewCreateAutoVniAllocationStrategy(kind AllocationStrategyKind, scope CreateResourceScope, ) *CreateAutoVniAllocationStrategy`

NewCreateAutoVniAllocationStrategy instantiates a new CreateAutoVniAllocationStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAutoVniAllocationStrategyWithDefaults

`func NewCreateAutoVniAllocationStrategyWithDefaults() *CreateAutoVniAllocationStrategy`

NewCreateAutoVniAllocationStrategyWithDefaults instantiates a new CreateAutoVniAllocationStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CreateAutoVniAllocationStrategy) GetKind() AllocationStrategyKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateAutoVniAllocationStrategy) GetKindOk() (*AllocationStrategyKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateAutoVniAllocationStrategy) SetKind(v AllocationStrategyKind)`

SetKind sets Kind field to given value.


### GetScope

`func (o *CreateAutoVniAllocationStrategy) GetScope() CreateResourceScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateAutoVniAllocationStrategy) GetScopeOk() (*CreateResourceScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateAutoVniAllocationStrategy) SetScope(v CreateResourceScope)`

SetScope sets Scope field to given value.


### GetGranularityLevel

`func (o *CreateAutoVniAllocationStrategy) GetGranularityLevel() VniAllocationGranularityLevel`

GetGranularityLevel returns the GranularityLevel field if non-nil, zero value otherwise.

### GetGranularityLevelOk

`func (o *CreateAutoVniAllocationStrategy) GetGranularityLevelOk() (*VniAllocationGranularityLevel, bool)`

GetGranularityLevelOk returns a tuple with the GranularityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularityLevel

`func (o *CreateAutoVniAllocationStrategy) SetGranularityLevel(v VniAllocationGranularityLevel)`

SetGranularityLevel sets GranularityLevel field to given value.

### HasGranularityLevel

`func (o *CreateAutoVniAllocationStrategy) HasGranularityLevel() bool`

HasGranularityLevel returns a boolean if a field has been set.

### SetGranularityLevelNil

`func (o *CreateAutoVniAllocationStrategy) SetGranularityLevelNil(b bool)`

 SetGranularityLevelNil sets the value for GranularityLevel to be an explicit nil

### UnsetGranularityLevel
`func (o *CreateAutoVniAllocationStrategy) UnsetGranularityLevel()`

UnsetGranularityLevel ensures that no value is present for GranularityLevel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


