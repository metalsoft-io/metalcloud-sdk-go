# CreateServerInstanceGroupNetworkConnectionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogicalNetworkId** | **string** | The logical network ID. | 
**Tagged** | **bool** | Whether the logical network is tagged. | 
**AccessMode** | **string** | The access mode of the network endpoint group | 
**Redundancy** | Pointer to [**NullableRedundancyConfig**](RedundancyConfig.md) | The redundancy configuration | [optional] 

## Methods

### NewCreateServerInstanceGroupNetworkConnectionDto

`func NewCreateServerInstanceGroupNetworkConnectionDto(logicalNetworkId string, tagged bool, accessMode string, ) *CreateServerInstanceGroupNetworkConnectionDto`

NewCreateServerInstanceGroupNetworkConnectionDto instantiates a new CreateServerInstanceGroupNetworkConnectionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServerInstanceGroupNetworkConnectionDtoWithDefaults

`func NewCreateServerInstanceGroupNetworkConnectionDtoWithDefaults() *CreateServerInstanceGroupNetworkConnectionDto`

NewCreateServerInstanceGroupNetworkConnectionDtoWithDefaults instantiates a new CreateServerInstanceGroupNetworkConnectionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogicalNetworkId

`func (o *CreateServerInstanceGroupNetworkConnectionDto) GetLogicalNetworkId() string`

GetLogicalNetworkId returns the LogicalNetworkId field if non-nil, zero value otherwise.

### GetLogicalNetworkIdOk

`func (o *CreateServerInstanceGroupNetworkConnectionDto) GetLogicalNetworkIdOk() (*string, bool)`

GetLogicalNetworkIdOk returns a tuple with the LogicalNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkId

`func (o *CreateServerInstanceGroupNetworkConnectionDto) SetLogicalNetworkId(v string)`

SetLogicalNetworkId sets LogicalNetworkId field to given value.


### GetTagged

`func (o *CreateServerInstanceGroupNetworkConnectionDto) GetTagged() bool`

GetTagged returns the Tagged field if non-nil, zero value otherwise.

### GetTaggedOk

`func (o *CreateServerInstanceGroupNetworkConnectionDto) GetTaggedOk() (*bool, bool)`

GetTaggedOk returns a tuple with the Tagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagged

`func (o *CreateServerInstanceGroupNetworkConnectionDto) SetTagged(v bool)`

SetTagged sets Tagged field to given value.


### GetAccessMode

`func (o *CreateServerInstanceGroupNetworkConnectionDto) GetAccessMode() string`

GetAccessMode returns the AccessMode field if non-nil, zero value otherwise.

### GetAccessModeOk

`func (o *CreateServerInstanceGroupNetworkConnectionDto) GetAccessModeOk() (*string, bool)`

GetAccessModeOk returns a tuple with the AccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMode

`func (o *CreateServerInstanceGroupNetworkConnectionDto) SetAccessMode(v string)`

SetAccessMode sets AccessMode field to given value.


### GetRedundancy

`func (o *CreateServerInstanceGroupNetworkConnectionDto) GetRedundancy() RedundancyConfig`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *CreateServerInstanceGroupNetworkConnectionDto) GetRedundancyOk() (*RedundancyConfig, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *CreateServerInstanceGroupNetworkConnectionDto) SetRedundancy(v RedundancyConfig)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *CreateServerInstanceGroupNetworkConnectionDto) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### SetRedundancyNil

`func (o *CreateServerInstanceGroupNetworkConnectionDto) SetRedundancyNil(b bool)`

 SetRedundancyNil sets the value for Redundancy to be an explicit nil

### UnsetRedundancy
`func (o *CreateServerInstanceGroupNetworkConnectionDto) UnsetRedundancy()`

UnsetRedundancy ensures that no value is present for Redundancy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


