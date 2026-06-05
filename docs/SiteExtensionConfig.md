# SiteExtensionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The extension ID | 
**Name** | **string** | The extension name | 
**Label** | Pointer to **string** | The extension unique label | [optional] 
**Version** | **string** | Version of the extension. | 
**Enabled** | **bool** | Indicator on whether the extension is enabled or not on the site | 
**Variables** | [**map[string]ExtensionSiteConfigVariablesValue**](ExtensionSiteConfigVariablesValue.md) | Object with property names matching extension config variable names | 
**ConfigVars** | Pointer to [**[]ExtensionDefinitionConfigVarsDataItem**](ExtensionDefinitionConfigVarsDataItem.md) | List of site configuration variables. | [optional] 

## Methods

### NewSiteExtensionConfig

`func NewSiteExtensionConfig(id int64, name string, version string, enabled bool, variables map[string]ExtensionSiteConfigVariablesValue, ) *SiteExtensionConfig`

NewSiteExtensionConfig instantiates a new SiteExtensionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteExtensionConfigWithDefaults

`func NewSiteExtensionConfigWithDefaults() *SiteExtensionConfig`

NewSiteExtensionConfigWithDefaults instantiates a new SiteExtensionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SiteExtensionConfig) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SiteExtensionConfig) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SiteExtensionConfig) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *SiteExtensionConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SiteExtensionConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SiteExtensionConfig) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *SiteExtensionConfig) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SiteExtensionConfig) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SiteExtensionConfig) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *SiteExtensionConfig) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetVersion

`func (o *SiteExtensionConfig) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SiteExtensionConfig) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SiteExtensionConfig) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetEnabled

`func (o *SiteExtensionConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SiteExtensionConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SiteExtensionConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetVariables

`func (o *SiteExtensionConfig) GetVariables() map[string]ExtensionSiteConfigVariablesValue`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *SiteExtensionConfig) GetVariablesOk() (*map[string]ExtensionSiteConfigVariablesValue, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *SiteExtensionConfig) SetVariables(v map[string]ExtensionSiteConfigVariablesValue)`

SetVariables sets Variables field to given value.


### GetConfigVars

`func (o *SiteExtensionConfig) GetConfigVars() []ExtensionDefinitionConfigVarsDataItem`

GetConfigVars returns the ConfigVars field if non-nil, zero value otherwise.

### GetConfigVarsOk

`func (o *SiteExtensionConfig) GetConfigVarsOk() (*[]ExtensionDefinitionConfigVarsDataItem, bool)`

GetConfigVarsOk returns a tuple with the ConfigVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigVars

`func (o *SiteExtensionConfig) SetConfigVars(v []ExtensionDefinitionConfigVarsDataItem)`

SetConfigVars sets ConfigVars field to given value.

### HasConfigVars

`func (o *SiteExtensionConfig) HasConfigVars() bool`

HasConfigVars returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


