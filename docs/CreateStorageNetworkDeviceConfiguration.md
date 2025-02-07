# CreateStorageNetworkDeviceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkEquipmentId** | **float32** | Id of the network equipment | 
**StoragePhysicalInterfaceIdentifier** | **string** | Identifier of the storage physical interface | 
**NetworkEquipmentInterfaceIdentifier** | **string** | Identifier of the network equipment interface | 
**NetworkEquipmentInterfaceVlans** | **[]float32** | Array of VLANS for the network equipment interface | 

## Methods

### NewCreateStorageNetworkDeviceConfiguration

`func NewCreateStorageNetworkDeviceConfiguration(networkEquipmentId float32, storagePhysicalInterfaceIdentifier string, networkEquipmentInterfaceIdentifier string, networkEquipmentInterfaceVlans []float32, ) *CreateStorageNetworkDeviceConfiguration`

NewCreateStorageNetworkDeviceConfiguration instantiates a new CreateStorageNetworkDeviceConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStorageNetworkDeviceConfigurationWithDefaults

`func NewCreateStorageNetworkDeviceConfigurationWithDefaults() *CreateStorageNetworkDeviceConfiguration`

NewCreateStorageNetworkDeviceConfigurationWithDefaults instantiates a new CreateStorageNetworkDeviceConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkEquipmentId

`func (o *CreateStorageNetworkDeviceConfiguration) GetNetworkEquipmentId() float32`

GetNetworkEquipmentId returns the NetworkEquipmentId field if non-nil, zero value otherwise.

### GetNetworkEquipmentIdOk

`func (o *CreateStorageNetworkDeviceConfiguration) GetNetworkEquipmentIdOk() (*float32, bool)`

GetNetworkEquipmentIdOk returns a tuple with the NetworkEquipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEquipmentId

`func (o *CreateStorageNetworkDeviceConfiguration) SetNetworkEquipmentId(v float32)`

SetNetworkEquipmentId sets NetworkEquipmentId field to given value.


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


### GetNetworkEquipmentInterfaceIdentifier

`func (o *CreateStorageNetworkDeviceConfiguration) GetNetworkEquipmentInterfaceIdentifier() string`

GetNetworkEquipmentInterfaceIdentifier returns the NetworkEquipmentInterfaceIdentifier field if non-nil, zero value otherwise.

### GetNetworkEquipmentInterfaceIdentifierOk

`func (o *CreateStorageNetworkDeviceConfiguration) GetNetworkEquipmentInterfaceIdentifierOk() (*string, bool)`

GetNetworkEquipmentInterfaceIdentifierOk returns a tuple with the NetworkEquipmentInterfaceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEquipmentInterfaceIdentifier

`func (o *CreateStorageNetworkDeviceConfiguration) SetNetworkEquipmentInterfaceIdentifier(v string)`

SetNetworkEquipmentInterfaceIdentifier sets NetworkEquipmentInterfaceIdentifier field to given value.


### GetNetworkEquipmentInterfaceVlans

`func (o *CreateStorageNetworkDeviceConfiguration) GetNetworkEquipmentInterfaceVlans() []float32`

GetNetworkEquipmentInterfaceVlans returns the NetworkEquipmentInterfaceVlans field if non-nil, zero value otherwise.

### GetNetworkEquipmentInterfaceVlansOk

`func (o *CreateStorageNetworkDeviceConfiguration) GetNetworkEquipmentInterfaceVlansOk() (*[]float32, bool)`

GetNetworkEquipmentInterfaceVlansOk returns a tuple with the NetworkEquipmentInterfaceVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEquipmentInterfaceVlans

`func (o *CreateStorageNetworkDeviceConfiguration) SetNetworkEquipmentInterfaceVlans(v []float32)`

SetNetworkEquipmentInterfaceVlans sets NetworkEquipmentInterfaceVlans field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


