# LogicalNetworkProfile

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

### NewLogicalNetworkProfile

`func NewLogicalNetworkProfile(id float32, revision float32, logicalNetworkType string, ) *LogicalNetworkProfile`

NewLogicalNetworkProfile instantiates a new LogicalNetworkProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogicalNetworkProfileWithDefaults

`func NewLogicalNetworkProfileWithDefaults() *LogicalNetworkProfile`

NewLogicalNetworkProfileWithDefaults instantiates a new LogicalNetworkProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LogicalNetworkProfile) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogicalNetworkProfile) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogicalNetworkProfile) SetId(v float32)`

SetId sets Id field to given value.


### GetRevision

`func (o *LogicalNetworkProfile) GetRevision() float32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *LogicalNetworkProfile) GetRevisionOk() (*float32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *LogicalNetworkProfile) SetRevision(v float32)`

SetRevision sets Revision field to given value.


### GetLabel

`func (o *LogicalNetworkProfile) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LogicalNetworkProfile) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LogicalNetworkProfile) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *LogicalNetworkProfile) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *LogicalNetworkProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogicalNetworkProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogicalNetworkProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LogicalNetworkProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *LogicalNetworkProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LogicalNetworkProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LogicalNetworkProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LogicalNetworkProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAnnotations

`func (o *LogicalNetworkProfile) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *LogicalNetworkProfile) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *LogicalNetworkProfile) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *LogicalNetworkProfile) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetLogicalNetworkType

`func (o *LogicalNetworkProfile) GetLogicalNetworkType() string`

GetLogicalNetworkType returns the LogicalNetworkType field if non-nil, zero value otherwise.

### GetLogicalNetworkTypeOk

`func (o *LogicalNetworkProfile) GetLogicalNetworkTypeOk() (*string, bool)`

GetLogicalNetworkTypeOk returns a tuple with the LogicalNetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalNetworkType

`func (o *LogicalNetworkProfile) SetLogicalNetworkType(v string)`

SetLogicalNetworkType sets LogicalNetworkType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


