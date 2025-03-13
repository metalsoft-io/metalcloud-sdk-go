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

// checks if the ExtensionInstanceConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtensionInstanceConfiguration{}

// ExtensionInstanceConfiguration struct for ExtensionInstanceConfiguration
type ExtensionInstanceConfiguration struct {
	// Revision number
	Revision float32 `json:"revision"`
	// The extension instance label. Will be automatically generated if not provided.
	Label string `json:"label"`
	// Flag specifying if the extension instance supports automatic management.
	AutomaticManagement float32 `json:"automaticManagement"`
	// Subdomain of the Extension Instance.
	Subdomain *string `json:"subdomain,omitempty"`
	// Id of the DNS subdomain for the Extension Instance.
	DnsSubdomainChangeId *float32 `json:"dnsSubdomainChangeId,omitempty"`
	// Deploy type of the Extension Instance
	DeployType string `json:"deployType"`
	// Deploy status of the Extension Instance
	DeployStatus string `json:"deployStatus"`
	// Id of the deployment for the Extension Instance.
	InfrastructureDeployId *float32 `json:"infrastructureDeployId,omitempty"`
	// Timestamp of the Extension Instance last update.
	UpdatedTimestamp string `json:"updatedTimestamp"`
	AdditionalProperties map[string]interface{}
}

type _ExtensionInstanceConfiguration ExtensionInstanceConfiguration

// NewExtensionInstanceConfiguration instantiates a new ExtensionInstanceConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtensionInstanceConfiguration(revision float32, label string, automaticManagement float32, deployType string, deployStatus string, updatedTimestamp string) *ExtensionInstanceConfiguration {
	this := ExtensionInstanceConfiguration{}
	this.Revision = revision
	this.Label = label
	this.AutomaticManagement = automaticManagement
	this.DeployType = deployType
	this.DeployStatus = deployStatus
	this.UpdatedTimestamp = updatedTimestamp
	return &this
}

// NewExtensionInstanceConfigurationWithDefaults instantiates a new ExtensionInstanceConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionInstanceConfigurationWithDefaults() *ExtensionInstanceConfiguration {
	this := ExtensionInstanceConfiguration{}
	var deployType string = "create"
	this.DeployType = deployType
	var deployStatus string = "not_started"
	this.DeployStatus = deployStatus
	return &this
}

// GetRevision returns the Revision field value
func (o *ExtensionInstanceConfiguration) GetRevision() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *ExtensionInstanceConfiguration) GetRevisionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *ExtensionInstanceConfiguration) SetRevision(v float32) {
	o.Revision = v
}

// GetLabel returns the Label field value
func (o *ExtensionInstanceConfiguration) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *ExtensionInstanceConfiguration) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *ExtensionInstanceConfiguration) SetLabel(v string) {
	o.Label = v
}

// GetAutomaticManagement returns the AutomaticManagement field value
func (o *ExtensionInstanceConfiguration) GetAutomaticManagement() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AutomaticManagement
}

// GetAutomaticManagementOk returns a tuple with the AutomaticManagement field value
// and a boolean to check if the value has been set.
func (o *ExtensionInstanceConfiguration) GetAutomaticManagementOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutomaticManagement, true
}

// SetAutomaticManagement sets field value
func (o *ExtensionInstanceConfiguration) SetAutomaticManagement(v float32) {
	o.AutomaticManagement = v
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *ExtensionInstanceConfiguration) GetSubdomain() string {
	if o == nil || IsNil(o.Subdomain) {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionInstanceConfiguration) GetSubdomainOk() (*string, bool) {
	if o == nil || IsNil(o.Subdomain) {
		return nil, false
	}
	return o.Subdomain, true
}

// HasSubdomain returns a boolean if a field has been set.
func (o *ExtensionInstanceConfiguration) HasSubdomain() bool {
	if o != nil && !IsNil(o.Subdomain) {
		return true
	}

	return false
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *ExtensionInstanceConfiguration) SetSubdomain(v string) {
	o.Subdomain = &v
}

// GetDnsSubdomainChangeId returns the DnsSubdomainChangeId field value if set, zero value otherwise.
func (o *ExtensionInstanceConfiguration) GetDnsSubdomainChangeId() float32 {
	if o == nil || IsNil(o.DnsSubdomainChangeId) {
		var ret float32
		return ret
	}
	return *o.DnsSubdomainChangeId
}

// GetDnsSubdomainChangeIdOk returns a tuple with the DnsSubdomainChangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionInstanceConfiguration) GetDnsSubdomainChangeIdOk() (*float32, bool) {
	if o == nil || IsNil(o.DnsSubdomainChangeId) {
		return nil, false
	}
	return o.DnsSubdomainChangeId, true
}

// HasDnsSubdomainChangeId returns a boolean if a field has been set.
func (o *ExtensionInstanceConfiguration) HasDnsSubdomainChangeId() bool {
	if o != nil && !IsNil(o.DnsSubdomainChangeId) {
		return true
	}

	return false
}

