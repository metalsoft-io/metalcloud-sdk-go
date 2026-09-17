# StorageScopedAccessUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Id of the storage scoped access user | 
**Revision** | **int64** | Revision of the Storage Interface | 
**StorageId** | **int64** | Id of the storage | 
**InfrastructureId** | **int64** | Id of the infrastructure where this user is deployed | 
**Username** | **string** | Username of the storage scoped access user | 
**CreatedTimestamp** | **time.Time** | Timestamp when the user was created | 
**UpdatedTimestamp** | **time.Time** | Timestamp when the user was updated | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewStorageScopedAccessUser

`func NewStorageScopedAccessUser(id int64, revision int64, storageId int64, infrastructureId int64, username string, createdTimestamp time.Time, updatedTimestamp time.Time, ) *StorageScopedAccessUser`

NewStorageScopedAccessUser instantiates a new StorageScopedAccessUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageScopedAccessUserWithDefaults

`func NewStorageScopedAccessUserWithDefaults() *StorageScopedAccessUser`

NewStorageScopedAccessUserWithDefaults instantiates a new StorageScopedAccessUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorageScopedAccessUser) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageScopedAccessUser) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageScopedAccessUser) SetId(v int64)`

SetId sets Id field to given value.


### GetRevision

`func (o *StorageScopedAccessUser) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *StorageScopedAccessUser) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *StorageScopedAccessUser) SetRevision(v int64)`

SetRevision sets Revision field to given value.


### GetStorageId

`func (o *StorageScopedAccessUser) GetStorageId() int64`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *StorageScopedAccessUser) GetStorageIdOk() (*int64, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *StorageScopedAccessUser) SetStorageId(v int64)`

SetStorageId sets StorageId field to given value.


### GetInfrastructureId

`func (o *StorageScopedAccessUser) GetInfrastructureId() int64`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *StorageScopedAccessUser) GetInfrastructureIdOk() (*int64, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *StorageScopedAccessUser) SetInfrastructureId(v int64)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetUsername

`func (o *StorageScopedAccessUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *StorageScopedAccessUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *StorageScopedAccessUser) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetCreatedTimestamp

`func (o *StorageScopedAccessUser) GetCreatedTimestamp() time.Time`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *StorageScopedAccessUser) GetCreatedTimestampOk() (*time.Time, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *StorageScopedAccessUser) SetCreatedTimestamp(v time.Time)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *StorageScopedAccessUser) GetUpdatedTimestamp() time.Time`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *StorageScopedAccessUser) GetUpdatedTimestampOk() (*time.Time, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *StorageScopedAccessUser) SetUpdatedTimestamp(v time.Time)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetLinks

`func (o *StorageScopedAccessUser) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StorageScopedAccessUser) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StorageScopedAccessUser) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StorageScopedAccessUser) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


