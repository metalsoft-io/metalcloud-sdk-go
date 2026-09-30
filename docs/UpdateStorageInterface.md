# UpdateStorageInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsUplink** | Pointer to **bool** | Specifies if the Storage Interface is an uplink. Uplink interfaces will be provisioned on network device. | [optional] 
**UseForDeploys** | Pointer to **bool** | Specifies if the Storage Interface is used to deploy storage resources (if multiple interfaces are marked, one will be chosen at random. If no interface is marked, the system will pick a random one automatically for each resource). | [optional] 
**NetworkEquipmentInterfaceIds** | Pointer to **[]float32** | Ids of every Network Equipment Interface this Storage Interface is reachable through. More than one is normal: an MLAG pair contributes one port per switch, and a bond on the storage side may face several independent switch ports. Sending this field replaces the whole set. | [optional] 

## Methods

### NewUpdateStorageInterface

`func NewUpdateStorageInterface() *UpdateStorageInterface`

NewUpdateStorageInterface instantiates a new UpdateStorageInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStorageInterfaceWithDefaults

`func NewUpdateStorageInterfaceWithDefaults() *UpdateStorageInterface`

NewUpdateStorageInterfaceWithDefaults instantiates a new UpdateStorageInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsUplink

`func (o *UpdateStorageInterface) GetIsUplink() bool`

GetIsUplink returns the IsUplink field if non-nil, zero value otherwise.

### GetIsUplinkOk

`func (o *UpdateStorageInterface) GetIsUplinkOk() (*bool, bool)`

GetIsUplinkOk returns a tuple with the IsUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUplink

`func (o *UpdateStorageInterface) SetIsUplink(v bool)`

SetIsUplink sets IsUplink field to given value.

### HasIsUplink

`func (o *UpdateStorageInterface) HasIsUplink() bool`

HasIsUplink returns a boolean if a field has been set.

### GetUseForDeploys

`func (o *UpdateStorageInterface) GetUseForDeploys() bool`

GetUseForDeploys returns the UseForDeploys field if non-nil, zero value otherwise.

### GetUseForDeploysOk

`func (o *UpdateStorageInterface) GetUseForDeploysOk() (*bool, bool)`

GetUseForDeploysOk returns a tuple with the UseForDeploys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseForDeploys

`func (o *UpdateStorageInterface) SetUseForDeploys(v bool)`

SetUseForDeploys sets UseForDeploys field to given value.

### HasUseForDeploys

`func (o *UpdateStorageInterface) HasUseForDeploys() bool`

HasUseForDeploys returns a boolean if a field has been set.

### GetNetworkEquipmentInterfaceIds

`func (o *UpdateStorageInterface) GetNetworkEquipmentInterfaceIds() []float32`

GetNetworkEquipmentInterfaceIds returns the NetworkEquipmentInterfaceIds field if non-nil, zero value otherwise.

### GetNetworkEquipmentInterfaceIdsOk

`func (o *UpdateStorageInterface) GetNetworkEquipmentInterfaceIdsOk() (*[]float32, bool)`

GetNetworkEquipmentInterfaceIdsOk returns a tuple with the NetworkEquipmentInterfaceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEquipmentInterfaceIds

`func (o *UpdateStorageInterface) SetNetworkEquipmentInterfaceIds(v []float32)`

SetNetworkEquipmentInterfaceIds sets NetworkEquipmentInterfaceIds field to given value.

### HasNetworkEquipmentInterfaceIds

`func (o *UpdateStorageInterface) HasNetworkEquipmentInterfaceIds() bool`

HasNetworkEquipmentInterfaceIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


