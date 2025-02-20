# TemplateAssetCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateId** | **float32** | The ID of the OS template that this template asset belongs to | 
**Usage** | **string** | The template asset usage | 
**File** | [**TemplateAssetFile**](TemplateAssetFile.md) |  | 
**Tags** | Pointer to **[]string** | The tags associated with the template asset | [optional] 

## Methods

### NewTemplateAssetCreate

`func NewTemplateAssetCreate(templateId float32, usage string, file TemplateAssetFile, ) *TemplateAssetCreate`

NewTemplateAssetCreate instantiates a new TemplateAssetCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateAssetCreateWithDefaults

`func NewTemplateAssetCreateWithDefaults() *TemplateAssetCreate`

NewTemplateAssetCreateWithDefaults instantiates a new TemplateAssetCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateId

`func (o *TemplateAssetCreate) GetTemplateId() float32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *TemplateAssetCreate) GetTemplateIdOk() (*float32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *TemplateAssetCreate) SetTemplateId(v float32)`

SetTemplateId sets TemplateId field to given value.


### GetUsage

`func (o *TemplateAssetCreate) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *TemplateAssetCreate) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *TemplateAssetCreate) SetUsage(v string)`

SetUsage sets Usage field to given value.


### GetFile

`func (o *TemplateAssetCreate) GetFile() TemplateAssetFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *TemplateAssetCreate) GetFileOk() (*TemplateAssetFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *TemplateAssetCreate) SetFile(v TemplateAssetFile)`

SetFile sets File field to given value.


### GetTags

`func (o *TemplateAssetCreate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TemplateAssetCreate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TemplateAssetCreate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TemplateAssetCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


