# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | Account ID | 
**Revision** | **float32** | Revision number | 
**ParentAccountId** | Pointer to **float32** | The ID of the parent account | [optional] 
**Name** | **string** | The name of the account | 
**Code** | Pointer to **string** | The code of the account | [optional] 
**FiscalNumber** | Pointer to **string** | The fiscal number of the account | [optional] 
**Address** | Pointer to [**AccountAddress**](AccountAddress.md) |  | [optional] 
**PrimaryContactId** | Pointer to **float32** | The user ID of the primary contact | [optional] 
**SecondaryContactId** | Pointer to **float32** | The user ID of the secondary contact | [optional] 
**Archived** | Pointer to **float32** | Whether the account is archived | [optional] 
**Limits** | [**AccountLimits**](AccountLimits.md) |  | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewAccount

`func NewAccount(id float32, revision float32, name string, limits AccountLimits, ) *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Account) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Account) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Account) SetId(v float32)`

SetId sets Id field to given value.


### GetRevision

`func (o *Account) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *Account) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *Account) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetParentAccountId

`func (o *Account) GetParentAccountId() float32`

GetParentAccountId returns the ParentAccountId field if non-nil, zero value otherwise.

### GetParentAccountIdOk

`func (o *Account) GetParentAccountIdOk() (*float32, bool)`

GetParentAccountIdOk returns a tuple with the ParentAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountId

`func (o *Account) SetParentAccountId(v float32)`

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

`func (o *Account) GetPrimaryContactId() float32`

GetPrimaryContactId returns the PrimaryContactId field if non-nil, zero value otherwise.

### GetPrimaryContactIdOk

`func (o *Account) GetPrimaryContactIdOk() (*float32, bool)`

GetPrimaryContactIdOk returns a tuple with the PrimaryContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryContactId

`func (o *Account) SetPrimaryContactId(v float32)`

SetPrimaryContactId sets PrimaryContactId field to given value.

### HasPrimaryContactId

`func (o *Account) HasPrimaryContactId() bool`

HasPrimaryContactId returns a boolean if a field has been set.

### GetSecondaryContactId

`func (o *Account) GetSecondaryContactId() float32`

GetSecondaryContactId returns the SecondaryContactId field if non-nil, zero value otherwise.

### GetSecondaryContactIdOk

`func (o *Account) GetSecondaryContactIdOk() (*float32, bool)`

GetSecondaryContactIdOk returns a tuple with the SecondaryContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryContactId

`func (o *Account) SetSecondaryContactId(v float32)`

SetSecondaryContactId sets SecondaryContactId field to given value.

### HasSecondaryContactId

`func (o *Account) HasSecondaryContactId() bool`

HasSecondaryContactId returns a boolean if a field has been set.

### GetArchived

`func (o *Account) GetArchived() float32`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *Account) GetArchivedOk() (*float32, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *Account) SetArchived(v float32)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *Account) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetLimits

`func (o *Account) GetLimits() AccountLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *Account) GetLimitsOk() (*AccountLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *Account) SetLimits(v AccountLimits)`

SetLimits sets Limits field to given value.


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


