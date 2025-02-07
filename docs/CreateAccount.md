# CreateAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentAccountId** | Pointer to **float32** | The ID of the parent account | [optional] 
**Name** | **string** | The name of the account | 
**Code** | Pointer to **string** | The code of the account | [optional] 
**FiscalNumber** | Pointer to **string** | The fiscal number of the account | [optional] 
**Address** | Pointer to [**AccountAddress**](AccountAddress.md) |  | [optional] 

## Methods

### NewCreateAccount

`func NewCreateAccount(name string, ) *CreateAccount`

NewCreateAccount instantiates a new CreateAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountWithDefaults

`func NewCreateAccountWithDefaults() *CreateAccount`

NewCreateAccountWithDefaults instantiates a new CreateAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentAccountId

`func (o *CreateAccount) GetParentAccountId() float32`

GetParentAccountId returns the ParentAccountId field if non-nil, zero value otherwise.

### GetParentAccountIdOk

`func (o *CreateAccount) GetParentAccountIdOk() (*float32, bool)`

GetParentAccountIdOk returns a tuple with the ParentAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountId

`func (o *CreateAccount) SetParentAccountId(v float32)`

SetParentAccountId sets ParentAccountId field to given value.

### HasParentAccountId

`func (o *CreateAccount) HasParentAccountId() bool`

HasParentAccountId returns a boolean if a field has been set.

### GetName

`func (o *CreateAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAccount) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *CreateAccount) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateAccount) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateAccount) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CreateAccount) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetFiscalNumber

`func (o *CreateAccount) GetFiscalNumber() string`

GetFiscalNumber returns the FiscalNumber field if non-nil, zero value otherwise.

### GetFiscalNumberOk

`func (o *CreateAccount) GetFiscalNumberOk() (*string, bool)`

GetFiscalNumberOk returns a tuple with the FiscalNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalNumber

`func (o *CreateAccount) SetFiscalNumber(v string)`

SetFiscalNumber sets FiscalNumber field to given value.

### HasFiscalNumber

`func (o *CreateAccount) HasFiscalNumber() bool`

HasFiscalNumber returns a boolean if a field has been set.

### GetAddress

`func (o *CreateAccount) GetAddress() AccountAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreateAccount) GetAddressOk() (*AccountAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreateAccount) SetAddress(v AccountAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CreateAccount) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


