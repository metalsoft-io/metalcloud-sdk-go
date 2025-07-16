# CreateLogicalNetworkProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**Kind** | [**LogicalNetworkKind**](LogicalNetworkKind.md) |  | 
**FabricId** | **int32** |  | 
**Vlan** | Pointer to [**CreateLogicalNetworkVlanProperties**](CreateLogicalNetworkVlanProperties.md) |  | [optional] 
**Vxlan** | Pointer to [**CreateLogicalNetworkVxlanProperties**](CreateLogicalNetworkVxlanProperties.md) |  | [optional] 
**Ipv4** | Pointer to [**CreateLogicalNetworkIpv4Properties**](CreateLogicalNetworkIpv4Properties.md) |  | [optional] 
**Ipv6** | Pointer to [**CreateLogicalNetworkIpv6Properties**](CreateLogicalNetworkIpv6Properties.md) |  | [optional] 
**RouteDomainId** | Pointer to **NullableInt32** |  | [optional] 
**Mtu** | Pointer to **NullableInt32** | Maximum Transmission Unit (MTU) in bytes | [optional] 

## Methods

### NewCreateLogicalNetworkProfile

`func NewCreateLogicalNetworkProfile(kind LogicalNetworkKind, fabricId int32, ) *CreateLogicalNetworkProfile`

NewCreateLogicalNetworkProfile instantiates a new CreateLogicalNetworkProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLogicalNetworkProfileWithDefaults

`func NewCreateLogicalNetworkProfileWithDefaults() *CreateLogicalNetworkProfile`

NewCreateLogicalNetworkProfileWithDefaults instantiates a new CreateLogicalNetworkProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CreateLogicalNetworkProfile) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateLogicalNetworkProfile) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateLogicalNetworkProfile) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateLogicalNetworkProfile) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *CreateLogicalNetworkProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLogicalNetworkProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLogicalNetworkProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateLogicalNetworkProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAnnotations

`func (o *CreateLogicalNetworkProfile) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *CreateLogicalNetworkProfile) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *CreateLogicalNetworkProfile) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *CreateLogicalNetworkProfile) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetKind

`func (o *CreateLogicalNetworkProfile) GetKind() LogicalNetworkKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateLogicalNetworkProfile) GetKindOk() (*LogicalNetworkKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateLogicalNetworkProfile) SetKind(v LogicalNetworkKind)`

SetKind sets Kind field to given value.


### GetFabricId

`func (o *CreateLogicalNetworkProfile) GetFabricId() int32`

GetFabricId returns the FabricId field if non-nil, zero value otherwise.

### GetFabricIdOk

`func (o *CreateLogicalNetworkProfile) GetFabricIdOk() (*int32, bool)`

GetFabricIdOk returns a tuple with the FabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricId

`func (o *CreateLogicalNetworkProfile) SetFabricId(v int32)`

SetFabricId sets FabricId field to given value.


### GetVlan

`func (o *CreateLogicalNetworkProfile) GetVlan() CreateLogicalNetworkVlanProperties`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *CreateLogicalNetworkProfile) GetVlanOk() (*CreateLogicalNetworkVlanProperties, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *CreateLogicalNetworkProfile) SetVlan(v CreateLogicalNetworkVlanProperties)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *CreateLogicalNetworkProfile) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVxlan

`func (o *CreateLogicalNetworkProfile) GetVxlan() CreateLogicalNetworkVxlanProperties`

GetVxlan returns the Vxlan field if non-nil, zero value otherwise.

### GetVxlanOk

`func (o *CreateLogicalNetworkProfile) GetVxlanOk() (*CreateLogicalNetworkVxlanProperties, bool)`

GetVxlanOk returns a tuple with the Vxlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVxlan

`func (o *CreateLogicalNetworkProfile) SetVxlan(v CreateLogicalNetworkVxlanProperties)`

SetVxlan sets Vxlan field to given value.

### HasVxlan

`func (o *CreateLogicalNetworkProfile) HasVxlan() bool`

HasVxlan returns a boolean if a field has been set.

### GetIpv4

`func (o *CreateLogicalNetworkProfile) GetIpv4() CreateLogicalNetworkIpv4Properties`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *CreateLogicalNetworkProfile) GetIpv4Ok() (*CreateLogicalNetworkIpv4Properties, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *CreateLogicalNetworkProfile) SetIpv4(v CreateLogicalNetworkIpv4Properties)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *CreateLogicalNetworkProfile) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *CreateLogicalNetworkProfile) GetIpv6() CreateLogicalNetworkIpv6Properties`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *CreateLogicalNetworkProfile) GetIpv6Ok() (*CreateLogicalNetworkIpv6Properties, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *CreateLogicalNetworkProfile) SetIpv6(v CreateLogicalNetworkIpv6Properties)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *CreateLogicalNetworkProfile) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetRouteDomainId

`func (o *CreateLogicalNetworkProfile) GetRouteDomainId() int32`

GetRouteDomainId returns the RouteDomainId field if non-nil, zero value otherwise.

### GetRouteDomainIdOk

`func (o *CreateLogicalNetworkProfile) GetRouteDomainIdOk() (*int32, bool)`

GetRouteDomainIdOk returns a tuple with the RouteDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteDomainId

`func (o *CreateLogicalNetworkProfile) SetRouteDomainId(v int32)`

SetRouteDomainId sets RouteDomainId field to given value.

### HasRouteDomainId

`func (o *CreateLogicalNetworkProfile) HasRouteDomainId() bool`

HasRouteDomainId returns a boolean if a field has been set.

### SetRouteDomainIdNil

`func (o *CreateLogicalNetworkProfile) SetRouteDomainIdNil(b bool)`

 SetRouteDomainIdNil sets the value for RouteDomainId to be an explicit nil

### UnsetRouteDomainId
`func (o *CreateLogicalNetworkProfile) UnsetRouteDomainId()`

UnsetRouteDomainId ensures that no value is present for RouteDomainId, not even an explicit nil
### GetMtu

`func (o *CreateLogicalNetworkProfile) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *CreateLogicalNetworkProfile) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *CreateLogicalNetworkProfile) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *CreateLogicalNetworkProfile) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *CreateLogicalNetworkProfile) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *CreateLogicalNetworkProfile) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


