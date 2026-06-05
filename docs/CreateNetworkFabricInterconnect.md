# CreateNetworkFabricInterconnect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterconnectType** | [**NetworkFabricInterconnectType**](NetworkFabricInterconnectType.md) | The network fabric interconnect type | 
**Label** | **string** | Unique label for the network fabric interconnect | 
**Name** | Pointer to **string** | Name of the network fabric interconnect | [optional] 
**Description** | Pointer to **string** | Short description of the network fabric interconnect | [optional] 
**BgpConfigurationTemplateId** | Pointer to **int64** | ID of the BGP interconnect configuration template assigned to this interconnect | [optional] 
**TransportId** | Pointer to **int64** | Transport ID allocated for this interconnect, if any. Allocated from range 65534-1 (descending) | [optional] 

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

### GetBgpConfigurationTemplateId

`func (o *CreateNetworkFabricInterconnect) GetBgpConfigurationTemplateId() int64`

GetBgpConfigurationTemplateId returns the BgpConfigurationTemplateId field if non-nil, zero value otherwise.

### GetBgpConfigurationTemplateIdOk

`func (o *CreateNetworkFabricInterconnect) GetBgpConfigurationTemplateIdOk() (*int64, bool)`

GetBgpConfigurationTemplateIdOk returns a tuple with the BgpConfigurationTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpConfigurationTemplateId

`func (o *CreateNetworkFabricInterconnect) SetBgpConfigurationTemplateId(v int64)`

SetBgpConfigurationTemplateId sets BgpConfigurationTemplateId field to given value.

### HasBgpConfigurationTemplateId

`func (o *CreateNetworkFabricInterconnect) HasBgpConfigurationTemplateId() bool`

HasBgpConfigurationTemplateId returns a boolean if a field has been set.

### GetTransportId

`func (o *CreateNetworkFabricInterconnect) GetTransportId() int64`

GetTransportId returns the TransportId field if non-nil, zero value otherwise.

### GetTransportIdOk

`func (o *CreateNetworkFabricInterconnect) GetTransportIdOk() (*int64, bool)`

GetTransportIdOk returns a tuple with the TransportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportId

`func (o *CreateNetworkFabricInterconnect) SetTransportId(v int64)`

SetTransportId sets TransportId field to given value.

### HasTransportId

`func (o *CreateNetworkFabricInterconnect) HasTransportId() bool`

HasTransportId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


