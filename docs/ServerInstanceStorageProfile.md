# ServerInstanceStorageProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Controllers** | Pointer to [**[]ServerInstanceStorageProfileController**](ServerInstanceStorageProfileController.md) | The information for each Storage Controller of the Instance Storage Profile. | [optional] 

## Methods

### NewServerInstanceStorageProfile

`func NewServerInstanceStorageProfile() *ServerInstanceStorageProfile`

NewServerInstanceStorageProfile instantiates a new ServerInstanceStorageProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceStorageProfileWithDefaults

`func NewServerInstanceStorageProfileWithDefaults() *ServerInstanceStorageProfile`

NewServerInstanceStorageProfileWithDefaults instantiates a new ServerInstanceStorageProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControllers

`func (o *ServerInstanceStorageProfile) GetControllers() []ServerInstanceStorageProfileController`

GetControllers returns the Controllers field if non-nil, zero value otherwise.

### GetControllersOk

`func (o *ServerInstanceStorageProfile) GetControllersOk() (*[]ServerInstanceStorageProfileController, bool)`

GetControllersOk returns a tuple with the Controllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllers

`func (o *ServerInstanceStorageProfile) SetControllers(v []ServerInstanceStorageProfileController)`

SetControllers sets Controllers field to given value.

### HasControllers

`func (o *ServerInstanceStorageProfile) HasControllers() bool`

HasControllers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


