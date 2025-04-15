# CreateRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Role label | 
**Description** | Pointer to **string** | Role description | [optional] 

## Methods

### NewCreateRole

`func NewCreateRole(label string, ) *CreateRole`

NewCreateRole instantiates a new CreateRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRoleWithDefaults

`func NewCreateRoleWithDefaults() *CreateRole`

NewCreateRoleWithDefaults instantiates a new CreateRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CreateRole) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateRole) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateRole) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *CreateRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


