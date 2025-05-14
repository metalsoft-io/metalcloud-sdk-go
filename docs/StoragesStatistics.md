# StoragesStatistics

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

### NewStoragesStatistics

`func NewStoragesStatistics(maintenanceCount float32, experimentalCount float32, lowSpaceCount float32, usedSpace float32, freeSpace float32, types map[string]interface{}, pendingCount float32, readyCount float32, activeCount float32, ) *StoragesStatistics`

NewStoragesStatistics instantiates a new StoragesStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragesStatisticsWithDefaults

`func NewStoragesStatisticsWithDefaults() *StoragesStatistics`

NewStoragesStatisticsWithDefaults instantiates a new StoragesStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaintenanceCount

`func (o *StoragesStatistics) GetMaintenanceCount() float32`

GetMaintenanceCount returns the MaintenanceCount field if non-nil, zero value otherwise.

### GetMaintenanceCountOk

`func (o *StoragesStatistics) GetMaintenanceCountOk() (*float32, bool)`

GetMaintenanceCountOk returns a tuple with the MaintenanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceCount

`func (o *StoragesStatistics) SetMaintenanceCount(v float32)`

SetMaintenanceCount sets MaintenanceCount field to given value.


### GetExperimentalCount

`func (o *StoragesStatistics) GetExperimentalCount() float32`

GetExperimentalCount returns the ExperimentalCount field if non-nil, zero value otherwise.

### GetExperimentalCountOk

`func (o *StoragesStatistics) GetExperimentalCountOk() (*float32, bool)`

GetExperimentalCountOk returns a tuple with the ExperimentalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentalCount

`func (o *StoragesStatistics) SetExperimentalCount(v float32)`

SetExperimentalCount sets ExperimentalCount field to given value.


### GetLowSpaceCount

`func (o *StoragesStatistics) GetLowSpaceCount() float32`

GetLowSpaceCount returns the LowSpaceCount field if non-nil, zero value otherwise.

### GetLowSpaceCountOk

`func (o *StoragesStatistics) GetLowSpaceCountOk() (*float32, bool)`

GetLowSpaceCountOk returns a tuple with the LowSpaceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowSpaceCount

`func (o *StoragesStatistics) SetLowSpaceCount(v float32)`

SetLowSpaceCount sets LowSpaceCount field to given value.


### GetUsedSpace

`func (o *StoragesStatistics) GetUsedSpace() float32`

GetUsedSpace returns the UsedSpace field if non-nil, zero value otherwise.

### GetUsedSpaceOk

`func (o *StoragesStatistics) GetUsedSpaceOk() (*float32, bool)`

GetUsedSpaceOk returns a tuple with the UsedSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSpace

`func (o *StoragesStatistics) SetUsedSpace(v float32)`

SetUsedSpace sets UsedSpace field to given value.


### GetFreeSpace

`func (o *StoragesStatistics) GetFreeSpace() float32`

GetFreeSpace returns the FreeSpace field if non-nil, zero value otherwise.

### GetFreeSpaceOk

`func (o *StoragesStatistics) GetFreeSpaceOk() (*float32, bool)`

GetFreeSpaceOk returns a tuple with the FreeSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpace

`func (o *StoragesStatistics) SetFreeSpace(v float32)`

SetFreeSpace sets FreeSpace field to given value.


### GetTypes

`func (o *StoragesStatistics) GetTypes() map[string]interface{}`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *StoragesStatistics) GetTypesOk() (*map[string]interface{}, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *StoragesStatistics) SetTypes(v map[string]interface{})`

SetTypes sets Types field to given value.


### GetPendingCount

`func (o *StoragesStatistics) GetPendingCount() float32`

GetPendingCount returns the PendingCount field if non-nil, zero value otherwise.

### GetPendingCountOk

`func (o *StoragesStatistics) GetPendingCountOk() (*float32, bool)`

GetPendingCountOk returns a tuple with the PendingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCount

`func (o *StoragesStatistics) SetPendingCount(v float32)`

SetPendingCount sets PendingCount field to given value.


### GetReadyCount

`func (o *StoragesStatistics) GetReadyCount() float32`

GetReadyCount returns the ReadyCount field if non-nil, zero value otherwise.

### GetReadyCountOk

`func (o *StoragesStatistics) GetReadyCountOk() (*float32, bool)`

GetReadyCountOk returns a tuple with the ReadyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyCount

`func (o *StoragesStatistics) SetReadyCount(v float32)`

SetReadyCount sets ReadyCount field to given value.


### GetActiveCount

`func (o *StoragesStatistics) GetActiveCount() float32`

GetActiveCount returns the ActiveCount field if non-nil, zero value otherwise.

### GetActiveCountOk

`func (o *StoragesStatistics) GetActiveCountOk() (*float32, bool)`

GetActiveCountOk returns a tuple with the ActiveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveCount

`func (o *StoragesStatistics) SetActiveCount(v float32)`

SetActiveCount sets ActiveCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


