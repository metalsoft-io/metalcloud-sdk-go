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

// checks if the CreateFirmwareBinary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFirmwareBinary{}

// CreateFirmwareBinary struct for CreateFirmwareBinary
type CreateFirmwareBinary struct {
	ServerFirmwareBinaryCatalogId float32 `json:"serverFirmwareBinaryCatalogId"`
	ServerFirmwareBinaryExternalId *string `json:"serverFirmwareBinaryExternalId,omitempty"`
	ServerFirmwareBinaryVendorInfoUrl *string `json:"serverFirmwareBinaryVendorInfoUrl,omitempty"`
	ServerFirmwareBinaryVendorDownloadUrl string `json:"serverFirmwareBinaryVendorDownloadUrl"`
	ServerFirmwareBinaryCacheDownloadUrl *string `json:"serverFirmwareBinaryCacheDownloadUrl,omitempty"`
	ServerFirmwareBinaryName string `json:"serverFirmwareBinaryName"`
	ServerFirmwareBinaryPackageId *string `json:"serverFirmwareBinaryPackageId,omitempty"`
	ServerFirmwareBinaryPackageVersion *string `json:"serverFirmwareBinaryPackageVersion,omitempty"`
	ServerFirmwareBinaryRebootRequired bool `json:"serverFirmwareBinaryRebootRequired"`
	ServerFirmwareBinaryUpdateSeverity FirmwareBinaryUpdateSeverity `json:"serverFirmwareBinaryUpdateSeverity"`
	ServerFirmwareBinaryVendorSupportedDevicesJson string `json:"serverFirmwareBinaryVendorSupportedDevicesJson"`
	ServerFirmwareBinaryVendorSupportedSystemsJson string `json:"serverFirmwareBinaryVendorSupportedSystemsJson"`
	ServerFirmwareBinaryVendorReleaseTimestamp *string `json:"serverFirmwareBinaryVendorReleaseTimestamp,omitempty"`
	ServerFirmwareBinaryVendorJson *string `json:"serverFirmwareBinaryVendorJson,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateFirmwareBinary CreateFirmwareBinary

// NewCreateFirmwareBinary instantiates a new CreateFirmwareBinary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFirmwareBinary(serverFirmwareBinaryCatalogId float32, serverFirmwareBinaryVendorDownloadUrl string, serverFirmwareBinaryName string, serverFirmwareBinaryRebootRequired bool, serverFirmwareBinaryUpdateSeverity FirmwareBinaryUpdateSeverity, serverFirmwareBinaryVendorSupportedDevicesJson string, serverFirmwareBinaryVendorSupportedSystemsJson string) *CreateFirmwareBinary {
	this := CreateFirmwareBinary{}
	this.ServerFirmwareBinaryCatalogId = serverFirmwareBinaryCatalogId
	this.ServerFirmwareBinaryVendorDownloadUrl = serverFirmwareBinaryVendorDownloadUrl
	this.ServerFirmwareBinaryName = serverFirmwareBinaryName
	this.ServerFirmwareBinaryRebootRequired = serverFirmwareBinaryRebootRequired
	this.ServerFirmwareBinaryUpdateSeverity = serverFirmwareBinaryUpdateSeverity
	this.ServerFirmwareBinaryVendorSupportedDevicesJson = serverFirmwareBinaryVendorSupportedDevicesJson
	this.ServerFirmwareBinaryVendorSupportedSystemsJson = serverFirmwareBinaryVendorSupportedSystemsJson
	return &this
}

// NewCreateFirmwareBinaryWithDefaults instantiates a new CreateFirmwareBinary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFirmwareBinaryWithDefaults() *CreateFirmwareBinary {
	this := CreateFirmwareBinary{}
	return &this
}

// GetServerFirmwareBinaryCatalogId returns the ServerFirmwareBinaryCatalogId field value
func (o *CreateFirmwareBinary) GetServerFirmwareBinaryCatalogId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ServerFirmwareBinaryCatalogId
}

// GetServerFirmwareBinaryCatalogIdOk returns a tuple with the ServerFirmwareBinaryCatalogId field value
// and a boolean to check if the value has been set.
func (o *CreateFirmwareBinary) GetServerFirmwareBinaryCatalogIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerFirmwareBinaryCatalogId, true
}

// SetServerFirmwareBinaryCatalogId sets field value
func (o *CreateFirmwareBinary) SetServerFirmwareBinaryCatalogId(v float32) {
	o.ServerFirmwareBinaryCatalogId = v
}

// GetServerFirmwareBinaryExternalId returns the ServerFirmwareBinaryExternalId field value if set, zero value otherwise.
func (o *CreateFirmwareBinary) GetServerFirmwareBinaryExternalId() string {
	if o == nil || IsNil(o.ServerFirmwareBinaryExternalId) {
		var ret string
		return ret
	}
	return *o.ServerFirmwareBinaryExternalId
}

// GetServerFirmwareBinaryExternalIdOk returns a tuple with the ServerFirmwareBinaryExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFirmwareBinary) GetServerFirmwareBinaryExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServerFirmwareBinaryExternalId) {
		return nil, false
	}
	return o.ServerFirmwareBinaryExternalId, true
}

// HasServerFirmwareBinaryExternalId returns a boolean if a field has been set.
func (o *CreateFirmwareBinary) HasServerFirmwareBinaryExternalId() bool {
	if o != nil && !IsNil(o.ServerFirmwareBinaryExternalId) {
		return true
	}

	return false
}

// SetServerFirmwareBinaryExternalId gets a reference to the given string and assigns it to the ServerFirmwareBinaryExternalId field.
func (o *CreateFirmwareBinary) SetServerFirmwareBinaryExternalId(v string) {
	o.ServerFirmwareBinaryExternalId = &v
}

// GetServerFirmwareBinaryVendorInfoUrl returns the ServerFirmwareBinaryVendorInfoUrl field value if set, zero value otherwise.
func (o *CreateFirmwareBinary) GetServerFirmwareBinaryVendorInfoUrl() string {
	if o == nil || IsNil(o.ServerFirmwareBinaryVendorInfoUrl) {
		var ret string
		return ret
	}
	return *o.ServerFirmwareBinaryVendorInfoUrl
}

// GetServerFirmwareBinaryVendorInfoUrlOk returns a tuple with the ServerFirmwareBinaryVendorInfoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFirmwareBinary) GetServerFirmwareBinaryVendorInfoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ServerFirmwareBinaryVendorInfoUrl) {
		return nil, false
	}
	return o.ServerFirmwareBinaryVendorInfoUrl, true
}

// HasServerFirmwareBinaryVendorInfoUrl returns a boolean if a field has been set.
func (o *CreateFirmwareBinary) HasServerFirmwareBinaryVendorInfoUrl() bool {
	if o != nil && !IsNil(o.ServerFirmwareBinaryVendorInfoUrl) {
		return true
	}

	return false
}

// SetServerFirmwareBinaryVendorInfoUrl gets a reference to the given string and assigns it to the ServerFirmwareBinaryVendorInfoUrl field.
func (o *CreateFirmwareBinary) SetServerFirmwareBinaryVendorInfoUrl(v string) {
	o.ServerFirmwareBinaryVendorInfoUrl = &v
}

// GetServerFirmwareBinaryVendorDownloadUrl returns the ServerFirmwareBinaryVendorDownloadUrl field value
func (o *CreateFirmwareBinary) GetServerFirmwareBinaryVendorDownloadUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerFirmwareBinaryVendorDownloadUrl
}

// GetServerFirmwareBinaryVendorDownloadUrlOk returns a tuple with the ServerFirmwareBinaryVendorDownloadUrl field value
// and a boolean to check if the value has been set.
func (o *CreateFirmwareBinary) GetServerFirmwareBinaryVendorDownloadUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerFirmwareBinaryVendorDownloadUrl, true
}

// SetServerFirmwareBinaryVendorDownloadUrl sets field value
func (o *CreateFirmwareBinary) SetServerFirmwareBinaryVendorDownloadUrl(v string) {
	o.ServerFirmwareBinaryVendorDownloadUrl = v
}

// GetServerFirmwareBinaryCacheDownloadUrl returns the ServerFirmwareBinaryCacheDownloadUrl field value if set, zero value otherwise.
func (o *CreateFirmwareBinary) GetServerFirmwareBinaryCacheDownloadUrl() string {
	if o == nil || IsNil(o.ServerFirmwareBinaryCacheDownloadUrl) {
		var ret string
		return ret
	}
	return *o.ServerFirmwareBinaryCacheDownloadUrl
}

// GetServerFirmwareBinaryCacheDownloadUrlOk returns a tuple with the ServerFirmwareBinaryCacheDownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFirmwareBinary) GetServerFirmwareBinaryCacheDownloadUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ServerFirmwareBinaryCacheDownloadUrl) {
		return nil, false
	}
	return o.ServerFirmwareBinaryCacheDownloadUrl, true
}

// HasServerFirmwareBinaryCacheDownloadUrl returns a boolean if a field has been set.
func (o *CreateFirmwareBinary) HasServerFirmwareBinaryCacheDownloadUrl() bool {
	if o != nil && !IsNil(o.ServerFirmwareBinaryCacheDownloadUrl) {
		return true
	}

	return false
}

// SetServerFirmwareBinaryCacheDownloadUrl gets a reference to the given string and assigns it to the ServerFirmwareBinaryCacheDownloadUrl field.
func (o *CreateFirmwareBinary) SetServerFirmwareBinaryCacheDownloadUrl(v string) {
	o.ServerFirmwareBinaryCacheDownloadUrl = &v
}

// GetServerFirmwareBinaryName returns the ServerFirmwareBinaryName field value
func (o *CreateFirmwareBinary) GetServerFirmwareBinaryName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerFirmwareBinaryName
}

// GetServerFirmwareBinaryNameOk returns a tuple with the ServerFirmwareBinaryName field value
// and a boolean to check if the value has been set.
func (o *CreateFirmwareBinary) GetServerFirmwareBinaryNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerFirmwareBinaryName, true
}

// SetServerFirmwareBinaryName sets field value
func (o *CreateFirmwareBinary) SetServerFirmwareBinaryName(v string) {
	o.ServerFirmwareBinaryName = v
}

// GetServerFirmwareBinaryPackageId returns the ServerFirmwareBinaryPackageId field value if set, zero value otherwise.
func (o *CreateFirmwareBinary) GetServerFirmwareBinaryPackageId() string {
	if o == nil || IsNil(o.ServerFirmwareBinaryPackageId) {
		var ret string
		return ret
	}
	return *o.ServerFirmwareBinaryPackageId
}

// GetServerFirmwareBinaryPackageIdOk returns a tuple with the ServerFirmwareBinaryPackageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFirmwareBinary) GetServerFirmwareBinaryPackageIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServerFirmwareBinaryPackageId) {
		return nil, false
	}
	return o.ServerFirmwareBinaryPackageId, true
}

// HasServerFirmwareBinaryPackageId returns a boolean if a field has been set.
func (o *CreateFirmwareBinary) HasServerFirmwareBinaryPackageId() bool {
	if o != nil && !IsNil(o.ServerFirmwareBinaryPackageId) {
		return true
	}

	return false
}

// SetServerFirmwareBinaryPackageId gets a reference to the given string and assigns it to the ServerFirmwareBinaryPackageId field.
func (o *CreateFirmwareBinary) SetServerFirmwareBinaryPackageId(v string) {
	o.ServerFirmwareBinaryPackageId = &v
}

// GetServerFirmwareBinaryPackageVersion returns the ServerFirmwareBinaryPackageVersion field value if set, zero value otherwise.
func (o *CreateFirmwareBinary) GetServerFirmwareBinaryPackageVersion() string {
	if o == nil || IsNil(o.ServerFirmwareBinaryPackageVersion) {
		var ret string
		return ret
	}
	return *o.ServerFirmwareBinaryPackageVersion
}

// GetServerFirmwareBinaryPackageVersionOk returns a tuple with the ServerFirmwareBinaryPackageVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFirmwareBinary) GetServerFirmwareBinaryPackageVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ServerFirmwareBinaryPackageVersion) {
		return nil, false
	}
	return o.ServerFirmwareBinaryPackageVersion, true
}

// HasServerFirmwareBinaryPackageVersion returns a boolean if a field has been set.
func (o *CreateFirmwareBinary) HasServerFirmwareBinaryPackageVersion() bool {
	if o != nil && !IsNil(o.ServerFirmwareBinaryPackageVersion) {
		return true
	}

	return false
}

// SetServerFirmwareBinaryPackageVersion gets a reference to the given string and assigns it to the ServerFirmwareBinaryPackageVersion field.
func (o *CreateFirmwareBinary) SetServerFirmwareBinaryPackageVersion(v string) {
	o.ServerFirmwareBinaryPackageVersion = &v
}

// GetServerFirmwareBinaryRebootRequired returns the ServerFirmwareBinaryRebootRequired field value
func (o *CreateFirmwareBinary) GetServerFirmwareBinaryRebootRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ServerFirmwareBinaryRebootRequired
}

// GetServerFirmwareBinaryRebootRequiredOk returns a tuple with the ServerFirmwareBinaryRebootRequired field value
// and a boolean to check if the value has been set.
func (o *CreateFirmwareBinary) GetServerFirmwareBinaryRebootRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerFirmwareBinaryRebootRequired, true
}

// SetServerFirmwareBinaryRebootRequired sets field value
func (o *CreateFirmwareBinary) SetServerFirmwareBinaryRebootRequired(v bool) {
	o.ServerFirmwareBinaryRebootRequired = v
}

// GetServerFirmwareBinaryUpdateSeverity returns the ServerFirmwareBinaryUpdateSeverity field value
func (o *CreateFirmwareBinary) GetServerFirmwareBinaryUpdateSeverity() FirmwareBinaryUpdateSeverity {
	if o == nil {
		var ret FirmwareBinaryUpdateSeverity
		return ret
	}

	return o.ServerFirmwareBinaryUpdateSeverity
}

// GetServerFirmwareBinaryUpdateSeverityOk returns a tuple with the ServerFirmwareBinaryUpdateSeverity field value
// and a boolean to check if the value has been set.
func (o *CreateFirmwareBinary) GetServerFirmwareBinaryUpdateSeverityOk() (*FirmwareBinaryUpdateSeverity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerFirmwareBinaryUpdateSeverity, true
}

// SetServerFirmwareBinaryUpdateSeverity sets field value
func (o *CreateFirmwareBinary) SetServerFirmwareBinaryUpdateSeverity(v FirmwareBinaryUpdateSeverity) {
	o.ServerFirmwareBinaryUpdateSeverity = v
}

// GetServerFirmwareBinaryVendorSupportedDevicesJson returns the ServerFirmwareBinaryVendorSupportedDevicesJson field value
func (o *CreateFirmwareBinary) GetServerFirmwareBinaryVendorSupportedDevicesJson() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerFirmwareBinaryVendorSupportedDevicesJson
}

// GetServerFirmwareBinaryVendorSupportedDevicesJsonOk returns a tuple with the ServerFirmwareBinaryVendorSupportedDevicesJson field value
// and a boolean to check if the value has been set.
func (o *CreateFirmwareBinary) GetServerFirmwareBinaryVendorSupportedDevicesJsonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerFirmwareBinaryVendorSupportedDevicesJson, true
}

// SetServerFirmwareBinaryVendorSupportedDevicesJson sets field value
func (o *CreateFirmwareBinary) SetServerFirmwareBinaryVendorSupportedDevicesJson(v string) {
	o.ServerFirmwareBinaryVendorSupportedDevicesJson = v
}

// GetServerFirmwareBinaryVendorSupportedSystemsJson returns the ServerFirmwareBinaryVendorSupportedSystemsJson field value
func (o *CreateFirmwareBinary) GetServerFirmwareBinaryVendorSupportedSystemsJson() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerFirmwareBinaryVendorSupportedSystemsJson
}

// GetServerFirmwareBinaryVendorSupportedSystemsJsonOk returns a tuple with the ServerFirmwareBinaryVendorSupportedSystemsJson field value
// and a boolean to check if the value has been set.
func (o *CreateFirmwareBinary) GetServerFirmwareBinaryVendorSupportedSystemsJsonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerFirmwareBinaryVendorSupportedSystemsJson, true
}

// SetServerFirmwareBinaryVendorSupportedSystemsJson sets field value
func (o *CreateFirmwareBinary) SetServerFirmwareBinaryVendorSupportedSystemsJson(v string) {
	o.ServerFirmwareBinaryVendorSupportedSystemsJson = v
}

// GetServerFirmwareBinaryVendorReleaseTimestamp returns the ServerFirmwareBinaryVendorReleaseTimestamp field value if set, zero value otherwise.
func (o *CreateFirmwareBinary) GetServerFirmwareBinaryVendorReleaseTimestamp() string {
	if o == nil || IsNil(o.ServerFirmwareBinaryVendorReleaseTimestamp) {
		var ret string
		return ret
	}
	return *o.ServerFirmwareBinaryVendorReleaseTimestamp
}

// GetServerFirmwareBinaryVendorReleaseTimestampOk returns a tuple with the ServerFirmwareBinaryVendorReleaseTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFirmwareBinary) GetServerFirmwareBinaryVendorReleaseTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.ServerFirmwareBinaryVendorReleaseTimestamp) {
		return nil, false
	}
	return o.ServerFirmwareBinaryVendorReleaseTimestamp, true
}

// HasServerFirmwareBinaryVendorReleaseTimestamp returns a boolean if a field has been set.
func (o *CreateFirmwareBinary) HasServerFirmwareBinaryVendorReleaseTimestamp() bool {
	if o != nil && !IsNil(o.ServerFirmwareBinaryVendorReleaseTimestamp) {
		return true
	}

	return false
}

// SetServerFirmwareBinaryVendorReleaseTimestamp gets a reference to the given string and assigns it to the ServerFirmwareBinaryVendorReleaseTimestamp field.
func (o *CreateFirmwareBinary) SetServerFirmwareBinaryVendorReleaseTimestamp(v string) {
	o.ServerFirmwareBinaryVendorReleaseTimestamp = &v
}

// GetServerFirmwareBinaryVendorJson returns the ServerFirmwareBinaryVendorJson field value if set, zero value otherwise.
func (o *CreateFirmwareBinary) GetServerFirmwareBinaryVendorJson() string {
	if o == nil || IsNil(o.ServerFirmwareBinaryVendorJson) {
		var ret string
		return ret
	}
	return *o.ServerFirmwareBinaryVendorJson
}

// GetServerFirmwareBinaryVendorJsonOk returns a tuple with the ServerFirmwareBinaryVendorJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFirmwareBinary) GetServerFirmwareBinaryVendorJsonOk() (*string, bool) {
	if o == nil || IsNil(o.ServerFirmwareBinaryVendorJson) {
		return nil, false
	}
	return o.ServerFirmwareBinaryVendorJson, true
}

// HasServerFirmwareBinaryVendorJson returns a boolean if a field has been set.
func (o *CreateFirmwareBinary) HasServerFirmwareBinaryVendorJson() bool {
	if o != nil && !IsNil(o.ServerFirmwareBinaryVendorJson) {
		return true
	}

	return false
}

// SetServerFirmwareBinaryVendorJson gets a reference to the given string and assigns it to the ServerFirmwareBinaryVendorJson field.
func (o *CreateFirmwareBinary) SetServerFirmwareBinaryVendorJson(v string) {
	o.ServerFirmwareBinaryVendorJson = &v
}

func (o CreateFirmwareBinary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFirmwareBinary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serverFirmwareBinaryCatalogId"] = o.ServerFirmwareBinaryCatalogId
	if !IsNil(o.ServerFirmwareBinaryExternalId) {
		toSerialize["serverFirmwareBinaryExternalId"] = o.ServerFirmwareBinaryExternalId
	}
	if !IsNil(o.ServerFirmwareBinaryVendorInfoUrl) {
		toSerialize["serverFirmwareBinaryVendorInfoUrl"] = o.ServerFirmwareBinaryVendorInfoUrl
	}
	toSerialize["serverFirmwareBinaryVendorDownloadUrl"] = o.ServerFirmwareBinaryVendorDownloadUrl
	if !IsNil(o.ServerFirmwareBinaryCacheDownloadUrl) {
		toSerialize["serverFirmwareBinaryCacheDownloadUrl"] = o.ServerFirmwareBinaryCacheDownloadUrl
	}
	toSerialize["serverFirmwareBinaryName"] = o.ServerFirmwareBinaryName
	if !IsNil(o.ServerFirmwareBinaryPackageId) {
		toSerialize["serverFirmwareBinaryPackageId"] = o.ServerFirmwareBinaryPackageId
	}
	if !IsNil(o.ServerFirmwareBinaryPackageVersion) {
		toSerialize["serverFirmwareBinaryPackageVersion"] = o.ServerFirmwareBinaryPackageVersion
	}
	toSerialize["serverFirmwareBinaryRebootRequired"] = o.ServerFirmwareBinaryRebootRequired
	toSerialize["serverFirmwareBinaryUpdateSeverity"] = o.ServerFirmwareBinaryUpdateSeverity
	toSerialize["serverFirmwareBinaryVendorSupportedDevicesJson"] = o.ServerFirmwareBinaryVendorSupportedDevicesJson
	toSerialize["serverFirmwareBinaryVendorSupportedSystemsJson"] = o.ServerFirmwareBinaryVendorSupportedSystemsJson
	if !IsNil(o.ServerFirmwareBinaryVendorReleaseTimestamp) {
		toSerialize["serverFirmwareBinaryVendorReleaseTimestamp"] = o.ServerFirmwareBinaryVendorReleaseTimestamp
	}
	if !IsNil(o.ServerFirmwareBinaryVendorJson) {
		toSerialize["serverFirmwareBinaryVendorJson"] = o.ServerFirmwareBinaryVendorJson
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateFirmwareBinary) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"serverFirmwareBinaryCatalogId",
		"serverFirmwareBinaryVendorDownloadUrl",
		"serverFirmwareBinaryName",
		"serverFirmwareBinaryRebootRequired",
		"serverFirmwareBinaryUpdateSeverity",
		"serverFirmwareBinaryVendorSupportedDevicesJson",
		"serverFirmwareBinaryVendorSupportedSystemsJson",
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

	varCreateFirmwareBinary := _CreateFirmwareBinary{}

	err = json.Unmarshal(data, &varCreateFirmwareBinary)

	if err != nil {
		return err
	}

	*o = CreateFirmwareBinary(varCreateFirmwareBinary)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "serverFirmwareBinaryCatalogId")
		delete(additionalProperties, "serverFirmwareBinaryExternalId")
		delete(additionalProperties, "serverFirmwareBinaryVendorInfoUrl")
		delete(additionalProperties, "serverFirmwareBinaryVendorDownloadUrl")
		delete(additionalProperties, "serverFirmwareBinaryCacheDownloadUrl")
		delete(additionalProperties, "serverFirmwareBinaryName")
		delete(additionalProperties, "serverFirmwareBinaryPackageId")
		delete(additionalProperties, "serverFirmwareBinaryPackageVersion")
		delete(additionalProperties, "serverFirmwareBinaryRebootRequired")
		delete(additionalProperties, "serverFirmwareBinaryUpdateSeverity")
		delete(additionalProperties, "serverFirmwareBinaryVendorSupportedDevicesJson")
		delete(additionalProperties, "serverFirmwareBinaryVendorSupportedSystemsJson")
		delete(additionalProperties, "serverFirmwareBinaryVendorReleaseTimestamp")
		delete(additionalProperties, "serverFirmwareBinaryVendorJson")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateFirmwareBinary struct {
	value *CreateFirmwareBinary
	isSet bool
}

func (v NullableCreateFirmwareBinary) Get() *CreateFirmwareBinary {
	return v.value
}

func (v *NullableCreateFirmwareBinary) Set(val *CreateFirmwareBinary) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFirmwareBinary) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFirmwareBinary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFirmwareBinary(val *CreateFirmwareBinary) *NullableCreateFirmwareBinary {
	return &NullableCreateFirmwareBinary{value: val, isSet: true}
}

func (v NullableCreateFirmwareBinary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFirmwareBinary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


