# RenderDeviceConfigurationTemplateProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkDeviceId** | **int64** | NetworkDevice id to render against. Must match the profile&#39;s bound device when the profile is device-scoped; required for fabric-scoped profiles. | 
**ExtraVariables** | Pointer to **map[string]interface{}** | Ad-hoc render-context overrides with HIGHEST user-supplied precedence. Precedence (low → high): template.customVariablesJson &lt; profile.variables &lt; extraVariables &lt; whitelisted device fields. The system-injected &#x60;interfaces&#x60; key (array of switch ports) is reserved. Whitelisted device fields (e.g. &#x60;identifierString&#x60;, &#x60;managementAddress&#x60;, &#x60;asn&#x60;) are spread at top level and override any user-provided keys with the same name. See &#x60;DEVICE_RENDER_FIELDS&#x60; / &#x60;INTERFACE_RENDER_FIELDS&#x60; in the inventory service for the exposed schema. | [optional] 
**Debug** | Pointer to **bool** | When true, the response additionally includes the decoded (non-base64) template source and the full variable bag passed to the engine. | [optional] [default to false]

## Methods

### NewRenderDeviceConfigurationTemplateProfile

`func NewRenderDeviceConfigurationTemplateProfile(networkDeviceId int64, ) *RenderDeviceConfigurationTemplateProfile`

NewRenderDeviceConfigurationTemplateProfile instantiates a new RenderDeviceConfigurationTemplateProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenderDeviceConfigurationTemplateProfileWithDefaults

`func NewRenderDeviceConfigurationTemplateProfileWithDefaults() *RenderDeviceConfigurationTemplateProfile`

NewRenderDeviceConfigurationTemplateProfileWithDefaults instantiates a new RenderDeviceConfigurationTemplateProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkDeviceId

`func (o *RenderDeviceConfigurationTemplateProfile) GetNetworkDeviceId() int64`

GetNetworkDeviceId returns the NetworkDeviceId field if non-nil, zero value otherwise.

### GetNetworkDeviceIdOk

`func (o *RenderDeviceConfigurationTemplateProfile) GetNetworkDeviceIdOk() (*int64, bool)`

GetNetworkDeviceIdOk returns a tuple with the NetworkDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceId

`func (o *RenderDeviceConfigurationTemplateProfile) SetNetworkDeviceId(v int64)`

SetNetworkDeviceId sets NetworkDeviceId field to given value.


### GetExtraVariables

`func (o *RenderDeviceConfigurationTemplateProfile) GetExtraVariables() map[string]interface{}`

GetExtraVariables returns the ExtraVariables field if non-nil, zero value otherwise.

### GetExtraVariablesOk

`func (o *RenderDeviceConfigurationTemplateProfile) GetExtraVariablesOk() (*map[string]interface{}, bool)`

GetExtraVariablesOk returns a tuple with the ExtraVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraVariables

`func (o *RenderDeviceConfigurationTemplateProfile) SetExtraVariables(v map[string]interface{})`

SetExtraVariables sets ExtraVariables field to given value.

### HasExtraVariables

`func (o *RenderDeviceConfigurationTemplateProfile) HasExtraVariables() bool`

HasExtraVariables returns a boolean if a field has been set.

### GetDebug

`func (o *RenderDeviceConfigurationTemplateProfile) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *RenderDeviceConfigurationTemplateProfile) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *RenderDeviceConfigurationTemplateProfile) SetDebug(v bool)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *RenderDeviceConfigurationTemplateProfile) HasDebug() bool`

HasDebug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


