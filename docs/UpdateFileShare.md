# UpdateFileShare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeGB** | Pointer to **float32** | Disk size in GB for File Share | [optional] 
**Label** | Pointer to **string** | Label of the File Share. | [optional] 
**GuiSettings** | Pointer to [**GenericGUISettings**](GenericGUISettings.md) |  | [optional] 

## Methods

### NewUpdateFileShare

`func NewUpdateFileShare() *UpdateFileShare`

NewUpdateFileShare instantiates a new UpdateFileShare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFileShareWithDefaults

`func NewUpdateFileShareWithDefaults() *UpdateFileShare`

NewUpdateFileShareWithDefaults instantiates a new UpdateFileShare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeGB

`func (o *UpdateFileShare) GetSizeGB() float32`

GetSizeGB returns the SizeGB field if non-nil, zero value otherwise.

### GetSizeGBOk

`func (o *UpdateFileShare) GetSizeGBOk() (*float32, bool)`

GetSizeGBOk returns a tuple with the SizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGB

`func (o *UpdateFileShare) SetSizeGB(v float32)`

SetSizeGB sets SizeGB field to given value.

### HasSizeGB

`func (o *UpdateFileShare) HasSizeGB() bool`

HasSizeGB returns a boolean if a field has been set.

### GetLabel

`func (o *UpdateFileShare) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateFileShare) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateFileShare) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpdateFileShare) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetGuiSettings

`func (o *UpdateFileShare) GetGuiSettings() GenericGUISettings`

GetGuiSettings returns the GuiSettings field if non-nil, zero value otherwise.

### GetGuiSettingsOk

`func (o *UpdateFileShare) GetGuiSettingsOk() (*GenericGUISettings, bool)`

GetGuiSettingsOk returns a tuple with the GuiSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuiSettings

`func (o *UpdateFileShare) SetGuiSettings(v GenericGUISettings)`

SetGuiSettings sets GuiSettings field to given value.

### HasGuiSettings

`func (o *UpdateFileShare) HasGuiSettings() bool`

HasGuiSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


