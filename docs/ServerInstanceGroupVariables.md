# ServerInstanceGroupVariables

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The Product Instance ID. | 
**Revision** | **int32** | Revision number | 
**Label** | **string** | The Product Instance label. Will be automatically generated if not provided. | 
**CreatedTimestamp** | **string** | Timestamp of the Product Instance creation. | 
**UpdatedTimestamp** | **string** | Timestamp of the latest update of the Product Instance. | 
**Subdomain** | Pointer to **string** | Subdomain of the Product Instance. | [optional] 
**SubdomainPermanent** | Pointer to **string** | Subdomain permanent of the Product Instance. | [optional] 
**DnsSubdomainId** | Pointer to **int32** | Id of the DNS subdomain for the Product Instance | [optional] 
**DnsSubdomainPermanentId** | Pointer to **int32** | Id of the permanent DNS subdomain for the Product Instance | [optional] 
**ServerGroupName** | Pointer to **string** |  | [optional] 
**InfrastructureId** | **int32** |  | 
**ExtensionInstanceId** | Pointer to **int32** |  | [optional] 
**InstanceCount** | **int32** | The number of instances to be created on the Instance Group. | [default to 1]
**DefaultServerTypeId** | **int32** | The server type ID that will be assigned to newly created instances. | 
**DefaultCustomStorageProfile** | Pointer to [**ServerInstanceStorageProfile**](ServerInstanceStorageProfile.md) | Default Custom Storage Profile for the newly created Instances. | [optional] 
**IpAllocateAuto** | **int32** | Automatically allocate IP addresses to child Instance&#x60;s Instance Interface elements. | [default to 1]
**Ipv4SubnetCreateAuto** | **int32** | Automatically create or expand Subnet elements until the necessary IPv4 addresses are allocated. | [default to 1]
**FirmwarePolicyIds** | Pointer to **[]float32** | Array of firmware policy ids containing associated firmware policies. | [optional] 
**Hostname** | Pointer to **string** | Custom hostname for the DNS Load Balancing record. If set, this will be used as the DNS Load Balancing record name instead of the default \&quot;instance-group\&quot;. The hostname must be a valid DNS subdomain and can only contain alphanumeric characters, hyphens, and underscores. This will only take effect if the property \&quot;provisionLoadBalancingDnsRecord\&quot; is true. It will be automatically suffixed with the server instance group ID (e.g., \&quot;-34\&quot;) to ensure the uniqueness of the resulting DNS name. | [optional] 
**OsTemplateId** | Pointer to **int32** | The volume template ID (or name) to use if the servers in the Instance Group have local disks. | [optional] 
**CustomVariables** | Pointer to **map[string]interface{}** | Object containing custom variables and variable overrides. | [optional] 
**ProcessorCount** | **int32** | The CPU count on each instance. | [default to 1]
**ProcessorCoreCount** | **int32** | The minimum cores of a CPU. | [default to 1]
**ProcessorCoreMhz** | **int32** | The minimum clock speed of a CPU. | [default to 1000]
**RamGbytes** | Pointer to **int32** | The minimum RAM capacity of each instance. | [optional] [default to 1]
**DiskCount** | **int32** | The minimum number of physical disks. | [default to 0]
**DiskSizeMbytes** | **int32** | The minimum size of a single disk. | [default to 0]
**DiskTypes** | **[]string** | The types of physical disks. | [default to []]
**VirtualInterfacesEnabled** | **int32** | Enable virtual interfaces | [default to 0]
**AdditionalWanIpv4Json** | Pointer to **map[string]interface{}** | Contains info about additional ips to be assigned to the WAN interfaces. | [optional] 
**NetworkProfileGroupId** | Pointer to **int32** |  | [optional] 
**NetworkProfileSnapshotId** | Pointer to **int32** |  | [optional] 
**OverrideIpv4WanVlanId** | Pointer to **int32** | The ipv4 vlan that should override the default from the WAN Network for the primary ip. | [optional] 
**NetworkEquipmentForceSubnetPoolIpv4WanId** | Pointer to **int32** | ID of a ipv4 WAN subnet-pool from which to force the subnet allocation for the Instance Interfaces associated with this Instance Group. | [optional] 
**ServiceStatus** | **string** | Current status of the Server Instance Group. | 
**ResourcePoolId** | Pointer to **int32** | The resource pool assigned to this instance array | [optional] 
**IsVmGroup** | **int32** | Flag to indicate if the Server Instance Group is belongs to a VM. | 
**IsEndpointInstanceGroup** | **int32** | Flag to indicate if the Server Instance Group is belongs to a Endpoint. | 
**VmInstanceGroupId** | Pointer to **int32** | Id of the VM Instance Group this Server Instance Group belongs to. | [optional] 
**NetworkEndpointGroupId** | Pointer to **int32** |  | [optional] 
**Config** | Pointer to [**ServerInstanceGroupConfiguration**](ServerInstanceGroupConfiguration.md) |  | [optional] 

