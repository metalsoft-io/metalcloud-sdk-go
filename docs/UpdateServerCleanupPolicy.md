# UpdateServerCleanupPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Server cleanup policy label | [optional] 
**CleanupDrivesForOobEnabledServer** | Pointer to **float32** | Cleanup drives for oob enabled server | [optional] 
**RecreateRaid** | Pointer to **float32** | Recreate raid | [optional] 
**ResetRaidControllers** | Pointer to **float32** | Reset raid controllers to default | [optional] 
**DisableEmbeddedNics** | Pointer to **float32** | Disable embedded nics | [optional] 
**RaidOneDrive** | Pointer to **string** | Raid one drive | [optional] 
**RaidTwoDrives** | Pointer to **string** | Raid two drives | [optional] 
**RaidEvenNumberMoreThanTwoDrives** | Pointer to **string** | Raid even number more than two drives | [optional] 
**RaidOddNumberMoreThanOneDrive** | Pointer to **string** | Raid odd number more than one drive | [optional] 
**SkipRaidActions** | Pointer to **[]string** | Skip raid actions | [optional] 

## Methods

### NewUpdateServerCleanupPolicy

`func NewUpdateServerCleanupPolicy() *UpdateServerCleanupPolicy`

NewUpdateServerCleanupPolicy instantiates a new UpdateServerCleanupPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServerCleanupPolicyWithDefaults

`func NewUpdateServerCleanupPolicyWithDefaults() *UpdateServerCleanupPolicy`

NewUpdateServerCleanupPolicyWithDefaults instantiates a new UpdateServerCleanupPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *UpdateServerCleanupPolicy) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateServerCleanupPolicy) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateServerCleanupPolicy) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpdateServerCleanupPolicy) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetCleanupDrivesForOobEnabledServer

`func (o *UpdateServerCleanupPolicy) GetCleanupDrivesForOobEnabledServer() float32`

GetCleanupDrivesForOobEnabledServer returns the CleanupDrivesForOobEnabledServer field if non-nil, zero value otherwise.

### GetCleanupDrivesForOobEnabledServerOk

`func (o *UpdateServerCleanupPolicy) GetCleanupDrivesForOobEnabledServerOk() (*float32, bool)`

GetCleanupDrivesForOobEnabledServerOk returns a tuple with the CleanupDrivesForOobEnabledServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupDrivesForOobEnabledServer

`func (o *UpdateServerCleanupPolicy) SetCleanupDrivesForOobEnabledServer(v float32)`

SetCleanupDrivesForOobEnabledServer sets CleanupDrivesForOobEnabledServer field to given value.

### HasCleanupDrivesForOobEnabledServer

`func (o *UpdateServerCleanupPolicy) HasCleanupDrivesForOobEnabledServer() bool`

HasCleanupDrivesForOobEnabledServer returns a boolean if a field has been set.

### GetRecreateRaid

`func (o *UpdateServerCleanupPolicy) GetRecreateRaid() float32`

GetRecreateRaid returns the RecreateRaid field if non-nil, zero value otherwise.

### GetRecreateRaidOk

`func (o *UpdateServerCleanupPolicy) GetRecreateRaidOk() (*float32, bool)`

GetRecreateRaidOk returns a tuple with the RecreateRaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecreateRaid

`func (o *UpdateServerCleanupPolicy) SetRecreateRaid(v float32)`

SetRecreateRaid sets RecreateRaid field to given value.

### HasRecreateRaid

`func (o *UpdateServerCleanupPolicy) HasRecreateRaid() bool`

HasRecreateRaid returns a boolean if a field has been set.

### GetResetRaidControllers

`func (o *UpdateServerCleanupPolicy) GetResetRaidControllers() float32`

GetResetRaidControllers returns the ResetRaidControllers field if non-nil, zero value otherwise.

### GetResetRaidControllersOk

`func (o *UpdateServerCleanupPolicy) GetResetRaidControllersOk() (*float32, bool)`

GetResetRaidControllersOk returns a tuple with the ResetRaidControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetRaidControllers

`func (o *UpdateServerCleanupPolicy) SetResetRaidControllers(v float32)`

SetResetRaidControllers sets ResetRaidControllers field to given value.

### HasResetRaidControllers

`func (o *UpdateServerCleanupPolicy) HasResetRaidControllers() bool`

HasResetRaidControllers returns a boolean if a field has been set.

### GetDisableEmbeddedNics

`func (o *UpdateServerCleanupPolicy) GetDisableEmbeddedNics() float32`

GetDisableEmbeddedNics returns the DisableEmbeddedNics field if non-nil, zero value otherwise.

### GetDisableEmbeddedNicsOk

`func (o *UpdateServerCleanupPolicy) GetDisableEmbeddedNicsOk() (*float32, bool)`

GetDisableEmbeddedNicsOk returns a tuple with the DisableEmbeddedNics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEmbeddedNics

`func (o *UpdateServerCleanupPolicy) SetDisableEmbeddedNics(v float32)`

