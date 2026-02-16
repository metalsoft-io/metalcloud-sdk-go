# UpdateStorageInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsUplink** | Pointer to **bool** | Specifies if the Storage Interface is an uplink. Uplink interfaces will be provisioned on network device. | [optional] 
**UseForDeploys** | Pointer to **bool** | Specifies if the Storage Interface is used to deploy storage resources (if multiple interfaces are marked, one will be chosen at random. If no interface is marked, the system will pick a random one automatically for each resource). | [optional] 
**NetworkEquipmentInterfaceId** | Pointer to **float32** | Id of the Network Equipment Interface associated to this Storage Interface | [optional] 

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

### GetNetworkEquipmentInterfaceId

`func (o *UpdateStorageInterface) GetNetworkEquipmentInterfaceId() float32`

GetNetworkEquipmentInterfaceId returns the NetworkEquipmentInterfaceId field if non-nil, zero value otherwise.

### GetNetworkEquipmentInterfaceIdOk

`func (o *UpdateStorageInterface) GetNetworkEquipmentInterfaceIdOk() (*float32, bool)`

GetNetworkEquipmentInterfaceIdOk returns a tuple with the NetworkEquipmentInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEquipmentInterfaceId

`func (o *UpdateStorageInterface) SetNetworkEquipmentInterfaceId(v float32)`

SetNetworkEquipmentInterfaceId sets NetworkEquipmentInterfaceId field to given value.

### HasNetworkEquipmentInterfaceId

`func (o *UpdateStorageInterface) HasNetworkEquipmentInterfaceId() bool`

HasNetworkEquipmentInterfaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


