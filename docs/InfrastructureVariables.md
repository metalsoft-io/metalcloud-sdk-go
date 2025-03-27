# InfrastructureVariables

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** |  | 
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
**CertificatesJson** | Pointer to **string** | Certificates in JSON format. | [optional] 
**DeployCookieJarJson** | Pointer to **map[string]interface{}** | Deploy cookie jar JSON. | [optional] 
**DeferredDeployAttemptLastErrorJson** | Pointer to **map[string]interface{}** | Last error of deferred deploy attempt. | [optional] 
**IsAutomanaged** | Pointer to **float32** | Whether the infrastructure is automanaged. | [optional] 
**UpdatedTimestamp** | **string** | Timestamp of the latest update for the Infrastructure. | 
**Id** | **float32** | Infrastructure Id | 
**Revision** | **float32** | Revision of the Infrastructure | 
**ServiceStatus** | [**GenericServiceStatus**](GenericServiceStatus.md) | Service status of the Infrastructure | 
**DatacenterName** | **string** | Datacenter name where the Infrastructure is located. | 
**SiteId** | **float32** | The ID of the site where the Infrastructure is located. | 
**CreatedTimestamp** | **string** | Timestamp of the Infrastructure creation. | 
**SubdomainPermanent** | Pointer to **string** | Permanent subdomain associated with the Infrastructure. | [optional] 
**DnsSubdomainId** | Pointer to **float32** | DNS Subdomain ID. | [optional] 
**DnsSubdomainPermanentId** | Pointer to **float32** | Permanent DNS Subdomain ID. | [optional] 
**DesignIsLocked** | **float32** | Infrastructure design locked flag. | 
**Config** | [**InfrastructureConfig**](InfrastructureConfig.md) | The current changes to be deployed for the Infrastructure. | 

## Methods

### NewInfrastructureVariables

`func NewInfrastructureVariables(label string, updatedTimestamp string, id float32, revision float32, serviceStatus GenericServiceStatus, datacenterName string, siteId float32, createdTimestamp string, designIsLocked float32, config InfrastructureConfig, ) *InfrastructureVariables`

NewInfrastructureVariables instantiates a new InfrastructureVariables object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureVariablesWithDefaults

`func NewInfrastructureVariablesWithDefaults() *InfrastructureVariables`

NewInfrastructureVariablesWithDefaults instantiates a new InfrastructureVariables object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *InfrastructureVariables) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *InfrastructureVariables) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *InfrastructureVariables) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetCustomVariables

`func (o *InfrastructureVariables) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *InfrastructureVariables) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *InfrastructureVariables) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *InfrastructureVariables) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.

### GetUserIdOwner

`func (o *InfrastructureVariables) GetUserIdOwner() float32`

GetUserIdOwner returns the UserIdOwner field if non-nil, zero value otherwise.

### GetUserIdOwnerOk

`func (o *InfrastructureVariables) GetUserIdOwnerOk() (*float32, bool)`

GetUserIdOwnerOk returns a tuple with the UserIdOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdOwner

`func (o *InfrastructureVariables) SetUserIdOwner(v float32)`

SetUserIdOwner sets UserIdOwner field to given value.

### HasUserIdOwner

`func (o *InfrastructureVariables) HasUserIdOwner() bool`

HasUserIdOwner returns a boolean if a field has been set.

### GetSubdomain

`func (o *InfrastructureVariables) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *InfrastructureVariables) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *InfrastructureVariables) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *InfrastructureVariables) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetInstancesCountActive

`func (o *InfrastructureVariables) GetInstancesCountActive() float32`

GetInstancesCountActive returns the InstancesCountActive field if non-nil, zero value otherwise.

### GetInstancesCountActiveOk

`func (o *InfrastructureVariables) GetInstancesCountActiveOk() (*float32, bool)`

GetInstancesCountActiveOk returns a tuple with the InstancesCountActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancesCountActive

`func (o *InfrastructureVariables) SetInstancesCountActive(v float32)`

SetInstancesCountActive sets InstancesCountActive field to given value.

### HasInstancesCountActive

`func (o *InfrastructureVariables) HasInstancesCountActive() bool`

HasInstancesCountActive returns a boolean if a field has been set.

### GetDrivesCountActive

`func (o *InfrastructureVariables) GetDrivesCountActive() float32`

GetDrivesCountActive returns the DrivesCountActive field if non-nil, zero value otherwise.

### GetDrivesCountActiveOk

`func (o *InfrastructureVariables) GetDrivesCountActiveOk() (*float32, bool)`

