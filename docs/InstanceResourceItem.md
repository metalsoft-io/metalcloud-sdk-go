# InstanceResourceItem

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
**ServerTypeId** | **float32** |  | 
**ServerId** | **float32** |  | 
**ServerTypeName** | **string** |  | 
**OperatingSystemType** | **string** |  | 
**OperatingSystemVersion** | **string** |  | 
**OperatingSystemDisplayName** | **string** |  | 
**OperatingSystemTemplateId** | **float32** |  | 
**OriginalStartTimestamp** | **string** |  | 

## Methods

### NewInstanceResourceItem

`func NewInstanceResourceItem(id float32, label string, startTimestamp string, endTimestamp string, measurementPeriod float32, measurementUnit string, quantity float32, serverTypeId float32, serverId float32, serverTypeName string, operatingSystemType string, operatingSystemVersion string, operatingSystemDisplayName string, operatingSystemTemplateId float32, originalStartTimestamp string, ) *InstanceResourceItem`

NewInstanceResourceItem instantiates a new InstanceResourceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceResourceItemWithDefaults

`func NewInstanceResourceItemWithDefaults() *InstanceResourceItem`

NewInstanceResourceItemWithDefaults instantiates a new InstanceResourceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceResourceItem) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceResourceItem) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceResourceItem) SetId(v float32)`

SetId sets Id field to given value.


### GetLabel

`func (o *InstanceResourceItem) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *InstanceResourceItem) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *InstanceResourceItem) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetStartTimestamp

`func (o *InstanceResourceItem) GetStartTimestamp() string`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *InstanceResourceItem) GetStartTimestampOk() (*string, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *InstanceResourceItem) SetStartTimestamp(v string)`

SetStartTimestamp sets StartTimestamp field to given value.


### GetEndTimestamp

`func (o *InstanceResourceItem) GetEndTimestamp() string`

GetEndTimestamp returns the EndTimestamp field if non-nil, zero value otherwise.

### GetEndTimestampOk

`func (o *InstanceResourceItem) GetEndTimestampOk() (*string, bool)`

GetEndTimestampOk returns a tuple with the EndTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimestamp

`func (o *InstanceResourceItem) SetEndTimestamp(v string)`

SetEndTimestamp sets EndTimestamp field to given value.


### GetMeasurementPeriod

`func (o *InstanceResourceItem) GetMeasurementPeriod() float32`

GetMeasurementPeriod returns the MeasurementPeriod field if non-nil, zero value otherwise.

### GetMeasurementPeriodOk

`func (o *InstanceResourceItem) GetMeasurementPeriodOk() (*float32, bool)`

GetMeasurementPeriodOk returns a tuple with the MeasurementPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurementPeriod

`func (o *InstanceResourceItem) SetMeasurementPeriod(v float32)`

SetMeasurementPeriod sets MeasurementPeriod field to given value.


### GetMeasurementUnit

`func (o *InstanceResourceItem) GetMeasurementUnit() string`

GetMeasurementUnit returns the MeasurementUnit field if non-nil, zero value otherwise.

### GetMeasurementUnitOk

`func (o *InstanceResourceItem) GetMeasurementUnitOk() (*string, bool)`

GetMeasurementUnitOk returns a tuple with the MeasurementUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurementUnit

`func (o *InstanceResourceItem) SetMeasurementUnit(v string)`

SetMeasurementUnit sets MeasurementUnit field to given value.


### GetQuantity

`func (o *InstanceResourceItem) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InstanceResourceItem) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InstanceResourceItem) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetTags

