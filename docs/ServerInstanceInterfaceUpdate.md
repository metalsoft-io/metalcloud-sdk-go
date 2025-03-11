# ServerInstanceInterfaceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapacityMbps** | Pointer to **int32** |  | [optional] 
**NetworkId** | Pointer to **int32** | The ID of the network to which this interface is to be attached to. | [optional] 
**Label** | Pointer to **string** | The server instance interface label. | [optional] 

## Methods

### NewServerInstanceInterfaceUpdate

`func NewServerInstanceInterfaceUpdate() *ServerInstanceInterfaceUpdate`

NewServerInstanceInterfaceUpdate instantiates a new ServerInstanceInterfaceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceInterfaceUpdateWithDefaults

`func NewServerInstanceInterfaceUpdateWithDefaults() *ServerInstanceInterfaceUpdate`

NewServerInstanceInterfaceUpdateWithDefaults instantiates a new ServerInstanceInterfaceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacityMbps

`func (o *ServerInstanceInterfaceUpdate) GetCapacityMbps() int32`

GetCapacityMbps returns the CapacityMbps field if non-nil, zero value otherwise.

### GetCapacityMbpsOk

`func (o *ServerInstanceInterfaceUpdate) GetCapacityMbpsOk() (*int32, bool)`

GetCapacityMbpsOk returns a tuple with the CapacityMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityMbps

`func (o *ServerInstanceInterfaceUpdate) SetCapacityMbps(v int32)`

SetCapacityMbps sets CapacityMbps field to given value.

### HasCapacityMbps

`func (o *ServerInstanceInterfaceUpdate) HasCapacityMbps() bool`

HasCapacityMbps returns a boolean if a field has been set.

### GetNetworkId

`func (o *ServerInstanceInterfaceUpdate) GetNetworkId() int32`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *ServerInstanceInterfaceUpdate) GetNetworkIdOk() (*int32, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *ServerInstanceInterfaceUpdate) SetNetworkId(v int32)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *ServerInstanceInterfaceUpdate) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetLabel

`func (o *ServerInstanceInterfaceUpdate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServerInstanceInterfaceUpdate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServerInstanceInterfaceUpdate) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ServerInstanceInterfaceUpdate) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


