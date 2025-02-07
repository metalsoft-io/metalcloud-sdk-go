# UserSpecialPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpecialPermission** | **string** | The resource permission for the user | 
**EnablePermission** | Pointer to **bool** | Whether to enable the permission | [optional] 

## Methods

### NewUserSpecialPermission

`func NewUserSpecialPermission(specialPermission string, ) *UserSpecialPermission`

NewUserSpecialPermission instantiates a new UserSpecialPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSpecialPermissionWithDefaults

`func NewUserSpecialPermissionWithDefaults() *UserSpecialPermission`

NewUserSpecialPermissionWithDefaults instantiates a new UserSpecialPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpecialPermission

`func (o *UserSpecialPermission) GetSpecialPermission() string`

GetSpecialPermission returns the SpecialPermission field if non-nil, zero value otherwise.

### GetSpecialPermissionOk

`func (o *UserSpecialPermission) GetSpecialPermissionOk() (*string, bool)`

GetSpecialPermissionOk returns a tuple with the SpecialPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialPermission

`func (o *UserSpecialPermission) SetSpecialPermission(v string)`

SetSpecialPermission sets SpecialPermission field to given value.


### GetEnablePermission

`func (o *UserSpecialPermission) GetEnablePermission() bool`

GetEnablePermission returns the EnablePermission field if non-nil, zero value otherwise.

### GetEnablePermissionOk

`func (o *UserSpecialPermission) GetEnablePermissionOk() (*bool, bool)`

GetEnablePermissionOk returns a tuple with the EnablePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePermission

`func (o *UserSpecialPermission) SetEnablePermission(v bool)`

SetEnablePermission sets EnablePermission field to given value.

### HasEnablePermission

`func (o *UserSpecialPermission) HasEnablePermission() bool`

HasEnablePermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


