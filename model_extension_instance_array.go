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

// checks if the ExtensionInstanceArray type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtensionInstanceArray{}

// ExtensionInstanceArray struct for ExtensionInstanceArray
type ExtensionInstanceArray struct {
	// Label of the instance array.
	Label string `json:"label"`
	// Instance count value or reference to variable.
	InstanceCount string `json:"instanceCount"`
	// Server type variable reference.
	ServerType string `json:"serverType"`
	// OS template variable reference.
	OsTemplate string `json:"osTemplate"`
	// Connected shared drive arrays.
	ConnectedSharedDrives []string `json:"connectedSharedDrives,omitempty"`
	// Custom variables. The value may be a reference to an input variable.
	CustomVariables []CustomVariable `json:"customVariables,omitempty"`
	// Labels of instance arrays this one depends on.
	Dependencies []string `json:"dependencies,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExtensionInstanceArray ExtensionInstanceArray

// NewExtensionInstanceArray instantiates a new ExtensionInstanceArray object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtensionInstanceArray(label string, instanceCount string, serverType string, osTemplate string) *ExtensionInstanceArray {
	this := ExtensionInstanceArray{}
	this.Label = label
	this.InstanceCount = instanceCount
	this.ServerType = serverType
	this.OsTemplate = osTemplate
	return &this
}

// NewExtensionInstanceArrayWithDefaults instantiates a new ExtensionInstanceArray object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionInstanceArrayWithDefaults() *ExtensionInstanceArray {
	this := ExtensionInstanceArray{}
	return &this
}

// GetLabel returns the Label field value
func (o *ExtensionInstanceArray) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *ExtensionInstanceArray) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *ExtensionInstanceArray) SetLabel(v string) {
	o.Label = v
}

// GetInstanceCount returns the InstanceCount field value
func (o *ExtensionInstanceArray) GetInstanceCount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceCount
}

// GetInstanceCountOk returns a tuple with the InstanceCount field value
// and a boolean to check if the value has been set.
func (o *ExtensionInstanceArray) GetInstanceCountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceCount, true
}

// SetInstanceCount sets field value
func (o *ExtensionInstanceArray) SetInstanceCount(v string) {
	o.InstanceCount = v
}

// GetServerType returns the ServerType field value
func (o *ExtensionInstanceArray) GetServerType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerType
}

// GetServerTypeOk returns a tuple with the ServerType field value
// and a boolean to check if the value has been set.
func (o *ExtensionInstanceArray) GetServerTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerType, true
}

// SetServerType sets field value
func (o *ExtensionInstanceArray) SetServerType(v string) {
	o.ServerType = v
}

// GetOsTemplate returns the OsTemplate field value
func (o *ExtensionInstanceArray) GetOsTemplate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OsTemplate
}

// GetOsTemplateOk returns a tuple with the OsTemplate field value
// and a boolean to check if the value has been set.
func (o *ExtensionInstanceArray) GetOsTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsTemplate, true
}

// SetOsTemplate sets field value
func (o *ExtensionInstanceArray) SetOsTemplate(v string) {
	o.OsTemplate = v
}

// GetConnectedSharedDrives returns the ConnectedSharedDrives field value if set, zero value otherwise.
func (o *ExtensionInstanceArray) GetConnectedSharedDrives() []string {
	if o == nil || IsNil(o.ConnectedSharedDrives) {
		var ret []string
		return ret
	}
	return o.ConnectedSharedDrives
}

// GetConnectedSharedDrivesOk returns a tuple with the ConnectedSharedDrives field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionInstanceArray) GetConnectedSharedDrivesOk() ([]string, bool) {
	if o == nil || IsNil(o.ConnectedSharedDrives) {
		return nil, false
	}
	return o.ConnectedSharedDrives, true
}

// HasConnectedSharedDrives returns a boolean if a field has been set.
func (o *ExtensionInstanceArray) HasConnectedSharedDrives() bool {
	if o != nil && !IsNil(o.ConnectedSharedDrives) {
		return true
	}

	return false
}

// SetConnectedSharedDrives gets a reference to the given []string and assigns it to the ConnectedSharedDrives field.
func (o *ExtensionInstanceArray) SetConnectedSharedDrives(v []string) {
	o.ConnectedSharedDrives = v
}

// GetCustomVariables returns the CustomVariables field value if set, zero value otherwise.
func (o *ExtensionInstanceArray) GetCustomVariables() []CustomVariable {
	if o == nil || IsNil(o.CustomVariables) {
		var ret []CustomVariable
		return ret
	}
	return o.CustomVariables
}

// GetCustomVariablesOk returns a tuple with the CustomVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionInstanceArray) GetCustomVariablesOk() ([]CustomVariable, bool) {
	if o == nil || IsNil(o.CustomVariables) {
		return nil, false
	}
	return o.CustomVariables, true
}

// HasCustomVariables returns a boolean if a field has been set.
func (o *ExtensionInstanceArray) HasCustomVariables() bool {
	if o != nil && !IsNil(o.CustomVariables) {
		return true
	}

	return false
}

// SetCustomVariables gets a reference to the given []CustomVariable and assigns it to the CustomVariables field.
func (o *ExtensionInstanceArray) SetCustomVariables(v []CustomVariable) {
	o.CustomVariables = v
}

// GetDependencies returns the Dependencies field value if set, zero value otherwise.
func (o *ExtensionInstanceArray) GetDependencies() []string {
	if o == nil || IsNil(o.Dependencies) {
		var ret []string
		return ret
	}
	return o.Dependencies
}

// GetDependenciesOk returns a tuple with the Dependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionInstanceArray) GetDependenciesOk() ([]string, bool) {
	if o == nil || IsNil(o.Dependencies) {
		return nil, false
	}
	return o.Dependencies, true
}

// HasDependencies returns a boolean if a field has been set.
func (o *ExtensionInstanceArray) HasDependencies() bool {
	if o != nil && !IsNil(o.Dependencies) {
		return true
	}

	return false
}

// SetDependencies gets a reference to the given []string and assigns it to the Dependencies field.
func (o *ExtensionInstanceArray) SetDependencies(v []string) {
	o.Dependencies = v
}

func (o ExtensionInstanceArray) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtensionInstanceArray) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	toSerialize["instanceCount"] = o.InstanceCount
	toSerialize["serverType"] = o.ServerType
	toSerialize["osTemplate"] = o.OsTemplate
	if !IsNil(o.ConnectedSharedDrives) {
		toSerialize["connectedSharedDrives"] = o.ConnectedSharedDrives
	}
	if !IsNil(o.CustomVariables) {
		toSerialize["customVariables"] = o.CustomVariables
	}
	if !IsNil(o.Dependencies) {
		toSerialize["dependencies"] = o.Dependencies
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExtensionInstanceArray) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"label",
		"instanceCount",
		"serverType",
		"osTemplate",
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

	varExtensionInstanceArray := _ExtensionInstanceArray{}

	err = json.Unmarshal(data, &varExtensionInstanceArray)

	if err != nil {
		return err
	}

	*o = ExtensionInstanceArray(varExtensionInstanceArray)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "instanceCount")
		delete(additionalProperties, "serverType")
		delete(additionalProperties, "osTemplate")
		delete(additionalProperties, "connectedSharedDrives")
		delete(additionalProperties, "customVariables")
		delete(additionalProperties, "dependencies")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExtensionInstanceArray struct {
	value *ExtensionInstanceArray
	isSet bool
}

func (v NullableExtensionInstanceArray) Get() *ExtensionInstanceArray {
	return v.value
}

func (v *NullableExtensionInstanceArray) Set(val *ExtensionInstanceArray) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionInstanceArray) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionInstanceArray) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionInstanceArray(val *ExtensionInstanceArray) *NullableExtensionInstanceArray {
	return &NullableExtensionInstanceArray{value: val, isSet: true}
}

func (v NullableExtensionInstanceArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionInstanceArray) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


