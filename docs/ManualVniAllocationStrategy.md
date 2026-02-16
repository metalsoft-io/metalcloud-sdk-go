# ManualVniAllocationStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Kind** | [**AllocationStrategyKind**](AllocationStrategyKind.md) |  | 
**Scope** | [**ResourceScope**](ResourceScope.md) |  | 
**GranularityLevel** | Pointer to [**NullableVniAllocationGranularityLevel**](VniAllocationGranularityLevel.md) |  | [optional] 
**Vni** | **int32** |  | 

## Methods

### NewManualVniAllocationStrategy

`func NewManualVniAllocationStrategy(id int32, createdAt time.Time, updatedAt time.Time, kind AllocationStrategyKind, scope ResourceScope, vni int32, ) *ManualVniAllocationStrategy`

NewManualVniAllocationStrategy instantiates a new ManualVniAllocationStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualVniAllocationStrategyWithDefaults

`func NewManualVniAllocationStrategyWithDefaults() *ManualVniAllocationStrategy`

NewManualVniAllocationStrategyWithDefaults instantiates a new ManualVniAllocationStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManualVniAllocationStrategy) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManualVniAllocationStrategy) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManualVniAllocationStrategy) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ManualVniAllocationStrategy) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ManualVniAllocationStrategy) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ManualVniAllocationStrategy) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ManualVniAllocationStrategy) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ManualVniAllocationStrategy) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ManualVniAllocationStrategy) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetKind

`func (o *ManualVniAllocationStrategy) GetKind() AllocationStrategyKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ManualVniAllocationStrategy) GetKindOk() (*AllocationStrategyKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ManualVniAllocationStrategy) SetKind(v AllocationStrategyKind)`

SetKind sets Kind field to given value.


### GetScope

`func (o *ManualVniAllocationStrategy) GetScope() ResourceScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ManualVniAllocationStrategy) GetScopeOk() (*ResourceScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ManualVniAllocationStrategy) SetScope(v ResourceScope)`

SetScope sets Scope field to given value.


### GetGranularityLevel

`func (o *ManualVniAllocationStrategy) GetGranularityLevel() VniAllocationGranularityLevel`

GetGranularityLevel returns the GranularityLevel field if non-nil, zero value otherwise.

### GetGranularityLevelOk

`func (o *ManualVniAllocationStrategy) GetGranularityLevelOk() (*VniAllocationGranularityLevel, bool)`

GetGranularityLevelOk returns a tuple with the GranularityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularityLevel

`func (o *ManualVniAllocationStrategy) SetGranularityLevel(v VniAllocationGranularityLevel)`

SetGranularityLevel sets GranularityLevel field to given value.

### HasGranularityLevel

`func (o *ManualVniAllocationStrategy) HasGranularityLevel() bool`

HasGranularityLevel returns a boolean if a field has been set.

### SetGranularityLevelNil

`func (o *ManualVniAllocationStrategy) SetGranularityLevelNil(b bool)`

 SetGranularityLevelNil sets the value for GranularityLevel to be an explicit nil

### UnsetGranularityLevel
`func (o *ManualVniAllocationStrategy) UnsetGranularityLevel()`

UnsetGranularityLevel ensures that no value is present for GranularityLevel, not even an explicit nil
### GetVni

`func (o *ManualVniAllocationStrategy) GetVni() int32`

GetVni returns the Vni field if non-nil, zero value otherwise.

### GetVniOk

`func (o *ManualVniAllocationStrategy) GetVniOk() (*int32, bool)`

GetVniOk returns a tuple with the Vni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVni

`func (o *ManualVniAllocationStrategy) SetVni(v int32)`

SetVni sets Vni field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


