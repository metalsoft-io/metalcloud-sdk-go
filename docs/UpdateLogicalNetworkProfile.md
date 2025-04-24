# UpdateLogicalNetworkProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**Vlan** | [**UpdateLogicalNetworkProfileVlanProperties**](UpdateLogicalNetworkProfileVlanProperties.md) |  | 
**Vxlan** | [**UpdateLogicalNetworkProfileVxlanProperties**](UpdateLogicalNetworkProfileVxlanProperties.md) |  | 
**Ipv4** | [**UpdateLogicalNetworkProfileIpv4Properties**](UpdateLogicalNetworkProfileIpv4Properties.md) |  | 
**Ipv6** | [**UpdateLogicalNetworkProfileIpv6Properties**](UpdateLogicalNetworkProfileIpv6Properties.md) |  | 
**RouteDomainId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewUpdateLogicalNetworkProfile

`func NewUpdateLogicalNetworkProfile(vlan UpdateLogicalNetworkProfileVlanProperties, vxlan UpdateLogicalNetworkProfileVxlanProperties, ipv4 UpdateLogicalNetworkProfileIpv4Properties, ipv6 UpdateLogicalNetworkProfileIpv6Properties, ) *UpdateLogicalNetworkProfile`

NewUpdateLogicalNetworkProfile instantiates a new UpdateLogicalNetworkProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLogicalNetworkProfileWithDefaults

`func NewUpdateLogicalNetworkProfileWithDefaults() *UpdateLogicalNetworkProfile`

NewUpdateLogicalNetworkProfileWithDefaults instantiates a new UpdateLogicalNetworkProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *UpdateLogicalNetworkProfile) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateLogicalNetworkProfile) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateLogicalNetworkProfile) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpdateLogicalNetworkProfile) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *UpdateLogicalNetworkProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateLogicalNetworkProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateLogicalNetworkProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateLogicalNetworkProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAnnotations

`func (o *UpdateLogicalNetworkProfile) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *UpdateLogicalNetworkProfile) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *UpdateLogicalNetworkProfile) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *UpdateLogicalNetworkProfile) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetVlan

`func (o *UpdateLogicalNetworkProfile) GetVlan() UpdateLogicalNetworkProfileVlanProperties`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *UpdateLogicalNetworkProfile) GetVlanOk() (*UpdateLogicalNetworkProfileVlanProperties, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *UpdateLogicalNetworkProfile) SetVlan(v UpdateLogicalNetworkProfileVlanProperties)`

SetVlan sets Vlan field to given value.


### GetVxlan

`func (o *UpdateLogicalNetworkProfile) GetVxlan() UpdateLogicalNetworkProfileVxlanProperties`

GetVxlan returns the Vxlan field if non-nil, zero value otherwise.

### GetVxlanOk

`func (o *UpdateLogicalNetworkProfile) GetVxlanOk() (*UpdateLogicalNetworkProfileVxlanProperties, bool)`

GetVxlanOk returns a tuple with the Vxlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVxlan

`func (o *UpdateLogicalNetworkProfile) SetVxlan(v UpdateLogicalNetworkProfileVxlanProperties)`

SetVxlan sets Vxlan field to given value.


### GetIpv4

`func (o *UpdateLogicalNetworkProfile) GetIpv4() UpdateLogicalNetworkProfileIpv4Properties`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *UpdateLogicalNetworkProfile) GetIpv4Ok() (*UpdateLogicalNetworkProfileIpv4Properties, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *UpdateLogicalNetworkProfile) SetIpv4(v UpdateLogicalNetworkProfileIpv4Properties)`

SetIpv4 sets Ipv4 field to given value.


### GetIpv6

`func (o *UpdateLogicalNetworkProfile) GetIpv6() UpdateLogicalNetworkProfileIpv6Properties`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *UpdateLogicalNetworkProfile) GetIpv6Ok() (*UpdateLogicalNetworkProfileIpv6Properties, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *UpdateLogicalNetworkProfile) SetIpv6(v UpdateLogicalNetworkProfileIpv6Properties)`

SetIpv6 sets Ipv6 field to given value.


### GetRouteDomainId

`func (o *UpdateLogicalNetworkProfile) GetRouteDomainId() int32`

GetRouteDomainId returns the RouteDomainId field if non-nil, zero value otherwise.

### GetRouteDomainIdOk

`func (o *UpdateLogicalNetworkProfile) GetRouteDomainIdOk() (*int32, bool)`

GetRouteDomainIdOk returns a tuple with the RouteDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteDomainId

`func (o *UpdateLogicalNetworkProfile) SetRouteDomainId(v int32)`

SetRouteDomainId sets RouteDomainId field to given value.

### HasRouteDomainId

`func (o *UpdateLogicalNetworkProfile) HasRouteDomainId() bool`

HasRouteDomainId returns a boolean if a field has been set.

### SetRouteDomainIdNil

`func (o *UpdateLogicalNetworkProfile) SetRouteDomainIdNil(b bool)`

 SetRouteDomainIdNil sets the value for RouteDomainId to be an explicit nil

### UnsetRouteDomainId
`func (o *UpdateLogicalNetworkProfile) UnsetRouteDomainId()`

UnsetRouteDomainId ensures that no value is present for RouteDomainId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


