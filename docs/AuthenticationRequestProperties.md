# AuthenticationRequestProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email of the user | 
**Password** | **string** | The password of the user | 
**Data** | **map[string]interface{}** | The data for the SAML authentication request. | 

## Methods

### NewAuthenticationRequestProperties

`func NewAuthenticationRequestProperties(email string, password string, data map[string]interface{}, ) *AuthenticationRequestProperties`

NewAuthenticationRequestProperties instantiates a new AuthenticationRequestProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationRequestPropertiesWithDefaults

`func NewAuthenticationRequestPropertiesWithDefaults() *AuthenticationRequestProperties`

NewAuthenticationRequestPropertiesWithDefaults instantiates a new AuthenticationRequestProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *AuthenticationRequestProperties) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AuthenticationRequestProperties) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AuthenticationRequestProperties) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *AuthenticationRequestProperties) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AuthenticationRequestProperties) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AuthenticationRequestProperties) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetData

`func (o *AuthenticationRequestProperties) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AuthenticationRequestProperties) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AuthenticationRequestProperties) SetData(v map[string]interface{})`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


