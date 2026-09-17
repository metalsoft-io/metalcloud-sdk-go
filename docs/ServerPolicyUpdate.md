# ServerPolicyUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultServerRegistrationProfileId** | Pointer to **int64** | ID of the default server registration profile | [optional] 
**DhcpOption82ToIPMapping** | Pointer to **map[string]interface{}** | DEPRECATED — use DHCP reservations instead. Option82 to IP mapping. | [optional] 
**DhcpAllocateUnmatched** | Pointer to **bool** | Whether to allocate an IP for a DHCP request that matches no reservation. Defaults to true (allocate the next free IP across all site OOB subnets); set to false to reject unmatched requests. | [optional] [default to true]
**DhcpBmcMacAddressWhitelistEnabled** | Pointer to **bool** | Whether to enable DHCP BMC MAC address whitelist | [optional] 
**DhcpBmcMacAddressWhitelist** | Pointer to **[]string** | List of DHCP BMC MAC address whitelist | [optional] 
**DefaultServerCleanupPolicyID** | Pointer to **int64** | Default server cleanup policy ID | [optional] 
**AutomaticallyAllocateServerTypes** | Pointer to **bool** | Automatically allocate server types | [optional] 
**AutomaticallySetServersAsAvailable** | Pointer to **bool** | Automatically set servers as available | [optional] 
**FirmwareBackupEnabled** | Pointer to **bool** | Whether periodic server firmware backup (snapshotting the observed firmware to git) is enabled for             the site. Disabling it also stops firmware drift detection. Defaults to true. | [optional] [default to true]
**FirmwareDriftDetectionEnabled** | Pointer to **bool** | Whether server drift detection is enabled for the site. Gates both checks: the configuration             drift alert (each snapshot against the previous one) and the firmware baseline check (installed             versions against the baseline minimums). Requires firmwareBackupEnabled. Defaults to true. | [optional] [default to true]

## Methods

### NewServerPolicyUpdate

`func NewServerPolicyUpdate() *ServerPolicyUpdate`

NewServerPolicyUpdate instantiates a new ServerPolicyUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerPolicyUpdateWithDefaults

`func NewServerPolicyUpdateWithDefaults() *ServerPolicyUpdate`

NewServerPolicyUpdateWithDefaults instantiates a new ServerPolicyUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultServerRegistrationProfileId

`func (o *ServerPolicyUpdate) GetDefaultServerRegistrationProfileId() int64`

GetDefaultServerRegistrationProfileId returns the DefaultServerRegistrationProfileId field if non-nil, zero value otherwise.

### GetDefaultServerRegistrationProfileIdOk

`func (o *ServerPolicyUpdate) GetDefaultServerRegistrationProfileIdOk() (*int64, bool)`

GetDefaultServerRegistrationProfileIdOk returns a tuple with the DefaultServerRegistrationProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultServerRegistrationProfileId

`func (o *ServerPolicyUpdate) SetDefaultServerRegistrationProfileId(v int64)`

SetDefaultServerRegistrationProfileId sets DefaultServerRegistrationProfileId field to given value.

### HasDefaultServerRegistrationProfileId

`func (o *ServerPolicyUpdate) HasDefaultServerRegistrationProfileId() bool`

HasDefaultServerRegistrationProfileId returns a boolean if a field has been set.

### GetDhcpOption82ToIPMapping

`func (o *ServerPolicyUpdate) GetDhcpOption82ToIPMapping() map[string]interface{}`

GetDhcpOption82ToIPMapping returns the DhcpOption82ToIPMapping field if non-nil, zero value otherwise.

### GetDhcpOption82ToIPMappingOk

`func (o *ServerPolicyUpdate) GetDhcpOption82ToIPMappingOk() (*map[string]interface{}, bool)`

GetDhcpOption82ToIPMappingOk returns a tuple with the DhcpOption82ToIPMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOption82ToIPMapping

`func (o *ServerPolicyUpdate) SetDhcpOption82ToIPMapping(v map[string]interface{})`

SetDhcpOption82ToIPMapping sets DhcpOption82ToIPMapping field to given value.

### HasDhcpOption82ToIPMapping

`func (o *ServerPolicyUpdate) HasDhcpOption82ToIPMapping() bool`

