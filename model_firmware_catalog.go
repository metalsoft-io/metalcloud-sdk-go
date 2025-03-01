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

// checks if the FirmwareCatalog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirmwareCatalog{}

// FirmwareCatalog struct for FirmwareCatalog
type FirmwareCatalog struct {
	// Unique identifier for the firmware catalog
	ServerFirmwareCatalogId float32 `json:"serverFirmwareCatalogId"`
	// Name of the catalog, must be unique
	ServerFirmwareCatalogName string `json:"serverFirmwareCatalogName"`
	// User description of the catalog
	ServerFirmwareCatalogDescription *string `json:"serverFirmwareCatalogDescription,omitempty"`
	// Firmware catalog vendor: dell, hp, lenovo
	ServerFirmwareCatalogVendor string `json:"serverFirmwareCatalogVendor"`
	// Vendor identifier for the catalog
	ServerFirmwareCatalogVendorId *string `json:"serverFirmwareCatalogVendorId,omitempty"`
	// Vendor URL for the firmware catalog
	ServerFirmwareCatalogVendorUrl *string `json:"serverFirmwareCatalogVendorUrl,omitempty"`
	// Vendor release timestamp for the catalog
	ServerFirmwareCatalogVendorReleaseTimestamp *string `json:"serverFirmwareCatalogVendorReleaseTimestamp,omitempty"`
	// Catalog update type: online or offline
	ServerFirmwareCatalogUpdateType string `json:"serverFirmwareCatalogUpdateType"`
	// List of supported Metalsoft server types for which firmware binaries are available
	ServerFirmwareCatalogMetalsoftServerTypesSupportedJson map[string]interface{} `json:"serverFirmwareCatalogMetalsoftServerTypesSupportedJson,omitempty"`
	// List of supported server types for which firmware binaries are available
	ServerFirmwareCatalogVendorServerTypesSupportedJson map[string]interface{} `json:"serverFirmwareCatalogVendorServerTypesSupportedJson,omitempty"`
	// Vendor configuration in JSON format
	ServerFirmwareCatalogVendorConfigurationJson map[string]interface{} `json:"serverFirmwareCatalogVendorConfigurationJson,omitempty"`
	// Timestamp when the catalog was created
	ServerFirmwareCatalogCreatedTimestamp string `json:"serverFirmwareCatalogCreatedTimestamp"`
	// Links to other resources
	Links map[string]interface{} `json:"links"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareCatalog FirmwareCatalog

// NewFirmwareCatalog instantiates a new FirmwareCatalog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareCatalog(serverFirmwareCatalogId float32, serverFirmwareCatalogName string, serverFirmwareCatalogVendor string, serverFirmwareCatalogUpdateType string, serverFirmwareCatalogCreatedTimestamp string, links map[string]interface{}) *FirmwareCatalog {
	this := FirmwareCatalog{}
	this.ServerFirmwareCatalogId = serverFirmwareCatalogId
	this.ServerFirmwareCatalogName = serverFirmwareCatalogName
	this.ServerFirmwareCatalogVendor = serverFirmwareCatalogVendor
	this.ServerFirmwareCatalogUpdateType = serverFirmwareCatalogUpdateType
	this.ServerFirmwareCatalogCreatedTimestamp = serverFirmwareCatalogCreatedTimestamp
	this.Links = links
	return &this
}

// NewFirmwareCatalogWithDefaults instantiates a new FirmwareCatalog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareCatalogWithDefaults() *FirmwareCatalog {
	this := FirmwareCatalog{}
	var serverFirmwareCatalogUpdateType string = "online"
	this.ServerFirmwareCatalogUpdateType = serverFirmwareCatalogUpdateType
	return &this
}

// GetServerFirmwareCatalogId returns the ServerFirmwareCatalogId field value
func (o *FirmwareCatalog) GetServerFirmwareCatalogId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ServerFirmwareCatalogId
}

// GetServerFirmwareCatalogIdOk returns a tuple with the ServerFirmwareCatalogId field value
// and a boolean to check if the value has been set.
func (o *FirmwareCatalog) GetServerFirmwareCatalogIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerFirmwareCatalogId, true
}

// SetServerFirmwareCatalogId sets field value
func (o *FirmwareCatalog) SetServerFirmwareCatalogId(v float32) {
	o.ServerFirmwareCatalogId = v
}

// GetServerFirmwareCatalogName returns the ServerFirmwareCatalogName field value
func (o *FirmwareCatalog) GetServerFirmwareCatalogName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerFirmwareCatalogName
}

// GetServerFirmwareCatalogNameOk returns a tuple with the ServerFirmwareCatalogName field value
// and a boolean to check if the value has been set.
func (o *FirmwareCatalog) GetServerFirmwareCatalogNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerFirmwareCatalogName, true
}

// SetServerFirmwareCatalogName sets field value
func (o *FirmwareCatalog) SetServerFirmwareCatalogName(v string) {
	o.ServerFirmwareCatalogName = v
}

// GetServerFirmwareCatalogDescription returns the ServerFirmwareCatalogDescription field value if set, zero value otherwise.
func (o *FirmwareCatalog) GetServerFirmwareCatalogDescription() string {
	if o == nil || IsNil(o.ServerFirmwareCatalogDescription) {
		var ret string
		return ret
	}
	return *o.ServerFirmwareCatalogDescription
}

// GetServerFirmwareCatalogDescriptionOk returns a tuple with the ServerFirmwareCatalogDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareCatalog) GetServerFirmwareCatalogDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ServerFirmwareCatalogDescription) {
		return nil, false
	}
	return o.ServerFirmwareCatalogDescription, true
}

// HasServerFirmwareCatalogDescription returns a boolean if a field has been set.
func (o *FirmwareCatalog) HasServerFirmwareCatalogDescription() bool {
	if o != nil && !IsNil(o.ServerFirmwareCatalogDescription) {
		return true
	}

	return false
}

// SetServerFirmwareCatalogDescription gets a reference to the given string and assigns it to the ServerFirmwareCatalogDescription field.
func (o *FirmwareCatalog) SetServerFirmwareCatalogDescription(v string) {
	o.ServerFirmwareCatalogDescription = &v
}

// GetServerFirmwareCatalogVendor returns the ServerFirmwareCatalogVendor field value
func (o *FirmwareCatalog) GetServerFirmwareCatalogVendor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerFirmwareCatalogVendor
}

// GetServerFirmwareCatalogVendorOk returns a tuple with the ServerFirmwareCatalogVendor field value
// and a boolean to check if the value has been set.
func (o *FirmwareCatalog) GetServerFirmwareCatalogVendorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerFirmwareCatalogVendor, true
}

// SetServerFirmwareCatalogVendor sets field value
func (o *FirmwareCatalog) SetServerFirmwareCatalogVendor(v string) {
	o.ServerFirmwareCatalogVendor = v
}

// GetServerFirmwareCatalogVendorId returns the ServerFirmwareCatalogVendorId field value if set, zero value otherwise.
func (o *FirmwareCatalog) GetServerFirmwareCatalogVendorId() string {
	if o == nil || IsNil(o.ServerFirmwareCatalogVendorId) {
		var ret string
		return ret
	}
	return *o.ServerFirmwareCatalogVendorId
}

// GetServerFirmwareCatalogVendorIdOk returns a tuple with the ServerFirmwareCatalogVendorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareCatalog) GetServerFirmwareCatalogVendorIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServerFirmwareCatalogVendorId) {
		return nil, false
	}
	return o.ServerFirmwareCatalogVendorId, true
}

// HasServerFirmwareCatalogVendorId returns a boolean if a field has been set.
func (o *FirmwareCatalog) HasServerFirmwareCatalogVendorId() bool {
	if o != nil && !IsNil(o.ServerFirmwareCatalogVendorId) {
		return true
	}

	return false
}

// SetServerFirmwareCatalogVendorId gets a reference to the given string and assigns it to the ServerFirmwareCatalogVendorId field.
func (o *FirmwareCatalog) SetServerFirmwareCatalogVendorId(v string) {
	o.ServerFirmwareCatalogVendorId = &v
}

// GetServerFirmwareCatalogVendorUrl returns the ServerFirmwareCatalogVendorUrl field value if set, zero value otherwise.
func (o *FirmwareCatalog) GetServerFirmwareCatalogVendorUrl() string {
	if o == nil || IsNil(o.ServerFirmwareCatalogVendorUrl) {
		var ret string
		return ret
	}
	return *o.ServerFirmwareCatalogVendorUrl
}

// GetServerFirmwareCatalogVendorUrlOk returns a tuple with the ServerFirmwareCatalogVendorUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareCatalog) GetServerFirmwareCatalogVendorUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ServerFirmwareCatalogVendorUrl) {
		return nil, false
	}
	return o.ServerFirmwareCatalogVendorUrl, true
}

// HasServerFirmwareCatalogVendorUrl returns a boolean if a field has been set.
func (o *FirmwareCatalog) HasServerFirmwareCatalogVendorUrl() bool {
	if o != nil && !IsNil(o.ServerFirmwareCatalogVendorUrl) {
		return true
	}

	return false
}

// SetServerFirmwareCatalogVendorUrl gets a reference to the given string and assigns it to the ServerFirmwareCatalogVendorUrl field.
func (o *FirmwareCatalog) SetServerFirmwareCatalogVendorUrl(v string) {
	o.ServerFirmwareCatalogVendorUrl = &v
}

// GetServerFirmwareCatalogVendorReleaseTimestamp returns the ServerFirmwareCatalogVendorReleaseTimestamp field value if set, zero value otherwise.
func (o *FirmwareCatalog) GetServerFirmwareCatalogVendorReleaseTimestamp() string {
	if o == nil || IsNil(o.ServerFirmwareCatalogVendorReleaseTimestamp) {
		var ret string
		return ret
	}
	return *o.ServerFirmwareCatalogVendorReleaseTimestamp
}

// GetServerFirmwareCatalogVendorReleaseTimestampOk returns a tuple with the ServerFirmwareCatalogVendorReleaseTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareCatalog) GetServerFirmwareCatalogVendorReleaseTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.ServerFirmwareCatalogVendorReleaseTimestamp) {
		return nil, false
	}
	return o.ServerFirmwareCatalogVendorReleaseTimestamp, true
}

// HasServerFirmwareCatalogVendorReleaseTimestamp returns a boolean if a field has been set.
func (o *FirmwareCatalog) HasServerFirmwareCatalogVendorReleaseTimestamp() bool {
	if o != nil && !IsNil(o.ServerFirmwareCatalogVendorReleaseTimestamp) {
		return true
	}

	return false
}

// SetServerFirmwareCatalogVendorReleaseTimestamp gets a reference to the given string and assigns it to the ServerFirmwareCatalogVendorReleaseTimestamp field.
func (o *FirmwareCatalog) SetServerFirmwareCatalogVendorReleaseTimestamp(v string) {
	o.ServerFirmwareCatalogVendorReleaseTimestamp = &v
}

// GetServerFirmwareCatalogUpdateType returns the ServerFirmwareCatalogUpdateType field value
func (o *FirmwareCatalog) GetServerFirmwareCatalogUpdateType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerFirmwareCatalogUpdateType
}

// GetServerFirmwareCatalogUpdateTypeOk returns a tuple with the ServerFirmwareCatalogUpdateType field value
// and a boolean to check if the value has been set.
func (o *FirmwareCatalog) GetServerFirmwareCatalogUpdateTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerFirmwareCatalogUpdateType, true
}

// SetServerFirmwareCatalogUpdateType sets field value
func (o *FirmwareCatalog) SetServerFirmwareCatalogUpdateType(v string) {
	o.ServerFirmwareCatalogUpdateType = v
}

// GetServerFirmwareCatalogMetalsoftServerTypesSupportedJson returns the ServerFirmwareCatalogMetalsoftServerTypesSupportedJson field value if set, zero value otherwise.
func (o *FirmwareCatalog) GetServerFirmwareCatalogMetalsoftServerTypesSupportedJson() map[string]interface{} {
	if o == nil || IsNil(o.ServerFirmwareCatalogMetalsoftServerTypesSupportedJson) {
		var ret map[string]interface{}
		return ret
	}
	return o.ServerFirmwareCatalogMetalsoftServerTypesSupportedJson
}

// GetServerFirmwareCatalogMetalsoftServerTypesSupportedJsonOk returns a tuple with the ServerFirmwareCatalogMetalsoftServerTypesSupportedJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareCatalog) GetServerFirmwareCatalogMetalsoftServerTypesSupportedJsonOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ServerFirmwareCatalogMetalsoftServerTypesSupportedJson) {
		return map[string]interface{}{}, false
	}
	return o.ServerFirmwareCatalogMetalsoftServerTypesSupportedJson, true
}

// HasServerFirmwareCatalogMetalsoftServerTypesSupportedJson returns a boolean if a field has been set.
func (o *FirmwareCatalog) HasServerFirmwareCatalogMetalsoftServerTypesSupportedJson() bool {
	if o != nil && !IsNil(o.ServerFirmwareCatalogMetalsoftServerTypesSupportedJson) {
		return true
	}

	return false
}

// SetServerFirmwareCatalogMetalsoftServerTypesSupportedJson gets a reference to the given map[string]interface{} and assigns it to the ServerFirmwareCatalogMetalsoftServerTypesSupportedJson field.
func (o *FirmwareCatalog) SetServerFirmwareCatalogMetalsoftServerTypesSupportedJson(v map[string]interface{}) {
	o.ServerFirmwareCatalogMetalsoftServerTypesSupportedJson = v
}

// GetServerFirmwareCatalogVendorServerTypesSupportedJson returns the ServerFirmwareCatalogVendorServerTypesSupportedJson field value if set, zero value otherwise.
func (o *FirmwareCatalog) GetServerFirmwareCatalogVendorServerTypesSupportedJson() map[string]interface{} {
	if o == nil || IsNil(o.ServerFirmwareCatalogVendorServerTypesSupportedJson) {
		var ret map[string]interface{}
		return ret
	}
	return o.ServerFirmwareCatalogVendorServerTypesSupportedJson
}

// GetServerFirmwareCatalogVendorServerTypesSupportedJsonOk returns a tuple with the ServerFirmwareCatalogVendorServerTypesSupportedJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareCatalog) GetServerFirmwareCatalogVendorServerTypesSupportedJsonOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ServerFirmwareCatalogVendorServerTypesSupportedJson) {
		return map[string]interface{}{}, false
	}
	return o.ServerFirmwareCatalogVendorServerTypesSupportedJson, true
}

// HasServerFirmwareCatalogVendorServerTypesSupportedJson returns a boolean if a field has been set.
func (o *FirmwareCatalog) HasServerFirmwareCatalogVendorServerTypesSupportedJson() bool {
	if o != nil && !IsNil(o.ServerFirmwareCatalogVendorServerTypesSupportedJson) {
		return true
	}

	return false
}

// SetServerFirmwareCatalogVendorServerTypesSupportedJson gets a reference to the given map[string]interface{} and assigns it to the ServerFirmwareCatalogVendorServerTypesSupportedJson field.
func (o *FirmwareCatalog) SetServerFirmwareCatalogVendorServerTypesSupportedJson(v map[string]interface{}) {
	o.ServerFirmwareCatalogVendorServerTypesSupportedJson = v
}

// GetServerFirmwareCatalogVendorConfigurationJson returns the ServerFirmwareCatalogVendorConfigurationJson field value if set, zero value otherwise.
func (o *FirmwareCatalog) GetServerFirmwareCatalogVendorConfigurationJson() map[string]interface{} {
	if o == nil || IsNil(o.ServerFirmwareCatalogVendorConfigurationJson) {
		var ret map[string]interface{}
		return ret
	}
	return o.ServerFirmwareCatalogVendorConfigurationJson
}

// GetServerFirmwareCatalogVendorConfigurationJsonOk returns a tuple with the ServerFirmwareCatalogVendorConfigurationJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareCatalog) GetServerFirmwareCatalogVendorConfigurationJsonOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ServerFirmwareCatalogVendorConfigurationJson) {
		return map[string]interface{}{}, false
	}
	return o.ServerFirmwareCatalogVendorConfigurationJson, true
}

// HasServerFirmwareCatalogVendorConfigurationJson returns a boolean if a field has been set.
func (o *FirmwareCatalog) HasServerFirmwareCatalogVendorConfigurationJson() bool {
	if o != nil && !IsNil(o.ServerFirmwareCatalogVendorConfigurationJson) {
		return true
	}

	return false
}

// SetServerFirmwareCatalogVendorConfigurationJson gets a reference to the given map[string]interface{} and assigns it to the ServerFirmwareCatalogVendorConfigurationJson field.
func (o *FirmwareCatalog) SetServerFirmwareCatalogVendorConfigurationJson(v map[string]interface{}) {
	o.ServerFirmwareCatalogVendorConfigurationJson = v
}

// GetServerFirmwareCatalogCreatedTimestamp returns the ServerFirmwareCatalogCreatedTimestamp field value
func (o *FirmwareCatalog) GetServerFirmwareCatalogCreatedTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerFirmwareCatalogCreatedTimestamp
}

// GetServerFirmwareCatalogCreatedTimestampOk returns a tuple with the ServerFirmwareCatalogCreatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *FirmwareCatalog) GetServerFirmwareCatalogCreatedTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerFirmwareCatalogCreatedTimestamp, true
}

// SetServerFirmwareCatalogCreatedTimestamp sets field value
func (o *FirmwareCatalog) SetServerFirmwareCatalogCreatedTimestamp(v string) {
	o.ServerFirmwareCatalogCreatedTimestamp = v
}

// GetLinks returns the Links field value
func (o *FirmwareCatalog) GetLinks() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *FirmwareCatalog) GetLinksOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *FirmwareCatalog) SetLinks(v map[string]interface{}) {
	o.Links = v
}

func (o FirmwareCatalog) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirmwareCatalog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serverFirmwareCatalogId"] = o.ServerFirmwareCatalogId
	toSerialize["serverFirmwareCatalogName"] = o.ServerFirmwareCatalogName
	if !IsNil(o.ServerFirmwareCatalogDescription) {
		toSerialize["serverFirmwareCatalogDescription"] = o.ServerFirmwareCatalogDescription
	}
	toSerialize["serverFirmwareCatalogVendor"] = o.ServerFirmwareCatalogVendor
	if !IsNil(o.ServerFirmwareCatalogVendorId) {
		toSerialize["serverFirmwareCatalogVendorId"] = o.ServerFirmwareCatalogVendorId
	}
	if !IsNil(o.ServerFirmwareCatalogVendorUrl) {
		toSerialize["serverFirmwareCatalogVendorUrl"] = o.ServerFirmwareCatalogVendorUrl
	}
	if !IsNil(o.ServerFirmwareCatalogVendorReleaseTimestamp) {
		toSerialize["serverFirmwareCatalogVendorReleaseTimestamp"] = o.ServerFirmwareCatalogVendorReleaseTimestamp
	}
	toSerialize["serverFirmwareCatalogUpdateType"] = o.ServerFirmwareCatalogUpdateType
	if !IsNil(o.ServerFirmwareCatalogMetalsoftServerTypesSupportedJson) {
		toSerialize["serverFirmwareCatalogMetalsoftServerTypesSupportedJson"] = o.ServerFirmwareCatalogMetalsoftServerTypesSupportedJson
	}
	if !IsNil(o.ServerFirmwareCatalogVendorServerTypesSupportedJson) {
		toSerialize["serverFirmwareCatalogVendorServerTypesSupportedJson"] = o.ServerFirmwareCatalogVendorServerTypesSupportedJson
	}
	if !IsNil(o.ServerFirmwareCatalogVendorConfigurationJson) {
		toSerialize["serverFirmwareCatalogVendorConfigurationJson"] = o.ServerFirmwareCatalogVendorConfigurationJson
	}
	toSerialize["serverFirmwareCatalogCreatedTimestamp"] = o.ServerFirmwareCatalogCreatedTimestamp
	toSerialize["links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FirmwareCatalog) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"serverFirmwareCatalogId",
		"serverFirmwareCatalogName",
		"serverFirmwareCatalogVendor",
		"serverFirmwareCatalogUpdateType",
		"serverFirmwareCatalogCreatedTimestamp",
		"links",
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

	varFirmwareCatalog := _FirmwareCatalog{}

	err = json.Unmarshal(data, &varFirmwareCatalog)

	if err != nil {
		return err
	}

	*o = FirmwareCatalog(varFirmwareCatalog)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "serverFirmwareCatalogId")
		delete(additionalProperties, "serverFirmwareCatalogName")
		delete(additionalProperties, "serverFirmwareCatalogDescription")
		delete(additionalProperties, "serverFirmwareCatalogVendor")
		delete(additionalProperties, "serverFirmwareCatalogVendorId")
		delete(additionalProperties, "serverFirmwareCatalogVendorUrl")
		delete(additionalProperties, "serverFirmwareCatalogVendorReleaseTimestamp")
		delete(additionalProperties, "serverFirmwareCatalogUpdateType")
		delete(additionalProperties, "serverFirmwareCatalogMetalsoftServerTypesSupportedJson")
		delete(additionalProperties, "serverFirmwareCatalogVendorServerTypesSupportedJson")
		delete(additionalProperties, "serverFirmwareCatalogVendorConfigurationJson")
		delete(additionalProperties, "serverFirmwareCatalogCreatedTimestamp")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFirmwareCatalog struct {
	value *FirmwareCatalog
	isSet bool
}

func (v NullableFirmwareCatalog) Get() *FirmwareCatalog {
	return v.value
}

func (v *NullableFirmwareCatalog) Set(val *FirmwareCatalog) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareCatalog) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareCatalog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareCatalog(val *FirmwareCatalog) *NullableFirmwareCatalog {
	return &NullableFirmwareCatalog{value: val, isSet: true}
}

func (v NullableFirmwareCatalog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareCatalog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


