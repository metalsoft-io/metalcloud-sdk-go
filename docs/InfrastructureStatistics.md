# InfrastructureStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobGroupStatistics** | Pointer to [**JobGroupStatistics**](JobGroupStatistics.md) | Statistics of job group for the infrastructure | [optional] 
**ServerTypesForUsage** | [**[]ServerTypesForUsage**](ServerTypesForUsage.md) | List of server types statistics | 
**VmPoolsForUsage** | [**[]VMPoolForUsage**](VMPoolForUsage.md) | List of VM Pools statistics | 
**StoragePoolsForUsage** | [**[]StoragePoolForUsage**](StoragePoolForUsage.md) | List of Storage Pools statistics | 

## Methods

### NewInfrastructureStatistics

`func NewInfrastructureStatistics(serverTypesForUsage []ServerTypesForUsage, vmPoolsForUsage []VMPoolForUsage, storagePoolsForUsage []StoragePoolForUsage, ) *InfrastructureStatistics`

NewInfrastructureStatistics instantiates a new InfrastructureStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureStatisticsWithDefaults

`func NewInfrastructureStatisticsWithDefaults() *InfrastructureStatistics`

NewInfrastructureStatisticsWithDefaults instantiates a new InfrastructureStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobGroupStatistics

`func (o *InfrastructureStatistics) GetJobGroupStatistics() JobGroupStatistics`

GetJobGroupStatistics returns the JobGroupStatistics field if non-nil, zero value otherwise.

### GetJobGroupStatisticsOk

`func (o *InfrastructureStatistics) GetJobGroupStatisticsOk() (*JobGroupStatistics, bool)`

GetJobGroupStatisticsOk returns a tuple with the JobGroupStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobGroupStatistics

`func (o *InfrastructureStatistics) SetJobGroupStatistics(v JobGroupStatistics)`

SetJobGroupStatistics sets JobGroupStatistics field to given value.

### HasJobGroupStatistics

`func (o *InfrastructureStatistics) HasJobGroupStatistics() bool`

HasJobGroupStatistics returns a boolean if a field has been set.

### GetServerTypesForUsage

`func (o *InfrastructureStatistics) GetServerTypesForUsage() []ServerTypesForUsage`

GetServerTypesForUsage returns the ServerTypesForUsage field if non-nil, zero value otherwise.

### GetServerTypesForUsageOk

`func (o *InfrastructureStatistics) GetServerTypesForUsageOk() (*[]ServerTypesForUsage, bool)`

GetServerTypesForUsageOk returns a tuple with the ServerTypesForUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypesForUsage

`func (o *InfrastructureStatistics) SetServerTypesForUsage(v []ServerTypesForUsage)`

SetServerTypesForUsage sets ServerTypesForUsage field to given value.


### GetVmPoolsForUsage

`func (o *InfrastructureStatistics) GetVmPoolsForUsage() []VMPoolForUsage`

GetVmPoolsForUsage returns the VmPoolsForUsage field if non-nil, zero value otherwise.

### GetVmPoolsForUsageOk

`func (o *InfrastructureStatistics) GetVmPoolsForUsageOk() (*[]VMPoolForUsage, bool)`

GetVmPoolsForUsageOk returns a tuple with the VmPoolsForUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmPoolsForUsage

`func (o *InfrastructureStatistics) SetVmPoolsForUsage(v []VMPoolForUsage)`

SetVmPoolsForUsage sets VmPoolsForUsage field to given value.


### GetStoragePoolsForUsage

`func (o *InfrastructureStatistics) GetStoragePoolsForUsage() []StoragePoolForUsage`

GetStoragePoolsForUsage returns the StoragePoolsForUsage field if non-nil, zero value otherwise.

### GetStoragePoolsForUsageOk

`func (o *InfrastructureStatistics) GetStoragePoolsForUsageOk() (*[]StoragePoolForUsage, bool)`

GetStoragePoolsForUsageOk returns a tuple with the StoragePoolsForUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePoolsForUsage

`func (o *InfrastructureStatistics) SetStoragePoolsForUsage(v []StoragePoolForUsage)`

SetStoragePoolsForUsage sets StoragePoolsForUsage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


