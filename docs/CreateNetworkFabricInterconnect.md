# CreateNetworkFabricInterconnect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterconnectType** | [**NetworkFabricInterconnectType**](NetworkFabricInterconnectType.md) | The network fabric interconnect type | 
**Label** | **string** | Unique label for the network fabric interconnect | 
**Name** | Pointer to **string** | Name of the network fabric interconnect | [optional] 
**Description** | Pointer to **string** | Short description of the network fabric interconnect | [optional] 
**BgpConfigurationTemplate** | Pointer to **string** | BGP configuration template for the interconnect | [optional] 
**BgpNeighborTemplate** | Pointer to **string** | BGP neighbor template for the interconnect | [optional] 
**TransportId** | Pointer to **int32** | Transport ID allocated for this interconnect, if any. Allocated from range 65534-1 (descending) | [optional] 

## Methods

### NewCreateNetworkFabricInterconnect

`func NewCreateNetworkFabricInterconnect(interconnectType NetworkFabricInterconnectType, label string, ) *CreateNetworkFabricInterconnect`

NewCreateNetworkFabricInterconnect instantiates a new CreateNetworkFabricInterconnect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkFabricInterconnectWithDefaults

`func NewCreateNetworkFabricInterconnectWithDefaults() *CreateNetworkFabricInterconnect`

NewCreateNetworkFabricInterconnectWithDefaults instantiates a new CreateNetworkFabricInterconnect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterconnectType

`func (o *CreateNetworkFabricInterconnect) GetInterconnectType() NetworkFabricInterconnectType`

GetInterconnectType returns the InterconnectType field if non-nil, zero value otherwise.

### GetInterconnectTypeOk

`func (o *CreateNetworkFabricInterconnect) GetInterconnectTypeOk() (*NetworkFabricInterconnectType, bool)`

GetInterconnectTypeOk returns a tuple with the InterconnectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterconnectType

`func (o *CreateNetworkFabricInterconnect) SetInterconnectType(v NetworkFabricInterconnectType)`

SetInterconnectType sets InterconnectType field to given value.


### GetLabel

`func (o *CreateNetworkFabricInterconnect) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateNetworkFabricInterconnect) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateNetworkFabricInterconnect) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *CreateNetworkFabricInterconnect) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkFabricInterconnect) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkFabricInterconnect) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateNetworkFabricInterconnect) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateNetworkFabricInterconnect) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateNetworkFabricInterconnect) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateNetworkFabricInterconnect) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateNetworkFabricInterconnect) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBgpConfigurationTemplate

`func (o *CreateNetworkFabricInterconnect) GetBgpConfigurationTemplate() string`

GetBgpConfigurationTemplate returns the BgpConfigurationTemplate field if non-nil, zero value otherwise.

### GetBgpConfigurationTemplateOk

`func (o *CreateNetworkFabricInterconnect) GetBgpConfigurationTemplateOk() (*string, bool)`

GetBgpConfigurationTemplateOk returns a tuple with the BgpConfigurationTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpConfigurationTemplate

`func (o *CreateNetworkFabricInterconnect) SetBgpConfigurationTemplate(v string)`

SetBgpConfigurationTemplate sets BgpConfigurationTemplate field to given value.

### HasBgpConfigurationTemplate

`func (o *CreateNetworkFabricInterconnect) HasBgpConfigurationTemplate() bool`

HasBgpConfigurationTemplate returns a boolean if a field has been set.

### GetBgpNeighborTemplate

`func (o *CreateNetworkFabricInterconnect) GetBgpNeighborTemplate() string`

GetBgpNeighborTemplate returns the BgpNeighborTemplate field if non-nil, zero value otherwise.

### GetBgpNeighborTemplateOk

`func (o *CreateNetworkFabricInterconnect) GetBgpNeighborTemplateOk() (*string, bool)`

GetBgpNeighborTemplateOk returns a tuple with the BgpNeighborTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpNeighborTemplate

`func (o *CreateNetworkFabricInterconnect) SetBgpNeighborTemplate(v string)`

SetBgpNeighborTemplate sets BgpNeighborTemplate field to given value.

### HasBgpNeighborTemplate

`func (o *CreateNetworkFabricInterconnect) HasBgpNeighborTemplate() bool`

HasBgpNeighborTemplate returns a boolean if a field has been set.

### GetTransportId

`func (o *CreateNetworkFabricInterconnect) GetTransportId() int32`

GetTransportId returns the TransportId field if non-nil, zero value otherwise.

### GetTransportIdOk

`func (o *CreateNetworkFabricInterconnect) GetTransportIdOk() (*int32, bool)`

GetTransportIdOk returns a tuple with the TransportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportId

`func (o *CreateNetworkFabricInterconnect) SetTransportId(v int32)`

SetTransportId sets TransportId field to given value.

### HasTransportId

`func (o *CreateNetworkFabricInterconnect) HasTransportId() bool`

HasTransportId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


