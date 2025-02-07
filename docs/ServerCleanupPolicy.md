# ServerCleanupPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | Server cleanup policy id | 
**Label** | **string** | Server cleanup policy label | 
**CleanupDrivesForOobEnabledServer** | **float32** | Cleanup drives for oob enabled server | 
**RecreateRaid** | **float32** | Recreate raid | 
**DisableEmbeddedNics** | **float32** | Disable embedded nics | 
**RaidOneDrive** | **string** | Raid one drive | 
**RaidTwoDrives** | **string** | Raid two drives | 
**RaidEvenNumberMoreThanTwoDrives** | **string** | Raid even number more than two drives | 
**RaidOddNumberMoreThanOneDrive** | **string** | Raid odd number more than one drive | 
**CreatedTimestamp** | **string** | Created timestamp | 
**UpdatedTimestamp** | **string** | Updated timestamp | 
**SkipRaidActions** | **[]string** | Skip raid actions | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewServerCleanupPolicy

`func NewServerCleanupPolicy(id float32, label string, cleanupDrivesForOobEnabledServer float32, recreateRaid float32, disableEmbeddedNics float32, raidOneDrive string, raidTwoDrives string, raidEvenNumberMoreThanTwoDrives string, raidOddNumberMoreThanOneDrive string, createdTimestamp string, updatedTimestamp string, skipRaidActions []string, ) *ServerCleanupPolicy`

NewServerCleanupPolicy instantiates a new ServerCleanupPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerCleanupPolicyWithDefaults

`func NewServerCleanupPolicyWithDefaults() *ServerCleanupPolicy`

NewServerCleanupPolicyWithDefaults instantiates a new ServerCleanupPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerCleanupPolicy) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerCleanupPolicy) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerCleanupPolicy) SetId(v float32)`

SetId sets Id field to given value.


### GetLabel

`func (o *ServerCleanupPolicy) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServerCleanupPolicy) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServerCleanupPolicy) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetCleanupDrivesForOobEnabledServer

`func (o *ServerCleanupPolicy) GetCleanupDrivesForOobEnabledServer() float32`

GetCleanupDrivesForOobEnabledServer returns the CleanupDrivesForOobEnabledServer field if non-nil, zero value otherwise.

### GetCleanupDrivesForOobEnabledServerOk

`func (o *ServerCleanupPolicy) GetCleanupDrivesForOobEnabledServerOk() (*float32, bool)`

GetCleanupDrivesForOobEnabledServerOk returns a tuple with the CleanupDrivesForOobEnabledServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupDrivesForOobEnabledServer

`func (o *ServerCleanupPolicy) SetCleanupDrivesForOobEnabledServer(v float32)`

SetCleanupDrivesForOobEnabledServer sets CleanupDrivesForOobEnabledServer field to given value.


### GetRecreateRaid

`func (o *ServerCleanupPolicy) GetRecreateRaid() float32`

GetRecreateRaid returns the RecreateRaid field if non-nil, zero value otherwise.

### GetRecreateRaidOk

`func (o *ServerCleanupPolicy) GetRecreateRaidOk() (*float32, bool)`

GetRecreateRaidOk returns a tuple with the RecreateRaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecreateRaid

`func (o *ServerCleanupPolicy) SetRecreateRaid(v float32)`

SetRecreateRaid sets RecreateRaid field to given value.


### GetDisableEmbeddedNics

`func (o *ServerCleanupPolicy) GetDisableEmbeddedNics() float32`

GetDisableEmbeddedNics returns the DisableEmbeddedNics field if non-nil, zero value otherwise.

### GetDisableEmbeddedNicsOk

`func (o *ServerCleanupPolicy) GetDisableEmbeddedNicsOk() (*float32, bool)`

GetDisableEmbeddedNicsOk returns a tuple with the DisableEmbeddedNics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEmbeddedNics

`func (o *ServerCleanupPolicy) SetDisableEmbeddedNics(v float32)`

SetDisableEmbeddedNics sets DisableEmbeddedNics field to given value.


