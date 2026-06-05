# DeviceConfigurationTemplateProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceConfigurationTemplateId** | **int64** | Id of the Device Configuration Template to bind. | 
**NetworkDeviceId** | Pointer to **NullableInt64** | Target NetworkDevice id (null &#x3D; fabric-wide intent). | [optional] 
**NetworkFabricId** | Pointer to **NullableInt64** | Optional fabric scope; null &#x3D; applies regardless of fabric. | [optional] 
**LifecycleStage** | Pointer to [**DeviceConfigurationProfileLifecycleStage**](DeviceConfigurationProfileLifecycleStage.md) | When this profile is consumed by the device lifecycle. Defaults to \&quot;configuration\&quot;. | [optional] 
**Variables** | Pointer to **map[string]interface{}** | Per-profile render variable overrides. Highest precedence at render time: profile &gt; device.customVariables &gt; template.customVariablesJson. | [optional] 
**IsEnabled** | Pointer to **bool** | Disable to skip applying this profile without deleting it from the system. | [optional] [default to true]
**Priority** | Pointer to **float32** | Lower applies first when multiple profiles match a (device, lifecycleStage) query. | [optional] [default to 100]
**CreatedTimestamp** | **time.Time** | Entity creation timestamp | 
**UpdatedTimestamp** | **time.Time** | Entity last update timestamp | 
**Revision** | **string** | Revision number of the entity | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 
**Id** | **string** | Id of the device configuration template profile. | 
**CreatedBy** | Pointer to **NullableFloat32** | Id of the user who created this profile (from the X-User-Id header at create time). | [optional] 
**UpdatedBy** | Pointer to **NullableFloat32** | Id of the user who last updated this profile (from the X-User-Id header at update time). | [optional] 

## Methods

### NewDeviceConfigurationTemplateProfile

`func NewDeviceConfigurationTemplateProfile(deviceConfigurationTemplateId int64, createdTimestamp time.Time, updatedTimestamp time.Time, revision string, id string, ) *DeviceConfigurationTemplateProfile`

NewDeviceConfigurationTemplateProfile instantiates a new DeviceConfigurationTemplateProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceConfigurationTemplateProfileWithDefaults

`func NewDeviceConfigurationTemplateProfileWithDefaults() *DeviceConfigurationTemplateProfile`

NewDeviceConfigurationTemplateProfileWithDefaults instantiates a new DeviceConfigurationTemplateProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceConfigurationTemplateId

`func (o *DeviceConfigurationTemplateProfile) GetDeviceConfigurationTemplateId() int64`

GetDeviceConfigurationTemplateId returns the DeviceConfigurationTemplateId field if non-nil, zero value otherwise.

### GetDeviceConfigurationTemplateIdOk

`func (o *DeviceConfigurationTemplateProfile) GetDeviceConfigurationTemplateIdOk() (*int64, bool)`

GetDeviceConfigurationTemplateIdOk returns a tuple with the DeviceConfigurationTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceConfigurationTemplateId

`func (o *DeviceConfigurationTemplateProfile) SetDeviceConfigurationTemplateId(v int64)`

SetDeviceConfigurationTemplateId sets DeviceConfigurationTemplateId field to given value.


### GetNetworkDeviceId

`func (o *DeviceConfigurationTemplateProfile) GetNetworkDeviceId() int64`

GetNetworkDeviceId returns the NetworkDeviceId field if non-nil, zero value otherwise.

### GetNetworkDeviceIdOk

`func (o *DeviceConfigurationTemplateProfile) GetNetworkDeviceIdOk() (*int64, bool)`

GetNetworkDeviceIdOk returns a tuple with the NetworkDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceId

`func (o *DeviceConfigurationTemplateProfile) SetNetworkDeviceId(v int64)`

SetNetworkDeviceId sets NetworkDeviceId field to given value.

### HasNetworkDeviceId

`func (o *DeviceConfigurationTemplateProfile) HasNetworkDeviceId() bool`

HasNetworkDeviceId returns a boolean if a field has been set.

### SetNetworkDeviceIdNil

`func (o *DeviceConfigurationTemplateProfile) SetNetworkDeviceIdNil(b bool)`

 SetNetworkDeviceIdNil sets the value for NetworkDeviceId to be an explicit nil

### UnsetNetworkDeviceId
`func (o *DeviceConfigurationTemplateProfile) UnsetNetworkDeviceId()`

UnsetNetworkDeviceId ensures that no value is present for NetworkDeviceId, not even an explicit nil
### GetNetworkFabricId

`func (o *DeviceConfigurationTemplateProfile) GetNetworkFabricId() int64`

GetNetworkFabricId returns the NetworkFabricId field if non-nil, zero value otherwise.

### GetNetworkFabricIdOk

`func (o *DeviceConfigurationTemplateProfile) GetNetworkFabricIdOk() (*int64, bool)`

GetNetworkFabricIdOk returns a tuple with the NetworkFabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFabricId

`func (o *DeviceConfigurationTemplateProfile) SetNetworkFabricId(v int64)`

SetNetworkFabricId sets NetworkFabricId field to given value.

### HasNetworkFabricId

`func (o *DeviceConfigurationTemplateProfile) HasNetworkFabricId() bool`

HasNetworkFabricId returns a boolean if a field has been set.

### SetNetworkFabricIdNil

`func (o *DeviceConfigurationTemplateProfile) SetNetworkFabricIdNil(b bool)`

 SetNetworkFabricIdNil sets the value for NetworkFabricId to be an explicit nil

### UnsetNetworkFabricId
`func (o *DeviceConfigurationTemplateProfile) UnsetNetworkFabricId()`

UnsetNetworkFabricId ensures that no value is present for NetworkFabricId, not even an explicit nil
### GetLifecycleStage

`func (o *DeviceConfigurationTemplateProfile) GetLifecycleStage() DeviceConfigurationProfileLifecycleStage`

GetLifecycleStage returns the LifecycleStage field if non-nil, zero value otherwise.

### GetLifecycleStageOk

`func (o *DeviceConfigurationTemplateProfile) GetLifecycleStageOk() (*DeviceConfigurationProfileLifecycleStage, bool)`

GetLifecycleStageOk returns a tuple with the LifecycleStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleStage

`func (o *DeviceConfigurationTemplateProfile) SetLifecycleStage(v DeviceConfigurationProfileLifecycleStage)`

SetLifecycleStage sets LifecycleStage field to given value.

### HasLifecycleStage

`func (o *DeviceConfigurationTemplateProfile) HasLifecycleStage() bool`

HasLifecycleStage returns a boolean if a field has been set.

### GetVariables

`func (o *DeviceConfigurationTemplateProfile) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *DeviceConfigurationTemplateProfile) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *DeviceConfigurationTemplateProfile) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *DeviceConfigurationTemplateProfile) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetIsEnabled

`func (o *DeviceConfigurationTemplateProfile) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *DeviceConfigurationTemplateProfile) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *DeviceConfigurationTemplateProfile) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *DeviceConfigurationTemplateProfile) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetPriority

