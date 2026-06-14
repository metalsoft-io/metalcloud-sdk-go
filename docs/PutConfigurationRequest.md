# PutConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**UsersDto**](UsersDto.md) | User management and authentication provider settings | [optional] 
**Ldap** | Pointer to [**LdapDto**](LdapDto.md) | LDAP directory integration settings | [optional] 
**Saml** | Pointer to [**SamlDto**](SamlDto.md) | SAML SSO integration settings | [optional] 
**Oauth** | Pointer to [**OauthDto**](OauthDto.md) | OAuth/OIDC token verification settings | [optional] 
**Microservices** | Pointer to [**MicroservicesDto**](MicroservicesDto.md) | Connected microservices configuration | [optional] 
**OauthWhitelistedPermissions** | Pointer to **[]string** | Permissions automatically granted to OAuth token holders regardless of their token scopes | [optional] 
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
**SyslogNotificationRules** | Pointer to [**[]SyslogRuleDto**](SyslogRuleDto.md) | Syslog notification rules for email alerting | [optional] 
**ServerHealthNotificationEmails** | Pointer to **[]string** | Email addresses to notify about server health events | [optional] 
**NetworkDeviceHealthNotificationEmails** | Pointer to **[]string** | Email addresses to notify about network device health events | [optional] 
**NetworkDeviceConfigurationDriftNotificationEmails** | Pointer to **[]string** | Email addresses to notify about network device configuration drift | [optional] 
**VmPoolHealthNotificationEmails** | Pointer to **[]string** | Email addresses to notify about VM pool health events | [optional] 
**EventRetention** | Pointer to **map[string]interface{}** | Event retention policy configuration | [optional] 
**SyslogBufferTtlMs** | Pointer to **float32** | Time-to-live for buffered syslog messages in milliseconds | [optional] 
**SyslogBufferMaxUnique** | Pointer to **float32** | Maximum number of unique syslog messages to buffer before flushing | [optional] 
**SharedSecret** | **string** | Shared secret used for tunnel authentication | 
**Bdk** | [**TunnelBdkDto**](TunnelBdkDto.md) | BDK service credentials | 
**Syslog** | [**TunnelSyslogDto**](TunnelSyslogDto.md) | Syslog message filtering configuration | 
**Branding** | Pointer to [**BrandingDto**](BrandingDto.md) | UI branding configuration (logos, colors, company name) | [optional] 
**LicenseMicroservice** | Pointer to [**LicenseMicroserviceDto**](LicenseMicroserviceDto.md) | License microservice connection settings | [optional] 
**RateLimiting** | Pointer to [**RateLimitingConfigDto**](RateLimitingConfigDto.md) | API rate limiting configuration | [optional] 
**ImageDownloadHttpProxy** | Pointer to **string** | Proxy for http:// image/asset downloads (e.g. http://proxy.example.com:8888). Empty or omitted means no proxy (direct download). | [optional] 
**ImageDownloadHttpsProxy** | Pointer to **string** | Proxy for https:// image/asset downloads (e.g. http://proxy.example.com:8888). Empty or omitted means no proxy (direct download). | [optional] 

## Methods

### NewPutConfigurationRequest

`func NewPutConfigurationRequest(sharedSecret string, bdk TunnelBdkDto, syslog TunnelSyslogDto, ) *PutConfigurationRequest`

NewPutConfigurationRequest instantiates a new PutConfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutConfigurationRequestWithDefaults

`func NewPutConfigurationRequestWithDefaults() *PutConfigurationRequest`

