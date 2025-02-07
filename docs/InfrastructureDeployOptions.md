# InfrastructureDeployOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowDataLoss** | **bool** | Allow data loss | 
**ShutdownOptions** | [**InfrastructureDeployShutdownOptions**](InfrastructureDeployShutdownOptions.md) |  | 
**ServerTypeIdToPreferredServerIds** | Pointer to **map[string]interface{}** | An object having as key the server type id and as value an array of preferred server ids | [optional] 

## Methods

### NewInfrastructureDeployOptions

`func NewInfrastructureDeployOptions(allowDataLoss bool, shutdownOptions InfrastructureDeployShutdownOptions, ) *InfrastructureDeployOptions`

NewInfrastructureDeployOptions instantiates a new InfrastructureDeployOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureDeployOptionsWithDefaults

`func NewInfrastructureDeployOptionsWithDefaults() *InfrastructureDeployOptions`

NewInfrastructureDeployOptionsWithDefaults instantiates a new InfrastructureDeployOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowDataLoss

`func (o *InfrastructureDeployOptions) GetAllowDataLoss() bool`

GetAllowDataLoss returns the AllowDataLoss field if non-nil, zero value otherwise.

### GetAllowDataLossOk

`func (o *InfrastructureDeployOptions) GetAllowDataLossOk() (*bool, bool)`

GetAllowDataLossOk returns a tuple with the AllowDataLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDataLoss

`func (o *InfrastructureDeployOptions) SetAllowDataLoss(v bool)`

SetAllowDataLoss sets AllowDataLoss field to given value.


### GetShutdownOptions

`func (o *InfrastructureDeployOptions) GetShutdownOptions() InfrastructureDeployShutdownOptions`

GetShutdownOptions returns the ShutdownOptions field if non-nil, zero value otherwise.

### GetShutdownOptionsOk

`func (o *InfrastructureDeployOptions) GetShutdownOptionsOk() (*InfrastructureDeployShutdownOptions, bool)`

GetShutdownOptionsOk returns a tuple with the ShutdownOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdownOptions

`func (o *InfrastructureDeployOptions) SetShutdownOptions(v InfrastructureDeployShutdownOptions)`

SetShutdownOptions sets ShutdownOptions field to given value.


### GetServerTypeIdToPreferredServerIds

`func (o *InfrastructureDeployOptions) GetServerTypeIdToPreferredServerIds() map[string]interface{}`

GetServerTypeIdToPreferredServerIds returns the ServerTypeIdToPreferredServerIds field if non-nil, zero value otherwise.

### GetServerTypeIdToPreferredServerIdsOk

`func (o *InfrastructureDeployOptions) GetServerTypeIdToPreferredServerIdsOk() (*map[string]interface{}, bool)`

GetServerTypeIdToPreferredServerIdsOk returns a tuple with the ServerTypeIdToPreferredServerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeIdToPreferredServerIds

`func (o *InfrastructureDeployOptions) SetServerTypeIdToPreferredServerIds(v map[string]interface{})`

SetServerTypeIdToPreferredServerIds sets ServerTypeIdToPreferredServerIds field to given value.

### HasServerTypeIdToPreferredServerIds

`func (o *InfrastructureDeployOptions) HasServerTypeIdToPreferredServerIds() bool`

HasServerTypeIdToPreferredServerIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