## Methods

### NewServerInstanceGroupVariables

`func NewServerInstanceGroupVariables(id int32, revision int32, label string, createdTimestamp string, updatedTimestamp string, infrastructureId int32, instanceCount int32, defaultServerTypeId int32, ipAllocateAuto int32, ipv4SubnetCreateAuto int32, processorCount int32, processorCoreCount int32, processorCoreMhz int32, diskCount int32, diskSizeMbytes int32, diskTypes []string, virtualInterfacesEnabled int32, serviceStatus string, isVmGroup int32, isEndpointInstanceGroup int32, ) *ServerInstanceGroupVariables`

NewServerInstanceGroupVariables instantiates a new ServerInstanceGroupVariables object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceGroupVariablesWithDefaults

`func NewServerInstanceGroupVariablesWithDefaults() *ServerInstanceGroupVariables`

NewServerInstanceGroupVariablesWithDefaults instantiates a new ServerInstanceGroupVariables object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerInstanceGroupVariables) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerInstanceGroupVariables) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerInstanceGroupVariables) SetId(v int32)`

SetId sets Id field to given value.


### GetRevision

`func (o *ServerInstanceGroupVariables) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ServerInstanceGroupVariables) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ServerInstanceGroupVariables) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetLabel

`func (o *ServerInstanceGroupVariables) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServerInstanceGroupVariables) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServerInstanceGroupVariables) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetCreatedTimestamp

`func (o *ServerInstanceGroupVariables) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *ServerInstanceGroupVariables) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *ServerInstanceGroupVariables) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *ServerInstanceGroupVariables) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *ServerInstanceGroupVariables) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *ServerInstanceGroupVariables) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetSubdomain

