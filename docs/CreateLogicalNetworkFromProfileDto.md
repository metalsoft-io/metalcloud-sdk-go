# CreateLogicalNetworkFromProfileDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**LogicalNetworkProfileId** | **int32** |  | 
**InfrastructureId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCreateLogicalNetworkFromProfileDto

`func NewCreateLogicalNetworkFromProfileDto(logicalNetworkProfileId int32, ) *CreateLogicalNetworkFromProfileDto`

NewCreateLogicalNetworkFromProfileDto instantiates a new CreateLogicalNetworkFromProfileDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLogicalNetworkFromProfileDtoWithDefaults

`func NewCreateLogicalNetworkFromProfileDtoWithDefaults() *CreateLogicalNetworkFromProfileDto`

NewCreateLogicalNetworkFromProfileDtoWithDefaults instantiates a new CreateLogicalNetworkFromProfileDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CreateLogicalNetworkFromProfileDto) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateLogicalNetworkFromProfileDto) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateLogicalNetworkFromProfileDto) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateLogicalNetworkFromProfileDto) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *CreateLogicalNetworkFromProfileDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLogicalNetworkFromProfileDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLogicalNetworkFromProfileDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateLogicalNetworkFromProfileDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAnnotations

`func (o *CreateLogicalNetworkFromProfileDto) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *CreateLogicalNetworkFromProfileDto) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *CreateLogicalNetworkFromProfileDto) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *CreateLogicalNetworkFromProfileDto) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetLogicalNetworkProfileId

`func (o *CreateLogicalNetworkFromProfileDto) GetLogicalNetworkProfileId() int32`

GetLogicalNetworkProfileId returns the LogicalNetworkProfileId field if non-nil, zero value otherwise.

### GetLogicalNetworkProfileIdOk

`func (o *CreateLogicalNetworkFromProfileDto) GetLogicalNetworkProfileIdOk() (*int32, bool)`

GetLogicalNetworkProfileIdOk returns a tuple with the LogicalNetworkProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkProfileId

`func (o *CreateLogicalNetworkFromProfileDto) SetLogicalNetworkProfileId(v int32)`

SetLogicalNetworkProfileId sets LogicalNetworkProfileId field to given value.


### GetInfrastructureId

`func (o *CreateLogicalNetworkFromProfileDto) GetInfrastructureId() int32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *CreateLogicalNetworkFromProfileDto) GetInfrastructureIdOk() (*int32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *CreateLogicalNetworkFromProfileDto) SetInfrastructureId(v int32)`

SetInfrastructureId sets InfrastructureId field to given value.

### HasInfrastructureId

`func (o *CreateLogicalNetworkFromProfileDto) HasInfrastructureId() bool`

HasInfrastructureId returns a boolean if a field has been set.

### SetInfrastructureIdNil

`func (o *CreateLogicalNetworkFromProfileDto) SetInfrastructureIdNil(b bool)`

 SetInfrastructureIdNil sets the value for InfrastructureId to be an explicit nil

### UnsetInfrastructureId
`func (o *CreateLogicalNetworkFromProfileDto) UnsetInfrastructureId()`

UnsetInfrastructureId ensures that no value is present for InfrastructureId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


