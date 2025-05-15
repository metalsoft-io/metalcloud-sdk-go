# UpdateInfrastructureMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GuiSettings** | Pointer to [**GenericGUISettings**](GenericGUISettings.md) |  | [optional] 
**Tags** | Pointer to **[]string** | Tags for the Infrastructure. | [optional] 
**Name** | **string** | name of the Infrastructure | 

## Methods

### NewUpdateInfrastructureMeta

`func NewUpdateInfrastructureMeta(name string, ) *UpdateInfrastructureMeta`

NewUpdateInfrastructureMeta instantiates a new UpdateInfrastructureMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInfrastructureMetaWithDefaults

`func NewUpdateInfrastructureMetaWithDefaults() *UpdateInfrastructureMeta`

NewUpdateInfrastructureMetaWithDefaults instantiates a new UpdateInfrastructureMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuiSettings

`func (o *UpdateInfrastructureMeta) GetGuiSettings() GenericGUISettings`

GetGuiSettings returns the GuiSettings field if non-nil, zero value otherwise.

### GetGuiSettingsOk

`func (o *UpdateInfrastructureMeta) GetGuiSettingsOk() (*GenericGUISettings, bool)`

GetGuiSettingsOk returns a tuple with the GuiSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuiSettings

`func (o *UpdateInfrastructureMeta) SetGuiSettings(v GenericGUISettings)`

SetGuiSettings sets GuiSettings field to given value.

### HasGuiSettings

`func (o *UpdateInfrastructureMeta) HasGuiSettings() bool`

HasGuiSettings returns a boolean if a field has been set.

### GetTags

`func (o *UpdateInfrastructureMeta) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateInfrastructureMeta) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateInfrastructureMeta) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateInfrastructureMeta) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetName

`func (o *UpdateInfrastructureMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateInfrastructureMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateInfrastructureMeta) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


