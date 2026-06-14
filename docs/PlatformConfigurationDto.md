# PlatformConfigurationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MailClient** | Pointer to [**PlatformMailClientDto**](PlatformMailClientDto.md) | Outbound email (SMTP) settings | [optional] 
**PowerDNS** | Pointer to [**PlatformPowerDNSDto**](PlatformPowerDNSDto.md) | PowerDNS integration settings | [optional] 
**BSIAdminURLRoot** | Pointer to **string** | Base URL for the platform admin interface | [optional] 
**RepoURLRootMaster** | Pointer to **string** | Base URL for the asset repository | [optional] 
**Uploads** | Pointer to [**PlatformUploadsDto**](PlatformUploadsDto.md) | File upload storage path settings | [optional] 
**StorageConfig** | Pointer to [**PlatformStorageConfigDto**](PlatformStorageConfigDto.md) | Storage overprovisioning settings | [optional] 
**AllowedPrefixSizesOnWAN** | Pointer to [**PlatformAllowedPrefixSizesOnWANDto**](PlatformAllowedPrefixSizesOnWANDto.md) | Allowed IP prefix sizes for WAN allocation | [optional] 
**CookieDomain** | Pointer to **string** | Domain scope for session cookies | [optional] 
**AllowTFTPBootThroughWAN** | Pointer to **bool** | Whether servers are allowed to TFTP-boot through the WAN interface | [optional] 
**AllowServersWithOneInterface** | Pointer to **bool** | Whether single-interface servers are permitted in the platform | [optional] 
**AFC** | Pointer to [**PlatformAFCDto**](PlatformAFCDto.md) | Automated Fulfillment Center notification settings | [optional] 
**DisablePublicUserSignup** | Pointer to **bool** | Whether self-service user registration is disabled | [optional] 

## Methods

### NewPlatformConfigurationDto

`func NewPlatformConfigurationDto() *PlatformConfigurationDto`

NewPlatformConfigurationDto instantiates a new PlatformConfigurationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformConfigurationDtoWithDefaults

`func NewPlatformConfigurationDtoWithDefaults() *PlatformConfigurationDto`

NewPlatformConfigurationDtoWithDefaults instantiates a new PlatformConfigurationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMailClient

`func (o *PlatformConfigurationDto) GetMailClient() PlatformMailClientDto`

GetMailClient returns the MailClient field if non-nil, zero value otherwise.

### GetMailClientOk

`func (o *PlatformConfigurationDto) GetMailClientOk() (*PlatformMailClientDto, bool)`

GetMailClientOk returns a tuple with the MailClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailClient

`func (o *PlatformConfigurationDto) SetMailClient(v PlatformMailClientDto)`

SetMailClient sets MailClient field to given value.

### HasMailClient

`func (o *PlatformConfigurationDto) HasMailClient() bool`

HasMailClient returns a boolean if a field has been set.

### GetPowerDNS

`func (o *PlatformConfigurationDto) GetPowerDNS() PlatformPowerDNSDto`

GetPowerDNS returns the PowerDNS field if non-nil, zero value otherwise.

### GetPowerDNSOk

`func (o *PlatformConfigurationDto) GetPowerDNSOk() (*PlatformPowerDNSDto, bool)`

GetPowerDNSOk returns a tuple with the PowerDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerDNS

`func (o *PlatformConfigurationDto) SetPowerDNS(v PlatformPowerDNSDto)`

SetPowerDNS sets PowerDNS field to given value.

### HasPowerDNS

`func (o *PlatformConfigurationDto) HasPowerDNS() bool`

HasPowerDNS returns a boolean if a field has been set.

### GetBSIAdminURLRoot

`func (o *PlatformConfigurationDto) GetBSIAdminURLRoot() string`

GetBSIAdminURLRoot returns the BSIAdminURLRoot field if non-nil, zero value otherwise.

### GetBSIAdminURLRootOk

`func (o *PlatformConfigurationDto) GetBSIAdminURLRootOk() (*string, bool)`

