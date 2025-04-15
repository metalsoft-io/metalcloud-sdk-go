# Role

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Role ID | 
**Name** | **string** | Role name identifier | 
**Label** | **string** | Display label for the role | 
**Description** | Pointer to **string** | Role description | [optional] 
**Type** | **string** | Role type | 
**Permissions** | **[]string** | List of permissions assigned to the role | 
**UsersWithRole** | Pointer to **float32** | Number of users with this role | [optional] 

## Methods

### NewRole

`func NewRole(id string, name string, label string, type_ string, permissions []string, ) *Role`

NewRole instantiates a new Role object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleWithDefaults

`func NewRoleWithDefaults() *Role`

NewRoleWithDefaults instantiates a new Role object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Role) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Role) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Role) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Role) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Role) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Role) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *Role) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Role) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Role) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *Role) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Role) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Role) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Role) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *Role) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Role) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Role) SetType(v string)`

SetType sets Type field to given value.


### GetPermissions

`func (o *Role) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *Role) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *Role) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.


### GetUsersWithRole

`func (o *Role) GetUsersWithRole() float32`

GetUsersWithRole returns the UsersWithRole field if non-nil, zero value otherwise.

### GetUsersWithRoleOk

`func (o *Role) GetUsersWithRoleOk() (*float32, bool)`

GetUsersWithRoleOk returns a tuple with the UsersWithRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersWithRole

`func (o *Role) SetUsersWithRole(v float32)`

SetUsersWithRole sets UsersWithRole field to given value.

### HasUsersWithRole

`func (o *Role) HasUsersWithRole() bool`

HasUsersWithRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


