# SwitchDNSRecordSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SwitchId** | **float32** | The id of the switch. | 
**Zone** | [**GenericDNSZoneInformation**](GenericDNSZoneInformation.md) | The DNS zone information. | 
**ManagementAddress** | **string** | The management address of the switch. | 
**Hostname** | **string** | The hostname of the switch DNS record. | 
**Fqdn** | **string** | The hostname of the switch DNS record. | 
**Ip** | **map[string]interface{}** | The IP address of the switch. | 
**Operation** | **string** | The operation to perform for the DNS record. Either \&quot;create\&quot; or \&quot;delete\&quot;. | 

## Methods

### NewSwitchDNSRecordSet

`func NewSwitchDNSRecordSet(switchId float32, zone GenericDNSZoneInformation, managementAddress string, hostname string, fqdn string, ip map[string]interface{}, operation string, ) *SwitchDNSRecordSet`

NewSwitchDNSRecordSet instantiates a new SwitchDNSRecordSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchDNSRecordSetWithDefaults

`func NewSwitchDNSRecordSetWithDefaults() *SwitchDNSRecordSet`

NewSwitchDNSRecordSetWithDefaults instantiates a new SwitchDNSRecordSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSwitchId

`func (o *SwitchDNSRecordSet) GetSwitchId() float32`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *SwitchDNSRecordSet) GetSwitchIdOk() (*float32, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *SwitchDNSRecordSet) SetSwitchId(v float32)`

SetSwitchId sets SwitchId field to given value.


### GetZone

`func (o *SwitchDNSRecordSet) GetZone() GenericDNSZoneInformation`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *SwitchDNSRecordSet) GetZoneOk() (*GenericDNSZoneInformation, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *SwitchDNSRecordSet) SetZone(v GenericDNSZoneInformation)`

SetZone sets Zone field to given value.


### GetManagementAddress

`func (o *SwitchDNSRecordSet) GetManagementAddress() string`

GetManagementAddress returns the ManagementAddress field if non-nil, zero value otherwise.

### GetManagementAddressOk

`func (o *SwitchDNSRecordSet) GetManagementAddressOk() (*string, bool)`

GetManagementAddressOk returns a tuple with the ManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddress

`func (o *SwitchDNSRecordSet) SetManagementAddress(v string)`

SetManagementAddress sets ManagementAddress field to given value.


### GetHostname

`func (o *SwitchDNSRecordSet) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *SwitchDNSRecordSet) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *SwitchDNSRecordSet) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetFqdn

`func (o *SwitchDNSRecordSet) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *SwitchDNSRecordSet) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *SwitchDNSRecordSet) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.


### GetIp

`func (o *SwitchDNSRecordSet) GetIp() map[string]interface{}`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *SwitchDNSRecordSet) GetIpOk() (*map[string]interface{}, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *SwitchDNSRecordSet) SetIp(v map[string]interface{})`

SetIp sets Ip field to given value.


### GetOperation

`func (o *SwitchDNSRecordSet) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *SwitchDNSRecordSet) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *SwitchDNSRecordSet) SetOperation(v string)`

SetOperation sets Operation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


