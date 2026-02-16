# NetworkFabricLinkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkFabricLinkId** | **float32** | Unique identifier for the network fabric link | 
**NetworkFabricLinkAggregationId** | Pointer to **float32** | Unique identifier for the network fabric link aggregation to be associated with that link | [optional] 
**NetworkFabricLink** | Pointer to [**NetworkFabricLink**](NetworkFabricLink.md) | Associated network fabric link | [optional] 
**LinkAggregation** | Pointer to [**NetworkFabricLinkAggregation**](NetworkFabricLinkAggregation.md) | Network Fabric Link Aggregation associated with this link | [optional] 
**Status** | **string** | Status of the network fabric link configuration | 

## Methods

### NewNetworkFabricLinkConfig

`func NewNetworkFabricLinkConfig(networkFabricLinkId float32, status string, ) *NetworkFabricLinkConfig`

NewNetworkFabricLinkConfig instantiates a new NetworkFabricLinkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFabricLinkConfigWithDefaults

`func NewNetworkFabricLinkConfigWithDefaults() *NetworkFabricLinkConfig`

NewNetworkFabricLinkConfigWithDefaults instantiates a new NetworkFabricLinkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkFabricLinkId

`func (o *NetworkFabricLinkConfig) GetNetworkFabricLinkId() float32`

GetNetworkFabricLinkId returns the NetworkFabricLinkId field if non-nil, zero value otherwise.

### GetNetworkFabricLinkIdOk

`func (o *NetworkFabricLinkConfig) GetNetworkFabricLinkIdOk() (*float32, bool)`

GetNetworkFabricLinkIdOk returns a tuple with the NetworkFabricLinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFabricLinkId

`func (o *NetworkFabricLinkConfig) SetNetworkFabricLinkId(v float32)`

SetNetworkFabricLinkId sets NetworkFabricLinkId field to given value.


### GetNetworkFabricLinkAggregationId

`func (o *NetworkFabricLinkConfig) GetNetworkFabricLinkAggregationId() float32`

GetNetworkFabricLinkAggregationId returns the NetworkFabricLinkAggregationId field if non-nil, zero value otherwise.

### GetNetworkFabricLinkAggregationIdOk

`func (o *NetworkFabricLinkConfig) GetNetworkFabricLinkAggregationIdOk() (*float32, bool)`

GetNetworkFabricLinkAggregationIdOk returns a tuple with the NetworkFabricLinkAggregationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFabricLinkAggregationId

`func (o *NetworkFabricLinkConfig) SetNetworkFabricLinkAggregationId(v float32)`

SetNetworkFabricLinkAggregationId sets NetworkFabricLinkAggregationId field to given value.

### HasNetworkFabricLinkAggregationId

`func (o *NetworkFabricLinkConfig) HasNetworkFabricLinkAggregationId() bool`

HasNetworkFabricLinkAggregationId returns a boolean if a field has been set.

### GetNetworkFabricLink

`func (o *NetworkFabricLinkConfig) GetNetworkFabricLink() NetworkFabricLink`

GetNetworkFabricLink returns the NetworkFabricLink field if non-nil, zero value otherwise.

### GetNetworkFabricLinkOk

`func (o *NetworkFabricLinkConfig) GetNetworkFabricLinkOk() (*NetworkFabricLink, bool)`

GetNetworkFabricLinkOk returns a tuple with the NetworkFabricLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFabricLink

`func (o *NetworkFabricLinkConfig) SetNetworkFabricLink(v NetworkFabricLink)`

SetNetworkFabricLink sets NetworkFabricLink field to given value.

### HasNetworkFabricLink

`func (o *NetworkFabricLinkConfig) HasNetworkFabricLink() bool`

HasNetworkFabricLink returns a boolean if a field has been set.

### GetLinkAggregation

`func (o *NetworkFabricLinkConfig) GetLinkAggregation() NetworkFabricLinkAggregation`

GetLinkAggregation returns the LinkAggregation field if non-nil, zero value otherwise.

### GetLinkAggregationOk

`func (o *NetworkFabricLinkConfig) GetLinkAggregationOk() (*NetworkFabricLinkAggregation, bool)`

GetLinkAggregationOk returns a tuple with the LinkAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkAggregation

`func (o *NetworkFabricLinkConfig) SetLinkAggregation(v NetworkFabricLinkAggregation)`

SetLinkAggregation sets LinkAggregation field to given value.

### HasLinkAggregation

`func (o *NetworkFabricLinkConfig) HasLinkAggregation() bool`

HasLinkAggregation returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkFabricLinkConfig) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkFabricLinkConfig) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkFabricLinkConfig) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


