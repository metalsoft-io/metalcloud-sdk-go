# LogicalNetworkDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID of the logical network | 
**Revision** | **float32** | Revision number of the logical network | 
**Label** | Pointer to **string** | Label for the logical network | [optional] 
**Name** | Pointer to **string** | Name of the logical network | [optional] 
**Description** | Pointer to **string** | Description of the logical network | [optional] 
**Annotations** | Pointer to **map[string]interface{}** | Annotations for the logical network | [optional] 
**FabricId** | **float32** | Fabric ID associated with the logical network | 
**InfrastructureId** | Pointer to **float32** | Infrastructure ID associated with the logical network | [optional] 
**LogicalNetworkType** | **string** | Type of the logical network | 

## Methods

### NewLogicalNetworkDto

`func NewLogicalNetworkDto(id float32, revision float32, fabricId float32, logicalNetworkType string, ) *LogicalNetworkDto`

NewLogicalNetworkDto instantiates a new LogicalNetworkDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogicalNetworkDtoWithDefaults

`func NewLogicalNetworkDtoWithDefaults() *LogicalNetworkDto`

NewLogicalNetworkDtoWithDefaults instantiates a new LogicalNetworkDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LogicalNetworkDto) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogicalNetworkDto) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogicalNetworkDto) SetId(v float32)`

SetId sets Id field to given value.


### GetRevision

`func (o *LogicalNetworkDto) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *LogicalNetworkDto) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *LogicalNetworkDto) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetLabel

`func (o *LogicalNetworkDto) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LogicalNetworkDto) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LogicalNetworkDto) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *LogicalNetworkDto) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *LogicalNetworkDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogicalNetworkDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogicalNetworkDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LogicalNetworkDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *LogicalNetworkDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LogicalNetworkDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LogicalNetworkDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LogicalNetworkDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAnnotations

`func (o *LogicalNetworkDto) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *LogicalNetworkDto) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *LogicalNetworkDto) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *LogicalNetworkDto) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetFabricId

`func (o *LogicalNetworkDto) GetFabricId() float32`

GetFabricId returns the FabricId field if non-nil, zero value otherwise.

### GetFabricIdOk

`func (o *LogicalNetworkDto) GetFabricIdOk() (*float32, bool)`

GetFabricIdOk returns a tuple with the FabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricId

`func (o *LogicalNetworkDto) SetFabricId(v float32)`

SetFabricId sets FabricId field to given value.


### GetInfrastructureId

`func (o *LogicalNetworkDto) GetInfrastructureId() float32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *LogicalNetworkDto) GetInfrastructureIdOk() (*float32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *LogicalNetworkDto) SetInfrastructureId(v float32)`

SetInfrastructureId sets InfrastructureId field to given value.

### HasInfrastructureId

`func (o *LogicalNetworkDto) HasInfrastructureId() bool`

HasInfrastructureId returns a boolean if a field has been set.

### GetLogicalNetworkType

`func (o *LogicalNetworkDto) GetLogicalNetworkType() string`

GetLogicalNetworkType returns the LogicalNetworkType field if non-nil, zero value otherwise.

### GetLogicalNetworkTypeOk

`func (o *LogicalNetworkDto) GetLogicalNetworkTypeOk() (*string, bool)`

GetLogicalNetworkTypeOk returns a tuple with the LogicalNetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkType

`func (o *LogicalNetworkDto) SetLogicalNetworkType(v string)`

SetLogicalNetworkType sets LogicalNetworkType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


