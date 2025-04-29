# ServerPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegisterCredentials** | **string** | Server registration type | 
**MinimumNumberOfConnectedInterfaces** | **float32** | Minimum number of switch-connected interfaces required | 
**DhcpOption82ToIPMapping** | **map[string]interface{}** | Option82 to IP mapping | 
**DhcpBmcMacAddressWhitelistEnabled** | **bool** | Whether to enable DHCP BMC MAC address whitelist | 
**DhcpBmcMacAddressWhitelist** | **[]string** | List of DHCP BMC MAC address whitelist | 
**RaidConfigurationEnabled** | **bool** | Whether RAID configuration is enabled | 
**DisableTpmAfterRegistration** | **bool** | Whether to disable TPM after registration | 
**SyslogMonitoringEnabled** | **bool** | Whether syslog monitoring is enabled | 
**DefaultServerCleanupPolicyID** | **float32** | Default server cleanup policy ID | 
**AutomaticallyAllocateServerTypes** | **bool** | Automatically allocate server types | 
**AutomaticallySetServersAsAvailable** | **bool** | Automatically set servers as available | 
**ServerRegistrationBiosProfile** | [**[]ServerRegistrationBiosProfile**](ServerRegistrationBiosProfile.md) | Server registration BIOS profile | 

## Methods

### NewServerPolicy

`func NewServerPolicy(registerCredentials string, minimumNumberOfConnectedInterfaces float32, dhcpOption82ToIPMapping map[string]interface{}, dhcpBmcMacAddressWhitelistEnabled bool, dhcpBmcMacAddressWhitelist []string, raidConfigurationEnabled bool, disableTpmAfterRegistration bool, syslogMonitoringEnabled bool, defaultServerCleanupPolicyID float32, automaticallyAllocateServerTypes bool, automaticallySetServersAsAvailable bool, serverRegistrationBiosProfile []ServerRegistrationBiosProfile, ) *ServerPolicy`

NewServerPolicy instantiates a new ServerPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerPolicyWithDefaults

`func NewServerPolicyWithDefaults() *ServerPolicy`

NewServerPolicyWithDefaults instantiates a new ServerPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegisterCredentials

`func (o *ServerPolicy) GetRegisterCredentials() string`

GetRegisterCredentials returns the RegisterCredentials field if non-nil, zero value otherwise.

### GetRegisterCredentialsOk

`func (o *ServerPolicy) GetRegisterCredentialsOk() (*string, bool)`

GetRegisterCredentialsOk returns a tuple with the RegisterCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterCredentials

`func (o *ServerPolicy) SetRegisterCredentials(v string)`

SetRegisterCredentials sets RegisterCredentials field to given value.


### GetMinimumNumberOfConnectedInterfaces

`func (o *ServerPolicy) GetMinimumNumberOfConnectedInterfaces() float32`

GetMinimumNumberOfConnectedInterfaces returns the MinimumNumberOfConnectedInterfaces field if non-nil, zero value otherwise.

### GetMinimumNumberOfConnectedInterfacesOk

`func (o *ServerPolicy) GetMinimumNumberOfConnectedInterfacesOk() (*float32, bool)`

GetMinimumNumberOfConnectedInterfacesOk returns a tuple with the MinimumNumberOfConnectedInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumNumberOfConnectedInterfaces

`func (o *ServerPolicy) SetMinimumNumberOfConnectedInterfaces(v float32)`

SetMinimumNumberOfConnectedInterfaces sets MinimumNumberOfConnectedInterfaces field to given value.


### GetDhcpOption82ToIPMapping

`func (o *ServerPolicy) GetDhcpOption82ToIPMapping() map[string]interface{}`

GetDhcpOption82ToIPMapping returns the DhcpOption82ToIPMapping field if non-nil, zero value otherwise.

### GetDhcpOption82ToIPMappingOk

`func (o *ServerPolicy) GetDhcpOption82ToIPMappingOk() (*map[string]interface{}, bool)`

GetDhcpOption82ToIPMappingOk returns a tuple with the DhcpOption82ToIPMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOption82ToIPMapping

`func (o *ServerPolicy) SetDhcpOption82ToIPMapping(v map[string]interface{})`

SetDhcpOption82ToIPMapping sets DhcpOption82ToIPMapping field to given value.


### GetDhcpBmcMacAddressWhitelistEnabled

`func (o *ServerPolicy) GetDhcpBmcMacAddressWhitelistEnabled() bool`

GetDhcpBmcMacAddressWhitelistEnabled returns the DhcpBmcMacAddressWhitelistEnabled field if non-nil, zero value otherwise.

### GetDhcpBmcMacAddressWhitelistEnabledOk

`func (o *ServerPolicy) GetDhcpBmcMacAddressWhitelistEnabledOk() (*bool, bool)`

GetDhcpBmcMacAddressWhitelistEnabledOk returns a tuple with the DhcpBmcMacAddressWhitelistEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpBmcMacAddressWhitelistEnabled

`func (o *ServerPolicy) SetDhcpBmcMacAddressWhitelistEnabled(v bool)`

SetDhcpBmcMacAddressWhitelistEnabled sets DhcpBmcMacAddressWhitelistEnabled field to given value.


### GetDhcpBmcMacAddressWhitelist

`func (o *ServerPolicy) GetDhcpBmcMacAddressWhitelist() []string`

GetDhcpBmcMacAddressWhitelist returns the DhcpBmcMacAddressWhitelist field if non-nil, zero value otherwise.

### GetDhcpBmcMacAddressWhitelistOk

