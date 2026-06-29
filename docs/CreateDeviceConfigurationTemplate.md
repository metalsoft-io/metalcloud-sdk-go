# CreateDeviceConfigurationTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Unique machine-friendly identifier (kebab-case recommended). Immutable; used in API filters and uniqueness checks. | 
**Name** | Pointer to **string** | Human-readable display name. | [optional] 
**Description** | Pointer to **string** | Optional free-text description for the template. | [optional] 
**DeviceDriver** | [**SwitchDriver**](SwitchDriver.md) | The device driver this template applies to. Used for filtering and validation purposes. | 
**ExecutionType** | [**NetworkTemplateExecutionType**](NetworkTemplateExecutionType.md) | The execution type for the template. | 
**TemplateContent** | Pointer to **string** | Base64-encoded Jinja2 template body. Rendered server-side by Nunjucks (Jinja2-compatible syntax). | [optional] 
**CustomVariablesJson** | Pointer to **map[string]interface{}** | Template-level custom variables merged into the Jinja2 render context as defaults. Request-time &#x60;variables&#x60; (on the render endpoints) override these on key conflict. | [optional] 
**Annotations** | Pointer to **map[string]string** | Key-value annotations for storing additional metadata. | [optional] 
**Tags** | Pointer to **[]string** | Tags associated with the template. Each entry is capped at 255 characters. | [optional] 

## Methods

### NewCreateDeviceConfigurationTemplate

`func NewCreateDeviceConfigurationTemplate(label string, deviceDriver SwitchDriver, executionType NetworkTemplateExecutionType, ) *CreateDeviceConfigurationTemplate`

NewCreateDeviceConfigurationTemplate instantiates a new CreateDeviceConfigurationTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeviceConfigurationTemplateWithDefaults

`func NewCreateDeviceConfigurationTemplateWithDefaults() *CreateDeviceConfigurationTemplate`

NewCreateDeviceConfigurationTemplateWithDefaults instantiates a new CreateDeviceConfigurationTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CreateDeviceConfigurationTemplate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateDeviceConfigurationTemplate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateDeviceConfigurationTemplate) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *CreateDeviceConfigurationTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDeviceConfigurationTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDeviceConfigurationTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateDeviceConfigurationTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateDeviceConfigurationTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateDeviceConfigurationTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateDeviceConfigurationTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateDeviceConfigurationTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeviceDriver

`func (o *CreateDeviceConfigurationTemplate) GetDeviceDriver() SwitchDriver`

GetDeviceDriver returns the DeviceDriver field if non-nil, zero value otherwise.

### GetDeviceDriverOk

`func (o *CreateDeviceConfigurationTemplate) GetDeviceDriverOk() (*SwitchDriver, bool)`

GetDeviceDriverOk returns a tuple with the DeviceDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDriver

`func (o *CreateDeviceConfigurationTemplate) SetDeviceDriver(v SwitchDriver)`

SetDeviceDriver sets DeviceDriver field to given value.


### GetExecutionType

`func (o *CreateDeviceConfigurationTemplate) GetExecutionType() NetworkTemplateExecutionType`

GetExecutionType returns the ExecutionType field if non-nil, zero value otherwise.

### GetExecutionTypeOk

`func (o *CreateDeviceConfigurationTemplate) GetExecutionTypeOk() (*NetworkTemplateExecutionType, bool)`

GetExecutionTypeOk returns a tuple with the ExecutionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionType

`func (o *CreateDeviceConfigurationTemplate) SetExecutionType(v NetworkTemplateExecutionType)`

SetExecutionType sets ExecutionType field to given value.


### GetTemplateContent

`func (o *CreateDeviceConfigurationTemplate) GetTemplateContent() string`

GetTemplateContent returns the TemplateContent field if non-nil, zero value otherwise.

### GetTemplateContentOk

`func (o *CreateDeviceConfigurationTemplate) GetTemplateContentOk() (*string, bool)`

GetTemplateContentOk returns a tuple with the TemplateContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateContent

`func (o *CreateDeviceConfigurationTemplate) SetTemplateContent(v string)`

SetTemplateContent sets TemplateContent field to given value.

### HasTemplateContent

`func (o *CreateDeviceConfigurationTemplate) HasTemplateContent() bool`

HasTemplateContent returns a boolean if a field has been set.

### GetCustomVariablesJson

`func (o *CreateDeviceConfigurationTemplate) GetCustomVariablesJson() map[string]interface{}`

GetCustomVariablesJson returns the CustomVariablesJson field if non-nil, zero value otherwise.

### GetCustomVariablesJsonOk

`func (o *CreateDeviceConfigurationTemplate) GetCustomVariablesJsonOk() (*map[string]interface{}, bool)`

GetCustomVariablesJsonOk returns a tuple with the CustomVariablesJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariablesJson

`func (o *CreateDeviceConfigurationTemplate) SetCustomVariablesJson(v map[string]interface{})`

SetCustomVariablesJson sets CustomVariablesJson field to given value.

### HasCustomVariablesJson

`func (o *CreateDeviceConfigurationTemplate) HasCustomVariablesJson() bool`

HasCustomVariablesJson returns a boolean if a field has been set.

### GetAnnotations

`func (o *CreateDeviceConfigurationTemplate) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *CreateDeviceConfigurationTemplate) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *CreateDeviceConfigurationTemplate) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *CreateDeviceConfigurationTemplate) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetTags

`func (o *CreateDeviceConfigurationTemplate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateDeviceConfigurationTemplate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateDeviceConfigurationTemplate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateDeviceConfigurationTemplate) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


