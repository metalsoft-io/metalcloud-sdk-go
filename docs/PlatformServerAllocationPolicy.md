# PlatformServerAllocationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreventAllocationOnCriticalHealthIssues** | Pointer to **bool** | Prevent server allocation when the server presents critical health issues | [optional] 
**PreventAllocationOnAnyHealthIssues** | Pointer to **bool** | Prevent server allocation when the server presents any health issues | [optional] 
**PreventAllocationOnFirmwareDrift** | Pointer to **bool** | Prevent server allocation when the server has firmware version drift from the assigned baseline | [optional] 

## Methods

### NewPlatformServerAllocationPolicy

`func NewPlatformServerAllocationPolicy() *PlatformServerAllocationPolicy`

NewPlatformServerAllocationPolicy instantiates a new PlatformServerAllocationPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformServerAllocationPolicyWithDefaults

`func NewPlatformServerAllocationPolicyWithDefaults() *PlatformServerAllocationPolicy`

NewPlatformServerAllocationPolicyWithDefaults instantiates a new PlatformServerAllocationPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreventAllocationOnCriticalHealthIssues

`func (o *PlatformServerAllocationPolicy) GetPreventAllocationOnCriticalHealthIssues() bool`

GetPreventAllocationOnCriticalHealthIssues returns the PreventAllocationOnCriticalHealthIssues field if non-nil, zero value otherwise.

### GetPreventAllocationOnCriticalHealthIssuesOk

`func (o *PlatformServerAllocationPolicy) GetPreventAllocationOnCriticalHealthIssuesOk() (*bool, bool)`

GetPreventAllocationOnCriticalHealthIssuesOk returns a tuple with the PreventAllocationOnCriticalHealthIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventAllocationOnCriticalHealthIssues

`func (o *PlatformServerAllocationPolicy) SetPreventAllocationOnCriticalHealthIssues(v bool)`

SetPreventAllocationOnCriticalHealthIssues sets PreventAllocationOnCriticalHealthIssues field to given value.

### HasPreventAllocationOnCriticalHealthIssues

`func (o *PlatformServerAllocationPolicy) HasPreventAllocationOnCriticalHealthIssues() bool`

HasPreventAllocationOnCriticalHealthIssues returns a boolean if a field has been set.

### GetPreventAllocationOnAnyHealthIssues

`func (o *PlatformServerAllocationPolicy) GetPreventAllocationOnAnyHealthIssues() bool`

GetPreventAllocationOnAnyHealthIssues returns the PreventAllocationOnAnyHealthIssues field if non-nil, zero value otherwise.

### GetPreventAllocationOnAnyHealthIssuesOk

`func (o *PlatformServerAllocationPolicy) GetPreventAllocationOnAnyHealthIssuesOk() (*bool, bool)`

GetPreventAllocationOnAnyHealthIssuesOk returns a tuple with the PreventAllocationOnAnyHealthIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventAllocationOnAnyHealthIssues

`func (o *PlatformServerAllocationPolicy) SetPreventAllocationOnAnyHealthIssues(v bool)`

SetPreventAllocationOnAnyHealthIssues sets PreventAllocationOnAnyHealthIssues field to given value.

### HasPreventAllocationOnAnyHealthIssues

`func (o *PlatformServerAllocationPolicy) HasPreventAllocationOnAnyHealthIssues() bool`

HasPreventAllocationOnAnyHealthIssues returns a boolean if a field has been set.

### GetPreventAllocationOnFirmwareDrift

`func (o *PlatformServerAllocationPolicy) GetPreventAllocationOnFirmwareDrift() bool`

GetPreventAllocationOnFirmwareDrift returns the PreventAllocationOnFirmwareDrift field if non-nil, zero value otherwise.

### GetPreventAllocationOnFirmwareDriftOk

`func (o *PlatformServerAllocationPolicy) GetPreventAllocationOnFirmwareDriftOk() (*bool, bool)`

GetPreventAllocationOnFirmwareDriftOk returns a tuple with the PreventAllocationOnFirmwareDrift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventAllocationOnFirmwareDrift

`func (o *PlatformServerAllocationPolicy) SetPreventAllocationOnFirmwareDrift(v bool)`

SetPreventAllocationOnFirmwareDrift sets PreventAllocationOnFirmwareDrift field to given value.

### HasPreventAllocationOnFirmwareDrift

`func (o *PlatformServerAllocationPolicy) HasPreventAllocationOnFirmwareDrift() bool`

HasPreventAllocationOnFirmwareDrift returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


