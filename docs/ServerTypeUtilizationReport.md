# ServerTypeUtilizationReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupByServerRamGb** | [**map[string]ServerTypeUtilizationReportGrouped**](ServerTypeUtilizationReportGrouped.md) | The utilization report for the server types grouped by server ram gb. | 
**GroupByServerTypeName** | [**map[string]ServerTypeUtilizationReportGrouped**](ServerTypeUtilizationReportGrouped.md) | The utilization report for the server types grouped by server type name. | 
**GroupByServerProductName** | [**map[string]ServerTypeUtilizationReportGrouped**](ServerTypeUtilizationReportGrouped.md) | The utilization report for the server types grouped by server product name. | 
**GroupByUserId** | [**map[string]ServerTypeUtilizationReportGrouped**](ServerTypeUtilizationReportGrouped.md) | The utilization report for the server types grouped by user id. | 

## Methods

### NewServerTypeUtilizationReport

`func NewServerTypeUtilizationReport(groupByServerRamGb map[string]ServerTypeUtilizationReportGrouped, groupByServerTypeName map[string]ServerTypeUtilizationReportGrouped, groupByServerProductName map[string]ServerTypeUtilizationReportGrouped, groupByUserId map[string]ServerTypeUtilizationReportGrouped, ) *ServerTypeUtilizationReport`

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

`func (o *ServerTypeUtilizationReport) GetGroupByServerRamGb() map[string]ServerTypeUtilizationReportGrouped`

GetGroupByServerRamGb returns the GroupByServerRamGb field if non-nil, zero value otherwise.

### GetGroupByServerRamGbOk

`func (o *ServerTypeUtilizationReport) GetGroupByServerRamGbOk() (*map[string]ServerTypeUtilizationReportGrouped, bool)`

GetGroupByServerRamGbOk returns a tuple with the GroupByServerRamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByServerRamGb

`func (o *ServerTypeUtilizationReport) SetGroupByServerRamGb(v map[string]ServerTypeUtilizationReportGrouped)`

SetGroupByServerRamGb sets GroupByServerRamGb field to given value.


### GetGroupByServerTypeName

`func (o *ServerTypeUtilizationReport) GetGroupByServerTypeName() map[string]ServerTypeUtilizationReportGrouped`

GetGroupByServerTypeName returns the GroupByServerTypeName field if non-nil, zero value otherwise.

### GetGroupByServerTypeNameOk

`func (o *ServerTypeUtilizationReport) GetGroupByServerTypeNameOk() (*map[string]ServerTypeUtilizationReportGrouped, bool)`

GetGroupByServerTypeNameOk returns a tuple with the GroupByServerTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByServerTypeName

`func (o *ServerTypeUtilizationReport) SetGroupByServerTypeName(v map[string]ServerTypeUtilizationReportGrouped)`

SetGroupByServerTypeName sets GroupByServerTypeName field to given value.


### GetGroupByServerProductName

`func (o *ServerTypeUtilizationReport) GetGroupByServerProductName() map[string]ServerTypeUtilizationReportGrouped`

GetGroupByServerProductName returns the GroupByServerProductName field if non-nil, zero value otherwise.

### GetGroupByServerProductNameOk

`func (o *ServerTypeUtilizationReport) GetGroupByServerProductNameOk() (*map[string]ServerTypeUtilizationReportGrouped, bool)`

GetGroupByServerProductNameOk returns a tuple with the GroupByServerProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByServerProductName

`func (o *ServerTypeUtilizationReport) SetGroupByServerProductName(v map[string]ServerTypeUtilizationReportGrouped)`

SetGroupByServerProductName sets GroupByServerProductName field to given value.


### GetGroupByUserId

`func (o *ServerTypeUtilizationReport) GetGroupByUserId() map[string]ServerTypeUtilizationReportGrouped`

GetGroupByUserId returns the GroupByUserId field if non-nil, zero value otherwise.

### GetGroupByUserIdOk

`func (o *ServerTypeUtilizationReport) GetGroupByUserIdOk() (*map[string]ServerTypeUtilizationReportGrouped, bool)`

GetGroupByUserIdOk returns a tuple with the GroupByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByUserId

`func (o *ServerTypeUtilizationReport) SetGroupByUserId(v map[string]ServerTypeUtilizationReportGrouped)`

SetGroupByUserId sets GroupByUserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


