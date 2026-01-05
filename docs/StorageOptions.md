# StorageOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableDataReduction** | Pointer to **bool** | Enable data reduction | [optional] 
**EnableAdvancedDeduplication** | Pointer to **bool** | Enable advanced deduplication | [optional] 
**EnableDataReductionSharedVolume** | Pointer to **bool** | Enable data reduction on a newly created shared volume | [optional] 
**VolumeName** | Pointer to **string** | Volume name | [optional] 
**ArrayId** | Pointer to **string** | Array id to use (for certain storage drivers) | [optional] 
**DirectorId** | Pointer to **string** | Director id to use (for certain storage drivers) | [optional] 
**S3Hostname** | Pointer to **string** | S3 Hostname to use (for certain storage drivers) | [optional] 
**S3Port** | Pointer to **float32** | Enable advanced deduplication | [optional] 
**DefaultServiceLevel** | Pointer to **string** | Default service level to use (for certain storage drivers) | [optional] 
**FibreChannelEnabled** | Pointer to **bool** | Fibre channel enabled | [optional] 
**Directors** | Pointer to **[]string** | Array of directors | [optional] 
**ResourcePool** | Pointer to **string** | Resource pool | [optional] 
**Version** | Pointer to **string** | Version of the storage | [optional] 
**MajorVersion** | Pointer to **string** | Major version of the storage | [optional] 
**Arrays** | Pointer to **[]string** | Arrays | [optional] 
**Info** | Pointer to **map[string]interface{}** | Storage info | [optional] 
**ServiceLevelNames** | Pointer to **[]string** | Service level names | [optional] 
**FibreChannelCapable** | Pointer to **bool** | Fibre channel capable | [optional] 
**PortsIscsi** | Pointer to [**[]StoragePort**](StoragePort.md) | ISCSI ports | [optional] 
**PortsScsiFc** | Pointer to [**[]StoragePort**](StoragePort.md) | SCSI FC ports | [optional] 
**PortsNvmeTcp** | Pointer to [**[]StoragePort**](StoragePort.md) | NVMe TCP ports | [optional] 
**PortsNvmeFc** | Pointer to [**[]StoragePort**](StoragePort.md) | NVMe FC ports | [optional] 
**PortsToUse** | Pointer to [**[]StoragePort**](StoragePort.md) | Array of storage ports to use | [optional] 
**Servers** | Pointer to [**[]StoragePort**](StoragePort.md) | Array of storage servers | [optional] 
**Nodes** | Pointer to [**[]StorageNodes**](StorageNodes.md) | Array of storage nodes | [optional] 
**InfoGatherError** | Pointer to **string** | Error message when gathering storage info | [optional] 
**ConfigureError** | Pointer to **string** | Error message when configuring storage | [optional] 

## Methods

### NewStorageOptions

`func NewStorageOptions() *StorageOptions`

NewStorageOptions instantiates a new StorageOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageOptionsWithDefaults

`func NewStorageOptionsWithDefaults() *StorageOptions`

NewStorageOptionsWithDefaults instantiates a new StorageOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableDataReduction

`func (o *StorageOptions) GetEnableDataReduction() bool`

GetEnableDataReduction returns the EnableDataReduction field if non-nil, zero value otherwise.

### GetEnableDataReductionOk

`func (o *StorageOptions) GetEnableDataReductionOk() (*bool, bool)`

GetEnableDataReductionOk returns a tuple with the EnableDataReduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDataReduction

`func (o *StorageOptions) SetEnableDataReduction(v bool)`

SetEnableDataReduction sets EnableDataReduction field to given value.

### HasEnableDataReduction

`func (o *StorageOptions) HasEnableDataReduction() bool`

HasEnableDataReduction returns a boolean if a field has been set.

### GetEnableAdvancedDeduplication

`func (o *StorageOptions) GetEnableAdvancedDeduplication() bool`

GetEnableAdvancedDeduplication returns the EnableAdvancedDeduplication field if non-nil, zero value otherwise.

### GetEnableAdvancedDeduplicationOk

`func (o *StorageOptions) GetEnableAdvancedDeduplicationOk() (*bool, bool)`

GetEnableAdvancedDeduplicationOk returns a tuple with the EnableAdvancedDeduplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAdvancedDeduplication

`func (o *StorageOptions) SetEnableAdvancedDeduplication(v bool)`

SetEnableAdvancedDeduplication sets EnableAdvancedDeduplication field to given value.

### HasEnableAdvancedDeduplication

`func (o *StorageOptions) HasEnableAdvancedDeduplication() bool`

HasEnableAdvancedDeduplication returns a boolean if a field has been set.

### GetEnableDataReductionSharedVolume

