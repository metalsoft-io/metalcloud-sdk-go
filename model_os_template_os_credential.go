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

// checks if the OSTemplateOsCredential type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSTemplateOsCredential{}

// OSTemplateOsCredential struct for OSTemplateOsCredential
type OSTemplateOsCredential struct {
	// The default username for the operating system
	Username string `json:"username"`
	// The default username password type.                     PLAIN means that the password is provided by the user in plain text.                     GENERATED means that the password is generated by the system.                     HASHED means that the password is provided by the user in hashed format.                     NONE means that the password is not needed
	PasswordType string `json:"passwordType"`
	// The default username password.                     Required if the password type is PLAIN or HASHED
	Password *string `json:"password,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OSTemplateOsCredential OSTemplateOsCredential

// NewOSTemplateOsCredential instantiates a new OSTemplateOsCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSTemplateOsCredential(username string, passwordType string) *OSTemplateOsCredential {
	this := OSTemplateOsCredential{}
	this.Username = username
	this.PasswordType = passwordType
	return &this
}

// NewOSTemplateOsCredentialWithDefaults instantiates a new OSTemplateOsCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSTemplateOsCredentialWithDefaults() *OSTemplateOsCredential {
	this := OSTemplateOsCredential{}
	return &this
}

// GetUsername returns the Username field value
func (o *OSTemplateOsCredential) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *OSTemplateOsCredential) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *OSTemplateOsCredential) SetUsername(v string) {
	o.Username = v
}

// GetPasswordType returns the PasswordType field value
func (o *OSTemplateOsCredential) GetPasswordType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PasswordType
}

// GetPasswordTypeOk returns a tuple with the PasswordType field value
// and a boolean to check if the value has been set.
func (o *OSTemplateOsCredential) GetPasswordTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PasswordType, true
}

// SetPasswordType sets field value
func (o *OSTemplateOsCredential) SetPasswordType(v string) {
	o.PasswordType = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *OSTemplateOsCredential) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSTemplateOsCredential) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *OSTemplateOsCredential) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *OSTemplateOsCredential) SetPassword(v string) {
	o.Password = &v
}

func (o OSTemplateOsCredential) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSTemplateOsCredential) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["username"] = o.Username
	toSerialize["passwordType"] = o.PasswordType
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OSTemplateOsCredential) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"username",
		"passwordType",
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

	varOSTemplateOsCredential := _OSTemplateOsCredential{}

	err = json.Unmarshal(data, &varOSTemplateOsCredential)

	if err != nil {
		return err
	}

	*o = OSTemplateOsCredential(varOSTemplateOsCredential)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "username")
		delete(additionalProperties, "passwordType")
		delete(additionalProperties, "password")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOSTemplateOsCredential struct {
	value *OSTemplateOsCredential
	isSet bool
}

func (v NullableOSTemplateOsCredential) Get() *OSTemplateOsCredential {
	return v.value
}

func (v *NullableOSTemplateOsCredential) Set(val *OSTemplateOsCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableOSTemplateOsCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableOSTemplateOsCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSTemplateOsCredential(val *OSTemplateOsCredential) *NullableOSTemplateOsCredential {
	return &NullableOSTemplateOsCredential{value: val, isSet: true}
}

func (v NullableOSTemplateOsCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSTemplateOsCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


