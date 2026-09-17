# CreateContainerType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the Container type | 
**DisplayName** | Pointer to **string** | Display name of the Container type | [optional] 
**Label** | Pointer to **string** | Label of the Container type | [optional] 
**CpuCores** | **float32** | Number of CPU cores for the Container type | 
**RamGB** | **float32** | RAM in GB for the Container type | 
**IsExperimental** | Pointer to **float32** | Flag to indicate if the Container is experimental. 1 for true, 0 for false. Default is 0. | [optional] 
**Tags** | Pointer to **[]string** | Tags for the Container Type. | [optional] 
**ForUnmanagedContainersOnly** | Pointer to **float32** | Flag to indicate if the Container is for unmanaged Containers only. 1 for true, 0 for false. Default is 0. | [optional] 
**GpuInfo** | Pointer to [**[]ContainerTypeGPUInfo**](ContainerTypeGPUInfo.md) | Information about GPUs available for this Container Type | [optional] 

## Methods

### NewCreateContainerType

`func NewCreateContainerType(name string, cpuCores float32, ramGB float32, ) *CreateContainerType`

NewCreateContainerType instantiates a new CreateContainerType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContainerTypeWithDefaults

`func NewCreateContainerTypeWithDefaults() *CreateContainerType`

NewCreateContainerTypeWithDefaults instantiates a new CreateContainerType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateContainerType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateContainerType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateContainerType) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *CreateContainerType) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateContainerType) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateContainerType) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateContainerType) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLabel

`func (o *CreateContainerType) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateContainerType) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateContainerType) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateContainerType) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetCpuCores

`func (o *CreateContainerType) GetCpuCores() float32`

GetCpuCores returns the CpuCores field if non-nil, zero value otherwise.

### GetCpuCoresOk

`func (o *CreateContainerType) GetCpuCoresOk() (*float32, bool)`

GetCpuCoresOk returns a tuple with the CpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCores

`func (o *CreateContainerType) SetCpuCores(v float32)`

SetCpuCores sets CpuCores field to given value.


### GetRamGB

`func (o *CreateContainerType) GetRamGB() float32`

GetRamGB returns the RamGB field if non-nil, zero value otherwise.

### GetRamGBOk

`func (o *CreateContainerType) GetRamGBOk() (*float32, bool)`

GetRamGBOk returns a tuple with the RamGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGB

`func (o *CreateContainerType) SetRamGB(v float32)`

SetRamGB sets RamGB field to given value.


### GetIsExperimental

`func (o *CreateContainerType) GetIsExperimental() float32`

GetIsExperimental returns the IsExperimental field if non-nil, zero value otherwise.

### GetIsExperimentalOk

`func (o *CreateContainerType) GetIsExperimentalOk() (*float32, bool)`

GetIsExperimentalOk returns a tuple with the IsExperimental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExperimental

`func (o *CreateContainerType) SetIsExperimental(v float32)`

SetIsExperimental sets IsExperimental field to given value.

### HasIsExperimental

`func (o *CreateContainerType) HasIsExperimental() bool`

HasIsExperimental returns a boolean if a field has been set.

### GetTags

`func (o *CreateContainerType) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateContainerType) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateContainerType) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateContainerType) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetForUnmanagedContainersOnly

`func (o *CreateContainerType) GetForUnmanagedContainersOnly() float32`

GetForUnmanagedContainersOnly returns the ForUnmanagedContainersOnly field if non-nil, zero value otherwise.

### GetForUnmanagedContainersOnlyOk

`func (o *CreateContainerType) GetForUnmanagedContainersOnlyOk() (*float32, bool)`

GetForUnmanagedContainersOnlyOk returns a tuple with the ForUnmanagedContainersOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForUnmanagedContainersOnly

`func (o *CreateContainerType) SetForUnmanagedContainersOnly(v float32)`

SetForUnmanagedContainersOnly sets ForUnmanagedContainersOnly field to given value.

### HasForUnmanagedContainersOnly

`func (o *CreateContainerType) HasForUnmanagedContainersOnly() bool`

HasForUnmanagedContainersOnly returns a boolean if a field has been set.

### GetGpuInfo

`func (o *CreateContainerType) GetGpuInfo() []ContainerTypeGPUInfo`

GetGpuInfo returns the GpuInfo field if non-nil, zero value otherwise.

### GetGpuInfoOk

`func (o *CreateContainerType) GetGpuInfoOk() (*[]ContainerTypeGPUInfo, bool)`

GetGpuInfoOk returns a tuple with the GpuInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuInfo

`func (o *CreateContainerType) SetGpuInfo(v []ContainerTypeGPUInfo)`

SetGpuInfo sets GpuInfo field to given value.

### HasGpuInfo

`func (o *CreateContainerType) HasGpuInfo() bool`

HasGpuInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


