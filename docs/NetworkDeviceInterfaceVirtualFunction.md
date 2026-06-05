# NetworkDeviceInterfaceVirtualFunction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The ID of the virtual function | 
**NetworkDeviceId** | **int64** | The ID of the associated network device | 
**InterfaceId** | **int64** | The ID of the associated network equipment interface | 
**Name** | **string** | The name of the virtual function | 
**Index** | **float32** | The index of the virtual function | 
**Status** | **string** | The status of the virtual function | 
**LogicalNetworkId** | Pointer to **int64** | The ID of the logical network associated with this virtual function | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewNetworkDeviceInterfaceVirtualFunction

`func NewNetworkDeviceInterfaceVirtualFunction(id int64, networkDeviceId int64, interfaceId int64, name string, index float32, status string, ) *NetworkDeviceInterfaceVirtualFunction`

NewNetworkDeviceInterfaceVirtualFunction instantiates a new NetworkDeviceInterfaceVirtualFunction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceInterfaceVirtualFunctionWithDefaults

`func NewNetworkDeviceInterfaceVirtualFunctionWithDefaults() *NetworkDeviceInterfaceVirtualFunction`

NewNetworkDeviceInterfaceVirtualFunctionWithDefaults instantiates a new NetworkDeviceInterfaceVirtualFunction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkDeviceInterfaceVirtualFunction) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkDeviceInterfaceVirtualFunction) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkDeviceInterfaceVirtualFunction) SetId(v int64)`

SetId sets Id field to given value.


### GetNetworkDeviceId

`func (o *NetworkDeviceInterfaceVirtualFunction) GetNetworkDeviceId() int64`

GetNetworkDeviceId returns the NetworkDeviceId field if non-nil, zero value otherwise.

### GetNetworkDeviceIdOk

`func (o *NetworkDeviceInterfaceVirtualFunction) GetNetworkDeviceIdOk() (*int64, bool)`

GetNetworkDeviceIdOk returns a tuple with the NetworkDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceId

`func (o *NetworkDeviceInterfaceVirtualFunction) SetNetworkDeviceId(v int64)`

SetNetworkDeviceId sets NetworkDeviceId field to given value.


### GetInterfaceId

`func (o *NetworkDeviceInterfaceVirtualFunction) GetInterfaceId() int64`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *NetworkDeviceInterfaceVirtualFunction) GetInterfaceIdOk() (*int64, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *NetworkDeviceInterfaceVirtualFunction) SetInterfaceId(v int64)`

SetInterfaceId sets InterfaceId field to given value.


### GetName

`func (o *NetworkDeviceInterfaceVirtualFunction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkDeviceInterfaceVirtualFunction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkDeviceInterfaceVirtualFunction) SetName(v string)`

SetName sets Name field to given value.


### GetIndex

`func (o *NetworkDeviceInterfaceVirtualFunction) GetIndex() float32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *NetworkDeviceInterfaceVirtualFunction) GetIndexOk() (*float32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *NetworkDeviceInterfaceVirtualFunction) SetIndex(v float32)`

SetIndex sets Index field to given value.


### GetStatus

`func (o *NetworkDeviceInterfaceVirtualFunction) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkDeviceInterfaceVirtualFunction) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkDeviceInterfaceVirtualFunction) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLogicalNetworkId

`func (o *NetworkDeviceInterfaceVirtualFunction) GetLogicalNetworkId() int64`

GetLogicalNetworkId returns the LogicalNetworkId field if non-nil, zero value otherwise.

### GetLogicalNetworkIdOk

`func (o *NetworkDeviceInterfaceVirtualFunction) GetLogicalNetworkIdOk() (*int64, bool)`

GetLogicalNetworkIdOk returns a tuple with the LogicalNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkId

`func (o *NetworkDeviceInterfaceVirtualFunction) SetLogicalNetworkId(v int64)`

SetLogicalNetworkId sets LogicalNetworkId field to given value.

### HasLogicalNetworkId

`func (o *NetworkDeviceInterfaceVirtualFunction) HasLogicalNetworkId() bool`

HasLogicalNetworkId returns a boolean if a field has been set.

### GetLinks

`func (o *NetworkDeviceInterfaceVirtualFunction) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NetworkDeviceInterfaceVirtualFunction) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NetworkDeviceInterfaceVirtualFunction) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NetworkDeviceInterfaceVirtualFunction) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


