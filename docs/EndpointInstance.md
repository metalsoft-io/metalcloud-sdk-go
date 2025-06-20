# EndpointInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The Product Instance ID. | 
**Revision** | **int32** | Revision number | 
**Label** | **string** | The Product Instance label. Will be automatically generated if not provided. | 
**CreatedTimestamp** | **string** | Timestamp of the Product Instance creation. | 
**UpdatedTimestamp** | **string** | Timestamp of the latest update of the Product Instance. | 
**Subdomain** | Pointer to **string** | Subdomain of the Product Instance. | [optional] 
**SubdomainPermanent** | Pointer to **string** | Subdomain permanent of the Product Instance. | [optional] 
**DnsSubdomainId** | Pointer to **int32** | Id of the DNS subdomain for the Product Instance | [optional] 
**DnsSubdomainPermanentId** | Pointer to **int32** | Id of the permanent DNS subdomain for the Product Instance | [optional] 
**InfrastructureId** | **int32** |  | 
**GroupId** | **int32** |  | 
**EndpointId** | Pointer to **int32** | The ID of the endpoint assigned to the instance. | [optional] 
**Hostname** | Pointer to **string** | The subdomain of the server instance. | [optional] 
**ServiceStatus** | **string** | Current status of the endpoint instance. | 
**Meta** | [**GenericMeta**](GenericMeta.md) |  | 
**Config** | Pointer to [**EndpointInstanceConfiguration**](EndpointInstanceConfiguration.md) |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewEndpointInstance

`func NewEndpointInstance(id int32, revision int32, label string, createdTimestamp string, updatedTimestamp string, infrastructureId int32, groupId int32, serviceStatus string, meta GenericMeta, ) *EndpointInstance`

NewEndpointInstance instantiates a new EndpointInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointInstanceWithDefaults

`func NewEndpointInstanceWithDefaults() *EndpointInstance`

NewEndpointInstanceWithDefaults instantiates a new EndpointInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EndpointInstance) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EndpointInstance) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EndpointInstance) SetId(v int32)`

SetId sets Id field to given value.


### GetRevision

`func (o *EndpointInstance) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *EndpointInstance) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *EndpointInstance) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetLabel

`func (o *EndpointInstance) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *EndpointInstance) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *EndpointInstance) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetCreatedTimestamp

`func (o *EndpointInstance) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *EndpointInstance) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *EndpointInstance) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *EndpointInstance) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *EndpointInstance) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *EndpointInstance) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetSubdomain

`func (o *EndpointInstance) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *EndpointInstance) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *EndpointInstance) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *EndpointInstance) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetSubdomainPermanent

`func (o *EndpointInstance) GetSubdomainPermanent() string`

GetSubdomainPermanent returns the SubdomainPermanent field if non-nil, zero value otherwise.

### GetSubdomainPermanentOk

`func (o *EndpointInstance) GetSubdomainPermanentOk() (*string, bool)`

GetSubdomainPermanentOk returns a tuple with the SubdomainPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomainPermanent

`func (o *EndpointInstance) SetSubdomainPermanent(v string)`

SetSubdomainPermanent sets SubdomainPermanent field to given value.

### HasSubdomainPermanent

`func (o *EndpointInstance) HasSubdomainPermanent() bool`

HasSubdomainPermanent returns a boolean if a field has been set.

### GetDnsSubdomainId

`func (o *EndpointInstance) GetDnsSubdomainId() int32`

GetDnsSubdomainId returns the DnsSubdomainId field if non-nil, zero value otherwise.

### GetDnsSubdomainIdOk

`func (o *EndpointInstance) GetDnsSubdomainIdOk() (*int32, bool)`

GetDnsSubdomainIdOk returns a tuple with the DnsSubdomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainId

`func (o *EndpointInstance) SetDnsSubdomainId(v int32)`

SetDnsSubdomainId sets DnsSubdomainId field to given value.

### HasDnsSubdomainId

`func (o *EndpointInstance) HasDnsSubdomainId() bool`

HasDnsSubdomainId returns a boolean if a field has been set.

### GetDnsSubdomainPermanentId

`func (o *EndpointInstance) GetDnsSubdomainPermanentId() int32`

GetDnsSubdomainPermanentId returns the DnsSubdomainPermanentId field if non-nil, zero value otherwise.

### GetDnsSubdomainPermanentIdOk

`func (o *EndpointInstance) GetDnsSubdomainPermanentIdOk() (*int32, bool)`

GetDnsSubdomainPermanentIdOk returns a tuple with the DnsSubdomainPermanentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainPermanentId

`func (o *EndpointInstance) SetDnsSubdomainPermanentId(v int32)`

SetDnsSubdomainPermanentId sets DnsSubdomainPermanentId field to given value.

### HasDnsSubdomainPermanentId

`func (o *EndpointInstance) HasDnsSubdomainPermanentId() bool`

HasDnsSubdomainPermanentId returns a boolean if a field has been set.

### GetInfrastructureId

`func (o *EndpointInstance) GetInfrastructureId() int32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *EndpointInstance) GetInfrastructureIdOk() (*int32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *EndpointInstance) SetInfrastructureId(v int32)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetGroupId

`func (o *EndpointInstance) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *EndpointInstance) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *EndpointInstance) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.


### GetEndpointId

`func (o *EndpointInstance) GetEndpointId() int32`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *EndpointInstance) GetEndpointIdOk() (*int32, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *EndpointInstance) SetEndpointId(v int32)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *EndpointInstance) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetHostname

`func (o *EndpointInstance) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *EndpointInstance) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *EndpointInstance) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *EndpointInstance) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetServiceStatus

`func (o *EndpointInstance) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *EndpointInstance) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *EndpointInstance) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetMeta

`func (o *EndpointInstance) GetMeta() GenericMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *EndpointInstance) GetMetaOk() (*GenericMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *EndpointInstance) SetMeta(v GenericMeta)`

SetMeta sets Meta field to given value.


### GetConfig

`func (o *EndpointInstance) GetConfig() EndpointInstanceConfiguration`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *EndpointInstance) GetConfigOk() (*EndpointInstanceConfiguration, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *EndpointInstance) SetConfig(v EndpointInstanceConfiguration)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *EndpointInstance) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetLinks

`func (o *EndpointInstance) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EndpointInstance) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EndpointInstance) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EndpointInstance) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


