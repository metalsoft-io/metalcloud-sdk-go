# ServerUnmanagedImport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManagementAddress** | Pointer to **string** | The management address of the server. | [optional] 
**SiteId** | **float32** | The site id where the server is located. | 
**ServerTypeId** | **float32** | The server type id. | 
**ServerSupportsOobProvisioning** | Pointer to **float32** | Flag to indicate if the server supports SOL. | [optional] 
**ServerInterfaces** | [**[]ServerUnmanagedImportInternalInterface**](ServerUnmanagedImportInternalInterface.md) | The interfaces of the server. | 
**ServerSerialNumber** | Pointer to **string** | The server Serial Number. | [optional] 
**ServerUUID** | Pointer to **string** | The server UUID. | [optional] 
**Hostname** | Pointer to **string** | The server hostname. | [optional] 
**ServerIpmiUsername** | Pointer to **string** | The server IPMI username. | [optional] 
**ServerIpmiPassword** | Pointer to **string** | The server IPMI password. | [optional] 

## Methods

### NewServerUnmanagedImport

`func NewServerUnmanagedImport(siteId float32, serverTypeId float32, serverInterfaces []ServerUnmanagedImportInternalInterface, ) *ServerUnmanagedImport`

NewServerUnmanagedImport instantiates a new ServerUnmanagedImport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerUnmanagedImportWithDefaults

`func NewServerUnmanagedImportWithDefaults() *ServerUnmanagedImport`

NewServerUnmanagedImportWithDefaults instantiates a new ServerUnmanagedImport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManagementAddress

`func (o *ServerUnmanagedImport) GetManagementAddress() string`

GetManagementAddress returns the ManagementAddress field if non-nil, zero value otherwise.

### GetManagementAddressOk

`func (o *ServerUnmanagedImport) GetManagementAddressOk() (*string, bool)`

GetManagementAddressOk returns a tuple with the ManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddress

`func (o *ServerUnmanagedImport) SetManagementAddress(v string)`

SetManagementAddress sets ManagementAddress field to given value.

### HasManagementAddress

`func (o *ServerUnmanagedImport) HasManagementAddress() bool`

HasManagementAddress returns a boolean if a field has been set.

### GetSiteId

`func (o *ServerUnmanagedImport) GetSiteId() float32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ServerUnmanagedImport) GetSiteIdOk() (*float32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ServerUnmanagedImport) SetSiteId(v float32)`

SetSiteId sets SiteId field to given value.


### GetServerTypeId

`func (o *ServerUnmanagedImport) GetServerTypeId() float32`

GetServerTypeId returns the ServerTypeId field if non-nil, zero value otherwise.

### GetServerTypeIdOk

`func (o *ServerUnmanagedImport) GetServerTypeIdOk() (*float32, bool)`

GetServerTypeIdOk returns a tuple with the ServerTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeId

`func (o *ServerUnmanagedImport) SetServerTypeId(v float32)`

SetServerTypeId sets ServerTypeId field to given value.


### GetServerSupportsOobProvisioning

`func (o *ServerUnmanagedImport) GetServerSupportsOobProvisioning() float32`

GetServerSupportsOobProvisioning returns the ServerSupportsOobProvisioning field if non-nil, zero value otherwise.

### GetServerSupportsOobProvisioningOk

`func (o *ServerUnmanagedImport) GetServerSupportsOobProvisioningOk() (*float32, bool)`

GetServerSupportsOobProvisioningOk returns a tuple with the ServerSupportsOobProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSupportsOobProvisioning

`func (o *ServerUnmanagedImport) SetServerSupportsOobProvisioning(v float32)`

SetServerSupportsOobProvisioning sets ServerSupportsOobProvisioning field to given value.

### HasServerSupportsOobProvisioning

`func (o *ServerUnmanagedImport) HasServerSupportsOobProvisioning() bool`

HasServerSupportsOobProvisioning returns a boolean if a field has been set.

### GetServerInterfaces

`func (o *ServerUnmanagedImport) GetServerInterfaces() []ServerUnmanagedImportInternalInterface`

GetServerInterfaces returns the ServerInterfaces field if non-nil, zero value otherwise.

### GetServerInterfacesOk

`func (o *ServerUnmanagedImport) GetServerInterfacesOk() (*[]ServerUnmanagedImportInternalInterface, bool)`

GetServerInterfacesOk returns a tuple with the ServerInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInterfaces

`func (o *ServerUnmanagedImport) SetServerInterfaces(v []ServerUnmanagedImportInternalInterface)`

SetServerInterfaces sets ServerInterfaces field to given value.


### GetServerSerialNumber

`func (o *ServerUnmanagedImport) GetServerSerialNumber() string`

GetServerSerialNumber returns the ServerSerialNumber field if non-nil, zero value otherwise.

### GetServerSerialNumberOk

`func (o *ServerUnmanagedImport) GetServerSerialNumberOk() (*string, bool)`

GetServerSerialNumberOk returns a tuple with the ServerSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSerialNumber

`func (o *ServerUnmanagedImport) SetServerSerialNumber(v string)`

SetServerSerialNumber sets ServerSerialNumber field to given value.

### HasServerSerialNumber

`func (o *ServerUnmanagedImport) HasServerSerialNumber() bool`

HasServerSerialNumber returns a boolean if a field has been set.

### GetServerUUID

`func (o *ServerUnmanagedImport) GetServerUUID() string`

GetServerUUID returns the ServerUUID field if non-nil, zero value otherwise.

### GetServerUUIDOk

`func (o *ServerUnmanagedImport) GetServerUUIDOk() (*string, bool)`

GetServerUUIDOk returns a tuple with the ServerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUUID

`func (o *ServerUnmanagedImport) SetServerUUID(v string)`

SetServerUUID sets ServerUUID field to given value.

### HasServerUUID

`func (o *ServerUnmanagedImport) HasServerUUID() bool`

HasServerUUID returns a boolean if a field has been set.

### GetHostname

`func (o *ServerUnmanagedImport) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ServerUnmanagedImport) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ServerUnmanagedImport) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ServerUnmanagedImport) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetServerIpmiUsername

`func (o *ServerUnmanagedImport) GetServerIpmiUsername() string`

GetServerIpmiUsername returns the ServerIpmiUsername field if non-nil, zero value otherwise.

### GetServerIpmiUsernameOk

`func (o *ServerUnmanagedImport) GetServerIpmiUsernameOk() (*string, bool)`

GetServerIpmiUsernameOk returns a tuple with the ServerIpmiUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIpmiUsername

`func (o *ServerUnmanagedImport) SetServerIpmiUsername(v string)`

SetServerIpmiUsername sets ServerIpmiUsername field to given value.

### HasServerIpmiUsername

`func (o *ServerUnmanagedImport) HasServerIpmiUsername() bool`

HasServerIpmiUsername returns a boolean if a field has been set.

### GetServerIpmiPassword

`func (o *ServerUnmanagedImport) GetServerIpmiPassword() string`

GetServerIpmiPassword returns the ServerIpmiPassword field if non-nil, zero value otherwise.

### GetServerIpmiPasswordOk

`func (o *ServerUnmanagedImport) GetServerIpmiPasswordOk() (*string, bool)`

GetServerIpmiPasswordOk returns a tuple with the ServerIpmiPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIpmiPassword

`func (o *ServerUnmanagedImport) SetServerIpmiPassword(v string)`

SetServerIpmiPassword sets ServerIpmiPassword field to given value.

### HasServerIpmiPassword

`func (o *ServerUnmanagedImport) HasServerIpmiPassword() bool`

HasServerIpmiPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


