# UsersProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mysql** | **map[string]interface{}** | Whether MySQL/local user authentication is enabled | 

## Methods

### NewUsersProvider

`func NewUsersProvider(mysql map[string]interface{}, ) *UsersProvider`

NewUsersProvider instantiates a new UsersProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersProviderWithDefaults

`func NewUsersProviderWithDefaults() *UsersProvider`

NewUsersProviderWithDefaults instantiates a new UsersProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMysql

`func (o *UsersProvider) GetMysql() map[string]interface{}`

GetMysql returns the Mysql field if non-nil, zero value otherwise.

### GetMysqlOk

`func (o *UsersProvider) GetMysqlOk() (*map[string]interface{}, bool)`

GetMysqlOk returns a tuple with the Mysql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysql

`func (o *UsersProvider) SetMysql(v map[string]interface{})`

SetMysql sets Mysql field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


