# RouteDomain

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
**Kind** | [**RouteDomainKind**](RouteDomainKind.md) |  | 
**ServiceStatus** | [**GenericServiceStatus**](GenericServiceStatus.md) |  | 
**PreventVrfCleanup** | **bool** | If true, VRFs belonging to this route domain will not be deleted from switches during cleanup. | [default to false]
**Vrfs** | [**[]VrfAllocation**](VrfAllocation.md) |  | 
**VrfAllocationStrategies** | [**[]VrfAllocationStrategy**](VrfAllocationStrategy.md) |  | 
**Config** | [**RouteDomainConfig**](RouteDomainConfig.md) |  | 
**L3Vlans** | Pointer to [**[]VlanAllocation**](VlanAllocation.md) |  | [optional] 
**L3VlanAllocationStrategies** | Pointer to [**[]VlanAllocationStrategy**](VlanAllocationStrategy.md) |  | [optional] 
**L3Vnis** | Pointer to [**[]VniAllocation**](VniAllocation.md) |  | [optional] 
**L3VniAllocationStrategies** | Pointer to [**[]VniAllocationStrategy**](VniAllocationStrategy.md) |  | [optional] 

## Methods

### NewRouteDomain

`func NewRouteDomain(id int32, label string, name string, annotations map[string]string, createdAt time.Time, updatedAt time.Time, revision int32, kind RouteDomainKind, serviceStatus GenericServiceStatus, preventVrfCleanup bool, vrfs []VrfAllocation, vrfAllocationStrategies []VrfAllocationStrategy, config RouteDomainConfig, ) *RouteDomain`

NewRouteDomain instantiates a new RouteDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteDomainWithDefaults

`func NewRouteDomainWithDefaults() *RouteDomain`

NewRouteDomainWithDefaults instantiates a new RouteDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RouteDomain) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RouteDomain) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RouteDomain) SetId(v int32)`

SetId sets Id field to given value.


### GetLabel

`func (o *RouteDomain) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *RouteDomain) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *RouteDomain) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *RouteDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RouteDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RouteDomain) SetName(v string)`

SetName sets Name field to given value.


### GetAnnotations

`func (o *RouteDomain) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *RouteDomain) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *RouteDomain) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.


### GetCreatedAt

`func (o *RouteDomain) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RouteDomain) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RouteDomain) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *RouteDomain) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RouteDomain) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RouteDomain) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevision

`func (o *RouteDomain) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *RouteDomain) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *RouteDomain) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetKind

`func (o *RouteDomain) GetKind() RouteDomainKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RouteDomain) GetKindOk() (*RouteDomainKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RouteDomain) SetKind(v RouteDomainKind)`

SetKind sets Kind field to given value.


### GetServiceStatus

