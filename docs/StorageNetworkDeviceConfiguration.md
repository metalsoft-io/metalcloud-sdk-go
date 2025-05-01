# StorageNetworkDeviceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | Id of the storage network device configuration | 
**StorageId** | **float32** | Id of the storage | 
**NetworkDeviceId** | **float32** | Id of the network device | 
**StoragePhysicalInterfaceIdentifier** | **string** | Identifier of the storage physical interface | 
**NetworkDeviceIdentifier** | **string** | Identifier of the network device | 
**NetworkDeviceInterfaceIdentifier** | **string** | Identifier of the network device interface | 
**NetworkDeviceInterfaceVlans** | Pointer to **[]float32** | Array of VLANS for the network device interface | [optional] 

## Methods

### NewStorageNetworkDeviceConfiguration

`func NewStorageNetworkDeviceConfiguration(id float32, storageId float32, networkDeviceId float32, storagePhysicalInterfaceIdentifier string, networkDeviceIdentifier string, networkDeviceInterfaceIdentifier string, ) *StorageNetworkDeviceConfiguration`

NewStorageNetworkDeviceConfiguration instantiates a new StorageNetworkDeviceConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetworkDeviceConfigurationWithDefaults

`func NewStorageNetworkDeviceConfigurationWithDefaults() *StorageNetworkDeviceConfiguration`

NewStorageNetworkDeviceConfigurationWithDefaults instantiates a new StorageNetworkDeviceConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorageNetworkDeviceConfiguration) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageNetworkDeviceConfiguration) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageNetworkDeviceConfiguration) SetId(v float32)`

SetId sets Id field to given value.


### GetStorageId

`func (o *StorageNetworkDeviceConfiguration) GetStorageId() float32`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *StorageNetworkDeviceConfiguration) GetStorageIdOk() (*float32, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *StorageNetworkDeviceConfiguration) SetStorageId(v float32)`

SetStorageId sets StorageId field to given value.


### GetNetworkDeviceId

`func (o *StorageNetworkDeviceConfiguration) GetNetworkDeviceId() float32`

GetNetworkDeviceId returns the NetworkDeviceId field if non-nil, zero value otherwise.

### GetNetworkDeviceIdOk

`func (o *StorageNetworkDeviceConfiguration) GetNetworkDeviceIdOk() (*float32, bool)`

GetNetworkDeviceIdOk returns a tuple with the NetworkDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceId

`func (o *StorageNetworkDeviceConfiguration) SetNetworkDeviceId(v float32)`

SetNetworkDeviceId sets NetworkDeviceId field to given value.


### GetStoragePhysicalInterfaceIdentifier

`func (o *StorageNetworkDeviceConfiguration) GetStoragePhysicalInterfaceIdentifier() string`

GetStoragePhysicalInterfaceIdentifier returns the StoragePhysicalInterfaceIdentifier field if non-nil, zero value otherwise.

### GetStoragePhysicalInterfaceIdentifierOk

`func (o *StorageNetworkDeviceConfiguration) GetStoragePhysicalInterfaceIdentifierOk() (*string, bool)`

GetStoragePhysicalInterfaceIdentifierOk returns a tuple with the StoragePhysicalInterfaceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePhysicalInterfaceIdentifier

`func (o *StorageNetworkDeviceConfiguration) SetStoragePhysicalInterfaceIdentifier(v string)`

SetStoragePhysicalInterfaceIdentifier sets StoragePhysicalInterfaceIdentifier field to given value.


### GetNetworkDeviceIdentifier

`func (o *StorageNetworkDeviceConfiguration) GetNetworkDeviceIdentifier() string`

GetNetworkDeviceIdentifier returns the NetworkDeviceIdentifier field if non-nil, zero value otherwise.

### GetNetworkDeviceIdentifierOk

`func (o *StorageNetworkDeviceConfiguration) GetNetworkDeviceIdentifierOk() (*string, bool)`

GetNetworkDeviceIdentifierOk returns a tuple with the NetworkDeviceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceIdentifier

`func (o *StorageNetworkDeviceConfiguration) SetNetworkDeviceIdentifier(v string)`

SetNetworkDeviceIdentifier sets NetworkDeviceIdentifier field to given value.


### GetNetworkDeviceInterfaceIdentifier

`func (o *StorageNetworkDeviceConfiguration) GetNetworkDeviceInterfaceIdentifier() string`

GetNetworkDeviceInterfaceIdentifier returns the NetworkDeviceInterfaceIdentifier field if non-nil, zero value otherwise.

### GetNetworkDeviceInterfaceIdentifierOk

`func (o *StorageNetworkDeviceConfiguration) GetNetworkDeviceInterfaceIdentifierOk() (*string, bool)`

GetNetworkDeviceInterfaceIdentifierOk returns a tuple with the NetworkDeviceInterfaceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceInterfaceIdentifier

`func (o *StorageNetworkDeviceConfiguration) SetNetworkDeviceInterfaceIdentifier(v string)`

SetNetworkDeviceInterfaceIdentifier sets NetworkDeviceInterfaceIdentifier field to given value.


### GetNetworkDeviceInterfaceVlans

`func (o *StorageNetworkDeviceConfiguration) GetNetworkDeviceInterfaceVlans() []float32`

GetNetworkDeviceInterfaceVlans returns the NetworkDeviceInterfaceVlans field if non-nil, zero value otherwise.

### GetNetworkDeviceInterfaceVlansOk

`func (o *StorageNetworkDeviceConfiguration) GetNetworkDeviceInterfaceVlansOk() (*[]float32, bool)`

GetNetworkDeviceInterfaceVlansOk returns a tuple with the NetworkDeviceInterfaceVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceInterfaceVlans

`func (o *StorageNetworkDeviceConfiguration) SetNetworkDeviceInterfaceVlans(v []float32)`

SetNetworkDeviceInterfaceVlans sets NetworkDeviceInterfaceVlans field to given value.

### HasNetworkDeviceInterfaceVlans

`func (o *StorageNetworkDeviceConfiguration) HasNetworkDeviceInterfaceVlans() bool`

HasNetworkDeviceInterfaceVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


