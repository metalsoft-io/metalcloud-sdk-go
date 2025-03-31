# Extension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | The extension ID | 
**Revision** | **float32** | Revision number | 
**Slug** | Pointer to **string** | The extension unique slug | [optional] 
**Name** | **string** | The extension name | 
**Label** | Pointer to **string** | The extension unique label | [optional] 
**Description** | **string** | The extension description | 
**Status** | **string** | Extension status | 
**Kind** | **string** | Extension kind | 
**Definition** | [**ExtensionDefinition**](ExtensionDefinition.md) |  | 
**Links** | Pointer to [**[]Link**](Link.md) | Reference links | [optional] 

## Methods

### NewExtension

`func NewExtension(id float32, revision float32, name string, description string, status string, kind string, definition ExtensionDefinition, ) *Extension`

NewExtension instantiates a new Extension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionWithDefaults

`func NewExtensionWithDefaults() *Extension`

NewExtensionWithDefaults instantiates a new Extension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Extension) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Extension) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Extension) SetId(v float32)`

SetId sets Id field to given value.


### GetRevision

`func (o *Extension) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *Extension) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *Extension) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetSlug

`func (o *Extension) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Extension) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Extension) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *Extension) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetName

`func (o *Extension) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Extension) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Extension) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *Extension) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Extension) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Extension) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Extension) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *Extension) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Extension) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Extension) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetStatus

`func (o *Extension) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Extension) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Extension) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetKind

`func (o *Extension) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Extension) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Extension) SetKind(v string)`

SetKind sets Kind field to given value.


### GetDefinition

`func (o *Extension) GetDefinition() ExtensionDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *Extension) GetDefinitionOk() (*ExtensionDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *Extension) SetDefinition(v ExtensionDefinition)`

SetDefinition sets Definition field to given value.


### GetLinks

`func (o *Extension) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Extension) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Extension) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Extension) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


