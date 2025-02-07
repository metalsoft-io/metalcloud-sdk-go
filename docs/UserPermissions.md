# UserPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminPasswordReveal** | **map[string]interface{}** | Admin password reveal permissions | 
**SpecialPermissions** | **map[string]interface{}** | Special permissions | 

## Methods

### NewUserPermissions

`func NewUserPermissions(adminPasswordReveal map[string]interface{}, specialPermissions map[string]interface{}, ) *UserPermissions`

NewUserPermissions instantiates a new UserPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPermissionsWithDefaults

`func NewUserPermissionsWithDefaults() *UserPermissions`

NewUserPermissionsWithDefaults instantiates a new UserPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminPasswordReveal

`func (o *UserPermissions) GetAdminPasswordReveal() map[string]interface{}`

GetAdminPasswordReveal returns the AdminPasswordReveal field if non-nil, zero value otherwise.

### GetAdminPasswordRevealOk

`func (o *UserPermissions) GetAdminPasswordRevealOk() (*map[string]interface{}, bool)`

GetAdminPasswordRevealOk returns a tuple with the AdminPasswordReveal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPasswordReveal

`func (o *UserPermissions) SetAdminPasswordReveal(v map[string]interface{})`

SetAdminPasswordReveal sets AdminPasswordReveal field to given value.


### GetSpecialPermissions

`func (o *UserPermissions) GetSpecialPermissions() map[string]interface{}`

GetSpecialPermissions returns the SpecialPermissions field if non-nil, zero value otherwise.

### GetSpecialPermissionsOk

`func (o *UserPermissions) GetSpecialPermissionsOk() (*map[string]interface{}, bool)`

GetSpecialPermissionsOk returns a tuple with the SpecialPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialPermissions

`func (o *UserPermissions) SetSpecialPermissions(v map[string]interface{})`

SetSpecialPermissions sets SpecialPermissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


