# UpdateVMType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | Display name of the VM Pool type | [optional] 
**Label** | Pointer to **string** | Label of the VM Pool type | [optional] 
**IsExperimental** | Pointer to **float32** | Flag to indicate if the VM Pool is experimental. 1 for true, 0 for false. Default is 0. | [optional] 
**Tags** | Pointer to **[]string** | Tags for the VM Type. | [optional] 
**ForUnmanagedVMsOnly** | Pointer to **float32** | Flag to indicate if the VM Pool is for unmanaged VMs only. 1 for true, 0 for false. Default is 0. | [optional] 

## Methods

### NewUpdateVMType

`func NewUpdateVMType() *UpdateVMType`

NewUpdateVMType instantiates a new UpdateVMType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVMTypeWithDefaults

`func NewUpdateVMTypeWithDefaults() *UpdateVMType`

NewUpdateVMTypeWithDefaults instantiates a new UpdateVMType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *UpdateVMType) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateVMType) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateVMType) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateVMType) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLabel

`func (o *UpdateVMType) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateVMType) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateVMType) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpdateVMType) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetIsExperimental

`func (o *UpdateVMType) GetIsExperimental() float32`

GetIsExperimental returns the IsExperimental field if non-nil, zero value otherwise.

### GetIsExperimentalOk

`func (o *UpdateVMType) GetIsExperimentalOk() (*float32, bool)`

GetIsExperimentalOk returns a tuple with the IsExperimental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExperimental

`func (o *UpdateVMType) SetIsExperimental(v float32)`

SetIsExperimental sets IsExperimental field to given value.

### HasIsExperimental

`func (o *UpdateVMType) HasIsExperimental() bool`

HasIsExperimental returns a boolean if a field has been set.

### GetTags

`func (o *UpdateVMType) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateVMType) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateVMType) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateVMType) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetForUnmanagedVMsOnly

`func (o *UpdateVMType) GetForUnmanagedVMsOnly() float32`

GetForUnmanagedVMsOnly returns the ForUnmanagedVMsOnly field if non-nil, zero value otherwise.

### GetForUnmanagedVMsOnlyOk

`func (o *UpdateVMType) GetForUnmanagedVMsOnlyOk() (*float32, bool)`

GetForUnmanagedVMsOnlyOk returns a tuple with the ForUnmanagedVMsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForUnmanagedVMsOnly

`func (o *UpdateVMType) SetForUnmanagedVMsOnly(v float32)`

SetForUnmanagedVMsOnly sets ForUnmanagedVMsOnly field to given value.

### HasForUnmanagedVMsOnly

`func (o *UpdateVMType) HasForUnmanagedVMsOnly() bool`

HasForUnmanagedVMsOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


