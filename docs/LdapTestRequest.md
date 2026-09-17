# LdapTestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | Username to authenticate against the configured LDAP directory | 
**Password** | **string** | Password for the supplied username | 

## Methods

### NewLdapTestRequest

`func NewLdapTestRequest(username string, password string, ) *LdapTestRequest`

NewLdapTestRequest instantiates a new LdapTestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapTestRequestWithDefaults

`func NewLdapTestRequestWithDefaults() *LdapTestRequest`

NewLdapTestRequestWithDefaults instantiates a new LdapTestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *LdapTestRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *LdapTestRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *LdapTestRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *LdapTestRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LdapTestRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LdapTestRequest) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


