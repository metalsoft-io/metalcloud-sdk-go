# ServerRegistrationProfileUpdateSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegisterCredentials** | Pointer to **string** | Server registration password settings. If using \&quot;user\&quot; then the password remains unchanged at the end of the registration process. If using \&quot;random\&quot; then a random password will be generated and set on the server during registration. | [optional] 
**MinimumNumberOfConnectedInterfaces** | Pointer to **float32** | Minimum number of switch-connected interfaces required | [optional] 
**AlwaysDiscoverInterfacesWithBDK** | Pointer to **bool** | Whether to always attempt to discover interfaces with BDK | [optional] 
**ClearTpm** | Pointer to **bool** | Whether to clear TPM | [optional] 
**EnableTpm** | Pointer to **bool** | Whether to enable TPM | [optional] 
**EnableIntelTxt** | Pointer to **bool** | Whether to enable Intel TXT | [optional] 
**EnableSyslogMonitoring** | Pointer to **bool** | Whether to enable syslog monitoring | [optional] 
**DisableTpmAfterRegistration** | Pointer to **bool** | Whether to disable TPM after registration | [optional] 
**DefaultVirtualMediaProtocol** | Pointer to **string** | Default protocol for virtual media | [optional] 
**FirmwareBaselineId** | Pointer to **int64** | Firmware baseline ID to apply during registration | [optional] 
**ResetRaidControllers** | Pointer to **bool** | Whether to reset RAID controllers to factory defaults | [optional] 
**CleanupDrives** | Pointer to **bool** | Whether to cleanup drives | [optional] 
**RecreateRaid** | Pointer to **bool** | Whether to recreate RAID | [optional] 
**DisableEmbeddedNics** | Pointer to **bool** | Whether to disable embedded NICs | [optional] 
**RaidOneDrive** | Pointer to **string** | Raid one drive | [optional] 
**RaidTwoDrives** | Pointer to **string** | Raid two drives | [optional] 
**RaidEvenNumberMoreThanTwoDrives** | Pointer to **string** | Raid even number more than two drives | [optional] 
**RaidOddNumberMoreThanOneDrive** | Pointer to **string** | Raid odd number more than one drive | [optional] 
**BiosProfile** | Pointer to [**[]ServerRegistrationBiosProfile**](ServerRegistrationBiosProfile.md) | Server registration BIOS profile | [optional] 
**DpuMode** | Pointer to **string** | Whether to register the server in DPU or NIC mode. | [optional] [default to "dpu"]
**DpuMaxVirtualFunctions** | Pointer to **float32** | Maximum number of virtual functions to create on DPUs. This setting is only applicable if dpuMode is set to DPU. | [optional] [default to 64]

## Methods

### NewServerRegistrationProfileUpdateSettings

`func NewServerRegistrationProfileUpdateSettings() *ServerRegistrationProfileUpdateSettings`

NewServerRegistrationProfileUpdateSettings instantiates a new ServerRegistrationProfileUpdateSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerRegistrationProfileUpdateSettingsWithDefaults

`func NewServerRegistrationProfileUpdateSettingsWithDefaults() *ServerRegistrationProfileUpdateSettings`

NewServerRegistrationProfileUpdateSettingsWithDefaults instantiates a new ServerRegistrationProfileUpdateSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegisterCredentials

`func (o *ServerRegistrationProfileUpdateSettings) GetRegisterCredentials() string`

GetRegisterCredentials returns the RegisterCredentials field if non-nil, zero value otherwise.

### GetRegisterCredentialsOk

`func (o *ServerRegistrationProfileUpdateSettings) GetRegisterCredentialsOk() (*string, bool)`

GetRegisterCredentialsOk returns a tuple with the RegisterCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterCredentials

`func (o *ServerRegistrationProfileUpdateSettings) SetRegisterCredentials(v string)`

SetRegisterCredentials sets RegisterCredentials field to given value.

### HasRegisterCredentials

`func (o *ServerRegistrationProfileUpdateSettings) HasRegisterCredentials() bool`

HasRegisterCredentials returns a boolean if a field has been set.

### GetMinimumNumberOfConnectedInterfaces

`func (o *ServerRegistrationProfileUpdateSettings) GetMinimumNumberOfConnectedInterfaces() float32`

GetMinimumNumberOfConnectedInterfaces returns the MinimumNumberOfConnectedInterfaces field if non-nil, zero value otherwise.

### GetMinimumNumberOfConnectedInterfacesOk

`func (o *ServerRegistrationProfileUpdateSettings) GetMinimumNumberOfConnectedInterfacesOk() (*float32, bool)`

