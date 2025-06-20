# ServerInstanceInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The Product Instance ID. | 
**Revision** | **int32** | Revision number | 
**Label** | **string** | The Product Instance label. Will be automatically generated if not provided. | 
**CreatedTimestamp** | **string** | Timestamp of the Product Instance creation. | 
**UpdatedTimestamp** | **string** | Timestamp of the latest update of the Product Instance. | 
**Subdomain** | Pointer to **string** | Subdomain of the Product Instance. | [optional] 
**SubdomainPermanent** | Pointer to **string** | Subdomain permanent of the Product Instance. | [optional] 
**DnsSubdomainId** | Pointer to **int32** | Id of the DNS subdomain for the Product Instance | [optional] 
**DnsSubdomainPermanentId** | Pointer to **int32** | Id of the permanent DNS subdomain for the Product Instance | [optional] 
**InfrastructureId** | **int32** |  | 
**InstanceId** | **int32** |  | 
**Index** | **int32** | The index of the interface (0-based) on this server. | 
**CapacityMbps** | **int32** |  | 
**NetworkId** | Pointer to **int32** | The ID of the network to which this interface is to be attached to. | [optional] 
**SubnetPoolSanIndex** | Pointer to **int32** |  | [optional] 
**VlanIdentifier** | Pointer to **int32** |  | [optional] 
**AclIdentifier** | Pointer to **int32** |  | [optional] 
**NetworkEquipmentSubnetPoolSanId** | Pointer to **int32** |  | [optional] 
**ServerInterfaceId** | Pointer to **int32** |  | [optional] 
**DirtyBit** | **bool** |  | 
**ServiceStatus** | **string** | Current status of the server instance interface. | 
**Config** | Pointer to [**ServerInstanceInterfaceConfiguration**](ServerInstanceInterfaceConfiguration.md) |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewServerInstanceInterface

`func NewServerInstanceInterface(id int32, revision int32, label string, createdTimestamp string, updatedTimestamp string, infrastructureId int32, instanceId int32, index int32, capacityMbps int32, dirtyBit bool, serviceStatus string, ) *ServerInstanceInterface`

NewServerInstanceInterface instantiates a new ServerInstanceInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceInterfaceWithDefaults

`func NewServerInstanceInterfaceWithDefaults() *ServerInstanceInterface`

NewServerInstanceInterfaceWithDefaults instantiates a new ServerInstanceInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerInstanceInterface) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerInstanceInterface) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerInstanceInterface) SetId(v int32)`

SetId sets Id field to given value.


### GetRevision

`func (o *ServerInstanceInterface) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ServerInstanceInterface) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ServerInstanceInterface) SetRevision(v int32)`

SetRevision sets Revision field to given value.


### GetLabel

`func (o *ServerInstanceInterface) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServerInstanceInterface) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServerInstanceInterface) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetCreatedTimestamp

`func (o *ServerInstanceInterface) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *ServerInstanceInterface) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *ServerInstanceInterface) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *ServerInstanceInterface) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *ServerInstanceInterface) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *ServerInstanceInterface) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetSubdomain

`func (o *ServerInstanceInterface) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *ServerInstanceInterface) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *ServerInstanceInterface) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *ServerInstanceInterface) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetSubdomainPermanent

`func (o *ServerInstanceInterface) GetSubdomainPermanent() string`

GetSubdomainPermanent returns the SubdomainPermanent field if non-nil, zero value otherwise.

### GetSubdomainPermanentOk

`func (o *ServerInstanceInterface) GetSubdomainPermanentOk() (*string, bool)`

GetSubdomainPermanentOk returns a tuple with the SubdomainPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomainPermanent

`func (o *ServerInstanceInterface) SetSubdomainPermanent(v string)`

SetSubdomainPermanent sets SubdomainPermanent field to given value.

### HasSubdomainPermanent

`func (o *ServerInstanceInterface) HasSubdomainPermanent() bool`

HasSubdomainPermanent returns a boolean if a field has been set.

### GetDnsSubdomainId

`func (o *ServerInstanceInterface) GetDnsSubdomainId() int32`

GetDnsSubdomainId returns the DnsSubdomainId field if non-nil, zero value otherwise.

