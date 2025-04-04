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
)

// checks if the ExtensionExecutionInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtensionExecutionInfo{}

// ExtensionExecutionInfo struct for ExtensionExecutionInfo
type ExtensionExecutionInfo struct {
	// The path towards the ansible bundle
	AnsibleBundlePath *string `json:"ansibleBundlePath,omitempty"`
	// The name of the ansible bundle file
	AnsibleBundleFileName *string `json:"ansibleBundleFileName,omitempty"`
	// The ansible bundle token
	AnsibleBundleToken *string `json:"ansibleBundleToken,omitempty"`
	// The success message of the ansible bundle
	SuccessMessage *string `json:"successMessage,omitempty"`
	// Success timestamp
	SuccessMessageTimestamp *string `json:"successMessageTimestamp,omitempty"`
	// The error message of the ansible bundle
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Error timestamp
	ErrorMessageTimestamp *string `json:"errorMessageTimestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExtensionExecutionInfo ExtensionExecutionInfo

// NewExtensionExecutionInfo instantiates a new ExtensionExecutionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtensionExecutionInfo() *ExtensionExecutionInfo {
	this := ExtensionExecutionInfo{}
	return &this
}

// NewExtensionExecutionInfoWithDefaults instantiates a new ExtensionExecutionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionExecutionInfoWithDefaults() *ExtensionExecutionInfo {
	this := ExtensionExecutionInfo{}
	return &this
}

// GetAnsibleBundlePath returns the AnsibleBundlePath field value if set, zero value otherwise.
func (o *ExtensionExecutionInfo) GetAnsibleBundlePath() string {
	if o == nil || IsNil(o.AnsibleBundlePath) {
		var ret string
		return ret
	}
	return *o.AnsibleBundlePath
}

// GetAnsibleBundlePathOk returns a tuple with the AnsibleBundlePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionExecutionInfo) GetAnsibleBundlePathOk() (*string, bool) {
	if o == nil || IsNil(o.AnsibleBundlePath) {
		return nil, false
	}
	return o.AnsibleBundlePath, true
}

// HasAnsibleBundlePath returns a boolean if a field has been set.
func (o *ExtensionExecutionInfo) HasAnsibleBundlePath() bool {
	if o != nil && !IsNil(o.AnsibleBundlePath) {
		return true
	}

	return false
}

// SetAnsibleBundlePath gets a reference to the given string and assigns it to the AnsibleBundlePath field.
func (o *ExtensionExecutionInfo) SetAnsibleBundlePath(v string) {
	o.AnsibleBundlePath = &v
}

// GetAnsibleBundleFileName returns the AnsibleBundleFileName field value if set, zero value otherwise.
func (o *ExtensionExecutionInfo) GetAnsibleBundleFileName() string {
	if o == nil || IsNil(o.AnsibleBundleFileName) {
		var ret string
		return ret
	}
	return *o.AnsibleBundleFileName
}

// GetAnsibleBundleFileNameOk returns a tuple with the AnsibleBundleFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionExecutionInfo) GetAnsibleBundleFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.AnsibleBundleFileName) {
		return nil, false
	}
	return o.AnsibleBundleFileName, true
}

// HasAnsibleBundleFileName returns a boolean if a field has been set.
func (o *ExtensionExecutionInfo) HasAnsibleBundleFileName() bool {
	if o != nil && !IsNil(o.AnsibleBundleFileName) {
		return true
	}

	return false
}

// SetAnsibleBundleFileName gets a reference to the given string and assigns it to the AnsibleBundleFileName field.
func (o *ExtensionExecutionInfo) SetAnsibleBundleFileName(v string) {
	o.AnsibleBundleFileName = &v
}

// GetAnsibleBundleToken returns the AnsibleBundleToken field value if set, zero value otherwise.
func (o *ExtensionExecutionInfo) GetAnsibleBundleToken() string {
	if o == nil || IsNil(o.AnsibleBundleToken) {
		var ret string
		return ret
	}
	return *o.AnsibleBundleToken
}

// GetAnsibleBundleTokenOk returns a tuple with the AnsibleBundleToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionExecutionInfo) GetAnsibleBundleTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AnsibleBundleToken) {
		return nil, false
	}
	return o.AnsibleBundleToken, true
}

// HasAnsibleBundleToken returns a boolean if a field has been set.
func (o *ExtensionExecutionInfo) HasAnsibleBundleToken() bool {
	if o != nil && !IsNil(o.AnsibleBundleToken) {
		return true
	}

	return false
}

// SetAnsibleBundleToken gets a reference to the given string and assigns it to the AnsibleBundleToken field.
func (o *ExtensionExecutionInfo) SetAnsibleBundleToken(v string) {
	o.AnsibleBundleToken = &v
}

// GetSuccessMessage returns the SuccessMessage field value if set, zero value otherwise.
func (o *ExtensionExecutionInfo) GetSuccessMessage() string {
	if o == nil || IsNil(o.SuccessMessage) {
		var ret string
		return ret
	}
	return *o.SuccessMessage
}

// GetSuccessMessageOk returns a tuple with the SuccessMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionExecutionInfo) GetSuccessMessageOk() (*string, bool) {
	if o == nil || IsNil(o.SuccessMessage) {
		return nil, false
	}
	return o.SuccessMessage, true
}

// HasSuccessMessage returns a boolean if a field has been set.
func (o *ExtensionExecutionInfo) HasSuccessMessage() bool {
	if o != nil && !IsNil(o.SuccessMessage) {
		return true
	}

	return false
}

// SetSuccessMessage gets a reference to the given string and assigns it to the SuccessMessage field.
func (o *ExtensionExecutionInfo) SetSuccessMessage(v string) {
	o.SuccessMessage = &v
}

// GetSuccessMessageTimestamp returns the SuccessMessageTimestamp field value if set, zero value otherwise.
func (o *ExtensionExecutionInfo) GetSuccessMessageTimestamp() string {
	if o == nil || IsNil(o.SuccessMessageTimestamp) {
		var ret string
		return ret
	}
	return *o.SuccessMessageTimestamp
}

// GetSuccessMessageTimestampOk returns a tuple with the SuccessMessageTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionExecutionInfo) GetSuccessMessageTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.SuccessMessageTimestamp) {
		return nil, false
	}
	return o.SuccessMessageTimestamp, true
}

// HasSuccessMessageTimestamp returns a boolean if a field has been set.
func (o *ExtensionExecutionInfo) HasSuccessMessageTimestamp() bool {
	if o != nil && !IsNil(o.SuccessMessageTimestamp) {
		return true
	}

	return false
}

// SetSuccessMessageTimestamp gets a reference to the given string and assigns it to the SuccessMessageTimestamp field.
func (o *ExtensionExecutionInfo) SetSuccessMessageTimestamp(v string) {
	o.SuccessMessageTimestamp = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *ExtensionExecutionInfo) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionExecutionInfo) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ExtensionExecutionInfo) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *ExtensionExecutionInfo) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetErrorMessageTimestamp returns the ErrorMessageTimestamp field value if set, zero value otherwise.
func (o *ExtensionExecutionInfo) GetErrorMessageTimestamp() string {
	if o == nil || IsNil(o.ErrorMessageTimestamp) {
		var ret string
		return ret
	}
	return *o.ErrorMessageTimestamp
}

// GetErrorMessageTimestampOk returns a tuple with the ErrorMessageTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionExecutionInfo) GetErrorMessageTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessageTimestamp) {
		return nil, false
	}
	return o.ErrorMessageTimestamp, true
}

// HasErrorMessageTimestamp returns a boolean if a field has been set.
func (o *ExtensionExecutionInfo) HasErrorMessageTimestamp() bool {
	if o != nil && !IsNil(o.ErrorMessageTimestamp) {
		return true
	}

	return false
}

// SetErrorMessageTimestamp gets a reference to the given string and assigns it to the ErrorMessageTimestamp field.
func (o *ExtensionExecutionInfo) SetErrorMessageTimestamp(v string) {
	o.ErrorMessageTimestamp = &v
}

func (o ExtensionExecutionInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtensionExecutionInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AnsibleBundlePath) {
		toSerialize["ansibleBundlePath"] = o.AnsibleBundlePath
	}
	if !IsNil(o.AnsibleBundleFileName) {
		toSerialize["ansibleBundleFileName"] = o.AnsibleBundleFileName
	}
	if !IsNil(o.AnsibleBundleToken) {
		toSerialize["ansibleBundleToken"] = o.AnsibleBundleToken
	}
	if !IsNil(o.SuccessMessage) {
		toSerialize["successMessage"] = o.SuccessMessage
	}
	if !IsNil(o.SuccessMessageTimestamp) {
		toSerialize["successMessageTimestamp"] = o.SuccessMessageTimestamp
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if !IsNil(o.ErrorMessageTimestamp) {
		toSerialize["errorMessageTimestamp"] = o.ErrorMessageTimestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExtensionExecutionInfo) UnmarshalJSON(data []byte) (err error) {
	varExtensionExecutionInfo := _ExtensionExecutionInfo{}

	err = json.Unmarshal(data, &varExtensionExecutionInfo)

	if err != nil {
		return err
	}

	*o = ExtensionExecutionInfo(varExtensionExecutionInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ansibleBundlePath")
		delete(additionalProperties, "ansibleBundleFileName")
		delete(additionalProperties, "ansibleBundleToken")
		delete(additionalProperties, "successMessage")
		delete(additionalProperties, "successMessageTimestamp")
		delete(additionalProperties, "errorMessage")
		delete(additionalProperties, "errorMessageTimestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExtensionExecutionInfo struct {
	value *ExtensionExecutionInfo
	isSet bool
}

func (v NullableExtensionExecutionInfo) Get() *ExtensionExecutionInfo {
	return v.value
}

func (v *NullableExtensionExecutionInfo) Set(val *ExtensionExecutionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionExecutionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionExecutionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionExecutionInfo(val *ExtensionExecutionInfo) *NullableExtensionExecutionInfo {
	return &NullableExtensionExecutionInfo{value: val, isSet: true}
}

func (v NullableExtensionExecutionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionExecutionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


