# ContainerInstanceGroupInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Name of the VM Instance Group Interface. | 
**Index** | **float32** | Interface index | 
**NetworkId** | Pointer to **int64** | Network ID | [optional] 
**Subdomain** | Pointer to **string** | Subdomain of the VM Instance Group Interface. | [optional] 
**UpdatedTimestamp** | **string** | Timestamp of the VM Instance Group Interface update. | 
**Id** | **int64** | Interface ID | 
**Revision** | **int64** | Revision of the VM Instance Group Interface Configuration | 
**ServiceStatus** | **string** | Service status of the VM Instance Group Interface. | 
**GroupId** | **int64** | VM Instance Group ID | 
**InfrastructureId** | **int64** | Infrastructure ID | 
**SubdomainPermanent** | Pointer to **string** | Subdomain permanent of the VM Instance Group Interface. | [optional] 
**DnsSubdomainId** | Pointer to **int64** | Id of the DNS subdomain for the VM Instance Group Interface. | [optional] 
**DnsSubdomainPermanentId** | Pointer to **int64** | Id of the permanent DNS subdomain for the VM Instance Group Interface. | [optional] 
**Config** | [**VMInstanceGroupInterfaceConfiguration**](VMInstanceGroupInterfaceConfiguration.md) | The current changes to be deployed for the VM Instance Group Interface. | 
**Meta** | **map[string]interface{}** | Meta information of the VM Instance Group Interface. | 
**CreatedTimestamp** | **string** | Timestamp of the VM Instance Group Interface creation. | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewContainerInstanceGroupInterface

`func NewContainerInstanceGroupInterface(label string, index float32, updatedTimestamp string, id int64, revision int64, serviceStatus string, groupId int64, infrastructureId int64, config VMInstanceGroupInterfaceConfiguration, meta map[string]interface{}, createdTimestamp string, ) *ContainerInstanceGroupInterface`

NewContainerInstanceGroupInterface instantiates a new ContainerInstanceGroupInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerInstanceGroupInterfaceWithDefaults

`func NewContainerInstanceGroupInterfaceWithDefaults() *ContainerInstanceGroupInterface`

NewContainerInstanceGroupInterfaceWithDefaults instantiates a new ContainerInstanceGroupInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ContainerInstanceGroupInterface) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ContainerInstanceGroupInterface) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ContainerInstanceGroupInterface) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetIndex

`func (o *ContainerInstanceGroupInterface) GetIndex() float32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ContainerInstanceGroupInterface) GetIndexOk() (*float32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ContainerInstanceGroupInterface) SetIndex(v float32)`

SetIndex sets Index field to given value.


### GetNetworkId

`func (o *ContainerInstanceGroupInterface) GetNetworkId() int64`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *ContainerInstanceGroupInterface) GetNetworkIdOk() (*int64, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *ContainerInstanceGroupInterface) SetNetworkId(v int64)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *ContainerInstanceGroupInterface) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetSubdomain

`func (o *ContainerInstanceGroupInterface) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *ContainerInstanceGroupInterface) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *ContainerInstanceGroupInterface) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *ContainerInstanceGroupInterface) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *ContainerInstanceGroupInterface) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *ContainerInstanceGroupInterface) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *ContainerInstanceGroupInterface) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetId

`func (o *ContainerInstanceGroupInterface) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContainerInstanceGroupInterface) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContainerInstanceGroupInterface) SetId(v int64)`

SetId sets Id field to given value.


### GetRevision

`func (o *ContainerInstanceGroupInterface) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ContainerInstanceGroupInterface) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ContainerInstanceGroupInterface) SetRevision(v int64)`

SetRevision sets Revision field to given value.


### GetServiceStatus

`func (o *ContainerInstanceGroupInterface) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *ContainerInstanceGroupInterface) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *ContainerInstanceGroupInterface) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetGroupId

`func (o *ContainerInstanceGroupInterface) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ContainerInstanceGroupInterface) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ContainerInstanceGroupInterface) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.


### GetInfrastructureId

`func (o *ContainerInstanceGroupInterface) GetInfrastructureId() int64`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *ContainerInstanceGroupInterface) GetInfrastructureIdOk() (*int64, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *ContainerInstanceGroupInterface) SetInfrastructureId(v int64)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetSubdomainPermanent

`func (o *ContainerInstanceGroupInterface) GetSubdomainPermanent() string`

GetSubdomainPermanent returns the SubdomainPermanent field if non-nil, zero value otherwise.

### GetSubdomainPermanentOk

`func (o *ContainerInstanceGroupInterface) GetSubdomainPermanentOk() (*string, bool)`

GetSubdomainPermanentOk returns a tuple with the SubdomainPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomainPermanent

`func (o *ContainerInstanceGroupInterface) SetSubdomainPermanent(v string)`

SetSubdomainPermanent sets SubdomainPermanent field to given value.

### HasSubdomainPermanent

`func (o *ContainerInstanceGroupInterface) HasSubdomainPermanent() bool`

HasSubdomainPermanent returns a boolean if a field has been set.

### GetDnsSubdomainId

`func (o *ContainerInstanceGroupInterface) GetDnsSubdomainId() int64`

GetDnsSubdomainId returns the DnsSubdomainId field if non-nil, zero value otherwise.

### GetDnsSubdomainIdOk

`func (o *ContainerInstanceGroupInterface) GetDnsSubdomainIdOk() (*int64, bool)`

GetDnsSubdomainIdOk returns a tuple with the DnsSubdomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainId

`func (o *ContainerInstanceGroupInterface) SetDnsSubdomainId(v int64)`

SetDnsSubdomainId sets DnsSubdomainId field to given value.

### HasDnsSubdomainId

`func (o *ContainerInstanceGroupInterface) HasDnsSubdomainId() bool`

HasDnsSubdomainId returns a boolean if a field has been set.

### GetDnsSubdomainPermanentId

`func (o *ContainerInstanceGroupInterface) GetDnsSubdomainPermanentId() int64`

GetDnsSubdomainPermanentId returns the DnsSubdomainPermanentId field if non-nil, zero value otherwise.

### GetDnsSubdomainPermanentIdOk

`func (o *ContainerInstanceGroupInterface) GetDnsSubdomainPermanentIdOk() (*int64, bool)`

GetDnsSubdomainPermanentIdOk returns a tuple with the DnsSubdomainPermanentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainPermanentId

`func (o *ContainerInstanceGroupInterface) SetDnsSubdomainPermanentId(v int64)`

SetDnsSubdomainPermanentId sets DnsSubdomainPermanentId field to given value.

### HasDnsSubdomainPermanentId

`func (o *ContainerInstanceGroupInterface) HasDnsSubdomainPermanentId() bool`

HasDnsSubdomainPermanentId returns a boolean if a field has been set.

### GetConfig

`func (o *ContainerInstanceGroupInterface) GetConfig() VMInstanceGroupInterfaceConfiguration`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ContainerInstanceGroupInterface) GetConfigOk() (*VMInstanceGroupInterfaceConfiguration, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ContainerInstanceGroupInterface) SetConfig(v VMInstanceGroupInterfaceConfiguration)`

SetConfig sets Config field to given value.


### GetMeta

`func (o *ContainerInstanceGroupInterface) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ContainerInstanceGroupInterface) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ContainerInstanceGroupInterface) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.


### GetCreatedTimestamp

`func (o *ContainerInstanceGroupInterface) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *ContainerInstanceGroupInterface) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *ContainerInstanceGroupInterface) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetLinks

`func (o *ContainerInstanceGroupInterface) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ContainerInstanceGroupInterface) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ContainerInstanceGroupInterface) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ContainerInstanceGroupInterface) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


