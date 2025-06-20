# ServerInstanceGroupInterface

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
**Index** | **int32** | The index of the interface (0-based) on this server. | 
**NetworkId** | Pointer to **int32** | The ID of the network to which this interface is to be attached to. | [optional] 
**ServiceStatus** | **string** | Current status of the server instance group interface. | 
**Config** | Pointer to [**ServerInstanceGroupInterfaceConfiguration**](ServerInstanceGroupInterfaceConfiguration.md) |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewServerInstanceGroupInterface

`func NewServerInstanceGroupInterface(id int32, revision int32, label string, createdTimestamp string, updatedTimestamp string, infrastructureId int32, groupId int32, index int32, serviceStatus string, ) *ServerInstanceGroupInterface`

NewServerInstanceGroupInterface instantiates a new ServerInstanceGroupInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceGroupInterfaceWithDefaults

`func NewServerInstanceGroupInterfaceWithDefaults() *ServerInstanceGroupInterface`

NewServerInstanceGroupInterfaceWithDefaults instantiates a new ServerInstanceGroupInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerInstanceGroupInterface) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerInstanceGroupInterface) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerInstanceGroupInterface) SetId(v int32)`

SetId sets Id field to given value.


### GetRevision

`func (o *ServerInstanceGroupInterface) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ServerInstanceGroupInterface) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ServerInstanceGroupInterface) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetLabel

`func (o *ServerInstanceGroupInterface) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServerInstanceGroupInterface) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServerInstanceGroupInterface) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetCreatedTimestamp

`func (o *ServerInstanceGroupInterface) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *ServerInstanceGroupInterface) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *ServerInstanceGroupInterface) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *ServerInstanceGroupInterface) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *ServerInstanceGroupInterface) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *ServerInstanceGroupInterface) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetSubdomain

`func (o *ServerInstanceGroupInterface) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *ServerInstanceGroupInterface) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *ServerInstanceGroupInterface) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *ServerInstanceGroupInterface) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetSubdomainPermanent

`func (o *ServerInstanceGroupInterface) GetSubdomainPermanent() string`

GetSubdomainPermanent returns the SubdomainPermanent field if non-nil, zero value otherwise.

### GetSubdomainPermanentOk

`func (o *ServerInstanceGroupInterface) GetSubdomainPermanentOk() (*string, bool)`

GetSubdomainPermanentOk returns a tuple with the SubdomainPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomainPermanent

`func (o *ServerInstanceGroupInterface) SetSubdomainPermanent(v string)`

SetSubdomainPermanent sets SubdomainPermanent field to given value.

### HasSubdomainPermanent

`func (o *ServerInstanceGroupInterface) HasSubdomainPermanent() bool`

HasSubdomainPermanent returns a boolean if a field has been set.

### GetDnsSubdomainId

`func (o *ServerInstanceGroupInterface) GetDnsSubdomainId() int32`

GetDnsSubdomainId returns the DnsSubdomainId field if non-nil, zero value otherwise.

### GetDnsSubdomainIdOk

`func (o *ServerInstanceGroupInterface) GetDnsSubdomainIdOk() (*int32, bool)`

GetDnsSubdomainIdOk returns a tuple with the DnsSubdomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainId

`func (o *ServerInstanceGroupInterface) SetDnsSubdomainId(v int32)`

SetDnsSubdomainId sets DnsSubdomainId field to given value.

### HasDnsSubdomainId

`func (o *ServerInstanceGroupInterface) HasDnsSubdomainId() bool`

HasDnsSubdomainId returns a boolean if a field has been set.

### GetDnsSubdomainPermanentId

`func (o *ServerInstanceGroupInterface) GetDnsSubdomainPermanentId() int32`

GetDnsSubdomainPermanentId returns the DnsSubdomainPermanentId field if non-nil, zero value otherwise.

### GetDnsSubdomainPermanentIdOk

`func (o *ServerInstanceGroupInterface) GetDnsSubdomainPermanentIdOk() (*int32, bool)`

GetDnsSubdomainPermanentIdOk returns a tuple with the DnsSubdomainPermanentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainPermanentId

`func (o *ServerInstanceGroupInterface) SetDnsSubdomainPermanentId(v int32)`

SetDnsSubdomainPermanentId sets DnsSubdomainPermanentId field to given value.

### HasDnsSubdomainPermanentId

`func (o *ServerInstanceGroupInterface) HasDnsSubdomainPermanentId() bool`

HasDnsSubdomainPermanentId returns a boolean if a field has been set.

### GetInfrastructureId

`func (o *ServerInstanceGroupInterface) GetInfrastructureId() int32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *ServerInstanceGroupInterface) GetInfrastructureIdOk() (*int32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *ServerInstanceGroupInterface) SetInfrastructureId(v int32)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetGroupId

`func (o *ServerInstanceGroupInterface) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ServerInstanceGroupInterface) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ServerInstanceGroupInterface) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.


### GetIndex

`func (o *ServerInstanceGroupInterface) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ServerInstanceGroupInterface) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ServerInstanceGroupInterface) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetNetworkId

`func (o *ServerInstanceGroupInterface) GetNetworkId() int32`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *ServerInstanceGroupInterface) GetNetworkIdOk() (*int32, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *ServerInstanceGroupInterface) SetNetworkId(v int32)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *ServerInstanceGroupInterface) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetServiceStatus

`func (o *ServerInstanceGroupInterface) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *ServerInstanceGroupInterface) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *ServerInstanceGroupInterface) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetConfig

`func (o *ServerInstanceGroupInterface) GetConfig() ServerInstanceGroupInterfaceConfiguration`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ServerInstanceGroupInterface) GetConfigOk() (*ServerInstanceGroupInterfaceConfiguration, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ServerInstanceGroupInterface) SetConfig(v ServerInstanceGroupInterfaceConfiguration)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ServerInstanceGroupInterface) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetLinks

`func (o *ServerInstanceGroupInterface) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServerInstanceGroupInterface) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServerInstanceGroupInterface) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServerInstanceGroupInterface) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


