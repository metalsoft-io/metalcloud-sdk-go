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

// checks if the VMInstanceGroupConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VMInstanceGroupConfiguration{}

// VMInstanceGroupConfiguration struct for VMInstanceGroupConfiguration
type VMInstanceGroupConfiguration struct {
	// Revision of the VM Instance Group Configuration
	Revision float32 `json:"revision"`
	// Name of the VM Instance Group.
	Label string `json:"label"`
	// Subdomain of the VM Instance Group.
	Subdomain *string `json:"subdomain,omitempty"`
	// Id of the DNS subdomain for the VM Instance Group.
	DnsSubdomainChangeId *float32 `json:"dnsSubdomainChangeId,omitempty"`
	InstanceCount *float32 `json:"instanceCount,omitempty"`
	// Deploy type of the VM Instance Group
	DeployType string `json:"deployType"`
	// Deploy status of the VM Instance Group
	DeployStatus string `json:"deployStatus"`
	// Id of the deployment for the VM Instance Group.
	InfrastructureDeployId *float32 `json:"infrastructureDeployId,omitempty"`
	// Custom variables for the VM Instance Group.
	CustomVariables map[string]interface{} `json:"customVariables,omitempty"`
	// Timestamp of the VM Instance Group last update.
	UpdatedTimestamp string `json:"updatedTimestamp"`
	AdditionalProperties map[string]interface{}
}

type _VMInstanceGroupConfiguration VMInstanceGroupConfiguration

// NewVMInstanceGroupConfiguration instantiates a new VMInstanceGroupConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVMInstanceGroupConfiguration(revision float32, label string, deployType string, deployStatus string, updatedTimestamp string) *VMInstanceGroupConfiguration {
	this := VMInstanceGroupConfiguration{}
	this.Revision = revision
	this.Label = label
	var instanceCount float32 = 1
	this.InstanceCount = &instanceCount
	this.DeployType = deployType
	this.DeployStatus = deployStatus
	this.UpdatedTimestamp = updatedTimestamp
	return &this
}

// NewVMInstanceGroupConfigurationWithDefaults instantiates a new VMInstanceGroupConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVMInstanceGroupConfigurationWithDefaults() *VMInstanceGroupConfiguration {
	this := VMInstanceGroupConfiguration{}
	var instanceCount float32 = 1
	this.InstanceCount = &instanceCount
	var deployType string = "create"
	this.DeployType = deployType
	var deployStatus string = "not_started"
	this.DeployStatus = deployStatus
	return &this
}

// GetRevision returns the Revision field value
func (o *VMInstanceGroupConfiguration) GetRevision() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *VMInstanceGroupConfiguration) GetRevisionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *VMInstanceGroupConfiguration) SetRevision(v float32) {
	o.Revision = v
}

// GetLabel returns the Label field value
func (o *VMInstanceGroupConfiguration) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *VMInstanceGroupConfiguration) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *VMInstanceGroupConfiguration) SetLabel(v string) {
	o.Label = v
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *VMInstanceGroupConfiguration) GetSubdomain() string {
	if o == nil || IsNil(o.Subdomain) {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMInstanceGroupConfiguration) GetSubdomainOk() (*string, bool) {
	if o == nil || IsNil(o.Subdomain) {
		return nil, false
	}
	return o.Subdomain, true
}

// HasSubdomain returns a boolean if a field has been set.
func (o *VMInstanceGroupConfiguration) HasSubdomain() bool {
	if o != nil && !IsNil(o.Subdomain) {
		return true
	}

	return false
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *VMInstanceGroupConfiguration) SetSubdomain(v string) {
	o.Subdomain = &v
}

// GetDnsSubdomainChangeId returns the DnsSubdomainChangeId field value if set, zero value otherwise.
func (o *VMInstanceGroupConfiguration) GetDnsSubdomainChangeId() float32 {
	if o == nil || IsNil(o.DnsSubdomainChangeId) {
		var ret float32
		return ret
	}
	return *o.DnsSubdomainChangeId
}

// GetDnsSubdomainChangeIdOk returns a tuple with the DnsSubdomainChangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMInstanceGroupConfiguration) GetDnsSubdomainChangeIdOk() (*float32, bool) {
	if o == nil || IsNil(o.DnsSubdomainChangeId) {
		return nil, false
	}
	return o.DnsSubdomainChangeId, true
}