HasDhcpOption82ToIPMapping returns a boolean if a field has been set.

### GetDhcpAllocateUnmatched

`func (o *ServerPolicyUpdate) GetDhcpAllocateUnmatched() bool`

GetDhcpAllocateUnmatched returns the DhcpAllocateUnmatched field if non-nil, zero value otherwise.

### GetDhcpAllocateUnmatchedOk

`func (o *ServerPolicyUpdate) GetDhcpAllocateUnmatchedOk() (*bool, bool)`

GetDhcpAllocateUnmatchedOk returns a tuple with the DhcpAllocateUnmatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpAllocateUnmatched

`func (o *ServerPolicyUpdate) SetDhcpAllocateUnmatched(v bool)`

SetDhcpAllocateUnmatched sets DhcpAllocateUnmatched field to given value.

### HasDhcpAllocateUnmatched

`func (o *ServerPolicyUpdate) HasDhcpAllocateUnmatched() bool`

HasDhcpAllocateUnmatched returns a boolean if a field has been set.

### GetDhcpBmcMacAddressWhitelistEnabled

`func (o *ServerPolicyUpdate) GetDhcpBmcMacAddressWhitelistEnabled() bool`

GetDhcpBmcMacAddressWhitelistEnabled returns the DhcpBmcMacAddressWhitelistEnabled field if non-nil, zero value otherwise.

### GetDhcpBmcMacAddressWhitelistEnabledOk

`func (o *ServerPolicyUpdate) GetDhcpBmcMacAddressWhitelistEnabledOk() (*bool, bool)`

GetDhcpBmcMacAddressWhitelistEnabledOk returns a tuple with the DhcpBmcMacAddressWhitelistEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpBmcMacAddressWhitelistEnabled

`func (o *ServerPolicyUpdate) SetDhcpBmcMacAddressWhitelistEnabled(v bool)`

SetDhcpBmcMacAddressWhitelistEnabled sets DhcpBmcMacAddressWhitelistEnabled field to given value.

### HasDhcpBmcMacAddressWhitelistEnabled

`func (o *ServerPolicyUpdate) HasDhcpBmcMacAddressWhitelistEnabled() bool`

HasDhcpBmcMacAddressWhitelistEnabled returns a boolean if a field has been set.

### GetDhcpBmcMacAddressWhitelist

`func (o *ServerPolicyUpdate) GetDhcpBmcMacAddressWhitelist() []string`

GetDhcpBmcMacAddressWhitelist returns the DhcpBmcMacAddressWhitelist field if non-nil, zero value otherwise.

### GetDhcpBmcMacAddressWhitelistOk

`func (o *ServerPolicyUpdate) GetDhcpBmcMacAddressWhitelistOk() (*[]string, bool)`

GetDhcpBmcMacAddressWhitelistOk returns a tuple with the DhcpBmcMacAddressWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpBmcMacAddressWhitelist

`func (o *ServerPolicyUpdate) SetDhcpBmcMacAddressWhitelist(v []string)`

SetDhcpBmcMacAddressWhitelist sets DhcpBmcMacAddressWhitelist field to given value.

### HasDhcpBmcMacAddressWhitelist

`func (o *ServerPolicyUpdate) HasDhcpBmcMacAddressWhitelist() bool`

HasDhcpBmcMacAddressWhitelist returns a boolean if a field has been set.

### GetDefaultServerCleanupPolicyID

`func (o *ServerPolicyUpdate) GetDefaultServerCleanupPolicyID() int64`

GetDefaultServerCleanupPolicyID returns the DefaultServerCleanupPolicyID field if non-nil, zero value otherwise.

### GetDefaultServerCleanupPolicyIDOk

`func (o *ServerPolicyUpdate) GetDefaultServerCleanupPolicyIDOk() (*int64, bool)`

GetDefaultServerCleanupPolicyIDOk returns a tuple with the DefaultServerCleanupPolicyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultServerCleanupPolicyID

`func (o *ServerPolicyUpdate) SetDefaultServerCleanupPolicyID(v int64)`

SetDefaultServerCleanupPolicyID sets DefaultServerCleanupPolicyID field to given value.

