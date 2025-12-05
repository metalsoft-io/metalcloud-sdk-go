# ServerInstanceAllocatedServerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | The id of the server allocated to the instance. | 
**Vendor** | **string** | The vendor of the server allocated to the instance. | 
**Model** | Pointer to **string** | The model of the server allocated to the instance. | [optional] 
**VendorInfo** | Pointer to [**ServerVendorInfo**](ServerVendorInfo.md) | The vendor specific information of the server allocated to the instance. | [optional] 

## Methods

### NewServerInstanceAllocatedServerInfo

`func NewServerInstanceAllocatedServerInfo(id float32, vendor string, ) *ServerInstanceAllocatedServerInfo`

NewServerInstanceAllocatedServerInfo instantiates a new ServerInstanceAllocatedServerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceAllocatedServerInfoWithDefaults

`func NewServerInstanceAllocatedServerInfoWithDefaults() *ServerInstanceAllocatedServerInfo`

NewServerInstanceAllocatedServerInfoWithDefaults instantiates a new ServerInstanceAllocatedServerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerInstanceAllocatedServerInfo) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerInstanceAllocatedServerInfo) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerInstanceAllocatedServerInfo) SetId(v float32)`

SetId sets Id field to given value.


### GetVendor

`func (o *ServerInstanceAllocatedServerInfo) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ServerInstanceAllocatedServerInfo) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ServerInstanceAllocatedServerInfo) SetVendor(v string)`

SetVendor sets Vendor field to given value.


### GetModel

`func (o *ServerInstanceAllocatedServerInfo) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ServerInstanceAllocatedServerInfo) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ServerInstanceAllocatedServerInfo) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ServerInstanceAllocatedServerInfo) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetVendorInfo

`func (o *ServerInstanceAllocatedServerInfo) GetVendorInfo() ServerVendorInfo`

GetVendorInfo returns the VendorInfo field if non-nil, zero value otherwise.

### GetVendorInfoOk

`func (o *ServerInstanceAllocatedServerInfo) GetVendorInfoOk() (*ServerVendorInfo, bool)`

GetVendorInfoOk returns a tuple with the VendorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInfo

`func (o *ServerInstanceAllocatedServerInfo) SetVendorInfo(v ServerVendorInfo)`

SetVendorInfo sets VendorInfo field to given value.

### HasVendorInfo

`func (o *ServerInstanceAllocatedServerInfo) HasVendorInfo() bool`

HasVendorInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


