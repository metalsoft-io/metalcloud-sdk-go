# PkeyAllocationStrategy2DataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Kind** | [**AllocationStrategyKind**](AllocationStrategyKind.md) |  | 
**Scope** | [**ResourceScope**](ResourceScope.md) |  | 
**Pkey** | **int32** |  | 

## Methods

### NewPkeyAllocationStrategy2DataItem

`func NewPkeyAllocationStrategy2DataItem(id int32, createdAt time.Time, updatedAt time.Time, kind AllocationStrategyKind, scope ResourceScope, pkey int32, ) *PkeyAllocationStrategy2DataItem`

NewPkeyAllocationStrategy2DataItem instantiates a new PkeyAllocationStrategy2DataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkeyAllocationStrategy2DataItemWithDefaults

`func NewPkeyAllocationStrategy2DataItemWithDefaults() *PkeyAllocationStrategy2DataItem`

NewPkeyAllocationStrategy2DataItemWithDefaults instantiates a new PkeyAllocationStrategy2DataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PkeyAllocationStrategy2DataItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PkeyAllocationStrategy2DataItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PkeyAllocationStrategy2DataItem) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *PkeyAllocationStrategy2DataItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PkeyAllocationStrategy2DataItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PkeyAllocationStrategy2DataItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PkeyAllocationStrategy2DataItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PkeyAllocationStrategy2DataItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PkeyAllocationStrategy2DataItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetKind

`func (o *PkeyAllocationStrategy2DataItem) GetKind() AllocationStrategyKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PkeyAllocationStrategy2DataItem) GetKindOk() (*AllocationStrategyKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PkeyAllocationStrategy2DataItem) SetKind(v AllocationStrategyKind)`

SetKind sets Kind field to given value.


### GetScope

`func (o *PkeyAllocationStrategy2DataItem) GetScope() ResourceScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *PkeyAllocationStrategy2DataItem) GetScopeOk() (*ResourceScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *PkeyAllocationStrategy2DataItem) SetScope(v ResourceScope)`

SetScope sets Scope field to given value.


### GetPkey

`func (o *PkeyAllocationStrategy2DataItem) GetPkey() int32`

GetPkey returns the Pkey field if non-nil, zero value otherwise.

### GetPkeyOk

`func (o *PkeyAllocationStrategy2DataItem) GetPkeyOk() (*int32, bool)`

GetPkeyOk returns a tuple with the Pkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkey

`func (o *PkeyAllocationStrategy2DataItem) SetPkey(v int32)`

SetPkey sets Pkey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


