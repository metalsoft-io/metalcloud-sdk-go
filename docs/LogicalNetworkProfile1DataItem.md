# LogicalNetworkProfile1DataItem

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
**Vlan** | [**VxlanLogicalNetworkProfileVlanProperties**](VxlanLogicalNetworkProfileVlanProperties.md) |  | 
**Ipv4** | [**VxlanLogicalNetworkProfileIpv4Properties**](VxlanLogicalNetworkProfileIpv4Properties.md) |  | 
**Ipv6** | [**VxlanLogicalNetworkProfileIpv6Properties**](VxlanLogicalNetworkProfileIpv6Properties.md) |  | 
**RouteDomainId** | **NullableInt32** |  | 
**Vxlan** | [**VxlanLogicalNetworkProfileVxlanProperties**](VxlanLogicalNetworkProfileVxlanProperties.md) |  | 

## Methods

### NewLogicalNetworkProfile1DataItem

`func NewLogicalNetworkProfile1DataItem(id int32, label string, name string, annotations map[string]string, createdAt time.Time, updatedAt time.Time, revision int32, kind LogicalNetworkKind, fabricId int32, vlan VxlanLogicalNetworkProfileVlanProperties, ipv4 VxlanLogicalNetworkProfileIpv4Properties, ipv6 VxlanLogicalNetworkProfileIpv6Properties, routeDomainId NullableInt32, vxlan VxlanLogicalNetworkProfileVxlanProperties, ) *LogicalNetworkProfile1DataItem`

NewLogicalNetworkProfile1DataItem instantiates a new LogicalNetworkProfile1DataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogicalNetworkProfile1DataItemWithDefaults

`func NewLogicalNetworkProfile1DataItemWithDefaults() *LogicalNetworkProfile1DataItem`

NewLogicalNetworkProfile1DataItemWithDefaults instantiates a new LogicalNetworkProfile1DataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LogicalNetworkProfile1DataItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogicalNetworkProfile1DataItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogicalNetworkProfile1DataItem) SetId(v int32)`

SetId sets Id field to given value.


### GetLabel

`func (o *LogicalNetworkProfile1DataItem) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LogicalNetworkProfile1DataItem) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LogicalNetworkProfile1DataItem) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *LogicalNetworkProfile1DataItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogicalNetworkProfile1DataItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogicalNetworkProfile1DataItem) SetName(v string)`

SetName sets Name field to given value.


### GetAnnotations

`func (o *LogicalNetworkProfile1DataItem) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *LogicalNetworkProfile1DataItem) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *LogicalNetworkProfile1DataItem) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.


### GetCreatedAt

`func (o *LogicalNetworkProfile1DataItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LogicalNetworkProfile1DataItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LogicalNetworkProfile1DataItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *LogicalNetworkProfile1DataItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LogicalNetworkProfile1DataItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LogicalNetworkProfile1DataItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevision

`func (o *LogicalNetworkProfile1DataItem) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *LogicalNetworkProfile1DataItem) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *LogicalNetworkProfile1DataItem) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetKind

`func (o *LogicalNetworkProfile1DataItem) GetKind() LogicalNetworkKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *LogicalNetworkProfile1DataItem) GetKindOk() (*LogicalNetworkKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *LogicalNetworkProfile1DataItem) SetKind(v LogicalNetworkKind)`

SetKind sets Kind field to given value.


### GetFabricId

`func (o *LogicalNetworkProfile1DataItem) GetFabricId() int32`

GetFabricId returns the FabricId field if non-nil, zero value otherwise.

### GetFabricIdOk

`func (o *LogicalNetworkProfile1DataItem) GetFabricIdOk() (*int32, bool)`

GetFabricIdOk returns a tuple with the FabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricId

`func (o *LogicalNetworkProfile1DataItem) SetFabricId(v int32)`

SetFabricId sets FabricId field to given value.


### GetVlan

`func (o *LogicalNetworkProfile1DataItem) GetVlan() VxlanLogicalNetworkProfileVlanProperties`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *LogicalNetworkProfile1DataItem) GetVlanOk() (*VxlanLogicalNetworkProfileVlanProperties, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *LogicalNetworkProfile1DataItem) SetVlan(v VxlanLogicalNetworkProfileVlanProperties)`

SetVlan sets Vlan field to given value.


### GetIpv4

`func (o *LogicalNetworkProfile1DataItem) GetIpv4() VxlanLogicalNetworkProfileIpv4Properties`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *LogicalNetworkProfile1DataItem) GetIpv4Ok() (*VxlanLogicalNetworkProfileIpv4Properties, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *LogicalNetworkProfile1DataItem) SetIpv4(v VxlanLogicalNetworkProfileIpv4Properties)`

SetIpv4 sets Ipv4 field to given value.


### GetIpv6

`func (o *LogicalNetworkProfile1DataItem) GetIpv6() VxlanLogicalNetworkProfileIpv6Properties`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *LogicalNetworkProfile1DataItem) GetIpv6Ok() (*VxlanLogicalNetworkProfileIpv6Properties, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *LogicalNetworkProfile1DataItem) SetIpv6(v VxlanLogicalNetworkProfileIpv6Properties)`

SetIpv6 sets Ipv6 field to given value.


### GetRouteDomainId

`func (o *LogicalNetworkProfile1DataItem) GetRouteDomainId() int32`

GetRouteDomainId returns the RouteDomainId field if non-nil, zero value otherwise.

### GetRouteDomainIdOk

`func (o *LogicalNetworkProfile1DataItem) GetRouteDomainIdOk() (*int32, bool)`

GetRouteDomainIdOk returns a tuple with the RouteDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteDomainId

`func (o *LogicalNetworkProfile1DataItem) SetRouteDomainId(v int32)`

SetRouteDomainId sets RouteDomainId field to given value.


### SetRouteDomainIdNil

`func (o *LogicalNetworkProfile1DataItem) SetRouteDomainIdNil(b bool)`

 SetRouteDomainIdNil sets the value for RouteDomainId to be an explicit nil

### UnsetRouteDomainId
`func (o *LogicalNetworkProfile1DataItem) UnsetRouteDomainId()`

UnsetRouteDomainId ensures that no value is present for RouteDomainId, not even an explicit nil
### GetVxlan

`func (o *LogicalNetworkProfile1DataItem) GetVxlan() VxlanLogicalNetworkProfileVxlanProperties`

GetVxlan returns the Vxlan field if non-nil, zero value otherwise.

### GetVxlanOk

`func (o *LogicalNetworkProfile1DataItem) GetVxlanOk() (*VxlanLogicalNetworkProfileVxlanProperties, bool)`

GetVxlanOk returns a tuple with the Vxlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVxlan

`func (o *LogicalNetworkProfile1DataItem) SetVxlan(v VxlanLogicalNetworkProfileVxlanProperties)`

SetVxlan sets Vxlan field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


