# ServerDNSRecordSetDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | **float32** | The id of the server. | 
**Zone** | [**GenericDNSZoneInformation**](GenericDNSZoneInformation.md) | The DNS zone information. | 
**SerialNumber** | Pointer to **string** | The serial number of the server. | [optional] 
**ManagementAddress** | **string** | The management address of the server. | 
**Hostname** | **string** | The hostname of the server DNS record. | 
**Fqdn** | **string** | The hostname of the server DNS record. | 
**Ip** | **map[string]interface{}** | The IP address of the server. | 
**Operation** | **string** | The operation to perform for the DNS record. Either \&quot;create\&quot; or \&quot;delete\&quot;. | 

## Methods

### NewServerDNSRecordSetDto

`func NewServerDNSRecordSetDto(serverId float32, zone GenericDNSZoneInformation, managementAddress string, hostname string, fqdn string, ip map[string]interface{}, operation string, ) *ServerDNSRecordSetDto`

NewServerDNSRecordSetDto instantiates a new ServerDNSRecordSetDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerDNSRecordSetDtoWithDefaults

`func NewServerDNSRecordSetDtoWithDefaults() *ServerDNSRecordSetDto`

NewServerDNSRecordSetDtoWithDefaults instantiates a new ServerDNSRecordSetDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *ServerDNSRecordSetDto) GetServerId() float32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ServerDNSRecordSetDto) GetServerIdOk() (*float32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ServerDNSRecordSetDto) SetServerId(v float32)`

SetServerId sets ServerId field to given value.


### GetZone

`func (o *ServerDNSRecordSetDto) GetZone() GenericDNSZoneInformation`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ServerDNSRecordSetDto) GetZoneOk() (*GenericDNSZoneInformation, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ServerDNSRecordSetDto) SetZone(v GenericDNSZoneInformation)`

SetZone sets Zone field to given value.


### GetSerialNumber

`func (o *ServerDNSRecordSetDto) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ServerDNSRecordSetDto) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *ServerDNSRecordSetDto) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *ServerDNSRecordSetDto) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetManagementAddress

`func (o *ServerDNSRecordSetDto) GetManagementAddress() string`

GetManagementAddress returns the ManagementAddress field if non-nil, zero value otherwise.

### GetManagementAddressOk

`func (o *ServerDNSRecordSetDto) GetManagementAddressOk() (*string, bool)`

GetManagementAddressOk returns a tuple with the ManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddress

`func (o *ServerDNSRecordSetDto) SetManagementAddress(v string)`

SetManagementAddress sets ManagementAddress field to given value.


### GetHostname

`func (o *ServerDNSRecordSetDto) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ServerDNSRecordSetDto) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ServerDNSRecordSetDto) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetFqdn

`func (o *ServerDNSRecordSetDto) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *ServerDNSRecordSetDto) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *ServerDNSRecordSetDto) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.


### GetIp

`func (o *ServerDNSRecordSetDto) GetIp() map[string]interface{}`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ServerDNSRecordSetDto) GetIpOk() (*map[string]interface{}, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ServerDNSRecordSetDto) SetIp(v map[string]interface{})`

SetIp sets Ip field to given value.


### GetOperation

`func (o *ServerDNSRecordSetDto) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ServerDNSRecordSetDto) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ServerDNSRecordSetDto) SetOperation(v string)`

SetOperation sets Operation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


