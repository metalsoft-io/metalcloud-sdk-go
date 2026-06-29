# QuotaLimitsBreakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effective** | [**QuotaProfileLimits**](QuotaProfileLimits.md) | The most restrictive limits merged across all applicable quota profiles | 
**Role** | Pointer to [**NullableQuotaProfileLimits**](QuotaProfileLimits.md) | Limits from the quota profile assigned to the user&#39;s role, or null if none applies | [optional] 
**Group** | Pointer to [**NullableQuotaProfileLimits**](QuotaProfileLimits.md) | Limits from the quota profile assigned to the user&#39;s LDAP/SAML group mapping, or null if none applies | [optional] 
**Account** | Pointer to [**NullableQuotaProfileLimits**](QuotaProfileLimits.md) | Limits from the quota profile assigned to the user&#39;s account, or null if none applies | [optional] 
**ParentAccount** | Pointer to [**NullableQuotaProfileLimits**](QuotaProfileLimits.md) | Limits from the quota profile assigned to the parent account, or null if none applies | [optional] 
**UserUsage** | Pointer to [**NullableQuotaProfileLimits**](QuotaProfileLimits.md) | Current aggregate resource counts for this user — only populated when ?includeUsage&#x3D;true | [optional] 
**GroupUsage** | Pointer to [**NullableQuotaProfileLimits**](QuotaProfileLimits.md) | Aggregate resource counts across all LDAP/SAML group members — only populated when ?includeUsage&#x3D;true | [optional] 
**AccountUsage** | Pointer to [**NullableQuotaProfileLimits**](QuotaProfileLimits.md) | Aggregate resource counts across all users in the same account — only populated when ?includeUsage&#x3D;true | [optional] 
**ParentAccountUsage** | Pointer to [**NullableQuotaProfileLimits**](QuotaProfileLimits.md) | Aggregate resource counts across all users in child accounts — only populated when ?includeUsage&#x3D;true | [optional] 

## Methods

### NewQuotaLimitsBreakdown

`func NewQuotaLimitsBreakdown(effective QuotaProfileLimits, ) *QuotaLimitsBreakdown`

NewQuotaLimitsBreakdown instantiates a new QuotaLimitsBreakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaLimitsBreakdownWithDefaults

`func NewQuotaLimitsBreakdownWithDefaults() *QuotaLimitsBreakdown`

NewQuotaLimitsBreakdownWithDefaults instantiates a new QuotaLimitsBreakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffective

`func (o *QuotaLimitsBreakdown) GetEffective() QuotaProfileLimits`

GetEffective returns the Effective field if non-nil, zero value otherwise.

### GetEffectiveOk

`func (o *QuotaLimitsBreakdown) GetEffectiveOk() (*QuotaProfileLimits, bool)`

GetEffectiveOk returns a tuple with the Effective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffective

`func (o *QuotaLimitsBreakdown) SetEffective(v QuotaProfileLimits)`

SetEffective sets Effective field to given value.


### GetRole

`func (o *QuotaLimitsBreakdown) GetRole() QuotaProfileLimits`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *QuotaLimitsBreakdown) GetRoleOk() (*QuotaProfileLimits, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *QuotaLimitsBreakdown) SetRole(v QuotaProfileLimits)`

SetRole sets Role field to given value.

### HasRole

`func (o *QuotaLimitsBreakdown) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *QuotaLimitsBreakdown) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *QuotaLimitsBreakdown) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetGroup

`func (o *QuotaLimitsBreakdown) GetGroup() QuotaProfileLimits`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *QuotaLimitsBreakdown) GetGroupOk() (*QuotaProfileLimits, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *QuotaLimitsBreakdown) SetGroup(v QuotaProfileLimits)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *QuotaLimitsBreakdown) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *QuotaLimitsBreakdown) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *QuotaLimitsBreakdown) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetAccount

`func (o *QuotaLimitsBreakdown) GetAccount() QuotaProfileLimits`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *QuotaLimitsBreakdown) GetAccountOk() (*QuotaProfileLimits, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *QuotaLimitsBreakdown) SetAccount(v QuotaProfileLimits)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *QuotaLimitsBreakdown) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *QuotaLimitsBreakdown) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *QuotaLimitsBreakdown) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetParentAccount

`func (o *QuotaLimitsBreakdown) GetParentAccount() QuotaProfileLimits`

GetParentAccount returns the ParentAccount field if non-nil, zero value otherwise.

### GetParentAccountOk

`func (o *QuotaLimitsBreakdown) GetParentAccountOk() (*QuotaProfileLimits, bool)`

GetParentAccountOk returns a tuple with the ParentAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccount

`func (o *QuotaLimitsBreakdown) SetParentAccount(v QuotaProfileLimits)`

SetParentAccount sets ParentAccount field to given value.

### HasParentAccount

`func (o *QuotaLimitsBreakdown) HasParentAccount() bool`

HasParentAccount returns a boolean if a field has been set.

### SetParentAccountNil