`func (o *InstanceResourceItem) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InstanceResourceItem) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InstanceResourceItem) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InstanceResourceItem) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetServerTypeId

`func (o *InstanceResourceItem) GetServerTypeId() float32`

GetServerTypeId returns the ServerTypeId field if non-nil, zero value otherwise.

### GetServerTypeIdOk

`func (o *InstanceResourceItem) GetServerTypeIdOk() (*float32, bool)`

GetServerTypeIdOk returns a tuple with the ServerTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeId

`func (o *InstanceResourceItem) SetServerTypeId(v float32)`

SetServerTypeId sets ServerTypeId field to given value.


### GetServerId

`func (o *InstanceResourceItem) GetServerId() float32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *InstanceResourceItem) GetServerIdOk() (*float32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *InstanceResourceItem) SetServerId(v float32)`

SetServerId sets ServerId field to given value.


### GetServerTypeName

`func (o *InstanceResourceItem) GetServerTypeName() string`

GetServerTypeName returns the ServerTypeName field if non-nil, zero value otherwise.

### GetServerTypeNameOk

`func (o *InstanceResourceItem) GetServerTypeNameOk() (*string, bool)`

GetServerTypeNameOk returns a tuple with the ServerTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeName

`func (o *InstanceResourceItem) SetServerTypeName(v string)`

SetServerTypeName sets ServerTypeName field to given value.


### GetOperatingSystemType

`func (o *InstanceResourceItem) GetOperatingSystemType() string`

GetOperatingSystemType returns the OperatingSystemType field if non-nil, zero value otherwise.

### GetOperatingSystemTypeOk

`func (o *InstanceResourceItem) GetOperatingSystemTypeOk() (*string, bool)`

GetOperatingSystemTypeOk returns a tuple with the OperatingSystemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemType

`func (o *InstanceResourceItem) SetOperatingSystemType(v string)`

SetOperatingSystemType sets OperatingSystemType field to given value.


### GetOperatingSystemVersion

`func (o *InstanceResourceItem) GetOperatingSystemVersion() string`

GetOperatingSystemVersion returns the OperatingSystemVersion field if non-nil, zero value otherwise.

### GetOperatingSystemVersionOk

`func (o *InstanceResourceItem) GetOperatingSystemVersionOk() (*string, bool)`

GetOperatingSystemVersionOk returns a tuple with the OperatingSystemVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemVersion

`func (o *InstanceResourceItem) SetOperatingSystemVersion(v string)`

SetOperatingSystemVersion sets OperatingSystemVersion field to given value.


### GetOperatingSystemDisplayName

`func (o *InstanceResourceItem) GetOperatingSystemDisplayName() string`

GetOperatingSystemDisplayName returns the OperatingSystemDisplayName field if non-nil, zero value otherwise.

### GetOperatingSystemDisplayNameOk

`func (o *InstanceResourceItem) GetOperatingSystemDisplayNameOk() (*string, bool)`

GetOperatingSystemDisplayNameOk returns a tuple with the OperatingSystemDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemDisplayName

`func (o *InstanceResourceItem) SetOperatingSystemDisplayName(v string)`

SetOperatingSystemDisplayName sets OperatingSystemDisplayName field to given value.


### GetOperatingSystemTemplateId

`func (o *InstanceResourceItem) GetOperatingSystemTemplateId() float32`

GetOperatingSystemTemplateId returns the OperatingSystemTemplateId field if non-nil, zero value otherwise.

### GetOperatingSystemTemplateIdOk

`func (o *InstanceResourceItem) GetOperatingSystemTemplateIdOk() (*float32, bool)`

GetOperatingSystemTemplateIdOk returns a tuple with the OperatingSystemTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemTemplateId

`func (o *InstanceResourceItem) SetOperatingSystemTemplateId(v float32)`

SetOperatingSystemTemplateId sets OperatingSystemTemplateId field to given value.


### GetOriginalStartTimestamp

`func (o *InstanceResourceItem) GetOriginalStartTimestamp() string`

GetOriginalStartTimestamp returns the OriginalStartTimestamp field if non-nil, zero value otherwise.

### GetOriginalStartTimestampOk

`func (o *InstanceResourceItem) GetOriginalStartTimestampOk() (*string, bool)`

GetOriginalStartTimestampOk returns a tuple with the OriginalStartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalStartTimestamp

`func (o *InstanceResourceItem) SetOriginalStartTimestamp(v string)`

SetOriginalStartTimestamp sets OriginalStartTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


