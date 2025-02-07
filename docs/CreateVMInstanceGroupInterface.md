# CreateVMInstanceGroupInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **float32** | Network ID for the interface. Can be null if the interface is not connected to any network | [optional] 

## Methods

### NewCreateVMInstanceGroupInterface

`func NewCreateVMInstanceGroupInterface() *CreateVMInstanceGroupInterface`

NewCreateVMInstanceGroupInterface instantiates a new CreateVMInstanceGroupInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVMInstanceGroupInterfaceWithDefaults

`func NewCreateVMInstanceGroupInterfaceWithDefaults() *CreateVMInstanceGroupInterface`

NewCreateVMInstanceGroupInterfaceWithDefaults instantiates a new CreateVMInstanceGroupInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *CreateVMInstanceGroupInterface) GetNetworkId() float32`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *CreateVMInstanceGroupInterface) GetNetworkIdOk() (*float32, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *CreateVMInstanceGroupInterface) SetNetworkId(v float32)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *CreateVMInstanceGroupInterface) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


