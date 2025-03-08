# Infrastructure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Label of the Infrastructure. | 
**CustomVariables** | Pointer to **map[string]interface{}** | Custom variables in JSON format. | [optional] 
**UserIdOwner** | Pointer to **float32** | User ID of the owner of the Infrastructure. | [optional] 
**Subdomain** | Pointer to **string** | Subdomain associated with the Infrastructure. | [optional] 
**InstancesCountActive** | Pointer to **float32** | Number of active instances. | [optional] 
**DrivesCountActive** | Pointer to **float32** | Number of active drives. | [optional] 
**Ipv4SubnetsCountActive** | Pointer to **float32** | Number of active IPv4 subnets. | [optional] 
**Ipv6SubnetsCountActive** | Pointer to **float32** | Number of active IPv6 subnets. | [optional] 
**Ipv4UnusedIpAddresses** | Pointer to **float32** | Number of unused IPv4 addresses. | [optional] 
**Description** | Pointer to **string** | Description of the infrastructure. | [optional] 
**Settings** | Pointer to **map[string]interface{}** | Settings in JSON format. | [optional] 
**IsApiPrivate** | Pointer to **float32** | Whether the infrastructure API is private. | [optional] 
**ExperimentalPriority** | Pointer to **string** | Experimental priority. | [optional] 
**IsPublicDesignsMember** | Pointer to **float32** | Whether the infrastructure is a member of public designs. | [optional] 
**CertificatesJson** | Pointer to **map[string]interface{}** | Certificates in JSON format. | [optional] 
**DeployCookieJarJson** | Pointer to **map[string]interface{}** | Deploy cookie jar JSON. | [optional] 
**DeferredDeployAttemptLastErrorJson** | Pointer to **map[string]interface{}** | Last error of deferred deploy attempt. | [optional] 
**IsAutomanaged** | Pointer to **float32** | Whether the infrastructure is automanaged. | [optional] 
**UpdatedTimestamp** | **string** | Timestamp of the latest update for the Infrastructure. | 
**Id** | **float32** | Infrastructure Id | 
**Revision** | **float32** | Revision of the Infrastructure | 
**ServiceStatus** | **string** | Service status of the Infrastructure | 
**DatacenterName** | **string** | Datacenter name where the Infrastructure is located. | 
**SiteId** | **float32** | The ID of the site where the Infrastructure is located. | 
**CreatedTimestamp** | **string** | Timestamp of the Infrastructure creation. | 
**SubdomainPermanent** | Pointer to **string** | Permanent subdomain associated with the Infrastructure. | [optional] 
**DnsSubdomainId** | Pointer to **float32** | DNS Subdomain ID. | [optional] 
**DnsSubdomainPermanentId** | Pointer to **float32** | Permanent DNS Subdomain ID. | [optional] 
**DesignIsLocked** | **float32** | Infrastructure design locked flag. | 
**Config** | [**InfrastructureConfig**](InfrastructureConfig.md) | The current changes to be deployed for the Infrastructure. | 
**Meta** | Pointer to [**InfrastructureMeta**](InfrastructureMeta.md) | Meta information for the Infrastructure | [optional] 
**Statistics** | Pointer to [**JobGroupStatisticsWithoutId**](JobGroupStatisticsWithoutId.md) |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Links to other resources | [optional] 

## Methods

### NewInfrastructure

`func NewInfrastructure(label string, updatedTimestamp string, id float32, revision float32, serviceStatus string, datacenterName string, siteId float32, createdTimestamp string, designIsLocked float32, config InfrastructureConfig, ) *Infrastructure`

NewInfrastructure instantiates a new Infrastructure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureWithDefaults

`func NewInfrastructureWithDefaults() *Infrastructure`

NewInfrastructureWithDefaults instantiates a new Infrastructure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *Infrastructure) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Infrastructure) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Infrastructure) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetCustomVariables

`func (o *Infrastructure) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *Infrastructure) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *Infrastructure) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *Infrastructure) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.

### GetUserIdOwner

`func (o *Infrastructure) GetUserIdOwner() float32`

GetUserIdOwner returns the UserIdOwner field if non-nil, zero value otherwise.

### GetUserIdOwnerOk

`func (o *Infrastructure) GetUserIdOwnerOk() (*float32, bool)`

GetUserIdOwnerOk returns a tuple with the UserIdOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdOwner

`func (o *Infrastructure) SetUserIdOwner(v float32)`

SetUserIdOwner sets UserIdOwner field to given value.

### HasUserIdOwner

`func (o *Infrastructure) HasUserIdOwner() bool`

HasUserIdOwner returns a boolean if a field has been set.

### GetSubdomain

`func (o *Infrastructure) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *Infrastructure) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *Infrastructure) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *Infrastructure) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetInstancesCountActive

`func (o *Infrastructure) GetInstancesCountActive() float32`

