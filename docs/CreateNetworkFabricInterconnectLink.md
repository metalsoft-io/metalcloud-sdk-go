# CreateNetworkFabricInterconnectLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FabricId** | **int64** | The ID of the network fabric this interconnect link belongs to | 
**NetworkEquipmentId** | **int64** | The ID of the network equipment in the fabric | 

## Methods

### NewCreateNetworkFabricInterconnectLink

`func NewCreateNetworkFabricInterconnectLink(fabricId int64, networkEquipmentId int64, ) *CreateNetworkFabricInterconnectLink`

NewCreateNetworkFabricInterconnectLink instantiates a new CreateNetworkFabricInterconnectLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkFabricInterconnectLinkWithDefaults

`func NewCreateNetworkFabricInterconnectLinkWithDefaults() *CreateNetworkFabricInterconnectLink`

NewCreateNetworkFabricInterconnectLinkWithDefaults instantiates a new CreateNetworkFabricInterconnectLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFabricId

`func (o *CreateNetworkFabricInterconnectLink) GetFabricId() int64`

GetFabricId returns the FabricId field if non-nil, zero value otherwise.

### GetFabricIdOk

`func (o *CreateNetworkFabricInterconnectLink) GetFabricIdOk() (*int64, bool)`

GetFabricIdOk returns a tuple with the FabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricId

`func (o *CreateNetworkFabricInterconnectLink) SetFabricId(v int64)`

SetFabricId sets FabricId field to given value.


### GetNetworkEquipmentId

`func (o *CreateNetworkFabricInterconnectLink) GetNetworkEquipmentId() int64`

GetNetworkEquipmentId returns the NetworkEquipmentId field if non-nil, zero value otherwise.

### GetNetworkEquipmentIdOk

`func (o *CreateNetworkFabricInterconnectLink) GetNetworkEquipmentIdOk() (*int64, bool)`

GetNetworkEquipmentIdOk returns a tuple with the NetworkEquipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEquipmentId

`func (o *CreateNetworkFabricInterconnectLink) SetNetworkEquipmentId(v int64)`

SetNetworkEquipmentId sets NetworkEquipmentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