GetDrivesCountActiveOk returns a tuple with the DrivesCountActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivesCountActive

`func (o *InfrastructureVariables) SetDrivesCountActive(v float32)`

SetDrivesCountActive sets DrivesCountActive field to given value.

### HasDrivesCountActive

`func (o *InfrastructureVariables) HasDrivesCountActive() bool`

HasDrivesCountActive returns a boolean if a field has been set.

### GetIpv4SubnetsCountActive

`func (o *InfrastructureVariables) GetIpv4SubnetsCountActive() float32`

GetIpv4SubnetsCountActive returns the Ipv4SubnetsCountActive field if non-nil, zero value otherwise.

### GetIpv4SubnetsCountActiveOk

`func (o *InfrastructureVariables) GetIpv4SubnetsCountActiveOk() (*float32, bool)`

GetIpv4SubnetsCountActiveOk returns a tuple with the Ipv4SubnetsCountActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4SubnetsCountActive

`func (o *InfrastructureVariables) SetIpv4SubnetsCountActive(v float32)`

SetIpv4SubnetsCountActive sets Ipv4SubnetsCountActive field to given value.

### HasIpv4SubnetsCountActive

`func (o *InfrastructureVariables) HasIpv4SubnetsCountActive() bool`

HasIpv4SubnetsCountActive returns a boolean if a field has been set.

### GetIpv6SubnetsCountActive

`func (o *InfrastructureVariables) GetIpv6SubnetsCountActive() float32`

GetIpv6SubnetsCountActive returns the Ipv6SubnetsCountActive field if non-nil, zero value otherwise.

### GetIpv6SubnetsCountActiveOk

`func (o *InfrastructureVariables) GetIpv6SubnetsCountActiveOk() (*float32, bool)`

GetIpv6SubnetsCountActiveOk returns a tuple with the Ipv6SubnetsCountActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6SubnetsCountActive

`func (o *InfrastructureVariables) SetIpv6SubnetsCountActive(v float32)`

SetIpv6SubnetsCountActive sets Ipv6SubnetsCountActive field to given value.

### HasIpv6SubnetsCountActive

`func (o *InfrastructureVariables) HasIpv6SubnetsCountActive() bool`

HasIpv6SubnetsCountActive returns a boolean if a field has been set.

### GetIpv4UnusedIpAddresses

`func (o *InfrastructureVariables) GetIpv4UnusedIpAddresses() float32`

GetIpv4UnusedIpAddresses returns the Ipv4UnusedIpAddresses field if non-nil, zero value otherwise.

### GetIpv4UnusedIpAddressesOk

`func (o *InfrastructureVariables) GetIpv4UnusedIpAddressesOk() (*float32, bool)`

GetIpv4UnusedIpAddressesOk returns a tuple with the Ipv4UnusedIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4UnusedIpAddresses

`func (o *InfrastructureVariables) SetIpv4UnusedIpAddresses(v float32)`

SetIpv4UnusedIpAddresses sets Ipv4UnusedIpAddresses field to given value.

### HasIpv4UnusedIpAddresses

`func (o *InfrastructureVariables) HasIpv4UnusedIpAddresses() bool`

HasIpv4UnusedIpAddresses returns a boolean if a field has been set.

### GetDescription

`func (o *InfrastructureVariables) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InfrastructureVariables) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InfrastructureVariables) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InfrastructureVariables) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSettings

`func (o *InfrastructureVariables) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *InfrastructureVariables) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *InfrastructureVariables) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *InfrastructureVariables) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetIsApiPrivate

`func (o *InfrastructureVariables) GetIsApiPrivate() float32`

GetIsApiPrivate returns the IsApiPrivate field if non-nil, zero value otherwise.

### GetIsApiPrivateOk

`func (o *InfrastructureVariables) GetIsApiPrivateOk() (*float32, bool)`

GetIsApiPrivateOk returns a tuple with the IsApiPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApiPrivate

`func (o *InfrastructureVariables) SetIsApiPrivate(v float32)`

SetIsApiPrivate sets IsApiPrivate field to given value.

### HasIsApiPrivate

`func (o *InfrastructureVariables) HasIsApiPrivate() bool`

HasIsApiPrivate returns a boolean if a field has been set.

### GetExperimentalPriority

`func (o *InfrastructureVariables) GetExperimentalPriority() string`

GetExperimentalPriority returns the ExperimentalPriority field if non-nil, zero value otherwise.

### GetExperimentalPriorityOk

`func (o *InfrastructureVariables) GetExperimentalPriorityOk() (*string, bool)`

