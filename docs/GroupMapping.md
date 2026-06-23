# GroupMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupName** | **map[string]interface{}** | LDAP or SAML group name | 
**RoleName** | **map[string]interface{}** | Application role assigned to members of this group | 
**Priority** | **map[string]interface{}** | Mapping priority — lower value wins when a user matches multiple groups | 
**QuotaProfileId** | **string** | Quota profile assigned to users matched by this group | 

## Methods

### NewGroupMapping

`func NewGroupMapping(groupName map[string]interface{}, roleName map[string]interface{}, priority map[string]interface{}, quotaProfileId string, ) *GroupMapping`

NewGroupMapping instantiates a new GroupMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupMappingWithDefaults

`func NewGroupMappingWithDefaults() *GroupMapping`

NewGroupMappingWithDefaults instantiates a new GroupMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupName

`func (o *GroupMapping) GetGroupName() map[string]interface{}`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *GroupMapping) GetGroupNameOk() (*map[string]interface{}, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *GroupMapping) SetGroupName(v map[string]interface{})`

SetGroupName sets GroupName field to given value.


### GetRoleName

`func (o *GroupMapping) GetRoleName() map[string]interface{}`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *GroupMapping) GetRoleNameOk() (*map[string]interface{}, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *GroupMapping) SetRoleName(v map[string]interface{})`

SetRoleName sets RoleName field to given value.


### GetPriority

`func (o *GroupMapping) GetPriority() map[string]interface{}`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *GroupMapping) GetPriorityOk() (*map[string]interface{}, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *GroupMapping) SetPriority(v map[string]interface{})`

SetPriority sets Priority field to given value.


### GetQuotaProfileId

`func (o *GroupMapping) GetQuotaProfileId() string`

GetQuotaProfileId returns the QuotaProfileId field if non-nil, zero value otherwise.

### GetQuotaProfileIdOk

`func (o *GroupMapping) GetQuotaProfileIdOk() (*string, bool)`

GetQuotaProfileIdOk returns a tuple with the QuotaProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaProfileId

`func (o *GroupMapping) SetQuotaProfileId(v string)`

SetQuotaProfileId sets QuotaProfileId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