GetInstancesCountActive returns the InstancesCountActive field if non-nil, zero value otherwise.

### GetInstancesCountActiveOk

`func (o *Infrastructure) GetInstancesCountActiveOk() (*float32, bool)`

GetInstancesCountActiveOk returns a tuple with the InstancesCountActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancesCountActive

`func (o *Infrastructure) SetInstancesCountActive(v float32)`

SetInstancesCountActive sets InstancesCountActive field to given value.

### HasInstancesCountActive

`func (o *Infrastructure) HasInstancesCountActive() bool`

HasInstancesCountActive returns a boolean if a field has been set.

### GetDrivesCountActive

`func (o *Infrastructure) GetDrivesCountActive() float32`

GetDrivesCountActive returns the DrivesCountActive field if non-nil, zero value otherwise.

### GetDrivesCountActiveOk

`func (o *Infrastructure) GetDrivesCountActiveOk() (*float32, bool)`

GetDrivesCountActiveOk returns a tuple with the DrivesCountActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivesCountActive

`func (o *Infrastructure) SetDrivesCountActive(v float32)`

SetDrivesCountActive sets DrivesCountActive field to given value.

### HasDrivesCountActive

`func (o *Infrastructure) HasDrivesCountActive() bool`

HasDrivesCountActive returns a boolean if a field has been set.

### GetIpv4SubnetsCountActive

`func (o *Infrastructure) GetIpv4SubnetsCountActive() float32`

GetIpv4SubnetsCountActive returns the Ipv4SubnetsCountActive field if non-nil, zero value otherwise.

### GetIpv4SubnetsCountActiveOk

`func (o *Infrastructure) GetIpv4SubnetsCountActiveOk() (*float32, bool)`

GetIpv4SubnetsCountActiveOk returns a tuple with the Ipv4SubnetsCountActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4SubnetsCountActive

`func (o *Infrastructure) SetIpv4SubnetsCountActive(v float32)`

SetIpv4SubnetsCountActive sets Ipv4SubnetsCountActive field to given value.

### HasIpv4SubnetsCountActive

`func (o *Infrastructure) HasIpv4SubnetsCountActive() bool`

HasIpv4SubnetsCountActive returns a boolean if a field has been set.

### GetIpv6SubnetsCountActive

`func (o *Infrastructure) GetIpv6SubnetsCountActive() float32`

GetIpv6SubnetsCountActive returns the Ipv6SubnetsCountActive field if non-nil, zero value otherwise.

### GetIpv6SubnetsCountActiveOk

`func (o *Infrastructure) GetIpv6SubnetsCountActiveOk() (*float32, bool)`

GetIpv6SubnetsCountActiveOk returns a tuple with the Ipv6SubnetsCountActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6SubnetsCountActive

`func (o *Infrastructure) SetIpv6SubnetsCountActive(v float32)`

SetIpv6SubnetsCountActive sets Ipv6SubnetsCountActive field to given value.

### HasIpv6SubnetsCountActive

`func (o *Infrastructure) HasIpv6SubnetsCountActive() bool`

HasIpv6SubnetsCountActive returns a boolean if a field has been set.

### GetIpv4UnusedIpAddresses

`func (o *Infrastructure) GetIpv4UnusedIpAddresses() float32`

GetIpv4UnusedIpAddresses returns the Ipv4UnusedIpAddresses field if non-nil, zero value otherwise.

### GetIpv4UnusedIpAddressesOk

`func (o *Infrastructure) GetIpv4UnusedIpAddressesOk() (*float32, bool)`

GetIpv4UnusedIpAddressesOk returns a tuple with the Ipv4UnusedIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4UnusedIpAddresses

`func (o *Infrastructure) SetIpv4UnusedIpAddresses(v float32)`

SetIpv4UnusedIpAddresses sets Ipv4UnusedIpAddresses field to given value.

### HasIpv4UnusedIpAddresses

`func (o *Infrastructure) HasIpv4UnusedIpAddresses() bool`

HasIpv4UnusedIpAddresses returns a boolean if a field has been set.

### GetDescription

`func (o *Infrastructure) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Infrastructure) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Infrastructure) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Infrastructure) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSettings

`func (o *Infrastructure) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *Infrastructure) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *Infrastructure) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *Infrastructure) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetIsApiPrivate

`func (o *Infrastructure) GetIsApiPrivate() float32`

GetIsApiPrivate returns the IsApiPrivate field if non-nil, zero value otherwise.

### GetIsApiPrivateOk

`func (o *Infrastructure) GetIsApiPrivateOk() (*float32, bool)`

GetIsApiPrivateOk returns a tuple with the IsApiPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApiPrivate

`func (o *Infrastructure) SetIsApiPrivate(v float32)`

SetIsApiPrivate sets IsApiPrivate field to given value.

