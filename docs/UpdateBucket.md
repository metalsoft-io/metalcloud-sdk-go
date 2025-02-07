# UpdateBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeGB** | Pointer to **float32** | Disk size in GB for Bucket | [optional] 
**Label** | Pointer to **string** | Label of the Bucket. | [optional] 
**GuiSettings** | Pointer to [**GenericGUISettings**](GenericGUISettings.md) |  | [optional] 

## Methods

### NewUpdateBucket

`func NewUpdateBucket() *UpdateBucket`

NewUpdateBucket instantiates a new UpdateBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBucketWithDefaults

`func NewUpdateBucketWithDefaults() *UpdateBucket`

NewUpdateBucketWithDefaults instantiates a new UpdateBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeGB

`func (o *UpdateBucket) GetSizeGB() float32`

GetSizeGB returns the SizeGB field if non-nil, zero value otherwise.

### GetSizeGBOk

`func (o *UpdateBucket) GetSizeGBOk() (*float32, bool)`

GetSizeGBOk returns a tuple with the SizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGB

`func (o *UpdateBucket) SetSizeGB(v float32)`

SetSizeGB sets SizeGB field to given value.

### HasSizeGB

`func (o *UpdateBucket) HasSizeGB() bool`

HasSizeGB returns a boolean if a field has been set.

### GetLabel

`func (o *UpdateBucket) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateBucket) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateBucket) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpdateBucket) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetGuiSettings

`func (o *UpdateBucket) GetGuiSettings() GenericGUISettings`

GetGuiSettings returns the GuiSettings field if non-nil, zero value otherwise.

### GetGuiSettingsOk

`func (o *UpdateBucket) GetGuiSettingsOk() (*GenericGUISettings, bool)`

GetGuiSettingsOk returns a tuple with the GuiSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuiSettings

`func (o *UpdateBucket) SetGuiSettings(v GenericGUISettings)`

SetGuiSettings sets GuiSettings field to given value.

### HasGuiSettings

`func (o *UpdateBucket) HasGuiSettings() bool`

HasGuiSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


