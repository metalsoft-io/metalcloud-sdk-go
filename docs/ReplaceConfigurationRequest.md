# ReplaceConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**Users**](Users.md) | User management and authentication provider settings | [optional] 
**Ldap** | Pointer to [**Ldap**](Ldap.md) | LDAP directory integration settings | [optional] 
**Saml** | Pointer to [**Saml**](Saml.md) | SAML SSO integration settings | [optional] 
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
**SyslogNotificationRules** | Pointer to [**[]SyslogRule**](SyslogRule.md) | Syslog notification rules for email alerting | [optional] 
**ServerHealthNotificationEmails** | Pointer to **[]string** | Email addresses to notify about server health events | [optional] 
**NetworkDeviceHealthNotificationEmails** | Pointer to **[]string** | Email addresses to notify about network device health events | [optional] 
**NetworkDeviceConfigurationDriftNotificationEmails** | Pointer to **[]string** | Email addresses to notify about network device configuration drift | [optional] 
**VmPoolHealthNotificationEmails** | Pointer to **[]string** | Email addresses to notify about VM pool health events | [optional] 
**EventRetention** | Pointer to **map[string]interface{}** | Event retention policy configuration | [optional] 
**SyslogBufferTtlMs** | Pointer to **float32** | Time-to-live for buffered syslog messages in milliseconds | [optional] 
**SyslogBufferMaxUnique** | Pointer to **float32** | Maximum number of unique syslog messages to buffer before flushing | [optional] 
**SharedSecret** | **string** | Shared secret used for tunnel authentication | 
**Bdk** | [**TunnelBdk**](TunnelBdk.md) | BDK service credentials | 
**Syslog** | [**TunnelSyslog**](TunnelSyslog.md) | Syslog message filtering configuration | 
**Branding** | Pointer to [**Branding**](Branding.md) | UI branding configuration (logos, colors, company name) | [optional] 
**LicenseMicroservice** | Pointer to [**LicenseMicroservice**](LicenseMicroservice.md) | License microservice connection settings | [optional] 
**Microservices** | Pointer to [**Microservices**](Microservices.md) | Connected microservices configuration | [optional] 
**RateLimiting** | Pointer to [**RateLimitingConfig**](RateLimitingConfig.md) | API rate limiting configuration | [optional] 
**ImageDownloadHttpProxy** | Pointer to **string** | Proxy for http:// image/asset downloads (e.g. http://proxy.example.com:8888). Empty or omitted means no proxy (direct download). | [optional] 
**ImageDownloadHttpsProxy** | Pointer to **string** | Proxy for https:// image/asset downloads (e.g. http://proxy.example.com:8888). Empty or omitted means no proxy (direct download). | [optional] 

## Methods

### NewReplaceConfigurationRequest

`func NewReplaceConfigurationRequest(sharedSecret string, bdk TunnelBdk, syslog TunnelSyslog, ) *ReplaceConfigurationRequest`

NewReplaceConfigurationRequest instantiates a new ReplaceConfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplaceConfigurationRequestWithDefaults

`func NewReplaceConfigurationRequestWithDefaults() *ReplaceConfigurationRequest`

