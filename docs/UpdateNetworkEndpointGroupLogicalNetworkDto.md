# UpdateNetworkEndpointGroupLogicalNetworkDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tagged** | Pointer to **bool** | Whether the logical network is tagged. | [optional] 
**AccessMode** | Pointer to **string** | The access mode of the network endpoint group | [optional] 
**Redundancy** | Pointer to [**NullableRedundancyConfig**](RedundancyConfig.md) | The redundancy configuration | [optional] 

## Methods

### NewUpdateNetworkEndpointGroupLogicalNetworkDto

`func NewUpdateNetworkEndpointGroupLogicalNetworkDto() *UpdateNetworkEndpointGroupLogicalNetworkDto`

NewUpdateNetworkEndpointGroupLogicalNetworkDto instantiates a new UpdateNetworkEndpointGroupLogicalNetworkDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkEndpointGroupLogicalNetworkDtoWithDefaults

`func NewUpdateNetworkEndpointGroupLogicalNetworkDtoWithDefaults() *UpdateNetworkEndpointGroupLogicalNetworkDto`

NewUpdateNetworkEndpointGroupLogicalNetworkDtoWithDefaults instantiates a new UpdateNetworkEndpointGroupLogicalNetworkDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagged

`func (o *UpdateNetworkEndpointGroupLogicalNetworkDto) GetTagged() bool`

GetTagged returns the Tagged field if non-nil, zero value otherwise.

### GetTaggedOk

`func (o *UpdateNetworkEndpointGroupLogicalNetworkDto) GetTaggedOk() (*bool, bool)`

GetTaggedOk returns a tuple with the Tagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagged

`func (o *UpdateNetworkEndpointGroupLogicalNetworkDto) SetTagged(v bool)`

SetTagged sets Tagged field to given value.

### HasTagged

`func (o *UpdateNetworkEndpointGroupLogicalNetworkDto) HasTagged() bool`

HasTagged returns a boolean if a field has been set.

### GetAccessMode

`func (o *UpdateNetworkEndpointGroupLogicalNetworkDto) GetAccessMode() string`

GetAccessMode returns the AccessMode field if non-nil, zero value otherwise.

### GetAccessModeOk

`func (o *UpdateNetworkEndpointGroupLogicalNetworkDto) GetAccessModeOk() (*string, bool)`

GetAccessModeOk returns a tuple with the AccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMode

`func (o *UpdateNetworkEndpointGroupLogicalNetworkDto) SetAccessMode(v string)`

SetAccessMode sets AccessMode field to given value.

### HasAccessMode

`func (o *UpdateNetworkEndpointGroupLogicalNetworkDto) HasAccessMode() bool`

HasAccessMode returns a boolean if a field has been set.

### GetRedundancy

`func (o *UpdateNetworkEndpointGroupLogicalNetworkDto) GetRedundancy() RedundancyConfig`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *UpdateNetworkEndpointGroupLogicalNetworkDto) GetRedundancyOk() (*RedundancyConfig, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *UpdateNetworkEndpointGroupLogicalNetworkDto) SetRedundancy(v RedundancyConfig)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *UpdateNetworkEndpointGroupLogicalNetworkDto) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### SetRedundancyNil

`func (o *UpdateNetworkEndpointGroupLogicalNetworkDto) SetRedundancyNil(b bool)`

 SetRedundancyNil sets the value for Redundancy to be an explicit nil

### UnsetRedundancy
`func (o *UpdateNetworkEndpointGroupLogicalNetworkDto) UnsetRedundancy()`

UnsetRedundancy ensures that no value is present for Redundancy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


