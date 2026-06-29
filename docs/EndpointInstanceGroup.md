# EndpointInstanceGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The Product Instance ID. | 
**Revision** | **int64** | Revision number | 
**Label** | **string** | The Product Instance label. Will be automatically generated if not provided. | 
**CreatedTimestamp** | **string** | Timestamp of the Product Instance creation. | 
**UpdatedTimestamp** | **string** | Timestamp of the latest update of the Product Instance. | 
**Subdomain** | Pointer to **string** | Subdomain of the Product Instance. | [optional] 
**SubdomainPermanent** | Pointer to **string** | Subdomain permanent of the Product Instance. | [optional] 
**DnsSubdomainId** | Pointer to **int64** | Id of the DNS subdomain for the Product Instance | [optional] 
**DnsSubdomainPermanentId** | Pointer to **int64** | Id of the permanent DNS subdomain for the Product Instance | [optional] 
**EndpointGroupName** | Pointer to **string** |  | [optional] 
**InfrastructureId** | **int64** |  | 
**ExtensionInstanceId** | Pointer to **int64** |  | [optional] 
**Hostname** | Pointer to **string** | Custom hostname for the DNS Load Balancing record. If set, this will be used as the DNS Load Balancing record name instead of the default \&quot;endpoint-instance-group\&quot;. The hostname must be a valid DNS subdomain and can only contain alphanumeric characters, hyphens, and underscores. This will only take effect if the property \&quot;dnsLoadBalancingRecord\&quot; is true. It will be automatically suffixed with the endpoint instance group ID (e.g., \&quot;-34\&quot;) to ensure the uniqueness of the resulting DNS name. | [optional] 
**ServiceStatus** | **string** | Current status of the Endpoint Instance Group. | 
**ResourcePoolId** | Pointer to **int64** | The resource pool assigned to this instance array | [optional] 
**NetworkEndpointGroupId** | Pointer to **int64** |  | [optional] 
**Meta** | [**GenericMeta**](GenericMeta.md) |  | 
**Config** | Pointer to [**EndpointInstanceGroupConfiguration**](EndpointInstanceGroupConfiguration.md) |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewEndpointInstanceGroup

`func NewEndpointInstanceGroup(id int64, revision int64, label string, createdTimestamp string, updatedTimestamp string, infrastructureId int64, serviceStatus string, meta GenericMeta, ) *EndpointInstanceGroup`

NewEndpointInstanceGroup instantiates a new EndpointInstanceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointInstanceGroupWithDefaults

`func NewEndpointInstanceGroupWithDefaults() *EndpointInstanceGroup`

NewEndpointInstanceGroupWithDefaults instantiates a new EndpointInstanceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EndpointInstanceGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EndpointInstanceGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EndpointInstanceGroup) SetId(v int64)`

SetId sets Id field to given value.


### GetRevision

`func (o *EndpointInstanceGroup) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *EndpointInstanceGroup) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *EndpointInstanceGroup) SetRevision(v int64)`

SetRevision sets Revision field to given value.


### GetLabel

`func (o *EndpointInstanceGroup) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *EndpointInstanceGroup) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *EndpointInstanceGroup) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetCreatedTimestamp

`func (o *EndpointInstanceGroup) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *EndpointInstanceGroup) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *EndpointInstanceGroup) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *EndpointInstanceGroup) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *EndpointInstanceGroup) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *EndpointInstanceGroup) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetSubdomain

`func (o *EndpointInstanceGroup) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *EndpointInstanceGroup) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *EndpointInstanceGroup) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *EndpointInstanceGroup) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetSubdomainPermanent

`func (o *EndpointInstanceGroup) GetSubdomainPermanent() string`

GetSubdomainPermanent returns the SubdomainPermanent field if non-nil, zero value otherwise.

### GetSubdomainPermanentOk

`func (o *EndpointInstanceGroup) GetSubdomainPermanentOk() (*string, bool)`

GetSubdomainPermanentOk returns a tuple with the SubdomainPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomainPermanent

`func (o *EndpointInstanceGroup) SetSubdomainPermanent(v string)`

SetSubdomainPermanent sets SubdomainPermanent field to given value.

### HasSubdomainPermanent

`func (o *EndpointInstanceGroup) HasSubdomainPermanent() bool`

HasSubdomainPermanent returns a boolean if a field has been set.

### GetDnsSubdomainId

`func (o *EndpointInstanceGroup) GetDnsSubdomainId() int64`

GetDnsSubdomainId returns the DnsSubdomainId field if non-nil, zero value otherwise.

### GetDnsSubdomainIdOk

`func (o *EndpointInstanceGroup) GetDnsSubdomainIdOk() (*int64, bool)`

GetDnsSubdomainIdOk returns a tuple with the DnsSubdomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainId

`func (o *EndpointInstanceGroup) SetDnsSubdomainId(v int64)`

SetDnsSubdomainId sets DnsSubdomainId field to given value.

### HasDnsSubdomainId

`func (o *EndpointInstanceGroup) HasDnsSubdomainId() bool`

