# LogicalNetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**DeployType** | **string** |  | 
**DeployStatus** | **string** |  | 
**CreatedAt** | **time.Time** | Creation timestamp for the entity. | 
**UpdatedAt** | **time.Time** | Last update timestamp for the entity. | 
**Revision** | **int32** |  | 
**Kind** | [**LogicalNetworkKind**](LogicalNetworkKind.md) |  | 
**Mtu** | Pointer to **NullableInt32** | Maximum Transmission Unit (MTU) in bytes | [optional] 
**Vlan** | Pointer to [**LogicalNetworkConfigVlanProperties**](LogicalNetworkConfigVlanProperties.md) |  | [optional] 
**Vxlan** | Pointer to [**LogicalNetworkConfigVxlanProperties**](LogicalNetworkConfigVxlanProperties.md) |  | [optional] 
**Pkey** | Pointer to [**LogicalNetworkConfigPkeyProperties**](LogicalNetworkConfigPkeyProperties.md) |  | [optional] 
**Zone** | Pointer to [**LogicalNetworkConfigZoneProperties**](LogicalNetworkConfigZoneProperties.md) |  | [optional] 
**Ipv4** | Pointer to [**LogicalNetworkConfigIpv4Properties**](LogicalNetworkConfigIpv4Properties.md) |  | [optional] 
**Ipv6** | Pointer to [**LogicalNetworkConfigIpv6Properties**](LogicalNetworkConfigIpv6Properties.md) |  | [optional] 

## Methods

### NewLogicalNetworkConfig

`func NewLogicalNetworkConfig(id int32, deployType string, deployStatus string, createdAt time.Time, updatedAt time.Time, revision int32, kind LogicalNetworkKind, ) *LogicalNetworkConfig`

NewLogicalNetworkConfig instantiates a new LogicalNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogicalNetworkConfigWithDefaults

`func NewLogicalNetworkConfigWithDefaults() *LogicalNetworkConfig`

NewLogicalNetworkConfigWithDefaults instantiates a new LogicalNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LogicalNetworkConfig) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogicalNetworkConfig) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogicalNetworkConfig) SetId(v int32)`

SetId sets Id field to given value.


### GetDeployType

`func (o *LogicalNetworkConfig) GetDeployType() string`

GetDeployType returns the DeployType field if non-nil, zero value otherwise.

### GetDeployTypeOk

`func (o *LogicalNetworkConfig) GetDeployTypeOk() (*string, bool)`

GetDeployTypeOk returns a tuple with the DeployType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployType

`func (o *LogicalNetworkConfig) SetDeployType(v string)`

SetDeployType sets DeployType field to given value.


### GetDeployStatus

`func (o *LogicalNetworkConfig) GetDeployStatus() string`

GetDeployStatus returns the DeployStatus field if non-nil, zero value otherwise.

### GetDeployStatusOk

`func (o *LogicalNetworkConfig) GetDeployStatusOk() (*string, bool)`

GetDeployStatusOk returns a tuple with the DeployStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStatus

`func (o *LogicalNetworkConfig) SetDeployStatus(v string)`

SetDeployStatus sets DeployStatus field to given value.


### GetCreatedAt

`func (o *LogicalNetworkConfig) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LogicalNetworkConfig) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LogicalNetworkConfig) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *LogicalNetworkConfig) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LogicalNetworkConfig) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LogicalNetworkConfig) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevision

`func (o *LogicalNetworkConfig) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *LogicalNetworkConfig) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *LogicalNetworkConfig) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetKind

`func (o *LogicalNetworkConfig) GetKind() LogicalNetworkKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *LogicalNetworkConfig) GetKindOk() (*LogicalNetworkKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *LogicalNetworkConfig) SetKind(v LogicalNetworkKind)`

SetKind sets Kind field to given value.


### GetMtu

`func (o *LogicalNetworkConfig) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *LogicalNetworkConfig) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *LogicalNetworkConfig) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *LogicalNetworkConfig) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *LogicalNetworkConfig) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *LogicalNetworkConfig) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
### GetVlan

`func (o *LogicalNetworkConfig) GetVlan() LogicalNetworkConfigVlanProperties`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *LogicalNetworkConfig) GetVlanOk() (*LogicalNetworkConfigVlanProperties, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *LogicalNetworkConfig) SetVlan(v LogicalNetworkConfigVlanProperties)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *LogicalNetworkConfig) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVxlan

`func (o *LogicalNetworkConfig) GetVxlan() LogicalNetworkConfigVxlanProperties`

GetVxlan returns the Vxlan field if non-nil, zero value otherwise.

### GetVxlanOk

`func (o *LogicalNetworkConfig) GetVxlanOk() (*LogicalNetworkConfigVxlanProperties, bool)`

GetVxlanOk returns a tuple with the Vxlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVxlan

`func (o *LogicalNetworkConfig) SetVxlan(v LogicalNetworkConfigVxlanProperties)`

SetVxlan sets Vxlan field to given value.

### HasVxlan

`func (o *LogicalNetworkConfig) HasVxlan() bool`

HasVxlan returns a boolean if a field has been set.

### GetPkey

`func (o *LogicalNetworkConfig) GetPkey() LogicalNetworkConfigPkeyProperties`

GetPkey returns the Pkey field if non-nil, zero value otherwise.

### GetPkeyOk

`func (o *LogicalNetworkConfig) GetPkeyOk() (*LogicalNetworkConfigPkeyProperties, bool)`

GetPkeyOk returns a tuple with the Pkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkey

`func (o *LogicalNetworkConfig) SetPkey(v LogicalNetworkConfigPkeyProperties)`

SetPkey sets Pkey field to given value.

### HasPkey

`func (o *LogicalNetworkConfig) HasPkey() bool`

HasPkey returns a boolean if a field has been set.

### GetZone

`func (o *LogicalNetworkConfig) GetZone() LogicalNetworkConfigZoneProperties`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *LogicalNetworkConfig) GetZoneOk() (*LogicalNetworkConfigZoneProperties, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *LogicalNetworkConfig) SetZone(v LogicalNetworkConfigZoneProperties)`

SetZone sets Zone field to given value.

### HasZone

`func (o *LogicalNetworkConfig) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetIpv4

`func (o *LogicalNetworkConfig) GetIpv4() LogicalNetworkConfigIpv4Properties`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *LogicalNetworkConfig) GetIpv4Ok() (*LogicalNetworkConfigIpv4Properties, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *LogicalNetworkConfig) SetIpv4(v LogicalNetworkConfigIpv4Properties)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *LogicalNetworkConfig) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *LogicalNetworkConfig) GetIpv6() LogicalNetworkConfigIpv6Properties`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *LogicalNetworkConfig) GetIpv6Ok() (*LogicalNetworkConfigIpv6Properties, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *LogicalNetworkConfig) SetIpv6(v LogicalNetworkConfigIpv6Properties)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *LogicalNetworkConfig) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


