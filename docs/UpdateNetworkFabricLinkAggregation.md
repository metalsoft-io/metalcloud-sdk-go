# UpdateNetworkFabricLinkAggregation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomVariables** | Pointer to **map[string]interface{}** | Custom variables exposed to network device configuration templates as link_aggregation_custom_variables | [optional] 
**LinkIds** | **[]int64** | List of link IDs to be associated with the link aggregation | 
**Options** | Pointer to [**NetworkFabricLinkAggregationOptions**](NetworkFabricLinkAggregationOptions.md) | Link aggregation options | [optional] 

## Methods

### NewUpdateNetworkFabricLinkAggregation

`func NewUpdateNetworkFabricLinkAggregation(linkIds []int64, ) *UpdateNetworkFabricLinkAggregation`

NewUpdateNetworkFabricLinkAggregation instantiates a new UpdateNetworkFabricLinkAggregation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkFabricLinkAggregationWithDefaults

`func NewUpdateNetworkFabricLinkAggregationWithDefaults() *UpdateNetworkFabricLinkAggregation`

NewUpdateNetworkFabricLinkAggregationWithDefaults instantiates a new UpdateNetworkFabricLinkAggregation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomVariables

`func (o *UpdateNetworkFabricLinkAggregation) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *UpdateNetworkFabricLinkAggregation) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *UpdateNetworkFabricLinkAggregation) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *UpdateNetworkFabricLinkAggregation) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.

### GetLinkIds

`func (o *UpdateNetworkFabricLinkAggregation) GetLinkIds() []int64`

GetLinkIds returns the LinkIds field if non-nil, zero value otherwise.

### GetLinkIdsOk

`func (o *UpdateNetworkFabricLinkAggregation) GetLinkIdsOk() (*[]int64, bool)`

GetLinkIdsOk returns a tuple with the LinkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkIds

`func (o *UpdateNetworkFabricLinkAggregation) SetLinkIds(v []int64)`

SetLinkIds sets LinkIds field to given value.


### GetOptions

`func (o *UpdateNetworkFabricLinkAggregation) GetOptions() NetworkFabricLinkAggregationOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *UpdateNetworkFabricLinkAggregation) GetOptionsOk() (*NetworkFabricLinkAggregationOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *UpdateNetworkFabricLinkAggregation) SetOptions(v NetworkFabricLinkAggregationOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *UpdateNetworkFabricLinkAggregation) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


