# InfrastructureResourceUtilizationSummaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTimestamp** | **string** | Start timestamp of the utilization period | 
**EndTimestamp** | **string** | End timestamp of the utilization period | 
**Infrastructures** | [**map[string]InfrastructureItem**](InfrastructureItem.md) | Infrastructures information | 
**ResourceUtilization** | [**map[string]UtilizationData**](UtilizationData.md) | Resource utilization summary | 
**Internet** | [**NetworkEntry**](NetworkEntry.md) | Internet usage data | 
**ReservationInstallments** | Pointer to [**ReservationInstallments**](ReservationInstallments.md) | Reservation installments information | [optional] 
**LicenseInstallments** | Pointer to [**LicenseInstallments**](LicenseInstallments.md) | License installments information | [optional] 

## Methods

### NewInfrastructureResourceUtilizationSummaryResponse

`func NewInfrastructureResourceUtilizationSummaryResponse(startTimestamp string, endTimestamp string, infrastructures map[string]InfrastructureItem, resourceUtilization map[string]UtilizationData, internet NetworkEntry, ) *InfrastructureResourceUtilizationSummaryResponse`

NewInfrastructureResourceUtilizationSummaryResponse instantiates a new InfrastructureResourceUtilizationSummaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureResourceUtilizationSummaryResponseWithDefaults

`func NewInfrastructureResourceUtilizationSummaryResponseWithDefaults() *InfrastructureResourceUtilizationSummaryResponse`

NewInfrastructureResourceUtilizationSummaryResponseWithDefaults instantiates a new InfrastructureResourceUtilizationSummaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTimestamp

`func (o *InfrastructureResourceUtilizationSummaryResponse) GetStartTimestamp() string`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *InfrastructureResourceUtilizationSummaryResponse) GetStartTimestampOk() (*string, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *InfrastructureResourceUtilizationSummaryResponse) SetStartTimestamp(v string)`

SetStartTimestamp sets StartTimestamp field to given value.


### GetEndTimestamp

`func (o *InfrastructureResourceUtilizationSummaryResponse) GetEndTimestamp() string`

GetEndTimestamp returns the EndTimestamp field if non-nil, zero value otherwise.

### GetEndTimestampOk

`func (o *InfrastructureResourceUtilizationSummaryResponse) GetEndTimestampOk() (*string, bool)`

GetEndTimestampOk returns a tuple with the EndTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimestamp

`func (o *InfrastructureResourceUtilizationSummaryResponse) SetEndTimestamp(v string)`

SetEndTimestamp sets EndTimestamp field to given value.


### GetInfrastructures

`func (o *InfrastructureResourceUtilizationSummaryResponse) GetInfrastructures() map[string]InfrastructureItem`

GetInfrastructures returns the Infrastructures field if non-nil, zero value otherwise.

### GetInfrastructuresOk

`func (o *InfrastructureResourceUtilizationSummaryResponse) GetInfrastructuresOk() (*map[string]InfrastructureItem, bool)`

GetInfrastructuresOk returns a tuple with the Infrastructures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructures

`func (o *InfrastructureResourceUtilizationSummaryResponse) SetInfrastructures(v map[string]InfrastructureItem)`

SetInfrastructures sets Infrastructures field to given value.


### GetResourceUtilization

`func (o *InfrastructureResourceUtilizationSummaryResponse) GetResourceUtilization() map[string]UtilizationData`

GetResourceUtilization returns the ResourceUtilization field if non-nil, zero value otherwise.

### GetResourceUtilizationOk

`func (o *InfrastructureResourceUtilizationSummaryResponse) GetResourceUtilizationOk() (*map[string]UtilizationData, bool)`

GetResourceUtilizationOk returns a tuple with the ResourceUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUtilization

`func (o *InfrastructureResourceUtilizationSummaryResponse) SetResourceUtilization(v map[string]UtilizationData)`

SetResourceUtilization sets ResourceUtilization field to given value.


### GetInternet

`func (o *InfrastructureResourceUtilizationSummaryResponse) GetInternet() NetworkEntry`

GetInternet returns the Internet field if non-nil, zero value otherwise.

### GetInternetOk

`func (o *InfrastructureResourceUtilizationSummaryResponse) GetInternetOk() (*NetworkEntry, bool)`

GetInternetOk returns a tuple with the Internet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternet

`func (o *InfrastructureResourceUtilizationSummaryResponse) SetInternet(v NetworkEntry)`

SetInternet sets Internet field to given value.


### GetReservationInstallments

`func (o *InfrastructureResourceUtilizationSummaryResponse) GetReservationInstallments() ReservationInstallments`

GetReservationInstallments returns the ReservationInstallments field if non-nil, zero value otherwise.

### GetReservationInstallmentsOk

`func (o *InfrastructureResourceUtilizationSummaryResponse) GetReservationInstallmentsOk() (*ReservationInstallments, bool)`

GetReservationInstallmentsOk returns a tuple with the ReservationInstallments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationInstallments

`func (o *InfrastructureResourceUtilizationSummaryResponse) SetReservationInstallments(v ReservationInstallments)`

SetReservationInstallments sets ReservationInstallments field to given value.

### HasReservationInstallments

`func (o *InfrastructureResourceUtilizationSummaryResponse) HasReservationInstallments() bool`

HasReservationInstallments returns a boolean if a field has been set.

### GetLicenseInstallments

`func (o *InfrastructureResourceUtilizationSummaryResponse) GetLicenseInstallments() LicenseInstallments`

GetLicenseInstallments returns the LicenseInstallments field if non-nil, zero value otherwise.

### GetLicenseInstallmentsOk

`func (o *InfrastructureResourceUtilizationSummaryResponse) GetLicenseInstallmentsOk() (*LicenseInstallments, bool)`

GetLicenseInstallmentsOk returns a tuple with the LicenseInstallments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseInstallments

`func (o *InfrastructureResourceUtilizationSummaryResponse) SetLicenseInstallments(v LicenseInstallments)`

SetLicenseInstallments sets LicenseInstallments field to given value.

### HasLicenseInstallments

`func (o *InfrastructureResourceUtilizationSummaryResponse) HasLicenseInstallments() bool`

HasLicenseInstallments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


