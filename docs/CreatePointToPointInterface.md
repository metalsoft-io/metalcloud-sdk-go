# CreatePointToPointInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**PointToPointInterfaceType**](PointToPointInterfaceType.md) | Interface type. Direct creation of a PointToPointLink supports &#x60;network_equipment_interface&#x60; only; &#x60;server_interface&#x60; interfaces are attached at deploy time via a NetworkEndpointGroup attachment. | 
**InterfaceId** | **int64** | ID of the referenced interface (switch interface id or server interface id, per &#x60;type&#x60;). | 

## Methods

### NewCreatePointToPointInterface

`func NewCreatePointToPointInterface(type_ PointToPointInterfaceType, interfaceId int64, ) *CreatePointToPointInterface`

NewCreatePointToPointInterface instantiates a new CreatePointToPointInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePointToPointInterfaceWithDefaults

`func NewCreatePointToPointInterfaceWithDefaults() *CreatePointToPointInterface`

NewCreatePointToPointInterfaceWithDefaults instantiates a new CreatePointToPointInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreatePointToPointInterface) GetType() PointToPointInterfaceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreatePointToPointInterface) GetTypeOk() (*PointToPointInterfaceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreatePointToPointInterface) SetType(v PointToPointInterfaceType)`

SetType sets Type field to given value.


### GetInterfaceId

`func (o *CreatePointToPointInterface) GetInterfaceId() int64`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *CreatePointToPointInterface) GetInterfaceIdOk() (*int64, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *CreatePointToPointInterface) SetInterfaceId(v int64)`

SetInterfaceId sets InterfaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


