# ExtensionInstanceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | **float32** | Revision number | 
**Label** | **string** | The extension instance label. Will be automatically generated if not provided. | 
**AutomaticManagement** | **float32** | Flag specifying if the extension instance supports automatic management. | 
**Subdomain** | Pointer to **string** | Subdomain of the Extension Instance. | [optional] 
**DnsSubdomainChangeId** | Pointer to **float32** | Id of the DNS subdomain for the Extension Instance. | [optional] 
**DeployType** | **string** | Deploy type of the Extension Instance | [default to "create"]
**DeployStatus** | **string** | Deploy status of the Extension Instance | [default to "not_started"]
**InfrastructureDeployId** | Pointer to **float32** | Id of the deployment for the Extension Instance. | [optional] 
**UpdatedTimestamp** | **string** | Timestamp of the Extension Instance last update. | 

## Methods

### NewExtensionInstanceConfiguration

`func NewExtensionInstanceConfiguration(revision float32, label string, automaticManagement float32, deployType string, deployStatus string, updatedTimestamp string, ) *ExtensionInstanceConfiguration`

NewExtensionInstanceConfiguration instantiates a new ExtensionInstanceConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionInstanceConfigurationWithDefaults

`func NewExtensionInstanceConfigurationWithDefaults() *ExtensionInstanceConfiguration`

NewExtensionInstanceConfigurationWithDefaults instantiates a new ExtensionInstanceConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *ExtensionInstanceConfiguration) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ExtensionInstanceConfiguration) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ExtensionInstanceConfiguration) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetLabel

`func (o *ExtensionInstanceConfiguration) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ExtensionInstanceConfiguration) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ExtensionInstanceConfiguration) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetAutomaticManagement

`func (o *ExtensionInstanceConfiguration) GetAutomaticManagement() float32`

GetAutomaticManagement returns the AutomaticManagement field if non-nil, zero value otherwise.

### GetAutomaticManagementOk

`func (o *ExtensionInstanceConfiguration) GetAutomaticManagementOk() (*float32, bool)`

GetAutomaticManagementOk returns a tuple with the AutomaticManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticManagement

`func (o *ExtensionInstanceConfiguration) SetAutomaticManagement(v float32)`

SetAutomaticManagement sets AutomaticManagement field to given value.


### GetSubdomain

`func (o *ExtensionInstanceConfiguration) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *ExtensionInstanceConfiguration) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *ExtensionInstanceConfiguration) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *ExtensionInstanceConfiguration) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetDnsSubdomainChangeId

`func (o *ExtensionInstanceConfiguration) GetDnsSubdomainChangeId() float32`

GetDnsSubdomainChangeId returns the DnsSubdomainChangeId field if non-nil, zero value otherwise.

### GetDnsSubdomainChangeIdOk

`func (o *ExtensionInstanceConfiguration) GetDnsSubdomainChangeIdOk() (*float32, bool)`

GetDnsSubdomainChangeIdOk returns a tuple with the DnsSubdomainChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainChangeId

`func (o *ExtensionInstanceConfiguration) SetDnsSubdomainChangeId(v float32)`

SetDnsSubdomainChangeId sets DnsSubdomainChangeId field to given value.

### HasDnsSubdomainChangeId

`func (o *ExtensionInstanceConfiguration) HasDnsSubdomainChangeId() bool`

HasDnsSubdomainChangeId returns a boolean if a field has been set.

### GetDeployType

`func (o *ExtensionInstanceConfiguration) GetDeployType() string`

GetDeployType returns the DeployType field if non-nil, zero value otherwise.

### GetDeployTypeOk

`func (o *ExtensionInstanceConfiguration) GetDeployTypeOk() (*string, bool)`

GetDeployTypeOk returns a tuple with the DeployType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployType

`func (o *ExtensionInstanceConfiguration) SetDeployType(v string)`

SetDeployType sets DeployType field to given value.


### GetDeployStatus

`func (o *ExtensionInstanceConfiguration) GetDeployStatus() string`

GetDeployStatus returns the DeployStatus field if non-nil, zero value otherwise.

### GetDeployStatusOk

`func (o *ExtensionInstanceConfiguration) GetDeployStatusOk() (*string, bool)`

GetDeployStatusOk returns a tuple with the DeployStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStatus

`func (o *ExtensionInstanceConfiguration) SetDeployStatus(v string)`

SetDeployStatus sets DeployStatus field to given value.


### GetInfrastructureDeployId

`func (o *ExtensionInstanceConfiguration) GetInfrastructureDeployId() float32`

GetInfrastructureDeployId returns the InfrastructureDeployId field if non-nil, zero value otherwise.

### GetInfrastructureDeployIdOk

`func (o *ExtensionInstanceConfiguration) GetInfrastructureDeployIdOk() (*float32, bool)`

GetInfrastructureDeployIdOk returns a tuple with the InfrastructureDeployId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureDeployId

`func (o *ExtensionInstanceConfiguration) SetInfrastructureDeployId(v float32)`

SetInfrastructureDeployId sets InfrastructureDeployId field to given value.

### HasInfrastructureDeployId

`func (o *ExtensionInstanceConfiguration) HasInfrastructureDeployId() bool`

HasInfrastructureDeployId returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *ExtensionInstanceConfiguration) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *ExtensionInstanceConfiguration) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *ExtensionInstanceConfiguration) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


