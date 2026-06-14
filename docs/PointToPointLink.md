# PointToPointLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Label** | **string** |  | 
**Name** | **string** |  | 
**Annotations** | **map[string]string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Revision** | **int64** |  | 
**RouteDomainId** | **NullableInt64** |  | 
**InterfaceA** | [**NullablePointToPointInterface**](PointToPointInterface.md) |  | 
**InterfaceB** | [**NullablePointToPointInterface**](PointToPointInterface.md) |  | 
**Description** | **NullableString** |  | 
**Mtu** | **NullableInt32** |  | 
**RoutingActivation** | [**PointToPointLinkRoutingActivation**](PointToPointLinkRoutingActivation.md) | When the link routing config (IP + VRF membership) is activated on the switch interface: &#x60;default&#x60; at P2P deploy, &#x60;while_transporting_logical_network&#x60; only while an l3Only logical network rides the link. | 
**ServiceStatus** | [**GenericServiceStatus**](GenericServiceStatus.md) | Lifecycle status: &#39;ordered&#39; from creation until the first successful deploy makes the link &#39;active&#39;. Deleting an &#39;ordered&#39; link removes it immediately; deleting an &#39;active&#39; link stages a DELETE deploy. | 
**Config** | [**PointToPointLinkConfig**](PointToPointLinkConfig.md) |  | 
**Ipv4** | Pointer to [**PointToPointLinkIpv4Properties**](PointToPointLinkIpv4Properties.md) |  | [optional] 
**Ipv6** | Pointer to [**PointToPointLinkIpv6Properties**](PointToPointLinkIpv6Properties.md) |  | [optional] 

## Methods

### NewPointToPointLink

`func NewPointToPointLink(id int64, label string, name string, annotations map[string]string, createdAt time.Time, updatedAt time.Time, revision int64, routeDomainId NullableInt64, interfaceA NullablePointToPointInterface, interfaceB NullablePointToPointInterface, description NullableString, mtu NullableInt32, routingActivation PointToPointLinkRoutingActivation, serviceStatus GenericServiceStatus, config PointToPointLinkConfig, ) *PointToPointLink`

NewPointToPointLink instantiates a new PointToPointLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPointToPointLinkWithDefaults

`func NewPointToPointLinkWithDefaults() *PointToPointLink`

NewPointToPointLinkWithDefaults instantiates a new PointToPointLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PointToPointLink) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PointToPointLink) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PointToPointLink) SetId(v int64)`

SetId sets Id field to given value.


### GetLabel

`func (o *PointToPointLink) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PointToPointLink) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PointToPointLink) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *PointToPointLink) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PointToPointLink) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PointToPointLink) SetName(v string)`

SetName sets Name field to given value.


### GetAnnotations

`func (o *PointToPointLink) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *PointToPointLink) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *PointToPointLink) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.


### GetCreatedAt