// HasDnsSubdomainChangeId returns a boolean if a field has been set.
func (o *VMInstanceGroupConfiguration) HasDnsSubdomainChangeId() bool {
	if o != nil && !IsNil(o.DnsSubdomainChangeId) {
		return true
	}

	return false
}

// SetDnsSubdomainChangeId gets a reference to the given float32 and assigns it to the DnsSubdomainChangeId field.
func (o *VMInstanceGroupConfiguration) SetDnsSubdomainChangeId(v float32) {
	o.DnsSubdomainChangeId = &v
}

// GetInstanceCount returns the InstanceCount field value if set, zero value otherwise.
func (o *VMInstanceGroupConfiguration) GetInstanceCount() float32 {
	if o == nil || IsNil(o.InstanceCount) {
		var ret float32
		return ret
	}
	return *o.InstanceCount
}

// GetInstanceCountOk returns a tuple with the InstanceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMInstanceGroupConfiguration) GetInstanceCountOk() (*float32, bool) {
	if o == nil || IsNil(o.InstanceCount) {
		return nil, false
	}
	return o.InstanceCount, true
}

// HasInstanceCount returns a boolean if a field has been set.
func (o *VMInstanceGroupConfiguration) HasInstanceCount() bool {
	if o != nil && !IsNil(o.InstanceCount) {
		return true
	}

	return false
}

// SetInstanceCount gets a reference to the given float32 and assigns it to the InstanceCount field.
func (o *VMInstanceGroupConfiguration) SetInstanceCount(v float32) {
	o.InstanceCount = &v
}

// GetDeployType returns the DeployType field value
func (o *VMInstanceGroupConfiguration) GetDeployType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeployType
}

// GetDeployTypeOk returns a tuple with the DeployType field value
// and a boolean to check if the value has been set.
func (o *VMInstanceGroupConfiguration) GetDeployTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeployType, true
}

// SetDeployType sets field value
func (o *VMInstanceGroupConfiguration) SetDeployType(v string) {
	o.DeployType = v
}

// GetDeployStatus returns the DeployStatus field value
func (o *VMInstanceGroupConfiguration) GetDeployStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeployStatus
}

// GetDeployStatusOk returns a tuple with the DeployStatus field value
// and a boolean to check if the value has been set.
func (o *VMInstanceGroupConfiguration) GetDeployStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeployStatus, true
}

// SetDeployStatus sets field value
func (o *VMInstanceGroupConfiguration) SetDeployStatus(v string) {
	o.DeployStatus = v
}

// GetInfrastructureDeployId returns the InfrastructureDeployId field value if set, zero value otherwise.
func (o *VMInstanceGroupConfiguration) GetInfrastructureDeployId() float32 {
	if o == nil || IsNil(o.InfrastructureDeployId) {
		var ret float32
		return ret
	}
	return *o.InfrastructureDeployId
}

// GetInfrastructureDeployIdOk returns a tuple with the InfrastructureDeployId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMInstanceGroupConfiguration) GetInfrastructureDeployIdOk() (*float32, bool) {
	if o == nil || IsNil(o.InfrastructureDeployId) {
		return nil, false
	}
	return o.InfrastructureDeployId, true
}

// HasInfrastructureDeployId returns a boolean if a field has been set.
func (o *VMInstanceGroupConfiguration) HasInfrastructureDeployId() bool {
	if o != nil && !IsNil(o.InfrastructureDeployId) {
		return true
	}

	return false
}

// SetInfrastructureDeployId gets a reference to the given float32 and assigns it to the InfrastructureDeployId field.
func (o *VMInstanceGroupConfiguration) SetInfrastructureDeployId(v float32) {
	o.InfrastructureDeployId = &v
}

