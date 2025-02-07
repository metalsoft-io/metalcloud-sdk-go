# ServerInstanceGroupConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | **float32** | Revision number | 
**Label** | Pointer to **string** | The server instance group label. Will be automatically generated if not provided. | [optional] 
**ServerGroupName** | Pointer to **string** |  | [optional] 
**UpdatedTimestamp** | **string** | Timestamp of the latest update for the Server Instance Group. | 
**Subdomain** | Pointer to **string** | Subdomain of the Server Group. | [optional] 
**InstanceCount** | **int32** |  | [default to 1]
**IpAllocateAuto** | **float32** |  | 
**Ipv4SubnetCreateAuto** | **float32** |  | 
**FirewallProfileId** | Pointer to **int32** |  | [optional] 
**FirewallRulesSetId** | Pointer to **int32** |  | [optional] 
**FirewallManaged** | **float32** |  | 
**FirmwarePoliciesJson** | **map[string]interface{}** |  | 
**VolumeTemplateId** | Pointer to **int32** |  | [optional] 
**DriveArrayIdBoot** | Pointer to **int32** | Id of the bootable drive for the Server Instance Group. | [optional] 
**InstanceArrayBootMethod** | **string** | Instance Array Boot Method | 
**CustomVariables** | Pointer to **map[string]interface{}** |  | [optional] 
**ProcessorCoreCount** | **int32** |  | 
**ProcessorCoreMhz** | **float32** |  | 
**ProcessorCount** | **int32** |  | 
**RamGbytes** | **float32** |  | 
**DiskCount** | **int32** |  | 
**DiskSizeMbytes** | **int32** |  | 
**DiskTypes** | **[]string** |  | 
**VirtualInterfacesEnabled** | **float32** |  | 
**AdditionalWanIpv4Json** | Pointer to **map[string]interface{}** |  | [optional] 
**NetworkProfileGroupId** | Pointer to **int32** |  | [optional] 
**OverrideIpv4WanVlanId** | Pointer to **int32** |  | [optional] 
**NetworkEquipmentForceSubnetPoolIpv4WanId** | Pointer to **int32** |  | [optional] 
**DefaultServerProfileID** | **int32** | The group&#39;s default server profile. Useful when creating a server instance with a group id set, the profile will be automatically applied. | 
**EmptyEdit** | Pointer to **int32** | Number of empty edits | [optional] 
**DeployType** | **string** | Server Instance Group deploy type | 
**DeployStatus** | **string** | Server Instance Group deploy status | 
**DnsSubdomainChangeId** | Pointer to **int32** | Id of the DNS subdomain for the Server Instance Group. | [optional] 
**InfrastructureDeployId** | Pointer to **int32** | Id of the deployment for the Instance Group. | [optional] 

## Methods

### NewServerInstanceGroupConfiguration

`func NewServerInstanceGroupConfiguration(revision float32, updatedTimestamp string, instanceCount int32, ipAllocateAuto float32, ipv4SubnetCreateAuto float32, firewallManaged float32, firmwarePoliciesJson map[string]interface{}, instanceArrayBootMethod string, processorCoreCount int32, processorCoreMhz float32, processorCount int32, ramGbytes float32, diskCount int32, diskSizeMbytes int32, diskTypes []string, virtualInterfacesEnabled float32, defaultServerProfileID int32, deployType string, deployStatus string, ) *ServerInstanceGroupConfiguration`

NewServerInstanceGroupConfiguration instantiates a new ServerInstanceGroupConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceGroupConfigurationWithDefaults

`func NewServerInstanceGroupConfigurationWithDefaults() *ServerInstanceGroupConfiguration`

NewServerInstanceGroupConfigurationWithDefaults instantiates a new ServerInstanceGroupConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *ServerInstanceGroupConfiguration) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ServerInstanceGroupConfiguration) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ServerInstanceGroupConfiguration) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetLabel

`func (o *ServerInstanceGroupConfiguration) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServerInstanceGroupConfiguration) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServerInstanceGroupConfiguration) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ServerInstanceGroupConfiguration) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetServerGroupName

`func (o *ServerInstanceGroupConfiguration) GetServerGroupName() string`

GetServerGroupName returns the ServerGroupName field if non-nil, zero value otherwise.

### GetServerGroupNameOk

`func (o *ServerInstanceGroupConfiguration) GetServerGroupNameOk() (*string, bool)`

GetServerGroupNameOk returns a tuple with the ServerGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroupName

`func (o *ServerInstanceGroupConfiguration) SetServerGroupName(v string)`

SetServerGroupName sets ServerGroupName field to given value.

### HasServerGroupName

`func (o *ServerInstanceGroupConfiguration) HasServerGroupName() bool`

HasServerGroupName returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *ServerInstanceGroupConfiguration) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *ServerInstanceGroupConfiguration) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *ServerInstanceGroupConfiguration) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetSubdomain

`func (o *ServerInstanceGroupConfiguration) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *ServerInstanceGroupConfiguration) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *ServerInstanceGroupConfiguration) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *ServerInstanceGroupConfiguration) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetInstanceCount

`func (o *ServerInstanceGroupConfiguration) GetInstanceCount() int32`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *ServerInstanceGroupConfiguration) GetInstanceCountOk() (*int32, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *ServerInstanceGroupConfiguration) SetInstanceCount(v int32)`

SetInstanceCount sets InstanceCount field to given value.


### GetIpAllocateAuto

`func (o *ServerInstanceGroupConfiguration) GetIpAllocateAuto() float32`

GetIpAllocateAuto returns the IpAllocateAuto field if non-nil, zero value otherwise.

### GetIpAllocateAutoOk

`func (o *ServerInstanceGroupConfiguration) GetIpAllocateAutoOk() (*float32, bool)`

GetIpAllocateAutoOk returns a tuple with the IpAllocateAuto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAllocateAuto

`func (o *ServerInstanceGroupConfiguration) SetIpAllocateAuto(v float32)`

SetIpAllocateAuto sets IpAllocateAuto field to given value.


### GetIpv4SubnetCreateAuto

`func (o *ServerInstanceGroupConfiguration) GetIpv4SubnetCreateAuto() float32`

GetIpv4SubnetCreateAuto returns the Ipv4SubnetCreateAuto field if non-nil, zero value otherwise.

### GetIpv4SubnetCreateAutoOk

`func (o *ServerInstanceGroupConfiguration) GetIpv4SubnetCreateAutoOk() (*float32, bool)`

GetIpv4SubnetCreateAutoOk returns a tuple with the Ipv4SubnetCreateAuto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4SubnetCreateAuto

`func (o *ServerInstanceGroupConfiguration) SetIpv4SubnetCreateAuto(v float32)`

SetIpv4SubnetCreateAuto sets Ipv4SubnetCreateAuto field to given value.


### GetFirewallProfileId

`func (o *ServerInstanceGroupConfiguration) GetFirewallProfileId() int32`

GetFirewallProfileId returns the FirewallProfileId field if non-nil, zero value otherwise.

### GetFirewallProfileIdOk

`func (o *ServerInstanceGroupConfiguration) GetFirewallProfileIdOk() (*int32, bool)`

GetFirewallProfileIdOk returns a tuple with the FirewallProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallProfileId

`func (o *ServerInstanceGroupConfiguration) SetFirewallProfileId(v int32)`

SetFirewallProfileId sets FirewallProfileId field to given value.

### HasFirewallProfileId

`func (o *ServerInstanceGroupConfiguration) HasFirewallProfileId() bool`

HasFirewallProfileId returns a boolean if a field has been set.

### GetFirewallRulesSetId

`func (o *ServerInstanceGroupConfiguration) GetFirewallRulesSetId() int32`

GetFirewallRulesSetId returns the FirewallRulesSetId field if non-nil, zero value otherwise.

### GetFirewallRulesSetIdOk

`func (o *ServerInstanceGroupConfiguration) GetFirewallRulesSetIdOk() (*int32, bool)`

GetFirewallRulesSetIdOk returns a tuple with the FirewallRulesSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallRulesSetId

`func (o *ServerInstanceGroupConfiguration) SetFirewallRulesSetId(v int32)`

SetFirewallRulesSetId sets FirewallRulesSetId field to given value.

### HasFirewallRulesSetId

`func (o *ServerInstanceGroupConfiguration) HasFirewallRulesSetId() bool`

HasFirewallRulesSetId returns a boolean if a field has been set.

### GetFirewallManaged

`func (o *ServerInstanceGroupConfiguration) GetFirewallManaged() float32`

GetFirewallManaged returns the FirewallManaged field if non-nil, zero value otherwise.

### GetFirewallManagedOk

`func (o *ServerInstanceGroupConfiguration) GetFirewallManagedOk() (*float32, bool)`

GetFirewallManagedOk returns a tuple with the FirewallManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallManaged

`func (o *ServerInstanceGroupConfiguration) SetFirewallManaged(v float32)`

SetFirewallManaged sets FirewallManaged field to given value.


### GetFirmwarePoliciesJson

`func (o *ServerInstanceGroupConfiguration) GetFirmwarePoliciesJson() map[string]interface{}`

GetFirmwarePoliciesJson returns the FirmwarePoliciesJson field if non-nil, zero value otherwise.

### GetFirmwarePoliciesJsonOk

`func (o *ServerInstanceGroupConfiguration) GetFirmwarePoliciesJsonOk() (*map[string]interface{}, bool)`

GetFirmwarePoliciesJsonOk returns a tuple with the FirmwarePoliciesJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwarePoliciesJson

`func (o *ServerInstanceGroupConfiguration) SetFirmwarePoliciesJson(v map[string]interface{})`

SetFirmwarePoliciesJson sets FirmwarePoliciesJson field to given value.


### GetVolumeTemplateId

`func (o *ServerInstanceGroupConfiguration) GetVolumeTemplateId() int32`

GetVolumeTemplateId returns the VolumeTemplateId field if non-nil, zero value otherwise.

### GetVolumeTemplateIdOk

`func (o *ServerInstanceGroupConfiguration) GetVolumeTemplateIdOk() (*int32, bool)`

GetVolumeTemplateIdOk returns a tuple with the VolumeTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeTemplateId

`func (o *ServerInstanceGroupConfiguration) SetVolumeTemplateId(v int32)`

SetVolumeTemplateId sets VolumeTemplateId field to given value.

### HasVolumeTemplateId

`func (o *ServerInstanceGroupConfiguration) HasVolumeTemplateId() bool`

HasVolumeTemplateId returns a boolean if a field has been set.

### GetDriveArrayIdBoot

`func (o *ServerInstanceGroupConfiguration) GetDriveArrayIdBoot() int32`

GetDriveArrayIdBoot returns the DriveArrayIdBoot field if non-nil, zero value otherwise.

### GetDriveArrayIdBootOk

`func (o *ServerInstanceGroupConfiguration) GetDriveArrayIdBootOk() (*int32, bool)`

GetDriveArrayIdBootOk returns a tuple with the DriveArrayIdBoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveArrayIdBoot

`func (o *ServerInstanceGroupConfiguration) SetDriveArrayIdBoot(v int32)`

SetDriveArrayIdBoot sets DriveArrayIdBoot field to given value.

### HasDriveArrayIdBoot

`func (o *ServerInstanceGroupConfiguration) HasDriveArrayIdBoot() bool`

HasDriveArrayIdBoot returns a boolean if a field has been set.

### GetInstanceArrayBootMethod

`func (o *ServerInstanceGroupConfiguration) GetInstanceArrayBootMethod() string`

GetInstanceArrayBootMethod returns the InstanceArrayBootMethod field if non-nil, zero value otherwise.

### GetInstanceArrayBootMethodOk

`func (o *ServerInstanceGroupConfiguration) GetInstanceArrayBootMethodOk() (*string, bool)`

GetInstanceArrayBootMethodOk returns a tuple with the InstanceArrayBootMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceArrayBootMethod

`func (o *ServerInstanceGroupConfiguration) SetInstanceArrayBootMethod(v string)`

SetInstanceArrayBootMethod sets InstanceArrayBootMethod field to given value.


### GetCustomVariables

`func (o *ServerInstanceGroupConfiguration) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *ServerInstanceGroupConfiguration) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *ServerInstanceGroupConfiguration) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *ServerInstanceGroupConfiguration) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.

