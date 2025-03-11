# ServerInstanceProfileNetworkInterfacesConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | **float32** | Revision number | 
**Label** | **string** | The server instance interface label. | 
**Subdomain** | Pointer to **string** | Subdomain of the Server Group. | [optional] 
**UpdatedTimestamp** | **string** | Timestamp of the latest update for the Server Instance. | 
**InstanceId** | **int32** |  | 
**Index** | **int32** | The index of the interface (0-based) on this server. | 
**CapacityMbps** | **int32** |  | 
**NetworkId** | Pointer to **int32** | The ID of the network to which this interface is to be attached to. | [optional] 
**ServerInterfaceId** | Pointer to **int32** |  | [optional] 
**EmptyEdit** | Pointer to **int32** | Number of empty edits | [optional] 
**DeployType** | **string** | Server Instance Interface deploy type | 
**DeployStatus** | **string** | Server Instance Interface deploy status | 
**DnsSubdomainChangeId** | Pointer to **int32** | Id of the DNS subdomain for the Server Instance Interface. | [optional] 
**InfrastructureDeployId** | Pointer to **int32** | Id of the deployment for the Server Instance Interface. | [optional] 

## Methods

### NewServerInstanceProfileNetworkInterfacesConfig

`func NewServerInstanceProfileNetworkInterfacesConfig(revision float32, label string, updatedTimestamp string, instanceId int32, index int32, capacityMbps int32, deployType string, deployStatus string, ) *ServerInstanceProfileNetworkInterfacesConfig`

NewServerInstanceProfileNetworkInterfacesConfig instantiates a new ServerInstanceProfileNetworkInterfacesConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceProfileNetworkInterfacesConfigWithDefaults

`func NewServerInstanceProfileNetworkInterfacesConfigWithDefaults() *ServerInstanceProfileNetworkInterfacesConfig`

NewServerInstanceProfileNetworkInterfacesConfigWithDefaults instantiates a new ServerInstanceProfileNetworkInterfacesConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *ServerInstanceProfileNetworkInterfacesConfig) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ServerInstanceProfileNetworkInterfacesConfig) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ServerInstanceProfileNetworkInterfacesConfig) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetLabel

`func (o *ServerInstanceProfileNetworkInterfacesConfig) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServerInstanceProfileNetworkInterfacesConfig) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServerInstanceProfileNetworkInterfacesConfig) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetSubdomain

`func (o *ServerInstanceProfileNetworkInterfacesConfig) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *ServerInstanceProfileNetworkInterfacesConfig) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *ServerInstanceProfileNetworkInterfacesConfig) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *ServerInstanceProfileNetworkInterfacesConfig) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *ServerInstanceProfileNetworkInterfacesConfig) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *ServerInstanceProfileNetworkInterfacesConfig) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *ServerInstanceProfileNetworkInterfacesConfig) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetInstanceId

`func (o *ServerInstanceProfileNetworkInterfacesConfig) GetInstanceId() int32`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ServerInstanceProfileNetworkInterfacesConfig) GetInstanceIdOk() (*int32, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ServerInstanceProfileNetworkInterfacesConfig) SetInstanceId(v int32)`

SetInstanceId sets InstanceId field to given value.


### GetIndex

`func (o *ServerInstanceProfileNetworkInterfacesConfig) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ServerInstanceProfileNetworkInterfacesConfig) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ServerInstanceProfileNetworkInterfacesConfig) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetCapacityMbps

`func (o *ServerInstanceProfileNetworkInterfacesConfig) GetCapacityMbps() int32`

GetCapacityMbps returns the CapacityMbps field if non-nil, zero value otherwise.

### GetCapacityMbpsOk

`func (o *ServerInstanceProfileNetworkInterfacesConfig) GetCapacityMbpsOk() (*int32, bool)`

GetCapacityMbpsOk returns a tuple with the CapacityMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityMbps

`func (o *ServerInstanceProfileNetworkInterfacesConfig) SetCapacityMbps(v int32)`

SetCapacityMbps sets CapacityMbps field to given value.


### GetNetworkId

`func (o *ServerInstanceProfileNetworkInterfacesConfig) GetNetworkId() int32`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *ServerInstanceProfileNetworkInterfacesConfig) GetNetworkIdOk() (*int32, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *ServerInstanceProfileNetworkInterfacesConfig) SetNetworkId(v int32)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *ServerInstanceProfileNetworkInterfacesConfig) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetServerInterfaceId

`func (o *ServerInstanceProfileNetworkInterfacesConfig) GetServerInterfaceId() int32`

