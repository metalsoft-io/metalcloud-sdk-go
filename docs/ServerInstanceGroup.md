# ServerInstanceGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The server instance group ID. | 
**Revision** | **float32** | Revision number | 
**Label** | Pointer to **string** | The server instance group label. Will be automatically generated if not provided. | [optional] 
**ServerGroupName** | Pointer to **string** |  | [optional] 
**InfrastructureId** | **int32** |  | 
**ClusterId** | **int32** |  | 
**ClusterRoleGroup** | **string** | Cluster Role Group of the Drive Group, if part of a cluster | [default to "none"]
**CreatedTimestamp** | **string** | Timestamp of the Server Instance Group creation. | 
**UpdatedTimestamp** | **string** | Timestamp of the latest update for the Server Instance Group. | 
**Meta** | Pointer to [**GenericGUISettings**](GenericGUISettings.md) | GUI settings in JSON format. | [optional] 
**Subdomain** | Pointer to **string** | Subdomain of the Server Group. | [optional] 
**SubdomainPermanent** | Pointer to **string** | Subdomain permanent of the Server Group. | [optional] 
**DnsSubdomainId** | Pointer to **int32** | Id of the DNS subdomain for the Server Group. | [optional] 
**DnsSubdomainPermanentId** | Pointer to **int32** | Id of the permanent DNS subdomain for the Server Group. | [optional] 
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
**NetworkProfileSnapshotId** | Pointer to **int32** |  | [optional] 
**OverrideIpv4WanVlanId** | Pointer to **int32** |  | [optional] 
**NetworkEquipmentForceSubnetPoolIpv4WanId** | Pointer to **int32** |  | [optional] 
**ServiceStatus** | **string** | Current status of the Server Instance Group. | 
**ResourcePoolId** | Pointer to **int32** |  | [optional] 
**IsVmGroup** | **float32** | Flag to indicate if the Server Instance Group is belongs to a VM. | 
**VmInstanceGroupId** | Pointer to **int32** | Id of the VM Instance Group this Server Instance Group belongs to. | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**DefaultServerProfileID** | **int32** | The group&#39;s default server profile. Useful when creating a server instance with a group id set, the profile will be automatically applied. | 
**Config** | Pointer to [**ServerInstanceGroupConfiguration**](ServerInstanceGroupConfiguration.md) |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewServerInstanceGroup

`func NewServerInstanceGroup(id int32, revision float32, infrastructureId int32, clusterId int32, clusterRoleGroup string, createdTimestamp string, updatedTimestamp string, instanceCount int32, ipAllocateAuto float32, ipv4SubnetCreateAuto float32, firewallManaged float32, firmwarePoliciesJson map[string]interface{}, instanceArrayBootMethod string, processorCoreCount int32, processorCoreMhz float32, processorCount int32, ramGbytes float32, diskCount int32, diskSizeMbytes int32, diskTypes []string, virtualInterfacesEnabled float32, serviceStatus string, isVmGroup float32, defaultServerProfileID int32, ) *ServerInstanceGroup`

NewServerInstanceGroup instantiates a new ServerInstanceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceGroupWithDefaults

`func NewServerInstanceGroupWithDefaults() *ServerInstanceGroup`

NewServerInstanceGroupWithDefaults instantiates a new ServerInstanceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerInstanceGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerInstanceGroup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerInstanceGroup) SetId(v int32)`

SetId sets Id field to given value.


### GetRevision

`func (o *ServerInstanceGroup) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ServerInstanceGroup) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ServerInstanceGroup) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetLabel

