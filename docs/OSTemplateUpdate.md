# OSTemplateUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The OS template name | 
**Description** | Pointer to **string** | The OS template description | [optional] 
**Label** | Pointer to **string** | The OS template label. It must be unique | [optional] 
**Device** | [**OSTemplateDevice**](OSTemplateDevice.md) |  | 
**Install** | [**OSTemplateInstall**](OSTemplateInstall.md) |  | 
**ImageBuild** | [**OSTemplateImageBuild**](OSTemplateImageBuild.md) |  | 
**ImageCertSerialNumber** | Pointer to **string** | The image boot certificate serial number associated with the OS template. Used for secure boot | [optional] 
**Os** | [**OSTemplateOs**](OSTemplateOs.md) |  | 
**Visibility** | Pointer to **string** | The visibility of the OS template.                     If the visibility is PUBLIC any user can use the OS template in deployments                     If the visibility is PRIVATE the OS template can be used in deployments only                     by the user who created and/or updated the template | [optional] 
**Status** | Pointer to **string** | The status, let the user to decide with templates to delete and when,                     and how much to keep them in the history (archived status). Also, it allows the user to                     resurrect the archived templates if needed.                     Status: READY                         - is the initial status of the template                         - the OS template is ready for deployment                         - the OS template can be deleted, use in deployments and updated                     Status: ACTIVE                         - the OS template is part of at least one ongoing deployment                         - can&#39;t be deleted (the template service will have validation for this)                         - the status can&#39;t be changed to ARCHIVED (the template service will have validation for this)                     Status: USED                         - the OS Template is part of at least one finished deployment, that is not deleted                         - can&#39;t be deleted (the template service will have validation for this)                         - can be updated, deploy or ARCHIVED                     Status: ARCHIVED                         - the OS Template is kept in the system for historical reasons                         - can&#39;t be deleted (the template service will have validation for this)                         - can&#39;t be updated or deployed                         - the status can be changed to READY or USED, if it needs to be used again or deleted | [optional] 
**Tags** | Pointer to **[]string** | The tags associated with the OS template | [optional] 

## Methods

### NewOSTemplateUpdate

`func NewOSTemplateUpdate(name string, device OSTemplateDevice, install OSTemplateInstall, imageBuild OSTemplateImageBuild, os OSTemplateOs, ) *OSTemplateUpdate`

NewOSTemplateUpdate instantiates a new OSTemplateUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSTemplateUpdateWithDefaults

`func NewOSTemplateUpdateWithDefaults() *OSTemplateUpdate`

NewOSTemplateUpdateWithDefaults instantiates a new OSTemplateUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OSTemplateUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OSTemplateUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OSTemplateUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *OSTemplateUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OSTemplateUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OSTemplateUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OSTemplateUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *OSTemplateUpdate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *OSTemplateUpdate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *OSTemplateUpdate) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *OSTemplateUpdate) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDevice

`func (o *OSTemplateUpdate) GetDevice() OSTemplateDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *OSTemplateUpdate) GetDeviceOk() (*OSTemplateDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *OSTemplateUpdate) SetDevice(v OSTemplateDevice)`

SetDevice sets Device field to given value.


### GetInstall

`func (o *OSTemplateUpdate) GetInstall() OSTemplateInstall`

GetInstall returns the Install field if non-nil, zero value otherwise.

### GetInstallOk

`func (o *OSTemplateUpdate) GetInstallOk() (*OSTemplateInstall, bool)`

GetInstallOk returns a tuple with the Install field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstall

`func (o *OSTemplateUpdate) SetInstall(v OSTemplateInstall)`

SetInstall sets Install field to given value.


### GetImageBuild

`func (o *OSTemplateUpdate) GetImageBuild() OSTemplateImageBuild`

GetImageBuild returns the ImageBuild field if non-nil, zero value otherwise.

### GetImageBuildOk

`func (o *OSTemplateUpdate) GetImageBuildOk() (*OSTemplateImageBuild, bool)`

GetImageBuildOk returns a tuple with the ImageBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBuild

`func (o *OSTemplateUpdate) SetImageBuild(v OSTemplateImageBuild)`

SetImageBuild sets ImageBuild field to given value.


### GetImageCertSerialNumber

`func (o *OSTemplateUpdate) GetImageCertSerialNumber() string`

GetImageCertSerialNumber returns the ImageCertSerialNumber field if non-nil, zero value otherwise.

### GetImageCertSerialNumberOk

`func (o *OSTemplateUpdate) GetImageCertSerialNumberOk() (*string, bool)`

GetImageCertSerialNumberOk returns a tuple with the ImageCertSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageCertSerialNumber

`func (o *OSTemplateUpdate) SetImageCertSerialNumber(v string)`

SetImageCertSerialNumber sets ImageCertSerialNumber field to given value.

### HasImageCertSerialNumber

`func (o *OSTemplateUpdate) HasImageCertSerialNumber() bool`

HasImageCertSerialNumber returns a boolean if a field has been set.

### GetOs

`func (o *OSTemplateUpdate) GetOs() OSTemplateOs`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *OSTemplateUpdate) GetOsOk() (*OSTemplateOs, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *OSTemplateUpdate) SetOs(v OSTemplateOs)`

SetOs sets Os field to given value.


### GetVisibility

`func (o *OSTemplateUpdate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *OSTemplateUpdate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *OSTemplateUpdate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *OSTemplateUpdate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetStatus

`func (o *OSTemplateUpdate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OSTemplateUpdate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OSTemplateUpdate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OSTemplateUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *OSTemplateUpdate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OSTemplateUpdate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OSTemplateUpdate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OSTemplateUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


