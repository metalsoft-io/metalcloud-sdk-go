# StorageStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalSpaceGB** | **float32** | Total capacity in GB | 
**UsedSpaceGB** | **float32** | Total used space in GB | 
**FreeSpaceGB** | **float32** | Total free space in GB | 

## Methods

### NewStorageStatistics

`func NewStorageStatistics(totalSpaceGB float32, usedSpaceGB float32, freeSpaceGB float32, ) *StorageStatistics`

NewStorageStatistics instantiates a new StorageStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageStatisticsWithDefaults

`func NewStorageStatisticsWithDefaults() *StorageStatistics`

NewStorageStatisticsWithDefaults instantiates a new StorageStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalSpaceGB

`func (o *StorageStatistics) GetTotalSpaceGB() float32`

GetTotalSpaceGB returns the TotalSpaceGB field if non-nil, zero value otherwise.

### GetTotalSpaceGBOk

`func (o *StorageStatistics) GetTotalSpaceGBOk() (*float32, bool)`

GetTotalSpaceGBOk returns a tuple with the TotalSpaceGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSpaceGB

`func (o *StorageStatistics) SetTotalSpaceGB(v float32)`

SetTotalSpaceGB sets TotalSpaceGB field to given value.


### GetUsedSpaceGB

`func (o *StorageStatistics) GetUsedSpaceGB() float32`

GetUsedSpaceGB returns the UsedSpaceGB field if non-nil, zero value otherwise.

### GetUsedSpaceGBOk

`func (o *StorageStatistics) GetUsedSpaceGBOk() (*float32, bool)`

GetUsedSpaceGBOk returns a tuple with the UsedSpaceGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSpaceGB

`func (o *StorageStatistics) SetUsedSpaceGB(v float32)`

SetUsedSpaceGB sets UsedSpaceGB field to given value.


### GetFreeSpaceGB

`func (o *StorageStatistics) GetFreeSpaceGB() float32`

GetFreeSpaceGB returns the FreeSpaceGB field if non-nil, zero value otherwise.

### GetFreeSpaceGBOk

`func (o *StorageStatistics) GetFreeSpaceGBOk() (*float32, bool)`

GetFreeSpaceGBOk returns a tuple with the FreeSpaceGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpaceGB

`func (o *StorageStatistics) SetFreeSpaceGB(v float32)`

SetFreeSpaceGB sets FreeSpaceGB field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