### GetDnsSubdomainIdOk

`func (o *ServerInstanceInterface) GetDnsSubdomainIdOk() (*int32, bool)`

GetDnsSubdomainIdOk returns a tuple with the DnsSubdomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainId

`func (o *ServerInstanceInterface) SetDnsSubdomainId(v int32)`

SetDnsSubdomainId sets DnsSubdomainId field to given value.

### HasDnsSubdomainId

`func (o *ServerInstanceInterface) HasDnsSubdomainId() bool`

HasDnsSubdomainId returns a boolean if a field has been set.

### GetDnsSubdomainPermanentId

`func (o *ServerInstanceInterface) GetDnsSubdomainPermanentId() int32`

GetDnsSubdomainPermanentId returns the DnsSubdomainPermanentId field if non-nil, zero value otherwise.

### GetDnsSubdomainPermanentIdOk

`func (o *ServerInstanceInterface) GetDnsSubdomainPermanentIdOk() (*int32, bool)`

GetDnsSubdomainPermanentIdOk returns a tuple with the DnsSubdomainPermanentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubdomainPermanentId

`func (o *ServerInstanceInterface) SetDnsSubdomainPermanentId(v int32)`

SetDnsSubdomainPermanentId sets DnsSubdomainPermanentId field to given value.

### HasDnsSubdomainPermanentId

`func (o *ServerInstanceInterface) HasDnsSubdomainPermanentId() bool`

HasDnsSubdomainPermanentId returns a boolean if a field has been set.

### GetInfrastructureId

`func (o *ServerInstanceInterface) GetInfrastructureId() int32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *ServerInstanceInterface) GetInfrastructureIdOk() (*int32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *ServerInstanceInterface) SetInfrastructureId(v int32)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetInstanceId

`func (o *ServerInstanceInterface) GetInstanceId() int32`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ServerInstanceInterface) GetInstanceIdOk() (*int32, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ServerInstanceInterface) SetInstanceId(v int32)`

SetInstanceId sets InstanceId field to given value.


### GetIndex

