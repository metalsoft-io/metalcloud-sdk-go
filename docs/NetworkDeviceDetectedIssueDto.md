# NetworkDeviceDetectedIssueDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailureType** | **string** | &#39;network-device&#39; &#x3D; device-level fault (hardware or software/control-plane); &#39;link&#39; &#x3D; fault scoped to a single port/link | 
**Severity** | **string** | Severity of this individual issue | 
**AffectedPort** | Pointer to **string** | Interface/port name — set for &#39;link&#39; failures | [optional] 
**LinkRole** | Pointer to **string** | For &#39;link&#39; failures: &#39;fabric&#39; &#x3D; inter-switch leaf-spine link; &#39;infrastructure&#39; &#x3D; downlink to a server in an active deployment; &#39;unused&#39; &#x3D; no live connection; &#39;unknown&#39; &#x3D; undetermined | [optional] 
**Description** | **string** | Short description of this specific failure | 

## Methods

### NewNetworkDeviceDetectedIssueDto

`func NewNetworkDeviceDetectedIssueDto(failureType string, severity string, description string, ) *NetworkDeviceDetectedIssueDto`

NewNetworkDeviceDetectedIssueDto instantiates a new NetworkDeviceDetectedIssueDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceDetectedIssueDtoWithDefaults

`func NewNetworkDeviceDetectedIssueDtoWithDefaults() *NetworkDeviceDetectedIssueDto`

NewNetworkDeviceDetectedIssueDtoWithDefaults instantiates a new NetworkDeviceDetectedIssueDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailureType

`func (o *NetworkDeviceDetectedIssueDto) GetFailureType() string`

GetFailureType returns the FailureType field if non-nil, zero value otherwise.

### GetFailureTypeOk

`func (o *NetworkDeviceDetectedIssueDto) GetFailureTypeOk() (*string, bool)`

GetFailureTypeOk returns a tuple with the FailureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureType

`func (o *NetworkDeviceDetectedIssueDto) SetFailureType(v string)`

SetFailureType sets FailureType field to given value.


### GetSeverity

`func (o *NetworkDeviceDetectedIssueDto) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *NetworkDeviceDetectedIssueDto) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *NetworkDeviceDetectedIssueDto) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetAffectedPort

`func (o *NetworkDeviceDetectedIssueDto) GetAffectedPort() string`

GetAffectedPort returns the AffectedPort field if non-nil, zero value otherwise.

### GetAffectedPortOk

`func (o *NetworkDeviceDetectedIssueDto) GetAffectedPortOk() (*string, bool)`

GetAffectedPortOk returns a tuple with the AffectedPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedPort

`func (o *NetworkDeviceDetectedIssueDto) SetAffectedPort(v string)`

SetAffectedPort sets AffectedPort field to given value.

### HasAffectedPort

`func (o *NetworkDeviceDetectedIssueDto) HasAffectedPort() bool`

HasAffectedPort returns a boolean if a field has been set.

### GetLinkRole

`func (o *NetworkDeviceDetectedIssueDto) GetLinkRole() string`

GetLinkRole returns the LinkRole field if non-nil, zero value otherwise.

### GetLinkRoleOk

`func (o *NetworkDeviceDetectedIssueDto) GetLinkRoleOk() (*string, bool)`

GetLinkRoleOk returns a tuple with the LinkRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkRole

`func (o *NetworkDeviceDetectedIssueDto) SetLinkRole(v string)`

SetLinkRole sets LinkRole field to given value.

### HasLinkRole

`func (o *NetworkDeviceDetectedIssueDto) HasLinkRole() bool`

HasLinkRole returns a boolean if a field has been set.

### GetDescription

`func (o *NetworkDeviceDetectedIssueDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkDeviceDetectedIssueDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkDeviceDetectedIssueDto) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


