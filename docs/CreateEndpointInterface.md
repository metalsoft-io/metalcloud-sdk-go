# CreateEndpointInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkEquipmentInterfaceId** | **float32** | Network equipment interface id | 
**MacAddress** | Pointer to **string** | Network equipment interface mac address | [optional] 

## Methods

### NewCreateEndpointInterface

`func NewCreateEndpointInterface(networkEquipmentInterfaceId float32, ) *CreateEndpointInterface`

NewCreateEndpointInterface instantiates a new CreateEndpointInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEndpointInterfaceWithDefaults

`func NewCreateEndpointInterfaceWithDefaults() *CreateEndpointInterface`

NewCreateEndpointInterfaceWithDefaults instantiates a new CreateEndpointInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkEquipmentInterfaceId

`func (o *CreateEndpointInterface) GetNetworkEquipmentInterfaceId() float32`

GetNetworkEquipmentInterfaceId returns the NetworkEquipmentInterfaceId field if non-nil, zero value otherwise.

### GetNetworkEquipmentInterfaceIdOk

`func (o *CreateEndpointInterface) GetNetworkEquipmentInterfaceIdOk() (*float32, bool)`

GetNetworkEquipmentInterfaceIdOk returns a tuple with the NetworkEquipmentInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEquipmentInterfaceId

`func (o *CreateEndpointInterface) SetNetworkEquipmentInterfaceId(v float32)`

SetNetworkEquipmentInterfaceId sets NetworkEquipmentInterfaceId field to given value.


### GetMacAddress

`func (o *CreateEndpointInterface) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *CreateEndpointInterface) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *CreateEndpointInterface) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *CreateEndpointInterface) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


