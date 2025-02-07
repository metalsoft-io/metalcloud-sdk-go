# ServerDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | The id of the disk. | 
**ServerId** | **float32** | The id of the server. | 
**Model** | **string** | The model of the disk | 
**DiskSizeGb** | **float32** | The size of the disk in GB | 
**Serial** | **string** | The serial of the disk | 
**Vendor** | **string** | The vendor of the disk | 
**Status** | **string** | The status of the disk | 
**Type** | **string** | The type of the disk | 
**ServerStorageControllerId** | Pointer to **float32** | The id of the storage controller | [optional] 

## Methods

### NewServerDisk

`func NewServerDisk(id float32, serverId float32, model string, diskSizeGb float32, serial string, vendor string, status string, type_ string, ) *ServerDisk`

NewServerDisk instantiates a new ServerDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerDiskWithDefaults

`func NewServerDiskWithDefaults() *ServerDisk`

NewServerDiskWithDefaults instantiates a new ServerDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerDisk) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerDisk) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerDisk) SetId(v float32)`

SetId sets Id field to given value.


### GetServerId

`func (o *ServerDisk) GetServerId() float32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ServerDisk) GetServerIdOk() (*float32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ServerDisk) SetServerId(v float32)`

SetServerId sets ServerId field to given value.


### GetModel

`func (o *ServerDisk) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ServerDisk) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ServerDisk) SetModel(v string)`

SetModel sets Model field to given value.


### GetDiskSizeGb

`func (o *ServerDisk) GetDiskSizeGb() float32`

GetDiskSizeGb returns the DiskSizeGb field if non-nil, zero value otherwise.

### GetDiskSizeGbOk

`func (o *ServerDisk) GetDiskSizeGbOk() (*float32, bool)`

GetDiskSizeGbOk returns a tuple with the DiskSizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeGb

`func (o *ServerDisk) SetDiskSizeGb(v float32)`

SetDiskSizeGb sets DiskSizeGb field to given value.


### GetSerial

`func (o *ServerDisk) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *ServerDisk) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *ServerDisk) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetVendor

`func (o *ServerDisk) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ServerDisk) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ServerDisk) SetVendor(v string)`

SetVendor sets Vendor field to given value.


### GetStatus

`func (o *ServerDisk) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServerDisk) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServerDisk) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *ServerDisk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServerDisk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServerDisk) SetType(v string)`

SetType sets Type field to given value.


### GetServerStorageControllerId

`func (o *ServerDisk) GetServerStorageControllerId() float32`

GetServerStorageControllerId returns the ServerStorageControllerId field if non-nil, zero value otherwise.

### GetServerStorageControllerIdOk

`func (o *ServerDisk) GetServerStorageControllerIdOk() (*float32, bool)`

GetServerStorageControllerIdOk returns a tuple with the ServerStorageControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStorageControllerId

`func (o *ServerDisk) SetServerStorageControllerId(v float32)`

SetServerStorageControllerId sets ServerStorageControllerId field to given value.

### HasServerStorageControllerId

`func (o *ServerDisk) HasServerStorageControllerId() bool`

HasServerStorageControllerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


