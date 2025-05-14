# EmailTemplateCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Email template name | 
**Subject** | **string** | Email template subject | 
**Description** | Pointer to **string** | Email template description | [optional] 
**Text** | **string** | Email template text | 
**Html** | **string** | Email template html | 

## Methods

### NewEmailTemplateCreate

`func NewEmailTemplateCreate(name string, subject string, text string, html string, ) *EmailTemplateCreate`

NewEmailTemplateCreate instantiates a new EmailTemplateCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailTemplateCreateWithDefaults

`func NewEmailTemplateCreateWithDefaults() *EmailTemplateCreate`

NewEmailTemplateCreateWithDefaults instantiates a new EmailTemplateCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EmailTemplateCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailTemplateCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailTemplateCreate) SetName(v string)`

SetName sets Name field to given value.


### GetSubject

`func (o *EmailTemplateCreate) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailTemplateCreate) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailTemplateCreate) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetDescription

`func (o *EmailTemplateCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EmailTemplateCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EmailTemplateCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EmailTemplateCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetText

`func (o *EmailTemplateCreate) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *EmailTemplateCreate) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *EmailTemplateCreate) SetText(v string)`

SetText sets Text field to given value.


### GetHtml

`func (o *EmailTemplateCreate) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *EmailTemplateCreate) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *EmailTemplateCreate) SetHtml(v string)`

SetHtml sets Html field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