NewReplaceConfigurationRequestWithDefaults instantiates a new ReplaceConfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *ReplaceConfigurationRequest) GetUsers() Users`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ReplaceConfigurationRequest) GetUsersOk() (*Users, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ReplaceConfigurationRequest) SetUsers(v Users)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *ReplaceConfigurationRequest) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetLdap

`func (o *ReplaceConfigurationRequest) GetLdap() Ldap`

GetLdap returns the Ldap field if non-nil, zero value otherwise.

### GetLdapOk

`func (o *ReplaceConfigurationRequest) GetLdapOk() (*Ldap, bool)`

GetLdapOk returns a tuple with the Ldap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdap

`func (o *ReplaceConfigurationRequest) SetLdap(v Ldap)`

SetLdap sets Ldap field to given value.

### HasLdap

`func (o *ReplaceConfigurationRequest) HasLdap() bool`

HasLdap returns a boolean if a field has been set.

### GetSaml

`func (o *ReplaceConfigurationRequest) GetSaml() Saml`

GetSaml returns the Saml field if non-nil, zero value otherwise.

### GetSamlOk

`func (o *ReplaceConfigurationRequest) GetSamlOk() (*Saml, bool)`

GetSamlOk returns a tuple with the Saml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml

`func (o *ReplaceConfigurationRequest) SetSaml(v Saml)`

SetSaml sets Saml field to given value.

### HasSaml

`func (o *ReplaceConfigurationRequest) HasSaml() bool`

HasSaml returns a boolean if a field has been set.

### GetMailClient

`func (o *ReplaceConfigurationRequest) GetMailClient() PlatformMailClient`

GetMailClient returns the MailClient field if non-nil, zero value otherwise.

### GetMailClientOk

`func (o *ReplaceConfigurationRequest) GetMailClientOk() (*PlatformMailClient, bool)`

GetMailClientOk returns a tuple with the MailClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailClient

`func (o *ReplaceConfigurationRequest) SetMailClient(v PlatformMailClient)`

SetMailClient sets MailClient field to given value.

### HasMailClient

`func (o *ReplaceConfigurationRequest) HasMailClient() bool`

HasMailClient returns a boolean if a field has been set.

### GetPowerDNS

`func (o *ReplaceConfigurationRequest) GetPowerDNS() PlatformPowerDNS`

GetPowerDNS returns the PowerDNS field if non-nil, zero value otherwise.

### GetPowerDNSOk

`func (o *ReplaceConfigurationRequest) GetPowerDNSOk() (*PlatformPowerDNS, bool)`

GetPowerDNSOk returns a tuple with the PowerDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerDNS

`func (o *ReplaceConfigurationRequest) SetPowerDNS(v PlatformPowerDNS)`

SetPowerDNS sets PowerDNS field to given value.

### HasPowerDNS

`func (o *ReplaceConfigurationRequest) HasPowerDNS() bool`

HasPowerDNS returns a boolean if a field has been set.

### GetBSIAdminURLRoot

`func (o *ReplaceConfigurationRequest) GetBSIAdminURLRoot() string`

GetBSIAdminURLRoot returns the BSIAdminURLRoot field if non-nil, zero value otherwise.

### GetBSIAdminURLRootOk

`func (o *ReplaceConfigurationRequest) GetBSIAdminURLRootOk() (*string, bool)`

GetBSIAdminURLRootOk returns a tuple with the BSIAdminURLRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBSIAdminURLRoot

`func (o *ReplaceConfigurationRequest) SetBSIAdminURLRoot(v string)`

SetBSIAdminURLRoot sets BSIAdminURLRoot field to given value.

### HasBSIAdminURLRoot

`func (o *ReplaceConfigurationRequest) HasBSIAdminURLRoot() bool`

HasBSIAdminURLRoot returns a boolean if a field has been set.

### GetRepoURLRootMaster

`func (o *ReplaceConfigurationRequest) GetRepoURLRootMaster() string`

GetRepoURLRootMaster returns the RepoURLRootMaster field if non-nil, zero value otherwise.

### GetRepoURLRootMasterOk

`func (o *ReplaceConfigurationRequest) GetRepoURLRootMasterOk() (*string, bool)`

GetRepoURLRootMasterOk returns a tuple with the RepoURLRootMaster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoURLRootMaster

`func (o *ReplaceConfigurationRequest) SetRepoURLRootMaster(v string)`

SetRepoURLRootMaster sets RepoURLRootMaster field to given value.

### HasRepoURLRootMaster

`func (o *ReplaceConfigurationRequest) HasRepoURLRootMaster() bool`

HasRepoURLRootMaster returns a boolean if a field has been set.

### GetUploads

`func (o *ReplaceConfigurationRequest) GetUploads() PlatformUploads`

GetUploads returns the Uploads field if non-nil, zero value otherwise.

### GetUploadsOk

`func (o *ReplaceConfigurationRequest) GetUploadsOk() (*PlatformUploads, bool)`

GetUploadsOk returns a tuple with the Uploads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploads

`func (o *ReplaceConfigurationRequest) SetUploads(v PlatformUploads)`

SetUploads sets Uploads field to given value.

### HasUploads

`func (o *ReplaceConfigurationRequest) HasUploads() bool`

HasUploads returns a boolean if a field has been set.

### GetStorageConfig

`func (o *ReplaceConfigurationRequest) GetStorageConfig() PlatformStorageConfig`

GetStorageConfig returns the StorageConfig field if non-nil, zero value otherwise.

### GetStorageConfigOk

`func (o *ReplaceConfigurationRequest) GetStorageConfigOk() (*PlatformStorageConfig, bool)`

GetStorageConfigOk returns a tuple with the StorageConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfig

`func (o *ReplaceConfigurationRequest) SetStorageConfig(v PlatformStorageConfig)`

SetStorageConfig sets StorageConfig field to given value.

### HasStorageConfig

`func (o *ReplaceConfigurationRequest) HasStorageConfig() bool`

HasStorageConfig returns a boolean if a field has been set.

### GetAllowedPrefixSizesOnWAN

`func (o *ReplaceConfigurationRequest) GetAllowedPrefixSizesOnWAN() PlatformAllowedPrefixSizesOnWAN`

GetAllowedPrefixSizesOnWAN returns the AllowedPrefixSizesOnWAN field if non-nil, zero value otherwise.

### GetAllowedPrefixSizesOnWANOk

`func (o *ReplaceConfigurationRequest) GetAllowedPrefixSizesOnWANOk() (*PlatformAllowedPrefixSizesOnWAN, bool)`

GetAllowedPrefixSizesOnWANOk returns a tuple with the AllowedPrefixSizesOnWAN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPrefixSizesOnWAN

`func (o *ReplaceConfigurationRequest) SetAllowedPrefixSizesOnWAN(v PlatformAllowedPrefixSizesOnWAN)`

SetAllowedPrefixSizesOnWAN sets AllowedPrefixSizesOnWAN field to given value.

### HasAllowedPrefixSizesOnWAN

`func (o *ReplaceConfigurationRequest) HasAllowedPrefixSizesOnWAN() bool`

HasAllowedPrefixSizesOnWAN returns a boolean if a field has been set.

### GetCookieDomain

`func (o *ReplaceConfigurationRequest) GetCookieDomain() string`

GetCookieDomain returns the CookieDomain field if non-nil, zero value otherwise.

### GetCookieDomainOk

`func (o *ReplaceConfigurationRequest) GetCookieDomainOk() (*string, bool)`

GetCookieDomainOk returns a tuple with the CookieDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieDomain

`func (o *ReplaceConfigurationRequest) SetCookieDomain(v string)`

SetCookieDomain sets CookieDomain field to given value.

### HasCookieDomain

`func (o *ReplaceConfigurationRequest) HasCookieDomain() bool`

HasCookieDomain returns a boolean if a field has been set.

### GetAllowTFTPBootThroughWAN

`func (o *ReplaceConfigurationRequest) GetAllowTFTPBootThroughWAN() bool`

GetAllowTFTPBootThroughWAN returns the AllowTFTPBootThroughWAN field if non-nil, zero value otherwise.

### GetAllowTFTPBootThroughWANOk

`func (o *ReplaceConfigurationRequest) GetAllowTFTPBootThroughWANOk() (*bool, bool)`

GetAllowTFTPBootThroughWANOk returns a tuple with the AllowTFTPBootThroughWAN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTFTPBootThroughWAN

`func (o *ReplaceConfigurationRequest) SetAllowTFTPBootThroughWAN(v bool)`

SetAllowTFTPBootThroughWAN sets AllowTFTPBootThroughWAN field to given value.

### HasAllowTFTPBootThroughWAN

`func (o *ReplaceConfigurationRequest) HasAllowTFTPBootThroughWAN() bool`

HasAllowTFTPBootThroughWAN returns a boolean if a field has been set.

### GetAllowServersWithOneInterface

`func (o *ReplaceConfigurationRequest) GetAllowServersWithOneInterface() bool`

GetAllowServersWithOneInterface returns the AllowServersWithOneInterface field if non-nil, zero value otherwise.

### GetAllowServersWithOneInterfaceOk

`func (o *ReplaceConfigurationRequest) GetAllowServersWithOneInterfaceOk() (*bool, bool)`

GetAllowServersWithOneInterfaceOk returns a tuple with the AllowServersWithOneInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowServersWithOneInterface

`func (o *ReplaceConfigurationRequest) SetAllowServersWithOneInterface(v bool)`

SetAllowServersWithOneInterface sets AllowServersWithOneInterface field to given value.

### HasAllowServersWithOneInterface

`func (o *ReplaceConfigurationRequest) HasAllowServersWithOneInterface() bool`

HasAllowServersWithOneInterface returns a boolean if a field has been set.

### GetAFC

`func (o *ReplaceConfigurationRequest) GetAFC() PlatformAFC`

GetAFC returns the AFC field if non-nil, zero value otherwise.

### GetAFCOk

`func (o *ReplaceConfigurationRequest) GetAFCOk() (*PlatformAFC, bool)`

GetAFCOk returns a tuple with the AFC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAFC

`func (o *ReplaceConfigurationRequest) SetAFC(v PlatformAFC)`

SetAFC sets AFC field to given value.

### HasAFC

`func (o *ReplaceConfigurationRequest) HasAFC() bool`

HasAFC returns a boolean if a field has been set.

### GetDisablePublicUserSignup

`func (o *ReplaceConfigurationRequest) GetDisablePublicUserSignup() bool`

GetDisablePublicUserSignup returns the DisablePublicUserSignup field if non-nil, zero value otherwise.

### GetDisablePublicUserSignupOk

`func (o *ReplaceConfigurationRequest) GetDisablePublicUserSignupOk() (*bool, bool)`

GetDisablePublicUserSignupOk returns a tuple with the DisablePublicUserSignup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisablePublicUserSignup

`func (o *ReplaceConfigurationRequest) SetDisablePublicUserSignup(v bool)`

SetDisablePublicUserSignup sets DisablePublicUserSignup field to given value.

### HasDisablePublicUserSignup

`func (o *ReplaceConfigurationRequest) HasDisablePublicUserSignup() bool`

HasDisablePublicUserSignup returns a boolean if a field has been set.

### GetSyslogNotificationRules

`func (o *ReplaceConfigurationRequest) GetSyslogNotificationRules() []SyslogRule`

GetSyslogNotificationRules returns the SyslogNotificationRules field if non-nil, zero value otherwise.

### GetSyslogNotificationRulesOk

`func (o *ReplaceConfigurationRequest) GetSyslogNotificationRulesOk() (*[]SyslogRule, bool)`

GetSyslogNotificationRulesOk returns a tuple with the SyslogNotificationRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogNotificationRules

`func (o *ReplaceConfigurationRequest) SetSyslogNotificationRules(v []SyslogRule)`

SetSyslogNotificationRules sets SyslogNotificationRules field to given value.

### HasSyslogNotificationRules

`func (o *ReplaceConfigurationRequest) HasSyslogNotificationRules() bool`

HasSyslogNotificationRules returns a boolean if a field has been set.

### GetServerHealthNotificationEmails

`func (o *ReplaceConfigurationRequest) GetServerHealthNotificationEmails() []string`

GetServerHealthNotificationEmails returns the ServerHealthNotificationEmails field if non-nil, zero value otherwise.

### GetServerHealthNotificationEmailsOk

`func (o *ReplaceConfigurationRequest) GetServerHealthNotificationEmailsOk() (*[]string, bool)`

GetServerHealthNotificationEmailsOk returns a tuple with the ServerHealthNotificationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHealthNotificationEmails

`func (o *ReplaceConfigurationRequest) SetServerHealthNotificationEmails(v []string)`

SetServerHealthNotificationEmails sets ServerHealthNotificationEmails field to given value.

### HasServerHealthNotificationEmails

`func (o *ReplaceConfigurationRequest) HasServerHealthNotificationEmails() bool`

HasServerHealthNotificationEmails returns a boolean if a field has been set.

### GetNetworkDeviceHealthNotificationEmails

`func (o *ReplaceConfigurationRequest) GetNetworkDeviceHealthNotificationEmails() []string`

GetNetworkDeviceHealthNotificationEmails returns the NetworkDeviceHealthNotificationEmails field if non-nil, zero value otherwise.

### GetNetworkDeviceHealthNotificationEmailsOk

`func (o *ReplaceConfigurationRequest) GetNetworkDeviceHealthNotificationEmailsOk() (*[]string, bool)`

GetNetworkDeviceHealthNotificationEmailsOk returns a tuple with the NetworkDeviceHealthNotificationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceHealthNotificationEmails

`func (o *ReplaceConfigurationRequest) SetNetworkDeviceHealthNotificationEmails(v []string)`

SetNetworkDeviceHealthNotificationEmails sets NetworkDeviceHealthNotificationEmails field to given value.

### HasNetworkDeviceHealthNotificationEmails

`func (o *ReplaceConfigurationRequest) HasNetworkDeviceHealthNotificationEmails() bool`

HasNetworkDeviceHealthNotificationEmails returns a boolean if a field has been set.

### GetNetworkDeviceConfigurationDriftNotificationEmails

`func (o *ReplaceConfigurationRequest) GetNetworkDeviceConfigurationDriftNotificationEmails() []string`

GetNetworkDeviceConfigurationDriftNotificationEmails returns the NetworkDeviceConfigurationDriftNotificationEmails field if non-nil, zero value otherwise.

### GetNetworkDeviceConfigurationDriftNotificationEmailsOk

`func (o *ReplaceConfigurationRequest) GetNetworkDeviceConfigurationDriftNotificationEmailsOk() (*[]string, bool)`

GetNetworkDeviceConfigurationDriftNotificationEmailsOk returns a tuple with the NetworkDeviceConfigurationDriftNotificationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeviceConfigurationDriftNotificationEmails

`func (o *ReplaceConfigurationRequest) SetNetworkDeviceConfigurationDriftNotificationEmails(v []string)`

SetNetworkDeviceConfigurationDriftNotificationEmails sets NetworkDeviceConfigurationDriftNotificationEmails field to given value.

### HasNetworkDeviceConfigurationDriftNotificationEmails

`func (o *ReplaceConfigurationRequest) HasNetworkDeviceConfigurationDriftNotificationEmails() bool`

HasNetworkDeviceConfigurationDriftNotificationEmails returns a boolean if a field has been set.

### GetVmPoolHealthNotificationEmails

`func (o *ReplaceConfigurationRequest) GetVmPoolHealthNotificationEmails() []string`

GetVmPoolHealthNotificationEmails returns the VmPoolHealthNotificationEmails field if non-nil, zero value otherwise.

### GetVmPoolHealthNotificationEmailsOk

`func (o *ReplaceConfigurationRequest) GetVmPoolHealthNotificationEmailsOk() (*[]string, bool)`

GetVmPoolHealthNotificationEmailsOk returns a tuple with the VmPoolHealthNotificationEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmPoolHealthNotificationEmails

`func (o *ReplaceConfigurationRequest) SetVmPoolHealthNotificationEmails(v []string)`

SetVmPoolHealthNotificationEmails sets VmPoolHealthNotificationEmails field to given value.

### HasVmPoolHealthNotificationEmails

`func (o *ReplaceConfigurationRequest) HasVmPoolHealthNotificationEmails() bool`

HasVmPoolHealthNotificationEmails returns a boolean if a field has been set.

### GetEventRetention

`func (o *ReplaceConfigurationRequest) GetEventRetention() map[string]interface{}`

GetEventRetention returns the EventRetention field if non-nil, zero value otherwise.

### GetEventRetentionOk

`func (o *ReplaceConfigurationRequest) GetEventRetentionOk() (*map[string]interface{}, bool)`

GetEventRetentionOk returns a tuple with the EventRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventRetention

`func (o *ReplaceConfigurationRequest) SetEventRetention(v map[string]interface{})`

SetEventRetention sets EventRetention field to given value.

### HasEventRetention

`func (o *ReplaceConfigurationRequest) HasEventRetention() bool`

HasEventRetention returns a boolean if a field has been set.

### GetSyslogBufferTtlMs

`func (o *ReplaceConfigurationRequest) GetSyslogBufferTtlMs() float32`

GetSyslogBufferTtlMs returns the SyslogBufferTtlMs field if non-nil, zero value otherwise.

### GetSyslogBufferTtlMsOk

`func (o *ReplaceConfigurationRequest) GetSyslogBufferTtlMsOk() (*float32, bool)`

GetSyslogBufferTtlMsOk returns a tuple with the SyslogBufferTtlMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogBufferTtlMs

`func (o *ReplaceConfigurationRequest) SetSyslogBufferTtlMs(v float32)`

SetSyslogBufferTtlMs sets SyslogBufferTtlMs field to given value.

### HasSyslogBufferTtlMs

`func (o *ReplaceConfigurationRequest) HasSyslogBufferTtlMs() bool`

HasSyslogBufferTtlMs returns a boolean if a field has been set.

### GetSyslogBufferMaxUnique

`func (o *ReplaceConfigurationRequest) GetSyslogBufferMaxUnique() float32`

GetSyslogBufferMaxUnique returns the SyslogBufferMaxUnique field if non-nil, zero value otherwise.

### GetSyslogBufferMaxUniqueOk

`func (o *ReplaceConfigurationRequest) GetSyslogBufferMaxUniqueOk() (*float32, bool)`

GetSyslogBufferMaxUniqueOk returns a tuple with the SyslogBufferMaxUnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogBufferMaxUnique

`func (o *ReplaceConfigurationRequest) SetSyslogBufferMaxUnique(v float32)`

SetSyslogBufferMaxUnique sets SyslogBufferMaxUnique field to given value.

### HasSyslogBufferMaxUnique

`func (o *ReplaceConfigurationRequest) HasSyslogBufferMaxUnique() bool`

HasSyslogBufferMaxUnique returns a boolean if a field has been set.

### GetSharedSecret

`func (o *ReplaceConfigurationRequest) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *ReplaceConfigurationRequest) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *ReplaceConfigurationRequest) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.


