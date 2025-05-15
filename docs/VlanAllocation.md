# VlanAllocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Scope** | [**ResourceScope**](ResourceScope.md) |  | 
**Status** | [**ResourceAllocationStatus**](ResourceAllocationStatus.md) |  | 
**VlanId** | **int32** |  | 

## Methods

### NewVlanAllocation

`func NewVlanAllocation(id int32, scope ResourceScope, status ResourceAllocationStatus, vlanId int32, ) *VlanAllocation`

NewVlanAllocation instantiates a new VlanAllocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanAllocationWithDefaults

`func NewVlanAllocationWithDefaults() *VlanAllocation`

NewVlanAllocationWithDefaults instantiates a new VlanAllocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VlanAllocation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VlanAllocation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VlanAllocation) SetId(v int32)`

SetId sets Id field to given value.


### GetScope

`func (o *VlanAllocation) GetScope() ResourceScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *VlanAllocation) GetScopeOk() (*ResourceScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *VlanAllocation) SetScope(v ResourceScope)`

SetScope sets Scope field to given value.


### GetStatus

`func (o *VlanAllocation) GetStatus() ResourceAllocationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VlanAllocation) GetStatusOk() (*ResourceAllocationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VlanAllocation) SetStatus(v ResourceAllocationStatus)`

SetStatus sets Status field to given value.


### GetVlanId

`func (o *VlanAllocation) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *VlanAllocation) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *VlanAllocation) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


