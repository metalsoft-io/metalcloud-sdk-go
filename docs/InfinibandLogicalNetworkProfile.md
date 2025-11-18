# InfinibandLogicalNetworkProfile

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
**Pkey** | [**InfinibandLogicalNetworkProfilePkeyProperties**](InfinibandLogicalNetworkProfilePkeyProperties.md) |  | 
**Ipv4** | [**InfinibandLogicalNetworkProfileIpv4Properties**](InfinibandLogicalNetworkProfileIpv4Properties.md) |  | 
**Ipv6** | [**InfinibandLogicalNetworkProfileIpv6Properties**](InfinibandLogicalNetworkProfileIpv6Properties.md) |  | 
**RouteDomainId** | **NullableInt32** |  | 
**Mtu** | Pointer to **NullableInt32** | Maximum Transmission Unit (MTU) in bytes | [optional] 

## Methods

### NewInfinibandLogicalNetworkProfile

`func NewInfinibandLogicalNetworkProfile(id int32, label string, name string, annotations map[string]string, createdAt time.Time, updatedAt time.Time, revision int32, kind LogicalNetworkKind, fabricId int32, pkey InfinibandLogicalNetworkProfilePkeyProperties, ipv4 InfinibandLogicalNetworkProfileIpv4Properties, ipv6 InfinibandLogicalNetworkProfileIpv6Properties, routeDomainId NullableInt32, ) *InfinibandLogicalNetworkProfile`

NewInfinibandLogicalNetworkProfile instantiates a new InfinibandLogicalNetworkProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfinibandLogicalNetworkProfileWithDefaults

`func NewInfinibandLogicalNetworkProfileWithDefaults() *InfinibandLogicalNetworkProfile`

NewInfinibandLogicalNetworkProfileWithDefaults instantiates a new InfinibandLogicalNetworkProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InfinibandLogicalNetworkProfile) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InfinibandLogicalNetworkProfile) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InfinibandLogicalNetworkProfile) SetId(v int32)`

SetId sets Id field to given value.


### GetLabel

`func (o *InfinibandLogicalNetworkProfile) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *InfinibandLogicalNetworkProfile) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *InfinibandLogicalNetworkProfile) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *InfinibandLogicalNetworkProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InfinibandLogicalNetworkProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InfinibandLogicalNetworkProfile) SetName(v string)`

SetName sets Name field to given value.


### GetAnnotations

`func (o *InfinibandLogicalNetworkProfile) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *InfinibandLogicalNetworkProfile) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *InfinibandLogicalNetworkProfile) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.


### GetCreatedAt

`func (o *InfinibandLogicalNetworkProfile) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InfinibandLogicalNetworkProfile) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InfinibandLogicalNetworkProfile) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *InfinibandLogicalNetworkProfile) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InfinibandLogicalNetworkProfile) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InfinibandLogicalNetworkProfile) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevision

`func (o *InfinibandLogicalNetworkProfile) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *InfinibandLogicalNetworkProfile) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *InfinibandLogicalNetworkProfile) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetKind

`func (o *InfinibandLogicalNetworkProfile) GetKind() LogicalNetworkKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *InfinibandLogicalNetworkProfile) GetKindOk() (*LogicalNetworkKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *InfinibandLogicalNetworkProfile) SetKind(v LogicalNetworkKind)`

SetKind sets Kind field to given value.


### GetFabricId

`func (o *InfinibandLogicalNetworkProfile) GetFabricId() int32`

GetFabricId returns the FabricId field if non-nil, zero value otherwise.

### GetFabricIdOk

`func (o *InfinibandLogicalNetworkProfile) GetFabricIdOk() (*int32, bool)`

GetFabricIdOk returns a tuple with the FabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricId

`func (o *InfinibandLogicalNetworkProfile) SetFabricId(v int32)`

SetFabricId sets FabricId field to given value.


### GetPkey

`func (o *InfinibandLogicalNetworkProfile) GetPkey() InfinibandLogicalNetworkProfilePkeyProperties`

GetPkey returns the Pkey field if non-nil, zero value otherwise.

### GetPkeyOk

`func (o *InfinibandLogicalNetworkProfile) GetPkeyOk() (*InfinibandLogicalNetworkProfilePkeyProperties, bool)`

GetPkeyOk returns a tuple with the Pkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkey

`func (o *InfinibandLogicalNetworkProfile) SetPkey(v InfinibandLogicalNetworkProfilePkeyProperties)`

SetPkey sets Pkey field to given value.


### GetIpv4

`func (o *InfinibandLogicalNetworkProfile) GetIpv4() InfinibandLogicalNetworkProfileIpv4Properties`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *InfinibandLogicalNetworkProfile) GetIpv4Ok() (*InfinibandLogicalNetworkProfileIpv4Properties, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *InfinibandLogicalNetworkProfile) SetIpv4(v InfinibandLogicalNetworkProfileIpv4Properties)`

SetIpv4 sets Ipv4 field to given value.


### GetIpv6

`func (o *InfinibandLogicalNetworkProfile) GetIpv6() InfinibandLogicalNetworkProfileIpv6Properties`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *InfinibandLogicalNetworkProfile) GetIpv6Ok() (*InfinibandLogicalNetworkProfileIpv6Properties, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *InfinibandLogicalNetworkProfile) SetIpv6(v InfinibandLogicalNetworkProfileIpv6Properties)`

SetIpv6 sets Ipv6 field to given value.


### GetRouteDomainId

`func (o *InfinibandLogicalNetworkProfile) GetRouteDomainId() int32`

GetRouteDomainId returns the RouteDomainId field if non-nil, zero value otherwise.

### GetRouteDomainIdOk

`func (o *InfinibandLogicalNetworkProfile) GetRouteDomainIdOk() (*int32, bool)`

GetRouteDomainIdOk returns a tuple with the RouteDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteDomainId

`func (o *InfinibandLogicalNetworkProfile) SetRouteDomainId(v int32)`

SetRouteDomainId sets RouteDomainId field to given value.


### SetRouteDomainIdNil

`func (o *InfinibandLogicalNetworkProfile) SetRouteDomainIdNil(b bool)`

 SetRouteDomainIdNil sets the value for RouteDomainId to be an explicit nil

### UnsetRouteDomainId
`func (o *InfinibandLogicalNetworkProfile) UnsetRouteDomainId()`

UnsetRouteDomainId ensures that no value is present for RouteDomainId, not even an explicit nil
### GetMtu

`func (o *InfinibandLogicalNetworkProfile) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *InfinibandLogicalNetworkProfile) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *InfinibandLogicalNetworkProfile) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *InfinibandLogicalNetworkProfile) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *InfinibandLogicalNetworkProfile) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *InfinibandLogicalNetworkProfile) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


