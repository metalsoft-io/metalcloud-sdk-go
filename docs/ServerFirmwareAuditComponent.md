# ServerFirmwareAuditComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentId** | **int64** | The id of the server component. | 
**Name** | **string** | The name of the server component. | 
**Type** | Pointer to **string** | The type of the server component. | [optional] 
**CurrentVersion** | Pointer to **string** | The firmware version currently installed. | [optional] 
**TargetVersion** | Pointer to **string** | The firmware version the component is currently set to be upgraded to, if any. | [optional] 
**LatestVersion** | Pointer to **string** | The highest version available in the baseline for this component. | [optional] 
**Status** | **string** | Summary of how the installed version compares to the versions available in the baseline. | 
**MinimumRequiredVersion** | Pointer to **string** | The minimum version the baseline requires for this component, when it defines one. | [optional] 
**BelowMinimumVersion** | Pointer to **bool** | True when the installed version is below the baseline minimum. Absent when the baseline defines no minimum for this component. | [optional] 
**Upgrades** | [**[]ServerFirmwareAuditVersionOption**](ServerFirmwareAuditVersionOption.md) | Versions newer than the installed one, newest first. | 
**Downgrades** | [**[]ServerFirmwareAuditVersionOption**](ServerFirmwareAuditVersionOption.md) | Versions older than the installed one, newest first. | 

## Methods

### NewServerFirmwareAuditComponent

`func NewServerFirmwareAuditComponent(componentId int64, name string, status string, upgrades []ServerFirmwareAuditVersionOption, downgrades []ServerFirmwareAuditVersionOption, ) *ServerFirmwareAuditComponent`

NewServerFirmwareAuditComponent instantiates a new ServerFirmwareAuditComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerFirmwareAuditComponentWithDefaults

`func NewServerFirmwareAuditComponentWithDefaults() *ServerFirmwareAuditComponent`

NewServerFirmwareAuditComponentWithDefaults instantiates a new ServerFirmwareAuditComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentId

`func (o *ServerFirmwareAuditComponent) GetComponentId() int64`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *ServerFirmwareAuditComponent) GetComponentIdOk() (*int64, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *ServerFirmwareAuditComponent) SetComponentId(v int64)`

SetComponentId sets ComponentId field to given value.


### GetName

`func (o *ServerFirmwareAuditComponent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerFirmwareAuditComponent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerFirmwareAuditComponent) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ServerFirmwareAuditComponent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServerFirmwareAuditComponent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServerFirmwareAuditComponent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ServerFirmwareAuditComponent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCurrentVersion

`func (o *ServerFirmwareAuditComponent) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *ServerFirmwareAuditComponent) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *ServerFirmwareAuditComponent) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *ServerFirmwareAuditComponent) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetTargetVersion

`func (o *ServerFirmwareAuditComponent) GetTargetVersion() string`

GetTargetVersion returns the TargetVersion field if non-nil, zero value otherwise.

### GetTargetVersionOk

`func (o *ServerFirmwareAuditComponent) GetTargetVersionOk() (*string, bool)`

GetTargetVersionOk returns a tuple with the TargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersion

`func (o *ServerFirmwareAuditComponent) SetTargetVersion(v string)`

SetTargetVersion sets TargetVersion field to given value.

### HasTargetVersion

`func (o *ServerFirmwareAuditComponent) HasTargetVersion() bool`

HasTargetVersion returns a boolean if a field has been set.

### GetLatestVersion

`func (o *ServerFirmwareAuditComponent) GetLatestVersion() string`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *ServerFirmwareAuditComponent) GetLatestVersionOk() (*string, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *ServerFirmwareAuditComponent) SetLatestVersion(v string)`

SetLatestVersion sets LatestVersion field to given value.

### HasLatestVersion

`func (o *ServerFirmwareAuditComponent) HasLatestVersion() bool`

HasLatestVersion returns a boolean if a field has been set.

### GetStatus

`func (o *ServerFirmwareAuditComponent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServerFirmwareAuditComponent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServerFirmwareAuditComponent) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMinimumRequiredVersion

`func (o *ServerFirmwareAuditComponent) GetMinimumRequiredVersion() string`

GetMinimumRequiredVersion returns the MinimumRequiredVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredVersionOk

`func (o *ServerFirmwareAuditComponent) GetMinimumRequiredVersionOk() (*string, bool)`

GetMinimumRequiredVersionOk returns a tuple with the MinimumRequiredVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumRequiredVersion

`func (o *ServerFirmwareAuditComponent) SetMinimumRequiredVersion(v string)`

SetMinimumRequiredVersion sets MinimumRequiredVersion field to given value.

### HasMinimumRequiredVersion

`func (o *ServerFirmwareAuditComponent) HasMinimumRequiredVersion() bool`

HasMinimumRequiredVersion returns a boolean if a field has been set.

### GetBelowMinimumVersion

`func (o *ServerFirmwareAuditComponent) GetBelowMinimumVersion() bool`

GetBelowMinimumVersion returns the BelowMinimumVersion field if non-nil, zero value otherwise.

### GetBelowMinimumVersionOk

`func (o *ServerFirmwareAuditComponent) GetBelowMinimumVersionOk() (*bool, bool)`

GetBelowMinimumVersionOk returns a tuple with the BelowMinimumVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBelowMinimumVersion

`func (o *ServerFirmwareAuditComponent) SetBelowMinimumVersion(v bool)`

SetBelowMinimumVersion sets BelowMinimumVersion field to given value.

### HasBelowMinimumVersion

`func (o *ServerFirmwareAuditComponent) HasBelowMinimumVersion() bool`

HasBelowMinimumVersion returns a boolean if a field has been set.

### GetUpgrades

`func (o *ServerFirmwareAuditComponent) GetUpgrades() []ServerFirmwareAuditVersionOption`

GetUpgrades returns the Upgrades field if non-nil, zero value otherwise.

### GetUpgradesOk

`func (o *ServerFirmwareAuditComponent) GetUpgradesOk() (*[]ServerFirmwareAuditVersionOption, bool)`

GetUpgradesOk returns a tuple with the Upgrades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgrades

`func (o *ServerFirmwareAuditComponent) SetUpgrades(v []ServerFirmwareAuditVersionOption)`

SetUpgrades sets Upgrades field to given value.


### GetDowngrades

`func (o *ServerFirmwareAuditComponent) GetDowngrades() []ServerFirmwareAuditVersionOption`

GetDowngrades returns the Downgrades field if non-nil, zero value otherwise.

### GetDowngradesOk

`func (o *ServerFirmwareAuditComponent) GetDowngradesOk() (*[]ServerFirmwareAuditVersionOption, bool)`

GetDowngradesOk returns a tuple with the Downgrades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowngrades

`func (o *ServerFirmwareAuditComponent) SetDowngrades(v []ServerFirmwareAuditVersionOption)`

SetDowngrades sets Downgrades field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


