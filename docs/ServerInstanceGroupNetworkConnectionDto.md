# ServerInstanceGroupNetworkConnectionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tagged** | **bool** | Whether the logical network is tagged. | 
**AccessMode** | **string** | The access mode of the network endpoint group | 
**Redundancy** | Pointer to [**NullableRedundancyConfig**](RedundancyConfig.md) | The redundancy configuration | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 
**Id** | **string** | The network connection ID. | 

## Methods

### NewServerInstanceGroupNetworkConnectionDto

`func NewServerInstanceGroupNetworkConnectionDto(tagged bool, accessMode string, id string, ) *ServerInstanceGroupNetworkConnectionDto`

NewServerInstanceGroupNetworkConnectionDto instantiates a new ServerInstanceGroupNetworkConnectionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceGroupNetworkConnectionDtoWithDefaults

`func NewServerInstanceGroupNetworkConnectionDtoWithDefaults() *ServerInstanceGroupNetworkConnectionDto`

NewServerInstanceGroupNetworkConnectionDtoWithDefaults instantiates a new ServerInstanceGroupNetworkConnectionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagged

`func (o *ServerInstanceGroupNetworkConnectionDto) GetTagged() bool`

GetTagged returns the Tagged field if non-nil, zero value otherwise.

### GetTaggedOk

`func (o *ServerInstanceGroupNetworkConnectionDto) GetTaggedOk() (*bool, bool)`

GetTaggedOk returns a tuple with the Tagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagged

`func (o *ServerInstanceGroupNetworkConnectionDto) SetTagged(v bool)`

SetTagged sets Tagged field to given value.


### GetAccessMode

`func (o *ServerInstanceGroupNetworkConnectionDto) GetAccessMode() string`

GetAccessMode returns the AccessMode field if non-nil, zero value otherwise.

### GetAccessModeOk

`func (o *ServerInstanceGroupNetworkConnectionDto) GetAccessModeOk() (*string, bool)`

GetAccessModeOk returns a tuple with the AccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMode

`func (o *ServerInstanceGroupNetworkConnectionDto) SetAccessMode(v string)`

SetAccessMode sets AccessMode field to given value.


### GetRedundancy

`func (o *ServerInstanceGroupNetworkConnectionDto) GetRedundancy() RedundancyConfig`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *ServerInstanceGroupNetworkConnectionDto) GetRedundancyOk() (*RedundancyConfig, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *ServerInstanceGroupNetworkConnectionDto) SetRedundancy(v RedundancyConfig)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *ServerInstanceGroupNetworkConnectionDto) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### SetRedundancyNil

`func (o *ServerInstanceGroupNetworkConnectionDto) SetRedundancyNil(b bool)`

 SetRedundancyNil sets the value for Redundancy to be an explicit nil

### UnsetRedundancy
`func (o *ServerInstanceGroupNetworkConnectionDto) UnsetRedundancy()`

UnsetRedundancy ensures that no value is present for Redundancy, not even an explicit nil
### GetLinks

`func (o *ServerInstanceGroupNetworkConnectionDto) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServerInstanceGroupNetworkConnectionDto) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServerInstanceGroupNetworkConnectionDto) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServerInstanceGroupNetworkConnectionDto) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *ServerInstanceGroupNetworkConnectionDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerInstanceGroupNetworkConnectionDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerInstanceGroupNetworkConnectionDto) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


