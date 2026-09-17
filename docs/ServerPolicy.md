# ServerPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultServerRegistrationProfileId** | Pointer to **int64** | ID of the default server registration profile | [optional] 
**DhcpOption82ToIPMapping** | Pointer to **map[string]interface{}** | DEPRECATED — use DHCP reservations instead. Option82 to IP mapping. | [optional] 
**DhcpAllocateUnmatched** | Pointer to **bool** | Whether to allocate an IP for a DHCP request that matches no reservation. Defaults to true (allocate the next free IP across all site OOB subnets); set to false to reject unmatched requests. | [optional] [default to true]
**DhcpBmcMacAddressWhitelistEnabled** | **bool** | Whether to enable DHCP BMC MAC address whitelist | 
**DhcpBmcMacAddressWhitelist** | **[]string** | List of DHCP BMC MAC address whitelist | 
**DefaultServerCleanupPolicyID** | Pointer to **int64** | Default server cleanup policy ID | [optional] 
**AutomaticallyAllocateServerTypes** | **bool** | Automatically allocate server types | 
**AutomaticallySetServersAsAvailable** | **bool** | Automatically set servers as available | 
**FirmwareBackupEnabled** | Pointer to **bool** | Whether periodic server firmware backup (snapshotting the observed firmware to git) is enabled for             the site. Disabling it also stops firmware drift detection. Defaults to true. | [optional] [default to true]
**FirmwareDriftDetectionEnabled** | Pointer to **bool** | Whether server drift detection is enabled for the site. Gates both checks: the configuration             drift alert (each snapshot against the previous one) and the firmware baseline check (installed             versions against the baseline minimums). Requires firmwareBackupEnabled. Defaults to true. | [optional] [default to true]

## Methods

### NewServerPolicy

`func NewServerPolicy(dhcpBmcMacAddressWhitelistEnabled bool, dhcpBmcMacAddressWhitelist []string, automaticallyAllocateServerTypes bool, automaticallySetServersAsAvailable bool, ) *ServerPolicy`

NewServerPolicy instantiates a new ServerPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerPolicyWithDefaults

`func NewServerPolicyWithDefaults() *ServerPolicy`

NewServerPolicyWithDefaults instantiates a new ServerPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultServerRegistrationProfileId

`func (o *ServerPolicy) GetDefaultServerRegistrationProfileId() int64`

GetDefaultServerRegistrationProfileId returns the DefaultServerRegistrationProfileId field if non-nil, zero value otherwise.

### GetDefaultServerRegistrationProfileIdOk

`func (o *ServerPolicy) GetDefaultServerRegistrationProfileIdOk() (*int64, bool)`

GetDefaultServerRegistrationProfileIdOk returns a tuple with the DefaultServerRegistrationProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultServerRegistrationProfileId

`func (o *ServerPolicy) SetDefaultServerRegistrationProfileId(v int64)`

SetDefaultServerRegistrationProfileId sets DefaultServerRegistrationProfileId field to given value.

### HasDefaultServerRegistrationProfileId

`func (o *ServerPolicy) HasDefaultServerRegistrationProfileId() bool`

HasDefaultServerRegistrationProfileId returns a boolean if a field has been set.

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

### HasDhcpOption82ToIPMapping

`func (o *ServerPolicy) HasDhcpOption82ToIPMapping() bool`

HasDhcpOption82ToIPMapping returns a boolean if a field has been set.

### GetDhcpAllocateUnmatched

`func (o *ServerPolicy) GetDhcpAllocateUnmatched() bool`

GetDhcpAllocateUnmatched returns the DhcpAllocateUnmatched field if non-nil, zero value otherwise.

### GetDhcpAllocateUnmatchedOk

`func (o *ServerPolicy) GetDhcpAllocateUnmatchedOk() (*bool, bool)`

GetDhcpAllocateUnmatchedOk returns a tuple with the DhcpAllocateUnmatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpAllocateUnmatched

`func (o *ServerPolicy) SetDhcpAllocateUnmatched(v bool)`

SetDhcpAllocateUnmatched sets DhcpAllocateUnmatched field to given value.

### HasDhcpAllocateUnmatched

`func (o *ServerPolicy) HasDhcpAllocateUnmatched() bool`

HasDhcpAllocateUnmatched returns a boolean if a field has been set.

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


### GetDefaultServerCleanupPolicyID

`func (o *ServerPolicy) GetDefaultServerCleanupPolicyID() int64`

GetDefaultServerCleanupPolicyID returns the DefaultServerCleanupPolicyID field if non-nil, zero value otherwise.

### GetDefaultServerCleanupPolicyIDOk

`func (o *ServerPolicy) GetDefaultServerCleanupPolicyIDOk() (*int64, bool)`

GetDefaultServerCleanupPolicyIDOk returns a tuple with the DefaultServerCleanupPolicyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultServerCleanupPolicyID

`func (o *ServerPolicy) SetDefaultServerCleanupPolicyID(v int64)`

SetDefaultServerCleanupPolicyID sets DefaultServerCleanupPolicyID field to given value.

### HasDefaultServerCleanupPolicyID

`func (o *ServerPolicy) HasDefaultServerCleanupPolicyID() bool`

HasDefaultServerCleanupPolicyID returns a boolean if a field has been set.

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


### GetFirmwareBackupEnabled

`func (o *ServerPolicy) GetFirmwareBackupEnabled() bool`

GetFirmwareBackupEnabled returns the FirmwareBackupEnabled field if non-nil, zero value otherwise.

### GetFirmwareBackupEnabledOk

`func (o *ServerPolicy) GetFirmwareBackupEnabledOk() (*bool, bool)`

GetFirmwareBackupEnabledOk returns a tuple with the FirmwareBackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareBackupEnabled

`func (o *ServerPolicy) SetFirmwareBackupEnabled(v bool)`

SetFirmwareBackupEnabled sets FirmwareBackupEnabled field to given value.

### HasFirmwareBackupEnabled

`func (o *ServerPolicy) HasFirmwareBackupEnabled() bool`

HasFirmwareBackupEnabled returns a boolean if a field has been set.

### GetFirmwareDriftDetectionEnabled

`func (o *ServerPolicy) GetFirmwareDriftDetectionEnabled() bool`

GetFirmwareDriftDetectionEnabled returns the FirmwareDriftDetectionEnabled field if non-nil, zero value otherwise.

### GetFirmwareDriftDetectionEnabledOk

`func (o *ServerPolicy) GetFirmwareDriftDetectionEnabledOk() (*bool, bool)`

GetFirmwareDriftDetectionEnabledOk returns a tuple with the FirmwareDriftDetectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareDriftDetectionEnabled

`func (o *ServerPolicy) SetFirmwareDriftDetectionEnabled(v bool)`

SetFirmwareDriftDetectionEnabled sets FirmwareDriftDetectionEnabled field to given value.

### HasFirmwareDriftDetectionEnabled

`func (o *ServerPolicy) HasFirmwareDriftDetectionEnabled() bool`

HasFirmwareDriftDetectionEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