GetMinimumNumberOfConnectedInterfacesOk returns a tuple with the MinimumNumberOfConnectedInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumNumberOfConnectedInterfaces

`func (o *ServerRegistrationProfileUpdateSettings) SetMinimumNumberOfConnectedInterfaces(v float32)`

SetMinimumNumberOfConnectedInterfaces sets MinimumNumberOfConnectedInterfaces field to given value.

### HasMinimumNumberOfConnectedInterfaces

`func (o *ServerRegistrationProfileUpdateSettings) HasMinimumNumberOfConnectedInterfaces() bool`

HasMinimumNumberOfConnectedInterfaces returns a boolean if a field has been set.

### GetAlwaysDiscoverInterfacesWithBDK

`func (o *ServerRegistrationProfileUpdateSettings) GetAlwaysDiscoverInterfacesWithBDK() bool`

GetAlwaysDiscoverInterfacesWithBDK returns the AlwaysDiscoverInterfacesWithBDK field if non-nil, zero value otherwise.

### GetAlwaysDiscoverInterfacesWithBDKOk

`func (o *ServerRegistrationProfileUpdateSettings) GetAlwaysDiscoverInterfacesWithBDKOk() (*bool, bool)`

GetAlwaysDiscoverInterfacesWithBDKOk returns a tuple with the AlwaysDiscoverInterfacesWithBDK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysDiscoverInterfacesWithBDK

`func (o *ServerRegistrationProfileUpdateSettings) SetAlwaysDiscoverInterfacesWithBDK(v bool)`

SetAlwaysDiscoverInterfacesWithBDK sets AlwaysDiscoverInterfacesWithBDK field to given value.

### HasAlwaysDiscoverInterfacesWithBDK

`func (o *ServerRegistrationProfileUpdateSettings) HasAlwaysDiscoverInterfacesWithBDK() bool`

HasAlwaysDiscoverInterfacesWithBDK returns a boolean if a field has been set.

### GetClearTpm

`func (o *ServerRegistrationProfileUpdateSettings) GetClearTpm() bool`

GetClearTpm returns the ClearTpm field if non-nil, zero value otherwise.

### GetClearTpmOk

`func (o *ServerRegistrationProfileUpdateSettings) GetClearTpmOk() (*bool, bool)`

GetClearTpmOk returns a tuple with the ClearTpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearTpm

`func (o *ServerRegistrationProfileUpdateSettings) SetClearTpm(v bool)`

SetClearTpm sets ClearTpm field to given value.

### HasClearTpm

`func (o *ServerRegistrationProfileUpdateSettings) HasClearTpm() bool`

HasClearTpm returns a boolean if a field has been set.

### GetEnableTpm

`func (o *ServerRegistrationProfileUpdateSettings) GetEnableTpm() bool`

GetEnableTpm returns the EnableTpm field if non-nil, zero value otherwise.

### GetEnableTpmOk

`func (o *ServerRegistrationProfileUpdateSettings) GetEnableTpmOk() (*bool, bool)`

GetEnableTpmOk returns a tuple with the EnableTpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTpm

`func (o *ServerRegistrationProfileUpdateSettings) SetEnableTpm(v bool)`

SetEnableTpm sets EnableTpm field to given value.

### HasEnableTpm

`func (o *ServerRegistrationProfileUpdateSettings) HasEnableTpm() bool`

HasEnableTpm returns a boolean if a field has been set.

### GetEnableIntelTxt

`func (o *ServerRegistrationProfileUpdateSettings) GetEnableIntelTxt() bool`

GetEnableIntelTxt returns the EnableIntelTxt field if non-nil, zero value otherwise.

### GetEnableIntelTxtOk

`func (o *ServerRegistrationProfileUpdateSettings) GetEnableIntelTxtOk() (*bool, bool)`

GetEnableIntelTxtOk returns a tuple with the EnableIntelTxt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIntelTxt

`func (o *ServerRegistrationProfileUpdateSettings) SetEnableIntelTxt(v bool)`

SetEnableIntelTxt sets EnableIntelTxt field to given value.

### HasEnableIntelTxt

`func (o *ServerRegistrationProfileUpdateSettings) HasEnableIntelTxt() bool`

HasEnableIntelTxt returns a boolean if a field has been set.

### GetEnableSyslogMonitoring

`func (o *ServerRegistrationProfileUpdateSettings) GetEnableSyslogMonitoring() bool`

GetEnableSyslogMonitoring returns the EnableSyslogMonitoring field if non-nil, zero value otherwise.

