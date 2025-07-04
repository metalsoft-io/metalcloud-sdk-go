# ServerInstanceGroupCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | The Product Instance label. Will be automatically generated if not provided. | [optional] 
**ServerGroupName** | Pointer to **string** |  | [optional] 
**ExtensionInstanceId** | Pointer to **int32** |  | [optional] 
**InstanceCount** | Pointer to **int32** | The number of instances to be created on the Instance Group. | [optional] [default to 1]
**DefaultCustomStorageProfile** | Pointer to [**ServerInstanceStorageProfile**](ServerInstanceStorageProfile.md) | Default Custom Storage Profile for the newly created Instances. | [optional] 
**IpAllocateAuto** | Pointer to **int32** | Automatically allocate IP addresses to child Instance&#x60;s Instance Interface elements. | [optional] [default to 1]
**Ipv4SubnetCreateAuto** | Pointer to **int32** | Automatically create or expand Subnet elements until the necessary IPv4 addresses are allocated. | [optional] [default to 1]
**Hostname** | Pointer to **string** | Custom hostname(subdomain) part of the fully qualified domain name (FQDN). If set, this will be used as the subdomain record part of the DNS Load Balancing record name instead of the default \&quot;instance-group\&quot;. The hostname must be a valid DNS subdomain and can only contain alphanumeric characters and hyphens. This will only take effect if the property \&quot;provisionLoadBalancingDnsRecord\&quot; is true.  | [optional] 
**OsTemplateId** | Pointer to **int32** | The volume template ID (or name) to use if the servers in the Instance Group have local disks. | [optional] 
**CustomVariables** | Pointer to **map[string]interface{}** | Object containing custom variables and variable overrides. | [optional] 
**ProcessorCount** | Pointer to **int32** | The CPU count on each instance. | [optional] [default to 1]
**ProcessorCoreCount** | Pointer to **int32** | The minimum cores of a CPU. | [optional] [default to 1]
**ProcessorCoreMhz** | Pointer to **int32** | The minimum clock speed of a CPU. | [optional] [default to 1000]
**RamGbytes** | Pointer to **int32** | The minimum RAM capacity of each instance. | [optional] [default to 1]
**DiskCount** | Pointer to **int32** | The minimum number of physical disks. | [optional] [default to 0]
**DiskSizeMbytes** | Pointer to **int32** | The minimum size of a single disk. | [optional] [default to 0]
**DiskTypes** | Pointer to **[]string** | The types of physical disks. | [optional] [default to []]
**VirtualInterfacesEnabled** | Pointer to **int32** | Enable virtual interfaces | [optional] [default to 0]
**AdditionalWanIpv4Json** | Pointer to **map[string]interface{}** | Contains info about additional ips to be assigned to the WAN interfaces. | [optional] 
**OverrideIpv4WanVlanId** | Pointer to **int32** | The ipv4 vlan that should override the default from the WAN Network for the primary ip. | [optional] 
**NetworkEquipmentForceSubnetPoolIpv4WanId** | Pointer to **int32** | ID of a ipv4 WAN subnet-pool from which to force the subnet allocation for the Instance Interfaces associated with this Instance Group. | [optional] 
**ResourcePoolId** | Pointer to **int32** | The resource pool assigned to this instance array | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**DefaultServerTypeId** | **int32** | The server type ID that will be assigned to newly created instances. | 

## Methods

### NewServerInstanceGroupCreate

`func NewServerInstanceGroupCreate(defaultServerTypeId int32, ) *ServerInstanceGroupCreate`

NewServerInstanceGroupCreate instantiates a new ServerInstanceGroupCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceGroupCreateWithDefaults

`func NewServerInstanceGroupCreateWithDefaults() *ServerInstanceGroupCreate`

NewServerInstanceGroupCreateWithDefaults instantiates a new ServerInstanceGroupCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ServerInstanceGroupCreate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServerInstanceGroupCreate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServerInstanceGroupCreate) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ServerInstanceGroupCreate) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetServerGroupName

`func (o *ServerInstanceGroupCreate) GetServerGroupName() string`

