# CreateInfinibandLogicalNetworkProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**Kind** | [**LogicalNetworkKind**](LogicalNetworkKind.md) |  | 
**FabricId** | **int32** |  | 
**Pkey** | [**CreateInfinibandLogicalNetworkPkeyProperties**](CreateInfinibandLogicalNetworkPkeyProperties.md) |  | 
**Ipv4** | [**CreateInfinibandLogicalNetworkIpv4Properties**](CreateInfinibandLogicalNetworkIpv4Properties.md) |  | 
**Ipv6** | Pointer to [**CreateInfinibandLogicalNetworkIpv6Properties**](CreateInfinibandLogicalNetworkIpv6Properties.md) |  | [optional] 
**RouteDomainId** | Pointer to **NullableInt32** |  | [optional] 
**Mtu** | Pointer to **NullableInt32** | Maximum Transmission Unit (MTU) in bytes | [optional] 

## Methods

### NewCreateInfinibandLogicalNetworkProfile

`func NewCreateInfinibandLogicalNetworkProfile(kind LogicalNetworkKind, fabricId int32, pkey CreateInfinibandLogicalNetworkPkeyProperties, ipv4 CreateInfinibandLogicalNetworkIpv4Properties, ) *CreateInfinibandLogicalNetworkProfile`

NewCreateInfinibandLogicalNetworkProfile instantiates a new CreateInfinibandLogicalNetworkProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInfinibandLogicalNetworkProfileWithDefaults

`func NewCreateInfinibandLogicalNetworkProfileWithDefaults() *CreateInfinibandLogicalNetworkProfile`

NewCreateInfinibandLogicalNetworkProfileWithDefaults instantiates a new CreateInfinibandLogicalNetworkProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CreateInfinibandLogicalNetworkProfile) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateInfinibandLogicalNetworkProfile) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateInfinibandLogicalNetworkProfile) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateInfinibandLogicalNetworkProfile) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *CreateInfinibandLogicalNetworkProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateInfinibandLogicalNetworkProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateInfinibandLogicalNetworkProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateInfinibandLogicalNetworkProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAnnotations

`func (o *CreateInfinibandLogicalNetworkProfile) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *CreateInfinibandLogicalNetworkProfile) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *CreateInfinibandLogicalNetworkProfile) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *CreateInfinibandLogicalNetworkProfile) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetKind

`func (o *CreateInfinibandLogicalNetworkProfile) GetKind() LogicalNetworkKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateInfinibandLogicalNetworkProfile) GetKindOk() (*LogicalNetworkKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateInfinibandLogicalNetworkProfile) SetKind(v LogicalNetworkKind)`

SetKind sets Kind field to given value.


### GetFabricId

`func (o *CreateInfinibandLogicalNetworkProfile) GetFabricId() int32`

GetFabricId returns the FabricId field if non-nil, zero value otherwise.

### GetFabricIdOk

`func (o *CreateInfinibandLogicalNetworkProfile) GetFabricIdOk() (*int32, bool)`

GetFabricIdOk returns a tuple with the FabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricId

`func (o *CreateInfinibandLogicalNetworkProfile) SetFabricId(v int32)`

SetFabricId sets FabricId field to given value.


### GetPkey

`func (o *CreateInfinibandLogicalNetworkProfile) GetPkey() CreateInfinibandLogicalNetworkPkeyProperties`

GetPkey returns the Pkey field if non-nil, zero value otherwise.

### GetPkeyOk

`func (o *CreateInfinibandLogicalNetworkProfile) GetPkeyOk() (*CreateInfinibandLogicalNetworkPkeyProperties, bool)`

GetPkeyOk returns a tuple with the Pkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkey

`func (o *CreateInfinibandLogicalNetworkProfile) SetPkey(v CreateInfinibandLogicalNetworkPkeyProperties)`

SetPkey sets Pkey field to given value.


### GetIpv4

`func (o *CreateInfinibandLogicalNetworkProfile) GetIpv4() CreateInfinibandLogicalNetworkIpv4Properties`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *CreateInfinibandLogicalNetworkProfile) GetIpv4Ok() (*CreateInfinibandLogicalNetworkIpv4Properties, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *CreateInfinibandLogicalNetworkProfile) SetIpv4(v CreateInfinibandLogicalNetworkIpv4Properties)`

SetIpv4 sets Ipv4 field to given value.


### GetIpv6

`func (o *CreateInfinibandLogicalNetworkProfile) GetIpv6() CreateInfinibandLogicalNetworkIpv6Properties`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *CreateInfinibandLogicalNetworkProfile) GetIpv6Ok() (*CreateInfinibandLogicalNetworkIpv6Properties, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *CreateInfinibandLogicalNetworkProfile) SetIpv6(v CreateInfinibandLogicalNetworkIpv6Properties)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *CreateInfinibandLogicalNetworkProfile) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetRouteDomainId

`func (o *CreateInfinibandLogicalNetworkProfile) GetRouteDomainId() int32`

GetRouteDomainId returns the RouteDomainId field if non-nil, zero value otherwise.

### GetRouteDomainIdOk

`func (o *CreateInfinibandLogicalNetworkProfile) GetRouteDomainIdOk() (*int32, bool)`

GetRouteDomainIdOk returns a tuple with the RouteDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteDomainId

`func (o *CreateInfinibandLogicalNetworkProfile) SetRouteDomainId(v int32)`

SetRouteDomainId sets RouteDomainId field to given value.

### HasRouteDomainId

`func (o *CreateInfinibandLogicalNetworkProfile) HasRouteDomainId() bool`

HasRouteDomainId returns a boolean if a field has been set.

### SetRouteDomainIdNil

`func (o *CreateInfinibandLogicalNetworkProfile) SetRouteDomainIdNil(b bool)`

 SetRouteDomainIdNil sets the value for RouteDomainId to be an explicit nil

### UnsetRouteDomainId
`func (o *CreateInfinibandLogicalNetworkProfile) UnsetRouteDomainId()`

UnsetRouteDomainId ensures that no value is present for RouteDomainId, not even an explicit nil
### GetMtu

`func (o *CreateInfinibandLogicalNetworkProfile) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *CreateInfinibandLogicalNetworkProfile) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *CreateInfinibandLogicalNetworkProfile) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *CreateInfinibandLogicalNetworkProfile) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *CreateInfinibandLogicalNetworkProfile) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *CreateInfinibandLogicalNetworkProfile) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


