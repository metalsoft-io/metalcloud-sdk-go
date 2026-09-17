# ContainerType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Container Type ID | 
**Name** | **string** | Name of the Container type | 
**DisplayName** | Pointer to **string** | Display name of the Container type | [optional] 
**Label** | Pointer to **string** | Label of the Container type | [optional] 
**CpuCores** | **float32** | Number of CPU cores for the Container type | 
**RamGB** | **float32** | RAM in GB for the Container type | 
**GpuInfo** | Pointer to [**[]ContainerTypeGPUInfo**](ContainerTypeGPUInfo.md) | Information about GPUs available for this Container Type | [optional] 
**IsExperimental** | Pointer to **float32** | Flag to indicate if the Container is experimental. 1 for true, 0 for false. Default is 0. | [optional] 
**Tags** | Pointer to **[]string** | Tags for the Container Type. | [optional] 
**ForUnmanagedContainersOnly** | Pointer to **float32** | Flag to indicate if the Container is for unmanaged Containers only. 1 for true, 0 for false. Default is 0. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewContainerType

`func NewContainerType(id int64, name string, cpuCores float32, ramGB float32, ) *ContainerType`

NewContainerType instantiates a new ContainerType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerTypeWithDefaults

`func NewContainerTypeWithDefaults() *ContainerType`

NewContainerTypeWithDefaults instantiates a new ContainerType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContainerType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContainerType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContainerType) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ContainerType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerType) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *ContainerType) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ContainerType) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ContainerType) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ContainerType) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLabel

`func (o *ContainerType) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ContainerType) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ContainerType) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ContainerType) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetCpuCores

`func (o *ContainerType) GetCpuCores() float32`

GetCpuCores returns the CpuCores field if non-nil, zero value otherwise.

### GetCpuCoresOk

`func (o *ContainerType) GetCpuCoresOk() (*float32, bool)`

GetCpuCoresOk returns a tuple with the CpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCores

`func (o *ContainerType) SetCpuCores(v float32)`

SetCpuCores sets CpuCores field to given value.


### GetRamGB

`func (o *ContainerType) GetRamGB() float32`

GetRamGB returns the RamGB field if non-nil, zero value otherwise.

### GetRamGBOk

`func (o *ContainerType) GetRamGBOk() (*float32, bool)`

GetRamGBOk returns a tuple with the RamGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGB

`func (o *ContainerType) SetRamGB(v float32)`

SetRamGB sets RamGB field to given value.


### GetGpuInfo

`func (o *ContainerType) GetGpuInfo() []ContainerTypeGPUInfo`

GetGpuInfo returns the GpuInfo field if non-nil, zero value otherwise.

### GetGpuInfoOk

`func (o *ContainerType) GetGpuInfoOk() (*[]ContainerTypeGPUInfo, bool)`

GetGpuInfoOk returns a tuple with the GpuInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuInfo

`func (o *ContainerType) SetGpuInfo(v []ContainerTypeGPUInfo)`

SetGpuInfo sets GpuInfo field to given value.

### HasGpuInfo

`func (o *ContainerType) HasGpuInfo() bool`

HasGpuInfo returns a boolean if a field has been set.

### GetIsExperimental

`func (o *ContainerType) GetIsExperimental() float32`

GetIsExperimental returns the IsExperimental field if non-nil, zero value otherwise.

### GetIsExperimentalOk

`func (o *ContainerType) GetIsExperimentalOk() (*float32, bool)`

GetIsExperimentalOk returns a tuple with the IsExperimental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExperimental

`func (o *ContainerType) SetIsExperimental(v float32)`

SetIsExperimental sets IsExperimental field to given value.

### HasIsExperimental

`func (o *ContainerType) HasIsExperimental() bool`

HasIsExperimental returns a boolean if a field has been set.

### GetTags

`func (o *ContainerType) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ContainerType) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ContainerType) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ContainerType) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetForUnmanagedContainersOnly

`func (o *ContainerType) GetForUnmanagedContainersOnly() float32`

GetForUnmanagedContainersOnly returns the ForUnmanagedContainersOnly field if non-nil, zero value otherwise.

### GetForUnmanagedContainersOnlyOk

`func (o *ContainerType) GetForUnmanagedContainersOnlyOk() (*float32, bool)`

GetForUnmanagedContainersOnlyOk returns a tuple with the ForUnmanagedContainersOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForUnmanagedContainersOnly

`func (o *ContainerType) SetForUnmanagedContainersOnly(v float32)`

SetForUnmanagedContainersOnly sets ForUnmanagedContainersOnly field to given value.

### HasForUnmanagedContainersOnly

`func (o *ContainerType) HasForUnmanagedContainersOnly() bool`

HasForUnmanagedContainersOnly returns a boolean if a field has been set.

### GetLinks

`func (o *ContainerType) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ContainerType) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ContainerType) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ContainerType) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


