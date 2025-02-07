# ServerType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | The id of the server type. | 
**RamGbytes** | **float32** | The RAM GB of the server type. | 
**ProcessorCount** | **float32** | The processor count of the server type. | 
**ProcessorCoreMhz** | **float32** | The processor core Mhz of the server type. | 
**ProcessorCoreCount** | **float32** | The processor core count of the server type. | 
**Name** | **string** | The display name of the server type. | 
**DisplayName** | Pointer to **string** | The display name of the server type. | [optional] 
**Label** | Pointer to **string** | The label of the server type. | [optional] 
**NetworkTotalCapacityMbps** | **float32** | The total network capacity of the server type. | 
**NetworkInterfaceCount** | **float32** | The number of interfaces of the server type. | 
**NetworkInterfaceSpeeds** | **[]float32** | The network speeds of each interface of the server type. | 
**ProcessorNames** | **[]string** | The name of each processor of the server type. | 
**AllowedVendorSkuIds** | Pointer to **[]string** | The list of allowed SKU ids for the server type. | [optional] 
**DiskCount** | **float32** | The number of disks of the server type. | 
**IsExperimental** | Pointer to **float32** | Flag specifying if the server type is experimental. | [optional] [default to 0]
**ForUnmanagedServersOnly** | Pointer to **float32** | Flag specifying if the server type is only for unmanaged servers. | [optional] [default to 0]
**ForGenericEndpointsOnly** | Pointer to **float32** | Flag specifying if the server type is only for generic endpoints. | [optional] [default to 0]
**SupportsOobProvisioning** | Pointer to **float32** | Flag specifying if the server type supports OOB provisioning. | [optional] [default to 0]
**ServerClass** | **string** | The class of servers allowed for the server type. | 
**BootType** | Pointer to **string** | The server boot type allowed for the server type. | [optional] 
**Tags** | Pointer to **[]string** | The tags for the server type. | [optional] 
**GpuCount** | Pointer to **float32** | The number of GPUs for the server type. | [optional] [default to 0]
**GpuInfo** | Pointer to [**[]ServerGpuInfo**](ServerGpuInfo.md) | The information for each GPU of the server type. | [optional] 
**DiskGroups** | Pointer to [**ServerTypeDiskGroup**](ServerTypeDiskGroup.md) |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewServerType

`func NewServerType(id float32, ramGbytes float32, processorCount float32, processorCoreMhz float32, processorCoreCount float32, name string, networkTotalCapacityMbps float32, networkInterfaceCount float32, networkInterfaceSpeeds []float32, processorNames []string, diskCount float32, serverClass string, ) *ServerType`

NewServerType instantiates a new ServerType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerTypeWithDefaults

`func NewServerTypeWithDefaults() *ServerType`

NewServerTypeWithDefaults instantiates a new ServerType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerType) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerType) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerType) SetId(v float32)`

SetId sets Id field to given value.


### GetRamGbytes

`func (o *ServerType) GetRamGbytes() float32`

GetRamGbytes returns the RamGbytes field if non-nil, zero value otherwise.

### GetRamGbytesOk

`func (o *ServerType) GetRamGbytesOk() (*float32, bool)`

GetRamGbytesOk returns a tuple with the RamGbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGbytes

`func (o *ServerType) SetRamGbytes(v float32)`

SetRamGbytes sets RamGbytes field to given value.


### GetProcessorCount

`func (o *ServerType) GetProcessorCount() float32`

GetProcessorCount returns the ProcessorCount field if non-nil, zero value otherwise.

### GetProcessorCountOk

`func (o *ServerType) GetProcessorCountOk() (*float32, bool)`

GetProcessorCountOk returns a tuple with the ProcessorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCount

`func (o *ServerType) SetProcessorCount(v float32)`

SetProcessorCount sets ProcessorCount field to given value.


### GetProcessorCoreMhz

`func (o *ServerType) GetProcessorCoreMhz() float32`

GetProcessorCoreMhz returns the ProcessorCoreMhz field if non-nil, zero value otherwise.

### GetProcessorCoreMhzOk

`func (o *ServerType) GetProcessorCoreMhzOk() (*float32, bool)`

GetProcessorCoreMhzOk returns a tuple with the ProcessorCoreMhz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCoreMhz

`func (o *ServerType) SetProcessorCoreMhz(v float32)`

SetProcessorCoreMhz sets ProcessorCoreMhz field to given value.


### GetProcessorCoreCount

`func (o *ServerType) GetProcessorCoreCount() float32`

GetProcessorCoreCount returns the ProcessorCoreCount field if non-nil, zero value otherwise.

### GetProcessorCoreCountOk

`func (o *ServerType) GetProcessorCoreCountOk() (*float32, bool)`

GetProcessorCoreCountOk returns a tuple with the ProcessorCoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCoreCount

`func (o *ServerType) SetProcessorCoreCount(v float32)`

SetProcessorCoreCount sets ProcessorCoreCount field to given value.


### GetName

`func (o *ServerType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerType) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *ServerType) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ServerType) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ServerType) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ServerType) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLabel

