# UserProviderDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Authentication provider identifier | 
**Enabled** | **bool** | Whether this provider is active | 
**Domains** | **[]string** | Email domains handled by this provider | 
**Subtrees** | Pointer to **[]string** | LDAP subtrees for directory lookups — for non-email domain formats (LDAP only) | [optional] 
**StoreMail** | Pointer to **bool** | Whether to store the LDAP mail attribute as the user email; when false the username is used instead (LDAP only) | [optional] 

## Methods

### NewUserProviderDto

`func NewUserProviderDto(name string, enabled bool, domains []string, ) *UserProviderDto`

NewUserProviderDto instantiates a new UserProviderDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserProviderDtoWithDefaults

`func NewUserProviderDtoWithDefaults() *UserProviderDto`

NewUserProviderDtoWithDefaults instantiates a new UserProviderDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UserProviderDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserProviderDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserProviderDto) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *UserProviderDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UserProviderDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UserProviderDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetDomains

`func (o *UserProviderDto) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *UserProviderDto) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *UserProviderDto) SetDomains(v []string)`

SetDomains sets Domains field to given value.


### GetSubtrees

`func (o *UserProviderDto) GetSubtrees() []string`

GetSubtrees returns the Subtrees field if non-nil, zero value otherwise.

### GetSubtreesOk

`func (o *UserProviderDto) GetSubtreesOk() (*[]string, bool)`

GetSubtreesOk returns a tuple with the Subtrees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtrees

`func (o *UserProviderDto) SetSubtrees(v []string)`

SetSubtrees sets Subtrees field to given value.

### HasSubtrees

`func (o *UserProviderDto) HasSubtrees() bool`

HasSubtrees returns a boolean if a field has been set.

### GetStoreMail

`func (o *UserProviderDto) GetStoreMail() bool`

GetStoreMail returns the StoreMail field if non-nil, zero value otherwise.

### GetStoreMailOk

`func (o *UserProviderDto) GetStoreMailOk() (*bool, bool)`

GetStoreMailOk returns a tuple with the StoreMail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreMail

`func (o *UserProviderDto) SetStoreMail(v bool)`

SetStoreMail sets StoreMail field to given value.

### HasStoreMail

`func (o *UserProviderDto) HasStoreMail() bool`

HasStoreMail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


