# VMPoolHostInterfaceNetworkDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Network device assignment ID | 
**HostInterfaceId** | **int64** | VM Pool Host Interface ID | 
**NetworkDeviceId** | **int64** | ID of the network equipment (switch) | 
**NetworkDeviceInterfaceName** | **string** | Name of the interface on the network equipment | 

## Methods

### NewVMPoolHostInterfaceNetworkDevice

`func NewVMPoolHostInterfaceNetworkDevice(id int64, hostInterfaceId int64, networkDeviceId int64, networkDeviceInterfaceName string, ) *VMPoolHostInterfaceNetworkDevice`

NewVMPoolHostInterfaceNetworkDevice instantiates a new VMPoolHostInterfaceNetworkDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMPoolHostInterfaceNetworkDeviceWithDefaults

`func NewVMPoolHostInterfaceNetworkDeviceWithDefaults() *VMPoolHostInterfaceNetworkDevice`

NewVMPoolHostInterfaceNetworkDeviceWithDefaults instantiates a new VMPoolHostInterfaceNetworkDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VMPoolHostInterfaceNetworkDevice) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VMPoolHostInterfaceNetworkDevice) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VMPoolHostInterfaceNetworkDevice) SetId(v int64)`

SetId sets Id field to given value.


### GetHostInterfaceId

`func (o *VMPoolHostInterfaceNetworkDevice) GetHostInterfaceId() int64`

GetHostInterfaceId returns the HostInterfaceId field if non-nil, zero value otherwise.

### GetHostInterfaceIdOk

`func (o *VMPoolHostInterfaceNetworkDevice) GetHostInterfaceIdOk() (*int64, bool)`

GetHostInterfaceIdOk returns a tuple with the HostInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostInterfaceId

`func (o *VMPoolHostInterfaceNetworkDevice) SetHostInterfaceId(v int64)`

SetHostInterfaceId sets HostInterfaceId field to given value.


### GetNetworkDeviceId

`func (o *VMPoolHostInterfaceNetworkDevice) GetNetworkDeviceId() int64`

GetNetworkDeviceId returns the NetworkDeviceId field if non-nil, zero value otherwise.

### GetNetworkDeviceIdOk

`func (o *VMPoolHostInterfaceNetworkDevice) GetNetworkDeviceIdOk() (*int64, bool)`

GetNetworkDeviceIdOk returns a tuple with the NetworkDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceId

`func (o *VMPoolHostInterfaceNetworkDevice) SetNetworkDeviceId(v int64)`

SetNetworkDeviceId sets NetworkDeviceId field to given value.


### GetNetworkDeviceInterfaceName

`func (o *VMPoolHostInterfaceNetworkDevice) GetNetworkDeviceInterfaceName() string`

GetNetworkDeviceInterfaceName returns the NetworkDeviceInterfaceName field if non-nil, zero value otherwise.

### GetNetworkDeviceInterfaceNameOk

`func (o *VMPoolHostInterfaceNetworkDevice) GetNetworkDeviceInterfaceNameOk() (*string, bool)`

GetNetworkDeviceInterfaceNameOk returns a tuple with the NetworkDeviceInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceInterfaceName

`func (o *VMPoolHostInterfaceNetworkDevice) SetNetworkDeviceInterfaceName(v string)`

SetNetworkDeviceInterfaceName sets NetworkDeviceInterfaceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