GetExperimentalPriorityOk returns a tuple with the ExperimentalPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentalPriority

`func (o *InfrastructureVariables) SetExperimentalPriority(v string)`

SetExperimentalPriority sets ExperimentalPriority field to given value.

### HasExperimentalPriority

`func (o *InfrastructureVariables) HasExperimentalPriority() bool`

HasExperimentalPriority returns a boolean if a field has been set.

### GetIsPublicDesignsMember

`func (o *InfrastructureVariables) GetIsPublicDesignsMember() float32`

GetIsPublicDesignsMember returns the IsPublicDesignsMember field if non-nil, zero value otherwise.

### GetIsPublicDesignsMemberOk

`func (o *InfrastructureVariables) GetIsPublicDesignsMemberOk() (*float32, bool)`

GetIsPublicDesignsMemberOk returns a tuple with the IsPublicDesignsMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublicDesignsMember

`func (o *InfrastructureVariables) SetIsPublicDesignsMember(v float32)`

SetIsPublicDesignsMember sets IsPublicDesignsMember field to given value.

### HasIsPublicDesignsMember

`func (o *InfrastructureVariables) HasIsPublicDesignsMember() bool`

HasIsPublicDesignsMember returns a boolean if a field has been set.

### GetCertificatesJson

`func (o *InfrastructureVariables) GetCertificatesJson() string`

GetCertificatesJson returns the CertificatesJson field if non-nil, zero value otherwise.

### GetCertificatesJsonOk

`func (o *InfrastructureVariables) GetCertificatesJsonOk() (*string, bool)`

GetCertificatesJsonOk returns a tuple with the CertificatesJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatesJson

`func (o *InfrastructureVariables) SetCertificatesJson(v string)`

SetCertificatesJson sets CertificatesJson field to given value.

### HasCertificatesJson

`func (o *InfrastructureVariables) HasCertificatesJson() bool`

HasCertificatesJson returns a boolean if a field has been set.

### GetDeployCookieJarJson

`func (o *InfrastructureVariables) GetDeployCookieJarJson() map[string]interface{}`

GetDeployCookieJarJson returns the DeployCookieJarJson field if non-nil, zero value otherwise.

### GetDeployCookieJarJsonOk

`func (o *InfrastructureVariables) GetDeployCookieJarJsonOk() (*map[string]interface{}, bool)`

GetDeployCookieJarJsonOk returns a tuple with the DeployCookieJarJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployCookieJarJson

`func (o *InfrastructureVariables) SetDeployCookieJarJson(v map[string]interface{})`

SetDeployCookieJarJson sets DeployCookieJarJson field to given value.

### HasDeployCookieJarJson

`func (o *InfrastructureVariables) HasDeployCookieJarJson() bool`

HasDeployCookieJarJson returns a boolean if a field has been set.

### GetDeferredDeployAttemptLastErrorJson

`func (o *InfrastructureVariables) GetDeferredDeployAttemptLastErrorJson() map[string]interface{}`

GetDeferredDeployAttemptLastErrorJson returns the DeferredDeployAttemptLastErrorJson field if non-nil, zero value otherwise.

### GetDeferredDeployAttemptLastErrorJsonOk

`func (o *InfrastructureVariables) GetDeferredDeployAttemptLastErrorJsonOk() (*map[string]interface{}, bool)`

GetDeferredDeployAttemptLastErrorJsonOk returns a tuple with the DeferredDeployAttemptLastErrorJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferredDeployAttemptLastErrorJson

`func (o *InfrastructureVariables) SetDeferredDeployAttemptLastErrorJson(v map[string]interface{})`

SetDeferredDeployAttemptLastErrorJson sets DeferredDeployAttemptLastErrorJson field to given value.

### HasDeferredDeployAttemptLastErrorJson

`func (o *InfrastructureVariables) HasDeferredDeployAttemptLastErrorJson() bool`

HasDeferredDeployAttemptLastErrorJson returns a boolean if a field has been set.

### GetIsAutomanaged

`func (o *InfrastructureVariables) GetIsAutomanaged() float32`

GetIsAutomanaged returns the IsAutomanaged field if non-nil, zero value otherwise.

### GetIsAutomanagedOk

`func (o *InfrastructureVariables) GetIsAutomanagedOk() (*float32, bool)`

GetIsAutomanagedOk returns a tuple with the IsAutomanaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomanaged

`func (o *InfrastructureVariables) SetIsAutomanaged(v float32)`

SetIsAutomanaged sets IsAutomanaged field to given value.

### HasIsAutomanaged