`func (o *ServerInstanceGroupVariables) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *ServerInstanceGroupVariables) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *ServerInstanceGroupVariables) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *ServerInstanceGroupVariables) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetSubdomainPermanent

`func (o *ServerInstanceGroupVariables) GetSubdomainPermanent() string`

GetSubdomainPermanent returns the SubdomainPermanent field if non-nil, zero value otherwise.

### GetSubdomainPermanentOk

`func (o *ServerInstanceGroupVariables) GetSubdomainPermanentOk() (*string, bool)`

GetSubdomainPermanentOk returns a tuple with the SubdomainPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomainPermanent

`func (o *ServerInstanceGroupVariables) SetSubdomainPermanent(v string)`

SetSubdomainPermanent sets SubdomainPermanent field to given value.

### HasSubdomainPermanent

`func (o *ServerInstanceGroupVariables) HasSubdomainPermanent() bool`

HasSubdomainPermanent returns a boolean if a field has been set.

### GetDnsSubdomainId

`func (o *ServerInstanceGroupVariables) GetDnsSubdomainId() int32`

GetDnsSubdomainId returns the DnsSubdomainId field if non-nil, zero value otherwise.

### GetDnsSubdomainIdOk

`func (o *ServerInstanceGroupVariables) GetDnsSubdomainIdOk() (*int32, bool)`

GetDnsSubdomainIdOk returns a tuple with the DnsSubdomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainId

`func (o *ServerInstanceGroupVariables) SetDnsSubdomainId(v int32)`

SetDnsSubdomainId sets DnsSubdomainId field to given value.

### HasDnsSubdomainId

`func (o *ServerInstanceGroupVariables) HasDnsSubdomainId() bool`

HasDnsSubdomainId returns a boolean if a field has been set.

### GetDnsSubdomainPermanentId

`func (o *ServerInstanceGroupVariables) GetDnsSubdomainPermanentId() int32`

GetDnsSubdomainPermanentId returns the DnsSubdomainPermanentId field if non-nil, zero value otherwise.

### GetDnsSubdomainPermanentIdOk

`func (o *ServerInstanceGroupVariables) GetDnsSubdomainPermanentIdOk() (*int32, bool)`

GetDnsSubdomainPermanentIdOk returns a tuple with the DnsSubdomainPermanentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainPermanentId

`func (o *ServerInstanceGroupVariables) SetDnsSubdomainPermanentId(v int32)`

SetDnsSubdomainPermanentId sets DnsSubdomainPermanentId field to given value.

### HasDnsSubdomainPermanentId

`func (o *ServerInstanceGroupVariables) HasDnsSubdomainPermanentId() bool`

HasDnsSubdomainPermanentId returns a boolean if a field has been set.

### GetServerGroupName

`func (o *ServerInstanceGroupVariables) GetServerGroupName() string`

GetServerGroupName returns the ServerGroupName field if non-nil, zero value otherwise.

### GetServerGroupNameOk

`func (o *ServerInstanceGroupVariables) GetServerGroupNameOk() (*string, bool)`

GetServerGroupNameOk returns a tuple with the ServerGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroupName

`func (o *ServerInstanceGroupVariables) SetServerGroupName(v string)`

SetServerGroupName sets ServerGroupName field to given value.

### HasServerGroupName

`func (o *ServerInstanceGroupVariables) HasServerGroupName() bool`

HasServerGroupName returns a boolean if a field has been set.

### GetInfrastructureId

`func (o *ServerInstanceGroupVariables) GetInfrastructureId() int32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *ServerInstanceGroupVariables) GetInfrastructureIdOk() (*int32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *ServerInstanceGroupVariables) SetInfrastructureId(v int32)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetExtensionInstanceId

`func (o *ServerInstanceGroupVariables) GetExtensionInstanceId() int32`

GetExtensionInstanceId returns the ExtensionInstanceId field if non-nil, zero value otherwise.

### GetExtensionInstanceIdOk

`func (o *ServerInstanceGroupVariables) GetExtensionInstanceIdOk() (*int32, bool)`

GetExtensionInstanceIdOk returns a tuple with the ExtensionInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionInstanceId

`func (o *ServerInstanceGroupVariables) SetExtensionInstanceId(v int32)`

SetExtensionInstanceId sets ExtensionInstanceId field to given value.

### HasExtensionInstanceId

`func (o *ServerInstanceGroupVariables) HasExtensionInstanceId() bool`

HasExtensionInstanceId returns a boolean if a field has been set.

### GetInstanceCount