### HasIsApiPrivate

`func (o *Infrastructure) HasIsApiPrivate() bool`

HasIsApiPrivate returns a boolean if a field has been set.

### GetExperimentalPriority

`func (o *Infrastructure) GetExperimentalPriority() string`

GetExperimentalPriority returns the ExperimentalPriority field if non-nil, zero value otherwise.

### GetExperimentalPriorityOk

`func (o *Infrastructure) GetExperimentalPriorityOk() (*string, bool)`

GetExperimentalPriorityOk returns a tuple with the ExperimentalPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentalPriority

`func (o *Infrastructure) SetExperimentalPriority(v string)`

SetExperimentalPriority sets ExperimentalPriority field to given value.

### HasExperimentalPriority

`func (o *Infrastructure) HasExperimentalPriority() bool`

HasExperimentalPriority returns a boolean if a field has been set.

### GetIsPublicDesignsMember

`func (o *Infrastructure) GetIsPublicDesignsMember() float32`

GetIsPublicDesignsMember returns the IsPublicDesignsMember field if non-nil, zero value otherwise.

### GetIsPublicDesignsMemberOk

`func (o *Infrastructure) GetIsPublicDesignsMemberOk() (*float32, bool)`

GetIsPublicDesignsMemberOk returns a tuple with the IsPublicDesignsMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublicDesignsMember

`func (o *Infrastructure) SetIsPublicDesignsMember(v float32)`

SetIsPublicDesignsMember sets IsPublicDesignsMember field to given value.

### HasIsPublicDesignsMember

`func (o *Infrastructure) HasIsPublicDesignsMember() bool`

HasIsPublicDesignsMember returns a boolean if a field has been set.

### GetCertificatesJson

`func (o *Infrastructure) GetCertificatesJson() map[string]interface{}`

GetCertificatesJson returns the CertificatesJson field if non-nil, zero value otherwise.

### GetCertificatesJsonOk

`func (o *Infrastructure) GetCertificatesJsonOk() (*map[string]interface{}, bool)`

GetCertificatesJsonOk returns a tuple with the CertificatesJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatesJson

`func (o *Infrastructure) SetCertificatesJson(v map[string]interface{})`

SetCertificatesJson sets CertificatesJson field to given value.

### HasCertificatesJson

`func (o *Infrastructure) HasCertificatesJson() bool`

HasCertificatesJson returns a boolean if a field has been set.

### GetDeployCookieJarJson

`func (o *Infrastructure) GetDeployCookieJarJson() map[string]interface{}`

GetDeployCookieJarJson returns the DeployCookieJarJson field if non-nil, zero value otherwise.

### GetDeployCookieJarJsonOk

`func (o *Infrastructure) GetDeployCookieJarJsonOk() (*map[string]interface{}, bool)`

GetDeployCookieJarJsonOk returns a tuple with the DeployCookieJarJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployCookieJarJson

`func (o *Infrastructure) SetDeployCookieJarJson(v map[string]interface{})`

SetDeployCookieJarJson sets DeployCookieJarJson field to given value.

### HasDeployCookieJarJson

`func (o *Infrastructure) HasDeployCookieJarJson() bool`

HasDeployCookieJarJson returns a boolean if a field has been set.

### GetDeferredDeployAttemptLastErrorJson

`func (o *Infrastructure) GetDeferredDeployAttemptLastErrorJson() map[string]interface{}`

GetDeferredDeployAttemptLastErrorJson returns the DeferredDeployAttemptLastErrorJson field if non-nil, zero value otherwise.

### GetDeferredDeployAttemptLastErrorJsonOk

`func (o *Infrastructure) GetDeferredDeployAttemptLastErrorJsonOk() (*map[string]interface{}, bool)`

GetDeferredDeployAttemptLastErrorJsonOk returns a tuple with the DeferredDeployAttemptLastErrorJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferredDeployAttemptLastErrorJson

`func (o *Infrastructure) SetDeferredDeployAttemptLastErrorJson(v map[string]interface{})`

SetDeferredDeployAttemptLastErrorJson sets DeferredDeployAttemptLastErrorJson field to given value.

### HasDeferredDeployAttemptLastErrorJson

`func (o *Infrastructure) HasDeferredDeployAttemptLastErrorJson() bool`

HasDeferredDeployAttemptLastErrorJson returns a boolean if a field has been set.

### GetIsAutomanaged

`func (o *Infrastructure) GetIsAutomanaged() float32`

GetIsAutomanaged returns the IsAutomanaged field if non-nil, zero value otherwise.

### GetIsAutomanagedOk

`func (o *Infrastructure) GetIsAutomanagedOk() (*float32, bool)`

GetIsAutomanagedOk returns a tuple with the IsAutomanaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomanaged

`func (o *Infrastructure) SetIsAutomanaged(v float32)`