// SetDnsSubdomainChangeId gets a reference to the given float32 and assigns it to the DnsSubdomainChangeId field.
func (o *ExtensionInstanceConfiguration) SetDnsSubdomainChangeId(v float32) {
	o.DnsSubdomainChangeId = &v
}

// GetDeployType returns the DeployType field value
func (o *ExtensionInstanceConfiguration) GetDeployType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeployType
}

// GetDeployTypeOk returns a tuple with the DeployType field value
// and a boolean to check if the value has been set.
func (o *ExtensionInstanceConfiguration) GetDeployTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeployType, true
}

// SetDeployType sets field value
func (o *ExtensionInstanceConfiguration) SetDeployType(v string) {
	o.DeployType = v
}

// GetDeployStatus returns the DeployStatus field value
func (o *ExtensionInstanceConfiguration) GetDeployStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeployStatus
}

// GetDeployStatusOk returns a tuple with the DeployStatus field value
// and a boolean to check if the value has been set.
func (o *ExtensionInstanceConfiguration) GetDeployStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeployStatus, true
}

// SetDeployStatus sets field value
func (o *ExtensionInstanceConfiguration) SetDeployStatus(v string) {
	o.DeployStatus = v
}

// GetInfrastructureDeployId returns the InfrastructureDeployId field value if set, zero value otherwise.
func (o *ExtensionInstanceConfiguration) GetInfrastructureDeployId() float32 {
	if o == nil || IsNil(o.InfrastructureDeployId) {
		var ret float32
		return ret
	}
	return *o.InfrastructureDeployId
}

// GetInfrastructureDeployIdOk returns a tuple with the InfrastructureDeployId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionInstanceConfiguration) GetInfrastructureDeployIdOk() (*float32, bool) {
	if o == nil || IsNil(o.InfrastructureDeployId) {
		return nil, false
	}
	return o.InfrastructureDeployId, true
}

// HasInfrastructureDeployId returns a boolean if a field has been set.
func (o *ExtensionInstanceConfiguration) HasInfrastructureDeployId() bool {
	if o != nil && !IsNil(o.InfrastructureDeployId) {
		return true
	}

	return false
}

// SetInfrastructureDeployId gets a reference to the given float32 and assigns it to the InfrastructureDeployId field.
func (o *ExtensionInstanceConfiguration) SetInfrastructureDeployId(v float32) {
	o.InfrastructureDeployId = &v
}

// GetUpdatedTimestamp returns the UpdatedTimestamp field value
func (o *ExtensionInstanceConfiguration) GetUpdatedTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedTimestamp
}

// GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *ExtensionInstanceConfiguration) GetUpdatedTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedTimestamp, true
}

// SetUpdatedTimestamp sets field value
func (o *ExtensionInstanceConfiguration) SetUpdatedTimestamp(v string) {
	o.UpdatedTimestamp = v
}

func (o ExtensionInstanceConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtensionInstanceConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["revision"] = o.Revision
	toSerialize["label"] = o.Label
	toSerialize["automaticManagement"] = o.AutomaticManagement
	if !IsNil(o.Subdomain) {
		toSerialize["subdomain"] = o.Subdomain
	}
	if !IsNil(o.DnsSubdomainChangeId) {
		toSerialize["dnsSubdomainChangeId"] = o.DnsSubdomainChangeId
	}
	toSerialize["deployType"] = o.DeployType
	toSerialize["deployStatus"] = o.DeployStatus
	if !IsNil(o.InfrastructureDeployId) {
		toSerialize["infrastructureDeployId"] = o.InfrastructureDeployId
	}
	toSerialize["updatedTimestamp"] = o.UpdatedTimestamp

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExtensionInstanceConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"revision",
		"label",
		"automaticManagement",
		"deployType",
		"deployStatus",
		"updatedTimestamp",
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

	varExtensionInstanceConfiguration := _ExtensionInstanceConfiguration{}

	err = json.Unmarshal(data, &varExtensionInstanceConfiguration)

	if err != nil {
		return err
	}

	*o = ExtensionInstanceConfiguration(varExtensionInstanceConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "revision")
		delete(additionalProperties, "label")
		delete(additionalProperties, "automaticManagement")
		delete(additionalProperties, "subdomain")
		delete(additionalProperties, "dnsSubdomainChangeId")
		delete(additionalProperties, "deployType")
		delete(additionalProperties, "deployStatus")
		delete(additionalProperties, "infrastructureDeployId")
		delete(additionalProperties, "updatedTimestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExtensionInstanceConfiguration struct {
	value *ExtensionInstanceConfiguration
	isSet bool
}

func (v NullableExtensionInstanceConfiguration) Get() *ExtensionInstanceConfiguration {
	return v.value
}

func (v *NullableExtensionInstanceConfiguration) Set(val *ExtensionInstanceConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionInstanceConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionInstanceConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionInstanceConfiguration(val *ExtensionInstanceConfiguration) *NullableExtensionInstanceConfiguration {
	return &NullableExtensionInstanceConfiguration{value: val, isSet: true}
}

func (v NullableExtensionInstanceConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionInstanceConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


