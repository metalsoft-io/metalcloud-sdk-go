# ExtensionAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Label of the asset. | 
**Name** | **string** | Name of the asset. | 
**AssetType** | [**ExtensionAssetType**](ExtensionAssetType.md) | Type of the asset. | 
**Url** | Pointer to **string** | URL of the asset. Required for AnsibleBundle assets. | [optional] 
**HostRegistry** | Pointer to **string** | Registry host for OCI image assets. | [optional] 
**PortRegistry** | Pointer to **float32** | Registry port for OCI image assets. | [optional] 
**NamespaceRegistry** | Pointer to **string** | Registry namespace for OCI image assets. | [optional] 
**RepositoryRegistry** | Pointer to **string** | Registry repository for OCI image assets. | [optional] 
**TagRegistry** | Pointer to **string** | Registry tag for OCI image assets. | [optional] 
**RequiredAssets** | Pointer to **[]string** | Required assets by this asset. | [optional] 

## Methods

### NewExtensionAsset

`func NewExtensionAsset(label string, name string, assetType ExtensionAssetType, ) *ExtensionAsset`

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

`func (o *ExtensionAsset) GetAssetType() ExtensionAssetType`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *ExtensionAsset) GetAssetTypeOk() (*ExtensionAssetType, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *ExtensionAsset) SetAssetType(v ExtensionAssetType)`

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

### HasUrl

`func (o *ExtensionAsset) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetHostRegistry

`func (o *ExtensionAsset) GetHostRegistry() string`

GetHostRegistry returns the HostRegistry field if non-nil, zero value otherwise.

### GetHostRegistryOk

`func (o *ExtensionAsset) GetHostRegistryOk() (*string, bool)`

GetHostRegistryOk returns a tuple with the HostRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostRegistry

`func (o *ExtensionAsset) SetHostRegistry(v string)`

SetHostRegistry sets HostRegistry field to given value.

### HasHostRegistry

`func (o *ExtensionAsset) HasHostRegistry() bool`

HasHostRegistry returns a boolean if a field has been set.

### GetPortRegistry

`func (o *ExtensionAsset) GetPortRegistry() float32`

GetPortRegistry returns the PortRegistry field if non-nil, zero value otherwise.

### GetPortRegistryOk

`func (o *ExtensionAsset) GetPortRegistryOk() (*float32, bool)`

GetPortRegistryOk returns a tuple with the PortRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRegistry

`func (o *ExtensionAsset) SetPortRegistry(v float32)`

SetPortRegistry sets PortRegistry field to given value.

### HasPortRegistry

`func (o *ExtensionAsset) HasPortRegistry() bool`

HasPortRegistry returns a boolean if a field has been set.

### GetNamespaceRegistry

`func (o *ExtensionAsset) GetNamespaceRegistry() string`

GetNamespaceRegistry returns the NamespaceRegistry field if non-nil, zero value otherwise.

### GetNamespaceRegistryOk

`func (o *ExtensionAsset) GetNamespaceRegistryOk() (*string, bool)`

GetNamespaceRegistryOk returns a tuple with the NamespaceRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceRegistry

`func (o *ExtensionAsset) SetNamespaceRegistry(v string)`

SetNamespaceRegistry sets NamespaceRegistry field to given value.

### HasNamespaceRegistry

`func (o *ExtensionAsset) HasNamespaceRegistry() bool`

HasNamespaceRegistry returns a boolean if a field has been set.

### GetRepositoryRegistry

`func (o *ExtensionAsset) GetRepositoryRegistry() string`

GetRepositoryRegistry returns the RepositoryRegistry field if non-nil, zero value otherwise.

### GetRepositoryRegistryOk

`func (o *ExtensionAsset) GetRepositoryRegistryOk() (*string, bool)`

GetRepositoryRegistryOk returns a tuple with the RepositoryRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryRegistry

`func (o *ExtensionAsset) SetRepositoryRegistry(v string)`

SetRepositoryRegistry sets RepositoryRegistry field to given value.

### HasRepositoryRegistry

`func (o *ExtensionAsset) HasRepositoryRegistry() bool`

HasRepositoryRegistry returns a boolean if a field has been set.

### GetTagRegistry

`func (o *ExtensionAsset) GetTagRegistry() string`

GetTagRegistry returns the TagRegistry field if non-nil, zero value otherwise.

### GetTagRegistryOk

`func (o *ExtensionAsset) GetTagRegistryOk() (*string, bool)`

GetTagRegistryOk returns a tuple with the TagRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagRegistry

`func (o *ExtensionAsset) SetTagRegistry(v string)`

SetTagRegistry sets TagRegistry field to given value.

### HasTagRegistry

`func (o *ExtensionAsset) HasTagRegistry() bool`

HasTagRegistry returns a boolean if a field has been set.

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


