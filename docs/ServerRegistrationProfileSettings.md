# ServerRegistrationProfileSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegisterCredentials** | Pointer to **string** | Server registration password settings. If using \&quot;user\&quot; then the password remains unchanged at the end of the registration process. If using \&quot;random\&quot; then a random password will be generated and set on the server during registration. | [optional] [default to "user"]
**MinimumNumberOfConnectedInterfaces** | Pointer to **float32** | Minimum number of switch-connected interfaces required | [optional] [default to 0]
**AlwaysDiscoverInterfacesWithBDK** | Pointer to **bool** | Whether to always attempt to discover interfaces with BDK | [optional] [default to true]
**EnableTpm** | Pointer to **bool** | Whether to enable TPM | [optional] [default to true]
**EnableIntelTxt** | Pointer to **bool** | Whether to enable Intel TXT | [optional] [default to true]
**EnableSyslogMonitoring** | Pointer to **bool** | Whether to enable syslog monitoring | [optional] [default to true]
**DisableTpmAfterRegistration** | Pointer to **bool** | Whether to disable TPM after registration | [optional] [default to false]
**BiosProfile** | Pointer to [**[]ServerRegistrationBiosProfile**](ServerRegistrationBiosProfile.md) | Server registration BIOS profile | [optional] 
**DefaultVirtualMediaProtocol** | Pointer to **string** | Default protocol for virtual media | [optional] [default to "HTTPS"]
**ResetRaidControllers** | Pointer to **bool** | Whether to reset RAID controllers to factory defaults | [optional] [default to true]
**CleanupDrives** | Pointer to **bool** | Whether to cleanup drives | [optional] [default to true]
**RecreateRaid** | Pointer to **bool** | Whether to recreate RAID | [optional] [default to true]
**DisableEmbeddedNics** | Pointer to **bool** | Whether to disable embedded NICs | [optional] [default to true]
**RaidOneDrive** | Pointer to **string** | Raid one drive | [optional] [default to "RAID0"]
**RaidTwoDrives** | Pointer to **string** | Raid two drives | [optional] [default to "RAID1"]
**RaidEvenNumberMoreThanTwoDrives** | Pointer to **string** | Raid even number more than two drives | [optional] [default to "RAID10"]
**RaidOddNumberMoreThanOneDrive** | Pointer to **string** | Raid odd number more than one drive | [optional] [default to "RAID5"]

## Methods

### NewServerRegistrationProfileSettings

`func NewServerRegistrationProfileSettings() *ServerRegistrationProfileSettings`

NewServerRegistrationProfileSettings instantiates a new ServerRegistrationProfileSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerRegistrationProfileSettingsWithDefaults

`func NewServerRegistrationProfileSettingsWithDefaults() *ServerRegistrationProfileSettings`

NewServerRegistrationProfileSettingsWithDefaults instantiates a new ServerRegistrationProfileSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegisterCredentials

`func (o *ServerRegistrationProfileSettings) GetRegisterCredentials() string`

GetRegisterCredentials returns the RegisterCredentials field if non-nil, zero value otherwise.

### GetRegisterCredentialsOk

`func (o *ServerRegistrationProfileSettings) GetRegisterCredentialsOk() (*string, bool)`

GetRegisterCredentialsOk returns a tuple with the RegisterCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterCredentials

`func (o *ServerRegistrationProfileSettings) SetRegisterCredentials(v string)`

SetRegisterCredentials sets RegisterCredentials field to given value.

### HasRegisterCredentials

`func (o *ServerRegistrationProfileSettings) HasRegisterCredentials() bool`

HasRegisterCredentials returns a boolean if a field has been set.

### GetMinimumNumberOfConnectedInterfaces

`func (o *ServerRegistrationProfileSettings) GetMinimumNumberOfConnectedInterfaces() float32`

GetMinimumNumberOfConnectedInterfaces returns the MinimumNumberOfConnectedInterfaces field if non-nil, zero value otherwise.

### GetMinimumNumberOfConnectedInterfacesOk

`func (o *ServerRegistrationProfileSettings) GetMinimumNumberOfConnectedInterfacesOk() (*float32, bool)`

GetMinimumNumberOfConnectedInterfacesOk returns a tuple with the MinimumNumberOfConnectedInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumNumberOfConnectedInterfaces

`func (o *ServerRegistrationProfileSettings) SetMinimumNumberOfConnectedInterfaces(v float32)`

SetMinimumNumberOfConnectedInterfaces sets MinimumNumberOfConnectedInterfaces field to given value.

### HasMinimumNumberOfConnectedInterfaces

