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

// checks if the EmailTemplateCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailTemplateCreate{}

// EmailTemplateCreate struct for EmailTemplateCreate
type EmailTemplateCreate struct {
	// Email template name
	Name string `json:"name"`
	// Email template subject
	Subject string `json:"subject"`
	// Email template description
	Description *string `json:"description,omitempty"`
	// Email template text
	Text string `json:"text"`
	// Email template html
	Html string `json:"html"`
	AdditionalProperties map[string]interface{}
}

type _EmailTemplateCreate EmailTemplateCreate

// NewEmailTemplateCreate instantiates a new EmailTemplateCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailTemplateCreate(name string, subject string, text string, html string) *EmailTemplateCreate {
	this := EmailTemplateCreate{}
	this.Name = name
	this.Subject = subject
	this.Text = text
	this.Html = html
	return &this
}

// NewEmailTemplateCreateWithDefaults instantiates a new EmailTemplateCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailTemplateCreateWithDefaults() *EmailTemplateCreate {
	this := EmailTemplateCreate{}
	return &this
}

// GetName returns the Name field value
func (o *EmailTemplateCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EmailTemplateCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EmailTemplateCreate) SetName(v string) {
	o.Name = v
}

// GetSubject returns the Subject field value
func (o *EmailTemplateCreate) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *EmailTemplateCreate) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *EmailTemplateCreate) SetSubject(v string) {
	o.Subject = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EmailTemplateCreate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplateCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EmailTemplateCreate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EmailTemplateCreate) SetDescription(v string) {
	o.Description = &v
}

// GetText returns the Text field value
func (o *EmailTemplateCreate) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *EmailTemplateCreate) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *EmailTemplateCreate) SetText(v string) {
	o.Text = v
}

// GetHtml returns the Html field value
func (o *EmailTemplateCreate) GetHtml() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Html
}

// GetHtmlOk returns a tuple with the Html field value
// and a boolean to check if the value has been set.
func (o *EmailTemplateCreate) GetHtmlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Html, true
}

// SetHtml sets field value
func (o *EmailTemplateCreate) SetHtml(v string) {
	o.Html = v
}

func (o EmailTemplateCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailTemplateCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["subject"] = o.Subject
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["text"] = o.Text
	toSerialize["html"] = o.Html

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EmailTemplateCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"subject",
		"text",
		"html",
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

	varEmailTemplateCreate := _EmailTemplateCreate{}

	err = json.Unmarshal(data, &varEmailTemplateCreate)

	if err != nil {
		return err
	}

	*o = EmailTemplateCreate(varEmailTemplateCreate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "subject")
		delete(additionalProperties, "description")
		delete(additionalProperties, "text")
		delete(additionalProperties, "html")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEmailTemplateCreate struct {
	value *EmailTemplateCreate
	isSet bool
}

func (v NullableEmailTemplateCreate) Get() *EmailTemplateCreate {
	return v.value
}

func (v *NullableEmailTemplateCreate) Set(val *EmailTemplateCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailTemplateCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailTemplateCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailTemplateCreate(val *EmailTemplateCreate) *NullableEmailTemplateCreate {
	return &NullableEmailTemplateCreate{value: val, isSet: true}
}

func (v NullableEmailTemplateCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailTemplateCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


