# AuthConfigValidationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Valid** | **bool** | True when the stored configuration contains no errors | 
**Issues** | [**[]AuthConfigValidationIssue**](AuthConfigValidationIssue.md) | Validation issues found in the stored configuration; empty when valid | 

## Methods

### NewAuthConfigValidationResult

`func NewAuthConfigValidationResult(valid bool, issues []AuthConfigValidationIssue, ) *AuthConfigValidationResult`

NewAuthConfigValidationResult instantiates a new AuthConfigValidationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthConfigValidationResultWithDefaults

`func NewAuthConfigValidationResultWithDefaults() *AuthConfigValidationResult`

NewAuthConfigValidationResultWithDefaults instantiates a new AuthConfigValidationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValid

`func (o *AuthConfigValidationResult) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *AuthConfigValidationResult) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *AuthConfigValidationResult) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetIssues

`func (o *AuthConfigValidationResult) GetIssues() []AuthConfigValidationIssue`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *AuthConfigValidationResult) GetIssuesOk() (*[]AuthConfigValidationIssue, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *AuthConfigValidationResult) SetIssues(v []AuthConfigValidationIssue)`

SetIssues sets Issues field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


