# InfrastructuresStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InfrastructureServiceStatus** | [**InfrastructureServiceStatusDto**](InfrastructureServiceStatusDto.md) | Infrastructure service status counts | 
**InfrastructureDeployOngoingStatusCount** | [**InfrastructureDeployOngoingStatusCountDto**](InfrastructureDeployOngoingStatusCountDto.md) | Infrastructure deploy ongoing status counts | 
**InfrastructureCount** | **float32** | Total count of infrastructures | 
**OngoingInfrastructureCount** | **float32** | Count of ongoing infrastructures | 
**InfrastructureCountEmpty** | **float32** | Count of empty infrastructures | 
**InfrastructureCountEmptyWithSubnets** | **float32** | Count of empty infrastructures with subnets | 

## Methods

### NewInfrastructuresStatistics

`func NewInfrastructuresStatistics(infrastructureServiceStatus InfrastructureServiceStatusDto, infrastructureDeployOngoingStatusCount InfrastructureDeployOngoingStatusCountDto, infrastructureCount float32, ongoingInfrastructureCount float32, infrastructureCountEmpty float32, infrastructureCountEmptyWithSubnets float32, ) *InfrastructuresStatistics`

NewInfrastructuresStatistics instantiates a new InfrastructuresStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructuresStatisticsWithDefaults

`func NewInfrastructuresStatisticsWithDefaults() *InfrastructuresStatistics`

NewInfrastructuresStatisticsWithDefaults instantiates a new InfrastructuresStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfrastructureServiceStatus

`func (o *InfrastructuresStatistics) GetInfrastructureServiceStatus() InfrastructureServiceStatusDto`

GetInfrastructureServiceStatus returns the InfrastructureServiceStatus field if non-nil, zero value otherwise.

### GetInfrastructureServiceStatusOk

`func (o *InfrastructuresStatistics) GetInfrastructureServiceStatusOk() (*InfrastructureServiceStatusDto, bool)`

GetInfrastructureServiceStatusOk returns a tuple with the InfrastructureServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureServiceStatus

`func (o *InfrastructuresStatistics) SetInfrastructureServiceStatus(v InfrastructureServiceStatusDto)`

SetInfrastructureServiceStatus sets InfrastructureServiceStatus field to given value.


### GetInfrastructureDeployOngoingStatusCount

`func (o *InfrastructuresStatistics) GetInfrastructureDeployOngoingStatusCount() InfrastructureDeployOngoingStatusCountDto`

GetInfrastructureDeployOngoingStatusCount returns the InfrastructureDeployOngoingStatusCount field if non-nil, zero value otherwise.

### GetInfrastructureDeployOngoingStatusCountOk

`func (o *InfrastructuresStatistics) GetInfrastructureDeployOngoingStatusCountOk() (*InfrastructureDeployOngoingStatusCountDto, bool)`

GetInfrastructureDeployOngoingStatusCountOk returns a tuple with the InfrastructureDeployOngoingStatusCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureDeployOngoingStatusCount

`func (o *InfrastructuresStatistics) SetInfrastructureDeployOngoingStatusCount(v InfrastructureDeployOngoingStatusCountDto)`

SetInfrastructureDeployOngoingStatusCount sets InfrastructureDeployOngoingStatusCount field to given value.


### GetInfrastructureCount

`func (o *InfrastructuresStatistics) GetInfrastructureCount() float32`

GetInfrastructureCount returns the InfrastructureCount field if non-nil, zero value otherwise.

### GetInfrastructureCountOk

`func (o *InfrastructuresStatistics) GetInfrastructureCountOk() (*float32, bool)`

GetInfrastructureCountOk returns a tuple with the InfrastructureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureCount

`func (o *InfrastructuresStatistics) SetInfrastructureCount(v float32)`

SetInfrastructureCount sets InfrastructureCount field to given value.


### GetOngoingInfrastructureCount

`func (o *InfrastructuresStatistics) GetOngoingInfrastructureCount() float32`

GetOngoingInfrastructureCount returns the OngoingInfrastructureCount field if non-nil, zero value otherwise.

### GetOngoingInfrastructureCountOk

`func (o *InfrastructuresStatistics) GetOngoingInfrastructureCountOk() (*float32, bool)`

GetOngoingInfrastructureCountOk returns a tuple with the OngoingInfrastructureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOngoingInfrastructureCount

`func (o *InfrastructuresStatistics) SetOngoingInfrastructureCount(v float32)`

SetOngoingInfrastructureCount sets OngoingInfrastructureCount field to given value.


### GetInfrastructureCountEmpty

`func (o *InfrastructuresStatistics) GetInfrastructureCountEmpty() float32`

GetInfrastructureCountEmpty returns the InfrastructureCountEmpty field if non-nil, zero value otherwise.

### GetInfrastructureCountEmptyOk

`func (o *InfrastructuresStatistics) GetInfrastructureCountEmptyOk() (*float32, bool)`

GetInfrastructureCountEmptyOk returns a tuple with the InfrastructureCountEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureCountEmpty

`func (o *InfrastructuresStatistics) SetInfrastructureCountEmpty(v float32)`

SetInfrastructureCountEmpty sets InfrastructureCountEmpty field to given value.


### GetInfrastructureCountEmptyWithSubnets

`func (o *InfrastructuresStatistics) GetInfrastructureCountEmptyWithSubnets() float32`

GetInfrastructureCountEmptyWithSubnets returns the InfrastructureCountEmptyWithSubnets field if non-nil, zero value otherwise.

### GetInfrastructureCountEmptyWithSubnetsOk

`func (o *InfrastructuresStatistics) GetInfrastructureCountEmptyWithSubnetsOk() (*float32, bool)`

GetInfrastructureCountEmptyWithSubnetsOk returns a tuple with the InfrastructureCountEmptyWithSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureCountEmptyWithSubnets

`func (o *InfrastructuresStatistics) SetInfrastructureCountEmptyWithSubnets(v float32)`

SetInfrastructureCountEmptyWithSubnets sets InfrastructureCountEmptyWithSubnets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