`func (o *ServerRegistrationProfileSettings) HasMinimumNumberOfConnectedInterfaces() bool`

HasMinimumNumberOfConnectedInterfaces returns a boolean if a field has been set.

### GetAlwaysDiscoverInterfacesWithBDK

`func (o *ServerRegistrationProfileSettings) GetAlwaysDiscoverInterfacesWithBDK() bool`

GetAlwaysDiscoverInterfacesWithBDK returns the AlwaysDiscoverInterfacesWithBDK field if non-nil, zero value otherwise.

### GetAlwaysDiscoverInterfacesWithBDKOk

`func (o *ServerRegistrationProfileSettings) GetAlwaysDiscoverInterfacesWithBDKOk() (*bool, bool)`

GetAlwaysDiscoverInterfacesWithBDKOk returns a tuple with the AlwaysDiscoverInterfacesWithBDK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysDiscoverInterfacesWithBDK

`func (o *ServerRegistrationProfileSettings) SetAlwaysDiscoverInterfacesWithBDK(v bool)`

SetAlwaysDiscoverInterfacesWithBDK sets AlwaysDiscoverInterfacesWithBDK field to given value.

### HasAlwaysDiscoverInterfacesWithBDK

`func (o *ServerRegistrationProfileSettings) HasAlwaysDiscoverInterfacesWithBDK() bool`

HasAlwaysDiscoverInterfacesWithBDK returns a boolean if a field has been set.

### GetEnableTpm

`func (o *ServerRegistrationProfileSettings) GetEnableTpm() bool`

GetEnableTpm returns the EnableTpm field if non-nil, zero value otherwise.

### GetEnableTpmOk

`func (o *ServerRegistrationProfileSettings) GetEnableTpmOk() (*bool, bool)`

GetEnableTpmOk returns a tuple with the EnableTpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTpm

`func (o *ServerRegistrationProfileSettings) SetEnableTpm(v bool)`

SetEnableTpm sets EnableTpm field to given value.

### HasEnableTpm

`func (o *ServerRegistrationProfileSettings) HasEnableTpm() bool`

HasEnableTpm returns a boolean if a field has been set.

### GetEnableIntelTxt

`func (o *ServerRegistrationProfileSettings) GetEnableIntelTxt() bool`

GetEnableIntelTxt returns the EnableIntelTxt field if non-nil, zero value otherwise.

### GetEnableIntelTxtOk

`func (o *ServerRegistrationProfileSettings) GetEnableIntelTxtOk() (*bool, bool)`

GetEnableIntelTxtOk returns a tuple with the EnableIntelTxt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIntelTxt

`func (o *ServerRegistrationProfileSettings) SetEnableIntelTxt(v bool)`

SetEnableIntelTxt sets EnableIntelTxt field to given value.

### HasEnableIntelTxt

`func (o *ServerRegistrationProfileSettings) HasEnableIntelTxt() bool`

HasEnableIntelTxt returns a boolean if a field has been set.

### GetEnableSyslogMonitoring

`func (o *ServerRegistrationProfileSettings) GetEnableSyslogMonitoring() bool`

GetEnableSyslogMonitoring returns the EnableSyslogMonitoring field if non-nil, zero value otherwise.

### GetEnableSyslogMonitoringOk

`func (o *ServerRegistrationProfileSettings) GetEnableSyslogMonitoringOk() (*bool, bool)`

GetEnableSyslogMonitoringOk returns a tuple with the EnableSyslogMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSyslogMonitoring

`func (o *ServerRegistrationProfileSettings) SetEnableSyslogMonitoring(v bool)`

SetEnableSyslogMonitoring sets EnableSyslogMonitoring field to given value.

### HasEnableSyslogMonitoring

`func (o *ServerRegistrationProfileSettings) HasEnableSyslogMonitoring() bool`

HasEnableSyslogMonitoring returns a boolean if a field has been set.

### GetDisableTpmAfterRegistration

`func (o *ServerRegistrationProfileSettings) GetDisableTpmAfterRegistration() bool`

GetDisableTpmAfterRegistration returns the DisableTpmAfterRegistration field if non-nil, zero value otherwise.

### GetDisableTpmAfterRegistrationOk

`func (o *ServerRegistrationProfileSettings) GetDisableTpmAfterRegistrationOk() (*bool, bool)`

GetDisableTpmAfterRegistrationOk returns a tuple with the DisableTpmAfterRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableTpmAfterRegistration

`func (o *ServerRegistrationProfileSettings) SetDisableTpmAfterRegistration(v bool)`

SetDisableTpmAfterRegistration sets DisableTpmAfterRegistration field to given value.

