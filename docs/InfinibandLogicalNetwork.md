# InfinibandLogicalNetwork

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
**ExtensionInstanceId** | Pointer to **NullableInt32** |  | [optional] 
**ServiceStatus** | [**GenericServiceStatus**](GenericServiceStatus.md) |  | 
**LastAppliedLogicalNetworkProfileId** | **NullableInt32** |  | 
**LastLogicalNetworkProfileAppliedAt** | **time.Time** |  | 
**Config** | [**InfinibandLogicalNetworkConfig**](InfinibandLogicalNetworkConfig.md) |  | 
**Pkey** | [**InfinibandLogicalNetworkPkeyProperties**](InfinibandLogicalNetworkPkeyProperties.md) |  | 
**Ipv4** | [**InfinibandLogicalNetworkIpv4Properties**](InfinibandLogicalNetworkIpv4Properties.md) |  | 
**Ipv6** | [**InfinibandLogicalNetworkIpv6Properties**](InfinibandLogicalNetworkIpv6Properties.md) |  | 
**RouteDomainId** | **NullableInt32** |  | 
**Mtu** | Pointer to **NullableInt32** | Maximum Transmission Unit (MTU) in bytes | [optional] 

## Methods

### NewInfinibandLogicalNetwork

`func NewInfinibandLogicalNetwork(id int32, label string, name string, annotations map[string]string, createdAt time.Time, updatedAt time.Time, revision int32, kind LogicalNetworkKind, fabricId int32, infrastructureId NullableInt32, serviceStatus GenericServiceStatus, lastAppliedLogicalNetworkProfileId NullableInt32, lastLogicalNetworkProfileAppliedAt time.Time, config InfinibandLogicalNetworkConfig, pkey InfinibandLogicalNetworkPkeyProperties, ipv4 InfinibandLogicalNetworkIpv4Properties, ipv6 InfinibandLogicalNetworkIpv6Properties, routeDomainId NullableInt32, ) *InfinibandLogicalNetwork`

NewInfinibandLogicalNetwork instantiates a new InfinibandLogicalNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfinibandLogicalNetworkWithDefaults

`func NewInfinibandLogicalNetworkWithDefaults() *InfinibandLogicalNetwork`

NewInfinibandLogicalNetworkWithDefaults instantiates a new InfinibandLogicalNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InfinibandLogicalNetwork) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InfinibandLogicalNetwork) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InfinibandLogicalNetwork) SetId(v int32)`

SetId sets Id field to given value.


### GetLabel

`func (o *InfinibandLogicalNetwork) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *InfinibandLogicalNetwork) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *InfinibandLogicalNetwork) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *InfinibandLogicalNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InfinibandLogicalNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InfinibandLogicalNetwork) SetName(v string)`

SetName sets Name field to given value.


### GetAnnotations

`func (o *InfinibandLogicalNetwork) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *InfinibandLogicalNetwork) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *InfinibandLogicalNetwork) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.


### GetCreatedAt

`func (o *InfinibandLogicalNetwork) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InfinibandLogicalNetwork) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InfinibandLogicalNetwork) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *InfinibandLogicalNetwork) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InfinibandLogicalNetwork) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InfinibandLogicalNetwork) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevision

`func (o *InfinibandLogicalNetwork) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *InfinibandLogicalNetwork) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *InfinibandLogicalNetwork) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetKind

`func (o *InfinibandLogicalNetwork) GetKind() LogicalNetworkKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *InfinibandLogicalNetwork) GetKindOk() (*LogicalNetworkKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *InfinibandLogicalNetwork) SetKind(v LogicalNetworkKind)`

SetKind sets Kind field to given value.


### GetFabricId

`func (o *InfinibandLogicalNetwork) GetFabricId() int32`

GetFabricId returns the FabricId field if non-nil, zero value otherwise.

### GetFabricIdOk

`func (o *InfinibandLogicalNetwork) GetFabricIdOk() (*int32, bool)`

GetFabricIdOk returns a tuple with the FabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricId

`func (o *InfinibandLogicalNetwork) SetFabricId(v int32)`

