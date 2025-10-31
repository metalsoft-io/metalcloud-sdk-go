# CreateServerType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RamGbytes** | **float32** | The RAM GB of the server type. | 
**ProcessorCount** | **float32** | The processor count of the server type. | 
**ProcessorCoreMhz** | **float32** | The processor core Mhz of the server type. | 
**ProcessorCoreCount** | **float32** | The processor core count of the server type. | 
**Name** | **string** | The human-readable name of the server type. | 
**Description** | Pointer to **string** | Description of the server type. | [optional] 
**Label** | **string** | The label of the server type. | 
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

## Methods

### NewCreateServerType

`func NewCreateServerType(ramGbytes float32, processorCount float32, processorCoreMhz float32, processorCoreCount float32, name string, label string, networkInterfaceSpeeds []float32, processorNames []string, diskCount float32, serverClass string, ) *CreateServerType`

NewCreateServerType instantiates a new CreateServerType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServerTypeWithDefaults

`func NewCreateServerTypeWithDefaults() *CreateServerType`

NewCreateServerTypeWithDefaults instantiates a new CreateServerType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRamGbytes

`func (o *CreateServerType) GetRamGbytes() float32`

GetRamGbytes returns the RamGbytes field if non-nil, zero value otherwise.

### GetRamGbytesOk

`func (o *CreateServerType) GetRamGbytesOk() (*float32, bool)`

GetRamGbytesOk returns a tuple with the RamGbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGbytes

`func (o *CreateServerType) SetRamGbytes(v float32)`

SetRamGbytes sets RamGbytes field to given value.


### GetProcessorCount

`func (o *CreateServerType) GetProcessorCount() float32`

GetProcessorCount returns the ProcessorCount field if non-nil, zero value otherwise.

### GetProcessorCountOk

`func (o *CreateServerType) GetProcessorCountOk() (*float32, bool)`

GetProcessorCountOk returns a tuple with the ProcessorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCount

`func (o *CreateServerType) SetProcessorCount(v float32)`

SetProcessorCount sets ProcessorCount field to given value.


### GetProcessorCoreMhz

`func (o *CreateServerType) GetProcessorCoreMhz() float32`

GetProcessorCoreMhz returns the ProcessorCoreMhz field if non-nil, zero value otherwise.

### GetProcessorCoreMhzOk

`func (o *CreateServerType) GetProcessorCoreMhzOk() (*float32, bool)`

GetProcessorCoreMhzOk returns a tuple with the ProcessorCoreMhz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCoreMhz

`func (o *CreateServerType) SetProcessorCoreMhz(v float32)`

SetProcessorCoreMhz sets ProcessorCoreMhz field to given value.


### GetProcessorCoreCount

`func (o *CreateServerType) GetProcessorCoreCount() float32`

GetProcessorCoreCount returns the ProcessorCoreCount field if non-nil, zero value otherwise.

### GetProcessorCoreCountOk

`func (o *CreateServerType) GetProcessorCoreCountOk() (*float32, bool)`

GetProcessorCoreCountOk returns a tuple with the ProcessorCoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorCoreCount

`func (o *CreateServerType) SetProcessorCoreCount(v float32)`

SetProcessorCoreCount sets ProcessorCoreCount field to given value.


### GetName