GetBSIAdminURLRootOk returns a tuple with the BSIAdminURLRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBSIAdminURLRoot

`func (o *PlatformConfigurationDto) SetBSIAdminURLRoot(v string)`

SetBSIAdminURLRoot sets BSIAdminURLRoot field to given value.

### HasBSIAdminURLRoot

`func (o *PlatformConfigurationDto) HasBSIAdminURLRoot() bool`

HasBSIAdminURLRoot returns a boolean if a field has been set.

### GetRepoURLRootMaster

`func (o *PlatformConfigurationDto) GetRepoURLRootMaster() string`

GetRepoURLRootMaster returns the RepoURLRootMaster field if non-nil, zero value otherwise.

### GetRepoURLRootMasterOk

`func (o *PlatformConfigurationDto) GetRepoURLRootMasterOk() (*string, bool)`

GetRepoURLRootMasterOk returns a tuple with the RepoURLRootMaster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoURLRootMaster

`func (o *PlatformConfigurationDto) SetRepoURLRootMaster(v string)`

SetRepoURLRootMaster sets RepoURLRootMaster field to given value.

### HasRepoURLRootMaster

`func (o *PlatformConfigurationDto) HasRepoURLRootMaster() bool`

HasRepoURLRootMaster returns a boolean if a field has been set.

### GetUploads

`func (o *PlatformConfigurationDto) GetUploads() PlatformUploadsDto`

GetUploads returns the Uploads field if non-nil, zero value otherwise.

### GetUploadsOk

`func (o *PlatformConfigurationDto) GetUploadsOk() (*PlatformUploadsDto, bool)`

GetUploadsOk returns a tuple with the Uploads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploads

`func (o *PlatformConfigurationDto) SetUploads(v PlatformUploadsDto)`

SetUploads sets Uploads field to given value.

### HasUploads

`func (o *PlatformConfigurationDto) HasUploads() bool`

HasUploads returns a boolean if a field has been set.

### GetStorageConfig

`func (o *PlatformConfigurationDto) GetStorageConfig() PlatformStorageConfigDto`

GetStorageConfig returns the StorageConfig field if non-nil, zero value otherwise.

### GetStorageConfigOk

`func (o *PlatformConfigurationDto) GetStorageConfigOk() (*PlatformStorageConfigDto, bool)`

GetStorageConfigOk returns a tuple with the StorageConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfig

`func (o *PlatformConfigurationDto) SetStorageConfig(v PlatformStorageConfigDto)`

SetStorageConfig sets StorageConfig field to given value.

### HasStorageConfig

`func (o *PlatformConfigurationDto) HasStorageConfig() bool`

HasStorageConfig returns a boolean if a field has been set.

### GetAllowedPrefixSizesOnWAN

`func (o *PlatformConfigurationDto) GetAllowedPrefixSizesOnWAN() PlatformAllowedPrefixSizesOnWANDto`

GetAllowedPrefixSizesOnWAN returns the AllowedPrefixSizesOnWAN field if non-nil, zero value otherwise.

### GetAllowedPrefixSizesOnWANOk

`func (o *PlatformConfigurationDto) GetAllowedPrefixSizesOnWANOk() (*PlatformAllowedPrefixSizesOnWANDto, bool)`

GetAllowedPrefixSizesOnWANOk returns a tuple with the AllowedPrefixSizesOnWAN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPrefixSizesOnWAN

`func (o *PlatformConfigurationDto) SetAllowedPrefixSizesOnWAN(v PlatformAllowedPrefixSizesOnWANDto)`

SetAllowedPrefixSizesOnWAN sets AllowedPrefixSizesOnWAN field to given value.

### HasAllowedPrefixSizesOnWAN

`func (o *PlatformConfigurationDto) HasAllowedPrefixSizesOnWAN() bool`

HasAllowedPrefixSizesOnWAN returns a boolean if a field has been set.

### GetCookieDomain

`func (o *PlatformConfigurationDto) GetCookieDomain() string`

