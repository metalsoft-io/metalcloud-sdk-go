# RegisterProductionServerSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InfrastructureId** | **float32** | The id of the infrastructure where the server is located. | 
**OsTemplateId** | Pointer to **float32** | The id of the template installed on the server. | [optional] 
**InterfaceConnections** | Pointer to **[]string** | Interface to network device connection settings. | [optional] 

## Methods

### NewRegisterProductionServerSettings

`func NewRegisterProductionServerSettings(infrastructureId float32, ) *RegisterProductionServerSettings`

NewRegisterProductionServerSettings instantiates a new RegisterProductionServerSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterProductionServerSettingsWithDefaults

`func NewRegisterProductionServerSettingsWithDefaults() *RegisterProductionServerSettings`

NewRegisterProductionServerSettingsWithDefaults instantiates a new RegisterProductionServerSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfrastructureId

`func (o *RegisterProductionServerSettings) GetInfrastructureId() float32`

GetInfrastructureId returns the InfrastructureId field if non-nil, zero value otherwise.

### GetInfrastructureIdOk

`func (o *RegisterProductionServerSettings) GetInfrastructureIdOk() (*float32, bool)`

GetInfrastructureIdOk returns a tuple with the InfrastructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureId

`func (o *RegisterProductionServerSettings) SetInfrastructureId(v float32)`

SetInfrastructureId sets InfrastructureId field to given value.


### GetOsTemplateId

`func (o *RegisterProductionServerSettings) GetOsTemplateId() float32`

GetOsTemplateId returns the OsTemplateId field if non-nil, zero value otherwise.

### GetOsTemplateIdOk

`func (o *RegisterProductionServerSettings) GetOsTemplateIdOk() (*float32, bool)`

GetOsTemplateIdOk returns a tuple with the OsTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsTemplateId

`func (o *RegisterProductionServerSettings) SetOsTemplateId(v float32)`

SetOsTemplateId sets OsTemplateId field to given value.

### HasOsTemplateId

`func (o *RegisterProductionServerSettings) HasOsTemplateId() bool`

HasOsTemplateId returns a boolean if a field has been set.

### GetInterfaceConnections

`func (o *RegisterProductionServerSettings) GetInterfaceConnections() []string`

GetInterfaceConnections returns the InterfaceConnections field if non-nil, zero value otherwise.

### GetInterfaceConnectionsOk

`func (o *RegisterProductionServerSettings) GetInterfaceConnectionsOk() (*[]string, bool)`

GetInterfaceConnectionsOk returns a tuple with the InterfaceConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceConnections

`func (o *RegisterProductionServerSettings) SetInterfaceConnections(v []string)`

SetInterfaceConnections sets InterfaceConnections field to given value.

### HasInterfaceConnections

`func (o *RegisterProductionServerSettings) HasInterfaceConnections() bool`

HasInterfaceConnections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