### HasDisableTpmAfterRegistration

`func (o *ServerRegistrationProfileSettings) HasDisableTpmAfterRegistration() bool`

HasDisableTpmAfterRegistration returns a boolean if a field has been set.

### GetBiosProfile

`func (o *ServerRegistrationProfileSettings) GetBiosProfile() []ServerRegistrationBiosProfile`

GetBiosProfile returns the BiosProfile field if non-nil, zero value otherwise.

### GetBiosProfileOk

`func (o *ServerRegistrationProfileSettings) GetBiosProfileOk() (*[]ServerRegistrationBiosProfile, bool)`

GetBiosProfileOk returns a tuple with the BiosProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosProfile

`func (o *ServerRegistrationProfileSettings) SetBiosProfile(v []ServerRegistrationBiosProfile)`

SetBiosProfile sets BiosProfile field to given value.

### HasBiosProfile

`func (o *ServerRegistrationProfileSettings) HasBiosProfile() bool`

HasBiosProfile returns a boolean if a field has been set.

### GetDefaultVirtualMediaProtocol

`func (o *ServerRegistrationProfileSettings) GetDefaultVirtualMediaProtocol() string`

GetDefaultVirtualMediaProtocol returns the DefaultVirtualMediaProtocol field if non-nil, zero value otherwise.

### GetDefaultVirtualMediaProtocolOk

`func (o *ServerRegistrationProfileSettings) GetDefaultVirtualMediaProtocolOk() (*string, bool)`

GetDefaultVirtualMediaProtocolOk returns a tuple with the DefaultVirtualMediaProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVirtualMediaProtocol

`func (o *ServerRegistrationProfileSettings) SetDefaultVirtualMediaProtocol(v string)`

SetDefaultVirtualMediaProtocol sets DefaultVirtualMediaProtocol field to given value.

### HasDefaultVirtualMediaProtocol

`func (o *ServerRegistrationProfileSettings) HasDefaultVirtualMediaProtocol() bool`

HasDefaultVirtualMediaProtocol returns a boolean if a field has been set.

### GetResetRaidControllers

`func (o *ServerRegistrationProfileSettings) GetResetRaidControllers() bool`

GetResetRaidControllers returns the ResetRaidControllers field if non-nil, zero value otherwise.

### GetResetRaidControllersOk

`func (o *ServerRegistrationProfileSettings) GetResetRaidControllersOk() (*bool, bool)`

GetResetRaidControllersOk returns a tuple with the ResetRaidControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetRaidControllers

`func (o *ServerRegistrationProfileSettings) SetResetRaidControllers(v bool)`

SetResetRaidControllers sets ResetRaidControllers field to given value.

### HasResetRaidControllers

`func (o *ServerRegistrationProfileSettings) HasResetRaidControllers() bool`

HasResetRaidControllers returns a boolean if a field has been set.

### GetCleanupDrives

`func (o *ServerRegistrationProfileSettings) GetCleanupDrives() bool`

GetCleanupDrives returns the CleanupDrives field if non-nil, zero value otherwise.

### GetCleanupDrivesOk

`func (o *ServerRegistrationProfileSettings) GetCleanupDrivesOk() (*bool, bool)`

GetCleanupDrivesOk returns a tuple with the CleanupDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupDrives

`func (o *ServerRegistrationProfileSettings) SetCleanupDrives(v bool)`

SetCleanupDrives sets CleanupDrives field to given value.

### HasCleanupDrives

`func (o *ServerRegistrationProfileSettings) HasCleanupDrives() bool`

HasCleanupDrives returns a boolean if a field has been set.

### GetRecreateRaid

`func (o *ServerRegistrationProfileSettings) GetRecreateRaid() bool`

GetRecreateRaid returns the RecreateRaid field if non-nil, zero value otherwise.

### GetRecreateRaidOk

`func (o *ServerRegistrationProfileSettings) GetRecreateRaidOk() (*bool, bool)`

GetRecreateRaidOk returns a tuple with the RecreateRaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecreateRaid

`func (o *ServerRegistrationProfileSettings) SetRecreateRaid(v bool)`

SetRecreateRaid sets RecreateRaid field to given value.

### HasRecreateRaid

`func (o *ServerRegistrationProfileSettings) HasRecreateRaid() bool`

HasRecreateRaid returns a boolean if a field has been set.

### GetDisableEmbeddedNics

`func (o *ServerRegistrationProfileSettings) GetDisableEmbeddedNics() bool`

GetDisableEmbeddedNics returns the DisableEmbeddedNics field if non-nil, zero value otherwise.

