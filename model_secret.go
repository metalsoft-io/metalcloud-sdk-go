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

// checks if the Secret type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Secret{}

// Secret struct for Secret
type Secret struct {
	// The secret ID.
	Id int32 `json:"id"`
	// ID of owner user.
	UserIdOwner float32 `json:"userIdOwner"`
	// The secret name.
	Name string `json:"name"`
	// The secret value encrypted.
	ValueEncrypted string `json:"valueEncrypted"`
	// Secret usage type.
	Usage *VariableUsageType `json:"usage,omitempty"`
	// Timestamp of creation.
	CreatedTimestamp string `json:"createdTimestamp"`
	// Timestamp of last update.
	UpdatedTimestamp string `json:"updatedTimestamp"`
	// Reference links
	Links []Link `json:"links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Secret Secret

// NewSecret instantiates a new Secret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecret(id int32, userIdOwner float32, name string, valueEncrypted string, createdTimestamp string, updatedTimestamp string) *Secret {
	this := Secret{}
	this.Id = id
	this.UserIdOwner = userIdOwner
	this.Name = name
	this.ValueEncrypted = valueEncrypted
	this.CreatedTimestamp = createdTimestamp
	this.UpdatedTimestamp = updatedTimestamp
	return &this
}

// NewSecretWithDefaults instantiates a new Secret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretWithDefaults() *Secret {
	this := Secret{}
	return &this
}

// GetId returns the Id field value
func (o *Secret) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Secret) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Secret) SetId(v int32) {
	o.Id = v
}

// GetUserIdOwner returns the UserIdOwner field value
func (o *Secret) GetUserIdOwner() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UserIdOwner
}

// GetUserIdOwnerOk returns a tuple with the UserIdOwner field value
// and a boolean to check if the value has been set.
func (o *Secret) GetUserIdOwnerOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserIdOwner, true
}

// SetUserIdOwner sets field value
func (o *Secret) SetUserIdOwner(v float32) {
	o.UserIdOwner = v
}

// GetName returns the Name field value
func (o *Secret) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Secret) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Secret) SetName(v string) {
	o.Name = v
}

// GetValueEncrypted returns the ValueEncrypted field value
func (o *Secret) GetValueEncrypted() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValueEncrypted
}

// GetValueEncryptedOk returns a tuple with the ValueEncrypted field value
// and a boolean to check if the value has been set.
func (o *Secret) GetValueEncryptedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueEncrypted, true
}

// SetValueEncrypted sets field value
func (o *Secret) SetValueEncrypted(v string) {
	o.ValueEncrypted = v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *Secret) GetUsage() VariableUsageType {
	if o == nil || IsNil(o.Usage) {
		var ret VariableUsageType
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Secret) GetUsageOk() (*VariableUsageType, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *Secret) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given VariableUsageType and assigns it to the Usage field.
func (o *Secret) SetUsage(v VariableUsageType) {
	o.Usage = &v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value
func (o *Secret) GetCreatedTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *Secret) GetCreatedTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTimestamp, true
}

// SetCreatedTimestamp sets field value
func (o *Secret) SetCreatedTimestamp(v string) {
	o.CreatedTimestamp = v
}

// GetUpdatedTimestamp returns the UpdatedTimestamp field value
func (o *Secret) GetUpdatedTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedTimestamp
}

// GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *Secret) GetUpdatedTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedTimestamp, true
}

// SetUpdatedTimestamp sets field value
func (o *Secret) SetUpdatedTimestamp(v string) {
	o.UpdatedTimestamp = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Secret) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Secret) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Secret) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *Secret) SetLinks(v []Link) {
	o.Links = v
}

func (o Secret) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Secret) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["userIdOwner"] = o.UserIdOwner
	toSerialize["name"] = o.Name
	toSerialize["valueEncrypted"] = o.ValueEncrypted
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	toSerialize["createdTimestamp"] = o.CreatedTimestamp
	toSerialize["updatedTimestamp"] = o.UpdatedTimestamp
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Secret) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"userIdOwner",
		"name",
		"valueEncrypted",
		"createdTimestamp",
		"updatedTimestamp",
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

	varSecret := _Secret{}

	err = json.Unmarshal(data, &varSecret)

	if err != nil {
		return err
	}

	*o = Secret(varSecret)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "userIdOwner")
		delete(additionalProperties, "name")
		delete(additionalProperties, "valueEncrypted")
		delete(additionalProperties, "usage")
		delete(additionalProperties, "createdTimestamp")
		delete(additionalProperties, "updatedTimestamp")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecret struct {
	value *Secret
	isSet bool
}

func (v NullableSecret) Get() *Secret {
	return v.value
}

func (v *NullableSecret) Set(val *Secret) {
	v.value = val
	v.isSet = true
}

func (v NullableSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecret(val *Secret) *NullableSecret {
	return &NullableSecret{value: val, isSet: true}
}

func (v NullableSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


