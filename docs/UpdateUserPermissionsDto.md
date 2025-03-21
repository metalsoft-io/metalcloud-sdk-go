# UpdateUserPermissionsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminPasswordRevealPermissions** | Pointer to [**AdminPasswordRevealPermissions**](AdminPasswordRevealPermissions.md) | Admin password reveal permissions | [optional] 
**SpecialPermissions** | Pointer to [**SpecialPermissions**](SpecialPermissions.md) | Special permissions | [optional] 

## Methods

### NewUpdateUserPermissionsDto

`func NewUpdateUserPermissionsDto() *UpdateUserPermissionsDto`

NewUpdateUserPermissionsDto instantiates a new UpdateUserPermissionsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserPermissionsDtoWithDefaults

`func NewUpdateUserPermissionsDtoWithDefaults() *UpdateUserPermissionsDto`

NewUpdateUserPermissionsDtoWithDefaults instantiates a new UpdateUserPermissionsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminPasswordRevealPermissions

`func (o *UpdateUserPermissionsDto) GetAdminPasswordRevealPermissions() AdminPasswordRevealPermissions`

GetAdminPasswordRevealPermissions returns the AdminPasswordRevealPermissions field if non-nil, zero value otherwise.

### GetAdminPasswordRevealPermissionsOk

`func (o *UpdateUserPermissionsDto) GetAdminPasswordRevealPermissionsOk() (*AdminPasswordRevealPermissions, bool)`

GetAdminPasswordRevealPermissionsOk returns a tuple with the AdminPasswordRevealPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPasswordRevealPermissions

`func (o *UpdateUserPermissionsDto) SetAdminPasswordRevealPermissions(v AdminPasswordRevealPermissions)`

SetAdminPasswordRevealPermissions sets AdminPasswordRevealPermissions field to given value.

### HasAdminPasswordRevealPermissions

`func (o *UpdateUserPermissionsDto) HasAdminPasswordRevealPermissions() bool`

HasAdminPasswordRevealPermissions returns a boolean if a field has been set.

### GetSpecialPermissions

`func (o *UpdateUserPermissionsDto) GetSpecialPermissions() SpecialPermissions`

GetSpecialPermissions returns the SpecialPermissions field if non-nil, zero value otherwise.

### GetSpecialPermissionsOk

`func (o *UpdateUserPermissionsDto) GetSpecialPermissionsOk() (*SpecialPermissions, bool)`

GetSpecialPermissionsOk returns a tuple with the SpecialPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialPermissions

`func (o *UpdateUserPermissionsDto) SetSpecialPermissions(v SpecialPermissions)`

SetSpecialPermissions sets SpecialPermissions field to given value.

### HasSpecialPermissions

`func (o *UpdateUserPermissionsDto) HasSpecialPermissions() bool`

HasSpecialPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


