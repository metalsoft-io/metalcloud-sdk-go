# OidcProfileMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserExternalIdentifier** | **map[string]interface{}** | ID token claim used as the unique user identifier | 
**Username** | **map[string]interface{}** | ID token claim mapped to the username | 
**Email** | **map[string]interface{}** | ID token claim mapped to the email address | 
**Role** | Pointer to **string** | ID token claim mapped to the user role — fallback when no group mapping matches | [optional] 
**Group** | **map[string]interface{}** | ID token claim containing group memberships | 

## Methods

### NewOidcProfileMapping

`func NewOidcProfileMapping(userExternalIdentifier map[string]interface{}, username map[string]interface{}, email map[string]interface{}, group map[string]interface{}, ) *OidcProfileMapping`

NewOidcProfileMapping instantiates a new OidcProfileMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOidcProfileMappingWithDefaults

`func NewOidcProfileMappingWithDefaults() *OidcProfileMapping`

NewOidcProfileMappingWithDefaults instantiates a new OidcProfileMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserExternalIdentifier

`func (o *OidcProfileMapping) GetUserExternalIdentifier() map[string]interface{}`

GetUserExternalIdentifier returns the UserExternalIdentifier field if non-nil, zero value otherwise.

### GetUserExternalIdentifierOk

`func (o *OidcProfileMapping) GetUserExternalIdentifierOk() (*map[string]interface{}, bool)`

GetUserExternalIdentifierOk returns a tuple with the UserExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserExternalIdentifier

`func (o *OidcProfileMapping) SetUserExternalIdentifier(v map[string]interface{})`

SetUserExternalIdentifier sets UserExternalIdentifier field to given value.


### GetUsername

`func (o *OidcProfileMapping) GetUsername() map[string]interface{}`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *OidcProfileMapping) GetUsernameOk() (*map[string]interface{}, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *OidcProfileMapping) SetUsername(v map[string]interface{})`

SetUsername sets Username field to given value.


### GetEmail

`func (o *OidcProfileMapping) GetEmail() map[string]interface{}`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OidcProfileMapping) GetEmailOk() (*map[string]interface{}, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OidcProfileMapping) SetEmail(v map[string]interface{})`

SetEmail sets Email field to given value.


### GetRole

`func (o *OidcProfileMapping) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OidcProfileMapping) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OidcProfileMapping) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OidcProfileMapping) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetGroup

`func (o *OidcProfileMapping) GetGroup() map[string]interface{}`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *OidcProfileMapping) GetGroupOk() (*map[string]interface{}, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *OidcProfileMapping) SetGroup(v map[string]interface{})`

SetGroup sets Group field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


