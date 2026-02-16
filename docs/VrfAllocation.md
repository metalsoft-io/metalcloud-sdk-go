# VrfAllocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Scope** | [**ResourceScope**](ResourceScope.md) |  | 
**Status** | [**ResourceAllocationStatus**](ResourceAllocationStatus.md) |  | 
**Name** | **string** |  | 

## Methods

### NewVrfAllocation

`func NewVrfAllocation(id int32, scope ResourceScope, status ResourceAllocationStatus, name string, ) *VrfAllocation`

NewVrfAllocation instantiates a new VrfAllocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVrfAllocationWithDefaults

`func NewVrfAllocationWithDefaults() *VrfAllocation`

NewVrfAllocationWithDefaults instantiates a new VrfAllocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VrfAllocation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VrfAllocation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VrfAllocation) SetId(v int32)`

SetId sets Id field to given value.


### GetScope

`func (o *VrfAllocation) GetScope() ResourceScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *VrfAllocation) GetScopeOk() (*ResourceScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *VrfAllocation) SetScope(v ResourceScope)`

SetScope sets Scope field to given value.


### GetStatus

`func (o *VrfAllocation) GetStatus() ResourceAllocationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VrfAllocation) GetStatusOk() (*ResourceAllocationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VrfAllocation) SetStatus(v ResourceAllocationStatus)`

SetStatus sets Status field to given value.


### GetName

`func (o *VrfAllocation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VrfAllocation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VrfAllocation) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


