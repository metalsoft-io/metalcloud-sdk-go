# CreateServerCleanupPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Server cleanup policy label | 
**CleanupDrivesForOobEnabledServer** | **float32** | Cleanup drives for oob enabled server | 
**RecreateRaid** | **float32** | Recreate raid | 
**DisableEmbeddedNics** | **float32** | Disable embedded nics | 
**RaidOneDrive** | **string** | Raid one drive | 
**RaidTwoDrives** | **string** | Raid two drives | 
**RaidEvenNumberMoreThanTwoDrives** | **string** | Raid even number more than two drives | 
**RaidOddNumberMoreThanOneDrive** | **string** | Raid odd number more than one drive | 
**SkipRaidActions** | **[]string** | Skip raid actions | 

## Methods

### NewCreateServerCleanupPolicy

`func NewCreateServerCleanupPolicy(label string, cleanupDrivesForOobEnabledServer float32, recreateRaid float32, disableEmbeddedNics float32, raidOneDrive string, raidTwoDrives string, raidEvenNumberMoreThanTwoDrives string, raidOddNumberMoreThanOneDrive string, skipRaidActions []string, ) *CreateServerCleanupPolicy`

NewCreateServerCleanupPolicy instantiates a new CreateServerCleanupPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServerCleanupPolicyWithDefaults

`func NewCreateServerCleanupPolicyWithDefaults() *CreateServerCleanupPolicy`

NewCreateServerCleanupPolicyWithDefaults instantiates a new CreateServerCleanupPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CreateServerCleanupPolicy) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateServerCleanupPolicy) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateServerCleanupPolicy) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetCleanupDrivesForOobEnabledServer

`func (o *CreateServerCleanupPolicy) GetCleanupDrivesForOobEnabledServer() float32`

GetCleanupDrivesForOobEnabledServer returns the CleanupDrivesForOobEnabledServer field if non-nil, zero value otherwise.

### GetCleanupDrivesForOobEnabledServerOk

`func (o *CreateServerCleanupPolicy) GetCleanupDrivesForOobEnabledServerOk() (*float32, bool)`

GetCleanupDrivesForOobEnabledServerOk returns a tuple with the CleanupDrivesForOobEnabledServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupDrivesForOobEnabledServer

`func (o *CreateServerCleanupPolicy) SetCleanupDrivesForOobEnabledServer(v float32)`

SetCleanupDrivesForOobEnabledServer sets CleanupDrivesForOobEnabledServer field to given value.


### GetRecreateRaid

`func (o *CreateServerCleanupPolicy) GetRecreateRaid() float32`

GetRecreateRaid returns the RecreateRaid field if non-nil, zero value otherwise.

### GetRecreateRaidOk

`func (o *CreateServerCleanupPolicy) GetRecreateRaidOk() (*float32, bool)`

GetRecreateRaidOk returns a tuple with the RecreateRaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecreateRaid

`func (o *CreateServerCleanupPolicy) SetRecreateRaid(v float32)`

SetRecreateRaid sets RecreateRaid field to given value.


### GetDisableEmbeddedNics

`func (o *CreateServerCleanupPolicy) GetDisableEmbeddedNics() float32`

GetDisableEmbeddedNics returns the DisableEmbeddedNics field if non-nil, zero value otherwise.

### GetDisableEmbeddedNicsOk

`func (o *CreateServerCleanupPolicy) GetDisableEmbeddedNicsOk() (*float32, bool)`

GetDisableEmbeddedNicsOk returns a tuple with the DisableEmbeddedNics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEmbeddedNics

`func (o *CreateServerCleanupPolicy) SetDisableEmbeddedNics(v float32)`

SetDisableEmbeddedNics sets DisableEmbeddedNics field to given value.


### GetRaidOneDrive

`func (o *CreateServerCleanupPolicy) GetRaidOneDrive() string`

GetRaidOneDrive returns the RaidOneDrive field if non-nil, zero value otherwise.

### GetRaidOneDriveOk

`func (o *CreateServerCleanupPolicy) GetRaidOneDriveOk() (*string, bool)`

GetRaidOneDriveOk returns a tuple with the RaidOneDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidOneDrive

`func (o *CreateServerCleanupPolicy) SetRaidOneDrive(v string)`

SetRaidOneDrive sets RaidOneDrive field to given value.


### GetRaidTwoDrives

`func (o *CreateServerCleanupPolicy) GetRaidTwoDrives() string`

GetRaidTwoDrives returns the RaidTwoDrives field if non-nil, zero value otherwise.

### GetRaidTwoDrivesOk

`func (o *CreateServerCleanupPolicy) GetRaidTwoDrivesOk() (*string, bool)`

GetRaidTwoDrivesOk returns a tuple with the RaidTwoDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidTwoDrives

`func (o *CreateServerCleanupPolicy) SetRaidTwoDrives(v string)`

SetRaidTwoDrives sets RaidTwoDrives field to given value.


### GetRaidEvenNumberMoreThanTwoDrives

`func (o *CreateServerCleanupPolicy) GetRaidEvenNumberMoreThanTwoDrives() string`

GetRaidEvenNumberMoreThanTwoDrives returns the RaidEvenNumberMoreThanTwoDrives field if non-nil, zero value otherwise.

### GetRaidEvenNumberMoreThanTwoDrivesOk

`func (o *CreateServerCleanupPolicy) GetRaidEvenNumberMoreThanTwoDrivesOk() (*string, bool)`

GetRaidEvenNumberMoreThanTwoDrivesOk returns a tuple with the RaidEvenNumberMoreThanTwoDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidEvenNumberMoreThanTwoDrives

`func (o *CreateServerCleanupPolicy) SetRaidEvenNumberMoreThanTwoDrives(v string)`

SetRaidEvenNumberMoreThanTwoDrives sets RaidEvenNumberMoreThanTwoDrives field to given value.


### GetRaidOddNumberMoreThanOneDrive

`func (o *CreateServerCleanupPolicy) GetRaidOddNumberMoreThanOneDrive() string`

GetRaidOddNumberMoreThanOneDrive returns the RaidOddNumberMoreThanOneDrive field if non-nil, zero value otherwise.

### GetRaidOddNumberMoreThanOneDriveOk

`func (o *CreateServerCleanupPolicy) GetRaidOddNumberMoreThanOneDriveOk() (*string, bool)`

GetRaidOddNumberMoreThanOneDriveOk returns a tuple with the RaidOddNumberMoreThanOneDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidOddNumberMoreThanOneDrive

`func (o *CreateServerCleanupPolicy) SetRaidOddNumberMoreThanOneDrive(v string)`

SetRaidOddNumberMoreThanOneDrive sets RaidOddNumberMoreThanOneDrive field to given value.


### GetSkipRaidActions

`func (o *CreateServerCleanupPolicy) GetSkipRaidActions() []string`

GetSkipRaidActions returns the SkipRaidActions field if non-nil, zero value otherwise.

### GetSkipRaidActionsOk

`func (o *CreateServerCleanupPolicy) GetSkipRaidActionsOk() (*[]string, bool)`

GetSkipRaidActionsOk returns a tuple with the SkipRaidActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipRaidActions

`func (o *CreateServerCleanupPolicy) SetSkipRaidActions(v []string)`

SetSkipRaidActions sets SkipRaidActions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