NewPutConfigurationRequestWithDefaults instantiates a new PutConfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *PutConfigurationRequest) GetUsers() UsersDto`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *PutConfigurationRequest) GetUsersOk() (*UsersDto, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *PutConfigurationRequest) SetUsers(v UsersDto)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *PutConfigurationRequest) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetLdap

`func (o *PutConfigurationRequest) GetLdap() LdapDto`

GetLdap returns the Ldap field if non-nil, zero value otherwise.

### GetLdapOk

`func (o *PutConfigurationRequest) GetLdapOk() (*LdapDto, bool)`

GetLdapOk returns a tuple with the Ldap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdap

`func (o *PutConfigurationRequest) SetLdap(v LdapDto)`

SetLdap sets Ldap field to given value.

### HasLdap

`func (o *PutConfigurationRequest) HasLdap() bool`

HasLdap returns a boolean if a field has been set.

### GetSaml

`func (o *PutConfigurationRequest) GetSaml() SamlDto`

GetSaml returns the Saml field if non-nil, zero value otherwise.

### GetSamlOk

`func (o *PutConfigurationRequest) GetSamlOk() (*SamlDto, bool)`

GetSamlOk returns a tuple with the Saml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml

`func (o *PutConfigurationRequest) SetSaml(v SamlDto)`

SetSaml sets Saml field to given value.

### HasSaml

`func (o *PutConfigurationRequest) HasSaml() bool`

HasSaml returns a boolean if a field has been set.

### GetOauth

`func (o *PutConfigurationRequest) GetOauth() OauthDto`

GetOauth returns the Oauth field if non-nil, zero value otherwise.

### GetOauthOk

`func (o *PutConfigurationRequest) GetOauthOk() (*OauthDto, bool)`

GetOauthOk returns a tuple with the Oauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth

`func (o *PutConfigurationRequest) SetOauth(v OauthDto)`

SetOauth sets Oauth field to given value.

### HasOauth

`func (o *PutConfigurationRequest) HasOauth() bool`

HasOauth returns a boolean if a field has been set.

### GetMicroservices

`func (o *PutConfigurationRequest) GetMicroservices() MicroservicesDto`

GetMicroservices returns the Microservices field if non-nil, zero value otherwise.

### GetMicroservicesOk

`func (o *PutConfigurationRequest) GetMicroservicesOk() (*MicroservicesDto, bool)`

GetMicroservicesOk returns a tuple with the Microservices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroservices

`func (o *PutConfigurationRequest) SetMicroservices(v MicroservicesDto)`

SetMicroservices sets Microservices field to given value.

### HasMicroservices

`func (o *PutConfigurationRequest) HasMicroservices() bool`

HasMicroservices returns a boolean if a field has been set.

### GetOauthWhitelistedPermissions

`func (o *PutConfigurationRequest) GetOauthWhitelistedPermissions() []string`

GetOauthWhitelistedPermissions returns the OauthWhitelistedPermissions field if non-nil, zero value otherwise.

### GetOauthWhitelistedPermissionsOk

`func (o *PutConfigurationRequest) GetOauthWhitelistedPermissionsOk() (*[]string, bool)`

GetOauthWhitelistedPermissionsOk returns a tuple with the OauthWhitelistedPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthWhitelistedPermissions

`func (o *PutConfigurationRequest) SetOauthWhitelistedPermissions(v []string)`

SetOauthWhitelistedPermissions sets OauthWhitelistedPermissions field to given value.

### HasOauthWhitelistedPermissions

`func (o *PutConfigurationRequest) HasOauthWhitelistedPermissions() bool`

HasOauthWhitelistedPermissions returns a boolean if a field has been set.

### GetMailClient

`func (o *PutConfigurationRequest) GetMailClient() PlatformMailClientDto`

GetMailClient returns the MailClient field if non-nil, zero value otherwise.

### GetMailClientOk

`func (o *PutConfigurationRequest) GetMailClientOk() (*PlatformMailClientDto, bool)`

GetMailClientOk returns a tuple with the MailClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailClient

`func (o *PutConfigurationRequest) SetMailClient(v PlatformMailClientDto)`

SetMailClient sets MailClient field to given value.

### HasMailClient

`func (o *PutConfigurationRequest) HasMailClient() bool`

HasMailClient returns a boolean if a field has been set.

### GetPowerDNS

`func (o *PutConfigurationRequest) GetPowerDNS() PlatformPowerDNSDto`

GetPowerDNS returns the PowerDNS field if non-nil, zero value otherwise.

### GetPowerDNSOk

`func (o *PutConfigurationRequest) GetPowerDNSOk() (*PlatformPowerDNSDto, bool)`

GetPowerDNSOk returns a tuple with the PowerDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerDNS

`func (o *PutConfigurationRequest) SetPowerDNS(v PlatformPowerDNSDto)`

SetPowerDNS sets PowerDNS field to given value.

### HasPowerDNS

`func (o *PutConfigurationRequest) HasPowerDNS() bool`

HasPowerDNS returns a boolean if a field has been set.

### GetBSIAdminURLRoot

`func (o *PutConfigurationRequest) GetBSIAdminURLRoot() string`

GetBSIAdminURLRoot returns the BSIAdminURLRoot field if non-nil, zero value otherwise.

### GetBSIAdminURLRootOk

`func (o *PutConfigurationRequest) GetBSIAdminURLRootOk() (*string, bool)`

GetBSIAdminURLRootOk returns a tuple with the BSIAdminURLRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBSIAdminURLRoot

`func (o *PutConfigurationRequest) SetBSIAdminURLRoot(v string)`

SetBSIAdminURLRoot sets BSIAdminURLRoot field to given value.

### HasBSIAdminURLRoot

`func (o *PutConfigurationRequest) HasBSIAdminURLRoot() bool`

HasBSIAdminURLRoot returns a boolean if a field has been set.

### GetRepoURLRootMaster

`func (o *PutConfigurationRequest) GetRepoURLRootMaster() string`

GetRepoURLRootMaster returns the RepoURLRootMaster field if non-nil, zero value otherwise.

### GetRepoURLRootMasterOk

`func (o *PutConfigurationRequest) GetRepoURLRootMasterOk() (*string, bool)`

GetRepoURLRootMasterOk returns a tuple with the RepoURLRootMaster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoURLRootMaster

`func (o *PutConfigurationRequest) SetRepoURLRootMaster(v string)`

SetRepoURLRootMaster sets RepoURLRootMaster field to given value.

### HasRepoURLRootMaster

`func (o *PutConfigurationRequest) HasRepoURLRootMaster() bool`

HasRepoURLRootMaster returns a boolean if a field has been set.

### GetUploads

`func (o *PutConfigurationRequest) GetUploads() PlatformUploadsDto`

GetUploads returns the Uploads field if non-nil, zero value otherwise.

### GetUploadsOk

`func (o *PutConfigurationRequest) GetUploadsOk() (*PlatformUploadsDto, bool)`

GetUploadsOk returns a tuple with the Uploads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploads

`func (o *PutConfigurationRequest) SetUploads(v PlatformUploadsDto)`

SetUploads sets Uploads field to given value.

### HasUploads

`func (o *PutConfigurationRequest) HasUploads() bool`

HasUploads returns a boolean if a field has been set.

### GetStorageConfig

`func (o *PutConfigurationRequest) GetStorageConfig() PlatformStorageConfigDto`

GetStorageConfig returns the StorageConfig field if non-nil, zero value otherwise.

### GetStorageConfigOk

`func (o *PutConfigurationRequest) GetStorageConfigOk() (*PlatformStorageConfigDto, bool)`

GetStorageConfigOk returns a tuple with the StorageConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfig

`func (o *PutConfigurationRequest) SetStorageConfig(v PlatformStorageConfigDto)`

SetStorageConfig sets StorageConfig field to given value.

### HasStorageConfig

`func (o *PutConfigurationRequest) HasStorageConfig() bool`

HasStorageConfig returns a boolean if a field has been set.

### GetAllowedPrefixSizesOnWAN

`func (o *PutConfigurationRequest) GetAllowedPrefixSizesOnWAN() PlatformAllowedPrefixSizesOnWANDto`

GetAllowedPrefixSizesOnWAN returns the AllowedPrefixSizesOnWAN field if non-nil, zero value otherwise.

### GetAllowedPrefixSizesOnWANOk

`func (o *PutConfigurationRequest) GetAllowedPrefixSizesOnWANOk() (*PlatformAllowedPrefixSizesOnWANDto, bool)`

GetAllowedPrefixSizesOnWANOk returns a tuple with the AllowedPrefixSizesOnWAN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPrefixSizesOnWAN

`func (o *PutConfigurationRequest) SetAllowedPrefixSizesOnWAN(v PlatformAllowedPrefixSizesOnWANDto)`

SetAllowedPrefixSizesOnWAN sets AllowedPrefixSizesOnWAN field to given value.

### HasAllowedPrefixSizesOnWAN

`func (o *PutConfigurationRequest) HasAllowedPrefixSizesOnWAN() bool`

HasAllowedPrefixSizesOnWAN returns a boolean if a field has been set.

### GetCookieDomain

`func (o *PutConfigurationRequest) GetCookieDomain() string`

GetCookieDomain returns the CookieDomain field if non-nil, zero value otherwise.

### GetCookieDomainOk

`func (o *PutConfigurationRequest) GetCookieDomainOk() (*string, bool)`

GetCookieDomainOk returns a tuple with the CookieDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieDomain

`func (o *PutConfigurationRequest) SetCookieDomain(v string)`

SetCookieDomain sets CookieDomain field to given value.

### HasCookieDomain

`func (o *PutConfigurationRequest) HasCookieDomain() bool`

HasCookieDomain returns a boolean if a field has been set.

### GetAllowTFTPBootThroughWAN

`func (o *PutConfigurationRequest) GetAllowTFTPBootThroughWAN() bool`

GetAllowTFTPBootThroughWAN returns the AllowTFTPBootThroughWAN field if non-nil, zero value otherwise.

### GetAllowTFTPBootThroughWANOk

`func (o *PutConfigurationRequest) GetAllowTFTPBootThroughWANOk() (*bool, bool)`

GetAllowTFTPBootThroughWANOk returns a tuple with the AllowTFTPBootThroughWAN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTFTPBootThroughWAN

`func (o *PutConfigurationRequest) SetAllowTFTPBootThroughWAN(v bool)`

SetAllowTFTPBootThroughWAN sets AllowTFTPBootThroughWAN field to given value.

### HasAllowTFTPBootThroughWAN

`func (o *PutConfigurationRequest) HasAllowTFTPBootThroughWAN() bool`

HasAllowTFTPBootThroughWAN returns a boolean if a field has been set.

### GetAllowServersWithOneInterface

`func (o *PutConfigurationRequest) GetAllowServersWithOneInterface() bool`

GetAllowServersWithOneInterface returns the AllowServersWithOneInterface field if non-nil, zero value otherwise.

### GetAllowServersWithOneInterfaceOk

`func (o *PutConfigurationRequest) GetAllowServersWithOneInterfaceOk() (*bool, bool)`

GetAllowServersWithOneInterfaceOk returns a tuple with the AllowServersWithOneInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowServersWithOneInterface

`func (o *PutConfigurationRequest) SetAllowServersWithOneInterface(v bool)`

SetAllowServersWithOneInterface sets AllowServersWithOneInterface field to given value.

### HasAllowServersWithOneInterface

`func (o *PutConfigurationRequest) HasAllowServersWithOneInterface() bool`

HasAllowServersWithOneInterface returns a boolean if a field has been set.

### GetAFC

`func (o *PutConfigurationRequest) GetAFC() PlatformAFCDto`

GetAFC returns the AFC field if non-nil, zero value otherwise.

### GetAFCOk

`func (o *PutConfigurationRequest) GetAFCOk() (*PlatformAFCDto, bool)`

GetAFCOk returns a tuple with the AFC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAFC

`func (o *PutConfigurationRequest) SetAFC(v PlatformAFCDto)`

SetAFC sets AFC field to given value.

### HasAFC

`func (o *PutConfigurationRequest) HasAFC() bool`

HasAFC returns a boolean if a field has been set.

### GetDisablePublicUserSignup

`func (o *PutConfigurationRequest) GetDisablePublicUserSignup() bool`

GetDisablePublicUserSignup returns the DisablePublicUserSignup field if non-nil, zero value otherwise.

### GetDisablePublicUserSignupOk

`func (o *PutConfigurationRequest) GetDisablePublicUserSignupOk() (*bool, bool)`

GetDisablePublicUserSignupOk returns a tuple with the DisablePublicUserSignup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisablePublicUserSignup

`func (o *PutConfigurationRequest) SetDisablePublicUserSignup(v bool)`

SetDisablePublicUserSignup sets DisablePublicUserSignup field to given value.

### HasDisablePublicUserSignup

`func (o *PutConfigurationRequest) HasDisablePublicUserSignup() bool`

HasDisablePublicUserSignup returns a boolean if a field has been set.

### GetSyslogNotificationRules

`func (o *PutConfigurationRequest) GetSyslogNotificationRules() []SyslogRuleDto`

GetSyslogNotificationRules returns the SyslogNotificationRules field if non-nil, zero value otherwise.

### GetSyslogNotificationRulesOk

`func (o *PutConfigurationRequest) GetSyslogNotificationRulesOk() (*[]SyslogRuleDto, bool)`

GetSyslogNotificationRulesOk returns a tuple with the SyslogNotificationRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogNotificationRules

`func (o *PutConfigurationRequest) SetSyslogNotificationRules(v []SyslogRuleDto)`

SetSyslogNotificationRules sets SyslogNotificationRules field to given value.

### HasSyslogNotificationRules

`func (o *PutConfigurationRequest) HasSyslogNotificationRules() bool`

HasSyslogNotificationRules returns a boolean if a field has been set.

### GetServerHealthNotificationEmails

`func (o *PutConfigurationRequest) GetServerHealthNotificationEmails() []string`

GetServerHealthNotificationEmails returns the ServerHealthNotificationEmails field if non-nil, zero value otherwise.

### GetServerHealthNotificationEmailsOk

`func (o *PutConfigurationRequest) GetServerHealthNotificationEmailsOk() (*[]string, bool)`

GetServerHealthNotificationEmailsOk returns a tuple with the ServerHealthNotificationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHealthNotificationEmails

`func (o *PutConfigurationRequest) SetServerHealthNotificationEmails(v []string)`

SetServerHealthNotificationEmails sets ServerHealthNotificationEmails field to given value.

### HasServerHealthNotificationEmails

`func (o *PutConfigurationRequest) HasServerHealthNotificationEmails() bool`

HasServerHealthNotificationEmails returns a boolean if a field has been set.

### GetNetworkDeviceHealthNotificationEmails

`func (o *PutConfigurationRequest) GetNetworkDeviceHealthNotificationEmails() []string`

GetNetworkDeviceHealthNotificationEmails returns the NetworkDeviceHealthNotificationEmails field if non-nil, zero value otherwise.

### GetNetworkDeviceHealthNotificationEmailsOk

`func (o *PutConfigurationRequest) GetNetworkDeviceHealthNotificationEmailsOk() (*[]string, bool)`

GetNetworkDeviceHealthNotificationEmailsOk returns a tuple with the NetworkDeviceHealthNotificationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceHealthNotificationEmails

`func (o *PutConfigurationRequest) SetNetworkDeviceHealthNotificationEmails(v []string)`

SetNetworkDeviceHealthNotificationEmails sets NetworkDeviceHealthNotificationEmails field to given value.

### HasNetworkDeviceHealthNotificationEmails

`func (o *PutConfigurationRequest) HasNetworkDeviceHealthNotificationEmails() bool`

HasNetworkDeviceHealthNotificationEmails returns a boolean if a field has been set.

### GetNetworkDeviceConfigurationDriftNotificationEmails

`func (o *PutConfigurationRequest) GetNetworkDeviceConfigurationDriftNotificationEmails() []string`

GetNetworkDeviceConfigurationDriftNotificationEmails returns the NetworkDeviceConfigurationDriftNotificationEmails field if non-nil, zero value otherwise.

### GetNetworkDeviceConfigurationDriftNotificationEmailsOk

`func (o *PutConfigurationRequest) GetNetworkDeviceConfigurationDriftNotificationEmailsOk() (*[]string, bool)`

GetNetworkDeviceConfigurationDriftNotificationEmailsOk returns a tuple with the NetworkDeviceConfigurationDriftNotificationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceConfigurationDriftNotificationEmails

`func (o *PutConfigurationRequest) SetNetworkDeviceConfigurationDriftNotificationEmails(v []string)`

SetNetworkDeviceConfigurationDriftNotificationEmails sets NetworkDeviceConfigurationDriftNotificationEmails field to given value.

### HasNetworkDeviceConfigurationDriftNotificationEmails

`func (o *PutConfigurationRequest) HasNetworkDeviceConfigurationDriftNotificationEmails() bool`

HasNetworkDeviceConfigurationDriftNotificationEmails returns a boolean if a field has been set.

### GetVmPoolHealthNotificationEmails

`func (o *PutConfigurationRequest) GetVmPoolHealthNotificationEmails() []string`

GetVmPoolHealthNotificationEmails returns the VmPoolHealthNotificationEmails field if non-nil, zero value otherwise.

### GetVmPoolHealthNotificationEmailsOk

`func (o *PutConfigurationRequest) GetVmPoolHealthNotificationEmailsOk() (*[]string, bool)`

GetVmPoolHealthNotificationEmailsOk returns a tuple with the VmPoolHealthNotificationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmPoolHealthNotificationEmails

`func (o *PutConfigurationRequest) SetVmPoolHealthNotificationEmails(v []string)`

SetVmPoolHealthNotificationEmails sets VmPoolHealthNotificationEmails field to given value.

### HasVmPoolHealthNotificationEmails

`func (o *PutConfigurationRequest) HasVmPoolHealthNotificationEmails() bool`

HasVmPoolHealthNotificationEmails returns a boolean if a field has been set.

### GetEventRetention

`func (o *PutConfigurationRequest) GetEventRetention() map[string]interface{}`

GetEventRetention returns the EventRetention field if non-nil, zero value otherwise.

### GetEventRetentionOk

`func (o *PutConfigurationRequest) GetEventRetentionOk() (*map[string]interface{}, bool)`

GetEventRetentionOk returns a tuple with the EventRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventRetention

`func (o *PutConfigurationRequest) SetEventRetention(v map[string]interface{})`

SetEventRetention sets EventRetention field to given value.

### HasEventRetention

`func (o *PutConfigurationRequest) HasEventRetention() bool`

HasEventRetention returns a boolean if a field has been set.

### GetSyslogBufferTtlMs

`func (o *PutConfigurationRequest) GetSyslogBufferTtlMs() float32`

GetSyslogBufferTtlMs returns the SyslogBufferTtlMs field if non-nil, zero value otherwise.

### GetSyslogBufferTtlMsOk

`func (o *PutConfigurationRequest) GetSyslogBufferTtlMsOk() (*float32, bool)`

GetSyslogBufferTtlMsOk returns a tuple with the SyslogBufferTtlMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogBufferTtlMs

`func (o *PutConfigurationRequest) SetSyslogBufferTtlMs(v float32)`

SetSyslogBufferTtlMs sets SyslogBufferTtlMs field to given value.

### HasSyslogBufferTtlMs

`func (o *PutConfigurationRequest) HasSyslogBufferTtlMs() bool`

HasSyslogBufferTtlMs returns a boolean if a field has been set.

### GetSyslogBufferMaxUnique

`func (o *PutConfigurationRequest) GetSyslogBufferMaxUnique() float32`

GetSyslogBufferMaxUnique returns the SyslogBufferMaxUnique field if non-nil, zero value otherwise.

### GetSyslogBufferMaxUniqueOk

`func (o *PutConfigurationRequest) GetSyslogBufferMaxUniqueOk() (*float32, bool)`

GetSyslogBufferMaxUniqueOk returns a tuple with the SyslogBufferMaxUnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogBufferMaxUnique

`func (o *PutConfigurationRequest) SetSyslogBufferMaxUnique(v float32)`

SetSyslogBufferMaxUnique sets SyslogBufferMaxUnique field to given value.

### HasSyslogBufferMaxUnique

`func (o *PutConfigurationRequest) HasSyslogBufferMaxUnique() bool`

HasSyslogBufferMaxUnique returns a boolean if a field has been set.

### GetSharedSecret

`func (o *PutConfigurationRequest) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *PutConfigurationRequest) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *PutConfigurationRequest) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.