`func (o *StorageOptions) GetEnableDataReductionSharedVolume() bool`

GetEnableDataReductionSharedVolume returns the EnableDataReductionSharedVolume field if non-nil, zero value otherwise.

### GetEnableDataReductionSharedVolumeOk

`func (o *StorageOptions) GetEnableDataReductionSharedVolumeOk() (*bool, bool)`

GetEnableDataReductionSharedVolumeOk returns a tuple with the EnableDataReductionSharedVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDataReductionSharedVolume

`func (o *StorageOptions) SetEnableDataReductionSharedVolume(v bool)`

SetEnableDataReductionSharedVolume sets EnableDataReductionSharedVolume field to given value.

### HasEnableDataReductionSharedVolume

`func (o *StorageOptions) HasEnableDataReductionSharedVolume() bool`

HasEnableDataReductionSharedVolume returns a boolean if a field has been set.

### GetVolumeName

`func (o *StorageOptions) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *StorageOptions) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *StorageOptions) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *StorageOptions) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.

### GetArrayId

`func (o *StorageOptions) GetArrayId() string`

GetArrayId returns the ArrayId field if non-nil, zero value otherwise.

### GetArrayIdOk

`func (o *StorageOptions) GetArrayIdOk() (*string, bool)`

GetArrayIdOk returns a tuple with the ArrayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayId

`func (o *StorageOptions) SetArrayId(v string)`

SetArrayId sets ArrayId field to given value.

### HasArrayId

`func (o *StorageOptions) HasArrayId() bool`

HasArrayId returns a boolean if a field has been set.

### GetDirectorId

`func (o *StorageOptions) GetDirectorId() string`

GetDirectorId returns the DirectorId field if non-nil, zero value otherwise.

### GetDirectorIdOk

`func (o *StorageOptions) GetDirectorIdOk() (*string, bool)`

GetDirectorIdOk returns a tuple with the DirectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectorId

`func (o *StorageOptions) SetDirectorId(v string)`

SetDirectorId sets DirectorId field to given value.

### HasDirectorId

`func (o *StorageOptions) HasDirectorId() bool`

HasDirectorId returns a boolean if a field has been set.

### GetS3Hostname

`func (o *StorageOptions) GetS3Hostname() string`

GetS3Hostname returns the S3Hostname field if non-nil, zero value otherwise.

### GetS3HostnameOk

`func (o *StorageOptions) GetS3HostnameOk() (*string, bool)`

GetS3HostnameOk returns a tuple with the S3Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Hostname

`func (o *StorageOptions) SetS3Hostname(v string)`

SetS3Hostname sets S3Hostname field to given value.

### HasS3Hostname

`func (o *StorageOptions) HasS3Hostname() bool`

HasS3Hostname returns a boolean if a field has been set.

### GetS3Port

`func (o *StorageOptions) GetS3Port() float32`

GetS3Port returns the S3Port field if non-nil, zero value otherwise.

### GetS3PortOk

`func (o *StorageOptions) GetS3PortOk() (*float32, bool)`

GetS3PortOk returns a tuple with the S3Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Port

`func (o *StorageOptions) SetS3Port(v float32)`

SetS3Port sets S3Port field to given value.

### HasS3Port

`func (o *StorageOptions) HasS3Port() bool`

HasS3Port returns a boolean if a field has been set.

### GetDefaultServiceLevel

`func (o *StorageOptions) GetDefaultServiceLevel() string`

GetDefaultServiceLevel returns the DefaultServiceLevel field if non-nil, zero value otherwise.

### GetDefaultServiceLevelOk

`func (o *StorageOptions) GetDefaultServiceLevelOk() (*string, bool)`

GetDefaultServiceLevelOk returns a tuple with the DefaultServiceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultServiceLevel

`func (o *StorageOptions) SetDefaultServiceLevel(v string)`

SetDefaultServiceLevel sets DefaultServiceLevel field to given value.

### HasDefaultServiceLevel

`func (o *StorageOptions) HasDefaultServiceLevel() bool`

HasDefaultServiceLevel returns a boolean if a field has been set.

### GetFibreChannelEnabled

`func (o *StorageOptions) GetFibreChannelEnabled() bool`

GetFibreChannelEnabled returns the FibreChannelEnabled field if non-nil, zero value otherwise.

### GetFibreChannelEnabledOk

`func (o *StorageOptions) GetFibreChannelEnabledOk() (*bool, bool)`

GetFibreChannelEnabledOk returns a tuple with the FibreChannelEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFibreChannelEnabled

`func (o *StorageOptions) SetFibreChannelEnabled(v bool)`

SetFibreChannelEnabled sets FibreChannelEnabled field to given value.

### HasFibreChannelEnabled

`func (o *StorageOptions) HasFibreChannelEnabled() bool`

HasFibreChannelEnabled returns a boolean if a field has been set.

### GetDirectors

`func (o *StorageOptions) GetDirectors() []string`

GetDirectors returns the Directors field if non-nil, zero value otherwise.

### GetDirectorsOk

`func (o *StorageOptions) GetDirectorsOk() (*[]string, bool)`

GetDirectorsOk returns a tuple with the Directors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectors

`func (o *StorageOptions) SetDirectors(v []string)`

SetDirectors sets Directors field to given value.

### HasDirectors

`func (o *StorageOptions) HasDirectors() bool`

HasDirectors returns a boolean if a field has been set.

### GetResourcePool

`func (o *StorageOptions) GetResourcePool() string`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *StorageOptions) GetResourcePoolOk() (*string, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *StorageOptions) SetResourcePool(v string)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *StorageOptions) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetVersion

`func (o *StorageOptions) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StorageOptions) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StorageOptions) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *StorageOptions) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetMajorVersion

`func (o *StorageOptions) GetMajorVersion() string`

GetMajorVersion returns the MajorVersion field if non-nil, zero value otherwise.

### GetMajorVersionOk

`func (o *StorageOptions) GetMajorVersionOk() (*string, bool)`

GetMajorVersionOk returns a tuple with the MajorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorVersion

`func (o *StorageOptions) SetMajorVersion(v string)`

SetMajorVersion sets MajorVersion field to given value.

### HasMajorVersion

`func (o *StorageOptions) HasMajorVersion() bool`

HasMajorVersion returns a boolean if a field has been set.

### GetArrays

`func (o *StorageOptions) GetArrays() []string`

GetArrays returns the Arrays field if non-nil, zero value otherwise.

### GetArraysOk

`func (o *StorageOptions) GetArraysOk() (*[]string, bool)`

GetArraysOk returns a tuple with the Arrays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrays

`func (o *StorageOptions) SetArrays(v []string)`

SetArrays sets Arrays field to given value.

### HasArrays

`func (o *StorageOptions) HasArrays() bool`

HasArrays returns a boolean if a field has been set.

### GetInfo

`func (o *StorageOptions) GetInfo() map[string]interface{}`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *StorageOptions) GetInfoOk() (*map[string]interface{}, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *StorageOptions) SetInfo(v map[string]interface{})`

SetInfo sets Info field to given value.

### HasInfo

`func (o *StorageOptions) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetServiceLevelNames

`func (o *StorageOptions) GetServiceLevelNames() []string`

GetServiceLevelNames returns the ServiceLevelNames field if non-nil, zero value otherwise.

### GetServiceLevelNamesOk

`func (o *StorageOptions) GetServiceLevelNamesOk() (*[]string, bool)`

GetServiceLevelNamesOk returns a tuple with the ServiceLevelNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevelNames

`func (o *StorageOptions) SetServiceLevelNames(v []string)`

SetServiceLevelNames sets ServiceLevelNames field to given value.

### HasServiceLevelNames

`func (o *StorageOptions) HasServiceLevelNames() bool`

HasServiceLevelNames returns a boolean if a field has been set.

### GetFibreChannelCapable

`func (o *StorageOptions) GetFibreChannelCapable() bool`

GetFibreChannelCapable returns the FibreChannelCapable field if non-nil, zero value otherwise.

### GetFibreChannelCapableOk

`func (o *StorageOptions) GetFibreChannelCapableOk() (*bool, bool)`

GetFibreChannelCapableOk returns a tuple with the FibreChannelCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFibreChannelCapable

`func (o *StorageOptions) SetFibreChannelCapable(v bool)`

SetFibreChannelCapable sets FibreChannelCapable field to given value.

### HasFibreChannelCapable

`func (o *StorageOptions) HasFibreChannelCapable() bool`

HasFibreChannelCapable returns a boolean if a field has been set.

### GetPortsIscsi

`func (o *StorageOptions) GetPortsIscsi() []StoragePort`

GetPortsIscsi returns the PortsIscsi field if non-nil, zero value otherwise.

### GetPortsIscsiOk

`func (o *StorageOptions) GetPortsIscsiOk() (*[]StoragePort, bool)`

GetPortsIscsiOk returns a tuple with the PortsIscsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortsIscsi

`func (o *StorageOptions) SetPortsIscsi(v []StoragePort)`

SetPortsIscsi sets PortsIscsi field to given value.

### HasPortsIscsi

`func (o *StorageOptions) HasPortsIscsi() bool`

HasPortsIscsi returns a boolean if a field has been set.

### GetPortsScsiFc