### GetBdk

`func (o *ReplaceConfigurationRequest) GetBdk() TunnelBdk`

GetBdk returns the Bdk field if non-nil, zero value otherwise.

### GetBdkOk

`func (o *ReplaceConfigurationRequest) GetBdkOk() (*TunnelBdk, bool)`

GetBdkOk returns a tuple with the Bdk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBdk

`func (o *ReplaceConfigurationRequest) SetBdk(v TunnelBdk)`

SetBdk sets Bdk field to given value.


### GetSyslog

`func (o *ReplaceConfigurationRequest) GetSyslog() TunnelSyslog`

GetSyslog returns the Syslog field if non-nil, zero value otherwise.

### GetSyslogOk

`func (o *ReplaceConfigurationRequest) GetSyslogOk() (*TunnelSyslog, bool)`

GetSyslogOk returns a tuple with the Syslog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslog

`func (o *ReplaceConfigurationRequest) SetSyslog(v TunnelSyslog)`

SetSyslog sets Syslog field to given value.


### GetBranding

`func (o *ReplaceConfigurationRequest) GetBranding() Branding`

GetBranding returns the Branding field if non-nil, zero value otherwise.

### GetBrandingOk

`func (o *ReplaceConfigurationRequest) GetBrandingOk() (*Branding, bool)`

