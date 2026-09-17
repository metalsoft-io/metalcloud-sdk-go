# ReturnNetworkDeviceToPlanned

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManagementMAC** | Pointer to **string** | Management MAC address of the physical device. Required only when the archived record has none. | [optional] 
**SerialNumber** | Pointer to **string** | Serial number of the physical device, used instead of the MAC address to identify it. | [optional] 
**OsTemplateId** | Pointer to **int64** | OS template for the zero-touch installation (ONIE install method). Overrides the stored one. | [optional] 

## Methods

### NewReturnNetworkDeviceToPlanned

`func NewReturnNetworkDeviceToPlanned() *ReturnNetworkDeviceToPlanned`

NewReturnNetworkDeviceToPlanned instantiates a new ReturnNetworkDeviceToPlanned object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnNetworkDeviceToPlannedWithDefaults

`func NewReturnNetworkDeviceToPlannedWithDefaults() *ReturnNetworkDeviceToPlanned`

NewReturnNetworkDeviceToPlannedWithDefaults instantiates a new ReturnNetworkDeviceToPlanned object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManagementMAC

`func (o *ReturnNetworkDeviceToPlanned) GetManagementMAC() string`

GetManagementMAC returns the ManagementMAC field if non-nil, zero value otherwise.

### GetManagementMACOk

`func (o *ReturnNetworkDeviceToPlanned) GetManagementMACOk() (*string, bool)`

GetManagementMACOk returns a tuple with the ManagementMAC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementMAC

`func (o *ReturnNetworkDeviceToPlanned) SetManagementMAC(v string)`

SetManagementMAC sets ManagementMAC field to given value.

### HasManagementMAC

`func (o *ReturnNetworkDeviceToPlanned) HasManagementMAC() bool`

HasManagementMAC returns a boolean if a field has been set.

### GetSerialNumber

`func (o *ReturnNetworkDeviceToPlanned) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ReturnNetworkDeviceToPlanned) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *ReturnNetworkDeviceToPlanned) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *ReturnNetworkDeviceToPlanned) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetOsTemplateId

`func (o *ReturnNetworkDeviceToPlanned) GetOsTemplateId() int64`

GetOsTemplateId returns the OsTemplateId field if non-nil, zero value otherwise.

### GetOsTemplateIdOk

`func (o *ReturnNetworkDeviceToPlanned) GetOsTemplateIdOk() (*int64, bool)`

GetOsTemplateIdOk returns a tuple with the OsTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsTemplateId

`func (o *ReturnNetworkDeviceToPlanned) SetOsTemplateId(v int64)`

SetOsTemplateId sets OsTemplateId field to given value.

### HasOsTemplateId

`func (o *ReturnNetworkDeviceToPlanned) HasOsTemplateId() bool`

HasOsTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


