# OSTemplateCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The OS template name | 
**Description** | Pointer to **string** | The OS template description | [optional] 
**Label** | Pointer to **string** | The OS template label | [optional] 
**Device** | [**OSTemplateDevice**](OSTemplateDevice.md) |  | 
**Install** | [**OSTemplateInstall**](OSTemplateInstall.md) |  | 
**ImageBuild** | [**OSTemplateImageBuild**](OSTemplateImageBuild.md) |  | 
**Os** | [**OSTemplateOs**](OSTemplateOs.md) |  | 
**Visibility** | Pointer to **string** | The visibility of the OS template.                     If the visibility is PUBLIC any user can use the OS template in deployments                     If the visibility is PRIVATE the OS template can be used in deployments only                     by the user who created and/or updated the template | [optional] [default to "private"]
**Tags** | Pointer to **[]string** | The tags associated with the OS template | [optional] 

## Methods

### NewOSTemplateCreate

`func NewOSTemplateCreate(name string, device OSTemplateDevice, install OSTemplateInstall, imageBuild OSTemplateImageBuild, os OSTemplateOs, ) *OSTemplateCreate`

NewOSTemplateCreate instantiates a new OSTemplateCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSTemplateCreateWithDefaults

`func NewOSTemplateCreateWithDefaults() *OSTemplateCreate`

NewOSTemplateCreateWithDefaults instantiates a new OSTemplateCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OSTemplateCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OSTemplateCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OSTemplateCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *OSTemplateCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OSTemplateCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OSTemplateCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OSTemplateCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *OSTemplateCreate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *OSTemplateCreate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *OSTemplateCreate) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *OSTemplateCreate) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDevice

`func (o *OSTemplateCreate) GetDevice() OSTemplateDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *OSTemplateCreate) GetDeviceOk() (*OSTemplateDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *OSTemplateCreate) SetDevice(v OSTemplateDevice)`

SetDevice sets Device field to given value.


### GetInstall

`func (o *OSTemplateCreate) GetInstall() OSTemplateInstall`

GetInstall returns the Install field if non-nil, zero value otherwise.

### GetInstallOk

`func (o *OSTemplateCreate) GetInstallOk() (*OSTemplateInstall, bool)`

GetInstallOk returns a tuple with the Install field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstall

`func (o *OSTemplateCreate) SetInstall(v OSTemplateInstall)`

SetInstall sets Install field to given value.


### GetImageBuild

`func (o *OSTemplateCreate) GetImageBuild() OSTemplateImageBuild`

GetImageBuild returns the ImageBuild field if non-nil, zero value otherwise.

### GetImageBuildOk

`func (o *OSTemplateCreate) GetImageBuildOk() (*OSTemplateImageBuild, bool)`

GetImageBuildOk returns a tuple with the ImageBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBuild

`func (o *OSTemplateCreate) SetImageBuild(v OSTemplateImageBuild)`

SetImageBuild sets ImageBuild field to given value.


### GetOs

`func (o *OSTemplateCreate) GetOs() OSTemplateOs`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *OSTemplateCreate) GetOsOk() (*OSTemplateOs, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *OSTemplateCreate) SetOs(v OSTemplateOs)`

SetOs sets Os field to given value.


### GetVisibility

`func (o *OSTemplateCreate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *OSTemplateCreate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *OSTemplateCreate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *OSTemplateCreate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTags

`func (o *OSTemplateCreate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OSTemplateCreate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OSTemplateCreate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OSTemplateCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


