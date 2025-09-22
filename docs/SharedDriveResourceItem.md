# SharedDriveResourceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**Label** | **string** |  | 
**StartTimestamp** | **string** |  | 
**EndTimestamp** | **string** |  | 
**MeasurementPeriod** | **float32** |  | 
**MeasurementUnit** | **string** |  | 
**Quantity** | **float32** |  | 
**Tags** | Pointer to **string** |  | [optional] 
**SharedDriveSizeMbytes** | **float32** |  | 
**SharedDriveStorageType** | **string** |  | 

## Methods

### NewSharedDriveResourceItem

`func NewSharedDriveResourceItem(id float32, label string, startTimestamp string, endTimestamp string, measurementPeriod float32, measurementUnit string, quantity float32, sharedDriveSizeMbytes float32, sharedDriveStorageType string, ) *SharedDriveResourceItem`

NewSharedDriveResourceItem instantiates a new SharedDriveResourceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedDriveResourceItemWithDefaults

`func NewSharedDriveResourceItemWithDefaults() *SharedDriveResourceItem`

NewSharedDriveResourceItemWithDefaults instantiates a new SharedDriveResourceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SharedDriveResourceItem) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SharedDriveResourceItem) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SharedDriveResourceItem) SetId(v float32)`

SetId sets Id field to given value.


### GetLabel

`func (o *SharedDriveResourceItem) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SharedDriveResourceItem) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SharedDriveResourceItem) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetStartTimestamp

`func (o *SharedDriveResourceItem) GetStartTimestamp() string`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *SharedDriveResourceItem) GetStartTimestampOk() (*string, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *SharedDriveResourceItem) SetStartTimestamp(v string)`

SetStartTimestamp sets StartTimestamp field to given value.


### GetEndTimestamp

`func (o *SharedDriveResourceItem) GetEndTimestamp() string`

GetEndTimestamp returns the EndTimestamp field if non-nil, zero value otherwise.

### GetEndTimestampOk

`func (o *SharedDriveResourceItem) GetEndTimestampOk() (*string, bool)`

GetEndTimestampOk returns a tuple with the EndTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimestamp

`func (o *SharedDriveResourceItem) SetEndTimestamp(v string)`

SetEndTimestamp sets EndTimestamp field to given value.


### GetMeasurementPeriod

`func (o *SharedDriveResourceItem) GetMeasurementPeriod() float32`

GetMeasurementPeriod returns the MeasurementPeriod field if non-nil, zero value otherwise.

### GetMeasurementPeriodOk

`func (o *SharedDriveResourceItem) GetMeasurementPeriodOk() (*float32, bool)`

GetMeasurementPeriodOk returns a tuple with the MeasurementPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurementPeriod

`func (o *SharedDriveResourceItem) SetMeasurementPeriod(v float32)`

SetMeasurementPeriod sets MeasurementPeriod field to given value.


### GetMeasurementUnit

`func (o *SharedDriveResourceItem) GetMeasurementUnit() string`

GetMeasurementUnit returns the MeasurementUnit field if non-nil, zero value otherwise.

### GetMeasurementUnitOk

`func (o *SharedDriveResourceItem) GetMeasurementUnitOk() (*string, bool)`

GetMeasurementUnitOk returns a tuple with the MeasurementUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurementUnit

`func (o *SharedDriveResourceItem) SetMeasurementUnit(v string)`

SetMeasurementUnit sets MeasurementUnit field to given value.


### GetQuantity

`func (o *SharedDriveResourceItem) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SharedDriveResourceItem) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SharedDriveResourceItem) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetTags

`func (o *SharedDriveResourceItem) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SharedDriveResourceItem) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SharedDriveResourceItem) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SharedDriveResourceItem) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSharedDriveSizeMbytes

`func (o *SharedDriveResourceItem) GetSharedDriveSizeMbytes() float32`

GetSharedDriveSizeMbytes returns the SharedDriveSizeMbytes field if non-nil, zero value otherwise.

### GetSharedDriveSizeMbytesOk

`func (o *SharedDriveResourceItem) GetSharedDriveSizeMbytesOk() (*float32, bool)`

GetSharedDriveSizeMbytesOk returns a tuple with the SharedDriveSizeMbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedDriveSizeMbytes

`func (o *SharedDriveResourceItem) SetSharedDriveSizeMbytes(v float32)`

SetSharedDriveSizeMbytes sets SharedDriveSizeMbytes field to given value.


### GetSharedDriveStorageType

`func (o *SharedDriveResourceItem) GetSharedDriveStorageType() string`

GetSharedDriveStorageType returns the SharedDriveStorageType field if non-nil, zero value otherwise.

### GetSharedDriveStorageTypeOk

`func (o *SharedDriveResourceItem) GetSharedDriveStorageTypeOk() (*string, bool)`

GetSharedDriveStorageTypeOk returns a tuple with the SharedDriveStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedDriveStorageType

`func (o *SharedDriveResourceItem) SetSharedDriveStorageType(v string)`

SetSharedDriveStorageType sets SharedDriveStorageType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


