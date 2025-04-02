# CreateLogicalNetwork

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

### NewCreateLogicalNetwork

`func NewCreateLogicalNetwork(fabricId float32, logicalNetworkType string, ) *CreateLogicalNetwork`

NewCreateLogicalNetwork instantiates a new CreateLogicalNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLogicalNetworkWithDefaults

`func NewCreateLogicalNetworkWithDefaults() *CreateLogicalNetwork`

NewCreateLogicalNetworkWithDefaults instantiates a new CreateLogicalNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CreateLogicalNetwork) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateLogicalNetwork) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateLogicalNetwork) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateLogicalNetwork) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *CreateLogicalNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLogicalNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLogicalNetwork) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateLogicalNetwork) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateLogicalNetwork) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateLogicalNetwork) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateLogicalNetwork) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateLogicalNetwork) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAnnotations

`func (o *CreateLogicalNetwork) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *CreateLogicalNetwork) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *CreateLogicalNetwork) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *CreateLogicalNetwork) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetFabricId

`func (o *CreateLogicalNetwork) GetFabricId() float32`

GetFabricId returns the FabricId field if non-nil, zero value otherwise.

### GetFabricIdOk

`func (o *CreateLogicalNetwork) GetFabricIdOk() (*float32, bool)`

GetFabricIdOk returns a tuple with the FabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricId

`func (o *CreateLogicalNetwork) SetFabricId(v float32)`

SetFabricId sets FabricId field to given value.


### GetInfrastructureId

`func (o *CreateLogicalNetwork) GetInfrastructureId() float32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *CreateLogicalNetwork) GetInfrastructureIdOk() (*float32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *CreateLogicalNetwork) SetInfrastructureId(v float32)`

SetInfrastructureId sets InfrastructureId field to given value.

### HasInfrastructureId

`func (o *CreateLogicalNetwork) HasInfrastructureId() bool`

HasInfrastructureId returns a boolean if a field has been set.

### GetLogicalNetworkType

`func (o *CreateLogicalNetwork) GetLogicalNetworkType() string`

GetLogicalNetworkType returns the LogicalNetworkType field if non-nil, zero value otherwise.

### GetLogicalNetworkTypeOk

`func (o *CreateLogicalNetwork) GetLogicalNetworkTypeOk() (*string, bool)`

GetLogicalNetworkTypeOk returns a tuple with the LogicalNetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkType

`func (o *CreateLogicalNetwork) SetLogicalNetworkType(v string)`

SetLogicalNetworkType sets LogicalNetworkType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