`func (o *ServerPolicy) GetDhcpBmcMacAddressWhitelistOk() (*[]string, bool)`

GetDhcpBmcMacAddressWhitelistOk returns a tuple with the DhcpBmcMacAddressWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpBmcMacAddressWhitelist

`func (o *ServerPolicy) SetDhcpBmcMacAddressWhitelist(v []string)`

SetDhcpBmcMacAddressWhitelist sets DhcpBmcMacAddressWhitelist field to given value.


### GetRaidConfigurationEnabled

`func (o *ServerPolicy) GetRaidConfigurationEnabled() bool`

GetRaidConfigurationEnabled returns the RaidConfigurationEnabled field if non-nil, zero value otherwise.

### GetRaidConfigurationEnabledOk

`func (o *ServerPolicy) GetRaidConfigurationEnabledOk() (*bool, bool)`

GetRaidConfigurationEnabledOk returns a tuple with the RaidConfigurationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidConfigurationEnabled

`func (o *ServerPolicy) SetRaidConfigurationEnabled(v bool)`

SetRaidConfigurationEnabled sets RaidConfigurationEnabled field to given value.


### GetDisableTpmAfterRegistration

`func (o *ServerPolicy) GetDisableTpmAfterRegistration() bool`

GetDisableTpmAfterRegistration returns the DisableTpmAfterRegistration field if non-nil, zero value otherwise.

### GetDisableTpmAfterRegistrationOk

`func (o *ServerPolicy) GetDisableTpmAfterRegistrationOk() (*bool, bool)`

GetDisableTpmAfterRegistrationOk returns a tuple with the DisableTpmAfterRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableTpmAfterRegistration

`func (o *ServerPolicy) SetDisableTpmAfterRegistration(v bool)`

SetDisableTpmAfterRegistration sets DisableTpmAfterRegistration field to given value.


### GetSyslogMonitoringEnabled

`func (o *ServerPolicy) GetSyslogMonitoringEnabled() bool`

GetSyslogMonitoringEnabled returns the SyslogMonitoringEnabled field if non-nil, zero value otherwise.

### GetSyslogMonitoringEnabledOk

`func (o *ServerPolicy) GetSyslogMonitoringEnabledOk() (*bool, bool)`

GetSyslogMonitoringEnabledOk returns a tuple with the SyslogMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogMonitoringEnabled

`func (o *ServerPolicy) SetSyslogMonitoringEnabled(v bool)`

SetSyslogMonitoringEnabled sets SyslogMonitoringEnabled field to given value.


### GetDefaultServerCleanupPolicyID

`func (o *ServerPolicy) GetDefaultServerCleanupPolicyID() float32`

GetDefaultServerCleanupPolicyID returns the DefaultServerCleanupPolicyID field if non-nil, zero value otherwise.

### GetDefaultServerCleanupPolicyIDOk

`func (o *ServerPolicy) GetDefaultServerCleanupPolicyIDOk() (*float32, bool)`

GetDefaultServerCleanupPolicyIDOk returns a tuple with the DefaultServerCleanupPolicyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultServerCleanupPolicyID

`func (o *ServerPolicy) SetDefaultServerCleanupPolicyID(v float32)`

SetDefaultServerCleanupPolicyID sets DefaultServerCleanupPolicyID field to given value.


### GetAutomaticallyAllocateServerTypes

`func (o *ServerPolicy) GetAutomaticallyAllocateServerTypes() bool`

GetAutomaticallyAllocateServerTypes returns the AutomaticallyAllocateServerTypes field if non-nil, zero value otherwise.

### GetAutomaticallyAllocateServerTypesOk

`func (o *ServerPolicy) GetAutomaticallyAllocateServerTypesOk() (*bool, bool)`

GetAutomaticallyAllocateServerTypesOk returns a tuple with the AutomaticallyAllocateServerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticallyAllocateServerTypes

`func (o *ServerPolicy) SetAutomaticallyAllocateServerTypes(v bool)`

SetAutomaticallyAllocateServerTypes sets AutomaticallyAllocateServerTypes field to given value.


### GetAutomaticallySetServersAsAvailable

`func (o *ServerPolicy) GetAutomaticallySetServersAsAvailable() bool`

GetAutomaticallySetServersAsAvailable returns the AutomaticallySetServersAsAvailable field if non-nil, zero value otherwise.

### GetAutomaticallySetServersAsAvailableOk

`func (o *ServerPolicy) GetAutomaticallySetServersAsAvailableOk() (*bool, bool)`

GetAutomaticallySetServersAsAvailableOk returns a tuple with the AutomaticallySetServersAsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticallySetServersAsAvailable

`func (o *ServerPolicy) SetAutomaticallySetServersAsAvailable(v bool)`

SetAutomaticallySetServersAsAvailable sets AutomaticallySetServersAsAvailable field to given value.


### GetServerRegistrationBiosProfile

`func (o *ServerPolicy) GetServerRegistrationBiosProfile() []ServerRegistrationBiosProfile`

GetServerRegistrationBiosProfile returns the ServerRegistrationBiosProfile field if non-nil, zero value otherwise.

### GetServerRegistrationBiosProfileOk

`func (o *ServerPolicy) GetServerRegistrationBiosProfileOk() (*[]ServerRegistrationBiosProfile, bool)`

GetServerRegistrationBiosProfileOk returns a tuple with the ServerRegistrationBiosProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRegistrationBiosProfile

`func (o *ServerPolicy) SetServerRegistrationBiosProfile(v []ServerRegistrationBiosProfile)`

SetServerRegistrationBiosProfile sets ServerRegistrationBiosProfile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


