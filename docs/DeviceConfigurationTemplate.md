# DeviceConfigurationTemplate

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
**CreatedTimestamp** | **time.Time** | Entity creation timestamp | 
**UpdatedTimestamp** | **time.Time** | Entity last update timestamp | 
**Revision** | **string** | Revision number of the entity | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 
**Id** | **int64** | Id | 

## Methods

### NewDeviceConfigurationTemplate

`func NewDeviceConfigurationTemplate(label string, deviceDriver SwitchDriver, executionType NetworkTemplateExecutionType, createdTimestamp time.Time, updatedTimestamp time.Time, revision string, id int64, ) *DeviceConfigurationTemplate`

NewDeviceConfigurationTemplate instantiates a new DeviceConfigurationTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceConfigurationTemplateWithDefaults

`func NewDeviceConfigurationTemplateWithDefaults() *DeviceConfigurationTemplate`

NewDeviceConfigurationTemplateWithDefaults instantiates a new DeviceConfigurationTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *DeviceConfigurationTemplate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DeviceConfigurationTemplate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DeviceConfigurationTemplate) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *DeviceConfigurationTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceConfigurationTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceConfigurationTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceConfigurationTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *DeviceConfigurationTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeviceConfigurationTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeviceConfigurationTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeviceConfigurationTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeviceDriver

`func (o *DeviceConfigurationTemplate) GetDeviceDriver() SwitchDriver`

GetDeviceDriver returns the DeviceDriver field if non-nil, zero value otherwise.

### GetDeviceDriverOk

`func (o *DeviceConfigurationTemplate) GetDeviceDriverOk() (*SwitchDriver, bool)`

GetDeviceDriverOk returns a tuple with the DeviceDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDriver

`func (o *DeviceConfigurationTemplate) SetDeviceDriver(v SwitchDriver)`

SetDeviceDriver sets DeviceDriver field to given value.


### GetExecutionType

`func (o *DeviceConfigurationTemplate) GetExecutionType() NetworkTemplateExecutionType`

GetExecutionType returns the ExecutionType field if non-nil, zero value otherwise.

### GetExecutionTypeOk

`func (o *DeviceConfigurationTemplate) GetExecutionTypeOk() (*NetworkTemplateExecutionType, bool)`

GetExecutionTypeOk returns a tuple with the ExecutionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionType

`func (o *DeviceConfigurationTemplate) SetExecutionType(v NetworkTemplateExecutionType)`

SetExecutionType sets ExecutionType field to given value.


### GetTemplateContent

`func (o *DeviceConfigurationTemplate) GetTemplateContent() string`

GetTemplateContent returns the TemplateContent field if non-nil, zero value otherwise.

### GetTemplateContentOk

`func (o *DeviceConfigurationTemplate) GetTemplateContentOk() (*string, bool)`

GetTemplateContentOk returns a tuple with the TemplateContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateContent

`func (o *DeviceConfigurationTemplate) SetTemplateContent(v string)`

SetTemplateContent sets TemplateContent field to given value.

### HasTemplateContent

`func (o *DeviceConfigurationTemplate) HasTemplateContent() bool`

HasTemplateContent returns a boolean if a field has been set.

### GetCustomVariablesJson

`func (o *DeviceConfigurationTemplate) GetCustomVariablesJson() map[string]interface{}`

GetCustomVariablesJson returns the CustomVariablesJson field if non-nil, zero value otherwise.

### GetCustomVariablesJsonOk

`func (o *DeviceConfigurationTemplate) GetCustomVariablesJsonOk() (*map[string]interface{}, bool)`

GetCustomVariablesJsonOk returns a tuple with the CustomVariablesJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariablesJson

`func (o *DeviceConfigurationTemplate) SetCustomVariablesJson(v map[string]interface{})`

SetCustomVariablesJson sets CustomVariablesJson field to given value.

### HasCustomVariablesJson

`func (o *DeviceConfigurationTemplate) HasCustomVariablesJson() bool`

HasCustomVariablesJson returns a boolean if a field has been set.

### GetAnnotations

`func (o *DeviceConfigurationTemplate) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *DeviceConfigurationTemplate) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *DeviceConfigurationTemplate) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *DeviceConfigurationTemplate) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetTags

`func (o *DeviceConfigurationTemplate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DeviceConfigurationTemplate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DeviceConfigurationTemplate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DeviceConfigurationTemplate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *DeviceConfigurationTemplate) GetCreatedTimestamp() time.Time`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *DeviceConfigurationTemplate) GetCreatedTimestampOk() (*time.Time, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *DeviceConfigurationTemplate) SetCreatedTimestamp(v time.Time)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *DeviceConfigurationTemplate) GetUpdatedTimestamp() time.Time`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *DeviceConfigurationTemplate) GetUpdatedTimestampOk() (*time.Time, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *DeviceConfigurationTemplate) SetUpdatedTimestamp(v time.Time)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetRevision

`func (o *DeviceConfigurationTemplate) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *DeviceConfigurationTemplate) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *DeviceConfigurationTemplate) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetLinks

`func (o *DeviceConfigurationTemplate) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DeviceConfigurationTemplate) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DeviceConfigurationTemplate) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DeviceConfigurationTemplate) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *DeviceConfigurationTemplate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceConfigurationTemplate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceConfigurationTemplate) SetId(v int64)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


