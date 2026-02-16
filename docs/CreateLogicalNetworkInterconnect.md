# CreateLogicalNetworkInterconnect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Unique label for the logical network interconnect | 
**Name** | **string** | Name of the logical network interconnect | 
**Annotations** | Pointer to **map[string]string** | JSON object containing additional metadata or annotations | [optional] 
**Kind** | Pointer to [**LogicalNetworkInterconnectKind**](LogicalNetworkInterconnectKind.md) | Kind of the logical network interconnect | [optional] 
**FabricInterconnectId** | **int32** | Fabric Interconnect identifier | 

## Methods

### NewCreateLogicalNetworkInterconnect

`func NewCreateLogicalNetworkInterconnect(label string, name string, fabricInterconnectId int32, ) *CreateLogicalNetworkInterconnect`

NewCreateLogicalNetworkInterconnect instantiates a new CreateLogicalNetworkInterconnect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLogicalNetworkInterconnectWithDefaults

`func NewCreateLogicalNetworkInterconnectWithDefaults() *CreateLogicalNetworkInterconnect`

NewCreateLogicalNetworkInterconnectWithDefaults instantiates a new CreateLogicalNetworkInterconnect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CreateLogicalNetworkInterconnect) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateLogicalNetworkInterconnect) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateLogicalNetworkInterconnect) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *CreateLogicalNetworkInterconnect) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLogicalNetworkInterconnect) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLogicalNetworkInterconnect) SetName(v string)`

SetName sets Name field to given value.


### GetAnnotations

`func (o *CreateLogicalNetworkInterconnect) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *CreateLogicalNetworkInterconnect) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *CreateLogicalNetworkInterconnect) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *CreateLogicalNetworkInterconnect) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetKind

`func (o *CreateLogicalNetworkInterconnect) GetKind() LogicalNetworkInterconnectKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateLogicalNetworkInterconnect) GetKindOk() (*LogicalNetworkInterconnectKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateLogicalNetworkInterconnect) SetKind(v LogicalNetworkInterconnectKind)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CreateLogicalNetworkInterconnect) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetFabricInterconnectId

`func (o *CreateLogicalNetworkInterconnect) GetFabricInterconnectId() int32`

GetFabricInterconnectId returns the FabricInterconnectId field if non-nil, zero value otherwise.

### GetFabricInterconnectIdOk

`func (o *CreateLogicalNetworkInterconnect) GetFabricInterconnectIdOk() (*int32, bool)`

GetFabricInterconnectIdOk returns a tuple with the FabricInterconnectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricInterconnectId

`func (o *CreateLogicalNetworkInterconnect) SetFabricInterconnectId(v int32)`

SetFabricInterconnectId sets FabricInterconnectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


