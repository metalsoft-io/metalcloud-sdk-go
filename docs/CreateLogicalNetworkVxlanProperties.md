# CreateLogicalNetworkVxlanProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VniAllocationStrategies** | [**[]CreateVniAllocationStrategy**](CreateVniAllocationStrategy.md) |  | 
**NeighSuppress** | Pointer to **bool** | Controls neigh-suppress on switches for this logical network. Defaults to true (enabled). Set to false to disable neigh-suppress | [optional] 
**AutoRouteDistinguisher** | Pointer to **bool** | When true, the switch auto-generates the EVPN Route Distinguisher instead of Metalsoft computing it. Defaults to false. | [optional] 
**AutoRouteTarget** | Pointer to **bool** | When true, the switch auto-generates EVPN Route Targets instead of Metalsoft computing them. Defaults to false. | [optional] 

## Methods

### NewCreateLogicalNetworkVxlanProperties

`func NewCreateLogicalNetworkVxlanProperties(vniAllocationStrategies []CreateVniAllocationStrategy, ) *CreateLogicalNetworkVxlanProperties`

NewCreateLogicalNetworkVxlanProperties instantiates a new CreateLogicalNetworkVxlanProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLogicalNetworkVxlanPropertiesWithDefaults

`func NewCreateLogicalNetworkVxlanPropertiesWithDefaults() *CreateLogicalNetworkVxlanProperties`

NewCreateLogicalNetworkVxlanPropertiesWithDefaults instantiates a new CreateLogicalNetworkVxlanProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVniAllocationStrategies

`func (o *CreateLogicalNetworkVxlanProperties) GetVniAllocationStrategies() []CreateVniAllocationStrategy`

GetVniAllocationStrategies returns the VniAllocationStrategies field if non-nil, zero value otherwise.

### GetVniAllocationStrategiesOk

`func (o *CreateLogicalNetworkVxlanProperties) GetVniAllocationStrategiesOk() (*[]CreateVniAllocationStrategy, bool)`

GetVniAllocationStrategiesOk returns a tuple with the VniAllocationStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVniAllocationStrategies

`func (o *CreateLogicalNetworkVxlanProperties) SetVniAllocationStrategies(v []CreateVniAllocationStrategy)`

SetVniAllocationStrategies sets VniAllocationStrategies field to given value.


### GetNeighSuppress

`func (o *CreateLogicalNetworkVxlanProperties) GetNeighSuppress() bool`

GetNeighSuppress returns the NeighSuppress field if non-nil, zero value otherwise.

### GetNeighSuppressOk

`func (o *CreateLogicalNetworkVxlanProperties) GetNeighSuppressOk() (*bool, bool)`

GetNeighSuppressOk returns a tuple with the NeighSuppress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighSuppress

`func (o *CreateLogicalNetworkVxlanProperties) SetNeighSuppress(v bool)`

SetNeighSuppress sets NeighSuppress field to given value.

### HasNeighSuppress

`func (o *CreateLogicalNetworkVxlanProperties) HasNeighSuppress() bool`

HasNeighSuppress returns a boolean if a field has been set.

### GetAutoRouteDistinguisher

`func (o *CreateLogicalNetworkVxlanProperties) GetAutoRouteDistinguisher() bool`

GetAutoRouteDistinguisher returns the AutoRouteDistinguisher field if non-nil, zero value otherwise.

### GetAutoRouteDistinguisherOk

`func (o *CreateLogicalNetworkVxlanProperties) GetAutoRouteDistinguisherOk() (*bool, bool)`

GetAutoRouteDistinguisherOk returns a tuple with the AutoRouteDistinguisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRouteDistinguisher

`func (o *CreateLogicalNetworkVxlanProperties) SetAutoRouteDistinguisher(v bool)`

SetAutoRouteDistinguisher sets AutoRouteDistinguisher field to given value.

### HasAutoRouteDistinguisher

`func (o *CreateLogicalNetworkVxlanProperties) HasAutoRouteDistinguisher() bool`

HasAutoRouteDistinguisher returns a boolean if a field has been set.

### GetAutoRouteTarget

`func (o *CreateLogicalNetworkVxlanProperties) GetAutoRouteTarget() bool`

GetAutoRouteTarget returns the AutoRouteTarget field if non-nil, zero value otherwise.

### GetAutoRouteTargetOk

`func (o *CreateLogicalNetworkVxlanProperties) GetAutoRouteTargetOk() (*bool, bool)`

GetAutoRouteTargetOk returns a tuple with the AutoRouteTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRouteTarget

`func (o *CreateLogicalNetworkVxlanProperties) SetAutoRouteTarget(v bool)`

SetAutoRouteTarget sets AutoRouteTarget field to given value.

### HasAutoRouteTarget

`func (o *CreateLogicalNetworkVxlanProperties) HasAutoRouteTarget() bool`

HasAutoRouteTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


