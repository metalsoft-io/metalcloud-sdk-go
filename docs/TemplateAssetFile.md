# TemplateAssetFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | File name of the template asset file | 
**MimeType** | **string** | The MIME type of the template asset file | 
**Checksum** | Pointer to **string** | The checksums of the template asset file (sha1 or sha256) | [optional] 
**ContentBase64** | Pointer to **string** | The base64 encoded contents of the template asset file.                     Required if the URL is not provided | [optional] 
**TemplatingEngine** | **bool** | If the template asset file needs to be processed by a templating engine (Nunjucks) | 
**Url** | Pointer to **string** | The URL from where to fetch the template asset file | [optional] 
**Path** | **string** | The path is unique for the same OS template ID.         OSTemplateDeviceType &#39;SERVER&#39; uses:             - If the template asset &#39;file.usage&#39; is &#39;build_source_image&#39;, the &#39;file.path&#39; is ignored and &#39;file.url&#39; is required.             - If the template asset &#39;file.usage&#39; is &#39;build_component&#39;, the &#39;file.path&#39; is the relative ISO location path where the template asset will be copied to.         OSTemplateDeviceType &#39;NETWORK_DEVICE&#39; uses:             - If the template asset &#39;file.usage&#39; is &#39;source_image&#39;, the &#39;file.url&#39; is the location from where to fetch the Network Operating Systems (NOS) image and                the &#39;file.path&#39; is a composable part of the HTTP URL, expected as a request by the Open Network Install Environment (ONIE) install process.             - If the template asset &#39;file.usage&#39; is &#39;switch_ztp_config&#39;, the &#39;file.path&#39; is a composable part of the HTTP URL, expected as a request by the ZTP process.             - If the template asset &#39;file.usage&#39; is &#39;generic&#39;, the application code doesn&#39;t have any specific logic applied to the template asset. The &#39;file.path&#39;                is a composable part of the HTTP URL, expected as a request by the ZTP process (e.g. Enterprise_SONiC_OS).         OSTemplateDeviceType &#39;VM&#39; uses:             - If the template asset &#39;file.usage&#39; is &#39;metadata_source_image&#39;, the template asset contains information about the VM image. The &#39;file.path&#39; is used to                identify the VM image metadata (e.g. Incus).             - If the template asset &#39;file.usage&#39; is &#39;generic&#39;, the application code doesn&#39;t have any specific logic applied to the template asset. The &#39;file.path&#39; is               used to identify the role of each of the template asset (e.g. Incus cloud-init: network-config, user-data, vendor-data).          | 

## Methods

### NewTemplateAssetFile

`func NewTemplateAssetFile(name string, mimeType string, templatingEngine bool, path string, ) *TemplateAssetFile`

NewTemplateAssetFile instantiates a new TemplateAssetFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateAssetFileWithDefaults

`func NewTemplateAssetFileWithDefaults() *TemplateAssetFile`

NewTemplateAssetFileWithDefaults instantiates a new TemplateAssetFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TemplateAssetFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateAssetFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateAssetFile) SetName(v string)`

SetName sets Name field to given value.


### GetMimeType

`func (o *TemplateAssetFile) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *TemplateAssetFile) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *TemplateAssetFile) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.


### GetChecksum

`func (o *TemplateAssetFile) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *TemplateAssetFile) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *TemplateAssetFile) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *TemplateAssetFile) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetContentBase64

`func (o *TemplateAssetFile) GetContentBase64() string`

GetContentBase64 returns the ContentBase64 field if non-nil, zero value otherwise.

### GetContentBase64Ok

`func (o *TemplateAssetFile) GetContentBase64Ok() (*string, bool)`

GetContentBase64Ok returns a tuple with the ContentBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentBase64

`func (o *TemplateAssetFile) SetContentBase64(v string)`

SetContentBase64 sets ContentBase64 field to given value.

### HasContentBase64

`func (o *TemplateAssetFile) HasContentBase64() bool`

HasContentBase64 returns a boolean if a field has been set.

### GetTemplatingEngine

`func (o *TemplateAssetFile) GetTemplatingEngine() bool`

GetTemplatingEngine returns the TemplatingEngine field if non-nil, zero value otherwise.

### GetTemplatingEngineOk

`func (o *TemplateAssetFile) GetTemplatingEngineOk() (*bool, bool)`

GetTemplatingEngineOk returns a tuple with the TemplatingEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatingEngine

`func (o *TemplateAssetFile) SetTemplatingEngine(v bool)`

SetTemplatingEngine sets TemplatingEngine field to given value.


### GetUrl

`func (o *TemplateAssetFile) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TemplateAssetFile) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TemplateAssetFile) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TemplateAssetFile) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetPath

`func (o *TemplateAssetFile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *TemplateAssetFile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *TemplateAssetFile) SetPath(v string)`

SetPath sets Path field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