SetFabricId sets FabricId field to given value.


### GetInfrastructureId

`func (o *InfinibandLogicalNetwork) GetInfrastructureId() int32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *InfinibandLogicalNetwork) GetInfrastructureIdOk() (*int32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *InfinibandLogicalNetwork) SetInfrastructureId(v int32)`

SetInfrastructureId sets InfrastructureId field to given value.


### SetInfrastructureIdNil

`func (o *InfinibandLogicalNetwork) SetInfrastructureIdNil(b bool)`

 SetInfrastructureIdNil sets the value for InfrastructureId to be an explicit nil

### UnsetInfrastructureId
`func (o *InfinibandLogicalNetwork) UnsetInfrastructureId()`

UnsetInfrastructureId ensures that no value is present for InfrastructureId, not even an explicit nil
### GetExtensionInstanceId

`func (o *InfinibandLogicalNetwork) GetExtensionInstanceId() int32`

GetExtensionInstanceId returns the ExtensionInstanceId field if non-nil, zero value otherwise.

### GetExtensionInstanceIdOk

`func (o *InfinibandLogicalNetwork) GetExtensionInstanceIdOk() (*int32, bool)`

GetExtensionInstanceIdOk returns a tuple with the ExtensionInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionInstanceId

`func (o *InfinibandLogicalNetwork) SetExtensionInstanceId(v int32)`

SetExtensionInstanceId sets ExtensionInstanceId field to given value.

### HasExtensionInstanceId

`func (o *InfinibandLogicalNetwork) HasExtensionInstanceId() bool`

HasExtensionInstanceId returns a boolean if a field has been set.

### SetExtensionInstanceIdNil

`func (o *InfinibandLogicalNetwork) SetExtensionInstanceIdNil(b bool)`

 SetExtensionInstanceIdNil sets the value for ExtensionInstanceId to be an explicit nil

### UnsetExtensionInstanceId
`func (o *InfinibandLogicalNetwork) UnsetExtensionInstanceId()`

UnsetExtensionInstanceId ensures that no value is present for ExtensionInstanceId, not even an explicit nil
### GetServiceStatus

