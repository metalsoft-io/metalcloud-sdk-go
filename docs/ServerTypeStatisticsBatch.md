# ServerTypeStatisticsBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerTypeIdToServerCount** | **map[string]interface{}** | An object having server type id as key and server count as value. | 
**ServerTypeIdToServerInformation** | **map[string]interface{}** | An object having server type id as key and server information as value. | 
**UtilizationReport** | [**ServerTypeUtilizationReport**](ServerTypeUtilizationReport.md) | The utilization report for the server types. | 

## Methods

### NewServerTypeStatisticsBatch

`func NewServerTypeStatisticsBatch(serverTypeIdToServerCount map[string]interface{}, serverTypeIdToServerInformation map[string]interface{}, utilizationReport ServerTypeUtilizationReport, ) *ServerTypeStatisticsBatch`

NewServerTypeStatisticsBatch instantiates a new ServerTypeStatisticsBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerTypeStatisticsBatchWithDefaults

`func NewServerTypeStatisticsBatchWithDefaults() *ServerTypeStatisticsBatch`

NewServerTypeStatisticsBatchWithDefaults instantiates a new ServerTypeStatisticsBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerTypeIdToServerCount

`func (o *ServerTypeStatisticsBatch) GetServerTypeIdToServerCount() map[string]interface{}`

GetServerTypeIdToServerCount returns the ServerTypeIdToServerCount field if non-nil, zero value otherwise.

### GetServerTypeIdToServerCountOk

`func (o *ServerTypeStatisticsBatch) GetServerTypeIdToServerCountOk() (*map[string]interface{}, bool)`

GetServerTypeIdToServerCountOk returns a tuple with the ServerTypeIdToServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeIdToServerCount

`func (o *ServerTypeStatisticsBatch) SetServerTypeIdToServerCount(v map[string]interface{})`

SetServerTypeIdToServerCount sets ServerTypeIdToServerCount field to given value.


### GetServerTypeIdToServerInformation

`func (o *ServerTypeStatisticsBatch) GetServerTypeIdToServerInformation() map[string]interface{}`

GetServerTypeIdToServerInformation returns the ServerTypeIdToServerInformation field if non-nil, zero value otherwise.

### GetServerTypeIdToServerInformationOk

`func (o *ServerTypeStatisticsBatch) GetServerTypeIdToServerInformationOk() (*map[string]interface{}, bool)`

GetServerTypeIdToServerInformationOk returns a tuple with the ServerTypeIdToServerInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeIdToServerInformation

`func (o *ServerTypeStatisticsBatch) SetServerTypeIdToServerInformation(v map[string]interface{})`

SetServerTypeIdToServerInformation sets ServerTypeIdToServerInformation field to given value.


### GetUtilizationReport

`func (o *ServerTypeStatisticsBatch) GetUtilizationReport() ServerTypeUtilizationReport`

GetUtilizationReport returns the UtilizationReport field if non-nil, zero value otherwise.

### GetUtilizationReportOk

`func (o *ServerTypeStatisticsBatch) GetUtilizationReportOk() (*ServerTypeUtilizationReport, bool)`

GetUtilizationReportOk returns a tuple with the UtilizationReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationReport

`func (o *ServerTypeStatisticsBatch) SetUtilizationReport(v ServerTypeUtilizationReport)`

SetUtilizationReport sets UtilizationReport field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


