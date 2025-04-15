# PermissionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | [**[]Permission**](Permission.md) | List of permissions | 

## Methods

### NewPermissionList

`func NewPermissionList(permissions []Permission, ) *PermissionList`

NewPermissionList instantiates a new PermissionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionListWithDefaults

`func NewPermissionListWithDefaults() *PermissionList`

NewPermissionListWithDefaults instantiates a new PermissionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *PermissionList) GetPermissions() []Permission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PermissionList) GetPermissionsOk() (*[]Permission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PermissionList) SetPermissions(v []Permission)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


