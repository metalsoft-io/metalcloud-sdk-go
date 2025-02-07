# OSTemplateImageBuild

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Required** | **bool** | Indicates whether the OS template requires an image build | [default to true]
**Provider** | Pointer to **string** | The OS template image build provider | [optional] 

## Methods

### NewOSTemplateImageBuild

`func NewOSTemplateImageBuild(required bool, ) *OSTemplateImageBuild`

NewOSTemplateImageBuild instantiates a new OSTemplateImageBuild object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSTemplateImageBuildWithDefaults

`func NewOSTemplateImageBuildWithDefaults() *OSTemplateImageBuild`

NewOSTemplateImageBuildWithDefaults instantiates a new OSTemplateImageBuild object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequired

`func (o *OSTemplateImageBuild) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *OSTemplateImageBuild) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *OSTemplateImageBuild) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetProvider

`func (o *OSTemplateImageBuild) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *OSTemplateImageBuild) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *OSTemplateImageBuild) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *OSTemplateImageBuild) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


