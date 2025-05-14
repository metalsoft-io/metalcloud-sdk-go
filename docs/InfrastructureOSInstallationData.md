# InfrastructureOSInstallationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** |  | 
**CustomVariables** | Pointer to **map[string]interface{}** | Custom variables in JSON format. | [optional] 

## Methods

### NewInfrastructureOSInstallationData

`func NewInfrastructureOSInstallationData(label string, ) *InfrastructureOSInstallationData`

NewInfrastructureOSInstallationData instantiates a new InfrastructureOSInstallationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureOSInstallationDataWithDefaults

`func NewInfrastructureOSInstallationDataWithDefaults() *InfrastructureOSInstallationData`

NewInfrastructureOSInstallationDataWithDefaults instantiates a new InfrastructureOSInstallationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *InfrastructureOSInstallationData) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *InfrastructureOSInstallationData) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *InfrastructureOSInstallationData) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetCustomVariables

`func (o *InfrastructureOSInstallationData) GetCustomVariables() map[string]interface{}`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *InfrastructureOSInstallationData) GetCustomVariablesOk() (*map[string]interface{}, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *InfrastructureOSInstallationData) SetCustomVariables(v map[string]interface{})`

SetCustomVariables sets CustomVariables field to given value.

### HasCustomVariables

`func (o *InfrastructureOSInstallationData) HasCustomVariables() bool`

HasCustomVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


