# VMInstanceGroupConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | **float32** | Revision of the VM Instance Group Configuration | 
**Label** | **string** | Name of the VM Instance Group. | 
**Subdomain** | Pointer to **string** | Subdomain of the VM Instance Group. | [optional] 
**DnsSubdomainChangeId** | Pointer to **float32** | Id of the DNS subdomain for the VM Instance Group. | [optional] 
**InstanceCount** | Pointer to **float32** |  | [optional] [default to 1]
**DeployType** | **string** | Deploy type of the VM Instance Group | [default to "create"]
**DeployStatus** | **string** | Deploy status of the VM Instance Group | [default to "not_started"]
**InfrastructureDeployId** | Pointer to **float32** | Id of the deployment for the VM Instance Group. | [optional] 
**CustomVariables** | Pointer to **map[string]interface{}** | Custom variables for the VM Instance Group. | [optional] 
**UpdatedTimestamp** | **string** | Timestamp of the VM Instance Group last update. | 

## Methods

### NewVMInstanceGroupConfiguration

`func NewVMInstanceGroupConfiguration(revision float32, label string, deployType string, deployStatus string, updatedTimestamp string, ) *VMInstanceGroupConfiguration`

NewVMInstanceGroupConfiguration instantiates a new VMInstanceGroupConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMInstanceGroupConfigurationWithDefaults

`func NewVMInstanceGroupConfigurationWithDefaults() *VMInstanceGroupConfiguration`

NewVMInstanceGroupConfigurationWithDefaults instantiates a new VMInstanceGroupConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *VMInstanceGroupConfiguration) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *VMInstanceGroupConfiguration) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *VMInstanceGroupConfiguration) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetLabel

`func (o *VMInstanceGroupConfiguration) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *VMInstanceGroupConfiguration) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *VMInstanceGroupConfiguration) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetSubdomain

`func (o *VMInstanceGroupConfiguration) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *VMInstanceGroupConfiguration) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *VMInstanceGroupConfiguration) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *VMInstanceGroupConfiguration) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetDnsSubdomainChangeId

`func (o *VMInstanceGroupConfiguration) GetDnsSubdomainChangeId() float32`

GetDnsSubdomainChangeId returns the DnsSubdomainChangeId field if non-nil, zero value otherwise.

### GetDnsSubdomainChangeIdOk

`func (o *VMInstanceGroupConfiguration) GetDnsSubdomainChangeIdOk() (*float32, bool)`

GetDnsSubdomainChangeIdOk returns a tuple with the DnsSubdomainChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainChangeId

`func (o *VMInstanceGroupConfiguration) SetDnsSubdomainChangeId(v float32)`

SetDnsSubdomainChangeId sets DnsSubdomainChangeId field to given value.

### HasDnsSubdomainChangeId

`func (o *VMInstanceGroupConfiguration) HasDnsSubdomainChangeId() bool`

HasDnsSubdomainChangeId returns a boolean if a field has been set.

### GetInstanceCount

`func (o *VMInstanceGroupConfiguration) GetInstanceCount() float32`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *VMInstanceGroupConfiguration) GetInstanceCountOk() (*float32, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *VMInstanceGroupConfiguration) SetInstanceCount(v float32)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *VMInstanceGroupConfiguration) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

### GetDeployType

`func (o *VMInstanceGroupConfiguration) GetDeployType() string`

GetDeployType returns the DeployType field if non-nil, zero value otherwise.

### GetDeployTypeOk

`func (o *VMInstanceGroupConfiguration) GetDeployTypeOk() (*string, bool)`

GetDeployTypeOk returns a tuple with the DeployType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployType

`func (o *VMInstanceGroupConfiguration) SetDeployType(v string)`

SetDeployType sets DeployType field to given value.


### GetDeployStatus

`func (o *VMInstanceGroupConfiguration) GetDeployStatus() string`

GetDeployStatus returns the DeployStatus field if non-nil, zero value otherwise.

### GetDeployStatusOk

`func (o *VMInstanceGroupConfiguration) GetDeployStatusOk() (*string, bool)`

GetDeployStatusOk returns a tuple with the DeployStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStatus

`func (o *VMInstanceGroupConfiguration) SetDeployStatus(v string)`

SetDeployStatus sets DeployStatus field to given value.


### GetInfrastructureDeployId

`func (o *VMInstanceGroupConfiguration) GetInfrastructureDeployId() float32`

GetInfrastructureDeployId returns the InfrastructureDeployId field if non-nil, zero value otherwise.

### GetInfrastructureDeployIdOk

`func (o *VMInstanceGroupConfiguration) GetInfrastructureDeployIdOk() (*float32, bool)`

GetInfrastructureDeployIdOk returns a tuple with the InfrastructureDeployId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureDeployId

`func (o *VMInstanceGroupConfiguration) SetInfrastructureDeployId(v float32)`

SetInfrastructureDeployId sets InfrastructureDeployId field to given value.

### HasInfrastructureDeployId

`func (o *VMInstanceGroupConfiguration) HasInfrastructureDeployId() bool`

HasInfrastructureDeployId returns a boolean if a field has been set.

### GetCustomVariables

`func (o *VMInstanceGroupConfiguration) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *VMInstanceGroupConfiguration) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *VMInstanceGroupConfiguration) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *VMInstanceGroupConfiguration) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *VMInstanceGroupConfiguration) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *VMInstanceGroupConfiguration) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *VMInstanceGroupConfiguration) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


