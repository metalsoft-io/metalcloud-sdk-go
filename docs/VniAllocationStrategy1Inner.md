# VniAllocationStrategy1Inner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Kind** | [**AllocationStrategyKind**](AllocationStrategyKind.md) |  | 
**Scope** | [**ResourceScope**](ResourceScope.md) |  | 
**Vni** | **int32** |  | 

## Methods

### NewVniAllocationStrategy1Inner

`func NewVniAllocationStrategy1Inner(id int32, createdAt time.Time, updatedAt time.Time, kind AllocationStrategyKind, scope ResourceScope, vni int32, ) *VniAllocationStrategy1Inner`

NewVniAllocationStrategy1Inner instantiates a new VniAllocationStrategy1Inner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVniAllocationStrategy1InnerWithDefaults

`func NewVniAllocationStrategy1InnerWithDefaults() *VniAllocationStrategy1Inner`

NewVniAllocationStrategy1InnerWithDefaults instantiates a new VniAllocationStrategy1Inner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VniAllocationStrategy1Inner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VniAllocationStrategy1Inner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VniAllocationStrategy1Inner) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *VniAllocationStrategy1Inner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VniAllocationStrategy1Inner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VniAllocationStrategy1Inner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *VniAllocationStrategy1Inner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VniAllocationStrategy1Inner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VniAllocationStrategy1Inner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetKind

`func (o *VniAllocationStrategy1Inner) GetKind() AllocationStrategyKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *VniAllocationStrategy1Inner) GetKindOk() (*AllocationStrategyKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *VniAllocationStrategy1Inner) SetKind(v AllocationStrategyKind)`

SetKind sets Kind field to given value.


### GetScope

`func (o *VniAllocationStrategy1Inner) GetScope() ResourceScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *VniAllocationStrategy1Inner) GetScopeOk() (*ResourceScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *VniAllocationStrategy1Inner) SetScope(v ResourceScope)`

SetScope sets Scope field to given value.


### GetVni

`func (o *VniAllocationStrategy1Inner) GetVni() int32`

GetVni returns the Vni field if non-nil, zero value otherwise.

### GetVniOk

`func (o *VniAllocationStrategy1Inner) GetVniOk() (*int32, bool)`

GetVniOk returns a tuple with the Vni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVni

`func (o *VniAllocationStrategy1Inner) SetVni(v int32)`

SetVni sets Vni field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