`func (o *ServerInstanceGroup) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServerInstanceGroup) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServerInstanceGroup) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ServerInstanceGroup) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetServerGroupName

`func (o *ServerInstanceGroup) GetServerGroupName() string`

GetServerGroupName returns the ServerGroupName field if non-nil, zero value otherwise.

### GetServerGroupNameOk

`func (o *ServerInstanceGroup) GetServerGroupNameOk() (*string, bool)`

GetServerGroupNameOk returns a tuple with the ServerGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroupName

`func (o *ServerInstanceGroup) SetServerGroupName(v string)`

SetServerGroupName sets ServerGroupName field to given value.

### HasServerGroupName

`func (o *ServerInstanceGroup) HasServerGroupName() bool`

HasServerGroupName returns a boolean if a field has been set.

### GetInfrastructureId

`func (o *ServerInstanceGroup) GetInfrastructureId() int32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *ServerInstanceGroup) GetInfrastructureIdOk() (*int32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *ServerInstanceGroup) SetInfrastructureId(v int32)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetClusterId

`func (o *ServerInstanceGroup) GetClusterId() int32`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ServerInstanceGroup) GetClusterIdOk() (*int32, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ServerInstanceGroup) SetClusterId(v int32)`

SetClusterId sets ClusterId field to given value.


### GetClusterRoleGroup

`func (o *ServerInstanceGroup) GetClusterRoleGroup() string`

GetClusterRoleGroup returns the ClusterRoleGroup field if non-nil, zero value otherwise.

### GetClusterRoleGroupOk

`func (o *ServerInstanceGroup) GetClusterRoleGroupOk() (*string, bool)`

GetClusterRoleGroupOk returns a tuple with the ClusterRoleGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterRoleGroup

`func (o *ServerInstanceGroup) SetClusterRoleGroup(v string)`

SetClusterRoleGroup sets ClusterRoleGroup field to given value.


### GetCreatedTimestamp

`func (o *ServerInstanceGroup) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *ServerInstanceGroup) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *ServerInstanceGroup) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *ServerInstanceGroup) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *ServerInstanceGroup) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *ServerInstanceGroup) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetMeta

`func (o *ServerInstanceGroup) GetMeta() GenericGUISettings`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ServerInstanceGroup) GetMetaOk() (*GenericGUISettings, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ServerInstanceGroup) SetMeta(v GenericGUISettings)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ServerInstanceGroup) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSubdomain

`func (o *ServerInstanceGroup) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *ServerInstanceGroup) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *ServerInstanceGroup) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *ServerInstanceGroup) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetSubdomainPermanent

`func (o *ServerInstanceGroup) GetSubdomainPermanent() string`

GetSubdomainPermanent returns the SubdomainPermanent field if non-nil, zero value otherwise.

### GetSubdomainPermanentOk

`func (o *ServerInstanceGroup) GetSubdomainPermanentOk() (*string, bool)`

GetSubdomainPermanentOk returns a tuple with the SubdomainPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomainPermanent

`func (o *ServerInstanceGroup) SetSubdomainPermanent(v string)`

SetSubdomainPermanent sets SubdomainPermanent field to given value.

### HasSubdomainPermanent

`func (o *ServerInstanceGroup) HasSubdomainPermanent() bool`

HasSubdomainPermanent returns a boolean if a field has been set.

### GetDnsSubdomainId

`func (o *ServerInstanceGroup) GetDnsSubdomainId() int32`

GetDnsSubdomainId returns the DnsSubdomainId field if non-nil, zero value otherwise.

### GetDnsSubdomainIdOk

`func (o *ServerInstanceGroup) GetDnsSubdomainIdOk() (*int32, bool)`

GetDnsSubdomainIdOk returns a tuple with the DnsSubdomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainId

`func (o *ServerInstanceGroup) SetDnsSubdomainId(v int32)`

SetDnsSubdomainId sets DnsSubdomainId field to given value.

### HasDnsSubdomainId

`func (o *ServerInstanceGroup) HasDnsSubdomainId() bool`

HasDnsSubdomainId returns a boolean if a field has been set.

### GetDnsSubdomainPermanentId

`func (o *ServerInstanceGroup) GetDnsSubdomainPermanentId() int32`

GetDnsSubdomainPermanentId returns the DnsSubdomainPermanentId field if non-nil, zero value otherwise.

### GetDnsSubdomainPermanentIdOk

`func (o *ServerInstanceGroup) GetDnsSubdomainPermanentIdOk() (*int32, bool)`

GetDnsSubdomainPermanentIdOk returns a tuple with the DnsSubdomainPermanentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainPermanentId

`func (o *ServerInstanceGroup) SetDnsSubdomainPermanentId(v int32)`

SetDnsSubdomainPermanentId sets DnsSubdomainPermanentId field to given value.

### HasDnsSubdomainPermanentId

`func (o *ServerInstanceGroup) HasDnsSubdomainPermanentId() bool`

HasDnsSubdomainPermanentId returns a boolean if a field has been set.

### GetInstanceCount

`func (o *ServerInstanceGroup) GetInstanceCount() int32`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *ServerInstanceGroup) GetInstanceCountOk() (*int32, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *ServerInstanceGroup) SetInstanceCount(v int32)`