`func (o *StorageOptions) GetPortsScsiFc() []StoragePort`

GetPortsScsiFc returns the PortsScsiFc field if non-nil, zero value otherwise.

### GetPortsScsiFcOk

`func (o *StorageOptions) GetPortsScsiFcOk() (*[]StoragePort, bool)`

GetPortsScsiFcOk returns a tuple with the PortsScsiFc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortsScsiFc

`func (o *StorageOptions) SetPortsScsiFc(v []StoragePort)`

SetPortsScsiFc sets PortsScsiFc field to given value.

### HasPortsScsiFc

`func (o *StorageOptions) HasPortsScsiFc() bool`

HasPortsScsiFc returns a boolean if a field has been set.

### GetPortsNvmeTcp

`func (o *StorageOptions) GetPortsNvmeTcp() []StoragePort`

GetPortsNvmeTcp returns the PortsNvmeTcp field if non-nil, zero value otherwise.

### GetPortsNvmeTcpOk

`func (o *StorageOptions) GetPortsNvmeTcpOk() (*[]StoragePort, bool)`

GetPortsNvmeTcpOk returns a tuple with the PortsNvmeTcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortsNvmeTcp

`func (o *StorageOptions) SetPortsNvmeTcp(v []StoragePort)`

SetPortsNvmeTcp sets PortsNvmeTcp field to given value.

### HasPortsNvmeTcp

`func (o *StorageOptions) HasPortsNvmeTcp() bool`

HasPortsNvmeTcp returns a boolean if a field has been set.

### GetPortsNvmeFc

`func (o *StorageOptions) GetPortsNvmeFc() []StoragePort`

GetPortsNvmeFc returns the PortsNvmeFc field if non-nil, zero value otherwise.

### GetPortsNvmeFcOk

`func (o *StorageOptions) GetPortsNvmeFcOk() (*[]StoragePort, bool)`

GetPortsNvmeFcOk returns a tuple with the PortsNvmeFc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortsNvmeFc

`func (o *StorageOptions) SetPortsNvmeFc(v []StoragePort)`

SetPortsNvmeFc sets PortsNvmeFc field to given value.

### HasPortsNvmeFc

`func (o *StorageOptions) HasPortsNvmeFc() bool`

HasPortsNvmeFc returns a boolean if a field has been set.

### GetPortsToUse

`func (o *StorageOptions) GetPortsToUse() []StoragePort`

GetPortsToUse returns the PortsToUse field if non-nil, zero value otherwise.

### GetPortsToUseOk

`func (o *StorageOptions) GetPortsToUseOk() (*[]StoragePort, bool)`

GetPortsToUseOk returns a tuple with the PortsToUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortsToUse

`func (o *StorageOptions) SetPortsToUse(v []StoragePort)`

SetPortsToUse sets PortsToUse field to given value.

### HasPortsToUse

`func (o *StorageOptions) HasPortsToUse() bool`

HasPortsToUse returns a boolean if a field has been set.

### GetServers

`func (o *StorageOptions) GetServers() []StoragePort`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *StorageOptions) GetServersOk() (*[]StoragePort, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *StorageOptions) SetServers(v []StoragePort)`

SetServers sets Servers field to given value.

### HasServers

`func (o *StorageOptions) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetNodes

`func (o *StorageOptions) GetNodes() []StorageNodes`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *StorageOptions) GetNodesOk() (*[]StorageNodes, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *StorageOptions) SetNodes(v []StorageNodes)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *StorageOptions) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetInfoGatherError

`func (o *StorageOptions) GetInfoGatherError() string`

GetInfoGatherError returns the InfoGatherError field if non-nil, zero value otherwise.

### GetInfoGatherErrorOk

`func (o *StorageOptions) GetInfoGatherErrorOk() (*string, bool)`

GetInfoGatherErrorOk returns a tuple with the InfoGatherError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoGatherError

`func (o *StorageOptions) SetInfoGatherError(v string)`

SetInfoGatherError sets InfoGatherError field to given value.

### HasInfoGatherError

`func (o *StorageOptions) HasInfoGatherError() bool`

HasInfoGatherError returns a boolean if a field has been set.

### GetConfigureError

`func (o *StorageOptions) GetConfigureError() string`

GetConfigureError returns the ConfigureError field if non-nil, zero value otherwise.

### GetConfigureErrorOk

`func (o *StorageOptions) GetConfigureErrorOk() (*string, bool)`

GetConfigureErrorOk returns a tuple with the ConfigureError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureError

`func (o *StorageOptions) SetConfigureError(v string)`

SetConfigureError sets ConfigureError field to given value.

### HasConfigureError

`func (o *StorageOptions) HasConfigureError() bool`

HasConfigureError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


