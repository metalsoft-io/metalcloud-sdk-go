# ContainerInstanceGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Name of the VM Instance Group. | 
**Subdomain** | Pointer to **string** | Subdomain of the VM Instance Group. | [optional] 
**InstanceCount** | Pointer to **float32** |  | [optional] [default to 1]
**CustomVariables** | Pointer to **map[string]interface{}** | Custom variables for the VM Instance Group. | [optional] 
**UpdatedTimestamp** | **string** | Timestamp of the VM Instance Group last update. | 
**Id** | **int64** | Id of the VM Instance Group. | 
**Revision** | **int64** | Revision of the VM Instance Group State | 
**InfrastructureId** | **int64** | Id of the Infrastructure. | 
**NetworkEndpointGroupId** | Pointer to **int64** | The network endpoint group identifier | [optional] 
**Infrastructure** | [**ParentInfrastructure**](ParentInfrastructure.md) | Infrastructure information | 
**ServiceStatus** | **string** | Status of the VM Instance Group. | 
**InstanceGroupType** | **string** | Type of the Instance Group. | 
**DiskSizeGB** | **float32** | Disk size in GB for each VM Instance in the VM Instance Group. | 
**SubdomainPermanent** | Pointer to **string** | Subdomain permanent of the VM Instance Group. | [optional] 
**DnsSubdomainId** | Pointer to **int64** | Id of the DNS subdomain for the VM Instance Group. | [optional] 
**DnsSubdomainPermanentId** | Pointer to **int64** | Id of the permanent DNS subdomain for the VM Instance Group. | [optional] 
**VmPoolId** | Pointer to **int64** | Id of the VM Pool used by the VM Instance Group. | [optional] 
**CreatedTimestamp** | **string** | Timestamp of the VM Instance Group creation. | 
**Config** | [**VMInstanceGroupConfiguration**](VMInstanceGroupConfiguration.md) | The current changes to be deployed for the VM Instance Group. | 
**Meta** | [**VMInstanceGroupMeta**](VMInstanceGroupMeta.md) | Meta information of the VM Instance Group. | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewContainerInstanceGroup

`func NewContainerInstanceGroup(label string, updatedTimestamp string, id int64, revision int64, infrastructureId int64, infrastructure ParentInfrastructure, serviceStatus string, instanceGroupType string, diskSizeGB float32, createdTimestamp string, config VMInstanceGroupConfiguration, meta VMInstanceGroupMeta, ) *ContainerInstanceGroup`

NewContainerInstanceGroup instantiates a new ContainerInstanceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerInstanceGroupWithDefaults

`func NewContainerInstanceGroupWithDefaults() *ContainerInstanceGroup`

NewContainerInstanceGroupWithDefaults instantiates a new ContainerInstanceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ContainerInstanceGroup) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ContainerInstanceGroup) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ContainerInstanceGroup) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetSubdomain

`func (o *ContainerInstanceGroup) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *ContainerInstanceGroup) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *ContainerInstanceGroup) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *ContainerInstanceGroup) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetInstanceCount

`func (o *ContainerInstanceGroup) GetInstanceCount() float32`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *ContainerInstanceGroup) GetInstanceCountOk() (*float32, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *ContainerInstanceGroup) SetInstanceCount(v float32)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *ContainerInstanceGroup) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

### GetCustomVariables

`func (o *ContainerInstanceGroup) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *ContainerInstanceGroup) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *ContainerInstanceGroup) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *ContainerInstanceGroup) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *ContainerInstanceGroup) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *ContainerInstanceGroup) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *ContainerInstanceGroup) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetId

`func (o *ContainerInstanceGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContainerInstanceGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContainerInstanceGroup) SetId(v int64)`

SetId sets Id field to given value.


### GetRevision

`func (o *ContainerInstanceGroup) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ContainerInstanceGroup) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ContainerInstanceGroup) SetRevision(v int64)`

SetRevision sets Revision field to given value.


### GetInfrastructureId

`func (o *ContainerInstanceGroup) GetInfrastructureId() int64`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *ContainerInstanceGroup) GetInfrastructureIdOk() (*int64, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *ContainerInstanceGroup) SetInfrastructureId(v int64)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetNetworkEndpointGroupId

`func (o *ContainerInstanceGroup) GetNetworkEndpointGroupId() int64`

GetNetworkEndpointGroupId returns the NetworkEndpointGroupId field if non-nil, zero value otherwise.

### GetNetworkEndpointGroupIdOk

`func (o *ContainerInstanceGroup) GetNetworkEndpointGroupIdOk() (*int64, bool)`

GetNetworkEndpointGroupIdOk returns a tuple with the NetworkEndpointGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEndpointGroupId

`func (o *ContainerInstanceGroup) SetNetworkEndpointGroupId(v int64)`

SetNetworkEndpointGroupId sets NetworkEndpointGroupId field to given value.

### HasNetworkEndpointGroupId

`func (o *ContainerInstanceGroup) HasNetworkEndpointGroupId() bool`

HasNetworkEndpointGroupId returns a boolean if a field has been set.

### GetInfrastructure

`func (o *ContainerInstanceGroup) GetInfrastructure() ParentInfrastructure`

GetInfrastructure returns the Infrastructure field if non-nil, zero value otherwise.

### GetInfrastructureOk

`func (o *ContainerInstanceGroup) GetInfrastructureOk() (*ParentInfrastructure, bool)`

