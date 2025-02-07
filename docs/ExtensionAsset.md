# ExtensionAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Label of the asset. | 
**Name** | **string** | Name of the asset. | 
**AssetType** | **string** | Type of the asset. | 
**Url** | **string** | URL of the asset. | 
**RequiredAssets** | Pointer to **[]string** | Required assets by this asset. | [optional] 

## Methods

### NewExtensionAsset

`func NewExtensionAsset(label string, name string, assetType string, url string, ) *ExtensionAsset`

NewExtensionAsset instantiates a new ExtensionAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionAssetWithDefaults

`func NewExtensionAssetWithDefaults() *ExtensionAsset`

NewExtensionAssetWithDefaults instantiates a new ExtensionAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ExtensionAsset) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ExtensionAsset) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ExtensionAsset) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *ExtensionAsset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtensionAsset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtensionAsset) SetName(v string)`

SetName sets Name field to given value.


### GetAssetType

`func (o *ExtensionAsset) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *ExtensionAsset) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *ExtensionAsset) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.


### GetUrl

`func (o *ExtensionAsset) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ExtensionAsset) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ExtensionAsset) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetRequiredAssets

`func (o *ExtensionAsset) GetRequiredAssets() []string`

GetRequiredAssets returns the RequiredAssets field if non-nil, zero value otherwise.

### GetRequiredAssetsOk

`func (o *ExtensionAsset) GetRequiredAssetsOk() (*[]string, bool)`

GetRequiredAssetsOk returns a tuple with the RequiredAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredAssets

`func (o *ExtensionAsset) SetRequiredAssets(v []string)`

SetRequiredAssets sets RequiredAssets field to given value.

### HasRequiredAssets

`func (o *ExtensionAsset) HasRequiredAssets() bool`

HasRequiredAssets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


