# UpdateNetworkEndpointGroupLogicalNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tagged** | Pointer to **bool** | Whether the logical network is tagged. | [optional] 
**AccessMode** | Pointer to [**NetworkEndpointGroupAllowedAccessMode**](NetworkEndpointGroupAllowedAccessMode.md) | The access mode of the network endpoint group | [optional] 
**Redundancy** | Pointer to [**NullableRedundancyConfig**](RedundancyConfig.md) | The redundancy configuration | [optional] 

## Methods

### NewUpdateNetworkEndpointGroupLogicalNetwork

`func NewUpdateNetworkEndpointGroupLogicalNetwork() *UpdateNetworkEndpointGroupLogicalNetwork`

NewUpdateNetworkEndpointGroupLogicalNetwork instantiates a new UpdateNetworkEndpointGroupLogicalNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkEndpointGroupLogicalNetworkWithDefaults

`func NewUpdateNetworkEndpointGroupLogicalNetworkWithDefaults() *UpdateNetworkEndpointGroupLogicalNetwork`

NewUpdateNetworkEndpointGroupLogicalNetworkWithDefaults instantiates a new UpdateNetworkEndpointGroupLogicalNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagged

`func (o *UpdateNetworkEndpointGroupLogicalNetwork) GetTagged() bool`

GetTagged returns the Tagged field if non-nil, zero value otherwise.

### GetTaggedOk

`func (o *UpdateNetworkEndpointGroupLogicalNetwork) GetTaggedOk() (*bool, bool)`

GetTaggedOk returns a tuple with the Tagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagged

`func (o *UpdateNetworkEndpointGroupLogicalNetwork) SetTagged(v bool)`

SetTagged sets Tagged field to given value.

### HasTagged

`func (o *UpdateNetworkEndpointGroupLogicalNetwork) HasTagged() bool`

HasTagged returns a boolean if a field has been set.

### GetAccessMode

`func (o *UpdateNetworkEndpointGroupLogicalNetwork) GetAccessMode() NetworkEndpointGroupAllowedAccessMode`

GetAccessMode returns the AccessMode field if non-nil, zero value otherwise.

### GetAccessModeOk

`func (o *UpdateNetworkEndpointGroupLogicalNetwork) GetAccessModeOk() (*NetworkEndpointGroupAllowedAccessMode, bool)`

GetAccessModeOk returns a tuple with the AccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMode

`func (o *UpdateNetworkEndpointGroupLogicalNetwork) SetAccessMode(v NetworkEndpointGroupAllowedAccessMode)`

SetAccessMode sets AccessMode field to given value.

### HasAccessMode

`func (o *UpdateNetworkEndpointGroupLogicalNetwork) HasAccessMode() bool`

HasAccessMode returns a boolean if a field has been set.

### GetRedundancy

`func (o *UpdateNetworkEndpointGroupLogicalNetwork) GetRedundancy() RedundancyConfig`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *UpdateNetworkEndpointGroupLogicalNetwork) GetRedundancyOk() (*RedundancyConfig, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *UpdateNetworkEndpointGroupLogicalNetwork) SetRedundancy(v RedundancyConfig)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *UpdateNetworkEndpointGroupLogicalNetwork) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### SetRedundancyNil

`func (o *UpdateNetworkEndpointGroupLogicalNetwork) SetRedundancyNil(b bool)`

 SetRedundancyNil sets the value for Redundancy to be an explicit nil

### UnsetRedundancy
`func (o *UpdateNetworkEndpointGroupLogicalNetwork) UnsetRedundancy()`

UnsetRedundancy ensures that no value is present for Redundancy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


