# NetworkDeviceSNMPConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **int32** | SNMP port | [optional] 
**Community** | Pointer to **string** | SNMP community string | [optional] 
**Contact** | Pointer to **string** | SNMP contact information | [optional] 

## Methods

### NewNetworkDeviceSNMPConfig

`func NewNetworkDeviceSNMPConfig() *NetworkDeviceSNMPConfig`

NewNetworkDeviceSNMPConfig instantiates a new NetworkDeviceSNMPConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceSNMPConfigWithDefaults

`func NewNetworkDeviceSNMPConfigWithDefaults() *NetworkDeviceSNMPConfig`

NewNetworkDeviceSNMPConfigWithDefaults instantiates a new NetworkDeviceSNMPConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *NetworkDeviceSNMPConfig) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *NetworkDeviceSNMPConfig) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *NetworkDeviceSNMPConfig) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *NetworkDeviceSNMPConfig) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetCommunity

`func (o *NetworkDeviceSNMPConfig) GetCommunity() string`

GetCommunity returns the Community field if non-nil, zero value otherwise.

### GetCommunityOk

`func (o *NetworkDeviceSNMPConfig) GetCommunityOk() (*string, bool)`

GetCommunityOk returns a tuple with the Community field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunity

`func (o *NetworkDeviceSNMPConfig) SetCommunity(v string)`

SetCommunity sets Community field to given value.

### HasCommunity

`func (o *NetworkDeviceSNMPConfig) HasCommunity() bool`

HasCommunity returns a boolean if a field has been set.

### GetContact

`func (o *NetworkDeviceSNMPConfig) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *NetworkDeviceSNMPConfig) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *NetworkDeviceSNMPConfig) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *NetworkDeviceSNMPConfig) HasContact() bool`

HasContact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


