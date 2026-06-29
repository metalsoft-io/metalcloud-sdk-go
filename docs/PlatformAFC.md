# PlatformAFC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminWarningsEmail** | Pointer to **[]string** | Email addresses for AFC admin warning notifications | [optional] 
**NoErrorEmailsForFunctionNames** | Pointer to **[]string** | AFC function names excluded from error email notifications | [optional] 
**NoAfcEventsForFunctionNames** | Pointer to **[]string** | AFC function names excluded from event generation | [optional] 

## Methods

### NewPlatformAFC

`func NewPlatformAFC() *PlatformAFC`

NewPlatformAFC instantiates a new PlatformAFC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformAFCWithDefaults

`func NewPlatformAFCWithDefaults() *PlatformAFC`

NewPlatformAFCWithDefaults instantiates a new PlatformAFC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminWarningsEmail

`func (o *PlatformAFC) GetAdminWarningsEmail() []string`

GetAdminWarningsEmail returns the AdminWarningsEmail field if non-nil, zero value otherwise.

### GetAdminWarningsEmailOk

`func (o *PlatformAFC) GetAdminWarningsEmailOk() (*[]string, bool)`

GetAdminWarningsEmailOk returns a tuple with the AdminWarningsEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminWarningsEmail

`func (o *PlatformAFC) SetAdminWarningsEmail(v []string)`

SetAdminWarningsEmail sets AdminWarningsEmail field to given value.

### HasAdminWarningsEmail

`func (o *PlatformAFC) HasAdminWarningsEmail() bool`

HasAdminWarningsEmail returns a boolean if a field has been set.

### GetNoErrorEmailsForFunctionNames

`func (o *PlatformAFC) GetNoErrorEmailsForFunctionNames() []string`

GetNoErrorEmailsForFunctionNames returns the NoErrorEmailsForFunctionNames field if non-nil, zero value otherwise.

### GetNoErrorEmailsForFunctionNamesOk

`func (o *PlatformAFC) GetNoErrorEmailsForFunctionNamesOk() (*[]string, bool)`

GetNoErrorEmailsForFunctionNamesOk returns a tuple with the NoErrorEmailsForFunctionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoErrorEmailsForFunctionNames

`func (o *PlatformAFC) SetNoErrorEmailsForFunctionNames(v []string)`

SetNoErrorEmailsForFunctionNames sets NoErrorEmailsForFunctionNames field to given value.

### HasNoErrorEmailsForFunctionNames

`func (o *PlatformAFC) HasNoErrorEmailsForFunctionNames() bool`

HasNoErrorEmailsForFunctionNames returns a boolean if a field has been set.

### GetNoAfcEventsForFunctionNames

`func (o *PlatformAFC) GetNoAfcEventsForFunctionNames() []string`

GetNoAfcEventsForFunctionNames returns the NoAfcEventsForFunctionNames field if non-nil, zero value otherwise.

### GetNoAfcEventsForFunctionNamesOk

`func (o *PlatformAFC) GetNoAfcEventsForFunctionNamesOk() (*[]string, bool)`

GetNoAfcEventsForFunctionNamesOk returns a tuple with the NoAfcEventsForFunctionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAfcEventsForFunctionNames

`func (o *PlatformAFC) SetNoAfcEventsForFunctionNames(v []string)`

SetNoAfcEventsForFunctionNames sets NoAfcEventsForFunctionNames field to given value.

### HasNoAfcEventsForFunctionNames

`func (o *PlatformAFC) HasNoAfcEventsForFunctionNames() bool`

HasNoAfcEventsForFunctionNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


