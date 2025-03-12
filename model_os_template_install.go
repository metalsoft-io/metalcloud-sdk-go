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

// checks if the OSTemplateInstall type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSTemplateInstall{}

// OSTemplateInstall struct for OSTemplateInstall
type OSTemplateInstall struct {
	// The OS template installation method                     OOB - Out of Band (virtual media)                     ONIE - Open Network Install Environment                     DS_LXD - Cloud-init with datasource LXD
	Method string `json:"method"`
	// The OS template installation drive type
	DriveType string `json:"driveType"`
	// The OS template installation ready method,                     The \"ready method\" is used to determine when the OS installation is complete.
	ReadyMethod string `json:"readyMethod"`
	// Used for selecting the OS template during network device ZTP
	OnieStrings []string `json:"onieStrings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OSTemplateInstall OSTemplateInstall

// NewOSTemplateInstall instantiates a new OSTemplateInstall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSTemplateInstall(method string, driveType string, readyMethod string) *OSTemplateInstall {
	this := OSTemplateInstall{}
	this.Method = method
	this.DriveType = driveType
	this.ReadyMethod = readyMethod
	return &this
}

// NewOSTemplateInstallWithDefaults instantiates a new OSTemplateInstall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSTemplateInstallWithDefaults() *OSTemplateInstall {
	this := OSTemplateInstall{}
	return &this
}

// GetMethod returns the Method field value
func (o *OSTemplateInstall) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *OSTemplateInstall) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *OSTemplateInstall) SetMethod(v string) {
	o.Method = v
}

// GetDriveType returns the DriveType field value
func (o *OSTemplateInstall) GetDriveType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DriveType
}

// GetDriveTypeOk returns a tuple with the DriveType field value
// and a boolean to check if the value has been set.
func (o *OSTemplateInstall) GetDriveTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DriveType, true
}

// SetDriveType sets field value
func (o *OSTemplateInstall) SetDriveType(v string) {
	o.DriveType = v
}

// GetReadyMethod returns the ReadyMethod field value
func (o *OSTemplateInstall) GetReadyMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReadyMethod
}

// GetReadyMethodOk returns a tuple with the ReadyMethod field value
// and a boolean to check if the value has been set.
func (o *OSTemplateInstall) GetReadyMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReadyMethod, true
}

// SetReadyMethod sets field value
func (o *OSTemplateInstall) SetReadyMethod(v string) {
	o.ReadyMethod = v
}

// GetOnieStrings returns the OnieStrings field value if set, zero value otherwise.
func (o *OSTemplateInstall) GetOnieStrings() []string {
	if o == nil || IsNil(o.OnieStrings) {
		var ret []string
		return ret
	}
	return o.OnieStrings
}

// GetOnieStringsOk returns a tuple with the OnieStrings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSTemplateInstall) GetOnieStringsOk() ([]string, bool) {
	if o == nil || IsNil(o.OnieStrings) {
		return nil, false
	}
	return o.OnieStrings, true
}

// HasOnieStrings returns a boolean if a field has been set.
func (o *OSTemplateInstall) HasOnieStrings() bool {
	if o != nil && !IsNil(o.OnieStrings) {
		return true
	}

	return false
}

// SetOnieStrings gets a reference to the given []string and assigns it to the OnieStrings field.
func (o *OSTemplateInstall) SetOnieStrings(v []string) {
	o.OnieStrings = v
}

func (o OSTemplateInstall) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSTemplateInstall) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["method"] = o.Method
	toSerialize["driveType"] = o.DriveType
	toSerialize["readyMethod"] = o.ReadyMethod
	if !IsNil(o.OnieStrings) {
		toSerialize["onieStrings"] = o.OnieStrings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OSTemplateInstall) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"method",
		"driveType",
		"readyMethod",
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

	varOSTemplateInstall := _OSTemplateInstall{}

	err = json.Unmarshal(data, &varOSTemplateInstall)

	if err != nil {
		return err
	}

	*o = OSTemplateInstall(varOSTemplateInstall)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "method")
		delete(additionalProperties, "driveType")
		delete(additionalProperties, "readyMethod")
		delete(additionalProperties, "onieStrings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOSTemplateInstall struct {
	value *OSTemplateInstall
	isSet bool
}

func (v NullableOSTemplateInstall) Get() *OSTemplateInstall {
	return v.value
}

func (v *NullableOSTemplateInstall) Set(val *OSTemplateInstall) {
	v.value = val
	v.isSet = true
}

func (v NullableOSTemplateInstall) IsSet() bool {
	return v.isSet
}

func (v *NullableOSTemplateInstall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSTemplateInstall(val *OSTemplateInstall) *NullableOSTemplateInstall {
	return &NullableOSTemplateInstall{value: val, isSet: true}
}

func (v NullableOSTemplateInstall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSTemplateInstall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


