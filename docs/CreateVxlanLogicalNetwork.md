# CreateVxlanLogicalNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**Kind** | [**LogicalNetworkKind**](LogicalNetworkKind.md) |  | 
**FabricId** | **int32** |  | 
**InfrastructureId** | Pointer to **NullableInt32** |  | [optional] 
**Vlan** | [**CreateVxlanLogicalNetworkVlanProperties**](CreateVxlanLogicalNetworkVlanProperties.md) |  | 
**Vxlan** | [**CreateVxlanLogicalNetworkVxlanProperties**](CreateVxlanLogicalNetworkVxlanProperties.md) |  | 
**Ipv4** | [**CreateVxlanLogicalNetworkIpv4Properties**](CreateVxlanLogicalNetworkIpv4Properties.md) |  | 
**Ipv6** | Pointer to [**CreateVxlanLogicalNetworkIpv6Properties**](CreateVxlanLogicalNetworkIpv6Properties.md) |  | [optional] 
**RouteDomainId** | Pointer to **NullableInt32** |  | [optional] 
**Mtu** | Pointer to **NullableInt32** | Maximum Transmission Unit (MTU) in bytes | [optional] 

## Methods

### NewCreateVxlanLogicalNetwork

`func NewCreateVxlanLogicalNetwork(kind LogicalNetworkKind, fabricId int32, vlan CreateVxlanLogicalNetworkVlanProperties, vxlan CreateVxlanLogicalNetworkVxlanProperties, ipv4 CreateVxlanLogicalNetworkIpv4Properties, ) *CreateVxlanLogicalNetwork`

NewCreateVxlanLogicalNetwork instantiates a new CreateVxlanLogicalNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVxlanLogicalNetworkWithDefaults

`func NewCreateVxlanLogicalNetworkWithDefaults() *CreateVxlanLogicalNetwork`

NewCreateVxlanLogicalNetworkWithDefaults instantiates a new CreateVxlanLogicalNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CreateVxlanLogicalNetwork) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateVxlanLogicalNetwork) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateVxlanLogicalNetwork) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateVxlanLogicalNetwork) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *CreateVxlanLogicalNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateVxlanLogicalNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateVxlanLogicalNetwork) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateVxlanLogicalNetwork) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAnnotations

`func (o *CreateVxlanLogicalNetwork) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *CreateVxlanLogicalNetwork) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *CreateVxlanLogicalNetwork) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *CreateVxlanLogicalNetwork) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetKind

`func (o *CreateVxlanLogicalNetwork) GetKind() LogicalNetworkKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateVxlanLogicalNetwork) GetKindOk() (*LogicalNetworkKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateVxlanLogicalNetwork) SetKind(v LogicalNetworkKind)`

SetKind sets Kind field to given value.


### GetFabricId

`func (o *CreateVxlanLogicalNetwork) GetFabricId() int32`

GetFabricId returns the FabricId field if non-nil, zero value otherwise.

### GetFabricIdOk

`func (o *CreateVxlanLogicalNetwork) GetFabricIdOk() (*int32, bool)`

GetFabricIdOk returns a tuple with the FabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricId

`func (o *CreateVxlanLogicalNetwork) SetFabricId(v int32)`

SetFabricId sets FabricId field to given value.


### GetInfrastructureId

`func (o *CreateVxlanLogicalNetwork) GetInfrastructureId() int32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *CreateVxlanLogicalNetwork) GetInfrastructureIdOk() (*int32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *CreateVxlanLogicalNetwork) SetInfrastructureId(v int32)`

SetInfrastructureId sets InfrastructureId field to given value.

### HasInfrastructureId

`func (o *CreateVxlanLogicalNetwork) HasInfrastructureId() bool`

HasInfrastructureId returns a boolean if a field has been set.

### SetInfrastructureIdNil

`func (o *CreateVxlanLogicalNetwork) SetInfrastructureIdNil(b bool)`

 SetInfrastructureIdNil sets the value for InfrastructureId to be an explicit nil

### UnsetInfrastructureId
`func (o *CreateVxlanLogicalNetwork) UnsetInfrastructureId()`

UnsetInfrastructureId ensures that no value is present for InfrastructureId, not even an explicit nil
### GetVlan

`func (o *CreateVxlanLogicalNetwork) GetVlan() CreateVxlanLogicalNetworkVlanProperties`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *CreateVxlanLogicalNetwork) GetVlanOk() (*CreateVxlanLogicalNetworkVlanProperties, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *CreateVxlanLogicalNetwork) SetVlan(v CreateVxlanLogicalNetworkVlanProperties)`

SetVlan sets Vlan field to given value.


### GetVxlan

`func (o *CreateVxlanLogicalNetwork) GetVxlan() CreateVxlanLogicalNetworkVxlanProperties`

GetVxlan returns the Vxlan field if non-nil, zero value otherwise.

### GetVxlanOk

`func (o *CreateVxlanLogicalNetwork) GetVxlanOk() (*CreateVxlanLogicalNetworkVxlanProperties, bool)`

GetVxlanOk returns a tuple with the Vxlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVxlan

`func (o *CreateVxlanLogicalNetwork) SetVxlan(v CreateVxlanLogicalNetworkVxlanProperties)`

SetVxlan sets Vxlan field to given value.


### GetIpv4

`func (o *CreateVxlanLogicalNetwork) GetIpv4() CreateVxlanLogicalNetworkIpv4Properties`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *CreateVxlanLogicalNetwork) GetIpv4Ok() (*CreateVxlanLogicalNetworkIpv4Properties, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *CreateVxlanLogicalNetwork) SetIpv4(v CreateVxlanLogicalNetworkIpv4Properties)`

SetIpv4 sets Ipv4 field to given value.


### GetIpv6

`func (o *CreateVxlanLogicalNetwork) GetIpv6() CreateVxlanLogicalNetworkIpv6Properties`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *CreateVxlanLogicalNetwork) GetIpv6Ok() (*CreateVxlanLogicalNetworkIpv6Properties, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *CreateVxlanLogicalNetwork) SetIpv6(v CreateVxlanLogicalNetworkIpv6Properties)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *CreateVxlanLogicalNetwork) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetRouteDomainId

`func (o *CreateVxlanLogicalNetwork) GetRouteDomainId() int32`

GetRouteDomainId returns the RouteDomainId field if non-nil, zero value otherwise.

### GetRouteDomainIdOk

`func (o *CreateVxlanLogicalNetwork) GetRouteDomainIdOk() (*int32, bool)`

GetRouteDomainIdOk returns a tuple with the RouteDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteDomainId

`func (o *CreateVxlanLogicalNetwork) SetRouteDomainId(v int32)`

SetRouteDomainId sets RouteDomainId field to given value.

### HasRouteDomainId

`func (o *CreateVxlanLogicalNetwork) HasRouteDomainId() bool`

HasRouteDomainId returns a boolean if a field has been set.

### SetRouteDomainIdNil

`func (o *CreateVxlanLogicalNetwork) SetRouteDomainIdNil(b bool)`

 SetRouteDomainIdNil sets the value for RouteDomainId to be an explicit nil

### UnsetRouteDomainId
`func (o *CreateVxlanLogicalNetwork) UnsetRouteDomainId()`

UnsetRouteDomainId ensures that no value is present for RouteDomainId, not even an explicit nil
### GetMtu

`func (o *CreateVxlanLogicalNetwork) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *CreateVxlanLogicalNetwork) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *CreateVxlanLogicalNetwork) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *CreateVxlanLogicalNetwork) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *CreateVxlanLogicalNetwork) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *CreateVxlanLogicalNetwork) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


