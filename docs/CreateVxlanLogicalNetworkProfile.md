# CreateVxlanLogicalNetworkProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**Kind** | [**LogicalNetworkKind**](LogicalNetworkKind.md) |  | 
**FabricId** | **int32** |  | 
**Vlan** | [**CreateVxlanLogicalNetworkVlanProperties**](CreateVxlanLogicalNetworkVlanProperties.md) |  | 
**Vxlan** | [**CreateVxlanLogicalNetworkVxlanProperties**](CreateVxlanLogicalNetworkVxlanProperties.md) |  | 
**Ipv4** | [**CreateVxlanLogicalNetworkIpv4Properties**](CreateVxlanLogicalNetworkIpv4Properties.md) |  | 
**Ipv6** | Pointer to [**CreateVxlanLogicalNetworkIpv6Properties**](CreateVxlanLogicalNetworkIpv6Properties.md) |  | [optional] 
**RouteDomainId** | Pointer to **NullableInt32** |  | [optional] 
**Mtu** | Pointer to **NullableInt32** | Maximum Transmission Unit (MTU) in bytes | [optional] 

## Methods

### NewCreateVxlanLogicalNetworkProfile

`func NewCreateVxlanLogicalNetworkProfile(kind LogicalNetworkKind, fabricId int32, vlan CreateVxlanLogicalNetworkVlanProperties, vxlan CreateVxlanLogicalNetworkVxlanProperties, ipv4 CreateVxlanLogicalNetworkIpv4Properties, ) *CreateVxlanLogicalNetworkProfile`

NewCreateVxlanLogicalNetworkProfile instantiates a new CreateVxlanLogicalNetworkProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVxlanLogicalNetworkProfileWithDefaults

`func NewCreateVxlanLogicalNetworkProfileWithDefaults() *CreateVxlanLogicalNetworkProfile`

NewCreateVxlanLogicalNetworkProfileWithDefaults instantiates a new CreateVxlanLogicalNetworkProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CreateVxlanLogicalNetworkProfile) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateVxlanLogicalNetworkProfile) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateVxlanLogicalNetworkProfile) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateVxlanLogicalNetworkProfile) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *CreateVxlanLogicalNetworkProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateVxlanLogicalNetworkProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateVxlanLogicalNetworkProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateVxlanLogicalNetworkProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAnnotations

`func (o *CreateVxlanLogicalNetworkProfile) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *CreateVxlanLogicalNetworkProfile) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *CreateVxlanLogicalNetworkProfile) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *CreateVxlanLogicalNetworkProfile) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetKind

`func (o *CreateVxlanLogicalNetworkProfile) GetKind() LogicalNetworkKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateVxlanLogicalNetworkProfile) GetKindOk() (*LogicalNetworkKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateVxlanLogicalNetworkProfile) SetKind(v LogicalNetworkKind)`

SetKind sets Kind field to given value.


### GetFabricId

`func (o *CreateVxlanLogicalNetworkProfile) GetFabricId() int32`

GetFabricId returns the FabricId field if non-nil, zero value otherwise.

### GetFabricIdOk

`func (o *CreateVxlanLogicalNetworkProfile) GetFabricIdOk() (*int32, bool)`

GetFabricIdOk returns a tuple with the FabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricId

`func (o *CreateVxlanLogicalNetworkProfile) SetFabricId(v int32)`

SetFabricId sets FabricId field to given value.


### GetVlan

`func (o *CreateVxlanLogicalNetworkProfile) GetVlan() CreateVxlanLogicalNetworkVlanProperties`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *CreateVxlanLogicalNetworkProfile) GetVlanOk() (*CreateVxlanLogicalNetworkVlanProperties, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *CreateVxlanLogicalNetworkProfile) SetVlan(v CreateVxlanLogicalNetworkVlanProperties)`

SetVlan sets Vlan field to given value.


### GetVxlan

`func (o *CreateVxlanLogicalNetworkProfile) GetVxlan() CreateVxlanLogicalNetworkVxlanProperties`

GetVxlan returns the Vxlan field if non-nil, zero value otherwise.

### GetVxlanOk

`func (o *CreateVxlanLogicalNetworkProfile) GetVxlanOk() (*CreateVxlanLogicalNetworkVxlanProperties, bool)`

GetVxlanOk returns a tuple with the Vxlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVxlan

`func (o *CreateVxlanLogicalNetworkProfile) SetVxlan(v CreateVxlanLogicalNetworkVxlanProperties)`

SetVxlan sets Vxlan field to given value.


### GetIpv4

`func (o *CreateVxlanLogicalNetworkProfile) GetIpv4() CreateVxlanLogicalNetworkIpv4Properties`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *CreateVxlanLogicalNetworkProfile) GetIpv4Ok() (*CreateVxlanLogicalNetworkIpv4Properties, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *CreateVxlanLogicalNetworkProfile) SetIpv4(v CreateVxlanLogicalNetworkIpv4Properties)`

SetIpv4 sets Ipv4 field to given value.


### GetIpv6

`func (o *CreateVxlanLogicalNetworkProfile) GetIpv6() CreateVxlanLogicalNetworkIpv6Properties`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *CreateVxlanLogicalNetworkProfile) GetIpv6Ok() (*CreateVxlanLogicalNetworkIpv6Properties, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *CreateVxlanLogicalNetworkProfile) SetIpv6(v CreateVxlanLogicalNetworkIpv6Properties)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *CreateVxlanLogicalNetworkProfile) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetRouteDomainId

`func (o *CreateVxlanLogicalNetworkProfile) GetRouteDomainId() int32`

GetRouteDomainId returns the RouteDomainId field if non-nil, zero value otherwise.

### GetRouteDomainIdOk

`func (o *CreateVxlanLogicalNetworkProfile) GetRouteDomainIdOk() (*int32, bool)`

GetRouteDomainIdOk returns a tuple with the RouteDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteDomainId

`func (o *CreateVxlanLogicalNetworkProfile) SetRouteDomainId(v int32)`

SetRouteDomainId sets RouteDomainId field to given value.

### HasRouteDomainId

`func (o *CreateVxlanLogicalNetworkProfile) HasRouteDomainId() bool`

HasRouteDomainId returns a boolean if a field has been set.

### SetRouteDomainIdNil

`func (o *CreateVxlanLogicalNetworkProfile) SetRouteDomainIdNil(b bool)`

 SetRouteDomainIdNil sets the value for RouteDomainId to be an explicit nil

### UnsetRouteDomainId
`func (o *CreateVxlanLogicalNetworkProfile) UnsetRouteDomainId()`

UnsetRouteDomainId ensures that no value is present for RouteDomainId, not even an explicit nil
### GetMtu

`func (o *CreateVxlanLogicalNetworkProfile) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *CreateVxlanLogicalNetworkProfile) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *CreateVxlanLogicalNetworkProfile) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *CreateVxlanLogicalNetworkProfile) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *CreateVxlanLogicalNetworkProfile) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *CreateVxlanLogicalNetworkProfile) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


