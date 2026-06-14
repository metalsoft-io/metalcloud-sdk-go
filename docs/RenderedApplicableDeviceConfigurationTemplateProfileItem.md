# RenderedApplicableDeviceConfigurationTemplateProfileItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Id of the rendered profile. | 
**DeviceConfigurationTemplateId** | **int64** |  | 
**NetworkDeviceId** | **NullableInt64** |  | 
**NetworkFabricId** | **NullableInt64** |  | 
**LifecycleStage** | [**DeviceConfigurationProfileLifecycleStage**](DeviceConfigurationProfileLifecycleStage.md) |  | 
**Priority** | **float32** |  | 
**Annotations** | **map[string]string** | Key-value annotations copied from the profile. | 
**Tags** | **[]string** | Tags copied from the profile. Empty array when the profile has no tags. | 
**Rendered** | **string** | Rendered template output as UTF-8 plain text. | 
**TemplateContent** | Pointer to **string** | Debug-only: decoded (non-base64) template source that was rendered. Present only when the request set &#x60;debug: true&#x60;. | [optional] 
**Variables** | Pointer to **map[string]interface{}** | Debug-only: full variable bag passed to the Jinja2 engine. Present only when the request set &#x60;debug: true&#x60;. | [optional] 

## Methods

### NewRenderedApplicableDeviceConfigurationTemplateProfileItem

`func NewRenderedApplicableDeviceConfigurationTemplateProfileItem(id int64, deviceConfigurationTemplateId int64, networkDeviceId NullableInt64, networkFabricId NullableInt64, lifecycleStage DeviceConfigurationProfileLifecycleStage, priority float32, annotations map[string]string, tags []string, rendered string, ) *RenderedApplicableDeviceConfigurationTemplateProfileItem`

NewRenderedApplicableDeviceConfigurationTemplateProfileItem instantiates a new RenderedApplicableDeviceConfigurationTemplateProfileItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenderedApplicableDeviceConfigurationTemplateProfileItemWithDefaults

`func NewRenderedApplicableDeviceConfigurationTemplateProfileItemWithDefaults() *RenderedApplicableDeviceConfigurationTemplateProfileItem`

NewRenderedApplicableDeviceConfigurationTemplateProfileItemWithDefaults instantiates a new RenderedApplicableDeviceConfigurationTemplateProfileItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) SetId(v int64)`

SetId sets Id field to given value.


### GetDeviceConfigurationTemplateId

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) GetDeviceConfigurationTemplateId() int64`

GetDeviceConfigurationTemplateId returns the DeviceConfigurationTemplateId field if non-nil, zero value otherwise.

### GetDeviceConfigurationTemplateIdOk

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) GetDeviceConfigurationTemplateIdOk() (*int64, bool)`

GetDeviceConfigurationTemplateIdOk returns a tuple with the DeviceConfigurationTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceConfigurationTemplateId

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) SetDeviceConfigurationTemplateId(v int64)`

SetDeviceConfigurationTemplateId sets DeviceConfigurationTemplateId field to given value.


### GetNetworkDeviceId

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) GetNetworkDeviceId() int64`

GetNetworkDeviceId returns the NetworkDeviceId field if non-nil, zero value otherwise.

### GetNetworkDeviceIdOk

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) GetNetworkDeviceIdOk() (*int64, bool)`

GetNetworkDeviceIdOk returns a tuple with the NetworkDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceId

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) SetNetworkDeviceId(v int64)`

SetNetworkDeviceId sets NetworkDeviceId field to given value.


### SetNetworkDeviceIdNil

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) SetNetworkDeviceIdNil(b bool)`

 SetNetworkDeviceIdNil sets the value for NetworkDeviceId to be an explicit nil

### UnsetNetworkDeviceId
`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) UnsetNetworkDeviceId()`

UnsetNetworkDeviceId ensures that no value is present for NetworkDeviceId, not even an explicit nil
### GetNetworkFabricId

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) GetNetworkFabricId() int64`

GetNetworkFabricId returns the NetworkFabricId field if non-nil, zero value otherwise.

### GetNetworkFabricIdOk

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) GetNetworkFabricIdOk() (*int64, bool)`

GetNetworkFabricIdOk returns a tuple with the NetworkFabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFabricId

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) SetNetworkFabricId(v int64)`

SetNetworkFabricId sets NetworkFabricId field to given value.


### SetNetworkFabricIdNil

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) SetNetworkFabricIdNil(b bool)`

 SetNetworkFabricIdNil sets the value for NetworkFabricId to be an explicit nil

### UnsetNetworkFabricId
`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) UnsetNetworkFabricId()`

UnsetNetworkFabricId ensures that no value is present for NetworkFabricId, not even an explicit nil
### GetLifecycleStage

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) GetLifecycleStage() DeviceConfigurationProfileLifecycleStage`

GetLifecycleStage returns the LifecycleStage field if non-nil, zero value otherwise.

### GetLifecycleStageOk

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) GetLifecycleStageOk() (*DeviceConfigurationProfileLifecycleStage, bool)`

GetLifecycleStageOk returns a tuple with the LifecycleStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleStage

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) SetLifecycleStage(v DeviceConfigurationProfileLifecycleStage)`

SetLifecycleStage sets LifecycleStage field to given value.


### GetPriority

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) GetPriority() float32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) GetPriorityOk() (*float32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) SetPriority(v float32)`

SetPriority sets Priority field to given value.


### GetAnnotations

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.


### GetTags

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetRendered

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) GetRendered() string`

GetRendered returns the Rendered field if non-nil, zero value otherwise.

### GetRenderedOk

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) GetRenderedOk() (*string, bool)`

GetRenderedOk returns a tuple with the Rendered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRendered

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) SetRendered(v string)`

SetRendered sets Rendered field to given value.


### GetTemplateContent

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) GetTemplateContent() string`

GetTemplateContent returns the TemplateContent field if non-nil, zero value otherwise.

### GetTemplateContentOk

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) GetTemplateContentOk() (*string, bool)`

GetTemplateContentOk returns a tuple with the TemplateContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateContent

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) SetTemplateContent(v string)`

SetTemplateContent sets TemplateContent field to given value.

### HasTemplateContent

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) HasTemplateContent() bool`

HasTemplateContent returns a boolean if a field has been set.

### GetVariables

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *RenderedApplicableDeviceConfigurationTemplateProfileItem) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


