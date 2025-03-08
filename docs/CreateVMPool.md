# CreateVMPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | **float32** | Id of the site for the VM | 
**ManagementHost** | **string** | Host of the VM Pool | 
**ManagementPort** | **float32** | Port of the VM Pool | 
**Name** | **string** | Name of the VM Pool | 
**Description** | Pointer to **string** | Description of the VM Pool | [optional] 
**Type** | **string** | Type of the VM Pool | 
**Certificate** | **string** | Certificate of the VM Pool | 
**InMaintenance** | Pointer to **float32** | Flag to indicate if the VM Pool is in maintenance mode. 1 for true, 0 for false. Default is 0. | [optional] 
**IsExperimental** | Pointer to **float32** | Flag to indicate if the VM Pool is experimental. 1 for true, 0 for false. Default is 0. | [optional] 
**Tags** | Pointer to **[]string** | Tags for the VM Pool. | [optional] 
**PrivateKey** | **string** | Private key of the VM Pool | 

## Methods

### NewCreateVMPool

`func NewCreateVMPool(siteId float32, managementHost string, managementPort float32, name string, type_ string, certificate string, privateKey string, ) *CreateVMPool`

NewCreateVMPool instantiates a new CreateVMPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVMPoolWithDefaults

`func NewCreateVMPoolWithDefaults() *CreateVMPool`

NewCreateVMPoolWithDefaults instantiates a new CreateVMPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *CreateVMPool) GetSiteId() float32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *CreateVMPool) GetSiteIdOk() (*float32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *CreateVMPool) SetSiteId(v float32)`

SetSiteId sets SiteId field to given value.


### GetManagementHost

`func (o *CreateVMPool) GetManagementHost() string`

GetManagementHost returns the ManagementHost field if non-nil, zero value otherwise.

### GetManagementHostOk

`func (o *CreateVMPool) GetManagementHostOk() (*string, bool)`

GetManagementHostOk returns a tuple with the ManagementHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementHost

`func (o *CreateVMPool) SetManagementHost(v string)`

SetManagementHost sets ManagementHost field to given value.


### GetManagementPort

`func (o *CreateVMPool) GetManagementPort() float32`

GetManagementPort returns the ManagementPort field if non-nil, zero value otherwise.

### GetManagementPortOk

`func (o *CreateVMPool) GetManagementPortOk() (*float32, bool)`

GetManagementPortOk returns a tuple with the ManagementPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementPort

`func (o *CreateVMPool) SetManagementPort(v float32)`

SetManagementPort sets ManagementPort field to given value.


### GetName

`func (o *CreateVMPool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateVMPool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateVMPool) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateVMPool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateVMPool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateVMPool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateVMPool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *CreateVMPool) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateVMPool) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateVMPool) SetType(v string)`

SetType sets Type field to given value.


### GetCertificate

`func (o *CreateVMPool) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CreateVMPool) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CreateVMPool) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetInMaintenance

`func (o *CreateVMPool) GetInMaintenance() float32`

GetInMaintenance returns the InMaintenance field if non-nil, zero value otherwise.

### GetInMaintenanceOk

`func (o *CreateVMPool) GetInMaintenanceOk() (*float32, bool)`

GetInMaintenanceOk returns a tuple with the InMaintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenance

`func (o *CreateVMPool) SetInMaintenance(v float32)`

SetInMaintenance sets InMaintenance field to given value.

### HasInMaintenance

`func (o *CreateVMPool) HasInMaintenance() bool`

HasInMaintenance returns a boolean if a field has been set.

### GetIsExperimental

`func (o *CreateVMPool) GetIsExperimental() float32`

GetIsExperimental returns the IsExperimental field if non-nil, zero value otherwise.

### GetIsExperimentalOk

`func (o *CreateVMPool) GetIsExperimentalOk() (*float32, bool)`

GetIsExperimentalOk returns a tuple with the IsExperimental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExperimental

`func (o *CreateVMPool) SetIsExperimental(v float32)`

SetIsExperimental sets IsExperimental field to given value.

### HasIsExperimental

`func (o *CreateVMPool) HasIsExperimental() bool`

HasIsExperimental returns a boolean if a field has been set.

### GetTags

`func (o *CreateVMPool) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateVMPool) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateVMPool) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateVMPool) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPrivateKey

`func (o *CreateVMPool) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *CreateVMPool) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *CreateVMPool) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


