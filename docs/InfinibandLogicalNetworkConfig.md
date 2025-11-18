# InfinibandLogicalNetworkConfig

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
**Pkey** | [**InfinibandLogicalNetworkConfigPkeyProperties**](InfinibandLogicalNetworkConfigPkeyProperties.md) |  | 
**Ipv4** | [**InfinibandLogicalNetworkConfigIpv4Properties**](InfinibandLogicalNetworkConfigIpv4Properties.md) |  | 
**Ipv6** | [**InfinibandLogicalNetworkConfigIpv6Properties**](InfinibandLogicalNetworkConfigIpv6Properties.md) |  | 

## Methods

### NewInfinibandLogicalNetworkConfig

`func NewInfinibandLogicalNetworkConfig(id int32, deployType string, deployStatus string, createdAt time.Time, updatedAt time.Time, revision int32, kind LogicalNetworkKind, pkey InfinibandLogicalNetworkConfigPkeyProperties, ipv4 InfinibandLogicalNetworkConfigIpv4Properties, ipv6 InfinibandLogicalNetworkConfigIpv6Properties, ) *InfinibandLogicalNetworkConfig`

NewInfinibandLogicalNetworkConfig instantiates a new InfinibandLogicalNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfinibandLogicalNetworkConfigWithDefaults

`func NewInfinibandLogicalNetworkConfigWithDefaults() *InfinibandLogicalNetworkConfig`

NewInfinibandLogicalNetworkConfigWithDefaults instantiates a new InfinibandLogicalNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InfinibandLogicalNetworkConfig) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InfinibandLogicalNetworkConfig) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InfinibandLogicalNetworkConfig) SetId(v int32)`

SetId sets Id field to given value.


### GetDeployType

`func (o *InfinibandLogicalNetworkConfig) GetDeployType() string`

GetDeployType returns the DeployType field if non-nil, zero value otherwise.

### GetDeployTypeOk

`func (o *InfinibandLogicalNetworkConfig) GetDeployTypeOk() (*string, bool)`

GetDeployTypeOk returns a tuple with the DeployType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployType

`func (o *InfinibandLogicalNetworkConfig) SetDeployType(v string)`

SetDeployType sets DeployType field to given value.


### GetDeployStatus

`func (o *InfinibandLogicalNetworkConfig) GetDeployStatus() string`

GetDeployStatus returns the DeployStatus field if non-nil, zero value otherwise.

### GetDeployStatusOk

`func (o *InfinibandLogicalNetworkConfig) GetDeployStatusOk() (*string, bool)`

GetDeployStatusOk returns a tuple with the DeployStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStatus

`func (o *InfinibandLogicalNetworkConfig) SetDeployStatus(v string)`

SetDeployStatus sets DeployStatus field to given value.


### GetCreatedAt

`func (o *InfinibandLogicalNetworkConfig) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InfinibandLogicalNetworkConfig) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InfinibandLogicalNetworkConfig) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *InfinibandLogicalNetworkConfig) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InfinibandLogicalNetworkConfig) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InfinibandLogicalNetworkConfig) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevision

`func (o *InfinibandLogicalNetworkConfig) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *InfinibandLogicalNetworkConfig) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *InfinibandLogicalNetworkConfig) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetKind

`func (o *InfinibandLogicalNetworkConfig) GetKind() LogicalNetworkKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *InfinibandLogicalNetworkConfig) GetKindOk() (*LogicalNetworkKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *InfinibandLogicalNetworkConfig) SetKind(v LogicalNetworkKind)`

SetKind sets Kind field to given value.


### GetMtu

`func (o *InfinibandLogicalNetworkConfig) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *InfinibandLogicalNetworkConfig) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *InfinibandLogicalNetworkConfig) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *InfinibandLogicalNetworkConfig) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *InfinibandLogicalNetworkConfig) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *InfinibandLogicalNetworkConfig) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
### GetPkey

`func (o *InfinibandLogicalNetworkConfig) GetPkey() InfinibandLogicalNetworkConfigPkeyProperties`

GetPkey returns the Pkey field if non-nil, zero value otherwise.

### GetPkeyOk

`func (o *InfinibandLogicalNetworkConfig) GetPkeyOk() (*InfinibandLogicalNetworkConfigPkeyProperties, bool)`

GetPkeyOk returns a tuple with the Pkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkey

`func (o *InfinibandLogicalNetworkConfig) SetPkey(v InfinibandLogicalNetworkConfigPkeyProperties)`

SetPkey sets Pkey field to given value.


### GetIpv4

`func (o *InfinibandLogicalNetworkConfig) GetIpv4() InfinibandLogicalNetworkConfigIpv4Properties`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *InfinibandLogicalNetworkConfig) GetIpv4Ok() (*InfinibandLogicalNetworkConfigIpv4Properties, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *InfinibandLogicalNetworkConfig) SetIpv4(v InfinibandLogicalNetworkConfigIpv4Properties)`

SetIpv4 sets Ipv4 field to given value.


### GetIpv6

`func (o *InfinibandLogicalNetworkConfig) GetIpv6() InfinibandLogicalNetworkConfigIpv6Properties`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *InfinibandLogicalNetworkConfig) GetIpv6Ok() (*InfinibandLogicalNetworkConfigIpv6Properties, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *InfinibandLogicalNetworkConfig) SetIpv6(v InfinibandLogicalNetworkConfigIpv6Properties)`

SetIpv6 sets Ipv6 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


