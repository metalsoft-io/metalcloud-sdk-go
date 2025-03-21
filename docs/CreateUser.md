# CreateUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | The display name of the user | 
**AccessLevel** | **string** | The access level of the user | 
**Email** | **string** | The email address of the user | 
**Password** | Pointer to **string** | The password of the user | [optional] 
**EmailVerified** | Pointer to **bool** | Whether the user has verified their email address | [optional] [default to false]
**CreateWithAccount** | Pointer to **bool** | Whether an account should be created with the user | [optional] [default to false]
**AccountId** | Pointer to **float32** | The account ID of the user | [optional] 

## Methods

### NewCreateUser

`func NewCreateUser(displayName string, accessLevel string, email string, ) *CreateUser`

NewCreateUser instantiates a new CreateUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserWithDefaults

`func NewCreateUserWithDefaults() *CreateUser`

NewCreateUserWithDefaults instantiates a new CreateUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *CreateUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetAccessLevel

`func (o *CreateUser) GetAccessLevel() string`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *CreateUser) GetAccessLevelOk() (*string, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *CreateUser) SetAccessLevel(v string)`

SetAccessLevel sets AccessLevel field to given value.


### GetEmail

`func (o *CreateUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *CreateUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateUser) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateUser) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetEmailVerified

`func (o *CreateUser) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *CreateUser) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *CreateUser) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.

### HasEmailVerified

`func (o *CreateUser) HasEmailVerified() bool`

HasEmailVerified returns a boolean if a field has been set.

### GetCreateWithAccount

`func (o *CreateUser) GetCreateWithAccount() bool`

GetCreateWithAccount returns the CreateWithAccount field if non-nil, zero value otherwise.

### GetCreateWithAccountOk

`func (o *CreateUser) GetCreateWithAccountOk() (*bool, bool)`

GetCreateWithAccountOk returns a tuple with the CreateWithAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateWithAccount

`func (o *CreateUser) SetCreateWithAccount(v bool)`

SetCreateWithAccount sets CreateWithAccount field to given value.

### HasCreateWithAccount

`func (o *CreateUser) HasCreateWithAccount() bool`

HasCreateWithAccount returns a boolean if a field has been set.

### GetAccountId

`func (o *CreateUser) GetAccountId() float32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateUser) GetAccountIdOk() (*float32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateUser) SetAccountId(v float32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CreateUser) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


