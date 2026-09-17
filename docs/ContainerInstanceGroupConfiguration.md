# ContainerInstanceGroupConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | **int64** | Revision of the VM Instance Group Configuration | 
**Label** | **string** | Name of the VM Instance Group. | 
**Subdomain** | Pointer to **string** | Subdomain of the VM Instance Group. | [optional] 
**DnsSubdomainChangeId** | Pointer to **int64** | Id of the DNS subdomain for the VM Instance Group. | [optional] 
**InstanceCount** | Pointer to **float32** |  | [optional] [default to 1]
**DeployType** | **string** | Deploy type of the VM Instance Group | [default to "create"]
**DeployStatus** | **string** | Deploy status of the VM Instance Group | [default to "not_started"]
**InfrastructureDeployId** | Pointer to **int64** | Id of the deployment for the VM Instance Group. | [optional] 
**CustomVariables** | Pointer to **map[string]interface{}** | Custom variables for the VM Instance Group. | [optional] 
**UpdatedTimestamp** | **string** | Timestamp of the VM Instance Group last update. | 

## Methods

### NewContainerInstanceGroupConfiguration

`func NewContainerInstanceGroupConfiguration(revision int64, label string, deployType string, deployStatus string, updatedTimestamp string, ) *ContainerInstanceGroupConfiguration`

NewContainerInstanceGroupConfiguration instantiates a new ContainerInstanceGroupConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerInstanceGroupConfigurationWithDefaults

`func NewContainerInstanceGroupConfigurationWithDefaults() *ContainerInstanceGroupConfiguration`

NewContainerInstanceGroupConfigurationWithDefaults instantiates a new ContainerInstanceGroupConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *ContainerInstanceGroupConfiguration) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ContainerInstanceGroupConfiguration) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ContainerInstanceGroupConfiguration) SetRevision(v int64)`

SetRevision sets Revision field to given value.


### GetLabel

`func (o *ContainerInstanceGroupConfiguration) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ContainerInstanceGroupConfiguration) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ContainerInstanceGroupConfiguration) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetSubdomain

`func (o *ContainerInstanceGroupConfiguration) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *ContainerInstanceGroupConfiguration) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *ContainerInstanceGroupConfiguration) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *ContainerInstanceGroupConfiguration) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetDnsSubdomainChangeId

`func (o *ContainerInstanceGroupConfiguration) GetDnsSubdomainChangeId() int64`

GetDnsSubdomainChangeId returns the DnsSubdomainChangeId field if non-nil, zero value otherwise.

### GetDnsSubdomainChangeIdOk

`func (o *ContainerInstanceGroupConfiguration) GetDnsSubdomainChangeIdOk() (*int64, bool)`

GetDnsSubdomainChangeIdOk returns a tuple with the DnsSubdomainChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainChangeId

`func (o *ContainerInstanceGroupConfiguration) SetDnsSubdomainChangeId(v int64)`

SetDnsSubdomainChangeId sets DnsSubdomainChangeId field to given value.

### HasDnsSubdomainChangeId

`func (o *ContainerInstanceGroupConfiguration) HasDnsSubdomainChangeId() bool`

HasDnsSubdomainChangeId returns a boolean if a field has been set.

### GetInstanceCount

`func (o *ContainerInstanceGroupConfiguration) GetInstanceCount() float32`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *ContainerInstanceGroupConfiguration) GetInstanceCountOk() (*float32, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *ContainerInstanceGroupConfiguration) SetInstanceCount(v float32)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *ContainerInstanceGroupConfiguration) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

### GetDeployType

`func (o *ContainerInstanceGroupConfiguration) GetDeployType() string`

GetDeployType returns the DeployType field if non-nil, zero value otherwise.

### GetDeployTypeOk

`func (o *ContainerInstanceGroupConfiguration) GetDeployTypeOk() (*string, bool)`

GetDeployTypeOk returns a tuple with the DeployType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployType

`func (o *ContainerInstanceGroupConfiguration) SetDeployType(v string)`

SetDeployType sets DeployType field to given value.


### GetDeployStatus

`func (o *ContainerInstanceGroupConfiguration) GetDeployStatus() string`

GetDeployStatus returns the DeployStatus field if non-nil, zero value otherwise.

### GetDeployStatusOk

`func (o *ContainerInstanceGroupConfiguration) GetDeployStatusOk() (*string, bool)`

GetDeployStatusOk returns a tuple with the DeployStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStatus

`func (o *ContainerInstanceGroupConfiguration) SetDeployStatus(v string)`

SetDeployStatus sets DeployStatus field to given value.


### GetInfrastructureDeployId

`func (o *ContainerInstanceGroupConfiguration) GetInfrastructureDeployId() int64`

GetInfrastructureDeployId returns the InfrastructureDeployId field if non-nil, zero value otherwise.

### GetInfrastructureDeployIdOk

`func (o *ContainerInstanceGroupConfiguration) GetInfrastructureDeployIdOk() (*int64, bool)`

GetInfrastructureDeployIdOk returns a tuple with the InfrastructureDeployId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureDeployId

`func (o *ContainerInstanceGroupConfiguration) SetInfrastructureDeployId(v int64)`

SetInfrastructureDeployId sets InfrastructureDeployId field to given value.

### HasInfrastructureDeployId

`func (o *ContainerInstanceGroupConfiguration) HasInfrastructureDeployId() bool`

HasInfrastructureDeployId returns a boolean if a field has been set.

### GetCustomVariables

`func (o *ContainerInstanceGroupConfiguration) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *ContainerInstanceGroupConfiguration) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *ContainerInstanceGroupConfiguration) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *ContainerInstanceGroupConfiguration) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *ContainerInstanceGroupConfiguration) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *ContainerInstanceGroupConfiguration) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *ContainerInstanceGroupConfiguration) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


