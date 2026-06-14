# RenderApplicableDeviceConfigurationTemplateProfiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkDeviceId** | Pointer to **int64** | NetworkDevice id to match against. Returns profiles bound to this device (profile.networkDeviceId &#x3D; id) plus fabric-bound profiles (profile.networkDeviceId IS NULL) for every fabric the device is attached to. Mutually exclusive with networkFabricId — exactly one of the two must be defined (enforced at the service layer). | [optional] 
**NetworkFabricId** | Pointer to **int64** | NetworkFabric id to match against. Returns fabric-bound profiles (profile.networkFabricId &#x3D; id AND profile.networkDeviceId IS NULL) plus device-bound profiles whose target device belongs to this fabric. Mutually exclusive with networkDeviceId — exactly one of the two must be defined (enforced at the service layer). | [optional] 
**LifecycleStage** | Pointer to [**DeviceConfigurationProfileLifecycleStage**](DeviceConfigurationProfileLifecycleStage.md) | Optional lifecycle-stage filter. When omitted, profiles of every stage are returned. | [optional] 
**IncludeDisabled** | Pointer to **bool** | When true, profiles with isEnabled&#x3D;false are included. Defaults to enabled-only. | [optional] [default to false]
**ExtraVariables** | Pointer to **map[string]interface{}** | Ad-hoc render-context overrides applied to every rendered profile. Precedence (low → high): template.customVariablesJson &lt; profile.variables &lt; extraVariables &lt; whitelisted device fields. Whitelisted device fields are only injected when networkDeviceId is provided. | [optional] 
**Debug** | Pointer to **bool** | When true, each rendered item additionally includes the decoded (non-base64) template source and the full variable bag passed to the engine. | [optional] [default to false]

## Methods

### NewRenderApplicableDeviceConfigurationTemplateProfiles

`func NewRenderApplicableDeviceConfigurationTemplateProfiles() *RenderApplicableDeviceConfigurationTemplateProfiles`

NewRenderApplicableDeviceConfigurationTemplateProfiles instantiates a new RenderApplicableDeviceConfigurationTemplateProfiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenderApplicableDeviceConfigurationTemplateProfilesWithDefaults

`func NewRenderApplicableDeviceConfigurationTemplateProfilesWithDefaults() *RenderApplicableDeviceConfigurationTemplateProfiles`

NewRenderApplicableDeviceConfigurationTemplateProfilesWithDefaults instantiates a new RenderApplicableDeviceConfigurationTemplateProfiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkDeviceId

`func (o *RenderApplicableDeviceConfigurationTemplateProfiles) GetNetworkDeviceId() int64`

GetNetworkDeviceId returns the NetworkDeviceId field if non-nil, zero value otherwise.

### GetNetworkDeviceIdOk

`func (o *RenderApplicableDeviceConfigurationTemplateProfiles) GetNetworkDeviceIdOk() (*int64, bool)`

GetNetworkDeviceIdOk returns a tuple with the NetworkDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceId

`func (o *RenderApplicableDeviceConfigurationTemplateProfiles) SetNetworkDeviceId(v int64)`

SetNetworkDeviceId sets NetworkDeviceId field to given value.

### HasNetworkDeviceId

`func (o *RenderApplicableDeviceConfigurationTemplateProfiles) HasNetworkDeviceId() bool`

HasNetworkDeviceId returns a boolean if a field has been set.

### GetNetworkFabricId

`func (o *RenderApplicableDeviceConfigurationTemplateProfiles) GetNetworkFabricId() int64`

GetNetworkFabricId returns the NetworkFabricId field if non-nil, zero value otherwise.

### GetNetworkFabricIdOk

`func (o *RenderApplicableDeviceConfigurationTemplateProfiles) GetNetworkFabricIdOk() (*int64, bool)`

GetNetworkFabricIdOk returns a tuple with the NetworkFabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFabricId

`func (o *RenderApplicableDeviceConfigurationTemplateProfiles) SetNetworkFabricId(v int64)`

SetNetworkFabricId sets NetworkFabricId field to given value.

### HasNetworkFabricId

`func (o *RenderApplicableDeviceConfigurationTemplateProfiles) HasNetworkFabricId() bool`

HasNetworkFabricId returns a boolean if a field has been set.

### GetLifecycleStage

`func (o *RenderApplicableDeviceConfigurationTemplateProfiles) GetLifecycleStage() DeviceConfigurationProfileLifecycleStage`

GetLifecycleStage returns the LifecycleStage field if non-nil, zero value otherwise.

### GetLifecycleStageOk

`func (o *RenderApplicableDeviceConfigurationTemplateProfiles) GetLifecycleStageOk() (*DeviceConfigurationProfileLifecycleStage, bool)`

GetLifecycleStageOk returns a tuple with the LifecycleStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleStage

`func (o *RenderApplicableDeviceConfigurationTemplateProfiles) SetLifecycleStage(v DeviceConfigurationProfileLifecycleStage)`

SetLifecycleStage sets LifecycleStage field to given value.

### HasLifecycleStage

`func (o *RenderApplicableDeviceConfigurationTemplateProfiles) HasLifecycleStage() bool`

HasLifecycleStage returns a boolean if a field has been set.

### GetIncludeDisabled

`func (o *RenderApplicableDeviceConfigurationTemplateProfiles) GetIncludeDisabled() bool`

GetIncludeDisabled returns the IncludeDisabled field if non-nil, zero value otherwise.

### GetIncludeDisabledOk

`func (o *RenderApplicableDeviceConfigurationTemplateProfiles) GetIncludeDisabledOk() (*bool, bool)`

GetIncludeDisabledOk returns a tuple with the IncludeDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDisabled

`func (o *RenderApplicableDeviceConfigurationTemplateProfiles) SetIncludeDisabled(v bool)`

SetIncludeDisabled sets IncludeDisabled field to given value.

### HasIncludeDisabled

`func (o *RenderApplicableDeviceConfigurationTemplateProfiles) HasIncludeDisabled() bool`

HasIncludeDisabled returns a boolean if a field has been set.

### GetExtraVariables

`func (o *RenderApplicableDeviceConfigurationTemplateProfiles) GetExtraVariables() map[string]interface{}`

GetExtraVariables returns the ExtraVariables field if non-nil, zero value otherwise.

### GetExtraVariablesOk

`func (o *RenderApplicableDeviceConfigurationTemplateProfiles) GetExtraVariablesOk() (*map[string]interface{}, bool)`

GetExtraVariablesOk returns a tuple with the ExtraVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraVariables

`func (o *RenderApplicableDeviceConfigurationTemplateProfiles) SetExtraVariables(v map[string]interface{})`

SetExtraVariables sets ExtraVariables field to given value.

### HasExtraVariables

`func (o *RenderApplicableDeviceConfigurationTemplateProfiles) HasExtraVariables() bool`

HasExtraVariables returns a boolean if a field has been set.

### GetDebug

`func (o *RenderApplicableDeviceConfigurationTemplateProfiles) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *RenderApplicableDeviceConfigurationTemplateProfiles) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *RenderApplicableDeviceConfigurationTemplateProfiles) SetDebug(v bool)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *RenderApplicableDeviceConfigurationTemplateProfiles) HasDebug() bool`

HasDebug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


