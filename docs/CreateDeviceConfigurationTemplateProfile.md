# CreateDeviceConfigurationTemplateProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceConfigurationTemplateId** | **int64** | Id of the Device Configuration Template to bind. | 
**NetworkDeviceId** | Pointer to **NullableInt64** | Target NetworkDevice id (null &#x3D; fabric-wide intent). | [optional] 
**NetworkFabricId** | Pointer to **NullableInt64** | Optional fabric scope; null &#x3D; applies regardless of fabric. | [optional] 
**LifecycleStage** | Pointer to [**DeviceConfigurationProfileLifecycleStage**](DeviceConfigurationProfileLifecycleStage.md) | When this profile is consumed by the device lifecycle. Defaults to \&quot;configuration\&quot;. | [optional] 
**Variables** | Pointer to **map[string]interface{}** | Per-profile render variable overrides. Highest precedence at render time: profile &gt; device.customVariables &gt; template.customVariablesJson. | [optional] 
**IsEnabled** | Pointer to **bool** | Disable to skip applying this profile without deleting it from the system. | [optional] [default to true]
**Priority** | Pointer to **float32** | Ordering hint when multiple enabled profiles match the same (device, lifecycleStage) tuple at deploy time — lower values apply first. Integer in the range [0, 1000]. Defaults to 100, which leaves headroom in both directions.  Recommended bands (convention, not enforced): - 0–49: system-level overrides — first thing on the wire after baseline reset. - 50–99: high-priority operator rules (e.g. emergency policy push). - 100: default — use this unless you have a reason not to. - 101–500: standard operator customizations. - 501–1000: low-priority fallbacks, applied last.  Ties (two profiles with the same priority for the same device + lifecycleStage) resolve in undefined order. Use distinct priorities if order matters. | [optional] [default to 100]
**ApplyMode** | Pointer to [**DeviceConfigurationProfileApplyMode**](DeviceConfigurationProfileApplyMode.md) | How often this profile applies during fabric deploys. - &#39;always&#39; (default): re-applied on every deploy that matches. - &#39;once&#39;: applied to each matching device exactly once across all deploys, then skipped. Tracking lives in &#x60;device_configuration_template_history&#x60; — one row per successful apply, keyed on the profile id and the network device id. Reset by deleting those history rows. | [optional] [default to DEVICECONFIGURATIONPROFILEAPPLYMODE_ALWAYS]
**Annotations** | Pointer to **map[string]string** | Key-value annotations for storing additional metadata. | [optional] 
**Tags** | Pointer to **[]string** | Tags associated with the profile. Each entry is capped at 255 characters. | [optional] 

## Methods

### NewCreateDeviceConfigurationTemplateProfile

`func NewCreateDeviceConfigurationTemplateProfile(deviceConfigurationTemplateId int64, ) *CreateDeviceConfigurationTemplateProfile`

NewCreateDeviceConfigurationTemplateProfile instantiates a new CreateDeviceConfigurationTemplateProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeviceConfigurationTemplateProfileWithDefaults

`func NewCreateDeviceConfigurationTemplateProfileWithDefaults() *CreateDeviceConfigurationTemplateProfile`

NewCreateDeviceConfigurationTemplateProfileWithDefaults instantiates a new CreateDeviceConfigurationTemplateProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceConfigurationTemplateId

`func (o *CreateDeviceConfigurationTemplateProfile) GetDeviceConfigurationTemplateId() int64`

GetDeviceConfigurationTemplateId returns the DeviceConfigurationTemplateId field if non-nil, zero value otherwise.

### GetDeviceConfigurationTemplateIdOk

`func (o *CreateDeviceConfigurationTemplateProfile) GetDeviceConfigurationTemplateIdOk() (*int64, bool)`

GetDeviceConfigurationTemplateIdOk returns a tuple with the DeviceConfigurationTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceConfigurationTemplateId

`func (o *CreateDeviceConfigurationTemplateProfile) SetDeviceConfigurationTemplateId(v int64)`

SetDeviceConfigurationTemplateId sets DeviceConfigurationTemplateId field to given value.


### GetNetworkDeviceId

`func (o *CreateDeviceConfigurationTemplateProfile) GetNetworkDeviceId() int64`

GetNetworkDeviceId returns the NetworkDeviceId field if non-nil, zero value otherwise.

### GetNetworkDeviceIdOk

`func (o *CreateDeviceConfigurationTemplateProfile) GetNetworkDeviceIdOk() (*int64, bool)`

GetNetworkDeviceIdOk returns a tuple with the NetworkDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceId

`func (o *CreateDeviceConfigurationTemplateProfile) SetNetworkDeviceId(v int64)`

SetNetworkDeviceId sets NetworkDeviceId field to given value.

### HasNetworkDeviceId

`func (o *CreateDeviceConfigurationTemplateProfile) HasNetworkDeviceId() bool`

HasNetworkDeviceId returns a boolean if a field has been set.

### SetNetworkDeviceIdNil

`func (o *CreateDeviceConfigurationTemplateProfile) SetNetworkDeviceIdNil(b bool)`

 SetNetworkDeviceIdNil sets the value for NetworkDeviceId to be an explicit nil

### UnsetNetworkDeviceId
`func (o *CreateDeviceConfigurationTemplateProfile) UnsetNetworkDeviceId()`

UnsetNetworkDeviceId ensures that no value is present for NetworkDeviceId, not even an explicit nil
### GetNetworkFabricId

`func (o *CreateDeviceConfigurationTemplateProfile) GetNetworkFabricId() int64`

GetNetworkFabricId returns the NetworkFabricId field if non-nil, zero value otherwise.

### GetNetworkFabricIdOk

`func (o *CreateDeviceConfigurationTemplateProfile) GetNetworkFabricIdOk() (*int64, bool)`

GetNetworkFabricIdOk returns a tuple with the NetworkFabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFabricId

`func (o *CreateDeviceConfigurationTemplateProfile) SetNetworkFabricId(v int64)`

SetNetworkFabricId sets NetworkFabricId field to given value.

### HasNetworkFabricId

`func (o *CreateDeviceConfigurationTemplateProfile) HasNetworkFabricId() bool`

HasNetworkFabricId returns a boolean if a field has been set.

### SetNetworkFabricIdNil

`func (o *CreateDeviceConfigurationTemplateProfile) SetNetworkFabricIdNil(b bool)`

 SetNetworkFabricIdNil sets the value for NetworkFabricId to be an explicit nil

### UnsetNetworkFabricId
`func (o *CreateDeviceConfigurationTemplateProfile) UnsetNetworkFabricId()`

UnsetNetworkFabricId ensures that no value is present for NetworkFabricId, not even an explicit nil
### GetLifecycleStage

`func (o *CreateDeviceConfigurationTemplateProfile) GetLifecycleStage() DeviceConfigurationProfileLifecycleStage`

GetLifecycleStage returns the LifecycleStage field if non-nil, zero value otherwise.

### GetLifecycleStageOk

`func (o *CreateDeviceConfigurationTemplateProfile) GetLifecycleStageOk() (*DeviceConfigurationProfileLifecycleStage, bool)`

GetLifecycleStageOk returns a tuple with the LifecycleStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleStage

`func (o *CreateDeviceConfigurationTemplateProfile) SetLifecycleStage(v DeviceConfigurationProfileLifecycleStage)`

SetLifecycleStage sets LifecycleStage field to given value.

### HasLifecycleStage

`func (o *CreateDeviceConfigurationTemplateProfile) HasLifecycleStage() bool`

HasLifecycleStage returns a boolean if a field has been set.

### GetVariables

`func (o *CreateDeviceConfigurationTemplateProfile) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *CreateDeviceConfigurationTemplateProfile) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *CreateDeviceConfigurationTemplateProfile) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *CreateDeviceConfigurationTemplateProfile) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetIsEnabled

`func (o *CreateDeviceConfigurationTemplateProfile) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *CreateDeviceConfigurationTemplateProfile) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *CreateDeviceConfigurationTemplateProfile) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *CreateDeviceConfigurationTemplateProfile) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetPriority

