# FindApplicableDeviceConfigurationTemplateProfiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkDeviceId** | Pointer to **int64** | NetworkDevice id to match against. Returns profiles bound to this device (profile.networkDeviceId &#x3D; id) plus fabric-bound profiles (profile.networkDeviceId IS NULL) for every fabric the device is attached to. Mutually exclusive with networkFabricId — exactly one of the two must be defined (enforced at the service layer). | [optional] 
**NetworkFabricId** | Pointer to **int64** | NetworkFabric id to match against. Returns fabric-bound profiles (profile.networkFabricId &#x3D; id AND profile.networkDeviceId IS NULL) plus device-bound profiles whose target device belongs to this fabric. Mutually exclusive with networkDeviceId — exactly one of the two must be defined (enforced at the service layer). | [optional] 
**LifecycleStage** | Pointer to [**DeviceConfigurationProfileLifecycleStage**](DeviceConfigurationProfileLifecycleStage.md) | Optional lifecycle-stage filter. When omitted, profiles of every stage are returned. | [optional] 
**IncludeDisabled** | Pointer to **bool** | When true, profiles with isEnabled&#x3D;false are included. Defaults to enabled-only. | [optional] [default to false]

## Methods

### NewFindApplicableDeviceConfigurationTemplateProfiles

`func NewFindApplicableDeviceConfigurationTemplateProfiles() *FindApplicableDeviceConfigurationTemplateProfiles`

NewFindApplicableDeviceConfigurationTemplateProfiles instantiates a new FindApplicableDeviceConfigurationTemplateProfiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindApplicableDeviceConfigurationTemplateProfilesWithDefaults

`func NewFindApplicableDeviceConfigurationTemplateProfilesWithDefaults() *FindApplicableDeviceConfigurationTemplateProfiles`

NewFindApplicableDeviceConfigurationTemplateProfilesWithDefaults instantiates a new FindApplicableDeviceConfigurationTemplateProfiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkDeviceId

`func (o *FindApplicableDeviceConfigurationTemplateProfiles) GetNetworkDeviceId() int64`

GetNetworkDeviceId returns the NetworkDeviceId field if non-nil, zero value otherwise.

### GetNetworkDeviceIdOk

`func (o *FindApplicableDeviceConfigurationTemplateProfiles) GetNetworkDeviceIdOk() (*int64, bool)`

GetNetworkDeviceIdOk returns a tuple with the NetworkDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceId

`func (o *FindApplicableDeviceConfigurationTemplateProfiles) SetNetworkDeviceId(v int64)`

SetNetworkDeviceId sets NetworkDeviceId field to given value.

### HasNetworkDeviceId

`func (o *FindApplicableDeviceConfigurationTemplateProfiles) HasNetworkDeviceId() bool`

HasNetworkDeviceId returns a boolean if a field has been set.

### GetNetworkFabricId

`func (o *FindApplicableDeviceConfigurationTemplateProfiles) GetNetworkFabricId() int64`

GetNetworkFabricId returns the NetworkFabricId field if non-nil, zero value otherwise.

### GetNetworkFabricIdOk

`func (o *FindApplicableDeviceConfigurationTemplateProfiles) GetNetworkFabricIdOk() (*int64, bool)`

GetNetworkFabricIdOk returns a tuple with the NetworkFabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFabricId

`func (o *FindApplicableDeviceConfigurationTemplateProfiles) SetNetworkFabricId(v int64)`

SetNetworkFabricId sets NetworkFabricId field to given value.

### HasNetworkFabricId

`func (o *FindApplicableDeviceConfigurationTemplateProfiles) HasNetworkFabricId() bool`

HasNetworkFabricId returns a boolean if a field has been set.

### GetLifecycleStage

`func (o *FindApplicableDeviceConfigurationTemplateProfiles) GetLifecycleStage() DeviceConfigurationProfileLifecycleStage`

GetLifecycleStage returns the LifecycleStage field if non-nil, zero value otherwise.

### GetLifecycleStageOk

`func (o *FindApplicableDeviceConfigurationTemplateProfiles) GetLifecycleStageOk() (*DeviceConfigurationProfileLifecycleStage, bool)`

GetLifecycleStageOk returns a tuple with the LifecycleStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleStage

`func (o *FindApplicableDeviceConfigurationTemplateProfiles) SetLifecycleStage(v DeviceConfigurationProfileLifecycleStage)`

SetLifecycleStage sets LifecycleStage field to given value.

### HasLifecycleStage

`func (o *FindApplicableDeviceConfigurationTemplateProfiles) HasLifecycleStage() bool`

HasLifecycleStage returns a boolean if a field has been set.

### GetIncludeDisabled

`func (o *FindApplicableDeviceConfigurationTemplateProfiles) GetIncludeDisabled() bool`

GetIncludeDisabled returns the IncludeDisabled field if non-nil, zero value otherwise.

### GetIncludeDisabledOk

`func (o *FindApplicableDeviceConfigurationTemplateProfiles) GetIncludeDisabledOk() (*bool, bool)`

GetIncludeDisabledOk returns a tuple with the IncludeDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDisabled

`func (o *FindApplicableDeviceConfigurationTemplateProfiles) SetIncludeDisabled(v bool)`

SetIncludeDisabled sets IncludeDisabled field to given value.

### HasIncludeDisabled

`func (o *FindApplicableDeviceConfigurationTemplateProfiles) HasIncludeDisabled() bool`

HasIncludeDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


