# PlatformConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MailClient** | Pointer to [**PlatformMailClient**](PlatformMailClient.md) | Outbound email (SMTP) settings | [optional] 
**PowerDNS** | Pointer to [**PlatformPowerDNS**](PlatformPowerDNS.md) | PowerDNS integration settings | [optional] 
**BSIAdminURLRoot** | Pointer to **string** | Base URL for the platform admin interface | [optional] 
**RepoURLRootMaster** | Pointer to **string** | Base URL for the asset repository | [optional] 
**Uploads** | Pointer to [**PlatformUploads**](PlatformUploads.md) | File upload storage path settings | [optional] 
**StorageConfig** | Pointer to [**PlatformStorageConfig**](PlatformStorageConfig.md) | Storage overprovisioning settings | [optional] 
**AllowedPrefixSizesOnWAN** | Pointer to [**PlatformAllowedPrefixSizesOnWAN**](PlatformAllowedPrefixSizesOnWAN.md) | Allowed IP prefix sizes for WAN allocation | [optional] 
**CookieDomain** | Pointer to **string** | Domain scope for session cookies | [optional] 
**AllowTFTPBootThroughWAN** | Pointer to **bool** | Whether servers are allowed to TFTP-boot through the WAN interface | [optional] 
**AllowServersWithOneInterface** | Pointer to **bool** | Whether single-interface servers are permitted in the platform | [optional] 
**AFC** | Pointer to [**PlatformAFC**](PlatformAFC.md) | Automated Fulfillment Center notification settings | [optional] 
**DisablePublicUserSignup** | Pointer to **bool** | Whether self-service user registration is disabled | [optional] 

## Methods

### NewPlatformConfiguration

`func NewPlatformConfiguration() *PlatformConfiguration`

NewPlatformConfiguration instantiates a new PlatformConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformConfigurationWithDefaults

`func NewPlatformConfigurationWithDefaults() *PlatformConfiguration`

NewPlatformConfigurationWithDefaults instantiates a new PlatformConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMailClient

`func (o *PlatformConfiguration) GetMailClient() PlatformMailClient`

GetMailClient returns the MailClient field if non-nil, zero value otherwise.

### GetMailClientOk

`func (o *PlatformConfiguration) GetMailClientOk() (*PlatformMailClient, bool)`

GetMailClientOk returns a tuple with the MailClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailClient

`func (o *PlatformConfiguration) SetMailClient(v PlatformMailClient)`

SetMailClient sets MailClient field to given value.

### HasMailClient

`func (o *PlatformConfiguration) HasMailClient() bool`

HasMailClient returns a boolean if a field has been set.

### GetPowerDNS

`func (o *PlatformConfiguration) GetPowerDNS() PlatformPowerDNS`

GetPowerDNS returns the PowerDNS field if non-nil, zero value otherwise.

### GetPowerDNSOk

`func (o *PlatformConfiguration) GetPowerDNSOk() (*PlatformPowerDNS, bool)`

GetPowerDNSOk returns a tuple with the PowerDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerDNS

`func (o *PlatformConfiguration) SetPowerDNS(v PlatformPowerDNS)`

SetPowerDNS sets PowerDNS field to given value.

### HasPowerDNS

`func (o *PlatformConfiguration) HasPowerDNS() bool`

HasPowerDNS returns a boolean if a field has been set.

### GetBSIAdminURLRoot

`func (o *PlatformConfiguration) GetBSIAdminURLRoot() string`

GetBSIAdminURLRoot returns the BSIAdminURLRoot field if non-nil, zero value otherwise.

### GetBSIAdminURLRootOk

`func (o *PlatformConfiguration) GetBSIAdminURLRootOk() (*string, bool)`

GetBSIAdminURLRootOk returns a tuple with the BSIAdminURLRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBSIAdminURLRoot

`func (o *PlatformConfiguration) SetBSIAdminURLRoot(v string)`

SetBSIAdminURLRoot sets BSIAdminURLRoot field to given value.

### HasBSIAdminURLRoot

`func (o *PlatformConfiguration) HasBSIAdminURLRoot() bool`

HasBSIAdminURLRoot returns a boolean if a field has been set.

### GetRepoURLRootMaster

`func (o *PlatformConfiguration) GetRepoURLRootMaster() string`

GetRepoURLRootMaster returns the RepoURLRootMaster field if non-nil, zero value otherwise.

### GetRepoURLRootMasterOk

`func (o *PlatformConfiguration) GetRepoURLRootMasterOk() (*string, bool)`

GetRepoURLRootMasterOk returns a tuple with the RepoURLRootMaster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoURLRootMaster

`func (o *PlatformConfiguration) SetRepoURLRootMaster(v string)`

SetRepoURLRootMaster sets RepoURLRootMaster field to given value.

### HasRepoURLRootMaster

`func (o *PlatformConfiguration) HasRepoURLRootMaster() bool`

HasRepoURLRootMaster returns a boolean if a field has been set.

### GetUploads

`func (o *PlatformConfiguration) GetUploads() PlatformUploads`

GetUploads returns the Uploads field if non-nil, zero value otherwise.

### GetUploadsOk

`func (o *PlatformConfiguration) GetUploadsOk() (*PlatformUploads, bool)`

GetUploadsOk returns a tuple with the Uploads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploads

`func (o *PlatformConfiguration) SetUploads(v PlatformUploads)`

SetUploads sets Uploads field to given value.

### HasUploads

`func (o *PlatformConfiguration) HasUploads() bool`

HasUploads returns a boolean if a field has been set.

### GetStorageConfig

`func (o *PlatformConfiguration) GetStorageConfig() PlatformStorageConfig`

