# NetworkDevicePorts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SwitchId** | **string** |  | 
**Ports** | [**[]NetworkDevicePort**](NetworkDevicePort.md) |  | 

## Methods

### NewNetworkDevicePorts

`func NewNetworkDevicePorts(switchId string, ports []NetworkDevicePort, ) *NetworkDevicePorts`

NewNetworkDevicePorts instantiates a new NetworkDevicePorts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDevicePortsWithDefaults

`func NewNetworkDevicePortsWithDefaults() *NetworkDevicePorts`

NewNetworkDevicePortsWithDefaults instantiates a new NetworkDevicePorts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSwitchId

`func (o *NetworkDevicePorts) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *NetworkDevicePorts) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *NetworkDevicePorts) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.


### GetPorts

`func (o *NetworkDevicePorts) GetPorts() []NetworkDevicePort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *NetworkDevicePorts) GetPortsOk() (*[]NetworkDevicePort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *NetworkDevicePorts) SetPorts(v []NetworkDevicePort)`

SetPorts sets Ports field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


