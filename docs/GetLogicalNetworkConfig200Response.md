# GetLogicalNetworkConfig200Response

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
**Vlan** | [**VxlanLogicalNetworkConfigVlanProperties**](VxlanLogicalNetworkConfigVlanProperties.md) |  | 
**Ipv4** | [**VxlanLogicalNetworkConfigIpv4Properties**](VxlanLogicalNetworkConfigIpv4Properties.md) |  | 
**Ipv6** | [**VxlanLogicalNetworkConfigIpv6Properties**](VxlanLogicalNetworkConfigIpv6Properties.md) |  | 
**Vxlan** | [**VxlanLogicalNetworkConfigVxlanProperties**](VxlanLogicalNetworkConfigVxlanProperties.md) |  | 

## Methods

### NewGetLogicalNetworkConfig200Response

`func NewGetLogicalNetworkConfig200Response(id int32, deployType string, deployStatus string, createdAt time.Time, updatedAt time.Time, revision int32, kind LogicalNetworkKind, vlan VxlanLogicalNetworkConfigVlanProperties, ipv4 VxlanLogicalNetworkConfigIpv4Properties, ipv6 VxlanLogicalNetworkConfigIpv6Properties, vxlan VxlanLogicalNetworkConfigVxlanProperties, ) *GetLogicalNetworkConfig200Response`

NewGetLogicalNetworkConfig200Response instantiates a new GetLogicalNetworkConfig200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLogicalNetworkConfig200ResponseWithDefaults

`func NewGetLogicalNetworkConfig200ResponseWithDefaults() *GetLogicalNetworkConfig200Response`

NewGetLogicalNetworkConfig200ResponseWithDefaults instantiates a new GetLogicalNetworkConfig200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetLogicalNetworkConfig200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetLogicalNetworkConfig200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetLogicalNetworkConfig200Response) SetId(v int32)`

SetId sets Id field to given value.


### GetDeployType

`func (o *GetLogicalNetworkConfig200Response) GetDeployType() string`

GetDeployType returns the DeployType field if non-nil, zero value otherwise.

### GetDeployTypeOk

`func (o *GetLogicalNetworkConfig200Response) GetDeployTypeOk() (*string, bool)`

GetDeployTypeOk returns a tuple with the DeployType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployType

`func (o *GetLogicalNetworkConfig200Response) SetDeployType(v string)`

SetDeployType sets DeployType field to given value.


### GetDeployStatus

`func (o *GetLogicalNetworkConfig200Response) GetDeployStatus() string`

GetDeployStatus returns the DeployStatus field if non-nil, zero value otherwise.

### GetDeployStatusOk

`func (o *GetLogicalNetworkConfig200Response) GetDeployStatusOk() (*string, bool)`

GetDeployStatusOk returns a tuple with the DeployStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStatus

`func (o *GetLogicalNetworkConfig200Response) SetDeployStatus(v string)`

SetDeployStatus sets DeployStatus field to given value.


### GetCreatedAt

`func (o *GetLogicalNetworkConfig200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetLogicalNetworkConfig200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetLogicalNetworkConfig200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *GetLogicalNetworkConfig200Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetLogicalNetworkConfig200Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetLogicalNetworkConfig200Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevision

`func (o *GetLogicalNetworkConfig200Response) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *GetLogicalNetworkConfig200Response) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *GetLogicalNetworkConfig200Response) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetKind

`func (o *GetLogicalNetworkConfig200Response) GetKind() LogicalNetworkKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetLogicalNetworkConfig200Response) GetKindOk() (*LogicalNetworkKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetLogicalNetworkConfig200Response) SetKind(v LogicalNetworkKind)`

SetKind sets Kind field to given value.


### GetVlan

`func (o *GetLogicalNetworkConfig200Response) GetVlan() VxlanLogicalNetworkConfigVlanProperties`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *GetLogicalNetworkConfig200Response) GetVlanOk() (*VxlanLogicalNetworkConfigVlanProperties, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *GetLogicalNetworkConfig200Response) SetVlan(v VxlanLogicalNetworkConfigVlanProperties)`

SetVlan sets Vlan field to given value.


### GetIpv4

`func (o *GetLogicalNetworkConfig200Response) GetIpv4() VxlanLogicalNetworkConfigIpv4Properties`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *GetLogicalNetworkConfig200Response) GetIpv4Ok() (*VxlanLogicalNetworkConfigIpv4Properties, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *GetLogicalNetworkConfig200Response) SetIpv4(v VxlanLogicalNetworkConfigIpv4Properties)`

SetIpv4 sets Ipv4 field to given value.


### GetIpv6

`func (o *GetLogicalNetworkConfig200Response) GetIpv6() VxlanLogicalNetworkConfigIpv6Properties`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *GetLogicalNetworkConfig200Response) GetIpv6Ok() (*VxlanLogicalNetworkConfigIpv6Properties, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *GetLogicalNetworkConfig200Response) SetIpv6(v VxlanLogicalNetworkConfigIpv6Properties)`

SetIpv6 sets Ipv6 field to given value.


### GetVxlan

`func (o *GetLogicalNetworkConfig200Response) GetVxlan() VxlanLogicalNetworkConfigVxlanProperties`

GetVxlan returns the Vxlan field if non-nil, zero value otherwise.

### GetVxlanOk

`func (o *GetLogicalNetworkConfig200Response) GetVxlanOk() (*VxlanLogicalNetworkConfigVxlanProperties, bool)`

GetVxlanOk returns a tuple with the Vxlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVxlan

`func (o *GetLogicalNetworkConfig200Response) SetVxlan(v VxlanLogicalNetworkConfigVxlanProperties)`

SetVxlan sets Vxlan field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


