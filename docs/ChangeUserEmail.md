# ChangeUserEmail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The new email address of the user | 
**RedirectUrl** | Pointer to **string** | The redirect URL after email verification | [optional] 

## Methods

### NewChangeUserEmail

`func NewChangeUserEmail(email string, ) *ChangeUserEmail`

NewChangeUserEmail instantiates a new ChangeUserEmail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeUserEmailWithDefaults

`func NewChangeUserEmailWithDefaults() *ChangeUserEmail`

NewChangeUserEmailWithDefaults instantiates a new ChangeUserEmail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ChangeUserEmail) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ChangeUserEmail) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ChangeUserEmail) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRedirectUrl

`func (o *ChangeUserEmail) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *ChangeUserEmail) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *ChangeUserEmail) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *ChangeUserEmail) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


