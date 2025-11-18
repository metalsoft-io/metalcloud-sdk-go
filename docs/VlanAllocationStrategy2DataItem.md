# VlanAllocationStrategy2DataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Kind** | [**AllocationStrategyKind**](AllocationStrategyKind.md) |  | 
**Scope** | [**ResourceScope**](ResourceScope.md) |  | 
**GranularityLevel** | [**NullableVlanAllocationGranularityLevel**](VlanAllocationGranularityLevel.md) |  | 
**VlanId** | **int32** |  | 

## Methods

### NewVlanAllocationStrategy2DataItem

`func NewVlanAllocationStrategy2DataItem(id int32, createdAt time.Time, updatedAt time.Time, kind AllocationStrategyKind, scope ResourceScope, granularityLevel NullableVlanAllocationGranularityLevel, vlanId int32, ) *VlanAllocationStrategy2DataItem`

NewVlanAllocationStrategy2DataItem instantiates a new VlanAllocationStrategy2DataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanAllocationStrategy2DataItemWithDefaults

`func NewVlanAllocationStrategy2DataItemWithDefaults() *VlanAllocationStrategy2DataItem`

NewVlanAllocationStrategy2DataItemWithDefaults instantiates a new VlanAllocationStrategy2DataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VlanAllocationStrategy2DataItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VlanAllocationStrategy2DataItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VlanAllocationStrategy2DataItem) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *VlanAllocationStrategy2DataItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VlanAllocationStrategy2DataItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VlanAllocationStrategy2DataItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *VlanAllocationStrategy2DataItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VlanAllocationStrategy2DataItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VlanAllocationStrategy2DataItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetKind

`func (o *VlanAllocationStrategy2DataItem) GetKind() AllocationStrategyKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *VlanAllocationStrategy2DataItem) GetKindOk() (*AllocationStrategyKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *VlanAllocationStrategy2DataItem) SetKind(v AllocationStrategyKind)`

SetKind sets Kind field to given value.


### GetScope

`func (o *VlanAllocationStrategy2DataItem) GetScope() ResourceScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *VlanAllocationStrategy2DataItem) GetScopeOk() (*ResourceScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *VlanAllocationStrategy2DataItem) SetScope(v ResourceScope)`

SetScope sets Scope field to given value.


### GetGranularityLevel

`func (o *VlanAllocationStrategy2DataItem) GetGranularityLevel() VlanAllocationGranularityLevel`

GetGranularityLevel returns the GranularityLevel field if non-nil, zero value otherwise.

### GetGranularityLevelOk

`func (o *VlanAllocationStrategy2DataItem) GetGranularityLevelOk() (*VlanAllocationGranularityLevel, bool)`

GetGranularityLevelOk returns a tuple with the GranularityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularityLevel

`func (o *VlanAllocationStrategy2DataItem) SetGranularityLevel(v VlanAllocationGranularityLevel)`

SetGranularityLevel sets GranularityLevel field to given value.


### SetGranularityLevelNil

`func (o *VlanAllocationStrategy2DataItem) SetGranularityLevelNil(b bool)`

 SetGranularityLevelNil sets the value for GranularityLevel to be an explicit nil

### UnsetGranularityLevel
`func (o *VlanAllocationStrategy2DataItem) UnsetGranularityLevel()`

UnsetGranularityLevel ensures that no value is present for GranularityLevel, not even an explicit nil
### GetVlanId

`func (o *VlanAllocationStrategy2DataItem) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *VlanAllocationStrategy2DataItem) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *VlanAllocationStrategy2DataItem) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


