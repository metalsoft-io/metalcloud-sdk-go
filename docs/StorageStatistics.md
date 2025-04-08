# StorageStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaintenanceCount** | **float32** | Number of storages in maintenance | 
**ExperimentalCount** | **float32** | Number of experimental storages | 
**LowSpaceCount** | **float32** | Number of storages low on space | 
**UsedSpace** | **float32** | Total used space across all storages | 
**FreeSpace** | **float32** | Total free space across all storages | 
**Types** | **map[string]interface{}** | Number of storages by type | 
**PendingCount** | **float32** | Number of storages in pending state | 
**ReadyCount** | **float32** | Number of storages in ready state | 
**ActiveCount** | **float32** | Number of storages in active state | 

## Methods

### NewStorageStatistics

`func NewStorageStatistics(maintenanceCount float32, experimentalCount float32, lowSpaceCount float32, usedSpace float32, freeSpace float32, types map[string]interface{}, pendingCount float32, readyCount float32, activeCount float32, ) *StorageStatistics`

NewStorageStatistics instantiates a new StorageStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageStatisticsWithDefaults

`func NewStorageStatisticsWithDefaults() *StorageStatistics`

NewStorageStatisticsWithDefaults instantiates a new StorageStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaintenanceCount

`func (o *StorageStatistics) GetMaintenanceCount() float32`

GetMaintenanceCount returns the MaintenanceCount field if non-nil, zero value otherwise.

### GetMaintenanceCountOk

`func (o *StorageStatistics) GetMaintenanceCountOk() (*float32, bool)`

GetMaintenanceCountOk returns a tuple with the MaintenanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceCount

`func (o *StorageStatistics) SetMaintenanceCount(v float32)`

SetMaintenanceCount sets MaintenanceCount field to given value.


### GetExperimentalCount

`func (o *StorageStatistics) GetExperimentalCount() float32`

GetExperimentalCount returns the ExperimentalCount field if non-nil, zero value otherwise.

### GetExperimentalCountOk

`func (o *StorageStatistics) GetExperimentalCountOk() (*float32, bool)`

GetExperimentalCountOk returns a tuple with the ExperimentalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentalCount

`func (o *StorageStatistics) SetExperimentalCount(v float32)`

SetExperimentalCount sets ExperimentalCount field to given value.


### GetLowSpaceCount

`func (o *StorageStatistics) GetLowSpaceCount() float32`

GetLowSpaceCount returns the LowSpaceCount field if non-nil, zero value otherwise.

### GetLowSpaceCountOk

`func (o *StorageStatistics) GetLowSpaceCountOk() (*float32, bool)`

GetLowSpaceCountOk returns a tuple with the LowSpaceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowSpaceCount

`func (o *StorageStatistics) SetLowSpaceCount(v float32)`

SetLowSpaceCount sets LowSpaceCount field to given value.


### GetUsedSpace

`func (o *StorageStatistics) GetUsedSpace() float32`

GetUsedSpace returns the UsedSpace field if non-nil, zero value otherwise.

### GetUsedSpaceOk

`func (o *StorageStatistics) GetUsedSpaceOk() (*float32, bool)`

GetUsedSpaceOk returns a tuple with the UsedSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSpace

`func (o *StorageStatistics) SetUsedSpace(v float32)`

SetUsedSpace sets UsedSpace field to given value.


### GetFreeSpace

`func (o *StorageStatistics) GetFreeSpace() float32`

GetFreeSpace returns the FreeSpace field if non-nil, zero value otherwise.

### GetFreeSpaceOk

`func (o *StorageStatistics) GetFreeSpaceOk() (*float32, bool)`

GetFreeSpaceOk returns a tuple with the FreeSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpace

`func (o *StorageStatistics) SetFreeSpace(v float32)`

SetFreeSpace sets FreeSpace field to given value.


### GetTypes

`func (o *StorageStatistics) GetTypes() map[string]interface{}`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *StorageStatistics) GetTypesOk() (*map[string]interface{}, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *StorageStatistics) SetTypes(v map[string]interface{})`

SetTypes sets Types field to given value.


### GetPendingCount

`func (o *StorageStatistics) GetPendingCount() float32`

GetPendingCount returns the PendingCount field if non-nil, zero value otherwise.

### GetPendingCountOk

`func (o *StorageStatistics) GetPendingCountOk() (*float32, bool)`

GetPendingCountOk returns a tuple with the PendingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCount

`func (o *StorageStatistics) SetPendingCount(v float32)`

SetPendingCount sets PendingCount field to given value.


### GetReadyCount

`func (o *StorageStatistics) GetReadyCount() float32`

GetReadyCount returns the ReadyCount field if non-nil, zero value otherwise.

### GetReadyCountOk

`func (o *StorageStatistics) GetReadyCountOk() (*float32, bool)`

GetReadyCountOk returns a tuple with the ReadyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyCount

`func (o *StorageStatistics) SetReadyCount(v float32)`

SetReadyCount sets ReadyCount field to given value.


### GetActiveCount

`func (o *StorageStatistics) GetActiveCount() float32`

GetActiveCount returns the ActiveCount field if non-nil, zero value otherwise.

### GetActiveCountOk

`func (o *StorageStatistics) GetActiveCountOk() (*float32, bool)`

GetActiveCountOk returns a tuple with the ActiveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveCount

`func (o *StorageStatistics) SetActiveCount(v float32)`

SetActiveCount sets ActiveCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


