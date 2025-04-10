# ServerPolicyUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegisterCredentials** | Pointer to **string** | Server registration type | [optional] 
**MinimumNumberOfConnectedInterfaces** | Pointer to **float32** | Minimum number of switch-connected interfaces required | [optional] 
**DhcpOption82ToIPMapping** | Pointer to **map[string]interface{}** | Option82 to IP mapping | [optional] 
**DhcpBmcMacAddressWhitelistEnabled** | Pointer to **bool** | Whether to enable DHCP BMC MAC address whitelist | [optional] 
**DhcpBmcMacAddressWhitelist** | Pointer to **[]string** | List of DHCP BMC MAC address whitelist | [optional] 
**RaidConfigurationEnabled** | Pointer to **bool** | Whether RAID configuration is enabled | [optional] 
**DisableTpmAfterRegistration** | Pointer to **bool** | Whether to disable TPM after registration | [optional] 
**SyslogMonitoringEnabled** | Pointer to **bool** | Whether syslog monitoring is enabled | [optional] 
**DefaultServerCleanupPolicyID** | Pointer to **float32** | Default server cleanup policy ID | [optional] 
**AutomaticallyAllocateServerTypes** | Pointer to **bool** | Automatically allocate server types | [optional] 
**AutomaticallySetServersAsAvailable** | Pointer to **bool** | Automatically set servers as available | [optional] 
**ServerRegistrationBiosProfile** | Pointer to **[]string** | Server registration BIOS profile | [optional] 

## Methods

### NewServerPolicyUpdateDto

`func NewServerPolicyUpdateDto() *ServerPolicyUpdateDto`

NewServerPolicyUpdateDto instantiates a new ServerPolicyUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerPolicyUpdateDtoWithDefaults

`func NewServerPolicyUpdateDtoWithDefaults() *ServerPolicyUpdateDto`

NewServerPolicyUpdateDtoWithDefaults instantiates a new ServerPolicyUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegisterCredentials

`func (o *ServerPolicyUpdateDto) GetRegisterCredentials() string`

GetRegisterCredentials returns the RegisterCredentials field if non-nil, zero value otherwise.

### GetRegisterCredentialsOk

`func (o *ServerPolicyUpdateDto) GetRegisterCredentialsOk() (*string, bool)`

GetRegisterCredentialsOk returns a tuple with the RegisterCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterCredentials

`func (o *ServerPolicyUpdateDto) SetRegisterCredentials(v string)`

SetRegisterCredentials sets RegisterCredentials field to given value.

### HasRegisterCredentials

`func (o *ServerPolicyUpdateDto) HasRegisterCredentials() bool`

HasRegisterCredentials returns a boolean if a field has been set.

### GetMinimumNumberOfConnectedInterfaces

`func (o *ServerPolicyUpdateDto) GetMinimumNumberOfConnectedInterfaces() float32`

GetMinimumNumberOfConnectedInterfaces returns the MinimumNumberOfConnectedInterfaces field if non-nil, zero value otherwise.

### GetMinimumNumberOfConnectedInterfacesOk

`func (o *ServerPolicyUpdateDto) GetMinimumNumberOfConnectedInterfacesOk() (*float32, bool)`

GetMinimumNumberOfConnectedInterfacesOk returns a tuple with the MinimumNumberOfConnectedInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumNumberOfConnectedInterfaces

`func (o *ServerPolicyUpdateDto) SetMinimumNumberOfConnectedInterfaces(v float32)`

SetMinimumNumberOfConnectedInterfaces sets MinimumNumberOfConnectedInterfaces field to given value.

### HasMinimumNumberOfConnectedInterfaces

`func (o *ServerPolicyUpdateDto) HasMinimumNumberOfConnectedInterfaces() bool`

HasMinimumNumberOfConnectedInterfaces returns a boolean if a field has been set.

### GetDhcpOption82ToIPMapping

`func (o *ServerPolicyUpdateDto) GetDhcpOption82ToIPMapping() map[string]interface{}`

GetDhcpOption82ToIPMapping returns the DhcpOption82ToIPMapping field if non-nil, zero value otherwise.

### GetDhcpOption82ToIPMappingOk

`func (o *ServerPolicyUpdateDto) GetDhcpOption82ToIPMappingOk() (*map[string]interface{}, bool)`

GetDhcpOption82ToIPMappingOk returns a tuple with the DhcpOption82ToIPMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOption82ToIPMapping

`func (o *ServerPolicyUpdateDto) SetDhcpOption82ToIPMapping(v map[string]interface{})`

SetDhcpOption82ToIPMapping sets DhcpOption82ToIPMapping field to given value.

### HasDhcpOption82ToIPMapping

