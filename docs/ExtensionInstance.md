# ExtensionInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | **float32** | Revision number | 
**Label** | **string** | The extension instance label. Will be automatically generated if not provided. | 
**AutomaticManagement** | **float32** | Flag specifying if the extension instance supports automatic management. | 
**Subdomain** | Pointer to **string** | Subdomain of the Extension Instance. | [optional] 
**UpdatedTimestamp** | **string** | Timestamp of the Extension Instance last update. | 
**Id** | **float32** | The extension instance ID. | 
**InfrastructureId** | **float32** | The infrastructure ID. | 
**Infrastructure** | [**ParentInfrastructureDto**](ParentInfrastructureDto.md) | Infrastructure information | 
**ExtensionId** | **float32** | The extension ID. | 
**ServiceStatus** | **string** | Service status of the Extension Instance | 
**SubdomainPermanent** | Pointer to **string** | Subdomain permanent of the Extension Instance. | [optional] 
**DnsSubdomainId** | Pointer to **float32** | Id of the DNS subdomain for the Extension Instance. | [optional] 
**DnsSubdomainPermanentId** | Pointer to **float32** | Id of the permanent DNS subdomain for the Extension Instance. | [optional] 
**InputVariables** | [**[]ExtensionVariable**](ExtensionVariable.md) | Input variables values. | 
**OutputVariables** | [**[]ExtensionVariable**](ExtensionVariable.md) | Output variables values. | 
**Config** | [**ExtensionInstanceConfiguration**](ExtensionInstanceConfiguration.md) | The current changes to be deployed for the Extension Instance. | 
**CreatedTimestamp** | **string** | Timestamp of the Extension Instance creation. | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewExtensionInstance

`func NewExtensionInstance(revision float32, label string, automaticManagement float32, updatedTimestamp string, id float32, infrastructureId float32, infrastructure ParentInfrastructureDto, extensionId float32, serviceStatus string, inputVariables []ExtensionVariable, outputVariables []ExtensionVariable, config ExtensionInstanceConfiguration, createdTimestamp string, ) *ExtensionInstance`

NewExtensionInstance instantiates a new ExtensionInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionInstanceWithDefaults

`func NewExtensionInstanceWithDefaults() *ExtensionInstance`

NewExtensionInstanceWithDefaults instantiates a new ExtensionInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *ExtensionInstance) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ExtensionInstance) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ExtensionInstance) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetLabel

`func (o *ExtensionInstance) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ExtensionInstance) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ExtensionInstance) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetAutomaticManagement

`func (o *ExtensionInstance) GetAutomaticManagement() float32`

GetAutomaticManagement returns the AutomaticManagement field if non-nil, zero value otherwise.

### GetAutomaticManagementOk

`func (o *ExtensionInstance) GetAutomaticManagementOk() (*float32, bool)`

GetAutomaticManagementOk returns a tuple with the AutomaticManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticManagement

`func (o *ExtensionInstance) SetAutomaticManagement(v float32)`

SetAutomaticManagement sets AutomaticManagement field to given value.


### GetSubdomain

`func (o *ExtensionInstance) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *ExtensionInstance) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *ExtensionInstance) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *ExtensionInstance) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *ExtensionInstance) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *ExtensionInstance) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *ExtensionInstance) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetId

`func (o *ExtensionInstance) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExtensionInstance) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExtensionInstance) SetId(v float32)`

SetId sets Id field to given value.


### GetInfrastructureId

`func (o *ExtensionInstance) GetInfrastructureId() float32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *ExtensionInstance) GetInfrastructureIdOk() (*float32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *ExtensionInstance) SetInfrastructureId(v float32)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetInfrastructure

`func (o *ExtensionInstance) GetInfrastructure() ParentInfrastructureDto`

GetInfrastructure returns the Infrastructure field if non-nil, zero value otherwise.

### GetInfrastructureOk

`func (o *ExtensionInstance) GetInfrastructureOk() (*ParentInfrastructureDto, bool)`

GetInfrastructureOk returns a tuple with the Infrastructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructure

`func (o *ExtensionInstance) SetInfrastructure(v ParentInfrastructureDto)`

SetInfrastructure sets Infrastructure field to given value.


### GetExtensionId

`func (o *ExtensionInstance) GetExtensionId() float32`

GetExtensionId returns the ExtensionId field if non-nil, zero value otherwise.

### GetExtensionIdOk

`func (o *ExtensionInstance) GetExtensionIdOk() (*float32, bool)`

GetExtensionIdOk returns a tuple with the ExtensionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionId

`func (o *ExtensionInstance) SetExtensionId(v float32)`

SetExtensionId sets ExtensionId field to given value.


### GetServiceStatus

`func (o *ExtensionInstance) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *ExtensionInstance) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *ExtensionInstance) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetSubdomainPermanent

