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

// checks if the UpdateVMPool type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateVMPool{}

// UpdateVMPool struct for UpdateVMPool
type UpdateVMPool struct {
	// Description of the VM Pool
	Description *string `json:"description,omitempty"`
	// Flag to indicate if the VM Pool is in maintenance mode. 1 for true, 0 for false. Default is 0.
	InMaintenance *float32 `json:"inMaintenance,omitempty"`
	// Flag to indicate if the VM Pool is experimental. 1 for true, 0 for false. Default is 0.
	IsExperimental *float32 `json:"isExperimental,omitempty"`
	// Tags for the VM Pool.
	Tags []string `json:"tags"`
	// Host of the VM Pool
	ManagementHost *string `json:"managementHost,omitempty"`
	// Port of the VM Pool
	ManagementPort *float32 `json:"managementPort,omitempty"`
	// Certificate of the VM Pool
	Certificate *string `json:"certificate,omitempty"`
	// Private key of the VM Pool
	PrivateKey *string `json:"privateKey,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateVMPool UpdateVMPool

// NewUpdateVMPool instantiates a new UpdateVMPool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVMPool(tags []string) *UpdateVMPool {
	this := UpdateVMPool{}
	this.Tags = tags
	return &this
}

// NewUpdateVMPoolWithDefaults instantiates a new UpdateVMPool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVMPoolWithDefaults() *UpdateVMPool {
	this := UpdateVMPool{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateVMPool) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVMPool) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateVMPool) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateVMPool) SetDescription(v string) {
	o.Description = &v
}

// GetInMaintenance returns the InMaintenance field value if set, zero value otherwise.
func (o *UpdateVMPool) GetInMaintenance() float32 {
	if o == nil || IsNil(o.InMaintenance) {
		var ret float32
		return ret
	}
	return *o.InMaintenance
}

// GetInMaintenanceOk returns a tuple with the InMaintenance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVMPool) GetInMaintenanceOk() (*float32, bool) {
	if o == nil || IsNil(o.InMaintenance) {
		return nil, false
	}
	return o.InMaintenance, true
}

// HasInMaintenance returns a boolean if a field has been set.
func (o *UpdateVMPool) HasInMaintenance() bool {
	if o != nil && !IsNil(o.InMaintenance) {
		return true
	}

	return false
}

// SetInMaintenance gets a reference to the given float32 and assigns it to the InMaintenance field.
func (o *UpdateVMPool) SetInMaintenance(v float32) {
	o.InMaintenance = &v
}

// GetIsExperimental returns the IsExperimental field value if set, zero value otherwise.
func (o *UpdateVMPool) GetIsExperimental() float32 {
	if o == nil || IsNil(o.IsExperimental) {
		var ret float32
		return ret
	}
	return *o.IsExperimental
}

// GetIsExperimentalOk returns a tuple with the IsExperimental field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVMPool) GetIsExperimentalOk() (*float32, bool) {
	if o == nil || IsNil(o.IsExperimental) {
		return nil, false
	}
	return o.IsExperimental, true
}

// HasIsExperimental returns a boolean if a field has been set.
func (o *UpdateVMPool) HasIsExperimental() bool {
	if o != nil && !IsNil(o.IsExperimental) {
		return true
	}

	return false
}

// SetIsExperimental gets a reference to the given float32 and assigns it to the IsExperimental field.
func (o *UpdateVMPool) SetIsExperimental(v float32) {
	o.IsExperimental = &v
}

// GetTags returns the Tags field value
func (o *UpdateVMPool) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *UpdateVMPool) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *UpdateVMPool) SetTags(v []string) {
	o.Tags = v
}

// GetManagementHost returns the ManagementHost field value if set, zero value otherwise.
func (o *UpdateVMPool) GetManagementHost() string {
	if o == nil || IsNil(o.ManagementHost) {
		var ret string
		return ret
	}
	return *o.ManagementHost
}

// GetManagementHostOk returns a tuple with the ManagementHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVMPool) GetManagementHostOk() (*string, bool) {
	if o == nil || IsNil(o.ManagementHost) {
		return nil, false
	}
	return o.ManagementHost, true
}

// HasManagementHost returns a boolean if a field has been set.
func (o *UpdateVMPool) HasManagementHost() bool {
	if o != nil && !IsNil(o.ManagementHost) {
		return true
	}

	return false
}

// SetManagementHost gets a reference to the given string and assigns it to the ManagementHost field.
func (o *UpdateVMPool) SetManagementHost(v string) {
	o.ManagementHost = &v
}

// GetManagementPort returns the ManagementPort field value if set, zero value otherwise.
func (o *UpdateVMPool) GetManagementPort() float32 {
	if o == nil || IsNil(o.ManagementPort) {
		var ret float32
		return ret
	}
	return *o.ManagementPort
}

// GetManagementPortOk returns a tuple with the ManagementPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVMPool) GetManagementPortOk() (*float32, bool) {
	if o == nil || IsNil(o.ManagementPort) {
		return nil, false
	}
	return o.ManagementPort, true
}

// HasManagementPort returns a boolean if a field has been set.
func (o *UpdateVMPool) HasManagementPort() bool {
	if o != nil && !IsNil(o.ManagementPort) {
		return true
	}

	return false
}

// SetManagementPort gets a reference to the given float32 and assigns it to the ManagementPort field.
func (o *UpdateVMPool) SetManagementPort(v float32) {
	o.ManagementPort = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *UpdateVMPool) GetCertificate() string {
	if o == nil || IsNil(o.Certificate) {
		var ret string
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVMPool) GetCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.Certificate) {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *UpdateVMPool) HasCertificate() bool {
	if o != nil && !IsNil(o.Certificate) {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given string and assigns it to the Certificate field.
func (o *UpdateVMPool) SetCertificate(v string) {
	o.Certificate = &v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *UpdateVMPool) GetPrivateKey() string {
	if o == nil || IsNil(o.PrivateKey) {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVMPool) GetPrivateKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateKey) {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *UpdateVMPool) HasPrivateKey() bool {
	if o != nil && !IsNil(o.PrivateKey) {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *UpdateVMPool) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

func (o UpdateVMPool) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateVMPool) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.InMaintenance) {
		toSerialize["inMaintenance"] = o.InMaintenance
	}
	if !IsNil(o.IsExperimental) {
		toSerialize["isExperimental"] = o.IsExperimental
	}
	toSerialize["tags"] = o.Tags
	if !IsNil(o.ManagementHost) {
		toSerialize["managementHost"] = o.ManagementHost
	}
	if !IsNil(o.ManagementPort) {
		toSerialize["managementPort"] = o.ManagementPort
	}
	if !IsNil(o.Certificate) {
		toSerialize["certificate"] = o.Certificate
	}
	if !IsNil(o.PrivateKey) {
		toSerialize["privateKey"] = o.PrivateKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateVMPool) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tags",
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

	varUpdateVMPool := _UpdateVMPool{}

	err = json.Unmarshal(data, &varUpdateVMPool)

	if err != nil {
		return err
	}

	*o = UpdateVMPool(varUpdateVMPool)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "inMaintenance")
		delete(additionalProperties, "isExperimental")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "managementHost")
		delete(additionalProperties, "managementPort")
		delete(additionalProperties, "certificate")
		delete(additionalProperties, "privateKey")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateVMPool struct {
	value *UpdateVMPool
	isSet bool
}

func (v NullableUpdateVMPool) Get() *UpdateVMPool {
	return v.value
}

func (v *NullableUpdateVMPool) Set(val *UpdateVMPool) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateVMPool) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateVMPool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateVMPool(val *UpdateVMPool) *NullableUpdateVMPool {
	return &NullableUpdateVMPool{value: val, isSet: true}
}

func (v NullableUpdateVMPool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateVMPool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


