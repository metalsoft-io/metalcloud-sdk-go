# SearchFirmwareBinaryServerComponentFilterDataItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vendor** | **string** | Discriminator field set to supermicro | 
**Id** | **int64** | ID of the Supermicro component filter | 
**ComponentId** | **string** | Component ID for Supermicro | 
**ServerModel** | **string** | Server model for Supermicro | 
**Submodel** | **string** | Submodel for Lenovo | 
**SerialNumber** | **string** | Serial number for Lenovo | 
**ComponentName** | **string** | Component name for Lenovo | 
**DeviceClass** | **string** | Device class for HP | 
**Targets** | **[]string** | Targets for HP | 

## Methods

### NewSearchFirmwareBinaryServerComponentFilterDataItem

`func NewSearchFirmwareBinaryServerComponentFilterDataItem(vendor string, id int64, componentId string, serverModel string, submodel string, serialNumber string, componentName string, deviceClass string, targets []string, ) *SearchFirmwareBinaryServerComponentFilterDataItem`

NewSearchFirmwareBinaryServerComponentFilterDataItem instantiates a new SearchFirmwareBinaryServerComponentFilterDataItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchFirmwareBinaryServerComponentFilterDataItemWithDefaults

`func NewSearchFirmwareBinaryServerComponentFilterDataItemWithDefaults() *SearchFirmwareBinaryServerComponentFilterDataItem`

NewSearchFirmwareBinaryServerComponentFilterDataItemWithDefaults instantiates a new SearchFirmwareBinaryServerComponentFilterDataItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendor

`func (o *SearchFirmwareBinaryServerComponentFilterDataItem) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *SearchFirmwareBinaryServerComponentFilterDataItem) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *SearchFirmwareBinaryServerComponentFilterDataItem) SetVendor(v string)`

SetVendor sets Vendor field to given value.


### GetId

`func (o *SearchFirmwareBinaryServerComponentFilterDataItem) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchFirmwareBinaryServerComponentFilterDataItem) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchFirmwareBinaryServerComponentFilterDataItem) SetId(v int64)`

SetId sets Id field to given value.


### GetComponentId

`func (o *SearchFirmwareBinaryServerComponentFilterDataItem) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *SearchFirmwareBinaryServerComponentFilterDataItem) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *SearchFirmwareBinaryServerComponentFilterDataItem) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.


### GetServerModel

`func (o *SearchFirmwareBinaryServerComponentFilterDataItem) GetServerModel() string`

GetServerModel returns the ServerModel field if non-nil, zero value otherwise.

### GetServerModelOk

`func (o *SearchFirmwareBinaryServerComponentFilterDataItem) GetServerModelOk() (*string, bool)`

GetServerModelOk returns a tuple with the ServerModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerModel

`func (o *SearchFirmwareBinaryServerComponentFilterDataItem) SetServerModel(v string)`

SetServerModel sets ServerModel field to given value.


### GetSubmodel

`func (o *SearchFirmwareBinaryServerComponentFilterDataItem) GetSubmodel() string`

GetSubmodel returns the Submodel field if non-nil, zero value otherwise.

### GetSubmodelOk

`func (o *SearchFirmwareBinaryServerComponentFilterDataItem) GetSubmodelOk() (*string, bool)`

GetSubmodelOk returns a tuple with the Submodel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmodel

`func (o *SearchFirmwareBinaryServerComponentFilterDataItem) SetSubmodel(v string)`

SetSubmodel sets Submodel field to given value.


### GetSerialNumber

`func (o *SearchFirmwareBinaryServerComponentFilterDataItem) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *SearchFirmwareBinaryServerComponentFilterDataItem) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *SearchFirmwareBinaryServerComponentFilterDataItem) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetComponentName

`func (o *SearchFirmwareBinaryServerComponentFilterDataItem) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *SearchFirmwareBinaryServerComponentFilterDataItem) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *SearchFirmwareBinaryServerComponentFilterDataItem) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.


### GetDeviceClass

`func (o *SearchFirmwareBinaryServerComponentFilterDataItem) GetDeviceClass() string`

GetDeviceClass returns the DeviceClass field if non-nil, zero value otherwise.

### GetDeviceClassOk

`func (o *SearchFirmwareBinaryServerComponentFilterDataItem) GetDeviceClassOk() (*string, bool)`

GetDeviceClassOk returns a tuple with the DeviceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClass

`func (o *SearchFirmwareBinaryServerComponentFilterDataItem) SetDeviceClass(v string)`

SetDeviceClass sets DeviceClass field to given value.


### GetTargets

`func (o *SearchFirmwareBinaryServerComponentFilterDataItem) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *SearchFirmwareBinaryServerComponentFilterDataItem) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *SearchFirmwareBinaryServerComponentFilterDataItem) SetTargets(v []string)`

SetTargets sets Targets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