`func (o *ServerInstanceInterface) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ServerInstanceInterface) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ServerInstanceInterface) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetCapacityMbps

`func (o *ServerInstanceInterface) GetCapacityMbps() int32`

GetCapacityMbps returns the CapacityMbps field if non-nil, zero value otherwise.

### GetCapacityMbpsOk

`func (o *ServerInstanceInterface) GetCapacityMbpsOk() (*int32, bool)`

GetCapacityMbpsOk returns a tuple with the CapacityMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityMbps

`func (o *ServerInstanceInterface) SetCapacityMbps(v int32)`

SetCapacityMbps sets CapacityMbps field to given value.


### GetNetworkId

`func (o *ServerInstanceInterface) GetNetworkId() int32`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *ServerInstanceInterface) GetNetworkIdOk() (*int32, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *ServerInstanceInterface) SetNetworkId(v int32)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *ServerInstanceInterface) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetSubnetPoolSanIndex

`func (o *ServerInstanceInterface) GetSubnetPoolSanIndex() int32`

GetSubnetPoolSanIndex returns the SubnetPoolSanIndex field if non-nil, zero value otherwise.

### GetSubnetPoolSanIndexOk

`func (o *ServerInstanceInterface) GetSubnetPoolSanIndexOk() (*int32, bool)`

GetSubnetPoolSanIndexOk returns a tuple with the SubnetPoolSanIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetPoolSanIndex

`func (o *ServerInstanceInterface) SetSubnetPoolSanIndex(v int32)`

SetSubnetPoolSanIndex sets SubnetPoolSanIndex field to given value.

### HasSubnetPoolSanIndex

`func (o *ServerInstanceInterface) HasSubnetPoolSanIndex() bool`

HasSubnetPoolSanIndex returns a boolean if a field has been set.

### GetVlanIdentifier

`func (o *ServerInstanceInterface) GetVlanIdentifier() int32`

GetVlanIdentifier returns the VlanIdentifier field if non-nil, zero value otherwise.

### GetVlanIdentifierOk

`func (o *ServerInstanceInterface) GetVlanIdentifierOk() (*int32, bool)`

GetVlanIdentifierOk returns a tuple with the VlanIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanIdentifier

`func (o *ServerInstanceInterface) SetVlanIdentifier(v int32)`

SetVlanIdentifier sets VlanIdentifier field to given value.

### HasVlanIdentifier

`func (o *ServerInstanceInterface) HasVlanIdentifier() bool`

HasVlanIdentifier returns a boolean if a field has been set.

### GetAclIdentifier

`func (o *ServerInstanceInterface) GetAclIdentifier() int32`

GetAclIdentifier returns the AclIdentifier field if non-nil, zero value otherwise.

### GetAclIdentifierOk

`func (o *ServerInstanceInterface) GetAclIdentifierOk() (*int32, bool)`

GetAclIdentifierOk returns a tuple with the AclIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclIdentifier

`func (o *ServerInstanceInterface) SetAclIdentifier(v int32)`

SetAclIdentifier sets AclIdentifier field to given value.

### HasAclIdentifier

`func (o *ServerInstanceInterface) HasAclIdentifier() bool`

HasAclIdentifier returns a boolean if a field has been set.

### GetNetworkEquipmentSubnetPoolSanId

`func (o *ServerInstanceInterface) GetNetworkEquipmentSubnetPoolSanId() int32`

GetNetworkEquipmentSubnetPoolSanId returns the NetworkEquipmentSubnetPoolSanId field if non-nil, zero value otherwise.

### GetNetworkEquipmentSubnetPoolSanIdOk

`func (o *ServerInstanceInterface) GetNetworkEquipmentSubnetPoolSanIdOk() (*int32, bool)`

GetNetworkEquipmentSubnetPoolSanIdOk returns a tuple with the NetworkEquipmentSubnetPoolSanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEquipmentSubnetPoolSanId

`func (o *ServerInstanceInterface) SetNetworkEquipmentSubnetPoolSanId(v int32)`

SetNetworkEquipmentSubnetPoolSanId sets NetworkEquipmentSubnetPoolSanId field to given value.

### HasNetworkEquipmentSubnetPoolSanId

`func (o *ServerInstanceInterface) HasNetworkEquipmentSubnetPoolSanId() bool`

HasNetworkEquipmentSubnetPoolSanId returns a boolean if a field has been set.

### GetServerInterfaceId

`func (o *ServerInstanceInterface) GetServerInterfaceId() int32`

GetServerInterfaceId returns the ServerInterfaceId field if non-nil, zero value otherwise.

### GetServerInterfaceIdOk

`func (o *ServerInstanceInterface) GetServerInterfaceIdOk() (*int32, bool)`

GetServerInterfaceIdOk returns a tuple with the ServerInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInterfaceId

`func (o *ServerInstanceInterface) SetServerInterfaceId(v int32)`

SetServerInterfaceId sets ServerInterfaceId field to given value.

### HasServerInterfaceId

`func (o *ServerInstanceInterface) HasServerInterfaceId() bool`

HasServerInterfaceId returns a boolean if a field has been set.

### GetDirtyBit

`func (o *ServerInstanceInterface) GetDirtyBit() bool`

GetDirtyBit returns the DirtyBit field if non-nil, zero value otherwise.

### GetDirtyBitOk

`func (o *ServerInstanceInterface) GetDirtyBitOk() (*bool, bool)`

GetDirtyBitOk returns a tuple with the DirtyBit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirtyBit

`func (o *ServerInstanceInterface) SetDirtyBit(v bool)`

SetDirtyBit sets DirtyBit field to given value.


### GetServiceStatus

`func (o *ServerInstanceInterface) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *ServerInstanceInterface) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *ServerInstanceInterface) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetConfig

`func (o *ServerInstanceInterface) GetConfig() ServerInstanceInterfaceConfiguration`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ServerInstanceInterface) GetConfigOk() (*ServerInstanceInterfaceConfiguration, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ServerInstanceInterface) SetConfig(v ServerInstanceInterfaceConfiguration)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ServerInstanceInterface) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetLinks

`func (o *ServerInstanceInterface) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServerInstanceInterface) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServerInstanceInterface) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServerInstanceInterface) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


