# RecordSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | Pointer to [**Server**](Server.md) | The server data. | [optional] 
**NetworkDevice** | Pointer to [**NetworkDevice**](NetworkDevice.md) | The network device. | [optional] 
**BgpTemplateRecordSet** | Pointer to [**BgpTemplateRecordSet**](BgpTemplateRecordSet.md) | The BGP template record set. | [optional] 
**ServerInstanceGroupDNSRecordSet** | Pointer to **map[string]interface{}** | The Instance Array DNS record set | [optional] 
**ServerInstanceDNSRecordSet** | Pointer to **map[string]interface{}** | The Instance DNS record set | [optional] 
**ExtensionInstanceRecordSet** | Pointer to **map[string]interface{}** | The extension instance. | [optional] 
**ServerDNSRecordSet** | Pointer to [**ServerDNSRecordSetDto**](ServerDNSRecordSetDto.md) | The server DNS record set. | [optional] 
**SwitchDNSRecordSet** | Pointer to [**SwitchDNSRecordSet**](SwitchDNSRecordSet.md) | The switch DNS record set. | [optional] 

## Methods

### NewRecordSet

`func NewRecordSet() *RecordSet`

NewRecordSet instantiates a new RecordSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordSetWithDefaults

`func NewRecordSetWithDefaults() *RecordSet`

NewRecordSetWithDefaults instantiates a new RecordSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *RecordSet) GetServer() Server`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *RecordSet) GetServerOk() (*Server, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *RecordSet) SetServer(v Server)`

SetServer sets Server field to given value.

### HasServer