GetInfrastructureOk returns a tuple with the Infrastructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructure

`func (o *ContainerInstanceGroup) SetInfrastructure(v ParentInfrastructure)`

SetInfrastructure sets Infrastructure field to given value.


### GetServiceStatus

`func (o *ContainerInstanceGroup) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *ContainerInstanceGroup) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *ContainerInstanceGroup) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetInstanceGroupType

`func (o *ContainerInstanceGroup) GetInstanceGroupType() string`

GetInstanceGroupType returns the InstanceGroupType field if non-nil, zero value otherwise.

### GetInstanceGroupTypeOk

`func (o *ContainerInstanceGroup) GetInstanceGroupTypeOk() (*string, bool)`

GetInstanceGroupTypeOk returns a tuple with the InstanceGroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroupType

`func (o *ContainerInstanceGroup) SetInstanceGroupType(v string)`

SetInstanceGroupType sets InstanceGroupType field to given value.


### GetDiskSizeGB

`func (o *ContainerInstanceGroup) GetDiskSizeGB() float32`

GetDiskSizeGB returns the DiskSizeGB field if non-nil, zero value otherwise.

### GetDiskSizeGBOk

`func (o *ContainerInstanceGroup) GetDiskSizeGBOk() (*float32, bool)`

GetDiskSizeGBOk returns a tuple with the DiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeGB

`func (o *ContainerInstanceGroup) SetDiskSizeGB(v float32)`

SetDiskSizeGB sets DiskSizeGB field to given value.


### GetSubdomainPermanent

`func (o *ContainerInstanceGroup) GetSubdomainPermanent() string`

GetSubdomainPermanent returns the SubdomainPermanent field if non-nil, zero value otherwise.

### GetSubdomainPermanentOk

`func (o *ContainerInstanceGroup) GetSubdomainPermanentOk() (*string, bool)`

GetSubdomainPermanentOk returns a tuple with the SubdomainPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomainPermanent

`func (o *ContainerInstanceGroup) SetSubdomainPermanent(v string)`

SetSubdomainPermanent sets SubdomainPermanent field to given value.

### HasSubdomainPermanent

`func (o *ContainerInstanceGroup) HasSubdomainPermanent() bool`

HasSubdomainPermanent returns a boolean if a field has been set.

### GetDnsSubdomainId

`func (o *ContainerInstanceGroup) GetDnsSubdomainId() int64`

GetDnsSubdomainId returns the DnsSubdomainId field if non-nil, zero value otherwise.

### GetDnsSubdomainIdOk

`func (o *ContainerInstanceGroup) GetDnsSubdomainIdOk() (*int64, bool)`

GetDnsSubdomainIdOk returns a tuple with the DnsSubdomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainId

`func (o *ContainerInstanceGroup) SetDnsSubdomainId(v int64)`

SetDnsSubdomainId sets DnsSubdomainId field to given value.

### HasDnsSubdomainId

`func (o *ContainerInstanceGroup) HasDnsSubdomainId() bool`

HasDnsSubdomainId returns a boolean if a field has been set.

### GetDnsSubdomainPermanentId

`func (o *ContainerInstanceGroup) GetDnsSubdomainPermanentId() int64`

GetDnsSubdomainPermanentId returns the DnsSubdomainPermanentId field if non-nil, zero value otherwise.

### GetDnsSubdomainPermanentIdOk

`func (o *ContainerInstanceGroup) GetDnsSubdomainPermanentIdOk() (*int64, bool)`

GetDnsSubdomainPermanentIdOk returns a tuple with the DnsSubdomainPermanentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainPermanentId

`func (o *ContainerInstanceGroup) SetDnsSubdomainPermanentId(v int64)`

SetDnsSubdomainPermanentId sets DnsSubdomainPermanentId field to given value.

### HasDnsSubdomainPermanentId

`func (o *ContainerInstanceGroup) HasDnsSubdomainPermanentId() bool`

HasDnsSubdomainPermanentId returns a boolean if a field has been set.

### GetVmPoolId

`func (o *ContainerInstanceGroup) GetVmPoolId() int64`

GetVmPoolId returns the VmPoolId field if non-nil, zero value otherwise.

### GetVmPoolIdOk

`func (o *ContainerInstanceGroup) GetVmPoolIdOk() (*int64, bool)`

GetVmPoolIdOk returns a tuple with the VmPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmPoolId

`func (o *ContainerInstanceGroup) SetVmPoolId(v int64)`

SetVmPoolId sets VmPoolId field to given value.

### HasVmPoolId

`func (o *ContainerInstanceGroup) HasVmPoolId() bool`

HasVmPoolId returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *ContainerInstanceGroup) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *ContainerInstanceGroup) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *ContainerInstanceGroup) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetConfig

`func (o *ContainerInstanceGroup) GetConfig() VMInstanceGroupConfiguration`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ContainerInstanceGroup) GetConfigOk() (*VMInstanceGroupConfiguration, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ContainerInstanceGroup) SetConfig(v VMInstanceGroupConfiguration)`

SetConfig sets Config field to given value.


### GetMeta

`func (o *ContainerInstanceGroup) GetMeta() VMInstanceGroupMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ContainerInstanceGroup) GetMetaOk() (*VMInstanceGroupMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ContainerInstanceGroup) SetMeta(v VMInstanceGroupMeta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *ContainerInstanceGroup) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ContainerInstanceGroup) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ContainerInstanceGroup) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ContainerInstanceGroup) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


