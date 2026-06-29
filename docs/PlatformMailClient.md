# PlatformMailClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StrHostname** | Pointer to **string** | SMTP server hostname | [optional] 
**NPort** | Pointer to **float32** | SMTP server port | [optional] 
**StrUsername** | Pointer to **string** | SMTP authentication username | [optional] 
**StrFromEmail** | Pointer to **string** | Sender email address | [optional] 
**StrPassword** | Pointer to **string** | SMTP authentication password | [optional] 
**StrDisplayName** | Pointer to **string** | Sender display name shown in email clients | [optional] 
**StrSubjectPrefix** | Pointer to **string** | Subject line prefix prepended to all outgoing emails | [optional] 

## Methods

### NewPlatformMailClient

`func NewPlatformMailClient() *PlatformMailClient`

NewPlatformMailClient instantiates a new PlatformMailClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformMailClientWithDefaults

`func NewPlatformMailClientWithDefaults() *PlatformMailClient`

NewPlatformMailClientWithDefaults instantiates a new PlatformMailClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStrHostname

`func (o *PlatformMailClient) GetStrHostname() string`

GetStrHostname returns the StrHostname field if non-nil, zero value otherwise.

### GetStrHostnameOk

`func (o *PlatformMailClient) GetStrHostnameOk() (*string, bool)`

GetStrHostnameOk returns a tuple with the StrHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrHostname

`func (o *PlatformMailClient) SetStrHostname(v string)`

SetStrHostname sets StrHostname field to given value.

### HasStrHostname

`func (o *PlatformMailClient) HasStrHostname() bool`

HasStrHostname returns a boolean if a field has been set.

### GetNPort

`func (o *PlatformMailClient) GetNPort() float32`

GetNPort returns the NPort field if non-nil, zero value otherwise.

### GetNPortOk

`func (o *PlatformMailClient) GetNPortOk() (*float32, bool)`

GetNPortOk returns a tuple with the NPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPort

`func (o *PlatformMailClient) SetNPort(v float32)`

SetNPort sets NPort field to given value.

### HasNPort

`func (o *PlatformMailClient) HasNPort() bool`

HasNPort returns a boolean if a field has been set.

### GetStrUsername

`func (o *PlatformMailClient) GetStrUsername() string`

GetStrUsername returns the StrUsername field if non-nil, zero value otherwise.

### GetStrUsernameOk

`func (o *PlatformMailClient) GetStrUsernameOk() (*string, bool)`

GetStrUsernameOk returns a tuple with the StrUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrUsername

`func (o *PlatformMailClient) SetStrUsername(v string)`

SetStrUsername sets StrUsername field to given value.

### HasStrUsername

`func (o *PlatformMailClient) HasStrUsername() bool`

HasStrUsername returns a boolean if a field has been set.

### GetStrFromEmail

`func (o *PlatformMailClient) GetStrFromEmail() string`

GetStrFromEmail returns the StrFromEmail field if non-nil, zero value otherwise.

### GetStrFromEmailOk

`func (o *PlatformMailClient) GetStrFromEmailOk() (*string, bool)`

GetStrFromEmailOk returns a tuple with the StrFromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrFromEmail

`func (o *PlatformMailClient) SetStrFromEmail(v string)`

SetStrFromEmail sets StrFromEmail field to given value.

### HasStrFromEmail

`func (o *PlatformMailClient) HasStrFromEmail() bool`

HasStrFromEmail returns a boolean if a field has been set.

### GetStrPassword

`func (o *PlatformMailClient) GetStrPassword() string`

GetStrPassword returns the StrPassword field if non-nil, zero value otherwise.

### GetStrPasswordOk

`func (o *PlatformMailClient) GetStrPasswordOk() (*string, bool)`

GetStrPasswordOk returns a tuple with the StrPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrPassword

`func (o *PlatformMailClient) SetStrPassword(v string)`

SetStrPassword sets StrPassword field to given value.

### HasStrPassword

`func (o *PlatformMailClient) HasStrPassword() bool`

HasStrPassword returns a boolean if a field has been set.

### GetStrDisplayName

`func (o *PlatformMailClient) GetStrDisplayName() string`

GetStrDisplayName returns the StrDisplayName field if non-nil, zero value otherwise.

### GetStrDisplayNameOk

`func (o *PlatformMailClient) GetStrDisplayNameOk() (*string, bool)`

GetStrDisplayNameOk returns a tuple with the StrDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrDisplayName

`func (o *PlatformMailClient) SetStrDisplayName(v string)`

SetStrDisplayName sets StrDisplayName field to given value.

### HasStrDisplayName

`func (o *PlatformMailClient) HasStrDisplayName() bool`

HasStrDisplayName returns a boolean if a field has been set.

### GetStrSubjectPrefix

`func (o *PlatformMailClient) GetStrSubjectPrefix() string`

GetStrSubjectPrefix returns the StrSubjectPrefix field if non-nil, zero value otherwise.

### GetStrSubjectPrefixOk

`func (o *PlatformMailClient) GetStrSubjectPrefixOk() (*string, bool)`

GetStrSubjectPrefixOk returns a tuple with the StrSubjectPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrSubjectPrefix

`func (o *PlatformMailClient) SetStrSubjectPrefix(v string)`

SetStrSubjectPrefix sets StrSubjectPrefix field to given value.

### HasStrSubjectPrefix

`func (o *PlatformMailClient) HasStrSubjectPrefix() bool`

HasStrSubjectPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


