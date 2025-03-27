# NetworkEndpointGroupLogicalNetworkDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogicalNetworkId** | **string** | The logical network ID. | 
**Tagged** | **bool** | Whether the logical network is tagged. | 
**AccessMode** | **string** | The access mode of the network endpoint group | 
**Redundancy** | Pointer to [**NullableRedundancyConfig**](RedundancyConfig.md) | The redundancy configuration | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 
**NetworkEndpointGroupId** | **string** | The network endpoint group ID. | 

## Methods

### NewNetworkEndpointGroupLogicalNetworkDto

`func NewNetworkEndpointGroupLogicalNetworkDto(logicalNetworkId string, tagged bool, accessMode string, networkEndpointGroupId string, ) *NetworkEndpointGroupLogicalNetworkDto`

NewNetworkEndpointGroupLogicalNetworkDto instantiates a new NetworkEndpointGroupLogicalNetworkDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkEndpointGroupLogicalNetworkDtoWithDefaults

`func NewNetworkEndpointGroupLogicalNetworkDtoWithDefaults() *NetworkEndpointGroupLogicalNetworkDto`

NewNetworkEndpointGroupLogicalNetworkDtoWithDefaults instantiates a new NetworkEndpointGroupLogicalNetworkDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogicalNetworkId

`func (o *NetworkEndpointGroupLogicalNetworkDto) GetLogicalNetworkId() string`

GetLogicalNetworkId returns the LogicalNetworkId field if non-nil, zero value otherwise.

### GetLogicalNetworkIdOk

`func (o *NetworkEndpointGroupLogicalNetworkDto) GetLogicalNetworkIdOk() (*string, bool)`

GetLogicalNetworkIdOk returns a tuple with the LogicalNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkId

`func (o *NetworkEndpointGroupLogicalNetworkDto) SetLogicalNetworkId(v string)`

SetLogicalNetworkId sets LogicalNetworkId field to given value.


### GetTagged

`func (o *NetworkEndpointGroupLogicalNetworkDto) GetTagged() bool`

GetTagged returns the Tagged field if non-nil, zero value otherwise.

### GetTaggedOk

`func (o *NetworkEndpointGroupLogicalNetworkDto) GetTaggedOk() (*bool, bool)`

GetTaggedOk returns a tuple with the Tagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagged

`func (o *NetworkEndpointGroupLogicalNetworkDto) SetTagged(v bool)`

SetTagged sets Tagged field to given value.


### GetAccessMode

`func (o *NetworkEndpointGroupLogicalNetworkDto) GetAccessMode() string`

GetAccessMode returns the AccessMode field if non-nil, zero value otherwise.

### GetAccessModeOk

`func (o *NetworkEndpointGroupLogicalNetworkDto) GetAccessModeOk() (*string, bool)`

GetAccessModeOk returns a tuple with the AccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMode

`func (o *NetworkEndpointGroupLogicalNetworkDto) SetAccessMode(v string)`

SetAccessMode sets AccessMode field to given value.


### GetRedundancy

`func (o *NetworkEndpointGroupLogicalNetworkDto) GetRedundancy() RedundancyConfig`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *NetworkEndpointGroupLogicalNetworkDto) GetRedundancyOk() (*RedundancyConfig, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *NetworkEndpointGroupLogicalNetworkDto) SetRedundancy(v RedundancyConfig)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *NetworkEndpointGroupLogicalNetworkDto) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### SetRedundancyNil

`func (o *NetworkEndpointGroupLogicalNetworkDto) SetRedundancyNil(b bool)`

 SetRedundancyNil sets the value for Redundancy to be an explicit nil

### UnsetRedundancy
`func (o *NetworkEndpointGroupLogicalNetworkDto) UnsetRedundancy()`

UnsetRedundancy ensures that no value is present for Redundancy, not even an explicit nil
### GetLinks

`func (o *NetworkEndpointGroupLogicalNetworkDto) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NetworkEndpointGroupLogicalNetworkDto) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NetworkEndpointGroupLogicalNetworkDto) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NetworkEndpointGroupLogicalNetworkDto) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetNetworkEndpointGroupId

`func (o *NetworkEndpointGroupLogicalNetworkDto) GetNetworkEndpointGroupId() string`

GetNetworkEndpointGroupId returns the NetworkEndpointGroupId field if non-nil, zero value otherwise.

### GetNetworkEndpointGroupIdOk

`func (o *NetworkEndpointGroupLogicalNetworkDto) GetNetworkEndpointGroupIdOk() (*string, bool)`

GetNetworkEndpointGroupIdOk returns a tuple with the NetworkEndpointGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEndpointGroupId

`func (o *NetworkEndpointGroupLogicalNetworkDto) SetNetworkEndpointGroupId(v string)`

SetNetworkEndpointGroupId sets NetworkEndpointGroupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


