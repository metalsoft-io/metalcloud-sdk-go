# UpdateVMPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the VM Pool | [optional] 
**InMaintenance** | Pointer to **float32** | Flag to indicate if the VM Pool is in maintenance mode. 1 for true, 0 for false. Default is 0. | [optional] 
**IsExperimental** | Pointer to **float32** | Flag to indicate if the VM Pool is experimental. 1 for true, 0 for false. Default is 0. | [optional] 
**Tags** | Pointer to **[]string** | Tags for the VM Pool. | [optional] 
**ManagementHost** | Pointer to **string** | Host of the VM Pool | [optional] 
**ManagementPort** | Pointer to **float32** | Port of the VM Pool | [optional] 
**Certificate** | Pointer to **string** | Certificate of the VM Pool | [optional] 
**PrivateKey** | Pointer to **string** | Private key of the VM Pool | [optional] 
**Username** | Pointer to **string** | Username of the VM Pool | [optional] 
**Password** | Pointer to **string** | Password of the VM Pool | [optional] 
**Options** | Pointer to [**UpdateVMPoolOptions**](UpdateVMPoolOptions.md) | Options for the VM Pool | [optional] 

## Methods

### NewUpdateVMPool

`func NewUpdateVMPool() *UpdateVMPool`

NewUpdateVMPool instantiates a new UpdateVMPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVMPoolWithDefaults

`func NewUpdateVMPoolWithDefaults() *UpdateVMPool`

NewUpdateVMPoolWithDefaults instantiates a new UpdateVMPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateVMPool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateVMPool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateVMPool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateVMPool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInMaintenance

`func (o *UpdateVMPool) GetInMaintenance() float32`

GetInMaintenance returns the InMaintenance field if non-nil, zero value otherwise.

### GetInMaintenanceOk

`func (o *UpdateVMPool) GetInMaintenanceOk() (*float32, bool)`

GetInMaintenanceOk returns a tuple with the InMaintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenance

`func (o *UpdateVMPool) SetInMaintenance(v float32)`

SetInMaintenance sets InMaintenance field to given value.

### HasInMaintenance

`func (o *UpdateVMPool) HasInMaintenance() bool`

HasInMaintenance returns a boolean if a field has been set.

### GetIsExperimental

`func (o *UpdateVMPool) GetIsExperimental() float32`

GetIsExperimental returns the IsExperimental field if non-nil, zero value otherwise.

### GetIsExperimentalOk

`func (o *UpdateVMPool) GetIsExperimentalOk() (*float32, bool)`

GetIsExperimentalOk returns a tuple with the IsExperimental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExperimental

`func (o *UpdateVMPool) SetIsExperimental(v float32)`

SetIsExperimental sets IsExperimental field to given value.

### HasIsExperimental

`func (o *UpdateVMPool) HasIsExperimental() bool`

HasIsExperimental returns a boolean if a field has been set.

### GetTags

`func (o *UpdateVMPool) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateVMPool) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateVMPool) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateVMPool) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetManagementHost

`func (o *UpdateVMPool) GetManagementHost() string`

GetManagementHost returns the ManagementHost field if non-nil, zero value otherwise.

### GetManagementHostOk

`func (o *UpdateVMPool) GetManagementHostOk() (*string, bool)`

GetManagementHostOk returns a tuple with the ManagementHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementHost

`func (o *UpdateVMPool) SetManagementHost(v string)`

SetManagementHost sets ManagementHost field to given value.

### HasManagementHost

`func (o *UpdateVMPool) HasManagementHost() bool`

HasManagementHost returns a boolean if a field has been set.

### GetManagementPort

`func (o *UpdateVMPool) GetManagementPort() float32`

GetManagementPort returns the ManagementPort field if non-nil, zero value otherwise.

### GetManagementPortOk

`func (o *UpdateVMPool) GetManagementPortOk() (*float32, bool)`

GetManagementPortOk returns a tuple with the ManagementPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementPort

`func (o *UpdateVMPool) SetManagementPort(v float32)`

SetManagementPort sets ManagementPort field to given value.

### HasManagementPort

`func (o *UpdateVMPool) HasManagementPort() bool`

HasManagementPort returns a boolean if a field has been set.

### GetCertificate

`func (o *UpdateVMPool) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *UpdateVMPool) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *UpdateVMPool) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *UpdateVMPool) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetPrivateKey

`func (o *UpdateVMPool) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *UpdateVMPool) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *UpdateVMPool) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *UpdateVMPool) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetUsername

`func (o *UpdateVMPool) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateVMPool) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateVMPool) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateVMPool) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateVMPool) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateVMPool) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateVMPool) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateVMPool) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetOptions

`func (o *UpdateVMPool) GetOptions() UpdateVMPoolOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *UpdateVMPool) GetOptionsOk() (*UpdateVMPoolOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *UpdateVMPool) SetOptions(v UpdateVMPoolOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *UpdateVMPool) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


