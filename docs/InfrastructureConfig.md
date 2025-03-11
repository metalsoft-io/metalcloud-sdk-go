# InfrastructureConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | Pointer to **float32** | Revision of the Infrastructure | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**DeployType** | Pointer to **string** | Deploy type of the Infrastructure | [optional] [default to "create"]
**DeployStatus** | Pointer to **string** | Deploy status of the Infrastructure | [optional] [default to "not_started"]
**ServerTypeIdToPreferredServerIds** | Pointer to **map[string]interface{}** | An object having as key the server type id and as value an array of preferred server ids | [optional] 
**InfrastructureDeployId** | Pointer to **float32** | Id of the deployment for the Infrastructure | [optional] 
**DnsSubdomainChangeId** | Pointer to **float32** | DNS Subdomain Change ID | [optional] 
**DatacenterName** | Pointer to **string** | Datacenter name where the Infrastructure is located. | [optional] 
**SiteId** | Pointer to **float32** | The ID of the site where the Infrastructure is located. | [optional] 
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
**EmptyEdit** | Pointer to **float32** | Number of empty edits | [optional] 
**ReservedLanIpRanges** | Pointer to **map[string]interface{}** | Reserved LAN IP ranges in JSON format. | [optional] 
**SubnetPoolLan** | Pointer to **map[string]interface{}** | Subnet pool for LAN in JSON format. | [optional] 
**UpdatedTimestamp** | Pointer to **string** | Timestamp of the latest update for the Infrastructure. | [optional] 
**DnsSubdomainId** | Pointer to **float32** | DNS Subdomain ID. | [optional] 

## Methods

### NewInfrastructureConfig

`func NewInfrastructureConfig() *InfrastructureConfig`

NewInfrastructureConfig instantiates a new InfrastructureConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureConfigWithDefaults

`func NewInfrastructureConfigWithDefaults() *InfrastructureConfig`

NewInfrastructureConfigWithDefaults instantiates a new InfrastructureConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *InfrastructureConfig) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *InfrastructureConfig) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *InfrastructureConfig) SetRevision(v float32)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *InfrastructureConfig) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetLabel

`func (o *InfrastructureConfig) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *InfrastructureConfig) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *InfrastructureConfig) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *InfrastructureConfig) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDeployType

`func (o *InfrastructureConfig) GetDeployType() string`

GetDeployType returns the DeployType field if non-nil, zero value otherwise.

### GetDeployTypeOk

`func (o *InfrastructureConfig) GetDeployTypeOk() (*string, bool)`

GetDeployTypeOk returns a tuple with the DeployType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployType

`func (o *InfrastructureConfig) SetDeployType(v string)`

SetDeployType sets DeployType field to given value.

### HasDeployType

`func (o *InfrastructureConfig) HasDeployType() bool`

HasDeployType returns a boolean if a field has been set.

### GetDeployStatus

`func (o *InfrastructureConfig) GetDeployStatus() string`

GetDeployStatus returns the DeployStatus field if non-nil, zero value otherwise.

### GetDeployStatusOk

`func (o *InfrastructureConfig) GetDeployStatusOk() (*string, bool)`

GetDeployStatusOk returns a tuple with the DeployStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStatus

`func (o *InfrastructureConfig) SetDeployStatus(v string)`

SetDeployStatus sets DeployStatus field to given value.

### HasDeployStatus

`func (o *InfrastructureConfig) HasDeployStatus() bool`

HasDeployStatus returns a boolean if a field has been set.

### GetServerTypeIdToPreferredServerIds

`func (o *InfrastructureConfig) GetServerTypeIdToPreferredServerIds() map[string]interface{}`

GetServerTypeIdToPreferredServerIds returns the ServerTypeIdToPreferredServerIds field if non-nil, zero value otherwise.

### GetServerTypeIdToPreferredServerIdsOk

`func (o *InfrastructureConfig) GetServerTypeIdToPreferredServerIdsOk() (*map[string]interface{}, bool)`

GetServerTypeIdToPreferredServerIdsOk returns a tuple with the ServerTypeIdToPreferredServerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeIdToPreferredServerIds

`func (o *InfrastructureConfig) SetServerTypeIdToPreferredServerIds(v map[string]interface{})`

SetServerTypeIdToPreferredServerIds sets ServerTypeIdToPreferredServerIds field to given value.

### HasServerTypeIdToPreferredServerIds

`func (o *InfrastructureConfig) HasServerTypeIdToPreferredServerIds() bool`

HasServerTypeIdToPreferredServerIds returns a boolean if a field has been set.

### GetInfrastructureDeployId

`func (o *InfrastructureConfig) GetInfrastructureDeployId() float32`

GetInfrastructureDeployId returns the InfrastructureDeployId field if non-nil, zero value otherwise.

### GetInfrastructureDeployIdOk

`func (o *InfrastructureConfig) GetInfrastructureDeployIdOk() (*float32, bool)`

GetInfrastructureDeployIdOk returns a tuple with the InfrastructureDeployId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureDeployId

`func (o *InfrastructureConfig) SetInfrastructureDeployId(v float32)`

SetInfrastructureDeployId sets InfrastructureDeployId field to given value.

### HasInfrastructureDeployId

`func (o *InfrastructureConfig) HasInfrastructureDeployId() bool`

HasInfrastructureDeployId returns a boolean if a field has been set.

### GetDnsSubdomainChangeId

`func (o *InfrastructureConfig) GetDnsSubdomainChangeId() float32`

GetDnsSubdomainChangeId returns the DnsSubdomainChangeId field if non-nil, zero value otherwise.

### GetDnsSubdomainChangeIdOk

`func (o *InfrastructureConfig) GetDnsSubdomainChangeIdOk() (*float32, bool)`

GetDnsSubdomainChangeIdOk returns a tuple with the DnsSubdomainChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainChangeId

`func (o *InfrastructureConfig) SetDnsSubdomainChangeId(v float32)`

SetDnsSubdomainChangeId sets DnsSubdomainChangeId field to given value.

### HasDnsSubdomainChangeId

`func (o *InfrastructureConfig) HasDnsSubdomainChangeId() bool`

HasDnsSubdomainChangeId returns a boolean if a field has been set.

### GetDatacenterName

`func (o *InfrastructureConfig) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *InfrastructureConfig) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *InfrastructureConfig) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *InfrastructureConfig) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.

### GetSiteId

`func (o *InfrastructureConfig) GetSiteId() float32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *InfrastructureConfig) GetSiteIdOk() (*float32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *InfrastructureConfig) SetSiteId(v float32)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *InfrastructureConfig) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetCustomVariables

`func (o *InfrastructureConfig) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *InfrastructureConfig) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *InfrastructureConfig) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *InfrastructureConfig) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.

### GetUserIdOwner

`func (o *InfrastructureConfig) GetUserIdOwner() float32`

GetUserIdOwner returns the UserIdOwner field if non-nil, zero value otherwise.

### GetUserIdOwnerOk

`func (o *InfrastructureConfig) GetUserIdOwnerOk() (*float32, bool)`

GetUserIdOwnerOk returns a tuple with the UserIdOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdOwner

`func (o *InfrastructureConfig) SetUserIdOwner(v float32)`

SetUserIdOwner sets UserIdOwner field to given value.

### HasUserIdOwner

`func (o *InfrastructureConfig) HasUserIdOwner() bool`

HasUserIdOwner returns a boolean if a field has been set.

### GetSubdomain

