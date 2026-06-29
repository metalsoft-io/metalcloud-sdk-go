# NetworkFabricInterconnectLinksDeployOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkIds** | **[]int64** | IDs of the interconnect links to activate, must be in draft state. After activation all links will be in active state. | 
**RequireConfirmation** | **bool** | Flag to indicate if the deployment should be confirmed before proceeding with the execution of template on network devices. | 

## Methods

### NewNetworkFabricInterconnectLinksDeployOptions

`func NewNetworkFabricInterconnectLinksDeployOptions(linkIds []int64, requireConfirmation bool, ) *NetworkFabricInterconnectLinksDeployOptions`

NewNetworkFabricInterconnectLinksDeployOptions instantiates a new NetworkFabricInterconnectLinksDeployOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFabricInterconnectLinksDeployOptionsWithDefaults

`func NewNetworkFabricInterconnectLinksDeployOptionsWithDefaults() *NetworkFabricInterconnectLinksDeployOptions`

NewNetworkFabricInterconnectLinksDeployOptionsWithDefaults instantiates a new NetworkFabricInterconnectLinksDeployOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkIds

`func (o *NetworkFabricInterconnectLinksDeployOptions) GetLinkIds() []int64`

GetLinkIds returns the LinkIds field if non-nil, zero value otherwise.

### GetLinkIdsOk

`func (o *NetworkFabricInterconnectLinksDeployOptions) GetLinkIdsOk() (*[]int64, bool)`

GetLinkIdsOk returns a tuple with the LinkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkIds

`func (o *NetworkFabricInterconnectLinksDeployOptions) SetLinkIds(v []int64)`

SetLinkIds sets LinkIds field to given value.


### GetRequireConfirmation

`func (o *NetworkFabricInterconnectLinksDeployOptions) GetRequireConfirmation() bool`

GetRequireConfirmation returns the RequireConfirmation field if non-nil, zero value otherwise.

### GetRequireConfirmationOk

`func (o *NetworkFabricInterconnectLinksDeployOptions) GetRequireConfirmationOk() (*bool, bool)`

GetRequireConfirmationOk returns a tuple with the RequireConfirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireConfirmation

`func (o *NetworkFabricInterconnectLinksDeployOptions) SetRequireConfirmation(v bool)`

SetRequireConfirmation sets RequireConfirmation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


