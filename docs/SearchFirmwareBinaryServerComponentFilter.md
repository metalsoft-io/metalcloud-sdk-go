# SearchFirmwareBinaryServerComponentFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vendor** | **string** | Discriminator field set to hp | 
**Id** | **float32** | ID of the HP component filter | 
**ComponentId** | **string** | Component ID for Dell | 
**ServerBrand** | **string** | Server brand for Dell | 
**ServerModel** | **string** | Server model for Dell | 
**Submodel** | **string** | Submodel for Lenovo | 
**SerialNumber** | **string** | Serial number for Lenovo | 
**ComponentName** | **string** | Component name for Lenovo | 
**DeviceClass** | **string** | Device class for HP | 
**Targets** | **[]string** | Targets for HP | 

## Methods

### NewSearchFirmwareBinaryServerComponentFilter

`func NewSearchFirmwareBinaryServerComponentFilter(vendor string, id float32, componentId string, serverBrand string, serverModel string, submodel string, serialNumber string, componentName string, deviceClass string, targets []string, ) *SearchFirmwareBinaryServerComponentFilter`

NewSearchFirmwareBinaryServerComponentFilter instantiates a new SearchFirmwareBinaryServerComponentFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchFirmwareBinaryServerComponentFilterWithDefaults

`func NewSearchFirmwareBinaryServerComponentFilterWithDefaults() *SearchFirmwareBinaryServerComponentFilter`

NewSearchFirmwareBinaryServerComponentFilterWithDefaults instantiates a new SearchFirmwareBinaryServerComponentFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendor

`func (o *SearchFirmwareBinaryServerComponentFilter) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *SearchFirmwareBinaryServerComponentFilter) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *SearchFirmwareBinaryServerComponentFilter) SetVendor(v string)`

SetVendor sets Vendor field to given value.


### GetId

`func (o *SearchFirmwareBinaryServerComponentFilter) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchFirmwareBinaryServerComponentFilter) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchFirmwareBinaryServerComponentFilter) SetId(v float32)`

SetId sets Id field to given value.


### GetComponentId

`func (o *SearchFirmwareBinaryServerComponentFilter) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *SearchFirmwareBinaryServerComponentFilter) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *SearchFirmwareBinaryServerComponentFilter) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.


### GetServerBrand

`func (o *SearchFirmwareBinaryServerComponentFilter) GetServerBrand() string`

GetServerBrand returns the ServerBrand field if non-nil, zero value otherwise.

### GetServerBrandOk

`func (o *SearchFirmwareBinaryServerComponentFilter) GetServerBrandOk() (*string, bool)`

GetServerBrandOk returns a tuple with the ServerBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerBrand

`func (o *SearchFirmwareBinaryServerComponentFilter) SetServerBrand(v string)`

SetServerBrand sets ServerBrand field to given value.


### GetServerModel

`func (o *SearchFirmwareBinaryServerComponentFilter) GetServerModel() string`

GetServerModel returns the ServerModel field if non-nil, zero value otherwise.

### GetServerModelOk

`func (o *SearchFirmwareBinaryServerComponentFilter) GetServerModelOk() (*string, bool)`

GetServerModelOk returns a tuple with the ServerModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerModel

`func (o *SearchFirmwareBinaryServerComponentFilter) SetServerModel(v string)`

SetServerModel sets ServerModel field to given value.


### GetSubmodel

`func (o *SearchFirmwareBinaryServerComponentFilter) GetSubmodel() string`

GetSubmodel returns the Submodel field if non-nil, zero value otherwise.

### GetSubmodelOk

`func (o *SearchFirmwareBinaryServerComponentFilter) GetSubmodelOk() (*string, bool)`

GetSubmodelOk returns a tuple with the Submodel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmodel

`func (o *SearchFirmwareBinaryServerComponentFilter) SetSubmodel(v string)`

SetSubmodel sets Submodel field to given value.


### GetSerialNumber

`func (o *SearchFirmwareBinaryServerComponentFilter) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *SearchFirmwareBinaryServerComponentFilter) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *SearchFirmwareBinaryServerComponentFilter) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetComponentName

`func (o *SearchFirmwareBinaryServerComponentFilter) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *SearchFirmwareBinaryServerComponentFilter) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *SearchFirmwareBinaryServerComponentFilter) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.


### GetDeviceClass

`func (o *SearchFirmwareBinaryServerComponentFilter) GetDeviceClass() string`

GetDeviceClass returns the DeviceClass field if non-nil, zero value otherwise.

### GetDeviceClassOk

`func (o *SearchFirmwareBinaryServerComponentFilter) GetDeviceClassOk() (*string, bool)`

GetDeviceClassOk returns a tuple with the DeviceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClass

`func (o *SearchFirmwareBinaryServerComponentFilter) SetDeviceClass(v string)`

SetDeviceClass sets DeviceClass field to given value.


### GetTargets

`func (o *SearchFirmwareBinaryServerComponentFilter) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *SearchFirmwareBinaryServerComponentFilter) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *SearchFirmwareBinaryServerComponentFilter) SetTargets(v []string)`

SetTargets sets Targets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


