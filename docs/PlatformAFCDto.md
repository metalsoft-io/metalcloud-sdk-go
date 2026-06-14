# PlatformAFCDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminWarningsEmail** | Pointer to **[]string** | Email addresses for AFC admin warning notifications | [optional] 
**NoErrorEmailsForFunctionNames** | Pointer to **[]string** | AFC function names excluded from error email notifications | [optional] 
**NoAfcEventsForFunctionNames** | Pointer to **[]string** | AFC function names excluded from event generation | [optional] 

## Methods

### NewPlatformAFCDto

`func NewPlatformAFCDto() *PlatformAFCDto`

NewPlatformAFCDto instantiates a new PlatformAFCDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformAFCDtoWithDefaults

`func NewPlatformAFCDtoWithDefaults() *PlatformAFCDto`

NewPlatformAFCDtoWithDefaults instantiates a new PlatformAFCDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminWarningsEmail

`func (o *PlatformAFCDto) GetAdminWarningsEmail() []string`

GetAdminWarningsEmail returns the AdminWarningsEmail field if non-nil, zero value otherwise.

### GetAdminWarningsEmailOk

`func (o *PlatformAFCDto) GetAdminWarningsEmailOk() (*[]string, bool)`

GetAdminWarningsEmailOk returns a tuple with the AdminWarningsEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminWarningsEmail

`func (o *PlatformAFCDto) SetAdminWarningsEmail(v []string)`

SetAdminWarningsEmail sets AdminWarningsEmail field to given value.

### HasAdminWarningsEmail

`func (o *PlatformAFCDto) HasAdminWarningsEmail() bool`

HasAdminWarningsEmail returns a boolean if a field has been set.

### GetNoErrorEmailsForFunctionNames

`func (o *PlatformAFCDto) GetNoErrorEmailsForFunctionNames() []string`

GetNoErrorEmailsForFunctionNames returns the NoErrorEmailsForFunctionNames field if non-nil, zero value otherwise.

### GetNoErrorEmailsForFunctionNamesOk

`func (o *PlatformAFCDto) GetNoErrorEmailsForFunctionNamesOk() (*[]string, bool)`

GetNoErrorEmailsForFunctionNamesOk returns a tuple with the NoErrorEmailsForFunctionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoErrorEmailsForFunctionNames

`func (o *PlatformAFCDto) SetNoErrorEmailsForFunctionNames(v []string)`

SetNoErrorEmailsForFunctionNames sets NoErrorEmailsForFunctionNames field to given value.

### HasNoErrorEmailsForFunctionNames

`func (o *PlatformAFCDto) HasNoErrorEmailsForFunctionNames() bool`

HasNoErrorEmailsForFunctionNames returns a boolean if a field has been set.

### GetNoAfcEventsForFunctionNames

`func (o *PlatformAFCDto) GetNoAfcEventsForFunctionNames() []string`

GetNoAfcEventsForFunctionNames returns the NoAfcEventsForFunctionNames field if non-nil, zero value otherwise.

### GetNoAfcEventsForFunctionNamesOk

`func (o *PlatformAFCDto) GetNoAfcEventsForFunctionNamesOk() (*[]string, bool)`

GetNoAfcEventsForFunctionNamesOk returns a tuple with the NoAfcEventsForFunctionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAfcEventsForFunctionNames

`func (o *PlatformAFCDto) SetNoAfcEventsForFunctionNames(v []string)`

SetNoAfcEventsForFunctionNames sets NoAfcEventsForFunctionNames field to given value.

### HasNoAfcEventsForFunctionNames

`func (o *PlatformAFCDto) HasNoAfcEventsForFunctionNames() bool`

HasNoAfcEventsForFunctionNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


