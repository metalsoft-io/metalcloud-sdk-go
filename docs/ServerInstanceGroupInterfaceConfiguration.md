# ServerInstanceGroupInterfaceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | **int32** | Revision number | 
**Label** | **string** | The Product Instance label. Will be automatically generated if not provided. | 
**UpdatedTimestamp** | **string** | Timestamp of the latest update of the Product Instance. | 
**Subdomain** | Pointer to **string** | Subdomain of the Product Instance. | [optional] 
**GroupId** | **int32** |  | 
**Index** | **int32** | The index of the interface (0-based) on this server. | 
**NetworkId** | Pointer to **int32** | The ID of the network to which this interface is to be attached to. | [optional] 
**DnsSubdomainChangeId** | Pointer to **int32** | Id of the DNS subdomain for the Product Instance | [optional] 
**InfrastructureDeployId** | Pointer to **int32** | Id of the deployment for the Product Instance | [optional] 
**EmptyEdit** | Pointer to **int32** | Number of empty edits | [optional] 
**DeployType** | **string** | Product Instance deploy type | [default to "create"]
**DeployStatus** | **string** | Product Instance deploy status | [default to "not_started"]

## Methods

### NewServerInstanceGroupInterfaceConfiguration

`func NewServerInstanceGroupInterfaceConfiguration(revision int32, label string, updatedTimestamp string, groupId int32, index int32, deployType string, deployStatus string, ) *ServerInstanceGroupInterfaceConfiguration`

NewServerInstanceGroupInterfaceConfiguration instantiates a new ServerInstanceGroupInterfaceConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceGroupInterfaceConfigurationWithDefaults

`func NewServerInstanceGroupInterfaceConfigurationWithDefaults() *ServerInstanceGroupInterfaceConfiguration`

NewServerInstanceGroupInterfaceConfigurationWithDefaults instantiates a new ServerInstanceGroupInterfaceConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *ServerInstanceGroupInterfaceConfiguration) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ServerInstanceGroupInterfaceConfiguration) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ServerInstanceGroupInterfaceConfiguration) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetLabel

`func (o *ServerInstanceGroupInterfaceConfiguration) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServerInstanceGroupInterfaceConfiguration) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServerInstanceGroupInterfaceConfiguration) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetUpdatedTimestamp

`func (o *ServerInstanceGroupInterfaceConfiguration) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *ServerInstanceGroupInterfaceConfiguration) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *ServerInstanceGroupInterfaceConfiguration) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetSubdomain

`func (o *ServerInstanceGroupInterfaceConfiguration) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *ServerInstanceGroupInterfaceConfiguration) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *ServerInstanceGroupInterfaceConfiguration) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *ServerInstanceGroupInterfaceConfiguration) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetGroupId

`func (o *ServerInstanceGroupInterfaceConfiguration) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ServerInstanceGroupInterfaceConfiguration) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ServerInstanceGroupInterfaceConfiguration) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.


### GetIndex

`func (o *ServerInstanceGroupInterfaceConfiguration) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ServerInstanceGroupInterfaceConfiguration) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ServerInstanceGroupInterfaceConfiguration) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetNetworkId

`func (o *ServerInstanceGroupInterfaceConfiguration) GetNetworkId() int32`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *ServerInstanceGroupInterfaceConfiguration) GetNetworkIdOk() (*int32, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *ServerInstanceGroupInterfaceConfiguration) SetNetworkId(v int32)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *ServerInstanceGroupInterfaceConfiguration) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetDnsSubdomainChangeId

`func (o *ServerInstanceGroupInterfaceConfiguration) GetDnsSubdomainChangeId() int32`

GetDnsSubdomainChangeId returns the DnsSubdomainChangeId field if non-nil, zero value otherwise.

### GetDnsSubdomainChangeIdOk

`func (o *ServerInstanceGroupInterfaceConfiguration) GetDnsSubdomainChangeIdOk() (*int32, bool)`

GetDnsSubdomainChangeIdOk returns a tuple with the DnsSubdomainChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainChangeId

`func (o *ServerInstanceGroupInterfaceConfiguration) SetDnsSubdomainChangeId(v int32)`

SetDnsSubdomainChangeId sets DnsSubdomainChangeId field to given value.

### HasDnsSubdomainChangeId

`func (o *ServerInstanceGroupInterfaceConfiguration) HasDnsSubdomainChangeId() bool`

HasDnsSubdomainChangeId returns a boolean if a field has been set.

### GetInfrastructureDeployId

`func (o *ServerInstanceGroupInterfaceConfiguration) GetInfrastructureDeployId() int32`

GetInfrastructureDeployId returns the InfrastructureDeployId field if non-nil, zero value otherwise.

### GetInfrastructureDeployIdOk

`func (o *ServerInstanceGroupInterfaceConfiguration) GetInfrastructureDeployIdOk() (*int32, bool)`

GetInfrastructureDeployIdOk returns a tuple with the InfrastructureDeployId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureDeployId

`func (o *ServerInstanceGroupInterfaceConfiguration) SetInfrastructureDeployId(v int32)`

SetInfrastructureDeployId sets InfrastructureDeployId field to given value.

### HasInfrastructureDeployId

`func (o *ServerInstanceGroupInterfaceConfiguration) HasInfrastructureDeployId() bool`

HasInfrastructureDeployId returns a boolean if a field has been set.

### GetEmptyEdit

`func (o *ServerInstanceGroupInterfaceConfiguration) GetEmptyEdit() int32`

GetEmptyEdit returns the EmptyEdit field if non-nil, zero value otherwise.

### GetEmptyEditOk

`func (o *ServerInstanceGroupInterfaceConfiguration) GetEmptyEditOk() (*int32, bool)`

GetEmptyEditOk returns a tuple with the EmptyEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyEdit

`func (o *ServerInstanceGroupInterfaceConfiguration) SetEmptyEdit(v int32)`

SetEmptyEdit sets EmptyEdit field to given value.

### HasEmptyEdit

`func (o *ServerInstanceGroupInterfaceConfiguration) HasEmptyEdit() bool`

HasEmptyEdit returns a boolean if a field has been set.

### GetDeployType

`func (o *ServerInstanceGroupInterfaceConfiguration) GetDeployType() string`

GetDeployType returns the DeployType field if non-nil, zero value otherwise.

### GetDeployTypeOk

`func (o *ServerInstanceGroupInterfaceConfiguration) GetDeployTypeOk() (*string, bool)`

GetDeployTypeOk returns a tuple with the DeployType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployType

`func (o *ServerInstanceGroupInterfaceConfiguration) SetDeployType(v string)`

SetDeployType sets DeployType field to given value.


### GetDeployStatus

`func (o *ServerInstanceGroupInterfaceConfiguration) GetDeployStatus() string`

GetDeployStatus returns the DeployStatus field if non-nil, zero value otherwise.

### GetDeployStatusOk

`func (o *ServerInstanceGroupInterfaceConfiguration) GetDeployStatusOk() (*string, bool)`

GetDeployStatusOk returns a tuple with the DeployStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStatus

`func (o *ServerInstanceGroupInterfaceConfiguration) SetDeployStatus(v string)`

SetDeployStatus sets DeployStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