`func (o *DeviceConfigurationTemplateProfile) GetPriority() float32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DeviceConfigurationTemplateProfile) GetPriorityOk() (*float32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DeviceConfigurationTemplateProfile) SetPriority(v float32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DeviceConfigurationTemplateProfile) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *DeviceConfigurationTemplateProfile) GetCreatedTimestamp() time.Time`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *DeviceConfigurationTemplateProfile) GetCreatedTimestampOk() (*time.Time, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *DeviceConfigurationTemplateProfile) SetCreatedTimestamp(v time.Time)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *DeviceConfigurationTemplateProfile) GetUpdatedTimestamp() time.Time`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *DeviceConfigurationTemplateProfile) GetUpdatedTimestampOk() (*time.Time, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *DeviceConfigurationTemplateProfile) SetUpdatedTimestamp(v time.Time)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetRevision

`func (o *DeviceConfigurationTemplateProfile) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *DeviceConfigurationTemplateProfile) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *DeviceConfigurationTemplateProfile) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetLinks

`func (o *DeviceConfigurationTemplateProfile) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DeviceConfigurationTemplateProfile) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DeviceConfigurationTemplateProfile) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DeviceConfigurationTemplateProfile) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *DeviceConfigurationTemplateProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceConfigurationTemplateProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceConfigurationTemplateProfile) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *DeviceConfigurationTemplateProfile) GetCreatedBy() float32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DeviceConfigurationTemplateProfile) GetCreatedByOk() (*float32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DeviceConfigurationTemplateProfile) SetCreatedBy(v float32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DeviceConfigurationTemplateProfile) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *DeviceConfigurationTemplateProfile) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *DeviceConfigurationTemplateProfile) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *DeviceConfigurationTemplateProfile) GetUpdatedBy() float32`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DeviceConfigurationTemplateProfile) GetUpdatedByOk() (*float32, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DeviceConfigurationTemplateProfile) SetUpdatedBy(v float32)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DeviceConfigurationTemplateProfile) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *DeviceConfigurationTemplateProfile) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *DeviceConfigurationTemplateProfile) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


