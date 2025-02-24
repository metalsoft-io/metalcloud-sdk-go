# AddUserToInfrastructure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserEmail** | **string** | The email of the user | 
**CreateIfNotExists** | **bool** | Create a new user if the user does not exist | [default to false]

## Methods

### NewAddUserToInfrastructure

`func NewAddUserToInfrastructure(userEmail string, createIfNotExists bool, ) *AddUserToInfrastructure`

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


