# ExtensionSiteConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The site ID | 
**Slug** | Pointer to **string** | The site unique slug | [optional] 
**Name** | Pointer to **string** | The site name | [optional] 
**Variables** | [**map[string]ExtensionSiteConfigVariablesValue**](ExtensionSiteConfigVariablesValue.md) | Object with property names matching extension config variable names | 

## Methods

### NewExtensionSiteConfig

`func NewExtensionSiteConfig(id int64, variables map[string]ExtensionSiteConfigVariablesValue, ) *ExtensionSiteConfig`

NewExtensionSiteConfig instantiates a new ExtensionSiteConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionSiteConfigWithDefaults

`func NewExtensionSiteConfigWithDefaults() *ExtensionSiteConfig`

NewExtensionSiteConfigWithDefaults instantiates a new ExtensionSiteConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExtensionSiteConfig) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExtensionSiteConfig) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExtensionSiteConfig) SetId(v int64)`

SetId sets Id field to given value.


### GetSlug

`func (o *ExtensionSiteConfig) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ExtensionSiteConfig) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ExtensionSiteConfig) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ExtensionSiteConfig) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetName

`func (o *ExtensionSiteConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtensionSiteConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtensionSiteConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExtensionSiteConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVariables

`func (o *ExtensionSiteConfig) GetVariables() map[string]ExtensionSiteConfigVariablesValue`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ExtensionSiteConfig) GetVariablesOk() (*map[string]ExtensionSiteConfigVariablesValue, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ExtensionSiteConfig) SetVariables(v map[string]ExtensionSiteConfigVariablesValue)`

SetVariables sets Variables field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


