# NetworkDeviceExternalConnectionInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkDeviceId** | **float32** | Network device id | 
**NetworkDeviceInterfaceId** | **float32** | Network device interface id | 
**NetworkDeviceInterfaceName** | Pointer to **string** | Network device interface name | [optional] 
**ExternalConnectionId** | Pointer to **float32** | Network external connection id, null if not connected to an existing external connection | [optional] 
**ExternalConnectionInterfaceId** | Pointer to **float32** | external connection interface id, null if not connected to an existing external connection interface | [optional] 

## Methods

### NewNetworkDeviceExternalConnectionInterface

`func NewNetworkDeviceExternalConnectionInterface(networkDeviceId float32, networkDeviceInterfaceId float32, ) *NetworkDeviceExternalConnectionInterface`

NewNetworkDeviceExternalConnectionInterface instantiates a new NetworkDeviceExternalConnectionInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceExternalConnectionInterfaceWithDefaults

`func NewNetworkDeviceExternalConnectionInterfaceWithDefaults() *NetworkDeviceExternalConnectionInterface`

NewNetworkDeviceExternalConnectionInterfaceWithDefaults instantiates a new NetworkDeviceExternalConnectionInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkDeviceId

`func (o *NetworkDeviceExternalConnectionInterface) GetNetworkDeviceId() float32`

GetNetworkDeviceId returns the NetworkDeviceId field if non-nil, zero value otherwise.

### GetNetworkDeviceIdOk

`func (o *NetworkDeviceExternalConnectionInterface) GetNetworkDeviceIdOk() (*float32, bool)`

GetNetworkDeviceIdOk returns a tuple with the NetworkDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceId

`func (o *NetworkDeviceExternalConnectionInterface) SetNetworkDeviceId(v float32)`

SetNetworkDeviceId sets NetworkDeviceId field to given value.


### GetNetworkDeviceInterfaceId

`func (o *NetworkDeviceExternalConnectionInterface) GetNetworkDeviceInterfaceId() float32`

GetNetworkDeviceInterfaceId returns the NetworkDeviceInterfaceId field if non-nil, zero value otherwise.

### GetNetworkDeviceInterfaceIdOk

`func (o *NetworkDeviceExternalConnectionInterface) GetNetworkDeviceInterfaceIdOk() (*float32, bool)`

GetNetworkDeviceInterfaceIdOk returns a tuple with the NetworkDeviceInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceInterfaceId

`func (o *NetworkDeviceExternalConnectionInterface) SetNetworkDeviceInterfaceId(v float32)`

SetNetworkDeviceInterfaceId sets NetworkDeviceInterfaceId field to given value.


### GetNetworkDeviceInterfaceName

`func (o *NetworkDeviceExternalConnectionInterface) GetNetworkDeviceInterfaceName() string`

GetNetworkDeviceInterfaceName returns the NetworkDeviceInterfaceName field if non-nil, zero value otherwise.

### GetNetworkDeviceInterfaceNameOk

`func (o *NetworkDeviceExternalConnectionInterface) GetNetworkDeviceInterfaceNameOk() (*string, bool)`

GetNetworkDeviceInterfaceNameOk returns a tuple with the NetworkDeviceInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceInterfaceName

`func (o *NetworkDeviceExternalConnectionInterface) SetNetworkDeviceInterfaceName(v string)`

SetNetworkDeviceInterfaceName sets NetworkDeviceInterfaceName field to given value.

### HasNetworkDeviceInterfaceName

`func (o *NetworkDeviceExternalConnectionInterface) HasNetworkDeviceInterfaceName() bool`

HasNetworkDeviceInterfaceName returns a boolean if a field has been set.

### GetExternalConnectionId

`func (o *NetworkDeviceExternalConnectionInterface) GetExternalConnectionId() float32`

GetExternalConnectionId returns the ExternalConnectionId field if non-nil, zero value otherwise.

### GetExternalConnectionIdOk

`func (o *NetworkDeviceExternalConnectionInterface) GetExternalConnectionIdOk() (*float32, bool)`

GetExternalConnectionIdOk returns a tuple with the ExternalConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalConnectionId

`func (o *NetworkDeviceExternalConnectionInterface) SetExternalConnectionId(v float32)`

SetExternalConnectionId sets ExternalConnectionId field to given value.

### HasExternalConnectionId

`func (o *NetworkDeviceExternalConnectionInterface) HasExternalConnectionId() bool`

HasExternalConnectionId returns a boolean if a field has been set.

### GetExternalConnectionInterfaceId

`func (o *NetworkDeviceExternalConnectionInterface) GetExternalConnectionInterfaceId() float32`

GetExternalConnectionInterfaceId returns the ExternalConnectionInterfaceId field if non-nil, zero value otherwise.

### GetExternalConnectionInterfaceIdOk

`func (o *NetworkDeviceExternalConnectionInterface) GetExternalConnectionInterfaceIdOk() (*float32, bool)`

GetExternalConnectionInterfaceIdOk returns a tuple with the ExternalConnectionInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalConnectionInterfaceId

`func (o *NetworkDeviceExternalConnectionInterface) SetExternalConnectionInterfaceId(v float32)`

SetExternalConnectionInterfaceId sets ExternalConnectionInterfaceId field to given value.

### HasExternalConnectionInterfaceId

`func (o *NetworkDeviceExternalConnectionInterface) HasExternalConnectionInterfaceId() bool`

HasExternalConnectionInterfaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


