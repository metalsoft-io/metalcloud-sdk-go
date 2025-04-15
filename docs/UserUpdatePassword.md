# UserUpdatePassword

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewPassword** | **string** | The new password of the user | 
**OldPassword** | Pointer to **string** | The old password of the user | [optional] 

## Methods

### NewUserUpdatePassword

`func NewUserUpdatePassword(newPassword string, ) *UserUpdatePassword`

NewUserUpdatePassword instantiates a new UserUpdatePassword object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserUpdatePasswordWithDefaults

`func NewUserUpdatePasswordWithDefaults() *UserUpdatePassword`

NewUserUpdatePasswordWithDefaults instantiates a new UserUpdatePassword object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewPassword

`func (o *UserUpdatePassword) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *UserUpdatePassword) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *UserUpdatePassword) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.


### GetOldPassword

`func (o *UserUpdatePassword) GetOldPassword() string`

GetOldPassword returns the OldPassword field if non-nil, zero value otherwise.

### GetOldPasswordOk

`func (o *UserUpdatePassword) GetOldPasswordOk() (*string, bool)`

GetOldPasswordOk returns a tuple with the OldPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPassword

`func (o *UserUpdatePassword) SetOldPassword(v string)`

SetOldPassword sets OldPassword field to given value.

### HasOldPassword

`func (o *UserUpdatePassword) HasOldPassword() bool`

HasOldPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