### GetProcessorCoreCount

`func (o *ServerInstanceGroupConfiguration) GetProcessorCoreCount() int32`

GetProcessorCoreCount returns the ProcessorCoreCount field if non-nil, zero value otherwise.

### GetProcessorCoreCountOk

`func (o *ServerInstanceGroupConfiguration) GetProcessorCoreCountOk() (*int32, bool)`

GetProcessorCoreCountOk returns a tuple with the ProcessorCoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCoreCount

`func (o *ServerInstanceGroupConfiguration) SetProcessorCoreCount(v int32)`

SetProcessorCoreCount sets ProcessorCoreCount field to given value.


### GetProcessorCoreMhz

`func (o *ServerInstanceGroupConfiguration) GetProcessorCoreMhz() float32`

GetProcessorCoreMhz returns the ProcessorCoreMhz field if non-nil, zero value otherwise.

### GetProcessorCoreMhzOk

`func (o *ServerInstanceGroupConfiguration) GetProcessorCoreMhzOk() (*float32, bool)`

GetProcessorCoreMhzOk returns a tuple with the ProcessorCoreMhz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCoreMhz

`func (o *ServerInstanceGroupConfiguration) SetProcessorCoreMhz(v float32)`

SetProcessorCoreMhz sets ProcessorCoreMhz field to given value.


