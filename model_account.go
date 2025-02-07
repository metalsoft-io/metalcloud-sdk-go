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

// checks if the Account type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Account{}

// Account struct for Account
type Account struct {
	// Account ID
	Id float32 `json:"id"`
	// Revision number
	Revision float32 `json:"revision"`
	// The ID of the parent account
	ParentAccountId *float32 `json:"parentAccountId,omitempty"`
	// The name of the account
	Name string `json:"name"`
	// The code of the account
	Code *string `json:"code,omitempty"`
	// The fiscal number of the account
	FiscalNumber *string `json:"fiscalNumber,omitempty"`
	Address *AccountAddress `json:"address,omitempty"`
	// The user ID of the primary contact
	PrimaryContactId *float32 `json:"primaryContactId,omitempty"`
	// The user ID of the secondary contact
	SecondaryContactId *float32 `json:"secondaryContactId,omitempty"`
	// Whether the account is archived
	Archived *float32 `json:"archived,omitempty"`
	Limits AccountLimits `json:"limits"`
	// Reference links
	Links []Link `json:"links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Account Account

// NewAccount instantiates a new Account object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccount(id float32, revision float32, name string, limits AccountLimits) *Account {
	this := Account{}
	this.Id = id
	this.Revision = revision
	this.Name = name
	this.Limits = limits
	return &this
}

// NewAccountWithDefaults instantiates a new Account object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountWithDefaults() *Account {
	this := Account{}
	return &this
}

// GetId returns the Id field value
func (o *Account) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Account) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Account) SetId(v float32) {
	o.Id = v
}

// GetRevision returns the Revision field value
func (o *Account) GetRevision() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *Account) GetRevisionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *Account) SetRevision(v float32) {
	o.Revision = v
}

// GetParentAccountId returns the ParentAccountId field value if set, zero value otherwise.
func (o *Account) GetParentAccountId() float32 {
	if o == nil || IsNil(o.ParentAccountId) {
		var ret float32
		return ret
	}
	return *o.ParentAccountId
}

// GetParentAccountIdOk returns a tuple with the ParentAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetParentAccountIdOk() (*float32, bool) {
	if o == nil || IsNil(o.ParentAccountId) {
		return nil, false
	}
	return o.ParentAccountId, true
}

// HasParentAccountId returns a boolean if a field has been set.
func (o *Account) HasParentAccountId() bool {
	if o != nil && !IsNil(o.ParentAccountId) {
		return true
	}

	return false
}

// SetParentAccountId gets a reference to the given float32 and assigns it to the ParentAccountId field.
func (o *Account) SetParentAccountId(v float32) {
	o.ParentAccountId = &v
}

// GetName returns the Name field value
func (o *Account) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Account) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Account) SetName(v string) {
	o.Name = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Account) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Account) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *Account) SetCode(v string) {
	o.Code = &v
}

// GetFiscalNumber returns the FiscalNumber field value if set, zero value otherwise.
func (o *Account) GetFiscalNumber() string {
	if o == nil || IsNil(o.FiscalNumber) {
		var ret string
		return ret
	}
	return *o.FiscalNumber
}

// GetFiscalNumberOk returns a tuple with the FiscalNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetFiscalNumberOk() (*string, bool) {
	if o == nil || IsNil(o.FiscalNumber) {
		return nil, false
	}
	return o.FiscalNumber, true
}

// HasFiscalNumber returns a boolean if a field has been set.
func (o *Account) HasFiscalNumber() bool {
	if o != nil && !IsNil(o.FiscalNumber) {
		return true
	}

	return false
}

// SetFiscalNumber gets a reference to the given string and assigns it to the FiscalNumber field.
func (o *Account) SetFiscalNumber(v string) {
	o.FiscalNumber = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Account) GetAddress() AccountAddress {
	if o == nil || IsNil(o.Address) {
		var ret AccountAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetAddressOk() (*AccountAddress, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Account) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given AccountAddress and assigns it to the Address field.
func (o *Account) SetAddress(v AccountAddress) {
	o.Address = &v
}

// GetPrimaryContactId returns the PrimaryContactId field value if set, zero value otherwise.
func (o *Account) GetPrimaryContactId() float32 {
	if o == nil || IsNil(o.PrimaryContactId) {
		var ret float32
		return ret
	}
	return *o.PrimaryContactId
}

// GetPrimaryContactIdOk returns a tuple with the PrimaryContactId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetPrimaryContactIdOk() (*float32, bool) {
	if o == nil || IsNil(o.PrimaryContactId) {
		return nil, false
	}
	return o.PrimaryContactId, true
}

// HasPrimaryContactId returns a boolean if a field has been set.
func (o *Account) HasPrimaryContactId() bool {
	if o != nil && !IsNil(o.PrimaryContactId) {
		return true
	}

	return false
}

// SetPrimaryContactId gets a reference to the given float32 and assigns it to the PrimaryContactId field.
func (o *Account) SetPrimaryContactId(v float32) {
	o.PrimaryContactId = &v
}

// GetSecondaryContactId returns the SecondaryContactId field value if set, zero value otherwise.
func (o *Account) GetSecondaryContactId() float32 {
	if o == nil || IsNil(o.SecondaryContactId) {
		var ret float32
		return ret
	}
	return *o.SecondaryContactId
}

// GetSecondaryContactIdOk returns a tuple with the SecondaryContactId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetSecondaryContactIdOk() (*float32, bool) {
	if o == nil || IsNil(o.SecondaryContactId) {
		return nil, false
	}
	return o.SecondaryContactId, true
}

// HasSecondaryContactId returns a boolean if a field has been set.
func (o *Account) HasSecondaryContactId() bool {
	if o != nil && !IsNil(o.SecondaryContactId) {
		return true
	}

	return false
}

// SetSecondaryContactId gets a reference to the given float32 and assigns it to the SecondaryContactId field.
func (o *Account) SetSecondaryContactId(v float32) {
	o.SecondaryContactId = &v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *Account) GetArchived() float32 {
	if o == nil || IsNil(o.Archived) {
		var ret float32
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetArchivedOk() (*float32, bool) {
	if o == nil || IsNil(o.Archived) {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *Account) HasArchived() bool {
	if o != nil && !IsNil(o.Archived) {
		return true
	}

	return false
}

// SetArchived gets a reference to the given float32 and assigns it to the Archived field.
func (o *Account) SetArchived(v float32) {
	o.Archived = &v
}

// GetLimits returns the Limits field value
func (o *Account) GetLimits() AccountLimits {
	if o == nil {
		var ret AccountLimits
		return ret
	}

	return o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value
// and a boolean to check if the value has been set.
func (o *Account) GetLimitsOk() (*AccountLimits, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limits, true
}

// SetLimits sets field value
func (o *Account) SetLimits(v AccountLimits) {
	o.Limits = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Account) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Account) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *Account) SetLinks(v []Link) {
	o.Links = v
}

func (o Account) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Account) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["revision"] = o.Revision
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
	if !IsNil(o.PrimaryContactId) {
		toSerialize["primaryContactId"] = o.PrimaryContactId
	}
	if !IsNil(o.SecondaryContactId) {
		toSerialize["secondaryContactId"] = o.SecondaryContactId
	}
	if !IsNil(o.Archived) {
		toSerialize["archived"] = o.Archived
	}
	toSerialize["limits"] = o.Limits
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Account) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"revision",
		"name",
		"limits",
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

	varAccount := _Account{}

	err = json.Unmarshal(data, &varAccount)

	if err != nil {
		return err
	}

	*o = Account(varAccount)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "revision")
		delete(additionalProperties, "parentAccountId")
		delete(additionalProperties, "name")
		delete(additionalProperties, "code")
		delete(additionalProperties, "fiscalNumber")
		delete(additionalProperties, "address")
		delete(additionalProperties, "primaryContactId")
		delete(additionalProperties, "secondaryContactId")
		delete(additionalProperties, "archived")
		delete(additionalProperties, "limits")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccount struct {
	value *Account
	isSet bool
}

func (v NullableAccount) Get() *Account {
	return v.value
}

func (v *NullableAccount) Set(val *Account) {
	v.value = val
	v.isSet = true
}

func (v NullableAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccount(val *Account) *NullableAccount {
	return &NullableAccount{value: val, isSet: true}
}

func (v NullableAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