### GetDisableEmbeddedNicsOk

`func (o *ServerRegistrationProfileSettings) GetDisableEmbeddedNicsOk() (*bool, bool)`

GetDisableEmbeddedNicsOk returns a tuple with the DisableEmbeddedNics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEmbeddedNics

`func (o *ServerRegistrationProfileSettings) SetDisableEmbeddedNics(v bool)`

SetDisableEmbeddedNics sets DisableEmbeddedNics field to given value.

### HasDisableEmbeddedNics

`func (o *ServerRegistrationProfileSettings) HasDisableEmbeddedNics() bool`

HasDisableEmbeddedNics returns a boolean if a field has been set.

### GetRaidOneDrive

`func (o *ServerRegistrationProfileSettings) GetRaidOneDrive() string`

GetRaidOneDrive returns the RaidOneDrive field if non-nil, zero value otherwise.

### GetRaidOneDriveOk

`func (o *ServerRegistrationProfileSettings) GetRaidOneDriveOk() (*string, bool)`

GetRaidOneDriveOk returns a tuple with the RaidOneDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidOneDrive

`func (o *ServerRegistrationProfileSettings) SetRaidOneDrive(v string)`

SetRaidOneDrive sets RaidOneDrive field to given value.

### HasRaidOneDrive

`func (o *ServerRegistrationProfileSettings) HasRaidOneDrive() bool`

HasRaidOneDrive returns a boolean if a field has been set.

### GetRaidTwoDrives

`func (o *ServerRegistrationProfileSettings) GetRaidTwoDrives() string`

GetRaidTwoDrives returns the RaidTwoDrives field if non-nil, zero value otherwise.

### GetRaidTwoDrivesOk

`func (o *ServerRegistrationProfileSettings) GetRaidTwoDrivesOk() (*string, bool)`

GetRaidTwoDrivesOk returns a tuple with the RaidTwoDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidTwoDrives

`func (o *ServerRegistrationProfileSettings) SetRaidTwoDrives(v string)`

SetRaidTwoDrives sets RaidTwoDrives field to given value.

### HasRaidTwoDrives

`func (o *ServerRegistrationProfileSettings) HasRaidTwoDrives() bool`

HasRaidTwoDrives returns a boolean if a field has been set.

### GetRaidEvenNumberMoreThanTwoDrives

`func (o *ServerRegistrationProfileSettings) GetRaidEvenNumberMoreThanTwoDrives() string`

GetRaidEvenNumberMoreThanTwoDrives returns the RaidEvenNumberMoreThanTwoDrives field if non-nil, zero value otherwise.

### GetRaidEvenNumberMoreThanTwoDrivesOk

`func (o *ServerRegistrationProfileSettings) GetRaidEvenNumberMoreThanTwoDrivesOk() (*string, bool)`

GetRaidEvenNumberMoreThanTwoDrivesOk returns a tuple with the RaidEvenNumberMoreThanTwoDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidEvenNumberMoreThanTwoDrives

`func (o *ServerRegistrationProfileSettings) SetRaidEvenNumberMoreThanTwoDrives(v string)`

SetRaidEvenNumberMoreThanTwoDrives sets RaidEvenNumberMoreThanTwoDrives field to given value.

### HasRaidEvenNumberMoreThanTwoDrives

`func (o *ServerRegistrationProfileSettings) HasRaidEvenNumberMoreThanTwoDrives() bool`

HasRaidEvenNumberMoreThanTwoDrives returns a boolean if a field has been set.

### GetRaidOddNumberMoreThanOneDrive

`func (o *ServerRegistrationProfileSettings) GetRaidOddNumberMoreThanOneDrive() string`

GetRaidOddNumberMoreThanOneDrive returns the RaidOddNumberMoreThanOneDrive field if non-nil, zero value otherwise.

### GetRaidOddNumberMoreThanOneDriveOk

`func (o *ServerRegistrationProfileSettings) GetRaidOddNumberMoreThanOneDriveOk() (*string, bool)`

GetRaidOddNumberMoreThanOneDriveOk returns a tuple with the RaidOddNumberMoreThanOneDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidOddNumberMoreThanOneDrive

`func (o *ServerRegistrationProfileSettings) SetRaidOddNumberMoreThanOneDrive(v string)`

SetRaidOddNumberMoreThanOneDrive sets RaidOddNumberMoreThanOneDrive field to given value.

### HasRaidOddNumberMoreThanOneDrive

`func (o *ServerRegistrationProfileSettings) HasRaidOddNumberMoreThanOneDrive() bool`

HasRaidOddNumberMoreThanOneDrive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


