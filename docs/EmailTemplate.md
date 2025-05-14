# EmailTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The email template ID | 
**Revision** | **float32** | Revision number | 
**Name** | **string** | Email template name | 
**Subject** | **string** | Email template subject | 
**Description** | Pointer to **string** | Email template description | [optional] 
**Text** | **string** | Email template text | 
**Html** | **string** | Email template html | 

## Methods

### NewEmailTemplate

`func NewEmailTemplate(id int32, revision float32, name string, subject string, text string, html string, ) *EmailTemplate`

NewEmailTemplate instantiates a new EmailTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailTemplateWithDefaults

`func NewEmailTemplateWithDefaults() *EmailTemplate`

NewEmailTemplateWithDefaults instantiates a new EmailTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmailTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailTemplate) SetId(v int32)`

SetId sets Id field to given value.


### GetRevision

`func (o *EmailTemplate) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *EmailTemplate) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *EmailTemplate) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetName

`func (o *EmailTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetSubject

`func (o *EmailTemplate) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailTemplate) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailTemplate) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetDescription

`func (o *EmailTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EmailTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EmailTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EmailTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetText

`func (o *EmailTemplate) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *EmailTemplate) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *EmailTemplate) SetText(v string)`

SetText sets Text field to given value.


### GetHtml

`func (o *EmailTemplate) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *EmailTemplate) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *EmailTemplate) SetHtml(v string)`

SetHtml sets Html field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


