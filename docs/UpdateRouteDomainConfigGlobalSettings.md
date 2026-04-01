# UpdateRouteDomainConfigGlobalSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoRouteDistinguisher** | Pointer to **bool** | When true, the switch auto-generates the EVPN Route Distinguisher for the L3 VNI instead of Metalsoft computing it. | [optional] 
**AutoRouteTarget** | Pointer to **bool** | When true, the switch auto-generates EVPN Route Targets for the L3 VNI instead of Metalsoft computing them. | [optional] 

## Methods

### NewUpdateRouteDomainConfigGlobalSettings

`func NewUpdateRouteDomainConfigGlobalSettings() *UpdateRouteDomainConfigGlobalSettings`

NewUpdateRouteDomainConfigGlobalSettings instantiates a new UpdateRouteDomainConfigGlobalSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRouteDomainConfigGlobalSettingsWithDefaults

`func NewUpdateRouteDomainConfigGlobalSettingsWithDefaults() *UpdateRouteDomainConfigGlobalSettings`

NewUpdateRouteDomainConfigGlobalSettingsWithDefaults instantiates a new UpdateRouteDomainConfigGlobalSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoRouteDistinguisher

`func (o *UpdateRouteDomainConfigGlobalSettings) GetAutoRouteDistinguisher() bool`

GetAutoRouteDistinguisher returns the AutoRouteDistinguisher field if non-nil, zero value otherwise.

### GetAutoRouteDistinguisherOk

`func (o *UpdateRouteDomainConfigGlobalSettings) GetAutoRouteDistinguisherOk() (*bool, bool)`

GetAutoRouteDistinguisherOk returns a tuple with the AutoRouteDistinguisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRouteDistinguisher

`func (o *UpdateRouteDomainConfigGlobalSettings) SetAutoRouteDistinguisher(v bool)`

SetAutoRouteDistinguisher sets AutoRouteDistinguisher field to given value.

### HasAutoRouteDistinguisher

`func (o *UpdateRouteDomainConfigGlobalSettings) HasAutoRouteDistinguisher() bool`

HasAutoRouteDistinguisher returns a boolean if a field has been set.

### GetAutoRouteTarget

`func (o *UpdateRouteDomainConfigGlobalSettings) GetAutoRouteTarget() bool`

GetAutoRouteTarget returns the AutoRouteTarget field if non-nil, zero value otherwise.

### GetAutoRouteTargetOk

`func (o *UpdateRouteDomainConfigGlobalSettings) GetAutoRouteTargetOk() (*bool, bool)`

GetAutoRouteTargetOk returns a tuple with the AutoRouteTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRouteTarget

`func (o *UpdateRouteDomainConfigGlobalSettings) SetAutoRouteTarget(v bool)`

SetAutoRouteTarget sets AutoRouteTarget field to given value.

### HasAutoRouteTarget

`func (o *UpdateRouteDomainConfigGlobalSettings) HasAutoRouteTarget() bool`

HasAutoRouteTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