GetServerGroupName returns the ServerGroupName field if non-nil, zero value otherwise.

### GetServerGroupNameOk

`func (o *ServerInstanceGroupCreate) GetServerGroupNameOk() (*string, bool)`

GetServerGroupNameOk returns a tuple with the ServerGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroupName

`func (o *ServerInstanceGroupCreate) SetServerGroupName(v string)`

SetServerGroupName sets ServerGroupName field to given value.

### HasServerGroupName

`func (o *ServerInstanceGroupCreate) HasServerGroupName() bool`

HasServerGroupName returns a boolean if a field has been set.

### GetExtensionInstanceId

`func (o *ServerInstanceGroupCreate) GetExtensionInstanceId() int32`

GetExtensionInstanceId returns the ExtensionInstanceId field if non-nil, zero value otherwise.

### GetExtensionInstanceIdOk

`func (o *ServerInstanceGroupCreate) GetExtensionInstanceIdOk() (*int32, bool)`

GetExtensionInstanceIdOk returns a tuple with the ExtensionInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionInstanceId

`func (o *ServerInstanceGroupCreate) SetExtensionInstanceId(v int32)`

SetExtensionInstanceId sets ExtensionInstanceId field to given value.

### HasExtensionInstanceId

`func (o *ServerInstanceGroupCreate) HasExtensionInstanceId() bool`

HasExtensionInstanceId returns a boolean if a field has been set.

### GetInstanceCount

`func (o *ServerInstanceGroupCreate) GetInstanceCount() int32`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *ServerInstanceGroupCreate) GetInstanceCountOk() (*int32, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *ServerInstanceGroupCreate) SetInstanceCount(v int32)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *ServerInstanceGroupCreate) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

### GetDefaultCustomStorageProfile

`func (o *ServerInstanceGroupCreate) GetDefaultCustomStorageProfile() ServerInstanceStorageProfile`

GetDefaultCustomStorageProfile returns the DefaultCustomStorageProfile field if non-nil, zero value otherwise.

### GetDefaultCustomStorageProfileOk

`func (o *ServerInstanceGroupCreate) GetDefaultCustomStorageProfileOk() (*ServerInstanceStorageProfile, bool)`

GetDefaultCustomStorageProfileOk returns a tuple with the DefaultCustomStorageProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCustomStorageProfile

`func (o *ServerInstanceGroupCreate) SetDefaultCustomStorageProfile(v ServerInstanceStorageProfile)`

SetDefaultCustomStorageProfile sets DefaultCustomStorageProfile field to given value.

### HasDefaultCustomStorageProfile

`func (o *ServerInstanceGroupCreate) HasDefaultCustomStorageProfile() bool`

HasDefaultCustomStorageProfile returns a boolean if a field has been set.

### GetIpAllocateAuto

`func (o *ServerInstanceGroupCreate) GetIpAllocateAuto() int32`

GetIpAllocateAuto returns the IpAllocateAuto field if non-nil, zero value otherwise.

### GetIpAllocateAutoOk

`func (o *ServerInstanceGroupCreate) GetIpAllocateAutoOk() (*int32, bool)`

GetIpAllocateAutoOk returns a tuple with the IpAllocateAuto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAllocateAuto

`func (o *ServerInstanceGroupCreate) SetIpAllocateAuto(v int32)`

SetIpAllocateAuto sets IpAllocateAuto field to given value.

### HasIpAllocateAuto

`func (o *ServerInstanceGroupCreate) HasIpAllocateAuto() bool`

HasIpAllocateAuto returns a boolean if a field has been set.

### GetIpv4SubnetCreateAuto

`func (o *ServerInstanceGroupCreate) GetIpv4SubnetCreateAuto() int32`

GetIpv4SubnetCreateAuto returns the Ipv4SubnetCreateAuto field if non-nil, zero value otherwise.

### GetIpv4SubnetCreateAutoOk

`func (o *ServerInstanceGroupCreate) GetIpv4SubnetCreateAutoOk() (*int32, bool)`

GetIpv4SubnetCreateAutoOk returns a tuple with the Ipv4SubnetCreateAuto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4SubnetCreateAuto

