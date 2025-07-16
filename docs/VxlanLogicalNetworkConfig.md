# VxlanLogicalNetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**DeployType** | **string** |  | 
**DeployStatus** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Revision** | **int32** |  | 
**Kind** | [**LogicalNetworkKind**](LogicalNetworkKind.md) |  | 
**Mtu** | Pointer to **NullableInt32** | Maximum Transmission Unit (MTU) in bytes | [optional] 
**Vlan** | [**VxlanLogicalNetworkConfigVlanProperties**](VxlanLogicalNetworkConfigVlanProperties.md) |  | 
**Vxlan** | [**VxlanLogicalNetworkConfigVxlanProperties**](VxlanLogicalNetworkConfigVxlanProperties.md) |  | 
**Ipv4** | [**VxlanLogicalNetworkConfigIpv4Properties**](VxlanLogicalNetworkConfigIpv4Properties.md) |  | 
**Ipv6** | [**VxlanLogicalNetworkConfigIpv6Properties**](VxlanLogicalNetworkConfigIpv6Properties.md) |  | 

## Methods

### NewVxlanLogicalNetworkConfig

`func NewVxlanLogicalNetworkConfig(id int32, deployType string, deployStatus string, createdAt time.Time, updatedAt time.Time, revision int32, kind LogicalNetworkKind, vlan VxlanLogicalNetworkConfigVlanProperties, vxlan VxlanLogicalNetworkConfigVxlanProperties, ipv4 VxlanLogicalNetworkConfigIpv4Properties, ipv6 VxlanLogicalNetworkConfigIpv6Properties, ) *VxlanLogicalNetworkConfig`

NewVxlanLogicalNetworkConfig instantiates a new VxlanLogicalNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVxlanLogicalNetworkConfigWithDefaults

`func NewVxlanLogicalNetworkConfigWithDefaults() *VxlanLogicalNetworkConfig`

NewVxlanLogicalNetworkConfigWithDefaults instantiates a new VxlanLogicalNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VxlanLogicalNetworkConfig) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VxlanLogicalNetworkConfig) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VxlanLogicalNetworkConfig) SetId(v int32)`

SetId sets Id field to given value.


### GetDeployType

`func (o *VxlanLogicalNetworkConfig) GetDeployType() string`

GetDeployType returns the DeployType field if non-nil, zero value otherwise.

### GetDeployTypeOk

`func (o *VxlanLogicalNetworkConfig) GetDeployTypeOk() (*string, bool)`

GetDeployTypeOk returns a tuple with the DeployType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployType

`func (o *VxlanLogicalNetworkConfig) SetDeployType(v string)`

SetDeployType sets DeployType field to given value.


### GetDeployStatus

`func (o *VxlanLogicalNetworkConfig) GetDeployStatus() string`

GetDeployStatus returns the DeployStatus field if non-nil, zero value otherwise.

### GetDeployStatusOk

`func (o *VxlanLogicalNetworkConfig) GetDeployStatusOk() (*string, bool)`

GetDeployStatusOk returns a tuple with the DeployStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStatus

`func (o *VxlanLogicalNetworkConfig) SetDeployStatus(v string)`

SetDeployStatus sets DeployStatus field to given value.


### GetCreatedAt

`func (o *VxlanLogicalNetworkConfig) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VxlanLogicalNetworkConfig) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VxlanLogicalNetworkConfig) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *VxlanLogicalNetworkConfig) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VxlanLogicalNetworkConfig) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VxlanLogicalNetworkConfig) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevision

`func (o *VxlanLogicalNetworkConfig) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *VxlanLogicalNetworkConfig) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *VxlanLogicalNetworkConfig) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetKind

`func (o *VxlanLogicalNetworkConfig) GetKind() LogicalNetworkKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *VxlanLogicalNetworkConfig) GetKindOk() (*LogicalNetworkKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *VxlanLogicalNetworkConfig) SetKind(v LogicalNetworkKind)`

SetKind sets Kind field to given value.


### GetMtu

`func (o *VxlanLogicalNetworkConfig) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *VxlanLogicalNetworkConfig) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *VxlanLogicalNetworkConfig) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *VxlanLogicalNetworkConfig) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *VxlanLogicalNetworkConfig) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *VxlanLogicalNetworkConfig) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
### GetVlan

`func (o *VxlanLogicalNetworkConfig) GetVlan() VxlanLogicalNetworkConfigVlanProperties`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *VxlanLogicalNetworkConfig) GetVlanOk() (*VxlanLogicalNetworkConfigVlanProperties, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *VxlanLogicalNetworkConfig) SetVlan(v VxlanLogicalNetworkConfigVlanProperties)`

SetVlan sets Vlan field to given value.


### GetVxlan

`func (o *VxlanLogicalNetworkConfig) GetVxlan() VxlanLogicalNetworkConfigVxlanProperties`

GetVxlan returns the Vxlan field if non-nil, zero value otherwise.

### GetVxlanOk

`func (o *VxlanLogicalNetworkConfig) GetVxlanOk() (*VxlanLogicalNetworkConfigVxlanProperties, bool)`

GetVxlanOk returns a tuple with the Vxlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVxlan

`func (o *VxlanLogicalNetworkConfig) SetVxlan(v VxlanLogicalNetworkConfigVxlanProperties)`

SetVxlan sets Vxlan field to given value.


### GetIpv4

`func (o *VxlanLogicalNetworkConfig) GetIpv4() VxlanLogicalNetworkConfigIpv4Properties`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *VxlanLogicalNetworkConfig) GetIpv4Ok() (*VxlanLogicalNetworkConfigIpv4Properties, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *VxlanLogicalNetworkConfig) SetIpv4(v VxlanLogicalNetworkConfigIpv4Properties)`

SetIpv4 sets Ipv4 field to given value.


### GetIpv6

`func (o *VxlanLogicalNetworkConfig) GetIpv6() VxlanLogicalNetworkConfigIpv6Properties`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *VxlanLogicalNetworkConfig) GetIpv6Ok() (*VxlanLogicalNetworkConfigIpv6Properties, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *VxlanLogicalNetworkConfig) SetIpv6(v VxlanLogicalNetworkConfigIpv6Properties)`

SetIpv6 sets Ipv6 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


