# VrrpGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualAddress** | Pointer to **NullableString** | VRRP virtual (gateway) address. | [optional] 
**Priority** | Pointer to **NullableInt32** | VRRP priority (higher wins mastership). | [optional] 

## Methods

### NewVrrpGroup

`func NewVrrpGroup() *VrrpGroup`

NewVrrpGroup instantiates a new VrrpGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVrrpGroupWithDefaults

`func NewVrrpGroupWithDefaults() *VrrpGroup`

NewVrrpGroupWithDefaults instantiates a new VrrpGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualAddress

`func (o *VrrpGroup) GetVirtualAddress() string`

GetVirtualAddress returns the VirtualAddress field if non-nil, zero value otherwise.

### GetVirtualAddressOk

`func (o *VrrpGroup) GetVirtualAddressOk() (*string, bool)`

GetVirtualAddressOk returns a tuple with the VirtualAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualAddress

`func (o *VrrpGroup) SetVirtualAddress(v string)`

SetVirtualAddress sets VirtualAddress field to given value.

### HasVirtualAddress

`func (o *VrrpGroup) HasVirtualAddress() bool`

HasVirtualAddress returns a boolean if a field has been set.

### SetVirtualAddressNil

`func (o *VrrpGroup) SetVirtualAddressNil(b bool)`

 SetVirtualAddressNil sets the value for VirtualAddress to be an explicit nil

### UnsetVirtualAddress
`func (o *VrrpGroup) UnsetVirtualAddress()`

UnsetVirtualAddress ensures that no value is present for VirtualAddress, not even an explicit nil
### GetPriority

`func (o *VrrpGroup) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *VrrpGroup) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *VrrpGroup) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *VrrpGroup) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *VrrpGroup) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *VrrpGroup) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