### GetProcessorCount

`func (o *ServerInstanceGroupConfiguration) GetProcessorCount() int32`

GetProcessorCount returns the ProcessorCount field if non-nil, zero value otherwise.

### GetProcessorCountOk

`func (o *ServerInstanceGroupConfiguration) GetProcessorCountOk() (*int32, bool)`

GetProcessorCountOk returns a tuple with the ProcessorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCount

`func (o *ServerInstanceGroupConfiguration) SetProcessorCount(v int32)`

SetProcessorCount sets ProcessorCount field to given value.


### GetRamGbytes

`func (o *ServerInstanceGroupConfiguration) GetRamGbytes() float32`

GetRamGbytes returns the RamGbytes field if non-nil, zero value otherwise.

### GetRamGbytesOk

`func (o *ServerInstanceGroupConfiguration) GetRamGbytesOk() (*float32, bool)`

GetRamGbytesOk returns a tuple with the RamGbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGbytes

`func (o *ServerInstanceGroupConfiguration) SetRamGbytes(v float32)`

SetRamGbytes sets RamGbytes field to given value.


### GetDiskCount

`func (o *ServerInstanceGroupConfiguration) GetDiskCount() int32`

GetDiskCount returns the DiskCount field if non-nil, zero value otherwise.

### GetDiskCountOk

`func (o *ServerInstanceGroupConfiguration) GetDiskCountOk() (*int32, bool)`

GetDiskCountOk returns a tuple with the DiskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskCount

`func (o *ServerInstanceGroupConfiguration) SetDiskCount(v int32)`

SetDiskCount sets DiskCount field to given value.


### GetDiskSizeMbytes

`func (o *ServerInstanceGroupConfiguration) GetDiskSizeMbytes() int32`

GetDiskSizeMbytes returns the DiskSizeMbytes field if non-nil, zero value otherwise.

### GetDiskSizeMbytesOk

`func (o *ServerInstanceGroupConfiguration) GetDiskSizeMbytesOk() (*int32, bool)`

GetDiskSizeMbytesOk returns a tuple with the DiskSizeMbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeMbytes

`func (o *ServerInstanceGroupConfiguration) SetDiskSizeMbytes(v int32)`

SetDiskSizeMbytes sets DiskSizeMbytes field to given value.


### GetDiskTypes

`func (o *ServerInstanceGroupConfiguration) GetDiskTypes() []string`

GetDiskTypes returns the DiskTypes field if non-nil, zero value otherwise.

### GetDiskTypesOk

`func (o *ServerInstanceGroupConfiguration) GetDiskTypesOk() (*[]string, bool)`

GetDiskTypesOk returns a tuple with the DiskTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskTypes

`func (o *ServerInstanceGroupConfiguration) SetDiskTypes(v []string)`

SetDiskTypes sets DiskTypes field to given value.


### GetVirtualInterfacesEnabled

`func (o *ServerInstanceGroupConfiguration) GetVirtualInterfacesEnabled() float32`

GetVirtualInterfacesEnabled returns the VirtualInterfacesEnabled field if non-nil, zero value otherwise.

### GetVirtualInterfacesEnabledOk

`func (o *ServerInstanceGroupConfiguration) GetVirtualInterfacesEnabledOk() (*float32, bool)`

GetVirtualInterfacesEnabledOk returns a tuple with the VirtualInterfacesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualInterfacesEnabled

`func (o *ServerInstanceGroupConfiguration) SetVirtualInterfacesEnabled(v float32)`

