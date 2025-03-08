# VMInstanceGroupInterfaceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | **float32** | Revision of the VM Instance Group Interface Configuration | 
**Label** | **string** | Name of the VM Instance Group Interface. | 
**Index** | **float32** | Interface index | 
**NetworkId** | Pointer to **float32** | Network ID | [optional] 
**Subdomain** | Pointer to **string** | Subdomain of the VM Instance Group Interface. | [optional] 
**DnsSubdomainChangeId** | Pointer to **float32** | Id of the DNS subdomain for the VM Instance Group Interface. | [optional] 
**DeployType** | **string** | Deploy type of the VM Instance Group Interface | [default to "create"]
**DeployStatus** | **string** | Deploy status of the VM Instance Group Interface | [default to "not_started"]
**InfrastructureDeployId** | Pointer to **float32** | Id of the deployment for the VM Instance Group Interface. | [optional] 
**UpdatedTimestamp** | **string** | Timestamp of the VM Instance Group Interface update. | 

## Methods

### NewVMInstanceGroupInterfaceConfiguration

`func NewVMInstanceGroupInterfaceConfiguration(revision float32, label string, index float32, deployType string, deployStatus string, updatedTimestamp string, ) *VMInstanceGroupInterfaceConfiguration`

NewVMInstanceGroupInterfaceConfiguration instantiates a new VMInstanceGroupInterfaceConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMInstanceGroupInterfaceConfigurationWithDefaults

`func NewVMInstanceGroupInterfaceConfigurationWithDefaults() *VMInstanceGroupInterfaceConfiguration`

NewVMInstanceGroupInterfaceConfigurationWithDefaults instantiates a new VMInstanceGroupInterfaceConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *VMInstanceGroupInterfaceConfiguration) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *VMInstanceGroupInterfaceConfiguration) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *VMInstanceGroupInterfaceConfiguration) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetLabel

`func (o *VMInstanceGroupInterfaceConfiguration) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *VMInstanceGroupInterfaceConfiguration) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *VMInstanceGroupInterfaceConfiguration) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetIndex

`func (o *VMInstanceGroupInterfaceConfiguration) GetIndex() float32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *VMInstanceGroupInterfaceConfiguration) GetIndexOk() (*float32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *VMInstanceGroupInterfaceConfiguration) SetIndex(v float32)`

SetIndex sets Index field to given value.


### GetNetworkId

`func (o *VMInstanceGroupInterfaceConfiguration) GetNetworkId() float32`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *VMInstanceGroupInterfaceConfiguration) GetNetworkIdOk() (*float32, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *VMInstanceGroupInterfaceConfiguration) SetNetworkId(v float32)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *VMInstanceGroupInterfaceConfiguration) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetSubdomain

`func (o *VMInstanceGroupInterfaceConfiguration) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *VMInstanceGroupInterfaceConfiguration) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *VMInstanceGroupInterfaceConfiguration) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *VMInstanceGroupInterfaceConfiguration) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetDnsSubdomainChangeId

`func (o *VMInstanceGroupInterfaceConfiguration) GetDnsSubdomainChangeId() float32`

GetDnsSubdomainChangeId returns the DnsSubdomainChangeId field if non-nil, zero value otherwise.

### GetDnsSubdomainChangeIdOk

`func (o *VMInstanceGroupInterfaceConfiguration) GetDnsSubdomainChangeIdOk() (*float32, bool)`

GetDnsSubdomainChangeIdOk returns a tuple with the DnsSubdomainChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainChangeId

`func (o *VMInstanceGroupInterfaceConfiguration) SetDnsSubdomainChangeId(v float32)`

SetDnsSubdomainChangeId sets DnsSubdomainChangeId field to given value.

### HasDnsSubdomainChangeId

`func (o *VMInstanceGroupInterfaceConfiguration) HasDnsSubdomainChangeId() bool`

HasDnsSubdomainChangeId returns a boolean if a field has been set.

### GetDeployType

`func (o *VMInstanceGroupInterfaceConfiguration) GetDeployType() string`

GetDeployType returns the DeployType field if non-nil, zero value otherwise.

### GetDeployTypeOk

`func (o *VMInstanceGroupInterfaceConfiguration) GetDeployTypeOk() (*string, bool)`

GetDeployTypeOk returns a tuple with the DeployType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployType

`func (o *VMInstanceGroupInterfaceConfiguration) SetDeployType(v string)`

SetDeployType sets DeployType field to given value.


### GetDeployStatus

`func (o *VMInstanceGroupInterfaceConfiguration) GetDeployStatus() string`

GetDeployStatus returns the DeployStatus field if non-nil, zero value otherwise.

### GetDeployStatusOk

`func (o *VMInstanceGroupInterfaceConfiguration) GetDeployStatusOk() (*string, bool)`

GetDeployStatusOk returns a tuple with the DeployStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStatus

`func (o *VMInstanceGroupInterfaceConfiguration) SetDeployStatus(v string)`

SetDeployStatus sets DeployStatus field to given value.


### GetInfrastructureDeployId

`func (o *VMInstanceGroupInterfaceConfiguration) GetInfrastructureDeployId() float32`

GetInfrastructureDeployId returns the InfrastructureDeployId field if non-nil, zero value otherwise.

### GetInfrastructureDeployIdOk

`func (o *VMInstanceGroupInterfaceConfiguration) GetInfrastructureDeployIdOk() (*float32, bool)`

GetInfrastructureDeployIdOk returns a tuple with the InfrastructureDeployId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureDeployId

`func (o *VMInstanceGroupInterfaceConfiguration) SetInfrastructureDeployId(v float32)`

SetInfrastructureDeployId sets InfrastructureDeployId field to given value.

### HasInfrastructureDeployId

`func (o *VMInstanceGroupInterfaceConfiguration) HasInfrastructureDeployId() bool`

HasInfrastructureDeployId returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *VMInstanceGroupInterfaceConfiguration) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *VMInstanceGroupInterfaceConfiguration) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *VMInstanceGroupInterfaceConfiguration) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


