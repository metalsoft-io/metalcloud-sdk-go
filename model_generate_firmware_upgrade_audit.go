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

// checks if the GenerateFirmwareUpgradeAudit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateFirmwareUpgradeAudit{}

// GenerateFirmwareUpgradeAudit struct for GenerateFirmwareUpgradeAudit
type GenerateFirmwareUpgradeAudit struct {
	// The list of server ids for which firmware upgrade audit should be generated.
	ServerIds []float32 `json:"serverIds"`
	// Filter the available firmware upgrades using the specified baseline id.
	BaselineId *float32 `json:"baselineId,omitempty"`
	// Filter the available firmware upgrades using the specified os template label.
	OsTemplateLabel *string `json:"osTemplateLabel,omitempty"`
	// Filter the available firmware upgrades using the specified server type name.
	ServerTypeName *string `json:"serverTypeName,omitempty"`
	// Filter the available firmware upgrades using the specified site id.
	SiteId *float32 `json:"siteId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GenerateFirmwareUpgradeAudit GenerateFirmwareUpgradeAudit

// NewGenerateFirmwareUpgradeAudit instantiates a new GenerateFirmwareUpgradeAudit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateFirmwareUpgradeAudit(serverIds []float32) *GenerateFirmwareUpgradeAudit {
	this := GenerateFirmwareUpgradeAudit{}
	this.ServerIds = serverIds
	return &this
}

// NewGenerateFirmwareUpgradeAuditWithDefaults instantiates a new GenerateFirmwareUpgradeAudit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateFirmwareUpgradeAuditWithDefaults() *GenerateFirmwareUpgradeAudit {
	this := GenerateFirmwareUpgradeAudit{}
	return &this
}

// GetServerIds returns the ServerIds field value
func (o *GenerateFirmwareUpgradeAudit) GetServerIds() []float32 {
	if o == nil {
		var ret []float32
		return ret
	}

	return o.ServerIds
}

// GetServerIdsOk returns a tuple with the ServerIds field value
// and a boolean to check if the value has been set.
func (o *GenerateFirmwareUpgradeAudit) GetServerIdsOk() ([]float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServerIds, true
}

// SetServerIds sets field value
func (o *GenerateFirmwareUpgradeAudit) SetServerIds(v []float32) {
	o.ServerIds = v
}

// GetBaselineId returns the BaselineId field value if set, zero value otherwise.
func (o *GenerateFirmwareUpgradeAudit) GetBaselineId() float32 {
	if o == nil || IsNil(o.BaselineId) {
		var ret float32
		return ret
	}
	return *o.BaselineId
}

// GetBaselineIdOk returns a tuple with the BaselineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateFirmwareUpgradeAudit) GetBaselineIdOk() (*float32, bool) {
	if o == nil || IsNil(o.BaselineId) {
		return nil, false
	}
	return o.BaselineId, true
}

// HasBaselineId returns a boolean if a field has been set.
func (o *GenerateFirmwareUpgradeAudit) HasBaselineId() bool {
	if o != nil && !IsNil(o.BaselineId) {
		return true
	}

	return false
}

// SetBaselineId gets a reference to the given float32 and assigns it to the BaselineId field.
func (o *GenerateFirmwareUpgradeAudit) SetBaselineId(v float32) {
	o.BaselineId = &v
}

// GetOsTemplateLabel returns the OsTemplateLabel field value if set, zero value otherwise.
func (o *GenerateFirmwareUpgradeAudit) GetOsTemplateLabel() string {
	if o == nil || IsNil(o.OsTemplateLabel) {
		var ret string
		return ret
	}
	return *o.OsTemplateLabel
}

// GetOsTemplateLabelOk returns a tuple with the OsTemplateLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateFirmwareUpgradeAudit) GetOsTemplateLabelOk() (*string, bool) {
	if o == nil || IsNil(o.OsTemplateLabel) {
		return nil, false
	}
	return o.OsTemplateLabel, true
}

// HasOsTemplateLabel returns a boolean if a field has been set.
func (o *GenerateFirmwareUpgradeAudit) HasOsTemplateLabel() bool {
	if o != nil && !IsNil(o.OsTemplateLabel) {
		return true
	}

	return false
}

// SetOsTemplateLabel gets a reference to the given string and assigns it to the OsTemplateLabel field.
func (o *GenerateFirmwareUpgradeAudit) SetOsTemplateLabel(v string) {
	o.OsTemplateLabel = &v
}

// GetServerTypeName returns the ServerTypeName field value if set, zero value otherwise.
func (o *GenerateFirmwareUpgradeAudit) GetServerTypeName() string {
	if o == nil || IsNil(o.ServerTypeName) {
		var ret string
		return ret
	}
	return *o.ServerTypeName
}

// GetServerTypeNameOk returns a tuple with the ServerTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateFirmwareUpgradeAudit) GetServerTypeNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServerTypeName) {
		return nil, false
	}
	return o.ServerTypeName, true
}

// HasServerTypeName returns a boolean if a field has been set.
func (o *GenerateFirmwareUpgradeAudit) HasServerTypeName() bool {
	if o != nil && !IsNil(o.ServerTypeName) {
		return true
	}

	return false
}

// SetServerTypeName gets a reference to the given string and assigns it to the ServerTypeName field.
func (o *GenerateFirmwareUpgradeAudit) SetServerTypeName(v string) {
	o.ServerTypeName = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *GenerateFirmwareUpgradeAudit) GetSiteId() float32 {
	if o == nil || IsNil(o.SiteId) {
		var ret float32
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateFirmwareUpgradeAudit) GetSiteIdOk() (*float32, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *GenerateFirmwareUpgradeAudit) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given float32 and assigns it to the SiteId field.
func (o *GenerateFirmwareUpgradeAudit) SetSiteId(v float32) {
	o.SiteId = &v
}

func (o GenerateFirmwareUpgradeAudit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenerateFirmwareUpgradeAudit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serverIds"] = o.ServerIds
	if !IsNil(o.BaselineId) {
		toSerialize["baselineId"] = o.BaselineId
	}
	if !IsNil(o.OsTemplateLabel) {
		toSerialize["osTemplateLabel"] = o.OsTemplateLabel
	}
	if !IsNil(o.ServerTypeName) {
		toSerialize["serverTypeName"] = o.ServerTypeName
	}
	if !IsNil(o.SiteId) {
		toSerialize["siteId"] = o.SiteId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GenerateFirmwareUpgradeAudit) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"serverIds",
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

	varGenerateFirmwareUpgradeAudit := _GenerateFirmwareUpgradeAudit{}

	err = json.Unmarshal(data, &varGenerateFirmwareUpgradeAudit)

	if err != nil {
		return err
	}

	*o = GenerateFirmwareUpgradeAudit(varGenerateFirmwareUpgradeAudit)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "serverIds")
		delete(additionalProperties, "baselineId")
		delete(additionalProperties, "osTemplateLabel")
		delete(additionalProperties, "serverTypeName")
		delete(additionalProperties, "siteId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGenerateFirmwareUpgradeAudit struct {
	value *GenerateFirmwareUpgradeAudit
	isSet bool
}

func (v NullableGenerateFirmwareUpgradeAudit) Get() *GenerateFirmwareUpgradeAudit {
	return v.value
}

func (v *NullableGenerateFirmwareUpgradeAudit) Set(val *GenerateFirmwareUpgradeAudit) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateFirmwareUpgradeAudit) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateFirmwareUpgradeAudit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateFirmwareUpgradeAudit(val *GenerateFirmwareUpgradeAudit) *NullableGenerateFirmwareUpgradeAudit {
	return &NullableGenerateFirmwareUpgradeAudit{value: val, isSet: true}
}

func (v NullableGenerateFirmwareUpgradeAudit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateFirmwareUpgradeAudit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


