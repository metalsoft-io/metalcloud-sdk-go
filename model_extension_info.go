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

// checks if the ExtensionInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtensionInfo{}

// ExtensionInfo struct for ExtensionInfo
type ExtensionInfo struct {
	// The extension ID
	Id *float32 `json:"id,omitempty"`
	// The extension unique slug
	Slug *string `json:"slug,omitempty"`
	// The extension name
	Name *string `json:"name,omitempty"`
	// The extension unique label
	Label *string `json:"label,omitempty" validate:"regexp=^(?!.*-$)[a-z]{1}[a-z0-9-]{0,62}$"`
	// The extension description
	Description *string `json:"description,omitempty"`
	// Extension status
	Status *string `json:"status,omitempty"`
	// Reference links
	Links []Link `json:"links,omitempty"`
	// Icon of the extension.
	Icon *string `json:"icon,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExtensionInfo ExtensionInfo

// NewExtensionInfo instantiates a new ExtensionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtensionInfo() *ExtensionInfo {
	this := ExtensionInfo{}
	return &this
}

// NewExtensionInfoWithDefaults instantiates a new ExtensionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionInfoWithDefaults() *ExtensionInfo {
	this := ExtensionInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExtensionInfo) GetId() float32 {
	if o == nil || IsNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionInfo) GetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExtensionInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *ExtensionInfo) SetId(v float32) {
	o.Id = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *ExtensionInfo) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionInfo) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *ExtensionInfo) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *ExtensionInfo) SetSlug(v string) {
	o.Slug = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExtensionInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExtensionInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExtensionInfo) SetName(v string) {
	o.Name = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ExtensionInfo) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionInfo) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ExtensionInfo) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ExtensionInfo) SetLabel(v string) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ExtensionInfo) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ExtensionInfo) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ExtensionInfo) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ExtensionInfo) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionInfo) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ExtensionInfo) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ExtensionInfo) SetStatus(v string) {
	o.Status = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ExtensionInfo) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionInfo) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ExtensionInfo) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *ExtensionInfo) SetLinks(v []Link) {
	o.Links = v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *ExtensionInfo) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionInfo) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *ExtensionInfo) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *ExtensionInfo) SetIcon(v string) {
	o.Icon = &v
}

func (o ExtensionInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtensionInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExtensionInfo) UnmarshalJSON(data []byte) (err error) {
	varExtensionInfo := _ExtensionInfo{}

	err = json.Unmarshal(data, &varExtensionInfo)

	if err != nil {
		return err
	}

	*o = ExtensionInfo(varExtensionInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		delete(additionalProperties, "status")
		delete(additionalProperties, "links")
		delete(additionalProperties, "icon")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExtensionInfo struct {
	value *ExtensionInfo
	isSet bool
}

func (v NullableExtensionInfo) Get() *ExtensionInfo {
	return v.value
}

func (v *NullableExtensionInfo) Set(val *ExtensionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionInfo(val *ExtensionInfo) *NullableExtensionInfo {
	return &NullableExtensionInfo{value: val, isSet: true}
}

func (v NullableExtensionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


