# UpdateVMInstanceGroupMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GuiSettings** | Pointer to [**GenericGUISettings**](GenericGUISettings.md) |  | [optional] 
**Tags** | Pointer to **[]string** | Tags for the VM Instance Group. | [optional] 

## Methods

### NewUpdateVMInstanceGroupMeta

`func NewUpdateVMInstanceGroupMeta() *UpdateVMInstanceGroupMeta`

NewUpdateVMInstanceGroupMeta instantiates a new UpdateVMInstanceGroupMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVMInstanceGroupMetaWithDefaults

`func NewUpdateVMInstanceGroupMetaWithDefaults() *UpdateVMInstanceGroupMeta`

NewUpdateVMInstanceGroupMetaWithDefaults instantiates a new UpdateVMInstanceGroupMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuiSettings

`func (o *UpdateVMInstanceGroupMeta) GetGuiSettings() GenericGUISettings`

GetGuiSettings returns the GuiSettings field if non-nil, zero value otherwise.

### GetGuiSettingsOk

`func (o *UpdateVMInstanceGroupMeta) GetGuiSettingsOk() (*GenericGUISettings, bool)`

GetGuiSettingsOk returns a tuple with the GuiSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuiSettings

`func (o *UpdateVMInstanceGroupMeta) SetGuiSettings(v GenericGUISettings)`

SetGuiSettings sets GuiSettings field to given value.

### HasGuiSettings

`func (o *UpdateVMInstanceGroupMeta) HasGuiSettings() bool`

HasGuiSettings returns a boolean if a field has been set.

### GetTags

`func (o *UpdateVMInstanceGroupMeta) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateVMInstanceGroupMeta) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateVMInstanceGroupMeta) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateVMInstanceGroupMeta) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


