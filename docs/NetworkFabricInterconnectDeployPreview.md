# NetworkFabricInterconnectDeployPreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkDeviceId** | **int64** | Unique identifier for the network device. | 
**NetworkInterconnectId** | **int64** | Unique identifier for the network fabric interconnect | 
**GlobalTemplateActivate** | **[]string** | Preview of global activation template | 
**NeighborTemplateActivate** | **[]string** | Preview of BGP neighbor configuration activation template | 
**GlobalTemplateDeactivate** | **[]string** | Preview of global deactivation template | 
**NeighborTemplateDeactivate** | **[]string** | Preview of BGP neighbor configuration deactivation template | 
**Changes** | **[]string** | Apply-ready command list. Concatenation of the mode-appropriate generated templates (activate or deactivate side) with the sonic-cli / configure terminal prefix prepended once. Sent to switch.ssh_exec verbatim by the apply step. | 

## Methods

### NewNetworkFabricInterconnectDeployPreview

`func NewNetworkFabricInterconnectDeployPreview(networkDeviceId int64, networkInterconnectId int64, globalTemplateActivate []string, neighborTemplateActivate []string, globalTemplateDeactivate []string, neighborTemplateDeactivate []string, changes []string, ) *NetworkFabricInterconnectDeployPreview`

NewNetworkFabricInterconnectDeployPreview instantiates a new NetworkFabricInterconnectDeployPreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFabricInterconnectDeployPreviewWithDefaults

`func NewNetworkFabricInterconnectDeployPreviewWithDefaults() *NetworkFabricInterconnectDeployPreview`

NewNetworkFabricInterconnectDeployPreviewWithDefaults instantiates a new NetworkFabricInterconnectDeployPreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkDeviceId

`func (o *NetworkFabricInterconnectDeployPreview) GetNetworkDeviceId() int64`

GetNetworkDeviceId returns the NetworkDeviceId field if non-nil, zero value otherwise.

### GetNetworkDeviceIdOk

`func (o *NetworkFabricInterconnectDeployPreview) GetNetworkDeviceIdOk() (*int64, bool)`

GetNetworkDeviceIdOk returns a tuple with the NetworkDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceId

`func (o *NetworkFabricInterconnectDeployPreview) SetNetworkDeviceId(v int64)`

SetNetworkDeviceId sets NetworkDeviceId field to given value.


### GetNetworkInterconnectId

`func (o *NetworkFabricInterconnectDeployPreview) GetNetworkInterconnectId() int64`

GetNetworkInterconnectId returns the NetworkInterconnectId field if non-nil, zero value otherwise.

### GetNetworkInterconnectIdOk

`func (o *NetworkFabricInterconnectDeployPreview) GetNetworkInterconnectIdOk() (*int64, bool)`

GetNetworkInterconnectIdOk returns a tuple with the NetworkInterconnectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterconnectId

`func (o *NetworkFabricInterconnectDeployPreview) SetNetworkInterconnectId(v int64)`

SetNetworkInterconnectId sets NetworkInterconnectId field to given value.


### GetGlobalTemplateActivate

`func (o *NetworkFabricInterconnectDeployPreview) GetGlobalTemplateActivate() []string`

GetGlobalTemplateActivate returns the GlobalTemplateActivate field if non-nil, zero value otherwise.

### GetGlobalTemplateActivateOk

`func (o *NetworkFabricInterconnectDeployPreview) GetGlobalTemplateActivateOk() (*[]string, bool)`

GetGlobalTemplateActivateOk returns a tuple with the GlobalTemplateActivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalTemplateActivate

`func (o *NetworkFabricInterconnectDeployPreview) SetGlobalTemplateActivate(v []string)`

SetGlobalTemplateActivate sets GlobalTemplateActivate field to given value.


### GetNeighborTemplateActivate

`func (o *NetworkFabricInterconnectDeployPreview) GetNeighborTemplateActivate() []string`

GetNeighborTemplateActivate returns the NeighborTemplateActivate field if non-nil, zero value otherwise.

### GetNeighborTemplateActivateOk

`func (o *NetworkFabricInterconnectDeployPreview) GetNeighborTemplateActivateOk() (*[]string, bool)`

GetNeighborTemplateActivateOk returns a tuple with the NeighborTemplateActivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborTemplateActivate

`func (o *NetworkFabricInterconnectDeployPreview) SetNeighborTemplateActivate(v []string)`

SetNeighborTemplateActivate sets NeighborTemplateActivate field to given value.


### GetGlobalTemplateDeactivate

`func (o *NetworkFabricInterconnectDeployPreview) GetGlobalTemplateDeactivate() []string`

GetGlobalTemplateDeactivate returns the GlobalTemplateDeactivate field if non-nil, zero value otherwise.

### GetGlobalTemplateDeactivateOk

`func (o *NetworkFabricInterconnectDeployPreview) GetGlobalTemplateDeactivateOk() (*[]string, bool)`

GetGlobalTemplateDeactivateOk returns a tuple with the GlobalTemplateDeactivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalTemplateDeactivate

`func (o *NetworkFabricInterconnectDeployPreview) SetGlobalTemplateDeactivate(v []string)`

SetGlobalTemplateDeactivate sets GlobalTemplateDeactivate field to given value.


### GetNeighborTemplateDeactivate

`func (o *NetworkFabricInterconnectDeployPreview) GetNeighborTemplateDeactivate() []string`

GetNeighborTemplateDeactivate returns the NeighborTemplateDeactivate field if non-nil, zero value otherwise.

### GetNeighborTemplateDeactivateOk

`func (o *NetworkFabricInterconnectDeployPreview) GetNeighborTemplateDeactivateOk() (*[]string, bool)`

GetNeighborTemplateDeactivateOk returns a tuple with the NeighborTemplateDeactivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborTemplateDeactivate

`func (o *NetworkFabricInterconnectDeployPreview) SetNeighborTemplateDeactivate(v []string)`

SetNeighborTemplateDeactivate sets NeighborTemplateDeactivate field to given value.


### GetChanges

`func (o *NetworkFabricInterconnectDeployPreview) GetChanges() []string`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *NetworkFabricInterconnectDeployPreview) GetChangesOk() (*[]string, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *NetworkFabricInterconnectDeployPreview) SetChanges(v []string)`

SetChanges sets Changes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