SetVirtualInterfacesEnabled sets VirtualInterfacesEnabled field to given value.


### GetAdditionalWanIpv4Json

`func (o *ServerInstanceGroupConfiguration) GetAdditionalWanIpv4Json() map[string]interface{}`

GetAdditionalWanIpv4Json returns the AdditionalWanIpv4Json field if non-nil, zero value otherwise.

### GetAdditionalWanIpv4JsonOk

`func (o *ServerInstanceGroupConfiguration) GetAdditionalWanIpv4JsonOk() (*map[string]interface{}, bool)`

GetAdditionalWanIpv4JsonOk returns a tuple with the AdditionalWanIpv4Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalWanIpv4Json

`func (o *ServerInstanceGroupConfiguration) SetAdditionalWanIpv4Json(v map[string]interface{})`

SetAdditionalWanIpv4Json sets AdditionalWanIpv4Json field to given value.

### HasAdditionalWanIpv4Json

`func (o *ServerInstanceGroupConfiguration) HasAdditionalWanIpv4Json() bool`

HasAdditionalWanIpv4Json returns a boolean if a field has been set.

### GetNetworkProfileGroupId

`func (o *ServerInstanceGroupConfiguration) GetNetworkProfileGroupId() int32`

GetNetworkProfileGroupId returns the NetworkProfileGroupId field if non-nil, zero value otherwise.

### GetNetworkProfileGroupIdOk

`func (o *ServerInstanceGroupConfiguration) GetNetworkProfileGroupIdOk() (*int32, bool)`

GetNetworkProfileGroupIdOk returns a tuple with the NetworkProfileGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProfileGroupId

`func (o *ServerInstanceGroupConfiguration) SetNetworkProfileGroupId(v int32)`

SetNetworkProfileGroupId sets NetworkProfileGroupId field to given value.

### HasNetworkProfileGroupId

`func (o *ServerInstanceGroupConfiguration) HasNetworkProfileGroupId() bool`

HasNetworkProfileGroupId returns a boolean if a field has been set.

### GetOverrideIpv4WanVlanId

`func (o *ServerInstanceGroupConfiguration) GetOverrideIpv4WanVlanId() int32`

GetOverrideIpv4WanVlanId returns the OverrideIpv4WanVlanId field if non-nil, zero value otherwise.

### GetOverrideIpv4WanVlanIdOk

`func (o *ServerInstanceGroupConfiguration) GetOverrideIpv4WanVlanIdOk() (*int32, bool)`

GetOverrideIpv4WanVlanIdOk returns a tuple with the OverrideIpv4WanVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideIpv4WanVlanId

`func (o *ServerInstanceGroupConfiguration) SetOverrideIpv4WanVlanId(v int32)`

SetOverrideIpv4WanVlanId sets OverrideIpv4WanVlanId field to given value.

### HasOverrideIpv4WanVlanId

`func (o *ServerInstanceGroupConfiguration) HasOverrideIpv4WanVlanId() bool`

HasOverrideIpv4WanVlanId returns a boolean if a field has been set.

### GetNetworkEquipmentForceSubnetPoolIpv4WanId

`func (o *ServerInstanceGroupConfiguration) GetNetworkEquipmentForceSubnetPoolIpv4WanId() int32`

GetNetworkEquipmentForceSubnetPoolIpv4WanId returns the NetworkEquipmentForceSubnetPoolIpv4WanId field if non-nil, zero value otherwise.

### GetNetworkEquipmentForceSubnetPoolIpv4WanIdOk

`func (o *ServerInstanceGroupConfiguration) GetNetworkEquipmentForceSubnetPoolIpv4WanIdOk() (*int32, bool)`

GetNetworkEquipmentForceSubnetPoolIpv4WanIdOk returns a tuple with the NetworkEquipmentForceSubnetPoolIpv4WanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEquipmentForceSubnetPoolIpv4WanId

`func (o *ServerInstanceGroupConfiguration) SetNetworkEquipmentForceSubnetPoolIpv4WanId(v int32)`

SetNetworkEquipmentForceSubnetPoolIpv4WanId sets NetworkEquipmentForceSubnetPoolIpv4WanId field to given value.

### HasNetworkEquipmentForceSubnetPoolIpv4WanId

`func (o *ServerInstanceGroupConfiguration) HasNetworkEquipmentForceSubnetPoolIpv4WanId() bool`

