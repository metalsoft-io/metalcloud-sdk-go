# EndpointInstanceGroupConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | **int32** | Revision number | 
**Label** | **string** | The Product Instance label. Will be automatically generated if not provided. | 
**UpdatedTimestamp** | **string** | Timestamp of the latest update of the Product Instance. | 
**Subdomain** | Pointer to **string** | Subdomain of the Product Instance. | [optional] 
**EndpointGroupName** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** | Custom hostname for the DNS Load Balancing record. If set, this will be used as the DNS Load Balancing record name instead of the default \&quot;endpoint-instance-group\&quot;. The hostname must be a valid DNS subdomain and can only contain alphanumeric characters, hyphens, and underscores. This will only take effect if the property \&quot;dnsLoadBalancingRecord\&quot; is true. It will be automatically suffixed with the endpoint instance group ID (e.g., \&quot;-34\&quot;) to ensure the uniqueness of the resulting DNS name. | [optional] 
**DnsSubdomainChangeId** | Pointer to **int32** | Id of the DNS subdomain for the Product Instance | [optional] 
**InfrastructureDeployId** | Pointer to **int32** | Id of the deployment for the Product Instance | [optional] 
**EmptyEdit** | Pointer to **int32** | Number of empty edits | [optional] 
**DeployType** | **string** | Product Instance deploy type | [default to "create"]
**DeployStatus** | **string** | Product Instance deploy status | [default to "not_started"]

## Methods

### NewEndpointInstanceGroupConfiguration

`func NewEndpointInstanceGroupConfiguration(revision int32, label string, updatedTimestamp string, deployType string, deployStatus string, ) *EndpointInstanceGroupConfiguration`

NewEndpointInstanceGroupConfiguration instantiates a new EndpointInstanceGroupConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointInstanceGroupConfigurationWithDefaults

`func NewEndpointInstanceGroupConfigurationWithDefaults() *EndpointInstanceGroupConfiguration`

NewEndpointInstanceGroupConfigurationWithDefaults instantiates a new EndpointInstanceGroupConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *EndpointInstanceGroupConfiguration) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *EndpointInstanceGroupConfiguration) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *EndpointInstanceGroupConfiguration) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetLabel

`func (o *EndpointInstanceGroupConfiguration) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *EndpointInstanceGroupConfiguration) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *EndpointInstanceGroupConfiguration) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetUpdatedTimestamp

`func (o *EndpointInstanceGroupConfiguration) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *EndpointInstanceGroupConfiguration) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *EndpointInstanceGroupConfiguration) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetSubdomain

`func (o *EndpointInstanceGroupConfiguration) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *EndpointInstanceGroupConfiguration) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *EndpointInstanceGroupConfiguration) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *EndpointInstanceGroupConfiguration) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetEndpointGroupName

`func (o *EndpointInstanceGroupConfiguration) GetEndpointGroupName() string`

GetEndpointGroupName returns the EndpointGroupName field if non-nil, zero value otherwise.

### GetEndpointGroupNameOk

`func (o *EndpointInstanceGroupConfiguration) GetEndpointGroupNameOk() (*string, bool)`

GetEndpointGroupNameOk returns a tuple with the EndpointGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointGroupName

`func (o *EndpointInstanceGroupConfiguration) SetEndpointGroupName(v string)`

SetEndpointGroupName sets EndpointGroupName field to given value.

### HasEndpointGroupName

`func (o *EndpointInstanceGroupConfiguration) HasEndpointGroupName() bool`

HasEndpointGroupName returns a boolean if a field has been set.

### GetHostname

`func (o *EndpointInstanceGroupConfiguration) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *EndpointInstanceGroupConfiguration) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *EndpointInstanceGroupConfiguration) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *EndpointInstanceGroupConfiguration) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetDnsSubdomainChangeId

`func (o *EndpointInstanceGroupConfiguration) GetDnsSubdomainChangeId() int32`

GetDnsSubdomainChangeId returns the DnsSubdomainChangeId field if non-nil, zero value otherwise.

### GetDnsSubdomainChangeIdOk

`func (o *EndpointInstanceGroupConfiguration) GetDnsSubdomainChangeIdOk() (*int32, bool)`

GetDnsSubdomainChangeIdOk returns a tuple with the DnsSubdomainChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainChangeId

`func (o *EndpointInstanceGroupConfiguration) SetDnsSubdomainChangeId(v int32)`

SetDnsSubdomainChangeId sets DnsSubdomainChangeId field to given value.

### HasDnsSubdomainChangeId

`func (o *EndpointInstanceGroupConfiguration) HasDnsSubdomainChangeId() bool`

HasDnsSubdomainChangeId returns a boolean if a field has been set.

### GetInfrastructureDeployId

`func (o *EndpointInstanceGroupConfiguration) GetInfrastructureDeployId() int32`

GetInfrastructureDeployId returns the InfrastructureDeployId field if non-nil, zero value otherwise.

### GetInfrastructureDeployIdOk

`func (o *EndpointInstanceGroupConfiguration) GetInfrastructureDeployIdOk() (*int32, bool)`

GetInfrastructureDeployIdOk returns a tuple with the InfrastructureDeployId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureDeployId

`func (o *EndpointInstanceGroupConfiguration) SetInfrastructureDeployId(v int32)`

SetInfrastructureDeployId sets InfrastructureDeployId field to given value.

### HasInfrastructureDeployId

`func (o *EndpointInstanceGroupConfiguration) HasInfrastructureDeployId() bool`

HasInfrastructureDeployId returns a boolean if a field has been set.

### GetEmptyEdit

`func (o *EndpointInstanceGroupConfiguration) GetEmptyEdit() int32`

GetEmptyEdit returns the EmptyEdit field if non-nil, zero value otherwise.

### GetEmptyEditOk

`func (o *EndpointInstanceGroupConfiguration) GetEmptyEditOk() (*int32, bool)`

GetEmptyEditOk returns a tuple with the EmptyEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyEdit

`func (o *EndpointInstanceGroupConfiguration) SetEmptyEdit(v int32)`

SetEmptyEdit sets EmptyEdit field to given value.

### HasEmptyEdit

`func (o *EndpointInstanceGroupConfiguration) HasEmptyEdit() bool`

HasEmptyEdit returns a boolean if a field has been set.

### GetDeployType

`func (o *EndpointInstanceGroupConfiguration) GetDeployType() string`

GetDeployType returns the DeployType field if non-nil, zero value otherwise.

### GetDeployTypeOk

`func (o *EndpointInstanceGroupConfiguration) GetDeployTypeOk() (*string, bool)`

GetDeployTypeOk returns a tuple with the DeployType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployType

`func (o *EndpointInstanceGroupConfiguration) SetDeployType(v string)`

SetDeployType sets DeployType field to given value.


### GetDeployStatus

`func (o *EndpointInstanceGroupConfiguration) GetDeployStatus() string`

GetDeployStatus returns the DeployStatus field if non-nil, zero value otherwise.

### GetDeployStatusOk

`func (o *EndpointInstanceGroupConfiguration) GetDeployStatusOk() (*string, bool)`

GetDeployStatusOk returns a tuple with the DeployStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStatus

`func (o *EndpointInstanceGroupConfiguration) SetDeployStatus(v string)`

SetDeployStatus sets DeployStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


