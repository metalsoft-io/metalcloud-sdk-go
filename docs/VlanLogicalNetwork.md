# VlanLogicalNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Label** | **string** |  | 
**Name** | **string** |  | 
**Annotations** | **map[string]string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Revision** | **int32** |  | 
**Kind** | [**LogicalNetworkKind**](LogicalNetworkKind.md) |  | 
**FabricId** | **int32** |  | 
**InfrastructureId** | **NullableInt32** |  | 
**ExtensionInstanceId** | **NullableInt32** |  | 
**ServiceStatus** | [**GenericServiceStatus**](GenericServiceStatus.md) |  | 
**LastAppliedLogicalNetworkProfileId** | **NullableInt32** |  | 
**LastLogicalNetworkProfileAppliedAt** | **time.Time** |  | 
**Config** | [**VlanLogicalNetworkConfig**](VlanLogicalNetworkConfig.md) |  | 
**Vlan** | [**VlanLogicalNetworkVlanProperties**](VlanLogicalNetworkVlanProperties.md) |  | 
**Ipv4** | [**VlanLogicalNetworkIpv4Properties**](VlanLogicalNetworkIpv4Properties.md) |  | 
**Ipv6** | [**VlanLogicalNetworkIpv6Properties**](VlanLogicalNetworkIpv6Properties.md) |  | 
**RouteDomainId** | **NullableInt32** |  | 
**Mtu** | Pointer to **NullableInt32** | Maximum Transmission Unit (MTU) in bytes | [optional] 

## Methods

### NewVlanLogicalNetwork

`func NewVlanLogicalNetwork(id int32, label string, name string, annotations map[string]string, createdAt time.Time, updatedAt time.Time, revision int32, kind LogicalNetworkKind, fabricId int32, infrastructureId NullableInt32, extensionInstanceId NullableInt32, serviceStatus GenericServiceStatus, lastAppliedLogicalNetworkProfileId NullableInt32, lastLogicalNetworkProfileAppliedAt time.Time, config VlanLogicalNetworkConfig, vlan VlanLogicalNetworkVlanProperties, ipv4 VlanLogicalNetworkIpv4Properties, ipv6 VlanLogicalNetworkIpv6Properties, routeDomainId NullableInt32, ) *VlanLogicalNetwork`

NewVlanLogicalNetwork instantiates a new VlanLogicalNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanLogicalNetworkWithDefaults

`func NewVlanLogicalNetworkWithDefaults() *VlanLogicalNetwork`

NewVlanLogicalNetworkWithDefaults instantiates a new VlanLogicalNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VlanLogicalNetwork) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VlanLogicalNetwork) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VlanLogicalNetwork) SetId(v int32)`

SetId sets Id field to given value.


### GetLabel

`func (o *VlanLogicalNetwork) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *VlanLogicalNetwork) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *VlanLogicalNetwork) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *VlanLogicalNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VlanLogicalNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VlanLogicalNetwork) SetName(v string)`

SetName sets Name field to given value.


### GetAnnotations

`func (o *VlanLogicalNetwork) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *VlanLogicalNetwork) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *VlanLogicalNetwork) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.


### GetCreatedAt

`func (o *VlanLogicalNetwork) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VlanLogicalNetwork) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VlanLogicalNetwork) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *VlanLogicalNetwork) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VlanLogicalNetwork) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VlanLogicalNetwork) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevision

`func (o *VlanLogicalNetwork) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *VlanLogicalNetwork) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *VlanLogicalNetwork) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetKind

`func (o *VlanLogicalNetwork) GetKind() LogicalNetworkKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *VlanLogicalNetwork) GetKindOk() (*LogicalNetworkKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *VlanLogicalNetwork) SetKind(v LogicalNetworkKind)`

SetKind sets Kind field to given value.


### GetFabricId

`func (o *VlanLogicalNetwork) GetFabricId() int32`

GetFabricId returns the FabricId field if non-nil, zero value otherwise.

### GetFabricIdOk

`func (o *VlanLogicalNetwork) GetFabricIdOk() (*int32, bool)`

GetFabricIdOk returns a tuple with the FabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricId

`func (o *VlanLogicalNetwork) SetFabricId(v int32)`

SetFabricId sets FabricId field to given value.


### GetInfrastructureId

`func (o *VlanLogicalNetwork) GetInfrastructureId() int32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *VlanLogicalNetwork) GetInfrastructureIdOk() (*int32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *VlanLogicalNetwork) SetInfrastructureId(v int32)`

SetInfrastructureId sets InfrastructureId field to given value.


### SetInfrastructureIdNil

`func (o *VlanLogicalNetwork) SetInfrastructureIdNil(b bool)`

 SetInfrastructureIdNil sets the value for InfrastructureId to be an explicit nil

### UnsetInfrastructureId
`func (o *VlanLogicalNetwork) UnsetInfrastructureId()`

UnsetInfrastructureId ensures that no value is present for InfrastructureId, not even an explicit nil
### GetExtensionInstanceId

`func (o *VlanLogicalNetwork) GetExtensionInstanceId() int32`

GetExtensionInstanceId returns the ExtensionInstanceId field if non-nil, zero value otherwise.

### GetExtensionInstanceIdOk

`func (o *VlanLogicalNetwork) GetExtensionInstanceIdOk() (*int32, bool)`

GetExtensionInstanceIdOk returns a tuple with the ExtensionInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionInstanceId

`func (o *VlanLogicalNetwork) SetExtensionInstanceId(v int32)`

SetExtensionInstanceId sets ExtensionInstanceId field to given value.


### SetExtensionInstanceIdNil

`func (o *VlanLogicalNetwork) SetExtensionInstanceIdNil(b bool)`

 SetExtensionInstanceIdNil sets the value for ExtensionInstanceId to be an explicit nil

