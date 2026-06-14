# UserPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | **[]string** | List of permissions of the logged in user | 

## Methods

### NewUserPermissions

`func NewUserPermissions(permissions []string, ) *UserPermissions`

NewUserPermissions instantiates a new UserPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPermissionsWithDefaults

`func NewUserPermissionsWithDefaults() *UserPermissions`

NewUserPermissionsWithDefaults instantiates a new UserPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *UserPermissions) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UserPermissions) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UserPermissions) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


