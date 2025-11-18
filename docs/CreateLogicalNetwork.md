# CreateLogicalNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**Kind** | [**LogicalNetworkKind**](LogicalNetworkKind.md) |  | 
**FabricId** | **int32** |  | 
**InfrastructureId** | Pointer to **NullableInt32** |  | [optional] 
**Vlan** | Pointer to [**CreateLogicalNetworkVlanProperties**](CreateLogicalNetworkVlanProperties.md) |  | [optional] 
**Vxlan** | Pointer to [**CreateLogicalNetworkVxlanProperties**](CreateLogicalNetworkVxlanProperties.md) |  | [optional] 
**Pkey** | Pointer to [**CreateLogicalNetworkPkeyProperties**](CreateLogicalNetworkPkeyProperties.md) |  | [optional] 
**Ipv4** | Pointer to [**CreateLogicalNetworkIpv4Properties**](CreateLogicalNetworkIpv4Properties.md) |  | [optional] 
**Ipv6** | Pointer to [**CreateLogicalNetworkIpv6Properties**](CreateLogicalNetworkIpv6Properties.md) |  | [optional] 
**RouteDomainId** | Pointer to **NullableInt32** |  | [optional] 
**Mtu** | Pointer to **NullableInt32** | Maximum Transmission Unit (MTU) in bytes | [optional] 

## Methods

### NewCreateLogicalNetwork

`func NewCreateLogicalNetwork(kind LogicalNetworkKind, fabricId int32, ) *CreateLogicalNetwork`

NewCreateLogicalNetwork instantiates a new CreateLogicalNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLogicalNetworkWithDefaults

`func NewCreateLogicalNetworkWithDefaults() *CreateLogicalNetwork`

NewCreateLogicalNetworkWithDefaults instantiates a new CreateLogicalNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CreateLogicalNetwork) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateLogicalNetwork) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateLogicalNetwork) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateLogicalNetwork) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *CreateLogicalNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLogicalNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLogicalNetwork) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateLogicalNetwork) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAnnotations

`func (o *CreateLogicalNetwork) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *CreateLogicalNetwork) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *CreateLogicalNetwork) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *CreateLogicalNetwork) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetKind

`func (o *CreateLogicalNetwork) GetKind() LogicalNetworkKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateLogicalNetwork) GetKindOk() (*LogicalNetworkKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateLogicalNetwork) SetKind(v LogicalNetworkKind)`

SetKind sets Kind field to given value.


### GetFabricId

`func (o *CreateLogicalNetwork) GetFabricId() int32`

GetFabricId returns the FabricId field if non-nil, zero value otherwise.

### GetFabricIdOk

`func (o *CreateLogicalNetwork) GetFabricIdOk() (*int32, bool)`

GetFabricIdOk returns a tuple with the FabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricId

`func (o *CreateLogicalNetwork) SetFabricId(v int32)`

SetFabricId sets FabricId field to given value.


### GetInfrastructureId

`func (o *CreateLogicalNetwork) GetInfrastructureId() int32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *CreateLogicalNetwork) GetInfrastructureIdOk() (*int32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *CreateLogicalNetwork) SetInfrastructureId(v int32)`

SetInfrastructureId sets InfrastructureId field to given value.

### HasInfrastructureId

`func (o *CreateLogicalNetwork) HasInfrastructureId() bool`

HasInfrastructureId returns a boolean if a field has been set.

### SetInfrastructureIdNil

`func (o *CreateLogicalNetwork) SetInfrastructureIdNil(b bool)`

 SetInfrastructureIdNil sets the value for InfrastructureId to be an explicit nil

### UnsetInfrastructureId
`func (o *CreateLogicalNetwork) UnsetInfrastructureId()`

UnsetInfrastructureId ensures that no value is present for InfrastructureId, not even an explicit nil
### GetVlan

`func (o *CreateLogicalNetwork) GetVlan() CreateLogicalNetworkVlanProperties`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *CreateLogicalNetwork) GetVlanOk() (*CreateLogicalNetworkVlanProperties, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *CreateLogicalNetwork) SetVlan(v CreateLogicalNetworkVlanProperties)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *CreateLogicalNetwork) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVxlan

`func (o *CreateLogicalNetwork) GetVxlan() CreateLogicalNetworkVxlanProperties`

GetVxlan returns the Vxlan field if non-nil, zero value otherwise.

### GetVxlanOk

`func (o *CreateLogicalNetwork) GetVxlanOk() (*CreateLogicalNetworkVxlanProperties, bool)`

GetVxlanOk returns a tuple with the Vxlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVxlan

`func (o *CreateLogicalNetwork) SetVxlan(v CreateLogicalNetworkVxlanProperties)`

SetVxlan sets Vxlan field to given value.

### HasVxlan

`func (o *CreateLogicalNetwork) HasVxlan() bool`

HasVxlan returns a boolean if a field has been set.

### GetPkey

`func (o *CreateLogicalNetwork) GetPkey() CreateLogicalNetworkPkeyProperties`

GetPkey returns the Pkey field if non-nil, zero value otherwise.

### GetPkeyOk

`func (o *CreateLogicalNetwork) GetPkeyOk() (*CreateLogicalNetworkPkeyProperties, bool)`

GetPkeyOk returns a tuple with the Pkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkey

`func (o *CreateLogicalNetwork) SetPkey(v CreateLogicalNetworkPkeyProperties)`

SetPkey sets Pkey field to given value.

### HasPkey

`func (o *CreateLogicalNetwork) HasPkey() bool`

HasPkey returns a boolean if a field has been set.

### GetIpv4

`func (o *CreateLogicalNetwork) GetIpv4() CreateLogicalNetworkIpv4Properties`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *CreateLogicalNetwork) GetIpv4Ok() (*CreateLogicalNetworkIpv4Properties, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *CreateLogicalNetwork) SetIpv4(v CreateLogicalNetworkIpv4Properties)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *CreateLogicalNetwork) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *CreateLogicalNetwork) GetIpv6() CreateLogicalNetworkIpv6Properties`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *CreateLogicalNetwork) GetIpv6Ok() (*CreateLogicalNetworkIpv6Properties, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *CreateLogicalNetwork) SetIpv6(v CreateLogicalNetworkIpv6Properties)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *CreateLogicalNetwork) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetRouteDomainId

`func (o *CreateLogicalNetwork) GetRouteDomainId() int32`

GetRouteDomainId returns the RouteDomainId field if non-nil, zero value otherwise.

### GetRouteDomainIdOk

`func (o *CreateLogicalNetwork) GetRouteDomainIdOk() (*int32, bool)`

GetRouteDomainIdOk returns a tuple with the RouteDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteDomainId

`func (o *CreateLogicalNetwork) SetRouteDomainId(v int32)`

SetRouteDomainId sets RouteDomainId field to given value.

### HasRouteDomainId

`func (o *CreateLogicalNetwork) HasRouteDomainId() bool`

HasRouteDomainId returns a boolean if a field has been set.

### SetRouteDomainIdNil

`func (o *CreateLogicalNetwork) SetRouteDomainIdNil(b bool)`

 SetRouteDomainIdNil sets the value for RouteDomainId to be an explicit nil

### UnsetRouteDomainId
`func (o *CreateLogicalNetwork) UnsetRouteDomainId()`

UnsetRouteDomainId ensures that no value is present for RouteDomainId, not even an explicit nil
### GetMtu

`func (o *CreateLogicalNetwork) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *CreateLogicalNetwork) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *CreateLogicalNetwork) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *CreateLogicalNetwork) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *CreateLogicalNetwork) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *CreateLogicalNetwork) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


