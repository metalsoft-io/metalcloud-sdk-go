# UpdateVMInstanceMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GuiSettings** | Pointer to [**GenericGUISettings**](GenericGUISettings.md) |  | [optional] 
**Tags** | Pointer to **[]string** | Tags for the VM Instance. | [optional] 

## Methods

### NewUpdateVMInstanceMeta

`func NewUpdateVMInstanceMeta() *UpdateVMInstanceMeta`

NewUpdateVMInstanceMeta instantiates a new UpdateVMInstanceMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVMInstanceMetaWithDefaults

`func NewUpdateVMInstanceMetaWithDefaults() *UpdateVMInstanceMeta`

NewUpdateVMInstanceMetaWithDefaults instantiates a new UpdateVMInstanceMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuiSettings

`func (o *UpdateVMInstanceMeta) GetGuiSettings() GenericGUISettings`

GetGuiSettings returns the GuiSettings field if non-nil, zero value otherwise.

### GetGuiSettingsOk

`func (o *UpdateVMInstanceMeta) GetGuiSettingsOk() (*GenericGUISettings, bool)`

GetGuiSettingsOk returns a tuple with the GuiSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuiSettings

`func (o *UpdateVMInstanceMeta) SetGuiSettings(v GenericGUISettings)`

SetGuiSettings sets GuiSettings field to given value.

### HasGuiSettings

`func (o *UpdateVMInstanceMeta) HasGuiSettings() bool`

HasGuiSettings returns a boolean if a field has been set.

### GetTags

`func (o *UpdateVMInstanceMeta) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateVMInstanceMeta) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateVMInstanceMeta) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateVMInstanceMeta) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


