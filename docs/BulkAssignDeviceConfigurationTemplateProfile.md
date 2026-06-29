# BulkAssignDeviceConfigurationTemplateProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceConfigurationTemplateId** | **int64** | Template to assign across the chosen target set. | 
**NetworkFabricId** | Pointer to **int64** | Fabric to bulk-assign across — resolves to ALL NetworkDevice rows attached to this fabric. Mutually exclusive with networkDeviceIds. | [optional] 
**NetworkDeviceIds** | Pointer to **[]int64** | Explicit list of NetworkDevice ids to assign. Mutually exclusive with networkFabricId. | [optional] 
**LifecycleStage** | Pointer to [**DeviceConfigurationProfileLifecycleStage**](DeviceConfigurationProfileLifecycleStage.md) | Lifecycle stage applied to every created profile. Defaults to \&quot;configuration\&quot;. | [optional] 
**Variables** | Pointer to **map[string]interface{}** | Per-profile variable overrides applied uniformly to every created row. | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] [default to true]
**Priority** | Pointer to **float32** | Ordering hint when multiple enabled profiles match the same (device, lifecycleStage) tuple at deploy time — lower values apply first. Integer in the range [0, 1000]. Defaults to 100, which leaves headroom in both directions.  Recommended bands (convention, not enforced): - 0–49: system-level overrides — first thing on the wire after baseline reset. - 50–99: high-priority operator rules (e.g. emergency policy push). - 100: default — use this unless you have a reason not to. - 101–500: standard operator customizations. - 501–1000: low-priority fallbacks, applied last.  Ties (two profiles with the same priority for the same device + lifecycleStage) resolve in undefined order. Use distinct priorities if order matters. | [optional] [default to 100]
**ApplyMode** | Pointer to [**DeviceConfigurationProfileApplyMode**](DeviceConfigurationProfileApplyMode.md) | How often this profile applies during fabric deploys. - &#39;always&#39; (default): re-applied on every deploy that matches. - &#39;once&#39;: applied to each matching device exactly once across all deploys, then skipped. Tracking lives in &#x60;device_configuration_template_history&#x60; — one row per successful apply, keyed on the profile id and the network device id. Reset by deleting those history rows. | [optional] [default to DEVICECONFIGURATIONPROFILEAPPLYMODE_ALWAYS]
**Annotations** | Pointer to **map[string]string** | Annotations applied uniformly to every created profile. | [optional] 
**Tags** | Pointer to **[]string** | Tags applied uniformly to every created profile. Each entry is capped at 255 characters. | [optional] 

## Methods

### NewBulkAssignDeviceConfigurationTemplateProfile

`func NewBulkAssignDeviceConfigurationTemplateProfile(deviceConfigurationTemplateId int64, ) *BulkAssignDeviceConfigurationTemplateProfile`

NewBulkAssignDeviceConfigurationTemplateProfile instantiates a new BulkAssignDeviceConfigurationTemplateProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkAssignDeviceConfigurationTemplateProfileWithDefaults

`func NewBulkAssignDeviceConfigurationTemplateProfileWithDefaults() *BulkAssignDeviceConfigurationTemplateProfile`

NewBulkAssignDeviceConfigurationTemplateProfileWithDefaults instantiates a new BulkAssignDeviceConfigurationTemplateProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceConfigurationTemplateId

`func (o *BulkAssignDeviceConfigurationTemplateProfile) GetDeviceConfigurationTemplateId() int64`

GetDeviceConfigurationTemplateId returns the DeviceConfigurationTemplateId field if non-nil, zero value otherwise.

### GetDeviceConfigurationTemplateIdOk

`func (o *BulkAssignDeviceConfigurationTemplateProfile) GetDeviceConfigurationTemplateIdOk() (*int64, bool)`

GetDeviceConfigurationTemplateIdOk returns a tuple with the DeviceConfigurationTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceConfigurationTemplateId

`func (o *BulkAssignDeviceConfigurationTemplateProfile) SetDeviceConfigurationTemplateId(v int64)`

SetDeviceConfigurationTemplateId sets DeviceConfigurationTemplateId field to given value.


### GetNetworkFabricId

`func (o *BulkAssignDeviceConfigurationTemplateProfile) GetNetworkFabricId() int64`

GetNetworkFabricId returns the NetworkFabricId field if non-nil, zero value otherwise.

### GetNetworkFabricIdOk

`func (o *BulkAssignDeviceConfigurationTemplateProfile) GetNetworkFabricIdOk() (*int64, bool)`

GetNetworkFabricIdOk returns a tuple with the NetworkFabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFabricId

