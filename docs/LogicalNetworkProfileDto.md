# LogicalNetworkProfileDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | ID of the logical network profile | 
**Revision** | **float32** | Revision number of the logical network profile | 
**Label** | Pointer to **string** | Label for the logical network profile | [optional] 
**Name** | Pointer to **string** | Name of the logical network profile | [optional] 
**Description** | Pointer to **string** | Description of the logical network profile | [optional] 
**Annotations** | Pointer to **map[string]interface{}** | Annotations for the logical network profile | [optional] 
**LogicalNetworkType** | **string** | Type of the logical network profile | 

## Methods

### NewLogicalNetworkProfileDto

`func NewLogicalNetworkProfileDto(id float32, revision float32, logicalNetworkType string, ) *LogicalNetworkProfileDto`

NewLogicalNetworkProfileDto instantiates a new LogicalNetworkProfileDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogicalNetworkProfileDtoWithDefaults

`func NewLogicalNetworkProfileDtoWithDefaults() *LogicalNetworkProfileDto`

NewLogicalNetworkProfileDtoWithDefaults instantiates a new LogicalNetworkProfileDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LogicalNetworkProfileDto) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogicalNetworkProfileDto) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogicalNetworkProfileDto) SetId(v float32)`

SetId sets Id field to given value.


### GetRevision

`func (o *LogicalNetworkProfileDto) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *LogicalNetworkProfileDto) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *LogicalNetworkProfileDto) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetLabel

`func (o *LogicalNetworkProfileDto) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LogicalNetworkProfileDto) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LogicalNetworkProfileDto) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *LogicalNetworkProfileDto) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *LogicalNetworkProfileDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogicalNetworkProfileDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogicalNetworkProfileDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LogicalNetworkProfileDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *LogicalNetworkProfileDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LogicalNetworkProfileDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LogicalNetworkProfileDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LogicalNetworkProfileDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAnnotations

`func (o *LogicalNetworkProfileDto) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *LogicalNetworkProfileDto) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *LogicalNetworkProfileDto) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *LogicalNetworkProfileDto) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetLogicalNetworkType

`func (o *LogicalNetworkProfileDto) GetLogicalNetworkType() string`

GetLogicalNetworkType returns the LogicalNetworkType field if non-nil, zero value otherwise.

### GetLogicalNetworkTypeOk

`func (o *LogicalNetworkProfileDto) GetLogicalNetworkTypeOk() (*string, bool)`

GetLogicalNetworkTypeOk returns a tuple with the LogicalNetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkType

`func (o *LogicalNetworkProfileDto) SetLogicalNetworkType(v string)`

SetLogicalNetworkType sets LogicalNetworkType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


