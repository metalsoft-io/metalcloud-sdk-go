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

// checks if the AuthenticationRequestPropertiesSamlDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticationRequestPropertiesSamlDto{}

// AuthenticationRequestPropertiesSamlDto struct for AuthenticationRequestPropertiesSamlDto
type AuthenticationRequestPropertiesSamlDto struct {
	// The one-time password.
	OneTimePassword *string `json:"oneTimePassword,omitempty"`
	// Whether to remember the login
	RememberLogin *bool `json:"rememberLogin,omitempty"`
	// Whether to test the credentials
	TestCredentials *bool `json:"testCredentials,omitempty"`
	// The data for the SAML authentication request.
	Data map[string]interface{} `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticationRequestPropertiesSamlDto AuthenticationRequestPropertiesSamlDto

// NewAuthenticationRequestPropertiesSamlDto instantiates a new AuthenticationRequestPropertiesSamlDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationRequestPropertiesSamlDto(data map[string]interface{}) *AuthenticationRequestPropertiesSamlDto {
	this := AuthenticationRequestPropertiesSamlDto{}
	var rememberLogin bool = true
	this.RememberLogin = &rememberLogin
	var testCredentials bool = false
	this.TestCredentials = &testCredentials
	this.Data = data
	return &this
}

// NewAuthenticationRequestPropertiesSamlDtoWithDefaults instantiates a new AuthenticationRequestPropertiesSamlDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationRequestPropertiesSamlDtoWithDefaults() *AuthenticationRequestPropertiesSamlDto {
	this := AuthenticationRequestPropertiesSamlDto{}
	var rememberLogin bool = true
	this.RememberLogin = &rememberLogin
	var testCredentials bool = false
	this.TestCredentials = &testCredentials
	return &this
}

// GetOneTimePassword returns the OneTimePassword field value if set, zero value otherwise.
func (o *AuthenticationRequestPropertiesSamlDto) GetOneTimePassword() string {
	if o == nil || IsNil(o.OneTimePassword) {
		var ret string
		return ret
	}
	return *o.OneTimePassword
}

// GetOneTimePasswordOk returns a tuple with the OneTimePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationRequestPropertiesSamlDto) GetOneTimePasswordOk() (*string, bool) {
	if o == nil || IsNil(o.OneTimePassword) {
		return nil, false
	}
	return o.OneTimePassword, true
}

// HasOneTimePassword returns a boolean if a field has been set.
func (o *AuthenticationRequestPropertiesSamlDto) HasOneTimePassword() bool {
	if o != nil && !IsNil(o.OneTimePassword) {
		return true
	}

	return false
}

// SetOneTimePassword gets a reference to the given string and assigns it to the OneTimePassword field.
func (o *AuthenticationRequestPropertiesSamlDto) SetOneTimePassword(v string) {
	o.OneTimePassword = &v
}

// GetRememberLogin returns the RememberLogin field value if set, zero value otherwise.
func (o *AuthenticationRequestPropertiesSamlDto) GetRememberLogin() bool {
	if o == nil || IsNil(o.RememberLogin) {
		var ret bool
		return ret
	}
	return *o.RememberLogin
}

// GetRememberLoginOk returns a tuple with the RememberLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationRequestPropertiesSamlDto) GetRememberLoginOk() (*bool, bool) {
	if o == nil || IsNil(o.RememberLogin) {
		return nil, false
	}
	return o.RememberLogin, true
}

// HasRememberLogin returns a boolean if a field has been set.
func (o *AuthenticationRequestPropertiesSamlDto) HasRememberLogin() bool {
	if o != nil && !IsNil(o.RememberLogin) {
		return true
	}

	return false
}

// SetRememberLogin gets a reference to the given bool and assigns it to the RememberLogin field.
func (o *AuthenticationRequestPropertiesSamlDto) SetRememberLogin(v bool) {
	o.RememberLogin = &v
}

// GetTestCredentials returns the TestCredentials field value if set, zero value otherwise.
func (o *AuthenticationRequestPropertiesSamlDto) GetTestCredentials() bool {
	if o == nil || IsNil(o.TestCredentials) {
		var ret bool
		return ret
	}
	return *o.TestCredentials
}

// GetTestCredentialsOk returns a tuple with the TestCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationRequestPropertiesSamlDto) GetTestCredentialsOk() (*bool, bool) {
	if o == nil || IsNil(o.TestCredentials) {
		return nil, false
	}
	return o.TestCredentials, true
}

// HasTestCredentials returns a boolean if a field has been set.
func (o *AuthenticationRequestPropertiesSamlDto) HasTestCredentials() bool {
	if o != nil && !IsNil(o.TestCredentials) {
		return true
	}

	return false
}

// SetTestCredentials gets a reference to the given bool and assigns it to the TestCredentials field.
func (o *AuthenticationRequestPropertiesSamlDto) SetTestCredentials(v bool) {
	o.TestCredentials = &v
}

// GetData returns the Data field value
func (o *AuthenticationRequestPropertiesSamlDto) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AuthenticationRequestPropertiesSamlDto) GetDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *AuthenticationRequestPropertiesSamlDto) SetData(v map[string]interface{}) {
	o.Data = v
}

func (o AuthenticationRequestPropertiesSamlDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationRequestPropertiesSamlDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OneTimePassword) {
		toSerialize["oneTimePassword"] = o.OneTimePassword
	}
	if !IsNil(o.RememberLogin) {
		toSerialize["rememberLogin"] = o.RememberLogin
	}
	if !IsNil(o.TestCredentials) {
		toSerialize["testCredentials"] = o.TestCredentials
	}
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticationRequestPropertiesSamlDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varAuthenticationRequestPropertiesSamlDto := _AuthenticationRequestPropertiesSamlDto{}

	err = json.Unmarshal(data, &varAuthenticationRequestPropertiesSamlDto)

	if err != nil {
		return err
	}

	*o = AuthenticationRequestPropertiesSamlDto(varAuthenticationRequestPropertiesSamlDto)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "oneTimePassword")
		delete(additionalProperties, "rememberLogin")
		delete(additionalProperties, "testCredentials")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthenticationRequestPropertiesSamlDto struct {
	value *AuthenticationRequestPropertiesSamlDto
	isSet bool
}

func (v NullableAuthenticationRequestPropertiesSamlDto) Get() *AuthenticationRequestPropertiesSamlDto {
	return v.value
}

func (v *NullableAuthenticationRequestPropertiesSamlDto) Set(val *AuthenticationRequestPropertiesSamlDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationRequestPropertiesSamlDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationRequestPropertiesSamlDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationRequestPropertiesSamlDto(val *AuthenticationRequestPropertiesSamlDto) *NullableAuthenticationRequestPropertiesSamlDto {
	return &NullableAuthenticationRequestPropertiesSamlDto{value: val, isSet: true}
}

func (v NullableAuthenticationRequestPropertiesSamlDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationRequestPropertiesSamlDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


