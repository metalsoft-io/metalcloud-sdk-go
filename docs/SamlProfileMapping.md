# SamlProfileMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserExternalIdentifier** | **map[string]interface{}** | SAML assertion attribute used as the unique user identifier | 
**Username** | **map[string]interface{}** | SAML assertion attribute mapped to the username | 
**Email** | **map[string]interface{}** | SAML assertion attribute mapped to the email address | 
**Role** | **map[string]interface{}** | SAML assertion attribute mapped to the user role | 
**Group** | **map[string]interface{}** | SAML assertion attribute containing group memberships | 

## Methods

### NewSamlProfileMapping

`func NewSamlProfileMapping(userExternalIdentifier map[string]interface{}, username map[string]interface{}, email map[string]interface{}, role map[string]interface{}, group map[string]interface{}, ) *SamlProfileMapping`

NewSamlProfileMapping instantiates a new SamlProfileMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlProfileMappingWithDefaults

`func NewSamlProfileMappingWithDefaults() *SamlProfileMapping`

NewSamlProfileMappingWithDefaults instantiates a new SamlProfileMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserExternalIdentifier

`func (o *SamlProfileMapping) GetUserExternalIdentifier() map[string]interface{}`

GetUserExternalIdentifier returns the UserExternalIdentifier field if non-nil, zero value otherwise.

### GetUserExternalIdentifierOk

`func (o *SamlProfileMapping) GetUserExternalIdentifierOk() (*map[string]interface{}, bool)`

GetUserExternalIdentifierOk returns a tuple with the UserExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserExternalIdentifier

`func (o *SamlProfileMapping) SetUserExternalIdentifier(v map[string]interface{})`

SetUserExternalIdentifier sets UserExternalIdentifier field to given value.


### GetUsername

`func (o *SamlProfileMapping) GetUsername() map[string]interface{}`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SamlProfileMapping) GetUsernameOk() (*map[string]interface{}, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SamlProfileMapping) SetUsername(v map[string]interface{})`

SetUsername sets Username field to given value.


### GetEmail

`func (o *SamlProfileMapping) GetEmail() map[string]interface{}`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SamlProfileMapping) GetEmailOk() (*map[string]interface{}, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SamlProfileMapping) SetEmail(v map[string]interface{})`

SetEmail sets Email field to given value.


### GetRole

`func (o *SamlProfileMapping) GetRole() map[string]interface{}`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SamlProfileMapping) GetRoleOk() (*map[string]interface{}, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SamlProfileMapping) SetRole(v map[string]interface{})`

SetRole sets Role field to given value.


### GetGroup

`func (o *SamlProfileMapping) GetGroup() map[string]interface{}`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *SamlProfileMapping) GetGroupOk() (*map[string]interface{}, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *SamlProfileMapping) SetGroup(v map[string]interface{})`

SetGroup sets Group field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


