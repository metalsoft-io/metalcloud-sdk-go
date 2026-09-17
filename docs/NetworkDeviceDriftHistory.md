# NetworkDeviceDriftHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The ID of the drift history | 
**NetworkDeviceId** | **int64** | The ID of the associated network device | 
**SnapshotId** | **string** | The identifier of the snapshot on which this drift was reported | 
**ConfigurationDrift** | **string** | The configuration drift for the network device.  | 
**CreatedTimestamp** | **time.Time** | Entity creation timestamp | 
**AcknowledgedBy** | Pointer to **int64** | The ID of the admin who acknowledged the drift | [optional] 
**AcknowledgeTimestamp** | Pointer to **time.Time** | Timestamp when the drift was acknowledge by an admin | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewNetworkDeviceDriftHistory

`func NewNetworkDeviceDriftHistory(id int64, networkDeviceId int64, snapshotId string, configurationDrift string, createdTimestamp time.Time, ) *NetworkDeviceDriftHistory`

NewNetworkDeviceDriftHistory instantiates a new NetworkDeviceDriftHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceDriftHistoryWithDefaults

`func NewNetworkDeviceDriftHistoryWithDefaults() *NetworkDeviceDriftHistory`

NewNetworkDeviceDriftHistoryWithDefaults instantiates a new NetworkDeviceDriftHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkDeviceDriftHistory) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkDeviceDriftHistory) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkDeviceDriftHistory) SetId(v int64)`

SetId sets Id field to given value.


### GetNetworkDeviceId

`func (o *NetworkDeviceDriftHistory) GetNetworkDeviceId() int64`

GetNetworkDeviceId returns the NetworkDeviceId field if non-nil, zero value otherwise.

### GetNetworkDeviceIdOk

`func (o *NetworkDeviceDriftHistory) GetNetworkDeviceIdOk() (*int64, bool)`

GetNetworkDeviceIdOk returns a tuple with the NetworkDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceId

`func (o *NetworkDeviceDriftHistory) SetNetworkDeviceId(v int64)`

SetNetworkDeviceId sets NetworkDeviceId field to given value.


### GetSnapshotId

`func (o *NetworkDeviceDriftHistory) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *NetworkDeviceDriftHistory) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *NetworkDeviceDriftHistory) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.


### GetConfigurationDrift

`func (o *NetworkDeviceDriftHistory) GetConfigurationDrift() string`

GetConfigurationDrift returns the ConfigurationDrift field if non-nil, zero value otherwise.

### GetConfigurationDriftOk

`func (o *NetworkDeviceDriftHistory) GetConfigurationDriftOk() (*string, bool)`

GetConfigurationDriftOk returns a tuple with the ConfigurationDrift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationDrift

`func (o *NetworkDeviceDriftHistory) SetConfigurationDrift(v string)`

SetConfigurationDrift sets ConfigurationDrift field to given value.


### GetCreatedTimestamp

`func (o *NetworkDeviceDriftHistory) GetCreatedTimestamp() time.Time`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *NetworkDeviceDriftHistory) GetCreatedTimestampOk() (*time.Time, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *NetworkDeviceDriftHistory) SetCreatedTimestamp(v time.Time)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetAcknowledgedBy

`func (o *NetworkDeviceDriftHistory) GetAcknowledgedBy() int64`

GetAcknowledgedBy returns the AcknowledgedBy field if non-nil, zero value otherwise.

### GetAcknowledgedByOk

`func (o *NetworkDeviceDriftHistory) GetAcknowledgedByOk() (*int64, bool)`

GetAcknowledgedByOk returns a tuple with the AcknowledgedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedBy

`func (o *NetworkDeviceDriftHistory) SetAcknowledgedBy(v int64)`

SetAcknowledgedBy sets AcknowledgedBy field to given value.

### HasAcknowledgedBy

`func (o *NetworkDeviceDriftHistory) HasAcknowledgedBy() bool`

HasAcknowledgedBy returns a boolean if a field has been set.

### GetAcknowledgeTimestamp

`func (o *NetworkDeviceDriftHistory) GetAcknowledgeTimestamp() time.Time`

GetAcknowledgeTimestamp returns the AcknowledgeTimestamp field if non-nil, zero value otherwise.

### GetAcknowledgeTimestampOk

`func (o *NetworkDeviceDriftHistory) GetAcknowledgeTimestampOk() (*time.Time, bool)`

GetAcknowledgeTimestampOk returns a tuple with the AcknowledgeTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgeTimestamp

`func (o *NetworkDeviceDriftHistory) SetAcknowledgeTimestamp(v time.Time)`

SetAcknowledgeTimestamp sets AcknowledgeTimestamp field to given value.

### HasAcknowledgeTimestamp

`func (o *NetworkDeviceDriftHistory) HasAcknowledgeTimestamp() bool`

HasAcknowledgeTimestamp returns a boolean if a field has been set.

### GetLinks

`func (o *NetworkDeviceDriftHistory) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NetworkDeviceDriftHistory) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NetworkDeviceDriftHistory) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NetworkDeviceDriftHistory) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