`func (o *ServerInstanceGroupCreate) SetIpv4SubnetCreateAuto(v int32)`

SetIpv4SubnetCreateAuto sets Ipv4SubnetCreateAuto field to given value.

### HasIpv4SubnetCreateAuto

`func (o *ServerInstanceGroupCreate) HasIpv4SubnetCreateAuto() bool`

HasIpv4SubnetCreateAuto returns a boolean if a field has been set.

### GetHostname

`func (o *ServerInstanceGroupCreate) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ServerInstanceGroupCreate) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ServerInstanceGroupCreate) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ServerInstanceGroupCreate) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetOsTemplateId

`func (o *ServerInstanceGroupCreate) GetOsTemplateId() int32`

GetOsTemplateId returns the OsTemplateId field if non-nil, zero value otherwise.

### GetOsTemplateIdOk

`func (o *ServerInstanceGroupCreate) GetOsTemplateIdOk() (*int32, bool)`

GetOsTemplateIdOk returns a tuple with the OsTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsTemplateId

`func (o *ServerInstanceGroupCreate) SetOsTemplateId(v int32)`

SetOsTemplateId sets OsTemplateId field to given value.

### HasOsTemplateId

`func (o *ServerInstanceGroupCreate) HasOsTemplateId() bool`

HasOsTemplateId returns a boolean if a field has been set.

### GetCustomVariables

`func (o *ServerInstanceGroupCreate) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *ServerInstanceGroupCreate) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *ServerInstanceGroupCreate) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *ServerInstanceGroupCreate) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.

### GetProcessorCount

`func (o *ServerInstanceGroupCreate) GetProcessorCount() int32`

GetProcessorCount returns the ProcessorCount field if non-nil, zero value otherwise.

### GetProcessorCountOk

`func (o *ServerInstanceGroupCreate) GetProcessorCountOk() (*int32, bool)`

GetProcessorCountOk returns a tuple with the ProcessorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCount

`func (o *ServerInstanceGroupCreate) SetProcessorCount(v int32)`

SetProcessorCount sets ProcessorCount field to given value.

### HasProcessorCount

`func (o *ServerInstanceGroupCreate) HasProcessorCount() bool`

HasProcessorCount returns a boolean if a field has been set.

### GetProcessorCoreCount

`func (o *ServerInstanceGroupCreate) GetProcessorCoreCount() int32`

GetProcessorCoreCount returns the ProcessorCoreCount field if non-nil, zero value otherwise.

### GetProcessorCoreCountOk

`func (o *ServerInstanceGroupCreate) GetProcessorCoreCountOk() (*int32, bool)`

GetProcessorCoreCountOk returns a tuple with the ProcessorCoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCoreCount

`func (o *ServerInstanceGroupCreate) SetProcessorCoreCount(v int32)`

SetProcessorCoreCount sets ProcessorCoreCount field to given value.

### HasProcessorCoreCount

`func (o *ServerInstanceGroupCreate) HasProcessorCoreCount() bool`

HasProcessorCoreCount returns a boolean if a field has been set.

### GetProcessorCoreMhz

`func (o *ServerInstanceGroupCreate) GetProcessorCoreMhz() int32`

GetProcessorCoreMhz returns the ProcessorCoreMhz field if non-nil, zero value otherwise.

### GetProcessorCoreMhzOk

`func (o *ServerInstanceGroupCreate) GetProcessorCoreMhzOk() (*int32, bool)`

GetProcessorCoreMhzOk returns a tuple with the ProcessorCoreMhz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCoreMhz

`func (o *ServerInstanceGroupCreate) SetProcessorCoreMhz(v int32)`

SetProcessorCoreMhz sets ProcessorCoreMhz field to given value.

### HasProcessorCoreMhz

`func (o *ServerInstanceGroupCreate) HasProcessorCoreMhz() bool`

HasProcessorCoreMhz returns a boolean if a field has been set.

### GetRamGbytes

`func (o *ServerInstanceGroupCreate) GetRamGbytes() int32`