### GetEnableSyslogMonitoringOk

`func (o *ServerRegistrationProfileUpdateSettings) GetEnableSyslogMonitoringOk() (*bool, bool)`

GetEnableSyslogMonitoringOk returns a tuple with the EnableSyslogMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSyslogMonitoring

`func (o *ServerRegistrationProfileUpdateSettings) SetEnableSyslogMonitoring(v bool)`

SetEnableSyslogMonitoring sets EnableSyslogMonitoring field to given value.

### HasEnableSyslogMonitoring

`func (o *ServerRegistrationProfileUpdateSettings) HasEnableSyslogMonitoring() bool`

HasEnableSyslogMonitoring returns a boolean if a field has been set.

### GetDisableTpmAfterRegistration

`func (o *ServerRegistrationProfileUpdateSettings) GetDisableTpmAfterRegistration() bool`

GetDisableTpmAfterRegistration returns the DisableTpmAfterRegistration field if non-nil, zero value otherwise.

### GetDisableTpmAfterRegistrationOk

`func (o *ServerRegistrationProfileUpdateSettings) GetDisableTpmAfterRegistrationOk() (*bool, bool)`

GetDisableTpmAfterRegistrationOk returns a tuple with the DisableTpmAfterRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableTpmAfterRegistration

`func (o *ServerRegistrationProfileUpdateSettings) SetDisableTpmAfterRegistration(v bool)`

SetDisableTpmAfterRegistration sets DisableTpmAfterRegistration field to given value.

### HasDisableTpmAfterRegistration

`func (o *ServerRegistrationProfileUpdateSettings) HasDisableTpmAfterRegistration() bool`

HasDisableTpmAfterRegistration returns a boolean if a field has been set.

### GetDefaultVirtualMediaProtocol

`func (o *ServerRegistrationProfileUpdateSettings) GetDefaultVirtualMediaProtocol() string`

GetDefaultVirtualMediaProtocol returns the DefaultVirtualMediaProtocol field if non-nil, zero value otherwise.

### GetDefaultVirtualMediaProtocolOk

`func (o *ServerRegistrationProfileUpdateSettings) GetDefaultVirtualMediaProtocolOk() (*string, bool)`

GetDefaultVirtualMediaProtocolOk returns a tuple with the DefaultVirtualMediaProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVirtualMediaProtocol

`func (o *ServerRegistrationProfileUpdateSettings) SetDefaultVirtualMediaProtocol(v string)`

SetDefaultVirtualMediaProtocol sets DefaultVirtualMediaProtocol field to given value.

### HasDefaultVirtualMediaProtocol

`func (o *ServerRegistrationProfileUpdateSettings) HasDefaultVirtualMediaProtocol() bool`

HasDefaultVirtualMediaProtocol returns a boolean if a field has been set.

### GetFirmwareBaselineId

`func (o *ServerRegistrationProfileUpdateSettings) GetFirmwareBaselineId() int64`

GetFirmwareBaselineId returns the FirmwareBaselineId field if non-nil, zero value otherwise.

### GetFirmwareBaselineIdOk

`func (o *ServerRegistrationProfileUpdateSettings) GetFirmwareBaselineIdOk() (*int64, bool)`

GetFirmwareBaselineIdOk returns a tuple with the FirmwareBaselineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareBaselineId

`func (o *ServerRegistrationProfileUpdateSettings) SetFirmwareBaselineId(v int64)`

SetFirmwareBaselineId sets FirmwareBaselineId field to given value.

### HasFirmwareBaselineId

`func (o *ServerRegistrationProfileUpdateSettings) HasFirmwareBaselineId() bool`

HasFirmwareBaselineId returns a boolean if a field has been set.

### GetResetRaidControllers

`func (o *ServerRegistrationProfileUpdateSettings) GetResetRaidControllers() bool`

GetResetRaidControllers returns the ResetRaidControllers field if non-nil, zero value otherwise.

### GetResetRaidControllersOk

`func (o *ServerRegistrationProfileUpdateSettings) GetResetRaidControllersOk() (*bool, bool)`

GetResetRaidControllersOk returns a tuple with the ResetRaidControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetRaidControllers

`func (o *ServerRegistrationProfileUpdateSettings) SetResetRaidControllers(v bool)`

SetResetRaidControllers sets ResetRaidControllers field to given value.

### HasResetRaidControllers

`func (o *ServerRegistrationProfileUpdateSettings) HasResetRaidControllers() bool`

HasResetRaidControllers returns a boolean if a field has been set.

### GetCleanupDrives

