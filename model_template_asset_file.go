/*
MetalSoft REST API

MetalSoft REST API documentation

API version: 2.0
Contact: support@metalsoft.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
	"fmt"
)

// checks if the TemplateAssetFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateAssetFile{}

// TemplateAssetFile struct for TemplateAssetFile
type TemplateAssetFile struct {
	// File name of the template asset file
	Name string `json:"name"`
	// The MIME type of the template asset file
	MimeType string `json:"mimeType"`
	// The checksums of the template asset file (sha1 or sha256)
	Checksum *string `json:"checksum,omitempty"`
	// The base64 encoded contents of the template asset file.                     Required if the URL is not provided
	ContentBase64 *string `json:"contentBase64,omitempty"`
	// If the template asset file needs to be processed by a templating engine (Nunjucks)
	TemplatingEngine bool `json:"templatingEngine"`
	// The URL from where to fetch the template asset file
	Url *string `json:"url,omitempty"`
	// The path is unique for the same OS template ID.         OSTemplateDeviceType 'SERVER' uses:             - If the template asset 'file.usage' is 'build_source_image', the 'file.path' is ignored and 'file.url' is required.             - If the template asset 'file.usage' is 'build_component', the 'file.path' is the relative ISO location path where the template asset will be copied to.             - If the template asset 'file.usage' is 'secondary_image', the 'file.url' is the URL of the image which to mount in the 2nd virtual media of the server.         OSTemplateDeviceType 'NETWORK_DEVICE' uses:             - If the template asset 'file.usage' is 'source_image', the 'file.url' is the location from where to fetch the Network Operating Systems (NOS) image and                the 'file.path' is a composable part of the HTTP URL, expected as a request by the Open Network Install Environment (ONIE) install process.             - If the template asset 'file.usage' is 'switch_ztp_config', the 'file.path' is a composable part of the HTTP URL, expected as a request by the ZTP process.             - If the template asset 'file.usage' is 'generic', the application code doesn't have any specific logic applied to the template asset. The 'file.path'                is a composable part of the HTTP URL, expected as a request by the ZTP process (e.g. Enterprise_SONiC_OS).         OSTemplateDeviceType 'VM' uses:             - If the template asset 'file.usage' is 'metadata_source_image', the template asset contains information about the VM image. The 'file.path' is used to                identify the VM image metadata (e.g. Incus).             - If the template asset 'file.usage' is 'generic', the application code doesn't have any specific logic applied to the template asset. The 'file.path' is               used to identify the role of each of the template asset (e.g. Incus cloud-init: network-config, user-data, vendor-data).         
	Path string `json:"path" validate:"regexp=^\\/([a-zA-Z0-9._-]+\\/){0,19}[a-zA-Z0-9._-]+$"`
	AdditionalProperties map[string]interface{}
}

type _TemplateAssetFile TemplateAssetFile

// NewTemplateAssetFile instantiates a new TemplateAssetFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateAssetFile(name string, mimeType string, templatingEngine bool, path string) *TemplateAssetFile {
	this := TemplateAssetFile{}
	this.Name = name
	this.MimeType = mimeType
	this.TemplatingEngine = templatingEngine
	this.Path = path
	return &this
}

// NewTemplateAssetFileWithDefaults instantiates a new TemplateAssetFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateAssetFileWithDefaults() *TemplateAssetFile {
	this := TemplateAssetFile{}
	return &this
}

// GetName returns the Name field value
func (o *TemplateAssetFile) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TemplateAssetFile) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TemplateAssetFile) SetName(v string) {
	o.Name = v
}

// GetMimeType returns the MimeType field value
func (o *TemplateAssetFile) GetMimeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value
// and a boolean to check if the value has been set.
func (o *TemplateAssetFile) GetMimeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MimeType, true
}

// SetMimeType sets field value
func (o *TemplateAssetFile) SetMimeType(v string) {
	o.MimeType = v
}

// GetChecksum returns the Checksum field value if set, zero value otherwise.
func (o *TemplateAssetFile) GetChecksum() string {
	if o == nil || IsNil(o.Checksum) {
		var ret string
		return ret
	}
	return *o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateAssetFile) GetChecksumOk() (*string, bool) {
	if o == nil || IsNil(o.Checksum) {
		return nil, false
	}
	return o.Checksum, true
}

// HasChecksum returns a boolean if a field has been set.
func (o *TemplateAssetFile) HasChecksum() bool {
	if o != nil && !IsNil(o.Checksum) {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given string and assigns it to the Checksum field.
func (o *TemplateAssetFile) SetChecksum(v string) {
	o.Checksum = &v
}

// GetContentBase64 returns the ContentBase64 field value if set, zero value otherwise.
func (o *TemplateAssetFile) GetContentBase64() string {
	if o == nil || IsNil(o.ContentBase64) {
		var ret string
		return ret
	}
	return *o.ContentBase64
}

// GetContentBase64Ok returns a tuple with the ContentBase64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateAssetFile) GetContentBase64Ok() (*string, bool) {
	if o == nil || IsNil(o.ContentBase64) {
		return nil, false
	}
	return o.ContentBase64, true
}

// HasContentBase64 returns a boolean if a field has been set.
func (o *TemplateAssetFile) HasContentBase64() bool {
	if o != nil && !IsNil(o.ContentBase64) {
		return true
	}

	return false
}

// SetContentBase64 gets a reference to the given string and assigns it to the ContentBase64 field.
func (o *TemplateAssetFile) SetContentBase64(v string) {
	o.ContentBase64 = &v
}

// GetTemplatingEngine returns the TemplatingEngine field value
func (o *TemplateAssetFile) GetTemplatingEngine() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.TemplatingEngine
}

// GetTemplatingEngineOk returns a tuple with the TemplatingEngine field value
// and a boolean to check if the value has been set.
func (o *TemplateAssetFile) GetTemplatingEngineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplatingEngine, true
}

// SetTemplatingEngine sets field value
func (o *TemplateAssetFile) SetTemplatingEngine(v bool) {
	o.TemplatingEngine = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *TemplateAssetFile) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateAssetFile) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *TemplateAssetFile) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *TemplateAssetFile) SetUrl(v string) {
	o.Url = &v
}

// GetPath returns the Path field value
func (o *TemplateAssetFile) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *TemplateAssetFile) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *TemplateAssetFile) SetPath(v string) {
	o.Path = v
}

func (o TemplateAssetFile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateAssetFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["mimeType"] = o.MimeType
	if !IsNil(o.Checksum) {
		toSerialize["checksum"] = o.Checksum
	}
	if !IsNil(o.ContentBase64) {
		toSerialize["contentBase64"] = o.ContentBase64
	}
	toSerialize["templatingEngine"] = o.TemplatingEngine
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	toSerialize["path"] = o.Path

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TemplateAssetFile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"mimeType",
		"templatingEngine",
		"path",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTemplateAssetFile := _TemplateAssetFile{}

	err = json.Unmarshal(data, &varTemplateAssetFile)

	if err != nil {
		return err
	}

	*o = TemplateAssetFile(varTemplateAssetFile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "mimeType")
		delete(additionalProperties, "checksum")
		delete(additionalProperties, "contentBase64")
		delete(additionalProperties, "templatingEngine")
		delete(additionalProperties, "url")
		delete(additionalProperties, "path")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTemplateAssetFile struct {
	value *TemplateAssetFile
	isSet bool
}

func (v NullableTemplateAssetFile) Get() *TemplateAssetFile {
	return v.value
}

func (v *NullableTemplateAssetFile) Set(val *TemplateAssetFile) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateAssetFile) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateAssetFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateAssetFile(val *TemplateAssetFile) *NullableTemplateAssetFile {
	return &NullableTemplateAssetFile{value: val, isSet: true}
}

func (v NullableTemplateAssetFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateAssetFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


