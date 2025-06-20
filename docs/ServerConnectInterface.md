# ServerConnectInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerInterfaceId** | **float32** | The id of the server interface | 
**NetworkDevicePortId** | **string** | The network device port name | 
**NetworkDeviceHostname** | **string** | The network device hostname | 

## Methods

### NewServerConnectInterface

`func NewServerConnectInterface(serverInterfaceId float32, networkDevicePortId string, networkDeviceHostname string, ) *ServerConnectInterface`

NewServerConnectInterface instantiates a new ServerConnectInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerConnectInterfaceWithDefaults

`func NewServerConnectInterfaceWithDefaults() *ServerConnectInterface`

NewServerConnectInterfaceWithDefaults instantiates a new ServerConnectInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerInterfaceId

`func (o *ServerConnectInterface) GetServerInterfaceId() float32`

GetServerInterfaceId returns the ServerInterfaceId field if non-nil, zero value otherwise.

### GetServerInterfaceIdOk

`func (o *ServerConnectInterface) GetServerInterfaceIdOk() (*float32, bool)`

GetServerInterfaceIdOk returns a tuple with the ServerInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInterfaceId

`func (o *ServerConnectInterface) SetServerInterfaceId(v float32)`

SetServerInterfaceId sets ServerInterfaceId field to given value.


### GetNetworkDevicePortId

`func (o *ServerConnectInterface) GetNetworkDevicePortId() string`

GetNetworkDevicePortId returns the NetworkDevicePortId field if non-nil, zero value otherwise.

### GetNetworkDevicePortIdOk

`func (o *ServerConnectInterface) GetNetworkDevicePortIdOk() (*string, bool)`

GetNetworkDevicePortIdOk returns a tuple with the NetworkDevicePortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDevicePortId

`func (o *ServerConnectInterface) SetNetworkDevicePortId(v string)`

SetNetworkDevicePortId sets NetworkDevicePortId field to given value.


### GetNetworkDeviceHostname

`func (o *ServerConnectInterface) GetNetworkDeviceHostname() string`

GetNetworkDeviceHostname returns the NetworkDeviceHostname field if non-nil, zero value otherwise.

### GetNetworkDeviceHostnameOk

`func (o *ServerConnectInterface) GetNetworkDeviceHostnameOk() (*string, bool)`

GetNetworkDeviceHostnameOk returns a tuple with the NetworkDeviceHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceHostname

`func (o *ServerConnectInterface) SetNetworkDeviceHostname(v string)`

SetNetworkDeviceHostname sets NetworkDeviceHostname field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


