# StoragePort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortId** | **string** | The ID of the port. | 
**DirectorId** | Pointer to **string** | Director id to use (for certain storage drivers) | [optional] 
**NodeId** | Pointer to **string** | Node id on which the port is located (for certain storage drivers) | [optional] 

## Methods

### NewStoragePort

`func NewStoragePort(portId string, ) *StoragePort`

NewStoragePort instantiates a new StoragePort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePortWithDefaults

`func NewStoragePortWithDefaults() *StoragePort`

NewStoragePortWithDefaults instantiates a new StoragePort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortId

`func (o *StoragePort) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *StoragePort) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *StoragePort) SetPortId(v string)`

SetPortId sets PortId field to given value.


### GetDirectorId

`func (o *StoragePort) GetDirectorId() string`

GetDirectorId returns the DirectorId field if non-nil, zero value otherwise.

### GetDirectorIdOk

`func (o *StoragePort) GetDirectorIdOk() (*string, bool)`

GetDirectorIdOk returns a tuple with the DirectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectorId

`func (o *StoragePort) SetDirectorId(v string)`

SetDirectorId sets DirectorId field to given value.

### HasDirectorId

`func (o *StoragePort) HasDirectorId() bool`

HasDirectorId returns a boolean if a field has been set.

### GetNodeId

`func (o *StoragePort) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *StoragePort) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *StoragePort) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *StoragePort) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