### GetRaidOneDrive

`func (o *ServerCleanupPolicy) GetRaidOneDrive() string`

GetRaidOneDrive returns the RaidOneDrive field if non-nil, zero value otherwise.

### GetRaidOneDriveOk

`func (o *ServerCleanupPolicy) GetRaidOneDriveOk() (*string, bool)`

GetRaidOneDriveOk returns a tuple with the RaidOneDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidOneDrive

`func (o *ServerCleanupPolicy) SetRaidOneDrive(v string)`

SetRaidOneDrive sets RaidOneDrive field to given value.


### GetRaidTwoDrives

`func (o *ServerCleanupPolicy) GetRaidTwoDrives() string`

GetRaidTwoDrives returns the RaidTwoDrives field if non-nil, zero value otherwise.

### GetRaidTwoDrivesOk

`func (o *ServerCleanupPolicy) GetRaidTwoDrivesOk() (*string, bool)`

GetRaidTwoDrivesOk returns a tuple with the RaidTwoDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidTwoDrives

`func (o *ServerCleanupPolicy) SetRaidTwoDrives(v string)`

SetRaidTwoDrives sets RaidTwoDrives field to given value.


### GetRaidEvenNumberMoreThanTwoDrives

`func (o *ServerCleanupPolicy) GetRaidEvenNumberMoreThanTwoDrives() string`

GetRaidEvenNumberMoreThanTwoDrives returns the RaidEvenNumberMoreThanTwoDrives field if non-nil, zero value otherwise.

### GetRaidEvenNumberMoreThanTwoDrivesOk

`func (o *ServerCleanupPolicy) GetRaidEvenNumberMoreThanTwoDrivesOk() (*string, bool)`

GetRaidEvenNumberMoreThanTwoDrivesOk returns a tuple with the RaidEvenNumberMoreThanTwoDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidEvenNumberMoreThanTwoDrives

`func (o *ServerCleanupPolicy) SetRaidEvenNumberMoreThanTwoDrives(v string)`

SetRaidEvenNumberMoreThanTwoDrives sets RaidEvenNumberMoreThanTwoDrives field to given value.


### GetRaidOddNumberMoreThanOneDrive

`func (o *ServerCleanupPolicy) GetRaidOddNumberMoreThanOneDrive() string`

GetRaidOddNumberMoreThanOneDrive returns the RaidOddNumberMoreThanOneDrive field if non-nil, zero value otherwise.

### GetRaidOddNumberMoreThanOneDriveOk

`func (o *ServerCleanupPolicy) GetRaidOddNumberMoreThanOneDriveOk() (*string, bool)`

GetRaidOddNumberMoreThanOneDriveOk returns a tuple with the RaidOddNumberMoreThanOneDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidOddNumberMoreThanOneDrive

`func (o *ServerCleanupPolicy) SetRaidOddNumberMoreThanOneDrive(v string)`

SetRaidOddNumberMoreThanOneDrive sets RaidOddNumberMoreThanOneDrive field to given value.


### GetCreatedTimestamp

`func (o *ServerCleanupPolicy) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *ServerCleanupPolicy) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *ServerCleanupPolicy) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *ServerCleanupPolicy) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *ServerCleanupPolicy) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *ServerCleanupPolicy) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetSkipRaidActions

`func (o *ServerCleanupPolicy) GetSkipRaidActions() []string`

GetSkipRaidActions returns the SkipRaidActions field if non-nil, zero value otherwise.

### GetSkipRaidActionsOk

`func (o *ServerCleanupPolicy) GetSkipRaidActionsOk() (*[]string, bool)`

GetSkipRaidActionsOk returns a tuple with the SkipRaidActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipRaidActions

`func (o *ServerCleanupPolicy) SetSkipRaidActions(v []string)`

SetSkipRaidActions sets SkipRaidActions field to given value.


### GetLinks

`func (o *ServerCleanupPolicy) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServerCleanupPolicy) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServerCleanupPolicy) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServerCleanupPolicy) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


