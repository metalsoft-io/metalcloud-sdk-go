# CreateStorageNetworkDeviceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkDeviceId** | **float32** | Id of the network device | 
**StoragePhysicalInterfaceIdentifier** | **string** | Identifier of the storage physical interface | 
**NetworkDeviceInterfaceIdentifier** | **string** | Identifier of the network device interface | 
**NetworkDeviceInterfaceVlans** | Pointer to **[]float32** | Array of VLANS for the network device interface | [optional] 

## Methods

### NewCreateStorageNetworkDeviceConfiguration

`func NewCreateStorageNetworkDeviceConfiguration(networkDeviceId float32, storagePhysicalInterfaceIdentifier string, networkDeviceInterfaceIdentifier string, ) *CreateStorageNetworkDeviceConfiguration`

NewCreateStorageNetworkDeviceConfiguration instantiates a new CreateStorageNetworkDeviceConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStorageNetworkDeviceConfigurationWithDefaults

`func NewCreateStorageNetworkDeviceConfigurationWithDefaults() *CreateStorageNetworkDeviceConfiguration`

NewCreateStorageNetworkDeviceConfigurationWithDefaults instantiates a new CreateStorageNetworkDeviceConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkDeviceId

`func (o *CreateStorageNetworkDeviceConfiguration) GetNetworkDeviceId() float32`

GetNetworkDeviceId returns the NetworkDeviceId field if non-nil, zero value otherwise.

### GetNetworkDeviceIdOk

`func (o *CreateStorageNetworkDeviceConfiguration) GetNetworkDeviceIdOk() (*float32, bool)`

GetNetworkDeviceIdOk returns a tuple with the NetworkDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceId

`func (o *CreateStorageNetworkDeviceConfiguration) SetNetworkDeviceId(v float32)`

SetNetworkDeviceId sets NetworkDeviceId field to given value.


### GetStoragePhysicalInterfaceIdentifier

`func (o *CreateStorageNetworkDeviceConfiguration) GetStoragePhysicalInterfaceIdentifier() string`

GetStoragePhysicalInterfaceIdentifier returns the StoragePhysicalInterfaceIdentifier field if non-nil, zero value otherwise.

### GetStoragePhysicalInterfaceIdentifierOk

`func (o *CreateStorageNetworkDeviceConfiguration) GetStoragePhysicalInterfaceIdentifierOk() (*string, bool)`

GetStoragePhysicalInterfaceIdentifierOk returns a tuple with the StoragePhysicalInterfaceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePhysicalInterfaceIdentifier

`func (o *CreateStorageNetworkDeviceConfiguration) SetStoragePhysicalInterfaceIdentifier(v string)`

SetStoragePhysicalInterfaceIdentifier sets StoragePhysicalInterfaceIdentifier field to given value.


### GetNetworkDeviceInterfaceIdentifier

`func (o *CreateStorageNetworkDeviceConfiguration) GetNetworkDeviceInterfaceIdentifier() string`

GetNetworkDeviceInterfaceIdentifier returns the NetworkDeviceInterfaceIdentifier field if non-nil, zero value otherwise.

### GetNetworkDeviceInterfaceIdentifierOk

`func (o *CreateStorageNetworkDeviceConfiguration) GetNetworkDeviceInterfaceIdentifierOk() (*string, bool)`

GetNetworkDeviceInterfaceIdentifierOk returns a tuple with the NetworkDeviceInterfaceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceInterfaceIdentifier

`func (o *CreateStorageNetworkDeviceConfiguration) SetNetworkDeviceInterfaceIdentifier(v string)`

SetNetworkDeviceInterfaceIdentifier sets NetworkDeviceInterfaceIdentifier field to given value.


### GetNetworkDeviceInterfaceVlans

`func (o *CreateStorageNetworkDeviceConfiguration) GetNetworkDeviceInterfaceVlans() []float32`

GetNetworkDeviceInterfaceVlans returns the NetworkDeviceInterfaceVlans field if non-nil, zero value otherwise.

### GetNetworkDeviceInterfaceVlansOk

`func (o *CreateStorageNetworkDeviceConfiguration) GetNetworkDeviceInterfaceVlansOk() (*[]float32, bool)`

GetNetworkDeviceInterfaceVlansOk returns a tuple with the NetworkDeviceInterfaceVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceInterfaceVlans

`func (o *CreateStorageNetworkDeviceConfiguration) SetNetworkDeviceInterfaceVlans(v []float32)`

SetNetworkDeviceInterfaceVlans sets NetworkDeviceInterfaceVlans field to given value.

### HasNetworkDeviceInterfaceVlans

`func (o *CreateStorageNetworkDeviceConfiguration) HasNetworkDeviceInterfaceVlans() bool`

HasNetworkDeviceInterfaceVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


