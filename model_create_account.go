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

// checks if the CreateAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAccount{}

// CreateAccount struct for CreateAccount
type CreateAccount struct {
	// The ID of the parent account
	ParentAccountId *float32 `json:"parentAccountId,omitempty"`
	// The name of the account
	Name string `json:"name"`
	// The code of the account
	Code *string `json:"code,omitempty"`
	// The fiscal number of the account
	FiscalNumber *string `json:"fiscalNumber,omitempty"`
	Address *AccountAddress `json:"address,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateAccount CreateAccount

// NewCreateAccount instantiates a new CreateAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAccount(name string) *CreateAccount {
	this := CreateAccount{}
	this.Name = name
	return &this
}

// NewCreateAccountWithDefaults instantiates a new CreateAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAccountWithDefaults() *CreateAccount {
	this := CreateAccount{}
	return &this
}

// GetParentAccountId returns the ParentAccountId field value if set, zero value otherwise.
func (o *CreateAccount) GetParentAccountId() float32 {
	if o == nil || IsNil(o.ParentAccountId) {
		var ret float32
		return ret
	}
	return *o.ParentAccountId
}

// GetParentAccountIdOk returns a tuple with the ParentAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccount) GetParentAccountIdOk() (*float32, bool) {
	if o == nil || IsNil(o.ParentAccountId) {
		return nil, false
	}
	return o.ParentAccountId, true
}

// HasParentAccountId returns a boolean if a field has been set.
func (o *CreateAccount) HasParentAccountId() bool {
	if o != nil && !IsNil(o.ParentAccountId) {
		return true
	}

	return false
}

// SetParentAccountId gets a reference to the given float32 and assigns it to the ParentAccountId field.
func (o *CreateAccount) SetParentAccountId(v float32) {
	o.ParentAccountId = &v
}

// GetName returns the Name field value
func (o *CreateAccount) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateAccount) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateAccount) SetName(v string) {
	o.Name = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CreateAccount) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccount) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CreateAccount) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CreateAccount) SetCode(v string) {
	o.Code = &v
}

// GetFiscalNumber returns the FiscalNumber field value if set, zero value otherwise.
func (o *CreateAccount) GetFiscalNumber() string {
	if o == nil || IsNil(o.FiscalNumber) {
		var ret string
		return ret
	}
	return *o.FiscalNumber
}

// GetFiscalNumberOk returns a tuple with the FiscalNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccount) GetFiscalNumberOk() (*string, bool) {
	if o == nil || IsNil(o.FiscalNumber) {
		return nil, false
	}
	return o.FiscalNumber, true
}

// HasFiscalNumber returns a boolean if a field has been set.
func (o *CreateAccount) HasFiscalNumber() bool {
	if o != nil && !IsNil(o.FiscalNumber) {
		return true
	}

	return false
}

// SetFiscalNumber gets a reference to the given string and assigns it to the FiscalNumber field.
func (o *CreateAccount) SetFiscalNumber(v string) {
	o.FiscalNumber = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *CreateAccount) GetAddress() AccountAddress {
	if o == nil || IsNil(o.Address) {
		var ret AccountAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccount) GetAddressOk() (*AccountAddress, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *CreateAccount) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given AccountAddress and assigns it to the Address field.
func (o *CreateAccount) SetAddress(v AccountAddress) {
	o.Address = &v
}

func (o CreateAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ParentAccountId) {
		toSerialize["parentAccountId"] = o.ParentAccountId
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.FiscalNumber) {
		toSerialize["fiscalNumber"] = o.FiscalNumber
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateAccount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varCreateAccount := _CreateAccount{}

	err = json.Unmarshal(data, &varCreateAccount)

	if err != nil {
		return err
	}

	*o = CreateAccount(varCreateAccount)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "parentAccountId")
		delete(additionalProperties, "name")
		delete(additionalProperties, "code")
		delete(additionalProperties, "fiscalNumber")
		delete(additionalProperties, "address")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateAccount struct {
	value *CreateAccount
	isSet bool
}

func (v NullableCreateAccount) Get() *CreateAccount {
	return v.value
}

func (v *NullableCreateAccount) Set(val *CreateAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAccount(val *CreateAccount) *NullableCreateAccount {
	return &NullableCreateAccount{value: val, isSet: true}
}

func (v NullableCreateAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


