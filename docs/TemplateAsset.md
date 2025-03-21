# TemplateAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The template asset ID | [readonly] 
**TemplateId** | **int32** | The ID of the OS template that this template asset belongs to | 
**Usage** | **string** | The template asset usage:         - build_source_image: The template asset represents the source image, used as the base image for constructing the install image.         - build_component: The template asset represents the file that will be part of the install image.         - switch_ztp_config: The template asset represents the switch ZTP configuration.         - metadata_source_image: The template asset contains data about the install image.         - generic: The template asset is used for generic purposes. | 
**File** | [**TemplateAssetFile**](TemplateAssetFile.md) |  | 
**Tags** | Pointer to **[]string** | The tags associated with the template asset | [optional] 
**Revision** | **int32** | The revision number of the template asset | [readonly] 
**CreatedBy** | **int32** | The user ID of the user who created the template asset | 
**ModifiedBy** | Pointer to **int32** | The user ID of the user who last modified the template asset | [optional] 
**CreatedAt** | **time.Time** | The date and time the template asset was created | [readonly] 
**ModifiedAt** | Pointer to **time.Time** | The date and time the template asset was last modified | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewTemplateAsset

`func NewTemplateAsset(id int32, templateId int32, usage string, file TemplateAssetFile, revision int32, createdBy int32, createdAt time.Time, ) *TemplateAsset`

NewTemplateAsset instantiates a new TemplateAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateAssetWithDefaults

`func NewTemplateAssetWithDefaults() *TemplateAsset`

NewTemplateAssetWithDefaults instantiates a new TemplateAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplateAsset) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateAsset) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateAsset) SetId(v int32)`

SetId sets Id field to given value.


### GetTemplateId

`func (o *TemplateAsset) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *TemplateAsset) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *TemplateAsset) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.


### GetUsage

`func (o *TemplateAsset) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *TemplateAsset) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *TemplateAsset) SetUsage(v string)`

SetUsage sets Usage field to given value.


### GetFile

`func (o *TemplateAsset) GetFile() TemplateAssetFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *TemplateAsset) GetFileOk() (*TemplateAssetFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *TemplateAsset) SetFile(v TemplateAssetFile)`

SetFile sets File field to given value.


### GetTags

`func (o *TemplateAsset) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TemplateAsset) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TemplateAsset) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TemplateAsset) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetRevision

`func (o *TemplateAsset) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *TemplateAsset) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *TemplateAsset) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetCreatedBy

`func (o *TemplateAsset) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *TemplateAsset) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *TemplateAsset) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedBy

`func (o *TemplateAsset) GetModifiedBy() int32`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *TemplateAsset) GetModifiedByOk() (*int32, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *TemplateAsset) SetModifiedBy(v int32)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *TemplateAsset) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TemplateAsset) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TemplateAsset) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TemplateAsset) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *TemplateAsset) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *TemplateAsset) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *TemplateAsset) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *TemplateAsset) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetLinks

`func (o *TemplateAsset) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TemplateAsset) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TemplateAsset) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TemplateAsset) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


