# VMInstanceGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Name of the VM Instance Group. | 
**Subdomain** | Pointer to **string** | Subdomain of the VM Instance Group. | [optional] 
**InstanceCount** | Pointer to **float32** |  | [optional] [default to 1]
**CustomVariables** | Pointer to **map[string]interface{}** | Custom variables for the VM Instance Group. | [optional] 
**UpdatedTimestamp** | **string** | Timestamp of the VM Instance Group last update. | 
**Id** | **float32** | Id of the VM Instance Group. | 
**Revision** | **float32** | Revision of the VM Instance Group State | 
**InfrastructureId** | **float32** | Id of the Infrastructure. | 
**ServiceStatus** | **string** | Status of the VM Instance Group. | 
**DiskSizeGB** | **float32** | Disk size in GB for each VM Instance in the VM Instance Group. | 
**SubdomainPermanent** | Pointer to **string** | Subdomain permanent of the VM Instance Group. | [optional] 
**DnsSubdomainId** | Pointer to **float32** | Id of the DNS subdomain for the VM Instance Group. | [optional] 
**DnsSubdomainPermanentId** | Pointer to **float32** | Id of the permanent DNS subdomain for the VM Instance Group. | [optional] 
**CreatedTimestamp** | **string** | Timestamp of the VM Instance Group creation. | 
**Config** | [**VMInstanceGroupConfiguration**](VMInstanceGroupConfiguration.md) | The current changes to be deployed for the VM Instance Group. | 
**Meta** | [**VMInstanceGroupMeta**](VMInstanceGroupMeta.md) | Meta information of the VM Instance Group. | 
**Links** | **map[string]interface{}** | Links to other resources | 

## Methods

### NewVMInstanceGroup

`func NewVMInstanceGroup(label string, updatedTimestamp string, id float32, revision float32, infrastructureId float32, serviceStatus string, diskSizeGB float32, createdTimestamp string, config VMInstanceGroupConfiguration, meta VMInstanceGroupMeta, links map[string]interface{}, ) *VMInstanceGroup`

NewVMInstanceGroup instantiates a new VMInstanceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMInstanceGroupWithDefaults

`func NewVMInstanceGroupWithDefaults() *VMInstanceGroup`

NewVMInstanceGroupWithDefaults instantiates a new VMInstanceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *VMInstanceGroup) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *VMInstanceGroup) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *VMInstanceGroup) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetSubdomain

`func (o *VMInstanceGroup) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *VMInstanceGroup) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *VMInstanceGroup) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *VMInstanceGroup) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetInstanceCount

`func (o *VMInstanceGroup) GetInstanceCount() float32`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *VMInstanceGroup) GetInstanceCountOk() (*float32, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *VMInstanceGroup) SetInstanceCount(v float32)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *VMInstanceGroup) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

### GetCustomVariables

`func (o *VMInstanceGroup) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *VMInstanceGroup) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *VMInstanceGroup) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *VMInstanceGroup) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *VMInstanceGroup) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *VMInstanceGroup) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *VMInstanceGroup) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetId

`func (o *VMInstanceGroup) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VMInstanceGroup) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VMInstanceGroup) SetId(v float32)`

SetId sets Id field to given value.


### GetRevision

`func (o *VMInstanceGroup) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *VMInstanceGroup) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *VMInstanceGroup) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetInfrastructureId

`func (o *VMInstanceGroup) GetInfrastructureId() float32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *VMInstanceGroup) GetInfrastructureIdOk() (*float32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *VMInstanceGroup) SetInfrastructureId(v float32)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetServiceStatus

`func (o *VMInstanceGroup) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *VMInstanceGroup) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *VMInstanceGroup) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetDiskSizeGB

`func (o *VMInstanceGroup) GetDiskSizeGB() float32`

GetDiskSizeGB returns the DiskSizeGB field if non-nil, zero value otherwise.

### GetDiskSizeGBOk

`func (o *VMInstanceGroup) GetDiskSizeGBOk() (*float32, bool)`

GetDiskSizeGBOk returns a tuple with the DiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeGB

`func (o *VMInstanceGroup) SetDiskSizeGB(v float32)`

SetDiskSizeGB sets DiskSizeGB field to given value.


### GetSubdomainPermanent

`func (o *VMInstanceGroup) GetSubdomainPermanent() string`

GetSubdomainPermanent returns the SubdomainPermanent field if non-nil, zero value otherwise.

### GetSubdomainPermanentOk

`func (o *VMInstanceGroup) GetSubdomainPermanentOk() (*string, bool)`

GetSubdomainPermanentOk returns a tuple with the SubdomainPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomainPermanent

`func (o *VMInstanceGroup) SetSubdomainPermanent(v string)`

SetSubdomainPermanent sets SubdomainPermanent field to given value.

### HasSubdomainPermanent

`func (o *VMInstanceGroup) HasSubdomainPermanent() bool`

HasSubdomainPermanent returns a boolean if a field has been set.

### GetDnsSubdomainId

`func (o *VMInstanceGroup) GetDnsSubdomainId() float32`

GetDnsSubdomainId returns the DnsSubdomainId field if non-nil, zero value otherwise.

### GetDnsSubdomainIdOk

`func (o *VMInstanceGroup) GetDnsSubdomainIdOk() (*float32, bool)`

GetDnsSubdomainIdOk returns a tuple with the DnsSubdomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainId

`func (o *VMInstanceGroup) SetDnsSubdomainId(v float32)`

SetDnsSubdomainId sets DnsSubdomainId field to given value.

### HasDnsSubdomainId

`func (o *VMInstanceGroup) HasDnsSubdomainId() bool`

HasDnsSubdomainId returns a boolean if a field has been set.

### GetDnsSubdomainPermanentId

`func (o *VMInstanceGroup) GetDnsSubdomainPermanentId() float32`

GetDnsSubdomainPermanentId returns the DnsSubdomainPermanentId field if non-nil, zero value otherwise.

### GetDnsSubdomainPermanentIdOk

`func (o *VMInstanceGroup) GetDnsSubdomainPermanentIdOk() (*float32, bool)`

GetDnsSubdomainPermanentIdOk returns a tuple with the DnsSubdomainPermanentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainPermanentId

`func (o *VMInstanceGroup) SetDnsSubdomainPermanentId(v float32)`

SetDnsSubdomainPermanentId sets DnsSubdomainPermanentId field to given value.

### HasDnsSubdomainPermanentId

`func (o *VMInstanceGroup) HasDnsSubdomainPermanentId() bool`

HasDnsSubdomainPermanentId returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *VMInstanceGroup) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *VMInstanceGroup) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *VMInstanceGroup) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetConfig

`func (o *VMInstanceGroup) GetConfig() VMInstanceGroupConfiguration`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *VMInstanceGroup) GetConfigOk() (*VMInstanceGroupConfiguration, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *VMInstanceGroup) SetConfig(v VMInstanceGroupConfiguration)`

SetConfig sets Config field to given value.


### GetMeta

`func (o *VMInstanceGroup) GetMeta() VMInstanceGroupMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *VMInstanceGroup) GetMetaOk() (*VMInstanceGroupMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *VMInstanceGroup) SetMeta(v VMInstanceGroupMeta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *VMInstanceGroup) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VMInstanceGroup) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VMInstanceGroup) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


