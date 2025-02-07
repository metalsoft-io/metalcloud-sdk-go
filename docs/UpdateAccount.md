# UpdateAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentAccountId** | Pointer to **float32** | The ID of the parent account | [optional] 
**Name** | **string** | The name of the account | 
**Code** | Pointer to **string** | The code of the account | [optional] 
**FiscalNumber** | Pointer to **string** | The fiscal number of the account | [optional] 
**Address** | Pointer to [**AccountAddress**](AccountAddress.md) |  | [optional] 
**PrimaryContactId** | Pointer to **float32** | The user ID of the primary contact | [optional] 
**SecondaryContactId** | Pointer to **float32** | The user ID of the secondary contact | [optional] 

## Methods

### NewUpdateAccount

`func NewUpdateAccount(name string, ) *UpdateAccount`

NewUpdateAccount instantiates a new UpdateAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccountWithDefaults

`func NewUpdateAccountWithDefaults() *UpdateAccount`

NewUpdateAccountWithDefaults instantiates a new UpdateAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentAccountId

`func (o *UpdateAccount) GetParentAccountId() float32`

GetParentAccountId returns the ParentAccountId field if non-nil, zero value otherwise.

### GetParentAccountIdOk

`func (o *UpdateAccount) GetParentAccountIdOk() (*float32, bool)`

GetParentAccountIdOk returns a tuple with the ParentAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountId

`func (o *UpdateAccount) SetParentAccountId(v float32)`

SetParentAccountId sets ParentAccountId field to given value.

### HasParentAccountId

`func (o *UpdateAccount) HasParentAccountId() bool`

HasParentAccountId returns a boolean if a field has been set.

### GetName

`func (o *UpdateAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAccount) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *UpdateAccount) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateAccount) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateAccount) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateAccount) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetFiscalNumber

`func (o *UpdateAccount) GetFiscalNumber() string`

GetFiscalNumber returns the FiscalNumber field if non-nil, zero value otherwise.

### GetFiscalNumberOk

`func (o *UpdateAccount) GetFiscalNumberOk() (*string, bool)`

GetFiscalNumberOk returns a tuple with the FiscalNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalNumber

`func (o *UpdateAccount) SetFiscalNumber(v string)`

SetFiscalNumber sets FiscalNumber field to given value.

### HasFiscalNumber

`func (o *UpdateAccount) HasFiscalNumber() bool`

HasFiscalNumber returns a boolean if a field has been set.

### GetAddress

`func (o *UpdateAccount) GetAddress() AccountAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UpdateAccount) GetAddressOk() (*AccountAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UpdateAccount) SetAddress(v AccountAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UpdateAccount) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPrimaryContactId

`func (o *UpdateAccount) GetPrimaryContactId() float32`

GetPrimaryContactId returns the PrimaryContactId field if non-nil, zero value otherwise.

### GetPrimaryContactIdOk

`func (o *UpdateAccount) GetPrimaryContactIdOk() (*float32, bool)`

GetPrimaryContactIdOk returns a tuple with the PrimaryContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryContactId

`func (o *UpdateAccount) SetPrimaryContactId(v float32)`

SetPrimaryContactId sets PrimaryContactId field to given value.

### HasPrimaryContactId

`func (o *UpdateAccount) HasPrimaryContactId() bool`

HasPrimaryContactId returns a boolean if a field has been set.

### GetSecondaryContactId

`func (o *UpdateAccount) GetSecondaryContactId() float32`

GetSecondaryContactId returns the SecondaryContactId field if non-nil, zero value otherwise.

### GetSecondaryContactIdOk

`func (o *UpdateAccount) GetSecondaryContactIdOk() (*float32, bool)`

GetSecondaryContactIdOk returns a tuple with the SecondaryContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryContactId

`func (o *UpdateAccount) SetSecondaryContactId(v float32)`

SetSecondaryContactId sets SecondaryContactId field to given value.

### HasSecondaryContactId

`func (o *UpdateAccount) HasSecondaryContactId() bool`

HasSecondaryContactId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