`func (o *InfinibandLogicalNetwork) GetServiceStatus() GenericServiceStatus`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *InfinibandLogicalNetwork) GetServiceStatusOk() (*GenericServiceStatus, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *InfinibandLogicalNetwork) SetServiceStatus(v GenericServiceStatus)`

SetServiceStatus sets ServiceStatus field to given value.


### GetLastAppliedLogicalNetworkProfileId

`func (o *InfinibandLogicalNetwork) GetLastAppliedLogicalNetworkProfileId() int32`

GetLastAppliedLogicalNetworkProfileId returns the LastAppliedLogicalNetworkProfileId field if non-nil, zero value otherwise.

### GetLastAppliedLogicalNetworkProfileIdOk

`func (o *InfinibandLogicalNetwork) GetLastAppliedLogicalNetworkProfileIdOk() (*int32, bool)`

GetLastAppliedLogicalNetworkProfileIdOk returns a tuple with the LastAppliedLogicalNetworkProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAppliedLogicalNetworkProfileId

`func (o *InfinibandLogicalNetwork) SetLastAppliedLogicalNetworkProfileId(v int32)`

SetLastAppliedLogicalNetworkProfileId sets LastAppliedLogicalNetworkProfileId field to given value.


### SetLastAppliedLogicalNetworkProfileIdNil

`func (o *InfinibandLogicalNetwork) SetLastAppliedLogicalNetworkProfileIdNil(b bool)`

 SetLastAppliedLogicalNetworkProfileIdNil sets the value for LastAppliedLogicalNetworkProfileId to be an explicit nil

### UnsetLastAppliedLogicalNetworkProfileId
`func (o *InfinibandLogicalNetwork) UnsetLastAppliedLogicalNetworkProfileId()`

UnsetLastAppliedLogicalNetworkProfileId ensures that no value is present for LastAppliedLogicalNetworkProfileId, not even an explicit nil
### GetLastLogicalNetworkProfileAppliedAt

`func (o *InfinibandLogicalNetwork) GetLastLogicalNetworkProfileAppliedAt() time.Time`

GetLastLogicalNetworkProfileAppliedAt returns the LastLogicalNetworkProfileAppliedAt field if non-nil, zero value otherwise.

### GetLastLogicalNetworkProfileAppliedAtOk

`func (o *InfinibandLogicalNetwork) GetLastLogicalNetworkProfileAppliedAtOk() (*time.Time, bool)`

GetLastLogicalNetworkProfileAppliedAtOk returns a tuple with the LastLogicalNetworkProfileAppliedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogicalNetworkProfileAppliedAt

`func (o *InfinibandLogicalNetwork) SetLastLogicalNetworkProfileAppliedAt(v time.Time)`

SetLastLogicalNetworkProfileAppliedAt sets LastLogicalNetworkProfileAppliedAt field to given value.


### GetConfig

`func (o *InfinibandLogicalNetwork) GetConfig() InfinibandLogicalNetworkConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *InfinibandLogicalNetwork) GetConfigOk() (*InfinibandLogicalNetworkConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *InfinibandLogicalNetwork) SetConfig(v InfinibandLogicalNetworkConfig)`

SetConfig sets Config field to given value.


### GetPkey

`func (o *InfinibandLogicalNetwork) GetPkey() InfinibandLogicalNetworkPkeyProperties`

GetPkey returns the Pkey field if non-nil, zero value otherwise.

### GetPkeyOk

`func (o *InfinibandLogicalNetwork) GetPkeyOk() (*InfinibandLogicalNetworkPkeyProperties, bool)`

GetPkeyOk returns a tuple with the Pkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkey

`func (o *InfinibandLogicalNetwork) SetPkey(v InfinibandLogicalNetworkPkeyProperties)`

SetPkey sets Pkey field to given value.


### GetIpv4

`func (o *InfinibandLogicalNetwork) GetIpv4() InfinibandLogicalNetworkIpv4Properties`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *InfinibandLogicalNetwork) GetIpv4Ok() (*InfinibandLogicalNetworkIpv4Properties, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *InfinibandLogicalNetwork) SetIpv4(v InfinibandLogicalNetworkIpv4Properties)`

SetIpv4 sets Ipv4 field to given value.


### GetIpv6

`func (o *InfinibandLogicalNetwork) GetIpv6() InfinibandLogicalNetworkIpv6Properties`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *InfinibandLogicalNetwork) GetIpv6Ok() (*InfinibandLogicalNetworkIpv6Properties, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *InfinibandLogicalNetwork) SetIpv6(v InfinibandLogicalNetworkIpv6Properties)`

SetIpv6 sets Ipv6 field to given value.


### GetRouteDomainId

`func (o *InfinibandLogicalNetwork) GetRouteDomainId() int32`

GetRouteDomainId returns the RouteDomainId field if non-nil, zero value otherwise.

### GetRouteDomainIdOk

`func (o *InfinibandLogicalNetwork) GetRouteDomainIdOk() (*int32, bool)`

GetRouteDomainIdOk returns a tuple with the RouteDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteDomainId

`func (o *InfinibandLogicalNetwork) SetRouteDomainId(v int32)`

SetRouteDomainId sets RouteDomainId field to given value.


### SetRouteDomainIdNil

`func (o *InfinibandLogicalNetwork) SetRouteDomainIdNil(b bool)`

 SetRouteDomainIdNil sets the value for RouteDomainId to be an explicit nil

### UnsetRouteDomainId
`func (o *InfinibandLogicalNetwork) UnsetRouteDomainId()`

UnsetRouteDomainId ensures that no value is present for RouteDomainId, not even an explicit nil
### GetMtu

`func (o *InfinibandLogicalNetwork) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *InfinibandLogicalNetwork) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *InfinibandLogicalNetwork) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *InfinibandLogicalNetwork) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *InfinibandLogicalNetwork) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *InfinibandLogicalNetwork) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