### GetBdk

`func (o *PutConfigurationRequest) GetBdk() TunnelBdkDto`

GetBdk returns the Bdk field if non-nil, zero value otherwise.

### GetBdkOk

`func (o *PutConfigurationRequest) GetBdkOk() (*TunnelBdkDto, bool)`

GetBdkOk returns a tuple with the Bdk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBdk

`func (o *PutConfigurationRequest) SetBdk(v TunnelBdkDto)`

SetBdk sets Bdk field to given value.


### GetSyslog

`func (o *PutConfigurationRequest) GetSyslog() TunnelSyslogDto`

GetSyslog returns the Syslog field if non-nil, zero value otherwise.

### GetSyslogOk

`func (o *PutConfigurationRequest) GetSyslogOk() (*TunnelSyslogDto, bool)`

GetSyslogOk returns a tuple with the Syslog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslog

`func (o *PutConfigurationRequest) SetSyslog(v TunnelSyslogDto)`

SetSyslog sets Syslog field to given value.


### GetBranding

`func (o *PutConfigurationRequest) GetBranding() BrandingDto`

GetBranding returns the Branding field if non-nil, zero value otherwise.

### GetBrandingOk

`func (o *PutConfigurationRequest) GetBrandingOk() (*BrandingDto, bool)`