`func (o *BulkAssignDeviceConfigurationTemplateProfile) SetNetworkFabricId(v int64)`

SetNetworkFabricId sets NetworkFabricId field to given value.

### HasNetworkFabricId

`func (o *BulkAssignDeviceConfigurationTemplateProfile) HasNetworkFabricId() bool`

HasNetworkFabricId returns a boolean if a field has been set.

### GetNetworkDeviceIds

`func (o *BulkAssignDeviceConfigurationTemplateProfile) GetNetworkDeviceIds() []int64`

GetNetworkDeviceIds returns the NetworkDeviceIds field if non-nil, zero value otherwise.

### GetNetworkDeviceIdsOk

`func (o *BulkAssignDeviceConfigurationTemplateProfile) GetNetworkDeviceIdsOk() (*[]int64, bool)`

GetNetworkDeviceIdsOk returns a tuple with the NetworkDeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceIds

`func (o *BulkAssignDeviceConfigurationTemplateProfile) SetNetworkDeviceIds(v []int64)`

SetNetworkDeviceIds sets NetworkDeviceIds field to given value.

### HasNetworkDeviceIds

`func (o *BulkAssignDeviceConfigurationTemplateProfile) HasNetworkDeviceIds() bool`

HasNetworkDeviceIds returns a boolean if a field has been set.

### GetLifecycleStage

`func (o *BulkAssignDeviceConfigurationTemplateProfile) GetLifecycleStage() DeviceConfigurationProfileLifecycleStage`

GetLifecycleStage returns the LifecycleStage field if non-nil, zero value otherwise.

### GetLifecycleStageOk

`func (o *BulkAssignDeviceConfigurationTemplateProfile) GetLifecycleStageOk() (*DeviceConfigurationProfileLifecycleStage, bool)`

GetLifecycleStageOk returns a tuple with the LifecycleStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleStage

`func (o *BulkAssignDeviceConfigurationTemplateProfile) SetLifecycleStage(v DeviceConfigurationProfileLifecycleStage)`

SetLifecycleStage sets LifecycleStage field to given value.

### HasLifecycleStage

`func (o *BulkAssignDeviceConfigurationTemplateProfile) HasLifecycleStage() bool`

HasLifecycleStage returns a boolean if a field has been set.

### GetVariables

`func (o *BulkAssignDeviceConfigurationTemplateProfile) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *BulkAssignDeviceConfigurationTemplateProfile) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *BulkAssignDeviceConfigurationTemplateProfile) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *BulkAssignDeviceConfigurationTemplateProfile) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetIsEnabled

`func (o *BulkAssignDeviceConfigurationTemplateProfile) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *BulkAssignDeviceConfigurationTemplateProfile) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *BulkAssignDeviceConfigurationTemplateProfile) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *BulkAssignDeviceConfigurationTemplateProfile) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetPriority

`func (o *BulkAssignDeviceConfigurationTemplateProfile) GetPriority() float32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *BulkAssignDeviceConfigurationTemplateProfile) GetPriorityOk() (*float32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *BulkAssignDeviceConfigurationTemplateProfile) SetPriority(v float32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *BulkAssignDeviceConfigurationTemplateProfile) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetApplyMode

`func (o *BulkAssignDeviceConfigurationTemplateProfile) GetApplyMode() DeviceConfigurationProfileApplyMode`

GetApplyMode returns the ApplyMode field if non-nil, zero value otherwise.

### GetApplyModeOk

`func (o *BulkAssignDeviceConfigurationTemplateProfile) GetApplyModeOk() (*DeviceConfigurationProfileApplyMode, bool)`

GetApplyModeOk returns a tuple with the ApplyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyMode

`func (o *BulkAssignDeviceConfigurationTemplateProfile) SetApplyMode(v DeviceConfigurationProfileApplyMode)`

SetApplyMode sets ApplyMode field to given value.

### HasApplyMode

`func (o *BulkAssignDeviceConfigurationTemplateProfile) HasApplyMode() bool`

HasApplyMode returns a boolean if a field has been set.

### GetAnnotations

`func (o *BulkAssignDeviceConfigurationTemplateProfile) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *BulkAssignDeviceConfigurationTemplateProfile) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *BulkAssignDeviceConfigurationTemplateProfile) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *BulkAssignDeviceConfigurationTemplateProfile) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetTags

`func (o *BulkAssignDeviceConfigurationTemplateProfile) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkAssignDeviceConfigurationTemplateProfile) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkAssignDeviceConfigurationTemplateProfile) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkAssignDeviceConfigurationTemplateProfile) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


