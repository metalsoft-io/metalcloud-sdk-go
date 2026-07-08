# AddUserToInfrastructure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserEmail** | Pointer to **string** | The email of the user. Exactly one of userEmail or userId must be provided. | [optional] 
**UserId** | Pointer to **int64** | The id of the user. Exactly one of userEmail or userId must be provided. Use userId to reference users without a valid email address (e.g. LDAP users). | [optional] 
**CreateIfNotExists** | **bool** | Create a new user if the user does not exist. Only allowed when userEmail is provided. | [default to false]

## Methods

### NewAddUserToInfrastructure

`func NewAddUserToInfrastructure(createIfNotExists bool, ) *AddUserToInfrastructure`

NewAddUserToInfrastructure instantiates a new AddUserToInfrastructure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUserToInfrastructureWithDefaults

`func NewAddUserToInfrastructureWithDefaults() *AddUserToInfrastructure`

NewAddUserToInfrastructureWithDefaults instantiates a new AddUserToInfrastructure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserEmail

`func (o *AddUserToInfrastructure) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *AddUserToInfrastructure) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *AddUserToInfrastructure) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *AddUserToInfrastructure) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### GetUserId

`func (o *AddUserToInfrastructure) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AddUserToInfrastructure) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AddUserToInfrastructure) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AddUserToInfrastructure) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCreateIfNotExists

`func (o *AddUserToInfrastructure) GetCreateIfNotExists() bool`

GetCreateIfNotExists returns the CreateIfNotExists field if non-nil, zero value otherwise.

### GetCreateIfNotExistsOk

`func (o *AddUserToInfrastructure) GetCreateIfNotExistsOk() (*bool, bool)`

GetCreateIfNotExistsOk returns a tuple with the CreateIfNotExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIfNotExists

`func (o *AddUserToInfrastructure) SetCreateIfNotExists(v bool)`

SetCreateIfNotExists sets CreateIfNotExists field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