GetBrandingOk returns a tuple with the Branding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranding

`func (o *ReplaceConfigurationRequest) SetBranding(v Branding)`

SetBranding sets Branding field to given value.

### HasBranding

`func (o *ReplaceConfigurationRequest) HasBranding() bool`

HasBranding returns a boolean if a field has been set.

### GetLicenseMicroservice

`func (o *ReplaceConfigurationRequest) GetLicenseMicroservice() LicenseMicroservice`

GetLicenseMicroservice returns the LicenseMicroservice field if non-nil, zero value otherwise.

### GetLicenseMicroserviceOk

`func (o *ReplaceConfigurationRequest) GetLicenseMicroserviceOk() (*LicenseMicroservice, bool)`

GetLicenseMicroserviceOk returns a tuple with the LicenseMicroservice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseMicroservice

`func (o *ReplaceConfigurationRequest) SetLicenseMicroservice(v LicenseMicroservice)`

SetLicenseMicroservice sets LicenseMicroservice field to given value.

### HasLicenseMicroservice

`func (o *ReplaceConfigurationRequest) HasLicenseMicroservice() bool`

HasLicenseMicroservice returns a boolean if a field has been set.

### GetMicroservices

`func (o *ReplaceConfigurationRequest) GetMicroservices() Microservices`

