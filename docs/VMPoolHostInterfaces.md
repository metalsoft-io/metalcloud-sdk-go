# VMPoolHostInterfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | VM Pool Host Interface ID | 
**HostId** | **int64** | VM Pool Host ID | 
**Status** | [**VMPoolHostInterfaceStatus**](VMPoolHostInterfaceStatus.md) | Status of the VM Pool Host Interface | 
**Name** | **string** | Name of the VM Pool Host Interface | 
**MacAddress** | **string** | MAC Address of the VM Pool Host Interface | 
**Fabric** | Pointer to [**GenericNetworkFabric**](GenericNetworkFabric.md) | Fabric type of the VM Pool Host Interface | [optional] 
**NetworkDevices** | Pointer to [**[]VMPoolHostInterfaceNetworkDevice**](VMPoolHostInterfaceNetworkDevice.md) | Network device assignments for this interface | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewVMPoolHostInterfaces

`func NewVMPoolHostInterfaces(id int64, hostId int64, status VMPoolHostInterfaceStatus, name string, macAddress string, ) *VMPoolHostInterfaces`

NewVMPoolHostInterfaces instantiates a new VMPoolHostInterfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMPoolHostInterfacesWithDefaults

`func NewVMPoolHostInterfacesWithDefaults() *VMPoolHostInterfaces`

NewVMPoolHostInterfacesWithDefaults instantiates a new VMPoolHostInterfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VMPoolHostInterfaces) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VMPoolHostInterfaces) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VMPoolHostInterfaces) SetId(v int64)`

SetId sets Id field to given value.


### GetHostId

`func (o *VMPoolHostInterfaces) GetHostId() int64`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *VMPoolHostInterfaces) GetHostIdOk() (*int64, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *VMPoolHostInterfaces) SetHostId(v int64)`

SetHostId sets HostId field to given value.


### GetStatus

`func (o *VMPoolHostInterfaces) GetStatus() VMPoolHostInterfaceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VMPoolHostInterfaces) GetStatusOk() (*VMPoolHostInterfaceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VMPoolHostInterfaces) SetStatus(v VMPoolHostInterfaceStatus)`

SetStatus sets Status field to given value.


### GetName

`func (o *VMPoolHostInterfaces) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VMPoolHostInterfaces) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VMPoolHostInterfaces) SetName(v string)`

SetName sets Name field to given value.


### GetMacAddress

`func (o *VMPoolHostInterfaces) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *VMPoolHostInterfaces) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *VMPoolHostInterfaces) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.


### GetFabric

`func (o *VMPoolHostInterfaces) GetFabric() GenericNetworkFabric`

GetFabric returns the Fabric field if non-nil, zero value otherwise.

### GetFabricOk

`func (o *VMPoolHostInterfaces) GetFabricOk() (*GenericNetworkFabric, bool)`

GetFabricOk returns a tuple with the Fabric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabric

`func (o *VMPoolHostInterfaces) SetFabric(v GenericNetworkFabric)`

SetFabric sets Fabric field to given value.

### HasFabric

`func (o *VMPoolHostInterfaces) HasFabric() bool`

HasFabric returns a boolean if a field has been set.

### GetNetworkDevices

`func (o *VMPoolHostInterfaces) GetNetworkDevices() []VMPoolHostInterfaceNetworkDevice`

GetNetworkDevices returns the NetworkDevices field if non-nil, zero value otherwise.

### GetNetworkDevicesOk

`func (o *VMPoolHostInterfaces) GetNetworkDevicesOk() (*[]VMPoolHostInterfaceNetworkDevice, bool)`

GetNetworkDevicesOk returns a tuple with the NetworkDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDevices

`func (o *VMPoolHostInterfaces) SetNetworkDevices(v []VMPoolHostInterfaceNetworkDevice)`

SetNetworkDevices sets NetworkDevices field to given value.

### HasNetworkDevices

`func (o *VMPoolHostInterfaces) HasNetworkDevices() bool`

HasNetworkDevices returns a boolean if a field has been set.

### GetLinks

`func (o *VMPoolHostInterfaces) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VMPoolHostInterfaces) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VMPoolHostInterfaces) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *VMPoolHostInterfaces) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