GetBrandingOk returns a tuple with the Branding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranding

`func (o *PutConfigurationRequest) SetBranding(v BrandingDto)`

SetBranding sets Branding field to given value.

### HasBranding

`func (o *PutConfigurationRequest) HasBranding() bool`

HasBranding returns a boolean if a field has been set.

### GetLicenseMicroservice

`func (o *PutConfigurationRequest) GetLicenseMicroservice() LicenseMicroserviceDto`

GetLicenseMicroservice returns the LicenseMicroservice field if non-nil, zero value otherwise.

### GetLicenseMicroserviceOk

`func (o *PutConfigurationRequest) GetLicenseMicroserviceOk() (*LicenseMicroserviceDto, bool)`

GetLicenseMicroserviceOk returns a tuple with the LicenseMicroservice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseMicroservice

`func (o *PutConfigurationRequest) SetLicenseMicroservice(v LicenseMicroserviceDto)`

SetLicenseMicroservice sets LicenseMicroservice field to given value.

### HasLicenseMicroservice

`func (o *PutConfigurationRequest) HasLicenseMicroservice() bool`

HasLicenseMicroservice returns a boolean if a field has been set.

### GetRateLimiting

`func (o *PutConfigurationRequest) GetRateLimiting() RateLimitingConfigDto`