HasDnsSubdomainId returns a boolean if a field has been set.

### GetDnsSubdomainPermanentId

`func (o *EndpointInstanceGroup) GetDnsSubdomainPermanentId() int64`

GetDnsSubdomainPermanentId returns the DnsSubdomainPermanentId field if non-nil, zero value otherwise.

### GetDnsSubdomainPermanentIdOk

`func (o *EndpointInstanceGroup) GetDnsSubdomainPermanentIdOk() (*int64, bool)`

GetDnsSubdomainPermanentIdOk returns a tuple with the DnsSubdomainPermanentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainPermanentId

`func (o *EndpointInstanceGroup) SetDnsSubdomainPermanentId(v int64)`

SetDnsSubdomainPermanentId sets DnsSubdomainPermanentId field to given value.

### HasDnsSubdomainPermanentId

`func (o *EndpointInstanceGroup) HasDnsSubdomainPermanentId() bool`

HasDnsSubdomainPermanentId returns a boolean if a field has been set.

### GetEndpointGroupName

`func (o *EndpointInstanceGroup) GetEndpointGroupName() string`

GetEndpointGroupName returns the EndpointGroupName field if non-nil, zero value otherwise.

### GetEndpointGroupNameOk

`func (o *EndpointInstanceGroup) GetEndpointGroupNameOk() (*string, bool)`

GetEndpointGroupNameOk returns a tuple with the EndpointGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointGroupName

`func (o *EndpointInstanceGroup) SetEndpointGroupName(v string)`

SetEndpointGroupName sets EndpointGroupName field to given value.

### HasEndpointGroupName

`func (o *EndpointInstanceGroup) HasEndpointGroupName() bool`

HasEndpointGroupName returns a boolean if a field has been set.

### GetInfrastructureId

`func (o *EndpointInstanceGroup) GetInfrastructureId() int64`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *EndpointInstanceGroup) GetInfrastructureIdOk() (*int64, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *EndpointInstanceGroup) SetInfrastructureId(v int64)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetExtensionInstanceId

`func (o *EndpointInstanceGroup) GetExtensionInstanceId() int64`

GetExtensionInstanceId returns the ExtensionInstanceId field if non-nil, zero value otherwise.

### GetExtensionInstanceIdOk

`func (o *EndpointInstanceGroup) GetExtensionInstanceIdOk() (*int64, bool)`

GetExtensionInstanceIdOk returns a tuple with the ExtensionInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionInstanceId

`func (o *EndpointInstanceGroup) SetExtensionInstanceId(v int64)`

SetExtensionInstanceId sets ExtensionInstanceId field to given value.

### HasExtensionInstanceId

`func (o *EndpointInstanceGroup) HasExtensionInstanceId() bool`

HasExtensionInstanceId returns a boolean if a field has been set.

### GetHostname

`func (o *EndpointInstanceGroup) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *EndpointInstanceGroup) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *EndpointInstanceGroup) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *EndpointInstanceGroup) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetServiceStatus

`func (o *EndpointInstanceGroup) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *EndpointInstanceGroup) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *EndpointInstanceGroup) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetResourcePoolId

`func (o *EndpointInstanceGroup) GetResourcePoolId() int64`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *EndpointInstanceGroup) GetResourcePoolIdOk() (*int64, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *EndpointInstanceGroup) SetResourcePoolId(v int64)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *EndpointInstanceGroup) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetNetworkEndpointGroupId

`func (o *EndpointInstanceGroup) GetNetworkEndpointGroupId() int64`

GetNetworkEndpointGroupId returns the NetworkEndpointGroupId field if non-nil, zero value otherwise.

### GetNetworkEndpointGroupIdOk

`func (o *EndpointInstanceGroup) GetNetworkEndpointGroupIdOk() (*int64, bool)`

GetNetworkEndpointGroupIdOk returns a tuple with the NetworkEndpointGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEndpointGroupId

`func (o *EndpointInstanceGroup) SetNetworkEndpointGroupId(v int64)`

SetNetworkEndpointGroupId sets NetworkEndpointGroupId field to given value.

### HasNetworkEndpointGroupId

`func (o *EndpointInstanceGroup) HasNetworkEndpointGroupId() bool`

HasNetworkEndpointGroupId returns a boolean if a field has been set.

### GetMeta

`func (o *EndpointInstanceGroup) GetMeta() GenericMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *EndpointInstanceGroup) GetMetaOk() (*GenericMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *EndpointInstanceGroup) SetMeta(v GenericMeta)`

SetMeta sets Meta field to given value.


### GetConfig

`func (o *EndpointInstanceGroup) GetConfig() EndpointInstanceGroupConfiguration`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *EndpointInstanceGroup) GetConfigOk() (*EndpointInstanceGroupConfiguration, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *EndpointInstanceGroup) SetConfig(v EndpointInstanceGroupConfiguration)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *EndpointInstanceGroup) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetLinks

`func (o *EndpointInstanceGroup) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EndpointInstanceGroup) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EndpointInstanceGroup) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EndpointInstanceGroup) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


