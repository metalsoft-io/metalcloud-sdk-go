# NetworkDeviceEndpointInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkDeviceId** | **float32** | Network device id | 
**NetworkDeviceInterfaceId** | **float32** | Network device interface id | 
**NetworkDeviceInterfaceName** | Pointer to **string** | Network device interface name | [optional] 
**EndpointId** | Pointer to **string** | Network endpoint id if the interface is connected to a network endpoint | [optional] 
**EndpointInterfaceId** | Pointer to **float32** | Endpoint interface id | [optional] 

## Methods

### NewNetworkDeviceEndpointInterface

`func NewNetworkDeviceEndpointInterface(networkDeviceId float32, networkDeviceInterfaceId float32, ) *NetworkDeviceEndpointInterface`

NewNetworkDeviceEndpointInterface instantiates a new NetworkDeviceEndpointInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceEndpointInterfaceWithDefaults

`func NewNetworkDeviceEndpointInterfaceWithDefaults() *NetworkDeviceEndpointInterface`

NewNetworkDeviceEndpointInterfaceWithDefaults instantiates a new NetworkDeviceEndpointInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkDeviceId

`func (o *NetworkDeviceEndpointInterface) GetNetworkDeviceId() float32`

GetNetworkDeviceId returns the NetworkDeviceId field if non-nil, zero value otherwise.

### GetNetworkDeviceIdOk

`func (o *NetworkDeviceEndpointInterface) GetNetworkDeviceIdOk() (*float32, bool)`

GetNetworkDeviceIdOk returns a tuple with the NetworkDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceId

`func (o *NetworkDeviceEndpointInterface) SetNetworkDeviceId(v float32)`

SetNetworkDeviceId sets NetworkDeviceId field to given value.


### GetNetworkDeviceInterfaceId

`func (o *NetworkDeviceEndpointInterface) GetNetworkDeviceInterfaceId() float32`

GetNetworkDeviceInterfaceId returns the NetworkDeviceInterfaceId field if non-nil, zero value otherwise.

### GetNetworkDeviceInterfaceIdOk

`func (o *NetworkDeviceEndpointInterface) GetNetworkDeviceInterfaceIdOk() (*float32, bool)`

GetNetworkDeviceInterfaceIdOk returns a tuple with the NetworkDeviceInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceInterfaceId

`func (o *NetworkDeviceEndpointInterface) SetNetworkDeviceInterfaceId(v float32)`

SetNetworkDeviceInterfaceId sets NetworkDeviceInterfaceId field to given value.


### GetNetworkDeviceInterfaceName

`func (o *NetworkDeviceEndpointInterface) GetNetworkDeviceInterfaceName() string`

GetNetworkDeviceInterfaceName returns the NetworkDeviceInterfaceName field if non-nil, zero value otherwise.

### GetNetworkDeviceInterfaceNameOk

`func (o *NetworkDeviceEndpointInterface) GetNetworkDeviceInterfaceNameOk() (*string, bool)`

GetNetworkDeviceInterfaceNameOk returns a tuple with the NetworkDeviceInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceInterfaceName

`func (o *NetworkDeviceEndpointInterface) SetNetworkDeviceInterfaceName(v string)`

SetNetworkDeviceInterfaceName sets NetworkDeviceInterfaceName field to given value.

### HasNetworkDeviceInterfaceName

`func (o *NetworkDeviceEndpointInterface) HasNetworkDeviceInterfaceName() bool`

HasNetworkDeviceInterfaceName returns a boolean if a field has been set.

### GetEndpointId

`func (o *NetworkDeviceEndpointInterface) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *NetworkDeviceEndpointInterface) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *NetworkDeviceEndpointInterface) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *NetworkDeviceEndpointInterface) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetEndpointInterfaceId

`func (o *NetworkDeviceEndpointInterface) GetEndpointInterfaceId() float32`

GetEndpointInterfaceId returns the EndpointInterfaceId field if non-nil, zero value otherwise.

### GetEndpointInterfaceIdOk

`func (o *NetworkDeviceEndpointInterface) GetEndpointInterfaceIdOk() (*float32, bool)`

GetEndpointInterfaceIdOk returns a tuple with the EndpointInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointInterfaceId

`func (o *NetworkDeviceEndpointInterface) SetEndpointInterfaceId(v float32)`

SetEndpointInterfaceId sets EndpointInterfaceId field to given value.

### HasEndpointInterfaceId

`func (o *NetworkDeviceEndpointInterface) HasEndpointInterfaceId() bool`

HasEndpointInterfaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


