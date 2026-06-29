# RenderedApplicableDeviceConfigurationTemplateProfilesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]RenderedApplicableDeviceConfigurationTemplateProfileItem**](RenderedApplicableDeviceConfigurationTemplateProfileItem.md) | Rendered output for profiles that would be applied on the next deploy, ordered by priority (ASC) then profile id (ASC). Excludes applyMode&#x3D;&#39;once&#39; profiles already recorded in device_configuration_template_history for the target device (those appear in &#x60;alreadyAppliedItems&#x60;). | 
**AlreadyAppliedItems** | [**[]RenderedApplicableDeviceConfigurationTemplateProfileItem**](RenderedApplicableDeviceConfigurationTemplateProfileItem.md) | Rendered output for applyMode&#x3D;&#39;once&#39; profiles that match the scope but are skipped on the next deploy because they already have a device_configuration_template_history row for the target device. Always empty when the query is fabric-scoped. | 
**Joined** | **string** | All &#x60;items[].rendered&#x60; outputs concatenated with &#39;\\n&#39; separators, in the same order. | 

## Methods

### NewRenderedApplicableDeviceConfigurationTemplateProfilesResult

`func NewRenderedApplicableDeviceConfigurationTemplateProfilesResult(items []RenderedApplicableDeviceConfigurationTemplateProfileItem, alreadyAppliedItems []RenderedApplicableDeviceConfigurationTemplateProfileItem, joined string, ) *RenderedApplicableDeviceConfigurationTemplateProfilesResult`

NewRenderedApplicableDeviceConfigurationTemplateProfilesResult instantiates a new RenderedApplicableDeviceConfigurationTemplateProfilesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenderedApplicableDeviceConfigurationTemplateProfilesResultWithDefaults

`func NewRenderedApplicableDeviceConfigurationTemplateProfilesResultWithDefaults() *RenderedApplicableDeviceConfigurationTemplateProfilesResult`

NewRenderedApplicableDeviceConfigurationTemplateProfilesResultWithDefaults instantiates a new RenderedApplicableDeviceConfigurationTemplateProfilesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *RenderedApplicableDeviceConfigurationTemplateProfilesResult) GetItems() []RenderedApplicableDeviceConfigurationTemplateProfileItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *RenderedApplicableDeviceConfigurationTemplateProfilesResult) GetItemsOk() (*[]RenderedApplicableDeviceConfigurationTemplateProfileItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *RenderedApplicableDeviceConfigurationTemplateProfilesResult) SetItems(v []RenderedApplicableDeviceConfigurationTemplateProfileItem)`

SetItems sets Items field to given value.


### GetAlreadyAppliedItems

`func (o *RenderedApplicableDeviceConfigurationTemplateProfilesResult) GetAlreadyAppliedItems() []RenderedApplicableDeviceConfigurationTemplateProfileItem`

GetAlreadyAppliedItems returns the AlreadyAppliedItems field if non-nil, zero value otherwise.

### GetAlreadyAppliedItemsOk

`func (o *RenderedApplicableDeviceConfigurationTemplateProfilesResult) GetAlreadyAppliedItemsOk() (*[]RenderedApplicableDeviceConfigurationTemplateProfileItem, bool)`

GetAlreadyAppliedItemsOk returns a tuple with the AlreadyAppliedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlreadyAppliedItems

`func (o *RenderedApplicableDeviceConfigurationTemplateProfilesResult) SetAlreadyAppliedItems(v []RenderedApplicableDeviceConfigurationTemplateProfileItem)`

SetAlreadyAppliedItems sets AlreadyAppliedItems field to given value.


### GetJoined

`func (o *RenderedApplicableDeviceConfigurationTemplateProfilesResult) GetJoined() string`

GetJoined returns the Joined field if non-nil, zero value otherwise.

### GetJoinedOk

`func (o *RenderedApplicableDeviceConfigurationTemplateProfilesResult) GetJoinedOk() (*string, bool)`

GetJoinedOk returns a tuple with the Joined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoined

`func (o *RenderedApplicableDeviceConfigurationTemplateProfilesResult) SetJoined(v string)`

SetJoined sets Joined field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


