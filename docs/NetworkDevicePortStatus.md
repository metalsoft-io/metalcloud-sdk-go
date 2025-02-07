# NetworkDevicePortStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ports** | **[]string** | The ports of the network device that will have their status changed | 
**Status** | **bool** | The new status of the ports | 

## Methods

### NewNetworkDevicePortStatus

`func NewNetworkDevicePortStatus(ports []string, status bool, ) *NetworkDevicePortStatus`

NewNetworkDevicePortStatus instantiates a new NetworkDevicePortStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDevicePortStatusWithDefaults

`func NewNetworkDevicePortStatusWithDefaults() *NetworkDevicePortStatus`

NewNetworkDevicePortStatusWithDefaults instantiates a new NetworkDevicePortStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPorts

`func (o *NetworkDevicePortStatus) GetPorts() []string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *NetworkDevicePortStatus) GetPortsOk() (*[]string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *NetworkDevicePortStatus) SetPorts(v []string)`

SetPorts sets Ports field to given value.


### GetStatus

`func (o *NetworkDevicePortStatus) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkDevicePortStatus) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkDevicePortStatus) SetStatus(v bool)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