`func (o *InfrastructureConfig) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *InfrastructureConfig) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *InfrastructureConfig) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *InfrastructureConfig) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetInstancesCountActive

`func (o *InfrastructureConfig) GetInstancesCountActive() float32`

GetInstancesCountActive returns the InstancesCountActive field if non-nil, zero value otherwise.

### GetInstancesCountActiveOk

`func (o *InfrastructureConfig) GetInstancesCountActiveOk() (*float32, bool)`

GetInstancesCountActiveOk returns a tuple with the InstancesCountActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancesCountActive

`func (o *InfrastructureConfig) SetInstancesCountActive(v float32)`

SetInstancesCountActive sets InstancesCountActive field to given value.

### HasInstancesCountActive

`func (o *InfrastructureConfig) HasInstancesCountActive() bool`

HasInstancesCountActive returns a boolean if a field has been set.

### GetDrivesCountActive

`func (o *InfrastructureConfig) GetDrivesCountActive() float32`

GetDrivesCountActive returns the DrivesCountActive field if non-nil, zero value otherwise.

### GetDrivesCountActiveOk

`func (o *InfrastructureConfig) GetDrivesCountActiveOk() (*float32, bool)`

GetDrivesCountActiveOk returns a tuple with the DrivesCountActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivesCountActive

`func (o *InfrastructureConfig) SetDrivesCountActive(v float32)`

SetDrivesCountActive sets DrivesCountActive field to given value.

### HasDrivesCountActive

`func (o *InfrastructureConfig) HasDrivesCountActive() bool`

HasDrivesCountActive returns a boolean if a field has been set.

### GetIpv4SubnetsCountActive

`func (o *InfrastructureConfig) GetIpv4SubnetsCountActive() float32`

GetIpv4SubnetsCountActive returns the Ipv4SubnetsCountActive field if non-nil, zero value otherwise.

### GetIpv4SubnetsCountActiveOk

`func (o *InfrastructureConfig) GetIpv4SubnetsCountActiveOk() (*float32, bool)`

GetIpv4SubnetsCountActiveOk returns a tuple with the Ipv4SubnetsCountActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4SubnetsCountActive

`func (o *InfrastructureConfig) SetIpv4SubnetsCountActive(v float32)`

SetIpv4SubnetsCountActive sets Ipv4SubnetsCountActive field to given value.

### HasIpv4SubnetsCountActive

`func (o *InfrastructureConfig) HasIpv4SubnetsCountActive() bool`

HasIpv4SubnetsCountActive returns a boolean if a field has been set.

### GetIpv6SubnetsCountActive

`func (o *InfrastructureConfig) GetIpv6SubnetsCountActive() float32`

GetIpv6SubnetsCountActive returns the Ipv6SubnetsCountActive field if non-nil, zero value otherwise.

### GetIpv6SubnetsCountActiveOk

`func (o *InfrastructureConfig) GetIpv6SubnetsCountActiveOk() (*float32, bool)`

GetIpv6SubnetsCountActiveOk returns a tuple with the Ipv6SubnetsCountActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6SubnetsCountActive

`func (o *InfrastructureConfig) SetIpv6SubnetsCountActive(v float32)`

SetIpv6SubnetsCountActive sets Ipv6SubnetsCountActive field to given value.

### HasIpv6SubnetsCountActive

`func (o *InfrastructureConfig) HasIpv6SubnetsCountActive() bool`

HasIpv6SubnetsCountActive returns a boolean if a field has been set.

### GetIpv4UnusedIpAddresses

`func (o *InfrastructureConfig) GetIpv4UnusedIpAddresses() float32`

GetIpv4UnusedIpAddresses returns the Ipv4UnusedIpAddresses field if non-nil, zero value otherwise.

### GetIpv4UnusedIpAddressesOk

`func (o *InfrastructureConfig) GetIpv4UnusedIpAddressesOk() (*float32, bool)`

GetIpv4UnusedIpAddressesOk returns a tuple with the Ipv4UnusedIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4UnusedIpAddresses

`func (o *InfrastructureConfig) SetIpv4UnusedIpAddresses(v float32)`

SetIpv4UnusedIpAddresses sets Ipv4UnusedIpAddresses field to given value.

### HasIpv4UnusedIpAddresses

`func (o *InfrastructureConfig) HasIpv4UnusedIpAddresses() bool`

HasIpv4UnusedIpAddresses returns a boolean if a field has been set.

### GetDescription

`func (o *InfrastructureConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InfrastructureConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InfrastructureConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InfrastructureConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSettings

`func (o *InfrastructureConfig) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *InfrastructureConfig) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *InfrastructureConfig) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *InfrastructureConfig) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetIsApiPrivate

`func (o *InfrastructureConfig) GetIsApiPrivate() float32`

GetIsApiPrivate returns the IsApiPrivate field if non-nil, zero value otherwise.

### GetIsApiPrivateOk

`func (o *InfrastructureConfig) GetIsApiPrivateOk() (*float32, bool)`

GetIsApiPrivateOk returns a tuple with the IsApiPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApiPrivate

`func (o *InfrastructureConfig) SetIsApiPrivate(v float32)`

SetIsApiPrivate sets IsApiPrivate field to given value.

### HasIsApiPrivate

`func (o *InfrastructureConfig) HasIsApiPrivate() bool`

HasIsApiPrivate returns a boolean if a field has been set.

### GetExperimentalPriority

`func (o *InfrastructureConfig) GetExperimentalPriority() string`

GetExperimentalPriority returns the ExperimentalPriority field if non-nil, zero value otherwise.

### GetExperimentalPriorityOk

`func (o *InfrastructureConfig) GetExperimentalPriorityOk() (*string, bool)`

GetExperimentalPriorityOk returns a tuple with the ExperimentalPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentalPriority

`func (o *InfrastructureConfig) SetExperimentalPriority(v string)`

SetExperimentalPriority sets ExperimentalPriority field to given value.

### HasExperimentalPriority

`func (o *InfrastructureConfig) HasExperimentalPriority() bool`

HasExperimentalPriority returns a boolean if a field has been set.

### GetIsPublicDesignsMember

`func (o *InfrastructureConfig) GetIsPublicDesignsMember() float32`

GetIsPublicDesignsMember returns the IsPublicDesignsMember field if non-nil, zero value otherwise.

### GetIsPublicDesignsMemberOk

`func (o *InfrastructureConfig) GetIsPublicDesignsMemberOk() (*float32, bool)`

GetIsPublicDesignsMemberOk returns a tuple with the IsPublicDesignsMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublicDesignsMember

`func (o *InfrastructureConfig) SetIsPublicDesignsMember(v float32)`

SetIsPublicDesignsMember sets IsPublicDesignsMember field to given value.

### HasIsPublicDesignsMember

`func (o *InfrastructureConfig) HasIsPublicDesignsMember() bool`

HasIsPublicDesignsMember returns a boolean if a field has been set.

### GetCertificatesJson

`func (o *InfrastructureConfig) GetCertificatesJson() string`

GetCertificatesJson returns the CertificatesJson field if non-nil, zero value otherwise.

### GetCertificatesJsonOk

`func (o *InfrastructureConfig) GetCertificatesJsonOk() (*string, bool)`

GetCertificatesJsonOk returns a tuple with the CertificatesJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatesJson

`func (o *InfrastructureConfig) SetCertificatesJson(v string)`

SetCertificatesJson sets CertificatesJson field to given value.

### HasCertificatesJson

`func (o *InfrastructureConfig) HasCertificatesJson() bool`

HasCertificatesJson returns a boolean if a field has been set.

### GetDeployCookieJarJson

`func (o *InfrastructureConfig) GetDeployCookieJarJson() map[string]interface{}`

GetDeployCookieJarJson returns the DeployCookieJarJson field if non-nil, zero value otherwise.

### GetDeployCookieJarJsonOk

`func (o *InfrastructureConfig) GetDeployCookieJarJsonOk() (*map[string]interface{}, bool)`

GetDeployCookieJarJsonOk returns a tuple with the DeployCookieJarJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployCookieJarJson

`func (o *InfrastructureConfig) SetDeployCookieJarJson(v map[string]interface{})`

SetDeployCookieJarJson sets DeployCookieJarJson field to given value.

### HasDeployCookieJarJson

`func (o *InfrastructureConfig) HasDeployCookieJarJson() bool`

HasDeployCookieJarJson returns a boolean if a field has been set.

### GetDeferredDeployAttemptLastErrorJson

`func (o *InfrastructureConfig) GetDeferredDeployAttemptLastErrorJson() map[string]interface{}`

GetDeferredDeployAttemptLastErrorJson returns the DeferredDeployAttemptLastErrorJson field if non-nil, zero value otherwise.

### GetDeferredDeployAttemptLastErrorJsonOk

`func (o *InfrastructureConfig) GetDeferredDeployAttemptLastErrorJsonOk() (*map[string]interface{}, bool)`

GetDeferredDeployAttemptLastErrorJsonOk returns a tuple with the DeferredDeployAttemptLastErrorJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferredDeployAttemptLastErrorJson

`func (o *InfrastructureConfig) SetDeferredDeployAttemptLastErrorJson(v map[string]interface{})`

SetDeferredDeployAttemptLastErrorJson sets DeferredDeployAttemptLastErrorJson field to given value.

### HasDeferredDeployAttemptLastErrorJson

`func (o *InfrastructureConfig) HasDeferredDeployAttemptLastErrorJson() bool`

HasDeferredDeployAttemptLastErrorJson returns a boolean if a field has been set.

### GetIsAutomanaged

`func (o *InfrastructureConfig) GetIsAutomanaged() float32`

GetIsAutomanaged returns the IsAutomanaged field if non-nil, zero value otherwise.

### GetIsAutomanagedOk

`func (o *InfrastructureConfig) GetIsAutomanagedOk() (*float32, bool)`

GetIsAutomanagedOk returns a tuple with the IsAutomanaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomanaged

`func (o *InfrastructureConfig) SetIsAutomanaged(v float32)`

SetIsAutomanaged sets IsAutomanaged field to given value.

### HasIsAutomanaged

`func (o *InfrastructureConfig) HasIsAutomanaged() bool`

HasIsAutomanaged returns a boolean if a field has been set.

### GetEmptyEdit

`func (o *InfrastructureConfig) GetEmptyEdit() float32`

GetEmptyEdit returns the EmptyEdit field if non-nil, zero value otherwise.

### GetEmptyEditOk

`func (o *InfrastructureConfig) GetEmptyEditOk() (*float32, bool)`

GetEmptyEditOk returns a tuple with the EmptyEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyEdit

`func (o *InfrastructureConfig) SetEmptyEdit(v float32)`

SetEmptyEdit sets EmptyEdit field to given value.

### HasEmptyEdit

`func (o *InfrastructureConfig) HasEmptyEdit() bool`

HasEmptyEdit returns a boolean if a field has been set.

### GetReservedLanIpRanges

`func (o *InfrastructureConfig) GetReservedLanIpRanges() map[string]interface{}`

GetReservedLanIpRanges returns the ReservedLanIpRanges field if non-nil, zero value otherwise.

### GetReservedLanIpRangesOk

`func (o *InfrastructureConfig) GetReservedLanIpRangesOk() (*map[string]interface{}, bool)`

GetReservedLanIpRangesOk returns a tuple with the ReservedLanIpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedLanIpRanges

`func (o *InfrastructureConfig) SetReservedLanIpRanges(v map[string]interface{})`

SetReservedLanIpRanges sets ReservedLanIpRanges field to given value.

### HasReservedLanIpRanges

`func (o *InfrastructureConfig) HasReservedLanIpRanges() bool`

HasReservedLanIpRanges returns a boolean if a field has been set.

### GetSubnetPoolLan

`func (o *InfrastructureConfig) GetSubnetPoolLan() map[string]interface{}`

GetSubnetPoolLan returns the SubnetPoolLan field if non-nil, zero value otherwise.

### GetSubnetPoolLanOk

`func (o *InfrastructureConfig) GetSubnetPoolLanOk() (*map[string]interface{}, bool)`

GetSubnetPoolLanOk returns a tuple with the SubnetPoolLan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetPoolLan

`func (o *InfrastructureConfig) SetSubnetPoolLan(v map[string]interface{})`

SetSubnetPoolLan sets SubnetPoolLan field to given value.

### HasSubnetPoolLan

`func (o *InfrastructureConfig) HasSubnetPoolLan() bool`

HasSubnetPoolLan returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *InfrastructureConfig) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *InfrastructureConfig) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *InfrastructureConfig) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *InfrastructureConfig) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.

### GetDnsSubdomainId

`func (o *InfrastructureConfig) GetDnsSubdomainId() float32`

GetDnsSubdomainId returns the DnsSubdomainId field if non-nil, zero value otherwise.

### GetDnsSubdomainIdOk

`func (o *InfrastructureConfig) GetDnsSubdomainIdOk() (*float32, bool)`

GetDnsSubdomainIdOk returns a tuple with the DnsSubdomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainId

`func (o *InfrastructureConfig) SetDnsSubdomainId(v float32)`

SetDnsSubdomainId sets DnsSubdomainId field to given value.

### HasDnsSubdomainId

`func (o *InfrastructureConfig) HasDnsSubdomainId() bool`

HasDnsSubdomainId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


