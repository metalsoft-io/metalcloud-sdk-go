# RouteDomainConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**DeployType** | **string** |  | 
**DeployStatus** | **string** |  | 
**CreatedAt** | **time.Time** | Creation timestamp for the entity. | 
**UpdatedAt** | **time.Time** | Last update timestamp for the entity. | 
**Revision** | **int32** |  | 
**Kind** | [**RouteDomainKind**](RouteDomainKind.md) |  | 
**VrfAllocationStrategies** | [**[]CreateVrfAllocationStrategy**](CreateVrfAllocationStrategy.md) |  | 
**L3VlanAllocationStrategies** | Pointer to [**[]VlanAllocationStrategy**](VlanAllocationStrategy.md) |  | [optional] 
**L3VniAllocationStrategies** | Pointer to [**[]VniAllocationStrategy**](VniAllocationStrategy.md) |  | [optional] 

## Methods

### NewRouteDomainConfig

`func NewRouteDomainConfig(id int32, deployType string, deployStatus string, createdAt time.Time, updatedAt time.Time, revision int32, kind RouteDomainKind, vrfAllocationStrategies []CreateVrfAllocationStrategy, ) *RouteDomainConfig`

NewRouteDomainConfig instantiates a new RouteDomainConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteDomainConfigWithDefaults

`func NewRouteDomainConfigWithDefaults() *RouteDomainConfig`

NewRouteDomainConfigWithDefaults instantiates a new RouteDomainConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RouteDomainConfig) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RouteDomainConfig) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RouteDomainConfig) SetId(v int32)`

SetId sets Id field to given value.


### GetDeployType

`func (o *RouteDomainConfig) GetDeployType() string`

GetDeployType returns the DeployType field if non-nil, zero value otherwise.

### GetDeployTypeOk

`func (o *RouteDomainConfig) GetDeployTypeOk() (*string, bool)`

GetDeployTypeOk returns a tuple with the DeployType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployType

`func (o *RouteDomainConfig) SetDeployType(v string)`

SetDeployType sets DeployType field to given value.


### GetDeployStatus

`func (o *RouteDomainConfig) GetDeployStatus() string`

GetDeployStatus returns the DeployStatus field if non-nil, zero value otherwise.

### GetDeployStatusOk

`func (o *RouteDomainConfig) GetDeployStatusOk() (*string, bool)`

GetDeployStatusOk returns a tuple with the DeployStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStatus

`func (o *RouteDomainConfig) SetDeployStatus(v string)`

SetDeployStatus sets DeployStatus field to given value.


### GetCreatedAt

`func (o *RouteDomainConfig) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RouteDomainConfig) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RouteDomainConfig) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *RouteDomainConfig) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RouteDomainConfig) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RouteDomainConfig) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevision

`func (o *RouteDomainConfig) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *RouteDomainConfig) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *RouteDomainConfig) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetKind

`func (o *RouteDomainConfig) GetKind() RouteDomainKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RouteDomainConfig) GetKindOk() (*RouteDomainKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RouteDomainConfig) SetKind(v RouteDomainKind)`

SetKind sets Kind field to given value.


### GetVrfAllocationStrategies

`func (o *RouteDomainConfig) GetVrfAllocationStrategies() []CreateVrfAllocationStrategy`

GetVrfAllocationStrategies returns the VrfAllocationStrategies field if non-nil, zero value otherwise.

### GetVrfAllocationStrategiesOk

`func (o *RouteDomainConfig) GetVrfAllocationStrategiesOk() (*[]CreateVrfAllocationStrategy, bool)`

GetVrfAllocationStrategiesOk returns a tuple with the VrfAllocationStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfAllocationStrategies

`func (o *RouteDomainConfig) SetVrfAllocationStrategies(v []CreateVrfAllocationStrategy)`

SetVrfAllocationStrategies sets VrfAllocationStrategies field to given value.


### GetL3VlanAllocationStrategies

`func (o *RouteDomainConfig) GetL3VlanAllocationStrategies() []VlanAllocationStrategy`

GetL3VlanAllocationStrategies returns the L3VlanAllocationStrategies field if non-nil, zero value otherwise.

### GetL3VlanAllocationStrategiesOk

`func (o *RouteDomainConfig) GetL3VlanAllocationStrategiesOk() (*[]VlanAllocationStrategy, bool)`

GetL3VlanAllocationStrategiesOk returns a tuple with the L3VlanAllocationStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL3VlanAllocationStrategies

`func (o *RouteDomainConfig) SetL3VlanAllocationStrategies(v []VlanAllocationStrategy)`

SetL3VlanAllocationStrategies sets L3VlanAllocationStrategies field to given value.

### HasL3VlanAllocationStrategies

`func (o *RouteDomainConfig) HasL3VlanAllocationStrategies() bool`

HasL3VlanAllocationStrategies returns a boolean if a field has been set.

### GetL3VniAllocationStrategies

`func (o *RouteDomainConfig) GetL3VniAllocationStrategies() []VniAllocationStrategy`

GetL3VniAllocationStrategies returns the L3VniAllocationStrategies field if non-nil, zero value otherwise.

### GetL3VniAllocationStrategiesOk

`func (o *RouteDomainConfig) GetL3VniAllocationStrategiesOk() (*[]VniAllocationStrategy, bool)`

GetL3VniAllocationStrategiesOk returns a tuple with the L3VniAllocationStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL3VniAllocationStrategies

`func (o *RouteDomainConfig) SetL3VniAllocationStrategies(v []VniAllocationStrategy)`

SetL3VniAllocationStrategies sets L3VniAllocationStrategies field to given value.

### HasL3VniAllocationStrategies

`func (o *RouteDomainConfig) HasL3VniAllocationStrategies() bool`

HasL3VniAllocationStrategies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