`func (o *CreateServerType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateServerType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateServerType) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateServerType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateServerType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateServerType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateServerType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *CreateServerType) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateServerType) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateServerType) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetNetworkInterfaceSpeeds

`func (o *CreateServerType) GetNetworkInterfaceSpeeds() []float32`

GetNetworkInterfaceSpeeds returns the NetworkInterfaceSpeeds field if non-nil, zero value otherwise.

### GetNetworkInterfaceSpeedsOk

`func (o *CreateServerType) GetNetworkInterfaceSpeedsOk() (*[]float32, bool)`

GetNetworkInterfaceSpeedsOk returns a tuple with the NetworkInterfaceSpeeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceSpeeds

`func (o *CreateServerType) SetNetworkInterfaceSpeeds(v []float32)`

SetNetworkInterfaceSpeeds sets NetworkInterfaceSpeeds field to given value.


### GetProcessorNames

`func (o *CreateServerType) GetProcessorNames() []string`

GetProcessorNames returns the ProcessorNames field if non-nil, zero value otherwise.

### GetProcessorNamesOk

`func (o *CreateServerType) GetProcessorNamesOk() (*[]string, bool)`

GetProcessorNamesOk returns a tuple with the ProcessorNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorNames

`func (o *CreateServerType) SetProcessorNames(v []string)`

SetProcessorNames sets ProcessorNames field to given value.


### GetAllowedVendorSkuIds

`func (o *CreateServerType) GetAllowedVendorSkuIds() []string`

GetAllowedVendorSkuIds returns the AllowedVendorSkuIds field if non-nil, zero value otherwise.

### GetAllowedVendorSkuIdsOk

`func (o *CreateServerType) GetAllowedVendorSkuIdsOk() (*[]string, bool)`

GetAllowedVendorSkuIdsOk returns a tuple with the AllowedVendorSkuIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedVendorSkuIds

`func (o *CreateServerType) SetAllowedVendorSkuIds(v []string)`

SetAllowedVendorSkuIds sets AllowedVendorSkuIds field to given value.

### HasAllowedVendorSkuIds

`func (o *CreateServerType) HasAllowedVendorSkuIds() bool`

HasAllowedVendorSkuIds returns a boolean if a field has been set.

### GetDiskCount

`func (o *CreateServerType) GetDiskCount() float32`

GetDiskCount returns the DiskCount field if non-nil, zero value otherwise.

### GetDiskCountOk

`func (o *CreateServerType) GetDiskCountOk() (*float32, bool)`

GetDiskCountOk returns a tuple with the DiskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskCount

`func (o *CreateServerType) SetDiskCount(v float32)`

SetDiskCount sets DiskCount field to given value.


### GetIsExperimental

`func (o *CreateServerType) GetIsExperimental() float32`

GetIsExperimental returns the IsExperimental field if non-nil, zero value otherwise.

### GetIsExperimentalOk

`func (o *CreateServerType) GetIsExperimentalOk() (*float32, bool)`

GetIsExperimentalOk returns a tuple with the IsExperimental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExperimental

`func (o *CreateServerType) SetIsExperimental(v float32)`

SetIsExperimental sets IsExperimental field to given value.

### HasIsExperimental

`func (o *CreateServerType) HasIsExperimental() bool`

HasIsExperimental returns a boolean if a field has been set.

### GetForUnmanagedServersOnly

`func (o *CreateServerType) GetForUnmanagedServersOnly() float32`

GetForUnmanagedServersOnly returns the ForUnmanagedServersOnly field if non-nil, zero value otherwise.

### GetForUnmanagedServersOnlyOk

`func (o *CreateServerType) GetForUnmanagedServersOnlyOk() (*float32, bool)`

GetForUnmanagedServersOnlyOk returns a tuple with the ForUnmanagedServersOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForUnmanagedServersOnly

`func (o *CreateServerType) SetForUnmanagedServersOnly(v float32)`

SetForUnmanagedServersOnly sets ForUnmanagedServersOnly field to given value.

### HasForUnmanagedServersOnly

`func (o *CreateServerType) HasForUnmanagedServersOnly() bool`

HasForUnmanagedServersOnly returns a boolean if a field has been set.

### GetForGenericEndpointsOnly

`func (o *CreateServerType) GetForGenericEndpointsOnly() float32`

GetForGenericEndpointsOnly returns the ForGenericEndpointsOnly field if non-nil, zero value otherwise.

### GetForGenericEndpointsOnlyOk

`func (o *CreateServerType) GetForGenericEndpointsOnlyOk() (*float32, bool)`

GetForGenericEndpointsOnlyOk returns a tuple with the ForGenericEndpointsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForGenericEndpointsOnly

`func (o *CreateServerType) SetForGenericEndpointsOnly(v float32)`

SetForGenericEndpointsOnly sets ForGenericEndpointsOnly field to given value.

### HasForGenericEndpointsOnly

`func (o *CreateServerType) HasForGenericEndpointsOnly() bool`

HasForGenericEndpointsOnly returns a boolean if a field has been set.

### GetSupportsOobProvisioning

`func (o *CreateServerType) GetSupportsOobProvisioning() float32`

GetSupportsOobProvisioning returns the SupportsOobProvisioning field if non-nil, zero value otherwise.

### GetSupportsOobProvisioningOk

`func (o *CreateServerType) GetSupportsOobProvisioningOk() (*float32, bool)`

GetSupportsOobProvisioningOk returns a tuple with the SupportsOobProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsOobProvisioning

`func (o *CreateServerType) SetSupportsOobProvisioning(v float32)`

SetSupportsOobProvisioning sets SupportsOobProvisioning field to given value.

### HasSupportsOobProvisioning

`func (o *CreateServerType) HasSupportsOobProvisioning() bool`

HasSupportsOobProvisioning returns a boolean if a field has been set.

### GetServerClass

`func (o *CreateServerType) GetServerClass() string`

GetServerClass returns the ServerClass field if non-nil, zero value otherwise.

### GetServerClassOk

`func (o *CreateServerType) GetServerClassOk() (*string, bool)`

GetServerClassOk returns a tuple with the ServerClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClass

`func (o *CreateServerType) SetServerClass(v string)`

SetServerClass sets ServerClass field to given value.


### GetBootType

`func (o *CreateServerType) GetBootType() string`

GetBootType returns the BootType field if non-nil, zero value otherwise.

### GetBootTypeOk

`func (o *CreateServerType) GetBootTypeOk() (*string, bool)`

GetBootTypeOk returns a tuple with the BootType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootType

`func (o *CreateServerType) SetBootType(v string)`

SetBootType sets BootType field to given value.

### HasBootType

`func (o *CreateServerType) HasBootType() bool`

HasBootType returns a boolean if a field has been set.

### GetTags

`func (o *CreateServerType) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateServerType) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateServerType) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateServerType) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetGpuCount

`func (o *CreateServerType) GetGpuCount() float32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *CreateServerType) GetGpuCountOk() (*float32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *CreateServerType) SetGpuCount(v float32)`

SetGpuCount sets GpuCount field to given value.

### HasGpuCount

`func (o *CreateServerType) HasGpuCount() bool`

HasGpuCount returns a boolean if a field has been set.

### GetGpuInfo

`func (o *CreateServerType) GetGpuInfo() []ServerGpuInfo`

GetGpuInfo returns the GpuInfo field if non-nil, zero value otherwise.

### GetGpuInfoOk

`func (o *CreateServerType) GetGpuInfoOk() (*[]ServerGpuInfo, bool)`

GetGpuInfoOk returns a tuple with the GpuInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuInfo

`func (o *CreateServerType) SetGpuInfo(v []ServerGpuInfo)`

SetGpuInfo sets GpuInfo field to given value.

### HasGpuInfo

`func (o *CreateServerType) HasGpuInfo() bool`

HasGpuInfo returns a boolean if a field has been set.

### GetDiskGroups

`func (o *CreateServerType) GetDiskGroups() ServerTypeDiskGroup`

GetDiskGroups returns the DiskGroups field if non-nil, zero value otherwise.

### GetDiskGroupsOk

`func (o *CreateServerType) GetDiskGroupsOk() (*ServerTypeDiskGroup, bool)`

GetDiskGroupsOk returns a tuple with the DiskGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskGroups

`func (o *CreateServerType) SetDiskGroups(v ServerTypeDiskGroup)`

SetDiskGroups sets DiskGroups field to given value.

### HasDiskGroups

`func (o *CreateServerType) HasDiskGroups() bool`

HasDiskGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


