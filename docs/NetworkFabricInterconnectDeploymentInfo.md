# NetworkFabricInterconnectDeploymentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**InterconnectStatus**](InterconnectStatus.md) | The status of the network interconnect, by default it is in draft mode. | [optional] 
**DeployId** | Pointer to **int64** | The deploy ID of the network fabric interconnect, if it is being deployed. | [optional] 
**DeployPreview** | Pointer to [**[]NetworkFabricInterconnectDeployPreview**](NetworkFabricInterconnectDeployPreview.md) | The deploy preview for the network fabric interconnect, if it is being deployed. | [optional] 

## Methods

### NewNetworkFabricInterconnectDeploymentInfo

`func NewNetworkFabricInterconnectDeploymentInfo() *NetworkFabricInterconnectDeploymentInfo`

NewNetworkFabricInterconnectDeploymentInfo instantiates a new NetworkFabricInterconnectDeploymentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFabricInterconnectDeploymentInfoWithDefaults

`func NewNetworkFabricInterconnectDeploymentInfoWithDefaults() *NetworkFabricInterconnectDeploymentInfo`

NewNetworkFabricInterconnectDeploymentInfoWithDefaults instantiates a new NetworkFabricInterconnectDeploymentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *NetworkFabricInterconnectDeploymentInfo) GetStatus() InterconnectStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkFabricInterconnectDeploymentInfo) GetStatusOk() (*InterconnectStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkFabricInterconnectDeploymentInfo) SetStatus(v InterconnectStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkFabricInterconnectDeploymentInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDeployId

`func (o *NetworkFabricInterconnectDeploymentInfo) GetDeployId() int64`

GetDeployId returns the DeployId field if non-nil, zero value otherwise.

### GetDeployIdOk

`func (o *NetworkFabricInterconnectDeploymentInfo) GetDeployIdOk() (*int64, bool)`

GetDeployIdOk returns a tuple with the DeployId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployId

`func (o *NetworkFabricInterconnectDeploymentInfo) SetDeployId(v int64)`

SetDeployId sets DeployId field to given value.

### HasDeployId

`func (o *NetworkFabricInterconnectDeploymentInfo) HasDeployId() bool`

HasDeployId returns a boolean if a field has been set.

### GetDeployPreview

`func (o *NetworkFabricInterconnectDeploymentInfo) GetDeployPreview() []NetworkFabricInterconnectDeployPreview`

GetDeployPreview returns the DeployPreview field if non-nil, zero value otherwise.

### GetDeployPreviewOk

`func (o *NetworkFabricInterconnectDeploymentInfo) GetDeployPreviewOk() (*[]NetworkFabricInterconnectDeployPreview, bool)`

GetDeployPreviewOk returns a tuple with the DeployPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployPreview

`func (o *NetworkFabricInterconnectDeploymentInfo) SetDeployPreview(v []NetworkFabricInterconnectDeployPreview)`

SetDeployPreview sets DeployPreview field to given value.

### HasDeployPreview

`func (o *NetworkFabricInterconnectDeploymentInfo) HasDeployPreview() bool`

HasDeployPreview returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


