# EditRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Role label | 
**Description** | Pointer to **string** | Role description | [optional] 
**Permissions** | **[]string** | List of permissions assigned to the role | 

## Methods

### NewEditRole

`func NewEditRole(label string, permissions []string, ) *EditRole`

NewEditRole instantiates a new EditRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditRoleWithDefaults

`func NewEditRoleWithDefaults() *EditRole`

NewEditRoleWithDefaults instantiates a new EditRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *EditRole) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *EditRole) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *EditRole) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *EditRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EditRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EditRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EditRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPermissions

`func (o *EditRole) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *EditRole) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *EditRole) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


