# VlanAllocationStrategy1

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

### NewVlanAllocationStrategy1

`func NewVlanAllocationStrategy1(id int32, createdAt time.Time, updatedAt time.Time, kind AllocationStrategyKind, scope ResourceScope, granularityLevel NullableVlanAllocationGranularityLevel, vlanId int32, ) *VlanAllocationStrategy1`

NewVlanAllocationStrategy1 instantiates a new VlanAllocationStrategy1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanAllocationStrategy1WithDefaults

`func NewVlanAllocationStrategy1WithDefaults() *VlanAllocationStrategy1`

NewVlanAllocationStrategy1WithDefaults instantiates a new VlanAllocationStrategy1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VlanAllocationStrategy1) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VlanAllocationStrategy1) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VlanAllocationStrategy1) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *VlanAllocationStrategy1) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VlanAllocationStrategy1) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VlanAllocationStrategy1) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *VlanAllocationStrategy1) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VlanAllocationStrategy1) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VlanAllocationStrategy1) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetKind

`func (o *VlanAllocationStrategy1) GetKind() AllocationStrategyKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *VlanAllocationStrategy1) GetKindOk() (*AllocationStrategyKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *VlanAllocationStrategy1) SetKind(v AllocationStrategyKind)`

SetKind sets Kind field to given value.


### GetScope

`func (o *VlanAllocationStrategy1) GetScope() ResourceScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *VlanAllocationStrategy1) GetScopeOk() (*ResourceScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *VlanAllocationStrategy1) SetScope(v ResourceScope)`

SetScope sets Scope field to given value.


### GetGranularityLevel

`func (o *VlanAllocationStrategy1) GetGranularityLevel() VlanAllocationGranularityLevel`

GetGranularityLevel returns the GranularityLevel field if non-nil, zero value otherwise.

### GetGranularityLevelOk

`func (o *VlanAllocationStrategy1) GetGranularityLevelOk() (*VlanAllocationGranularityLevel, bool)`

GetGranularityLevelOk returns a tuple with the GranularityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularityLevel

`func (o *VlanAllocationStrategy1) SetGranularityLevel(v VlanAllocationGranularityLevel)`

SetGranularityLevel sets GranularityLevel field to given value.


### SetGranularityLevelNil

`func (o *VlanAllocationStrategy1) SetGranularityLevelNil(b bool)`

 SetGranularityLevelNil sets the value for GranularityLevel to be an explicit nil

### UnsetGranularityLevel
`func (o *VlanAllocationStrategy1) UnsetGranularityLevel()`

UnsetGranularityLevel ensures that no value is present for GranularityLevel, not even an explicit nil
### GetVlanId

`func (o *VlanAllocationStrategy1) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *VlanAllocationStrategy1) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *VlanAllocationStrategy1) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


