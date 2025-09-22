# StoragePortDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortId** | **string** | The ID of the port. | 
**DirectorId** | Pointer to **string** | Director id to use (for certain storage drivers) | [optional] 

## Methods

### NewStoragePortDto

`func NewStoragePortDto(portId string, ) *StoragePortDto`

NewStoragePortDto instantiates a new StoragePortDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePortDtoWithDefaults

`func NewStoragePortDtoWithDefaults() *StoragePortDto`

NewStoragePortDtoWithDefaults instantiates a new StoragePortDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortId

`func (o *StoragePortDto) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *StoragePortDto) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *StoragePortDto) SetPortId(v string)`

SetPortId sets PortId field to given value.


### GetDirectorId

`func (o *StoragePortDto) GetDirectorId() string`

GetDirectorId returns the DirectorId field if non-nil, zero value otherwise.

### GetDirectorIdOk

`func (o *StoragePortDto) GetDirectorIdOk() (*string, bool)`

GetDirectorIdOk returns a tuple with the DirectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectorId

`func (o *StoragePortDto) SetDirectorId(v string)`

SetDirectorId sets DirectorId field to given value.

### HasDirectorId

`func (o *StoragePortDto) HasDirectorId() bool`

HasDirectorId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