GetMicroservices returns the Microservices field if non-nil, zero value otherwise.

### GetMicroservicesOk

`func (o *ReplaceConfigurationRequest) GetMicroservicesOk() (*Microservices, bool)`

GetMicroservicesOk returns a tuple with the Microservices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroservices

`func (o *ReplaceConfigurationRequest) SetMicroservices(v Microservices)`

SetMicroservices sets Microservices field to given value.

### HasMicroservices

`func (o *ReplaceConfigurationRequest) HasMicroservices() bool`

HasMicroservices returns a boolean if a field has been set.

### GetRateLimiting

`func (o *ReplaceConfigurationRequest) GetRateLimiting() RateLimitingConfig`

GetRateLimiting returns the RateLimiting field if non-nil, zero value otherwise.

### GetRateLimitingOk

`func (o *ReplaceConfigurationRequest) GetRateLimitingOk() (*RateLimitingConfig, bool)`

GetRateLimitingOk returns a tuple with the RateLimiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimiting

`func (o *ReplaceConfigurationRequest) SetRateLimiting(v RateLimitingConfig)`

SetRateLimiting sets RateLimiting field to given value.

### HasRateLimiting

`func (o *ReplaceConfigurationRequest) HasRateLimiting() bool`

HasRateLimiting returns a boolean if a field has been set.

