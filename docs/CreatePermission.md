# CreatePermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Permission name | 
**Label** | **string** | Permission label | 
**Description** | Pointer to **string** | Permission description | [optional] 
**Type** | Pointer to **string** | Permission type | [optional] 
**Policies** | Pointer to [**[]PolicyEntry**](PolicyEntry.md) | List of Policy Entries | [optional] 

## Methods

### NewCreatePermission

`func NewCreatePermission(name string, label string, ) *CreatePermission`

NewCreatePermission instantiates a new CreatePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePermissionWithDefaults

`func NewCreatePermissionWithDefaults() *CreatePermission`

NewCreatePermissionWithDefaults instantiates a new CreatePermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreatePermission) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePermission) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePermission) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *CreatePermission) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreatePermission) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreatePermission) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *CreatePermission) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatePermission) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatePermission) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreatePermission) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *CreatePermission) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreatePermission) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreatePermission) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreatePermission) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPolicies

`func (o *CreatePermission) GetPolicies() []PolicyEntry`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *CreatePermission) GetPoliciesOk() (*[]PolicyEntry, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *CreatePermission) SetPolicies(v []PolicyEntry)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *CreatePermission) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