`func (o *QuotaLimitsBreakdown) SetParentAccountNil(b bool)`

 SetParentAccountNil sets the value for ParentAccount to be an explicit nil

### UnsetParentAccount
`func (o *QuotaLimitsBreakdown) UnsetParentAccount()`

UnsetParentAccount ensures that no value is present for ParentAccount, not even an explicit nil
### GetUserUsage

`func (o *QuotaLimitsBreakdown) GetUserUsage() QuotaProfileLimits`

GetUserUsage returns the UserUsage field if non-nil, zero value otherwise.

### GetUserUsageOk

`func (o *QuotaLimitsBreakdown) GetUserUsageOk() (*QuotaProfileLimits, bool)`

GetUserUsageOk returns a tuple with the UserUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUsage

`func (o *QuotaLimitsBreakdown) SetUserUsage(v QuotaProfileLimits)`

SetUserUsage sets UserUsage field to given value.

### HasUserUsage

`func (o *QuotaLimitsBreakdown) HasUserUsage() bool`

HasUserUsage returns a boolean if a field has been set.

### SetUserUsageNil

`func (o *QuotaLimitsBreakdown) SetUserUsageNil(b bool)`

 SetUserUsageNil sets the value for UserUsage to be an explicit nil

### UnsetUserUsage
`func (o *QuotaLimitsBreakdown) UnsetUserUsage()`

UnsetUserUsage ensures that no value is present for UserUsage, not even an explicit nil
### GetGroupUsage

`func (o *QuotaLimitsBreakdown) GetGroupUsage() QuotaProfileLimits`

GetGroupUsage returns the GroupUsage field if non-nil, zero value otherwise.

### GetGroupUsageOk

`func (o *QuotaLimitsBreakdown) GetGroupUsageOk() (*QuotaProfileLimits, bool)`

GetGroupUsageOk returns a tuple with the GroupUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupUsage

`func (o *QuotaLimitsBreakdown) SetGroupUsage(v QuotaProfileLimits)`

SetGroupUsage sets GroupUsage field to given value.

### HasGroupUsage

`func (o *QuotaLimitsBreakdown) HasGroupUsage() bool`

HasGroupUsage returns a boolean if a field has been set.

### SetGroupUsageNil

`func (o *QuotaLimitsBreakdown) SetGroupUsageNil(b bool)`

 SetGroupUsageNil sets the value for GroupUsage to be an explicit nil

### UnsetGroupUsage
`func (o *QuotaLimitsBreakdown) UnsetGroupUsage()`

UnsetGroupUsage ensures that no value is present for GroupUsage, not even an explicit nil
### GetAccountUsage

`func (o *QuotaLimitsBreakdown) GetAccountUsage() QuotaProfileLimits`

GetAccountUsage returns the AccountUsage field if non-nil, zero value otherwise.

### GetAccountUsageOk

`func (o *QuotaLimitsBreakdown) GetAccountUsageOk() (*QuotaProfileLimits, bool)`

GetAccountUsageOk returns a tuple with the AccountUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUsage

`func (o *QuotaLimitsBreakdown) SetAccountUsage(v QuotaProfileLimits)`

SetAccountUsage sets AccountUsage field to given value.

### HasAccountUsage

`func (o *QuotaLimitsBreakdown) HasAccountUsage() bool`

HasAccountUsage returns a boolean if a field has been set.

### SetAccountUsageNil

`func (o *QuotaLimitsBreakdown) SetAccountUsageNil(b bool)`

 SetAccountUsageNil sets the value for AccountUsage to be an explicit nil

### UnsetAccountUsage
`func (o *QuotaLimitsBreakdown) UnsetAccountUsage()`

UnsetAccountUsage ensures that no value is present for AccountUsage, not even an explicit nil
### GetParentAccountUsage

`func (o *QuotaLimitsBreakdown) GetParentAccountUsage() QuotaProfileLimits`

GetParentAccountUsage returns the ParentAccountUsage field if non-nil, zero value otherwise.

### GetParentAccountUsageOk

`func (o *QuotaLimitsBreakdown) GetParentAccountUsageOk() (*QuotaProfileLimits, bool)`

GetParentAccountUsageOk returns a tuple with the ParentAccountUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountUsage

`func (o *QuotaLimitsBreakdown) SetParentAccountUsage(v QuotaProfileLimits)`

SetParentAccountUsage sets ParentAccountUsage field to given value.

### HasParentAccountUsage

`func (o *QuotaLimitsBreakdown) HasParentAccountUsage() bool`

HasParentAccountUsage returns a boolean if a field has been set.

### SetParentAccountUsageNil

`func (o *QuotaLimitsBreakdown) SetParentAccountUsageNil(b bool)`

 SetParentAccountUsageNil sets the value for ParentAccountUsage to be an explicit nil

### UnsetParentAccountUsage
`func (o *QuotaLimitsBreakdown) UnsetParentAccountUsage()`

UnsetParentAccountUsage ensures that no value is present for ParentAccountUsage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