SetInstanceCount sets InstanceCount field to given value.


### GetIpAllocateAuto

`func (o *ServerInstanceGroup) GetIpAllocateAuto() float32`

GetIpAllocateAuto returns the IpAllocateAuto field if non-nil, zero value otherwise.

### GetIpAllocateAutoOk

`func (o *ServerInstanceGroup) GetIpAllocateAutoOk() (*float32, bool)`

GetIpAllocateAutoOk returns a tuple with the IpAllocateAuto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAllocateAuto

`func (o *ServerInstanceGroup) SetIpAllocateAuto(v float32)`

SetIpAllocateAuto sets IpAllocateAuto field to given value.


### GetIpv4SubnetCreateAuto

`func (o *ServerInstanceGroup) GetIpv4SubnetCreateAuto() float32`

GetIpv4SubnetCreateAuto returns the Ipv4SubnetCreateAuto field if non-nil, zero value otherwise.

### GetIpv4SubnetCreateAutoOk

`func (o *ServerInstanceGroup) GetIpv4SubnetCreateAutoOk() (*float32, bool)`

GetIpv4SubnetCreateAutoOk returns a tuple with the Ipv4SubnetCreateAuto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4SubnetCreateAuto

`func (o *ServerInstanceGroup) SetIpv4SubnetCreateAuto(v float32)`

SetIpv4SubnetCreateAuto sets Ipv4SubnetCreateAuto field to given value.


### GetFirewallProfileId

`func (o *ServerInstanceGroup) GetFirewallProfileId() int32`

GetFirewallProfileId returns the FirewallProfileId field if non-nil, zero value otherwise.

### GetFirewallProfileIdOk

`func (o *ServerInstanceGroup) GetFirewallProfileIdOk() (*int32, bool)`

GetFirewallProfileIdOk returns a tuple with the FirewallProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallProfileId

`func (o *ServerInstanceGroup) SetFirewallProfileId(v int32)`

SetFirewallProfileId sets FirewallProfileId field to given value.

### HasFirewallProfileId

`func (o *ServerInstanceGroup) HasFirewallProfileId() bool`

HasFirewallProfileId returns a boolean if a field has been set.

### GetFirewallRulesSetId

`func (o *ServerInstanceGroup) GetFirewallRulesSetId() int32`

GetFirewallRulesSetId returns the FirewallRulesSetId field if non-nil, zero value otherwise.

### GetFirewallRulesSetIdOk

`func (o *ServerInstanceGroup) GetFirewallRulesSetIdOk() (*int32, bool)`

GetFirewallRulesSetIdOk returns a tuple with the FirewallRulesSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallRulesSetId

`func (o *ServerInstanceGroup) SetFirewallRulesSetId(v int32)`

SetFirewallRulesSetId sets FirewallRulesSetId field to given value.

### HasFirewallRulesSetId

`func (o *ServerInstanceGroup) HasFirewallRulesSetId() bool`

HasFirewallRulesSetId returns a boolean if a field has been set.

### GetFirewallManaged

`func (o *ServerInstanceGroup) GetFirewallManaged() float32`

GetFirewallManaged returns the FirewallManaged field if non-nil, zero value otherwise.

### GetFirewallManagedOk

`func (o *ServerInstanceGroup) GetFirewallManagedOk() (*float32, bool)`

