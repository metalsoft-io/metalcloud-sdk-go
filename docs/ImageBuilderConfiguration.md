# ImageBuilderConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageDownloadHttpProxy** | Pointer to **string** | Proxy for http:// image/asset downloads (e.g. http://proxy.example.com:8888). Empty or omitted means no proxy (direct download). | [optional] 
**ImageDownloadHttpsProxy** | Pointer to **string** | Proxy for https:// image/asset downloads (e.g. http://proxy.example.com:8888). Empty or omitted means no proxy (direct download). | [optional] 

## Methods

### NewImageBuilderConfiguration

`func NewImageBuilderConfiguration() *ImageBuilderConfiguration`

NewImageBuilderConfiguration instantiates a new ImageBuilderConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageBuilderConfigurationWithDefaults

`func NewImageBuilderConfigurationWithDefaults() *ImageBuilderConfiguration`

NewImageBuilderConfigurationWithDefaults instantiates a new ImageBuilderConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageDownloadHttpProxy

`func (o *ImageBuilderConfiguration) GetImageDownloadHttpProxy() string`

GetImageDownloadHttpProxy returns the ImageDownloadHttpProxy field if non-nil, zero value otherwise.

### GetImageDownloadHttpProxyOk

`func (o *ImageBuilderConfiguration) GetImageDownloadHttpProxyOk() (*string, bool)`

GetImageDownloadHttpProxyOk returns a tuple with the ImageDownloadHttpProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDownloadHttpProxy

`func (o *ImageBuilderConfiguration) SetImageDownloadHttpProxy(v string)`

SetImageDownloadHttpProxy sets ImageDownloadHttpProxy field to given value.

### HasImageDownloadHttpProxy

`func (o *ImageBuilderConfiguration) HasImageDownloadHttpProxy() bool`

HasImageDownloadHttpProxy returns a boolean if a field has been set.

### GetImageDownloadHttpsProxy

`func (o *ImageBuilderConfiguration) GetImageDownloadHttpsProxy() string`

GetImageDownloadHttpsProxy returns the ImageDownloadHttpsProxy field if non-nil, zero value otherwise.

### GetImageDownloadHttpsProxyOk

`func (o *ImageBuilderConfiguration) GetImageDownloadHttpsProxyOk() (*string, bool)`

GetImageDownloadHttpsProxyOk returns a tuple with the ImageDownloadHttpsProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDownloadHttpsProxy

`func (o *ImageBuilderConfiguration) SetImageDownloadHttpsProxy(v string)`

SetImageDownloadHttpsProxy sets ImageDownloadHttpsProxy field to given value.

### HasImageDownloadHttpsProxy

`func (o *ImageBuilderConfiguration) HasImageDownloadHttpsProxy() bool`

HasImageDownloadHttpsProxy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


