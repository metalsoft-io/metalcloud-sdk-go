# UpdateExternalSystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | The external system unique label | [optional] 
**Name** | Pointer to **string** | The external system name | [optional] 
**Annotations** | Pointer to **map[string]interface{}** | Annotations - key/value pairs of additional metadata | [optional] 

## Methods

### NewUpdateExternalSystem

`func NewUpdateExternalSystem() *UpdateExternalSystem`

NewUpdateExternalSystem instantiates a new UpdateExternalSystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateExternalSystemWithDefaults

`func NewUpdateExternalSystemWithDefaults() *UpdateExternalSystem`

NewUpdateExternalSystemWithDefaults instantiates a new UpdateExternalSystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *UpdateExternalSystem) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateExternalSystem) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateExternalSystem) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpdateExternalSystem) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *UpdateExternalSystem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateExternalSystem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateExternalSystem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateExternalSystem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAnnotations

`func (o *UpdateExternalSystem) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *UpdateExternalSystem) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *UpdateExternalSystem) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *UpdateExternalSystem) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


