# CreateRouteDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**Kind** | [**RouteDomainKind**](RouteDomainKind.md) |  | 
**VrfAllocationStrategies** | [**[]CreateVrfAllocationStrategy**](CreateVrfAllocationStrategy.md) |  | 
**PreventVrfCleanup** | Pointer to **bool** | If true, VRFs belonging to this route domain will not be deleted from switches during cleanup. | [optional] [default to false]
**ExternalId** | Pointer to **string** | External identifier for this route domain on the controller | [optional] 
**AutoRouteDistinguisher** | Pointer to **bool** | When true, the switch auto-generates the EVPN Route Distinguisher for the L3 VNI. Only applicable to EVPN_L3VPN route domains. Defaults to false. | [optional] 
**AutoRouteTarget** | Pointer to **bool** | When true, the switch auto-generates EVPN Route Targets for the L3 VNI. Only applicable to EVPN_L3VPN route domains. Defaults to false. | [optional] 
**L3VlanAllocationStrategies** | Pointer to [**[]CreateVlanAllocationStrategy**](CreateVlanAllocationStrategy.md) |  | [optional] 
**L3VniAllocationStrategies** | Pointer to [**[]CreateVniAllocationStrategy**](CreateVniAllocationStrategy.md) |  | [optional] 

## Methods

### NewCreateRouteDomain

`func NewCreateRouteDomain(kind RouteDomainKind, vrfAllocationStrategies []CreateVrfAllocationStrategy, ) *CreateRouteDomain`

NewCreateRouteDomain instantiates a new CreateRouteDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRouteDomainWithDefaults

`func NewCreateRouteDomainWithDefaults() *CreateRouteDomain`

NewCreateRouteDomainWithDefaults instantiates a new CreateRouteDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CreateRouteDomain) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateRouteDomain) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateRouteDomain) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateRouteDomain) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *CreateRouteDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRouteDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRouteDomain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateRouteDomain) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAnnotations

`func (o *CreateRouteDomain) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *CreateRouteDomain) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *CreateRouteDomain) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *CreateRouteDomain) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetKind

`func (o *CreateRouteDomain) GetKind() RouteDomainKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateRouteDomain) GetKindOk() (*RouteDomainKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateRouteDomain) SetKind(v RouteDomainKind)`

SetKind sets Kind field to given value.


### GetVrfAllocationStrategies

`func (o *CreateRouteDomain) GetVrfAllocationStrategies() []CreateVrfAllocationStrategy`

GetVrfAllocationStrategies returns the VrfAllocationStrategies field if non-nil, zero value otherwise.

### GetVrfAllocationStrategiesOk

`func (o *CreateRouteDomain) GetVrfAllocationStrategiesOk() (*[]CreateVrfAllocationStrategy, bool)`

GetVrfAllocationStrategiesOk returns a tuple with the VrfAllocationStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfAllocationStrategies

`func (o *CreateRouteDomain) SetVrfAllocationStrategies(v []CreateVrfAllocationStrategy)`

SetVrfAllocationStrategies sets VrfAllocationStrategies field to given value.


### GetPreventVrfCleanup

`func (o *CreateRouteDomain) GetPreventVrfCleanup() bool`

GetPreventVrfCleanup returns the PreventVrfCleanup field if non-nil, zero value otherwise.

### GetPreventVrfCleanupOk

`func (o *CreateRouteDomain) GetPreventVrfCleanupOk() (*bool, bool)`

GetPreventVrfCleanupOk returns a tuple with the PreventVrfCleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventVrfCleanup

`func (o *CreateRouteDomain) SetPreventVrfCleanup(v bool)`

SetPreventVrfCleanup sets PreventVrfCleanup field to given value.

### HasPreventVrfCleanup

`func (o *CreateRouteDomain) HasPreventVrfCleanup() bool`

HasPreventVrfCleanup returns a boolean if a field has been set.

### GetExternalId

`func (o *CreateRouteDomain) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateRouteDomain) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateRouteDomain) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CreateRouteDomain) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetAutoRouteDistinguisher

`func (o *CreateRouteDomain) GetAutoRouteDistinguisher() bool`

GetAutoRouteDistinguisher returns the AutoRouteDistinguisher field if non-nil, zero value otherwise.

### GetAutoRouteDistinguisherOk

`func (o *CreateRouteDomain) GetAutoRouteDistinguisherOk() (*bool, bool)`

GetAutoRouteDistinguisherOk returns a tuple with the AutoRouteDistinguisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRouteDistinguisher

`func (o *CreateRouteDomain) SetAutoRouteDistinguisher(v bool)`

SetAutoRouteDistinguisher sets AutoRouteDistinguisher field to given value.

### HasAutoRouteDistinguisher

`func (o *CreateRouteDomain) HasAutoRouteDistinguisher() bool`

HasAutoRouteDistinguisher returns a boolean if a field has been set.

### GetAutoRouteTarget

`func (o *CreateRouteDomain) GetAutoRouteTarget() bool`

GetAutoRouteTarget returns the AutoRouteTarget field if non-nil, zero value otherwise.

### GetAutoRouteTargetOk

`func (o *CreateRouteDomain) GetAutoRouteTargetOk() (*bool, bool)`

GetAutoRouteTargetOk returns a tuple with the AutoRouteTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRouteTarget

`func (o *CreateRouteDomain) SetAutoRouteTarget(v bool)`

SetAutoRouteTarget sets AutoRouteTarget field to given value.

### HasAutoRouteTarget

`func (o *CreateRouteDomain) HasAutoRouteTarget() bool`

HasAutoRouteTarget returns a boolean if a field has been set.

### GetL3VlanAllocationStrategies

`func (o *CreateRouteDomain) GetL3VlanAllocationStrategies() []CreateVlanAllocationStrategy`

GetL3VlanAllocationStrategies returns the L3VlanAllocationStrategies field if non-nil, zero value otherwise.

### GetL3VlanAllocationStrategiesOk

`func (o *CreateRouteDomain) GetL3VlanAllocationStrategiesOk() (*[]CreateVlanAllocationStrategy, bool)`

GetL3VlanAllocationStrategiesOk returns a tuple with the L3VlanAllocationStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL3VlanAllocationStrategies

`func (o *CreateRouteDomain) SetL3VlanAllocationStrategies(v []CreateVlanAllocationStrategy)`

SetL3VlanAllocationStrategies sets L3VlanAllocationStrategies field to given value.

### HasL3VlanAllocationStrategies

`func (o *CreateRouteDomain) HasL3VlanAllocationStrategies() bool`

HasL3VlanAllocationStrategies returns a boolean if a field has been set.

### GetL3VniAllocationStrategies

`func (o *CreateRouteDomain) GetL3VniAllocationStrategies() []CreateVniAllocationStrategy`

GetL3VniAllocationStrategies returns the L3VniAllocationStrategies field if non-nil, zero value otherwise.

### GetL3VniAllocationStrategiesOk

`func (o *CreateRouteDomain) GetL3VniAllocationStrategiesOk() (*[]CreateVniAllocationStrategy, bool)`

GetL3VniAllocationStrategiesOk returns a tuple with the L3VniAllocationStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL3VniAllocationStrategies

`func (o *CreateRouteDomain) SetL3VniAllocationStrategies(v []CreateVniAllocationStrategy)`

SetL3VniAllocationStrategies sets L3VniAllocationStrategies field to given value.

### HasL3VniAllocationStrategies

`func (o *CreateRouteDomain) HasL3VniAllocationStrategies() bool`

HasL3VniAllocationStrategies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


