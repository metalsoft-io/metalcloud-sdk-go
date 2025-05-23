# CreateVMInstanceGroupNetworkConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogicalNetworkId** | **string** | The logical network ID. | 
**Tagged** | **bool** | Whether the logical network is tagged. | 
**AccessMode** | [**NetworkEndpointGroupAllowedAccessMode**](NetworkEndpointGroupAllowedAccessMode.md) | The access mode of the network endpoint group | 
**Mtu** | Pointer to **int32** | The MTU of the logical network | [optional] 
**Redundancy** | Pointer to [**NullableRedundancyConfig**](RedundancyConfig.md) | The redundancy configuration | [optional] 

## Methods

### NewCreateVMInstanceGroupNetworkConnection

`func NewCreateVMInstanceGroupNetworkConnection(logicalNetworkId string, tagged bool, accessMode NetworkEndpointGroupAllowedAccessMode, ) *CreateVMInstanceGroupNetworkConnection`

NewCreateVMInstanceGroupNetworkConnection instantiates a new CreateVMInstanceGroupNetworkConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVMInstanceGroupNetworkConnectionWithDefaults

`func NewCreateVMInstanceGroupNetworkConnectionWithDefaults() *CreateVMInstanceGroupNetworkConnection`

NewCreateVMInstanceGroupNetworkConnectionWithDefaults instantiates a new CreateVMInstanceGroupNetworkConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogicalNetworkId

`func (o *CreateVMInstanceGroupNetworkConnection) GetLogicalNetworkId() string`

GetLogicalNetworkId returns the LogicalNetworkId field if non-nil, zero value otherwise.

### GetLogicalNetworkIdOk

`func (o *CreateVMInstanceGroupNetworkConnection) GetLogicalNetworkIdOk() (*string, bool)`

GetLogicalNetworkIdOk returns a tuple with the LogicalNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkId

`func (o *CreateVMInstanceGroupNetworkConnection) SetLogicalNetworkId(v string)`

SetLogicalNetworkId sets LogicalNetworkId field to given value.


### GetTagged

`func (o *CreateVMInstanceGroupNetworkConnection) GetTagged() bool`

GetTagged returns the Tagged field if non-nil, zero value otherwise.

### GetTaggedOk

`func (o *CreateVMInstanceGroupNetworkConnection) GetTaggedOk() (*bool, bool)`

GetTaggedOk returns a tuple with the Tagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagged

`func (o *CreateVMInstanceGroupNetworkConnection) SetTagged(v bool)`

SetTagged sets Tagged field to given value.


### GetAccessMode

`func (o *CreateVMInstanceGroupNetworkConnection) GetAccessMode() NetworkEndpointGroupAllowedAccessMode`

GetAccessMode returns the AccessMode field if non-nil, zero value otherwise.

### GetAccessModeOk

`func (o *CreateVMInstanceGroupNetworkConnection) GetAccessModeOk() (*NetworkEndpointGroupAllowedAccessMode, bool)`

GetAccessModeOk returns a tuple with the AccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMode

`func (o *CreateVMInstanceGroupNetworkConnection) SetAccessMode(v NetworkEndpointGroupAllowedAccessMode)`

SetAccessMode sets AccessMode field to given value.


### GetMtu

`func (o *CreateVMInstanceGroupNetworkConnection) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *CreateVMInstanceGroupNetworkConnection) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *CreateVMInstanceGroupNetworkConnection) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *CreateVMInstanceGroupNetworkConnection) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetRedundancy

`func (o *CreateVMInstanceGroupNetworkConnection) GetRedundancy() RedundancyConfig`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *CreateVMInstanceGroupNetworkConnection) GetRedundancyOk() (*RedundancyConfig, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *CreateVMInstanceGroupNetworkConnection) SetRedundancy(v RedundancyConfig)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *CreateVMInstanceGroupNetworkConnection) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### SetRedundancyNil

`func (o *CreateVMInstanceGroupNetworkConnection) SetRedundancyNil(b bool)`

 SetRedundancyNil sets the value for Redundancy to be an explicit nil

### UnsetRedundancy
`func (o *CreateVMInstanceGroupNetworkConnection) UnsetRedundancy()`

UnsetRedundancy ensures that no value is present for Redundancy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


