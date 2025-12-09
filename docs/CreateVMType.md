# CreateVMType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the VM Pool type | 
**DisplayName** | Pointer to **string** | Display name of the VM Pool type | [optional] 
**Label** | Pointer to **string** | Label of the VM Pool type | [optional] 
**CpuCores** | **float32** | Number of CPU cores for the VM Pool type | 
**RamGB** | **float32** | RAM in GB for the VM Pool type | 
**IsExperimental** | Pointer to **float32** | Flag to indicate if the VM Pool is experimental. 1 for true, 0 for false. Default is 0. | [optional] 
**Tags** | Pointer to **[]string** | Tags for the VM Type. | [optional] 
**ForUnmanagedVMsOnly** | Pointer to **float32** | Flag to indicate if the VM Pool is for unmanaged VMs only. 1 for true, 0 for false. Default is 0. | [optional] 
**GpuInfo** | Pointer to [**[]VMTypeGPUInfo**](VMTypeGPUInfo.md) | Information about GPUs available for this VM Type | [optional] 

## Methods

### NewCreateVMType

`func NewCreateVMType(name string, cpuCores float32, ramGB float32, ) *CreateVMType`

NewCreateVMType instantiates a new CreateVMType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVMTypeWithDefaults

`func NewCreateVMTypeWithDefaults() *CreateVMType`

NewCreateVMTypeWithDefaults instantiates a new CreateVMType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateVMType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateVMType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateVMType) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *CreateVMType) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateVMType) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateVMType) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateVMType) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLabel

`func (o *CreateVMType) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateVMType) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateVMType) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateVMType) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetCpuCores

`func (o *CreateVMType) GetCpuCores() float32`

GetCpuCores returns the CpuCores field if non-nil, zero value otherwise.

### GetCpuCoresOk

`func (o *CreateVMType) GetCpuCoresOk() (*float32, bool)`

GetCpuCoresOk returns a tuple with the CpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCores

`func (o *CreateVMType) SetCpuCores(v float32)`

SetCpuCores sets CpuCores field to given value.


### GetRamGB

`func (o *CreateVMType) GetRamGB() float32`

GetRamGB returns the RamGB field if non-nil, zero value otherwise.

### GetRamGBOk

`func (o *CreateVMType) GetRamGBOk() (*float32, bool)`

GetRamGBOk returns a tuple with the RamGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGB

`func (o *CreateVMType) SetRamGB(v float32)`

SetRamGB sets RamGB field to given value.


### GetIsExperimental

`func (o *CreateVMType) GetIsExperimental() float32`

GetIsExperimental returns the IsExperimental field if non-nil, zero value otherwise.

### GetIsExperimentalOk

`func (o *CreateVMType) GetIsExperimentalOk() (*float32, bool)`

GetIsExperimentalOk returns a tuple with the IsExperimental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExperimental

`func (o *CreateVMType) SetIsExperimental(v float32)`

SetIsExperimental sets IsExperimental field to given value.

### HasIsExperimental

`func (o *CreateVMType) HasIsExperimental() bool`

HasIsExperimental returns a boolean if a field has been set.

### GetTags

`func (o *CreateVMType) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateVMType) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateVMType) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateVMType) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetForUnmanagedVMsOnly

`func (o *CreateVMType) GetForUnmanagedVMsOnly() float32`

GetForUnmanagedVMsOnly returns the ForUnmanagedVMsOnly field if non-nil, zero value otherwise.

### GetForUnmanagedVMsOnlyOk

`func (o *CreateVMType) GetForUnmanagedVMsOnlyOk() (*float32, bool)`

GetForUnmanagedVMsOnlyOk returns a tuple with the ForUnmanagedVMsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForUnmanagedVMsOnly

`func (o *CreateVMType) SetForUnmanagedVMsOnly(v float32)`

SetForUnmanagedVMsOnly sets ForUnmanagedVMsOnly field to given value.

### HasForUnmanagedVMsOnly

`func (o *CreateVMType) HasForUnmanagedVMsOnly() bool`

HasForUnmanagedVMsOnly returns a boolean if a field has been set.

### GetGpuInfo

`func (o *CreateVMType) GetGpuInfo() []VMTypeGPUInfo`

GetGpuInfo returns the GpuInfo field if non-nil, zero value otherwise.

### GetGpuInfoOk

`func (o *CreateVMType) GetGpuInfoOk() (*[]VMTypeGPUInfo, bool)`

GetGpuInfoOk returns a tuple with the GpuInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuInfo

`func (o *CreateVMType) SetGpuInfo(v []VMTypeGPUInfo)`

SetGpuInfo sets GpuInfo field to given value.

### HasGpuInfo

`func (o *CreateVMType) HasGpuInfo() bool`

HasGpuInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


