# DellComponentFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vendor** | **string** | Discriminator field set to dell | 
**Id** | **float32** | ID of the Dell component filter | 
**ComponentId** | **string** | Component ID for Dell | 
**ServerModel** | **string** | Server model for Dell | 

## Methods

### NewDellComponentFilter

`func NewDellComponentFilter(vendor string, id float32, componentId string, serverModel string, ) *DellComponentFilter`

NewDellComponentFilter instantiates a new DellComponentFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDellComponentFilterWithDefaults

`func NewDellComponentFilterWithDefaults() *DellComponentFilter`

NewDellComponentFilterWithDefaults instantiates a new DellComponentFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendor

`func (o *DellComponentFilter) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *DellComponentFilter) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *DellComponentFilter) SetVendor(v string)`

SetVendor sets Vendor field to given value.


### GetId

`func (o *DellComponentFilter) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DellComponentFilter) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DellComponentFilter) SetId(v float32)`

SetId sets Id field to given value.


### GetComponentId

`func (o *DellComponentFilter) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *DellComponentFilter) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *DellComponentFilter) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.


### GetServerModel

`func (o *DellComponentFilter) GetServerModel() string`

GetServerModel returns the ServerModel field if non-nil, zero value otherwise.

### GetServerModelOk

`func (o *DellComponentFilter) GetServerModelOk() (*string, bool)`

GetServerModelOk returns a tuple with the ServerModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerModel

`func (o *DellComponentFilter) SetServerModel(v string)`

SetServerModel sets ServerModel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