GetFirewallManagedOk returns a tuple with the FirewallManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallManaged

`func (o *ServerInstanceGroup) SetFirewallManaged(v float32)`

SetFirewallManaged sets FirewallManaged field to given value.


### GetFirmwarePoliciesJson

`func (o *ServerInstanceGroup) GetFirmwarePoliciesJson() map[string]interface{}`

GetFirmwarePoliciesJson returns the FirmwarePoliciesJson field if non-nil, zero value otherwise.

### GetFirmwarePoliciesJsonOk

`func (o *ServerInstanceGroup) GetFirmwarePoliciesJsonOk() (*map[string]interface{}, bool)`

GetFirmwarePoliciesJsonOk returns a tuple with the FirmwarePoliciesJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwarePoliciesJson

`func (o *ServerInstanceGroup) SetFirmwarePoliciesJson(v map[string]interface{})`

SetFirmwarePoliciesJson sets FirmwarePoliciesJson field to given value.


### GetVolumeTemplateId

`func (o *ServerInstanceGroup) GetVolumeTemplateId() int32`

GetVolumeTemplateId returns the VolumeTemplateId field if non-nil, zero value otherwise.

### GetVolumeTemplateIdOk

`func (o *ServerInstanceGroup) GetVolumeTemplateIdOk() (*int32, bool)`

GetVolumeTemplateIdOk returns a tuple with the VolumeTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeTemplateId

`func (o *ServerInstanceGroup) SetVolumeTemplateId(v int32)`

SetVolumeTemplateId sets VolumeTemplateId field to given value.

### HasVolumeTemplateId

`func (o *ServerInstanceGroup) HasVolumeTemplateId() bool`

HasVolumeTemplateId returns a boolean if a field has been set.

### GetDriveArrayIdBoot

`func (o *ServerInstanceGroup) GetDriveArrayIdBoot() int32`

GetDriveArrayIdBoot returns the DriveArrayIdBoot field if non-nil, zero value otherwise.

### GetDriveArrayIdBootOk

`func (o *ServerInstanceGroup) GetDriveArrayIdBootOk() (*int32, bool)`

GetDriveArrayIdBootOk returns a tuple with the DriveArrayIdBoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveArrayIdBoot

`func (o *ServerInstanceGroup) SetDriveArrayIdBoot(v int32)`

SetDriveArrayIdBoot sets DriveArrayIdBoot field to given value.

### HasDriveArrayIdBoot

`func (o *ServerInstanceGroup) HasDriveArrayIdBoot() bool`

HasDriveArrayIdBoot returns a boolean if a field has been set.

### GetInstanceArrayBootMethod

`func (o *ServerInstanceGroup) GetInstanceArrayBootMethod() string`

GetInstanceArrayBootMethod returns the InstanceArrayBootMethod field if non-nil, zero value otherwise.

### GetInstanceArrayBootMethodOk

`func (o *ServerInstanceGroup) GetInstanceArrayBootMethodOk() (*string, bool)`

GetInstanceArrayBootMethodOk returns a tuple with the InstanceArrayBootMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceArrayBootMethod

`func (o *ServerInstanceGroup) SetInstanceArrayBootMethod(v string)`

SetInstanceArrayBootMethod sets InstanceArrayBootMethod field to given value.


### GetCustomVariables