`func (o *ServerInstanceGroupVariables) GetInstanceCount() int32`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *ServerInstanceGroupVariables) GetInstanceCountOk() (*int32, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *ServerInstanceGroupVariables) SetInstanceCount(v int32)`

SetInstanceCount sets InstanceCount field to given value.


### GetDefaultServerTypeId

`func (o *ServerInstanceGroupVariables) GetDefaultServerTypeId() int32`

GetDefaultServerTypeId returns the DefaultServerTypeId field if non-nil, zero value otherwise.

### GetDefaultServerTypeIdOk

`func (o *ServerInstanceGroupVariables) GetDefaultServerTypeIdOk() (*int32, bool)`

GetDefaultServerTypeIdOk returns a tuple with the DefaultServerTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultServerTypeId

`func (o *ServerInstanceGroupVariables) SetDefaultServerTypeId(v int32)`

SetDefaultServerTypeId sets DefaultServerTypeId field to given value.


### GetDefaultCustomStorageProfile

`func (o *ServerInstanceGroupVariables) GetDefaultCustomStorageProfile() ServerInstanceStorageProfile`

GetDefaultCustomStorageProfile returns the DefaultCustomStorageProfile field if non-nil, zero value otherwise.

### GetDefaultCustomStorageProfileOk

`func (o *ServerInstanceGroupVariables) GetDefaultCustomStorageProfileOk() (*ServerInstanceStorageProfile, bool)`

GetDefaultCustomStorageProfileOk returns a tuple with the DefaultCustomStorageProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCustomStorageProfile

`func (o *ServerInstanceGroupVariables) SetDefaultCustomStorageProfile(v ServerInstanceStorageProfile)`

SetDefaultCustomStorageProfile sets DefaultCustomStorageProfile field to given value.

### HasDefaultCustomStorageProfile

`func (o *ServerInstanceGroupVariables) HasDefaultCustomStorageProfile() bool`

HasDefaultCustomStorageProfile returns a boolean if a field has been set.

### GetIpAllocateAuto

`func (o *ServerInstanceGroupVariables) GetIpAllocateAuto() int32`

GetIpAllocateAuto returns the IpAllocateAuto field if non-nil, zero value otherwise.

### GetIpAllocateAutoOk

`func (o *ServerInstanceGroupVariables) GetIpAllocateAutoOk() (*int32, bool)`

GetIpAllocateAutoOk returns a tuple with the IpAllocateAuto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAllocateAuto

`func (o *ServerInstanceGroupVariables) SetIpAllocateAuto(v int32)`

SetIpAllocateAuto sets IpAllocateAuto field to given value.


### GetIpv4SubnetCreateAuto

`func (o *ServerInstanceGroupVariables) GetIpv4SubnetCreateAuto() int32`

GetIpv4SubnetCreateAuto returns the Ipv4SubnetCreateAuto field if non-nil, zero value otherwise.

### GetIpv4SubnetCreateAutoOk

`func (o *ServerInstanceGroupVariables) GetIpv4SubnetCreateAutoOk() (*int32, bool)`

GetIpv4SubnetCreateAutoOk returns a tuple with the Ipv4SubnetCreateAuto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4SubnetCreateAuto

`func (o *ServerInstanceGroupVariables) SetIpv4SubnetCreateAuto(v int32)`

SetIpv4SubnetCreateAuto sets Ipv4SubnetCreateAuto field to given value.


### GetFirmwarePolicyIds

`func (o *ServerInstanceGroupVariables) GetFirmwarePolicyIds() []float32`

GetFirmwarePolicyIds returns the FirmwarePolicyIds field if non-nil, zero value otherwise.

### GetFirmwarePolicyIdsOk

`func (o *ServerInstanceGroupVariables) GetFirmwarePolicyIdsOk() (*[]float32, bool)`

GetFirmwarePolicyIdsOk returns a tuple with the FirmwarePolicyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwarePolicyIds

`func (o *ServerInstanceGroupVariables) SetFirmwarePolicyIds(v []float32)`

SetFirmwarePolicyIds sets FirmwarePolicyIds field to given value.

### HasFirmwarePolicyIds

`func (o *ServerInstanceGroupVariables) HasFirmwarePolicyIds() bool`

HasFirmwarePolicyIds returns a boolean if a field has been set.

### SetFirmwarePolicyIdsNil

`func (o *ServerInstanceGroupVariables) SetFirmwarePolicyIdsNil(b bool)`

 SetFirmwarePolicyIdsNil sets the value for FirmwarePolicyIds to be an explicit nil

### UnsetFirmwarePolicyIds
`func (o *ServerInstanceGroupVariables) UnsetFirmwarePolicyIds()`

UnsetFirmwarePolicyIds ensures that no value is present for FirmwarePolicyIds, not even an explicit nil
### GetHostname

`func (o *ServerInstanceGroupVariables) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ServerInstanceGroupVariables) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ServerInstanceGroupVariables) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ServerInstanceGroupVariables) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetOsTemplateId

`func (o *ServerInstanceGroupVariables) GetOsTemplateId() int32`

GetOsTemplateId returns the OsTemplateId field if non-nil, zero value otherwise.

### GetOsTemplateIdOk

`func (o *ServerInstanceGroupVariables) GetOsTemplateIdOk() (*int32, bool)`

GetOsTemplateIdOk returns a tuple with the OsTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsTemplateId

`func (o *ServerInstanceGroupVariables) SetOsTemplateId(v int32)`

SetOsTemplateId sets OsTemplateId field to given value.

### HasOsTemplateId

`func (o *ServerInstanceGroupVariables) HasOsTemplateId() bool`

HasOsTemplateId returns a boolean if a field has been set.

### GetCustomVariables

`func (o *ServerInstanceGroupVariables) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *ServerInstanceGroupVariables) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *ServerInstanceGroupVariables) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *ServerInstanceGroupVariables) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.

### GetProcessorCount

`func (o *ServerInstanceGroupVariables) GetProcessorCount() int32`

GetProcessorCount returns the ProcessorCount field if non-nil, zero value otherwise.

### GetProcessorCountOk

`func (o *ServerInstanceGroupVariables) GetProcessorCountOk() (*int32, bool)`

GetProcessorCountOk returns a tuple with the ProcessorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCount

`func (o *ServerInstanceGroupVariables) SetProcessorCount(v int32)`

SetProcessorCount sets ProcessorCount field to given value.


### GetProcessorCoreCount

`func (o *ServerInstanceGroupVariables) GetProcessorCoreCount() int32`

GetProcessorCoreCount returns the ProcessorCoreCount field if non-nil, zero value otherwise.

### GetProcessorCoreCountOk

`func (o *ServerInstanceGroupVariables) GetProcessorCoreCountOk() (*int32, bool)`

GetProcessorCoreCountOk returns a tuple with the ProcessorCoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCoreCount

`func (o *ServerInstanceGroupVariables) SetProcessorCoreCount(v int32)`

SetProcessorCoreCount sets ProcessorCoreCount field to given value.


### GetProcessorCoreMhz

`func (o *ServerInstanceGroupVariables) GetProcessorCoreMhz() int32`

GetProcessorCoreMhz returns the ProcessorCoreMhz field if non-nil, zero value otherwise.

### GetProcessorCoreMhzOk

`func (o *ServerInstanceGroupVariables) GetProcessorCoreMhzOk() (*int32, bool)`

GetProcessorCoreMhzOk returns a tuple with the ProcessorCoreMhz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCoreMhz

`func (o *ServerInstanceGroupVariables) SetProcessorCoreMhz(v int32)`

SetProcessorCoreMhz sets ProcessorCoreMhz field to given value.


### GetRamGbytes

`func (o *ServerInstanceGroupVariables) GetRamGbytes() int32`

GetRamGbytes returns the RamGbytes field if non-nil, zero value otherwise.

### GetRamGbytesOk

`func (o *ServerInstanceGroupVariables) GetRamGbytesOk() (*int32, bool)`

GetRamGbytesOk returns a tuple with the RamGbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGbytes

`func (o *ServerInstanceGroupVariables) SetRamGbytes(v int32)`

SetRamGbytes sets RamGbytes field to given value.

### HasRamGbytes

`func (o *ServerInstanceGroupVariables) HasRamGbytes() bool`

HasRamGbytes returns a boolean if a field has been set.

### GetDiskCount

`func (o *ServerInstanceGroupVariables) GetDiskCount() int32`

GetDiskCount returns the DiskCount field if non-nil, zero value otherwise.

### GetDiskCountOk

`func (o *ServerInstanceGroupVariables) GetDiskCountOk() (*int32, bool)`

GetDiskCountOk returns a tuple with the DiskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskCount

`func (o *ServerInstanceGroupVariables) SetDiskCount(v int32)`

SetDiskCount sets DiskCount field to given value.


### GetDiskSizeMbytes

`func (o *ServerInstanceGroupVariables) GetDiskSizeMbytes() int32`

GetDiskSizeMbytes returns the DiskSizeMbytes field if non-nil, zero value otherwise.

### GetDiskSizeMbytesOk

`func (o *ServerInstanceGroupVariables) GetDiskSizeMbytesOk() (*int32, bool)`

GetDiskSizeMbytesOk returns a tuple with the DiskSizeMbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeMbytes

`func (o *ServerInstanceGroupVariables) SetDiskSizeMbytes(v int32)`

SetDiskSizeMbytes sets DiskSizeMbytes field to given value.


### GetDiskTypes

`func (o *ServerInstanceGroupVariables) GetDiskTypes() []string`

GetDiskTypes returns the DiskTypes field if non-nil, zero value otherwise.

### GetDiskTypesOk

`func (o *ServerInstanceGroupVariables) GetDiskTypesOk() (*[]string, bool)`

GetDiskTypesOk returns a tuple with the DiskTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskTypes

`func (o *ServerInstanceGroupVariables) SetDiskTypes(v []string)`

SetDiskTypes sets DiskTypes field to given value.


### GetVirtualInterfacesEnabled

`func (o *ServerInstanceGroupVariables) GetVirtualInterfacesEnabled() int32`

GetVirtualInterfacesEnabled returns the VirtualInterfacesEnabled field if non-nil, zero value otherwise.

### GetVirtualInterfacesEnabledOk

`func (o *ServerInstanceGroupVariables) GetVirtualInterfacesEnabledOk() (*int32, bool)`

GetVirtualInterfacesEnabledOk returns a tuple with the VirtualInterfacesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualInterfacesEnabled

`func (o *ServerInstanceGroupVariables) SetVirtualInterfacesEnabled(v int32)`

SetVirtualInterfacesEnabled sets VirtualInterfacesEnabled field to given value.


### GetAdditionalWanIpv4Json

`func (o *ServerInstanceGroupVariables) GetAdditionalWanIpv4Json() map[string]interface{}`

GetAdditionalWanIpv4Json returns the AdditionalWanIpv4Json field if non-nil, zero value otherwise.

### GetAdditionalWanIpv4JsonOk

`func (o *ServerInstanceGroupVariables) GetAdditionalWanIpv4JsonOk() (*map[string]interface{}, bool)`

GetAdditionalWanIpv4JsonOk returns a tuple with the AdditionalWanIpv4Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalWanIpv4Json

`func (o *ServerInstanceGroupVariables) SetAdditionalWanIpv4Json(v map[string]interface{})`

SetAdditionalWanIpv4Json sets AdditionalWanIpv4Json field to given value.

### HasAdditionalWanIpv4Json

`func (o *ServerInstanceGroupVariables) HasAdditionalWanIpv4Json() bool`

HasAdditionalWanIpv4Json returns a boolean if a field has been set.

### GetNetworkProfileGroupId

`func (o *ServerInstanceGroupVariables) GetNetworkProfileGroupId() int32`

GetNetworkProfileGroupId returns the NetworkProfileGroupId field if non-nil, zero value otherwise.

### GetNetworkProfileGroupIdOk

`func (o *ServerInstanceGroupVariables) GetNetworkProfileGroupIdOk() (*int32, bool)`

GetNetworkProfileGroupIdOk returns a tuple with the NetworkProfileGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProfileGroupId

`func (o *ServerInstanceGroupVariables) SetNetworkProfileGroupId(v int32)`

SetNetworkProfileGroupId sets NetworkProfileGroupId field to given value.

### HasNetworkProfileGroupId

`func (o *ServerInstanceGroupVariables) HasNetworkProfileGroupId() bool`

HasNetworkProfileGroupId returns a boolean if a field has been set.

### GetNetworkProfileSnapshotId

`func (o *ServerInstanceGroupVariables) GetNetworkProfileSnapshotId() int32`

GetNetworkProfileSnapshotId returns the NetworkProfileSnapshotId field if non-nil, zero value otherwise.

### GetNetworkProfileSnapshotIdOk

`func (o *ServerInstanceGroupVariables) GetNetworkProfileSnapshotIdOk() (*int32, bool)`

GetNetworkProfileSnapshotIdOk returns a tuple with the NetworkProfileSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProfileSnapshotId

`func (o *ServerInstanceGroupVariables) SetNetworkProfileSnapshotId(v int32)`

SetNetworkProfileSnapshotId sets NetworkProfileSnapshotId field to given value.

### HasNetworkProfileSnapshotId

`func (o *ServerInstanceGroupVariables) HasNetworkProfileSnapshotId() bool`

HasNetworkProfileSnapshotId returns a boolean if a field has been set.

### GetOverrideIpv4WanVlanId

`func (o *ServerInstanceGroupVariables) GetOverrideIpv4WanVlanId() int32`

GetOverrideIpv4WanVlanId returns the OverrideIpv4WanVlanId field if non-nil, zero value otherwise.

### GetOverrideIpv4WanVlanIdOk

`func (o *ServerInstanceGroupVariables) GetOverrideIpv4WanVlanIdOk() (*int32, bool)`

GetOverrideIpv4WanVlanIdOk returns a tuple with the OverrideIpv4WanVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideIpv4WanVlanId

`func (o *ServerInstanceGroupVariables) SetOverrideIpv4WanVlanId(v int32)`

SetOverrideIpv4WanVlanId sets OverrideIpv4WanVlanId field to given value.

### HasOverrideIpv4WanVlanId

`func (o *ServerInstanceGroupVariables) HasOverrideIpv4WanVlanId() bool`

HasOverrideIpv4WanVlanId returns a boolean if a field has been set.

### GetNetworkEquipmentForceSubnetPoolIpv4WanId

`func (o *ServerInstanceGroupVariables) GetNetworkEquipmentForceSubnetPoolIpv4WanId() int32`

GetNetworkEquipmentForceSubnetPoolIpv4WanId returns the NetworkEquipmentForceSubnetPoolIpv4WanId field if non-nil, zero value otherwise.

### GetNetworkEquipmentForceSubnetPoolIpv4WanIdOk

`func (o *ServerInstanceGroupVariables) GetNetworkEquipmentForceSubnetPoolIpv4WanIdOk() (*int32, bool)`

GetNetworkEquipmentForceSubnetPoolIpv4WanIdOk returns a tuple with the NetworkEquipmentForceSubnetPoolIpv4WanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEquipmentForceSubnetPoolIpv4WanId

`func (o *ServerInstanceGroupVariables) SetNetworkEquipmentForceSubnetPoolIpv4WanId(v int32)`

SetNetworkEquipmentForceSubnetPoolIpv4WanId sets NetworkEquipmentForceSubnetPoolIpv4WanId field to given value.

### HasNetworkEquipmentForceSubnetPoolIpv4WanId

`func (o *ServerInstanceGroupVariables) HasNetworkEquipmentForceSubnetPoolIpv4WanId() bool`

HasNetworkEquipmentForceSubnetPoolIpv4WanId returns a boolean if a field has been set.

### GetServiceStatus

`func (o *ServerInstanceGroupVariables) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *ServerInstanceGroupVariables) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *ServerInstanceGroupVariables) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetResourcePoolId

`func (o *ServerInstanceGroupVariables) GetResourcePoolId() int32`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *ServerInstanceGroupVariables) GetResourcePoolIdOk() (*int32, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *ServerInstanceGroupVariables) SetResourcePoolId(v int32)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *ServerInstanceGroupVariables) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetIsVmGroup

`func (o *ServerInstanceGroupVariables) GetIsVmGroup() int32`

GetIsVmGroup returns the IsVmGroup field if non-nil, zero value otherwise.

### GetIsVmGroupOk

`func (o *ServerInstanceGroupVariables) GetIsVmGroupOk() (*int32, bool)`

GetIsVmGroupOk returns a tuple with the IsVmGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVmGroup

`func (o *ServerInstanceGroupVariables) SetIsVmGroup(v int32)`

SetIsVmGroup sets IsVmGroup field to given value.


### GetIsEndpointInstanceGroup

`func (o *ServerInstanceGroupVariables) GetIsEndpointInstanceGroup() int32`

GetIsEndpointInstanceGroup returns the IsEndpointInstanceGroup field if non-nil, zero value otherwise.

### GetIsEndpointInstanceGroupOk

`func (o *ServerInstanceGroupVariables) GetIsEndpointInstanceGroupOk() (*int32, bool)`

GetIsEndpointInstanceGroupOk returns a tuple with the IsEndpointInstanceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEndpointInstanceGroup

`func (o *ServerInstanceGroupVariables) SetIsEndpointInstanceGroup(v int32)`

SetIsEndpointInstanceGroup sets IsEndpointInstanceGroup field to given value.


### GetVmInstanceGroupId

`func (o *ServerInstanceGroupVariables) GetVmInstanceGroupId() int32`

GetVmInstanceGroupId returns the VmInstanceGroupId field if non-nil, zero value otherwise.

### GetVmInstanceGroupIdOk

`func (o *ServerInstanceGroupVariables) GetVmInstanceGroupIdOk() (*int32, bool)`

GetVmInstanceGroupIdOk returns a tuple with the VmInstanceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmInstanceGroupId

`func (o *ServerInstanceGroupVariables) SetVmInstanceGroupId(v int32)`

SetVmInstanceGroupId sets VmInstanceGroupId field to given value.

### HasVmInstanceGroupId

`func (o *ServerInstanceGroupVariables) HasVmInstanceGroupId() bool`

HasVmInstanceGroupId returns a boolean if a field has been set.

### GetNetworkEndpointGroupId

`func (o *ServerInstanceGroupVariables) GetNetworkEndpointGroupId() int32`

GetNetworkEndpointGroupId returns the NetworkEndpointGroupId field if non-nil, zero value otherwise.

### GetNetworkEndpointGroupIdOk

`func (o *ServerInstanceGroupVariables) GetNetworkEndpointGroupIdOk() (*int32, bool)`

GetNetworkEndpointGroupIdOk returns a tuple with the NetworkEndpointGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEndpointGroupId

`func (o *ServerInstanceGroupVariables) SetNetworkEndpointGroupId(v int32)`

SetNetworkEndpointGroupId sets NetworkEndpointGroupId field to given value.

### HasNetworkEndpointGroupId

`func (o *ServerInstanceGroupVariables) HasNetworkEndpointGroupId() bool`

HasNetworkEndpointGroupId returns a boolean if a field has been set.

### GetConfig

`func (o *ServerInstanceGroupVariables) GetConfig() ServerInstanceGroupConfiguration`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ServerInstanceGroupVariables) GetConfigOk() (*ServerInstanceGroupConfiguration, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ServerInstanceGroupVariables) SetConfig(v ServerInstanceGroupConfiguration)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ServerInstanceGroupVariables) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


