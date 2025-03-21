# CreateLogicalNetworkProfileDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Label for the logical network profile | [optional] 
**Name** | Pointer to **string** | Name of the logical network profile | [optional] 
**Description** | Pointer to **string** | Description of the logical network profile | [optional] 
**Annotations** | Pointer to **map[string]interface{}** | Annotations for the logical network profile | [optional] 
**LogicalNetworkType** | **string** | Type of the logical network profile | 

## Methods

### NewCreateLogicalNetworkProfileDto

`func NewCreateLogicalNetworkProfileDto(logicalNetworkType string, ) *CreateLogicalNetworkProfileDto`

NewCreateLogicalNetworkProfileDto instantiates a new CreateLogicalNetworkProfileDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLogicalNetworkProfileDtoWithDefaults

`func NewCreateLogicalNetworkProfileDtoWithDefaults() *CreateLogicalNetworkProfileDto`

NewCreateLogicalNetworkProfileDtoWithDefaults instantiates a new CreateLogicalNetworkProfileDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CreateLogicalNetworkProfileDto) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateLogicalNetworkProfileDto) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateLogicalNetworkProfileDto) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateLogicalNetworkProfileDto) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *CreateLogicalNetworkProfileDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLogicalNetworkProfileDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLogicalNetworkProfileDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateLogicalNetworkProfileDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateLogicalNetworkProfileDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateLogicalNetworkProfileDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateLogicalNetworkProfileDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateLogicalNetworkProfileDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAnnotations

`func (o *CreateLogicalNetworkProfileDto) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *CreateLogicalNetworkProfileDto) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *CreateLogicalNetworkProfileDto) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *CreateLogicalNetworkProfileDto) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetLogicalNetworkType

`func (o *CreateLogicalNetworkProfileDto) GetLogicalNetworkType() string`

GetLogicalNetworkType returns the LogicalNetworkType field if non-nil, zero value otherwise.

### GetLogicalNetworkTypeOk

`func (o *CreateLogicalNetworkProfileDto) GetLogicalNetworkTypeOk() (*string, bool)`

GetLogicalNetworkTypeOk returns a tuple with the LogicalNetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkType

`func (o *CreateLogicalNetworkProfileDto) SetLogicalNetworkType(v string)`

SetLogicalNetworkType sets LogicalNetworkType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


