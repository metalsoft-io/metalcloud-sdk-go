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

// checks if the Extension type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Extension{}

// Extension struct for Extension
type Extension struct {
	// The extension ID
	Id float32 `json:"id"`
	// Revision number
	Revision float32 `json:"revision"`
	// The extension unique slug
	Slug string `json:"slug"`
	// The extension name
	Name string `json:"name"`
	// The extension unique label
	Label *string `json:"label,omitempty" validate:"regexp=^(?!.*-$)[a-z]{1}[a-z0-9-]{0,62}$"`
	// The extension description
	Description string `json:"description"`
	// Extension status
	Status string `json:"status"`
	Definition ExtensionDefinition `json:"definition"`
	// Reference links
	Links []Link `json:"links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Extension Extension

// NewExtension instantiates a new Extension object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtension(id float32, revision float32, slug string, name string, description string, status string, definition ExtensionDefinition) *Extension {
	this := Extension{}
	this.Id = id
	this.Revision = revision
	this.Slug = slug
	this.Name = name
	this.Description = description
	this.Status = status
	this.Definition = definition
	return &this
}

// NewExtensionWithDefaults instantiates a new Extension object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionWithDefaults() *Extension {
	this := Extension{}
	return &this
}

// GetId returns the Id field value
func (o *Extension) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Extension) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Extension) SetId(v float32) {
	o.Id = v
}

// GetRevision returns the Revision field value
func (o *Extension) GetRevision() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *Extension) GetRevisionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *Extension) SetRevision(v float32) {
	o.Revision = v
}

// GetSlug returns the Slug field value
func (o *Extension) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *Extension) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *Extension) SetSlug(v string) {
	o.Slug = v
}

// GetName returns the Name field value
func (o *Extension) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Extension) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Extension) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *Extension) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Extension) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *Extension) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *Extension) SetLabel(v string) {
	o.Label = &v
}

// GetDescription returns the Description field value
func (o *Extension) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Extension) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Extension) SetDescription(v string) {
	o.Description = v
}

// GetStatus returns the Status field value
func (o *Extension) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Extension) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Extension) SetStatus(v string) {
	o.Status = v
}

// GetDefinition returns the Definition field value
func (o *Extension) GetDefinition() ExtensionDefinition {
	if o == nil {
		var ret ExtensionDefinition
		return ret
	}

	return o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value
// and a boolean to check if the value has been set.
func (o *Extension) GetDefinitionOk() (*ExtensionDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Definition, true
}

// SetDefinition sets field value
func (o *Extension) SetDefinition(v ExtensionDefinition) {
	o.Definition = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Extension) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Extension) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Extension) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *Extension) SetLinks(v []Link) {
	o.Links = v
}

func (o Extension) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Extension) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["revision"] = o.Revision
	toSerialize["slug"] = o.Slug
	toSerialize["name"] = o.Name
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	toSerialize["description"] = o.Description
	toSerialize["status"] = o.Status
	toSerialize["definition"] = o.Definition
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Extension) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"revision",
		"slug",
		"name",
		"description",
		"status",
		"definition",
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

	varExtension := _Extension{}

	err = json.Unmarshal(data, &varExtension)

	if err != nil {
		return err
	}

	*o = Extension(varExtension)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "revision")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		delete(additionalProperties, "status")
		delete(additionalProperties, "definition")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExtension struct {
	value *Extension
	isSet bool
}

func (v NullableExtension) Get() *Extension {
	return v.value
}

func (v *NullableExtension) Set(val *Extension) {
	v.value = val
	v.isSet = true
}

func (v NullableExtension) IsSet() bool {
	return v.isSet
}

func (v *NullableExtension) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtension(val *Extension) *NullableExtension {
	return &NullableExtension{value: val, isSet: true}
}

func (v NullableExtension) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtension) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