`func (o *PointToPointLink) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PointToPointLink) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PointToPointLink) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PointToPointLink) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PointToPointLink) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PointToPointLink) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevision

`func (o *PointToPointLink) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *PointToPointLink) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *PointToPointLink) SetRevision(v int64)`

SetRevision sets Revision field to given value.


### GetRouteDomainId

`func (o *PointToPointLink) GetRouteDomainId() int64`

GetRouteDomainId returns the RouteDomainId field if non-nil, zero value otherwise.

### GetRouteDomainIdOk

`func (o *PointToPointLink) GetRouteDomainIdOk() (*int64, bool)`

GetRouteDomainIdOk returns a tuple with the RouteDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteDomainId

`func (o *PointToPointLink) SetRouteDomainId(v int64)`

SetRouteDomainId sets RouteDomainId field to given value.


### SetRouteDomainIdNil

`func (o *PointToPointLink) SetRouteDomainIdNil(b bool)`

 SetRouteDomainIdNil sets the value for RouteDomainId to be an explicit nil

### UnsetRouteDomainId
`func (o *PointToPointLink) UnsetRouteDomainId()`

UnsetRouteDomainId ensures that no value is present for RouteDomainId, not even an explicit nil
### GetInterfaceA

`func (o *PointToPointLink) GetInterfaceA() PointToPointInterface`

GetInterfaceA returns the InterfaceA field if non-nil, zero value otherwise.

### GetInterfaceAOk

`func (o *PointToPointLink) GetInterfaceAOk() (*PointToPointInterface, bool)`

GetInterfaceAOk returns a tuple with the InterfaceA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceA

`func (o *PointToPointLink) SetInterfaceA(v PointToPointInterface)`

SetInterfaceA sets InterfaceA field to given value.


### SetInterfaceANil

`func (o *PointToPointLink) SetInterfaceANil(b bool)`

 SetInterfaceANil sets the value for InterfaceA to be an explicit nil

### UnsetInterfaceA
`func (o *PointToPointLink) UnsetInterfaceA()`

UnsetInterfaceA ensures that no value is present for InterfaceA, not even an explicit nil
### GetInterfaceB

`func (o *PointToPointLink) GetInterfaceB() PointToPointInterface`

GetInterfaceB returns the InterfaceB field if non-nil, zero value otherwise.

### GetInterfaceBOk

`func (o *PointToPointLink) GetInterfaceBOk() (*PointToPointInterface, bool)`

GetInterfaceBOk returns a tuple with the InterfaceB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceB

`func (o *PointToPointLink) SetInterfaceB(v PointToPointInterface)`

SetInterfaceB sets InterfaceB field to given value.


### SetInterfaceBNil

`func (o *PointToPointLink) SetInterfaceBNil(b bool)`

 SetInterfaceBNil sets the value for InterfaceB to be an explicit nil

### UnsetInterfaceB
`func (o *PointToPointLink) UnsetInterfaceB()`

UnsetInterfaceB ensures that no value is present for InterfaceB, not even an explicit nil
### GetDescription

`func (o *PointToPointLink) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PointToPointLink) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PointToPointLink) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *PointToPointLink) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PointToPointLink) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMtu

`func (o *PointToPointLink) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *PointToPointLink) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *PointToPointLink) SetMtu(v int32)`

SetMtu sets Mtu field to given value.


### SetMtuNil

`func (o *PointToPointLink) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *PointToPointLink) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
### GetRoutingActivation

`func (o *PointToPointLink) GetRoutingActivation() PointToPointLinkRoutingActivation`

GetRoutingActivation returns the RoutingActivation field if non-nil, zero value otherwise.

### GetRoutingActivationOk

`func (o *PointToPointLink) GetRoutingActivationOk() (*PointToPointLinkRoutingActivation, bool)`

GetRoutingActivationOk returns a tuple with the RoutingActivation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingActivation

`func (o *PointToPointLink) SetRoutingActivation(v PointToPointLinkRoutingActivation)`

SetRoutingActivation sets RoutingActivation field to given value.


### GetServiceStatus

`func (o *PointToPointLink) GetServiceStatus() GenericServiceStatus`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *PointToPointLink) GetServiceStatusOk() (*GenericServiceStatus, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *PointToPointLink) SetServiceStatus(v GenericServiceStatus)`

SetServiceStatus sets ServiceStatus field to given value.


### GetConfig

`func (o *PointToPointLink) GetConfig() PointToPointLinkConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *PointToPointLink) GetConfigOk() (*PointToPointLinkConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *PointToPointLink) SetConfig(v PointToPointLinkConfig)`

SetConfig sets Config field to given value.


### GetIpv4

`func (o *PointToPointLink) GetIpv4() PointToPointLinkIpv4Properties`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *PointToPointLink) GetIpv4Ok() (*PointToPointLinkIpv4Properties, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *PointToPointLink) SetIpv4(v PointToPointLinkIpv4Properties)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *PointToPointLink) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *PointToPointLink) GetIpv6() PointToPointLinkIpv6Properties`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *PointToPointLink) GetIpv6Ok() (*PointToPointLinkIpv6Properties, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *PointToPointLink) SetIpv6(v PointToPointLinkIpv6Properties)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *PointToPointLink) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