GetServerInterfaceId returns the ServerInterfaceId field if non-nil, zero value otherwise.

### GetServerInterfaceIdOk

`func (o *ServerInstanceProfileNetworkInterfacesConfig) GetServerInterfaceIdOk() (*int32, bool)`

GetServerInterfaceIdOk returns a tuple with the ServerInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInterfaceId

`func (o *ServerInstanceProfileNetworkInterfacesConfig) SetServerInterfaceId(v int32)`

SetServerInterfaceId sets ServerInterfaceId field to given value.

### HasServerInterfaceId

`func (o *ServerInstanceProfileNetworkInterfacesConfig) HasServerInterfaceId() bool`

HasServerInterfaceId returns a boolean if a field has been set.

### GetEmptyEdit

`func (o *ServerInstanceProfileNetworkInterfacesConfig) GetEmptyEdit() int32`

GetEmptyEdit returns the EmptyEdit field if non-nil, zero value otherwise.

### GetEmptyEditOk

`func (o *ServerInstanceProfileNetworkInterfacesConfig) GetEmptyEditOk() (*int32, bool)`

GetEmptyEditOk returns a tuple with the EmptyEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyEdit

`func (o *ServerInstanceProfileNetworkInterfacesConfig) SetEmptyEdit(v int32)`

SetEmptyEdit sets EmptyEdit field to given value.

### HasEmptyEdit

`func (o *ServerInstanceProfileNetworkInterfacesConfig) HasEmptyEdit() bool`

HasEmptyEdit returns a boolean if a field has been set.

### GetDeployType

`func (o *ServerInstanceProfileNetworkInterfacesConfig) GetDeployType() string`

GetDeployType returns the DeployType field if non-nil, zero value otherwise.

### GetDeployTypeOk

`func (o *ServerInstanceProfileNetworkInterfacesConfig) GetDeployTypeOk() (*string, bool)`

GetDeployTypeOk returns a tuple with the DeployType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployType

`func (o *ServerInstanceProfileNetworkInterfacesConfig) SetDeployType(v string)`

SetDeployType sets DeployType field to given value.


### GetDeployStatus

`func (o *ServerInstanceProfileNetworkInterfacesConfig) GetDeployStatus() string`

GetDeployStatus returns the DeployStatus field if non-nil, zero value otherwise.

### GetDeployStatusOk

`func (o *ServerInstanceProfileNetworkInterfacesConfig) GetDeployStatusOk() (*string, bool)`

GetDeployStatusOk returns a tuple with the DeployStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStatus

`func (o *ServerInstanceProfileNetworkInterfacesConfig) SetDeployStatus(v string)`

SetDeployStatus sets DeployStatus field to given value.


### GetDnsSubdomainChangeId

`func (o *ServerInstanceProfileNetworkInterfacesConfig) GetDnsSubdomainChangeId() int32`

GetDnsSubdomainChangeId returns the DnsSubdomainChangeId field if non-nil, zero value otherwise.

### GetDnsSubdomainChangeIdOk

`func (o *ServerInstanceProfileNetworkInterfacesConfig) GetDnsSubdomainChangeIdOk() (*int32, bool)`

GetDnsSubdomainChangeIdOk returns a tuple with the DnsSubdomainChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainChangeId

`func (o *ServerInstanceProfileNetworkInterfacesConfig) SetDnsSubdomainChangeId(v int32)`

SetDnsSubdomainChangeId sets DnsSubdomainChangeId field to given value.

### HasDnsSubdomainChangeId

`func (o *ServerInstanceProfileNetworkInterfacesConfig) HasDnsSubdomainChangeId() bool`

HasDnsSubdomainChangeId returns a boolean if a field has been set.

### GetInfrastructureDeployId

`func (o *ServerInstanceProfileNetworkInterfacesConfig) GetInfrastructureDeployId() int32`

GetInfrastructureDeployId returns the InfrastructureDeployId field if non-nil, zero value otherwise.

### GetInfrastructureDeployIdOk

`func (o *ServerInstanceProfileNetworkInterfacesConfig) GetInfrastructureDeployIdOk() (*int32, bool)`

GetInfrastructureDeployIdOk returns a tuple with the InfrastructureDeployId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureDeployId

`func (o *ServerInstanceProfileNetworkInterfacesConfig) SetInfrastructureDeployId(v int32)`

SetInfrastructureDeployId sets InfrastructureDeployId field to given value.

### HasInfrastructureDeployId

`func (o *ServerInstanceProfileNetworkInterfacesConfig) HasInfrastructureDeployId() bool`

HasInfrastructureDeployId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