GetCookieDomain returns the CookieDomain field if non-nil, zero value otherwise.

### GetCookieDomainOk

`func (o *PlatformConfigurationDto) GetCookieDomainOk() (*string, bool)`

GetCookieDomainOk returns a tuple with the CookieDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieDomain

`func (o *PlatformConfigurationDto) SetCookieDomain(v string)`

SetCookieDomain sets CookieDomain field to given value.

### HasCookieDomain

`func (o *PlatformConfigurationDto) HasCookieDomain() bool`

HasCookieDomain returns a boolean if a field has been set.

### GetAllowTFTPBootThroughWAN

`func (o *PlatformConfigurationDto) GetAllowTFTPBootThroughWAN() bool`

GetAllowTFTPBootThroughWAN returns the AllowTFTPBootThroughWAN field if non-nil, zero value otherwise.

### GetAllowTFTPBootThroughWANOk

`func (o *PlatformConfigurationDto) GetAllowTFTPBootThroughWANOk() (*bool, bool)`

GetAllowTFTPBootThroughWANOk returns a tuple with the AllowTFTPBootThroughWAN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTFTPBootThroughWAN

`func (o *PlatformConfigurationDto) SetAllowTFTPBootThroughWAN(v bool)`

SetAllowTFTPBootThroughWAN sets AllowTFTPBootThroughWAN field to given value.

### HasAllowTFTPBootThroughWAN

`func (o *PlatformConfigurationDto) HasAllowTFTPBootThroughWAN() bool`

HasAllowTFTPBootThroughWAN returns a boolean if a field has been set.

### GetAllowServersWithOneInterface

`func (o *PlatformConfigurationDto) GetAllowServersWithOneInterface() bool`

GetAllowServersWithOneInterface returns the AllowServersWithOneInterface field if non-nil, zero value otherwise.

### GetAllowServersWithOneInterfaceOk

`func (o *PlatformConfigurationDto) GetAllowServersWithOneInterfaceOk() (*bool, bool)`

GetAllowServersWithOneInterfaceOk returns a tuple with the AllowServersWithOneInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowServersWithOneInterface

`func (o *PlatformConfigurationDto) SetAllowServersWithOneInterface(v bool)`

SetAllowServersWithOneInterface sets AllowServersWithOneInterface field to given value.

### HasAllowServersWithOneInterface

`func (o *PlatformConfigurationDto) HasAllowServersWithOneInterface() bool`

HasAllowServersWithOneInterface returns a boolean if a field has been set.

### GetAFC

`func (o *PlatformConfigurationDto) GetAFC() PlatformAFCDto`

GetAFC returns the AFC field if non-nil, zero value otherwise.

### GetAFCOk

`func (o *PlatformConfigurationDto) GetAFCOk() (*PlatformAFCDto, bool)`

GetAFCOk returns a tuple with the AFC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAFC

`func (o *PlatformConfigurationDto) SetAFC(v PlatformAFCDto)`

SetAFC sets AFC field to given value.

### HasAFC

`func (o *PlatformConfigurationDto) HasAFC() bool`

HasAFC returns a boolean if a field has been set.

### GetDisablePublicUserSignup

`func (o *PlatformConfigurationDto) GetDisablePublicUserSignup() bool`

GetDisablePublicUserSignup returns the DisablePublicUserSignup field if non-nil, zero value otherwise.

### GetDisablePublicUserSignupOk

`func (o *PlatformConfigurationDto) GetDisablePublicUserSignupOk() (*bool, bool)`

GetDisablePublicUserSignupOk returns a tuple with the DisablePublicUserSignup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisablePublicUserSignup

`func (o *PlatformConfigurationDto) SetDisablePublicUserSignup(v bool)`

SetDisablePublicUserSignup sets DisablePublicUserSignup field to given value.

### HasDisablePublicUserSignup

`func (o *PlatformConfigurationDto) HasDisablePublicUserSignup() bool`

HasDisablePublicUserSignup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


