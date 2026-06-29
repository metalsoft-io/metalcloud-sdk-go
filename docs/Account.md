# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentAccountId** | Pointer to **int64** | The ID of the parent account | [optional] 
**Name** | **string** | The name of the account | 
**Code** | Pointer to **string** | The code of the account | [optional] 
**FiscalNumber** | Pointer to **string** | The fiscal number of the account | [optional] 
**Address** | Pointer to [**AccountAddress**](AccountAddress.md) |  | [optional] 
**PrimaryContactId** | Pointer to **int64** | The user ID of the primary contact | [optional] 
**SecondaryContactId** | Pointer to **int64** | The user ID of the secondary contact | [optional] 
**QuotaProfileId** | Pointer to **string** | The ID of the quota profile assigned to this account | [optional] 
**IsArchived** | Pointer to **bool** | Whether the account is archived | [optional] 
**Id** | **int64** | Account ID | 
**Revision** | **int64** | Revision number | 
**Config** | [**AccountConfig**](AccountConfig.md) | The new configuration of the account. | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewAccount

`func NewAccount(name string, id int64, revision int64, config AccountConfig, ) *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentAccountId

`func (o *Account) GetParentAccountId() int64`

GetParentAccountId returns the ParentAccountId field if non-nil, zero value otherwise.

### GetParentAccountIdOk

`func (o *Account) GetParentAccountIdOk() (*int64, bool)`

GetParentAccountIdOk returns a tuple with the ParentAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountId

`func (o *Account) SetParentAccountId(v int64)`

SetParentAccountId sets ParentAccountId field to given value.

### HasParentAccountId

`func (o *Account) HasParentAccountId() bool`

HasParentAccountId returns a boolean if a field has been set.

### GetName

`func (o *Account) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Account) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Account) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *Account) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Account) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Account) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Account) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetFiscalNumber

`func (o *Account) GetFiscalNumber() string`

GetFiscalNumber returns the FiscalNumber field if non-nil, zero value otherwise.

### GetFiscalNumberOk

`func (o *Account) GetFiscalNumberOk() (*string, bool)`

GetFiscalNumberOk returns a tuple with the FiscalNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalNumber

`func (o *Account) SetFiscalNumber(v string)`

SetFiscalNumber sets FiscalNumber field to given value.

### HasFiscalNumber

`func (o *Account) HasFiscalNumber() bool`

HasFiscalNumber returns a boolean if a field has been set.

### GetAddress

`func (o *Account) GetAddress() AccountAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Account) GetAddressOk() (*AccountAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Account) SetAddress(v AccountAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Account) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPrimaryContactId

`func (o *Account) GetPrimaryContactId() int64`

GetPrimaryContactId returns the PrimaryContactId field if non-nil, zero value otherwise.

### GetPrimaryContactIdOk

`func (o *Account) GetPrimaryContactIdOk() (*int64, bool)`

GetPrimaryContactIdOk returns a tuple with the PrimaryContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryContactId

`func (o *Account) SetPrimaryContactId(v int64)`

SetPrimaryContactId sets PrimaryContactId field to given value.

### HasPrimaryContactId

`func (o *Account) HasPrimaryContactId() bool`

HasPrimaryContactId returns a boolean if a field has been set.

### GetSecondaryContactId

`func (o *Account) GetSecondaryContactId() int64`

GetSecondaryContactId returns the SecondaryContactId field if non-nil, zero value otherwise.

### GetSecondaryContactIdOk

`func (o *Account) GetSecondaryContactIdOk() (*int64, bool)`

GetSecondaryContactIdOk returns a tuple with the SecondaryContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryContactId

`func (o *Account) SetSecondaryContactId(v int64)`

SetSecondaryContactId sets SecondaryContactId field to given value.

### HasSecondaryContactId

`func (o *Account) HasSecondaryContactId() bool`

HasSecondaryContactId returns a boolean if a field has been set.

### GetQuotaProfileId

`func (o *Account) GetQuotaProfileId() string`

GetQuotaProfileId returns the QuotaProfileId field if non-nil, zero value otherwise.

### GetQuotaProfileIdOk

`func (o *Account) GetQuotaProfileIdOk() (*string, bool)`

GetQuotaProfileIdOk returns a tuple with the QuotaProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaProfileId

`func (o *Account) SetQuotaProfileId(v string)`

SetQuotaProfileId sets QuotaProfileId field to given value.

### HasQuotaProfileId

`func (o *Account) HasQuotaProfileId() bool`

HasQuotaProfileId returns a boolean if a field has been set.

### GetIsArchived

`func (o *Account) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *Account) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *Account) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *Account) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### GetId

`func (o *Account) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Account) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Account) SetId(v int64)`

SetId sets Id field to given value.


### GetRevision

`func (o *Account) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *Account) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *Account) SetRevision(v int64)`

SetRevision sets Revision field to given value.


### GetConfig

`func (o *Account) GetConfig() AccountConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Account) GetConfigOk() (*AccountConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Account) SetConfig(v AccountConfig)`

SetConfig sets Config field to given value.


### GetLinks

`func (o *Account) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Account) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Account) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Account) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


