# VMInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Name of the VM Instance. | 
**VmId** | Pointer to **float32** | Id of the VM. | [optional] 
**TypeId** | **float32** | Id of the VM Type. | 
**Subdomain** | Pointer to **string** | Subdomain of the VM Instance. | [optional] 
**VmPoolId** | Pointer to **float32** | Id of the VM Pool. | [optional] 
**DiskSizeGB** | **float32** | Disk size in GB of the VM Instance. | 
**RamGB** | **float32** | RAM size in GB of the VM Instance. | 
**CpuCores** | **float32** | Number of CPU cores for the VM Instance. | 
**OsTemplateId** | Pointer to **float32** | Id of the template used by the VM Instance. | [optional] 
**CustomVariables** | Pointer to **map[string]interface{}** | Custom variables for the VM Instance. | [optional] 
**UpdatedTimestamp** | **string** | Timestamp of the VM Instance last update. | 
**Id** | **float32** | VM Instance ID | 
**Revision** | **float32** | Revision of the VM Instance | 
**GroupId** | **float32** | Id of the VM Instance Group. | 
**InfrastructureId** | **float32** | Id of the Infrastructure. | 
**ServiceStatus** | **string** | Service status of the VM Instance. | 
**SubdomainPermanent** | Pointer to **string** | Subdomain permanent of the VM Instance. | [optional] 
**DnsSubdomainId** | Pointer to **float32** | Id of the DNS subdomain for the VM Instance. | [optional] 
**DnsSubdomainPermanentId** | Pointer to **float32** | Id of the permanent DNS subdomain for the VM Instance. | [optional] 
**UniqueIndex** | Pointer to **float32** | Unique index of the VM Instance. | [optional] 
**CreatedTimestamp** | **string** | Timestamp of the VM Instance creation. | 
**Config** | [**VMInstanceConfiguration**](VMInstanceConfiguration.md) | The current changes to be deployed for the VM Instance. | 
**Meta** | [**VMInstanceMeta**](VMInstanceMeta.md) | Meta information of the VM Instance. | 
**Links** | **map[string]interface{}** | Links to other resources | 

## Methods

### NewVMInstance

`func NewVMInstance(label string, typeId float32, diskSizeGB float32, ramGB float32, cpuCores float32, updatedTimestamp string, id float32, revision float32, groupId float32, infrastructureId float32, serviceStatus string, createdTimestamp string, config VMInstanceConfiguration, meta VMInstanceMeta, links map[string]interface{}, ) *VMInstance`

NewVMInstance instantiates a new VMInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMInstanceWithDefaults

`func NewVMInstanceWithDefaults() *VMInstance`

NewVMInstanceWithDefaults instantiates a new VMInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *VMInstance) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *VMInstance) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *VMInstance) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetVmId

`func (o *VMInstance) GetVmId() float32`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *VMInstance) GetVmIdOk() (*float32, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *VMInstance) SetVmId(v float32)`

SetVmId sets VmId field to given value.

### HasVmId

`func (o *VMInstance) HasVmId() bool`

HasVmId returns a boolean if a field has been set.

### GetTypeId

`func (o *VMInstance) GetTypeId() float32`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *VMInstance) GetTypeIdOk() (*float32, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *VMInstance) SetTypeId(v float32)`

SetTypeId sets TypeId field to given value.


### GetSubdomain

`func (o *VMInstance) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *VMInstance) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *VMInstance) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *VMInstance) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetVmPoolId

`func (o *VMInstance) GetVmPoolId() float32`

GetVmPoolId returns the VmPoolId field if non-nil, zero value otherwise.

### GetVmPoolIdOk

`func (o *VMInstance) GetVmPoolIdOk() (*float32, bool)`

GetVmPoolIdOk returns a tuple with the VmPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmPoolId

`func (o *VMInstance) SetVmPoolId(v float32)`

SetVmPoolId sets VmPoolId field to given value.

### HasVmPoolId

`func (o *VMInstance) HasVmPoolId() bool`

HasVmPoolId returns a boolean if a field has been set.

### GetDiskSizeGB

`func (o *VMInstance) GetDiskSizeGB() float32`

GetDiskSizeGB returns the DiskSizeGB field if non-nil, zero value otherwise.

### GetDiskSizeGBOk

`func (o *VMInstance) GetDiskSizeGBOk() (*float32, bool)`

GetDiskSizeGBOk returns a tuple with the DiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeGB

`func (o *VMInstance) SetDiskSizeGB(v float32)`

SetDiskSizeGB sets DiskSizeGB field to given value.


### GetRamGB

`func (o *VMInstance) GetRamGB() float32`

GetRamGB returns the RamGB field if non-nil, zero value otherwise.

### GetRamGBOk

`func (o *VMInstance) GetRamGBOk() (*float32, bool)`

GetRamGBOk returns a tuple with the RamGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGB

`func (o *VMInstance) SetRamGB(v float32)`

SetRamGB sets RamGB field to given value.


### GetCpuCores

`func (o *VMInstance) GetCpuCores() float32`

GetCpuCores returns the CpuCores field if non-nil, zero value otherwise.

### GetCpuCoresOk

`func (o *VMInstance) GetCpuCoresOk() (*float32, bool)`

GetCpuCoresOk returns a tuple with the CpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCores

`func (o *VMInstance) SetCpuCores(v float32)`

SetCpuCores sets CpuCores field to given value.


### GetOsTemplateId

`func (o *VMInstance) GetOsTemplateId() float32`

GetOsTemplateId returns the OsTemplateId field if non-nil, zero value otherwise.

### GetOsTemplateIdOk

`func (o *VMInstance) GetOsTemplateIdOk() (*float32, bool)`

GetOsTemplateIdOk returns a tuple with the OsTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsTemplateId

`func (o *VMInstance) SetOsTemplateId(v float32)`

SetOsTemplateId sets OsTemplateId field to given value.

### HasOsTemplateId

`func (o *VMInstance) HasOsTemplateId() bool`

HasOsTemplateId returns a boolean if a field has been set.

### GetCustomVariables

`func (o *VMInstance) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *VMInstance) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *VMInstance) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *VMInstance) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *VMInstance) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *VMInstance) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *VMInstance) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetId

`func (o *VMInstance) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VMInstance) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VMInstance) SetId(v float32)`

SetId sets Id field to given value.


### GetRevision

`func (o *VMInstance) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *VMInstance) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *VMInstance) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetGroupId

`func (o *VMInstance) GetGroupId() float32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *VMInstance) GetGroupIdOk() (*float32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *VMInstance) SetGroupId(v float32)`

SetGroupId sets GroupId field to given value.


### GetInfrastructureId

`func (o *VMInstance) GetInfrastructureId() float32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *VMInstance) GetInfrastructureIdOk() (*float32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *VMInstance) SetInfrastructureId(v float32)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetServiceStatus

`func (o *VMInstance) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *VMInstance) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *VMInstance) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetSubdomainPermanent

`func (o *VMInstance) GetSubdomainPermanent() string`

GetSubdomainPermanent returns the SubdomainPermanent field if non-nil, zero value otherwise.

### GetSubdomainPermanentOk

`func (o *VMInstance) GetSubdomainPermanentOk() (*string, bool)`

GetSubdomainPermanentOk returns a tuple with the SubdomainPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomainPermanent

`func (o *VMInstance) SetSubdomainPermanent(v string)`

SetSubdomainPermanent sets SubdomainPermanent field to given value.

### HasSubdomainPermanent

`func (o *VMInstance) HasSubdomainPermanent() bool`

HasSubdomainPermanent returns a boolean if a field has been set.

### GetDnsSubdomainId

`func (o *VMInstance) GetDnsSubdomainId() float32`

GetDnsSubdomainId returns the DnsSubdomainId field if non-nil, zero value otherwise.

### GetDnsSubdomainIdOk

`func (o *VMInstance) GetDnsSubdomainIdOk() (*float32, bool)`

GetDnsSubdomainIdOk returns a tuple with the DnsSubdomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainId

`func (o *VMInstance) SetDnsSubdomainId(v float32)`

SetDnsSubdomainId sets DnsSubdomainId field to given value.

### HasDnsSubdomainId

`func (o *VMInstance) HasDnsSubdomainId() bool`

HasDnsSubdomainId returns a boolean if a field has been set.

### GetDnsSubdomainPermanentId

`func (o *VMInstance) GetDnsSubdomainPermanentId() float32`

GetDnsSubdomainPermanentId returns the DnsSubdomainPermanentId field if non-nil, zero value otherwise.

### GetDnsSubdomainPermanentIdOk

`func (o *VMInstance) GetDnsSubdomainPermanentIdOk() (*float32, bool)`

GetDnsSubdomainPermanentIdOk returns a tuple with the DnsSubdomainPermanentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainPermanentId

`func (o *VMInstance) SetDnsSubdomainPermanentId(v float32)`

SetDnsSubdomainPermanentId sets DnsSubdomainPermanentId field to given value.

### HasDnsSubdomainPermanentId

`func (o *VMInstance) HasDnsSubdomainPermanentId() bool`

HasDnsSubdomainPermanentId returns a boolean if a field has been set.

### GetUniqueIndex

`func (o *VMInstance) GetUniqueIndex() float32`

GetUniqueIndex returns the UniqueIndex field if non-nil, zero value otherwise.

### GetUniqueIndexOk

`func (o *VMInstance) GetUniqueIndexOk() (*float32, bool)`

GetUniqueIndexOk returns a tuple with the UniqueIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIndex

`func (o *VMInstance) SetUniqueIndex(v float32)`

SetUniqueIndex sets UniqueIndex field to given value.

### HasUniqueIndex

`func (o *VMInstance) HasUniqueIndex() bool`

HasUniqueIndex returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *VMInstance) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *VMInstance) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *VMInstance) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetConfig

`func (o *VMInstance) GetConfig() VMInstanceConfiguration`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *VMInstance) GetConfigOk() (*VMInstanceConfiguration, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *VMInstance) SetConfig(v VMInstanceConfiguration)`

SetConfig sets Config field to given value.


### GetMeta

`func (o *VMInstance) GetMeta() VMInstanceMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *VMInstance) GetMetaOk() (*VMInstanceMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *VMInstance) SetMeta(v VMInstanceMeta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *VMInstance) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VMInstance) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VMInstance) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


