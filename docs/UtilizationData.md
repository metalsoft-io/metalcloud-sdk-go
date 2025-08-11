# UtilizationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | **float32** |  | 
**MeasurementUnit** | **string** |  | 
**MeteredWaypoints** | Pointer to [**[]ResourceItem**](ResourceItem.md) | Collection of metered waypoints | [optional] 

## Methods

### NewUtilizationData

`func NewUtilizationData(quantity float32, measurementUnit string, ) *UtilizationData`

NewUtilizationData instantiates a new UtilizationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilizationDataWithDefaults

`func NewUtilizationDataWithDefaults() *UtilizationData`

NewUtilizationDataWithDefaults instantiates a new UtilizationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *UtilizationData) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UtilizationData) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UtilizationData) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetMeasurementUnit

`func (o *UtilizationData) GetMeasurementUnit() string`

GetMeasurementUnit returns the MeasurementUnit field if non-nil, zero value otherwise.

### GetMeasurementUnitOk

`func (o *UtilizationData) GetMeasurementUnitOk() (*string, bool)`

GetMeasurementUnitOk returns a tuple with the MeasurementUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurementUnit

`func (o *UtilizationData) SetMeasurementUnit(v string)`

SetMeasurementUnit sets MeasurementUnit field to given value.


### GetMeteredWaypoints

`func (o *UtilizationData) GetMeteredWaypoints() []ResourceItem`

GetMeteredWaypoints returns the MeteredWaypoints field if non-nil, zero value otherwise.

### GetMeteredWaypointsOk

`func (o *UtilizationData) GetMeteredWaypointsOk() (*[]ResourceItem, bool)`

GetMeteredWaypointsOk returns a tuple with the MeteredWaypoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeteredWaypoints

`func (o *UtilizationData) SetMeteredWaypoints(v []ResourceItem)`

SetMeteredWaypoints sets MeteredWaypoints field to given value.

### HasMeteredWaypoints

`func (o *UtilizationData) HasMeteredWaypoints() bool`

HasMeteredWaypoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


