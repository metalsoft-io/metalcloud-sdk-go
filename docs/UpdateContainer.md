# UpdateContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comments** | Pointer to **float32** | Container comments. | [optional] 
**Tags** | Pointer to **[]string** | Tags for the Container. This is a JSON object. | [optional] 

## Methods

### NewUpdateContainer

`func NewUpdateContainer() *UpdateContainer`

NewUpdateContainer instantiates a new UpdateContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateContainerWithDefaults

`func NewUpdateContainerWithDefaults() *UpdateContainer`

NewUpdateContainerWithDefaults instantiates a new UpdateContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComments

`func (o *UpdateContainer) GetComments() float32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *UpdateContainer) GetCommentsOk() (*float32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *UpdateContainer) SetComments(v float32)`

SetComments sets Comments field to given value.

### HasComments

`func (o *UpdateContainer) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *UpdateContainer) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateContainer) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateContainer) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateContainer) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


