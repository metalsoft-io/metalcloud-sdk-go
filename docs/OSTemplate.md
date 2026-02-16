# OSTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The OS template ID | [readonly] 
**Name** | **string** | The OS template name | 
**Description** | Pointer to **string** | The OS template description | [optional] 
**Label** | Pointer to **string** | The OS template label. It must be unique | [optional] 
**Device** | [**OSTemplateDevice**](OSTemplateDevice.md) |  | 
**Install** | [**OSTemplateInstall**](OSTemplateInstall.md) |  | 
**ImageBuild** | Pointer to [**OSTemplateImageBuild**](OSTemplateImageBuild.md) |  | [optional] 
**ImageCertSerialNumber** | Pointer to **string** | The image boot certificate serial number associated with the OS template. Used for secure boot | [optional] 
**Os** | [**OSTemplateOs**](OSTemplateOs.md) |  | 
**Visibility** | **string** | The visibility of the OS template.                     If the visibility is PUBLIC any user can use the OS template in deployments                     If the visibility is PRIVATE the OS template can be used in deployments only                     by the user who created and/or updated the template | [default to "private"]
**Status** | **string** | The status, let the user to decide with templates to delete and when,                     and how much to keep them in the history (archived status). Also, it allows the user to                     resurrect the archived templates if needed.                     Status: READY                         - is the initial status of the template                         - the OS template is ready for deployment                         - the OS template can be deleted, use in deployments and updated                     Status: ACTIVE                         - the OS template is part of at least one ongoing deployment                         - can&#39;t be deleted (the template service will have validation for this)                         - the status can&#39;t be changed to ARCHIVED (the template service will have validation for this)                     Status: USED                         - the OS Template is part of at least one finished deployment, that is not deleted                         - can&#39;t be deleted (the template service will have validation for this)                         - can be updated, deploy or ARCHIVED                     Status: ARCHIVED                         - the OS Template is kept in the system for historical reasons                         - can&#39;t be deleted (the template service will have validation for this)                         - can&#39;t be updated or deployed                         - the status can be changed to READY or USED, if it needs to be used again or deleted | [default to "ready"]
**TemplateAssetIDs** | Pointer to **[]int32** | The template asset IDs associated with the OS template | [optional] 
**Tags** | Pointer to **[]string** | The tags associated with the OS template | [optional] 
**Revision** | **int32** | The revision number of the OS template | [readonly] 
**CreatedBy** | **int32** | The user ID of the user who created the OS template | 
**ModifiedBy** | Pointer to **int32** | The user ID of the user who last modified the OS template | [optional] 
**CreatedAt** | **string** | The date and time the OS template was created | [readonly] 
**ModifiedAt** | Pointer to **string** | The date and time the OS template was last modified | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewOSTemplate

`func NewOSTemplate(id int32, name string, device OSTemplateDevice, install OSTemplateInstall, os OSTemplateOs, visibility string, status string, revision int32, createdBy int32, createdAt string, ) *OSTemplate`

NewOSTemplate instantiates a new OSTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSTemplateWithDefaults

`func NewOSTemplateWithDefaults() *OSTemplate`

NewOSTemplateWithDefaults instantiates a new OSTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OSTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OSTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OSTemplate) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *OSTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OSTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OSTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *OSTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OSTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OSTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OSTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *OSTemplate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *OSTemplate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *OSTemplate) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *OSTemplate) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDevice

`func (o *OSTemplate) GetDevice() OSTemplateDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *OSTemplate) GetDeviceOk() (*OSTemplateDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *OSTemplate) SetDevice(v OSTemplateDevice)`

SetDevice sets Device field to given value.


### GetInstall

`func (o *OSTemplate) GetInstall() OSTemplateInstall`

GetInstall returns the Install field if non-nil, zero value otherwise.

### GetInstallOk

`func (o *OSTemplate) GetInstallOk() (*OSTemplateInstall, bool)`

GetInstallOk returns a tuple with the Install field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstall

`func (o *OSTemplate) SetInstall(v OSTemplateInstall)`

SetInstall sets Install field to given value.


### GetImageBuild

`func (o *OSTemplate) GetImageBuild() OSTemplateImageBuild`

GetImageBuild returns the ImageBuild field if non-nil, zero value otherwise.

### GetImageBuildOk

`func (o *OSTemplate) GetImageBuildOk() (*OSTemplateImageBuild, bool)`

GetImageBuildOk returns a tuple with the ImageBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBuild

`func (o *OSTemplate) SetImageBuild(v OSTemplateImageBuild)`

SetImageBuild sets ImageBuild field to given value.

### HasImageBuild

`func (o *OSTemplate) HasImageBuild() bool`

HasImageBuild returns a boolean if a field has been set.

### GetImageCertSerialNumber

`func (o *OSTemplate) GetImageCertSerialNumber() string`

GetImageCertSerialNumber returns the ImageCertSerialNumber field if non-nil, zero value otherwise.

### GetImageCertSerialNumberOk

`func (o *OSTemplate) GetImageCertSerialNumberOk() (*string, bool)`

GetImageCertSerialNumberOk returns a tuple with the ImageCertSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageCertSerialNumber

`func (o *OSTemplate) SetImageCertSerialNumber(v string)`

SetImageCertSerialNumber sets ImageCertSerialNumber field to given value.

### HasImageCertSerialNumber

`func (o *OSTemplate) HasImageCertSerialNumber() bool`

HasImageCertSerialNumber returns a boolean if a field has been set.

### GetOs

`func (o *OSTemplate) GetOs() OSTemplateOs`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *OSTemplate) GetOsOk() (*OSTemplateOs, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *OSTemplate) SetOs(v OSTemplateOs)`

SetOs sets Os field to given value.


### GetVisibility

`func (o *OSTemplate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *OSTemplate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *OSTemplate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.


### GetStatus

`func (o *OSTemplate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OSTemplate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OSTemplate) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTemplateAssetIDs

`func (o *OSTemplate) GetTemplateAssetIDs() []int32`

GetTemplateAssetIDs returns the TemplateAssetIDs field if non-nil, zero value otherwise.

### GetTemplateAssetIDsOk

`func (o *OSTemplate) GetTemplateAssetIDsOk() (*[]int32, bool)`

GetTemplateAssetIDsOk returns a tuple with the TemplateAssetIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateAssetIDs

`func (o *OSTemplate) SetTemplateAssetIDs(v []int32)`

SetTemplateAssetIDs sets TemplateAssetIDs field to given value.

### HasTemplateAssetIDs

`func (o *OSTemplate) HasTemplateAssetIDs() bool`

HasTemplateAssetIDs returns a boolean if a field has been set.

### GetTags

`func (o *OSTemplate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OSTemplate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OSTemplate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OSTemplate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetRevision

`func (o *OSTemplate) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *OSTemplate) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *OSTemplate) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetCreatedBy

`func (o *OSTemplate) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *OSTemplate) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *OSTemplate) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedBy

`func (o *OSTemplate) GetModifiedBy() int32`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *OSTemplate) GetModifiedByOk() (*int32, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *OSTemplate) SetModifiedBy(v int32)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *OSTemplate) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OSTemplate) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OSTemplate) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OSTemplate) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *OSTemplate) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *OSTemplate) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *OSTemplate) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *OSTemplate) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetLinks

`func (o *OSTemplate) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OSTemplate) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OSTemplate) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OSTemplate) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