`func (o *RecordSet) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetNetworkDevice

`func (o *RecordSet) GetNetworkDevice() NetworkDevice`

GetNetworkDevice returns the NetworkDevice field if non-nil, zero value otherwise.

### GetNetworkDeviceOk

`func (o *RecordSet) GetNetworkDeviceOk() (*NetworkDevice, bool)`

GetNetworkDeviceOk returns a tuple with the NetworkDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDevice

`func (o *RecordSet) SetNetworkDevice(v NetworkDevice)`

SetNetworkDevice sets NetworkDevice field to given value.

### HasNetworkDevice

`func (o *RecordSet) HasNetworkDevice() bool`

HasNetworkDevice returns a boolean if a field has been set.

### GetBgpTemplateRecordSet

`func (o *RecordSet) GetBgpTemplateRecordSet() BgpTemplateRecordSet`

GetBgpTemplateRecordSet returns the BgpTemplateRecordSet field if non-nil, zero value otherwise.

### GetBgpTemplateRecordSetOk

`func (o *RecordSet) GetBgpTemplateRecordSetOk() (*BgpTemplateRecordSet, bool)`

GetBgpTemplateRecordSetOk returns a tuple with the BgpTemplateRecordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpTemplateRecordSet

`func (o *RecordSet) SetBgpTemplateRecordSet(v BgpTemplateRecordSet)`

SetBgpTemplateRecordSet sets BgpTemplateRecordSet field to given value.

### HasBgpTemplateRecordSet

`func (o *RecordSet) HasBgpTemplateRecordSet() bool`

HasBgpTemplateRecordSet returns a boolean if a field has been set.

### GetServerInstanceGroupDNSRecordSet

`func (o *RecordSet) GetServerInstanceGroupDNSRecordSet() map[string]interface{}`

GetServerInstanceGroupDNSRecordSet returns the ServerInstanceGroupDNSRecordSet field if non-nil, zero value otherwise.

### GetServerInstanceGroupDNSRecordSetOk

`func (o *RecordSet) GetServerInstanceGroupDNSRecordSetOk() (*map[string]interface{}, bool)`

GetServerInstanceGroupDNSRecordSetOk returns a tuple with the ServerInstanceGroupDNSRecordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInstanceGroupDNSRecordSet

`func (o *RecordSet) SetServerInstanceGroupDNSRecordSet(v map[string]interface{})`

SetServerInstanceGroupDNSRecordSet sets ServerInstanceGroupDNSRecordSet field to given value.

### HasServerInstanceGroupDNSRecordSet

`func (o *RecordSet) HasServerInstanceGroupDNSRecordSet() bool`

HasServerInstanceGroupDNSRecordSet returns a boolean if a field has been set.

### GetServerInstanceDNSRecordSet

`func (o *RecordSet) GetServerInstanceDNSRecordSet() map[string]interface{}`

GetServerInstanceDNSRecordSet returns the ServerInstanceDNSRecordSet field if non-nil, zero value otherwise.

### GetServerInstanceDNSRecordSetOk

`func (o *RecordSet) GetServerInstanceDNSRecordSetOk() (*map[string]interface{}, bool)`

GetServerInstanceDNSRecordSetOk returns a tuple with the ServerInstanceDNSRecordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInstanceDNSRecordSet

`func (o *RecordSet) SetServerInstanceDNSRecordSet(v map[string]interface{})`

SetServerInstanceDNSRecordSet sets ServerInstanceDNSRecordSet field to given value.

### HasServerInstanceDNSRecordSet

`func (o *RecordSet) HasServerInstanceDNSRecordSet() bool`

HasServerInstanceDNSRecordSet returns a boolean if a field has been set.

### GetExtensionInstanceRecordSet

`func (o *RecordSet) GetExtensionInstanceRecordSet() map[string]interface{}`

GetExtensionInstanceRecordSet returns the ExtensionInstanceRecordSet field if non-nil, zero value otherwise.

### GetExtensionInstanceRecordSetOk

`func (o *RecordSet) GetExtensionInstanceRecordSetOk() (*map[string]interface{}, bool)`

GetExtensionInstanceRecordSetOk returns a tuple with the ExtensionInstanceRecordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionInstanceRecordSet

`func (o *RecordSet) SetExtensionInstanceRecordSet(v map[string]interface{})`

SetExtensionInstanceRecordSet sets ExtensionInstanceRecordSet field to given value.

### HasExtensionInstanceRecordSet

`func (o *RecordSet) HasExtensionInstanceRecordSet() bool`

HasExtensionInstanceRecordSet returns a boolean if a field has been set.

### GetServerDNSRecordSet

`func (o *RecordSet) GetServerDNSRecordSet() ServerDNSRecordSetDto`

GetServerDNSRecordSet returns the ServerDNSRecordSet field if non-nil, zero value otherwise.

### GetServerDNSRecordSetOk

`func (o *RecordSet) GetServerDNSRecordSetOk() (*ServerDNSRecordSetDto, bool)`

GetServerDNSRecordSetOk returns a tuple with the ServerDNSRecordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerDNSRecordSet

`func (o *RecordSet) SetServerDNSRecordSet(v ServerDNSRecordSetDto)`

SetServerDNSRecordSet sets ServerDNSRecordSet field to given value.

### HasServerDNSRecordSet

`func (o *RecordSet) HasServerDNSRecordSet() bool`

HasServerDNSRecordSet returns a boolean if a field has been set.

### GetSwitchDNSRecordSet

`func (o *RecordSet) GetSwitchDNSRecordSet() SwitchDNSRecordSet`

GetSwitchDNSRecordSet returns the SwitchDNSRecordSet field if non-nil, zero value otherwise.

### GetSwitchDNSRecordSetOk

`func (o *RecordSet) GetSwitchDNSRecordSetOk() (*SwitchDNSRecordSet, bool)`

GetSwitchDNSRecordSetOk returns a tuple with the SwitchDNSRecordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchDNSRecordSet

`func (o *RecordSet) SetSwitchDNSRecordSet(v SwitchDNSRecordSet)`

SetSwitchDNSRecordSet sets SwitchDNSRecordSet field to given value.

### HasSwitchDNSRecordSet

`func (o *RecordSet) HasSwitchDNSRecordSet() bool`

HasSwitchDNSRecordSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