HasNetworkEquipmentForceSubnetPoolIpv4WanId returns a boolean if a field has been set.

### GetDefaultServerProfileID

`func (o *ServerInstanceGroupConfiguration) GetDefaultServerProfileID() int32`

GetDefaultServerProfileID returns the DefaultServerProfileID field if non-nil, zero value otherwise.

### GetDefaultServerProfileIDOk

`func (o *ServerInstanceGroupConfiguration) GetDefaultServerProfileIDOk() (*int32, bool)`

GetDefaultServerProfileIDOk returns a tuple with the DefaultServerProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultServerProfileID

`func (o *ServerInstanceGroupConfiguration) SetDefaultServerProfileID(v int32)`

SetDefaultServerProfileID sets DefaultServerProfileID field to given value.


### GetEmptyEdit

`func (o *ServerInstanceGroupConfiguration) GetEmptyEdit() int32`

GetEmptyEdit returns the EmptyEdit field if non-nil, zero value otherwise.

### GetEmptyEditOk

`func (o *ServerInstanceGroupConfiguration) GetEmptyEditOk() (*int32, bool)`

GetEmptyEditOk returns a tuple with the EmptyEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyEdit

`func (o *ServerInstanceGroupConfiguration) SetEmptyEdit(v int32)`

SetEmptyEdit sets EmptyEdit field to given value.

### HasEmptyEdit

`func (o *ServerInstanceGroupConfiguration) HasEmptyEdit() bool`

HasEmptyEdit returns a boolean if a field has been set.

### GetDeployType

`func (o *ServerInstanceGroupConfiguration) GetDeployType() string`

GetDeployType returns the DeployType field if non-nil, zero value otherwise.

### GetDeployTypeOk

`func (o *ServerInstanceGroupConfiguration) GetDeployTypeOk() (*string, bool)`

GetDeployTypeOk returns a tuple with the DeployType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployType

`func (o *ServerInstanceGroupConfiguration) SetDeployType(v string)`

SetDeployType sets DeployType field to given value.


### GetDeployStatus

`func (o *ServerInstanceGroupConfiguration) GetDeployStatus() string`

GetDeployStatus returns the DeployStatus field if non-nil, zero value otherwise.

### GetDeployStatusOk

`func (o *ServerInstanceGroupConfiguration) GetDeployStatusOk() (*string, bool)`

GetDeployStatusOk returns a tuple with the DeployStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStatus

`func (o *ServerInstanceGroupConfiguration) SetDeployStatus(v string)`

SetDeployStatus sets DeployStatus field to given value.


### GetDnsSubdomainChangeId

`func (o *ServerInstanceGroupConfiguration) GetDnsSubdomainChangeId() int32`

GetDnsSubdomainChangeId returns the DnsSubdomainChangeId field if non-nil, zero value otherwise.

### GetDnsSubdomainChangeIdOk

`func (o *ServerInstanceGroupConfiguration) GetDnsSubdomainChangeIdOk() (*int32, bool)`

GetDnsSubdomainChangeIdOk returns a tuple with the DnsSubdomainChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainChangeId

`func (o *ServerInstanceGroupConfiguration) SetDnsSubdomainChangeId(v int32)`

SetDnsSubdomainChangeId sets DnsSubdomainChangeId field to given value.

### HasDnsSubdomainChangeId

`func (o *ServerInstanceGroupConfiguration) HasDnsSubdomainChangeId() bool`

HasDnsSubdomainChangeId returns a boolean if a field has been set.

### GetInfrastructureDeployId

`func (o *ServerInstanceGroupConfiguration) GetInfrastructureDeployId() int32`

GetInfrastructureDeployId returns the InfrastructureDeployId field if non-nil, zero value otherwise.

### GetInfrastructureDeployIdOk

`func (o *ServerInstanceGroupConfiguration) GetInfrastructureDeployIdOk() (*int32, bool)`

GetInfrastructureDeployIdOk returns a tuple with the InfrastructureDeployId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureDeployId

`func (o *ServerInstanceGroupConfiguration) SetInfrastructureDeployId(v int32)`

SetInfrastructureDeployId sets InfrastructureDeployId field to given value.

### HasInfrastructureDeployId

`func (o *ServerInstanceGroupConfiguration) HasInfrastructureDeployId() bool`

HasInfrastructureDeployId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


