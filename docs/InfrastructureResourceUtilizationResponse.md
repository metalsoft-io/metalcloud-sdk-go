# InfrastructureResourceUtilizationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetailedReport** | **map[string]interface{}** | Detailed report of resource utilization | 
**NetworkReport** | Pointer to **map[string]interface{}** | Network utilization report | [optional] 
**ReservationInstallments** | Pointer to **map[string]interface{}** | Reservation installments information | [optional] 
**Infrastructures** | **map[string]interface{}** | Infrastructures information | 
**StartTimestamp** | **string** | Start timestamp of the utilization period | 
**EndTimestamp** | **string** | End timestamp of the utilization period | 
**LicenseInstallments** | Pointer to [**LicenseInstallments**](LicenseInstallments.md) | License installments information | [optional] 

## Methods

### NewInfrastructureResourceUtilizationResponse

`func NewInfrastructureResourceUtilizationResponse(detailedReport map[string]interface{}, infrastructures map[string]interface{}, startTimestamp string, endTimestamp string, ) *InfrastructureResourceUtilizationResponse`

NewInfrastructureResourceUtilizationResponse instantiates a new InfrastructureResourceUtilizationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureResourceUtilizationResponseWithDefaults

`func NewInfrastructureResourceUtilizationResponseWithDefaults() *InfrastructureResourceUtilizationResponse`

NewInfrastructureResourceUtilizationResponseWithDefaults instantiates a new InfrastructureResourceUtilizationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetailedReport

`func (o *InfrastructureResourceUtilizationResponse) GetDetailedReport() map[string]interface{}`

GetDetailedReport returns the DetailedReport field if non-nil, zero value otherwise.

### GetDetailedReportOk

`func (o *InfrastructureResourceUtilizationResponse) GetDetailedReportOk() (*map[string]interface{}, bool)`

GetDetailedReportOk returns a tuple with the DetailedReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedReport

`func (o *InfrastructureResourceUtilizationResponse) SetDetailedReport(v map[string]interface{})`

SetDetailedReport sets DetailedReport field to given value.


### GetNetworkReport

`func (o *InfrastructureResourceUtilizationResponse) GetNetworkReport() map[string]interface{}`

GetNetworkReport returns the NetworkReport field if non-nil, zero value otherwise.

### GetNetworkReportOk

`func (o *InfrastructureResourceUtilizationResponse) GetNetworkReportOk() (*map[string]interface{}, bool)`

GetNetworkReportOk returns a tuple with the NetworkReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkReport

`func (o *InfrastructureResourceUtilizationResponse) SetNetworkReport(v map[string]interface{})`

SetNetworkReport sets NetworkReport field to given value.

### HasNetworkReport

`func (o *InfrastructureResourceUtilizationResponse) HasNetworkReport() bool`

HasNetworkReport returns a boolean if a field has been set.

### GetReservationInstallments

`func (o *InfrastructureResourceUtilizationResponse) GetReservationInstallments() map[string]interface{}`

GetReservationInstallments returns the ReservationInstallments field if non-nil, zero value otherwise.

### GetReservationInstallmentsOk

`func (o *InfrastructureResourceUtilizationResponse) GetReservationInstallmentsOk() (*map[string]interface{}, bool)`

GetReservationInstallmentsOk returns a tuple with the ReservationInstallments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationInstallments

`func (o *InfrastructureResourceUtilizationResponse) SetReservationInstallments(v map[string]interface{})`

SetReservationInstallments sets ReservationInstallments field to given value.

### HasReservationInstallments

`func (o *InfrastructureResourceUtilizationResponse) HasReservationInstallments() bool`

HasReservationInstallments returns a boolean if a field has been set.

### GetInfrastructures

`func (o *InfrastructureResourceUtilizationResponse) GetInfrastructures() map[string]interface{}`

GetInfrastructures returns the Infrastructures field if non-nil, zero value otherwise.

### GetInfrastructuresOk

`func (o *InfrastructureResourceUtilizationResponse) GetInfrastructuresOk() (*map[string]interface{}, bool)`

GetInfrastructuresOk returns a tuple with the Infrastructures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructures

`func (o *InfrastructureResourceUtilizationResponse) SetInfrastructures(v map[string]interface{})`

SetInfrastructures sets Infrastructures field to given value.


### GetStartTimestamp

`func (o *InfrastructureResourceUtilizationResponse) GetStartTimestamp() string`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *InfrastructureResourceUtilizationResponse) GetStartTimestampOk() (*string, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *InfrastructureResourceUtilizationResponse) SetStartTimestamp(v string)`

SetStartTimestamp sets StartTimestamp field to given value.


### GetEndTimestamp

`func (o *InfrastructureResourceUtilizationResponse) GetEndTimestamp() string`

GetEndTimestamp returns the EndTimestamp field if non-nil, zero value otherwise.

### GetEndTimestampOk

`func (o *InfrastructureResourceUtilizationResponse) GetEndTimestampOk() (*string, bool)`

GetEndTimestampOk returns a tuple with the EndTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimestamp

`func (o *InfrastructureResourceUtilizationResponse) SetEndTimestamp(v string)`

SetEndTimestamp sets EndTimestamp field to given value.


### GetLicenseInstallments

`func (o *InfrastructureResourceUtilizationResponse) GetLicenseInstallments() LicenseInstallments`

GetLicenseInstallments returns the LicenseInstallments field if non-nil, zero value otherwise.

### GetLicenseInstallmentsOk

`func (o *InfrastructureResourceUtilizationResponse) GetLicenseInstallmentsOk() (*LicenseInstallments, bool)`

GetLicenseInstallmentsOk returns a tuple with the LicenseInstallments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseInstallments

`func (o *InfrastructureResourceUtilizationResponse) SetLicenseInstallments(v LicenseInstallments)`

SetLicenseInstallments sets LicenseInstallments field to given value.

### HasLicenseInstallments

`func (o *InfrastructureResourceUtilizationResponse) HasLicenseInstallments() bool`

HasLicenseInstallments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