// GetCustomVariables returns the CustomVariables field value if set, zero value otherwise.
func (o *VMInstanceGroupConfiguration) GetCustomVariables() map[string]interface{} {
	if o == nil || IsNil(o.CustomVariables) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomVariables
}

// GetCustomVariablesOk returns a tuple with the CustomVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMInstanceGroupConfiguration) GetCustomVariablesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomVariables) {
		return map[string]interface{}{}, false
	}
	return o.CustomVariables, true
}

// HasCustomVariables returns a boolean if a field has been set.
func (o *VMInstanceGroupConfiguration) HasCustomVariables() bool {
	if o != nil && !IsNil(o.CustomVariables) {
		return true
	}

	return false
}

// SetCustomVariables gets a reference to the given map[string]interface{} and assigns it to the CustomVariables field.
func (o *VMInstanceGroupConfiguration) SetCustomVariables(v map[string]interface{}) {
	o.CustomVariables = v
}

// GetUpdatedTimestamp returns the UpdatedTimestamp field value
func (o *VMInstanceGroupConfiguration) GetUpdatedTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedTimestamp
}

// GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *VMInstanceGroupConfiguration) GetUpdatedTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedTimestamp, true
}

// SetUpdatedTimestamp sets field value
func (o *VMInstanceGroupConfiguration) SetUpdatedTimestamp(v string) {
	o.UpdatedTimestamp = v
}

func (o VMInstanceGroupConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VMInstanceGroupConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["revision"] = o.Revision
	toSerialize["label"] = o.Label
	if !IsNil(o.Subdomain) {
		toSerialize["subdomain"] = o.Subdomain
	}
	if !IsNil(o.DnsSubdomainChangeId) {
		toSerialize["dnsSubdomainChangeId"] = o.DnsSubdomainChangeId
	}
	if !IsNil(o.InstanceCount) {
		toSerialize["instanceCount"] = o.InstanceCount
	}
	toSerialize["deployType"] = o.DeployType
	toSerialize["deployStatus"] = o.DeployStatus
	if !IsNil(o.InfrastructureDeployId) {
		toSerialize["infrastructureDeployId"] = o.InfrastructureDeployId
	}
	if !IsNil(o.CustomVariables) {
		toSerialize["customVariables"] = o.CustomVariables
	}
	toSerialize["updatedTimestamp"] = o.UpdatedTimestamp

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VMInstanceGroupConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"revision",
		"label",
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

	varVMInstanceGroupConfiguration := _VMInstanceGroupConfiguration{}

	err = json.Unmarshal(data, &varVMInstanceGroupConfiguration)

	if err != nil {
		return err
	}

	*o = VMInstanceGroupConfiguration(varVMInstanceGroupConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "revision")
		delete(additionalProperties, "label")
		delete(additionalProperties, "subdomain")
		delete(additionalProperties, "dnsSubdomainChangeId")
		delete(additionalProperties, "instanceCount")
		delete(additionalProperties, "deployType")
		delete(additionalProperties, "deployStatus")
		delete(additionalProperties, "infrastructureDeployId")
		delete(additionalProperties, "customVariables")
		delete(additionalProperties, "updatedTimestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVMInstanceGroupConfiguration struct {
	value *VMInstanceGroupConfiguration
	isSet bool
}

func (v NullableVMInstanceGroupConfiguration) Get() *VMInstanceGroupConfiguration {
	return v.value
}

func (v *NullableVMInstanceGroupConfiguration) Set(val *VMInstanceGroupConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableVMInstanceGroupConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableVMInstanceGroupConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVMInstanceGroupConfiguration(val *VMInstanceGroupConfiguration) *NullableVMInstanceGroupConfiguration {
	return &NullableVMInstanceGroupConfiguration{value: val, isSet: true}
}

func (v NullableVMInstanceGroupConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVMInstanceGroupConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


