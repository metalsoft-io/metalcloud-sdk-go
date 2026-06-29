# CreateVMPoolHostInterfaceNetworkDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkDeviceId** | **int64** | ID of the network equipment (switch). Must be active and in leaf position. | 
**NetworkDeviceInterfaceName** | **string** | Name of the interface on the network equipment | 

## Methods

### NewCreateVMPoolHostInterfaceNetworkDevice

`func NewCreateVMPoolHostInterfaceNetworkDevice(networkDeviceId int64, networkDeviceInterfaceName string, ) *CreateVMPoolHostInterfaceNetworkDevice`

NewCreateVMPoolHostInterfaceNetworkDevice instantiates a new CreateVMPoolHostInterfaceNetworkDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVMPoolHostInterfaceNetworkDeviceWithDefaults

`func NewCreateVMPoolHostInterfaceNetworkDeviceWithDefaults() *CreateVMPoolHostInterfaceNetworkDevice`

NewCreateVMPoolHostInterfaceNetworkDeviceWithDefaults instantiates a new CreateVMPoolHostInterfaceNetworkDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkDeviceId

`func (o *CreateVMPoolHostInterfaceNetworkDevice) GetNetworkDeviceId() int64`

GetNetworkDeviceId returns the NetworkDeviceId field if non-nil, zero value otherwise.

### GetNetworkDeviceIdOk

`func (o *CreateVMPoolHostInterfaceNetworkDevice) GetNetworkDeviceIdOk() (*int64, bool)`

GetNetworkDeviceIdOk returns a tuple with the NetworkDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceId

`func (o *CreateVMPoolHostInterfaceNetworkDevice) SetNetworkDeviceId(v int64)`

SetNetworkDeviceId sets NetworkDeviceId field to given value.


### GetNetworkDeviceInterfaceName

`func (o *CreateVMPoolHostInterfaceNetworkDevice) GetNetworkDeviceInterfaceName() string`

GetNetworkDeviceInterfaceName returns the NetworkDeviceInterfaceName field if non-nil, zero value otherwise.

### GetNetworkDeviceInterfaceNameOk

`func (o *CreateVMPoolHostInterfaceNetworkDevice) GetNetworkDeviceInterfaceNameOk() (*string, bool)`

GetNetworkDeviceInterfaceNameOk returns a tuple with the NetworkDeviceInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceInterfaceName

`func (o *CreateVMPoolHostInterfaceNetworkDevice) SetNetworkDeviceInterfaceName(v string)`

SetNetworkDeviceInterfaceName sets NetworkDeviceInterfaceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


