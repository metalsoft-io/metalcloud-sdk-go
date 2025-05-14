# CreateLogicalNetwork201Response

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
**ServiceStatus** | [**GenericServiceStatus**](GenericServiceStatus.md) |  | 
**LastAppliedLogicalNetworkProfileId** | **NullableInt32** |  | 
**LastLogicalNetworkProfileAppliedAt** | **time.Time** |  | 
**Config** | [**VxlanLogicalNetworkConfig**](VxlanLogicalNetworkConfig.md) |  | 
**Vlan** | [**VxlanLogicalNetworkVlanProperties**](VxlanLogicalNetworkVlanProperties.md) |  | 
**Ipv4** | [**VxlanLogicalNetworkIpv4Properties**](VxlanLogicalNetworkIpv4Properties.md) |  | 
**Ipv6** | [**VxlanLogicalNetworkIpv6Properties**](VxlanLogicalNetworkIpv6Properties.md) |  | 
**RouteDomainId** | **NullableInt32** |  | 
**Vxlan** | [**VxlanLogicalNetworkVxlanProperties**](VxlanLogicalNetworkVxlanProperties.md) |  | 

## Methods

### NewCreateLogicalNetwork201Response

`func NewCreateLogicalNetwork201Response(id int32, label string, name string, annotations map[string]string, createdAt time.Time, updatedAt time.Time, revision int32, kind LogicalNetworkKind, fabricId int32, infrastructureId NullableInt32, serviceStatus GenericServiceStatus, lastAppliedLogicalNetworkProfileId NullableInt32, lastLogicalNetworkProfileAppliedAt time.Time, config VxlanLogicalNetworkConfig, vlan VxlanLogicalNetworkVlanProperties, ipv4 VxlanLogicalNetworkIpv4Properties, ipv6 VxlanLogicalNetworkIpv6Properties, routeDomainId NullableInt32, vxlan VxlanLogicalNetworkVxlanProperties, ) *CreateLogicalNetwork201Response`

NewCreateLogicalNetwork201Response instantiates a new CreateLogicalNetwork201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLogicalNetwork201ResponseWithDefaults

`func NewCreateLogicalNetwork201ResponseWithDefaults() *CreateLogicalNetwork201Response`

NewCreateLogicalNetwork201ResponseWithDefaults instantiates a new CreateLogicalNetwork201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateLogicalNetwork201Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateLogicalNetwork201Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateLogicalNetwork201Response) SetId(v int32)`

SetId sets Id field to given value.


### GetLabel

`func (o *CreateLogicalNetwork201Response) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateLogicalNetwork201Response) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateLogicalNetwork201Response) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *CreateLogicalNetwork201Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLogicalNetwork201Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLogicalNetwork201Response) SetName(v string)`

SetName sets Name field to given value.


### GetAnnotations

`func (o *CreateLogicalNetwork201Response) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *CreateLogicalNetwork201Response) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *CreateLogicalNetwork201Response) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.


### GetCreatedAt

`func (o *CreateLogicalNetwork201Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateLogicalNetwork201Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateLogicalNetwork201Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CreateLogicalNetwork201Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateLogicalNetwork201Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateLogicalNetwork201Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevision

`func (o *CreateLogicalNetwork201Response) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *CreateLogicalNetwork201Response) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *CreateLogicalNetwork201Response) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetKind

`func (o *CreateLogicalNetwork201Response) GetKind() LogicalNetworkKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateLogicalNetwork201Response) GetKindOk() (*LogicalNetworkKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateLogicalNetwork201Response) SetKind(v LogicalNetworkKind)`

SetKind sets Kind field to given value.


### GetFabricId

`func (o *CreateLogicalNetwork201Response) GetFabricId() int32`

GetFabricId returns the FabricId field if non-nil, zero value otherwise.

### GetFabricIdOk

`func (o *CreateLogicalNetwork201Response) GetFabricIdOk() (*int32, bool)`

GetFabricIdOk returns a tuple with the FabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricId

`func (o *CreateLogicalNetwork201Response) SetFabricId(v int32)`

SetFabricId sets FabricId field to given value.


### GetInfrastructureId

`func (o *CreateLogicalNetwork201Response) GetInfrastructureId() int32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *CreateLogicalNetwork201Response) GetInfrastructureIdOk() (*int32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *CreateLogicalNetwork201Response) SetInfrastructureId(v int32)`

SetInfrastructureId sets InfrastructureId field to given value.


### SetInfrastructureIdNil

`func (o *CreateLogicalNetwork201Response) SetInfrastructureIdNil(b bool)`

 SetInfrastructureIdNil sets the value for InfrastructureId to be an explicit nil

### UnsetInfrastructureId
`func (o *CreateLogicalNetwork201Response) UnsetInfrastructureId()`

UnsetInfrastructureId ensures that no value is present for InfrastructureId, not even an explicit nil
### GetServiceStatus

