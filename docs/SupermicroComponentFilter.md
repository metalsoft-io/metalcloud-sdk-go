# SupermicroComponentFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vendor** | **string** | Discriminator field set to supermicro | 
**Id** | **int64** | ID of the Supermicro component filter | 
**ComponentId** | **string** | Component ID for Supermicro | 
**ServerModel** | **string** | Server model for Supermicro | 

## Methods

### NewSupermicroComponentFilter

`func NewSupermicroComponentFilter(vendor string, id int64, componentId string, serverModel string, ) *SupermicroComponentFilter`

NewSupermicroComponentFilter instantiates a new SupermicroComponentFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupermicroComponentFilterWithDefaults

`func NewSupermicroComponentFilterWithDefaults() *SupermicroComponentFilter`

NewSupermicroComponentFilterWithDefaults instantiates a new SupermicroComponentFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendor

`func (o *SupermicroComponentFilter) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *SupermicroComponentFilter) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *SupermicroComponentFilter) SetVendor(v string)`

SetVendor sets Vendor field to given value.


### GetId

`func (o *SupermicroComponentFilter) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupermicroComponentFilter) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupermicroComponentFilter) SetId(v int64)`

SetId sets Id field to given value.


### GetComponentId

`func (o *SupermicroComponentFilter) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *SupermicroComponentFilter) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *SupermicroComponentFilter) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.


### GetServerModel

`func (o *SupermicroComponentFilter) GetServerModel() string`

GetServerModel returns the ServerModel field if non-nil, zero value otherwise.

### GetServerModelOk

`func (o *SupermicroComponentFilter) GetServerModelOk() (*string, bool)`

GetServerModelOk returns a tuple with the ServerModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerModel

`func (o *SupermicroComponentFilter) SetServerModel(v string)`

SetServerModel sets ServerModel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