### GetImageDownloadHttpProxy

`func (o *ReplaceConfigurationRequest) GetImageDownloadHttpProxy() string`

GetImageDownloadHttpProxy returns the ImageDownloadHttpProxy field if non-nil, zero value otherwise.

### GetImageDownloadHttpProxyOk

`func (o *ReplaceConfigurationRequest) GetImageDownloadHttpProxyOk() (*string, bool)`

GetImageDownloadHttpProxyOk returns a tuple with the ImageDownloadHttpProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDownloadHttpProxy

`func (o *ReplaceConfigurationRequest) SetImageDownloadHttpProxy(v string)`

SetImageDownloadHttpProxy sets ImageDownloadHttpProxy field to given value.

### HasImageDownloadHttpProxy

`func (o *ReplaceConfigurationRequest) HasImageDownloadHttpProxy() bool`

HasImageDownloadHttpProxy returns a boolean if a field has been set.

### GetImageDownloadHttpsProxy

`func (o *ReplaceConfigurationRequest) GetImageDownloadHttpsProxy() string`

GetImageDownloadHttpsProxy returns the ImageDownloadHttpsProxy field if non-nil, zero value otherwise.

### GetImageDownloadHttpsProxyOk

`func (o *ReplaceConfigurationRequest) GetImageDownloadHttpsProxyOk() (*string, bool)`

GetImageDownloadHttpsProxyOk returns a tuple with the ImageDownloadHttpsProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDownloadHttpsProxy

`func (o *ReplaceConfigurationRequest) SetImageDownloadHttpsProxy(v string)`

SetImageDownloadHttpsProxy sets ImageDownloadHttpsProxy field to given value.

### HasImageDownloadHttpsProxy

`func (o *ReplaceConfigurationRequest) HasImageDownloadHttpsProxy() bool`

HasImageDownloadHttpsProxy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


