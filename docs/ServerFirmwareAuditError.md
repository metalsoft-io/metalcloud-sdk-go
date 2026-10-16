# ServerFirmwareAuditError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | **float32** | The HTTP status code describing why the server could not be audited. | 
**Message** | **string** | The reason the server could not be audited. | 

## Methods

### NewServerFirmwareAuditError

`func NewServerFirmwareAuditError(statusCode float32, message string, ) *ServerFirmwareAuditError`

NewServerFirmwareAuditError instantiates a new ServerFirmwareAuditError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerFirmwareAuditErrorWithDefaults

`func NewServerFirmwareAuditErrorWithDefaults() *ServerFirmwareAuditError`

NewServerFirmwareAuditErrorWithDefaults instantiates a new ServerFirmwareAuditError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *ServerFirmwareAuditError) GetStatusCode() float32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ServerFirmwareAuditError) GetStatusCodeOk() (*float32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ServerFirmwareAuditError) SetStatusCode(v float32)`

SetStatusCode sets StatusCode field to given value.


### GetMessage

`func (o *ServerFirmwareAuditError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ServerFirmwareAuditError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ServerFirmwareAuditError) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


