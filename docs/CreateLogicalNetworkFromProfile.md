# CreateLogicalNetworkFromProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**LogicalNetworkProfileId** | **int32** |  | 
**InfrastructureId** | Pointer to **NullableInt32** |  | [optional] 
**Mtu** | Pointer to **NullableInt32** | Maximum Transmission Unit (MTU) in bytes | [optional] 
**ExtensionInstanceId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCreateLogicalNetworkFromProfile

`func NewCreateLogicalNetworkFromProfile(logicalNetworkProfileId int32, ) *CreateLogicalNetworkFromProfile`

NewCreateLogicalNetworkFromProfile instantiates a new CreateLogicalNetworkFromProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLogicalNetworkFromProfileWithDefaults

`func NewCreateLogicalNetworkFromProfileWithDefaults() *CreateLogicalNetworkFromProfile`

NewCreateLogicalNetworkFromProfileWithDefaults instantiates a new CreateLogicalNetworkFromProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CreateLogicalNetworkFromProfile) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateLogicalNetworkFromProfile) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateLogicalNetworkFromProfile) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateLogicalNetworkFromProfile) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *CreateLogicalNetworkFromProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLogicalNetworkFromProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLogicalNetworkFromProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateLogicalNetworkFromProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAnnotations

`func (o *CreateLogicalNetworkFromProfile) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *CreateLogicalNetworkFromProfile) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *CreateLogicalNetworkFromProfile) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *CreateLogicalNetworkFromProfile) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetLogicalNetworkProfileId

`func (o *CreateLogicalNetworkFromProfile) GetLogicalNetworkProfileId() int32`

GetLogicalNetworkProfileId returns the LogicalNetworkProfileId field if non-nil, zero value otherwise.

### GetLogicalNetworkProfileIdOk

`func (o *CreateLogicalNetworkFromProfile) GetLogicalNetworkProfileIdOk() (*int32, bool)`

GetLogicalNetworkProfileIdOk returns a tuple with the LogicalNetworkProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkProfileId

`func (o *CreateLogicalNetworkFromProfile) SetLogicalNetworkProfileId(v int32)`

SetLogicalNetworkProfileId sets LogicalNetworkProfileId field to given value.


### GetInfrastructureId

`func (o *CreateLogicalNetworkFromProfile) GetInfrastructureId() int32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *CreateLogicalNetworkFromProfile) GetInfrastructureIdOk() (*int32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *CreateLogicalNetworkFromProfile) SetInfrastructureId(v int32)`

SetInfrastructureId sets InfrastructureId field to given value.

### HasInfrastructureId

`func (o *CreateLogicalNetworkFromProfile) HasInfrastructureId() bool`

HasInfrastructureId returns a boolean if a field has been set.

### SetInfrastructureIdNil

`func (o *CreateLogicalNetworkFromProfile) SetInfrastructureIdNil(b bool)`

 SetInfrastructureIdNil sets the value for InfrastructureId to be an explicit nil

### UnsetInfrastructureId
`func (o *CreateLogicalNetworkFromProfile) UnsetInfrastructureId()`

UnsetInfrastructureId ensures that no value is present for InfrastructureId, not even an explicit nil
### GetMtu

`func (o *CreateLogicalNetworkFromProfile) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *CreateLogicalNetworkFromProfile) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *CreateLogicalNetworkFromProfile) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *CreateLogicalNetworkFromProfile) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *CreateLogicalNetworkFromProfile) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *CreateLogicalNetworkFromProfile) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
### GetExtensionInstanceId

`func (o *CreateLogicalNetworkFromProfile) GetExtensionInstanceId() int32`

GetExtensionInstanceId returns the ExtensionInstanceId field if non-nil, zero value otherwise.

### GetExtensionInstanceIdOk

`func (o *CreateLogicalNetworkFromProfile) GetExtensionInstanceIdOk() (*int32, bool)`

GetExtensionInstanceIdOk returns a tuple with the ExtensionInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionInstanceId

`func (o *CreateLogicalNetworkFromProfile) SetExtensionInstanceId(v int32)`

SetExtensionInstanceId sets ExtensionInstanceId field to given value.

### HasExtensionInstanceId

`func (o *CreateLogicalNetworkFromProfile) HasExtensionInstanceId() bool`

HasExtensionInstanceId returns a boolean if a field has been set.

### SetExtensionInstanceIdNil

`func (o *CreateLogicalNetworkFromProfile) SetExtensionInstanceIdNil(b bool)`

 SetExtensionInstanceIdNil sets the value for ExtensionInstanceId to be an explicit nil

### UnsetExtensionInstanceId
`func (o *CreateLogicalNetworkFromProfile) UnsetExtensionInstanceId()`

UnsetExtensionInstanceId ensures that no value is present for ExtensionInstanceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


