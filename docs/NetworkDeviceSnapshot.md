# NetworkDeviceSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Oid** | **string** | The OID of the commit the snapshot is based on | 
**Message** | **string** | The commit message associated with the snapshot | 
**Timestamp** | **string** | The timestamp when the snapshot was taken | 
**Kind** | Pointer to **string** | The classification of this commit based on its message | [optional] 

## Methods

### NewNetworkDeviceSnapshot

`func NewNetworkDeviceSnapshot(oid string, message string, timestamp string, ) *NetworkDeviceSnapshot`

NewNetworkDeviceSnapshot instantiates a new NetworkDeviceSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceSnapshotWithDefaults

`func NewNetworkDeviceSnapshotWithDefaults() *NetworkDeviceSnapshot`

NewNetworkDeviceSnapshotWithDefaults instantiates a new NetworkDeviceSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOid

`func (o *NetworkDeviceSnapshot) GetOid() string`

GetOid returns the Oid field if non-nil, zero value otherwise.

### GetOidOk

`func (o *NetworkDeviceSnapshot) GetOidOk() (*string, bool)`

GetOidOk returns a tuple with the Oid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOid

`func (o *NetworkDeviceSnapshot) SetOid(v string)`

SetOid sets Oid field to given value.


### GetMessage

`func (o *NetworkDeviceSnapshot) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NetworkDeviceSnapshot) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NetworkDeviceSnapshot) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTimestamp

`func (o *NetworkDeviceSnapshot) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *NetworkDeviceSnapshot) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *NetworkDeviceSnapshot) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetKind

`func (o *NetworkDeviceSnapshot) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkDeviceSnapshot) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkDeviceSnapshot) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NetworkDeviceSnapshot) HasKind() bool`

HasKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


