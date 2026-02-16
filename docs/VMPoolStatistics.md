# VMPoolStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalRamGB** | **float32** | Total RAM in GB | 
**FreeRamGB** | **float32** | Free RAM in GB | 
**UsedRamGB** | **float32** | Used RAM in GB | 
**TotalSpaceGB** | **float32** | Total disk space in GB | 
**UsedSpaceGB** | **float32** | Used disk space in GB | 
**FreeSpaceGB** | **float32** | Free disk space in GB | 
**GpuInfo** | Pointer to [**[]VMTypeGPUInfo**](VMTypeGPUInfo.md) | GPU information | [optional] 

## Methods

### NewVMPoolStatistics

`func NewVMPoolStatistics(totalRamGB float32, freeRamGB float32, usedRamGB float32, totalSpaceGB float32, usedSpaceGB float32, freeSpaceGB float32, ) *VMPoolStatistics`

NewVMPoolStatistics instantiates a new VMPoolStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMPoolStatisticsWithDefaults

`func NewVMPoolStatisticsWithDefaults() *VMPoolStatistics`

NewVMPoolStatisticsWithDefaults instantiates a new VMPoolStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalRamGB

`func (o *VMPoolStatistics) GetTotalRamGB() float32`

GetTotalRamGB returns the TotalRamGB field if non-nil, zero value otherwise.

### GetTotalRamGBOk

`func (o *VMPoolStatistics) GetTotalRamGBOk() (*float32, bool)`

GetTotalRamGBOk returns a tuple with the TotalRamGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRamGB

`func (o *VMPoolStatistics) SetTotalRamGB(v float32)`

SetTotalRamGB sets TotalRamGB field to given value.


### GetFreeRamGB

`func (o *VMPoolStatistics) GetFreeRamGB() float32`

GetFreeRamGB returns the FreeRamGB field if non-nil, zero value otherwise.

### GetFreeRamGBOk

`func (o *VMPoolStatistics) GetFreeRamGBOk() (*float32, bool)`

GetFreeRamGBOk returns a tuple with the FreeRamGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeRamGB

`func (o *VMPoolStatistics) SetFreeRamGB(v float32)`

SetFreeRamGB sets FreeRamGB field to given value.


### GetUsedRamGB

`func (o *VMPoolStatistics) GetUsedRamGB() float32`

GetUsedRamGB returns the UsedRamGB field if non-nil, zero value otherwise.

### GetUsedRamGBOk

`func (o *VMPoolStatistics) GetUsedRamGBOk() (*float32, bool)`

GetUsedRamGBOk returns a tuple with the UsedRamGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedRamGB

`func (o *VMPoolStatistics) SetUsedRamGB(v float32)`

SetUsedRamGB sets UsedRamGB field to given value.


### GetTotalSpaceGB

`func (o *VMPoolStatistics) GetTotalSpaceGB() float32`

GetTotalSpaceGB returns the TotalSpaceGB field if non-nil, zero value otherwise.

### GetTotalSpaceGBOk

`func (o *VMPoolStatistics) GetTotalSpaceGBOk() (*float32, bool)`

GetTotalSpaceGBOk returns a tuple with the TotalSpaceGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSpaceGB

`func (o *VMPoolStatistics) SetTotalSpaceGB(v float32)`

SetTotalSpaceGB sets TotalSpaceGB field to given value.


### GetUsedSpaceGB

`func (o *VMPoolStatistics) GetUsedSpaceGB() float32`

GetUsedSpaceGB returns the UsedSpaceGB field if non-nil, zero value otherwise.

### GetUsedSpaceGBOk

`func (o *VMPoolStatistics) GetUsedSpaceGBOk() (*float32, bool)`

GetUsedSpaceGBOk returns a tuple with the UsedSpaceGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSpaceGB

`func (o *VMPoolStatistics) SetUsedSpaceGB(v float32)`

SetUsedSpaceGB sets UsedSpaceGB field to given value.


### GetFreeSpaceGB

`func (o *VMPoolStatistics) GetFreeSpaceGB() float32`

GetFreeSpaceGB returns the FreeSpaceGB field if non-nil, zero value otherwise.

### GetFreeSpaceGBOk

`func (o *VMPoolStatistics) GetFreeSpaceGBOk() (*float32, bool)`

GetFreeSpaceGBOk returns a tuple with the FreeSpaceGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpaceGB

`func (o *VMPoolStatistics) SetFreeSpaceGB(v float32)`

SetFreeSpaceGB sets FreeSpaceGB field to given value.


### GetGpuInfo

`func (o *VMPoolStatistics) GetGpuInfo() []VMTypeGPUInfo`

GetGpuInfo returns the GpuInfo field if non-nil, zero value otherwise.

### GetGpuInfoOk

`func (o *VMPoolStatistics) GetGpuInfoOk() (*[]VMTypeGPUInfo, bool)`

GetGpuInfoOk returns a tuple with the GpuInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuInfo

`func (o *VMPoolStatistics) SetGpuInfo(v []VMTypeGPUInfo)`

SetGpuInfo sets GpuInfo field to given value.

### HasGpuInfo

`func (o *VMPoolStatistics) HasGpuInfo() bool`

HasGpuInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


