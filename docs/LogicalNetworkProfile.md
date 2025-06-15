# LogicalNetworkProfile

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
**Vlan** | Pointer to [**LogicalNetworkProfileVlanProperties**](LogicalNetworkProfileVlanProperties.md) |  | [optional] 
**Vxlan** | Pointer to [**LogicalNetworkProfileVxlanProperties**](LogicalNetworkProfileVxlanProperties.md) |  | [optional] 
**Ipv4** | Pointer to [**LogicalNetworkProfileIpv4Properties**](LogicalNetworkProfileIpv4Properties.md) |  | [optional] 
**Ipv6** | Pointer to [**LogicalNetworkProfileIpv6Properties**](LogicalNetworkProfileIpv6Properties.md) |  | [optional] 
**RouteDomainId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewLogicalNetworkProfile

`func NewLogicalNetworkProfile(id int32, label string, name string, annotations map[string]string, createdAt time.Time, updatedAt time.Time, revision int32, kind LogicalNetworkKind, fabricId int32, ) *LogicalNetworkProfile`

NewLogicalNetworkProfile instantiates a new LogicalNetworkProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogicalNetworkProfileWithDefaults

`func NewLogicalNetworkProfileWithDefaults() *LogicalNetworkProfile`

NewLogicalNetworkProfileWithDefaults instantiates a new LogicalNetworkProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LogicalNetworkProfile) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogicalNetworkProfile) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogicalNetworkProfile) SetId(v int32)`

SetId sets Id field to given value.


### GetLabel

`func (o *LogicalNetworkProfile) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LogicalNetworkProfile) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LogicalNetworkProfile) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *LogicalNetworkProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogicalNetworkProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogicalNetworkProfile) SetName(v string)`

SetName sets Name field to given value.


### GetAnnotations

`func (o *LogicalNetworkProfile) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *LogicalNetworkProfile) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *LogicalNetworkProfile) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.


### GetCreatedAt

`func (o *LogicalNetworkProfile) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LogicalNetworkProfile) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LogicalNetworkProfile) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *LogicalNetworkProfile) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LogicalNetworkProfile) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LogicalNetworkProfile) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevision

`func (o *LogicalNetworkProfile) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *LogicalNetworkProfile) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *LogicalNetworkProfile) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetKind

`func (o *LogicalNetworkProfile) GetKind() LogicalNetworkKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *LogicalNetworkProfile) GetKindOk() (*LogicalNetworkKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *LogicalNetworkProfile) SetKind(v LogicalNetworkKind)`

SetKind sets Kind field to given value.


### GetFabricId

`func (o *LogicalNetworkProfile) GetFabricId() int32`

GetFabricId returns the FabricId field if non-nil, zero value otherwise.

### GetFabricIdOk

`func (o *LogicalNetworkProfile) GetFabricIdOk() (*int32, bool)`

GetFabricIdOk returns a tuple with the FabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricId

`func (o *LogicalNetworkProfile) SetFabricId(v int32)`

SetFabricId sets FabricId field to given value.


### GetVlan

`func (o *LogicalNetworkProfile) GetVlan() LogicalNetworkProfileVlanProperties`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *LogicalNetworkProfile) GetVlanOk() (*LogicalNetworkProfileVlanProperties, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *LogicalNetworkProfile) SetVlan(v LogicalNetworkProfileVlanProperties)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *LogicalNetworkProfile) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVxlan

`func (o *LogicalNetworkProfile) GetVxlan() LogicalNetworkProfileVxlanProperties`

GetVxlan returns the Vxlan field if non-nil, zero value otherwise.

### GetVxlanOk

`func (o *LogicalNetworkProfile) GetVxlanOk() (*LogicalNetworkProfileVxlanProperties, bool)`

GetVxlanOk returns a tuple with the Vxlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVxlan

`func (o *LogicalNetworkProfile) SetVxlan(v LogicalNetworkProfileVxlanProperties)`

SetVxlan sets Vxlan field to given value.

### HasVxlan

`func (o *LogicalNetworkProfile) HasVxlan() bool`

HasVxlan returns a boolean if a field has been set.

### GetIpv4

`func (o *LogicalNetworkProfile) GetIpv4() LogicalNetworkProfileIpv4Properties`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *LogicalNetworkProfile) GetIpv4Ok() (*LogicalNetworkProfileIpv4Properties, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *LogicalNetworkProfile) SetIpv4(v LogicalNetworkProfileIpv4Properties)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *LogicalNetworkProfile) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *LogicalNetworkProfile) GetIpv6() LogicalNetworkProfileIpv6Properties`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *LogicalNetworkProfile) GetIpv6Ok() (*LogicalNetworkProfileIpv6Properties, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *LogicalNetworkProfile) SetIpv6(v LogicalNetworkProfileIpv6Properties)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *LogicalNetworkProfile) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetRouteDomainId

`func (o *LogicalNetworkProfile) GetRouteDomainId() int32`

GetRouteDomainId returns the RouteDomainId field if non-nil, zero value otherwise.

### GetRouteDomainIdOk

`func (o *LogicalNetworkProfile) GetRouteDomainIdOk() (*int32, bool)`

GetRouteDomainIdOk returns a tuple with the RouteDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteDomainId

`func (o *LogicalNetworkProfile) SetRouteDomainId(v int32)`

SetRouteDomainId sets RouteDomainId field to given value.

### HasRouteDomainId

`func (o *LogicalNetworkProfile) HasRouteDomainId() bool`

HasRouteDomainId returns a boolean if a field has been set.

### SetRouteDomainIdNil

`func (o *LogicalNetworkProfile) SetRouteDomainIdNil(b bool)`

 SetRouteDomainIdNil sets the value for RouteDomainId to be an explicit nil

### UnsetRouteDomainId
`func (o *LogicalNetworkProfile) UnsetRouteDomainId()`

UnsetRouteDomainId ensures that no value is present for RouteDomainId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


