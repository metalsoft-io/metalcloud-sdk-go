# RenderSavedDeviceConfigurationTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Variables** | Pointer to **map[string]interface{}** | Render-time variables. Merged into the context after the template-level &#x60;customVariablesJson&#x60; defaults, so these keys win on key conflict. | [optional] 
**Debug** | Pointer to **bool** | When true, the response additionally includes the decoded (non-base64) template source and the full variable bag passed to the engine. | [optional] [default to false]

## Methods

### NewRenderSavedDeviceConfigurationTemplate

`func NewRenderSavedDeviceConfigurationTemplate() *RenderSavedDeviceConfigurationTemplate`

NewRenderSavedDeviceConfigurationTemplate instantiates a new RenderSavedDeviceConfigurationTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenderSavedDeviceConfigurationTemplateWithDefaults

`func NewRenderSavedDeviceConfigurationTemplateWithDefaults() *RenderSavedDeviceConfigurationTemplate`

NewRenderSavedDeviceConfigurationTemplateWithDefaults instantiates a new RenderSavedDeviceConfigurationTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariables

`func (o *RenderSavedDeviceConfigurationTemplate) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *RenderSavedDeviceConfigurationTemplate) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *RenderSavedDeviceConfigurationTemplate) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *RenderSavedDeviceConfigurationTemplate) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetDebug

`func (o *RenderSavedDeviceConfigurationTemplate) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *RenderSavedDeviceConfigurationTemplate) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *RenderSavedDeviceConfigurationTemplate) SetDebug(v bool)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *RenderSavedDeviceConfigurationTemplate) HasDebug() bool`

HasDebug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


