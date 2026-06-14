# PlatformMailClientDto

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

### NewPlatformMailClientDto

`func NewPlatformMailClientDto() *PlatformMailClientDto`

NewPlatformMailClientDto instantiates a new PlatformMailClientDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformMailClientDtoWithDefaults

`func NewPlatformMailClientDtoWithDefaults() *PlatformMailClientDto`

NewPlatformMailClientDtoWithDefaults instantiates a new PlatformMailClientDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStrHostname

`func (o *PlatformMailClientDto) GetStrHostname() string`

GetStrHostname returns the StrHostname field if non-nil, zero value otherwise.

### GetStrHostnameOk

`func (o *PlatformMailClientDto) GetStrHostnameOk() (*string, bool)`

GetStrHostnameOk returns a tuple with the StrHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrHostname

`func (o *PlatformMailClientDto) SetStrHostname(v string)`

SetStrHostname sets StrHostname field to given value.

### HasStrHostname

`func (o *PlatformMailClientDto) HasStrHostname() bool`

HasStrHostname returns a boolean if a field has been set.

### GetNPort

`func (o *PlatformMailClientDto) GetNPort() float32`

GetNPort returns the NPort field if non-nil, zero value otherwise.

### GetNPortOk

`func (o *PlatformMailClientDto) GetNPortOk() (*float32, bool)`

GetNPortOk returns a tuple with the NPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNPort

`func (o *PlatformMailClientDto) SetNPort(v float32)`

SetNPort sets NPort field to given value.

### HasNPort

`func (o *PlatformMailClientDto) HasNPort() bool`

HasNPort returns a boolean if a field has been set.

### GetStrUsername

`func (o *PlatformMailClientDto) GetStrUsername() string`

GetStrUsername returns the StrUsername field if non-nil, zero value otherwise.

### GetStrUsernameOk

`func (o *PlatformMailClientDto) GetStrUsernameOk() (*string, bool)`

GetStrUsernameOk returns a tuple with the StrUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrUsername

`func (o *PlatformMailClientDto) SetStrUsername(v string)`

SetStrUsername sets StrUsername field to given value.

### HasStrUsername

`func (o *PlatformMailClientDto) HasStrUsername() bool`

HasStrUsername returns a boolean if a field has been set.

### GetStrFromEmail

`func (o *PlatformMailClientDto) GetStrFromEmail() string`

GetStrFromEmail returns the StrFromEmail field if non-nil, zero value otherwise.

### GetStrFromEmailOk

`func (o *PlatformMailClientDto) GetStrFromEmailOk() (*string, bool)`

GetStrFromEmailOk returns a tuple with the StrFromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrFromEmail

`func (o *PlatformMailClientDto) SetStrFromEmail(v string)`

SetStrFromEmail sets StrFromEmail field to given value.

### HasStrFromEmail

`func (o *PlatformMailClientDto) HasStrFromEmail() bool`

HasStrFromEmail returns a boolean if a field has been set.

### GetStrPassword

`func (o *PlatformMailClientDto) GetStrPassword() string`

GetStrPassword returns the StrPassword field if non-nil, zero value otherwise.

### GetStrPasswordOk

`func (o *PlatformMailClientDto) GetStrPasswordOk() (*string, bool)`

GetStrPasswordOk returns a tuple with the StrPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrPassword

`func (o *PlatformMailClientDto) SetStrPassword(v string)`

SetStrPassword sets StrPassword field to given value.

### HasStrPassword

`func (o *PlatformMailClientDto) HasStrPassword() bool`

HasStrPassword returns a boolean if a field has been set.

### GetStrDisplayName

`func (o *PlatformMailClientDto) GetStrDisplayName() string`

GetStrDisplayName returns the StrDisplayName field if non-nil, zero value otherwise.

### GetStrDisplayNameOk

`func (o *PlatformMailClientDto) GetStrDisplayNameOk() (*string, bool)`

GetStrDisplayNameOk returns a tuple with the StrDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrDisplayName

`func (o *PlatformMailClientDto) SetStrDisplayName(v string)`

SetStrDisplayName sets StrDisplayName field to given value.

### HasStrDisplayName

`func (o *PlatformMailClientDto) HasStrDisplayName() bool`

HasStrDisplayName returns a boolean if a field has been set.

### GetStrSubjectPrefix

`func (o *PlatformMailClientDto) GetStrSubjectPrefix() string`

GetStrSubjectPrefix returns the StrSubjectPrefix field if non-nil, zero value otherwise.

### GetStrSubjectPrefixOk

`func (o *PlatformMailClientDto) GetStrSubjectPrefixOk() (*string, bool)`

GetStrSubjectPrefixOk returns a tuple with the StrSubjectPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrSubjectPrefix

`func (o *PlatformMailClientDto) SetStrSubjectPrefix(v string)`

SetStrSubjectPrefix sets StrSubjectPrefix field to given value.

### HasStrSubjectPrefix

`func (o *PlatformMailClientDto) HasStrSubjectPrefix() bool`

HasStrSubjectPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


