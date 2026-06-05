# UpdateDeviceConfigurationTemplateProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkFabricId** | Pointer to **NullableInt64** | Optional fabric scope; null &#x3D; applies regardless of fabric. | [optional] 
**LifecycleStage** | Pointer to [**DeviceConfigurationProfileLifecycleStage**](DeviceConfigurationProfileLifecycleStage.md) | When this profile is consumed by the device lifecycle. Defaults to \&quot;configuration\&quot;. | [optional] 
**Variables** | Pointer to **map[string]interface{}** | Per-profile render variable overrides. Highest precedence at render time: profile &gt; device.customVariables &gt; template.customVariablesJson. | [optional] 
**IsEnabled** | Pointer to **bool** | Disable to skip applying this profile without deleting it from the system. | [optional] [default to true]
**Priority** | Pointer to **float32** | Lower applies first when multiple profiles match a (device, lifecycleStage) query. | [optional] [default to 100]

## Methods

### NewUpdateDeviceConfigurationTemplateProfile

`func NewUpdateDeviceConfigurationTemplateProfile() *UpdateDeviceConfigurationTemplateProfile`

NewUpdateDeviceConfigurationTemplateProfile instantiates a new UpdateDeviceConfigurationTemplateProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeviceConfigurationTemplateProfileWithDefaults

`func NewUpdateDeviceConfigurationTemplateProfileWithDefaults() *UpdateDeviceConfigurationTemplateProfile`

NewUpdateDeviceConfigurationTemplateProfileWithDefaults instantiates a new UpdateDeviceConfigurationTemplateProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkFabricId

`func (o *UpdateDeviceConfigurationTemplateProfile) GetNetworkFabricId() int64`

GetNetworkFabricId returns the NetworkFabricId field if non-nil, zero value otherwise.

### GetNetworkFabricIdOk

`func (o *UpdateDeviceConfigurationTemplateProfile) GetNetworkFabricIdOk() (*int64, bool)`

GetNetworkFabricIdOk returns a tuple with the NetworkFabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFabricId

`func (o *UpdateDeviceConfigurationTemplateProfile) SetNetworkFabricId(v int64)`

SetNetworkFabricId sets NetworkFabricId field to given value.

### HasNetworkFabricId

`func (o *UpdateDeviceConfigurationTemplateProfile) HasNetworkFabricId() bool`

HasNetworkFabricId returns a boolean if a field has been set.

### SetNetworkFabricIdNil

`func (o *UpdateDeviceConfigurationTemplateProfile) SetNetworkFabricIdNil(b bool)`

 SetNetworkFabricIdNil sets the value for NetworkFabricId to be an explicit nil

### UnsetNetworkFabricId
`func (o *UpdateDeviceConfigurationTemplateProfile) UnsetNetworkFabricId()`

UnsetNetworkFabricId ensures that no value is present for NetworkFabricId, not even an explicit nil
### GetLifecycleStage

`func (o *UpdateDeviceConfigurationTemplateProfile) GetLifecycleStage() DeviceConfigurationProfileLifecycleStage`

GetLifecycleStage returns the LifecycleStage field if non-nil, zero value otherwise.

### GetLifecycleStageOk

`func (o *UpdateDeviceConfigurationTemplateProfile) GetLifecycleStageOk() (*DeviceConfigurationProfileLifecycleStage, bool)`

GetLifecycleStageOk returns a tuple with the LifecycleStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleStage

`func (o *UpdateDeviceConfigurationTemplateProfile) SetLifecycleStage(v DeviceConfigurationProfileLifecycleStage)`

SetLifecycleStage sets LifecycleStage field to given value.

### HasLifecycleStage

`func (o *UpdateDeviceConfigurationTemplateProfile) HasLifecycleStage() bool`

HasLifecycleStage returns a boolean if a field has been set.

### GetVariables

`func (o *UpdateDeviceConfigurationTemplateProfile) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *UpdateDeviceConfigurationTemplateProfile) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *UpdateDeviceConfigurationTemplateProfile) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *UpdateDeviceConfigurationTemplateProfile) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetIsEnabled

`func (o *UpdateDeviceConfigurationTemplateProfile) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *UpdateDeviceConfigurationTemplateProfile) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *UpdateDeviceConfigurationTemplateProfile) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *UpdateDeviceConfigurationTemplateProfile) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetPriority

`func (o *UpdateDeviceConfigurationTemplateProfile) GetPriority() float32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *UpdateDeviceConfigurationTemplateProfile) GetPriorityOk() (*float32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *UpdateDeviceConfigurationTemplateProfile) SetPriority(v float32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *UpdateDeviceConfigurationTemplateProfile) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


