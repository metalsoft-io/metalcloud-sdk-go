# UpdateVMInstanceGroupInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | Interface ID | 
**NetworkId** | Pointer to **float32** | Network ID for the interface. Can be null if the interface is not connected to any network | [optional] 

## Methods

### NewUpdateVMInstanceGroupInterface

`func NewUpdateVMInstanceGroupInterface(id float32, ) *UpdateVMInstanceGroupInterface`

NewUpdateVMInstanceGroupInterface instantiates a new UpdateVMInstanceGroupInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVMInstanceGroupInterfaceWithDefaults

`func NewUpdateVMInstanceGroupInterfaceWithDefaults() *UpdateVMInstanceGroupInterface`

NewUpdateVMInstanceGroupInterfaceWithDefaults instantiates a new UpdateVMInstanceGroupInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateVMInstanceGroupInterface) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateVMInstanceGroupInterface) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateVMInstanceGroupInterface) SetId(v float32)`

SetId sets Id field to given value.


### GetNetworkId

`func (o *UpdateVMInstanceGroupInterface) GetNetworkId() float32`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *UpdateVMInstanceGroupInterface) GetNetworkIdOk() (*float32, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *UpdateVMInstanceGroupInterface) SetNetworkId(v float32)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *UpdateVMInstanceGroupInterface) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


