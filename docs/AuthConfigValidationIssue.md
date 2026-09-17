# AuthConfigValidationIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Configuration field the issue relates to, as a dot-notation path | 
**Severity** | [**AuthConfigValidationSeverity**](AuthConfigValidationSeverity.md) | Severity of the issue; errors make the configuration invalid, warnings do not | 
**Message** | **string** | Human-readable description of the issue | 

## Methods

### NewAuthConfigValidationIssue

`func NewAuthConfigValidationIssue(field string, severity AuthConfigValidationSeverity, message string, ) *AuthConfigValidationIssue`

NewAuthConfigValidationIssue instantiates a new AuthConfigValidationIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthConfigValidationIssueWithDefaults

`func NewAuthConfigValidationIssueWithDefaults() *AuthConfigValidationIssue`

NewAuthConfigValidationIssueWithDefaults instantiates a new AuthConfigValidationIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *AuthConfigValidationIssue) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *AuthConfigValidationIssue) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *AuthConfigValidationIssue) SetField(v string)`

SetField sets Field field to given value.


### GetSeverity

`func (o *AuthConfigValidationIssue) GetSeverity() AuthConfigValidationSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AuthConfigValidationIssue) GetSeverityOk() (*AuthConfigValidationSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AuthConfigValidationIssue) SetSeverity(v AuthConfigValidationSeverity)`

SetSeverity sets Severity field to given value.


### GetMessage

`func (o *AuthConfigValidationIssue) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AuthConfigValidationIssue) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AuthConfigValidationIssue) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


