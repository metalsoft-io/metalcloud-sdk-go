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

// checks if the ServerCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerCredentials{}

// ServerCredentials struct for ServerCredentials
type ServerCredentials struct {
	// The username of the server.
	Username *string `json:"username,omitempty"`
	// The password of the server.
	Password *string `json:"password,omitempty"`
	// The VNC password of the server.
	VncPassword NullableString `json:"vncPassword,omitempty"`
	// The SNMP password of the server.
	SnmpPassword NullableString `json:"snmpPassword,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServerCredentials ServerCredentials

// NewServerCredentials instantiates a new ServerCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerCredentials() *ServerCredentials {
	this := ServerCredentials{}
	return &this
}

// NewServerCredentialsWithDefaults instantiates a new ServerCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerCredentialsWithDefaults() *ServerCredentials {
	this := ServerCredentials{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ServerCredentials) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerCredentials) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ServerCredentials) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ServerCredentials) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ServerCredentials) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerCredentials) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ServerCredentials) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ServerCredentials) SetPassword(v string) {
	o.Password = &v
}

// GetVncPassword returns the VncPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerCredentials) GetVncPassword() string {
	if o == nil || IsNil(o.VncPassword.Get()) {
		var ret string
		return ret
	}
	return *o.VncPassword.Get()
}

// GetVncPasswordOk returns a tuple with the VncPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerCredentials) GetVncPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VncPassword.Get(), o.VncPassword.IsSet()
}

// HasVncPassword returns a boolean if a field has been set.
func (o *ServerCredentials) HasVncPassword() bool {
	if o != nil && o.VncPassword.IsSet() {
		return true
	}

	return false
}

// SetVncPassword gets a reference to the given NullableString and assigns it to the VncPassword field.
func (o *ServerCredentials) SetVncPassword(v string) {
	o.VncPassword.Set(&v)
}
// SetVncPasswordNil sets the value for VncPassword to be an explicit nil
func (o *ServerCredentials) SetVncPasswordNil() {
	o.VncPassword.Set(nil)
}

// UnsetVncPassword ensures that no value is present for VncPassword, not even an explicit nil
func (o *ServerCredentials) UnsetVncPassword() {
	o.VncPassword.Unset()
}

// GetSnmpPassword returns the SnmpPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerCredentials) GetSnmpPassword() string {
	if o == nil || IsNil(o.SnmpPassword.Get()) {
		var ret string
		return ret
	}
	return *o.SnmpPassword.Get()
}

// GetSnmpPasswordOk returns a tuple with the SnmpPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerCredentials) GetSnmpPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SnmpPassword.Get(), o.SnmpPassword.IsSet()
}

// HasSnmpPassword returns a boolean if a field has been set.
func (o *ServerCredentials) HasSnmpPassword() bool {
	if o != nil && o.SnmpPassword.IsSet() {
		return true
	}

	return false
}

// SetSnmpPassword gets a reference to the given NullableString and assigns it to the SnmpPassword field.
func (o *ServerCredentials) SetSnmpPassword(v string) {
	o.SnmpPassword.Set(&v)
}
// SetSnmpPasswordNil sets the value for SnmpPassword to be an explicit nil
func (o *ServerCredentials) SetSnmpPasswordNil() {
	o.SnmpPassword.Set(nil)
}

// UnsetSnmpPassword ensures that no value is present for SnmpPassword, not even an explicit nil
func (o *ServerCredentials) UnsetSnmpPassword() {
	o.SnmpPassword.Unset()
}

func (o ServerCredentials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if o.VncPassword.IsSet() {
		toSerialize["vncPassword"] = o.VncPassword.Get()
	}
	if o.SnmpPassword.IsSet() {
		toSerialize["snmpPassword"] = o.SnmpPassword.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServerCredentials) UnmarshalJSON(data []byte) (err error) {
	varServerCredentials := _ServerCredentials{}

	err = json.Unmarshal(data, &varServerCredentials)

	if err != nil {
		return err
	}

	*o = ServerCredentials(varServerCredentials)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "username")
		delete(additionalProperties, "password")
		delete(additionalProperties, "vncPassword")
		delete(additionalProperties, "snmpPassword")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServerCredentials struct {
	value *ServerCredentials
	isSet bool
}

func (v NullableServerCredentials) Get() *ServerCredentials {
	return v.value
}

func (v *NullableServerCredentials) Set(val *ServerCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableServerCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableServerCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerCredentials(val *ServerCredentials) *NullableServerCredentials {
	return &NullableServerCredentials{value: val, isSet: true}
}

func (v NullableServerCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


