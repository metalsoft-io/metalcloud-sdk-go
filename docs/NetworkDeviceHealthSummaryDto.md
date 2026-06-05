# NetworkDeviceHealthSummaryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastUpdated** | **string** | ISO 8601 timestamp of when the summary was last written | 
**OverallSeverity** | **string** | Overall severity — the device is flagged as having a problem at this level | 
**TrendDirection** | **string** | The network device health trend change over time: improving, stable, or degrading | 
**SuspectedRootCauses** | **[]string** | Suspected root causes identified across sessions | 
**KeyFindings** | **string** | Narrative summary of the key findings and patterns observed | 
**NotableEvents** | [**[]NetworkDeviceHealthSummaryNotableEventDto**](NetworkDeviceHealthSummaryNotableEventDto.md) | Significant events preserved across sessions | 
**DetectedIssues** | Pointer to [**[]NetworkDeviceDetectedIssueDto**](NetworkDeviceDetectedIssueDto.md) | Failures detected on the device | [optional] 

## Methods

### NewNetworkDeviceHealthSummaryDto

`func NewNetworkDeviceHealthSummaryDto(lastUpdated string, overallSeverity string, trendDirection string, suspectedRootCauses []string, keyFindings string, notableEvents []NetworkDeviceHealthSummaryNotableEventDto, ) *NetworkDeviceHealthSummaryDto`

NewNetworkDeviceHealthSummaryDto instantiates a new NetworkDeviceHealthSummaryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceHealthSummaryDtoWithDefaults

`func NewNetworkDeviceHealthSummaryDtoWithDefaults() *NetworkDeviceHealthSummaryDto`

NewNetworkDeviceHealthSummaryDtoWithDefaults instantiates a new NetworkDeviceHealthSummaryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastUpdated

`func (o *NetworkDeviceHealthSummaryDto) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *NetworkDeviceHealthSummaryDto) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *NetworkDeviceHealthSummaryDto) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.


### GetOverallSeverity

`func (o *NetworkDeviceHealthSummaryDto) GetOverallSeverity() string`

GetOverallSeverity returns the OverallSeverity field if non-nil, zero value otherwise.

### GetOverallSeverityOk

`func (o *NetworkDeviceHealthSummaryDto) GetOverallSeverityOk() (*string, bool)`

GetOverallSeverityOk returns a tuple with the OverallSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallSeverity

`func (o *NetworkDeviceHealthSummaryDto) SetOverallSeverity(v string)`

SetOverallSeverity sets OverallSeverity field to given value.


### GetTrendDirection

`func (o *NetworkDeviceHealthSummaryDto) GetTrendDirection() string`

GetTrendDirection returns the TrendDirection field if non-nil, zero value otherwise.

### GetTrendDirectionOk

`func (o *NetworkDeviceHealthSummaryDto) GetTrendDirectionOk() (*string, bool)`

GetTrendDirectionOk returns a tuple with the TrendDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrendDirection

`func (o *NetworkDeviceHealthSummaryDto) SetTrendDirection(v string)`

SetTrendDirection sets TrendDirection field to given value.


### GetSuspectedRootCauses

`func (o *NetworkDeviceHealthSummaryDto) GetSuspectedRootCauses() []string`

GetSuspectedRootCauses returns the SuspectedRootCauses field if non-nil, zero value otherwise.

### GetSuspectedRootCausesOk

`func (o *NetworkDeviceHealthSummaryDto) GetSuspectedRootCausesOk() (*[]string, bool)`

GetSuspectedRootCausesOk returns a tuple with the SuspectedRootCauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspectedRootCauses

`func (o *NetworkDeviceHealthSummaryDto) SetSuspectedRootCauses(v []string)`

SetSuspectedRootCauses sets SuspectedRootCauses field to given value.


### GetKeyFindings

`func (o *NetworkDeviceHealthSummaryDto) GetKeyFindings() string`

GetKeyFindings returns the KeyFindings field if non-nil, zero value otherwise.

### GetKeyFindingsOk

`func (o *NetworkDeviceHealthSummaryDto) GetKeyFindingsOk() (*string, bool)`

GetKeyFindingsOk returns a tuple with the KeyFindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyFindings

`func (o *NetworkDeviceHealthSummaryDto) SetKeyFindings(v string)`

SetKeyFindings sets KeyFindings field to given value.


### GetNotableEvents

`func (o *NetworkDeviceHealthSummaryDto) GetNotableEvents() []NetworkDeviceHealthSummaryNotableEventDto`

GetNotableEvents returns the NotableEvents field if non-nil, zero value otherwise.

### GetNotableEventsOk

`func (o *NetworkDeviceHealthSummaryDto) GetNotableEventsOk() (*[]NetworkDeviceHealthSummaryNotableEventDto, bool)`

GetNotableEventsOk returns a tuple with the NotableEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotableEvents

`func (o *NetworkDeviceHealthSummaryDto) SetNotableEvents(v []NetworkDeviceHealthSummaryNotableEventDto)`

SetNotableEvents sets NotableEvents field to given value.


### GetDetectedIssues

`func (o *NetworkDeviceHealthSummaryDto) GetDetectedIssues() []NetworkDeviceDetectedIssueDto`

GetDetectedIssues returns the DetectedIssues field if non-nil, zero value otherwise.

### GetDetectedIssuesOk

`func (o *NetworkDeviceHealthSummaryDto) GetDetectedIssuesOk() (*[]NetworkDeviceDetectedIssueDto, bool)`

GetDetectedIssuesOk returns a tuple with the DetectedIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedIssues

`func (o *NetworkDeviceHealthSummaryDto) SetDetectedIssues(v []NetworkDeviceDetectedIssueDto)`

SetDetectedIssues sets DetectedIssues field to given value.

### HasDetectedIssues

`func (o *NetworkDeviceHealthSummaryDto) HasDetectedIssues() bool`

HasDetectedIssues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


