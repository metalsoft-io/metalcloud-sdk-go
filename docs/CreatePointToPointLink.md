# CreatePointToPointLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**RouteDomainId** | Pointer to **NullableInt64** | RouteDomain (VRF) the link belongs to. NULL &#x3D; the default RD. | [optional] 
**InterfaceA** | Pointer to [**NullableCreatePointToPointInterface**](CreatePointToPointInterface.md) |  | [optional] 
**InterfaceB** | Pointer to [**NullableCreatePointToPointInterface**](CreatePointToPointInterface.md) |  | [optional] 
**Description** | Pointer to **string** | Optional human-readable description. | [optional] 
**Mtu** | Pointer to **NullableInt32** | Maximum Transmission Unit (MTU) in bytes. | [optional] 
**RoutingActivation** | Pointer to [**PointToPointLinkRoutingActivation**](PointToPointLinkRoutingActivation.md) | When the link routing config (interface IP + route-domain VRF membership) is activated on the switch interface. The interface is put in routed (L3) mode and the /31(/127) is allocated at P2P deploy in both modes. &#x60;default&#x60; also installs the IP at P2P deploy; &#x60;while_transporting_logical_network&#x60; installs it only while an l3Only logical network is transported over the link — arriving with the network VRF membership and removed again (interface left in L3 mode) when the last such network detaches. Create-only — cannot be changed after creation. | [optional] [default to POINTTOPOINTLINKROUTINGACTIVATION_DEFAULT]
**Ipv4** | Pointer to [**CreatePointToPointLinkIpv4Properties**](CreatePointToPointLinkIpv4Properties.md) | IPv4 staged properties: subnet allocation strategies to stage on the candidate at creation (equivalent to adding them via the config endpoints afterwards). | [optional] 
**Ipv6** | Pointer to [**CreatePointToPointLinkIpv6Properties**](CreatePointToPointLinkIpv6Properties.md) | IPv6 staged properties: subnet allocation strategies to stage on the candidate at creation (equivalent to adding them via the config endpoints afterwards). | [optional] 

## Methods

### NewCreatePointToPointLink

`func NewCreatePointToPointLink() *CreatePointToPointLink`

NewCreatePointToPointLink instantiates a new CreatePointToPointLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePointToPointLinkWithDefaults

`func NewCreatePointToPointLinkWithDefaults() *CreatePointToPointLink`

NewCreatePointToPointLinkWithDefaults instantiates a new CreatePointToPointLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CreatePointToPointLink) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreatePointToPointLink) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreatePointToPointLink) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreatePointToPointLink) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *CreatePointToPointLink) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePointToPointLink) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePointToPointLink) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreatePointToPointLink) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAnnotations

`func (o *CreatePointToPointLink) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *CreatePointToPointLink) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *CreatePointToPointLink) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *CreatePointToPointLink) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetRouteDomainId

`func (o *CreatePointToPointLink) GetRouteDomainId() int64`

GetRouteDomainId returns the RouteDomainId field if non-nil, zero value otherwise.

### GetRouteDomainIdOk

`func (o *CreatePointToPointLink) GetRouteDomainIdOk() (*int64, bool)`

GetRouteDomainIdOk returns a tuple with the RouteDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteDomainId

`func (o *CreatePointToPointLink) SetRouteDomainId(v int64)`

SetRouteDomainId sets RouteDomainId field to given value.

### HasRouteDomainId

`func (o *CreatePointToPointLink) HasRouteDomainId() bool`

HasRouteDomainId returns a boolean if a field has been set.

### SetRouteDomainIdNil

`func (o *CreatePointToPointLink) SetRouteDomainIdNil(b bool)`

 SetRouteDomainIdNil sets the value for RouteDomainId to be an explicit nil

### UnsetRouteDomainId
`func (o *CreatePointToPointLink) UnsetRouteDomainId()`

UnsetRouteDomainId ensures that no value is present for RouteDomainId, not even an explicit nil
### GetInterfaceA

`func (o *CreatePointToPointLink) GetInterfaceA() CreatePointToPointInterface`

GetInterfaceA returns the InterfaceA field if non-nil, zero value otherwise.

### GetInterfaceAOk

`func (o *CreatePointToPointLink) GetInterfaceAOk() (*CreatePointToPointInterface, bool)`

GetInterfaceAOk returns a tuple with the InterfaceA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceA

`func (o *CreatePointToPointLink) SetInterfaceA(v CreatePointToPointInterface)`

SetInterfaceA sets InterfaceA field to given value.

### HasInterfaceA

`func (o *CreatePointToPointLink) HasInterfaceA() bool`

HasInterfaceA returns a boolean if a field has been set.

### SetInterfaceANil

`func (o *CreatePointToPointLink) SetInterfaceANil(b bool)`

 SetInterfaceANil sets the value for InterfaceA to be an explicit nil

### UnsetInterfaceA
`func (o *CreatePointToPointLink) UnsetInterfaceA()`

UnsetInterfaceA ensures that no value is present for InterfaceA, not even an explicit nil
### GetInterfaceB

`func (o *CreatePointToPointLink) GetInterfaceB() CreatePointToPointInterface`

GetInterfaceB returns the InterfaceB field if non-nil, zero value otherwise.

### GetInterfaceBOk

`func (o *CreatePointToPointLink) GetInterfaceBOk() (*CreatePointToPointInterface, bool)`

GetInterfaceBOk returns a tuple with the InterfaceB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceB

`func (o *CreatePointToPointLink) SetInterfaceB(v CreatePointToPointInterface)`

SetInterfaceB sets InterfaceB field to given value.

### HasInterfaceB

`func (o *CreatePointToPointLink) HasInterfaceB() bool`

HasInterfaceB returns a boolean if a field has been set.

### SetInterfaceBNil

`func (o *CreatePointToPointLink) SetInterfaceBNil(b bool)`

 SetInterfaceBNil sets the value for InterfaceB to be an explicit nil

### UnsetInterfaceB
`func (o *CreatePointToPointLink) UnsetInterfaceB()`

UnsetInterfaceB ensures that no value is present for InterfaceB, not even an explicit nil
### GetDescription

`func (o *CreatePointToPointLink) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatePointToPointLink) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatePointToPointLink) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreatePointToPointLink) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMtu

`func (o *CreatePointToPointLink) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *CreatePointToPointLink) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *CreatePointToPointLink) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *CreatePointToPointLink) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *CreatePointToPointLink) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *CreatePointToPointLink) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
### GetRoutingActivation

`func (o *CreatePointToPointLink) GetRoutingActivation() PointToPointLinkRoutingActivation`

GetRoutingActivation returns the RoutingActivation field if non-nil, zero value otherwise.

### GetRoutingActivationOk

`func (o *CreatePointToPointLink) GetRoutingActivationOk() (*PointToPointLinkRoutingActivation, bool)`

GetRoutingActivationOk returns a tuple with the RoutingActivation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingActivation

`func (o *CreatePointToPointLink) SetRoutingActivation(v PointToPointLinkRoutingActivation)`

SetRoutingActivation sets RoutingActivation field to given value.

### HasRoutingActivation

`func (o *CreatePointToPointLink) HasRoutingActivation() bool`

HasRoutingActivation returns a boolean if a field has been set.

### GetIpv4

`func (o *CreatePointToPointLink) GetIpv4() CreatePointToPointLinkIpv4Properties`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *CreatePointToPointLink) GetIpv4Ok() (*CreatePointToPointLinkIpv4Properties, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *CreatePointToPointLink) SetIpv4(v CreatePointToPointLinkIpv4Properties)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *CreatePointToPointLink) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *CreatePointToPointLink) GetIpv6() CreatePointToPointLinkIpv6Properties`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *CreatePointToPointLink) GetIpv6Ok() (*CreatePointToPointLinkIpv6Properties, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *CreatePointToPointLink) SetIpv6(v CreatePointToPointLinkIpv6Properties)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *CreatePointToPointLink) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


