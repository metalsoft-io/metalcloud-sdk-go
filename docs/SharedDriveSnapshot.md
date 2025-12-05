# SharedDriveSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the Snapshot. | 
**CreatedTimestamp** | **string** | Timestamp of the Snapshot creation. | 

## Methods

### NewSharedDriveSnapshot

`func NewSharedDriveSnapshot(name string, createdTimestamp string, ) *SharedDriveSnapshot`

NewSharedDriveSnapshot instantiates a new SharedDriveSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedDriveSnapshotWithDefaults

`func NewSharedDriveSnapshotWithDefaults() *SharedDriveSnapshot`

NewSharedDriveSnapshotWithDefaults instantiates a new SharedDriveSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SharedDriveSnapshot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SharedDriveSnapshot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SharedDriveSnapshot) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedTimestamp

`func (o *SharedDriveSnapshot) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *SharedDriveSnapshot) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *SharedDriveSnapshot) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