`func (o *CreateLogicalNetwork201Response) GetServiceStatus() GenericServiceStatus`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *CreateLogicalNetwork201Response) GetServiceStatusOk() (*GenericServiceStatus, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *CreateLogicalNetwork201Response) SetServiceStatus(v GenericServiceStatus)`

SetServiceStatus sets ServiceStatus field to given value.


### GetLastAppliedLogicalNetworkProfileId

`func (o *CreateLogicalNetwork201Response) GetLastAppliedLogicalNetworkProfileId() int32`

GetLastAppliedLogicalNetworkProfileId returns the LastAppliedLogicalNetworkProfileId field if non-nil, zero value otherwise.

### GetLastAppliedLogicalNetworkProfileIdOk

`func (o *CreateLogicalNetwork201Response) GetLastAppliedLogicalNetworkProfileIdOk() (*int32, bool)`

GetLastAppliedLogicalNetworkProfileIdOk returns a tuple with the LastAppliedLogicalNetworkProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAppliedLogicalNetworkProfileId

`func (o *CreateLogicalNetwork201Response) SetLastAppliedLogicalNetworkProfileId(v int32)`

SetLastAppliedLogicalNetworkProfileId sets LastAppliedLogicalNetworkProfileId field to given value.


### SetLastAppliedLogicalNetworkProfileIdNil

`func (o *CreateLogicalNetwork201Response) SetLastAppliedLogicalNetworkProfileIdNil(b bool)`

 SetLastAppliedLogicalNetworkProfileIdNil sets the value for LastAppliedLogicalNetworkProfileId to be an explicit nil

### UnsetLastAppliedLogicalNetworkProfileId
`func (o *CreateLogicalNetwork201Response) UnsetLastAppliedLogicalNetworkProfileId()`

UnsetLastAppliedLogicalNetworkProfileId ensures that no value is present for LastAppliedLogicalNetworkProfileId, not even an explicit nil
### GetLastLogicalNetworkProfileAppliedAt

`func (o *CreateLogicalNetwork201Response) GetLastLogicalNetworkProfileAppliedAt() time.Time`

GetLastLogicalNetworkProfileAppliedAt returns the LastLogicalNetworkProfileAppliedAt field if non-nil, zero value otherwise.

### GetLastLogicalNetworkProfileAppliedAtOk

`func (o *CreateLogicalNetwork201Response) GetLastLogicalNetworkProfileAppliedAtOk() (*time.Time, bool)`

GetLastLogicalNetworkProfileAppliedAtOk returns a tuple with the LastLogicalNetworkProfileAppliedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogicalNetworkProfileAppliedAt

`func (o *CreateLogicalNetwork201Response) SetLastLogicalNetworkProfileAppliedAt(v time.Time)`

SetLastLogicalNetworkProfileAppliedAt sets LastLogicalNetworkProfileAppliedAt field to given value.


### GetConfig

`func (o *CreateLogicalNetwork201Response) GetConfig() VxlanLogicalNetworkConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateLogicalNetwork201Response) GetConfigOk() (*VxlanLogicalNetworkConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateLogicalNetwork201Response) SetConfig(v VxlanLogicalNetworkConfig)`

SetConfig sets Config field to given value.


### GetVlan

`func (o *CreateLogicalNetwork201Response) GetVlan() VxlanLogicalNetworkVlanProperties`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *CreateLogicalNetwork201Response) GetVlanOk() (*VxlanLogicalNetworkVlanProperties, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *CreateLogicalNetwork201Response) SetVlan(v VxlanLogicalNetworkVlanProperties)`

SetVlan sets Vlan field to given value.


### GetIpv4

`func (o *CreateLogicalNetwork201Response) GetIpv4() VxlanLogicalNetworkIpv4Properties`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *CreateLogicalNetwork201Response) GetIpv4Ok() (*VxlanLogicalNetworkIpv4Properties, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *CreateLogicalNetwork201Response) SetIpv4(v VxlanLogicalNetworkIpv4Properties)`

SetIpv4 sets Ipv4 field to given value.


### GetIpv6

`func (o *CreateLogicalNetwork201Response) GetIpv6() VxlanLogicalNetworkIpv6Properties`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *CreateLogicalNetwork201Response) GetIpv6Ok() (*VxlanLogicalNetworkIpv6Properties, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *CreateLogicalNetwork201Response) SetIpv6(v VxlanLogicalNetworkIpv6Properties)`

SetIpv6 sets Ipv6 field to given value.


### GetRouteDomainId

`func (o *CreateLogicalNetwork201Response) GetRouteDomainId() int32`

GetRouteDomainId returns the RouteDomainId field if non-nil, zero value otherwise.

### GetRouteDomainIdOk

`func (o *CreateLogicalNetwork201Response) GetRouteDomainIdOk() (*int32, bool)`

GetRouteDomainIdOk returns a tuple with the RouteDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteDomainId

`func (o *CreateLogicalNetwork201Response) SetRouteDomainId(v int32)`

SetRouteDomainId sets RouteDomainId field to given value.


### SetRouteDomainIdNil

`func (o *CreateLogicalNetwork201Response) SetRouteDomainIdNil(b bool)`

 SetRouteDomainIdNil sets the value for RouteDomainId to be an explicit nil

### UnsetRouteDomainId
`func (o *CreateLogicalNetwork201Response) UnsetRouteDomainId()`

UnsetRouteDomainId ensures that no value is present for RouteDomainId, not even an explicit nil
### GetVxlan

`func (o *CreateLogicalNetwork201Response) GetVxlan() VxlanLogicalNetworkVxlanProperties`

GetVxlan returns the Vxlan field if non-nil, zero value otherwise.

### GetVxlanOk

`func (o *CreateLogicalNetwork201Response) GetVxlanOk() (*VxlanLogicalNetworkVxlanProperties, bool)`

GetVxlanOk returns a tuple with the Vxlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVxlan

`func (o *CreateLogicalNetwork201Response) SetVxlan(v VxlanLogicalNetworkVxlanProperties)`

SetVxlan sets Vxlan field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


