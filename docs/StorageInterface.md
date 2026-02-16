# StorageInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | Id of the Storage Interface | 
**Revision** | **float32** | Revision of the Storage Interface | 
**StorageId** | **float32** | Id of the Storage | 
**Name** | **string** | Name of the Storage Interface | 
**NodeIds** | **[]string** | Node IDs on which the Storage Interface is located | 
**Protocols** | **[]string** | Types of the Storage Interface | 
**IsUplink** | **bool** | Specifies if the Storage Interface is an uplink. Uplink interfaces will be provisioned on network device. | 
**UseForDeploys** | **bool** | Specifies if the Storage Interface is used to deploy storage resources (if multiple interfaces are marked, one will be chosen at random. If no interface is marked, the system will pick a random one automatically for each resource). | 
**NetworkEquipmentInterfaceId** | Pointer to **float32** | Id of the Network Equipment Interface associated to this Storage Interface | [optional] 
**LinkedNetworkDeviceInformation** | Pointer to [**StorageInterfaceNetworkDeviceInformation**](StorageInterfaceNetworkDeviceInformation.md) | Network device information | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewStorageInterface

`func NewStorageInterface(id float32, revision float32, storageId float32, name string, nodeIds []string, protocols []string, isUplink bool, useForDeploys bool, ) *StorageInterface`

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

`func (o *StorageInterface) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageInterface) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageInterface) SetId(v float32)`

SetId sets Id field to given value.


### GetRevision

`func (o *StorageInterface) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *StorageInterface) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *StorageInterface) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetStorageId

`func (o *StorageInterface) GetStorageId() float32`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *StorageInterface) GetStorageIdOk() (*float32, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *StorageInterface) SetStorageId(v float32)`

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


### GetNetworkEquipmentInterfaceId

`func (o *StorageInterface) GetNetworkEquipmentInterfaceId() float32`

GetNetworkEquipmentInterfaceId returns the NetworkEquipmentInterfaceId field if non-nil, zero value otherwise.

### GetNetworkEquipmentInterfaceIdOk

`func (o *StorageInterface) GetNetworkEquipmentInterfaceIdOk() (*float32, bool)`

GetNetworkEquipmentInterfaceIdOk returns a tuple with the NetworkEquipmentInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEquipmentInterfaceId

`func (o *StorageInterface) SetNetworkEquipmentInterfaceId(v float32)`

SetNetworkEquipmentInterfaceId sets NetworkEquipmentInterfaceId field to given value.

### HasNetworkEquipmentInterfaceId

`func (o *StorageInterface) HasNetworkEquipmentInterfaceId() bool`

HasNetworkEquipmentInterfaceId returns a boolean if a field has been set.

### GetLinkedNetworkDeviceInformation

`func (o *StorageInterface) GetLinkedNetworkDeviceInformation() StorageInterfaceNetworkDeviceInformation`

GetLinkedNetworkDeviceInformation returns the LinkedNetworkDeviceInformation field if non-nil, zero value otherwise.

### GetLinkedNetworkDeviceInformationOk

`func (o *StorageInterface) GetLinkedNetworkDeviceInformationOk() (*StorageInterfaceNetworkDeviceInformation, bool)`

GetLinkedNetworkDeviceInformationOk returns a tuple with the LinkedNetworkDeviceInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedNetworkDeviceInformation

`func (o *StorageInterface) SetLinkedNetworkDeviceInformation(v StorageInterfaceNetworkDeviceInformation)`

SetLinkedNetworkDeviceInformation sets LinkedNetworkDeviceInformation field to given value.

### HasLinkedNetworkDeviceInformation

`func (o *StorageInterface) HasLinkedNetworkDeviceInformation() bool`

HasLinkedNetworkDeviceInformation returns a boolean if a field has been set.

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


