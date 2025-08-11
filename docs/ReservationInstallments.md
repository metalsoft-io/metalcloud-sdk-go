# ReservationInstallments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerTypes** | **[]map[string]interface{}** | List of server types | 
**Drives** | **[]map[string]interface{}** | List of drives | 
**Subnets** | **[]map[string]interface{}** | List of subnets | 

## Methods

### NewReservationInstallments

`func NewReservationInstallments(serverTypes []map[string]interface{}, drives []map[string]interface{}, subnets []map[string]interface{}, ) *ReservationInstallments`

NewReservationInstallments instantiates a new ReservationInstallments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReservationInstallmentsWithDefaults

`func NewReservationInstallmentsWithDefaults() *ReservationInstallments`

NewReservationInstallmentsWithDefaults instantiates a new ReservationInstallments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerTypes

`func (o *ReservationInstallments) GetServerTypes() []map[string]interface{}`

GetServerTypes returns the ServerTypes field if non-nil, zero value otherwise.

### GetServerTypesOk

`func (o *ReservationInstallments) GetServerTypesOk() (*[]map[string]interface{}, bool)`

GetServerTypesOk returns a tuple with the ServerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypes

`func (o *ReservationInstallments) SetServerTypes(v []map[string]interface{})`

SetServerTypes sets ServerTypes field to given value.


### GetDrives

`func (o *ReservationInstallments) GetDrives() []map[string]interface{}`

GetDrives returns the Drives field if non-nil, zero value otherwise.

### GetDrivesOk

`func (o *ReservationInstallments) GetDrivesOk() (*[]map[string]interface{}, bool)`

GetDrivesOk returns a tuple with the Drives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrives

`func (o *ReservationInstallments) SetDrives(v []map[string]interface{})`

SetDrives sets Drives field to given value.


### GetSubnets

`func (o *ReservationInstallments) GetSubnets() []map[string]interface{}`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *ReservationInstallments) GetSubnetsOk() (*[]map[string]interface{}, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *ReservationInstallments) SetSubnets(v []map[string]interface{})`

SetSubnets sets Subnets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


