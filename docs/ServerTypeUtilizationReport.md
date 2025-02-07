# ServerTypeUtilizationReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupByServerRamGb** | [**ServerTypeUtilizationReportGrouped**](ServerTypeUtilizationReportGrouped.md) | The utilization report for the server types grouped by server ram gb. | 
**GroupByServerTypeName** | [**ServerTypeUtilizationReportGrouped**](ServerTypeUtilizationReportGrouped.md) | The utilization report for the server types grouped by server type name. | 
**GroupByServerProductName** | [**ServerTypeUtilizationReportGrouped**](ServerTypeUtilizationReportGrouped.md) | The utilization report for the server types grouped by server product name. | 
**GroupByUserId** | [**ServerTypeUtilizationReportGrouped**](ServerTypeUtilizationReportGrouped.md) | The utilization report for the server types grouped by user id. | 

## Methods

### NewServerTypeUtilizationReport

`func NewServerTypeUtilizationReport(groupByServerRamGb ServerTypeUtilizationReportGrouped, groupByServerTypeName ServerTypeUtilizationReportGrouped, groupByServerProductName ServerTypeUtilizationReportGrouped, groupByUserId ServerTypeUtilizationReportGrouped, ) *ServerTypeUtilizationReport`

NewServerTypeUtilizationReport instantiates a new ServerTypeUtilizationReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerTypeUtilizationReportWithDefaults

`func NewServerTypeUtilizationReportWithDefaults() *ServerTypeUtilizationReport`

NewServerTypeUtilizationReportWithDefaults instantiates a new ServerTypeUtilizationReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupByServerRamGb

`func (o *ServerTypeUtilizationReport) GetGroupByServerRamGb() ServerTypeUtilizationReportGrouped`

GetGroupByServerRamGb returns the GroupByServerRamGb field if non-nil, zero value otherwise.

### GetGroupByServerRamGbOk

`func (o *ServerTypeUtilizationReport) GetGroupByServerRamGbOk() (*ServerTypeUtilizationReportGrouped, bool)`

GetGroupByServerRamGbOk returns a tuple with the GroupByServerRamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByServerRamGb

`func (o *ServerTypeUtilizationReport) SetGroupByServerRamGb(v ServerTypeUtilizationReportGrouped)`

SetGroupByServerRamGb sets GroupByServerRamGb field to given value.


### GetGroupByServerTypeName

`func (o *ServerTypeUtilizationReport) GetGroupByServerTypeName() ServerTypeUtilizationReportGrouped`

GetGroupByServerTypeName returns the GroupByServerTypeName field if non-nil, zero value otherwise.

### GetGroupByServerTypeNameOk

`func (o *ServerTypeUtilizationReport) GetGroupByServerTypeNameOk() (*ServerTypeUtilizationReportGrouped, bool)`

GetGroupByServerTypeNameOk returns a tuple with the GroupByServerTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByServerTypeName

`func (o *ServerTypeUtilizationReport) SetGroupByServerTypeName(v ServerTypeUtilizationReportGrouped)`

SetGroupByServerTypeName sets GroupByServerTypeName field to given value.


### GetGroupByServerProductName

`func (o *ServerTypeUtilizationReport) GetGroupByServerProductName() ServerTypeUtilizationReportGrouped`

GetGroupByServerProductName returns the GroupByServerProductName field if non-nil, zero value otherwise.

### GetGroupByServerProductNameOk

`func (o *ServerTypeUtilizationReport) GetGroupByServerProductNameOk() (*ServerTypeUtilizationReportGrouped, bool)`

GetGroupByServerProductNameOk returns a tuple with the GroupByServerProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByServerProductName

`func (o *ServerTypeUtilizationReport) SetGroupByServerProductName(v ServerTypeUtilizationReportGrouped)`

SetGroupByServerProductName sets GroupByServerProductName field to given value.


### GetGroupByUserId

`func (o *ServerTypeUtilizationReport) GetGroupByUserId() ServerTypeUtilizationReportGrouped`

GetGroupByUserId returns the GroupByUserId field if non-nil, zero value otherwise.

### GetGroupByUserIdOk

`func (o *ServerTypeUtilizationReport) GetGroupByUserIdOk() (*ServerTypeUtilizationReportGrouped, bool)`

GetGroupByUserIdOk returns a tuple with the GroupByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByUserId

`func (o *ServerTypeUtilizationReport) SetGroupByUserId(v ServerTypeUtilizationReportGrouped)`

SetGroupByUserId sets GroupByUserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


