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

// checks if the UpdateServerIpmiCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateServerIpmiCredentials{}

// UpdateServerIpmiCredentials struct for UpdateServerIpmiCredentials
type UpdateServerIpmiCredentials struct {
	// The host of the server.
	Host string `json:"host"`
	// The username of the server.
	Username string `json:"username"`
	// The password of the server.
	Password string `json:"password"`
	// Flag to indicate if the credentials should be updated on the server as well.
	UpdateOnServer bool `json:"updateOnServer"`
	AdditionalProperties map[string]interface{}
}

type _UpdateServerIpmiCredentials UpdateServerIpmiCredentials

// NewUpdateServerIpmiCredentials instantiates a new UpdateServerIpmiCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateServerIpmiCredentials(host string, username string, password string, updateOnServer bool) *UpdateServerIpmiCredentials {
	this := UpdateServerIpmiCredentials{}
	this.Host = host
	this.Username = username
	this.Password = password
	this.UpdateOnServer = updateOnServer
	return &this
}

// NewUpdateServerIpmiCredentialsWithDefaults instantiates a new UpdateServerIpmiCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateServerIpmiCredentialsWithDefaults() *UpdateServerIpmiCredentials {
	this := UpdateServerIpmiCredentials{}
	return &this
}

// GetHost returns the Host field value
func (o *UpdateServerIpmiCredentials) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *UpdateServerIpmiCredentials) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *UpdateServerIpmiCredentials) SetHost(v string) {
	o.Host = v
}

// GetUsername returns the Username field value
func (o *UpdateServerIpmiCredentials) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *UpdateServerIpmiCredentials) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *UpdateServerIpmiCredentials) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *UpdateServerIpmiCredentials) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *UpdateServerIpmiCredentials) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *UpdateServerIpmiCredentials) SetPassword(v string) {
	o.Password = v
}

// GetUpdateOnServer returns the UpdateOnServer field value
func (o *UpdateServerIpmiCredentials) GetUpdateOnServer() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UpdateOnServer
}

// GetUpdateOnServerOk returns a tuple with the UpdateOnServer field value
// and a boolean to check if the value has been set.
func (o *UpdateServerIpmiCredentials) GetUpdateOnServerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdateOnServer, true
}

// SetUpdateOnServer sets field value
func (o *UpdateServerIpmiCredentials) SetUpdateOnServer(v bool) {
	o.UpdateOnServer = v
}

func (o UpdateServerIpmiCredentials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateServerIpmiCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["host"] = o.Host
	toSerialize["username"] = o.Username
	toSerialize["password"] = o.Password
	toSerialize["updateOnServer"] = o.UpdateOnServer

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateServerIpmiCredentials) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"host",
		"username",
		"password",
		"updateOnServer",
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

	varUpdateServerIpmiCredentials := _UpdateServerIpmiCredentials{}

	err = json.Unmarshal(data, &varUpdateServerIpmiCredentials)

	if err != nil {
		return err
	}

	*o = UpdateServerIpmiCredentials(varUpdateServerIpmiCredentials)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "host")
		delete(additionalProperties, "username")
		delete(additionalProperties, "password")
		delete(additionalProperties, "updateOnServer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateServerIpmiCredentials struct {
	value *UpdateServerIpmiCredentials
	isSet bool
}

func (v NullableUpdateServerIpmiCredentials) Get() *UpdateServerIpmiCredentials {
	return v.value
}

func (v *NullableUpdateServerIpmiCredentials) Set(val *UpdateServerIpmiCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateServerIpmiCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateServerIpmiCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateServerIpmiCredentials(val *UpdateServerIpmiCredentials) *NullableUpdateServerIpmiCredentials {
	return &NullableUpdateServerIpmiCredentials{value: val, isSet: true}
}

func (v NullableUpdateServerIpmiCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateServerIpmiCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


