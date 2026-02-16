# CreateNetworkFabricLinkAggregation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the network fabric link aggregation | 
**LinkIds** | **[]float32** | List of link IDs to be associated with the link aggregation | 

## Methods

### NewCreateNetworkFabricLinkAggregation

`func NewCreateNetworkFabricLinkAggregation(type_ string, linkIds []float32, ) *CreateNetworkFabricLinkAggregation`

NewCreateNetworkFabricLinkAggregation instantiates a new CreateNetworkFabricLinkAggregation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkFabricLinkAggregationWithDefaults

`func NewCreateNetworkFabricLinkAggregationWithDefaults() *CreateNetworkFabricLinkAggregation`

NewCreateNetworkFabricLinkAggregationWithDefaults instantiates a new CreateNetworkFabricLinkAggregation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateNetworkFabricLinkAggregation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateNetworkFabricLinkAggregation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateNetworkFabricLinkAggregation) SetType(v string)`

SetType sets Type field to given value.


### GetLinkIds

`func (o *CreateNetworkFabricLinkAggregation) GetLinkIds() []float32`

GetLinkIds returns the LinkIds field if non-nil, zero value otherwise.

### GetLinkIdsOk

`func (o *CreateNetworkFabricLinkAggregation) GetLinkIdsOk() (*[]float32, bool)`

GetLinkIdsOk returns a tuple with the LinkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkIds

`func (o *CreateNetworkFabricLinkAggregation) SetLinkIds(v []float32)`

SetLinkIds sets LinkIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


