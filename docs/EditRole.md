# EditRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Role label. Omit to keep the current value. Cannot be changed on built-in roles. | [optional] 
**Description** | Pointer to **string** | Role description. Omit to keep the current value. Cannot be changed on built-in roles. | [optional] 
**Permissions** | Pointer to [**[]MetalsoftPermissions**](MetalsoftPermissions.md) | List of permissions assigned to the role. Omit to keep the current value. Cannot be changed on built-in roles. | [optional] 
**QuotaProfileId** | Pointer to **string** | Quota profile assigned to this role | [optional] 

## Methods

### NewEditRole

`func NewEditRole() *EditRole`

NewEditRole instantiates a new EditRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditRoleWithDefaults

`func NewEditRoleWithDefaults() *EditRole`

NewEditRoleWithDefaults instantiates a new EditRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *EditRole) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *EditRole) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *EditRole) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *EditRole) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *EditRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EditRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EditRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EditRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPermissions

`func (o *EditRole) GetPermissions() []MetalsoftPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *EditRole) GetPermissionsOk() (*[]MetalsoftPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *EditRole) SetPermissions(v []MetalsoftPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *EditRole) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetQuotaProfileId

`func (o *EditRole) GetQuotaProfileId() string`

GetQuotaProfileId returns the QuotaProfileId field if non-nil, zero value otherwise.

### GetQuotaProfileIdOk

`func (o *EditRole) GetQuotaProfileIdOk() (*string, bool)`

GetQuotaProfileIdOk returns a tuple with the QuotaProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaProfileId

`func (o *EditRole) SetQuotaProfileId(v string)`

SetQuotaProfileId sets QuotaProfileId field to given value.

### HasQuotaProfileId

`func (o *EditRole) HasQuotaProfileId() bool`

HasQuotaProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


