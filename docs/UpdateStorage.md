# UpdateStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InMaintenance** | Pointer to **float32** | Specifies if the storage is in maintenance | [optional] 
**IsExperimental** | Pointer to **float32** | Specifies if the storage is experimental | [optional] 
**DrivePriority** | Pointer to **float32** | Specifies the drive priority | [optional] 
**SharedDrivePriority** | Pointer to **float32** | Specifies the shared drive priority | [optional] 
**Tags** | Pointer to **[]string** | Tags | [optional] 
**Options** | Pointer to [**UpdateStorageOptions**](UpdateStorageOptions.md) | Options for the storage | [optional] 
**Password** | Pointer to **string** | The password to use. | [optional] 

## Methods

### NewUpdateStorage

`func NewUpdateStorage() *UpdateStorage`

NewUpdateStorage instantiates a new UpdateStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStorageWithDefaults

`func NewUpdateStorageWithDefaults() *UpdateStorage`

NewUpdateStorageWithDefaults instantiates a new UpdateStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInMaintenance

`func (o *UpdateStorage) GetInMaintenance() float32`

GetInMaintenance returns the InMaintenance field if non-nil, zero value otherwise.

### GetInMaintenanceOk

`func (o *UpdateStorage) GetInMaintenanceOk() (*float32, bool)`

GetInMaintenanceOk returns a tuple with the InMaintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenance

`func (o *UpdateStorage) SetInMaintenance(v float32)`

SetInMaintenance sets InMaintenance field to given value.

### HasInMaintenance

`func (o *UpdateStorage) HasInMaintenance() bool`

HasInMaintenance returns a boolean if a field has been set.

### GetIsExperimental

`func (o *UpdateStorage) GetIsExperimental() float32`

GetIsExperimental returns the IsExperimental field if non-nil, zero value otherwise.

### GetIsExperimentalOk

`func (o *UpdateStorage) GetIsExperimentalOk() (*float32, bool)`

GetIsExperimentalOk returns a tuple with the IsExperimental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExperimental

`func (o *UpdateStorage) SetIsExperimental(v float32)`

SetIsExperimental sets IsExperimental field to given value.

### HasIsExperimental

`func (o *UpdateStorage) HasIsExperimental() bool`

HasIsExperimental returns a boolean if a field has been set.

### GetDrivePriority

`func (o *UpdateStorage) GetDrivePriority() float32`

GetDrivePriority returns the DrivePriority field if non-nil, zero value otherwise.

### GetDrivePriorityOk

`func (o *UpdateStorage) GetDrivePriorityOk() (*float32, bool)`

GetDrivePriorityOk returns a tuple with the DrivePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivePriority

`func (o *UpdateStorage) SetDrivePriority(v float32)`

SetDrivePriority sets DrivePriority field to given value.

### HasDrivePriority

`func (o *UpdateStorage) HasDrivePriority() bool`

HasDrivePriority returns a boolean if a field has been set.

### GetSharedDrivePriority

`func (o *UpdateStorage) GetSharedDrivePriority() float32`

GetSharedDrivePriority returns the SharedDrivePriority field if non-nil, zero value otherwise.

### GetSharedDrivePriorityOk

`func (o *UpdateStorage) GetSharedDrivePriorityOk() (*float32, bool)`

GetSharedDrivePriorityOk returns a tuple with the SharedDrivePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedDrivePriority

`func (o *UpdateStorage) SetSharedDrivePriority(v float32)`

SetSharedDrivePriority sets SharedDrivePriority field to given value.

### HasSharedDrivePriority

`func (o *UpdateStorage) HasSharedDrivePriority() bool`

HasSharedDrivePriority returns a boolean if a field has been set.

### GetTags

`func (o *UpdateStorage) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateStorage) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateStorage) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateStorage) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOptions

`func (o *UpdateStorage) GetOptions() UpdateStorageOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *UpdateStorage) GetOptionsOk() (*UpdateStorageOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *UpdateStorage) SetOptions(v UpdateStorageOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *UpdateStorage) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateStorage) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateStorage) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateStorage) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateStorage) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


