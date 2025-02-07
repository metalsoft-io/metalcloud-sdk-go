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

// checks if the ServerInstanceProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerInstanceProfile{}

// ServerInstanceProfile struct for ServerInstanceProfile
type ServerInstanceProfile struct {
	// The Server profile ID.
	Id int32 `json:"id"`
	// Revision number
	Revision float32 `json:"revision"`
	// The Server profile owner ID. Can be null if the profile is public. Will be set to the currently logged user.
	OwnerId *int32 `json:"ownerId,omitempty"`
	// The Server profile label. Will be automatically generated if not provided.
	Label *string `json:"label,omitempty"`
	// Network profiles mapping for each network in this infrastructure. Changes to this configuration will be duplicated on each vm-instance of this group.
	NetworkProfiles []ServerInstanceConfigurationNetworkProfilesInner `json:"networkProfiles,omitempty"`
	NetworkInterfaces *ServerInstanceProfileNetworkInterfaces `json:"networkInterfaces,omitempty"`
	// The template id of the operating system to deploy on the Server. Can be null in which case no OS will be deployed but all operations will continue as normal.
	TemplateId *int32 `json:"templateId,omitempty"`
	// Reference links
	Links []Link `json:"links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServerInstanceProfile ServerInstanceProfile

// NewServerInstanceProfile instantiates a new ServerInstanceProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerInstanceProfile(id int32, revision float32) *ServerInstanceProfile {
	this := ServerInstanceProfile{}
	this.Id = id
	this.Revision = revision
	return &this
}

// NewServerInstanceProfileWithDefaults instantiates a new ServerInstanceProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerInstanceProfileWithDefaults() *ServerInstanceProfile {
	this := ServerInstanceProfile{}
	return &this
}

// GetId returns the Id field value
func (o *ServerInstanceProfile) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServerInstanceProfile) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServerInstanceProfile) SetId(v int32) {
	o.Id = v
}

// GetRevision returns the Revision field value
func (o *ServerInstanceProfile) GetRevision() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *ServerInstanceProfile) GetRevisionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *ServerInstanceProfile) SetRevision(v float32) {
	o.Revision = v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *ServerInstanceProfile) GetOwnerId() int32 {
	if o == nil || IsNil(o.OwnerId) {
		var ret int32
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceProfile) GetOwnerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *ServerInstanceProfile) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given int32 and assigns it to the OwnerId field.
func (o *ServerInstanceProfile) SetOwnerId(v int32) {
	o.OwnerId = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ServerInstanceProfile) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceProfile) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ServerInstanceProfile) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ServerInstanceProfile) SetLabel(v string) {
	o.Label = &v
}

// GetNetworkProfiles returns the NetworkProfiles field value if set, zero value otherwise.
func (o *ServerInstanceProfile) GetNetworkProfiles() []ServerInstanceConfigurationNetworkProfilesInner {
	if o == nil || IsNil(o.NetworkProfiles) {
		var ret []ServerInstanceConfigurationNetworkProfilesInner
		return ret
	}
	return o.NetworkProfiles
}

// GetNetworkProfilesOk returns a tuple with the NetworkProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceProfile) GetNetworkProfilesOk() ([]ServerInstanceConfigurationNetworkProfilesInner, bool) {
	if o == nil || IsNil(o.NetworkProfiles) {
		return nil, false
	}
	return o.NetworkProfiles, true
}

// HasNetworkProfiles returns a boolean if a field has been set.
func (o *ServerInstanceProfile) HasNetworkProfiles() bool {
	if o != nil && !IsNil(o.NetworkProfiles) {
		return true
	}

	return false
}

// SetNetworkProfiles gets a reference to the given []ServerInstanceConfigurationNetworkProfilesInner and assigns it to the NetworkProfiles field.
func (o *ServerInstanceProfile) SetNetworkProfiles(v []ServerInstanceConfigurationNetworkProfilesInner) {
	o.NetworkProfiles = v
}

// GetNetworkInterfaces returns the NetworkInterfaces field value if set, zero value otherwise.
func (o *ServerInstanceProfile) GetNetworkInterfaces() ServerInstanceProfileNetworkInterfaces {
	if o == nil || IsNil(o.NetworkInterfaces) {
		var ret ServerInstanceProfileNetworkInterfaces
		return ret
	}
	return *o.NetworkInterfaces
}

// GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceProfile) GetNetworkInterfacesOk() (*ServerInstanceProfileNetworkInterfaces, bool) {
	if o == nil || IsNil(o.NetworkInterfaces) {
		return nil, false
	}
	return o.NetworkInterfaces, true
}

// HasNetworkInterfaces returns a boolean if a field has been set.
func (o *ServerInstanceProfile) HasNetworkInterfaces() bool {
	if o != nil && !IsNil(o.NetworkInterfaces) {
		return true
	}

	return false
}

// SetNetworkInterfaces gets a reference to the given ServerInstanceProfileNetworkInterfaces and assigns it to the NetworkInterfaces field.
func (o *ServerInstanceProfile) SetNetworkInterfaces(v ServerInstanceProfileNetworkInterfaces) {
	o.NetworkInterfaces = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *ServerInstanceProfile) GetTemplateId() int32 {
	if o == nil || IsNil(o.TemplateId) {
		var ret int32
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceProfile) GetTemplateIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *ServerInstanceProfile) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given int32 and assigns it to the TemplateId field.
func (o *ServerInstanceProfile) SetTemplateId(v int32) {
	o.TemplateId = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ServerInstanceProfile) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerInstanceProfile) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ServerInstanceProfile) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *ServerInstanceProfile) SetLinks(v []Link) {
	o.Links = v
}

func (o ServerInstanceProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerInstanceProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["revision"] = o.Revision
	if !IsNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.NetworkProfiles) {
		toSerialize["networkProfiles"] = o.NetworkProfiles
	}
	if !IsNil(o.NetworkInterfaces) {
		toSerialize["networkInterfaces"] = o.NetworkInterfaces
	}
	if !IsNil(o.TemplateId) {
		toSerialize["templateId"] = o.TemplateId
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServerInstanceProfile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"revision",
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

	varServerInstanceProfile := _ServerInstanceProfile{}

	err = json.Unmarshal(data, &varServerInstanceProfile)

	if err != nil {
		return err
	}

	*o = ServerInstanceProfile(varServerInstanceProfile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "revision")
		delete(additionalProperties, "ownerId")
		delete(additionalProperties, "label")
		delete(additionalProperties, "networkProfiles")
		delete(additionalProperties, "networkInterfaces")
		delete(additionalProperties, "templateId")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServerInstanceProfile struct {
	value *ServerInstanceProfile
	isSet bool
}

func (v NullableServerInstanceProfile) Get() *ServerInstanceProfile {
	return v.value
}

func (v *NullableServerInstanceProfile) Set(val *ServerInstanceProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableServerInstanceProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableServerInstanceProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerInstanceProfile(val *ServerInstanceProfile) *NullableServerInstanceProfile {
	return &NullableServerInstanceProfile{value: val, isSet: true}
}

func (v NullableServerInstanceProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerInstanceProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


