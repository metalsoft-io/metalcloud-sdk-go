# TemplateAssetCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateId** | **float32** | The ID of the OS template that this template asset belongs to | 
**Usage** | **string** | The template asset usage | 
**File** | [**TemplateAssetFile**](TemplateAssetFile.md) |  | 
**Tags** | Pointer to **[]string** | The tags associated with the template asset | [optional] 

## Methods

### NewTemplateAssetCreateDto

`func NewTemplateAssetCreateDto(templateId float32, usage string, file TemplateAssetFile, ) *TemplateAssetCreateDto`

NewTemplateAssetCreateDto instantiates a new TemplateAssetCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateAssetCreateDtoWithDefaults

`func NewTemplateAssetCreateDtoWithDefaults() *TemplateAssetCreateDto`

NewTemplateAssetCreateDtoWithDefaults instantiates a new TemplateAssetCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateId

`func (o *TemplateAssetCreateDto) GetTemplateId() float32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *TemplateAssetCreateDto) GetTemplateIdOk() (*float32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *TemplateAssetCreateDto) SetTemplateId(v float32)`

SetTemplateId sets TemplateId field to given value.


### GetUsage

`func (o *TemplateAssetCreateDto) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *TemplateAssetCreateDto) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *TemplateAssetCreateDto) SetUsage(v string)`

SetUsage sets Usage field to given value.


### GetFile

`func (o *TemplateAssetCreateDto) GetFile() TemplateAssetFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *TemplateAssetCreateDto) GetFileOk() (*TemplateAssetFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *TemplateAssetCreateDto) SetFile(v TemplateAssetFile)`

SetFile sets File field to given value.


### GetTags

`func (o *TemplateAssetCreateDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TemplateAssetCreateDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TemplateAssetCreateDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TemplateAssetCreateDto) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