`func (o *ServerRegistrationProfileUpdateSettings) GetCleanupDrives() bool`

GetCleanupDrives returns the CleanupDrives field if non-nil, zero value otherwise.

### GetCleanupDrivesOk

`func (o *ServerRegistrationProfileUpdateSettings) GetCleanupDrivesOk() (*bool, bool)`

GetCleanupDrivesOk returns a tuple with the CleanupDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupDrives

`func (o *ServerRegistrationProfileUpdateSettings) SetCleanupDrives(v bool)`

SetCleanupDrives sets CleanupDrives field to given value.

### HasCleanupDrives

`func (o *ServerRegistrationProfileUpdateSettings) HasCleanupDrives() bool`

HasCleanupDrives returns a boolean if a field has been set.

### GetRecreateRaid

`func (o *ServerRegistrationProfileUpdateSettings) GetRecreateRaid() bool`

GetRecreateRaid returns the RecreateRaid field if non-nil, zero value otherwise.

### GetRecreateRaidOk

`func (o *ServerRegistrationProfileUpdateSettings) GetRecreateRaidOk() (*bool, bool)`

GetRecreateRaidOk returns a tuple with the RecreateRaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecreateRaid

`func (o *ServerRegistrationProfileUpdateSettings) SetRecreateRaid(v bool)`

SetRecreateRaid sets RecreateRaid field to given value.

### HasRecreateRaid

`func (o *ServerRegistrationProfileUpdateSettings) HasRecreateRaid() bool`

HasRecreateRaid returns a boolean if a field has been set.

### GetDisableEmbeddedNics

`func (o *ServerRegistrationProfileUpdateSettings) GetDisableEmbeddedNics() bool`

GetDisableEmbeddedNics returns the DisableEmbeddedNics field if non-nil, zero value otherwise.

### GetDisableEmbeddedNicsOk

`func (o *ServerRegistrationProfileUpdateSettings) GetDisableEmbeddedNicsOk() (*bool, bool)`

GetDisableEmbeddedNicsOk returns a tuple with the DisableEmbeddedNics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEmbeddedNics

`func (o *ServerRegistrationProfileUpdateSettings) SetDisableEmbeddedNics(v bool)`

SetDisableEmbeddedNics sets DisableEmbeddedNics field to given value.

### HasDisableEmbeddedNics

`func (o *ServerRegistrationProfileUpdateSettings) HasDisableEmbeddedNics() bool`

HasDisableEmbeddedNics returns a boolean if a field has been set.

### GetRaidOneDrive

`func (o *ServerRegistrationProfileUpdateSettings) GetRaidOneDrive() string`

GetRaidOneDrive returns the RaidOneDrive field if non-nil, zero value otherwise.

### GetRaidOneDriveOk

`func (o *ServerRegistrationProfileUpdateSettings) GetRaidOneDriveOk() (*string, bool)`

GetRaidOneDriveOk returns a tuple with the RaidOneDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidOneDrive

`func (o *ServerRegistrationProfileUpdateSettings) SetRaidOneDrive(v string)`

SetRaidOneDrive sets RaidOneDrive field to given value.

### HasRaidOneDrive

`func (o *ServerRegistrationProfileUpdateSettings) HasRaidOneDrive() bool`

HasRaidOneDrive returns a boolean if a field has been set.

### GetRaidTwoDrives

`func (o *ServerRegistrationProfileUpdateSettings) GetRaidTwoDrives() string`

GetRaidTwoDrives returns the RaidTwoDrives field if non-nil, zero value otherwise.

### GetRaidTwoDrivesOk

`func (o *ServerRegistrationProfileUpdateSettings) GetRaidTwoDrivesOk() (*string, bool)`

GetRaidTwoDrivesOk returns a tuple with the RaidTwoDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidTwoDrives

`func (o *ServerRegistrationProfileUpdateSettings) SetRaidTwoDrives(v string)`

SetRaidTwoDrives sets RaidTwoDrives field to given value.

### HasRaidTwoDrives

`func (o *ServerRegistrationProfileUpdateSettings) HasRaidTwoDrives() bool`

HasRaidTwoDrives returns a boolean if a field has been set.

### GetRaidEvenNumberMoreThanTwoDrives

`func (o *ServerRegistrationProfileUpdateSettings) GetRaidEvenNumberMoreThanTwoDrives() string`

GetRaidEvenNumberMoreThanTwoDrives returns the RaidEvenNumberMoreThanTwoDrives field if non-nil, zero value otherwise.

### GetRaidEvenNumberMoreThanTwoDrivesOk