`func (o *ExtensionInstance) GetSubdomainPermanent() string`

GetSubdomainPermanent returns the SubdomainPermanent field if non-nil, zero value otherwise.

### GetSubdomainPermanentOk

`func (o *ExtensionInstance) GetSubdomainPermanentOk() (*string, bool)`

GetSubdomainPermanentOk returns a tuple with the SubdomainPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomainPermanent

`func (o *ExtensionInstance) SetSubdomainPermanent(v string)`

SetSubdomainPermanent sets SubdomainPermanent field to given value.

### HasSubdomainPermanent

`func (o *ExtensionInstance) HasSubdomainPermanent() bool`

HasSubdomainPermanent returns a boolean if a field has been set.

### GetDnsSubdomainId

`func (o *ExtensionInstance) GetDnsSubdomainId() float32`

GetDnsSubdomainId returns the DnsSubdomainId field if non-nil, zero value otherwise.

### GetDnsSubdomainIdOk

`func (o *ExtensionInstance) GetDnsSubdomainIdOk() (*float32, bool)`

GetDnsSubdomainIdOk returns a tuple with the DnsSubdomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainId

`func (o *ExtensionInstance) SetDnsSubdomainId(v float32)`

SetDnsSubdomainId sets DnsSubdomainId field to given value.

### HasDnsSubdomainId

`func (o *ExtensionInstance) HasDnsSubdomainId() bool`

HasDnsSubdomainId returns a boolean if a field has been set.

### GetDnsSubdomainPermanentId

`func (o *ExtensionInstance) GetDnsSubdomainPermanentId() float32`

GetDnsSubdomainPermanentId returns the DnsSubdomainPermanentId field if non-nil, zero value otherwise.

### GetDnsSubdomainPermanentIdOk

`func (o *ExtensionInstance) GetDnsSubdomainPermanentIdOk() (*float32, bool)`

GetDnsSubdomainPermanentIdOk returns a tuple with the DnsSubdomainPermanentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainPermanentId

`func (o *ExtensionInstance) SetDnsSubdomainPermanentId(v float32)`

SetDnsSubdomainPermanentId sets DnsSubdomainPermanentId field to given value.

### HasDnsSubdomainPermanentId

`func (o *ExtensionInstance) HasDnsSubdomainPermanentId() bool`

HasDnsSubdomainPermanentId returns a boolean if a field has been set.

### GetInputVariables

`func (o *ExtensionInstance) GetInputVariables() []ExtensionVariable`

GetInputVariables returns the InputVariables field if non-nil, zero value otherwise.

### GetInputVariablesOk

`func (o *ExtensionInstance) GetInputVariablesOk() (*[]ExtensionVariable, bool)`

GetInputVariablesOk returns a tuple with the InputVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputVariables

`func (o *ExtensionInstance) SetInputVariables(v []ExtensionVariable)`

SetInputVariables sets InputVariables field to given value.


### GetOutputVariables

`func (o *ExtensionInstance) GetOutputVariables() []ExtensionVariable`

GetOutputVariables returns the OutputVariables field if non-nil, zero value otherwise.

### GetOutputVariablesOk

`func (o *ExtensionInstance) GetOutputVariablesOk() (*[]ExtensionVariable, bool)`

GetOutputVariablesOk returns a tuple with the OutputVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputVariables

`func (o *ExtensionInstance) SetOutputVariables(v []ExtensionVariable)`

SetOutputVariables sets OutputVariables field to given value.


### GetConfig

`func (o *ExtensionInstance) GetConfig() ExtensionInstanceConfiguration`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ExtensionInstance) GetConfigOk() (*ExtensionInstanceConfiguration, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ExtensionInstance) SetConfig(v ExtensionInstanceConfiguration)`

SetConfig sets Config field to given value.


### GetCreatedTimestamp

`func (o *ExtensionInstance) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *ExtensionInstance) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *ExtensionInstance) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetLinks

`func (o *ExtensionInstance) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExtensionInstance) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExtensionInstance) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExtensionInstance) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


