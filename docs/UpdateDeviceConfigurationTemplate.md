# UpdateDeviceConfigurationTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Unique machine-friendly identifier (kebab-case recommended). Immutable; used in API filters and uniqueness checks. | [optional] 
**Name** | Pointer to **string** | Human-readable display name. | [optional] 
**Description** | Pointer to **string** | Optional free-text description for the template. | [optional] 
**TemplateContent** | Pointer to **string** | Base64-encoded Jinja2 template body. Rendered server-side by Nunjucks (Jinja2-compatible syntax). | [optional] 
**CustomVariablesJson** | Pointer to **map[string]interface{}** | Template-level custom variables merged into the Jinja2 render context as defaults. Request-time &#x60;variables&#x60; (on the render endpoints) override these on key conflict. | [optional] 
**Annotations** | Pointer to **map[string]string** | Key-value annotations for storing additional metadata. | [optional] 
**Tags** | Pointer to **[]string** | Tags associated with the template. Each entry is capped at 255 characters. | [optional] 

## Methods

### NewUpdateDeviceConfigurationTemplate

`func NewUpdateDeviceConfigurationTemplate() *UpdateDeviceConfigurationTemplate`

NewUpdateDeviceConfigurationTemplate instantiates a new UpdateDeviceConfigurationTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeviceConfigurationTemplateWithDefaults

`func NewUpdateDeviceConfigurationTemplateWithDefaults() *UpdateDeviceConfigurationTemplate`

NewUpdateDeviceConfigurationTemplateWithDefaults instantiates a new UpdateDeviceConfigurationTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *UpdateDeviceConfigurationTemplate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateDeviceConfigurationTemplate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateDeviceConfigurationTemplate) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpdateDeviceConfigurationTemplate) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *UpdateDeviceConfigurationTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDeviceConfigurationTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDeviceConfigurationTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateDeviceConfigurationTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateDeviceConfigurationTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateDeviceConfigurationTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateDeviceConfigurationTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateDeviceConfigurationTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTemplateContent

`func (o *UpdateDeviceConfigurationTemplate) GetTemplateContent() string`

GetTemplateContent returns the TemplateContent field if non-nil, zero value otherwise.

### GetTemplateContentOk

`func (o *UpdateDeviceConfigurationTemplate) GetTemplateContentOk() (*string, bool)`

GetTemplateContentOk returns a tuple with the TemplateContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateContent

`func (o *UpdateDeviceConfigurationTemplate) SetTemplateContent(v string)`

SetTemplateContent sets TemplateContent field to given value.

### HasTemplateContent

`func (o *UpdateDeviceConfigurationTemplate) HasTemplateContent() bool`

HasTemplateContent returns a boolean if a field has been set.

### GetCustomVariablesJson

`func (o *UpdateDeviceConfigurationTemplate) GetCustomVariablesJson() map[string]interface{}`

GetCustomVariablesJson returns the CustomVariablesJson field if non-nil, zero value otherwise.

### GetCustomVariablesJsonOk

`func (o *UpdateDeviceConfigurationTemplate) GetCustomVariablesJsonOk() (*map[string]interface{}, bool)`

GetCustomVariablesJsonOk returns a tuple with the CustomVariablesJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariablesJson

`func (o *UpdateDeviceConfigurationTemplate) SetCustomVariablesJson(v map[string]interface{})`

SetCustomVariablesJson sets CustomVariablesJson field to given value.

### HasCustomVariablesJson

`func (o *UpdateDeviceConfigurationTemplate) HasCustomVariablesJson() bool`

HasCustomVariablesJson returns a boolean if a field has been set.

### GetAnnotations

`func (o *UpdateDeviceConfigurationTemplate) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *UpdateDeviceConfigurationTemplate) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *UpdateDeviceConfigurationTemplate) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *UpdateDeviceConfigurationTemplate) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetTags

`func (o *UpdateDeviceConfigurationTemplate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateDeviceConfigurationTemplate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateDeviceConfigurationTemplate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateDeviceConfigurationTemplate) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


