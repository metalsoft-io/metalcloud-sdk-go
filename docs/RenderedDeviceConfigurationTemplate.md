# RenderedDeviceConfigurationTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rendered** | **string** | Rendered template output as UTF-8 plain text (**not** base64-encoded — the response is the literal rendered string). | 
**TemplateContent** | Pointer to **string** | Debug-only: the decoded (non-base64) template source that was rendered. Present only when the request set &#x60;debug: true&#x60;. | [optional] 
**Variables** | Pointer to **map[string]interface{}** | Debug-only: the full variable bag passed to the Jinja2 engine, after merging all precedence layers. Present only when the request set &#x60;debug: true&#x60;. | [optional] 

## Methods

### NewRenderedDeviceConfigurationTemplate

`func NewRenderedDeviceConfigurationTemplate(rendered string, ) *RenderedDeviceConfigurationTemplate`

NewRenderedDeviceConfigurationTemplate instantiates a new RenderedDeviceConfigurationTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenderedDeviceConfigurationTemplateWithDefaults

`func NewRenderedDeviceConfigurationTemplateWithDefaults() *RenderedDeviceConfigurationTemplate`

NewRenderedDeviceConfigurationTemplateWithDefaults instantiates a new RenderedDeviceConfigurationTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRendered

`func (o *RenderedDeviceConfigurationTemplate) GetRendered() string`

GetRendered returns the Rendered field if non-nil, zero value otherwise.

### GetRenderedOk

`func (o *RenderedDeviceConfigurationTemplate) GetRenderedOk() (*string, bool)`

GetRenderedOk returns a tuple with the Rendered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRendered

`func (o *RenderedDeviceConfigurationTemplate) SetRendered(v string)`

SetRendered sets Rendered field to given value.


### GetTemplateContent

`func (o *RenderedDeviceConfigurationTemplate) GetTemplateContent() string`

GetTemplateContent returns the TemplateContent field if non-nil, zero value otherwise.

### GetTemplateContentOk

`func (o *RenderedDeviceConfigurationTemplate) GetTemplateContentOk() (*string, bool)`

GetTemplateContentOk returns a tuple with the TemplateContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateContent

`func (o *RenderedDeviceConfigurationTemplate) SetTemplateContent(v string)`

SetTemplateContent sets TemplateContent field to given value.

### HasTemplateContent

`func (o *RenderedDeviceConfigurationTemplate) HasTemplateContent() bool`

HasTemplateContent returns a boolean if a field has been set.

### GetVariables

`func (o *RenderedDeviceConfigurationTemplate) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *RenderedDeviceConfigurationTemplate) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *RenderedDeviceConfigurationTemplate) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *RenderedDeviceConfigurationTemplate) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