GetRamGbytes returns the RamGbytes field if non-nil, zero value otherwise.

### GetRamGbytesOk

`func (o *ServerInstanceGroupCreate) GetRamGbytesOk() (*int32, bool)`

GetRamGbytesOk returns a tuple with the RamGbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGbytes

`func (o *ServerInstanceGroupCreate) SetRamGbytes(v int32)`

SetRamGbytes sets RamGbytes field to given value.

### HasRamGbytes

`func (o *ServerInstanceGroupCreate) HasRamGbytes() bool`

HasRamGbytes returns a boolean if a field has been set.

### GetDiskCount

`func (o *ServerInstanceGroupCreate) GetDiskCount() int32`

GetDiskCount returns the DiskCount field if non-nil, zero value otherwise.

### GetDiskCountOk

`func (o *ServerInstanceGroupCreate) GetDiskCountOk() (*int32, bool)`

GetDiskCountOk returns a tuple with the DiskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskCount

`func (o *ServerInstanceGroupCreate) SetDiskCount(v int32)`

SetDiskCount sets DiskCount field to given value.

### HasDiskCount

`func (o *ServerInstanceGroupCreate) HasDiskCount() bool`

HasDiskCount returns a boolean if a field has been set.

### GetDiskSizeMbytes

`func (o *ServerInstanceGroupCreate) GetDiskSizeMbytes() int32`

GetDiskSizeMbytes returns the DiskSizeMbytes field if non-nil, zero value otherwise.

### GetDiskSizeMbytesOk

`func (o *ServerInstanceGroupCreate) GetDiskSizeMbytesOk() (*int32, bool)`

GetDiskSizeMbytesOk returns a tuple with the DiskSizeMbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeMbytes

`func (o *ServerInstanceGroupCreate) SetDiskSizeMbytes(v int32)`

SetDiskSizeMbytes sets DiskSizeMbytes field to given value.

### HasDiskSizeMbytes

`func (o *ServerInstanceGroupCreate) HasDiskSizeMbytes() bool`

HasDiskSizeMbytes returns a boolean if a field has been set.

### GetDiskTypes

`func (o *ServerInstanceGroupCreate) GetDiskTypes() []string`

GetDiskTypes returns the DiskTypes field if non-nil, zero value otherwise.

### GetDiskTypesOk

`func (o *ServerInstanceGroupCreate) GetDiskTypesOk() (*[]string, bool)`

GetDiskTypesOk returns a tuple with the DiskTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskTypes

`func (o *ServerInstanceGroupCreate) SetDiskTypes(v []string)`

SetDiskTypes sets DiskTypes field to given value.

### HasDiskTypes

`func (o *ServerInstanceGroupCreate) HasDiskTypes() bool`

HasDiskTypes returns a boolean if a field has been set.

### GetVirtualInterfacesEnabled

`func (o *ServerInstanceGroupCreate) GetVirtualInterfacesEnabled() int32`

GetVirtualInterfacesEnabled returns the VirtualInterfacesEnabled field if non-nil, zero value otherwise.

### GetVirtualInterfacesEnabledOk

`func (o *ServerInstanceGroupCreate) GetVirtualInterfacesEnabledOk() (*int32, bool)`

GetVirtualInterfacesEnabledOk returns a tuple with the VirtualInterfacesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualInterfacesEnabled

`func (o *ServerInstanceGroupCreate) SetVirtualInterfacesEnabled(v int32)`

SetVirtualInterfacesEnabled sets VirtualInterfacesEnabled field to given value.

### HasVirtualInterfacesEnabled

`func (o *ServerInstanceGroupCreate) HasVirtualInterfacesEnabled() bool`

HasVirtualInterfacesEnabled returns a boolean if a field has been set.

### GetAdditionalWanIpv4Json

`func (o *ServerInstanceGroupCreate) GetAdditionalWanIpv4Json() map[string]interface{}`

GetAdditionalWanIpv4Json returns the AdditionalWanIpv4Json field if non-nil, zero value otherwise.

### GetAdditionalWanIpv4JsonOk

