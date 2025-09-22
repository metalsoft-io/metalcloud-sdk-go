# DriveResourceItem

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
**DriveSizeMbytes** | **float32** |  | 
**DriveStorageType** | **string** |  | 

## Methods

### NewDriveResourceItem

`func NewDriveResourceItem(id float32, label string, startTimestamp string, endTimestamp string, measurementPeriod float32, measurementUnit string, quantity float32, driveSizeMbytes float32, driveStorageType string, ) *DriveResourceItem`

NewDriveResourceItem instantiates a new DriveResourceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriveResourceItemWithDefaults

`func NewDriveResourceItemWithDefaults() *DriveResourceItem`

NewDriveResourceItemWithDefaults instantiates a new DriveResourceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DriveResourceItem) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DriveResourceItem) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DriveResourceItem) SetId(v float32)`

SetId sets Id field to given value.


### GetLabel

`func (o *DriveResourceItem) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DriveResourceItem) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DriveResourceItem) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetStartTimestamp

`func (o *DriveResourceItem) GetStartTimestamp() string`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *DriveResourceItem) GetStartTimestampOk() (*string, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *DriveResourceItem) SetStartTimestamp(v string)`

SetStartTimestamp sets StartTimestamp field to given value.


### GetEndTimestamp

`func (o *DriveResourceItem) GetEndTimestamp() string`

GetEndTimestamp returns the EndTimestamp field if non-nil, zero value otherwise.

### GetEndTimestampOk

`func (o *DriveResourceItem) GetEndTimestampOk() (*string, bool)`

GetEndTimestampOk returns a tuple with the EndTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimestamp

`func (o *DriveResourceItem) SetEndTimestamp(v string)`

SetEndTimestamp sets EndTimestamp field to given value.


### GetMeasurementPeriod

`func (o *DriveResourceItem) GetMeasurementPeriod() float32`

GetMeasurementPeriod returns the MeasurementPeriod field if non-nil, zero value otherwise.

### GetMeasurementPeriodOk

`func (o *DriveResourceItem) GetMeasurementPeriodOk() (*float32, bool)`

GetMeasurementPeriodOk returns a tuple with the MeasurementPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurementPeriod

`func (o *DriveResourceItem) SetMeasurementPeriod(v float32)`

SetMeasurementPeriod sets MeasurementPeriod field to given value.


### GetMeasurementUnit

`func (o *DriveResourceItem) GetMeasurementUnit() string`

GetMeasurementUnit returns the MeasurementUnit field if non-nil, zero value otherwise.

### GetMeasurementUnitOk

`func (o *DriveResourceItem) GetMeasurementUnitOk() (*string, bool)`

GetMeasurementUnitOk returns a tuple with the MeasurementUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurementUnit

`func (o *DriveResourceItem) SetMeasurementUnit(v string)`

SetMeasurementUnit sets MeasurementUnit field to given value.


### GetQuantity

`func (o *DriveResourceItem) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *DriveResourceItem) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *DriveResourceItem) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetTags

`func (o *DriveResourceItem) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DriveResourceItem) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DriveResourceItem) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DriveResourceItem) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDriveSizeMbytes

`func (o *DriveResourceItem) GetDriveSizeMbytes() float32`

GetDriveSizeMbytes returns the DriveSizeMbytes field if non-nil, zero value otherwise.

### GetDriveSizeMbytesOk

`func (o *DriveResourceItem) GetDriveSizeMbytesOk() (*float32, bool)`

GetDriveSizeMbytesOk returns a tuple with the DriveSizeMbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveSizeMbytes

`func (o *DriveResourceItem) SetDriveSizeMbytes(v float32)`

SetDriveSizeMbytes sets DriveSizeMbytes field to given value.


### GetDriveStorageType

`func (o *DriveResourceItem) GetDriveStorageType() string`

GetDriveStorageType returns the DriveStorageType field if non-nil, zero value otherwise.

### GetDriveStorageTypeOk

`func (o *DriveResourceItem) GetDriveStorageTypeOk() (*string, bool)`

GetDriveStorageTypeOk returns a tuple with the DriveStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveStorageType

`func (o *DriveResourceItem) SetDriveStorageType(v string)`

SetDriveStorageType sets DriveStorageType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


