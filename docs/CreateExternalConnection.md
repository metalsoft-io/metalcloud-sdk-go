# CreateExternalConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | The external connection unique label | 
**Name** | **string** | The external connection name | 
**FabricId** | **float32** | The ID of the Fabric identifier this entity belongs to. | 
**ExternalConnectionInterfaces** | Pointer to [**[]CreateExternalConnectionInterface**](CreateExternalConnectionInterface.md) | The external connection interfaces associated with this external connection | [optional] 

## Methods

### NewCreateExternalConnection

`func NewCreateExternalConnection(label string, name string, fabricId float32, ) *CreateExternalConnection`

NewCreateExternalConnection instantiates a new CreateExternalConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateExternalConnectionWithDefaults

`func NewCreateExternalConnectionWithDefaults() *CreateExternalConnection`

NewCreateExternalConnectionWithDefaults instantiates a new CreateExternalConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CreateExternalConnection) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateExternalConnection) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateExternalConnection) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *CreateExternalConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateExternalConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateExternalConnection) SetName(v string)`

SetName sets Name field to given value.


### GetFabricId

`func (o *CreateExternalConnection) GetFabricId() float32`

GetFabricId returns the FabricId field if non-nil, zero value otherwise.

### GetFabricIdOk

`func (o *CreateExternalConnection) GetFabricIdOk() (*float32, bool)`

GetFabricIdOk returns a tuple with the FabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricId

`func (o *CreateExternalConnection) SetFabricId(v float32)`

SetFabricId sets FabricId field to given value.


### GetExternalConnectionInterfaces

`func (o *CreateExternalConnection) GetExternalConnectionInterfaces() []CreateExternalConnectionInterface`

GetExternalConnectionInterfaces returns the ExternalConnectionInterfaces field if non-nil, zero value otherwise.

### GetExternalConnectionInterfacesOk

`func (o *CreateExternalConnection) GetExternalConnectionInterfacesOk() (*[]CreateExternalConnectionInterface, bool)`

GetExternalConnectionInterfacesOk returns a tuple with the ExternalConnectionInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalConnectionInterfaces

`func (o *CreateExternalConnection) SetExternalConnectionInterfaces(v []CreateExternalConnectionInterface)`

SetExternalConnectionInterfaces sets ExternalConnectionInterfaces field to given value.

### HasExternalConnectionInterfaces

`func (o *CreateExternalConnection) HasExternalConnectionInterfaces() bool`

HasExternalConnectionInterfaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


