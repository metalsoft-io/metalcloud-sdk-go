# UserResourcePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PasswordRevealPermissions** | Pointer to [**[]UserPasswordRevealPermission**](UserPasswordRevealPermission.md) | The new password reveal permissions of the user. | [optional] 
**SpecialPermissions** | Pointer to [**[]UserSpecialPermission**](UserSpecialPermission.md) | The new special permissions of the user. | [optional] 

## Methods

### NewUserResourcePermissions

`func NewUserResourcePermissions() *UserResourcePermissions`

NewUserResourcePermissions instantiates a new UserResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResourcePermissionsWithDefaults

`func NewUserResourcePermissionsWithDefaults() *UserResourcePermissions`

NewUserResourcePermissionsWithDefaults instantiates a new UserResourcePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPasswordRevealPermissions

`func (o *UserResourcePermissions) GetPasswordRevealPermissions() []UserPasswordRevealPermission`

GetPasswordRevealPermissions returns the PasswordRevealPermissions field if non-nil, zero value otherwise.

### GetPasswordRevealPermissionsOk

`func (o *UserResourcePermissions) GetPasswordRevealPermissionsOk() (*[]UserPasswordRevealPermission, bool)`

GetPasswordRevealPermissionsOk returns a tuple with the PasswordRevealPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordRevealPermissions

`func (o *UserResourcePermissions) SetPasswordRevealPermissions(v []UserPasswordRevealPermission)`

SetPasswordRevealPermissions sets PasswordRevealPermissions field to given value.

### HasPasswordRevealPermissions

`func (o *UserResourcePermissions) HasPasswordRevealPermissions() bool`

HasPasswordRevealPermissions returns a boolean if a field has been set.

### GetSpecialPermissions

`func (o *UserResourcePermissions) GetSpecialPermissions() []UserSpecialPermission`

GetSpecialPermissions returns the SpecialPermissions field if non-nil, zero value otherwise.

### GetSpecialPermissionsOk

`func (o *UserResourcePermissions) GetSpecialPermissionsOk() (*[]UserSpecialPermission, bool)`

GetSpecialPermissionsOk returns a tuple with the SpecialPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialPermissions

`func (o *UserResourcePermissions) SetSpecialPermissions(v []UserSpecialPermission)`

SetSpecialPermissions sets SpecialPermissions field to given value.

### HasSpecialPermissions

`func (o *UserResourcePermissions) HasSpecialPermissions() bool`

HasSpecialPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


