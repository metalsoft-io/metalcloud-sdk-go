# StorageScopedAccessUserCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | Username of the storage scoped access user | 
**Password** | Pointer to **string** | Password of the storage scoped access user | [optional] 
**ApiToken** | Pointer to **string** | API token of the storage scoped access user | [optional] 

## Methods

### NewStorageScopedAccessUserCredentials

`func NewStorageScopedAccessUserCredentials(username string, ) *StorageScopedAccessUserCredentials`

NewStorageScopedAccessUserCredentials instantiates a new StorageScopedAccessUserCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageScopedAccessUserCredentialsWithDefaults

`func NewStorageScopedAccessUserCredentialsWithDefaults() *StorageScopedAccessUserCredentials`

NewStorageScopedAccessUserCredentialsWithDefaults instantiates a new StorageScopedAccessUserCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *StorageScopedAccessUserCredentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *StorageScopedAccessUserCredentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *StorageScopedAccessUserCredentials) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *StorageScopedAccessUserCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *StorageScopedAccessUserCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *StorageScopedAccessUserCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *StorageScopedAccessUserCredentials) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetApiToken

`func (o *StorageScopedAccessUserCredentials) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *StorageScopedAccessUserCredentials) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *StorageScopedAccessUserCredentials) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.

### HasApiToken

`func (o *StorageScopedAccessUserCredentials) HasApiToken() bool`

HasApiToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