SetIsAutomanaged sets IsAutomanaged field to given value.

### HasIsAutomanaged

`func (o *Infrastructure) HasIsAutomanaged() bool`

HasIsAutomanaged returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *Infrastructure) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *Infrastructure) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *Infrastructure) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetId

`func (o *Infrastructure) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Infrastructure) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Infrastructure) SetId(v float32)`

SetId sets Id field to given value.


### GetRevision

`func (o *Infrastructure) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *Infrastructure) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *Infrastructure) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetServiceStatus

`func (o *Infrastructure) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *Infrastructure) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *Infrastructure) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetDatacenterName

`func (o *Infrastructure) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *Infrastructure) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *Infrastructure) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.


### GetSiteId

`func (o *Infrastructure) GetSiteId() float32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Infrastructure) GetSiteIdOk() (*float32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Infrastructure) SetSiteId(v float32)`

SetSiteId sets SiteId field to given value.


### GetCreatedTimestamp

`func (o *Infrastructure) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *Infrastructure) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *Infrastructure) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetSubdomainPermanent

`func (o *Infrastructure) GetSubdomainPermanent() string`

GetSubdomainPermanent returns the SubdomainPermanent field if non-nil, zero value otherwise.

### GetSubdomainPermanentOk

`func (o *Infrastructure) GetSubdomainPermanentOk() (*string, bool)`

GetSubdomainPermanentOk returns a tuple with the SubdomainPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomainPermanent

`func (o *Infrastructure) SetSubdomainPermanent(v string)`

SetSubdomainPermanent sets SubdomainPermanent field to given value.

### HasSubdomainPermanent

`func (o *Infrastructure) HasSubdomainPermanent() bool`

HasSubdomainPermanent returns a boolean if a field has been set.

### GetDnsSubdomainId

`func (o *Infrastructure) GetDnsSubdomainId() float32`

GetDnsSubdomainId returns the DnsSubdomainId field if non-nil, zero value otherwise.

### GetDnsSubdomainIdOk

`func (o *Infrastructure) GetDnsSubdomainIdOk() (*float32, bool)`

GetDnsSubdomainIdOk returns a tuple with the DnsSubdomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainId

`func (o *Infrastructure) SetDnsSubdomainId(v float32)`

SetDnsSubdomainId sets DnsSubdomainId field to given value.

### HasDnsSubdomainId

`func (o *Infrastructure) HasDnsSubdomainId() bool`

HasDnsSubdomainId returns a boolean if a field has been set.

### GetDnsSubdomainPermanentId

`func (o *Infrastructure) GetDnsSubdomainPermanentId() float32`

GetDnsSubdomainPermanentId returns the DnsSubdomainPermanentId field if non-nil, zero value otherwise.

### GetDnsSubdomainPermanentIdOk

`func (o *Infrastructure) GetDnsSubdomainPermanentIdOk() (*float32, bool)`

GetDnsSubdomainPermanentIdOk returns a tuple with the DnsSubdomainPermanentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainPermanentId

`func (o *Infrastructure) SetDnsSubdomainPermanentId(v float32)`

SetDnsSubdomainPermanentId sets DnsSubdomainPermanentId field to given value.

### HasDnsSubdomainPermanentId

`func (o *Infrastructure) HasDnsSubdomainPermanentId() bool`

HasDnsSubdomainPermanentId returns a boolean if a field has been set.

### GetDesignIsLocked

`func (o *Infrastructure) GetDesignIsLocked() float32`

GetDesignIsLocked returns the DesignIsLocked field if non-nil, zero value otherwise.

### GetDesignIsLockedOk

`func (o *Infrastructure) GetDesignIsLockedOk() (*float32, bool)`

GetDesignIsLockedOk returns a tuple with the DesignIsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesignIsLocked

`func (o *Infrastructure) SetDesignIsLocked(v float32)`

SetDesignIsLocked sets DesignIsLocked field to given value.


### GetConfig

`func (o *Infrastructure) GetConfig() InfrastructureConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Infrastructure) GetConfigOk() (*InfrastructureConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Infrastructure) SetConfig(v InfrastructureConfig)`

SetConfig sets Config field to given value.


### GetMeta

`func (o *Infrastructure) GetMeta() InfrastructureMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Infrastructure) GetMetaOk() (*InfrastructureMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Infrastructure) SetMeta(v InfrastructureMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Infrastructure) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetStatistics

`func (o *Infrastructure) GetStatistics() JobGroupStatisticsWithoutId`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *Infrastructure) GetStatisticsOk() (*JobGroupStatisticsWithoutId, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *Infrastructure) SetStatistics(v JobGroupStatisticsWithoutId)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *Infrastructure) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetLinks

`func (o *Infrastructure) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Infrastructure) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Infrastructure) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Infrastructure) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


