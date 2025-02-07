# UpdateVMInstanceGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** | Tags for the VM Instance Group. | [optional] 
**VolumeTemplateId** | Pointer to **float32** | Id of the template used by the VM Instance Group. | [optional] 
**Label** | Pointer to **string** | Label for the VM Instance Group. | [optional] 
**GuiSettings** | Pointer to [**GenericGUISettings**](GenericGUISettings.md) |  | [optional] 
**VmInstanceGroupInterfaces** | Pointer to [**[]UpdateVMInstanceGroupInterface**](UpdateVMInstanceGroupInterface.md) | Interfaces for the VM Instance Group | [optional] 
**CustomVariables** | Pointer to **map[string]interface{}** | Custom variables for the VM Instance. | [optional] 

## Methods

### NewUpdateVMInstanceGroup

`func NewUpdateVMInstanceGroup() *UpdateVMInstanceGroup`

NewUpdateVMInstanceGroup instantiates a new UpdateVMInstanceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVMInstanceGroupWithDefaults

`func NewUpdateVMInstanceGroupWithDefaults() *UpdateVMInstanceGroup`

NewUpdateVMInstanceGroupWithDefaults instantiates a new UpdateVMInstanceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *UpdateVMInstanceGroup) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateVMInstanceGroup) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateVMInstanceGroup) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateVMInstanceGroup) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVolumeTemplateId

`func (o *UpdateVMInstanceGroup) GetVolumeTemplateId() float32`

GetVolumeTemplateId returns the VolumeTemplateId field if non-nil, zero value otherwise.

### GetVolumeTemplateIdOk

`func (o *UpdateVMInstanceGroup) GetVolumeTemplateIdOk() (*float32, bool)`

GetVolumeTemplateIdOk returns a tuple with the VolumeTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeTemplateId

`func (o *UpdateVMInstanceGroup) SetVolumeTemplateId(v float32)`

SetVolumeTemplateId sets VolumeTemplateId field to given value.

### HasVolumeTemplateId

`func (o *UpdateVMInstanceGroup) HasVolumeTemplateId() bool`

HasVolumeTemplateId returns a boolean if a field has been set.

### GetLabel

`func (o *UpdateVMInstanceGroup) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateVMInstanceGroup) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateVMInstanceGroup) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpdateVMInstanceGroup) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetGuiSettings

`func (o *UpdateVMInstanceGroup) GetGuiSettings() GenericGUISettings`

GetGuiSettings returns the GuiSettings field if non-nil, zero value otherwise.

### GetGuiSettingsOk

`func (o *UpdateVMInstanceGroup) GetGuiSettingsOk() (*GenericGUISettings, bool)`

GetGuiSettingsOk returns a tuple with the GuiSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuiSettings

`func (o *UpdateVMInstanceGroup) SetGuiSettings(v GenericGUISettings)`

SetGuiSettings sets GuiSettings field to given value.

### HasGuiSettings

`func (o *UpdateVMInstanceGroup) HasGuiSettings() bool`

HasGuiSettings returns a boolean if a field has been set.

### GetVmInstanceGroupInterfaces

`func (o *UpdateVMInstanceGroup) GetVmInstanceGroupInterfaces() []UpdateVMInstanceGroupInterface`

GetVmInstanceGroupInterfaces returns the VmInstanceGroupInterfaces field if non-nil, zero value otherwise.

### GetVmInstanceGroupInterfacesOk

`func (o *UpdateVMInstanceGroup) GetVmInstanceGroupInterfacesOk() (*[]UpdateVMInstanceGroupInterface, bool)`

GetVmInstanceGroupInterfacesOk returns a tuple with the VmInstanceGroupInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmInstanceGroupInterfaces

`func (o *UpdateVMInstanceGroup) SetVmInstanceGroupInterfaces(v []UpdateVMInstanceGroupInterface)`

SetVmInstanceGroupInterfaces sets VmInstanceGroupInterfaces field to given value.

### HasVmInstanceGroupInterfaces

`func (o *UpdateVMInstanceGroup) HasVmInstanceGroupInterfaces() bool`

HasVmInstanceGroupInterfaces returns a boolean if a field has been set.

### GetCustomVariables

`func (o *UpdateVMInstanceGroup) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *UpdateVMInstanceGroup) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *UpdateVMInstanceGroup) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *UpdateVMInstanceGroup) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