### HasDefaultServerCleanupPolicyID

`func (o *ServerPolicyUpdate) HasDefaultServerCleanupPolicyID() bool`

HasDefaultServerCleanupPolicyID returns a boolean if a field has been set.

### GetAutomaticallyAllocateServerTypes

`func (o *ServerPolicyUpdate) GetAutomaticallyAllocateServerTypes() bool`

GetAutomaticallyAllocateServerTypes returns the AutomaticallyAllocateServerTypes field if non-nil, zero value otherwise.

### GetAutomaticallyAllocateServerTypesOk

`func (o *ServerPolicyUpdate) GetAutomaticallyAllocateServerTypesOk() (*bool, bool)`

GetAutomaticallyAllocateServerTypesOk returns a tuple with the AutomaticallyAllocateServerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticallyAllocateServerTypes

`func (o *ServerPolicyUpdate) SetAutomaticallyAllocateServerTypes(v bool)`

SetAutomaticallyAllocateServerTypes sets AutomaticallyAllocateServerTypes field to given value.

### HasAutomaticallyAllocateServerTypes

`func (o *ServerPolicyUpdate) HasAutomaticallyAllocateServerTypes() bool`

HasAutomaticallyAllocateServerTypes returns a boolean if a field has been set.

### GetAutomaticallySetServersAsAvailable

`func (o *ServerPolicyUpdate) GetAutomaticallySetServersAsAvailable() bool`

GetAutomaticallySetServersAsAvailable returns the AutomaticallySetServersAsAvailable field if non-nil, zero value otherwise.

### GetAutomaticallySetServersAsAvailableOk

`func (o *ServerPolicyUpdate) GetAutomaticallySetServersAsAvailableOk() (*bool, bool)`

GetAutomaticallySetServersAsAvailableOk returns a tuple with the AutomaticallySetServersAsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticallySetServersAsAvailable

`func (o *ServerPolicyUpdate) SetAutomaticallySetServersAsAvailable(v bool)`

SetAutomaticallySetServersAsAvailable sets AutomaticallySetServersAsAvailable field to given value.

### HasAutomaticallySetServersAsAvailable

`func (o *ServerPolicyUpdate) HasAutomaticallySetServersAsAvailable() bool`

HasAutomaticallySetServersAsAvailable returns a boolean if a field has been set.

### GetFirmwareBackupEnabled

`func (o *ServerPolicyUpdate) GetFirmwareBackupEnabled() bool`

GetFirmwareBackupEnabled returns the FirmwareBackupEnabled field if non-nil, zero value otherwise.

### GetFirmwareBackupEnabledOk

`func (o *ServerPolicyUpdate) GetFirmwareBackupEnabledOk() (*bool, bool)`

GetFirmwareBackupEnabledOk returns a tuple with the FirmwareBackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareBackupEnabled

`func (o *ServerPolicyUpdate) SetFirmwareBackupEnabled(v bool)`

SetFirmwareBackupEnabled sets FirmwareBackupEnabled field to given value.

### HasFirmwareBackupEnabled

`func (o *ServerPolicyUpdate) HasFirmwareBackupEnabled() bool`

HasFirmwareBackupEnabled returns a boolean if a field has been set.

### GetFirmwareDriftDetectionEnabled

`func (o *ServerPolicyUpdate) GetFirmwareDriftDetectionEnabled() bool`

GetFirmwareDriftDetectionEnabled returns the FirmwareDriftDetectionEnabled field if non-nil, zero value otherwise.

### GetFirmwareDriftDetectionEnabledOk

`func (o *ServerPolicyUpdate) GetFirmwareDriftDetectionEnabledOk() (*bool, bool)`

GetFirmwareDriftDetectionEnabledOk returns a tuple with the FirmwareDriftDetectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareDriftDetectionEnabled

`func (o *ServerPolicyUpdate) SetFirmwareDriftDetectionEnabled(v bool)`

SetFirmwareDriftDetectionEnabled sets FirmwareDriftDetectionEnabled field to given value.

### HasFirmwareDriftDetectionEnabled

`func (o *ServerPolicyUpdate) HasFirmwareDriftDetectionEnabled() bool`

HasFirmwareDriftDetectionEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


