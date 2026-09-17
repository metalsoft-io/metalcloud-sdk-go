# LdapTestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | True when the supplied credentials authenticated successfully against the directory | 
**Message** | **string** | Human-readable description of the test outcome | 

## Methods

### NewLdapTestResult

`func NewLdapTestResult(success bool, message string, ) *LdapTestResult`

NewLdapTestResult instantiates a new LdapTestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapTestResultWithDefaults

`func NewLdapTestResultWithDefaults() *LdapTestResult`

NewLdapTestResultWithDefaults instantiates a new LdapTestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *LdapTestResult) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *LdapTestResult) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *LdapTestResult) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetMessage

`func (o *LdapTestResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LdapTestResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LdapTestResult) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