`func (o *ServerPolicyUpdateDto) HasDhcpOption82ToIPMapping() bool`

HasDhcpOption82ToIPMapping returns a boolean if a field has been set.

### GetDhcpBmcMacAddressWhitelistEnabled

`func (o *ServerPolicyUpdateDto) GetDhcpBmcMacAddressWhitelistEnabled() bool`

GetDhcpBmcMacAddressWhitelistEnabled returns the DhcpBmcMacAddressWhitelistEnabled field if non-nil, zero value otherwise.

### GetDhcpBmcMacAddressWhitelistEnabledOk

`func (o *ServerPolicyUpdateDto) GetDhcpBmcMacAddressWhitelistEnabledOk() (*bool, bool)`

GetDhcpBmcMacAddressWhitelistEnabledOk returns a tuple with the DhcpBmcMacAddressWhitelistEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpBmcMacAddressWhitelistEnabled

`func (o *ServerPolicyUpdateDto) SetDhcpBmcMacAddressWhitelistEnabled(v bool)`

SetDhcpBmcMacAddressWhitelistEnabled sets DhcpBmcMacAddressWhitelistEnabled field to given value.

### HasDhcpBmcMacAddressWhitelistEnabled

`func (o *ServerPolicyUpdateDto) HasDhcpBmcMacAddressWhitelistEnabled() bool`

HasDhcpBmcMacAddressWhitelistEnabled returns a boolean if a field has been set.

### GetDhcpBmcMacAddressWhitelist

`func (o *ServerPolicyUpdateDto) GetDhcpBmcMacAddressWhitelist() []string`

GetDhcpBmcMacAddressWhitelist returns the DhcpBmcMacAddressWhitelist field if non-nil, zero value otherwise.

### GetDhcpBmcMacAddressWhitelistOk

`func (o *ServerPolicyUpdateDto) GetDhcpBmcMacAddressWhitelistOk() (*[]string, bool)`

GetDhcpBmcMacAddressWhitelistOk returns a tuple with the DhcpBmcMacAddressWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpBmcMacAddressWhitelist

`func (o *ServerPolicyUpdateDto) SetDhcpBmcMacAddressWhitelist(v []string)`

SetDhcpBmcMacAddressWhitelist sets DhcpBmcMacAddressWhitelist field to given value.

### HasDhcpBmcMacAddressWhitelist

`func (o *ServerPolicyUpdateDto) HasDhcpBmcMacAddressWhitelist() bool`

HasDhcpBmcMacAddressWhitelist returns a boolean if a field has been set.

### GetRaidConfigurationEnabled

`func (o *ServerPolicyUpdateDto) GetRaidConfigurationEnabled() bool`

GetRaidConfigurationEnabled returns the RaidConfigurationEnabled field if non-nil, zero value otherwise.

### GetRaidConfigurationEnabledOk

`func (o *ServerPolicyUpdateDto) GetRaidConfigurationEnabledOk() (*bool, bool)`

GetRaidConfigurationEnabledOk returns a tuple with the RaidConfigurationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidConfigurationEnabled

`func (o *ServerPolicyUpdateDto) SetRaidConfigurationEnabled(v bool)`

SetRaidConfigurationEnabled sets RaidConfigurationEnabled field to given value.

### HasRaidConfigurationEnabled

`func (o *ServerPolicyUpdateDto) HasRaidConfigurationEnabled() bool`

HasRaidConfigurationEnabled returns a boolean if a field has been set.

### GetDisableTpmAfterRegistration

`func (o *ServerPolicyUpdateDto) GetDisableTpmAfterRegistration() bool`

GetDisableTpmAfterRegistration returns the DisableTpmAfterRegistration field if non-nil, zero value otherwise.

### GetDisableTpmAfterRegistrationOk

`func (o *ServerPolicyUpdateDto) GetDisableTpmAfterRegistrationOk() (*bool, bool)`

GetDisableTpmAfterRegistrationOk returns a tuple with the DisableTpmAfterRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableTpmAfterRegistration

`func (o *ServerPolicyUpdateDto) SetDisableTpmAfterRegistration(v bool)`

SetDisableTpmAfterRegistration sets DisableTpmAfterRegistration field to given value.

### HasDisableTpmAfterRegistration

`func (o *ServerPolicyUpdateDto) HasDisableTpmAfterRegistration() bool`

HasDisableTpmAfterRegistration returns a boolean if a field has been set.

### GetSyslogMonitoringEnabled

`func (o *ServerPolicyUpdateDto) GetSyslogMonitoringEnabled() bool`

GetSyslogMonitoringEnabled returns the SyslogMonitoringEnabled field if non-nil, zero value otherwise.

### GetSyslogMonitoringEnabledOk

