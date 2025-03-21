# AccountConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | **float32** | Revision number | 
**ParentAccountId** | Pointer to **float32** | The ID of the parent account | [optional] 
**Name** | **string** | The name of the account | 
**Code** | Pointer to **string** | The code of the account | [optional] 
**FiscalNumber** | Pointer to **string** | The fiscal number of the account | [optional] 
**Address** | Pointer to [**AccountAddress**](AccountAddress.md) |  | [optional] 
**PrimaryContactId** | Pointer to **float32** | The user ID of the primary contact | [optional] 
**SecondaryContactId** | Pointer to **float32** | The user ID of the secondary contact | [optional] 
**IsArchived** | Pointer to **bool** | Whether the account is archived | [optional] 

## Methods

### NewAccountConfig

`func NewAccountConfig(revision float32, name string, ) *AccountConfig`

NewAccountConfig instantiates a new AccountConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountConfigWithDefaults

`func NewAccountConfigWithDefaults() *AccountConfig`

NewAccountConfigWithDefaults instantiates a new AccountConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *AccountConfig) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *AccountConfig) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *AccountConfig) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetParentAccountId

`func (o *AccountConfig) GetParentAccountId() float32`

GetParentAccountId returns the ParentAccountId field if non-nil, zero value otherwise.

### GetParentAccountIdOk

`func (o *AccountConfig) GetParentAccountIdOk() (*float32, bool)`

GetParentAccountIdOk returns a tuple with the ParentAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountId

`func (o *AccountConfig) SetParentAccountId(v float32)`

SetParentAccountId sets ParentAccountId field to given value.

### HasParentAccountId

`func (o *AccountConfig) HasParentAccountId() bool`

HasParentAccountId returns a boolean if a field has been set.

### GetName

`func (o *AccountConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountConfig) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *AccountConfig) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AccountConfig) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AccountConfig) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AccountConfig) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetFiscalNumber

`func (o *AccountConfig) GetFiscalNumber() string`

GetFiscalNumber returns the FiscalNumber field if non-nil, zero value otherwise.

### GetFiscalNumberOk

`func (o *AccountConfig) GetFiscalNumberOk() (*string, bool)`

GetFiscalNumberOk returns a tuple with the FiscalNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalNumber

`func (o *AccountConfig) SetFiscalNumber(v string)`

SetFiscalNumber sets FiscalNumber field to given value.

### HasFiscalNumber

`func (o *AccountConfig) HasFiscalNumber() bool`

HasFiscalNumber returns a boolean if a field has been set.

### GetAddress

`func (o *AccountConfig) GetAddress() AccountAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AccountConfig) GetAddressOk() (*AccountAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AccountConfig) SetAddress(v AccountAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AccountConfig) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPrimaryContactId

`func (o *AccountConfig) GetPrimaryContactId() float32`

GetPrimaryContactId returns the PrimaryContactId field if non-nil, zero value otherwise.

### GetPrimaryContactIdOk

`func (o *AccountConfig) GetPrimaryContactIdOk() (*float32, bool)`

GetPrimaryContactIdOk returns a tuple with the PrimaryContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryContactId

`func (o *AccountConfig) SetPrimaryContactId(v float32)`

SetPrimaryContactId sets PrimaryContactId field to given value.

### HasPrimaryContactId

`func (o *AccountConfig) HasPrimaryContactId() bool`

HasPrimaryContactId returns a boolean if a field has been set.

### GetSecondaryContactId

`func (o *AccountConfig) GetSecondaryContactId() float32`

GetSecondaryContactId returns the SecondaryContactId field if non-nil, zero value otherwise.

### GetSecondaryContactIdOk

`func (o *AccountConfig) GetSecondaryContactIdOk() (*float32, bool)`

GetSecondaryContactIdOk returns a tuple with the SecondaryContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryContactId

`func (o *AccountConfig) SetSecondaryContactId(v float32)`

SetSecondaryContactId sets SecondaryContactId field to given value.

### HasSecondaryContactId

`func (o *AccountConfig) HasSecondaryContactId() bool`

HasSecondaryContactId returns a boolean if a field has been set.

### GetIsArchived

`func (o *AccountConfig) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *AccountConfig) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *AccountConfig) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *AccountConfig) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


