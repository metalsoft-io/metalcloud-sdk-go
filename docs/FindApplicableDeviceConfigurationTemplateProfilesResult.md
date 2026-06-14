# FindApplicableDeviceConfigurationTemplateProfilesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]DeviceConfigurationTemplateProfile**](DeviceConfigurationTemplateProfile.md) | Applicable profiles that would be applied on the next deploy, ordered by priority (ASC) then profile id (ASC). Excludes applyMode&#x3D;&#39;once&#39; profiles already recorded in device_configuration_template_history for the target device. | 
**AlreadyApplied** | [**[]DeviceConfigurationTemplateProfile**](DeviceConfigurationTemplateProfile.md) | Profiles that match the scope but are skipped on the next deploy because they are applyMode&#x3D;&#39;once&#39; and already have a device_configuration_template_history row for the target device. Always empty when the query is fabric-scoped (history is per-(profile, device); the orchestrator runs the per-device split itself). | 

## Methods

### NewFindApplicableDeviceConfigurationTemplateProfilesResult

`func NewFindApplicableDeviceConfigurationTemplateProfilesResult(items []DeviceConfigurationTemplateProfile, alreadyApplied []DeviceConfigurationTemplateProfile, ) *FindApplicableDeviceConfigurationTemplateProfilesResult`

NewFindApplicableDeviceConfigurationTemplateProfilesResult instantiates a new FindApplicableDeviceConfigurationTemplateProfilesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindApplicableDeviceConfigurationTemplateProfilesResultWithDefaults

`func NewFindApplicableDeviceConfigurationTemplateProfilesResultWithDefaults() *FindApplicableDeviceConfigurationTemplateProfilesResult`

NewFindApplicableDeviceConfigurationTemplateProfilesResultWithDefaults instantiates a new FindApplicableDeviceConfigurationTemplateProfilesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *FindApplicableDeviceConfigurationTemplateProfilesResult) GetItems() []DeviceConfigurationTemplateProfile`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *FindApplicableDeviceConfigurationTemplateProfilesResult) GetItemsOk() (*[]DeviceConfigurationTemplateProfile, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *FindApplicableDeviceConfigurationTemplateProfilesResult) SetItems(v []DeviceConfigurationTemplateProfile)`

SetItems sets Items field to given value.


### GetAlreadyApplied

`func (o *FindApplicableDeviceConfigurationTemplateProfilesResult) GetAlreadyApplied() []DeviceConfigurationTemplateProfile`

GetAlreadyApplied returns the AlreadyApplied field if non-nil, zero value otherwise.

### GetAlreadyAppliedOk

`func (o *FindApplicableDeviceConfigurationTemplateProfilesResult) GetAlreadyAppliedOk() (*[]DeviceConfigurationTemplateProfile, bool)`

GetAlreadyAppliedOk returns a tuple with the AlreadyApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlreadyApplied

`func (o *FindApplicableDeviceConfigurationTemplateProfilesResult) SetAlreadyApplied(v []DeviceConfigurationTemplateProfile)`

SetAlreadyApplied sets AlreadyApplied field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


