# VMType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | VM Type ID | 
**Name** | **string** | Name of the VM Pool type | 
**DisplayName** | Pointer to **string** | Display name of the VM Pool type | [optional] 
**Label** | Pointer to **string** | Label of the VM Pool type | [optional] 
**CpuCores** | **float32** | Number of CPU cores for the VM Pool type | 
**RamGB** | **float32** | RAM in GB for the VM Pool type | 
**GpuInfo** | Pointer to [**[]VMTypeGPUInfo**](VMTypeGPUInfo.md) | Information about GPUs available for this VM Type | [optional] 
**IsExperimental** | Pointer to **float32** | Flag to indicate if the VM Pool is experimental. 1 for true, 0 for false. Default is 0. | [optional] 
**Tags** | Pointer to **[]string** | Tags for the VM Type. | [optional] 
**ForUnmanagedVMsOnly** | Pointer to **float32** | Flag to indicate if the VM Pool is for unmanaged VMs only. 1 for true, 0 for false. Default is 0. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewVMType

`func NewVMType(id float32, name string, cpuCores float32, ramGB float32, ) *VMType`

NewVMType instantiates a new VMType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMTypeWithDefaults

`func NewVMTypeWithDefaults() *VMType`

NewVMTypeWithDefaults instantiates a new VMType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VMType) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VMType) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VMType) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *VMType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VMType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VMType) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *VMType) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *VMType) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *VMType) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *VMType) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLabel

`func (o *VMType) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *VMType) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *VMType) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *VMType) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetCpuCores

`func (o *VMType) GetCpuCores() float32`

GetCpuCores returns the CpuCores field if non-nil, zero value otherwise.

### GetCpuCoresOk

`func (o *VMType) GetCpuCoresOk() (*float32, bool)`

GetCpuCoresOk returns a tuple with the CpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCores

`func (o *VMType) SetCpuCores(v float32)`

SetCpuCores sets CpuCores field to given value.


### GetRamGB

`func (o *VMType) GetRamGB() float32`

GetRamGB returns the RamGB field if non-nil, zero value otherwise.

### GetRamGBOk

`func (o *VMType) GetRamGBOk() (*float32, bool)`

GetRamGBOk returns a tuple with the RamGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGB

`func (o *VMType) SetRamGB(v float32)`

SetRamGB sets RamGB field to given value.


### GetGpuInfo

`func (o *VMType) GetGpuInfo() []VMTypeGPUInfo`

GetGpuInfo returns the GpuInfo field if non-nil, zero value otherwise.

### GetGpuInfoOk

`func (o *VMType) GetGpuInfoOk() (*[]VMTypeGPUInfo, bool)`

GetGpuInfoOk returns a tuple with the GpuInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuInfo

`func (o *VMType) SetGpuInfo(v []VMTypeGPUInfo)`

SetGpuInfo sets GpuInfo field to given value.

### HasGpuInfo

`func (o *VMType) HasGpuInfo() bool`

HasGpuInfo returns a boolean if a field has been set.

### GetIsExperimental

`func (o *VMType) GetIsExperimental() float32`

GetIsExperimental returns the IsExperimental field if non-nil, zero value otherwise.

### GetIsExperimentalOk

`func (o *VMType) GetIsExperimentalOk() (*float32, bool)`

GetIsExperimentalOk returns a tuple with the IsExperimental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExperimental

`func (o *VMType) SetIsExperimental(v float32)`

SetIsExperimental sets IsExperimental field to given value.

### HasIsExperimental

`func (o *VMType) HasIsExperimental() bool`

HasIsExperimental returns a boolean if a field has been set.

### GetTags

`func (o *VMType) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VMType) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VMType) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VMType) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetForUnmanagedVMsOnly

`func (o *VMType) GetForUnmanagedVMsOnly() float32`

GetForUnmanagedVMsOnly returns the ForUnmanagedVMsOnly field if non-nil, zero value otherwise.

### GetForUnmanagedVMsOnlyOk

`func (o *VMType) GetForUnmanagedVMsOnlyOk() (*float32, bool)`

GetForUnmanagedVMsOnlyOk returns a tuple with the ForUnmanagedVMsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForUnmanagedVMsOnly

`func (o *VMType) SetForUnmanagedVMsOnly(v float32)`

SetForUnmanagedVMsOnly sets ForUnmanagedVMsOnly field to given value.

### HasForUnmanagedVMsOnly

`func (o *VMType) HasForUnmanagedVMsOnly() bool`

HasForUnmanagedVMsOnly returns a boolean if a field has been set.

### GetLinks

`func (o *VMType) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VMType) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VMType) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *VMType) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