`func (o *RouteDomain) GetServiceStatus() GenericServiceStatus`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *RouteDomain) GetServiceStatusOk() (*GenericServiceStatus, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *RouteDomain) SetServiceStatus(v GenericServiceStatus)`

SetServiceStatus sets ServiceStatus field to given value.


### GetPreventVrfCleanup

`func (o *RouteDomain) GetPreventVrfCleanup() bool`

GetPreventVrfCleanup returns the PreventVrfCleanup field if non-nil, zero value otherwise.

### GetPreventVrfCleanupOk

`func (o *RouteDomain) GetPreventVrfCleanupOk() (*bool, bool)`

GetPreventVrfCleanupOk returns a tuple with the PreventVrfCleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventVrfCleanup

`func (o *RouteDomain) SetPreventVrfCleanup(v bool)`

SetPreventVrfCleanup sets PreventVrfCleanup field to given value.


### GetVrfs

`func (o *RouteDomain) GetVrfs() []VrfAllocation`

GetVrfs returns the Vrfs field if non-nil, zero value otherwise.

### GetVrfsOk

`func (o *RouteDomain) GetVrfsOk() (*[]VrfAllocation, bool)`

GetVrfsOk returns a tuple with the Vrfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfs

`func (o *RouteDomain) SetVrfs(v []VrfAllocation)`

SetVrfs sets Vrfs field to given value.


### GetVrfAllocationStrategies

`func (o *RouteDomain) GetVrfAllocationStrategies() []VrfAllocationStrategy`

GetVrfAllocationStrategies returns the VrfAllocationStrategies field if non-nil, zero value otherwise.

### GetVrfAllocationStrategiesOk

`func (o *RouteDomain) GetVrfAllocationStrategiesOk() (*[]VrfAllocationStrategy, bool)`

GetVrfAllocationStrategiesOk returns a tuple with the VrfAllocationStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfAllocationStrategies

`func (o *RouteDomain) SetVrfAllocationStrategies(v []VrfAllocationStrategy)`

SetVrfAllocationStrategies sets VrfAllocationStrategies field to given value.


### GetConfig

`func (o *RouteDomain) GetConfig() RouteDomainConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *RouteDomain) GetConfigOk() (*RouteDomainConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *RouteDomain) SetConfig(v RouteDomainConfig)`

SetConfig sets Config field to given value.


### GetL3Vlans

`func (o *RouteDomain) GetL3Vlans() []VlanAllocation`

GetL3Vlans returns the L3Vlans field if non-nil, zero value otherwise.

### GetL3VlansOk

`func (o *RouteDomain) GetL3VlansOk() (*[]VlanAllocation, bool)`

GetL3VlansOk returns a tuple with the L3Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL3Vlans

`func (o *RouteDomain) SetL3Vlans(v []VlanAllocation)`

SetL3Vlans sets L3Vlans field to given value.

### HasL3Vlans

`func (o *RouteDomain) HasL3Vlans() bool`

HasL3Vlans returns a boolean if a field has been set.

### GetL3VlanAllocationStrategies

`func (o *RouteDomain) GetL3VlanAllocationStrategies() []VlanAllocationStrategy`

GetL3VlanAllocationStrategies returns the L3VlanAllocationStrategies field if non-nil, zero value otherwise.

### GetL3VlanAllocationStrategiesOk

`func (o *RouteDomain) GetL3VlanAllocationStrategiesOk() (*[]VlanAllocationStrategy, bool)`

GetL3VlanAllocationStrategiesOk returns a tuple with the L3VlanAllocationStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL3VlanAllocationStrategies

`func (o *RouteDomain) SetL3VlanAllocationStrategies(v []VlanAllocationStrategy)`

SetL3VlanAllocationStrategies sets L3VlanAllocationStrategies field to given value.

### HasL3VlanAllocationStrategies

`func (o *RouteDomain) HasL3VlanAllocationStrategies() bool`

HasL3VlanAllocationStrategies returns a boolean if a field has been set.

### GetL3Vnis

`func (o *RouteDomain) GetL3Vnis() []VniAllocation`

GetL3Vnis returns the L3Vnis field if non-nil, zero value otherwise.

### GetL3VnisOk

`func (o *RouteDomain) GetL3VnisOk() (*[]VniAllocation, bool)`

GetL3VnisOk returns a tuple with the L3Vnis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL3Vnis

`func (o *RouteDomain) SetL3Vnis(v []VniAllocation)`

SetL3Vnis sets L3Vnis field to given value.

### HasL3Vnis

`func (o *RouteDomain) HasL3Vnis() bool`

HasL3Vnis returns a boolean if a field has been set.

### GetL3VniAllocationStrategies

`func (o *RouteDomain) GetL3VniAllocationStrategies() []VniAllocationStrategy`

GetL3VniAllocationStrategies returns the L3VniAllocationStrategies field if non-nil, zero value otherwise.

### GetL3VniAllocationStrategiesOk

`func (o *RouteDomain) GetL3VniAllocationStrategiesOk() (*[]VniAllocationStrategy, bool)`

GetL3VniAllocationStrategiesOk returns a tuple with the L3VniAllocationStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL3VniAllocationStrategies

`func (o *RouteDomain) SetL3VniAllocationStrategies(v []VniAllocationStrategy)`

SetL3VniAllocationStrategies sets L3VniAllocationStrategies field to given value.

### HasL3VniAllocationStrategies

`func (o *RouteDomain) HasL3VniAllocationStrategies() bool`

HasL3VniAllocationStrategies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


