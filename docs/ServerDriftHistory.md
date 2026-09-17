# ServerDriftHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The ID of the drift history | 
**ServerId** | **int64** | The ID of the associated server | 
**SnapshotId** | **string** | The identifier of the snapshot on which this drift was reported | 
**ConfigurationDrift** | **string** | What changed on the server: the unified diff between the previously accepted snapshot and this one. | 
**CreatedTimestamp** | **time.Time** | Entity creation timestamp | 
**AcknowledgedBy** | Pointer to **int64** | The ID of the admin who acknowledged the drift | [optional] 
**AcknowledgeTimestamp** | Pointer to **time.Time** | Timestamp when the drift was acknowledge by an admin | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewServerDriftHistory

`func NewServerDriftHistory(id int64, serverId int64, snapshotId string, configurationDrift string, createdTimestamp time.Time, ) *ServerDriftHistory`

NewServerDriftHistory instantiates a new ServerDriftHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerDriftHistoryWithDefaults

`func NewServerDriftHistoryWithDefaults() *ServerDriftHistory`

NewServerDriftHistoryWithDefaults instantiates a new ServerDriftHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerDriftHistory) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerDriftHistory) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerDriftHistory) SetId(v int64)`

SetId sets Id field to given value.


### GetServerId

`func (o *ServerDriftHistory) GetServerId() int64`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ServerDriftHistory) GetServerIdOk() (*int64, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ServerDriftHistory) SetServerId(v int64)`

SetServerId sets ServerId field to given value.


### GetSnapshotId

`func (o *ServerDriftHistory) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *ServerDriftHistory) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *ServerDriftHistory) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.


### GetConfigurationDrift

`func (o *ServerDriftHistory) GetConfigurationDrift() string`

GetConfigurationDrift returns the ConfigurationDrift field if non-nil, zero value otherwise.

### GetConfigurationDriftOk

`func (o *ServerDriftHistory) GetConfigurationDriftOk() (*string, bool)`

GetConfigurationDriftOk returns a tuple with the ConfigurationDrift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationDrift

`func (o *ServerDriftHistory) SetConfigurationDrift(v string)`

SetConfigurationDrift sets ConfigurationDrift field to given value.


### GetCreatedTimestamp

`func (o *ServerDriftHistory) GetCreatedTimestamp() time.Time`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *ServerDriftHistory) GetCreatedTimestampOk() (*time.Time, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *ServerDriftHistory) SetCreatedTimestamp(v time.Time)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetAcknowledgedBy

`func (o *ServerDriftHistory) GetAcknowledgedBy() int64`

GetAcknowledgedBy returns the AcknowledgedBy field if non-nil, zero value otherwise.

### GetAcknowledgedByOk

`func (o *ServerDriftHistory) GetAcknowledgedByOk() (*int64, bool)`

GetAcknowledgedByOk returns a tuple with the AcknowledgedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedBy

`func (o *ServerDriftHistory) SetAcknowledgedBy(v int64)`

SetAcknowledgedBy sets AcknowledgedBy field to given value.

### HasAcknowledgedBy

`func (o *ServerDriftHistory) HasAcknowledgedBy() bool`

HasAcknowledgedBy returns a boolean if a field has been set.

### GetAcknowledgeTimestamp

`func (o *ServerDriftHistory) GetAcknowledgeTimestamp() time.Time`

GetAcknowledgeTimestamp returns the AcknowledgeTimestamp field if non-nil, zero value otherwise.

### GetAcknowledgeTimestampOk

`func (o *ServerDriftHistory) GetAcknowledgeTimestampOk() (*time.Time, bool)`

GetAcknowledgeTimestampOk returns a tuple with the AcknowledgeTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgeTimestamp

`func (o *ServerDriftHistory) SetAcknowledgeTimestamp(v time.Time)`

SetAcknowledgeTimestamp sets AcknowledgeTimestamp field to given value.

### HasAcknowledgeTimestamp

`func (o *ServerDriftHistory) HasAcknowledgeTimestamp() bool`

HasAcknowledgeTimestamp returns a boolean if a field has been set.

### GetLinks

`func (o *ServerDriftHistory) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServerDriftHistory) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServerDriftHistory) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServerDriftHistory) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


