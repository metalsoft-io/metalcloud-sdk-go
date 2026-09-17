# ExtensionOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Label of the output. | 
**Name** | **string** | Name of the output. | 
**HelpText** | Pointer to **string** | Help text describing the output. | [optional] 
**OutputType** | **string** | Type of the output to be used as UI hint. Well known values include: boolean, integer, string, base64, password, secret, text, url, object, unknown | 
**Template** | Pointer to **string** | Template of the output. | [optional] 
**IsPassword** | Pointer to **bool** | Flag to indicate if the output is a password/secret. | [optional] [default to false]

## Methods

### NewExtensionOutput

`func NewExtensionOutput(label string, name string, outputType string, ) *ExtensionOutput`

NewExtensionOutput instantiates a new ExtensionOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionOutputWithDefaults

`func NewExtensionOutputWithDefaults() *ExtensionOutput`

NewExtensionOutputWithDefaults instantiates a new ExtensionOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ExtensionOutput) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ExtensionOutput) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ExtensionOutput) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *ExtensionOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtensionOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtensionOutput) SetName(v string)`

SetName sets Name field to given value.


### GetHelpText

`func (o *ExtensionOutput) GetHelpText() string`

GetHelpText returns the HelpText field if non-nil, zero value otherwise.

### GetHelpTextOk

`func (o *ExtensionOutput) GetHelpTextOk() (*string, bool)`

GetHelpTextOk returns a tuple with the HelpText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpText

`func (o *ExtensionOutput) SetHelpText(v string)`

SetHelpText sets HelpText field to given value.

### HasHelpText

`func (o *ExtensionOutput) HasHelpText() bool`

HasHelpText returns a boolean if a field has been set.

### GetOutputType

`func (o *ExtensionOutput) GetOutputType() string`

GetOutputType returns the OutputType field if non-nil, zero value otherwise.

### GetOutputTypeOk

`func (o *ExtensionOutput) GetOutputTypeOk() (*string, bool)`

GetOutputTypeOk returns a tuple with the OutputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputType

`func (o *ExtensionOutput) SetOutputType(v string)`

SetOutputType sets OutputType field to given value.


### GetTemplate

`func (o *ExtensionOutput) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ExtensionOutput) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ExtensionOutput) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *ExtensionOutput) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetIsPassword

`func (o *ExtensionOutput) GetIsPassword() bool`

GetIsPassword returns the IsPassword field if non-nil, zero value otherwise.

### GetIsPasswordOk

`func (o *ExtensionOutput) GetIsPasswordOk() (*bool, bool)`

GetIsPasswordOk returns a tuple with the IsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPassword

`func (o *ExtensionOutput) SetIsPassword(v bool)`

SetIsPassword sets IsPassword field to given value.

### HasIsPassword

`func (o *ExtensionOutput) HasIsPassword() bool`

HasIsPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


