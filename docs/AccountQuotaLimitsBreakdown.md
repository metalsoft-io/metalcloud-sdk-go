# AccountQuotaLimitsBreakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effective** | [**QuotaProfileLimits**](QuotaProfileLimits.md) | The most restrictive limits merged across the account and parent account quota profiles. Individual users may be further restricted by their role or group quota profiles. | 
**Account** | Pointer to [**NullableQuotaProfileLimits**](QuotaProfileLimits.md) | Limits from the quota profile assigned to the account, or null if none applies | [optional] 
**ParentAccount** | Pointer to [**NullableQuotaProfileLimits**](QuotaProfileLimits.md) | Limits from the quota profile assigned to the parent account, or null if none applies | [optional] 
**AccountUsage** | Pointer to [**NullableQuotaUsage**](QuotaUsage.md) | Aggregate resource counts across all users in the account — only populated when ?includeUsage&#x3D;true | [optional] 
**ParentAccountUsage** | Pointer to [**NullableQuotaUsage**](QuotaUsage.md) | Aggregate resource counts across all users in all child accounts of the parent account — only populated when ?includeUsage&#x3D;true | [optional] 

## Methods

### NewAccountQuotaLimitsBreakdown

`func NewAccountQuotaLimitsBreakdown(effective QuotaProfileLimits, ) *AccountQuotaLimitsBreakdown`

NewAccountQuotaLimitsBreakdown instantiates a new AccountQuotaLimitsBreakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountQuotaLimitsBreakdownWithDefaults

`func NewAccountQuotaLimitsBreakdownWithDefaults() *AccountQuotaLimitsBreakdown`

NewAccountQuotaLimitsBreakdownWithDefaults instantiates a new AccountQuotaLimitsBreakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffective

`func (o *AccountQuotaLimitsBreakdown) GetEffective() QuotaProfileLimits`

GetEffective returns the Effective field if non-nil, zero value otherwise.

### GetEffectiveOk

`func (o *AccountQuotaLimitsBreakdown) GetEffectiveOk() (*QuotaProfileLimits, bool)`

GetEffectiveOk returns a tuple with the Effective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffective

`func (o *AccountQuotaLimitsBreakdown) SetEffective(v QuotaProfileLimits)`

SetEffective sets Effective field to given value.


### GetAccount

`func (o *AccountQuotaLimitsBreakdown) GetAccount() QuotaProfileLimits`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AccountQuotaLimitsBreakdown) GetAccountOk() (*QuotaProfileLimits, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AccountQuotaLimitsBreakdown) SetAccount(v QuotaProfileLimits)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AccountQuotaLimitsBreakdown) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *AccountQuotaLimitsBreakdown) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *AccountQuotaLimitsBreakdown) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetParentAccount

`func (o *AccountQuotaLimitsBreakdown) GetParentAccount() QuotaProfileLimits`

GetParentAccount returns the ParentAccount field if non-nil, zero value otherwise.

### GetParentAccountOk

`func (o *AccountQuotaLimitsBreakdown) GetParentAccountOk() (*QuotaProfileLimits, bool)`

GetParentAccountOk returns a tuple with the ParentAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccount

`func (o *AccountQuotaLimitsBreakdown) SetParentAccount(v QuotaProfileLimits)`

SetParentAccount sets ParentAccount field to given value.

### HasParentAccount

`func (o *AccountQuotaLimitsBreakdown) HasParentAccount() bool`

HasParentAccount returns a boolean if a field has been set.

### SetParentAccountNil

`func (o *AccountQuotaLimitsBreakdown) SetParentAccountNil(b bool)`

 SetParentAccountNil sets the value for ParentAccount to be an explicit nil

### UnsetParentAccount
`func (o *AccountQuotaLimitsBreakdown) UnsetParentAccount()`

UnsetParentAccount ensures that no value is present for ParentAccount, not even an explicit nil
### GetAccountUsage

`func (o *AccountQuotaLimitsBreakdown) GetAccountUsage() QuotaUsage`

GetAccountUsage returns the AccountUsage field if non-nil, zero value otherwise.

### GetAccountUsageOk

`func (o *AccountQuotaLimitsBreakdown) GetAccountUsageOk() (*QuotaUsage, bool)`

GetAccountUsageOk returns a tuple with the AccountUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUsage

`func (o *AccountQuotaLimitsBreakdown) SetAccountUsage(v QuotaUsage)`

SetAccountUsage sets AccountUsage field to given value.

### HasAccountUsage

`func (o *AccountQuotaLimitsBreakdown) HasAccountUsage() bool`

HasAccountUsage returns a boolean if a field has been set.

### SetAccountUsageNil

`func (o *AccountQuotaLimitsBreakdown) SetAccountUsageNil(b bool)`

 SetAccountUsageNil sets the value for AccountUsage to be an explicit nil

### UnsetAccountUsage
`func (o *AccountQuotaLimitsBreakdown) UnsetAccountUsage()`

UnsetAccountUsage ensures that no value is present for AccountUsage, not even an explicit nil
### GetParentAccountUsage

`func (o *AccountQuotaLimitsBreakdown) GetParentAccountUsage() QuotaUsage`

GetParentAccountUsage returns the ParentAccountUsage field if non-nil, zero value otherwise.

### GetParentAccountUsageOk

`func (o *AccountQuotaLimitsBreakdown) GetParentAccountUsageOk() (*QuotaUsage, bool)`

GetParentAccountUsageOk returns a tuple with the ParentAccountUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountUsage

`func (o *AccountQuotaLimitsBreakdown) SetParentAccountUsage(v QuotaUsage)`

SetParentAccountUsage sets ParentAccountUsage field to given value.

### HasParentAccountUsage

`func (o *AccountQuotaLimitsBreakdown) HasParentAccountUsage() bool`

HasParentAccountUsage returns a boolean if a field has been set.

### SetParentAccountUsageNil

`func (o *AccountQuotaLimitsBreakdown) SetParentAccountUsageNil(b bool)`

 SetParentAccountUsageNil sets the value for ParentAccountUsage to be an explicit nil

### UnsetParentAccountUsage
`func (o *AccountQuotaLimitsBreakdown) UnsetParentAccountUsage()`

UnsetParentAccountUsage ensures that no value is present for ParentAccountUsage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


