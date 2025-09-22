# VMPoolRecordSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmPoolId** | **float32** | The ID of the VM Pool. | 
**Hostname** | **string** | The hostname of the VM Pool. | 
**Operation** | **string** | The operation of the VM Pool. | 

## Methods

### NewVMPoolRecordSet

`func NewVMPoolRecordSet(vmPoolId float32, hostname string, operation string, ) *VMPoolRecordSet`

NewVMPoolRecordSet instantiates a new VMPoolRecordSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMPoolRecordSetWithDefaults

`func NewVMPoolRecordSetWithDefaults() *VMPoolRecordSet`

NewVMPoolRecordSetWithDefaults instantiates a new VMPoolRecordSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmPoolId

`func (o *VMPoolRecordSet) GetVmPoolId() float32`

GetVmPoolId returns the VmPoolId field if non-nil, zero value otherwise.

### GetVmPoolIdOk

`func (o *VMPoolRecordSet) GetVmPoolIdOk() (*float32, bool)`

GetVmPoolIdOk returns a tuple with the VmPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmPoolId

`func (o *VMPoolRecordSet) SetVmPoolId(v float32)`

SetVmPoolId sets VmPoolId field to given value.


### GetHostname

`func (o *VMPoolRecordSet) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *VMPoolRecordSet) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *VMPoolRecordSet) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetOperation

`func (o *VMPoolRecordSet) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *VMPoolRecordSet) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *VMPoolRecordSet) SetOperation(v string)`

SetOperation sets Operation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