SetDisableEmbeddedNics sets DisableEmbeddedNics field to given value.

### HasDisableEmbeddedNics

`func (o *UpdateServerCleanupPolicy) HasDisableEmbeddedNics() bool`

HasDisableEmbeddedNics returns a boolean if a field has been set.

### GetRaidOneDrive

`func (o *UpdateServerCleanupPolicy) GetRaidOneDrive() string`

GetRaidOneDrive returns the RaidOneDrive field if non-nil, zero value otherwise.

### GetRaidOneDriveOk

`func (o *UpdateServerCleanupPolicy) GetRaidOneDriveOk() (*string, bool)`

GetRaidOneDriveOk returns a tuple with the RaidOneDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidOneDrive

`func (o *UpdateServerCleanupPolicy) SetRaidOneDrive(v string)`

SetRaidOneDrive sets RaidOneDrive field to given value.

### HasRaidOneDrive

`func (o *UpdateServerCleanupPolicy) HasRaidOneDrive() bool`

HasRaidOneDrive returns a boolean if a field has been set.

### GetRaidTwoDrives

`func (o *UpdateServerCleanupPolicy) GetRaidTwoDrives() string`

GetRaidTwoDrives returns the RaidTwoDrives field if non-nil, zero value otherwise.

### GetRaidTwoDrivesOk

`func (o *UpdateServerCleanupPolicy) GetRaidTwoDrivesOk() (*string, bool)`

GetRaidTwoDrivesOk returns a tuple with the RaidTwoDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidTwoDrives

`func (o *UpdateServerCleanupPolicy) SetRaidTwoDrives(v string)`

SetRaidTwoDrives sets RaidTwoDrives field to given value.

### HasRaidTwoDrives

`func (o *UpdateServerCleanupPolicy) HasRaidTwoDrives() bool`

HasRaidTwoDrives returns a boolean if a field has been set.

### GetRaidEvenNumberMoreThanTwoDrives

`func (o *UpdateServerCleanupPolicy) GetRaidEvenNumberMoreThanTwoDrives() string`

GetRaidEvenNumberMoreThanTwoDrives returns the RaidEvenNumberMoreThanTwoDrives field if non-nil, zero value otherwise.

### GetRaidEvenNumberMoreThanTwoDrivesOk

`func (o *UpdateServerCleanupPolicy) GetRaidEvenNumberMoreThanTwoDrivesOk() (*string, bool)`

GetRaidEvenNumberMoreThanTwoDrivesOk returns a tuple with the RaidEvenNumberMoreThanTwoDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidEvenNumberMoreThanTwoDrives

`func (o *UpdateServerCleanupPolicy) SetRaidEvenNumberMoreThanTwoDrives(v string)`

SetRaidEvenNumberMoreThanTwoDrives sets RaidEvenNumberMoreThanTwoDrives field to given value.

### HasRaidEvenNumberMoreThanTwoDrives

`func (o *UpdateServerCleanupPolicy) HasRaidEvenNumberMoreThanTwoDrives() bool`

HasRaidEvenNumberMoreThanTwoDrives returns a boolean if a field has been set.

### GetRaidOddNumberMoreThanOneDrive

`func (o *UpdateServerCleanupPolicy) GetRaidOddNumberMoreThanOneDrive() string`

GetRaidOddNumberMoreThanOneDrive returns the RaidOddNumberMoreThanOneDrive field if non-nil, zero value otherwise.

### GetRaidOddNumberMoreThanOneDriveOk

`func (o *UpdateServerCleanupPolicy) GetRaidOddNumberMoreThanOneDriveOk() (*string, bool)`

GetRaidOddNumberMoreThanOneDriveOk returns a tuple with the RaidOddNumberMoreThanOneDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidOddNumberMoreThanOneDrive

`func (o *UpdateServerCleanupPolicy) SetRaidOddNumberMoreThanOneDrive(v string)`

SetRaidOddNumberMoreThanOneDrive sets RaidOddNumberMoreThanOneDrive field to given value.

### HasRaidOddNumberMoreThanOneDrive

`func (o *UpdateServerCleanupPolicy) HasRaidOddNumberMoreThanOneDrive() bool`

HasRaidOddNumberMoreThanOneDrive returns a boolean if a field has been set.

### GetSkipRaidActions

`func (o *UpdateServerCleanupPolicy) GetSkipRaidActions() []string`

GetSkipRaidActions returns the SkipRaidActions field if non-nil, zero value otherwise.

### GetSkipRaidActionsOk

`func (o *UpdateServerCleanupPolicy) GetSkipRaidActionsOk() (*[]string, bool)`

GetSkipRaidActionsOk returns a tuple with the SkipRaidActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipRaidActions

`func (o *UpdateServerCleanupPolicy) SetSkipRaidActions(v []string)`

SetSkipRaidActions sets SkipRaidActions field to given value.

### HasSkipRaidActions

`func (o *UpdateServerCleanupPolicy) HasSkipRaidActions() bool`

HasSkipRaidActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


