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

// checks if the ChangeUserEmail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeUserEmail{}

// ChangeUserEmail struct for ChangeUserEmail
type ChangeUserEmail struct {
	// The new email address of the user
	Email string `json:"email"`
	// The redirect URL after email verification
	RedirectUrl *string `json:"redirectUrl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChangeUserEmail ChangeUserEmail

// NewChangeUserEmail instantiates a new ChangeUserEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeUserEmail(email string) *ChangeUserEmail {
	this := ChangeUserEmail{}
	this.Email = email
	return &this
}

// NewChangeUserEmailWithDefaults instantiates a new ChangeUserEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeUserEmailWithDefaults() *ChangeUserEmail {
	this := ChangeUserEmail{}
	return &this
}

// GetEmail returns the Email field value
func (o *ChangeUserEmail) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ChangeUserEmail) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ChangeUserEmail) SetEmail(v string) {
	o.Email = v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *ChangeUserEmail) GetRedirectUrl() string {
	if o == nil || IsNil(o.RedirectUrl) {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeUserEmail) GetRedirectUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RedirectUrl) {
		return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *ChangeUserEmail) HasRedirectUrl() bool {
	if o != nil && !IsNil(o.RedirectUrl) {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *ChangeUserEmail) SetRedirectUrl(v string) {
	o.RedirectUrl = &v
}

func (o ChangeUserEmail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeUserEmail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	if !IsNil(o.RedirectUrl) {
		toSerialize["redirectUrl"] = o.RedirectUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChangeUserEmail) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
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

	varChangeUserEmail := _ChangeUserEmail{}

	err = json.Unmarshal(data, &varChangeUserEmail)

	if err != nil {
		return err
	}

	*o = ChangeUserEmail(varChangeUserEmail)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "redirectUrl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChangeUserEmail struct {
	value *ChangeUserEmail
	isSet bool
}

func (v NullableChangeUserEmail) Get() *ChangeUserEmail {
	return v.value
}

func (v *NullableChangeUserEmail) Set(val *ChangeUserEmail) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeUserEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeUserEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeUserEmail(val *ChangeUserEmail) *NullableChangeUserEmail {
	return &NullableChangeUserEmail{value: val, isSet: true}
}

func (v NullableChangeUserEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeUserEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


