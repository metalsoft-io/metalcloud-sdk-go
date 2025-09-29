# ServerInstanceRecordSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceStatus** | **string** |  | 
**DeployType** | **string** |  | 
**DeployStatus** | **string** |  | 

## Methods

### NewServerInstanceRecordSet

`func NewServerInstanceRecordSet(serviceStatus string, deployType string, deployStatus string, ) *ServerInstanceRecordSet`

NewServerInstanceRecordSet instantiates a new ServerInstanceRecordSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceRecordSetWithDefaults

`func NewServerInstanceRecordSetWithDefaults() *ServerInstanceRecordSet`

NewServerInstanceRecordSetWithDefaults instantiates a new ServerInstanceRecordSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceStatus

`func (o *ServerInstanceRecordSet) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *ServerInstanceRecordSet) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *ServerInstanceRecordSet) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.


### GetDeployType

`func (o *ServerInstanceRecordSet) GetDeployType() string`

GetDeployType returns the DeployType field if non-nil, zero value otherwise.

### GetDeployTypeOk

`func (o *ServerInstanceRecordSet) GetDeployTypeOk() (*string, bool)`

GetDeployTypeOk returns a tuple with the DeployType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployType

`func (o *ServerInstanceRecordSet) SetDeployType(v string)`

SetDeployType sets DeployType field to given value.


### GetDeployStatus

`func (o *ServerInstanceRecordSet) GetDeployStatus() string`

GetDeployStatus returns the DeployStatus field if non-nil, zero value otherwise.

### GetDeployStatusOk

`func (o *ServerInstanceRecordSet) GetDeployStatusOk() (*string, bool)`

GetDeployStatusOk returns a tuple with the DeployStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStatus

`func (o *ServerInstanceRecordSet) SetDeployStatus(v string)`

SetDeployStatus sets DeployStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


