# RenderDeviceConfigurationTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateContent** | **string** | Base64-encoded Jinja2 template body to render. Rendered server-side by Nunjucks (Jinja2-compatible syntax). | 
**Variables** | Pointer to **map[string]interface{}** | Render context passed to the Jinja2 template. Keys must match &#x60;{{ varName }}&#x60; references in the template body. | [optional] 
**Debug** | Pointer to **bool** | When true, the response additionally includes the decoded (non-base64) template source and the full variable bag passed to the engine. | [optional] [default to false]

## Methods

### NewRenderDeviceConfigurationTemplate

`func NewRenderDeviceConfigurationTemplate(templateContent string, ) *RenderDeviceConfigurationTemplate`

NewRenderDeviceConfigurationTemplate instantiates a new RenderDeviceConfigurationTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenderDeviceConfigurationTemplateWithDefaults

`func NewRenderDeviceConfigurationTemplateWithDefaults() *RenderDeviceConfigurationTemplate`

NewRenderDeviceConfigurationTemplateWithDefaults instantiates a new RenderDeviceConfigurationTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateContent

`func (o *RenderDeviceConfigurationTemplate) GetTemplateContent() string`

GetTemplateContent returns the TemplateContent field if non-nil, zero value otherwise.

### GetTemplateContentOk

`func (o *RenderDeviceConfigurationTemplate) GetTemplateContentOk() (*string, bool)`

GetTemplateContentOk returns a tuple with the TemplateContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateContent

`func (o *RenderDeviceConfigurationTemplate) SetTemplateContent(v string)`

SetTemplateContent sets TemplateContent field to given value.


### GetVariables

`func (o *RenderDeviceConfigurationTemplate) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *RenderDeviceConfigurationTemplate) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *RenderDeviceConfigurationTemplate) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *RenderDeviceConfigurationTemplate) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetDebug

`func (o *RenderDeviceConfigurationTemplate) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *RenderDeviceConfigurationTemplate) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *RenderDeviceConfigurationTemplate) SetDebug(v bool)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *RenderDeviceConfigurationTemplate) HasDebug() bool`

HasDebug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


