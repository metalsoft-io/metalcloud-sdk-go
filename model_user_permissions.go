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

// checks if the UserPermissions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserPermissions{}

// UserPermissions struct for UserPermissions
type UserPermissions struct {
	// Admin password reveal permissions
	AdminPasswordRevealPermissions *AdminPasswordRevealPermissions `json:"adminPasswordRevealPermissions,omitempty"`
	// Special permissions
	SpecialPermissions *SpecialPermissions `json:"specialPermissions,omitempty"`
	// Role permissions
	RolePermissions []string `json:"rolePermissions"`
	AdditionalProperties map[string]interface{}
}

type _UserPermissions UserPermissions

// NewUserPermissions instantiates a new UserPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPermissions(rolePermissions []string) *UserPermissions {
	this := UserPermissions{}
	this.RolePermissions = rolePermissions
	return &this
}

// NewUserPermissionsWithDefaults instantiates a new UserPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPermissionsWithDefaults() *UserPermissions {
	this := UserPermissions{}
	return &this
}

// GetAdminPasswordRevealPermissions returns the AdminPasswordRevealPermissions field value if set, zero value otherwise.
func (o *UserPermissions) GetAdminPasswordRevealPermissions() AdminPasswordRevealPermissions {
	if o == nil || IsNil(o.AdminPasswordRevealPermissions) {
		var ret AdminPasswordRevealPermissions
		return ret
	}
	return *o.AdminPasswordRevealPermissions
}

// GetAdminPasswordRevealPermissionsOk returns a tuple with the AdminPasswordRevealPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPermissions) GetAdminPasswordRevealPermissionsOk() (*AdminPasswordRevealPermissions, bool) {
	if o == nil || IsNil(o.AdminPasswordRevealPermissions) {
		return nil, false
	}
	return o.AdminPasswordRevealPermissions, true
}

// HasAdminPasswordRevealPermissions returns a boolean if a field has been set.
func (o *UserPermissions) HasAdminPasswordRevealPermissions() bool {
	if o != nil && !IsNil(o.AdminPasswordRevealPermissions) {
		return true
	}

	return false
}

// SetAdminPasswordRevealPermissions gets a reference to the given AdminPasswordRevealPermissions and assigns it to the AdminPasswordRevealPermissions field.
func (o *UserPermissions) SetAdminPasswordRevealPermissions(v AdminPasswordRevealPermissions) {
	o.AdminPasswordRevealPermissions = &v
}

// GetSpecialPermissions returns the SpecialPermissions field value if set, zero value otherwise.
func (o *UserPermissions) GetSpecialPermissions() SpecialPermissions {
	if o == nil || IsNil(o.SpecialPermissions) {
		var ret SpecialPermissions
		return ret
	}
	return *o.SpecialPermissions
}

// GetSpecialPermissionsOk returns a tuple with the SpecialPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPermissions) GetSpecialPermissionsOk() (*SpecialPermissions, bool) {
	if o == nil || IsNil(o.SpecialPermissions) {
		return nil, false
	}
	return o.SpecialPermissions, true
}

// HasSpecialPermissions returns a boolean if a field has been set.
func (o *UserPermissions) HasSpecialPermissions() bool {
	if o != nil && !IsNil(o.SpecialPermissions) {
		return true
	}

	return false
}

// SetSpecialPermissions gets a reference to the given SpecialPermissions and assigns it to the SpecialPermissions field.
func (o *UserPermissions) SetSpecialPermissions(v SpecialPermissions) {
	o.SpecialPermissions = &v
}

// GetRolePermissions returns the RolePermissions field value
func (o *UserPermissions) GetRolePermissions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RolePermissions
}

// GetRolePermissionsOk returns a tuple with the RolePermissions field value
// and a boolean to check if the value has been set.
func (o *UserPermissions) GetRolePermissionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RolePermissions, true
}

// SetRolePermissions sets field value
func (o *UserPermissions) SetRolePermissions(v []string) {
	o.RolePermissions = v
}

func (o UserPermissions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserPermissions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdminPasswordRevealPermissions) {
		toSerialize["adminPasswordRevealPermissions"] = o.AdminPasswordRevealPermissions
	}
	if !IsNil(o.SpecialPermissions) {
		toSerialize["specialPermissions"] = o.SpecialPermissions
	}
	toSerialize["rolePermissions"] = o.RolePermissions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserPermissions) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"rolePermissions",
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

	varUserPermissions := _UserPermissions{}

	err = json.Unmarshal(data, &varUserPermissions)

	if err != nil {
		return err
	}

	*o = UserPermissions(varUserPermissions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "adminPasswordRevealPermissions")
		delete(additionalProperties, "specialPermissions")
		delete(additionalProperties, "rolePermissions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserPermissions struct {
	value *UserPermissions
	isSet bool
}

func (v NullableUserPermissions) Get() *UserPermissions {
	return v.value
}

func (v *NullableUserPermissions) Set(val *UserPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPermissions(val *UserPermissions) *NullableUserPermissions {
	return &NullableUserPermissions{value: val, isSet: true}
}

func (v NullableUserPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


