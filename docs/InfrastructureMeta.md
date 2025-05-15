# InfrastructureMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GuiSettings** | Pointer to [**GenericGUISettings**](GenericGUISettings.md) |  | [optional] 
**Tags** | Pointer to **[]string** | Tags for the Infrastructure. | [optional] 
**Name** | **string** | name of the Infrastructure | 

## Methods

### NewInfrastructureMeta

`func NewInfrastructureMeta(name string, ) *InfrastructureMeta`

NewInfrastructureMeta instantiates a new InfrastructureMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureMetaWithDefaults

`func NewInfrastructureMetaWithDefaults() *InfrastructureMeta`

NewInfrastructureMetaWithDefaults instantiates a new InfrastructureMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuiSettings

`func (o *InfrastructureMeta) GetGuiSettings() GenericGUISettings`

GetGuiSettings returns the GuiSettings field if non-nil, zero value otherwise.

### GetGuiSettingsOk

`func (o *InfrastructureMeta) GetGuiSettingsOk() (*GenericGUISettings, bool)`

GetGuiSettingsOk returns a tuple with the GuiSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuiSettings

`func (o *InfrastructureMeta) SetGuiSettings(v GenericGUISettings)`

SetGuiSettings sets GuiSettings field to given value.

### HasGuiSettings

`func (o *InfrastructureMeta) HasGuiSettings() bool`

HasGuiSettings returns a boolean if a field has been set.

### GetTags

`func (o *InfrastructureMeta) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InfrastructureMeta) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InfrastructureMeta) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InfrastructureMeta) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetName

`func (o *InfrastructureMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InfrastructureMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InfrastructureMeta) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


