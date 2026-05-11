# HardwareRescanServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RebootAllowed** | Pointer to **bool** | Whether a server reboot is allowed to perform full LLDP interface discovery. If false, existing switch connections are preserved. | [optional] [default to false]

## Methods

### NewHardwareRescanServerRequest

`func NewHardwareRescanServerRequest() *HardwareRescanServerRequest`

NewHardwareRescanServerRequest instantiates a new HardwareRescanServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHardwareRescanServerRequestWithDefaults

`func NewHardwareRescanServerRequestWithDefaults() *HardwareRescanServerRequest`

NewHardwareRescanServerRequestWithDefaults instantiates a new HardwareRescanServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRebootAllowed

`func (o *HardwareRescanServerRequest) GetRebootAllowed() bool`

GetRebootAllowed returns the RebootAllowed field if non-nil, zero value otherwise.

### GetRebootAllowedOk

`func (o *HardwareRescanServerRequest) GetRebootAllowedOk() (*bool, bool)`

GetRebootAllowedOk returns a tuple with the RebootAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootAllowed

`func (o *HardwareRescanServerRequest) SetRebootAllowed(v bool)`

SetRebootAllowed sets RebootAllowed field to given value.

### HasRebootAllowed

`func (o *HardwareRescanServerRequest) HasRebootAllowed() bool`

HasRebootAllowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


