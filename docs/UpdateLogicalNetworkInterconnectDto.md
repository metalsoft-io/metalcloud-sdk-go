# UpdateLogicalNetworkInterconnectDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Unique label for the logical network interconnect | [optional] 
**Name** | Pointer to **string** | Name of the logical network interconnect | [optional] 
**Annotations** | Pointer to **map[string]string** | JSON object containing additional metadata or annotations | [optional] 
**Kind** | Pointer to [**LogicalNetworkInterconnectKind**](LogicalNetworkInterconnectKind.md) | Kind of the logical network interconnect | [optional] 
**FabricInterconnectId** | Pointer to **int32** | Transport ID allocated from range 999999999-900000000 (descending) | [optional] 

## Methods

### NewUpdateLogicalNetworkInterconnectDto

`func NewUpdateLogicalNetworkInterconnectDto() *UpdateLogicalNetworkInterconnectDto`

NewUpdateLogicalNetworkInterconnectDto instantiates a new UpdateLogicalNetworkInterconnectDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLogicalNetworkInterconnectDtoWithDefaults

`func NewUpdateLogicalNetworkInterconnectDtoWithDefaults() *UpdateLogicalNetworkInterconnectDto`

NewUpdateLogicalNetworkInterconnectDtoWithDefaults instantiates a new UpdateLogicalNetworkInterconnectDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *UpdateLogicalNetworkInterconnectDto) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateLogicalNetworkInterconnectDto) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateLogicalNetworkInterconnectDto) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpdateLogicalNetworkInterconnectDto) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *UpdateLogicalNetworkInterconnectDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateLogicalNetworkInterconnectDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateLogicalNetworkInterconnectDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateLogicalNetworkInterconnectDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAnnotations

`func (o *UpdateLogicalNetworkInterconnectDto) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *UpdateLogicalNetworkInterconnectDto) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *UpdateLogicalNetworkInterconnectDto) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *UpdateLogicalNetworkInterconnectDto) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetKind

`func (o *UpdateLogicalNetworkInterconnectDto) GetKind() LogicalNetworkInterconnectKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *UpdateLogicalNetworkInterconnectDto) GetKindOk() (*LogicalNetworkInterconnectKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *UpdateLogicalNetworkInterconnectDto) SetKind(v LogicalNetworkInterconnectKind)`

SetKind sets Kind field to given value.

### HasKind

`func (o *UpdateLogicalNetworkInterconnectDto) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetFabricInterconnectId

`func (o *UpdateLogicalNetworkInterconnectDto) GetFabricInterconnectId() int32`

GetFabricInterconnectId returns the FabricInterconnectId field if non-nil, zero value otherwise.

### GetFabricInterconnectIdOk

`func (o *UpdateLogicalNetworkInterconnectDto) GetFabricInterconnectIdOk() (*int32, bool)`

GetFabricInterconnectIdOk returns a tuple with the FabricInterconnectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricInterconnectId

`func (o *UpdateLogicalNetworkInterconnectDto) SetFabricInterconnectId(v int32)`

SetFabricInterconnectId sets FabricInterconnectId field to given value.

### HasFabricInterconnectId

`func (o *UpdateLogicalNetworkInterconnectDto) HasFabricInterconnectId() bool`

HasFabricInterconnectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


