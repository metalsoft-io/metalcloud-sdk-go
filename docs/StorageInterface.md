# StorageInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Id of the Storage Interface | 
**Revision** | **int64** | Revision of the Storage Interface | 
**StorageId** | **int64** | Id of the Storage | 
**Name** | **string** | Name of the Storage Interface | 
**NodeIds** | **[]string** | Node IDs on which the Storage Interface is located | 
**Protocols** | **[]string** | Types of the Storage Interface | 
**IsUplink** | **bool** | Specifies if the Storage Interface is an uplink. Uplink interfaces will be provisioned on network device. | 
**UseForDeploys** | **bool** | Specifies if the Storage Interface is used to deploy storage resources (if multiple interfaces are marked, one will be chosen at random. If no interface is marked, the system will pick a random one automatically for each resource). | 
**NetworkEquipmentInterfaceIds** | Pointer to **[]float32** | Ids of every Network Equipment Interface this Storage Interface is reachable through. More than one is normal: an MLAG pair contributes one port per switch, and a bond on the storage side may face several independent switch ports. Sending this field replaces the whole set. | [optional] 
**LinkedNetworkDeviceInformation** | Pointer to [**[]StorageInterfaceNetworkDeviceInformation**](StorageInterfaceNetworkDeviceInformation.md) | Network device information, one entry per linked Network Equipment Interface | [optional] 
**Aggregation** | Pointer to **string** | How the linked switch ports are aggregated. Read-only and derived from the switch side: ports sharing a link aggregation group report \&quot;lag\&quot; (or \&quot;mlag\&quot; when that group is distributed across two switches), and ports in no group report \&quot;none\&quot;. A set of ports expected to be bonded but reporting \&quot;none\&quot; means the link aggregation group is missing on the switches. | [optional] 
**LagIdentifier** | Pointer to **int32** | Switch-facing port-channel number of the derived link aggregation group, if any. Read-only. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewStorageInterface

`func NewStorageInterface(id int64, revision int64, storageId int64, name string, nodeIds []string, protocols []string, isUplink bool, useForDeploys bool, ) *StorageInterface`

NewStorageInterface instantiates a new StorageInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageInterfaceWithDefaults

`func NewStorageInterfaceWithDefaults() *StorageInterface`

NewStorageInterfaceWithDefaults instantiates a new StorageInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorageInterface) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageInterface) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageInterface) SetId(v int64)`

SetId sets Id field to given value.


### GetRevision

`func (o *StorageInterface) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *StorageInterface) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *StorageInterface) SetRevision(v int64)`

SetRevision sets Revision field to given value.


### GetStorageId

`func (o *StorageInterface) GetStorageId() int64`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *StorageInterface) GetStorageIdOk() (*int64, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *StorageInterface) SetStorageId(v int64)`

SetStorageId sets StorageId field to given value.


### GetName

`func (o *StorageInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageInterface) SetName(v string)`

SetName sets Name field to given value.


### GetNodeIds

`func (o *StorageInterface) GetNodeIds() []string`

GetNodeIds returns the NodeIds field if non-nil, zero value otherwise.

### GetNodeIdsOk

`func (o *StorageInterface) GetNodeIdsOk() (*[]string, bool)`

GetNodeIdsOk returns a tuple with the NodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIds

`func (o *StorageInterface) SetNodeIds(v []string)`

SetNodeIds sets NodeIds field to given value.


### GetProtocols

`func (o *StorageInterface) GetProtocols() []string`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *StorageInterface) GetProtocolsOk() (*[]string, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *StorageInterface) SetProtocols(v []string)`

SetProtocols sets Protocols field to given value.


### GetIsUplink

`func (o *StorageInterface) GetIsUplink() bool`

GetIsUplink returns the IsUplink field if non-nil, zero value otherwise.

### GetIsUplinkOk

`func (o *StorageInterface) GetIsUplinkOk() (*bool, bool)`

GetIsUplinkOk returns a tuple with the IsUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUplink

`func (o *StorageInterface) SetIsUplink(v bool)`

SetIsUplink sets IsUplink field to given value.


### GetUseForDeploys

`func (o *StorageInterface) GetUseForDeploys() bool`

GetUseForDeploys returns the UseForDeploys field if non-nil, zero value otherwise.

### GetUseForDeploysOk

`func (o *StorageInterface) GetUseForDeploysOk() (*bool, bool)`

GetUseForDeploysOk returns a tuple with the UseForDeploys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseForDeploys

`func (o *StorageInterface) SetUseForDeploys(v bool)`

SetUseForDeploys sets UseForDeploys field to given value.


### GetNetworkEquipmentInterfaceIds

`func (o *StorageInterface) GetNetworkEquipmentInterfaceIds() []float32`

GetNetworkEquipmentInterfaceIds returns the NetworkEquipmentInterfaceIds field if non-nil, zero value otherwise.

### GetNetworkEquipmentInterfaceIdsOk

`func (o *StorageInterface) GetNetworkEquipmentInterfaceIdsOk() (*[]float32, bool)`

GetNetworkEquipmentInterfaceIdsOk returns a tuple with the NetworkEquipmentInterfaceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEquipmentInterfaceIds

`func (o *StorageInterface) SetNetworkEquipmentInterfaceIds(v []float32)`

SetNetworkEquipmentInterfaceIds sets NetworkEquipmentInterfaceIds field to given value.

### HasNetworkEquipmentInterfaceIds

`func (o *StorageInterface) HasNetworkEquipmentInterfaceIds() bool`

HasNetworkEquipmentInterfaceIds returns a boolean if a field has been set.

### GetLinkedNetworkDeviceInformation

`func (o *StorageInterface) GetLinkedNetworkDeviceInformation() []StorageInterfaceNetworkDeviceInformation`

GetLinkedNetworkDeviceInformation returns the LinkedNetworkDeviceInformation field if non-nil, zero value otherwise.

### GetLinkedNetworkDeviceInformationOk

`func (o *StorageInterface) GetLinkedNetworkDeviceInformationOk() (*[]StorageInterfaceNetworkDeviceInformation, bool)`

GetLinkedNetworkDeviceInformationOk returns a tuple with the LinkedNetworkDeviceInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedNetworkDeviceInformation

`func (o *StorageInterface) SetLinkedNetworkDeviceInformation(v []StorageInterfaceNetworkDeviceInformation)`

SetLinkedNetworkDeviceInformation sets LinkedNetworkDeviceInformation field to given value.

### HasLinkedNetworkDeviceInformation

`func (o *StorageInterface) HasLinkedNetworkDeviceInformation() bool`

HasLinkedNetworkDeviceInformation returns a boolean if a field has been set.

### GetAggregation

`func (o *StorageInterface) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *StorageInterface) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *StorageInterface) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *StorageInterface) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetLagIdentifier

`func (o *StorageInterface) GetLagIdentifier() int32`

GetLagIdentifier returns the LagIdentifier field if non-nil, zero value otherwise.

### GetLagIdentifierOk

`func (o *StorageInterface) GetLagIdentifierOk() (*int32, bool)`

GetLagIdentifierOk returns a tuple with the LagIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagIdentifier

`func (o *StorageInterface) SetLagIdentifier(v int32)`

SetLagIdentifier sets LagIdentifier field to given value.

### HasLagIdentifier

`func (o *StorageInterface) HasLagIdentifier() bool`

HasLagIdentifier returns a boolean if a field has been set.

### GetLinks

`func (o *StorageInterface) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StorageInterface) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StorageInterface) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StorageInterface) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


