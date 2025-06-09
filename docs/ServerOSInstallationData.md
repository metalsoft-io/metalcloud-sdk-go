# ServerOSInstallationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | **float32** | The id of the server. | 
**ServerTypeId** | Pointer to **float32** | The id of the server type. | [optional] 
**ServerUUID** | Pointer to **string** | The UUID of the server. | [optional] 
**SerialNumber** | Pointer to **string** | The Serial Number of the server. | [optional] 
**ManagementAddress** | Pointer to **string** | The Management Address of the server. | [optional] 
**Username** | Pointer to **string** | The username to use. | [optional] 
**PasswordEncrypted** | **string** | The encrypted password. | 
**DiskCount** | Pointer to **float32** | The disk count of the server. | [optional] 
**BmcMacAddress** | Pointer to **string** | The MAC address of the server. | [optional] 
**Vendor** | Pointer to **string** | The vendor of the server. | [optional] 
**VendorSkuId** | Pointer to **string** | The vendor sku id of the server. | [optional] 
**Model** | Pointer to **string** | The model of the server. | [optional] 
**RackName** | Pointer to **string** | The chassis rack name of the server. | [optional] 
**GpuCount** | Pointer to **float32** | The GPU count of the server. | [optional] 
**GpuInfo** | Pointer to [**[]ServerGpuInfo**](ServerGpuInfo.md) | The GPU info of the server. | [optional] 
**Interfaces** | Pointer to [**[]ServerInterface**](ServerInterface.md) | The interfaces of the server. | [optional] 
**Disks** | Pointer to [**[]ServerDisk**](ServerDisk.md) | The disks of the server. | [optional] 
**StorageControllers** | Pointer to [**[]ServerStorageController**](ServerStorageController.md) | The storage controllers of the server. | [optional] 
**Tags** | Pointer to **[]string** | Tags for the Server. | [optional] 

## Methods

### NewServerOSInstallationData

`func NewServerOSInstallationData(serverId float32, passwordEncrypted string, ) *ServerOSInstallationData`

NewServerOSInstallationData instantiates a new ServerOSInstallationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerOSInstallationDataWithDefaults

`func NewServerOSInstallationDataWithDefaults() *ServerOSInstallationData`

NewServerOSInstallationDataWithDefaults instantiates a new ServerOSInstallationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *ServerOSInstallationData) GetServerId() float32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ServerOSInstallationData) GetServerIdOk() (*float32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ServerOSInstallationData) SetServerId(v float32)`

SetServerId sets ServerId field to given value.


### GetServerTypeId

`func (o *ServerOSInstallationData) GetServerTypeId() float32`

GetServerTypeId returns the ServerTypeId field if non-nil, zero value otherwise.

### GetServerTypeIdOk

`func (o *ServerOSInstallationData) GetServerTypeIdOk() (*float32, bool)`

GetServerTypeIdOk returns a tuple with the ServerTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeId

`func (o *ServerOSInstallationData) SetServerTypeId(v float32)`

SetServerTypeId sets ServerTypeId field to given value.

### HasServerTypeId

`func (o *ServerOSInstallationData) HasServerTypeId() bool`

HasServerTypeId returns a boolean if a field has been set.

### GetServerUUID

`func (o *ServerOSInstallationData) GetServerUUID() string`

GetServerUUID returns the ServerUUID field if non-nil, zero value otherwise.

### GetServerUUIDOk

`func (o *ServerOSInstallationData) GetServerUUIDOk() (*string, bool)`

GetServerUUIDOk returns a tuple with the ServerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUUID

`func (o *ServerOSInstallationData) SetServerUUID(v string)`

SetServerUUID sets ServerUUID field to given value.

### HasServerUUID

`func (o *ServerOSInstallationData) HasServerUUID() bool`

HasServerUUID returns a boolean if a field has been set.

### GetSerialNumber

`func (o *ServerOSInstallationData) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ServerOSInstallationData) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *ServerOSInstallationData) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *ServerOSInstallationData) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetManagementAddress

`func (o *ServerOSInstallationData) GetManagementAddress() string`

GetManagementAddress returns the ManagementAddress field if non-nil, zero value otherwise.

### GetManagementAddressOk

`func (o *ServerOSInstallationData) GetManagementAddressOk() (*string, bool)`

GetManagementAddressOk returns a tuple with the ManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddress

`func (o *ServerOSInstallationData) SetManagementAddress(v string)`

SetManagementAddress sets ManagementAddress field to given value.

### HasManagementAddress

`func (o *ServerOSInstallationData) HasManagementAddress() bool`

HasManagementAddress returns a boolean if a field has been set.

### GetUsername

