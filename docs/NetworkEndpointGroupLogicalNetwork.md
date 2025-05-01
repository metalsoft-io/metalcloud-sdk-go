# NetworkEndpointGroupLogicalNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogicalNetworkId** | **string** | The logical network ID. | 
**Tagged** | **bool** | Whether the logical network is tagged. | 
**AccessMode** | [**NetworkEndpointGroupAllowedAccessMode**](NetworkEndpointGroupAllowedAccessMode.md) | The access mode of the network endpoint group | 
**Mtu** | Pointer to **int32** | The MTU of the logical network | [optional] 
**Redundancy** | Pointer to [**NullableRedundancyConfig**](RedundancyConfig.md) | The redundancy configuration | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 
**NetworkEndpointGroupId** | **string** | The network endpoint group ID. | 

## Methods

### NewNetworkEndpointGroupLogicalNetwork

`func NewNetworkEndpointGroupLogicalNetwork(logicalNetworkId string, tagged bool, accessMode NetworkEndpointGroupAllowedAccessMode, networkEndpointGroupId string, ) *NetworkEndpointGroupLogicalNetwork`

NewNetworkEndpointGroupLogicalNetwork instantiates a new NetworkEndpointGroupLogicalNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkEndpointGroupLogicalNetworkWithDefaults

`func NewNetworkEndpointGroupLogicalNetworkWithDefaults() *NetworkEndpointGroupLogicalNetwork`

NewNetworkEndpointGroupLogicalNetworkWithDefaults instantiates a new NetworkEndpointGroupLogicalNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogicalNetworkId

`func (o *NetworkEndpointGroupLogicalNetwork) GetLogicalNetworkId() string`

GetLogicalNetworkId returns the LogicalNetworkId field if non-nil, zero value otherwise.

### GetLogicalNetworkIdOk

`func (o *NetworkEndpointGroupLogicalNetwork) GetLogicalNetworkIdOk() (*string, bool)`

GetLogicalNetworkIdOk returns a tuple with the LogicalNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkId

`func (o *NetworkEndpointGroupLogicalNetwork) SetLogicalNetworkId(v string)`

SetLogicalNetworkId sets LogicalNetworkId field to given value.


### GetTagged

`func (o *NetworkEndpointGroupLogicalNetwork) GetTagged() bool`

GetTagged returns the Tagged field if non-nil, zero value otherwise.

### GetTaggedOk

`func (o *NetworkEndpointGroupLogicalNetwork) GetTaggedOk() (*bool, bool)`

GetTaggedOk returns a tuple with the Tagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagged

`func (o *NetworkEndpointGroupLogicalNetwork) SetTagged(v bool)`

SetTagged sets Tagged field to given value.


### GetAccessMode

`func (o *NetworkEndpointGroupLogicalNetwork) GetAccessMode() NetworkEndpointGroupAllowedAccessMode`

GetAccessMode returns the AccessMode field if non-nil, zero value otherwise.

### GetAccessModeOk

`func (o *NetworkEndpointGroupLogicalNetwork) GetAccessModeOk() (*NetworkEndpointGroupAllowedAccessMode, bool)`

GetAccessModeOk returns a tuple with the AccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMode

`func (o *NetworkEndpointGroupLogicalNetwork) SetAccessMode(v NetworkEndpointGroupAllowedAccessMode)`

SetAccessMode sets AccessMode field to given value.


### GetMtu

`func (o *NetworkEndpointGroupLogicalNetwork) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *NetworkEndpointGroupLogicalNetwork) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *NetworkEndpointGroupLogicalNetwork) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *NetworkEndpointGroupLogicalNetwork) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetRedundancy

`func (o *NetworkEndpointGroupLogicalNetwork) GetRedundancy() RedundancyConfig`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *NetworkEndpointGroupLogicalNetwork) GetRedundancyOk() (*RedundancyConfig, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *NetworkEndpointGroupLogicalNetwork) SetRedundancy(v RedundancyConfig)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *NetworkEndpointGroupLogicalNetwork) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### SetRedundancyNil

`func (o *NetworkEndpointGroupLogicalNetwork) SetRedundancyNil(b bool)`

 SetRedundancyNil sets the value for Redundancy to be an explicit nil

### UnsetRedundancy
`func (o *NetworkEndpointGroupLogicalNetwork) UnsetRedundancy()`

UnsetRedundancy ensures that no value is present for Redundancy, not even an explicit nil
### GetLinks

`func (o *NetworkEndpointGroupLogicalNetwork) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NetworkEndpointGroupLogicalNetwork) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NetworkEndpointGroupLogicalNetwork) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NetworkEndpointGroupLogicalNetwork) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetNetworkEndpointGroupId

`func (o *NetworkEndpointGroupLogicalNetwork) GetNetworkEndpointGroupId() string`

GetNetworkEndpointGroupId returns the NetworkEndpointGroupId field if non-nil, zero value otherwise.

### GetNetworkEndpointGroupIdOk

`func (o *NetworkEndpointGroupLogicalNetwork) GetNetworkEndpointGroupIdOk() (*string, bool)`

GetNetworkEndpointGroupIdOk returns a tuple with the NetworkEndpointGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEndpointGroupId

`func (o *NetworkEndpointGroupLogicalNetwork) SetNetworkEndpointGroupId(v string)`

SetNetworkEndpointGroupId sets NetworkEndpointGroupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


