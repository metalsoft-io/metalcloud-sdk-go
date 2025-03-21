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

// checks if the UpdateServerFirmwareUpgradePolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateServerFirmwareUpgradePolicy{}

// UpdateServerFirmwareUpgradePolicy struct for UpdateServerFirmwareUpgradePolicy
type UpdateServerFirmwareUpgradePolicy struct {
	// The rules of the firmware upgrade policy.
	Rules []ServerFirmwareUpgradePolicyRule `json:"rules,omitempty"`
	// The unique identifiers of the server instance groups associated with the firmware upgrade policy.
	ServerInstanceGroupIds []float32 `json:"serverInstanceGroupIds,omitempty"`
	// The label of the firmware upgrade policy.
	Label *string `json:"label,omitempty"`
	// The status of the firmware upgrade policy.
	Status *string `json:"status,omitempty"`
	// The action of the firmware upgrade policy.
	Action *string `json:"action,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateServerFirmwareUpgradePolicy UpdateServerFirmwareUpgradePolicy

// NewUpdateServerFirmwareUpgradePolicy instantiates a new UpdateServerFirmwareUpgradePolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateServerFirmwareUpgradePolicy() *UpdateServerFirmwareUpgradePolicy {
	this := UpdateServerFirmwareUpgradePolicy{}
	return &this
}

// NewUpdateServerFirmwareUpgradePolicyWithDefaults instantiates a new UpdateServerFirmwareUpgradePolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateServerFirmwareUpgradePolicyWithDefaults() *UpdateServerFirmwareUpgradePolicy {
	this := UpdateServerFirmwareUpgradePolicy{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *UpdateServerFirmwareUpgradePolicy) GetRules() []ServerFirmwareUpgradePolicyRule {
	if o == nil || IsNil(o.Rules) {
		var ret []ServerFirmwareUpgradePolicyRule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServerFirmwareUpgradePolicy) GetRulesOk() ([]ServerFirmwareUpgradePolicyRule, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *UpdateServerFirmwareUpgradePolicy) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []ServerFirmwareUpgradePolicyRule and assigns it to the Rules field.
func (o *UpdateServerFirmwareUpgradePolicy) SetRules(v []ServerFirmwareUpgradePolicyRule) {
	o.Rules = v
}

// GetServerInstanceGroupIds returns the ServerInstanceGroupIds field value if set, zero value otherwise.
func (o *UpdateServerFirmwareUpgradePolicy) GetServerInstanceGroupIds() []float32 {
	if o == nil || IsNil(o.ServerInstanceGroupIds) {
		var ret []float32
		return ret
	}
	return o.ServerInstanceGroupIds
}

// GetServerInstanceGroupIdsOk returns a tuple with the ServerInstanceGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServerFirmwareUpgradePolicy) GetServerInstanceGroupIdsOk() ([]float32, bool) {
	if o == nil || IsNil(o.ServerInstanceGroupIds) {
		return nil, false
	}
	return o.ServerInstanceGroupIds, true
}

// HasServerInstanceGroupIds returns a boolean if a field has been set.
func (o *UpdateServerFirmwareUpgradePolicy) HasServerInstanceGroupIds() bool {
	if o != nil && !IsNil(o.ServerInstanceGroupIds) {
		return true
	}

	return false
}

// SetServerInstanceGroupIds gets a reference to the given []float32 and assigns it to the ServerInstanceGroupIds field.
func (o *UpdateServerFirmwareUpgradePolicy) SetServerInstanceGroupIds(v []float32) {
	o.ServerInstanceGroupIds = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *UpdateServerFirmwareUpgradePolicy) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServerFirmwareUpgradePolicy) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *UpdateServerFirmwareUpgradePolicy) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *UpdateServerFirmwareUpgradePolicy) SetLabel(v string) {
	o.Label = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UpdateServerFirmwareUpgradePolicy) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServerFirmwareUpgradePolicy) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UpdateServerFirmwareUpgradePolicy) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UpdateServerFirmwareUpgradePolicy) SetStatus(v string) {
	o.Status = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *UpdateServerFirmwareUpgradePolicy) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServerFirmwareUpgradePolicy) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *UpdateServerFirmwareUpgradePolicy) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *UpdateServerFirmwareUpgradePolicy) SetAction(v string) {
	o.Action = &v
}

func (o UpdateServerFirmwareUpgradePolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateServerFirmwareUpgradePolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !IsNil(o.ServerInstanceGroupIds) {
		toSerialize["serverInstanceGroupIds"] = o.ServerInstanceGroupIds
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateServerFirmwareUpgradePolicy) UnmarshalJSON(data []byte) (err error) {
	varUpdateServerFirmwareUpgradePolicy := _UpdateServerFirmwareUpgradePolicy{}

	err = json.Unmarshal(data, &varUpdateServerFirmwareUpgradePolicy)

	if err != nil {
		return err
	}

	*o = UpdateServerFirmwareUpgradePolicy(varUpdateServerFirmwareUpgradePolicy)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rules")
		delete(additionalProperties, "serverInstanceGroupIds")
		delete(additionalProperties, "label")
		delete(additionalProperties, "status")
		delete(additionalProperties, "action")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateServerFirmwareUpgradePolicy struct {
	value *UpdateServerFirmwareUpgradePolicy
	isSet bool
}

func (v NullableUpdateServerFirmwareUpgradePolicy) Get() *UpdateServerFirmwareUpgradePolicy {
	return v.value
}

func (v *NullableUpdateServerFirmwareUpgradePolicy) Set(val *UpdateServerFirmwareUpgradePolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateServerFirmwareUpgradePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateServerFirmwareUpgradePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateServerFirmwareUpgradePolicy(val *UpdateServerFirmwareUpgradePolicy) *NullableUpdateServerFirmwareUpgradePolicy {
	return &NullableUpdateServerFirmwareUpgradePolicy{value: val, isSet: true}
}

func (v NullableUpdateServerFirmwareUpgradePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateServerFirmwareUpgradePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


