# NetworkFabricInterconnectDeploymentCheckRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkIds** | Pointer to **[]int64** | IDs of the links to validate. If omitted or empty, all links of the interconnect are validated. | [optional] 

## Methods

### NewNetworkFabricInterconnectDeploymentCheckRequest

`func NewNetworkFabricInterconnectDeploymentCheckRequest() *NetworkFabricInterconnectDeploymentCheckRequest`

NewNetworkFabricInterconnectDeploymentCheckRequest instantiates a new NetworkFabricInterconnectDeploymentCheckRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFabricInterconnectDeploymentCheckRequestWithDefaults

`func NewNetworkFabricInterconnectDeploymentCheckRequestWithDefaults() *NetworkFabricInterconnectDeploymentCheckRequest`

NewNetworkFabricInterconnectDeploymentCheckRequestWithDefaults instantiates a new NetworkFabricInterconnectDeploymentCheckRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkIds

`func (o *NetworkFabricInterconnectDeploymentCheckRequest) GetLinkIds() []int64`

GetLinkIds returns the LinkIds field if non-nil, zero value otherwise.

### GetLinkIdsOk

`func (o *NetworkFabricInterconnectDeploymentCheckRequest) GetLinkIdsOk() (*[]int64, bool)`

GetLinkIdsOk returns a tuple with the LinkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkIds

`func (o *NetworkFabricInterconnectDeploymentCheckRequest) SetLinkIds(v []int64)`

SetLinkIds sets LinkIds field to given value.

### HasLinkIds

`func (o *NetworkFabricInterconnectDeploymentCheckRequest) HasLinkIds() bool`

HasLinkIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


