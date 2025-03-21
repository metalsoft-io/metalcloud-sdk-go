# UserPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminPasswordRevealPermissions** | Pointer to [**AdminPasswordRevealPermissions**](AdminPasswordRevealPermissions.md) | Admin password reveal permissions | [optional] 
**SpecialPermissions** | Pointer to [**SpecialPermissions**](SpecialPermissions.md) | Special permissions | [optional] 
**RolePermissions** | **[]string** | Role permissions | 

## Methods

### NewUserPermissions

`func NewUserPermissions(rolePermissions []string, ) *UserPermissions`

NewUserPermissions instantiates a new UserPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPermissionsWithDefaults

`func NewUserPermissionsWithDefaults() *UserPermissions`

NewUserPermissionsWithDefaults instantiates a new UserPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminPasswordRevealPermissions

`func (o *UserPermissions) GetAdminPasswordRevealPermissions() AdminPasswordRevealPermissions`

GetAdminPasswordRevealPermissions returns the AdminPasswordRevealPermissions field if non-nil, zero value otherwise.

### GetAdminPasswordRevealPermissionsOk

`func (o *UserPermissions) GetAdminPasswordRevealPermissionsOk() (*AdminPasswordRevealPermissions, bool)`

GetAdminPasswordRevealPermissionsOk returns a tuple with the AdminPasswordRevealPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPasswordRevealPermissions

`func (o *UserPermissions) SetAdminPasswordRevealPermissions(v AdminPasswordRevealPermissions)`

SetAdminPasswordRevealPermissions sets AdminPasswordRevealPermissions field to given value.

### HasAdminPasswordRevealPermissions

`func (o *UserPermissions) HasAdminPasswordRevealPermissions() bool`

HasAdminPasswordRevealPermissions returns a boolean if a field has been set.

### GetSpecialPermissions

`func (o *UserPermissions) GetSpecialPermissions() SpecialPermissions`

GetSpecialPermissions returns the SpecialPermissions field if non-nil, zero value otherwise.

### GetSpecialPermissionsOk

`func (o *UserPermissions) GetSpecialPermissionsOk() (*SpecialPermissions, bool)`

GetSpecialPermissionsOk returns a tuple with the SpecialPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialPermissions

`func (o *UserPermissions) SetSpecialPermissions(v SpecialPermissions)`

SetSpecialPermissions sets SpecialPermissions field to given value.

### HasSpecialPermissions

`func (o *UserPermissions) HasSpecialPermissions() bool`

HasSpecialPermissions returns a boolean if a field has been set.

### GetRolePermissions

`func (o *UserPermissions) GetRolePermissions() []string`

GetRolePermissions returns the RolePermissions field if non-nil, zero value otherwise.

### GetRolePermissionsOk

`func (o *UserPermissions) GetRolePermissionsOk() (*[]string, bool)`

GetRolePermissionsOk returns a tuple with the RolePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolePermissions

`func (o *UserPermissions) SetRolePermissions(v []string)`

SetRolePermissions sets RolePermissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