`func (o *ServerOSInstallationData) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ServerOSInstallationData) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ServerOSInstallationData) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ServerOSInstallationData) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPasswordEncrypted

`func (o *ServerOSInstallationData) GetPasswordEncrypted() string`

GetPasswordEncrypted returns the PasswordEncrypted field if non-nil, zero value otherwise.

### GetPasswordEncryptedOk

`func (o *ServerOSInstallationData) GetPasswordEncryptedOk() (*string, bool)`

GetPasswordEncryptedOk returns a tuple with the PasswordEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordEncrypted

`func (o *ServerOSInstallationData) SetPasswordEncrypted(v string)`

SetPasswordEncrypted sets PasswordEncrypted field to given value.


### GetDiskCount

`func (o *ServerOSInstallationData) GetDiskCount() float32`

GetDiskCount returns the DiskCount field if non-nil, zero value otherwise.

### GetDiskCountOk

`func (o *ServerOSInstallationData) GetDiskCountOk() (*float32, bool)`

GetDiskCountOk returns a tuple with the DiskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskCount

`func (o *ServerOSInstallationData) SetDiskCount(v float32)`

SetDiskCount sets DiskCount field to given value.

### HasDiskCount

`func (o *ServerOSInstallationData) HasDiskCount() bool`

HasDiskCount returns a boolean if a field has been set.

### GetBmcMacAddress

`func (o *ServerOSInstallationData) GetBmcMacAddress() string`

GetBmcMacAddress returns the BmcMacAddress field if non-nil, zero value otherwise.

### GetBmcMacAddressOk

`func (o *ServerOSInstallationData) GetBmcMacAddressOk() (*string, bool)`

GetBmcMacAddressOk returns a tuple with the BmcMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmcMacAddress

`func (o *ServerOSInstallationData) SetBmcMacAddress(v string)`

SetBmcMacAddress sets BmcMacAddress field to given value.

### HasBmcMacAddress

`func (o *ServerOSInstallationData) HasBmcMacAddress() bool`

HasBmcMacAddress returns a boolean if a field has been set.

### GetVendor

`func (o *ServerOSInstallationData) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ServerOSInstallationData) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ServerOSInstallationData) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *ServerOSInstallationData) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVendorSkuId

`func (o *ServerOSInstallationData) GetVendorSkuId() string`

GetVendorSkuId returns the VendorSkuId field if non-nil, zero value otherwise.

### GetVendorSkuIdOk

`func (o *ServerOSInstallationData) GetVendorSkuIdOk() (*string, bool)`

GetVendorSkuIdOk returns a tuple with the VendorSkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorSkuId

`func (o *ServerOSInstallationData) SetVendorSkuId(v string)`

SetVendorSkuId sets VendorSkuId field to given value.

### HasVendorSkuId

`func (o *ServerOSInstallationData) HasVendorSkuId() bool`

HasVendorSkuId returns a boolean if a field has been set.

### GetModel

`func (o *ServerOSInstallationData) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ServerOSInstallationData) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ServerOSInstallationData) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ServerOSInstallationData) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetRackName

`func (o *ServerOSInstallationData) GetRackName() string`

GetRackName returns the RackName field if non-nil, zero value otherwise.

### GetRackNameOk

`func (o *ServerOSInstallationData) GetRackNameOk() (*string, bool)`

GetRackNameOk returns a tuple with the RackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackName

`func (o *ServerOSInstallationData) SetRackName(v string)`

SetRackName sets RackName field to given value.

### HasRackName

`func (o *ServerOSInstallationData) HasRackName() bool`

HasRackName returns a boolean if a field has been set.

### GetGpuCount

`func (o *ServerOSInstallationData) GetGpuCount() float32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *ServerOSInstallationData) GetGpuCountOk() (*float32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *ServerOSInstallationData) SetGpuCount(v float32)`

SetGpuCount sets GpuCount field to given value.

### HasGpuCount

`func (o *ServerOSInstallationData) HasGpuCount() bool`

HasGpuCount returns a boolean if a field has been set.

### GetGpuInfo

`func (o *ServerOSInstallationData) GetGpuInfo() []ServerGpuInfo`

GetGpuInfo returns the GpuInfo field if non-nil, zero value otherwise.

### GetGpuInfoOk

`func (o *ServerOSInstallationData) GetGpuInfoOk() (*[]ServerGpuInfo, bool)`

GetGpuInfoOk returns a tuple with the GpuInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuInfo

`func (o *ServerOSInstallationData) SetGpuInfo(v []ServerGpuInfo)`

SetGpuInfo sets GpuInfo field to given value.

### HasGpuInfo

`func (o *ServerOSInstallationData) HasGpuInfo() bool`

HasGpuInfo returns a boolean if a field has been set.

### GetInterfaces

`func (o *ServerOSInstallationData) GetInterfaces() []ServerInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *ServerOSInstallationData) GetInterfacesOk() (*[]ServerInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *ServerOSInstallationData) SetInterfaces(v []ServerInterface)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *ServerOSInstallationData) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetDisks

`func (o *ServerOSInstallationData) GetDisks() []ServerDisk`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *ServerOSInstallationData) GetDisksOk() (*[]ServerDisk, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *ServerOSInstallationData) SetDisks(v []ServerDisk)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *ServerOSInstallationData) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetStorageControllers

`func (o *ServerOSInstallationData) GetStorageControllers() []ServerStorageController`

GetStorageControllers returns the StorageControllers field if non-nil, zero value otherwise.

### GetStorageControllersOk

`func (o *ServerOSInstallationData) GetStorageControllersOk() (*[]ServerStorageController, bool)`

GetStorageControllersOk returns a tuple with the StorageControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllers

`func (o *ServerOSInstallationData) SetStorageControllers(v []ServerStorageController)`

SetStorageControllers sets StorageControllers field to given value.

### HasStorageControllers

`func (o *ServerOSInstallationData) HasStorageControllers() bool`

HasStorageControllers returns a boolean if a field has been set.

### GetTags

`func (o *ServerOSInstallationData) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServerOSInstallationData) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServerOSInstallationData) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServerOSInstallationData) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