`func (o *ServerInstanceGroup) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *ServerInstanceGroup) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *ServerInstanceGroup) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *ServerInstanceGroup) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.

### GetProcessorCoreCount

`func (o *ServerInstanceGroup) GetProcessorCoreCount() int32`

GetProcessorCoreCount returns the ProcessorCoreCount field if non-nil, zero value otherwise.

### GetProcessorCoreCountOk

`func (o *ServerInstanceGroup) GetProcessorCoreCountOk() (*int32, bool)`

GetProcessorCoreCountOk returns a tuple with the ProcessorCoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCoreCount

`func (o *ServerInstanceGroup) SetProcessorCoreCount(v int32)`

SetProcessorCoreCount sets ProcessorCoreCount field to given value.


### GetProcessorCoreMhz

`func (o *ServerInstanceGroup) GetProcessorCoreMhz() float32`

GetProcessorCoreMhz returns the ProcessorCoreMhz field if non-nil, zero value otherwise.

### GetProcessorCoreMhzOk

`func (o *ServerInstanceGroup) GetProcessorCoreMhzOk() (*float32, bool)`

GetProcessorCoreMhzOk returns a tuple with the ProcessorCoreMhz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCoreMhz

`func (o *ServerInstanceGroup) SetProcessorCoreMhz(v float32)`

SetProcessorCoreMhz sets ProcessorCoreMhz field to given value.


### GetProcessorCount

`func (o *ServerInstanceGroup) GetProcessorCount() int32`

GetProcessorCount returns the ProcessorCount field if non-nil, zero value otherwise.

### GetProcessorCountOk

`func (o *ServerInstanceGroup) GetProcessorCountOk() (*int32, bool)`

GetProcessorCountOk returns a tuple with the ProcessorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCount

`func (o *ServerInstanceGroup) SetProcessorCount(v int32)`

SetProcessorCount sets ProcessorCount field to given value.


### GetRamGbytes

`func (o *ServerInstanceGroup) GetRamGbytes() float32`

GetRamGbytes returns the RamGbytes field if non-nil, zero value otherwise.

### GetRamGbytesOk

`func (o *ServerInstanceGroup) GetRamGbytesOk() (*float32, bool)`

GetRamGbytesOk returns a tuple with the RamGbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGbytes

`func (o *ServerInstanceGroup) SetRamGbytes(v float32)`

SetRamGbytes sets RamGbytes field to given value.


### GetDiskCount

`func (o *ServerInstanceGroup) GetDiskCount() int32`

GetDiskCount returns the DiskCount field if non-nil, zero value otherwise.

### GetDiskCountOk

`func (o *ServerInstanceGroup) GetDiskCountOk() (*int32, bool)`

GetDiskCountOk returns a tuple with the DiskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskCount

`func (o *ServerInstanceGroup) SetDiskCount(v int32)`

SetDiskCount sets DiskCount field to given value.


### GetDiskSizeMbytes

`func (o *ServerInstanceGroup) GetDiskSizeMbytes() int32`

GetDiskSizeMbytes returns the DiskSizeMbytes field if non-nil, zero value otherwise.

### GetDiskSizeMbytesOk

`func (o *ServerInstanceGroup) GetDiskSizeMbytesOk() (*int32, bool)`

GetDiskSizeMbytesOk returns a tuple with the DiskSizeMbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeMbytes

`func (o *ServerInstanceGroup) SetDiskSizeMbytes(v int32)`

SetDiskSizeMbytes sets DiskSizeMbytes field to given value.


### GetDiskTypes

`func (o *ServerInstanceGroup) GetDiskTypes() []string`

GetDiskTypes returns the DiskTypes field if non-nil, zero value otherwise.

### GetDiskTypesOk

`func (o *ServerInstanceGroup) GetDiskTypesOk() (*[]string, bool)`

GetDiskTypesOk returns a tuple with the DiskTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskTypes

`func (o *ServerInstanceGroup) SetDiskTypes(v []string)`

SetDiskTypes sets DiskTypes field to given value.


### GetVirtualInterfacesEnabled

`func (o *ServerInstanceGroup) GetVirtualInterfacesEnabled() float32`

GetVirtualInterfacesEnabled returns the VirtualInterfacesEnabled field if non-nil, zero value otherwise.

### GetVirtualInterfacesEnabledOk

`func (o *ServerInstanceGroup) GetVirtualInterfacesEnabledOk() (*float32, bool)`

GetVirtualInterfacesEnabledOk returns a tuple with the VirtualInterfacesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualInterfacesEnabled

`func (o *ServerInstanceGroup) SetVirtualInterfacesEnabled(v float32)`

SetVirtualInterfacesEnabled sets VirtualInterfacesEnabled field to given value.


### GetAdditionalWanIpv4Json

`func (o *ServerInstanceGroup) GetAdditionalWanIpv4Json() map[string]interface{}`

GetAdditionalWanIpv4Json returns the AdditionalWanIpv4Json field if non-nil, zero value otherwise.

### GetAdditionalWanIpv4JsonOk

`func (o *ServerInstanceGroup) GetAdditionalWanIpv4JsonOk() (*map[string]interface{}, bool)`

GetAdditionalWanIpv4JsonOk returns a tuple with the AdditionalWanIpv4Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalWanIpv4Json

`func (o *ServerInstanceGroup) SetAdditionalWanIpv4Json(v map[string]interface{})`

SetAdditionalWanIpv4Json sets AdditionalWanIpv4Json field to given value.

### HasAdditionalWanIpv4Json

`func (o *ServerInstanceGroup) HasAdditionalWanIpv4Json() bool`

HasAdditionalWanIpv4Json returns a boolean if a field has been set.

### GetNetworkProfileGroupId

`func (o *ServerInstanceGroup) GetNetworkProfileGroupId() int32`

GetNetworkProfileGroupId returns the NetworkProfileGroupId field if non-nil, zero value otherwise.

### GetNetworkProfileGroupIdOk

`func (o *ServerInstanceGroup) GetNetworkProfileGroupIdOk() (*int32, bool)`

GetNetworkProfileGroupIdOk returns a tuple with the NetworkProfileGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProfileGroupId

`func (o *ServerInstanceGroup) SetNetworkProfileGroupId(v int32)`

SetNetworkProfileGroupId sets NetworkProfileGroupId field to given value.

### HasNetworkProfileGroupId

`func (o *ServerInstanceGroup) HasNetworkProfileGroupId() bool`

HasNetworkProfileGroupId returns a boolean if a field has been set.

### GetNetworkProfileSnapshotId

`func (o *ServerInstanceGroup) GetNetworkProfileSnapshotId() int32`

GetNetworkProfileSnapshotId returns the NetworkProfileSnapshotId field if non-nil, zero value otherwise.

### GetNetworkProfileSnapshotIdOk

`func (o *ServerInstanceGroup) GetNetworkProfileSnapshotIdOk() (*int32, bool)`

GetNetworkProfileSnapshotIdOk returns a tuple with the NetworkProfileSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProfileSnapshotId

`func (o *ServerInstanceGroup) SetNetworkProfileSnapshotId(v int32)`

SetNetworkProfileSnapshotId sets NetworkProfileSnapshotId field to given value.

### HasNetworkProfileSnapshotId

`func (o *ServerInstanceGroup) HasNetworkProfileSnapshotId() bool`

HasNetworkProfileSnapshotId returns a boolean if a field has been set.

### GetOverrideIpv4WanVlanId

`func (o *ServerInstanceGroup) GetOverrideIpv4WanVlanId() int32`

GetOverrideIpv4WanVlanId returns the OverrideIpv4WanVlanId field if non-nil, zero value otherwise.

### GetOverrideIpv4WanVlanIdOk

`func (o *ServerInstanceGroup) GetOverrideIpv4WanVlanIdOk() (*int32, bool)`

GetOverrideIpv4WanVlanIdOk returns a tuple with the OverrideIpv4WanVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideIpv4WanVlanId

`func (o *ServerInstanceGroup) SetOverrideIpv4WanVlanId(v int32)`

SetOverrideIpv4WanVlanId sets OverrideIpv4WanVlanId field to given value.

### HasOverrideIpv4WanVlanId

`func (o *ServerInstanceGroup) HasOverrideIpv4WanVlanId() bool`

HasOverrideIpv4WanVlanId returns a boolean if a field has been set.

### GetNetworkEquipmentForceSubnetPoolIpv4WanId

`func (o *ServerInstanceGroup) GetNetworkEquipmentForceSubnetPoolIpv4WanId() int32`

GetNetworkEquipmentForceSubnetPoolIpv4WanId returns the NetworkEquipmentForceSubnetPoolIpv4WanId field if non-nil, zero value otherwise.

### GetNetworkEquipmentForceSubnetPoolIpv4WanIdOk

`func (o *ServerInstanceGroup) GetNetworkEquipmentForceSubnetPoolIpv4WanIdOk() (*int32, bool)`

GetNetworkEquipmentForceSubnetPoolIpv4WanIdOk returns a tuple with the NetworkEquipmentForceSubnetPoolIpv4WanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEquipmentForceSubnetPoolIpv4WanId

`func (o *ServerInstanceGroup) SetNetworkEquipmentForceSubnetPoolIpv4WanId(v int32)`

SetNetworkEquipmentForceSubnetPoolIpv4WanId sets NetworkEquipmentForceSubnetPoolIpv4WanId field to given value.

### HasNetworkEquipmentForceSubnetPoolIpv4WanId

`func (o *ServerInstanceGroup) HasNetworkEquipmentForceSubnetPoolIpv4WanId() bool`

HasNetworkEquipmentForceSubnetPoolIpv4WanId returns a boolean if a field has been set.

### GetServiceStatus

`func (o *ServerInstanceGroup) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *ServerInstanceGroup) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *ServerInstanceGroup) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetResourcePoolId

`func (o *ServerInstanceGroup) GetResourcePoolId() int32`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *ServerInstanceGroup) GetResourcePoolIdOk() (*int32, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *ServerInstanceGroup) SetResourcePoolId(v int32)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *ServerInstanceGroup) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetIsVmGroup

`func (o *ServerInstanceGroup) GetIsVmGroup() float32`

GetIsVmGroup returns the IsVmGroup field if non-nil, zero value otherwise.

### GetIsVmGroupOk

`func (o *ServerInstanceGroup) GetIsVmGroupOk() (*float32, bool)`

GetIsVmGroupOk returns a tuple with the IsVmGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVmGroup

`func (o *ServerInstanceGroup) SetIsVmGroup(v float32)`

SetIsVmGroup sets IsVmGroup field to given value.


### GetVmInstanceGroupId

`func (o *ServerInstanceGroup) GetVmInstanceGroupId() int32`

GetVmInstanceGroupId returns the VmInstanceGroupId field if non-nil, zero value otherwise.

### GetVmInstanceGroupIdOk

`func (o *ServerInstanceGroup) GetVmInstanceGroupIdOk() (*int32, bool)`

GetVmInstanceGroupIdOk returns a tuple with the VmInstanceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmInstanceGroupId

`func (o *ServerInstanceGroup) SetVmInstanceGroupId(v int32)`