`func (o *InfrastructureVariables) HasIsAutomanaged() bool`

HasIsAutomanaged returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *InfrastructureVariables) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *InfrastructureVariables) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *InfrastructureVariables) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetId

`func (o *InfrastructureVariables) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InfrastructureVariables) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InfrastructureVariables) SetId(v float32)`

SetId sets Id field to given value.


### GetRevision

`func (o *InfrastructureVariables) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *InfrastructureVariables) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *InfrastructureVariables) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetServiceStatus

`func (o *InfrastructureVariables) GetServiceStatus() GenericServiceStatus`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *InfrastructureVariables) GetServiceStatusOk() (*GenericServiceStatus, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *InfrastructureVariables) SetServiceStatus(v GenericServiceStatus)`

SetServiceStatus sets ServiceStatus field to given value.


### GetDatacenterName

`func (o *InfrastructureVariables) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *InfrastructureVariables) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *InfrastructureVariables) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.


### GetSiteId

`func (o *InfrastructureVariables) GetSiteId() float32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *InfrastructureVariables) GetSiteIdOk() (*float32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *InfrastructureVariables) SetSiteId(v float32)`

SetSiteId sets SiteId field to given value.


### GetCreatedTimestamp

`func (o *InfrastructureVariables) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *InfrastructureVariables) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *InfrastructureVariables) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetSubdomainPermanent

`func (o *InfrastructureVariables) GetSubdomainPermanent() string`

GetSubdomainPermanent returns the SubdomainPermanent field if non-nil, zero value otherwise.

### GetSubdomainPermanentOk

`func (o *InfrastructureVariables) GetSubdomainPermanentOk() (*string, bool)`

GetSubdomainPermanentOk returns a tuple with the SubdomainPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomainPermanent

`func (o *InfrastructureVariables) SetSubdomainPermanent(v string)`

SetSubdomainPermanent sets SubdomainPermanent field to given value.

### HasSubdomainPermanent

`func (o *InfrastructureVariables) HasSubdomainPermanent() bool`

HasSubdomainPermanent returns a boolean if a field has been set.

### GetDnsSubdomainId

`func (o *InfrastructureVariables) GetDnsSubdomainId() float32`

GetDnsSubdomainId returns the DnsSubdomainId field if non-nil, zero value otherwise.

### GetDnsSubdomainIdOk

`func (o *InfrastructureVariables) GetDnsSubdomainIdOk() (*float32, bool)`

GetDnsSubdomainIdOk returns a tuple with the DnsSubdomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainId

`func (o *InfrastructureVariables) SetDnsSubdomainId(v float32)`

SetDnsSubdomainId sets DnsSubdomainId field to given value.

### HasDnsSubdomainId

`func (o *InfrastructureVariables) HasDnsSubdomainId() bool`

HasDnsSubdomainId returns a boolean if a field has been set.

### GetDnsSubdomainPermanentId

`func (o *InfrastructureVariables) GetDnsSubdomainPermanentId() float32`

GetDnsSubdomainPermanentId returns the DnsSubdomainPermanentId field if non-nil, zero value otherwise.

### GetDnsSubdomainPermanentIdOk

`func (o *InfrastructureVariables) GetDnsSubdomainPermanentIdOk() (*float32, bool)`

GetDnsSubdomainPermanentIdOk returns a tuple with the DnsSubdomainPermanentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainPermanentId

`func (o *InfrastructureVariables) SetDnsSubdomainPermanentId(v float32)`

SetDnsSubdomainPermanentId sets DnsSubdomainPermanentId field to given value.

### HasDnsSubdomainPermanentId

`func (o *InfrastructureVariables) HasDnsSubdomainPermanentId() bool`

HasDnsSubdomainPermanentId returns a boolean if a field has been set.

### GetDesignIsLocked

`func (o *InfrastructureVariables) GetDesignIsLocked() float32`

GetDesignIsLocked returns the DesignIsLocked field if non-nil, zero value otherwise.

### GetDesignIsLockedOk

`func (o *InfrastructureVariables) GetDesignIsLockedOk() (*float32, bool)`

GetDesignIsLockedOk returns a tuple with the DesignIsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesignIsLocked

`func (o *InfrastructureVariables) SetDesignIsLocked(v float32)`

SetDesignIsLocked sets DesignIsLocked field to given value.


### GetConfig

`func (o *InfrastructureVariables) GetConfig() InfrastructureConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *InfrastructureVariables) GetConfigOk() (*InfrastructureConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *InfrastructureVariables) SetConfig(v InfrastructureConfig)`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


