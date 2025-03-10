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

// checks if the UserDelegate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserDelegate{}

// UserDelegate struct for UserDelegate
type UserDelegate struct {
	// The user delegate e-mail
	Email *string `json:"email,omitempty"`
	// Whether to create the delegate if it does not exist. Can only be used when creating a new delegate.
	CreateIfNotExists *bool `json:"createIfNotExists,omitempty"`
	// Whether to remove the delegate. Cannot be used when creating a new delegate.
	RemoveDelegate *bool `json:"removeDelegate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserDelegate UserDelegate

// NewUserDelegate instantiates a new UserDelegate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDelegate() *UserDelegate {
	this := UserDelegate{}
	var createIfNotExists bool = false
	this.CreateIfNotExists = &createIfNotExists
	var removeDelegate bool = false
	this.RemoveDelegate = &removeDelegate
	return &this
}

// NewUserDelegateWithDefaults instantiates a new UserDelegate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDelegateWithDefaults() *UserDelegate {
	this := UserDelegate{}
	var createIfNotExists bool = false
	this.CreateIfNotExists = &createIfNotExists
	var removeDelegate bool = false
	this.RemoveDelegate = &removeDelegate
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserDelegate) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDelegate) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserDelegate) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserDelegate) SetEmail(v string) {
	o.Email = &v
}

// GetCreateIfNotExists returns the CreateIfNotExists field value if set, zero value otherwise.
func (o *UserDelegate) GetCreateIfNotExists() bool {
	if o == nil || IsNil(o.CreateIfNotExists) {
		var ret bool
		return ret
	}
	return *o.CreateIfNotExists
}

// GetCreateIfNotExistsOk returns a tuple with the CreateIfNotExists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDelegate) GetCreateIfNotExistsOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateIfNotExists) {
		return nil, false
	}
	return o.CreateIfNotExists, true
}

// HasCreateIfNotExists returns a boolean if a field has been set.
func (o *UserDelegate) HasCreateIfNotExists() bool {
	if o != nil && !IsNil(o.CreateIfNotExists) {
		return true
	}

	return false
}

// SetCreateIfNotExists gets a reference to the given bool and assigns it to the CreateIfNotExists field.
func (o *UserDelegate) SetCreateIfNotExists(v bool) {
	o.CreateIfNotExists = &v
}

// GetRemoveDelegate returns the RemoveDelegate field value if set, zero value otherwise.
func (o *UserDelegate) GetRemoveDelegate() bool {
	if o == nil || IsNil(o.RemoveDelegate) {
		var ret bool
		return ret
	}
	return *o.RemoveDelegate
}

// GetRemoveDelegateOk returns a tuple with the RemoveDelegate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDelegate) GetRemoveDelegateOk() (*bool, bool) {
	if o == nil || IsNil(o.RemoveDelegate) {
		return nil, false
	}
	return o.RemoveDelegate, true
}

// HasRemoveDelegate returns a boolean if a field has been set.
func (o *UserDelegate) HasRemoveDelegate() bool {
	if o != nil && !IsNil(o.RemoveDelegate) {
		return true
	}

	return false
}

// SetRemoveDelegate gets a reference to the given bool and assigns it to the RemoveDelegate field.
func (o *UserDelegate) SetRemoveDelegate(v bool) {
	o.RemoveDelegate = &v
}

func (o UserDelegate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserDelegate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.CreateIfNotExists) {
		toSerialize["createIfNotExists"] = o.CreateIfNotExists
	}
	if !IsNil(o.RemoveDelegate) {
		toSerialize["removeDelegate"] = o.RemoveDelegate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserDelegate) UnmarshalJSON(data []byte) (err error) {
	varUserDelegate := _UserDelegate{}

	err = json.Unmarshal(data, &varUserDelegate)

	if err != nil {
		return err
	}

	*o = UserDelegate(varUserDelegate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "createIfNotExists")
		delete(additionalProperties, "removeDelegate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserDelegate struct {
	value *UserDelegate
	isSet bool
}

func (v NullableUserDelegate) Get() *UserDelegate {
	return v.value
}

func (v *NullableUserDelegate) Set(val *UserDelegate) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDelegate) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDelegate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDelegate(val *UserDelegate) *NullableUserDelegate {
	return &NullableUserDelegate{value: val, isSet: true}
}

func (v NullableUserDelegate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDelegate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


