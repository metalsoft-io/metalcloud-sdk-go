# UserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | The display name of the user | 
**EmailStatus** | **string** | The email status of the user | 
**Language** | **string** | The language of the user | 
**LastLoginTimestamp** | **string** | The timestamp when the user logged in last | 
**AccessLevel** | **string** | The access level of the user | 
**IsArchived** | **bool** | Whether the user is a archived | 
**AccountId** | Pointer to **float32** | The account ID of the user | [optional] 
**Id** | **float32** | User ID | 
**Revision** | **float32** | Revision of the user | 
**Email** | **string** | The email address of the user | 
**CreatedTimestamp** | **string** | The timestamp when the user was created | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewUserInfo

`func NewUserInfo(displayName string, emailStatus string, language string, lastLoginTimestamp string, accessLevel string, isArchived bool, id float32, revision float32, email string, createdTimestamp string, ) *UserInfo`

NewUserInfo instantiates a new UserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInfoWithDefaults

`func NewUserInfoWithDefaults() *UserInfo`

NewUserInfoWithDefaults instantiates a new UserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *UserInfo) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserInfo) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserInfo) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetEmailStatus

`func (o *UserInfo) GetEmailStatus() string`

GetEmailStatus returns the EmailStatus field if non-nil, zero value otherwise.

### GetEmailStatusOk

`func (o *UserInfo) GetEmailStatusOk() (*string, bool)`

GetEmailStatusOk returns a tuple with the EmailStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailStatus

`func (o *UserInfo) SetEmailStatus(v string)`

SetEmailStatus sets EmailStatus field to given value.


### GetLanguage

`func (o *UserInfo) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *UserInfo) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *UserInfo) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetLastLoginTimestamp

`func (o *UserInfo) GetLastLoginTimestamp() string`

GetLastLoginTimestamp returns the LastLoginTimestamp field if non-nil, zero value otherwise.

### GetLastLoginTimestampOk

`func (o *UserInfo) GetLastLoginTimestampOk() (*string, bool)`

GetLastLoginTimestampOk returns a tuple with the LastLoginTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginTimestamp

`func (o *UserInfo) SetLastLoginTimestamp(v string)`

SetLastLoginTimestamp sets LastLoginTimestamp field to given value.


### GetAccessLevel

`func (o *UserInfo) GetAccessLevel() string`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *UserInfo) GetAccessLevelOk() (*string, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *UserInfo) SetAccessLevel(v string)`

SetAccessLevel sets AccessLevel field to given value.


### GetIsArchived

`func (o *UserInfo) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *UserInfo) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *UserInfo) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.


### GetAccountId

`func (o *UserInfo) GetAccountId() float32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UserInfo) GetAccountIdOk() (*float32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UserInfo) SetAccountId(v float32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UserInfo) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetId

`func (o *UserInfo) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserInfo) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserInfo) SetId(v float32)`

SetId sets Id field to given value.


### GetRevision

`func (o *UserInfo) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *UserInfo) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *UserInfo) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetEmail

`func (o *UserInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserInfo) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetCreatedTimestamp

`func (o *UserInfo) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *UserInfo) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *UserInfo) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetLinks

`func (o *UserInfo) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserInfo) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserInfo) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UserInfo) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