`func (o *ServerPolicyUpdateDto) GetSyslogMonitoringEnabledOk() (*bool, bool)`

GetSyslogMonitoringEnabledOk returns a tuple with the SyslogMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogMonitoringEnabled

`func (o *ServerPolicyUpdateDto) SetSyslogMonitoringEnabled(v bool)`

SetSyslogMonitoringEnabled sets SyslogMonitoringEnabled field to given value.

### HasSyslogMonitoringEnabled

`func (o *ServerPolicyUpdateDto) HasSyslogMonitoringEnabled() bool`

HasSyslogMonitoringEnabled returns a boolean if a field has been set.

### GetDefaultServerCleanupPolicyID

`func (o *ServerPolicyUpdateDto) GetDefaultServerCleanupPolicyID() float32`

GetDefaultServerCleanupPolicyID returns the DefaultServerCleanupPolicyID field if non-nil, zero value otherwise.

### GetDefaultServerCleanupPolicyIDOk

`func (o *ServerPolicyUpdateDto) GetDefaultServerCleanupPolicyIDOk() (*float32, bool)`

GetDefaultServerCleanupPolicyIDOk returns a tuple with the DefaultServerCleanupPolicyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultServerCleanupPolicyID

`func (o *ServerPolicyUpdateDto) SetDefaultServerCleanupPolicyID(v float32)`

SetDefaultServerCleanupPolicyID sets DefaultServerCleanupPolicyID field to given value.

### HasDefaultServerCleanupPolicyID

`func (o *ServerPolicyUpdateDto) HasDefaultServerCleanupPolicyID() bool`

HasDefaultServerCleanupPolicyID returns a boolean if a field has been set.

### GetAutomaticallyAllocateServerTypes

`func (o *ServerPolicyUpdateDto) GetAutomaticallyAllocateServerTypes() bool`

GetAutomaticallyAllocateServerTypes returns the AutomaticallyAllocateServerTypes field if non-nil, zero value otherwise.

### GetAutomaticallyAllocateServerTypesOk

`func (o *ServerPolicyUpdateDto) GetAutomaticallyAllocateServerTypesOk() (*bool, bool)`

GetAutomaticallyAllocateServerTypesOk returns a tuple with the AutomaticallyAllocateServerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticallyAllocateServerTypes

`func (o *ServerPolicyUpdateDto) SetAutomaticallyAllocateServerTypes(v bool)`

SetAutomaticallyAllocateServerTypes sets AutomaticallyAllocateServerTypes field to given value.

### HasAutomaticallyAllocateServerTypes

`func (o *ServerPolicyUpdateDto) HasAutomaticallyAllocateServerTypes() bool`

HasAutomaticallyAllocateServerTypes returns a boolean if a field has been set.

### GetAutomaticallySetServersAsAvailable

`func (o *ServerPolicyUpdateDto) GetAutomaticallySetServersAsAvailable() bool`

GetAutomaticallySetServersAsAvailable returns the AutomaticallySetServersAsAvailable field if non-nil, zero value otherwise.

### GetAutomaticallySetServersAsAvailableOk

`func (o *ServerPolicyUpdateDto) GetAutomaticallySetServersAsAvailableOk() (*bool, bool)`

GetAutomaticallySetServersAsAvailableOk returns a tuple with the AutomaticallySetServersAsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticallySetServersAsAvailable

`func (o *ServerPolicyUpdateDto) SetAutomaticallySetServersAsAvailable(v bool)`

SetAutomaticallySetServersAsAvailable sets AutomaticallySetServersAsAvailable field to given value.

### HasAutomaticallySetServersAsAvailable

`func (o *ServerPolicyUpdateDto) HasAutomaticallySetServersAsAvailable() bool`

HasAutomaticallySetServersAsAvailable returns a boolean if a field has been set.

### GetServerRegistrationBiosProfile

`func (o *ServerPolicyUpdateDto) GetServerRegistrationBiosProfile() []string`

GetServerRegistrationBiosProfile returns the ServerRegistrationBiosProfile field if non-nil, zero value otherwise.

### GetServerRegistrationBiosProfileOk

`func (o *ServerPolicyUpdateDto) GetServerRegistrationBiosProfileOk() (*[]string, bool)`

GetServerRegistrationBiosProfileOk returns a tuple with the ServerRegistrationBiosProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRegistrationBiosProfile

`func (o *ServerPolicyUpdateDto) SetServerRegistrationBiosProfile(v []string)`

SetServerRegistrationBiosProfile sets ServerRegistrationBiosProfile field to given value.

### HasServerRegistrationBiosProfile

`func (o *ServerPolicyUpdateDto) HasServerRegistrationBiosProfile() bool`

HasServerRegistrationBiosProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


