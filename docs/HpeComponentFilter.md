# HpeComponentFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vendor** | **string** | Discriminator field set to hpe | 
**Id** | **float32** | ID of the HP component filter | 
**DeviceClass** | **string** | Device class for HP | 
**Targets** | **[]string** | Targets for HP | 

## Methods

### NewHpeComponentFilter

`func NewHpeComponentFilter(vendor string, id float32, deviceClass string, targets []string, ) *HpeComponentFilter`

NewHpeComponentFilter instantiates a new HpeComponentFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHpeComponentFilterWithDefaults

`func NewHpeComponentFilterWithDefaults() *HpeComponentFilter`

NewHpeComponentFilterWithDefaults instantiates a new HpeComponentFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendor

`func (o *HpeComponentFilter) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *HpeComponentFilter) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *HpeComponentFilter) SetVendor(v string)`

SetVendor sets Vendor field to given value.


### GetId

`func (o *HpeComponentFilter) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HpeComponentFilter) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HpeComponentFilter) SetId(v float32)`

SetId sets Id field to given value.


### GetDeviceClass

`func (o *HpeComponentFilter) GetDeviceClass() string`

GetDeviceClass returns the DeviceClass field if non-nil, zero value otherwise.

### GetDeviceClassOk

`func (o *HpeComponentFilter) GetDeviceClassOk() (*string, bool)`

GetDeviceClassOk returns a tuple with the DeviceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClass

`func (o *HpeComponentFilter) SetDeviceClass(v string)`

SetDeviceClass sets DeviceClass field to given value.


### GetTargets

`func (o *HpeComponentFilter) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *HpeComponentFilter) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *HpeComponentFilter) SetTargets(v []string)`

SetTargets sets Targets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


