# CreateLogicalNetworkDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Label for the logical network | [optional] 
**Name** | Pointer to **string** | Name of the logical network | [optional] 
**Description** | Pointer to **string** | Description of the logical network | [optional] 
**Annotations** | Pointer to **map[string]interface{}** | Annotations for the logical network | [optional] 
**FabricId** | **float32** | Fabric ID associated with the logical network | 
**InfrastructureId** | Pointer to **float32** | Infrastructure ID associated with the logical network | [optional] 
**LogicalNetworkType** | **string** | Type of the logical network | 

## Methods

### NewCreateLogicalNetworkDto

`func NewCreateLogicalNetworkDto(fabricId float32, logicalNetworkType string, ) *CreateLogicalNetworkDto`

NewCreateLogicalNetworkDto instantiates a new CreateLogicalNetworkDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLogicalNetworkDtoWithDefaults

`func NewCreateLogicalNetworkDtoWithDefaults() *CreateLogicalNetworkDto`

NewCreateLogicalNetworkDtoWithDefaults instantiates a new CreateLogicalNetworkDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CreateLogicalNetworkDto) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateLogicalNetworkDto) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateLogicalNetworkDto) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateLogicalNetworkDto) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *CreateLogicalNetworkDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLogicalNetworkDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLogicalNetworkDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateLogicalNetworkDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateLogicalNetworkDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateLogicalNetworkDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateLogicalNetworkDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateLogicalNetworkDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAnnotations

`func (o *CreateLogicalNetworkDto) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *CreateLogicalNetworkDto) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *CreateLogicalNetworkDto) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *CreateLogicalNetworkDto) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetFabricId

`func (o *CreateLogicalNetworkDto) GetFabricId() float32`

GetFabricId returns the FabricId field if non-nil, zero value otherwise.

### GetFabricIdOk

`func (o *CreateLogicalNetworkDto) GetFabricIdOk() (*float32, bool)`

GetFabricIdOk returns a tuple with the FabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricId

`func (o *CreateLogicalNetworkDto) SetFabricId(v float32)`

SetFabricId sets FabricId field to given value.


### GetInfrastructureId

`func (o *CreateLogicalNetworkDto) GetInfrastructureId() float32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *CreateLogicalNetworkDto) GetInfrastructureIdOk() (*float32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *CreateLogicalNetworkDto) SetInfrastructureId(v float32)`

SetInfrastructureId sets InfrastructureId field to given value.

### HasInfrastructureId

`func (o *CreateLogicalNetworkDto) HasInfrastructureId() bool`

HasInfrastructureId returns a boolean if a field has been set.

### GetLogicalNetworkType

`func (o *CreateLogicalNetworkDto) GetLogicalNetworkType() string`

GetLogicalNetworkType returns the LogicalNetworkType field if non-nil, zero value otherwise.

### GetLogicalNetworkTypeOk

`func (o *CreateLogicalNetworkDto) GetLogicalNetworkTypeOk() (*string, bool)`

GetLogicalNetworkTypeOk returns a tuple with the LogicalNetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkType

`func (o *CreateLogicalNetworkDto) SetLogicalNetworkType(v string)`

SetLogicalNetworkType sets LogicalNetworkType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


