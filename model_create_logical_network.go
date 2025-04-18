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

// checks if the CreateLogicalNetwork type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLogicalNetwork{}

// CreateLogicalNetwork struct for CreateLogicalNetwork
type CreateLogicalNetwork struct {
	// Label for the logical network
	Label *string `json:"label,omitempty"`
	// Name of the logical network
	Name *string `json:"name,omitempty"`
	// Description of the logical network
	Description *string `json:"description,omitempty"`
	// Annotations for the logical network
	Annotations map[string]interface{} `json:"annotations,omitempty"`
	// Fabric ID associated with the logical network
	FabricId float32 `json:"fabricId"`
	// Infrastructure ID associated with the logical network
	InfrastructureId *float32 `json:"infrastructureId,omitempty"`
	// Type of the logical network
	LogicalNetworkType string `json:"logicalNetworkType"`
	AdditionalProperties map[string]interface{}
}

type _CreateLogicalNetwork CreateLogicalNetwork

// NewCreateLogicalNetwork instantiates a new CreateLogicalNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLogicalNetwork(fabricId float32, logicalNetworkType string) *CreateLogicalNetwork {
	this := CreateLogicalNetwork{}
	this.FabricId = fabricId
	this.LogicalNetworkType = logicalNetworkType
	return &this
}

// NewCreateLogicalNetworkWithDefaults instantiates a new CreateLogicalNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLogicalNetworkWithDefaults() *CreateLogicalNetwork {
	this := CreateLogicalNetwork{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CreateLogicalNetwork) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogicalNetwork) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CreateLogicalNetwork) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *CreateLogicalNetwork) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateLogicalNetwork) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogicalNetwork) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateLogicalNetwork) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateLogicalNetwork) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateLogicalNetwork) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogicalNetwork) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateLogicalNetwork) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateLogicalNetwork) SetDescription(v string) {
	o.Description = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *CreateLogicalNetwork) GetAnnotations() map[string]interface{} {
	if o == nil || IsNil(o.Annotations) {
		var ret map[string]interface{}
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogicalNetwork) GetAnnotationsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Annotations) {
		return map[string]interface{}{}, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *CreateLogicalNetwork) HasAnnotations() bool {
	if o != nil && !IsNil(o.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]interface{} and assigns it to the Annotations field.
func (o *CreateLogicalNetwork) SetAnnotations(v map[string]interface{}) {
	o.Annotations = v
}

// GetFabricId returns the FabricId field value
func (o *CreateLogicalNetwork) GetFabricId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.FabricId
}

// GetFabricIdOk returns a tuple with the FabricId field value
// and a boolean to check if the value has been set.
func (o *CreateLogicalNetwork) GetFabricIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FabricId, true
}

// SetFabricId sets field value
func (o *CreateLogicalNetwork) SetFabricId(v float32) {
	o.FabricId = v
}

// GetInfrastructureId returns the InfrastructureId field value if set, zero value otherwise.
func (o *CreateLogicalNetwork) GetInfrastructureId() float32 {
	if o == nil || IsNil(o.InfrastructureId) {
		var ret float32
		return ret
	}
	return *o.InfrastructureId
}

// GetInfrastructureIdOk returns a tuple with the InfrastructureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLogicalNetwork) GetInfrastructureIdOk() (*float32, bool) {
	if o == nil || IsNil(o.InfrastructureId) {
		return nil, false
	}
	return o.InfrastructureId, true
}

// HasInfrastructureId returns a boolean if a field has been set.
func (o *CreateLogicalNetwork) HasInfrastructureId() bool {
	if o != nil && !IsNil(o.InfrastructureId) {
		return true
	}

	return false
}

// SetInfrastructureId gets a reference to the given float32 and assigns it to the InfrastructureId field.
func (o *CreateLogicalNetwork) SetInfrastructureId(v float32) {
	o.InfrastructureId = &v
}

// GetLogicalNetworkType returns the LogicalNetworkType field value
func (o *CreateLogicalNetwork) GetLogicalNetworkType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogicalNetworkType
}

// GetLogicalNetworkTypeOk returns a tuple with the LogicalNetworkType field value
// and a boolean to check if the value has been set.
func (o *CreateLogicalNetwork) GetLogicalNetworkTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogicalNetworkType, true
}

// SetLogicalNetworkType sets field value
func (o *CreateLogicalNetwork) SetLogicalNetworkType(v string) {
	o.LogicalNetworkType = v
}

func (o CreateLogicalNetwork) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLogicalNetwork) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Annotations) {
		toSerialize["annotations"] = o.Annotations
	}
	toSerialize["fabricId"] = o.FabricId
	if !IsNil(o.InfrastructureId) {
		toSerialize["infrastructureId"] = o.InfrastructureId
	}
	toSerialize["logicalNetworkType"] = o.LogicalNetworkType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateLogicalNetwork) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fabricId",
		"logicalNetworkType",
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

	varCreateLogicalNetwork := _CreateLogicalNetwork{}

	err = json.Unmarshal(data, &varCreateLogicalNetwork)

	if err != nil {
		return err
	}

	*o = CreateLogicalNetwork(varCreateLogicalNetwork)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "annotations")
		delete(additionalProperties, "fabricId")
		delete(additionalProperties, "infrastructureId")
		delete(additionalProperties, "logicalNetworkType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateLogicalNetwork struct {
	value *CreateLogicalNetwork
	isSet bool
}

func (v NullableCreateLogicalNetwork) Get() *CreateLogicalNetwork {
	return v.value
}

func (v *NullableCreateLogicalNetwork) Set(val *CreateLogicalNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLogicalNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLogicalNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLogicalNetwork(val *CreateLogicalNetwork) *NullableCreateLogicalNetwork {
	return &NullableCreateLogicalNetwork{value: val, isSet: true}
}

func (v NullableCreateLogicalNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLogicalNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


