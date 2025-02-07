# UserDelegate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The user delegate e-mail | [optional] 
**CreateIfNotExists** | Pointer to **bool** | Whether to create the delegate if it does not exist. Can only be used when creating a new delegate. | [optional] [default to false]
**RemoveDelegate** | Pointer to **bool** | Whether to remove the delegate. Cannot be used when creating a new delegate. | [optional] [default to false]

## Methods

### NewUserDelegate

`func NewUserDelegate() *UserDelegate`

NewUserDelegate instantiates a new UserDelegate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDelegateWithDefaults

`func NewUserDelegateWithDefaults() *UserDelegate`

NewUserDelegateWithDefaults instantiates a new UserDelegate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserDelegate) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserDelegate) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserDelegate) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserDelegate) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCreateIfNotExists

`func (o *UserDelegate) GetCreateIfNotExists() bool`

GetCreateIfNotExists returns the CreateIfNotExists field if non-nil, zero value otherwise.

### GetCreateIfNotExistsOk

`func (o *UserDelegate) GetCreateIfNotExistsOk() (*bool, bool)`

GetCreateIfNotExistsOk returns a tuple with the CreateIfNotExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIfNotExists

`func (o *UserDelegate) SetCreateIfNotExists(v bool)`

SetCreateIfNotExists sets CreateIfNotExists field to given value.

### HasCreateIfNotExists

`func (o *UserDelegate) HasCreateIfNotExists() bool`

HasCreateIfNotExists returns a boolean if a field has been set.

### GetRemoveDelegate

`func (o *UserDelegate) GetRemoveDelegate() bool`

GetRemoveDelegate returns the RemoveDelegate field if non-nil, zero value otherwise.

### GetRemoveDelegateOk

`func (o *UserDelegate) GetRemoveDelegateOk() (*bool, bool)`

GetRemoveDelegateOk returns a tuple with the RemoveDelegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveDelegate

`func (o *UserDelegate) SetRemoveDelegate(v bool)`

SetRemoveDelegate sets RemoveDelegate field to given value.

### HasRemoveDelegate

`func (o *UserDelegate) HasRemoveDelegate() bool`

HasRemoveDelegate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


