# TemplateAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | The template asset ID | [readonly] 
**TemplateId** | **float32** | The ID of the OS template that this template asset belongs to | 
**Usage** | **string** | The template asset usage | 
**File** | [**TemplateAssetFile**](TemplateAssetFile.md) |  | 
**Tags** | Pointer to **[]string** | The tags associated with the template asset | [optional] 
**Revision** | **float32** | The revision number of the template asset | [readonly] 
**CreatedBy** | **float32** | The user ID of the user who created the template asset | 
**ModifiedBy** | Pointer to **float32** | The user ID of the user who last modified the template asset | [optional] 
**CreatedAt** | **time.Time** | The date and time the template asset was created | [readonly] 
**ModifiedAt** | Pointer to **time.Time** | The date and time the template asset was last modified | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewTemplateAsset

`func NewTemplateAsset(id float32, templateId float32, usage string, file TemplateAssetFile, revision float32, createdBy float32, createdAt time.Time, ) *TemplateAsset`

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

`func (o *TemplateAsset) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateAsset) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateAsset) SetId(v float32)`

SetId sets Id field to given value.


### GetTemplateId

`func (o *TemplateAsset) GetTemplateId() float32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *TemplateAsset) GetTemplateIdOk() (*float32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *TemplateAsset) SetTemplateId(v float32)`

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

`func (o *TemplateAsset) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *TemplateAsset) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *TemplateAsset) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetCreatedBy

`func (o *TemplateAsset) GetCreatedBy() float32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *TemplateAsset) GetCreatedByOk() (*float32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *TemplateAsset) SetCreatedBy(v float32)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedBy

`func (o *TemplateAsset) GetModifiedBy() float32`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *TemplateAsset) GetModifiedByOk() (*float32, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *TemplateAsset) SetModifiedBy(v float32)`

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


