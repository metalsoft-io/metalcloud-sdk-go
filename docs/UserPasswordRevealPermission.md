# UserPasswordRevealPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PasswordRevealPermission** | **string** | The password reveal permission for the user | 
**EnablePermission** | Pointer to **bool** | Whether to enable the permission | [optional] 

## Methods

### NewUserPasswordRevealPermission

`func NewUserPasswordRevealPermission(passwordRevealPermission string, ) *UserPasswordRevealPermission`

NewUserPasswordRevealPermission instantiates a new UserPasswordRevealPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPasswordRevealPermissionWithDefaults

`func NewUserPasswordRevealPermissionWithDefaults() *UserPasswordRevealPermission`

NewUserPasswordRevealPermissionWithDefaults instantiates a new UserPasswordRevealPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPasswordRevealPermission

`func (o *UserPasswordRevealPermission) GetPasswordRevealPermission() string`

GetPasswordRevealPermission returns the PasswordRevealPermission field if non-nil, zero value otherwise.

### GetPasswordRevealPermissionOk

`func (o *UserPasswordRevealPermission) GetPasswordRevealPermissionOk() (*string, bool)`

GetPasswordRevealPermissionOk returns a tuple with the PasswordRevealPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordRevealPermission

`func (o *UserPasswordRevealPermission) SetPasswordRevealPermission(v string)`

SetPasswordRevealPermission sets PasswordRevealPermission field to given value.


### GetEnablePermission

`func (o *UserPasswordRevealPermission) GetEnablePermission() bool`

GetEnablePermission returns the EnablePermission field if non-nil, zero value otherwise.

### GetEnablePermissionOk

`func (o *UserPasswordRevealPermission) GetEnablePermissionOk() (*bool, bool)`

GetEnablePermissionOk returns a tuple with the EnablePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePermission

`func (o *UserPasswordRevealPermission) SetEnablePermission(v bool)`

SetEnablePermission sets EnablePermission field to given value.

### HasEnablePermission

`func (o *UserPasswordRevealPermission) HasEnablePermission() bool`

HasEnablePermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


