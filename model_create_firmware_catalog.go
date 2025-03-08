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
	"time"
	"fmt"
)

// checks if the CreateFirmwareCatalog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFirmwareCatalog{}

// CreateFirmwareCatalog struct for CreateFirmwareCatalog
type CreateFirmwareCatalog struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Vendor FirmwareVendorType `json:"vendor"`
	UpdateType CatalogUpdateType `json:"updateType"`
	VendorId *string `json:"vendorId,omitempty"`
	VendorUrl *string `json:"vendorUrl,omitempty"`
	VendorReleaseTimestamp *time.Time `json:"vendorReleaseTimestamp,omitempty"`
	MetalsoftServerTypesSupported []string `json:"metalsoftServerTypesSupported,omitempty"`
	// Array of the server types supported by the vendor for this catalog
	VendorServerTypesSupported []string `json:"vendorServerTypesSupported,omitempty"`
	// Record of the vendor configuration for this catalog
	VendorConfiguration map[string]interface{} `json:"vendorConfiguration,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateFirmwareCatalog CreateFirmwareCatalog

// NewCreateFirmwareCatalog instantiates a new CreateFirmwareCatalog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFirmwareCatalog(name string, vendor FirmwareVendorType, updateType CatalogUpdateType) *CreateFirmwareCatalog {
	this := CreateFirmwareCatalog{}
	this.Name = name
	this.Vendor = vendor
	this.UpdateType = updateType
	return &this
}

// NewCreateFirmwareCatalogWithDefaults instantiates a new CreateFirmwareCatalog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFirmwareCatalogWithDefaults() *CreateFirmwareCatalog {
	this := CreateFirmwareCatalog{}
	return &this
}

// GetName returns the Name field value
func (o *CreateFirmwareCatalog) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateFirmwareCatalog) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateFirmwareCatalog) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateFirmwareCatalog) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFirmwareCatalog) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateFirmwareCatalog) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateFirmwareCatalog) SetDescription(v string) {
	o.Description = &v
}

// GetVendor returns the Vendor field value
func (o *CreateFirmwareCatalog) GetVendor() FirmwareVendorType {
	if o == nil {
		var ret FirmwareVendorType
		return ret
	}

	return o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value
// and a boolean to check if the value has been set.
func (o *CreateFirmwareCatalog) GetVendorOk() (*FirmwareVendorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vendor, true
}

// SetVendor sets field value
func (o *CreateFirmwareCatalog) SetVendor(v FirmwareVendorType) {
	o.Vendor = v
}

// GetUpdateType returns the UpdateType field value
func (o *CreateFirmwareCatalog) GetUpdateType() CatalogUpdateType {
	if o == nil {
		var ret CatalogUpdateType
		return ret
	}

	return o.UpdateType
}

// GetUpdateTypeOk returns a tuple with the UpdateType field value
// and a boolean to check if the value has been set.
func (o *CreateFirmwareCatalog) GetUpdateTypeOk() (*CatalogUpdateType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdateType, true
}

// SetUpdateType sets field value
func (o *CreateFirmwareCatalog) SetUpdateType(v CatalogUpdateType) {
	o.UpdateType = v
}

// GetVendorId returns the VendorId field value if set, zero value otherwise.
func (o *CreateFirmwareCatalog) GetVendorId() string {
	if o == nil || IsNil(o.VendorId) {
		var ret string
		return ret
	}
	return *o.VendorId
}

// GetVendorIdOk returns a tuple with the VendorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFirmwareCatalog) GetVendorIdOk() (*string, bool) {
	if o == nil || IsNil(o.VendorId) {
		return nil, false
	}
	return o.VendorId, true
}

// HasVendorId returns a boolean if a field has been set.
func (o *CreateFirmwareCatalog) HasVendorId() bool {
	if o != nil && !IsNil(o.VendorId) {
		return true
	}

	return false
}

// SetVendorId gets a reference to the given string and assigns it to the VendorId field.
func (o *CreateFirmwareCatalog) SetVendorId(v string) {
	o.VendorId = &v
}

// GetVendorUrl returns the VendorUrl field value if set, zero value otherwise.
func (o *CreateFirmwareCatalog) GetVendorUrl() string {
	if o == nil || IsNil(o.VendorUrl) {
		var ret string
		return ret
	}
	return *o.VendorUrl
}

// GetVendorUrlOk returns a tuple with the VendorUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFirmwareCatalog) GetVendorUrlOk() (*string, bool) {
	if o == nil || IsNil(o.VendorUrl) {
		return nil, false
	}
	return o.VendorUrl, true
}

// HasVendorUrl returns a boolean if a field has been set.
func (o *CreateFirmwareCatalog) HasVendorUrl() bool {
	if o != nil && !IsNil(o.VendorUrl) {
		return true
	}

	return false
}

// SetVendorUrl gets a reference to the given string and assigns it to the VendorUrl field.
func (o *CreateFirmwareCatalog) SetVendorUrl(v string) {
	o.VendorUrl = &v
}

// GetVendorReleaseTimestamp returns the VendorReleaseTimestamp field value if set, zero value otherwise.
func (o *CreateFirmwareCatalog) GetVendorReleaseTimestamp() time.Time {
	if o == nil || IsNil(o.VendorReleaseTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.VendorReleaseTimestamp
}

// GetVendorReleaseTimestampOk returns a tuple with the VendorReleaseTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFirmwareCatalog) GetVendorReleaseTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.VendorReleaseTimestamp) {
		return nil, false
	}
	return o.VendorReleaseTimestamp, true
}

// HasVendorReleaseTimestamp returns a boolean if a field has been set.
func (o *CreateFirmwareCatalog) HasVendorReleaseTimestamp() bool {
	if o != nil && !IsNil(o.VendorReleaseTimestamp) {
		return true
	}

	return false
}

// SetVendorReleaseTimestamp gets a reference to the given time.Time and assigns it to the VendorReleaseTimestamp field.
func (o *CreateFirmwareCatalog) SetVendorReleaseTimestamp(v time.Time) {
	o.VendorReleaseTimestamp = &v
}

// GetMetalsoftServerTypesSupported returns the MetalsoftServerTypesSupported field value if set, zero value otherwise.
func (o *CreateFirmwareCatalog) GetMetalsoftServerTypesSupported() []string {
	if o == nil || IsNil(o.MetalsoftServerTypesSupported) {
		var ret []string
		return ret
	}
	return o.MetalsoftServerTypesSupported
}

// GetMetalsoftServerTypesSupportedOk returns a tuple with the MetalsoftServerTypesSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFirmwareCatalog) GetMetalsoftServerTypesSupportedOk() ([]string, bool) {
	if o == nil || IsNil(o.MetalsoftServerTypesSupported) {
		return nil, false
	}
	return o.MetalsoftServerTypesSupported, true
}

// HasMetalsoftServerTypesSupported returns a boolean if a field has been set.
func (o *CreateFirmwareCatalog) HasMetalsoftServerTypesSupported() bool {
	if o != nil && !IsNil(o.MetalsoftServerTypesSupported) {
		return true
	}

	return false
}

// SetMetalsoftServerTypesSupported gets a reference to the given []string and assigns it to the MetalsoftServerTypesSupported field.
func (o *CreateFirmwareCatalog) SetMetalsoftServerTypesSupported(v []string) {
	o.MetalsoftServerTypesSupported = v
}

// GetVendorServerTypesSupported returns the VendorServerTypesSupported field value if set, zero value otherwise.
func (o *CreateFirmwareCatalog) GetVendorServerTypesSupported() []string {
	if o == nil || IsNil(o.VendorServerTypesSupported) {
		var ret []string
		return ret
	}
	return o.VendorServerTypesSupported
}

// GetVendorServerTypesSupportedOk returns a tuple with the VendorServerTypesSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFirmwareCatalog) GetVendorServerTypesSupportedOk() ([]string, bool) {
	if o == nil || IsNil(o.VendorServerTypesSupported) {
		return nil, false
	}
	return o.VendorServerTypesSupported, true
}

// HasVendorServerTypesSupported returns a boolean if a field has been set.
func (o *CreateFirmwareCatalog) HasVendorServerTypesSupported() bool {
	if o != nil && !IsNil(o.VendorServerTypesSupported) {
		return true
	}

	return false
}

// SetVendorServerTypesSupported gets a reference to the given []string and assigns it to the VendorServerTypesSupported field.
func (o *CreateFirmwareCatalog) SetVendorServerTypesSupported(v []string) {
	o.VendorServerTypesSupported = v
}

// GetVendorConfiguration returns the VendorConfiguration field value if set, zero value otherwise.
func (o *CreateFirmwareCatalog) GetVendorConfiguration() map[string]interface{} {
	if o == nil || IsNil(o.VendorConfiguration) {
		var ret map[string]interface{}
		return ret
	}
	return o.VendorConfiguration
}

// GetVendorConfigurationOk returns a tuple with the VendorConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFirmwareCatalog) GetVendorConfigurationOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.VendorConfiguration) {
		return map[string]interface{}{}, false
	}
	return o.VendorConfiguration, true
}

// HasVendorConfiguration returns a boolean if a field has been set.
func (o *CreateFirmwareCatalog) HasVendorConfiguration() bool {
	if o != nil && !IsNil(o.VendorConfiguration) {
		return true
	}

	return false
}

// SetVendorConfiguration gets a reference to the given map[string]interface{} and assigns it to the VendorConfiguration field.
func (o *CreateFirmwareCatalog) SetVendorConfiguration(v map[string]interface{}) {
	o.VendorConfiguration = v
}

func (o CreateFirmwareCatalog) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFirmwareCatalog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["vendor"] = o.Vendor
	toSerialize["updateType"] = o.UpdateType
	if !IsNil(o.VendorId) {
		toSerialize["vendorId"] = o.VendorId
	}
	if !IsNil(o.VendorUrl) {
		toSerialize["vendorUrl"] = o.VendorUrl
	}
	if !IsNil(o.VendorReleaseTimestamp) {
		toSerialize["vendorReleaseTimestamp"] = o.VendorReleaseTimestamp
	}
	if !IsNil(o.MetalsoftServerTypesSupported) {
		toSerialize["metalsoftServerTypesSupported"] = o.MetalsoftServerTypesSupported
	}
	if !IsNil(o.VendorServerTypesSupported) {
		toSerialize["vendorServerTypesSupported"] = o.VendorServerTypesSupported
	}
	if !IsNil(o.VendorConfiguration) {
		toSerialize["vendorConfiguration"] = o.VendorConfiguration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateFirmwareCatalog) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"vendor",
		"updateType",
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

	varCreateFirmwareCatalog := _CreateFirmwareCatalog{}

	err = json.Unmarshal(data, &varCreateFirmwareCatalog)

	if err != nil {
		return err
	}

	*o = CreateFirmwareCatalog(varCreateFirmwareCatalog)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "vendor")
		delete(additionalProperties, "updateType")
		delete(additionalProperties, "vendorId")
		delete(additionalProperties, "vendorUrl")
		delete(additionalProperties, "vendorReleaseTimestamp")
		delete(additionalProperties, "metalsoftServerTypesSupported")
		delete(additionalProperties, "vendorServerTypesSupported")
		delete(additionalProperties, "vendorConfiguration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateFirmwareCatalog struct {
	value *CreateFirmwareCatalog
	isSet bool
}

func (v NullableCreateFirmwareCatalog) Get() *CreateFirmwareCatalog {
	return v.value
}

func (v *NullableCreateFirmwareCatalog) Set(val *CreateFirmwareCatalog) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFirmwareCatalog) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFirmwareCatalog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFirmwareCatalog(val *CreateFirmwareCatalog) *NullableCreateFirmwareCatalog {
	return &NullableCreateFirmwareCatalog{value: val, isSet: true}
}

func (v NullableCreateFirmwareCatalog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFirmwareCatalog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


