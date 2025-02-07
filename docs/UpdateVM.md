# UpdateVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comments** | Pointer to **float32** | VM comments. | [optional] 
**Tags** | Pointer to **[]string** | Tags for the VM. This is a JSON object. | [optional] 

## Methods

### NewUpdateVM

`func NewUpdateVM() *UpdateVM`

NewUpdateVM instantiates a new UpdateVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVMWithDefaults

`func NewUpdateVMWithDefaults() *UpdateVM`

NewUpdateVMWithDefaults instantiates a new UpdateVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComments

`func (o *UpdateVM) GetComments() float32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *UpdateVM) GetCommentsOk() (*float32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *UpdateVM) SetComments(v float32)`

SetComments sets Comments field to given value.

### HasComments

`func (o *UpdateVM) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *UpdateVM) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateVM) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateVM) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateVM) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