`func (o *ServerType) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServerType) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServerType) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ServerType) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetNetworkTotalCapacityMbps

`func (o *ServerType) GetNetworkTotalCapacityMbps() float32`

GetNetworkTotalCapacityMbps returns the NetworkTotalCapacityMbps field if non-nil, zero value otherwise.

### GetNetworkTotalCapacityMbpsOk

`func (o *ServerType) GetNetworkTotalCapacityMbpsOk() (*float32, bool)`

GetNetworkTotalCapacityMbpsOk returns a tuple with the NetworkTotalCapacityMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTotalCapacityMbps

`func (o *ServerType) SetNetworkTotalCapacityMbps(v float32)`

SetNetworkTotalCapacityMbps sets NetworkTotalCapacityMbps field to given value.


### GetNetworkInterfaceCount

`func (o *ServerType) GetNetworkInterfaceCount() float32`

GetNetworkInterfaceCount returns the NetworkInterfaceCount field if non-nil, zero value otherwise.

### GetNetworkInterfaceCountOk

`func (o *ServerType) GetNetworkInterfaceCountOk() (*float32, bool)`

GetNetworkInterfaceCountOk returns a tuple with the NetworkInterfaceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceCount

`func (o *ServerType) SetNetworkInterfaceCount(v float32)`

SetNetworkInterfaceCount sets NetworkInterfaceCount field to given value.


### GetNetworkInterfaceSpeeds

`func (o *ServerType) GetNetworkInterfaceSpeeds() []float32`

GetNetworkInterfaceSpeeds returns the NetworkInterfaceSpeeds field if non-nil, zero value otherwise.

### GetNetworkInterfaceSpeedsOk

`func (o *ServerType) GetNetworkInterfaceSpeedsOk() (*[]float32, bool)`

GetNetworkInterfaceSpeedsOk returns a tuple with the NetworkInterfaceSpeeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceSpeeds

`func (o *ServerType) SetNetworkInterfaceSpeeds(v []float32)`

SetNetworkInterfaceSpeeds sets NetworkInterfaceSpeeds field to given value.


### GetProcessorNames

`func (o *ServerType) GetProcessorNames() []string`

GetProcessorNames returns the ProcessorNames field if non-nil, zero value otherwise.

### GetProcessorNamesOk

`func (o *ServerType) GetProcessorNamesOk() (*[]string, bool)`

GetProcessorNamesOk returns a tuple with the ProcessorNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorNames

`func (o *ServerType) SetProcessorNames(v []string)`

SetProcessorNames sets ProcessorNames field to given value.


### GetAllowedVendorSkuIds

`func (o *ServerType) GetAllowedVendorSkuIds() []string`

GetAllowedVendorSkuIds returns the AllowedVendorSkuIds field if non-nil, zero value otherwise.

### GetAllowedVendorSkuIdsOk

`func (o *ServerType) GetAllowedVendorSkuIdsOk() (*[]string, bool)`

GetAllowedVendorSkuIdsOk returns a tuple with the AllowedVendorSkuIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedVendorSkuIds

`func (o *ServerType) SetAllowedVendorSkuIds(v []string)`

SetAllowedVendorSkuIds sets AllowedVendorSkuIds field to given value.

### HasAllowedVendorSkuIds

`func (o *ServerType) HasAllowedVendorSkuIds() bool`

HasAllowedVendorSkuIds returns a boolean if a field has been set.

### GetDiskCount

`func (o *ServerType) GetDiskCount() float32`

GetDiskCount returns the DiskCount field if non-nil, zero value otherwise.

### GetDiskCountOk

`func (o *ServerType) GetDiskCountOk() (*float32, bool)`

GetDiskCountOk returns a tuple with the DiskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskCount

`func (o *ServerType) SetDiskCount(v float32)`

SetDiskCount sets DiskCount field to given value.


### GetIsExperimental

`func (o *ServerType) GetIsExperimental() float32`

GetIsExperimental returns the IsExperimental field if non-nil, zero value otherwise.

### GetIsExperimentalOk

`func (o *ServerType) GetIsExperimentalOk() (*float32, bool)`

GetIsExperimentalOk returns a tuple with the IsExperimental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExperimental

`func (o *ServerType) SetIsExperimental(v float32)`

SetIsExperimental sets IsExperimental field to given value.

### HasIsExperimental

`func (o *ServerType) HasIsExperimental() bool`

HasIsExperimental returns a boolean if a field has been set.

### GetForUnmanagedServersOnly

`func (o *ServerType) GetForUnmanagedServersOnly() float32`

GetForUnmanagedServersOnly returns the ForUnmanagedServersOnly field if non-nil, zero value otherwise.

### GetForUnmanagedServersOnlyOk

`func (o *ServerType) GetForUnmanagedServersOnlyOk() (*float32, bool)`

GetForUnmanagedServersOnlyOk returns a tuple with the ForUnmanagedServersOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForUnmanagedServersOnly

`func (o *ServerType) SetForUnmanagedServersOnly(v float32)`

SetForUnmanagedServersOnly sets ForUnmanagedServersOnly field to given value.

### HasForUnmanagedServersOnly

`func (o *ServerType) HasForUnmanagedServersOnly() bool`

HasForUnmanagedServersOnly returns a boolean if a field has been set.

### GetForGenericEndpointsOnly

`func (o *ServerType) GetForGenericEndpointsOnly() float32`

GetForGenericEndpointsOnly returns the ForGenericEndpointsOnly field if non-nil, zero value otherwise.

### GetForGenericEndpointsOnlyOk

`func (o *ServerType) GetForGenericEndpointsOnlyOk() (*float32, bool)`

GetForGenericEndpointsOnlyOk returns a tuple with the ForGenericEndpointsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForGenericEndpointsOnly

`func (o *ServerType) SetForGenericEndpointsOnly(v float32)`

SetForGenericEndpointsOnly sets ForGenericEndpointsOnly field to given value.

### HasForGenericEndpointsOnly

`func (o *ServerType) HasForGenericEndpointsOnly() bool`

HasForGenericEndpointsOnly returns a boolean if a field has been set.

### GetSupportsOobProvisioning

`func (o *ServerType) GetSupportsOobProvisioning() float32`

GetSupportsOobProvisioning returns the SupportsOobProvisioning field if non-nil, zero value otherwise.

### GetSupportsOobProvisioningOk

`func (o *ServerType) GetSupportsOobProvisioningOk() (*float32, bool)`

GetSupportsOobProvisioningOk returns a tuple with the SupportsOobProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsOobProvisioning

`func (o *ServerType) SetSupportsOobProvisioning(v float32)`

SetSupportsOobProvisioning sets SupportsOobProvisioning field to given value.

### HasSupportsOobProvisioning

`func (o *ServerType) HasSupportsOobProvisioning() bool`

HasSupportsOobProvisioning returns a boolean if a field has been set.

### GetServerClass

`func (o *ServerType) GetServerClass() string`

GetServerClass returns the ServerClass field if non-nil, zero value otherwise.

### GetServerClassOk

`func (o *ServerType) GetServerClassOk() (*string, bool)`

GetServerClassOk returns a tuple with the ServerClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClass

`func (o *ServerType) SetServerClass(v string)`

SetServerClass sets ServerClass field to given value.


### GetBootType

`func (o *ServerType) GetBootType() string`

GetBootType returns the BootType field if non-nil, zero value otherwise.

### GetBootTypeOk

`func (o *ServerType) GetBootTypeOk() (*string, bool)`

GetBootTypeOk returns a tuple with the BootType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootType

`func (o *ServerType) SetBootType(v string)`

SetBootType sets BootType field to given value.

### HasBootType

`func (o *ServerType) HasBootType() bool`

HasBootType returns a boolean if a field has been set.

### GetTags

`func (o *ServerType) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServerType) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServerType) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServerType) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetGpuCount

`func (o *ServerType) GetGpuCount() float32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *ServerType) GetGpuCountOk() (*float32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *ServerType) SetGpuCount(v float32)`

SetGpuCount sets GpuCount field to given value.

### HasGpuCount

`func (o *ServerType) HasGpuCount() bool`

HasGpuCount returns a boolean if a field has been set.

### GetGpuInfo

`func (o *ServerType) GetGpuInfo() []ServerGpuInfo`

GetGpuInfo returns the GpuInfo field if non-nil, zero value otherwise.

### GetGpuInfoOk

`func (o *ServerType) GetGpuInfoOk() (*[]ServerGpuInfo, bool)`

GetGpuInfoOk returns a tuple with the GpuInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuInfo

`func (o *ServerType) SetGpuInfo(v []ServerGpuInfo)`

SetGpuInfo sets GpuInfo field to given value.

### HasGpuInfo

`func (o *ServerType) HasGpuInfo() bool`

HasGpuInfo returns a boolean if a field has been set.

### GetDiskGroups

`func (o *ServerType) GetDiskGroups() ServerTypeDiskGroup`

GetDiskGroups returns the DiskGroups field if non-nil, zero value otherwise.

### GetDiskGroupsOk

`func (o *ServerType) GetDiskGroupsOk() (*ServerTypeDiskGroup, bool)`

GetDiskGroupsOk returns a tuple with the DiskGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskGroups

`func (o *ServerType) SetDiskGroups(v ServerTypeDiskGroup)`

SetDiskGroups sets DiskGroups field to given value.

### HasDiskGroups

`func (o *ServerType) HasDiskGroups() bool`

HasDiskGroups returns a boolean if a field has been set.

### GetLinks

`func (o *ServerType) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServerType) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServerType) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServerType) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


