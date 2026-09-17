# ServerSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Oid** | **string** | The OID of the commit the snapshot is based on | 
**Message** | **string** | The commit message associated with the snapshot | 
**Timestamp** | **string** | The timestamp when the snapshot was taken | 
**Kind** | Pointer to **string** | The classification of this commit based on its message | [optional] 

## Methods

### NewServerSnapshot

`func NewServerSnapshot(oid string, message string, timestamp string, ) *ServerSnapshot`

NewServerSnapshot instantiates a new ServerSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerSnapshotWithDefaults

`func NewServerSnapshotWithDefaults() *ServerSnapshot`

NewServerSnapshotWithDefaults instantiates a new ServerSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOid

`func (o *ServerSnapshot) GetOid() string`

GetOid returns the Oid field if non-nil, zero value otherwise.

### GetOidOk

`func (o *ServerSnapshot) GetOidOk() (*string, bool)`

GetOidOk returns a tuple with the Oid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOid

`func (o *ServerSnapshot) SetOid(v string)`

SetOid sets Oid field to given value.


### GetMessage

`func (o *ServerSnapshot) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ServerSnapshot) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ServerSnapshot) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTimestamp

`func (o *ServerSnapshot) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ServerSnapshot) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ServerSnapshot) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetKind

`func (o *ServerSnapshot) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ServerSnapshot) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ServerSnapshot) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ServerSnapshot) HasKind() bool`

HasKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


