# OSTemplateInstall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | The OS template installation method | 
**DriveType** | **string** | The OS template installation drive type | 
**ReadyMethod** | **string** | The OS template installation ready method,                     The \&quot;ready method\&quot; is used to determine when the OS installation is complete. | 

## Methods

### NewOSTemplateInstall

`func NewOSTemplateInstall(method string, driveType string, readyMethod string, ) *OSTemplateInstall`

NewOSTemplateInstall instantiates a new OSTemplateInstall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSTemplateInstallWithDefaults

`func NewOSTemplateInstallWithDefaults() *OSTemplateInstall`

NewOSTemplateInstallWithDefaults instantiates a new OSTemplateInstall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *OSTemplateInstall) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *OSTemplateInstall) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *OSTemplateInstall) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetDriveType

`func (o *OSTemplateInstall) GetDriveType() string`

GetDriveType returns the DriveType field if non-nil, zero value otherwise.

### GetDriveTypeOk

`func (o *OSTemplateInstall) GetDriveTypeOk() (*string, bool)`

GetDriveTypeOk returns a tuple with the DriveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveType

`func (o *OSTemplateInstall) SetDriveType(v string)`

SetDriveType sets DriveType field to given value.


### GetReadyMethod

`func (o *OSTemplateInstall) GetReadyMethod() string`

GetReadyMethod returns the ReadyMethod field if non-nil, zero value otherwise.

### GetReadyMethodOk

`func (o *OSTemplateInstall) GetReadyMethodOk() (*string, bool)`

GetReadyMethodOk returns a tuple with the ReadyMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyMethod

`func (o *OSTemplateInstall) SetReadyMethod(v string)`

SetReadyMethod sets ReadyMethod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


