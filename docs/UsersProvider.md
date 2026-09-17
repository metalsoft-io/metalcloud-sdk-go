# UsersProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuiltIn** | **map[string]interface{}** | Whether built-in user authentication is enabled | 

## Methods

### NewUsersProvider

`func NewUsersProvider(builtIn map[string]interface{}, ) *UsersProvider`

NewUsersProvider instantiates a new UsersProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersProviderWithDefaults

`func NewUsersProviderWithDefaults() *UsersProvider`

NewUsersProviderWithDefaults instantiates a new UsersProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuiltIn

`func (o *UsersProvider) GetBuiltIn() map[string]interface{}`

GetBuiltIn returns the BuiltIn field if non-nil, zero value otherwise.

### GetBuiltInOk

`func (o *UsersProvider) GetBuiltInOk() (*map[string]interface{}, bool)`

GetBuiltInOk returns a tuple with the BuiltIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltIn

`func (o *UsersProvider) SetBuiltIn(v map[string]interface{})`

SetBuiltIn sets BuiltIn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