`func (o *ServerRegistrationProfileUpdateSettings) GetRaidEvenNumberMoreThanTwoDrivesOk() (*string, bool)`

GetRaidEvenNumberMoreThanTwoDrivesOk returns a tuple with the RaidEvenNumberMoreThanTwoDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidEvenNumberMoreThanTwoDrives

`func (o *ServerRegistrationProfileUpdateSettings) SetRaidEvenNumberMoreThanTwoDrives(v string)`

SetRaidEvenNumberMoreThanTwoDrives sets RaidEvenNumberMoreThanTwoDrives field to given value.

### HasRaidEvenNumberMoreThanTwoDrives

`func (o *ServerRegistrationProfileUpdateSettings) HasRaidEvenNumberMoreThanTwoDrives() bool`

HasRaidEvenNumberMoreThanTwoDrives returns a boolean if a field has been set.

### GetRaidOddNumberMoreThanOneDrive

`func (o *ServerRegistrationProfileUpdateSettings) GetRaidOddNumberMoreThanOneDrive() string`

GetRaidOddNumberMoreThanOneDrive returns the RaidOddNumberMoreThanOneDrive field if non-nil, zero value otherwise.

### GetRaidOddNumberMoreThanOneDriveOk

`func (o *ServerRegistrationProfileUpdateSettings) GetRaidOddNumberMoreThanOneDriveOk() (*string, bool)`

GetRaidOddNumberMoreThanOneDriveOk returns a tuple with the RaidOddNumberMoreThanOneDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidOddNumberMoreThanOneDrive

`func (o *ServerRegistrationProfileUpdateSettings) SetRaidOddNumberMoreThanOneDrive(v string)`

SetRaidOddNumberMoreThanOneDrive sets RaidOddNumberMoreThanOneDrive field to given value.

### HasRaidOddNumberMoreThanOneDrive

`func (o *ServerRegistrationProfileUpdateSettings) HasRaidOddNumberMoreThanOneDrive() bool`

HasRaidOddNumberMoreThanOneDrive returns a boolean if a field has been set.

### GetBiosProfile

`func (o *ServerRegistrationProfileUpdateSettings) GetBiosProfile() []ServerRegistrationBiosProfile`

GetBiosProfile returns the BiosProfile field if non-nil, zero value otherwise.

### GetBiosProfileOk

`func (o *ServerRegistrationProfileUpdateSettings) GetBiosProfileOk() (*[]ServerRegistrationBiosProfile, bool)`

GetBiosProfileOk returns a tuple with the BiosProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosProfile

`func (o *ServerRegistrationProfileUpdateSettings) SetBiosProfile(v []ServerRegistrationBiosProfile)`

SetBiosProfile sets BiosProfile field to given value.

### HasBiosProfile

`func (o *ServerRegistrationProfileUpdateSettings) HasBiosProfile() bool`

HasBiosProfile returns a boolean if a field has been set.

### GetDpuMode

`func (o *ServerRegistrationProfileUpdateSettings) GetDpuMode() string`

GetDpuMode returns the DpuMode field if non-nil, zero value otherwise.

### GetDpuModeOk

`func (o *ServerRegistrationProfileUpdateSettings) GetDpuModeOk() (*string, bool)`

GetDpuModeOk returns a tuple with the DpuMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpuMode

`func (o *ServerRegistrationProfileUpdateSettings) SetDpuMode(v string)`

SetDpuMode sets DpuMode field to given value.

### HasDpuMode

`func (o *ServerRegistrationProfileUpdateSettings) HasDpuMode() bool`

HasDpuMode returns a boolean if a field has been set.

### GetDpuMaxVirtualFunctions

`func (o *ServerRegistrationProfileUpdateSettings) GetDpuMaxVirtualFunctions() float32`

GetDpuMaxVirtualFunctions returns the DpuMaxVirtualFunctions field if non-nil, zero value otherwise.

### GetDpuMaxVirtualFunctionsOk

`func (o *ServerRegistrationProfileUpdateSettings) GetDpuMaxVirtualFunctionsOk() (*float32, bool)`

GetDpuMaxVirtualFunctionsOk returns a tuple with the DpuMaxVirtualFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpuMaxVirtualFunctions

`func (o *ServerRegistrationProfileUpdateSettings) SetDpuMaxVirtualFunctions(v float32)`

SetDpuMaxVirtualFunctions sets DpuMaxVirtualFunctions field to given value.

### HasDpuMaxVirtualFunctions

`func (o *ServerRegistrationProfileUpdateSettings) HasDpuMaxVirtualFunctions() bool`

HasDpuMaxVirtualFunctions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