SetVmInstanceGroupId sets VmInstanceGroupId field to given value.

### HasVmInstanceGroupId

`func (o *ServerInstanceGroup) HasVmInstanceGroupId() bool`

HasVmInstanceGroupId returns a boolean if a field has been set.

### GetTags

`func (o *ServerInstanceGroup) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServerInstanceGroup) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServerInstanceGroup) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServerInstanceGroup) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDefaultServerProfileID

`func (o *ServerInstanceGroup) GetDefaultServerProfileID() int32`

GetDefaultServerProfileID returns the DefaultServerProfileID field if non-nil, zero value otherwise.

### GetDefaultServerProfileIDOk

`func (o *ServerInstanceGroup) GetDefaultServerProfileIDOk() (*int32, bool)`

GetDefaultServerProfileIDOk returns a tuple with the DefaultServerProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultServerProfileID

`func (o *ServerInstanceGroup) SetDefaultServerProfileID(v int32)`

SetDefaultServerProfileID sets DefaultServerProfileID field to given value.


### GetConfig

`func (o *ServerInstanceGroup) GetConfig() ServerInstanceGroupConfiguration`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ServerInstanceGroup) GetConfigOk() (*ServerInstanceGroupConfiguration, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ServerInstanceGroup) SetConfig(v ServerInstanceGroupConfiguration)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ServerInstanceGroup) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetLinks

`func (o *ServerInstanceGroup) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServerInstanceGroup) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServerInstanceGroup) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServerInstanceGroup) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


