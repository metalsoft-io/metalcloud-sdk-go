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
	Id float32 `json:"id"`
	// Name of the catalog, must be unique
	Name string `json:"name"`
	// User description of the catalog
	Description *string `json:"description,omitempty"`
	// Firmware catalog vendor: dell, lenovo, hp
	Vendor string `json:"vendor"`
	// Vendor identifier for the catalog
	VendorId *string `json:"vendorId,omitempty"`
	// Vendor URL for the firmware catalog
	VendorUrl *string `json:"vendorUrl,omitempty"`
	// Vendor release timestamp for the catalog
	VendorReleaseTimestamp *string `json:"vendorReleaseTimestamp,omitempty"`
	// Catalog update type: online or offline
	UpdateType string `json:"updateType"`
	// List of supported Metalsoft server types for which firmware binaries are available
	MetalsoftServerTypesSupported []string `json:"metalsoftServerTypesSupported,omitempty"`
	// List of supported server types for which firmware binaries are available
	VendorServerTypesSupported []string `json:"vendorServerTypesSupported,omitempty"`
	// Vendor configuration
	VendorConfiguration map[string]interface{} `json:"vendorConfiguration,omitempty"`
	// Timestamp when the catalog was created
	CreatedTimestamp string `json:"createdTimestamp"`
	// Links to other resources
	Links []Link `json:"links"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareCatalog FirmwareCatalog

// NewFirmwareCatalog instantiates a new FirmwareCatalog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareCatalog(id float32, name string, vendor string, updateType string, createdTimestamp string, links []Link) *FirmwareCatalog {
	this := FirmwareCatalog{}
	this.Id = id
	this.Name = name
	this.Vendor = vendor
	this.UpdateType = updateType
	this.CreatedTimestamp = createdTimestamp
	this.Links = links
	return &this
}

// NewFirmwareCatalogWithDefaults instantiates a new FirmwareCatalog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareCatalogWithDefaults() *FirmwareCatalog {
	this := FirmwareCatalog{}
	var updateType string = "online"
	this.UpdateType = updateType
	return &this
}

// GetId returns the Id field value
func (o *FirmwareCatalog) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FirmwareCatalog) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FirmwareCatalog) SetId(v float32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *FirmwareCatalog) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FirmwareCatalog) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FirmwareCatalog) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FirmwareCatalog) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareCatalog) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FirmwareCatalog) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FirmwareCatalog) SetDescription(v string) {
	o.Description = &v
}

// GetVendor returns the Vendor field value
func (o *FirmwareCatalog) GetVendor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value
// and a boolean to check if the value has been set.
func (o *FirmwareCatalog) GetVendorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vendor, true
}

// SetVendor sets field value
func (o *FirmwareCatalog) SetVendor(v string) {
	o.Vendor = v
}

// GetVendorId returns the VendorId field value if set, zero value otherwise.
func (o *FirmwareCatalog) GetVendorId() string {
	if o == nil || IsNil(o.VendorId) {
		var ret string
		return ret
	}
	return *o.VendorId
}

// GetVendorIdOk returns a tuple with the VendorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareCatalog) GetVendorIdOk() (*string, bool) {
	if o == nil || IsNil(o.VendorId) {
		return nil, false
	}
	return o.VendorId, true
}

// HasVendorId returns a boolean if a field has been set.
func (o *FirmwareCatalog) HasVendorId() bool {
	if o != nil && !IsNil(o.VendorId) {
		return true
	}

	return false
}

// SetVendorId gets a reference to the given string and assigns it to the VendorId field.
func (o *FirmwareCatalog) SetVendorId(v string) {
	o.VendorId = &v
}

// GetVendorUrl returns the VendorUrl field value if set, zero value otherwise.
func (o *FirmwareCatalog) GetVendorUrl() string {
	if o == nil || IsNil(o.VendorUrl) {
		var ret string
		return ret
	}
	return *o.VendorUrl
}

// GetVendorUrlOk returns a tuple with the VendorUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareCatalog) GetVendorUrlOk() (*string, bool) {
	if o == nil || IsNil(o.VendorUrl) {
		return nil, false
	}
	return o.VendorUrl, true
}

// HasVendorUrl returns a boolean if a field has been set.
func (o *FirmwareCatalog) HasVendorUrl() bool {
	if o != nil && !IsNil(o.VendorUrl) {
		return true
	}

	return false
}

// SetVendorUrl gets a reference to the given string and assigns it to the VendorUrl field.
func (o *FirmwareCatalog) SetVendorUrl(v string) {
	o.VendorUrl = &v
}

// GetVendorReleaseTimestamp returns the VendorReleaseTimestamp field value if set, zero value otherwise.
func (o *FirmwareCatalog) GetVendorReleaseTimestamp() string {
	if o == nil || IsNil(o.VendorReleaseTimestamp) {
		var ret string
		return ret
	}
	return *o.VendorReleaseTimestamp
}

// GetVendorReleaseTimestampOk returns a tuple with the VendorReleaseTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareCatalog) GetVendorReleaseTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.VendorReleaseTimestamp) {
		return nil, false
	}
	return o.VendorReleaseTimestamp, true
}

// HasVendorReleaseTimestamp returns a boolean if a field has been set.
func (o *FirmwareCatalog) HasVendorReleaseTimestamp() bool {
	if o != nil && !IsNil(o.VendorReleaseTimestamp) {
		return true
	}

	return false
}

// SetVendorReleaseTimestamp gets a reference to the given string and assigns it to the VendorReleaseTimestamp field.
func (o *FirmwareCatalog) SetVendorReleaseTimestamp(v string) {
	o.VendorReleaseTimestamp = &v
}

// GetUpdateType returns the UpdateType field value
func (o *FirmwareCatalog) GetUpdateType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdateType
}

// GetUpdateTypeOk returns a tuple with the UpdateType field value
// and a boolean to check if the value has been set.
func (o *FirmwareCatalog) GetUpdateTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdateType, true
}

// SetUpdateType sets field value
func (o *FirmwareCatalog) SetUpdateType(v string) {
	o.UpdateType = v
}

// GetMetalsoftServerTypesSupported returns the MetalsoftServerTypesSupported field value if set, zero value otherwise.
func (o *FirmwareCatalog) GetMetalsoftServerTypesSupported() []string {
	if o == nil || IsNil(o.MetalsoftServerTypesSupported) {
		var ret []string
		return ret
	}
	return o.MetalsoftServerTypesSupported
}

// GetMetalsoftServerTypesSupportedOk returns a tuple with the MetalsoftServerTypesSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareCatalog) GetMetalsoftServerTypesSupportedOk() ([]string, bool) {
	if o == nil || IsNil(o.MetalsoftServerTypesSupported) {
		return nil, false
	}
	return o.MetalsoftServerTypesSupported, true
}

// HasMetalsoftServerTypesSupported returns a boolean if a field has been set.
func (o *FirmwareCatalog) HasMetalsoftServerTypesSupported() bool {
	if o != nil && !IsNil(o.MetalsoftServerTypesSupported) {
		return true
	}

	return false
}

// SetMetalsoftServerTypesSupported gets a reference to the given []string and assigns it to the MetalsoftServerTypesSupported field.
func (o *FirmwareCatalog) SetMetalsoftServerTypesSupported(v []string) {
	o.MetalsoftServerTypesSupported = v
}

// GetVendorServerTypesSupported returns the VendorServerTypesSupported field value if set, zero value otherwise.
func (o *FirmwareCatalog) GetVendorServerTypesSupported() []string {
	if o == nil || IsNil(o.VendorServerTypesSupported) {
		var ret []string
		return ret
	}
	return o.VendorServerTypesSupported
}

// GetVendorServerTypesSupportedOk returns a tuple with the VendorServerTypesSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareCatalog) GetVendorServerTypesSupportedOk() ([]string, bool) {
	if o == nil || IsNil(o.VendorServerTypesSupported) {
		return nil, false
	}
	return o.VendorServerTypesSupported, true
}

// HasVendorServerTypesSupported returns a boolean if a field has been set.
func (o *FirmwareCatalog) HasVendorServerTypesSupported() bool {
	if o != nil && !IsNil(o.VendorServerTypesSupported) {
		return true
	}

	return false
}

// SetVendorServerTypesSupported gets a reference to the given []string and assigns it to the VendorServerTypesSupported field.
func (o *FirmwareCatalog) SetVendorServerTypesSupported(v []string) {
	o.VendorServerTypesSupported = v
}

// GetVendorConfiguration returns the VendorConfiguration field value if set, zero value otherwise.
func (o *FirmwareCatalog) GetVendorConfiguration() map[string]interface{} {
	if o == nil || IsNil(o.VendorConfiguration) {
		var ret map[string]interface{}
		return ret
	}
	return o.VendorConfiguration
}

// GetVendorConfigurationOk returns a tuple with the VendorConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareCatalog) GetVendorConfigurationOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.VendorConfiguration) {
		return map[string]interface{}{}, false
	}
	return o.VendorConfiguration, true
}

// HasVendorConfiguration returns a boolean if a field has been set.
func (o *FirmwareCatalog) HasVendorConfiguration() bool {
	if o != nil && !IsNil(o.VendorConfiguration) {
		return true
	}

	return false
}

// SetVendorConfiguration gets a reference to the given map[string]interface{} and assigns it to the VendorConfiguration field.
func (o *FirmwareCatalog) SetVendorConfiguration(v map[string]interface{}) {
	o.VendorConfiguration = v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value
func (o *FirmwareCatalog) GetCreatedTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *FirmwareCatalog) GetCreatedTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTimestamp, true
}

// SetCreatedTimestamp sets field value
func (o *FirmwareCatalog) SetCreatedTimestamp(v string) {
	o.CreatedTimestamp = v
}

// GetLinks returns the Links field value
func (o *FirmwareCatalog) GetLinks() []Link {
	if o == nil {
		var ret []Link
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *FirmwareCatalog) GetLinksOk() ([]Link, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *FirmwareCatalog) SetLinks(v []Link) {
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
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["vendor"] = o.Vendor
	if !IsNil(o.VendorId) {
		toSerialize["vendorId"] = o.VendorId
	}
	if !IsNil(o.VendorUrl) {
		toSerialize["vendorUrl"] = o.VendorUrl
	}
	if !IsNil(o.VendorReleaseTimestamp) {
		toSerialize["vendorReleaseTimestamp"] = o.VendorReleaseTimestamp
	}
	toSerialize["updateType"] = o.UpdateType
	if !IsNil(o.MetalsoftServerTypesSupported) {
		toSerialize["metalsoftServerTypesSupported"] = o.MetalsoftServerTypesSupported
	}
	if !IsNil(o.VendorServerTypesSupported) {
		toSerialize["vendorServerTypesSupported"] = o.VendorServerTypesSupported
	}
	if !IsNil(o.VendorConfiguration) {
		toSerialize["vendorConfiguration"] = o.VendorConfiguration
	}
	toSerialize["createdTimestamp"] = o.CreatedTimestamp
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
		"id",
		"name",
		"vendor",
		"updateType",
		"createdTimestamp",
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
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "vendor")
		delete(additionalProperties, "vendorId")
		delete(additionalProperties, "vendorUrl")
		delete(additionalProperties, "vendorReleaseTimestamp")
		delete(additionalProperties, "updateType")
		delete(additionalProperties, "metalsoftServerTypesSupported")
		delete(additionalProperties, "vendorServerTypesSupported")
		delete(additionalProperties, "vendorConfiguration")
		delete(additionalProperties, "createdTimestamp")
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