### UnsetExtensionInstanceId
`func (o *VlanLogicalNetwork) UnsetExtensionInstanceId()`

UnsetExtensionInstanceId ensures that no value is present for ExtensionInstanceId, not even an explicit nil
### GetServiceStatus

`func (o *VlanLogicalNetwork) GetServiceStatus() GenericServiceStatus`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *VlanLogicalNetwork) GetServiceStatusOk() (*GenericServiceStatus, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *VlanLogicalNetwork) SetServiceStatus(v GenericServiceStatus)`

SetServiceStatus sets ServiceStatus field to given value.


### GetLastAppliedLogicalNetworkProfileId

`func (o *VlanLogicalNetwork) GetLastAppliedLogicalNetworkProfileId() int32`

GetLastAppliedLogicalNetworkProfileId returns the LastAppliedLogicalNetworkProfileId field if non-nil, zero value otherwise.

### GetLastAppliedLogicalNetworkProfileIdOk

`func (o *VlanLogicalNetwork) GetLastAppliedLogicalNetworkProfileIdOk() (*int32, bool)`

GetLastAppliedLogicalNetworkProfileIdOk returns a tuple with the LastAppliedLogicalNetworkProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAppliedLogicalNetworkProfileId

`func (o *VlanLogicalNetwork) SetLastAppliedLogicalNetworkProfileId(v int32)`

SetLastAppliedLogicalNetworkProfileId sets LastAppliedLogicalNetworkProfileId field to given value.


### SetLastAppliedLogicalNetworkProfileIdNil

`func (o *VlanLogicalNetwork) SetLastAppliedLogicalNetworkProfileIdNil(b bool)`

 SetLastAppliedLogicalNetworkProfileIdNil sets the value for LastAppliedLogicalNetworkProfileId to be an explicit nil

### UnsetLastAppliedLogicalNetworkProfileId
`func (o *VlanLogicalNetwork) UnsetLastAppliedLogicalNetworkProfileId()`

UnsetLastAppliedLogicalNetworkProfileId ensures that no value is present for LastAppliedLogicalNetworkProfileId, not even an explicit nil
### GetLastLogicalNetworkProfileAppliedAt

`func (o *VlanLogicalNetwork) GetLastLogicalNetworkProfileAppliedAt() time.Time`

GetLastLogicalNetworkProfileAppliedAt returns the LastLogicalNetworkProfileAppliedAt field if non-nil, zero value otherwise.

### GetLastLogicalNetworkProfileAppliedAtOk

`func (o *VlanLogicalNetwork) GetLastLogicalNetworkProfileAppliedAtOk() (*time.Time, bool)`

GetLastLogicalNetworkProfileAppliedAtOk returns a tuple with the LastLogicalNetworkProfileAppliedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogicalNetworkProfileAppliedAt

`func (o *VlanLogicalNetwork) SetLastLogicalNetworkProfileAppliedAt(v time.Time)`

SetLastLogicalNetworkProfileAppliedAt sets LastLogicalNetworkProfileAppliedAt field to given value.


### GetConfig

`func (o *VlanLogicalNetwork) GetConfig() VlanLogicalNetworkConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *VlanLogicalNetwork) GetConfigOk() (*VlanLogicalNetworkConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *VlanLogicalNetwork) SetConfig(v VlanLogicalNetworkConfig)`

SetConfig sets Config field to given value.


### GetVlan

`func (o *VlanLogicalNetwork) GetVlan() VlanLogicalNetworkVlanProperties`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *VlanLogicalNetwork) GetVlanOk() (*VlanLogicalNetworkVlanProperties, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *VlanLogicalNetwork) SetVlan(v VlanLogicalNetworkVlanProperties)`

SetVlan sets Vlan field to given value.


### GetIpv4

`func (o *VlanLogicalNetwork) GetIpv4() VlanLogicalNetworkIpv4Properties`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *VlanLogicalNetwork) GetIpv4Ok() (*VlanLogicalNetworkIpv4Properties, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *VlanLogicalNetwork) SetIpv4(v VlanLogicalNetworkIpv4Properties)`

SetIpv4 sets Ipv4 field to given value.


### GetIpv6

`func (o *VlanLogicalNetwork) GetIpv6() VlanLogicalNetworkIpv6Properties`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *VlanLogicalNetwork) GetIpv6Ok() (*VlanLogicalNetworkIpv6Properties, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *VlanLogicalNetwork) SetIpv6(v VlanLogicalNetworkIpv6Properties)`

SetIpv6 sets Ipv6 field to given value.


### GetRouteDomainId

`func (o *VlanLogicalNetwork) GetRouteDomainId() int32`

GetRouteDomainId returns the RouteDomainId field if non-nil, zero value otherwise.

### GetRouteDomainIdOk

`func (o *VlanLogicalNetwork) GetRouteDomainIdOk() (*int32, bool)`

GetRouteDomainIdOk returns a tuple with the RouteDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteDomainId

`func (o *VlanLogicalNetwork) SetRouteDomainId(v int32)`

SetRouteDomainId sets RouteDomainId field to given value.


### SetRouteDomainIdNil

`func (o *VlanLogicalNetwork) SetRouteDomainIdNil(b bool)`

 SetRouteDomainIdNil sets the value for RouteDomainId to be an explicit nil

### UnsetRouteDomainId
`func (o *VlanLogicalNetwork) UnsetRouteDomainId()`

UnsetRouteDomainId ensures that no value is present for RouteDomainId, not even an explicit nil
### GetMtu

`func (o *VlanLogicalNetwork) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *VlanLogicalNetwork) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *VlanLogicalNetwork) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *VlanLogicalNetwork) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *VlanLogicalNetwork) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *VlanLogicalNetwork) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


