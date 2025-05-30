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

// checks if the FirmwareBinaryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirmwareBinaryResponse{}

// FirmwareBinaryResponse struct for FirmwareBinaryResponse
type FirmwareBinaryResponse struct {
	// Unique identifier for the firmware binary
	Id float32 `json:"id"`
	// Unique identifier for the firmware catalog
	CatalogId float32 `json:"catalogId"`
	// External/vendor identifier for the firmware binary
	ExternalId *string `json:"externalId,omitempty"`
	// Vendor URL with information about the firmware binary
	VendorInfoUrl *string `json:"vendorInfoUrl,omitempty"`
	// Vendor download URL for the firmware binary
	VendorDownloadUrl string `json:"vendorDownloadUrl"`
	// Local cache download URL for the firmware binary
	CacheDownloadUrl *string `json:"cacheDownloadUrl,omitempty"`
	// Name of the firmware binary as specified by the vendor
	Name string `json:"name"`
	// Vendor package identifier
	PackageId *string `json:"packageId,omitempty"`
	// Vendor package version
	PackageVersion string `json:"packageVersion"`
	// Indicates if a reboot is required for this binary
	RebootRequired bool `json:"rebootRequired"`
	// Severity of the firmware update
	UpdateSeverity string `json:"updateSeverity"`
	// Vendor specific record for supported devices
	VendorSupportedDevices map[string]interface{} `json:"vendorSupportedDevices"`
	// Vendor specific records for supported systems
	VendorSupportedSystems []map[string]interface{} `json:"vendorSupportedSystems"`
	// Vendor specific record containing all other firmware binary information
	Vendor map[string]interface{} `json:"vendor"`
	// Vendor release timestamp for the firmware binary
	VendorReleaseTimestamp *string `json:"vendorReleaseTimestamp,omitempty"`
	// Timestamp when the firmware binary was created
	CreatedTimestamp string `json:"createdTimestamp"`
	// Associated firmware catalog
	Catalog *FirmwareCatalog `json:"catalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareBinaryResponse FirmwareBinaryResponse

// NewFirmwareBinaryResponse instantiates a new FirmwareBinaryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareBinaryResponse(id float32, catalogId float32, vendorDownloadUrl string, name string, packageVersion string, rebootRequired bool, updateSeverity string, vendorSupportedDevices map[string]interface{}, vendorSupportedSystems []map[string]interface{}, vendor map[string]interface{}, createdTimestamp string) *FirmwareBinaryResponse {
	this := FirmwareBinaryResponse{}
	this.Id = id
	this.CatalogId = catalogId
	this.VendorDownloadUrl = vendorDownloadUrl
	this.Name = name
	this.PackageVersion = packageVersion
	this.RebootRequired = rebootRequired
	this.UpdateSeverity = updateSeverity
	this.VendorSupportedDevices = vendorSupportedDevices
	this.VendorSupportedSystems = vendorSupportedSystems
	this.Vendor = vendor
	this.CreatedTimestamp = createdTimestamp
	return &this
}

// NewFirmwareBinaryResponseWithDefaults instantiates a new FirmwareBinaryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareBinaryResponseWithDefaults() *FirmwareBinaryResponse {
	this := FirmwareBinaryResponse{}
	return &this
}

// GetId returns the Id field value
func (o *FirmwareBinaryResponse) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FirmwareBinaryResponse) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FirmwareBinaryResponse) SetId(v float32) {
	o.Id = v
}

// GetCatalogId returns the CatalogId field value
func (o *FirmwareBinaryResponse) GetCatalogId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CatalogId
}

// GetCatalogIdOk returns a tuple with the CatalogId field value
// and a boolean to check if the value has been set.
func (o *FirmwareBinaryResponse) GetCatalogIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CatalogId, true
}

// SetCatalogId sets field value
func (o *FirmwareBinaryResponse) SetCatalogId(v float32) {
	o.CatalogId = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *FirmwareBinaryResponse) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareBinaryResponse) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *FirmwareBinaryResponse) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *FirmwareBinaryResponse) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetVendorInfoUrl returns the VendorInfoUrl field value if set, zero value otherwise.
func (o *FirmwareBinaryResponse) GetVendorInfoUrl() string {
	if o == nil || IsNil(o.VendorInfoUrl) {
		var ret string
		return ret
	}
	return *o.VendorInfoUrl
}

// GetVendorInfoUrlOk returns a tuple with the VendorInfoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareBinaryResponse) GetVendorInfoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.VendorInfoUrl) {
		return nil, false
	}
	return o.VendorInfoUrl, true
}

// HasVendorInfoUrl returns a boolean if a field has been set.
func (o *FirmwareBinaryResponse) HasVendorInfoUrl() bool {
	if o != nil && !IsNil(o.VendorInfoUrl) {
		return true
	}

	return false
}

// SetVendorInfoUrl gets a reference to the given string and assigns it to the VendorInfoUrl field.
func (o *FirmwareBinaryResponse) SetVendorInfoUrl(v string) {
	o.VendorInfoUrl = &v
}

// GetVendorDownloadUrl returns the VendorDownloadUrl field value
func (o *FirmwareBinaryResponse) GetVendorDownloadUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VendorDownloadUrl
}

// GetVendorDownloadUrlOk returns a tuple with the VendorDownloadUrl field value
// and a boolean to check if the value has been set.
func (o *FirmwareBinaryResponse) GetVendorDownloadUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VendorDownloadUrl, true
}

// SetVendorDownloadUrl sets field value
func (o *FirmwareBinaryResponse) SetVendorDownloadUrl(v string) {
	o.VendorDownloadUrl = v
}

// GetCacheDownloadUrl returns the CacheDownloadUrl field value if set, zero value otherwise.
func (o *FirmwareBinaryResponse) GetCacheDownloadUrl() string {
	if o == nil || IsNil(o.CacheDownloadUrl) {
		var ret string
		return ret
	}
	return *o.CacheDownloadUrl
}

// GetCacheDownloadUrlOk returns a tuple with the CacheDownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareBinaryResponse) GetCacheDownloadUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CacheDownloadUrl) {
		return nil, false
	}
	return o.CacheDownloadUrl, true
}

// HasCacheDownloadUrl returns a boolean if a field has been set.
func (o *FirmwareBinaryResponse) HasCacheDownloadUrl() bool {
	if o != nil && !IsNil(o.CacheDownloadUrl) {
		return true
	}

	return false
}

// SetCacheDownloadUrl gets a reference to the given string and assigns it to the CacheDownloadUrl field.
func (o *FirmwareBinaryResponse) SetCacheDownloadUrl(v string) {
	o.CacheDownloadUrl = &v
}

// GetName returns the Name field value
func (o *FirmwareBinaryResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FirmwareBinaryResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FirmwareBinaryResponse) SetName(v string) {
	o.Name = v
}

// GetPackageId returns the PackageId field value if set, zero value otherwise.
func (o *FirmwareBinaryResponse) GetPackageId() string {
	if o == nil || IsNil(o.PackageId) {
		var ret string
		return ret
	}
	return *o.PackageId
}

// GetPackageIdOk returns a tuple with the PackageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareBinaryResponse) GetPackageIdOk() (*string, bool) {
	if o == nil || IsNil(o.PackageId) {
		return nil, false
	}
	return o.PackageId, true
}

// HasPackageId returns a boolean if a field has been set.
func (o *FirmwareBinaryResponse) HasPackageId() bool {
	if o != nil && !IsNil(o.PackageId) {
		return true
	}

	return false
}

// SetPackageId gets a reference to the given string and assigns it to the PackageId field.
func (o *FirmwareBinaryResponse) SetPackageId(v string) {
	o.PackageId = &v
}

// GetPackageVersion returns the PackageVersion field value
func (o *FirmwareBinaryResponse) GetPackageVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageVersion
}

// GetPackageVersionOk returns a tuple with the PackageVersion field value
// and a boolean to check if the value has been set.
func (o *FirmwareBinaryResponse) GetPackageVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageVersion, true
}

// SetPackageVersion sets field value
func (o *FirmwareBinaryResponse) SetPackageVersion(v string) {
	o.PackageVersion = v
}

// GetRebootRequired returns the RebootRequired field value
func (o *FirmwareBinaryResponse) GetRebootRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RebootRequired
}

// GetRebootRequiredOk returns a tuple with the RebootRequired field value
// and a boolean to check if the value has been set.
func (o *FirmwareBinaryResponse) GetRebootRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RebootRequired, true
}

// SetRebootRequired sets field value
func (o *FirmwareBinaryResponse) SetRebootRequired(v bool) {
	o.RebootRequired = v
}

// GetUpdateSeverity returns the UpdateSeverity field value
func (o *FirmwareBinaryResponse) GetUpdateSeverity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdateSeverity
}

// GetUpdateSeverityOk returns a tuple with the UpdateSeverity field value
// and a boolean to check if the value has been set.
func (o *FirmwareBinaryResponse) GetUpdateSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdateSeverity, true
}

// SetUpdateSeverity sets field value
func (o *FirmwareBinaryResponse) SetUpdateSeverity(v string) {
	o.UpdateSeverity = v
}

// GetVendorSupportedDevices returns the VendorSupportedDevices field value
func (o *FirmwareBinaryResponse) GetVendorSupportedDevices() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.VendorSupportedDevices
}

// GetVendorSupportedDevicesOk returns a tuple with the VendorSupportedDevices field value
// and a boolean to check if the value has been set.
func (o *FirmwareBinaryResponse) GetVendorSupportedDevicesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.VendorSupportedDevices, true
}

// SetVendorSupportedDevices sets field value
func (o *FirmwareBinaryResponse) SetVendorSupportedDevices(v map[string]interface{}) {
	o.VendorSupportedDevices = v
}

// GetVendorSupportedSystems returns the VendorSupportedSystems field value
func (o *FirmwareBinaryResponse) GetVendorSupportedSystems() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.VendorSupportedSystems
}

// GetVendorSupportedSystemsOk returns a tuple with the VendorSupportedSystems field value
// and a boolean to check if the value has been set.
func (o *FirmwareBinaryResponse) GetVendorSupportedSystemsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.VendorSupportedSystems, true
}

// SetVendorSupportedSystems sets field value
func (o *FirmwareBinaryResponse) SetVendorSupportedSystems(v []map[string]interface{}) {
	o.VendorSupportedSystems = v
}

// GetVendor returns the Vendor field value
func (o *FirmwareBinaryResponse) GetVendor() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value
// and a boolean to check if the value has been set.
func (o *FirmwareBinaryResponse) GetVendorOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Vendor, true
}

// SetVendor sets field value
func (o *FirmwareBinaryResponse) SetVendor(v map[string]interface{}) {
	o.Vendor = v
}

// GetVendorReleaseTimestamp returns the VendorReleaseTimestamp field value if set, zero value otherwise.
func (o *FirmwareBinaryResponse) GetVendorReleaseTimestamp() string {
	if o == nil || IsNil(o.VendorReleaseTimestamp) {
		var ret string
		return ret
	}
	return *o.VendorReleaseTimestamp
}

// GetVendorReleaseTimestampOk returns a tuple with the VendorReleaseTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareBinaryResponse) GetVendorReleaseTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.VendorReleaseTimestamp) {
		return nil, false
	}
	return o.VendorReleaseTimestamp, true
}

// HasVendorReleaseTimestamp returns a boolean if a field has been set.
func (o *FirmwareBinaryResponse) HasVendorReleaseTimestamp() bool {
	if o != nil && !IsNil(o.VendorReleaseTimestamp) {
		return true
	}

	return false
}

// SetVendorReleaseTimestamp gets a reference to the given string and assigns it to the VendorReleaseTimestamp field.
func (o *FirmwareBinaryResponse) SetVendorReleaseTimestamp(v string) {
	o.VendorReleaseTimestamp = &v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value
func (o *FirmwareBinaryResponse) GetCreatedTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *FirmwareBinaryResponse) GetCreatedTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTimestamp, true
}

// SetCreatedTimestamp sets field value
func (o *FirmwareBinaryResponse) SetCreatedTimestamp(v string) {
	o.CreatedTimestamp = v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *FirmwareBinaryResponse) GetCatalog() FirmwareCatalog {
	if o == nil || IsNil(o.Catalog) {
		var ret FirmwareCatalog
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareBinaryResponse) GetCatalogOk() (*FirmwareCatalog, bool) {
	if o == nil || IsNil(o.Catalog) {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *FirmwareBinaryResponse) HasCatalog() bool {
	if o != nil && !IsNil(o.Catalog) {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given FirmwareCatalog and assigns it to the Catalog field.
func (o *FirmwareBinaryResponse) SetCatalog(v FirmwareCatalog) {
	o.Catalog = &v
}

func (o FirmwareBinaryResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirmwareBinaryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["catalogId"] = o.CatalogId
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.VendorInfoUrl) {
		toSerialize["vendorInfoUrl"] = o.VendorInfoUrl
	}
	toSerialize["vendorDownloadUrl"] = o.VendorDownloadUrl
	if !IsNil(o.CacheDownloadUrl) {
		toSerialize["cacheDownloadUrl"] = o.CacheDownloadUrl
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.PackageId) {
		toSerialize["packageId"] = o.PackageId
	}
	toSerialize["packageVersion"] = o.PackageVersion
	toSerialize["rebootRequired"] = o.RebootRequired
	toSerialize["updateSeverity"] = o.UpdateSeverity
	toSerialize["vendorSupportedDevices"] = o.VendorSupportedDevices
	toSerialize["vendorSupportedSystems"] = o.VendorSupportedSystems
	toSerialize["vendor"] = o.Vendor
	if !IsNil(o.VendorReleaseTimestamp) {
		toSerialize["vendorReleaseTimestamp"] = o.VendorReleaseTimestamp
	}
	toSerialize["createdTimestamp"] = o.CreatedTimestamp
	if !IsNil(o.Catalog) {
		toSerialize["catalog"] = o.Catalog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FirmwareBinaryResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"catalogId",
		"vendorDownloadUrl",
		"name",
		"packageVersion",
		"rebootRequired",
		"updateSeverity",
		"vendorSupportedDevices",
		"vendorSupportedSystems",
		"vendor",
		"createdTimestamp",
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

	varFirmwareBinaryResponse := _FirmwareBinaryResponse{}

	err = json.Unmarshal(data, &varFirmwareBinaryResponse)

	if err != nil {
		return err
	}

	*o = FirmwareBinaryResponse(varFirmwareBinaryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "catalogId")
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "vendorInfoUrl")
		delete(additionalProperties, "vendorDownloadUrl")
		delete(additionalProperties, "cacheDownloadUrl")
		delete(additionalProperties, "name")
		delete(additionalProperties, "packageId")
		delete(additionalProperties, "packageVersion")
		delete(additionalProperties, "rebootRequired")
		delete(additionalProperties, "updateSeverity")
		delete(additionalProperties, "vendorSupportedDevices")
		delete(additionalProperties, "vendorSupportedSystems")
		delete(additionalProperties, "vendor")
		delete(additionalProperties, "vendorReleaseTimestamp")
		delete(additionalProperties, "createdTimestamp")
		delete(additionalProperties, "catalog")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFirmwareBinaryResponse struct {
	value *FirmwareBinaryResponse
	isSet bool
}

func (v NullableFirmwareBinaryResponse) Get() *FirmwareBinaryResponse {
	return v.value
}

func (v *NullableFirmwareBinaryResponse) Set(val *FirmwareBinaryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareBinaryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareBinaryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareBinaryResponse(val *FirmwareBinaryResponse) *NullableFirmwareBinaryResponse {
	return &NullableFirmwareBinaryResponse{value: val, isSet: true}
}

func (v NullableFirmwareBinaryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareBinaryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


