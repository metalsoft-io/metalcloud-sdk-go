# NetworkFabricInterconnectDeactivateLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkIds** | **[]int64** | IDs of the ACTIVE interconnect links to deactivate. Links will be put back in DRAFT state. | 
**RequireConfirmation** | **bool** | Flag to indicate if the deactivation should be confirmed before proceeding with the execution of templates on network devices. | 

## Methods

### NewNetworkFabricInterconnectDeactivateLinks

`func NewNetworkFabricInterconnectDeactivateLinks(linkIds []int64, requireConfirmation bool, ) *NetworkFabricInterconnectDeactivateLinks`

NewNetworkFabricInterconnectDeactivateLinks instantiates a new NetworkFabricInterconnectDeactivateLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFabricInterconnectDeactivateLinksWithDefaults

`func NewNetworkFabricInterconnectDeactivateLinksWithDefaults() *NetworkFabricInterconnectDeactivateLinks`

NewNetworkFabricInterconnectDeactivateLinksWithDefaults instantiates a new NetworkFabricInterconnectDeactivateLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkIds

`func (o *NetworkFabricInterconnectDeactivateLinks) GetLinkIds() []int64`

GetLinkIds returns the LinkIds field if non-nil, zero value otherwise.

### GetLinkIdsOk

`func (o *NetworkFabricInterconnectDeactivateLinks) GetLinkIdsOk() (*[]int64, bool)`

GetLinkIdsOk returns a tuple with the LinkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkIds

`func (o *NetworkFabricInterconnectDeactivateLinks) SetLinkIds(v []int64)`

SetLinkIds sets LinkIds field to given value.


### GetRequireConfirmation

`func (o *NetworkFabricInterconnectDeactivateLinks) GetRequireConfirmation() bool`

GetRequireConfirmation returns the RequireConfirmation field if non-nil, zero value otherwise.

### GetRequireConfirmationOk

`func (o *NetworkFabricInterconnectDeactivateLinks) GetRequireConfirmationOk() (*bool, bool)`

GetRequireConfirmationOk returns a tuple with the RequireConfirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireConfirmation

`func (o *NetworkFabricInterconnectDeactivateLinks) SetRequireConfirmation(v bool)`

SetRequireConfirmation sets RequireConfirmation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


