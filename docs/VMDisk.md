# VMDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | VM ID | 
**VmId** | **float32** | VM ID | 
**Datastore** | **string** | Name of the VM Disk datastore | 
**SizeGB** | **float32** | Size of the VM Disk in GB | 

## Methods

### NewVMDisk

`func NewVMDisk(id float32, vmId float32, datastore string, sizeGB float32, ) *VMDisk`

NewVMDisk instantiates a new VMDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMDiskWithDefaults

`func NewVMDiskWithDefaults() *VMDisk`

NewVMDiskWithDefaults instantiates a new VMDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VMDisk) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VMDisk) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VMDisk) SetId(v float32)`

SetId sets Id field to given value.


### GetVmId

`func (o *VMDisk) GetVmId() float32`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *VMDisk) GetVmIdOk() (*float32, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *VMDisk) SetVmId(v float32)`

SetVmId sets VmId field to given value.


### GetDatastore

`func (o *VMDisk) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *VMDisk) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *VMDisk) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.


### GetSizeGB

`func (o *VMDisk) GetSizeGB() float32`

GetSizeGB returns the SizeGB field if non-nil, zero value otherwise.

### GetSizeGBOk

`func (o *VMDisk) GetSizeGBOk() (*float32, bool)`

GetSizeGBOk returns a tuple with the SizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGB

`func (o *VMDisk) SetSizeGB(v float32)`

SetSizeGB sets SizeGB field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


