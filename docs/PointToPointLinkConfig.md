# PointToPointLinkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**DeployType** | **string** |  | 
**DeployStatus** | **string** |  | 
**CreatedAt** | **time.Time** | Creation timestamp for the entity. | 
**UpdatedAt** | **time.Time** | Last update timestamp for the entity. | 
**Revision** | **int64** |  | 
**Mtu** | Pointer to **NullableInt32** | Maximum Transmission Unit (MTU) in bytes. | [optional] 
**Ipv4** | Pointer to [**PointToPointLinkConfigIpv4Properties**](PointToPointLinkConfigIpv4Properties.md) |  | [optional] 
**Ipv6** | Pointer to [**PointToPointLinkConfigIpv6Properties**](PointToPointLinkConfigIpv6Properties.md) |  | [optional] 

## Methods

### NewPointToPointLinkConfig

`func NewPointToPointLinkConfig(id int64, deployType string, deployStatus string, createdAt time.Time, updatedAt time.Time, revision int64, ) *PointToPointLinkConfig`

NewPointToPointLinkConfig instantiates a new PointToPointLinkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPointToPointLinkConfigWithDefaults

`func NewPointToPointLinkConfigWithDefaults() *PointToPointLinkConfig`

NewPointToPointLinkConfigWithDefaults instantiates a new PointToPointLinkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PointToPointLinkConfig) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PointToPointLinkConfig) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PointToPointLinkConfig) SetId(v int64)`

SetId sets Id field to given value.


### GetDeployType

`func (o *PointToPointLinkConfig) GetDeployType() string`

GetDeployType returns the DeployType field if non-nil, zero value otherwise.

### GetDeployTypeOk

`func (o *PointToPointLinkConfig) GetDeployTypeOk() (*string, bool)`

GetDeployTypeOk returns a tuple with the DeployType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployType

`func (o *PointToPointLinkConfig) SetDeployType(v string)`

SetDeployType sets DeployType field to given value.


### GetDeployStatus

`func (o *PointToPointLinkConfig) GetDeployStatus() string`

GetDeployStatus returns the DeployStatus field if non-nil, zero value otherwise.

### GetDeployStatusOk

`func (o *PointToPointLinkConfig) GetDeployStatusOk() (*string, bool)`

GetDeployStatusOk returns a tuple with the DeployStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStatus

`func (o *PointToPointLinkConfig) SetDeployStatus(v string)`

SetDeployStatus sets DeployStatus field to given value.


### GetCreatedAt

`func (o *PointToPointLinkConfig) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PointToPointLinkConfig) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PointToPointLinkConfig) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PointToPointLinkConfig) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PointToPointLinkConfig) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PointToPointLinkConfig) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevision

`func (o *PointToPointLinkConfig) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *PointToPointLinkConfig) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *PointToPointLinkConfig) SetRevision(v int64)`

SetRevision sets Revision field to given value.


### GetMtu

`func (o *PointToPointLinkConfig) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *PointToPointLinkConfig) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *PointToPointLinkConfig) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *PointToPointLinkConfig) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *PointToPointLinkConfig) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *PointToPointLinkConfig) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
### GetIpv4

`func (o *PointToPointLinkConfig) GetIpv4() PointToPointLinkConfigIpv4Properties`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *PointToPointLinkConfig) GetIpv4Ok() (*PointToPointLinkConfigIpv4Properties, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *PointToPointLinkConfig) SetIpv4(v PointToPointLinkConfigIpv4Properties)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *PointToPointLinkConfig) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *PointToPointLinkConfig) GetIpv6() PointToPointLinkConfigIpv6Properties`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *PointToPointLinkConfig) GetIpv6Ok() (*PointToPointLinkConfigIpv6Properties, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *PointToPointLinkConfig) SetIpv6(v PointToPointLinkConfigIpv6Properties)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *PointToPointLinkConfig) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