GetStorageConfig returns the StorageConfig field if non-nil, zero value otherwise.

### GetStorageConfigOk

`func (o *PlatformConfiguration) GetStorageConfigOk() (*PlatformStorageConfig, bool)`

GetStorageConfigOk returns a tuple with the StorageConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfig

`func (o *PlatformConfiguration) SetStorageConfig(v PlatformStorageConfig)`

SetStorageConfig sets StorageConfig field to given value.

### HasStorageConfig

`func (o *PlatformConfiguration) HasStorageConfig() bool`

HasStorageConfig returns a boolean if a field has been set.

### GetAllowedPrefixSizesOnWAN

`func (o *PlatformConfiguration) GetAllowedPrefixSizesOnWAN() PlatformAllowedPrefixSizesOnWAN`

GetAllowedPrefixSizesOnWAN returns the AllowedPrefixSizesOnWAN field if non-nil, zero value otherwise.

### GetAllowedPrefixSizesOnWANOk

`func (o *PlatformConfiguration) GetAllowedPrefixSizesOnWANOk() (*PlatformAllowedPrefixSizesOnWAN, bool)`

GetAllowedPrefixSizesOnWANOk returns a tuple with the AllowedPrefixSizesOnWAN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPrefixSizesOnWAN

`func (o *PlatformConfiguration) SetAllowedPrefixSizesOnWAN(v PlatformAllowedPrefixSizesOnWAN)`

SetAllowedPrefixSizesOnWAN sets AllowedPrefixSizesOnWAN field to given value.

### HasAllowedPrefixSizesOnWAN

`func (o *PlatformConfiguration) HasAllowedPrefixSizesOnWAN() bool`

HasAllowedPrefixSizesOnWAN returns a boolean if a field has been set.

### GetCookieDomain

`func (o *PlatformConfiguration) GetCookieDomain() string`

GetCookieDomain returns the CookieDomain field if non-nil, zero value otherwise.

### GetCookieDomainOk

`func (o *PlatformConfiguration) GetCookieDomainOk() (*string, bool)`

GetCookieDomainOk returns a tuple with the CookieDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieDomain

`func (o *PlatformConfiguration) SetCookieDomain(v string)`

SetCookieDomain sets CookieDomain field to given value.

### HasCookieDomain

`func (o *PlatformConfiguration) HasCookieDomain() bool`

HasCookieDomain returns a boolean if a field has been set.

### GetAllowTFTPBootThroughWAN

`func (o *PlatformConfiguration) GetAllowTFTPBootThroughWAN() bool`

GetAllowTFTPBootThroughWAN returns the AllowTFTPBootThroughWAN field if non-nil, zero value otherwise.

### GetAllowTFTPBootThroughWANOk

`func (o *PlatformConfiguration) GetAllowTFTPBootThroughWANOk() (*bool, bool)`

GetAllowTFTPBootThroughWANOk returns a tuple with the AllowTFTPBootThroughWAN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTFTPBootThroughWAN

`func (o *PlatformConfiguration) SetAllowTFTPBootThroughWAN(v bool)`

SetAllowTFTPBootThroughWAN sets AllowTFTPBootThroughWAN field to given value.

### HasAllowTFTPBootThroughWAN

`func (o *PlatformConfiguration) HasAllowTFTPBootThroughWAN() bool`

HasAllowTFTPBootThroughWAN returns a boolean if a field has been set.

### GetAllowServersWithOneInterface

`func (o *PlatformConfiguration) GetAllowServersWithOneInterface() bool`

GetAllowServersWithOneInterface returns the AllowServersWithOneInterface field if non-nil, zero value otherwise.

### GetAllowServersWithOneInterfaceOk

`func (o *PlatformConfiguration) GetAllowServersWithOneInterfaceOk() (*bool, bool)`

GetAllowServersWithOneInterfaceOk returns a tuple with the AllowServersWithOneInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowServersWithOneInterface

`func (o *PlatformConfiguration) SetAllowServersWithOneInterface(v bool)`

SetAllowServersWithOneInterface sets AllowServersWithOneInterface field to given value.

### HasAllowServersWithOneInterface

`func (o *PlatformConfiguration) HasAllowServersWithOneInterface() bool`

HasAllowServersWithOneInterface returns a boolean if a field has been set.

### GetAFC

`func (o *PlatformConfiguration) GetAFC() PlatformAFC`

GetAFC returns the AFC field if non-nil, zero value otherwise.

### GetAFCOk

`func (o *PlatformConfiguration) GetAFCOk() (*PlatformAFC, bool)`

GetAFCOk returns a tuple with the AFC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAFC

`func (o *PlatformConfiguration) SetAFC(v PlatformAFC)`

SetAFC sets AFC field to given value.

### HasAFC

`func (o *PlatformConfiguration) HasAFC() bool`

HasAFC returns a boolean if a field has been set.

### GetDisablePublicUserSignup

`func (o *PlatformConfiguration) GetDisablePublicUserSignup() bool`

GetDisablePublicUserSignup returns the DisablePublicUserSignup field if non-nil, zero value otherwise.

### GetDisablePublicUserSignupOk

`func (o *PlatformConfiguration) GetDisablePublicUserSignupOk() (*bool, bool)`

GetDisablePublicUserSignupOk returns a tuple with the DisablePublicUserSignup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisablePublicUserSignup

`func (o *PlatformConfiguration) SetDisablePublicUserSignup(v bool)`

SetDisablePublicUserSignup sets DisablePublicUserSignup field to given value.

### HasDisablePublicUserSignup

`func (o *PlatformConfiguration) HasDisablePublicUserSignup() bool`

HasDisablePublicUserSignup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