`func (o *ServerInstanceGroupCreate) GetAdditionalWanIpv4JsonOk() (*map[string]interface{}, bool)`

GetAdditionalWanIpv4JsonOk returns a tuple with the AdditionalWanIpv4Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalWanIpv4Json

`func (o *ServerInstanceGroupCreate) SetAdditionalWanIpv4Json(v map[string]interface{})`

SetAdditionalWanIpv4Json sets AdditionalWanIpv4Json field to given value.

### HasAdditionalWanIpv4Json

`func (o *ServerInstanceGroupCreate) HasAdditionalWanIpv4Json() bool`

HasAdditionalWanIpv4Json returns a boolean if a field has been set.

### GetOverrideIpv4WanVlanId

`func (o *ServerInstanceGroupCreate) GetOverrideIpv4WanVlanId() int32`

GetOverrideIpv4WanVlanId returns the OverrideIpv4WanVlanId field if non-nil, zero value otherwise.

### GetOverrideIpv4WanVlanIdOk

`func (o *ServerInstanceGroupCreate) GetOverrideIpv4WanVlanIdOk() (*int32, bool)`

GetOverrideIpv4WanVlanIdOk returns a tuple with the OverrideIpv4WanVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideIpv4WanVlanId

`func (o *ServerInstanceGroupCreate) SetOverrideIpv4WanVlanId(v int32)`

SetOverrideIpv4WanVlanId sets OverrideIpv4WanVlanId field to given value.

### HasOverrideIpv4WanVlanId

`func (o *ServerInstanceGroupCreate) HasOverrideIpv4WanVlanId() bool`

HasOverrideIpv4WanVlanId returns a boolean if a field has been set.

### GetNetworkEquipmentForceSubnetPoolIpv4WanId

`func (o *ServerInstanceGroupCreate) GetNetworkEquipmentForceSubnetPoolIpv4WanId() int32`

GetNetworkEquipmentForceSubnetPoolIpv4WanId returns the NetworkEquipmentForceSubnetPoolIpv4WanId field if non-nil, zero value otherwise.

### GetNetworkEquipmentForceSubnetPoolIpv4WanIdOk

`func (o *ServerInstanceGroupCreate) GetNetworkEquipmentForceSubnetPoolIpv4WanIdOk() (*int32, bool)`

GetNetworkEquipmentForceSubnetPoolIpv4WanIdOk returns a tuple with the NetworkEquipmentForceSubnetPoolIpv4WanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEquipmentForceSubnetPoolIpv4WanId

`func (o *ServerInstanceGroupCreate) SetNetworkEquipmentForceSubnetPoolIpv4WanId(v int32)`

SetNetworkEquipmentForceSubnetPoolIpv4WanId sets NetworkEquipmentForceSubnetPoolIpv4WanId field to given value.

### HasNetworkEquipmentForceSubnetPoolIpv4WanId

`func (o *ServerInstanceGroupCreate) HasNetworkEquipmentForceSubnetPoolIpv4WanId() bool`

HasNetworkEquipmentForceSubnetPoolIpv4WanId returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *ServerInstanceGroupCreate) GetResourcePoolId() int32`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *ServerInstanceGroupCreate) GetResourcePoolIdOk() (*int32, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *ServerInstanceGroupCreate) SetResourcePoolId(v int32)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *ServerInstanceGroupCreate) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetTags

`func (o *ServerInstanceGroupCreate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServerInstanceGroupCreate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServerInstanceGroupCreate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServerInstanceGroupCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDefaultServerTypeId

`func (o *ServerInstanceGroupCreate) GetDefaultServerTypeId() int32`

GetDefaultServerTypeId returns the DefaultServerTypeId field if non-nil, zero value otherwise.

### GetDefaultServerTypeIdOk

`func (o *ServerInstanceGroupCreate) GetDefaultServerTypeIdOk() (*int32, bool)`

GetDefaultServerTypeIdOk returns a tuple with the DefaultServerTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultServerTypeId

`func (o *ServerInstanceGroupCreate) SetDefaultServerTypeId(v int32)`

SetDefaultServerTypeId sets DefaultServerTypeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


