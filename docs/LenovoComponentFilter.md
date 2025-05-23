# LenovoComponentFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vendor** | **string** | Discriminator field set to lenovo | 
**Id** | **float32** | ID of the Lenovo component filter | 
**Submodel** | **string** | Submodel for Lenovo | 
**SerialNumber** | **string** | Serial number for Lenovo | 
**ComponentName** | **string** | Component name for Lenovo | 

## Methods

### NewLenovoComponentFilter

`func NewLenovoComponentFilter(vendor string, id float32, submodel string, serialNumber string, componentName string, ) *LenovoComponentFilter`

NewLenovoComponentFilter instantiates a new LenovoComponentFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLenovoComponentFilterWithDefaults

`func NewLenovoComponentFilterWithDefaults() *LenovoComponentFilter`

NewLenovoComponentFilterWithDefaults instantiates a new LenovoComponentFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendor

`func (o *LenovoComponentFilter) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *LenovoComponentFilter) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *LenovoComponentFilter) SetVendor(v string)`

SetVendor sets Vendor field to given value.


### GetId

`func (o *LenovoComponentFilter) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LenovoComponentFilter) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LenovoComponentFilter) SetId(v float32)`

SetId sets Id field to given value.


### GetSubmodel

`func (o *LenovoComponentFilter) GetSubmodel() string`

GetSubmodel returns the Submodel field if non-nil, zero value otherwise.

### GetSubmodelOk

`func (o *LenovoComponentFilter) GetSubmodelOk() (*string, bool)`

GetSubmodelOk returns a tuple with the Submodel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmodel

`func (o *LenovoComponentFilter) SetSubmodel(v string)`

SetSubmodel sets Submodel field to given value.


### GetSerialNumber

`func (o *LenovoComponentFilter) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *LenovoComponentFilter) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *LenovoComponentFilter) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetComponentName

`func (o *LenovoComponentFilter) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *LenovoComponentFilter) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *LenovoComponentFilter) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


