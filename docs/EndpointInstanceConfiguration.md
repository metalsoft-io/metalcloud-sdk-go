# EndpointInstanceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | **int32** | Revision number | 
**Label** | **string** | The Product Instance label. Will be automatically generated if not provided. | 
**UpdatedTimestamp** | **string** | Timestamp of the latest update of the Product Instance. | 
**Subdomain** | Pointer to **string** | Subdomain of the Product Instance. | [optional] 
**GroupId** | **int32** |  | 
**EndpointId** | Pointer to **int32** | The ID of the endpoint assigned to the instance. | [optional] 
**Hostname** | Pointer to **string** | The subdomain of the server instance. | [optional] 
**DnsSubdomainChangeId** | Pointer to **int32** | Id of the DNS subdomain for the Product Instance | [optional] 
**InfrastructureDeployId** | Pointer to **int32** | Id of the deployment for the Product Instance | [optional] 
**EmptyEdit** | Pointer to **int32** | Number of empty edits | [optional] 
**DeployType** | **string** | Product Instance deploy type | [default to "create"]
**DeployStatus** | **string** | Product Instance deploy status | [default to "not_started"]

## Methods

### NewEndpointInstanceConfiguration

`func NewEndpointInstanceConfiguration(revision int32, label string, updatedTimestamp string, groupId int32, deployType string, deployStatus string, ) *EndpointInstanceConfiguration`

NewEndpointInstanceConfiguration instantiates a new EndpointInstanceConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointInstanceConfigurationWithDefaults

`func NewEndpointInstanceConfigurationWithDefaults() *EndpointInstanceConfiguration`

NewEndpointInstanceConfigurationWithDefaults instantiates a new EndpointInstanceConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *EndpointInstanceConfiguration) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *EndpointInstanceConfiguration) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *EndpointInstanceConfiguration) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetLabel

`func (o *EndpointInstanceConfiguration) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *EndpointInstanceConfiguration) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *EndpointInstanceConfiguration) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetUpdatedTimestamp

`func (o *EndpointInstanceConfiguration) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *EndpointInstanceConfiguration) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *EndpointInstanceConfiguration) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetSubdomain

`func (o *EndpointInstanceConfiguration) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *EndpointInstanceConfiguration) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *EndpointInstanceConfiguration) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *EndpointInstanceConfiguration) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetGroupId

`func (o *EndpointInstanceConfiguration) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *EndpointInstanceConfiguration) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *EndpointInstanceConfiguration) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.


### GetEndpointId

`func (o *EndpointInstanceConfiguration) GetEndpointId() int32`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *EndpointInstanceConfiguration) GetEndpointIdOk() (*int32, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *EndpointInstanceConfiguration) SetEndpointId(v int32)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *EndpointInstanceConfiguration) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetHostname

`func (o *EndpointInstanceConfiguration) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *EndpointInstanceConfiguration) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *EndpointInstanceConfiguration) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *EndpointInstanceConfiguration) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetDnsSubdomainChangeId

`func (o *EndpointInstanceConfiguration) GetDnsSubdomainChangeId() int32`

GetDnsSubdomainChangeId returns the DnsSubdomainChangeId field if non-nil, zero value otherwise.

### GetDnsSubdomainChangeIdOk

`func (o *EndpointInstanceConfiguration) GetDnsSubdomainChangeIdOk() (*int32, bool)`

GetDnsSubdomainChangeIdOk returns a tuple with the DnsSubdomainChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainChangeId

`func (o *EndpointInstanceConfiguration) SetDnsSubdomainChangeId(v int32)`

SetDnsSubdomainChangeId sets DnsSubdomainChangeId field to given value.

### HasDnsSubdomainChangeId

`func (o *EndpointInstanceConfiguration) HasDnsSubdomainChangeId() bool`

HasDnsSubdomainChangeId returns a boolean if a field has been set.

### GetInfrastructureDeployId

`func (o *EndpointInstanceConfiguration) GetInfrastructureDeployId() int32`

GetInfrastructureDeployId returns the InfrastructureDeployId field if non-nil, zero value otherwise.

### GetInfrastructureDeployIdOk

`func (o *EndpointInstanceConfiguration) GetInfrastructureDeployIdOk() (*int32, bool)`

GetInfrastructureDeployIdOk returns a tuple with the InfrastructureDeployId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureDeployId

`func (o *EndpointInstanceConfiguration) SetInfrastructureDeployId(v int32)`

SetInfrastructureDeployId sets InfrastructureDeployId field to given value.

### HasInfrastructureDeployId

`func (o *EndpointInstanceConfiguration) HasInfrastructureDeployId() bool`

HasInfrastructureDeployId returns a boolean if a field has been set.

### GetEmptyEdit

`func (o *EndpointInstanceConfiguration) GetEmptyEdit() int32`

GetEmptyEdit returns the EmptyEdit field if non-nil, zero value otherwise.

### GetEmptyEditOk

`func (o *EndpointInstanceConfiguration) GetEmptyEditOk() (*int32, bool)`

GetEmptyEditOk returns a tuple with the EmptyEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyEdit

`func (o *EndpointInstanceConfiguration) SetEmptyEdit(v int32)`

SetEmptyEdit sets EmptyEdit field to given value.

### HasEmptyEdit

`func (o *EndpointInstanceConfiguration) HasEmptyEdit() bool`

HasEmptyEdit returns a boolean if a field has been set.

### GetDeployType

`func (o *EndpointInstanceConfiguration) GetDeployType() string`

GetDeployType returns the DeployType field if non-nil, zero value otherwise.

### GetDeployTypeOk

`func (o *EndpointInstanceConfiguration) GetDeployTypeOk() (*string, bool)`

GetDeployTypeOk returns a tuple with the DeployType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployType

`func (o *EndpointInstanceConfiguration) SetDeployType(v string)`

SetDeployType sets DeployType field to given value.


### GetDeployStatus

`func (o *EndpointInstanceConfiguration) GetDeployStatus() string`

GetDeployStatus returns the DeployStatus field if non-nil, zero value otherwise.

### GetDeployStatusOk

`func (o *EndpointInstanceConfiguration) GetDeployStatusOk() (*string, bool)`

GetDeployStatusOk returns a tuple with the DeployStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStatus

`func (o *EndpointInstanceConfiguration) SetDeployStatus(v string)`

SetDeployStatus sets DeployStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