GetRateLimiting returns the RateLimiting field if non-nil, zero value otherwise.

### GetRateLimitingOk

`func (o *PutConfigurationRequest) GetRateLimitingOk() (*RateLimitingConfigDto, bool)`

GetRateLimitingOk returns a tuple with the RateLimiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimiting

`func (o *PutConfigurationRequest) SetRateLimiting(v RateLimitingConfigDto)`

SetRateLimiting sets RateLimiting field to given value.

### HasRateLimiting

`func (o *PutConfigurationRequest) HasRateLimiting() bool`

HasRateLimiting returns a boolean if a field has been set.

### GetImageDownloadHttpProxy

`func (o *PutConfigurationRequest) GetImageDownloadHttpProxy() string`

GetImageDownloadHttpProxy returns the ImageDownloadHttpProxy field if non-nil, zero value otherwise.

### GetImageDownloadHttpProxyOk

`func (o *PutConfigurationRequest) GetImageDownloadHttpProxyOk() (*string, bool)`

GetImageDownloadHttpProxyOk returns a tuple with the ImageDownloadHttpProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDownloadHttpProxy

`func (o *PutConfigurationRequest) SetImageDownloadHttpProxy(v string)`

SetImageDownloadHttpProxy sets ImageDownloadHttpProxy field to given value.

### HasImageDownloadHttpProxy

`func (o *PutConfigurationRequest) HasImageDownloadHttpProxy() bool`

HasImageDownloadHttpProxy returns a boolean if a field has been set.

### GetImageDownloadHttpsProxy

`func (o *PutConfigurationRequest) GetImageDownloadHttpsProxy() string`

GetImageDownloadHttpsProxy returns the ImageDownloadHttpsProxy field if non-nil, zero value otherwise.

### GetImageDownloadHttpsProxyOk

`func (o *PutConfigurationRequest) GetImageDownloadHttpsProxyOk() (*string, bool)`

GetImageDownloadHttpsProxyOk returns a tuple with the ImageDownloadHttpsProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDownloadHttpsProxy

`func (o *PutConfigurationRequest) SetImageDownloadHttpsProxy(v string)`

SetImageDownloadHttpsProxy sets ImageDownloadHttpsProxy field to given value.

### HasImageDownloadHttpsProxy

`func (o *PutConfigurationRequest) HasImageDownloadHttpsProxy() bool`

HasImageDownloadHttpsProxy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


