# VlanLogicalNetworkConfig

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
**Vlan** | [**VlanLogicalNetworkConfigVlanProperties**](VlanLogicalNetworkConfigVlanProperties.md) |  | 
**Ipv4** | [**VlanLogicalNetworkConfigIpv4Properties**](VlanLogicalNetworkConfigIpv4Properties.md) |  | 
**Ipv6** | [**VlanLogicalNetworkConfigIpv6Properties**](VlanLogicalNetworkConfigIpv6Properties.md) |  | 

## Methods

### NewVlanLogicalNetworkConfig

`func NewVlanLogicalNetworkConfig(id int32, deployType string, deployStatus string, createdAt time.Time, updatedAt time.Time, revision int32, kind LogicalNetworkKind, vlan VlanLogicalNetworkConfigVlanProperties, ipv4 VlanLogicalNetworkConfigIpv4Properties, ipv6 VlanLogicalNetworkConfigIpv6Properties, ) *VlanLogicalNetworkConfig`

NewVlanLogicalNetworkConfig instantiates a new VlanLogicalNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanLogicalNetworkConfigWithDefaults

`func NewVlanLogicalNetworkConfigWithDefaults() *VlanLogicalNetworkConfig`

NewVlanLogicalNetworkConfigWithDefaults instantiates a new VlanLogicalNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VlanLogicalNetworkConfig) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VlanLogicalNetworkConfig) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VlanLogicalNetworkConfig) SetId(v int32)`

SetId sets Id field to given value.


### GetDeployType

`func (o *VlanLogicalNetworkConfig) GetDeployType() string`

GetDeployType returns the DeployType field if non-nil, zero value otherwise.

### GetDeployTypeOk

`func (o *VlanLogicalNetworkConfig) GetDeployTypeOk() (*string, bool)`

GetDeployTypeOk returns a tuple with the DeployType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployType

`func (o *VlanLogicalNetworkConfig) SetDeployType(v string)`

SetDeployType sets DeployType field to given value.


### GetDeployStatus

`func (o *VlanLogicalNetworkConfig) GetDeployStatus() string`

GetDeployStatus returns the DeployStatus field if non-nil, zero value otherwise.

### GetDeployStatusOk

`func (o *VlanLogicalNetworkConfig) GetDeployStatusOk() (*string, bool)`

GetDeployStatusOk returns a tuple with the DeployStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStatus

`func (o *VlanLogicalNetworkConfig) SetDeployStatus(v string)`

SetDeployStatus sets DeployStatus field to given value.


### GetCreatedAt

`func (o *VlanLogicalNetworkConfig) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VlanLogicalNetworkConfig) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VlanLogicalNetworkConfig) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *VlanLogicalNetworkConfig) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VlanLogicalNetworkConfig) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VlanLogicalNetworkConfig) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevision

`func (o *VlanLogicalNetworkConfig) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *VlanLogicalNetworkConfig) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *VlanLogicalNetworkConfig) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetKind

`func (o *VlanLogicalNetworkConfig) GetKind() LogicalNetworkKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *VlanLogicalNetworkConfig) GetKindOk() (*LogicalNetworkKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *VlanLogicalNetworkConfig) SetKind(v LogicalNetworkKind)`

SetKind sets Kind field to given value.


### GetMtu

`func (o *VlanLogicalNetworkConfig) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *VlanLogicalNetworkConfig) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *VlanLogicalNetworkConfig) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *VlanLogicalNetworkConfig) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *VlanLogicalNetworkConfig) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *VlanLogicalNetworkConfig) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
### GetVlan

`func (o *VlanLogicalNetworkConfig) GetVlan() VlanLogicalNetworkConfigVlanProperties`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *VlanLogicalNetworkConfig) GetVlanOk() (*VlanLogicalNetworkConfigVlanProperties, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *VlanLogicalNetworkConfig) SetVlan(v VlanLogicalNetworkConfigVlanProperties)`

SetVlan sets Vlan field to given value.


### GetIpv4

`func (o *VlanLogicalNetworkConfig) GetIpv4() VlanLogicalNetworkConfigIpv4Properties`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *VlanLogicalNetworkConfig) GetIpv4Ok() (*VlanLogicalNetworkConfigIpv4Properties, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *VlanLogicalNetworkConfig) SetIpv4(v VlanLogicalNetworkConfigIpv4Properties)`

SetIpv4 sets Ipv4 field to given value.


### GetIpv6

`func (o *VlanLogicalNetworkConfig) GetIpv6() VlanLogicalNetworkConfigIpv6Properties`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *VlanLogicalNetworkConfig) GetIpv6Ok() (*VlanLogicalNetworkConfigIpv6Properties, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *VlanLogicalNetworkConfig) SetIpv6(v VlanLogicalNetworkConfigIpv6Properties)`

SetIpv6 sets Ipv6 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


