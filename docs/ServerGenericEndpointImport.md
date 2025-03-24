# ServerGenericEndpointImport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManagementAddress** | **string** | The management address of the server. | 
**SiteId** | **float32** | The site id where the server is located. | 
**ServerTypeId** | **float32** | The server type id. | 
**ServerSupportsOobProvisioning** | **bool** | Flag to indicate if the server supports SOL. | 
**ServerInterfaces** | [**[]ServerUnmanagedImportInternalInterfaceDto**](ServerUnmanagedImportInternalInterfaceDto.md) | The interfaces of the server. | 
**ServerSerialNumber** | Pointer to **string** | The server Serial Number. | [optional] 
**ServerUUID** | Pointer to **string** | The server UUID. | [optional] 
**Hostname** | Pointer to **string** | The server hostname. | [optional] 
**ServerIpmiUsername** | Pointer to **string** | The server IPMI username. | [optional] 
**ServerIpmiPassword** | Pointer to **string** | The server IPMI password. | [optional] 

## Methods

### NewServerGenericEndpointImport

`func NewServerGenericEndpointImport(managementAddress string, siteId float32, serverTypeId float32, serverSupportsOobProvisioning bool, serverInterfaces []ServerUnmanagedImportInternalInterfaceDto, ) *ServerGenericEndpointImport`

NewServerGenericEndpointImport instantiates a new ServerGenericEndpointImport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerGenericEndpointImportWithDefaults

`func NewServerGenericEndpointImportWithDefaults() *ServerGenericEndpointImport`

NewServerGenericEndpointImportWithDefaults instantiates a new ServerGenericEndpointImport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManagementAddress

`func (o *ServerGenericEndpointImport) GetManagementAddress() string`

GetManagementAddress returns the ManagementAddress field if non-nil, zero value otherwise.

### GetManagementAddressOk

`func (o *ServerGenericEndpointImport) GetManagementAddressOk() (*string, bool)`

GetManagementAddressOk returns a tuple with the ManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddress

`func (o *ServerGenericEndpointImport) SetManagementAddress(v string)`

SetManagementAddress sets ManagementAddress field to given value.


### GetSiteId

`func (o *ServerGenericEndpointImport) GetSiteId() float32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ServerGenericEndpointImport) GetSiteIdOk() (*float32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ServerGenericEndpointImport) SetSiteId(v float32)`

SetSiteId sets SiteId field to given value.


### GetServerTypeId

`func (o *ServerGenericEndpointImport) GetServerTypeId() float32`

GetServerTypeId returns the ServerTypeId field if non-nil, zero value otherwise.

### GetServerTypeIdOk

`func (o *ServerGenericEndpointImport) GetServerTypeIdOk() (*float32, bool)`

GetServerTypeIdOk returns a tuple with the ServerTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeId

`func (o *ServerGenericEndpointImport) SetServerTypeId(v float32)`

SetServerTypeId sets ServerTypeId field to given value.


### GetServerSupportsOobProvisioning

`func (o *ServerGenericEndpointImport) GetServerSupportsOobProvisioning() bool`

GetServerSupportsOobProvisioning returns the ServerSupportsOobProvisioning field if non-nil, zero value otherwise.

### GetServerSupportsOobProvisioningOk

`func (o *ServerGenericEndpointImport) GetServerSupportsOobProvisioningOk() (*bool, bool)`

GetServerSupportsOobProvisioningOk returns a tuple with the ServerSupportsOobProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSupportsOobProvisioning

`func (o *ServerGenericEndpointImport) SetServerSupportsOobProvisioning(v bool)`

SetServerSupportsOobProvisioning sets ServerSupportsOobProvisioning field to given value.


### GetServerInterfaces

`func (o *ServerGenericEndpointImport) GetServerInterfaces() []ServerUnmanagedImportInternalInterfaceDto`

GetServerInterfaces returns the ServerInterfaces field if non-nil, zero value otherwise.

### GetServerInterfacesOk

`func (o *ServerGenericEndpointImport) GetServerInterfacesOk() (*[]ServerUnmanagedImportInternalInterfaceDto, bool)`

GetServerInterfacesOk returns a tuple with the ServerInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInterfaces

`func (o *ServerGenericEndpointImport) SetServerInterfaces(v []ServerUnmanagedImportInternalInterfaceDto)`

SetServerInterfaces sets ServerInterfaces field to given value.


### GetServerSerialNumber

`func (o *ServerGenericEndpointImport) GetServerSerialNumber() string`

GetServerSerialNumber returns the ServerSerialNumber field if non-nil, zero value otherwise.

### GetServerSerialNumberOk

`func (o *ServerGenericEndpointImport) GetServerSerialNumberOk() (*string, bool)`

GetServerSerialNumberOk returns a tuple with the ServerSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSerialNumber

`func (o *ServerGenericEndpointImport) SetServerSerialNumber(v string)`

SetServerSerialNumber sets ServerSerialNumber field to given value.

### HasServerSerialNumber

`func (o *ServerGenericEndpointImport) HasServerSerialNumber() bool`

HasServerSerialNumber returns a boolean if a field has been set.

### GetServerUUID

`func (o *ServerGenericEndpointImport) GetServerUUID() string`

GetServerUUID returns the ServerUUID field if non-nil, zero value otherwise.

### GetServerUUIDOk

`func (o *ServerGenericEndpointImport) GetServerUUIDOk() (*string, bool)`

GetServerUUIDOk returns a tuple with the ServerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUUID

`func (o *ServerGenericEndpointImport) SetServerUUID(v string)`

SetServerUUID sets ServerUUID field to given value.

### HasServerUUID

`func (o *ServerGenericEndpointImport) HasServerUUID() bool`

HasServerUUID returns a boolean if a field has been set.

### GetHostname

`func (o *ServerGenericEndpointImport) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ServerGenericEndpointImport) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ServerGenericEndpointImport) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ServerGenericEndpointImport) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetServerIpmiUsername

`func (o *ServerGenericEndpointImport) GetServerIpmiUsername() string`

GetServerIpmiUsername returns the ServerIpmiUsername field if non-nil, zero value otherwise.

### GetServerIpmiUsernameOk

`func (o *ServerGenericEndpointImport) GetServerIpmiUsernameOk() (*string, bool)`

GetServerIpmiUsernameOk returns a tuple with the ServerIpmiUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIpmiUsername

`func (o *ServerGenericEndpointImport) SetServerIpmiUsername(v string)`

SetServerIpmiUsername sets ServerIpmiUsername field to given value.

### HasServerIpmiUsername

`func (o *ServerGenericEndpointImport) HasServerIpmiUsername() bool`

HasServerIpmiUsername returns a boolean if a field has been set.

### GetServerIpmiPassword

`func (o *ServerGenericEndpointImport) GetServerIpmiPassword() string`

GetServerIpmiPassword returns the ServerIpmiPassword field if non-nil, zero value otherwise.

### GetServerIpmiPasswordOk

`func (o *ServerGenericEndpointImport) GetServerIpmiPasswordOk() (*string, bool)`

GetServerIpmiPasswordOk returns a tuple with the ServerIpmiPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIpmiPassword

`func (o *ServerGenericEndpointImport) SetServerIpmiPassword(v string)`

SetServerIpmiPassword sets ServerIpmiPassword field to given value.

### HasServerIpmiPassword

`func (o *ServerGenericEndpointImport) HasServerIpmiPassword() bool`

HasServerIpmiPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