`func (o *CreateDeviceConfigurationTemplateProfile) GetPriority() float32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CreateDeviceConfigurationTemplateProfile) GetPriorityOk() (*float32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CreateDeviceConfigurationTemplateProfile) SetPriority(v float32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CreateDeviceConfigurationTemplateProfile) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetApplyMode

`func (o *CreateDeviceConfigurationTemplateProfile) GetApplyMode() DeviceConfigurationProfileApplyMode`

GetApplyMode returns the ApplyMode field if non-nil, zero value otherwise.

### GetApplyModeOk

`func (o *CreateDeviceConfigurationTemplateProfile) GetApplyModeOk() (*DeviceConfigurationProfileApplyMode, bool)`

GetApplyModeOk returns a tuple with the ApplyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyMode

`func (o *CreateDeviceConfigurationTemplateProfile) SetApplyMode(v DeviceConfigurationProfileApplyMode)`

SetApplyMode sets ApplyMode field to given value.

### HasApplyMode

`func (o *CreateDeviceConfigurationTemplateProfile) HasApplyMode() bool`

HasApplyMode returns a boolean if a field has been set.

### GetAnnotations

`func (o *CreateDeviceConfigurationTemplateProfile) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *CreateDeviceConfigurationTemplateProfile) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *CreateDeviceConfigurationTemplateProfile) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *CreateDeviceConfigurationTemplateProfile) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetTags

`func (o *CreateDeviceConfigurationTemplateProfile) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateDeviceConfigurationTemplateProfile) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateDeviceConfigurationTemplateProfile) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateDeviceConfigurationTemplateProfile) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


